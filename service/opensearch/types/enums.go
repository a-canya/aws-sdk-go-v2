// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActionSeverity string

// Enum values for ActionSeverity
const (
	ActionSeverityHigh   ActionSeverity = "HIGH"
	ActionSeverityMedium ActionSeverity = "MEDIUM"
	ActionSeverityLow    ActionSeverity = "LOW"
)

// Values returns all known values for ActionSeverity. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ActionSeverity) Values() []ActionSeverity {
	return []ActionSeverity{
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

type ActionStatus string

// Enum values for ActionStatus
const (
	ActionStatusPendingUpdate ActionStatus = "PENDING_UPDATE"
	ActionStatusInProgress    ActionStatus = "IN_PROGRESS"
	ActionStatusFailed        ActionStatus = "FAILED"
	ActionStatusCompleted     ActionStatus = "COMPLETED"
	ActionStatusNotEligible   ActionStatus = "NOT_ELIGIBLE"
	ActionStatusEligible      ActionStatus = "ELIGIBLE"
)

// Values returns all known values for ActionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ActionStatus) Values() []ActionStatus {
	return []ActionStatus{
		"PENDING_UPDATE",
		"IN_PROGRESS",
		"FAILED",
		"COMPLETED",
		"NOT_ELIGIBLE",
		"ELIGIBLE",
	}
}

type ActionType string

// Enum values for ActionType
const (
	ActionTypeServiceSoftwareUpdate ActionType = "SERVICE_SOFTWARE_UPDATE"
	ActionTypeJvmHeapSizeTuning     ActionType = "JVM_HEAP_SIZE_TUNING"
	ActionTypeJvmYoungGenTuning     ActionType = "JVM_YOUNG_GEN_TUNING"
)

// Values returns all known values for ActionType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ActionType) Values() []ActionType {
	return []ActionType{
		"SERVICE_SOFTWARE_UPDATE",
		"JVM_HEAP_SIZE_TUNING",
		"JVM_YOUNG_GEN_TUNING",
	}
}

type AutoTuneDesiredState string

// Enum values for AutoTuneDesiredState
const (
	AutoTuneDesiredStateEnabled  AutoTuneDesiredState = "ENABLED"
	AutoTuneDesiredStateDisabled AutoTuneDesiredState = "DISABLED"
)

// Values returns all known values for AutoTuneDesiredState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AutoTuneDesiredState) Values() []AutoTuneDesiredState {
	return []AutoTuneDesiredState{
		"ENABLED",
		"DISABLED",
	}
}

type AutoTuneState string

// Enum values for AutoTuneState
const (
	AutoTuneStateEnabled                       AutoTuneState = "ENABLED"
	AutoTuneStateDisabled                      AutoTuneState = "DISABLED"
	AutoTuneStateEnableInProgress              AutoTuneState = "ENABLE_IN_PROGRESS"
	AutoTuneStateDisableInProgress             AutoTuneState = "DISABLE_IN_PROGRESS"
	AutoTuneStateDisabledAndRollbackScheduled  AutoTuneState = "DISABLED_AND_ROLLBACK_SCHEDULED"
	AutoTuneStateDisabledAndRollbackInProgress AutoTuneState = "DISABLED_AND_ROLLBACK_IN_PROGRESS"
	AutoTuneStateDisabledAndRollbackComplete   AutoTuneState = "DISABLED_AND_ROLLBACK_COMPLETE"
	AutoTuneStateDisabledAndRollbackError      AutoTuneState = "DISABLED_AND_ROLLBACK_ERROR"
	AutoTuneStateError                         AutoTuneState = "ERROR"
)

// Values returns all known values for AutoTuneState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AutoTuneState) Values() []AutoTuneState {
	return []AutoTuneState{
		"ENABLED",
		"DISABLED",
		"ENABLE_IN_PROGRESS",
		"DISABLE_IN_PROGRESS",
		"DISABLED_AND_ROLLBACK_SCHEDULED",
		"DISABLED_AND_ROLLBACK_IN_PROGRESS",
		"DISABLED_AND_ROLLBACK_COMPLETE",
		"DISABLED_AND_ROLLBACK_ERROR",
		"ERROR",
	}
}

type AutoTuneType string

// Enum values for AutoTuneType
const (
	AutoTuneTypeScheduledAction AutoTuneType = "SCHEDULED_ACTION"
)

// Values returns all known values for AutoTuneType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AutoTuneType) Values() []AutoTuneType {
	return []AutoTuneType{
		"SCHEDULED_ACTION",
	}
}

type ConnectionMode string

// Enum values for ConnectionMode
const (
	ConnectionModeDirect      ConnectionMode = "DIRECT"
	ConnectionModeVpcEndpoint ConnectionMode = "VPC_ENDPOINT"
)

// Values returns all known values for ConnectionMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConnectionMode) Values() []ConnectionMode {
	return []ConnectionMode{
		"DIRECT",
		"VPC_ENDPOINT",
	}
}

type DeploymentStatus string

