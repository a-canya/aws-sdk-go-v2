// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Amazon Security Lake collects logs and events from supported Amazon Web
// Services and custom sources. For the list of supported Amazon Web Services, see
// the Amazon Security Lake User Guide (https://docs.aws.amazon.com/security-lake/latest/userguide/internal-sources.html)
// .
type AccountSources struct {

	// The ID of the Security Lake account for which logs are collected.
	//
	// This member is required.
	Account *string

	// The supported Amazon Web Services from which logs and events are collected.
	// Amazon Security Lake supports log and event collection for natively supported
	// Amazon Web Services.
	//
	// This member is required.
	SourceType *string

	// Initializes a new instance of the Event class.
	EventClass OcsfEventClass

	// The log status for the Security Lake account.
	LogsStatus []LogsStatus

	noSmithyDocumentSerde
}

// Automatically enable new organization accounts as member accounts from an
// Amazon Security Lake administrator account.
type AutoEnableNewRegionConfiguration struct {

	// The Amazon Web Services Regions where Security Lake is automatically enabled.
	//
	// This member is required.
	Region Region

	// The Amazon Web Services sources that are automatically enabled in Security Lake.
	//
	// This member is required.
	Sources []AwsLogSourceType

	noSmithyDocumentSerde
}

// List of all failures.
type Failures struct {

	// List of all exception messages.
	//
	// This member is required.
	ExceptionMessage *string

	// List of all remediation steps for failures.
	//
	// This member is required.
	Remediation *string

	// This error can occur if you configure the wrong timestamp format, or if the
	// subset of entries used for validation had errors or missing values.
	//
	// This member is required.
	Timestamp *time.Time

	noSmithyDocumentSerde
}

// Response element for actions that make changes, namely create, update, or
// delete actions.
type FailuresResponse struct {

	// List of all failures.
	Failures []Failures

	// List of Amazon Web Services Regions where the failure occurred.
	Region *string

	noSmithyDocumentSerde
}

// Provides details of Amazon Security Lake configuration object.
type LakeConfigurationRequest struct {

	// The type of encryption key used by Amazon Security Lake to encrypt the Security
	// Lake configuration object.
	EncryptionKey *string

	// Replication enables automatic, asynchronous copying of objects across Amazon S3
	// buckets. Amazon S3 buckets that are configured for object replication can be
	// owned by the same Amazon Web Services account or by different accounts. You can
	// replicate objects to a single destination bucket or to multiple destination
	// buckets. The destination buckets can be in different Amazon Web Services Regions
	// or within the same Region as the source bucket. Set up one or more rollup
	// Regions by providing the Region or Regions that should contribute to the central
	// rollup Region.
	ReplicationDestinationRegions []Region

	// Replication settings for the Amazon S3 buckets. This parameter uses the
	// Identity and Access Management (IAM) role you created that is managed by
	// Security Lake, to ensure the replication setting is correct.
	ReplicationRoleArn *string

	// Retention settings for the destination Amazon S3 buckets.
	RetentionSettings []RetentionSetting

	// A tag is a label that you assign to an Amazon Web Services resource. Each tag
	// consists of a key and an optional value, both of which you define.
	TagsMap map[string]string

	noSmithyDocumentSerde
}

// Provides details of Amazon Security Lake lake configuration object.
type LakeConfigurationResponse struct {

	// The type of encryption key used by secure the Security Lake configuration
	// object.
	EncryptionKey *string

	// Replication enables automatic, asynchronous copying of objects across Amazon S3
	// buckets. Amazon S3 buckets that are configured for object replication can be
	// owned by the same Amazon Web Services account or by different accounts. You can
	// replicate objects to a single destination bucket or to multiple destination
	// buckets. The destination buckets can be in different Amazon Web Services Regions
	// or within the same Region as the source bucket. Set up one or more rollup
	// Regions by providing the Region or Regions that should contribute to the central
	// rollup Region.
	ReplicationDestinationRegions []Region

	// Replication settings for the Amazon S3 buckets. This parameter uses the IAM
	// role you created that is managed by Security Lake, to ensure the replication
	// setting is correct.
	ReplicationRoleArn *string

	// Retention settings for the destination Amazon S3 buckets.
	RetentionSettings []RetentionSetting

	// Amazon Resource Names (ARNs) uniquely identify Amazon Web Services resources.
	// Security Lake requires an ARN when you need to specify a resource unambiguously
	// across all of Amazon Web Services, such as in IAM policies, Amazon Relational
	// Database Service (Amazon RDS) tags, and API calls.
	S3BucketArn *string

	// Retrieves the status of the configuration operation for an account in Amazon
	// Security Lake.
	Status SettingsStatus

	// A tag is a label that you assign to an Amazon Web Services resource. Each tag
	// consists of a key and an optional value, both of which you define.
	TagsMap map[string]string

	// The status of the last UpdateDatalake or DeleteDatalake API request.
	UpdateStatus *UpdateStatus

	noSmithyDocumentSerde
}

// The details of the last UpdateDatalake or DeleteDatalake API request which
// failed.
type LastUpdateFailure struct {

	// The reason code for the failure of the last UpdateDatalake or DeleteDatalake
	// API request.
	Code *string

	// The reason for the failure of the last UpdateDatalake or DeleteDatalake API
	// request.
	Reason *string

	noSmithyDocumentSerde
}

