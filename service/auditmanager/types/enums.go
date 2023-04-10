// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccountStatus string

// Enum values for AccountStatus
const (
	AccountStatusActive            AccountStatus = "ACTIVE"
	AccountStatusInactive          AccountStatus = "INACTIVE"
	AccountStatusPendingActivation AccountStatus = "PENDING_ACTIVATION"
)

// Values returns all known values for AccountStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AccountStatus) Values() []AccountStatus {
	return []AccountStatus{
		"ACTIVE",
		"INACTIVE",
		"PENDING_ACTIVATION",
	}
}

type ActionEnum string

// Enum values for ActionEnum
const (
	ActionEnumCreate         ActionEnum = "CREATE"
	ActionEnumUpdateMetadata ActionEnum = "UPDATE_METADATA"
	ActionEnumActive         ActionEnum = "ACTIVE"
	ActionEnumInactive       ActionEnum = "INACTIVE"
	ActionEnumDelete         ActionEnum = "DELETE"
	ActionEnumUnderReview    ActionEnum = "UNDER_REVIEW"
	ActionEnumReviewed       ActionEnum = "REVIEWED"
	ActionEnumImportEvidence ActionEnum = "IMPORT_EVIDENCE"
)

// Values returns all known values for ActionEnum. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ActionEnum) Values() []ActionEnum {
	return []ActionEnum{
		"CREATE",
		"UPDATE_METADATA",
		"ACTIVE",
		"INACTIVE",
		"DELETE",
		"UNDER_REVIEW",
		"REVIEWED",
		"IMPORT_EVIDENCE",
	}
}

type AssessmentReportDestinationType string

// Enum values for AssessmentReportDestinationType
const (
	AssessmentReportDestinationTypeS3 AssessmentReportDestinationType = "S3"
)

// Values returns all known values for AssessmentReportDestinationType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (AssessmentReportDestinationType) Values() []AssessmentReportDestinationType {
	return []AssessmentReportDestinationType{
		"S3",
	}
}

type AssessmentReportStatus string

// Enum values for AssessmentReportStatus
const (
	AssessmentReportStatusComplete   AssessmentReportStatus = "COMPLETE"
	AssessmentReportStatusInProgress AssessmentReportStatus = "IN_PROGRESS"
	AssessmentReportStatusFailed     AssessmentReportStatus = "FAILED"
)

// Values returns all known values for AssessmentReportStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssessmentReportStatus) Values() []AssessmentReportStatus {
	return []AssessmentReportStatus{
		"COMPLETE",
		"IN_PROGRESS",
		"FAILED",
	}
}

type AssessmentStatus string

// Enum values for AssessmentStatus
const (
	AssessmentStatusActive   AssessmentStatus = "ACTIVE"
	AssessmentStatusInactive AssessmentStatus = "INACTIVE"
)

// Values returns all known values for AssessmentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssessmentStatus) Values() []AssessmentStatus {
	return []AssessmentStatus{
		"ACTIVE",
		"INACTIVE",
	}
}

type ControlResponse string

// Enum values for ControlResponse
const (
	ControlResponseManual   ControlResponse = "MANUAL"
	ControlResponseAutomate ControlResponse = "AUTOMATE"
	ControlResponseDefer    ControlResponse = "DEFER"
	ControlResponseIgnore   ControlResponse = "IGNORE"
)

// Values returns all known values for ControlResponse. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ControlResponse) Values() []ControlResponse {
	return []ControlResponse{
		"MANUAL",
		"AUTOMATE",
		"DEFER",
		"IGNORE",
	}
}

type ControlSetStatus string

// Enum values for ControlSetStatus
const (
	ControlSetStatusActive      ControlSetStatus = "ACTIVE"
	ControlSetStatusUnderReview ControlSetStatus = "UNDER_REVIEW"
	ControlSetStatusReviewed    ControlSetStatus = "REVIEWED"
)

// Values returns all known values for ControlSetStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ControlSetStatus) Values() []ControlSetStatus {
	return []ControlSetStatus{
		"ACTIVE",
		"UNDER_REVIEW",
		"REVIEWED",
	}
}

type ControlStatus string

// Enum values for ControlStatus
const (
	ControlStatusUnderReview ControlStatus = "UNDER_REVIEW"
	ControlStatusReviewed    ControlStatus = "REVIEWED"
	ControlStatusInactive    ControlStatus = "INACTIVE"
)

// Values returns all known values for ControlStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ControlStatus) Values() []ControlStatus {
	return []ControlStatus{
		"UNDER_REVIEW",
		"REVIEWED",
		"INACTIVE",
	}
}

type ControlType string

// Enum values for ControlType
const (
	ControlTypeStandard ControlType = "Standard"
	ControlTypeCustom   ControlType = "Custom"
)

// Values returns all known values for ControlType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ControlType) Values() []ControlType {
	return []ControlType{
		"Standard",
		"Custom",
	}
}

type DelegationStatus string