// Enum values for DeploymentStatus
const (
	DeploymentStatusPendingUpdate DeploymentStatus = "PENDING_UPDATE"
	DeploymentStatusInProgress    DeploymentStatus = "IN_PROGRESS"
	DeploymentStatusCompleted     DeploymentStatus = "COMPLETED"
	DeploymentStatusNotEligible   DeploymentStatus = "NOT_ELIGIBLE"
	DeploymentStatusEligible      DeploymentStatus = "ELIGIBLE"
)

// Values returns all known values for DeploymentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentStatus) Values() []DeploymentStatus {
	return []DeploymentStatus{
		"PENDING_UPDATE",
		"IN_PROGRESS",
		"COMPLETED",
		"NOT_ELIGIBLE",
		"ELIGIBLE",
	}
}

type DescribePackagesFilterName string

// Enum values for DescribePackagesFilterName
const (
	DescribePackagesFilterNamePackageID     DescribePackagesFilterName = "PackageID"
	DescribePackagesFilterNamePackageName   DescribePackagesFilterName = "PackageName"
	DescribePackagesFilterNamePackageStatus DescribePackagesFilterName = "PackageStatus"
)

// Values returns all known values for DescribePackagesFilterName. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DescribePackagesFilterName) Values() []DescribePackagesFilterName {
	return []DescribePackagesFilterName{
		"PackageID",
		"PackageName",
		"PackageStatus",
	}
}

type DomainPackageStatus string

// Enum values for DomainPackageStatus
const (
	DomainPackageStatusAssociating        DomainPackageStatus = "ASSOCIATING"
	DomainPackageStatusAssociationFailed  DomainPackageStatus = "ASSOCIATION_FAILED"
	DomainPackageStatusActive             DomainPackageStatus = "ACTIVE"
	DomainPackageStatusDissociating       DomainPackageStatus = "DISSOCIATING"
	DomainPackageStatusDissociationFailed DomainPackageStatus = "DISSOCIATION_FAILED"
)

// Values returns all known values for DomainPackageStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DomainPackageStatus) Values() []DomainPackageStatus {
	return []DomainPackageStatus{
		"ASSOCIATING",
		"ASSOCIATION_FAILED",
		"ACTIVE",
		"DISSOCIATING",
		"DISSOCIATION_FAILED",
	}
}

type DryRunMode string

// Enum values for DryRunMode
const (
	DryRunModeBasic   DryRunMode = "Basic"
	DryRunModeVerbose DryRunMode = "Verbose"
)

// Values returns all known values for DryRunMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DryRunMode) Values() []DryRunMode {
	return []DryRunMode{
		"Basic",
		"Verbose",
	}
}

type EngineType string

// Enum values for EngineType
const (
	EngineTypeOpenSearch    EngineType = "OpenSearch"
	EngineTypeElasticsearch EngineType = "Elasticsearch"
)

// Values returns all known values for EngineType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EngineType) Values() []EngineType {
	return []EngineType{
		"OpenSearch",
		"Elasticsearch",
	}
}

type InboundConnectionStatusCode string

// Enum values for InboundConnectionStatusCode
const (
	InboundConnectionStatusCodePendingAcceptance InboundConnectionStatusCode = "PENDING_ACCEPTANCE"
	InboundConnectionStatusCodeApproved          InboundConnectionStatusCode = "APPROVED"
	InboundConnectionStatusCodeProvisioning      InboundConnectionStatusCode = "PROVISIONING"
	InboundConnectionStatusCodeActive            InboundConnectionStatusCode = "ACTIVE"
	InboundConnectionStatusCodeRejecting         InboundConnectionStatusCode = "REJECTING"
	InboundConnectionStatusCodeRejected          InboundConnectionStatusCode = "REJECTED"
	InboundConnectionStatusCodeDeleting          InboundConnectionStatusCode = "DELETING"
	InboundConnectionStatusCodeDeleted           InboundConnectionStatusCode = "DELETED"
)

// Values returns all known values for InboundConnectionStatusCode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (InboundConnectionStatusCode) Values() []InboundConnectionStatusCode {
	return []InboundConnectionStatusCode{
		"PENDING_ACCEPTANCE",
		"APPROVED",
		"PROVISIONING",
		"ACTIVE",
		"REJECTING",
		"REJECTED",
		"DELETING",
		"DELETED",
	}
}

type LogType string

// Enum values for LogType
const (
	LogTypeIndexSlowLogs     LogType = "INDEX_SLOW_LOGS"
	LogTypeSearchSlowLogs    LogType = "SEARCH_SLOW_LOGS"
	LogTypeEsApplicationLogs LogType = "ES_APPLICATION_LOGS"
	LogTypeAuditLogs         LogType = "AUDIT_LOGS"
)

// Values returns all known values for LogType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LogType) Values() []LogType {
	return []LogType{
		"INDEX_SLOW_LOGS",
		"SEARCH_SLOW_LOGS",
		"ES_APPLICATION_LOGS",
		"AUDIT_LOGS",
	}
}

type OpenSearchPartitionInstanceType string

