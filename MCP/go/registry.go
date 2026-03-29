package main

import (
	"github.com/amazon-simple-storage-service/mcp-server/config"
	"github.com/amazon-simple-storage-service/mcp-server/models"
	tools_bucket_publicaccessblock "github.com/amazon-simple-storage-service/mcp-server/tools/bucket_publicaccessblock"
	tools_bucket_lifecycle "github.com/amazon-simple-storage-service/mcp-server/tools/bucket_lifecycle"
	tools_bucket_ownershipcontrols "github.com/amazon-simple-storage-service/mcp-server/tools/bucket_ownershipcontrols"
	tools_bucket_cors "github.com/amazon-simple-storage-service/mcp-server/tools/bucket_cors"
	tools_bucket_intelligent_tiering_id "github.com/amazon-simple-storage-service/mcp-server/tools/bucket_intelligent_tiering_id"
	tools_bucket_metrics_id "github.com/amazon-simple-storage-service/mcp-server/tools/bucket_metrics_id"
	tools_bucket_encryption "github.com/amazon-simple-storage-service/mcp-server/tools/bucket_encryption"
	tools_bucket_tagging "github.com/amazon-simple-storage-service/mcp-server/tools/bucket_tagging"
	tools_bucket_policy "github.com/amazon-simple-storage-service/mcp-server/tools/bucket_policy"
	tools_bucket "github.com/amazon-simple-storage-service/mcp-server/tools/bucket"
	tools_bucket_analytics_id "github.com/amazon-simple-storage-service/mcp-server/tools/bucket_analytics_id"
	tools_bucket_website "github.com/amazon-simple-storage-service/mcp-server/tools/bucket_website"
	tools_bucket_replication "github.com/amazon-simple-storage-service/mcp-server/tools/bucket_replication"
	tools_bucket_inventory_id "github.com/amazon-simple-storage-service/mcp-server/tools/bucket_inventory_id"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_bucket_publicaccessblock.CreateDeletepublicaccessblockTool(cfg),
		tools_bucket_lifecycle.CreateDeletebucketlifecycleTool(cfg),
		tools_bucket_ownershipcontrols.CreateDeletebucketownershipcontrolsTool(cfg),
		tools_bucket_cors.CreateDeletebucketcorsTool(cfg),
		tools_bucket_intelligent_tiering_id.CreateDeletebucketintelligenttieringconfigurationTool(cfg),
		tools_bucket_metrics_id.CreateDeletebucketmetricsconfigurationTool(cfg),
		tools_bucket_encryption.CreateDeletebucketencryptionTool(cfg),
		tools_bucket_tagging.CreateDeletebuckettaggingTool(cfg),
		tools_bucket_policy.CreateDeletebucketpolicyTool(cfg),
		tools_bucket.CreateDeletebucketTool(cfg),
		tools_bucket_analytics_id.CreateDeletebucketanalyticsconfigurationTool(cfg),
		tools_bucket_website.CreateDeletebucketwebsiteTool(cfg),
		tools_bucket_replication.CreateDeletebucketreplicationTool(cfg),
		tools_bucket_inventory_id.CreateDeletebucketinventoryconfigurationTool(cfg),
	}
}