// Enum values for DelegationStatus
const (
	DelegationStatusInProgress  DelegationStatus = "IN_PROGRESS"
	DelegationStatusUnderReview DelegationStatus = "UNDER_REVIEW"
	DelegationStatusComplete    DelegationStatus = "COMPLETE"
)

// Values returns all known values for DelegationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DelegationStatus) Values() []DelegationStatus {
	return []DelegationStatus{
		"IN_PROGRESS",
		"UNDER_REVIEW",
		"COMPLETE",
	}
}

type DeleteResources string

// Enum values for DeleteResources
const (
	DeleteResourcesAll     DeleteResources = "ALL"
	DeleteResourcesDefault DeleteResources = "DEFAULT"
)

// Values returns all known values for DeleteResources. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeleteResources) Values() []DeleteResources {
	return []DeleteResources{
		"ALL",
		"DEFAULT",
	}
}

type EvidenceFinderBackfillStatus string

// Enum values for EvidenceFinderBackfillStatus
const (
	EvidenceFinderBackfillStatusNotStarted EvidenceFinderBackfillStatus = "NOT_STARTED"
	EvidenceFinderBackfillStatusInProgress EvidenceFinderBackfillStatus = "IN_PROGRESS"
	EvidenceFinderBackfillStatusCompleted  EvidenceFinderBackfillStatus = "COMPLETED"
)

// Values returns all known values for EvidenceFinderBackfillStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (EvidenceFinderBackfillStatus) Values() []EvidenceFinderBackfillStatus {
	return []EvidenceFinderBackfillStatus{
		"NOT_STARTED",
		"IN_PROGRESS",
		"COMPLETED",
	}
}

type EvidenceFinderEnablementStatus string

// Enum values for EvidenceFinderEnablementStatus
const (
	EvidenceFinderEnablementStatusEnabled           EvidenceFinderEnablementStatus = "ENABLED"
	EvidenceFinderEnablementStatusDisabled          EvidenceFinderEnablementStatus = "DISABLED"
	EvidenceFinderEnablementStatusEnableInProgress  EvidenceFinderEnablementStatus = "ENABLE_IN_PROGRESS"
	EvidenceFinderEnablementStatusDisableInProgress EvidenceFinderEnablementStatus = "DISABLE_IN_PROGRESS"
)

// Values returns all known values for EvidenceFinderEnablementStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (EvidenceFinderEnablementStatus) Values() []EvidenceFinderEnablementStatus {
	return []EvidenceFinderEnablementStatus{
		"ENABLED",
		"DISABLED",
		"ENABLE_IN_PROGRESS",
		"DISABLE_IN_PROGRESS",
	}
}

type FrameworkType string

// Enum values for FrameworkType
const (
	FrameworkTypeStandard FrameworkType = "Standard"
	FrameworkTypeCustom   FrameworkType = "Custom"
)

// Values returns all known values for FrameworkType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FrameworkType) Values() []FrameworkType {
	return []FrameworkType{
		"Standard",
		"Custom",
	}
}

type KeywordInputType string

// Enum values for KeywordInputType
const (
	KeywordInputTypeSelectFromList KeywordInputType = "SELECT_FROM_LIST"
)

// Values returns all known values for KeywordInputType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (KeywordInputType) Values() []KeywordInputType {
	return []KeywordInputType{
		"SELECT_FROM_LIST",
	}
}

type ObjectTypeEnum string

// Enum values for ObjectTypeEnum
const (
	ObjectTypeEnumAssessment       ObjectTypeEnum = "ASSESSMENT"
	ObjectTypeEnumControlSet       ObjectTypeEnum = "CONTROL_SET"
	ObjectTypeEnumControl          ObjectTypeEnum = "CONTROL"
	ObjectTypeEnumDelegation       ObjectTypeEnum = "DELEGATION"
	ObjectTypeEnumAssessmentReport ObjectTypeEnum = "ASSESSMENT_REPORT"
)

// Values returns all known values for ObjectTypeEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ObjectTypeEnum) Values() []ObjectTypeEnum {
	return []ObjectTypeEnum{
		"ASSESSMENT",
		"CONTROL_SET",
		"CONTROL",
		"DELEGATION",
		"ASSESSMENT_REPORT",
	}
}

type RoleType string

// Enum values for RoleType
const (
	RoleTypeProcessOwner  RoleType = "PROCESS_OWNER"
	RoleTypeResourceOwner RoleType = "RESOURCE_OWNER"
)

// Values returns all known values for RoleType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (RoleType) Values() []RoleType {
	return []RoleType{
		"PROCESS_OWNER",
		"RESOURCE_OWNER",
	}
}

type SettingAttribute string

