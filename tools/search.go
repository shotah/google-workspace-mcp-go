package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	customsearch "google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"

	"github.com/magks/google-workspace-mcp-go/server"
)

// RegisterSearchTools registers all Search tools with the MCP server.
func RegisterSearchTools(s *mcpserver.MCPServer, _ server.Config) {
	registerSearchCustom(s)
	registerGetSearchEngineInfo(s)
	registerSearchCustomSiterestrict(s)
}

// newCustomSearchService creates a customsearch.Service using the API key from env.
func newCustomSearchService(ctx context.Context) (*customsearch.Service, string, error) {
	apiKey := os.Getenv("GOOGLE_PSE_API_KEY")
	if apiKey == "" {
		return nil, "", fmt.Errorf("GOOGLE_PSE_API_KEY environment variable not set. Please set it to your Google Custom Search API key")
	}

	cx := os.Getenv("GOOGLE_PSE_ENGINE_ID")
	if cx == "" {
		return nil, "", fmt.Errorf("GOOGLE_PSE_ENGINE_ID environment variable not set. Please set it to your Programmable Search Engine ID")
	}

	svc, err := customsearch.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, "", fmt.Errorf("creating Custom Search service: %w", err)
	}
	return svc, cx, nil
}

// executeSearch runs a Custom Search query and formats the results.
func executeSearch(ctx context.Context, email, q string, num, start int, safe string, searchType, siteSearch, siteSearchFilter, dateRestrict, fileType, language, country string) (string, error) {
	svc, cx, err := newCustomSearchService(ctx)
	if err != nil {
		return "", err
	}

	call := svc.Cse.List().Q(q).Cx(cx).Num(int64(num)).Start(int64(start)).Safe(safe)

	if searchType != "" {
		call = call.SearchType(searchType)
	}
	if siteSearch != "" {
		call = call.SiteSearch(siteSearch)
	}
	if siteSearchFilter != "" {
		call = call.SiteSearchFilter(siteSearchFilter)
	}
	if dateRestrict != "" {
		call = call.DateRestrict(dateRestrict)
	}
	if fileType != "" {
		call = call.FileType(fileType)
	}
	if language != "" {
		call = call.Lr(language)
	}
	if country != "" {
		call = call.Cr(country)
	}

	result, err := call.Context(ctx).Do()
	if err != nil {
		return "", fmt.Errorf("executing search: %w", err)
	}

	totalResults := "0"
	searchTime := 0.0
	if result.SearchInformation != nil {
		totalResults = result.SearchInformation.TotalResults
		searchTime = result.SearchInformation.SearchTime
	}

	itemCount := len(result.Items)

	var sb strings.Builder
	fmt.Fprintf(&sb, "Search Results for %s:\n", email)
	fmt.Fprintf(&sb, "- Query: \"%s\"\n", q)
	fmt.Fprintf(&sb, "- Search Engine ID: %s\n", cx)
	fmt.Fprintf(&sb, "- Total Results: %s\n", totalResults)
	fmt.Fprintf(&sb, "- Search Time: %.3f seconds\n", searchTime)
	fmt.Fprintf(&sb, "- Results Returned: %d (showing %d to %d)\n\n", itemCount, start, start+itemCount-1)

	if itemCount > 0 {
		sb.WriteString("Results:\n")
		for i, item := range result.Items {
			title := item.Title
			if title == "" {
				title = "No title"
			}
			link := item.Link
			if link == "" {
				link = "No link"
			}
			snippet := item.Snippet
			if snippet == "" {
				snippet = "No description available"
			}
			snippet = strings.ReplaceAll(snippet, "\n", " ")

			fmt.Fprintf(&sb, "\n%d. %s\n", start+i, title)
			fmt.Fprintf(&sb, "   URL: %s\n", link)
			fmt.Fprintf(&sb, "   Snippet: %s\n", snippet)

			// Add metadata from pagemap if available
			if len(item.Pagemap) > 0 {
				var pagemap map[string]any
				if err := json.Unmarshal(item.Pagemap, &pagemap); err == nil {
					if metatags, ok := pagemap["metatags"].([]any); ok && len(metatags) > 0 {
						if meta, ok := metatags[0].(map[string]any); ok {
							if ogType, ok := meta["og:type"].(string); ok {
								fmt.Fprintf(&sb, "   Type: %s\n", ogType)
							}
							if pubTime, ok := meta["article:published_time"].(string); ok && len(pubTime) >= 10 {
								fmt.Fprintf(&sb, "   Published: %s\n", pubTime[:10])
							}
						}
					}
				}
			}
		}
	} else {
		sb.WriteString("\nNo results found.")
	}

	// Pagination info
	if result.Queries != nil {
		if len(result.Queries.NextPage) > 0 {
			nextStart := result.Queries.NextPage[0].StartIndex
			fmt.Fprintf(&sb, "\n\nTo see more results, search again with start=%d", nextStart)
		}
	}

	return sb.String(), nil
}