// Enum values for OpenSearchPartitionInstanceType
const (
	OpenSearchPartitionInstanceTypeM3MediumSearch         OpenSearchPartitionInstanceType = "m3.medium.search"
	OpenSearchPartitionInstanceTypeM3LargeSearch          OpenSearchPartitionInstanceType = "m3.large.search"
	OpenSearchPartitionInstanceTypeM3XlargeSearch         OpenSearchPartitionInstanceType = "m3.xlarge.search"
	OpenSearchPartitionInstanceTypeM32xlargeSearch        OpenSearchPartitionInstanceType = "m3.2xlarge.search"
	OpenSearchPartitionInstanceTypeM4LargeSearch          OpenSearchPartitionInstanceType = "m4.large.search"
	OpenSearchPartitionInstanceTypeM4XlargeSearch         OpenSearchPartitionInstanceType = "m4.xlarge.search"
	OpenSearchPartitionInstanceTypeM42xlargeSearch        OpenSearchPartitionInstanceType = "m4.2xlarge.search"
	OpenSearchPartitionInstanceTypeM44xlargeSearch        OpenSearchPartitionInstanceType = "m4.4xlarge.search"
	OpenSearchPartitionInstanceTypeM410xlargeSearch       OpenSearchPartitionInstanceType = "m4.10xlarge.search"
	OpenSearchPartitionInstanceTypeM5LargeSearch          OpenSearchPartitionInstanceType = "m5.large.search"
	OpenSearchPartitionInstanceTypeM5XlargeSearch         OpenSearchPartitionInstanceType = "m5.xlarge.search"
	OpenSearchPartitionInstanceTypeM52xlargeSearch        OpenSearchPartitionInstanceType = "m5.2xlarge.search"
	OpenSearchPartitionInstanceTypeM54xlargeSearch        OpenSearchPartitionInstanceType = "m5.4xlarge.search"
	OpenSearchPartitionInstanceTypeM512xlargeSearch       OpenSearchPartitionInstanceType = "m5.12xlarge.search"
	OpenSearchPartitionInstanceTypeM524xlargeSearch       OpenSearchPartitionInstanceType = "m5.24xlarge.search"
	OpenSearchPartitionInstanceTypeR5LargeSearch          OpenSearchPartitionInstanceType = "r5.large.search"
	OpenSearchPartitionInstanceTypeR5XlargeSearch         OpenSearchPartitionInstanceType = "r5.xlarge.search"
	OpenSearchPartitionInstanceTypeR52xlargeSearch        OpenSearchPartitionInstanceType = "r5.2xlarge.search"
	OpenSearchPartitionInstanceTypeR54xlargeSearch        OpenSearchPartitionInstanceType = "r5.4xlarge.search"
	OpenSearchPartitionInstanceTypeR512xlargeSearch       OpenSearchPartitionInstanceType = "r5.12xlarge.search"
	OpenSearchPartitionInstanceTypeR524xlargeSearch       OpenSearchPartitionInstanceType = "r5.24xlarge.search"
	OpenSearchPartitionInstanceTypeC5LargeSearch          OpenSearchPartitionInstanceType = "c5.large.search"
	OpenSearchPartitionInstanceTypeC5XlargeSearch         OpenSearchPartitionInstanceType = "c5.xlarge.search"
	OpenSearchPartitionInstanceTypeC52xlargeSearch        OpenSearchPartitionInstanceType = "c5.2xlarge.search"
	OpenSearchPartitionInstanceTypeC54xlargeSearch        OpenSearchPartitionInstanceType = "c5.4xlarge.search"
	OpenSearchPartitionInstanceTypeC59xlargeSearch        OpenSearchPartitionInstanceType = "c5.9xlarge.search"
	OpenSearchPartitionInstanceTypeC518xlargeSearch       OpenSearchPartitionInstanceType = "c5.18xlarge.search"
	OpenSearchPartitionInstanceTypeT3NanoSearch           OpenSearchPartitionInstanceType = "t3.nano.search"
	OpenSearchPartitionInstanceTypeT3MicroSearch          OpenSearchPartitionInstanceType = "t3.micro.search"
	OpenSearchPartitionInstanceTypeT3SmallSearch          OpenSearchPartitionInstanceType = "t3.small.search"
	OpenSearchPartitionInstanceTypeT3MediumSearch         OpenSearchPartitionInstanceType = "t3.medium.search"
	OpenSearchPartitionInstanceTypeT3LargeSearch          OpenSearchPartitionInstanceType = "t3.large.search"
	OpenSearchPartitionInstanceTypeT3XlargeSearch         OpenSearchPartitionInstanceType = "t3.xlarge.search"
	OpenSearchPartitionInstanceTypeT32xlargeSearch        OpenSearchPartitionInstanceType = "t3.2xlarge.search"
	OpenSearchPartitionInstanceTypeUltrawarm1MediumSearch OpenSearchPartitionInstanceType = "ultrawarm1.medium.search"
	OpenSearchPartitionInstanceTypeUltrawarm1LargeSearch  OpenSearchPartitionInstanceType = "ultrawarm1.large.search"
	OpenSearchPartitionInstanceTypeUltrawarm1XlargeSearch OpenSearchPartitionInstanceType = "ultrawarm1.xlarge.search"
	OpenSearchPartitionInstanceTypeT2MicroSearch          OpenSearchPartitionInstanceType = "t2.micro.search"
	OpenSearchPartitionInstanceTypeT2SmallSearch          OpenSearchPartitionInstanceType = "t2.small.search"
	OpenSearchPartitionInstanceTypeT2MediumSearch         OpenSearchPartitionInstanceType = "t2.medium.search"
	OpenSearchPartitionInstanceTypeR3LargeSearch          OpenSearchPartitionInstanceType = "r3.large.search"
	OpenSearchPartitionInstanceTypeR3XlargeSearch         OpenSearchPartitionInstanceType = "r3.xlarge.search"
	OpenSearchPartitionInstanceTypeR32xlargeSearch        OpenSearchPartitionInstanceType = "r3.2xlarge.search"
	OpenSearchPartitionInstanceTypeR34xlargeSearch        OpenSearchPartitionInstanceType = "r3.4xlarge.search"
	OpenSearchPartitionInstanceTypeR38xlargeSearch        OpenSearchPartitionInstanceType = "r3.8xlarge.search"
	OpenSearchPartitionInstanceTypeI2XlargeSearch         OpenSearchPartitionInstanceType = "i2.xlarge.search"
	OpenSearchPartitionInstanceTypeI22xlargeSearch        OpenSearchPartitionInstanceType = "i2.2xlarge.search"
	OpenSearchPartitionInstanceTypeD2XlargeSearch         OpenSearchPartitionInstanceType = "d2.xlarge.search"
	OpenSearchPartitionInstanceTypeD22xlargeSearch        OpenSearchPartitionInstanceType = "d2.2xlarge.search"
	OpenSearchPartitionInstanceTypeD24xlargeSearch        OpenSearchPartitionInstanceType = "d2.4xlarge.search"
	OpenSearchPartitionInstanceTypeD28xlargeSearch        OpenSearchPartitionInstanceType = "d2.8xlarge.search"
	OpenSearchPartitionInstanceTypeC4LargeSearch          OpenSearchPartitionInstanceType = "c4.large.search"
	OpenSearchPartitionInstanceTypeC4XlargeSearch         OpenSearchPartitionInstanceType = "c4.xlarge.search"
	OpenSearchPartitionInstanceTypeC42xlargeSearch        OpenSearchPartitionInstanceType = "c4.2xlarge.search"
	OpenSearchPartitionInstanceTypeC44xlargeSearch        OpenSearchPartitionInstanceType = "c4.4xlarge.search"
	OpenSearchPartitionInstanceTypeC48xlargeSearch        OpenSearchPartitionInstanceType = "c4.8xlarge.search"
	OpenSearchPartitionInstanceTypeR4LargeSearch          OpenSearchPartitionInstanceType = "r4.large.search"
	OpenSearchPartitionInstanceTypeR4XlargeSearch         OpenSearchPartitionInstanceType = "r4.xlarge.search"
	OpenSearchPartitionInstanceTypeR42xlargeSearch        OpenSearchPartitionInstanceType = "r4.2xlarge.search"
	OpenSearchPartitionInstanceTypeR44xlargeSearch        OpenSearchPartitionInstanceType = "r4.4xlarge.search"
	OpenSearchPartitionInstanceTypeR48xlargeSearch        OpenSearchPartitionInstanceType = "r4.8xlarge.search"
	OpenSearchPartitionInstanceTypeR416xlargeSearch       OpenSearchPartitionInstanceType = "r4.16xlarge.search"
	OpenSearchPartitionInstanceTypeI3LargeSearch          OpenSearchPartitionInstanceType = "i3.large.search"
	OpenSearchPartitionInstanceTypeI3XlargeSearch         OpenSearchPartitionInstanceType = "i3.xlarge.search"
	OpenSearchPartitionInstanceTypeI32xlargeSearch        OpenSearchPartitionInstanceType = "i3.2xlarge.search"
	OpenSearchPartitionInstanceTypeI34xlargeSearch        OpenSearchPartitionInstanceType = "i3.4xlarge.search"
	OpenSearchPartitionInstanceTypeI38xlargeSearch        OpenSearchPartitionInstanceType = "i3.8xlarge.search"
	OpenSearchPartitionInstanceTypeI316xlargeSearch       OpenSearchPartitionInstanceType = "i3.16xlarge.search"
	OpenSearchPartitionInstanceTypeR6gLargeSearch         OpenSearchPartitionInstanceType = "r6g.large.search"
	OpenSearchPartitionInstanceTypeR6gXlargeSearch        OpenSearchPartitionInstanceType = "r6g.xlarge.search"
	OpenSearchPartitionInstanceTypeR6g2xlargeSearch       OpenSearchPartitionInstanceType = "r6g.2xlarge.search"
	OpenSearchPartitionInstanceTypeR6g4xlargeSearch       OpenSearchPartitionInstanceType = "r6g.4xlarge.search"
	OpenSearchPartitionInstanceTypeR6g8xlargeSearch       OpenSearchPartitionInstanceType = "r6g.8xlarge.search"
	OpenSearchPartitionInstanceTypeR6g12xlargeSearch      OpenSearchPartitionInstanceType = "r6g.12xlarge.search"
	OpenSearchPartitionInstanceTypeM6gLargeSearch         OpenSearchPartitionInstanceType = "m6g.large.search"
	OpenSearchPartitionInstanceTypeM6gXlargeSearch        OpenSearchPartitionInstanceType = "m6g.xlarge.search"
	OpenSearchPartitionInstanceTypeM6g2xlargeSearch       OpenSearchPartitionInstanceType = "m6g.2xlarge.search"
	OpenSearchPartitionInstanceTypeM6g4xlargeSearch       OpenSearchPartitionInstanceType = "m6g.4xlarge.search"
	OpenSearchPartitionInstanceTypeM6g8xlargeSearch       OpenSearchPartitionInstanceType = "m6g.8xlarge.search"
	OpenSearchPartitionInstanceTypeM6g12xlargeSearch      OpenSearchPartitionInstanceType = "m6g.12xlarge.search"
	OpenSearchPartitionInstanceTypeC6gLargeSearch         OpenSearchPartitionInstanceType = "c6g.large.search"
	OpenSearchPartitionInstanceTypeC6gXlargeSearch        OpenSearchPartitionInstanceType = "c6g.xlarge.search"
	OpenSearchPartitionInstanceTypeC6g2xlargeSearch       OpenSearchPartitionInstanceType = "c6g.2xlarge.search"
	OpenSearchPartitionInstanceTypeC6g4xlargeSearch       OpenSearchPartitionInstanceType = "c6g.4xlarge.search"
	OpenSearchPartitionInstanceTypeC6g8xlargeSearch       OpenSearchPartitionInstanceType = "c6g.8xlarge.search"
	OpenSearchPartitionInstanceTypeC6g12xlargeSearch      OpenSearchPartitionInstanceType = "c6g.12xlarge.search"
	OpenSearchPartitionInstanceTypeR6gdLargeSearch        OpenSearchPartitionInstanceType = "r6gd.large.search"
	OpenSearchPartitionInstanceTypeR6gdXlargeSearch       OpenSearchPartitionInstanceType = "r6gd.xlarge.search"
	OpenSearchPartitionInstanceTypeR6gd2xlargeSearch      OpenSearchPartitionInstanceType = "r6gd.2xlarge.search"
	OpenSearchPartitionInstanceTypeR6gd4xlargeSearch      OpenSearchPartitionInstanceType = "r6gd.4xlarge.search"
	OpenSearchPartitionInstanceTypeR6gd8xlargeSearch      OpenSearchPartitionInstanceType = "r6gd.8xlarge.search"
	OpenSearchPartitionInstanceTypeR6gd12xlargeSearch     OpenSearchPartitionInstanceType = "r6gd.12xlarge.search"
	OpenSearchPartitionInstanceTypeR6gd16xlargeSearch     OpenSearchPartitionInstanceType = "r6gd.16xlarge.search"
	OpenSearchPartitionInstanceTypeT4gSmallSearch         OpenSearchPartitionInstanceType = "t4g.small.search"
	OpenSearchPartitionInstanceTypeT4gMediumSearch        OpenSearchPartitionInstanceType = "t4g.medium.search"
)

