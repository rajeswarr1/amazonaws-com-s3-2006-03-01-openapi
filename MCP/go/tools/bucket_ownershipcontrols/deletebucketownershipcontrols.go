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

func DeletebucketownershipcontrolsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["ownershipControls"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ownershipControls=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/%s#ownershipControls%s", cfg.BaseURL, Bucket, queryString)
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

func CreateDeletebucketownershipcontrolsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_Bucket#ownershipControls",
		mcp.WithDescription("<p>Removes <code>OwnershipControls</code> for an Amazon S3 bucket. To use this operation, you must have the <code>s3:PutBucketOwnershipControls</code> permission. For more information about Amazon S3 permissions, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html">Specifying Permissions in a Policy</a>.</p> <p>For information about Amazon S3 Object Ownership, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/about-object-ownership.html">Using Object Ownership</a>. </p> <p>The following operations are related to <code>DeleteBucketOwnershipControls</code>:</p> <ul> <li> <p> <a>GetBucketOwnershipControls</a> </p> </li> <li> <p> <a>PutBucketOwnershipControls</a> </p> </li> </ul>"),
		mcp.WithString("Bucket", mcp.Required(), mcp.Description("The Amazon S3 bucket whose <code>OwnershipControls</code> you want to delete. ")),
		mcp.WithString("x-amz-expected-bucket-owner", mcp.Description("The account ID of the expected bucket owner. If the bucket is owned by a different account, the request fails with the HTTP status code <code>403 Forbidden</code> (access denied).")),
		mcp.WithBoolean("ownershipControls", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletebucketownershipcontrolsHandler(cfg),
	}
}
