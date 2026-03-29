package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-simple-storage-service/mcp-server/config"
	"github.com/amazon-simple-storage-service/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DeletebuckettaggingHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		BucketVal, ok := args["Bucket"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: Bucket"), nil
		}
		Bucket, ok := BucketVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: Bucket"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["tagging"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tagging=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/%s#tagging%s", cfg.BaseURL, Bucket, queryString)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("x-amz-security-token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["x-amz-expected-bucket-owner"]; ok {
			req.Header.Set("x-amz-expected-bucket-owner", fmt.Sprintf("%v", val))
		}

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
		var result map[string]interface{}
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

func CreateDeletebuckettaggingTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_Bucket#tagging",
		mcp.WithDescription("<p>Deletes the tags from the bucket.</p> <p>To use this operation, you must have permission to perform the <code>s3:PutBucketTagging</code> action. By default, the bucket owner has this permission and can grant this permission to others. </p> <p>The following operations are related to <code>DeleteBucketTagging</code>:</p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketTagging.html">GetBucketTagging</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html">PutBucketTagging</a> </p> </li> </ul>"),
		mcp.WithString("Bucket", mcp.Required(), mcp.Description("The bucket that has the tag set to be removed.")),
		mcp.WithString("x-amz-expected-bucket-owner", mcp.Description("The account ID of the expected bucket owner. If the bucket is owned by a different account, the request fails with the HTTP status code <code>403 Forbidden</code> (access denied).")),
		mcp.WithBoolean("tagging", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletebuckettaggingHandler(cfg),
	}
}