// Values returns all known values for OpenSearchPartitionInstanceType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (OpenSearchPartitionInstanceType) Values() []OpenSearchPartitionInstanceType {
	return []OpenSearchPartitionInstanceType{
		"m3.medium.search",
		"m3.large.search",
		"m3.xlarge.search",
		"m3.2xlarge.search",
		"m4.large.search",
		"m4.xlarge.search",
		"m4.2xlarge.search",
		"m4.4xlarge.search",
		"m4.10xlarge.search",
		"m5.large.search",
		"m5.xlarge.search",
		"m5.2xlarge.search",
		"m5.4xlarge.search",
		"m5.12xlarge.search",
		"m5.24xlarge.search",
		"r5.large.search",
		"r5.xlarge.search",
		"r5.2xlarge.search",
		"r5.4xlarge.search",
		"r5.12xlarge.search",
		"r5.24xlarge.search",
		"c5.large.search",
		"c5.xlarge.search",
		"c5.2xlarge.search",
		"c5.4xlarge.search",
		"c5.9xlarge.search",
		"c5.18xlarge.search",
		"t3.nano.search",
		"t3.micro.search",
		"t3.small.search",
		"t3.medium.search",
		"t3.large.search",
		"t3.xlarge.search",
		"t3.2xlarge.search",
		"ultrawarm1.medium.search",
		"ultrawarm1.large.search",
		"ultrawarm1.xlarge.search",
		"t2.micro.search",
		"t2.small.search",
		"t2.medium.search",
		"r3.large.search",
		"r3.xlarge.search",
		"r3.2xlarge.search",
		"r3.4xlarge.search",
		"r3.8xlarge.search",
		"i2.xlarge.search",
		"i2.2xlarge.search",
		"d2.xlarge.search",
		"d2.2xlarge.search",
		"d2.4xlarge.search",
		"d2.8xlarge.search",
		"c4.large.search",
		"c4.xlarge.search",
		"c4.2xlarge.search",
		"c4.4xlarge.search",
		"c4.8xlarge.search",
		"r4.large.search",
		"r4.xlarge.search",
		"r4.2xlarge.search",
		"r4.4xlarge.search",
		"r4.8xlarge.search",
		"r4.16xlarge.search",
		"i3.large.search",
		"i3.xlarge.search",
		"i3.2xlarge.search",
		"i3.4xlarge.search",
		"i3.8xlarge.search",
		"i3.16xlarge.search",
		"r6g.large.search",
		"r6g.xlarge.search",
		"r6g.2xlarge.search",
		"r6g.4xlarge.search",
		"r6g.8xlarge.search",
		"r6g.12xlarge.search",
		"m6g.large.search",
		"m6g.xlarge.search",
		"m6g.2xlarge.search",
		"m6g.4xlarge.search",
		"m6g.8xlarge.search",
		"m6g.12xlarge.search",
		"c6g.large.search",
		"c6g.xlarge.search",
		"c6g.2xlarge.search",
		"c6g.4xlarge.search",
		"c6g.8xlarge.search",
		"c6g.12xlarge.search",
		"r6gd.large.search",
		"r6gd.xlarge.search",
		"r6gd.2xlarge.search",
		"r6gd.4xlarge.search",
		"r6gd.8xlarge.search",
		"r6gd.12xlarge.search",
		"r6gd.16xlarge.search",
		"t4g.small.search",
		"t4g.medium.search",
	}
}

