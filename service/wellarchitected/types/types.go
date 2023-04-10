// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The choice level additional resources for a custom lens. This field does not
// apply to Amazon Web Services official lenses.
type AdditionalResources struct {

	// The URLs for additional resources, either helpful resources or improvement
	// plans, for a custom lens. Up to five additional URLs can be specified.
	Content []ChoiceContent

	// Type of additional resource for a custom lens.
	Type AdditionalResourceType

	noSmithyDocumentSerde
}

// An answer of the question.
type Answer struct {

	// A list of selected choices to a question in your workload.
	ChoiceAnswers []ChoiceAnswer

	// List of choices available for a question.
	Choices []Choice

	// The helpful resource text to be displayed for a custom lens. This field does
	// not apply to Amazon Web Services official lenses.
	HelpfulResourceDisplayText *string

	// The helpful resource URL. For Amazon Web Services official lenses, this is the
	// helpful resource URL for a question or choice. For custom lenses, this is the
	// helpful resource URL for a question and is only provided if
	// HelpfulResourceDisplayText was specified for the question.
	HelpfulResourceUrl *string

	// The improvement plan URL for a question in an Amazon Web Services official
	// lenses. This value is only available if the question has been answered. This
	// value does not apply to custom lenses.
	ImprovementPlanUrl *string

	// Defines whether this question is applicable to a lens review.
	IsApplicable bool

	// The notes associated with the workload.
	Notes *string

	// The ID used to identify a pillar, for example, security . A pillar is identified
	// by its PillarReviewSummary$PillarId .
	PillarId *string

	// The description of the question.
	QuestionDescription *string

	// The ID of the question.
	QuestionId *string

	// The title of the question.
	QuestionTitle *string

	// The reason why the question is not applicable to your workload.
	Reason AnswerReason

	// The risk for a given workload, lens review, pillar, or question.
	Risk Risk

	// List of selected choice IDs in a question answer. The values entered replace
	// the previously selected choices.
	SelectedChoices []string

	noSmithyDocumentSerde
}

// An answer summary of a lens review in a workload.
type AnswerSummary struct {

	// A list of selected choices to a question in your workload.
	ChoiceAnswerSummaries []ChoiceAnswerSummary

	// List of choices available for a question.
	Choices []Choice

	// Defines whether this question is applicable to a lens review.
	IsApplicable bool

	// The ID used to identify a pillar, for example, security . A pillar is identified
	// by its PillarReviewSummary$PillarId .
	PillarId *string

	// The ID of the question.
	QuestionId *string

	// The title of the question.
	QuestionTitle *string

	// The reason why a choice is non-applicable to a question in your workload.
	Reason AnswerReason

	// The risk for a given workload, lens review, pillar, or question.
	Risk Risk

	// List of selected choice IDs in a question answer. The values entered replace
	// the previously selected choices.
	SelectedChoices []string

	noSmithyDocumentSerde
}

// A best practice, or question choice, that has been identified as a risk in this
// question.
type BestPractice struct {

	// The ID of a choice.
	ChoiceId *string

	// The title of a choice.
	ChoiceTitle *string

	noSmithyDocumentSerde
}

