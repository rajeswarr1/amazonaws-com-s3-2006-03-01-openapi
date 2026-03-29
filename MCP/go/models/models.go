package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ServerSideEncryptionByDefault represents the ServerSideEncryptionByDefault schema from the OpenAPI specification
type ServerSideEncryptionByDefault struct {
	Kmsmasterkeyid interface{} `json:"KMSMasterKeyID,omitempty"`
	Ssealgorithm interface{} `json:"SSEAlgorithm"`
}

// PutBucketOwnershipControlsRequest represents the PutBucketOwnershipControlsRequest schema from the OpenAPI specification
type PutBucketOwnershipControlsRequest struct {
	Ownershipcontrols interface{} `json:"OwnershipControls"`
}

// SSES3 represents the SSES3 schema from the OpenAPI specification
type SSES3 struct {
}

// ErrorDocument represents the ErrorDocument schema from the OpenAPI specification
type ErrorDocument struct {
	Key interface{} `json:"Key"`
}

// FilterRule represents the FilterRule schema from the OpenAPI specification
type FilterRule struct {
	Name interface{} `json:"Name,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// CompletedPart represents the CompletedPart schema from the OpenAPI specification
type CompletedPart struct {
	Checksumsha256 interface{} `json:"ChecksumSHA256,omitempty"`
	Etag interface{} `json:"ETag,omitempty"`
	Partnumber interface{} `json:"PartNumber,omitempty"`
	Checksumcrc32 interface{} `json:"ChecksumCRC32,omitempty"`
	Checksumcrc32c interface{} `json:"ChecksumCRC32C,omitempty"`
	Checksumsha1 interface{} `json:"ChecksumSHA1,omitempty"`
}

// GetBucketOwnershipControlsRequest represents the GetBucketOwnershipControlsRequest schema from the OpenAPI specification
type GetBucketOwnershipControlsRequest struct {
}

// PutBucketNotificationConfigurationRequest represents the PutBucketNotificationConfigurationRequest schema from the OpenAPI specification
type PutBucketNotificationConfigurationRequest struct {
	Notificationconfiguration NotificationConfiguration `json:"NotificationConfiguration"` // A container for specifying the notification configuration of the bucket. If this element is empty, notifications are turned off for the bucket.
}

// DeleteBucketRequest represents the DeleteBucketRequest schema from the OpenAPI specification
type DeleteBucketRequest struct {
}

// Redirect represents the Redirect schema from the OpenAPI specification
type Redirect struct {
	Replacekeyprefixwith interface{} `json:"ReplaceKeyPrefixWith,omitempty"`
	Replacekeywith interface{} `json:"ReplaceKeyWith,omitempty"`
	Hostname interface{} `json:"HostName,omitempty"`
	Httpredirectcode interface{} `json:"HttpRedirectCode,omitempty"`
	Protocol interface{} `json:"Protocol,omitempty"`
}

// UploadPartCopyRequest represents the UploadPartCopyRequest schema from the OpenAPI specification
type UploadPartCopyRequest struct {
}

// GetBucketLocationOutput represents the GetBucketLocationOutput schema from the OpenAPI specification
type GetBucketLocationOutput struct {
	Locationconstraint interface{} `json:"LocationConstraint,omitempty"`
}

// SelectObjectContentEventStream represents the SelectObjectContentEventStream schema from the OpenAPI specification
type SelectObjectContentEventStream struct {
	Progress interface{} `json:"Progress,omitempty"`
	Records interface{} `json:"Records,omitempty"`
	Stats interface{} `json:"Stats,omitempty"`
	Cont interface{} `json:"Cont,omitempty"`
	End interface{} `json:"End,omitempty"`
}

// ReplicationRule represents the ReplicationRule schema from the OpenAPI specification
type ReplicationRule struct {
	Deletemarkerreplication DeleteMarkerReplication `json:"DeleteMarkerReplication,omitempty"` // <p>Specifies whether Amazon S3 replicates delete markers. If you specify a <code>Filter</code> in your replication configuration, you must also include a <code>DeleteMarkerReplication</code> element. If your <code>Filter</code> includes a <code>Tag</code> element, the <code>DeleteMarkerReplication</code> <code>Status</code> must be set to Disabled, because Amazon S3 does not support replicating delete markers for tag-based rules. For an example configuration, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-config-min-rule-config">Basic Rule Configuration</a>. </p> <p>For more information about delete marker replication, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/delete-marker-replication.html">Basic Rule Configuration</a>. </p> <note> <p>If you are using an earlier version of the replication configuration, Amazon S3 handles replication of delete markers differently. For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-backward-compat-considerations">Backward Compatibility</a>.</p> </note>
	Destination interface{} `json:"Destination"`
	Priority interface{} `json:"Priority,omitempty"`
	Existingobjectreplication interface{} `json:"ExistingObjectReplication,omitempty"`
	Id interface{} `json:"ID,omitempty"`
	Prefix interface{} `json:"Prefix,omitempty"`
	Sourceselectioncriteria interface{} `json:"SourceSelectionCriteria,omitempty"`
	Filter ReplicationRuleFilter `json:"Filter,omitempty"` // A filter that identifies the subset of objects to which the replication rule applies. A <code>Filter</code> must specify exactly one <code>Prefix</code>, <code>Tag</code>, or an <code>And</code> child element.
	Status interface{} `json:"Status"`
}

// ObjectLockRetention represents the ObjectLockRetention schema from the OpenAPI specification
type ObjectLockRetention struct {
	Mode interface{} `json:"Mode,omitempty"`
	Retainuntildate interface{} `json:"RetainUntilDate,omitempty"`
}

// PutBucketWebsiteRequest represents the PutBucketWebsiteRequest schema from the OpenAPI specification
type PutBucketWebsiteRequest struct {
	Websiteconfiguration interface{} `json:"WebsiteConfiguration"`
}

// AbortMultipartUploadRequest represents the AbortMultipartUploadRequest schema from the OpenAPI specification
type AbortMultipartUploadRequest struct {
}

// ListObjectVersionsOutput represents the ListObjectVersionsOutput schema from the OpenAPI specification
type ListObjectVersionsOutput struct {
	Istruncated interface{} `json:"IsTruncated,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Nextkeymarker interface{} `json:"NextKeyMarker,omitempty"`
	Versionidmarker interface{} `json:"VersionIdMarker,omitempty"`
	Versions interface{} `json:"Versions,omitempty"`
	Commonprefixes interface{} `json:"CommonPrefixes,omitempty"`
	Maxkeys interface{} `json:"MaxKeys,omitempty"`
	Encodingtype interface{} `json:"EncodingType,omitempty"`
	Nextversionidmarker interface{} `json:"NextVersionIdMarker,omitempty"`
	Deletemarkers interface{} `json:"DeleteMarkers,omitempty"`
	Delimiter interface{} `json:"Delimiter,omitempty"`
	Keymarker interface{} `json:"KeyMarker,omitempty"`
	Prefix interface{} `json:"Prefix,omitempty"`
}