type OpenSearchWarmPartitionInstanceType string

// Enum values for OpenSearchWarmPartitionInstanceType
const (
	OpenSearchWarmPartitionInstanceTypeUltrawarm1MediumSearch OpenSearchWarmPartitionInstanceType = "ultrawarm1.medium.search"
	OpenSearchWarmPartitionInstanceTypeUltrawarm1LargeSearch  OpenSearchWarmPartitionInstanceType = "ultrawarm1.large.search"
	OpenSearchWarmPartitionInstanceTypeUltrawarm1XlargeSearch OpenSearchWarmPartitionInstanceType = "ultrawarm1.xlarge.search"
)

// Values returns all known values for OpenSearchWarmPartitionInstanceType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (OpenSearchWarmPartitionInstanceType) Values() []OpenSearchWarmPartitionInstanceType {
	return []OpenSearchWarmPartitionInstanceType{
		"ultrawarm1.medium.search",
		"ultrawarm1.large.search",
		"ultrawarm1.xlarge.search",
	}
}

type OptionState string

// Enum values for OptionState
const (
	OptionStateRequiresIndexDocuments OptionState = "RequiresIndexDocuments"
	OptionStateProcessing             OptionState = "Processing"
	OptionStateActive                 OptionState = "Active"
)

// Values returns all known values for OptionState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OptionState) Values() []OptionState {
	return []OptionState{
		"RequiresIndexDocuments",
		"Processing",
		"Active",
	}
}

