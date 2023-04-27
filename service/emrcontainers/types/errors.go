// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// This is an internal server exception.
type InternalServerException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServerException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The request throttled.
type RequestThrottledException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *RequestThrottledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RequestThrottledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RequestThrottledException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RequestThrottledException"
	}
	return *e.ErrorCodeOverride
}
func (e *RequestThrottledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource was not found.
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

// There are invalid parameters in the client request.
type ValidationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ValidationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