// --- search_custom ---

func registerSearchCustom(s *mcpserver.MCPServer) {
	tool := mcp.NewTool("search_custom",
		mcp.WithDescription("Performs a search using Google Custom Search JSON API."),
		mcp.WithString("user_google_email",
			mcp.Required(),
			mcp.Description("The user's Google email address."),
		),
		mcp.WithString("q",
			mcp.Required(),
			mcp.Description("The search query."),
		),
		mcp.WithNumber("num",
			mcp.Description("Number of results to return (1-10). Defaults to 10."),
		),
		mcp.WithNumber("start",
			mcp.Description("The index of the first result to return (1-based). Defaults to 1."),
		),
		mcp.WithString("safe",
			mcp.Description("Safe search level. Defaults to \"off\"."),
			mcp.Enum("active", "moderate", "off"),
		),
		mcp.WithString("search_type",
			mcp.Description("Search for images if set to \"image\"."),
			mcp.Enum("image"),
		),
		mcp.WithString("site_search",
			mcp.Description("Restrict search to a specific site/domain."),
		),
		mcp.WithString("site_search_filter",
			mcp.Description("Exclude (\"e\") or include (\"i\") site_search results."),
			mcp.Enum("e", "i"),
		),
		mcp.WithString("date_restrict",
			mcp.Description("Restrict results by date (e.g., \"d5\" for past 5 days, \"m3\" for past 3 months)."),
		),
		mcp.WithString("file_type",
			mcp.Description("Filter by file type (e.g., \"pdf\", \"doc\")."),
		),
		mcp.WithString("language",
			mcp.Description("Language code for results (e.g., \"lang_en\")."),
		),
		mcp.WithString("country",
			mcp.Description("Country code for results (e.g., \"countryUS\")."),
		),
	)
	RegisterTool(s, tool, handleSearchCustom)
}

func handleSearchCustom(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	email, err := resolveEmail(request)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	q, err := request.RequireString("q")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	num := request.GetInt("num", 10)
	start := request.GetInt("start", 1)
	safe := request.GetString("safe", "off")
	searchType := request.GetString("search_type", "")
	siteSearch := request.GetString("site_search", "")
	siteSearchFilter := request.GetString("site_search_filter", "")
	dateRestrict := request.GetString("date_restrict", "")
	fileType := request.GetString("file_type", "")
	language := request.GetString("language", "")
	country := request.GetString("country", "")

	text, err := executeSearch(ctx, email, q, num, start, safe, searchType, siteSearch, siteSearchFilter, dateRestrict, fileType, language, country)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(text), nil
}

// --- get_search_engine_info ---

func registerGetSearchEngineInfo(s *mcpserver.MCPServer) {
	tool := mcp.NewTool("get_search_engine_info",
		mcp.WithDescription("Retrieves metadata about a Programmable Search Engine."),
		mcp.WithString("user_google_email",
			mcp.Required(),
			mcp.Description("The user's Google email address."),
		),
	)
	RegisterTool(s, tool, handleGetSearchEngineInfo)
}