type OutboundConnectionStatusCode string

// Enum values for OutboundConnectionStatusCode
const (
	OutboundConnectionStatusCodeValidating        OutboundConnectionStatusCode = "VALIDATING"
	OutboundConnectionStatusCodeValidationFailed  OutboundConnectionStatusCode = "VALIDATION_FAILED"
	OutboundConnectionStatusCodePendingAcceptance OutboundConnectionStatusCode = "PENDING_ACCEPTANCE"
	OutboundConnectionStatusCodeApproved          OutboundConnectionStatusCode = "APPROVED"
	OutboundConnectionStatusCodeProvisioning      OutboundConnectionStatusCode = "PROVISIONING"
	OutboundConnectionStatusCodeActive            OutboundConnectionStatusCode = "ACTIVE"
	OutboundConnectionStatusCodeRejecting         OutboundConnectionStatusCode = "REJECTING"
	OutboundConnectionStatusCodeRejected          OutboundConnectionStatusCode = "REJECTED"
	OutboundConnectionStatusCodeDeleting          OutboundConnectionStatusCode = "DELETING"
	OutboundConnectionStatusCodeDeleted           OutboundConnectionStatusCode = "DELETED"
)

// Values returns all known values for OutboundConnectionStatusCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (OutboundConnectionStatusCode) Values() []OutboundConnectionStatusCode {
	return []OutboundConnectionStatusCode{
		"VALIDATING",
		"VALIDATION_FAILED",
		"PENDING_ACCEPTANCE",
		"APPROVED",
		"PROVISIONING",
		"ACTIVE",
		"REJECTING",
		"REJECTED",
		"DELETING",
		"DELETED",
	}
}

type OverallChangeStatus string

// Enum values for OverallChangeStatus
const (
	OverallChangeStatusPending    OverallChangeStatus = "PENDING"
	OverallChangeStatusProcessing OverallChangeStatus = "PROCESSING"
	OverallChangeStatusCompleted  OverallChangeStatus = "COMPLETED"
	OverallChangeStatusFailed     OverallChangeStatus = "FAILED"
)

