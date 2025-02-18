// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An occurrence during a stage session.
type Event struct {

	// If the event is an error event, the error code is provided to give insight into
	// the specific error that occurred. If the event is not an error event, this field
	// is null. INSUFFICIENT_CAPABILITIES indicates that the participant tried to take
	// an action that the participant’s token is not allowed to do. For more
	// information about participant capabilities, see the capabilities field in
	// CreateParticipantToken . QUOTA_EXCEEDED indicates that the number of
	// participants who want to publish/subscribe to a stage exceeds the quota; for
	// more information, see Service Quotas (https://docs.aws.amazon.com/ivs/latest/RealTimeUserGuide/service-quotas.html)
	// . PUBLISHER_NOT_FOUND indicates that the participant tried to subscribe to a
	// publisher that doesn’t exist.
	ErrorCode EventErrorCode

	// ISO 8601 timestamp (returned as a string) for when the event occurred.
	EventTime *time.Time

	// The name of the event.
	Name EventName

	// Unique identifier for the participant who triggered the event. This is assigned
	// by IVS.
	ParticipantId *string

	// Unique identifier for the remote participant. For a subscribe event, this is
	// the publisher. For a publish or join event, this is null. This is assigned by
	// IVS.
	RemoteParticipantId *string

	noSmithyDocumentSerde
}

// Object describing a participant that has joined a stage.
type Participant struct {

	// Application-provided attributes to encode into the token and attach to a stage.
	// Map keys and values can contain UTF-8 encoded text. The maximum length of this
	// field is 1 KB total. This field is exposed to all stage participants and should
	// not be used for personally identifying, confidential, or sensitive information.
	Attributes map[string]string

	// ISO 8601 timestamp (returned as a string) when the participant first joined the
	// stage session.
	FirstJoinTime *time.Time

	// Unique identifier for this participant, assigned by IVS.
	ParticipantId *string

	// Whether the participant ever published to the stage session.
	Published bool

	// Whether the participant is connected to or disconnected from the stage.
	State ParticipantState

	// Customer-assigned name to help identify the token; this can be used to link a
	// participant to a user in the customer’s own systems. This can be any UTF-8
	// encoded text. This field is exposed to all stage participants and should not be
	// used for personally identifying, confidential, or sensitive information.
	UserId *string

	noSmithyDocumentSerde
}

// Summary object describing a participant that has joined a stage.
type ParticipantSummary struct {

	// ISO 8601 timestamp (returned as a string) when the participant first joined the
	// stage session.
	FirstJoinTime *time.Time

	// Unique identifier for this participant, assigned by IVS.
	ParticipantId *string

	// Whether the participant ever published to the stage session.
	Published bool

	// Whether the participant is connected to or disconnected from the stage.
	State ParticipantState

	// Customer-assigned name to help identify the token; this can be used to link a
	// participant to a user in the customer’s own systems. This can be any UTF-8
	// encoded text. This field is exposed to all stage participants and should not be
	// used for personally identifying, confidential, or sensitive information.
	UserId *string

	noSmithyDocumentSerde
}

// Object specifying a participant token in a stage. Important: Treat tokens as
// opaque; i.e., do not build functionality based on token contents. The format of
// tokens could change in the future.
type ParticipantToken struct {

	// Application-provided attributes to encode into the token and attach to a stage.
	// This field is exposed to all stage participants and should not be used for
	// personally identifying, confidential, or sensitive information.
	Attributes map[string]string

	// Set of capabilities that the user is allowed to perform in the stage.
	Capabilities []ParticipantTokenCapability

	// Duration (in minutes), after which the participant token expires. Default: 720
	// (12 hours).
	Duration int32

	// ISO 8601 timestamp (returned as a string) for when this token expires.
	ExpirationTime *time.Time

	// Unique identifier for this participant token, assigned by IVS.
	ParticipantId *string

	// The issued client token, encrypted.
	Token *string

	// Customer-assigned name to help identify the token; this can be used to link a
	// participant to a user in the customer’s own systems. This can be any UTF-8
	// encoded text. This field is exposed to all stage participants and should not be
	// used for personally identifying, confidential, or sensitive information.
	UserId *string

	noSmithyDocumentSerde
}

// Object specifying a participant token configuration in a stage.
type ParticipantTokenConfiguration struct {

	// Application-provided attributes to encode into the corresponding participant
	// token and attach to a stage. Map keys and values can contain UTF-8 encoded text.
	// The maximum length of this field is 1 KB total. This field is exposed to all
	// stage participants and should not be used for personally identifying,
	// confidential, or sensitive information.
	Attributes map[string]string

	// Set of capabilities that the user is allowed to perform in the stage.
	Capabilities []ParticipantTokenCapability

	// Duration (in minutes), after which the corresponding participant token expires.
	// Default: 720 (12 hours).
	Duration int32

	// Customer-assigned name to help identify the token; this can be used to link a
	// participant to a user in the customer’s own systems. This can be any UTF-8
	// encoded text. This field is exposed to all stage participants and should not be
	// used for personally identifying, confidential, or sensitive information.
	UserId *string

	noSmithyDocumentSerde
}

// Object specifying a stage.
type Stage struct {

	// Stage ARN.
	//
	// This member is required.
	Arn *string

	// ID of the active session within the stage.
	ActiveSessionId *string

	// Stage name.
	Name *string

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) . See Tagging AWS Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for details, including restrictions that apply to tags and "Tag naming limits
	// and requirements"; Amazon IVS has no constraints on tags beyond what is
	// documented there.
	Tags map[string]string

	noSmithyDocumentSerde
}

// A stage session begins when the first participant joins a stage and ends after
// the last participant leaves the stage. A stage session helps with debugging
// stages by grouping events and participants into shorter periods of time (i.e., a
// session), which is helpful when stages are used over long periods of time.
type StageSession struct {

	// ISO 8601 timestamp (returned as a string) when the stage session ended. This is
	// null if the stage is active.
	EndTime *time.Time

	// ID of the session within the stage.
	SessionId *string

	// ISO 8601 timestamp (returned as a string) when this stage session began.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// Summary information about a stage session.
type StageSessionSummary struct {

	// ISO 8601 timestamp (returned as a string) when the stage session ended. This is
	// null if the stage is active.
	EndTime *time.Time

	// ID of the session within the stage.
	SessionId *string

	// ISO 8601 timestamp (returned as a string) when this stage session began.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// Summary information about a stage.
type StageSummary struct {

	// Stage ARN.
	//
	// This member is required.
	Arn *string

	// ID of the active session within the stage.
	ActiveSessionId *string

	// Stage name.
	Name *string

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) . See Tagging AWS Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for details, including restrictions that apply to tags and "Tag naming limits
	// and requirements"; Amazon IVS has no constraints on tags beyond what is
	// documented there.
	Tags map[string]string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