// Retrieves the Logs status for the Amazon Security Lake account.
type LogsStatus struct {

	// The health status of services, including error codes and patterns.
	//
	// This member is required.
	HealthStatus SourceStatus

	// Defines path the stored logs are available which has information on your
	// systems, applications, and services.
	//
	// This member is required.
	PathToLogs *string

	noSmithyDocumentSerde
}

// Protocol used in Amazon Security Lake that dictates how notifications are
// posted at the endpoint.
type ProtocolAndNotificationEndpoint struct {

	// The account that is subscribed to receive exception notifications.
	Endpoint *string

	// The protocol to which notification messages are posted.
	Protocol *string

	noSmithyDocumentSerde
}

// Retention settings for the destination Amazon S3 buckets in Amazon Security
// Lake.
type RetentionSetting struct {

	// The retention period specifies a fixed period of time during which the Security
	// Lake object remains locked. You can specify the retention period in days for one
	// or more sources.
	RetentionPeriod *int32

	// The range of storage classes that you can choose from based on the data access,
	// resiliency, and cost requirements of your workloads.
	StorageClass StorageClass

	noSmithyDocumentSerde
}

// The supported source types from which logs and events are collected in Amazon
// Security Lake. For the list of supported Amazon Web Services, see the Amazon
// Security Lake User Guide (https://docs.aws.amazon.com/security-lake/latest/userguide/internal-sources.html)
// .
//
// The following types satisfy this interface:
//
//	SourceTypeMemberAwsSourceType
//	SourceTypeMemberCustomSourceType
type SourceType interface {
	isSourceType()
}

// Amazon Security Lake supports log and event collection for natively supported
// Amazon Web Services.
type SourceTypeMemberAwsSourceType struct {
	Value AwsLogSourceType

	noSmithyDocumentSerde
}

func (*SourceTypeMemberAwsSourceType) isSourceType() {}

// Amazon Security Lake supports custom source types. For a detailed list, see the
// Amazon Security Lake User Guide.
type SourceTypeMemberCustomSourceType struct {
	Value string

	noSmithyDocumentSerde
}

func (*SourceTypeMemberCustomSourceType) isSourceType() {}

// Provides details about the Amazon Security Lake account subscription.
// Subscribers are notified of new objects for a source as the data is written to
// your Amazon S3 bucket for Security Lake.
type SubscriberResource struct {

	// The Amazon Web Services account ID you are using to create your Amazon Security
	// Lake account.
	//
	// This member is required.
	AccountId *string

	// Amazon Security Lake supports log and event collection for natively supported
	// Amazon Web Services. For more information, see the Amazon Security Lake User
	// Guide.
	//
	// This member is required.
	SourceTypes []SourceType

	// The subscription ID of the Amazon Security Lake subscriber account.
	//
	// This member is required.
	SubscriptionId *string

	// You can choose to notify subscribers of new objects with an Amazon Simple Queue
	// Service (Amazon SQS) queue or through messaging to an HTTPS endpoint provided by
	// the subscriber. Subscribers can consume data by directly querying Lake Formation
	// tables in your Amazon S3 bucket through services like Amazon Athena. This
	// subscription type is defined as LAKEFORMATION .
	AccessTypes []AccessType

	// The date and time when the subscription was created.
	CreatedAt *time.Time

	// The external ID of the subscriber. The external ID lets the user that is
	// assuming the role assert the circumstances in which they are operating. It also
	// provides a way for the account owner to permit the role to be assumed only under
	// specific circumstances.
	ExternalId *string

	// The Amazon Resource Name (ARN) which uniquely defines the AWS RAM resource
	// share. Before accepting the RAM resource share invitation, you can view details
	// related to the RAM resource share. This field is available only for Lake
	// Formation subscribers created after March 8, 2023.
	ResourceShareArn *string

	// The name of the resource share.
	ResourceShareName *string

	// The Amazon Resource Name (ARN) specifying the role of the subscriber.
	RoleArn *string

	// The ARN for the Amazon S3 bucket.
	S3BucketArn *string

	// The ARN for the Amazon Simple Notification Service.
	SnsArn *string

	// The subscriber descriptions for a subscriber account. The description for a
	// subscriber includes subscriberName , accountID , externalID , and subscriptionId
	// .
	SubscriberDescription *string

	// The name of your Amazon Security Lake subscriber account.
	SubscriberName *string

	// The subscription endpoint to which exception messages are posted.
	SubscriptionEndpoint *string

	// The subscription protocol to which exception messages are posted.
	SubscriptionProtocol EndpointProtocol

	// The subscription status of the Amazon Security Lake subscriber account.
	SubscriptionStatus SubscriptionStatus

	// The date and time when the subscription was created.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// The status of the last UpdateDatalake or DeleteDatalake API request. This is
// set to Completed after the configuration is updated, or removed if deletion of
// the data lake is successful.
type UpdateStatus struct {

	// The details of the last UpdateDatalake or DeleteDatalake API request which
	// failed.
	LastUpdateFailure *LastUpdateFailure

	// The unique ID for the UpdateDatalake or DeleteDatalake API request.
	LastUpdateRequestId *string

	// The status of the last UpdateDatalake or DeleteDatalake API request that was
	// requested.
	LastUpdateStatus SettingsStatus

	noSmithyDocumentSerde
}

// The input fails to meet the constraints specified in Amazon Security Lake.
type ValidationExceptionField struct {

	// Describes the error encountered.
	//
	// This member is required.
	Message *string

	// Name of the validation exception.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isSourceType() {}