// Values returns all known values for OverallChangeStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OverallChangeStatus) Values() []OverallChangeStatus {
	return []OverallChangeStatus{
		"PENDING",
		"PROCESSING",
		"COMPLETED",
		"FAILED",
	}
}

type PackageStatus string

// Enum values for PackageStatus
const (
	PackageStatusCopying          PackageStatus = "COPYING"
	PackageStatusCopyFailed       PackageStatus = "COPY_FAILED"
	PackageStatusValidating       PackageStatus = "VALIDATING"
	PackageStatusValidationFailed PackageStatus = "VALIDATION_FAILED"
	PackageStatusAvailable        PackageStatus = "AVAILABLE"
	PackageStatusDeleting         PackageStatus = "DELETING"
	PackageStatusDeleted          PackageStatus = "DELETED"
	PackageStatusDeleteFailed     PackageStatus = "DELETE_FAILED"
)

// Values returns all known values for PackageStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PackageStatus) Values() []PackageStatus {
	return []PackageStatus{
		"COPYING",
		"COPY_FAILED",
		"VALIDATING",
		"VALIDATION_FAILED",
		"AVAILABLE",
		"DELETING",
		"DELETED",
		"DELETE_FAILED",
	}
}

type PackageType string

// Enum values for PackageType
const (
	PackageTypeTxtDictionary PackageType = "TXT-DICTIONARY"
)

// Values returns all known values for PackageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PackageType) Values() []PackageType {
	return []PackageType{
		"TXT-DICTIONARY",
	}
}

type PrincipalType string

// Enum values for PrincipalType
const (
	PrincipalTypeAwsAccount PrincipalType = "AWS_ACCOUNT"
	PrincipalTypeAwsService PrincipalType = "AWS_SERVICE"
)

// Values returns all known values for PrincipalType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PrincipalType) Values() []PrincipalType {
	return []PrincipalType{
		"AWS_ACCOUNT",
		"AWS_SERVICE",
	}
}

type ReservedInstancePaymentOption string

// Enum values for ReservedInstancePaymentOption
const (
	ReservedInstancePaymentOptionAllUpfront     ReservedInstancePaymentOption = "ALL_UPFRONT"
	ReservedInstancePaymentOptionPartialUpfront ReservedInstancePaymentOption = "PARTIAL_UPFRONT"
	ReservedInstancePaymentOptionNoUpfront      ReservedInstancePaymentOption = "NO_UPFRONT"
)

// Values returns all known values for ReservedInstancePaymentOption. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ReservedInstancePaymentOption) Values() []ReservedInstancePaymentOption {
	return []ReservedInstancePaymentOption{
		"ALL_UPFRONT",
		"PARTIAL_UPFRONT",
		"NO_UPFRONT",
	}
}

type RollbackOnDisable string

// Enum values for RollbackOnDisable
const (
	RollbackOnDisableNoRollback      RollbackOnDisable = "NO_ROLLBACK"
	RollbackOnDisableDefaultRollback RollbackOnDisable = "DEFAULT_ROLLBACK"
)

// Values returns all known values for RollbackOnDisable. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RollbackOnDisable) Values() []RollbackOnDisable {
	return []RollbackOnDisable{
		"NO_ROLLBACK",
		"DEFAULT_ROLLBACK",
	}
}

type ScheduleAt string

// Enum values for ScheduleAt
const (
	ScheduleAtNow           ScheduleAt = "NOW"
	ScheduleAtTimestamp     ScheduleAt = "TIMESTAMP"
	ScheduleAtOffPeakWindow ScheduleAt = "OFF_PEAK_WINDOW"
)

// Values returns all known values for ScheduleAt. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ScheduleAt) Values() []ScheduleAt {
	return []ScheduleAt{
		"NOW",
		"TIMESTAMP",
		"OFF_PEAK_WINDOW",
	}
}

type ScheduledAutoTuneActionType string

// Enum values for ScheduledAutoTuneActionType
const (
	ScheduledAutoTuneActionTypeJvmHeapSizeTuning ScheduledAutoTuneActionType = "JVM_HEAP_SIZE_TUNING"
	ScheduledAutoTuneActionTypeJvmYoungGenTuning ScheduledAutoTuneActionType = "JVM_YOUNG_GEN_TUNING"
)

// Values returns all known values for ScheduledAutoTuneActionType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScheduledAutoTuneActionType) Values() []ScheduledAutoTuneActionType {
	return []ScheduledAutoTuneActionType{
		"JVM_HEAP_SIZE_TUNING",
		"JVM_YOUNG_GEN_TUNING",
	}
}

type ScheduledAutoTuneSeverityType string

// Enum values for ScheduledAutoTuneSeverityType
const (
	ScheduledAutoTuneSeverityTypeLow    ScheduledAutoTuneSeverityType = "LOW"
	ScheduledAutoTuneSeverityTypeMedium ScheduledAutoTuneSeverityType = "MEDIUM"
	ScheduledAutoTuneSeverityTypeHigh   ScheduledAutoTuneSeverityType = "HIGH"
)

