package tools

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"

	"github.com/shotah/google-workspace-mcp-go/auth"
	google "github.com/shotah/google-workspace-mcp-go/internal/google"
)

// RegisterAuthTools registers the start_google_auth meta-tool.
// It is filtered out when MCP_ENABLE_OAUTH21=true.
func RegisterAuthTools(s *mcpserver.MCPServer) {
	if isOAuth21Enabled() {
		return
	}

	store := auth.NewCredentialStore()

	s.AddTool(
		mcp.NewTool("start_google_auth",
			mcp.WithDescription(
				"Manually initiate Google OAuth authentication flow. "+
					"NOTE: This is a legacy OAuth 2.0 tool and is disabled when OAuth 2.1 is enabled. "+
					"The authentication system automatically handles credential checks and prompts for "+
					"authentication when needed. Only use this tool if: "+
					"1. You need to re-authenticate with different credentials. "+
					"2. You want to proactively authenticate before using other tools. "+
					"3. The automatic authentication flow failed and you need to retry. "+
					"In most cases, simply try calling the Google Workspace tool you need - it will "+
					"automatically handle authentication if required.",
			),
			mcp.WithString("service_name",
				mcp.Required(),
				mcp.Description("Name of the Google service requiring authentication (e.g., 'Google Calendar', 'Gmail')"),
			),
			mcp.WithString("user_google_email",
				mcp.Description("User's Google email address for authentication"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			serviceName, err := request.RequireString("service_name")
			if err != nil {
				return mcp.NewToolResultError("service_name is required"), nil
			}

			userEmail, emailErr := resolveEmail(request)
			if emailErr != nil {
				return mcp.NewToolResultError(emailErr.Error()), nil
			}

			msg, err := auth.StartAuthFlow(ctx, serviceName, userEmail, store, func(email string) {
				// Invalidate cached client so next tool call picks up new credentials.
				google.DefaultClientCache().Invalidate(email)
			})
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("**Authentication Error:** %v", err)), nil
			}

			return mcp.NewToolResultText(msg), nil
		},
	)
}

// isOAuth21Enabled checks if the MCP_ENABLE_OAUTH21 env var is set to "true".
func isOAuth21Enabled() bool {
	return strings.ToLower(os.Getenv("MCP_ENABLE_OAUTH21")) == "true"
}
