package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/authorized-buyers-marketplace-api/mcp-server/config"
	"github.com/authorized-buyers-marketplace-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Authorizedbuyersmarketplace_buyers_publisherprofiles_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		parentVal, ok := args["parent"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: parent"), nil
		}
		parent, ok := parentVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: parent"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["filter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter=%v", val))
		}
		if val, ok := args["pageSize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageSize=%v", val))
		}
		if val, ok := args["pageToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageToken=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("access_token=%s", cfg.BearerToken))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("oauth_token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/%s/publisherProfiles%s", cfg.BaseURL, parent, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.ListPublisherProfilesResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateAuthorizedbuyersmarketplace_buyers_publisherprofiles_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_parent_publisherProfiles",
		mcp.WithDescription("Lists publisher profiles. The returned publisher profiles aren't in any defined order. The order of the results might change. A new publisher profile can appear in any place in the list of returned results."),
		mcp.WithString("parent", mcp.Required(), mcp.Description("Required. Parent that owns the collection of publisher profiles Format: `buyers/{buyerId}`")),
		mcp.WithString("filter", mcp.Description("Optional query string using the [Cloud API list filtering] (https://developers.google.com/authorized-buyers/apis/guides/list-filters) syntax.")),
		mcp.WithNumber("pageSize", mcp.Description("Requested page size. The server may return fewer results than requested. If requested more than 500, the server will return 500 results per page. If unspecified, the server will pick a default page size of 100.")),
		mcp.WithString("pageToken", mcp.Description("The page token as returned from a previous ListPublisherProfilesResponse.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Authorizedbuyersmarketplace_buyers_publisherprofiles_listHandler(cfg),
	}
}