// Values returns all known values for ScheduledAutoTuneSeverityType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ScheduledAutoTuneSeverityType) Values() []ScheduledAutoTuneSeverityType {
	return []ScheduledAutoTuneSeverityType{
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

type ScheduledBy string

// Enum values for ScheduledBy
const (
	ScheduledByCustomer ScheduledBy = "CUSTOMER"
	ScheduledBySystem   ScheduledBy = "SYSTEM"
)

// Values returns all known values for ScheduledBy. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ScheduledBy) Values() []ScheduledBy {
	return []ScheduledBy{
		"CUSTOMER",
		"SYSTEM",
	}
}

type TimeUnit string

// Enum values for TimeUnit
const (
	TimeUnitHours TimeUnit = "HOURS"
)

// Values returns all known values for TimeUnit. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TimeUnit) Values() []TimeUnit {
	return []TimeUnit{
		"HOURS",
	}
}

type TLSSecurityPolicy string

// Enum values for TLSSecurityPolicy
const (
	TLSSecurityPolicyPolicyMinTls10201907 TLSSecurityPolicy = "Policy-Min-TLS-1-0-2019-07"
	TLSSecurityPolicyPolicyMinTls12201907 TLSSecurityPolicy = "Policy-Min-TLS-1-2-2019-07"
)

// Values returns all known values for TLSSecurityPolicy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TLSSecurityPolicy) Values() []TLSSecurityPolicy {
	return []TLSSecurityPolicy{
		"Policy-Min-TLS-1-0-2019-07",
		"Policy-Min-TLS-1-2-2019-07",
	}
}

type UpgradeStatus string

// Enum values for UpgradeStatus
const (
	UpgradeStatusInProgress          UpgradeStatus = "IN_PROGRESS"
	UpgradeStatusSucceeded           UpgradeStatus = "SUCCEEDED"
	UpgradeStatusSucceededWithIssues UpgradeStatus = "SUCCEEDED_WITH_ISSUES"
	UpgradeStatusFailed              UpgradeStatus = "FAILED"
)

// Values returns all known values for UpgradeStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UpgradeStatus) Values() []UpgradeStatus {
	return []UpgradeStatus{
		"IN_PROGRESS",
		"SUCCEEDED",
		"SUCCEEDED_WITH_ISSUES",
		"FAILED",
	}
}

type UpgradeStep string

// Enum values for UpgradeStep
const (
	UpgradeStepPreUpgradeCheck UpgradeStep = "PRE_UPGRADE_CHECK"
	UpgradeStepSnapshot        UpgradeStep = "SNAPSHOT"
	UpgradeStepUpgrade         UpgradeStep = "UPGRADE"
)

// Values returns all known values for UpgradeStep. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (UpgradeStep) Values() []UpgradeStep {
	return []UpgradeStep{
		"PRE_UPGRADE_CHECK",
		"SNAPSHOT",
		"UPGRADE",
	}
}

type VolumeType string

// Enum values for VolumeType
const (
	VolumeTypeStandard VolumeType = "standard"
	VolumeTypeGp2      VolumeType = "gp2"
	VolumeTypeIo1      VolumeType = "io1"
	VolumeTypeGp3      VolumeType = "gp3"
)

// Values returns all known values for VolumeType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (VolumeType) Values() []VolumeType {
	return []VolumeType{
		"standard",
		"gp2",
		"io1",
		"gp3",
	}
}

type VpcEndpointErrorCode string

// Enum values for VpcEndpointErrorCode
const (
	VpcEndpointErrorCodeEndpointNotFound VpcEndpointErrorCode = "ENDPOINT_NOT_FOUND"
	VpcEndpointErrorCodeServerError      VpcEndpointErrorCode = "SERVER_ERROR"
)

// Values returns all known values for VpcEndpointErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VpcEndpointErrorCode) Values() []VpcEndpointErrorCode {
	return []VpcEndpointErrorCode{
		"ENDPOINT_NOT_FOUND",
		"SERVER_ERROR",
	}
}

type VpcEndpointStatus string

// Enum values for VpcEndpointStatus
const (
	VpcEndpointStatusCreating     VpcEndpointStatus = "CREATING"
	VpcEndpointStatusCreateFailed VpcEndpointStatus = "CREATE_FAILED"
	VpcEndpointStatusActive       VpcEndpointStatus = "ACTIVE"
	VpcEndpointStatusUpdating     VpcEndpointStatus = "UPDATING"
	VpcEndpointStatusUpdateFailed VpcEndpointStatus = "UPDATE_FAILED"
	VpcEndpointStatusDeleting     VpcEndpointStatus = "DELETING"
	VpcEndpointStatusDeleteFailed VpcEndpointStatus = "DELETE_FAILED"
)

// Values returns all known values for VpcEndpointStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VpcEndpointStatus) Values() []VpcEndpointStatus {
	return []VpcEndpointStatus{
		"CREATING",
		"CREATE_FAILED",
		"ACTIVE",
		"UPDATING",
		"UPDATE_FAILED",
		"DELETING",
		"DELETE_FAILED",
	}
}
