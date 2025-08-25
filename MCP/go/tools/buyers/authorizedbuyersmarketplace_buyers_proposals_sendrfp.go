package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/authorized-buyers-marketplace-api/mcp-server/config"
	"github.com/authorized-buyers-marketplace-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Authorizedbuyersmarketplace_buyers_proposals_sendrfpHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		buyerVal, ok := args["buyer"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: buyer"), nil
		}
		buyer, ok := buyerVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: buyer"), nil
		}
		queryParams := make([]string, 0)
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
		// Create properly typed request body using the generated schema
		var requestBody models.SendRfpRequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/v1/%s/proposals:sendRfp%s", cfg.BaseURL, buyer, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.Proposal
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

func CreateAuthorizedbuyersmarketplace_buyers_proposals_sendrfpTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_buyer_proposals_sendRfp",
		mcp.WithDescription("Sends a request for proposal (RFP) to a publisher to initiate the negotiation regarding certain inventory. In the RFP, buyers can specify the deal type, deal terms, start and end dates, targeting, and a message to the publisher. Once the RFP is sent, a proposal in `SELLER_REVIEW_REQUESTED` state will be created and returned in the response. The publisher may review your request and respond with detailed deals in the proposal."),
		mcp.WithString("buyer", mcp.Required(), mcp.Description("Required. The current buyer who is sending the RFP in the format: `buyers/{accountId}`.")),
		mcp.WithString("displayName", mcp.Description("Input parameter: Required. The display name of the proposal being created by this RFP.")),
		mcp.WithString("flightStartTime", mcp.Description("Input parameter: Required. Proposed flight start time of the RFP. A timestamp in RFC3339 UTC \"Zulu\" format. Note that the specified value will be truncated to a granularity of one second.")),
		mcp.WithObject("inventorySizeTargeting", mcp.Description("Input parameter: Represents the size of an ad unit that can be targeted on a bid request.")),
		mcp.WithObject("preferredDealTerms", mcp.Description("Input parameter: Pricing terms for Preferred Deals.")),
		mcp.WithObject("programmaticGuaranteedTerms", mcp.Description("Input parameter: Pricing terms for Programmatic Guaranteed Deals.")),
		mcp.WithString("flightEndTime", mcp.Description("Input parameter: Required. Proposed flight end time of the RFP. A timestamp in RFC3339 UTC \"Zulu\" format. Note that the specified value will be truncated to a granularity of one second.")),
		mcp.WithObject("estimatedGrossSpend", mcp.Description("Input parameter: Represents an amount of money with its currency type.")),
		mcp.WithObject("geoTargeting", mcp.Description("Input parameter: Generic targeting used for targeting dimensions that contains a list of included and excluded numeric IDs. This cannot be filtered using list filter syntax.")),
		mcp.WithArray("buyerContacts", mcp.Description("Input parameter: Contact information for the buyer.")),
		mcp.WithString("note", mcp.Description("Input parameter: A message that is sent to the publisher. Maximum length is 1024 characters.")),
		mcp.WithString("publisherProfile", mcp.Description("Input parameter: Required. The profile of the publisher who will receive this RFP in the format: `buyers/{accountId}/publisherProfiles/{publisherProfileId}`.")),
		mcp.WithString("client", mcp.Description("Input parameter: If the current buyer is sending the RFP on behalf of its client, use this field to specify the name of the client in the format: `buyers/{accountId}/clients/{clientAccountid}`.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Authorizedbuyersmarketplace_buyers_proposals_sendrfpHandler(cfg),
	}
}