// Enum values for SettingAttribute
const (
	SettingAttributeAll                                 SettingAttribute = "ALL"
	SettingAttributeIsAwsOrgEnabled                     SettingAttribute = "IS_AWS_ORG_ENABLED"
	SettingAttributeSnsTopic                            SettingAttribute = "SNS_TOPIC"
	SettingAttributeDefaultAssessmentReportsDestination SettingAttribute = "DEFAULT_ASSESSMENT_REPORTS_DESTINATION"
	SettingAttributeDefaultProcessOwners                SettingAttribute = "DEFAULT_PROCESS_OWNERS"
	SettingAttributeEvidenceFinderEnablement            SettingAttribute = "EVIDENCE_FINDER_ENABLEMENT"
	SettingAttributeDeregistrationPolicy                SettingAttribute = "DEREGISTRATION_POLICY"
)

// Values returns all known values for SettingAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SettingAttribute) Values() []SettingAttribute {
	return []SettingAttribute{
		"ALL",
		"IS_AWS_ORG_ENABLED",
		"SNS_TOPIC",
		"DEFAULT_ASSESSMENT_REPORTS_DESTINATION",
		"DEFAULT_PROCESS_OWNERS",
		"EVIDENCE_FINDER_ENABLEMENT",
		"DEREGISTRATION_POLICY",
	}
}

type ShareRequestAction string

// Enum values for ShareRequestAction
const (
	ShareRequestActionAccept  ShareRequestAction = "ACCEPT"
	ShareRequestActionDecline ShareRequestAction = "DECLINE"
	ShareRequestActionRevoke  ShareRequestAction = "REVOKE"
)

// Values returns all known values for ShareRequestAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ShareRequestAction) Values() []ShareRequestAction {
	return []ShareRequestAction{
		"ACCEPT",
		"DECLINE",
		"REVOKE",
	}
}

type ShareRequestStatus string

// Enum values for ShareRequestStatus
const (
	ShareRequestStatusActive      ShareRequestStatus = "ACTIVE"
	ShareRequestStatusReplicating ShareRequestStatus = "REPLICATING"
	ShareRequestStatusShared      ShareRequestStatus = "SHARED"
	ShareRequestStatusExpiring    ShareRequestStatus = "EXPIRING"
	ShareRequestStatusFailed      ShareRequestStatus = "FAILED"
	ShareRequestStatusExpired     ShareRequestStatus = "EXPIRED"
	ShareRequestStatusDeclined    ShareRequestStatus = "DECLINED"
	ShareRequestStatusRevoked     ShareRequestStatus = "REVOKED"
)

// Values returns all known values for ShareRequestStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ShareRequestStatus) Values() []ShareRequestStatus {
	return []ShareRequestStatus{
		"ACTIVE",
		"REPLICATING",
		"SHARED",
		"EXPIRING",
		"FAILED",
		"EXPIRED",
		"DECLINED",
		"REVOKED",
	}
}

type ShareRequestType string

// Enum values for ShareRequestType
const (
	ShareRequestTypeSent     ShareRequestType = "SENT"
	ShareRequestTypeReceived ShareRequestType = "RECEIVED"
)

// Values returns all known values for ShareRequestType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ShareRequestType) Values() []ShareRequestType {
	return []ShareRequestType{
		"SENT",
		"RECEIVED",
	}
}

type SourceFrequency string

// Enum values for SourceFrequency
const (
	SourceFrequencyDaily   SourceFrequency = "DAILY"
	SourceFrequencyWeekly  SourceFrequency = "WEEKLY"
	SourceFrequencyMonthly SourceFrequency = "MONTHLY"
)

// Values returns all known values for SourceFrequency. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SourceFrequency) Values() []SourceFrequency {
	return []SourceFrequency{
		"DAILY",
		"WEEKLY",
		"MONTHLY",
	}
}

type SourceSetUpOption string

// Enum values for SourceSetUpOption
const (
	SourceSetUpOptionSystemControlsMapping     SourceSetUpOption = "System_Controls_Mapping"
	SourceSetUpOptionProceduralControlsMapping SourceSetUpOption = "Procedural_Controls_Mapping"
)

// Values returns all known values for SourceSetUpOption. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SourceSetUpOption) Values() []SourceSetUpOption {
	return []SourceSetUpOption{
		"System_Controls_Mapping",
		"Procedural_Controls_Mapping",
	}
}

type SourceType string

// Enum values for SourceType
const (
	SourceTypeAwsCloudtrail  SourceType = "AWS_Cloudtrail"
	SourceTypeAwsConfig      SourceType = "AWS_Config"
	SourceTypeAwsSecurityHub SourceType = "AWS_Security_Hub"
	SourceTypeAwsApiCall     SourceType = "AWS_API_Call"
	SourceTypeManual         SourceType = "MANUAL"
)

// Values returns all known values for SourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SourceType) Values() []SourceType {
	return []SourceType{
		"AWS_Cloudtrail",
		"AWS_Config",
		"AWS_Security_Hub",
		"AWS_API_Call",
		"MANUAL",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "unknownOperation"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "cannotParse"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "fieldValidationFailed"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "other"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"unknownOperation",
		"cannotParse",
		"fieldValidationFailed",
		"other",
	}
}
