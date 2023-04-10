// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AttachmentStatus string

// Enum values for AttachmentStatus
const (
	AttachmentStatusAttaching AttachmentStatus = "ATTACHING"
	AttachmentStatusAttached  AttachmentStatus = "ATTACHED"
	AttachmentStatusDetaching AttachmentStatus = "DETACHING"
	AttachmentStatusDetached  AttachmentStatus = "DETACHED"
)

// Values returns all known values for AttachmentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AttachmentStatus) Values() []AttachmentStatus {
	return []AttachmentStatus{
		"ATTACHING",
		"ATTACHED",
		"DETACHING",
		"DETACHED",
	}
}

type ExecutionState string

// Enum values for ExecutionState
const (
	ExecutionStateQueued     ExecutionState = "QUEUED"
	ExecutionStateInProgress ExecutionState = "IN_PROGRESS"
	ExecutionStateCanceled   ExecutionState = "CANCELED"
	ExecutionStateFailed     ExecutionState = "FAILED"
	ExecutionStateSucceeded  ExecutionState = "SUCCEEDED"
	ExecutionStateRejected   ExecutionState = "REJECTED"
	ExecutionStateTimedOut   ExecutionState = "TIMED_OUT"
)

// Values returns all known values for ExecutionState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExecutionState) Values() []ExecutionState {
	return []ExecutionState{
		"QUEUED",
		"IN_PROGRESS",
		"CANCELED",
		"FAILED",
		"SUCCEEDED",
		"REJECTED",
		"TIMED_OUT",
	}
}

type InstanceStateName string

// Enum values for InstanceStateName
const (
	InstanceStateNamePending      InstanceStateName = "PENDING"
	InstanceStateNameRunning      InstanceStateName = "RUNNING"
	InstanceStateNameShuttingDown InstanceStateName = "SHUTTING_DOWN"
	InstanceStateNameTerminated   InstanceStateName = "TERMINATED"
	InstanceStateNameStopping     InstanceStateName = "STOPPING"
	InstanceStateNameStopped      InstanceStateName = "STOPPED"
)

// Values returns all known values for InstanceStateName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InstanceStateName) Values() []InstanceStateName {
	return []InstanceStateName{
		"PENDING",
		"RUNNING",
		"SHUTTING_DOWN",
		"TERMINATED",
		"STOPPING",
		"STOPPED",
	}
}

type IpAddressAssignment string

// Enum values for IpAddressAssignment
const (
	IpAddressAssignmentDhcp   IpAddressAssignment = "DHCP"
	IpAddressAssignmentStatic IpAddressAssignment = "STATIC"
)

// Values returns all known values for IpAddressAssignment. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IpAddressAssignment) Values() []IpAddressAssignment {
	return []IpAddressAssignment{
		"DHCP",
		"STATIC",
	}
}

type PhysicalConnectorType string

// Enum values for PhysicalConnectorType
const (
	PhysicalConnectorTypeRj45    PhysicalConnectorType = "RJ45"
	PhysicalConnectorTypeSfpPlus PhysicalConnectorType = "SFP_PLUS"
	PhysicalConnectorTypeQsfp    PhysicalConnectorType = "QSFP"
	PhysicalConnectorTypeRj452   PhysicalConnectorType = "RJ45_2"
	PhysicalConnectorTypeWifi    PhysicalConnectorType = "WIFI"
)

// Values returns all known values for PhysicalConnectorType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PhysicalConnectorType) Values() []PhysicalConnectorType {
	return []PhysicalConnectorType{
		"RJ45",
		"SFP_PLUS",
		"QSFP",
		"RJ45_2",
		"WIFI",
	}
}

type TaskState string

// Enum values for TaskState
const (
	TaskStateInProgress TaskState = "IN_PROGRESS"
	TaskStateCanceled   TaskState = "CANCELED"
	TaskStateCompleted  TaskState = "COMPLETED"
)

// Values returns all known values for TaskState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TaskState) Values() []TaskState {
	return []TaskState{
		"IN_PROGRESS",
		"CANCELED",
		"COMPLETED",
	}
}

type UnlockState string

// Enum values for UnlockState
const (
	UnlockStateUnlocked  UnlockState = "UNLOCKED"
	UnlockStateLocked    UnlockState = "LOCKED"
	UnlockStateUnlocking UnlockState = "UNLOCKING"
)

// Values returns all known values for UnlockState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (UnlockState) Values() []UnlockState {
	return []UnlockState{
		"UNLOCKED",
		"LOCKED",
		"UNLOCKING",
	}
}