// Account details for a Well-Architected best practice in relation to Trusted
// Advisor checks.
type CheckDetail struct {

	// An Amazon Web Services account ID.
	AccountId *string

	// The ID of a choice.
	ChoiceId *string

	// Trusted Advisor check description.
	Description *string

	// Count of flagged resources associated to the check.
	FlaggedResources int32

	// Trusted Advisor check ID.
	Id *string

	// Well-Architected Lens ARN associated to the check.
	LensArn *string

	// Trusted Advisor check name.
	Name *string

	// The ID used to identify a pillar, for example, security . A pillar is identified
	// by its PillarReviewSummary$PillarId .
	PillarId *string

	// Provider of the check related to the best practice.
	Provider CheckProvider

	// The ID of the question.
	QuestionId *string

	// Reason associated to the check.
	Reason CheckFailureReason

	// Status associated to the check.
	Status CheckStatus

	// The date and time recorded.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// Trusted Advisor check summary.
type CheckSummary struct {

	// Account summary associated to the check.
	AccountSummary map[string]int32

	// The ID of a choice.
	ChoiceId *string

	// Trusted Advisor check description.
	Description *string

	// Trusted Advisor check ID.
	Id *string

	// Well-Architected Lens ARN associated to the check.
	LensArn *string

	// Trusted Advisor check name.
	Name *string

	// The ID used to identify a pillar, for example, security . A pillar is identified
	// by its PillarReviewSummary$PillarId .
	PillarId *string

	// Provider of the check related to the best practice.
	Provider CheckProvider

	// The ID of the question.
	QuestionId *string

	// Status associated to the check.
	Status CheckStatus

	// The date and time recorded.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// A choice available to answer question.
type Choice struct {

	// The additional resources for a choice in a custom lens. A choice can have up to
	// two additional resources: one of type HELPFUL_RESOURCE , one of type
	// IMPROVEMENT_PLAN , or both.
	AdditionalResources []AdditionalResources

	// The ID of a choice.
	ChoiceId *string

	// The description of a choice.
	Description *string

	// The helpful resource (both text and URL) for a particular choice. This field
	// only applies to custom lenses. Each choice can have only one helpful resource.
	HelpfulResource *ChoiceContent

	// The improvement plan (both text and URL) for a particular choice. This field
	// only applies to custom lenses. Each choice can have only one improvement plan.
	ImprovementPlan *ChoiceContent

	// The title of a choice.
	Title *string

	noSmithyDocumentSerde
}

// A choice that has been answered on a question in your workload.
type ChoiceAnswer struct {

	// The ID of a choice.
	ChoiceId *string

	// The notes associated with a choice.
	Notes *string

	// The reason why a choice is non-applicable to a question in your workload.
	Reason ChoiceReason

	// The status of a choice.
	Status ChoiceStatus

	noSmithyDocumentSerde
}

// A choice summary that has been answered on a question in your workload.
type ChoiceAnswerSummary struct {

	// The ID of a choice.
	ChoiceId *string

	// The reason why a choice is non-applicable to a question in your workload.
	Reason ChoiceReason

	// The status of a choice.
	Status ChoiceStatus

	noSmithyDocumentSerde
}

// The choice content.
type ChoiceContent struct {

	// The display text for the choice content.
	DisplayText *string

	// The URL for the choice content.
	Url *string

	noSmithyDocumentSerde
}

// The choice level improvement plan.
type ChoiceImprovementPlan struct {

	// The ID of a choice.
	ChoiceId *string

	// The display text for the improvement plan.
	DisplayText *string

	// The improvement plan URL for a question in an Amazon Web Services official
	// lenses. This value is only available if the question has been answered. This
	// value does not apply to custom lenses.
	ImprovementPlanUrl *string

	noSmithyDocumentSerde
}

// A list of choices to be updated.
type ChoiceUpdate struct {

	// The status of a choice.
	//
	// This member is required.
	Status ChoiceStatus

	// The notes associated with a choice.
	Notes *string

	// The reason why a choice is non-applicable to a question in your workload.
	Reason ChoiceReason

	noSmithyDocumentSerde
}

// A metric that contributes to the consolidated report.
type ConsolidatedReportMetric struct {

	// The metrics for the lenses in the workload.
	Lenses []LensMetric

	// The total number of lenses applied to the workload.
	LensesAppliedCount int32

	// The metric type of a metric in the consolidated report. Currently only WORKLOAD
	// metric types are supported.
	MetricType MetricType

	// A map from risk names to the count of how many questions have that rating.
	RiskCounts map[string]int32

	// The date and time recorded.
	UpdatedAt *time.Time

	// The ARN for the workload.
	WorkloadArn *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	WorkloadId *string

	// The name of the workload. The name must be unique within an account within an
	// Amazon Web Services Region. Spaces and capitalization are ignored when checking
	// for uniqueness.
	WorkloadName *string

	noSmithyDocumentSerde
}

// An improvement summary of a lens review in a workload.
type ImprovementSummary struct {

	// The improvement plan URL for a question in an Amazon Web Services official
	// lenses. This value is only available if the question has been answered. This
	// value does not apply to custom lenses.
	ImprovementPlanUrl *string

	// The improvement plan details.
	ImprovementPlans []ChoiceImprovementPlan

	// The ID used to identify a pillar, for example, security . A pillar is identified
	// by its PillarReviewSummary$PillarId .
	PillarId *string

	// The ID of the question.
	QuestionId *string

	// The title of the question.
	QuestionTitle *string

	// The risk for a given workload, lens review, pillar, or question.
	Risk Risk

	noSmithyDocumentSerde
}

// A lens return object.
type Lens struct {

	// The description of the lens.
	Description *string

	// The ARN of a lens.
	LensArn *string

	// The version of a lens.
	LensVersion *string

	// The full name of the lens.
	Name *string

	// The Amazon Web Services account ID that owns the lens.
	Owner *string

	// The ID assigned to the share invitation.
	ShareInvitationId *string

	// The tags assigned to the lens.
	Tags map[string]string

	noSmithyDocumentSerde
}

// A metric for a particular lens in a workload.
type LensMetric struct {

	// The lens ARN.
	LensArn *string

	// The metrics for the pillars in a lens.
	Pillars []PillarMetric

	// A map from risk names to the count of how many questions have that rating.
	RiskCounts map[string]int32

	noSmithyDocumentSerde
}

// A lens review of a question.
type LensReview struct {

	// The alias of the lens. For Amazon Web Services official lenses, this is either
	// the lens alias, such as serverless , or the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1::lens/serverless . Note that some operations
	// (such as ExportLens and CreateLensShare) are not permitted on Amazon Web
	// Services official lenses. For custom lenses, this is the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2:123456789012:lens/0123456789abcdef01234567890abcdef
	// . Each lens is identified by its LensSummary$LensAlias .
	LensAlias *string

	// The ARN for the lens.
	LensArn *string

	// The full name of the lens.
	LensName *string

	// The status of the lens.
	LensStatus LensStatus

	// The version of the lens.
	LensVersion *string

	// The token to use to retrieve the next set of results.
	NextToken *string

	// The notes associated with the workload.
	Notes *string

	// List of pillar review summaries of lens review in a workload.
	PillarReviewSummaries []PillarReviewSummary

	// A map from risk names to the count of how many questions have that rating.
	RiskCounts map[string]int32

	// The date and time recorded.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// A report of a lens review.
type LensReviewReport struct {

	// The Base64-encoded string representation of a lens review report. This data can
	// be used to create a PDF file. Only returned by GetConsolidatedReport when PDF
	// format is requested.
	Base64String *string

	// The alias of the lens. For Amazon Web Services official lenses, this is either
	// the lens alias, such as serverless , or the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1::lens/serverless . Note that some operations
	// (such as ExportLens and CreateLensShare) are not permitted on Amazon Web
	// Services official lenses. For custom lenses, this is the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2:123456789012:lens/0123456789abcdef01234567890abcdef
	// . Each lens is identified by its LensSummary$LensAlias .
	LensAlias *string

	// The ARN for the lens.
	LensArn *string

	noSmithyDocumentSerde
}

// A lens review summary of a workload.
type LensReviewSummary struct {

	// The alias of the lens. For Amazon Web Services official lenses, this is either
	// the lens alias, such as serverless , or the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1::lens/serverless . Note that some operations
	// (such as ExportLens and CreateLensShare) are not permitted on Amazon Web
	// Services official lenses. For custom lenses, this is the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2:123456789012:lens/0123456789abcdef01234567890abcdef
	// . Each lens is identified by its LensSummary$LensAlias .
	LensAlias *string

	// The ARN for the lens.
	LensArn *string

	// The full name of the lens.
	LensName *string

	// The status of the lens.
	LensStatus LensStatus

	// The version of the lens.
	LensVersion *string

	// A map from risk names to the count of how many questions have that rating.
	RiskCounts map[string]int32

	// The date and time recorded.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// A lens share summary return object.
type LensShareSummary struct {

	// The ID associated with the workload share.
	ShareId *string

	// The Amazon Web Services account ID, IAM role, organization ID, or
	// organizational unit (OU) ID with which the workload is shared.
	SharedWith *string

	// The status of a workload share.
	Status ShareStatus

	// Optional message to compliment the Status field.
	StatusMessage *string

	noSmithyDocumentSerde
}

// A lens summary of a lens.
type LensSummary struct {

	// The date and time recorded.
	CreatedAt *time.Time

	// The description of the lens.
	Description *string

	// The alias of the lens. For Amazon Web Services official lenses, this is either
	// the lens alias, such as serverless , or the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1::lens/serverless . Note that some operations
	// (such as ExportLens and CreateLensShare) are not permitted on Amazon Web
	// Services official lenses. For custom lenses, this is the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2:123456789012:lens/0123456789abcdef01234567890abcdef
	// . Each lens is identified by its LensSummary$LensAlias .
	LensAlias *string

	// The ARN of the lens.
	LensArn *string

	// The full name of the lens.
	LensName *string

	// The status of the lens.
	LensStatus LensStatus

	// The type of the lens.
	LensType LensType

	// The version of the lens.
	LensVersion *string

	// An Amazon Web Services account ID.
	Owner *string

	// The date and time recorded.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// Lens upgrade summary return object.
type LensUpgradeSummary struct {

	// The current version of the lens.
	CurrentLensVersion *string

	// The latest version of the lens.
	LatestLensVersion *string

	// The alias of the lens. For Amazon Web Services official lenses, this is either
	// the lens alias, such as serverless , or the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1::lens/serverless . Note that some operations
	// (such as ExportLens and CreateLensShare) are not permitted on Amazon Web
	// Services official lenses. For custom lenses, this is the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2:123456789012:lens/0123456789abcdef01234567890abcdef
	// . Each lens is identified by its LensSummary$LensAlias .
	LensAlias *string

	// The ARN for the lens.
	LensArn *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	WorkloadId *string

	// The name of the workload. The name must be unique within an account within an
	// Amazon Web Services Region. Spaces and capitalization are ignored when checking
	// for uniqueness.
	WorkloadName *string

	noSmithyDocumentSerde
}

// A milestone return object.
type Milestone struct {

	// The name of the milestone in a workload. Milestone names must be unique within
	// a workload.
	MilestoneName *string

	// The milestone number. A workload can have a maximum of 100 milestones.
	MilestoneNumber int32

	// The date and time recorded.
	RecordedAt *time.Time

	// A workload return object.
	Workload *Workload

	noSmithyDocumentSerde
}

// A milestone summary return object.
type MilestoneSummary struct {

	// The name of the milestone in a workload. Milestone names must be unique within
	// a workload.
	MilestoneName *string

	// The milestone number. A workload can have a maximum of 100 milestones.
	MilestoneNumber int32

	// The date and time recorded.
	RecordedAt *time.Time

	// A workload summary return object.
	WorkloadSummary *WorkloadSummary

	noSmithyDocumentSerde
}

// A notification summary return object.
type NotificationSummary struct {

	// Summary of lens upgrade.
	LensUpgradeSummary *LensUpgradeSummary

	// The type of notification.
	Type NotificationType

	noSmithyDocumentSerde
}

// A pillar difference return object.
type PillarDifference struct {

	// Indicates the type of change to the pillar.
	DifferenceStatus DifferenceStatus

	// The ID used to identify a pillar, for example, security . A pillar is identified
	// by its PillarReviewSummary$PillarId .
	PillarId *string

	// The name of the pillar.
	PillarName *string

	// List of question differences.
	QuestionDifferences []QuestionDifference

	noSmithyDocumentSerde
}

// A metric for a particular pillar in a lens.
type PillarMetric struct {

	// The ID used to identify a pillar, for example, security . A pillar is identified
	// by its PillarReviewSummary$PillarId .
	PillarId *string

	// The questions that have been identified as risks in the pillar.
	Questions []QuestionMetric

	// A map from risk names to the count of how many questions have that rating.
	RiskCounts map[string]int32

	noSmithyDocumentSerde
}

// A pillar review summary of a lens review.
type PillarReviewSummary struct {

	// The notes associated with the workload.
	Notes *string

	// The ID used to identify a pillar, for example, security . A pillar is identified
	// by its PillarReviewSummary$PillarId .
	PillarId *string

	// The name of the pillar.
	PillarName *string

	// A map from risk names to the count of how many questions have that rating.
	RiskCounts map[string]int32

	noSmithyDocumentSerde
}

// A question difference return object.
type QuestionDifference struct {

	// Indicates the type of change to the question.
	DifferenceStatus DifferenceStatus

	// The ID of the question.
	QuestionId *string

	// The title of the question.
	QuestionTitle *string

	noSmithyDocumentSerde
}

// A metric for a particular question in the pillar.
type QuestionMetric struct {

	// The best practices, or choices, that have been identified as contributing to
	// risk in a question.
	BestPractices []BestPractice

	// The ID of the question.
	QuestionId *string

	// The risk for a given workload, lens review, pillar, or question.
	Risk Risk

	noSmithyDocumentSerde
}

// The share invitation.
type ShareInvitation struct {

	// The alias of the lens. For Amazon Web Services official lenses, this is either
	// the lens alias, such as serverless , or the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1::lens/serverless . Note that some operations
	// (such as ExportLens and CreateLensShare) are not permitted on Amazon Web
	// Services official lenses. For custom lenses, this is the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2:123456789012:lens/0123456789abcdef01234567890abcdef
	// . Each lens is identified by its LensSummary$LensAlias .
	LensAlias *string

	// The ARN for the lens.
	LensArn *string

	// The ID assigned to the share invitation.
	ShareInvitationId *string

	// The resource type of the share invitation.
	ShareResourceType ShareResourceType

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	WorkloadId *string

	noSmithyDocumentSerde
}

// A share invitation summary return object.
type ShareInvitationSummary struct {

	// The ARN for the lens.
	LensArn *string

	// The full name of the lens.
	LensName *string

	// Permission granted on a workload share.
	PermissionType PermissionType

	// The ID assigned to the share invitation.
	ShareInvitationId *string

	// The resource type of the share invitation.
	ShareResourceType ShareResourceType

	// An Amazon Web Services account ID.
	SharedBy *string

	// The Amazon Web Services account ID, IAM role, organization ID, or
	// organizational unit (OU) ID with which the workload is shared.
	SharedWith *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	WorkloadId *string

	// The name of the workload. The name must be unique within an account within an
	// Amazon Web Services Region. Spaces and capitalization are ignored when checking
	// for uniqueness.
	WorkloadName *string

	noSmithyDocumentSerde
}

// Stores information about a field passed inside a request that resulted in an
// exception.
type ValidationExceptionField struct {

	// Description of the error.
	//
	// This member is required.
	Message *string

	// The field name for which validation failed.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// The differences between the base and latest versions of the lens.
type VersionDifferences struct {

	// The differences between the base and latest versions of the lens.
	PillarDifferences []PillarDifference

	noSmithyDocumentSerde
}

// A workload return object.
type Workload struct {

	// The list of Amazon Web Services account IDs associated with the workload.
	AccountIds []string

	// List of AppRegistry application ARNs associated to the workload.
	Applications []string

	// The URL of the architectural design for the workload.
	ArchitecturalDesign *string

	// The list of Amazon Web Services Regions associated with the workload, for
	// example, us-east-2 , or ca-central-1 .
	AwsRegions []string

	// The description for the workload.
	Description *string

	// Discovery configuration associated to the workload.
	DiscoveryConfig *WorkloadDiscoveryConfig

	// The environment for the workload.
	Environment WorkloadEnvironment

	// The improvement status for a workload.
	ImprovementStatus WorkloadImprovementStatus

	// The industry for the workload.
	Industry *string

	// The industry type for the workload. If specified, must be one of the following:
	//   - Agriculture
	//   - Automobile
	//   - Defense
	//   - Design and Engineering
	//   - Digital Advertising
	//   - Education
	//   - Environmental Protection
	//   - Financial Services
	//   - Gaming
	//   - General Public Services
	//   - Healthcare
	//   - Hospitality
	//   - InfoTech
	//   - Justice and Public Safety
	//   - Life Sciences
	//   - Manufacturing
	//   - Media & Entertainment
	//   - Mining & Resources
	//   - Oil & Gas
	//   - Power & Utilities
	//   - Professional Services
	//   - Real Estate & Construction
	//   - Retail & Wholesale
	//   - Social Protection
	//   - Telecommunications
	//   - Travel, Transportation & Logistics
	//   - Other
	IndustryType *string

	// Flag indicating whether the workload owner has acknowledged that the Review
	// owner field is required. If a Review owner is not added to the workload within
	// 60 days of acknowledgement, access to the workload is restricted until an owner
	// is added.
	IsReviewOwnerUpdateAcknowledged bool

	// The list of lenses associated with the workload. Each lens is identified by its
	// LensSummary$LensAlias .
	Lenses []string

	// The list of non-Amazon Web Services Regions associated with the workload.
	NonAwsRegions []string

	// The notes associated with the workload.
	Notes *string

	// An Amazon Web Services account ID.
	Owner *string

	// The priorities of the pillars, which are used to order items in the improvement
	// plan. Each pillar is represented by its PillarReviewSummary$PillarId .
	PillarPriorities []string

	// The review owner of the workload. The name, email address, or identifier for
	// the primary group or individual that owns the workload review process.
	ReviewOwner *string

	// The date and time recorded.
	ReviewRestrictionDate *time.Time

	// A map from risk names to the count of how many questions have that rating.
	RiskCounts map[string]int32

	// The ID assigned to the share invitation.
	ShareInvitationId *string

	// The tags associated with the workload.
	Tags map[string]string

	// The date and time recorded.
	UpdatedAt *time.Time

	// The ARN for the workload.
	WorkloadArn *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	WorkloadId *string

	// The name of the workload. The name must be unique within an account within an
	// Amazon Web Services Region. Spaces and capitalization are ignored when checking
	// for uniqueness.
	WorkloadName *string

	noSmithyDocumentSerde
}

// Discovery configuration associated to the workload.
type WorkloadDiscoveryConfig struct {

	// Discovery integration status in respect to Trusted Advisor for the workload.
	TrustedAdvisorIntegrationStatus TrustedAdvisorIntegrationStatus

	noSmithyDocumentSerde
}

// A workload share return object.
type WorkloadShare struct {

	// Permission granted on a workload share.
	PermissionType PermissionType

	// The ID associated with the workload share.
	ShareId *string

	// An Amazon Web Services account ID.
	SharedBy *string

	// The Amazon Web Services account ID, IAM role, organization ID, or
	// organizational unit (OU) ID with which the workload is shared.
	SharedWith *string

	// The status of a workload share.
	Status ShareStatus

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	WorkloadId *string

	// The name of the workload. The name must be unique within an account within an
	// Amazon Web Services Region. Spaces and capitalization are ignored when checking
	// for uniqueness.
	WorkloadName *string

	noSmithyDocumentSerde
}

// A workload share summary return object.
type WorkloadShareSummary struct {

	// Permission granted on a workload share.
	PermissionType PermissionType

	// The ID associated with the workload share.
	ShareId *string

	// The Amazon Web Services account ID, IAM role, organization ID, or
	// organizational unit (OU) ID with which the workload is shared.
	SharedWith *string

	// The status of a workload share.
	Status ShareStatus

	// Optional message to compliment the Status field.
	StatusMessage *string

	noSmithyDocumentSerde
}

// A workload summary return object.
type WorkloadSummary struct {

	// The improvement status for a workload.
	ImprovementStatus WorkloadImprovementStatus

	// The list of lenses associated with the workload. Each lens is identified by its
	// LensSummary$LensAlias .
	Lenses []string

	// An Amazon Web Services account ID.
	Owner *string

	// A map from risk names to the count of how many questions have that rating.
	RiskCounts map[string]int32

	// The date and time recorded.
	UpdatedAt *time.Time

	// The ARN for the workload.
	WorkloadArn *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	WorkloadId *string

	// The name of the workload. The name must be unique within an account within an
	// Amazon Web Services Region. Spaces and capitalization are ignored when checking
	// for uniqueness.
	WorkloadName *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
