// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AcceptCodeValidation string

// Enum values for AcceptCodeValidation
const (
	AcceptCodeValidationIgnore  AcceptCodeValidation = "IGNORE"
	AcceptCodeValidationEnforce AcceptCodeValidation = "ENFORCE"
)

// Values returns all known values for AcceptCodeValidation. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AcceptCodeValidation) Values() []AcceptCodeValidation {
	return []AcceptCodeValidation{
		"IGNORE",
		"ENFORCE",
	}
}

type AcceptType string

// Enum values for AcceptType
const (
	AcceptTypeDelivered AcceptType = "DELIVERED"
	AcceptTypeRead      AcceptType = "READ"
)

// Values returns all known values for AcceptType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AcceptType) Values() []AcceptType {
	return []AcceptType{
		"DELIVERED",
		"READ",
	}
}

type ActivationStatus string

// Enum values for ActivationStatus
const (
	ActivationStatusActivated    ActivationStatus = "ACTIVATED"
	ActivationStatusNotActivated ActivationStatus = "NOT_ACTIVATED"
)

// Values returns all known values for ActivationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ActivationStatus) Values() []ActivationStatus {
	return []ActivationStatus{
		"ACTIVATED",
		"NOT_ACTIVATED",
	}
}

type ChannelType string

// Enum values for ChannelType
const (
	ChannelTypeSms   ChannelType = "SMS"
	ChannelTypeVoice ChannelType = "VOICE"
	ChannelTypeEmail ChannelType = "EMAIL"
)

// Values returns all known values for ChannelType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChannelType) Values() []ChannelType {
	return []ChannelType{
		"SMS",
		"VOICE",
		"EMAIL",
	}
}

type ContactType string

// Enum values for ContactType
const (
	ContactTypePersonal       ContactType = "PERSONAL"
	ContactTypeEscalation     ContactType = "ESCALATION"
	ContactTypeOncallSchedule ContactType = "ONCALL_SCHEDULE"
)

// Values returns all known values for ContactType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ContactType) Values() []ContactType {
	return []ContactType{
		"PERSONAL",
		"ESCALATION",
		"ONCALL_SCHEDULE",
	}
}

type DayOfWeek string

// Enum values for DayOfWeek
const (
	DayOfWeekMon DayOfWeek = "MON"
	DayOfWeekTue DayOfWeek = "TUE"
	DayOfWeekWed DayOfWeek = "WED"
	DayOfWeekThu DayOfWeek = "THU"
	DayOfWeekFri DayOfWeek = "FRI"
	DayOfWeekSat DayOfWeek = "SAT"
	DayOfWeekSun DayOfWeek = "SUN"
)

// Values returns all known values for DayOfWeek. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DayOfWeek) Values() []DayOfWeek {
	return []DayOfWeek{
		"MON",
		"TUE",
		"WED",
		"THU",
		"FRI",
		"SAT",
		"SUN",
	}
}

type ReceiptType string

// Enum values for ReceiptType
const (
	ReceiptTypeDelivered ReceiptType = "DELIVERED"
	ReceiptTypeError     ReceiptType = "ERROR"
	ReceiptTypeRead      ReceiptType = "READ"
	ReceiptTypeSent      ReceiptType = "SENT"
	ReceiptTypeStop      ReceiptType = "STOP"
)

// Values returns all known values for ReceiptType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ReceiptType) Values() []ReceiptType {
	return []ReceiptType{
		"DELIVERED",
		"ERROR",
		"READ",
		"SENT",
		"STOP",
	}
}

type ShiftType string

// Enum values for ShiftType
const (
	ShiftTypeRegular    ShiftType = "REGULAR"
	ShiftTypeOverridden ShiftType = "OVERRIDDEN"
)

// Values returns all known values for ShiftType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ShiftType) Values() []ShiftType {
	return []ShiftType{
		"REGULAR",
		"OVERRIDDEN",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "UNKNOWN_OPERATION"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "CANNOT_PARSE"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "OTHER"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"UNKNOWN_OPERATION",
		"CANNOT_PARSE",
		"FIELD_VALIDATION_FAILED",
		"OTHER",
	}
}
