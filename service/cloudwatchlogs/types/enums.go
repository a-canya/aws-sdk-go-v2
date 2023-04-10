// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DataProtectionStatus string

// Enum values for DataProtectionStatus
const (
	DataProtectionStatusActivated DataProtectionStatus = "ACTIVATED"
	DataProtectionStatusDeleted   DataProtectionStatus = "DELETED"
	DataProtectionStatusArchived  DataProtectionStatus = "ARCHIVED"
	DataProtectionStatusDisabled  DataProtectionStatus = "DISABLED"
)

// Values returns all known values for DataProtectionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataProtectionStatus) Values() []DataProtectionStatus {
	return []DataProtectionStatus{
		"ACTIVATED",
		"DELETED",
		"ARCHIVED",
		"DISABLED",
	}
}

type Distribution string

// Enum values for Distribution
const (
	DistributionRandom      Distribution = "Random"
	DistributionByLogStream Distribution = "ByLogStream"
)

// Values returns all known values for Distribution. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (Distribution) Values() []Distribution {
	return []Distribution{
		"Random",
		"ByLogStream",
	}
}

type ExportTaskStatusCode string

// Enum values for ExportTaskStatusCode
const (
	ExportTaskStatusCodeCancelled     ExportTaskStatusCode = "CANCELLED"
	ExportTaskStatusCodeCompleted     ExportTaskStatusCode = "COMPLETED"
	ExportTaskStatusCodeFailed        ExportTaskStatusCode = "FAILED"
	ExportTaskStatusCodePending       ExportTaskStatusCode = "PENDING"
	ExportTaskStatusCodePendingCancel ExportTaskStatusCode = "PENDING_CANCEL"
	ExportTaskStatusCodeRunning       ExportTaskStatusCode = "RUNNING"
)

// Values returns all known values for ExportTaskStatusCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExportTaskStatusCode) Values() []ExportTaskStatusCode {
	return []ExportTaskStatusCode{
		"CANCELLED",
		"COMPLETED",
		"FAILED",
		"PENDING",
		"PENDING_CANCEL",
		"RUNNING",
	}
}

type OrderBy string

// Enum values for OrderBy
const (
	OrderByLogStreamName OrderBy = "LogStreamName"
	OrderByLastEventTime OrderBy = "LastEventTime"
)

// Values returns all known values for OrderBy. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (OrderBy) Values() []OrderBy {
	return []OrderBy{
		"LogStreamName",
		"LastEventTime",
	}
}

type QueryStatus string

// Enum values for QueryStatus
const (
	QueryStatusScheduled QueryStatus = "Scheduled"
	QueryStatusRunning   QueryStatus = "Running"
	QueryStatusComplete  QueryStatus = "Complete"
	QueryStatusFailed    QueryStatus = "Failed"
	QueryStatusCancelled QueryStatus = "Cancelled"
	QueryStatusTimeout   QueryStatus = "Timeout"
	QueryStatusUnknown   QueryStatus = "Unknown"
)

// Values returns all known values for QueryStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (QueryStatus) Values() []QueryStatus {
	return []QueryStatus{
		"Scheduled",
		"Running",
		"Complete",
		"Failed",
		"Cancelled",
		"Timeout",
		"Unknown",
	}
}

type StandardUnit string

// Enum values for StandardUnit
const (
	StandardUnitSeconds         StandardUnit = "Seconds"
	StandardUnitMicroseconds    StandardUnit = "Microseconds"
	StandardUnitMilliseconds    StandardUnit = "Milliseconds"
	StandardUnitBytes           StandardUnit = "Bytes"
	StandardUnitKilobytes       StandardUnit = "Kilobytes"
	StandardUnitMegabytes       StandardUnit = "Megabytes"
	StandardUnitGigabytes       StandardUnit = "Gigabytes"
	StandardUnitTerabytes       StandardUnit = "Terabytes"
	StandardUnitBits            StandardUnit = "Bits"
	StandardUnitKilobits        StandardUnit = "Kilobits"
	StandardUnitMegabits        StandardUnit = "Megabits"
	StandardUnitGigabits        StandardUnit = "Gigabits"
	StandardUnitTerabits        StandardUnit = "Terabits"
	StandardUnitPercent         StandardUnit = "Percent"
	StandardUnitCount           StandardUnit = "Count"
	StandardUnitBytesSecond     StandardUnit = "Bytes/Second"
	StandardUnitKilobytesSecond StandardUnit = "Kilobytes/Second"
	StandardUnitMegabytesSecond StandardUnit = "Megabytes/Second"
	StandardUnitGigabytesSecond StandardUnit = "Gigabytes/Second"
	StandardUnitTerabytesSecond StandardUnit = "Terabytes/Second"
	StandardUnitBitsSecond      StandardUnit = "Bits/Second"
	StandardUnitKilobitsSecond  StandardUnit = "Kilobits/Second"
	StandardUnitMegabitsSecond  StandardUnit = "Megabits/Second"
	StandardUnitGigabitsSecond  StandardUnit = "Gigabits/Second"
	StandardUnitTerabitsSecond  StandardUnit = "Terabits/Second"
	StandardUnitCountSecond     StandardUnit = "Count/Second"
	StandardUnitNone            StandardUnit = "None"
)

// Values returns all known values for StandardUnit. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StandardUnit) Values() []StandardUnit {
	return []StandardUnit{
		"Seconds",
		"Microseconds",
		"Milliseconds",
		"Bytes",
		"Kilobytes",
		"Megabytes",
		"Gigabytes",
		"Terabytes",
		"Bits",
		"Kilobits",
		"Megabits",
		"Gigabits",
		"Terabits",
		"Percent",
		"Count",
		"Bytes/Second",
		"Kilobytes/Second",
		"Megabytes/Second",
		"Gigabytes/Second",
		"Terabytes/Second",
		"Bits/Second",
		"Kilobits/Second",
		"Megabits/Second",
		"Gigabits/Second",
		"Terabits/Second",
		"Count/Second",
		"None",
	}
}
