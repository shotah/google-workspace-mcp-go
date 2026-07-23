package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

// CallbackResult holds the result from the OAuth callback.
type CallbackResult struct {
	Code  string
	State string
	Error string
}

// StartAuthFlow initiates the Google OAuth flow for the given user.
// It starts a local HTTP server to handle the callback, generates the
// authorization URL, and returns a message for the user to follow.
//
// The callback server runs in the background and stores credentials
// when the user completes the flow.
func StartAuthFlow(
	ctx context.Context,
	serviceName string,
	userEmail string,
	store *LocalDirectoryCredentialStore,
	onCredentialStored ...func(email string),
) (string, error) {
	clientID := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		return "", errors.New("OAuth client credentials not found. Please set GOOGLE_OAUTH_CLIENT_ID and " +
			"GOOGLE_OAUTH_CLIENT_SECRET environment variables",
		)
	}

	// Generate CSRF state token
	stateBytes := make([]byte, 16)
	if _, err := rand.Read(stateBytes); err != nil {
		return "", fmt.Errorf("generating state token: %w", err)
	}
	state := hex.EncodeToString(stateBytes)

	// Find an available port for the callback server
	port, err := findAvailablePort()
	if err != nil {
		return "", fmt.Errorf("finding available port: %w", err)
	}

	redirectURI := fmt.Sprintf("http://localhost:%d/oauth2callback", port)

	oauthConfig := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
		Scopes:      DefaultScopes,
		RedirectURL: redirectURI,
	}

	authURL := oauthConfig.AuthCodeURL(state,
		oauth2.AccessTypeOffline,
		oauth2.SetAuthURLParam("prompt", "consent"),
	)

	// Start callback server in the background
	resultCh := make(chan CallbackResult, 1)
	srv := startCallbackServer(port, state, resultCh)

	// OAuth callback outlives the tool call; keep values but not cancellation.
	oauthCtx := context.WithoutCancel(ctx)

	// Handle the callback asynchronously
	go func() {
		defer func() {
			shutdownCtx, cancel := context.WithTimeout(oauthCtx, 3*time.Second)
			defer cancel()
			if err := srv.Shutdown(shutdownCtx); err != nil {
				log.Printf("OAuth callback server shutdown error: %v", err)
			}
		}()

		select {
		case result := <-resultCh:
			if result.Error != "" {
				log.Printf("OAuth callback error: %s", result.Error)
				return
			}

			tok, err := oauthConfig.Exchange(oauthCtx, result.Code)
			if err != nil {
				log.Printf("OAuth token exchange error: %v", err)
				return
			}

			// Fetch user email from the token
			email, err := fetchUserEmail(tok)
			if err != nil {
				log.Printf("Failed to fetch user email: %v", err)
				if userEmail != "" {
					email = userEmail
				} else {
					return
				}
			}

			// Store credentials
			cred := &StoredCredential{
				Token:  tok,
				Config: oauthConfig,
			}
			if err := store.StoreCredential(email, cred); err != nil {
				log.Printf("Failed to store credentials for %s: %v", email, err)
				return
			}
			log.Printf("Successfully stored credentials for %s", email)

			// Notify caller so it can invalidate caches.
			for _, fn := range onCredentialStored {
				fn(email)
			}

		case <-time.After(5 * time.Minute):
			log.Printf("OAuth callback timed out after 5 minutes")
		}
	}()

	// Build user-facing message
	initialEmailProvided := userEmail != "" &&
		strings.TrimSpace(userEmail) != "" &&
		!strings.EqualFold(userEmail, "default")

	userDisplayName := serviceName
	if initialEmailProvided {
		userDisplayName = fmt.Sprintf("%s for '%s'", serviceName, userEmail)
	}

	var msg strings.Builder
	fmt.Fprintf(&msg, "**ACTION REQUIRED: Google Authentication Needed for %s**\n\n", userDisplayName)
	fmt.Fprintf(&msg, "To proceed, the user must authorize this application for %s access using all required permissions.\n", serviceName)
	msg.WriteString("**LLM, please present this exact authorization URL to the user as a clickable hyperlink:**\n")
	fmt.Fprintf(&msg, "Authorization URL: %s\n", authURL)
	fmt.Fprintf(&msg, "Markdown for hyperlink: [Click here to authorize %s access](%s)\n\n", serviceName, authURL)
	msg.WriteString("**LLM, after presenting the link, instruct the user as follows:**\n")
	msg.WriteString("1. Click the link and complete the authorization in their browser.\n")

	if !initialEmailProvided {
		msg.WriteString("2. After successful authorization, the browser page will display the authenticated email address.\n")
		msg.WriteString("   **LLM: Instruct the user to provide you with this email address.**\n")
		msg.WriteString("3. Once you have the email, **retry their original command, ensuring you include this `user_google_email`.**\n")
	} else {
		msg.WriteString("2. After successful authorization, **retry their original command**.\n")
	}

	fmt.Fprintf(&msg, "\nThe application will use the new credentials. If '%s' was provided, it must match the authenticated account.", userEmail)

	return msg.String(), nil
}

// findAvailablePort finds a free TCP port on localhost.
func findAvailablePort() (int, error) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	tcpAddr, ok := l.Addr().(*net.TCPAddr)
	if !ok {
		_ = l.Close()
		return 0, errors.New("unexpected listener address type")
	}
	port := tcpAddr.Port
	_ = l.Close()
	return port, nil
}

// startCallbackServer starts a minimal HTTP server that handles the OAuth callback.
func startCallbackServer(port int, expectedState string, resultCh chan<- CallbackResult) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/oauth2callback", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		errParam := query.Get("error")
		if errParam != "" {
			resultCh <- CallbackResult{Error: "Google returned error: " + errParam}
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, "<html><body><h1>Authentication Failed</h1><p>Error: %s</p><p>You can close this window.</p></body></html>", html.EscapeString(errParam))
			return
		}

		code := query.Get("code")
		state := query.Get("state")

		if code == "" {
			resultCh <- CallbackResult{Error: "no authorization code received"}
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, "<html><body><h1>Authentication Failed</h1><p>No authorization code received.</p><p>You can close this window.</p></body></html>")
			return
		}

		if state != expectedState {
			resultCh <- CallbackResult{Error: "state mismatch — possible CSRF attack"}
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, "<html><body><h1>Authentication Failed</h1><p>State mismatch.</p><p>You can close this window.</p></body></html>")
			return
		}

		resultCh <- CallbackResult{Code: code, State: state}
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "<html><body><h1>Authentication Successful!</h1><p>You can close this window and return to the application.</p></body></html>")
	})

	srv := &http.Server{
		Addr:              fmt.Sprintf("localhost:%d", port),
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("OAuth callback server error: %v", err)
		}
	}()

	return srv
}

// fetchUserEmail fetches the authenticated user's email address using the OAuth2 userinfo endpoint.
func fetchUserEmail(tok *oauth2.Token) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodGet, "https://www.googleapis.com/oauth2/v2/userinfo", http.NoBody)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+tok.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("fetching user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("userinfo returned %d: %s", resp.StatusCode, body)
	}

	var info struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return "", fmt.Errorf("decoding user info: %w", err)
	}
	if info.Email == "" {
		return "", errors.New("no email in userinfo response")
	}
	return info.Email, nil
}
