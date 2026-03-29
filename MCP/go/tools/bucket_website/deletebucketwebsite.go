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

func DeletebucketwebsiteHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["website"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("website=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/%s#website%s", cfg.BaseURL, Bucket, queryString)
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

func CreateDeletebucketwebsiteTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_Bucket#website",
		mcp.WithDescription("<p>This action removes the website configuration for a bucket. Amazon S3 returns a <code>200 OK</code> response upon successfully deleting a website configuration on the specified bucket. You will get a <code>200 OK</code> response if the website configuration you are trying to delete does not exist on the bucket. Amazon S3 returns a <code>404</code> response if the bucket specified in the request does not exist.</p> <p>This DELETE action requires the <code>S3:DeleteBucketWebsite</code> permission. By default, only the bucket owner can delete the website configuration attached to a bucket. However, bucket owners can grant other users permission to delete the website configuration by writing a bucket policy granting them the <code>S3:DeleteBucketWebsite</code> permission. </p> <p>For more information about hosting websites, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html">Hosting Websites on Amazon S3</a>. </p> <p>The following operations are related to <code>DeleteBucketWebsite</code>:</p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketWebsite.html">GetBucketWebsite</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketWebsite.html">PutBucketWebsite</a> </p> </li> </ul>"),
		mcp.WithString("Bucket", mcp.Required(), mcp.Description("The bucket name for which you want to remove the website configuration. ")),
		mcp.WithString("x-amz-expected-bucket-owner", mcp.Description("The account ID of the expected bucket owner. If the bucket is owned by a different account, the request fails with the HTTP status code <code>403 Forbidden</code> (access denied).")),
		mcp.WithBoolean("website", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletebucketwebsiteHandler(cfg),
	}
}
