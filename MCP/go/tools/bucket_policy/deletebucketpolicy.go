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

func DeletebucketpolicyHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["policy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("policy=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/%s#policy%s", cfg.BaseURL, Bucket, queryString)
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

func CreateDeletebucketpolicyTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_Bucket#policy",
		mcp.WithDescription("<p>This implementation of the DELETE action uses the policy subresource to delete the policy of a specified bucket. If you are using an identity other than the root user of the Amazon Web Services account that owns the bucket, the calling identity must have the <code>DeleteBucketPolicy</code> permissions on the specified bucket and belong to the bucket owner's account to use this operation. </p> <p>If you don't have <code>DeleteBucketPolicy</code> permissions, Amazon S3 returns a <code>403 Access Denied</code> error. If you have the correct permissions, but you're not using an identity that belongs to the bucket owner's account, Amazon S3 returns a <code>405 Method Not Allowed</code> error. </p> <important> <p>To ensure that bucket owners don't inadvertently lock themselves out of their own buckets, the root principal in a bucket owner's Amazon Web Services account can perform the <code>GetBucketPolicy</code>, <code>PutBucketPolicy</code>, and <code>DeleteBucketPolicy</code> API actions, even if their bucket policy explicitly denies the root principal's access. Bucket owner root principals can only be blocked from performing these API actions by VPC endpoint policies and Amazon Web Services Organizations policies.</p> </important> <p>For more information about bucket policies, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html">Using Bucket Policies and UserPolicies</a>. </p> <p>The following operations are related to <code>DeleteBucketPolicy</code> </p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html">CreateBucket</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html">DeleteObject</a> </p> </li> </ul>"),
		mcp.WithString("Bucket", mcp.Required(), mcp.Description("The bucket name.")),
		mcp.WithString("x-amz-expected-bucket-owner", mcp.Description("The account ID of the expected bucket owner. If the bucket is owned by a different account, the request fails with the HTTP status code <code>403 Forbidden</code> (access denied).")),
		mcp.WithBoolean("policy", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletebucketpolicyHandler(cfg),
	}
}
