// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Your request was throttled because you have exceeded the limit of allowed
// client calls. Try making the call later.
type ClientLimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ClientLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClientLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClientLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ClientLimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *ClientLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The value for this input parameter is invalid.
type InvalidArgumentException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidArgumentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArgumentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArgumentException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidArgumentException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidArgumentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified client is invalid.
type InvalidClientException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidClientException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidClientException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidClientException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidClientException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidClientException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The caller is not authorized to perform this operation.
type NotAuthorizedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NotAuthorizedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotAuthorizedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotAuthorizedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NotAuthorizedException"
	}
	return *e.ErrorCodeOverride
}
func (e *NotAuthorizedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource is not found.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// If the client session is expired. Once the client is connected, the session is
// valid for 45 minutes. Client should reconnect to the channel to continue
// sending/receiving messages.
type SessionExpiredException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SessionExpiredException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SessionExpiredException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SessionExpiredException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SessionExpiredException"
	}
	return *e.ErrorCodeOverride
}
func (e *SessionExpiredException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
