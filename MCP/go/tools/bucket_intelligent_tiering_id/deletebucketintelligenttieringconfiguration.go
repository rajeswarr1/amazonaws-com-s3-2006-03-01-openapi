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

func DeletebucketintelligenttieringconfigurationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("id=%v", val))
		}
		if val, ok := args["intelligent-tiering"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("intelligent-tiering=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/%s#intelligent-tiering&id%s", cfg.BaseURL, Bucket, queryString)
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

func CreateDeletebucketintelligenttieringconfigurationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_Bucket#intelligent-tiering&id",
		mcp.WithDescription("<p>Deletes the S3 Intelligent-Tiering configuration from the specified bucket.</p> <p>The S3 Intelligent-Tiering storage class is designed to optimize storage costs by automatically moving data to the most cost-effective storage access tier, without performance impact or operational overhead. S3 Intelligent-Tiering delivers automatic cost savings in three low latency and high throughput access tiers. To get the lowest storage cost on data that can be accessed in minutes to hours, you can choose to activate additional archiving capabilities.</p> <p>The S3 Intelligent-Tiering storage class is the ideal storage class for data with unknown, changing, or unpredictable access patterns, independent of object size or retention period. If the size of an object is less than 128 KB, it is not monitored and not eligible for auto-tiering. Smaller objects can be stored, but they are always charged at the Frequent Access tier rates in the S3 Intelligent-Tiering storage class.</p> <p>For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access">Storage class for automatically optimizing frequently and infrequently accessed objects</a>.</p> <p>Operations related to <code>DeleteBucketIntelligentTieringConfiguration</code> include: </p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketIntelligentTieringConfiguration.html">GetBucketIntelligentTieringConfiguration</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketIntelligentTieringConfiguration.html">PutBucketIntelligentTieringConfiguration</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketIntelligentTieringConfigurations.html">ListBucketIntelligentTieringConfigurations</a> </p> </li> </ul>"),
		mcp.WithString("Bucket", mcp.Required(), mcp.Description("The name of the Amazon S3 bucket whose configuration you want to modify or retrieve.")),
		mcp.WithString("id", mcp.Required(), mcp.Description("The ID used to identify the S3 Intelligent-Tiering configuration.")),
		mcp.WithBoolean("intelligent-tiering", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletebucketintelligenttieringconfigurationHandler(cfg),
	}
}