// ListObjectsV2Output represents the ListObjectsV2Output schema from the OpenAPI specification
type ListObjectsV2Output struct {
	Prefix interface{} `json:"Prefix,omitempty"`
	Startafter interface{} `json:"StartAfter,omitempty"`
	Contents interface{} `json:"Contents,omitempty"`
	Continuationtoken interface{} `json:"ContinuationToken,omitempty"`
	Delimiter interface{} `json:"Delimiter,omitempty"`
	Istruncated interface{} `json:"IsTruncated,omitempty"`
	Nextcontinuationtoken interface{} `json:"NextContinuationToken,omitempty"`
	Encodingtype interface{} `json:"EncodingType,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Commonprefixes interface{} `json:"CommonPrefixes,omitempty"`
	Keycount interface{} `json:"KeyCount,omitempty"`
	Maxkeys interface{} `json:"MaxKeys,omitempty"`
}

// PutBucketAnalyticsConfigurationRequest represents the PutBucketAnalyticsConfigurationRequest schema from the OpenAPI specification
type PutBucketAnalyticsConfigurationRequest struct {
	Analyticsconfiguration interface{} `json:"AnalyticsConfiguration"`
}

// DeleteBucketReplicationRequest represents the DeleteBucketReplicationRequest schema from the OpenAPI specification
type DeleteBucketReplicationRequest struct {
}

// OwnershipControlsRule represents the OwnershipControlsRule schema from the OpenAPI specification
type OwnershipControlsRule struct {
	Objectownership string `json:"ObjectOwnership"` // <p>The container element for object ownership for a bucket's ownership controls.</p> <p>BucketOwnerPreferred - Objects uploaded to the bucket change ownership to the bucket owner if the objects are uploaded with the <code>bucket-owner-full-control</code> canned ACL.</p> <p>ObjectWriter - The uploading account will own the object if the object is uploaded with the <code>bucket-owner-full-control</code> canned ACL.</p> <p>BucketOwnerEnforced - Access control lists (ACLs) are disabled and no longer affect permissions. The bucket owner automatically owns and has full control over every object in the bucket. The bucket only accepts PUT requests that don't specify an ACL or bucket owner full control ACLs, such as the <code>bucket-owner-full-control</code> canned ACL or an equivalent form of this ACL expressed in the XML format.</p>
}

// DeleteBucketInventoryConfigurationRequest represents the DeleteBucketInventoryConfigurationRequest schema from the OpenAPI specification
type DeleteBucketInventoryConfigurationRequest struct {
}

// GetBucketEncryptionRequest represents the GetBucketEncryptionRequest schema from the OpenAPI specification
type GetBucketEncryptionRequest struct {
}

// PutBucketEncryptionRequest represents the PutBucketEncryptionRequest schema from the OpenAPI specification
type PutBucketEncryptionRequest struct {
	Serversideencryptionconfiguration ServerSideEncryptionConfiguration `json:"ServerSideEncryptionConfiguration"` // Specifies the default server-side-encryption configuration.
}

// DeleteObjectsOutput represents the DeleteObjectsOutput schema from the OpenAPI specification
type DeleteObjectsOutput struct {
	Deleted interface{} `json:"Deleted,omitempty"`
	Errors interface{} `json:"Errors,omitempty"`
}

// ObjectVersion represents the ObjectVersion schema from the OpenAPI specification
type ObjectVersion struct {
	Lastmodified interface{} `json:"LastModified,omitempty"`
	Versionid interface{} `json:"VersionId,omitempty"`
	Etag interface{} `json:"ETag,omitempty"`
	Islatest interface{} `json:"IsLatest,omitempty"`
	Key interface{} `json:"Key,omitempty"`
	Owner interface{} `json:"Owner,omitempty"`
	Size interface{} `json:"Size,omitempty"`
	Storageclass interface{} `json:"StorageClass,omitempty"`
	Restorestatus interface{} `json:"RestoreStatus,omitempty"`
	Checksumalgorithm interface{} `json:"ChecksumAlgorithm,omitempty"`
}

// JSONInput represents the JSONInput schema from the OpenAPI specification
type JSONInput struct {
	TypeField interface{} `json:"Type,omitempty"`
}

// PutBucketAccelerateConfigurationRequest represents the PutBucketAccelerateConfigurationRequest schema from the OpenAPI specification
type PutBucketAccelerateConfigurationRequest struct {
	Accelerateconfiguration interface{} `json:"AccelerateConfiguration"`
}

// CompleteMultipartUploadRequest represents the CompleteMultipartUploadRequest schema from the OpenAPI specification
type CompleteMultipartUploadRequest struct {
	Multipartupload interface{} `json:"MultipartUpload,omitempty"`
}

// GetObjectAttributesOutput represents the GetObjectAttributesOutput schema from the OpenAPI specification
type GetObjectAttributesOutput struct {
	Etag interface{} `json:"ETag,omitempty"`
	Objectparts interface{} `json:"ObjectParts,omitempty"`
	Objectsize interface{} `json:"ObjectSize,omitempty"`
	Storageclass interface{} `json:"StorageClass,omitempty"`
	Checksum interface{} `json:"Checksum,omitempty"`
}

// NotificationConfigurationDeprecated represents the NotificationConfigurationDeprecated schema from the OpenAPI specification
type NotificationConfigurationDeprecated struct {
	Topicconfiguration interface{} `json:"TopicConfiguration,omitempty"`
	Cloudfunctionconfiguration interface{} `json:"CloudFunctionConfiguration,omitempty"`
	Queueconfiguration interface{} `json:"QueueConfiguration,omitempty"`
}

// LifecycleConfiguration represents the LifecycleConfiguration schema from the OpenAPI specification
type LifecycleConfiguration struct {
	Rules interface{} `json:"Rules"`
}

// ListObjectsOutput represents the ListObjectsOutput schema from the OpenAPI specification
type ListObjectsOutput struct {
	Prefix interface{} `json:"Prefix,omitempty"`
	Commonprefixes interface{} `json:"CommonPrefixes,omitempty"`
	Delimiter interface{} `json:"Delimiter,omitempty"`
	Maxkeys interface{} `json:"MaxKeys,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Istruncated interface{} `json:"IsTruncated,omitempty"`
	Encodingtype interface{} `json:"EncodingType,omitempty"`
	Marker interface{} `json:"Marker,omitempty"`
	Contents interface{} `json:"Contents,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// GetObjectRetentionRequest represents the GetObjectRetentionRequest schema from the OpenAPI specification
type GetObjectRetentionRequest struct {
}

// GetObjectAttributesParts represents the GetObjectAttributesParts schema from the OpenAPI specification
type GetObjectAttributesParts struct {
	Totalpartscount interface{} `json:"TotalPartsCount,omitempty"`
	Istruncated interface{} `json:"IsTruncated,omitempty"`
	Maxparts interface{} `json:"MaxParts,omitempty"`
	Nextpartnumbermarker interface{} `json:"NextPartNumberMarker,omitempty"`
	Partnumbermarker interface{} `json:"PartNumberMarker,omitempty"`
	Parts interface{} `json:"Parts,omitempty"`
}

// DeleteBucketLifecycleRequest represents the DeleteBucketLifecycleRequest schema from the OpenAPI specification
type DeleteBucketLifecycleRequest struct {
}

// GetBucketPolicyStatusOutput represents the GetBucketPolicyStatusOutput schema from the OpenAPI specification
type GetBucketPolicyStatusOutput struct {
	Policystatus interface{} `json:"PolicyStatus,omitempty"`
}

// LifecycleRule represents the LifecycleRule schema from the OpenAPI specification
type LifecycleRule struct {
	Noncurrentversionexpiration NoncurrentVersionExpiration `json:"NoncurrentVersionExpiration,omitempty"` // Specifies when noncurrent object versions expire. Upon expiration, Amazon S3 permanently deletes the noncurrent object versions. You set this lifecycle configuration action on a bucket that has versioning enabled (or suspended) to request that Amazon S3 delete noncurrent object versions at a specific period in the object's lifetime.
	Noncurrentversiontransitions interface{} `json:"NoncurrentVersionTransitions,omitempty"`
	Status interface{} `json:"Status"`
	Transitions interface{} `json:"Transitions,omitempty"`
	Expiration interface{} `json:"Expiration,omitempty"`
	Filter interface{} `json:"Filter,omitempty"`
	Id interface{} `json:"ID,omitempty"`
	Prefix interface{} `json:"Prefix,omitempty"`
	Abortincompletemultipartupload AbortIncompleteMultipartUpload `json:"AbortIncompleteMultipartUpload,omitempty"` // Specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload. For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config"> Aborting Incomplete Multipart Uploads Using a Bucket Lifecycle Configuration</a> in the <i>Amazon S3 User Guide</i>.
}

// Stats represents the Stats schema from the OpenAPI specification
type Stats struct {
	Bytesscanned interface{} `json:"BytesScanned,omitempty"`
	Bytesprocessed interface{} `json:"BytesProcessed,omitempty"`
	Bytesreturned interface{} `json:"BytesReturned,omitempty"`
}

// GetBucketPolicyStatusRequest represents the GetBucketPolicyStatusRequest schema from the OpenAPI specification
type GetBucketPolicyStatusRequest struct {
}

// DeleteMarkerEntry represents the DeleteMarkerEntry schema from the OpenAPI specification
type DeleteMarkerEntry struct {
	Owner interface{} `json:"Owner,omitempty"`
	Versionid interface{} `json:"VersionId,omitempty"`
	Islatest interface{} `json:"IsLatest,omitempty"`
	Key interface{} `json:"Key,omitempty"`
	Lastmodified interface{} `json:"LastModified,omitempty"`
}

// StorageClassAnalysisDataExport represents the StorageClassAnalysisDataExport schema from the OpenAPI specification
type StorageClassAnalysisDataExport struct {
	Destination interface{} `json:"Destination"`
	Outputschemaversion interface{} `json:"OutputSchemaVersion"`
}

// DeleteMarkerReplication represents the DeleteMarkerReplication schema from the OpenAPI specification
type DeleteMarkerReplication struct {
	Status interface{} `json:"Status,omitempty"`
}

// RedirectAllRequestsTo represents the RedirectAllRequestsTo schema from the OpenAPI specification
type RedirectAllRequestsTo struct {
	Hostname interface{} `json:"HostName"`
	Protocol interface{} `json:"Protocol,omitempty"`
}

// GetBucketEncryptionOutput represents the GetBucketEncryptionOutput schema from the OpenAPI specification
type GetBucketEncryptionOutput struct {
	Serversideencryptionconfiguration ServerSideEncryptionConfiguration `json:"ServerSideEncryptionConfiguration,omitempty"` // Specifies the default server-side-encryption configuration.
}

// GetBucketAclRequest represents the GetBucketAclRequest schema from the OpenAPI specification
type GetBucketAclRequest struct {
}

// Destination represents the Destination schema from the OpenAPI specification
type Destination struct {
	Bucket interface{} `json:"Bucket"`
	Encryptionconfiguration interface{} `json:"EncryptionConfiguration,omitempty"`
	Metrics interface{} `json:"Metrics,omitempty"`
	Replicationtime interface{} `json:"ReplicationTime,omitempty"`
	Storageclass interface{} `json:"StorageClass,omitempty"`
	Accesscontroltranslation interface{} `json:"AccessControlTranslation,omitempty"`
	Account interface{} `json:"Account,omitempty"`
}

// GetBucketLifecycleConfigurationRequest represents the GetBucketLifecycleConfigurationRequest schema from the OpenAPI specification
type GetBucketLifecycleConfigurationRequest struct {
}

// GetBucketVersioningRequest represents the GetBucketVersioningRequest schema from the OpenAPI specification
type GetBucketVersioningRequest struct {
}

// IntelligentTieringConfiguration represents the IntelligentTieringConfiguration schema from the OpenAPI specification
type IntelligentTieringConfiguration struct {
	Filter interface{} `json:"Filter,omitempty"`
	Id interface{} `json:"Id"`
	Status interface{} `json:"Status"`
	Tierings interface{} `json:"Tierings"`
}

// PutObjectLegalHoldRequest represents the PutObjectLegalHoldRequest schema from the OpenAPI specification
type PutObjectLegalHoldRequest struct {
	Legalhold interface{} `json:"LegalHold,omitempty"`
}

// NotificationConfiguration represents the NotificationConfiguration schema from the OpenAPI specification
type NotificationConfiguration struct {
	Lambdafunctionconfigurations interface{} `json:"LambdaFunctionConfigurations,omitempty"`
	Queueconfigurations interface{} `json:"QueueConfigurations,omitempty"`
	Topicconfigurations interface{} `json:"TopicConfigurations,omitempty"`
	Eventbridgeconfiguration interface{} `json:"EventBridgeConfiguration,omitempty"`
}

// ExistingObjectReplication represents the ExistingObjectReplication schema from the OpenAPI specification
type ExistingObjectReplication struct {
	Status interface{} `json:"Status"`
}

// GetPublicAccessBlockRequest represents the GetPublicAccessBlockRequest schema from the OpenAPI specification
type GetPublicAccessBlockRequest struct {
}

// PublicAccessBlockConfiguration represents the PublicAccessBlockConfiguration schema from the OpenAPI specification
type PublicAccessBlockConfiguration struct {
	Blockpublicacls interface{} `json:"BlockPublicAcls,omitempty"`
	Blockpublicpolicy interface{} `json:"BlockPublicPolicy,omitempty"`
	Ignorepublicacls interface{} `json:"IgnorePublicAcls,omitempty"`
	Restrictpublicbuckets interface{} `json:"RestrictPublicBuckets,omitempty"`
}

// GetBucketLifecycleOutput represents the GetBucketLifecycleOutput schema from the OpenAPI specification
type GetBucketLifecycleOutput struct {
	Rules interface{} `json:"Rules,omitempty"`
}

// AccessControlPolicy represents the AccessControlPolicy schema from the OpenAPI specification
type AccessControlPolicy struct {
	Grants interface{} `json:"Grants,omitempty"`
	Owner interface{} `json:"Owner,omitempty"`
}

// CORSRule represents the CORSRule schema from the OpenAPI specification
type CORSRule struct {
	Id interface{} `json:"ID,omitempty"`
	Maxageseconds interface{} `json:"MaxAgeSeconds,omitempty"`
	Allowedheaders interface{} `json:"AllowedHeaders,omitempty"`
	Allowedmethods interface{} `json:"AllowedMethods"`
	Allowedorigins interface{} `json:"AllowedOrigins"`
	Exposeheaders interface{} `json:"ExposeHeaders,omitempty"`
}

// DeletePublicAccessBlockRequest represents the DeletePublicAccessBlockRequest schema from the OpenAPI specification
type DeletePublicAccessBlockRequest struct {
}

// GetObjectRetentionOutput represents the GetObjectRetentionOutput schema from the OpenAPI specification
type GetObjectRetentionOutput struct {
	Retention interface{} `json:"Retention,omitempty"`
}

// GetBucketCorsRequest represents the GetBucketCorsRequest schema from the OpenAPI specification
type GetBucketCorsRequest struct {
}

// ListBucketAnalyticsConfigurationsRequest represents the ListBucketAnalyticsConfigurationsRequest schema from the OpenAPI specification
type ListBucketAnalyticsConfigurationsRequest struct {
}

// RestoreObjectRequest represents the RestoreObjectRequest schema from the OpenAPI specification
type RestoreObjectRequest struct {
	Restorerequest RestoreRequest `json:"RestoreRequest,omitempty"` // Container for restore job parameters.
}

// IntelligentTieringAndOperator represents the IntelligentTieringAndOperator schema from the OpenAPI specification
type IntelligentTieringAndOperator struct {
	Prefix interface{} `json:"Prefix,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ContinuationEvent represents the ContinuationEvent schema from the OpenAPI specification
type ContinuationEvent struct {
}

// ScanRange represents the ScanRange schema from the OpenAPI specification
type ScanRange struct {
	Start interface{} `json:"Start,omitempty"`
	End interface{} `json:"End,omitempty"`
}

// PutPublicAccessBlockRequest represents the PutPublicAccessBlockRequest schema from the OpenAPI specification
type PutPublicAccessBlockRequest struct {
	Publicaccessblockconfiguration interface{} `json:"PublicAccessBlockConfiguration"`
}

// SelectParameters represents the SelectParameters schema from the OpenAPI specification
type SelectParameters struct {
	Expressiontype interface{} `json:"ExpressionType"`
	Inputserialization interface{} `json:"InputSerialization"`
	Outputserialization interface{} `json:"OutputSerialization"`
	Expression interface{} `json:"Expression"`
}

// PutBucketLifecycleConfigurationRequest represents the PutBucketLifecycleConfigurationRequest schema from the OpenAPI specification
type PutBucketLifecycleConfigurationRequest struct {
	Lifecycleconfiguration interface{} `json:"LifecycleConfiguration,omitempty"`
}

// GetBucketVersioningOutput represents the GetBucketVersioningOutput schema from the OpenAPI specification
type GetBucketVersioningOutput struct {
	Mfadelete interface{} `json:"MFADelete,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// PutBucketVersioningRequest represents the PutBucketVersioningRequest schema from the OpenAPI specification
type PutBucketVersioningRequest struct {
	Versioningconfiguration interface{} `json:"VersioningConfiguration"`
}

// ProgressEvent represents the ProgressEvent schema from the OpenAPI specification
type ProgressEvent struct {
	Details interface{} `json:"Details,omitempty"`
}

// ListBucketMetricsConfigurationsOutput represents the ListBucketMetricsConfigurationsOutput schema from the OpenAPI specification
type ListBucketMetricsConfigurationsOutput struct {
	Continuationtoken interface{} `json:"ContinuationToken,omitempty"`
	Istruncated interface{} `json:"IsTruncated,omitempty"`
	Metricsconfigurationlist interface{} `json:"MetricsConfigurationList,omitempty"`
	Nextcontinuationtoken interface{} `json:"NextContinuationToken,omitempty"`
}

// MultipartUpload represents the MultipartUpload schema from the OpenAPI specification
type MultipartUpload struct {
	Checksumalgorithm interface{} `json:"ChecksumAlgorithm,omitempty"`
	Initiated interface{} `json:"Initiated,omitempty"`
	Initiator interface{} `json:"Initiator,omitempty"`
	Key interface{} `json:"Key,omitempty"`
	Owner interface{} `json:"Owner,omitempty"`
	Storageclass interface{} `json:"StorageClass,omitempty"`
	Uploadid interface{} `json:"UploadId,omitempty"`
}

// TargetGrant represents the TargetGrant schema from the OpenAPI specification
type TargetGrant struct {
	Grantee interface{} `json:"Grantee,omitempty"`
	Permission interface{} `json:"Permission,omitempty"`
}

// GlacierJobParameters represents the GlacierJobParameters schema from the OpenAPI specification
type GlacierJobParameters struct {
	Tier interface{} `json:"Tier"`
}

// AbortMultipartUploadOutput represents the AbortMultipartUploadOutput schema from the OpenAPI specification
type AbortMultipartUploadOutput struct {
}

// ReplicationConfiguration represents the ReplicationConfiguration schema from the OpenAPI specification
type ReplicationConfiguration struct {
	Role interface{} `json:"Role"`
	Rules interface{} `json:"Rules"`
}

// CompleteMultipartUploadOutput represents the CompleteMultipartUploadOutput schema from the OpenAPI specification
type CompleteMultipartUploadOutput struct {
	Location interface{} `json:"Location,omitempty"`
	Bucket interface{} `json:"Bucket,omitempty"`
	Checksumcrc32 interface{} `json:"ChecksumCRC32,omitempty"`
	Checksumcrc32c interface{} `json:"ChecksumCRC32C,omitempty"`
	Checksumsha1 interface{} `json:"ChecksumSHA1,omitempty"`
	Checksumsha256 interface{} `json:"ChecksumSHA256,omitempty"`
	Etag interface{} `json:"ETag,omitempty"`
	Key interface{} `json:"Key,omitempty"`
}

// PutBucketTaggingRequest represents the PutBucketTaggingRequest schema from the OpenAPI specification
type PutBucketTaggingRequest struct {
	Tagging interface{} `json:"Tagging"`
}

// DeleteBucketPolicyRequest represents the DeleteBucketPolicyRequest schema from the OpenAPI specification
type DeleteBucketPolicyRequest struct {
}

// S3KeyFilter represents the S3KeyFilter schema from the OpenAPI specification
type S3KeyFilter struct {
	Filterrules interface{} `json:"FilterRules,omitempty"`
}

// GetBucketRequestPaymentRequest represents the GetBucketRequestPaymentRequest schema from the OpenAPI specification
type GetBucketRequestPaymentRequest struct {
}

// ReplicationTime represents the ReplicationTime schema from the OpenAPI specification
type ReplicationTime struct {
	Time interface{} `json:"Time"`
	Status interface{} `json:"Status"`
}

// GetObjectTorrentOutput represents the GetObjectTorrentOutput schema from the OpenAPI specification
type GetObjectTorrentOutput struct {
	Body interface{} `json:"Body,omitempty"`
}

// ReplicaModifications represents the ReplicaModifications schema from the OpenAPI specification
type ReplicaModifications struct {
	Status interface{} `json:"Status"`
}

// PutObjectRequest represents the PutObjectRequest schema from the OpenAPI specification
type PutObjectRequest struct {
	Body interface{} `json:"Body,omitempty"`
	Metadata interface{} `json:"Metadata,omitempty"`
}

// GetBucketReplicationRequest represents the GetBucketReplicationRequest schema from the OpenAPI specification
type GetBucketReplicationRequest struct {
}

// GetBucketAnalyticsConfigurationRequest represents the GetBucketAnalyticsConfigurationRequest schema from the OpenAPI specification
type GetBucketAnalyticsConfigurationRequest struct {
}

// OwnershipControls represents the OwnershipControls schema from the OpenAPI specification
type OwnershipControls struct {
	Rules interface{} `json:"Rules"`
}

// ListBucketsOutput represents the ListBucketsOutput schema from the OpenAPI specification
type ListBucketsOutput struct {
	Buckets interface{} `json:"Buckets,omitempty"`
	Owner interface{} `json:"Owner,omitempty"`
}

// GetBucketPolicyRequest represents the GetBucketPolicyRequest schema from the OpenAPI specification
type GetBucketPolicyRequest struct {
}

// CreateBucketOutput represents the CreateBucketOutput schema from the OpenAPI specification
type CreateBucketOutput struct {
}

// GetBucketLifecycleRequest represents the GetBucketLifecycleRequest schema from the OpenAPI specification
type GetBucketLifecycleRequest struct {
}

// CommonPrefix represents the CommonPrefix schema from the OpenAPI specification
type CommonPrefix struct {
	Prefix interface{} `json:"Prefix,omitempty"`
}

// IndexDocument represents the IndexDocument schema from the OpenAPI specification
type IndexDocument struct {
	Suffix interface{} `json:"Suffix"`
}

// GetBucketOwnershipControlsOutput represents the GetBucketOwnershipControlsOutput schema from the OpenAPI specification
type GetBucketOwnershipControlsOutput struct {
	Ownershipcontrols interface{} `json:"OwnershipControls,omitempty"`
}

// BucketLoggingStatus represents the BucketLoggingStatus schema from the OpenAPI specification
type BucketLoggingStatus struct {
	Loggingenabled LoggingEnabled `json:"LoggingEnabled,omitempty"` // Describes where logs are stored and the prefix that Amazon S3 assigns to all log object keys for a bucket. For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlogging.html">PUT Bucket logging</a> in the <i>Amazon S3 API Reference</i>.
}

// DeleteBucketTaggingRequest represents the DeleteBucketTaggingRequest schema from the OpenAPI specification
type DeleteBucketTaggingRequest struct {
}

// CreateMultipartUploadOutput represents the CreateMultipartUploadOutput schema from the OpenAPI specification
type CreateMultipartUploadOutput struct {
	Uploadid interface{} `json:"UploadId,omitempty"`
	Bucket interface{} `json:"Bucket,omitempty"`
	Key interface{} `json:"Key,omitempty"`
}

// MetricsFilter represents the MetricsFilter schema from the OpenAPI specification
type MetricsFilter struct {
	Prefix interface{} `json:"Prefix,omitempty"`
	Tag interface{} `json:"Tag,omitempty"`
	Accesspointarn interface{} `json:"AccessPointArn,omitempty"`
	And interface{} `json:"And,omitempty"`
}

// AccessControlTranslation represents the AccessControlTranslation schema from the OpenAPI specification
type AccessControlTranslation struct {
	Owner interface{} `json:"Owner"`
}

// GetObjectLegalHoldOutput represents the GetObjectLegalHoldOutput schema from the OpenAPI specification
type GetObjectLegalHoldOutput struct {
	Legalhold interface{} `json:"LegalHold,omitempty"`
}

// IntelligentTieringFilter represents the IntelligentTieringFilter schema from the OpenAPI specification
type IntelligentTieringFilter struct {
	And interface{} `json:"And,omitempty"`
	Prefix interface{} `json:"Prefix,omitempty"`
	Tag Tag `json:"Tag,omitempty"` // A container of a key value name pair.
}

// ListPartsOutput represents the ListPartsOutput schema from the OpenAPI specification
type ListPartsOutput struct {
	Key interface{} `json:"Key,omitempty"`
	Maxparts interface{} `json:"MaxParts,omitempty"`
	Partnumbermarker interface{} `json:"PartNumberMarker,omitempty"`
	Storageclass interface{} `json:"StorageClass,omitempty"`
	Initiator interface{} `json:"Initiator,omitempty"`
	Parts interface{} `json:"Parts,omitempty"`
	Uploadid interface{} `json:"UploadId,omitempty"`
	Bucket interface{} `json:"Bucket,omitempty"`
	Checksumalgorithm interface{} `json:"ChecksumAlgorithm,omitempty"`
	Istruncated interface{} `json:"IsTruncated,omitempty"`
	Nextpartnumbermarker interface{} `json:"NextPartNumberMarker,omitempty"`
	Owner interface{} `json:"Owner,omitempty"`
}

// ObjectLockLegalHold represents the ObjectLockLegalHold schema from the OpenAPI specification
type ObjectLockLegalHold struct {
	Status interface{} `json:"Status,omitempty"`
}

// AbortIncompleteMultipartUpload represents the AbortIncompleteMultipartUpload schema from the OpenAPI specification
type AbortIncompleteMultipartUpload struct {
	Daysafterinitiation interface{} `json:"DaysAfterInitiation,omitempty"`
}

// HeadObjectOutput represents the HeadObjectOutput schema from the OpenAPI specification
type HeadObjectOutput struct {
	Metadata interface{} `json:"Metadata,omitempty"`
}

// PutObjectOutput represents the PutObjectOutput schema from the OpenAPI specification
type PutObjectOutput struct {
}

// HeadBucketRequest represents the HeadBucketRequest schema from the OpenAPI specification
type HeadBucketRequest struct {
}

// InventorySchedule represents the InventorySchedule schema from the OpenAPI specification
type InventorySchedule struct {
	Frequency interface{} `json:"Frequency"`
}

// LoggingEnabled represents the LoggingEnabled schema from the OpenAPI specification
type LoggingEnabled struct {
	Targetbucket interface{} `json:"TargetBucket"`
	Targetgrants interface{} `json:"TargetGrants,omitempty"`
	Targetprefix interface{} `json:"TargetPrefix"`
}

// GetObjectLegalHoldRequest represents the GetObjectLegalHoldRequest schema from the OpenAPI specification
type GetObjectLegalHoldRequest struct {
}

// GetObjectAttributesRequest represents the GetObjectAttributesRequest schema from the OpenAPI specification
type GetObjectAttributesRequest struct {
}

// ReplicationTimeValue represents the ReplicationTimeValue schema from the OpenAPI specification
type ReplicationTimeValue struct {
	Minutes interface{} `json:"Minutes,omitempty"`
}

// DefaultRetention represents the DefaultRetention schema from the OpenAPI specification
type DefaultRetention struct {
	Days interface{} `json:"Days,omitempty"`
	Mode interface{} `json:"Mode,omitempty"`
	Years interface{} `json:"Years,omitempty"`
}

// Tagging represents the Tagging schema from the OpenAPI specification
type Tagging struct {
	Tagset interface{} `json:"TagSet"`
}

// ListBucketInventoryConfigurationsOutput represents the ListBucketInventoryConfigurationsOutput schema from the OpenAPI specification
type ListBucketInventoryConfigurationsOutput struct {
	Continuationtoken interface{} `json:"ContinuationToken,omitempty"`
	Inventoryconfigurationlist interface{} `json:"InventoryConfigurationList,omitempty"`
	Istruncated interface{} `json:"IsTruncated,omitempty"`
	Nextcontinuationtoken interface{} `json:"NextContinuationToken,omitempty"`
}

// CloudFunctionConfiguration represents the CloudFunctionConfiguration schema from the OpenAPI specification
type CloudFunctionConfiguration struct {
	Invocationrole interface{} `json:"InvocationRole,omitempty"`
	Cloudfunction interface{} `json:"CloudFunction,omitempty"`
	Event interface{} `json:"Event,omitempty"`
	Events interface{} `json:"Events,omitempty"`
	Id string `json:"Id,omitempty"` // An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.
}

// GetBucketCorsOutput represents the GetBucketCorsOutput schema from the OpenAPI specification
type GetBucketCorsOutput struct {
	Corsrules interface{} `json:"CORSRules,omitempty"`
}

// PutObjectTaggingOutput represents the PutObjectTaggingOutput schema from the OpenAPI specification
type PutObjectTaggingOutput struct {
}

// DeleteBucketMetricsConfigurationRequest represents the DeleteBucketMetricsConfigurationRequest schema from the OpenAPI specification
type DeleteBucketMetricsConfigurationRequest struct {
}

// OutputLocation represents the OutputLocation schema from the OpenAPI specification
type OutputLocation struct {
	S3 interface{} `json:"S3,omitempty"`
}

// DeleteObjectTaggingRequest represents the DeleteObjectTaggingRequest schema from the OpenAPI specification
type DeleteObjectTaggingRequest struct {
}

// Owner represents the Owner schema from the OpenAPI specification
type Owner struct {
	Displayname interface{} `json:"DisplayName,omitempty"`
	Id interface{} `json:"ID,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// InventoryS3BucketDestination represents the InventoryS3BucketDestination schema from the OpenAPI specification
type InventoryS3BucketDestination struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Bucket interface{} `json:"Bucket"`
	Encryption interface{} `json:"Encryption,omitempty"`
	Format interface{} `json:"Format"`
	Prefix interface{} `json:"Prefix,omitempty"`
}

// InventoryFilter represents the InventoryFilter schema from the OpenAPI specification
type InventoryFilter struct {
	Prefix interface{} `json:"Prefix"`
}

// GetBucketAccelerateConfigurationOutput represents the GetBucketAccelerateConfigurationOutput schema from the OpenAPI specification
type GetBucketAccelerateConfigurationOutput struct {
	Status interface{} `json:"Status,omitempty"`
}

// LifecycleExpiration represents the LifecycleExpiration schema from the OpenAPI specification
type LifecycleExpiration struct {
	Date interface{} `json:"Date,omitempty"`
	Days interface{} `json:"Days,omitempty"`
	Expiredobjectdeletemarker interface{} `json:"ExpiredObjectDeleteMarker,omitempty"`
}

// ObjectLockRule represents the ObjectLockRule schema from the OpenAPI specification
type ObjectLockRule struct {
	Defaultretention interface{} `json:"DefaultRetention,omitempty"`
}

// EndEvent represents the EndEvent schema from the OpenAPI specification
type EndEvent struct {
}

// ObjectIdentifier represents the ObjectIdentifier schema from the OpenAPI specification
type ObjectIdentifier struct {
	Key interface{} `json:"Key"`
	Versionid interface{} `json:"VersionId,omitempty"`
}

// Bucket represents the Bucket schema from the OpenAPI specification
type Bucket struct {
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// GetBucketReplicationOutput represents the GetBucketReplicationOutput schema from the OpenAPI specification
type GetBucketReplicationOutput struct {
	Replicationconfiguration ReplicationConfiguration `json:"ReplicationConfiguration,omitempty"` // A container for replication rules. You can add up to 1,000 rules. The maximum size of a replication configuration is 2 MB.
}

// PutBucketCorsRequest represents the PutBucketCorsRequest schema from the OpenAPI specification
type PutBucketCorsRequest struct {
	Corsconfiguration interface{} `json:"CORSConfiguration"`
}

// PutObjectAclRequest represents the PutObjectAclRequest schema from the OpenAPI specification
type PutObjectAclRequest struct {
	Accesscontrolpolicy interface{} `json:"AccessControlPolicy,omitempty"`
}

// OutputSerialization represents the OutputSerialization schema from the OpenAPI specification
type OutputSerialization struct {
	Csv interface{} `json:"CSV,omitempty"`
	Json interface{} `json:"JSON,omitempty"`
}

// BucketLifecycleConfiguration represents the BucketLifecycleConfiguration schema from the OpenAPI specification
type BucketLifecycleConfiguration struct {
	Rules interface{} `json:"Rules"`
}

// SSEKMS represents the SSEKMS schema from the OpenAPI specification
type SSEKMS struct {
	Keyid interface{} `json:"KeyId"`
}

// InventoryConfiguration represents the InventoryConfiguration schema from the OpenAPI specification
type InventoryConfiguration struct {
	Optionalfields interface{} `json:"OptionalFields,omitempty"`
	Schedule interface{} `json:"Schedule"`
	Destination interface{} `json:"Destination"`
	Filter interface{} `json:"Filter,omitempty"`
	Id interface{} `json:"Id"`
	Includedobjectversions interface{} `json:"IncludedObjectVersions"`
	Isenabled interface{} `json:"IsEnabled"`
}

// ParquetInput represents the ParquetInput schema from the OpenAPI specification
type ParquetInput struct {
}

// NoncurrentVersionExpiration represents the NoncurrentVersionExpiration schema from the OpenAPI specification
type NoncurrentVersionExpiration struct {
	Newernoncurrentversions interface{} `json:"NewerNoncurrentVersions,omitempty"`
	Noncurrentdays interface{} `json:"NoncurrentDays,omitempty"`
}

// DeleteBucketAnalyticsConfigurationRequest represents the DeleteBucketAnalyticsConfigurationRequest schema from the OpenAPI specification
type DeleteBucketAnalyticsConfigurationRequest struct {
}

// PutBucketMetricsConfigurationRequest represents the PutBucketMetricsConfigurationRequest schema from the OpenAPI specification
type PutBucketMetricsConfigurationRequest struct {
	Metricsconfiguration interface{} `json:"MetricsConfiguration"`
}

// InputSerialization represents the InputSerialization schema from the OpenAPI specification
type InputSerialization struct {
	Csv interface{} `json:"CSV,omitempty"`
	Compressiontype interface{} `json:"CompressionType,omitempty"`
	Json interface{} `json:"JSON,omitempty"`
	Parquet interface{} `json:"Parquet,omitempty"`
}

// SelectObjectContentRequest represents the SelectObjectContentRequest schema from the OpenAPI specification
type SelectObjectContentRequest struct {
	Inputserialization interface{} `json:"InputSerialization"`
	Outputserialization interface{} `json:"OutputSerialization"`
	Requestprogress interface{} `json:"RequestProgress,omitempty"`
	Scanrange interface{} `json:"ScanRange,omitempty"`
	Expression interface{} `json:"Expression"`
	Expressiontype interface{} `json:"ExpressionType"`
}

// UploadPartOutput represents the UploadPartOutput schema from the OpenAPI specification
type UploadPartOutput struct {
}

// GetBucketWebsiteOutput represents the GetBucketWebsiteOutput schema from the OpenAPI specification
type GetBucketWebsiteOutput struct {
	Errordocument interface{} `json:"ErrorDocument,omitempty"`
	Indexdocument interface{} `json:"IndexDocument,omitempty"`
	Redirectallrequeststo interface{} `json:"RedirectAllRequestsTo,omitempty"`
	Routingrules interface{} `json:"RoutingRules,omitempty"`
}

// SelectObjectContentOutput represents the SelectObjectContentOutput schema from the OpenAPI specification
type SelectObjectContentOutput struct {
	Payload interface{} `json:"Payload,omitempty"`
}

// GetBucketLoggingRequest represents the GetBucketLoggingRequest schema from the OpenAPI specification
type GetBucketLoggingRequest struct {
}

// ListObjectVersionsRequest represents the ListObjectVersionsRequest schema from the OpenAPI specification
type ListObjectVersionsRequest struct {
}

// GetBucketMetricsConfigurationRequest represents the GetBucketMetricsConfigurationRequest schema from the OpenAPI specification
type GetBucketMetricsConfigurationRequest struct {
}

// NoncurrentVersionTransition represents the NoncurrentVersionTransition schema from the OpenAPI specification
type NoncurrentVersionTransition struct {
	Storageclass interface{} `json:"StorageClass,omitempty"`
	Newernoncurrentversions interface{} `json:"NewerNoncurrentVersions,omitempty"`
	Noncurrentdays interface{} `json:"NoncurrentDays,omitempty"`
}

// Delete represents the Delete schema from the OpenAPI specification
type Delete struct {
	Objects interface{} `json:"Objects"`
	Quiet interface{} `json:"Quiet,omitempty"`
}

// SseKmsEncryptedObjects represents the SseKmsEncryptedObjects schema from the OpenAPI specification
type SseKmsEncryptedObjects struct {
	Status interface{} `json:"Status"`
}

// PutBucketIntelligentTieringConfigurationRequest represents the PutBucketIntelligentTieringConfigurationRequest schema from the OpenAPI specification
type PutBucketIntelligentTieringConfigurationRequest struct {
	Intelligenttieringconfiguration interface{} `json:"IntelligentTieringConfiguration"`
}

// CSVOutput represents the CSVOutput schema from the OpenAPI specification
type CSVOutput struct {
	Recorddelimiter interface{} `json:"RecordDelimiter,omitempty"`
	Fielddelimiter interface{} `json:"FieldDelimiter,omitempty"`
	Quotecharacter interface{} `json:"QuoteCharacter,omitempty"`
	Quoteescapecharacter interface{} `json:"QuoteEscapeCharacter,omitempty"`
	Quotefields interface{} `json:"QuoteFields,omitempty"`
}

// UploadPartRequest represents the UploadPartRequest schema from the OpenAPI specification
type UploadPartRequest struct {
	Body interface{} `json:"Body,omitempty"`
}

// MetricsConfiguration represents the MetricsConfiguration schema from the OpenAPI specification
type MetricsConfiguration struct {
	Id interface{} `json:"Id"`
	Filter interface{} `json:"Filter,omitempty"`
}

// Initiator represents the Initiator schema from the OpenAPI specification
type Initiator struct {
	Id interface{} `json:"ID,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
}

// MetricsAndOperator represents the MetricsAndOperator schema from the OpenAPI specification
type MetricsAndOperator struct {
	Tags interface{} `json:"Tags,omitempty"`
	Accesspointarn interface{} `json:"AccessPointArn,omitempty"`
	Prefix interface{} `json:"Prefix,omitempty"`
}

// GetPublicAccessBlockOutput represents the GetPublicAccessBlockOutput schema from the OpenAPI specification
type GetPublicAccessBlockOutput struct {
	Publicaccessblockconfiguration interface{} `json:"PublicAccessBlockConfiguration,omitempty"`
}

// Condition represents the Condition schema from the OpenAPI specification
type Condition struct {
	Keyprefixequals interface{} `json:"KeyPrefixEquals,omitempty"`
	Httperrorcodereturnedequals interface{} `json:"HttpErrorCodeReturnedEquals,omitempty"`
}

// QueueConfigurationDeprecated represents the QueueConfigurationDeprecated schema from the OpenAPI specification
type QueueConfigurationDeprecated struct {
	Event interface{} `json:"Event,omitempty"`
	Events interface{} `json:"Events,omitempty"`
	Id string `json:"Id,omitempty"` // An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.
	Queue interface{} `json:"Queue,omitempty"`
}

// Part represents the Part schema from the OpenAPI specification
type Part struct {
	Partnumber interface{} `json:"PartNumber,omitempty"`
	Size interface{} `json:"Size,omitempty"`
	Checksumcrc32 interface{} `json:"ChecksumCRC32,omitempty"`
	Checksumcrc32c interface{} `json:"ChecksumCRC32C,omitempty"`
	Checksumsha1 interface{} `json:"ChecksumSHA1,omitempty"`
	Checksumsha256 interface{} `json:"ChecksumSHA256,omitempty"`
	Etag interface{} `json:"ETag,omitempty"`
	Lastmodified interface{} `json:"LastModified,omitempty"`
}

// UploadPartCopyOutput represents the UploadPartCopyOutput schema from the OpenAPI specification
type UploadPartCopyOutput struct {
	Copypartresult interface{} `json:"CopyPartResult,omitempty"`
}

// GetObjectAclRequest represents the GetObjectAclRequest schema from the OpenAPI specification
type GetObjectAclRequest struct {
}

// GetBucketMetricsConfigurationOutput represents the GetBucketMetricsConfigurationOutput schema from the OpenAPI specification
type GetBucketMetricsConfigurationOutput struct {
	Metricsconfiguration interface{} `json:"MetricsConfiguration,omitempty"`
}

// RestoreStatus represents the RestoreStatus schema from the OpenAPI specification
type RestoreStatus struct {
	Restoreexpirydate interface{} `json:"RestoreExpiryDate,omitempty"`
	Isrestoreinprogress interface{} `json:"IsRestoreInProgress,omitempty"`
}

// DeleteObjectsRequest represents the DeleteObjectsRequest schema from the OpenAPI specification
type DeleteObjectsRequest struct {
	DeleteField interface{} `json:"Delete"`
}

// CopyObjectOutput represents the CopyObjectOutput schema from the OpenAPI specification
type CopyObjectOutput struct {
	Copyobjectresult interface{} `json:"CopyObjectResult,omitempty"`
}

// WebsiteConfiguration represents the WebsiteConfiguration schema from the OpenAPI specification
type WebsiteConfiguration struct {
	Errordocument interface{} `json:"ErrorDocument,omitempty"`
	Indexdocument interface{} `json:"IndexDocument,omitempty"`
	Redirectallrequeststo interface{} `json:"RedirectAllRequestsTo,omitempty"`
	Routingrules interface{} `json:"RoutingRules,omitempty"`
}

// PutBucketReplicationRequest represents the PutBucketReplicationRequest schema from the OpenAPI specification
type PutBucketReplicationRequest struct {
	Replicationconfiguration ReplicationConfiguration `json:"ReplicationConfiguration"` // A container for replication rules. You can add up to 1,000 rules. The maximum size of a replication configuration is 2 MB.
}

// SourceSelectionCriteria represents the SourceSelectionCriteria schema from the OpenAPI specification
type SourceSelectionCriteria struct {
	Replicamodifications interface{} `json:"ReplicaModifications,omitempty"`
	Ssekmsencryptedobjects interface{} `json:"SseKmsEncryptedObjects,omitempty"`
}

// GetBucketIntelligentTieringConfigurationRequest represents the GetBucketIntelligentTieringConfigurationRequest schema from the OpenAPI specification
type GetBucketIntelligentTieringConfigurationRequest struct {
}

// AnalyticsAndOperator represents the AnalyticsAndOperator schema from the OpenAPI specification
type AnalyticsAndOperator struct {
	Prefix interface{} `json:"Prefix,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// DeleteBucketCorsRequest represents the DeleteBucketCorsRequest schema from the OpenAPI specification
type DeleteBucketCorsRequest struct {
}

// GetBucketAclOutput represents the GetBucketAclOutput schema from the OpenAPI specification
type GetBucketAclOutput struct {
	Grants interface{} `json:"Grants,omitempty"`
	Owner interface{} `json:"Owner,omitempty"`
}

// ListBucketInventoryConfigurationsRequest represents the ListBucketInventoryConfigurationsRequest schema from the OpenAPI specification
type ListBucketInventoryConfigurationsRequest struct {
}

// ListBucketMetricsConfigurationsRequest represents the ListBucketMetricsConfigurationsRequest schema from the OpenAPI specification
type ListBucketMetricsConfigurationsRequest struct {
}

// StatsEvent represents the StatsEvent schema from the OpenAPI specification
type StatsEvent struct {
	Details interface{} `json:"Details,omitempty"`
}

// DeletedObject represents the DeletedObject schema from the OpenAPI specification
type DeletedObject struct {
	Deletemarker interface{} `json:"DeleteMarker,omitempty"`
	Deletemarkerversionid interface{} `json:"DeleteMarkerVersionId,omitempty"`
	Key interface{} `json:"Key,omitempty"`
	Versionid interface{} `json:"VersionId,omitempty"`
}

// Encryption represents the Encryption schema from the OpenAPI specification
type Encryption struct {
	Encryptiontype interface{} `json:"EncryptionType"`
	Kmscontext interface{} `json:"KMSContext,omitempty"`
	Kmskeyid interface{} `json:"KMSKeyId,omitempty"`
}

// PutBucketLifecycleRequest represents the PutBucketLifecycleRequest schema from the OpenAPI specification
type PutBucketLifecycleRequest struct {
	Lifecycleconfiguration interface{} `json:"LifecycleConfiguration,omitempty"`
}

// PutBucketLoggingRequest represents the PutBucketLoggingRequest schema from the OpenAPI specification
type PutBucketLoggingRequest struct {
	Bucketloggingstatus interface{} `json:"BucketLoggingStatus"`
}

// PutObjectLockConfigurationRequest represents the PutObjectLockConfigurationRequest schema from the OpenAPI specification
type PutObjectLockConfigurationRequest struct {
	Objectlockconfiguration interface{} `json:"ObjectLockConfiguration,omitempty"`
}

// PutObjectLockConfigurationOutput represents the PutObjectLockConfigurationOutput schema from the OpenAPI specification
type PutObjectLockConfigurationOutput struct {
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Code interface{} `json:"Code,omitempty"`
	Key interface{} `json:"Key,omitempty"`
	Message interface{} `json:"Message,omitempty"`
	Versionid interface{} `json:"VersionId,omitempty"`
}

// DeleteObjectRequest represents the DeleteObjectRequest schema from the OpenAPI specification
type DeleteObjectRequest struct {
}

// LifecycleRuleAndOperator represents the LifecycleRuleAndOperator schema from the OpenAPI specification
type LifecycleRuleAndOperator struct {
	Objectsizegreaterthan interface{} `json:"ObjectSizeGreaterThan,omitempty"`
	Objectsizelessthan interface{} `json:"ObjectSizeLessThan,omitempty"`
	Prefix interface{} `json:"Prefix,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// GetObjectOutput represents the GetObjectOutput schema from the OpenAPI specification
type GetObjectOutput struct {
	Body interface{} `json:"Body,omitempty"`
	Metadata interface{} `json:"Metadata,omitempty"`
}

// RequestProgress represents the RequestProgress schema from the OpenAPI specification
type RequestProgress struct {
	Enabled interface{} `json:"Enabled,omitempty"`
}

// CreateBucketRequest represents the CreateBucketRequest schema from the OpenAPI specification
type CreateBucketRequest struct {
	Createbucketconfiguration interface{} `json:"CreateBucketConfiguration,omitempty"`
}

// ReplicationRuleFilter represents the ReplicationRuleFilter schema from the OpenAPI specification
type ReplicationRuleFilter struct {
	And interface{} `json:"And,omitempty"`
	Prefix interface{} `json:"Prefix,omitempty"`
	Tag interface{} `json:"Tag,omitempty"`
}

// AnalyticsFilter represents the AnalyticsFilter schema from the OpenAPI specification
type AnalyticsFilter struct {
	Prefix interface{} `json:"Prefix,omitempty"`
	Tag interface{} `json:"Tag,omitempty"`
	And interface{} `json:"And,omitempty"`
}

// GetBucketAccelerateConfigurationRequest represents the GetBucketAccelerateConfigurationRequest schema from the OpenAPI specification
type GetBucketAccelerateConfigurationRequest struct {
}

// Metrics represents the Metrics schema from the OpenAPI specification
type Metrics struct {
	Eventthreshold interface{} `json:"EventThreshold,omitempty"`
	Status interface{} `json:"Status"`
}

// Grant represents the Grant schema from the OpenAPI specification
type Grant struct {
	Grantee interface{} `json:"Grantee,omitempty"`
	Permission interface{} `json:"Permission,omitempty"`
}

// PutObjectAclOutput represents the PutObjectAclOutput schema from the OpenAPI specification
type PutObjectAclOutput struct {
}

// CopyObjectRequest represents the CopyObjectRequest schema from the OpenAPI specification
type CopyObjectRequest struct {
	Metadata interface{} `json:"Metadata,omitempty"`
}

// RecordsEvent represents the RecordsEvent schema from the OpenAPI specification
type RecordsEvent struct {
	Payload interface{} `json:"Payload,omitempty"`
}

// EncryptionConfiguration represents the EncryptionConfiguration schema from the OpenAPI specification
type EncryptionConfiguration struct {
	Replicakmskeyid interface{} `json:"ReplicaKmsKeyID,omitempty"`
}

// DeleteBucketOwnershipControlsRequest represents the DeleteBucketOwnershipControlsRequest schema from the OpenAPI specification
type DeleteBucketOwnershipControlsRequest struct {
}

// GetObjectTorrentRequest represents the GetObjectTorrentRequest schema from the OpenAPI specification
type GetObjectTorrentRequest struct {
}

// DeleteBucketWebsiteRequest represents the DeleteBucketWebsiteRequest schema from the OpenAPI specification
type DeleteBucketWebsiteRequest struct {
}

// LambdaFunctionConfiguration represents the LambdaFunctionConfiguration schema from the OpenAPI specification
type LambdaFunctionConfiguration struct {
	Events interface{} `json:"Events"`
	Filter NotificationConfigurationFilter `json:"Filter,omitempty"` // Specifies object key name filtering rules. For information about key name filtering, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/notification-how-to-filtering.html">Configuring event notifications using object key name filtering</a> in the <i>Amazon S3 User Guide</i>.
	Id string `json:"Id,omitempty"` // An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.
	Lambdafunctionarn interface{} `json:"LambdaFunctionArn"`
}

// GetBucketTaggingRequest represents the GetBucketTaggingRequest schema from the OpenAPI specification
type GetBucketTaggingRequest struct {
}

// GetBucketNotificationConfigurationRequest represents the GetBucketNotificationConfigurationRequest schema from the OpenAPI specification
type GetBucketNotificationConfigurationRequest struct {
}

// Rule represents the Rule schema from the OpenAPI specification
type Rule struct {
	Abortincompletemultipartupload AbortIncompleteMultipartUpload `json:"AbortIncompleteMultipartUpload,omitempty"` // Specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload. For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config"> Aborting Incomplete Multipart Uploads Using a Bucket Lifecycle Configuration</a> in the <i>Amazon S3 User Guide</i>.
	Expiration interface{} `json:"Expiration,omitempty"`
	Id interface{} `json:"ID,omitempty"`
	Noncurrentversionexpiration NoncurrentVersionExpiration `json:"NoncurrentVersionExpiration,omitempty"` // Specifies when noncurrent object versions expire. Upon expiration, Amazon S3 permanently deletes the noncurrent object versions. You set this lifecycle configuration action on a bucket that has versioning enabled (or suspended) to request that Amazon S3 delete noncurrent object versions at a specific period in the object's lifetime.
	Noncurrentversiontransition NoncurrentVersionTransition `json:"NoncurrentVersionTransition,omitempty"` // Container for the transition rule that describes when noncurrent objects transition to the <code>STANDARD_IA</code>, <code>ONEZONE_IA</code>, <code>INTELLIGENT_TIERING</code>, <code>GLACIER_IR</code>, <code>GLACIER</code>, or <code>DEEP_ARCHIVE</code> storage class. If your bucket is versioning-enabled (or versioning is suspended), you can set this action to request that Amazon S3 transition noncurrent object versions to the <code>STANDARD_IA</code>, <code>ONEZONE_IA</code>, <code>INTELLIGENT_TIERING</code>, <code>GLACIER_IR</code>, <code>GLACIER</code>, or <code>DEEP_ARCHIVE</code> storage class at a specific period in the object's lifetime.
	Prefix interface{} `json:"Prefix"`
	Status interface{} `json:"Status"`
	Transition interface{} `json:"Transition,omitempty"`
}

// CopyPartResult represents the CopyPartResult schema from the OpenAPI specification
type CopyPartResult struct {
	Checksumsha1 interface{} `json:"ChecksumSHA1,omitempty"`
	Checksumsha256 interface{} `json:"ChecksumSHA256,omitempty"`
	Etag interface{} `json:"ETag,omitempty"`
	Lastmodified interface{} `json:"LastModified,omitempty"`
	Checksumcrc32 interface{} `json:"ChecksumCRC32,omitempty"`
	Checksumcrc32c interface{} `json:"ChecksumCRC32C,omitempty"`
}

// VersioningConfiguration represents the VersioningConfiguration schema from the OpenAPI specification
type VersioningConfiguration struct {
	Mfadelete interface{} `json:"MFADelete,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// GetObjectTaggingOutput represents the GetObjectTaggingOutput schema from the OpenAPI specification
type GetObjectTaggingOutput struct {
	Tagset interface{} `json:"TagSet"`
}

// InventoryEncryption represents the InventoryEncryption schema from the OpenAPI specification
type InventoryEncryption struct {
	Ssekms interface{} `json:"SSEKMS,omitempty"`
	Sses3 interface{} `json:"SSES3,omitempty"`
}

// HeadObjectRequest represents the HeadObjectRequest schema from the OpenAPI specification
type HeadObjectRequest struct {
}

// CreateMultipartUploadRequest represents the CreateMultipartUploadRequest schema from the OpenAPI specification
type CreateMultipartUploadRequest struct {
	Metadata interface{} `json:"Metadata,omitempty"`
}

// ListObjectsRequest represents the ListObjectsRequest schema from the OpenAPI specification
type ListObjectsRequest struct {
}

// RestoreObjectOutput represents the RestoreObjectOutput schema from the OpenAPI specification
type RestoreObjectOutput struct {
}

// ListPartsRequest represents the ListPartsRequest schema from the OpenAPI specification
type ListPartsRequest struct {
}

// GetBucketLifecycleConfigurationOutput represents the GetBucketLifecycleConfigurationOutput schema from the OpenAPI specification
type GetBucketLifecycleConfigurationOutput struct {
	Rules interface{} `json:"Rules,omitempty"`
}

// DeleteObjectOutput represents the DeleteObjectOutput schema from the OpenAPI specification
type DeleteObjectOutput struct {
}

// PutObjectLegalHoldOutput represents the PutObjectLegalHoldOutput schema from the OpenAPI specification
type PutObjectLegalHoldOutput struct {
}

// RestoreRequest represents the RestoreRequest schema from the OpenAPI specification
type RestoreRequest struct {
	Days interface{} `json:"Days,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Glacierjobparameters interface{} `json:"GlacierJobParameters,omitempty"`
	Outputlocation interface{} `json:"OutputLocation,omitempty"`
	Selectparameters interface{} `json:"SelectParameters,omitempty"`
	Tier interface{} `json:"Tier,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// CreateBucketConfiguration represents the CreateBucketConfiguration schema from the OpenAPI specification
type CreateBucketConfiguration struct {
	Locationconstraint interface{} `json:"LocationConstraint,omitempty"`
}

// WriteGetObjectResponseRequest represents the WriteGetObjectResponseRequest schema from the OpenAPI specification
type WriteGetObjectResponseRequest struct {
	Body interface{} `json:"Body,omitempty"`
	Metadata interface{} `json:"Metadata,omitempty"`
}

// ListBucketIntelligentTieringConfigurationsOutput represents the ListBucketIntelligentTieringConfigurationsOutput schema from the OpenAPI specification
type ListBucketIntelligentTieringConfigurationsOutput struct {
	Nextcontinuationtoken interface{} `json:"NextContinuationToken,omitempty"`
	Continuationtoken interface{} `json:"ContinuationToken,omitempty"`
	Intelligenttieringconfigurationlist interface{} `json:"IntelligentTieringConfigurationList,omitempty"`
	Istruncated interface{} `json:"IsTruncated,omitempty"`
}

// ServerSideEncryptionRule represents the ServerSideEncryptionRule schema from the OpenAPI specification
type ServerSideEncryptionRule struct {
	Applyserversideencryptionbydefault interface{} `json:"ApplyServerSideEncryptionByDefault,omitempty"`
	Bucketkeyenabled interface{} `json:"BucketKeyEnabled,omitempty"`
}

// ListObjectsV2Request represents the ListObjectsV2Request schema from the OpenAPI specification
type ListObjectsV2Request struct {
}

// PolicyStatus represents the PolicyStatus schema from the OpenAPI specification
type PolicyStatus struct {
	Ispublic interface{} `json:"IsPublic,omitempty"`
}

// StorageClassAnalysis represents the StorageClassAnalysis schema from the OpenAPI specification
type StorageClassAnalysis struct {
	Dataexport interface{} `json:"DataExport,omitempty"`
}

// PutBucketAclRequest represents the PutBucketAclRequest schema from the OpenAPI specification
type PutBucketAclRequest struct {
	Accesscontrolpolicy interface{} `json:"AccessControlPolicy,omitempty"`
}

// TopicConfiguration represents the TopicConfiguration schema from the OpenAPI specification
type TopicConfiguration struct {
	Topicarn interface{} `json:"TopicArn"`
	Events interface{} `json:"Events"`
	Filter NotificationConfigurationFilter `json:"Filter,omitempty"` // Specifies object key name filtering rules. For information about key name filtering, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/notification-how-to-filtering.html">Configuring event notifications using object key name filtering</a> in the <i>Amazon S3 User Guide</i>.
	Id string `json:"Id,omitempty"` // An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.
}

// DeleteObjectTaggingOutput represents the DeleteObjectTaggingOutput schema from the OpenAPI specification
type DeleteObjectTaggingOutput struct {
}

// GetObjectRequest represents the GetObjectRequest schema from the OpenAPI specification
type GetObjectRequest struct {
}

// PutObjectTaggingRequest represents the PutObjectTaggingRequest schema from the OpenAPI specification
type PutObjectTaggingRequest struct {
	Tagging interface{} `json:"Tagging"`
}

// PutObjectRetentionOutput represents the PutObjectRetentionOutput schema from the OpenAPI specification
type PutObjectRetentionOutput struct {
}

// ListBucketIntelligentTieringConfigurationsRequest represents the ListBucketIntelligentTieringConfigurationsRequest schema from the OpenAPI specification
type ListBucketIntelligentTieringConfigurationsRequest struct {
}

// S3Location represents the S3Location schema from the OpenAPI specification
type S3Location struct {
	Storageclass interface{} `json:"StorageClass,omitempty"`
	Tagging interface{} `json:"Tagging,omitempty"`
	Usermetadata interface{} `json:"UserMetadata,omitempty"`
	Accesscontrollist interface{} `json:"AccessControlList,omitempty"`
	Bucketname interface{} `json:"BucketName"`
	Cannedacl interface{} `json:"CannedACL,omitempty"`
	Encryption Encryption `json:"Encryption,omitempty"` // Contains the type of server-side encryption used.
	Prefix interface{} `json:"Prefix"`
}

// ListBucketAnalyticsConfigurationsOutput represents the ListBucketAnalyticsConfigurationsOutput schema from the OpenAPI specification
type ListBucketAnalyticsConfigurationsOutput struct {
	Istruncated interface{} `json:"IsTruncated,omitempty"`
	Nextcontinuationtoken interface{} `json:"NextContinuationToken,omitempty"`
	Analyticsconfigurationlist interface{} `json:"AnalyticsConfigurationList,omitempty"`
	Continuationtoken interface{} `json:"ContinuationToken,omitempty"`
}

// Object represents the Object schema from the OpenAPI specification
type Object struct {
	Owner interface{} `json:"Owner,omitempty"`
	Restorestatus interface{} `json:"RestoreStatus,omitempty"`
	Size interface{} `json:"Size,omitempty"`
	Storageclass interface{} `json:"StorageClass,omitempty"`
	Checksumalgorithm interface{} `json:"ChecksumAlgorithm,omitempty"`
	Etag interface{} `json:"ETag,omitempty"`
	Key interface{} `json:"Key,omitempty"`
	Lastmodified interface{} `json:"LastModified,omitempty"`
}

// RoutingRule represents the RoutingRule schema from the OpenAPI specification
type RoutingRule struct {
	Condition interface{} `json:"Condition,omitempty"`
	Redirect interface{} `json:"Redirect"`
}

// CopyObjectResult represents the CopyObjectResult schema from the OpenAPI specification
type CopyObjectResult struct {
	Checksumcrc32 interface{} `json:"ChecksumCRC32,omitempty"`
	Checksumcrc32c interface{} `json:"ChecksumCRC32C,omitempty"`
	Checksumsha1 interface{} `json:"ChecksumSHA1,omitempty"`
	Checksumsha256 interface{} `json:"ChecksumSHA256,omitempty"`
	Etag interface{} `json:"ETag,omitempty"`
	Lastmodified interface{} `json:"LastModified,omitempty"`
}

// Progress represents the Progress schema from the OpenAPI specification
type Progress struct {
	Bytesprocessed interface{} `json:"BytesProcessed,omitempty"`
	Bytesreturned interface{} `json:"BytesReturned,omitempty"`
	Bytesscanned interface{} `json:"BytesScanned,omitempty"`
}

// Transition represents the Transition schema from the OpenAPI specification
type Transition struct {
	Days interface{} `json:"Days,omitempty"`
	Storageclass interface{} `json:"StorageClass,omitempty"`
	Date interface{} `json:"Date,omitempty"`
}

// AnalyticsConfiguration represents the AnalyticsConfiguration schema from the OpenAPI specification
type AnalyticsConfiguration struct {
	Filter interface{} `json:"Filter,omitempty"`
	Id interface{} `json:"Id"`
	Storageclassanalysis interface{} `json:"StorageClassAnalysis"`
}

// DeleteBucketEncryptionRequest represents the DeleteBucketEncryptionRequest schema from the OpenAPI specification
type DeleteBucketEncryptionRequest struct {
}

// RequestPaymentConfiguration represents the RequestPaymentConfiguration schema from the OpenAPI specification
type RequestPaymentConfiguration struct {
	Payer interface{} `json:"Payer"`
}

// NotificationConfigurationFilter represents the NotificationConfigurationFilter schema from the OpenAPI specification
type NotificationConfigurationFilter struct {
	Key interface{} `json:"Key,omitempty"`
}

// GetObjectAclOutput represents the GetObjectAclOutput schema from the OpenAPI specification
type GetObjectAclOutput struct {
	Grants interface{} `json:"Grants,omitempty"`
	Owner interface{} `json:"Owner,omitempty"`
}

// PutBucketRequestPaymentRequest represents the PutBucketRequestPaymentRequest schema from the OpenAPI specification
type PutBucketRequestPaymentRequest struct {
	Requestpaymentconfiguration interface{} `json:"RequestPaymentConfiguration"`
}

// GetObjectTaggingRequest represents the GetObjectTaggingRequest schema from the OpenAPI specification
type GetObjectTaggingRequest struct {
}

// EventBridgeConfiguration represents the EventBridgeConfiguration schema from the OpenAPI specification
type EventBridgeConfiguration struct {
}

// GetBucketInventoryConfigurationRequest represents the GetBucketInventoryConfigurationRequest schema from the OpenAPI specification
type GetBucketInventoryConfigurationRequest struct {
}

// GetBucketInventoryConfigurationOutput represents the GetBucketInventoryConfigurationOutput schema from the OpenAPI specification
type GetBucketInventoryConfigurationOutput struct {
	Inventoryconfiguration interface{} `json:"InventoryConfiguration,omitempty"`
}

// AnalyticsExportDestination represents the AnalyticsExportDestination schema from the OpenAPI specification
type AnalyticsExportDestination struct {
	S3bucketdestination interface{} `json:"S3BucketDestination"`
}

// AccelerateConfiguration represents the AccelerateConfiguration schema from the OpenAPI specification
type AccelerateConfiguration struct {
	Status interface{} `json:"Status,omitempty"`
}

// GetBucketIntelligentTieringConfigurationOutput represents the GetBucketIntelligentTieringConfigurationOutput schema from the OpenAPI specification
type GetBucketIntelligentTieringConfigurationOutput struct {
	Intelligenttieringconfiguration interface{} `json:"IntelligentTieringConfiguration,omitempty"`
}

// ReplicationRuleAndOperator represents the ReplicationRuleAndOperator schema from the OpenAPI specification
type ReplicationRuleAndOperator struct {
	Prefix interface{} `json:"Prefix,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// Checksum represents the Checksum schema from the OpenAPI specification
type Checksum struct {
	Checksumsha1 interface{} `json:"ChecksumSHA1,omitempty"`
	Checksumsha256 interface{} `json:"ChecksumSHA256,omitempty"`
	Checksumcrc32 interface{} `json:"ChecksumCRC32,omitempty"`
	Checksumcrc32c interface{} `json:"ChecksumCRC32C,omitempty"`
}

// Metadata represents the Metadata schema from the OpenAPI specification
type Metadata struct {
}

// GetBucketTaggingOutput represents the GetBucketTaggingOutput schema from the OpenAPI specification
type GetBucketTaggingOutput struct {
	Tagset interface{} `json:"TagSet"`
}

// MetadataEntry represents the MetadataEntry schema from the OpenAPI specification
type MetadataEntry struct {
	Name interface{} `json:"Name,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// ObjectPart represents the ObjectPart schema from the OpenAPI specification
type ObjectPart struct {
	Checksumsha1 interface{} `json:"ChecksumSHA1,omitempty"`
	Checksumsha256 interface{} `json:"ChecksumSHA256,omitempty"`
	Partnumber interface{} `json:"PartNumber,omitempty"`
	Size interface{} `json:"Size,omitempty"`
	Checksumcrc32 interface{} `json:"ChecksumCRC32,omitempty"`
	Checksumcrc32c interface{} `json:"ChecksumCRC32C,omitempty"`
}

// GetBucketRequestPaymentOutput represents the GetBucketRequestPaymentOutput schema from the OpenAPI specification
type GetBucketRequestPaymentOutput struct {
	Payer interface{} `json:"Payer,omitempty"`
}

// Tiering represents the Tiering schema from the OpenAPI specification
type Tiering struct {
	Accesstier interface{} `json:"AccessTier"`
	Days interface{} `json:"Days"`
}

// Grantee represents the Grantee schema from the OpenAPI specification
type Grantee struct {
	Displayname interface{} `json:"DisplayName,omitempty"`
	Emailaddress interface{} `json:"EmailAddress,omitempty"`
	Id interface{} `json:"ID,omitempty"`
	TypeField interface{} `json:"Type"`
	Uri interface{} `json:"URI,omitempty"`
}

// CORSConfiguration represents the CORSConfiguration schema from the OpenAPI specification
type CORSConfiguration struct {
	Corsrules interface{} `json:"CORSRules"`
}

// QueueConfiguration represents the QueueConfiguration schema from the OpenAPI specification
type QueueConfiguration struct {
	Events interface{} `json:"Events"`
	Filter NotificationConfigurationFilter `json:"Filter,omitempty"` // Specifies object key name filtering rules. For information about key name filtering, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/notification-how-to-filtering.html">Configuring event notifications using object key name filtering</a> in the <i>Amazon S3 User Guide</i>.
	Id string `json:"Id,omitempty"` // An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.
	Queuearn interface{} `json:"QueueArn"`
}

// ListMultipartUploadsRequest represents the ListMultipartUploadsRequest schema from the OpenAPI specification
type ListMultipartUploadsRequest struct {
}

// JSONOutput represents the JSONOutput schema from the OpenAPI specification
type JSONOutput struct {
	Recorddelimiter interface{} `json:"RecordDelimiter,omitempty"`
}

// CompletedMultipartUpload represents the CompletedMultipartUpload schema from the OpenAPI specification
type CompletedMultipartUpload struct {
	Parts interface{} `json:"Parts,omitempty"`
}

// AnalyticsS3BucketDestination represents the AnalyticsS3BucketDestination schema from the OpenAPI specification
type AnalyticsS3BucketDestination struct {
	Bucketaccountid interface{} `json:"BucketAccountId,omitempty"`
	Format interface{} `json:"Format"`
	Prefix interface{} `json:"Prefix,omitempty"`
	Bucket interface{} `json:"Bucket"`
}

// LifecycleRuleFilter represents the LifecycleRuleFilter schema from the OpenAPI specification
type LifecycleRuleFilter struct {
	And LifecycleRuleAndOperator `json:"And,omitempty"` // This is used in a Lifecycle Rule Filter to apply a logical AND to two or more predicates. The Lifecycle Rule will apply to any object matching all of the predicates configured inside the And operator.
	Objectsizegreaterthan interface{} `json:"ObjectSizeGreaterThan,omitempty"`
	Objectsizelessthan interface{} `json:"ObjectSizeLessThan,omitempty"`
	Prefix interface{} `json:"Prefix,omitempty"`
	Tag interface{} `json:"Tag,omitempty"`
}

// ObjectLockConfiguration represents the ObjectLockConfiguration schema from the OpenAPI specification
type ObjectLockConfiguration struct {
	Objectlockenabled interface{} `json:"ObjectLockEnabled,omitempty"`
	Rule interface{} `json:"Rule,omitempty"`
}

// PutBucketInventoryConfigurationRequest represents the PutBucketInventoryConfigurationRequest schema from the OpenAPI specification
type PutBucketInventoryConfigurationRequest struct {
	Inventoryconfiguration interface{} `json:"InventoryConfiguration"`
}

// GetBucketAnalyticsConfigurationOutput represents the GetBucketAnalyticsConfigurationOutput schema from the OpenAPI specification
type GetBucketAnalyticsConfigurationOutput struct {
	Analyticsconfiguration interface{} `json:"AnalyticsConfiguration,omitempty"`
}

// GetObjectLockConfigurationRequest represents the GetObjectLockConfigurationRequest schema from the OpenAPI specification
type GetObjectLockConfigurationRequest struct {
}

// GetObjectLockConfigurationOutput represents the GetObjectLockConfigurationOutput schema from the OpenAPI specification
type GetObjectLockConfigurationOutput struct {
	Objectlockconfiguration interface{} `json:"ObjectLockConfiguration,omitempty"`
}

// ServerSideEncryptionConfiguration represents the ServerSideEncryptionConfiguration schema from the OpenAPI specification
type ServerSideEncryptionConfiguration struct {
	Rules interface{} `json:"Rules"`
}

// GetBucketPolicyOutput represents the GetBucketPolicyOutput schema from the OpenAPI specification
type GetBucketPolicyOutput struct {
	Policy interface{} `json:"Policy,omitempty"`
}

// PutBucketNotificationRequest represents the PutBucketNotificationRequest schema from the OpenAPI specification
type PutBucketNotificationRequest struct {
	Notificationconfiguration interface{} `json:"NotificationConfiguration"`
}

// GetBucketLocationRequest represents the GetBucketLocationRequest schema from the OpenAPI specification
type GetBucketLocationRequest struct {
}

// ListMultipartUploadsOutput represents the ListMultipartUploadsOutput schema from the OpenAPI specification
type ListMultipartUploadsOutput struct {
	Istruncated interface{} `json:"IsTruncated,omitempty"`
	Nextuploadidmarker interface{} `json:"NextUploadIdMarker,omitempty"`
	Uploadidmarker interface{} `json:"UploadIdMarker,omitempty"`
	Bucket interface{} `json:"Bucket,omitempty"`
	Encodingtype interface{} `json:"EncodingType,omitempty"`
	Keymarker interface{} `json:"KeyMarker,omitempty"`
	Maxuploads interface{} `json:"MaxUploads,omitempty"`
	Nextkeymarker interface{} `json:"NextKeyMarker,omitempty"`
	Commonprefixes interface{} `json:"CommonPrefixes,omitempty"`
	Prefix interface{} `json:"Prefix,omitempty"`
	Uploads interface{} `json:"Uploads,omitempty"`
	Delimiter interface{} `json:"Delimiter,omitempty"`
}

// InventoryDestination represents the InventoryDestination schema from the OpenAPI specification
type InventoryDestination struct {
	S3bucketdestination interface{} `json:"S3BucketDestination"`
}

// PutObjectRetentionRequest represents the PutObjectRetentionRequest schema from the OpenAPI specification
type PutObjectRetentionRequest struct {
	Retention interface{} `json:"Retention,omitempty"`
}

// GetBucketLoggingOutput represents the GetBucketLoggingOutput schema from the OpenAPI specification
type GetBucketLoggingOutput struct {
	Loggingenabled LoggingEnabled `json:"LoggingEnabled,omitempty"` // Describes where logs are stored and the prefix that Amazon S3 assigns to all log object keys for a bucket. For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlogging.html">PUT Bucket logging</a> in the <i>Amazon S3 API Reference</i>.
}

// PutBucketPolicyRequest represents the PutBucketPolicyRequest schema from the OpenAPI specification
type PutBucketPolicyRequest struct {
	Policy interface{} `json:"Policy"`
}

// CSVInput represents the CSVInput schema from the OpenAPI specification
type CSVInput struct {
	Comments interface{} `json:"Comments,omitempty"`
	Fielddelimiter interface{} `json:"FieldDelimiter,omitempty"`
	Fileheaderinfo interface{} `json:"FileHeaderInfo,omitempty"`
	Quotecharacter interface{} `json:"QuoteCharacter,omitempty"`
	Quoteescapecharacter interface{} `json:"QuoteEscapeCharacter,omitempty"`
	Recorddelimiter interface{} `json:"RecordDelimiter,omitempty"`
	Allowquotedrecorddelimiter interface{} `json:"AllowQuotedRecordDelimiter,omitempty"`
}

// DeleteBucketIntelligentTieringConfigurationRequest represents the DeleteBucketIntelligentTieringConfigurationRequest schema from the OpenAPI specification
type DeleteBucketIntelligentTieringConfigurationRequest struct {
}

// GetBucketWebsiteRequest represents the GetBucketWebsiteRequest schema from the OpenAPI specification
type GetBucketWebsiteRequest struct {
}

// TopicConfigurationDeprecated represents the TopicConfigurationDeprecated schema from the OpenAPI specification
type TopicConfigurationDeprecated struct {
	Event interface{} `json:"Event,omitempty"`
	Events interface{} `json:"Events,omitempty"`
	Id string `json:"Id,omitempty"` // An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.
	Topic interface{} `json:"Topic,omitempty"`
}
