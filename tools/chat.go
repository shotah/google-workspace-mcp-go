package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	chat "google.golang.org/api/chat/v1"

	"github.com/magks/google-workspace-mcp-go/internal/google"
	"github.com/magks/google-workspace-mcp-go/server"
)

// RegisterChatTools registers all Chat tools with the MCP server.
func RegisterChatTools(s *mcpserver.MCPServer, _ server.Config) {
	getClient := clientFuncFromCache(google.DefaultClientCache())

	registerListSpaces(s, getClient)
	registerGetMessages(s, getClient)
	registerSendMessage(s, getClient)
	registerSearchMessages(s, getClient)
}

// newChatService creates a chat.Service for the given user email.
func newChatService(ctx context.Context, getClient httpClientFunc, email string) (*chat.Service, error) {
	httpClient, err := getClient(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("authenticating for %s: %w", email, err)
	}
	svc, err := chat.New(httpClient)
	if err != nil {
		return nil, fmt.Errorf("creating Chat service: %w", err)
	}
	return svc, nil
}

// --- list_spaces ---

func registerListSpaces(s *mcpserver.MCPServer, getClient httpClientFunc) {
	tool := mcp.NewTool("list_spaces",
		mcp.WithDescription("Lists Google Chat spaces (rooms and direct messages) accessible to the user."),
		mcp.WithString("user_google_email", mcp.Description("User's Google email address")),
		mcp.WithNumber("page_size", mcp.Description("Number of spaces to return per page. Defaults to 100.")),
		mcp.WithString("space_type", mcp.Description("Type of spaces to filter: 'all', 'room', or 'dm'. Defaults to 'all'.")),
	)
	s.AddTool(tool, handleListSpaces(getClient))
}

func handleListSpaces(getClient httpClientFunc) mcpserver.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		email, err := resolveEmail(request)
		if err != nil {
			return mcp.NewToolResultError("user_google_email is required"), nil
		}

		pageSize := request.GetInt("page_size", 100)
		spaceType := request.GetString("space_type", "all")

		svc, err := newChatService(ctx, getClient, email)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		call := svc.Spaces.List().PageSize(int64(pageSize))

		// Build filter based on space_type
		switch spaceType {
		case "room":
			call = call.Filter("spaceType = SPACE")
		case "dm":
			call = call.Filter("spaceType = DIRECT_MESSAGE")
		}

		resp, err := call.Do()
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("listing spaces: %v", err)), nil
		}

		spaces := resp.Spaces
		if len(spaces) == 0 {
			return mcp.NewToolResultText(fmt.Sprintf("No Chat spaces found for type '%s'.", spaceType)), nil
		}

		var out strings.Builder
		fmt.Fprintf(&out, "Found %d Chat spaces (type: %s):", len(spaces), spaceType)
		for _, space := range spaces {
			name := space.DisplayName
			if name == "" {
				name = "Unnamed Space"
			}
			spaceTypeActual := space.SpaceType
			if spaceTypeActual == "" {
				spaceTypeActual = "UNKNOWN"
			}
			fmt.Fprintf(&out, "\n- %s (ID: %s, Type: %s)", name, space.Name, spaceTypeActual)
		}

		return mcp.NewToolResultText(out.String()), nil
	}
}

// --- get_messages ---

func registerGetMessages(s *mcpserver.MCPServer, getClient httpClientFunc) {
	tool := mcp.NewTool("get_messages",
		mcp.WithDescription("Retrieves messages from a Google Chat space."),
		mcp.WithString("user_google_email", mcp.Description("User's Google email address")),
		mcp.WithString("space_id", mcp.Required(), mcp.Description("The ID of the Chat space to retrieve messages from")),
		mcp.WithNumber("page_size", mcp.Description("Number of messages to return per page. Defaults to 50.")),
		mcp.WithString("order_by", mcp.Description("Sort order for messages. Defaults to 'createTime desc'.")),
	)
	s.AddTool(tool, handleGetMessages(getClient))
}

func handleGetMessages(getClient httpClientFunc) mcpserver.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		email, err := resolveEmail(request)
		if err != nil {
			return mcp.NewToolResultError("user_google_email is required"), nil
		}
		spaceID, err := request.RequireString("space_id")
		if err != nil {
			return mcp.NewToolResultError("space_id is required"), nil
		}

		pageSize := request.GetInt("page_size", 50)
		orderBy := request.GetString("order_by", "createTime desc")

		svc, err := newChatService(ctx, getClient, email)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		// Get space info first
		spaceInfo, err := svc.Spaces.Get(spaceID).Do()
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("getting space info: %v", err)), nil
		}
		spaceName := spaceInfo.DisplayName
		if spaceName == "" {
			spaceName = "Unknown Space"
		}

		// Get messages
		resp, err := svc.Spaces.Messages.List(spaceID).
			PageSize(int64(pageSize)).
			OrderBy(orderBy).
			Do()
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("listing messages: %v", err)), nil
		}

		messages := resp.Messages
		if len(messages) == 0 {
			return mcp.NewToolResultText(fmt.Sprintf("No messages found in space '%s' (ID: %s).", spaceName, spaceID)), nil
		}

		var out strings.Builder
		fmt.Fprintf(&out, "Messages from '%s' (ID: %s):\n", spaceName, spaceID)
		for _, msg := range messages {
			sender := "Unknown Sender"
			if msg.Sender != nil && msg.Sender.DisplayName != "" {
				sender = msg.Sender.DisplayName
			}
			createTime := msg.CreateTime
			if createTime == "" {
				createTime = "Unknown Time"
			}
			text := msg.Text
			if text == "" {
				text = "No text content"
			}

			fmt.Fprintf(&out, "\n[%s] %s:", createTime, sender)
			fmt.Fprintf(&out, "\n  %s", text)
			fmt.Fprintf(&out, "\n  (Message ID: %s)\n", msg.Name)
		}

		return mcp.NewToolResultText(out.String()), nil
	}
}