func handleGetSearchEngineInfo(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	email, err := resolveEmail(request)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	svc, cx, err := newCustomSearchService(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	// Perform a minimal search to get the search engine context.
	result, err := svc.Cse.List().Q("test").Cx(cx).Num(1).Context(ctx).Do()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("querying search engine info: %v", err)), nil
	}

	// Context is googleapi.RawMessage (JSON bytes) — parse manually.
	title := "Unknown"
	var contextData map[string]any
	if len(result.Context) > 0 {
		_ = json.Unmarshal(result.Context, &contextData)
	}
	if contextData != nil {
		if t, ok := contextData["title"].(string); ok && t != "" {
			title = t
		}
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, "Search Engine Information for %s:\n", email)
	fmt.Fprintf(&sb, "- Search Engine ID: %s\n", cx)
	fmt.Fprintf(&sb, "- Title: %s\n", title)

	// Add facet/refinement information if available.
	if contextData != nil {
		if facets, ok := contextData["facets"].([]any); ok && len(facets) > 0 {
			sb.WriteString("\nAvailable Refinements:\n")
			for _, facet := range facets {
				if facetList, ok := facet.([]any); ok {
					for _, item := range facetList {
						if itemMap, ok := item.(map[string]any); ok {
							label := "Unknown"
							anchor := "Unknown"
							if l, ok := itemMap["label"].(string); ok && l != "" {
								label = l
							}
							if a, ok := itemMap["anchor"].(string); ok && a != "" {
								anchor = a
							}
							fmt.Fprintf(&sb, "  - %s (anchor: %s)\n", label, anchor)
						}
					}
				}
			}
		}
	}

	// Add search statistics.
	if result.SearchInformation != nil {
		totalResults := result.SearchInformation.TotalResults
		if totalResults == "" {
			totalResults = "Unknown"
		}
		sb.WriteString("\nSearch Statistics:\n")
		fmt.Fprintf(&sb, "  - Total indexed results: %s\n", totalResults)
	}

	return mcp.NewToolResultText(sb.String()), nil
}

// --- search_custom_siterestrict ---

func registerSearchCustomSiterestrict(s *mcpserver.MCPServer) {
	tool := mcp.NewTool("search_custom_siterestrict",
		mcp.WithDescription("Performs a search restricted to specific sites using Google Custom Search."),
		mcp.WithString("user_google_email",
			mcp.Required(),
			mcp.Description("The user's Google email address."),
		),
		mcp.WithString("q",
			mcp.Required(),
			mcp.Description("The search query."),
		),
		mcp.WithArray("sites",
			mcp.Required(),
			mcp.Description("List of sites/domains to search within."),
			mcp.Items(map[string]any{"type": "string"}),
		),
		mcp.WithNumber("num",
			mcp.Description("Number of results to return (1-10). Defaults to 10."),
		),
		mcp.WithNumber("start",
			mcp.Description("The index of the first result to return (1-based). Defaults to 1."),
		),
		mcp.WithString("safe",
			mcp.Description("Safe search level. Defaults to \"off\"."),
			mcp.Enum("active", "moderate", "off"),
		),
	)
	RegisterTool(s, tool, handleSearchCustomSiterestrict)
}

func handleSearchCustomSiterestrict(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	email, err := resolveEmail(request)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	q, err := request.RequireString("q")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	sites, err := request.RequireStringSlice("sites")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	num := request.GetInt("num", 10)
	start := request.GetInt("start", 1)
	safe := request.GetString("safe", "off")

	// Build site restriction query.
	siteParts := make([]string, len(sites))
	for i, site := range sites {
		siteParts[i] = "site:" + site
	}
	siteQuery := strings.Join(siteParts, " OR ")
	fullQuery := fmt.Sprintf("%s (%s)", q, siteQuery)

	text, err := executeSearch(ctx, email, fullQuery, num, start, safe, "", "", "", "", "", "", "")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(text), nil
}