// --- send_message ---

func registerSendMessage(s *mcpserver.MCPServer, getClient httpClientFunc) {
	tool := mcp.NewTool("send_message",
		mcp.WithDescription("Sends a message to a Google Chat space."),
		mcp.WithString("user_google_email", mcp.Description("User's Google email address")),
		mcp.WithString("space_id", mcp.Required(), mcp.Description("The ID of the Chat space to send the message to")),
		mcp.WithString("message_text", mcp.Required(), mcp.Description("The text content of the message to send")),
		mcp.WithString("thread_key", mcp.Description("Optional thread key for threaded replies")),
	)
	s.AddTool(tool, handleSendMessage(getClient))
}

func handleSendMessage(getClient httpClientFunc) mcpserver.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		email, err := resolveEmail(request)
		if err != nil {
			return mcp.NewToolResultError("user_google_email is required"), nil
		}
		spaceID, err := request.RequireString("space_id")
		if err != nil {
			return mcp.NewToolResultError("space_id is required"), nil
		}
		messageText, err := request.RequireString("message_text")
		if err != nil {
			return mcp.NewToolResultError("message_text is required"), nil
		}

		threadKey := request.GetString("thread_key", "")

		svc, err := newChatService(ctx, getClient, email)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		msgBody := &chat.Message{
			Text: messageText,
		}

		call := svc.Spaces.Messages.Create(spaceID, msgBody)
		if threadKey != "" {
			call = call.ThreadKey(threadKey)
		}

		message, err := call.Do()
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("sending message: %v", err)), nil
		}

		result := fmt.Sprintf("Message sent to space '%s' by %s. Message ID: %s, Time: %s",
			spaceID, email, message.Name, message.CreateTime)

		return mcp.NewToolResultText(result), nil
	}
}

// --- search_messages ---

func registerSearchMessages(s *mcpserver.MCPServer, getClient httpClientFunc) {
	tool := mcp.NewTool("search_messages",
		mcp.WithDescription("Searches for messages in Google Chat spaces by text content."),
		mcp.WithString("user_google_email", mcp.Description("User's Google email address")),
		mcp.WithString("query", mcp.Required(), mcp.Description("Search query text to find in messages")),
		mcp.WithString("space_id", mcp.Description("Optional space ID to limit search to a specific space; if not provided, searches all accessible spaces")),
		mcp.WithNumber("page_size", mcp.Description("Number of messages to return per page. Defaults to 25.")),
	)
	s.AddTool(tool, handleSearchMessages(getClient))
}

func handleSearchMessages(getClient httpClientFunc) mcpserver.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		email, err := resolveEmail(request)
		if err != nil {
			return mcp.NewToolResultError("user_google_email is required"), nil
		}
		query, err := request.RequireString("query")
		if err != nil {
			return mcp.NewToolResultError("query is required"), nil
		}

		spaceID := request.GetString("space_id", "")
		pageSize := request.GetInt("page_size", 25)

		svc, err := newChatService(ctx, getClient, email)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		type messageWithSpace struct {
			msg       *chat.Message
			spaceName string
		}

		var results []messageWithSpace
		var searchContext string

		filter := fmt.Sprintf(`text:"%s"`, query)

		if spaceID != "" {
			// Search within a specific space
			resp, err := svc.Spaces.Messages.List(spaceID).
				PageSize(int64(pageSize)).
				Filter(filter).
				Do()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("searching messages: %v", err)), nil
			}
			for _, msg := range resp.Messages {
				results = append(results, messageWithSpace{msg: msg, spaceName: ""})
			}
			searchContext = fmt.Sprintf("space '%s'", spaceID)
		} else {
			// Search across all accessible spaces
			spacesResp, err := svc.Spaces.List().PageSize(100).Do()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("listing spaces for search: %v", err)), nil
			}

			// Limit to first 10 spaces to avoid timeout
			spaces := spacesResp.Spaces
			if len(spaces) > 10 {
				spaces = spaces[:10]
			}

			for _, space := range spaces {
				resp, err := svc.Spaces.Messages.List(space.Name).
					PageSize(5).
					Filter(filter).
					Do()
				if err != nil {
					continue // Skip spaces we can't access
				}
				displayName := space.DisplayName
				if displayName == "" {
					displayName = "Unknown"
				}
				for _, msg := range resp.Messages {
					results = append(results, messageWithSpace{msg: msg, spaceName: displayName})
				}
			}
			searchContext = "all accessible spaces"
		}

		if len(results) == 0 {
			return mcp.NewToolResultText(fmt.Sprintf("No messages found matching '%s' in %s.", query, searchContext)), nil
		}

		var out strings.Builder
		fmt.Fprintf(&out, "Found %d messages matching '%s' in %s:", len(results), query, searchContext)
		for _, r := range results {
			sender := "Unknown Sender"
			if r.msg.Sender != nil && r.msg.Sender.DisplayName != "" {
				sender = r.msg.Sender.DisplayName
			}
			createTime := r.msg.CreateTime
			if createTime == "" {
				createTime = "Unknown Time"
			}
			text := r.msg.Text
			if text == "" {
				text = "No text content"
			}
			spaceName := r.spaceName
			if spaceName == "" {
				spaceName = "Unknown Space"
			}

			// Truncate long messages
			if len(text) > 100 {
				text = text[:100] + "..."
			}

			fmt.Fprintf(&out, "\n- [%s] %s in '%s': %s", createTime, sender, spaceName, text)
		}

		return mcp.NewToolResultText(out.String()), nil
	}
}
