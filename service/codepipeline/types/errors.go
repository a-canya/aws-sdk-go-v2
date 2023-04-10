// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The specified action cannot be found.
type ActionNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ActionNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ActionNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ActionNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ActionNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ActionNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified action type cannot be found.
type ActionTypeNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ActionTypeNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ActionTypeNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ActionTypeNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ActionTypeNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ActionTypeNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The approval action has already been approved or rejected.
type ApprovalAlreadyCompletedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ApprovalAlreadyCompletedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ApprovalAlreadyCompletedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ApprovalAlreadyCompletedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ApprovalAlreadyCompletedException"
	}
	return *e.ErrorCodeOverride
}
func (e *ApprovalAlreadyCompletedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Unable to modify the tag due to a simultaneous update request.
type ConcurrentModificationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConcurrentModificationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Your request cannot be handled because the pipeline is busy handling ongoing
// activities. Try again later.
type ConflictException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConflictException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The pipeline execution is already in a Stopping state. If you already chose to
// stop and wait, you cannot make that request again. You can choose to stop and
// abandon now, but be aware that this option can lead to failed tasks or out of
// sequence tasks. If you already chose to stop and abandon, you cannot make that
// request again.
type DuplicatedStopRequestException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DuplicatedStopRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicatedStopRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicatedStopRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DuplicatedStopRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *DuplicatedStopRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The action declaration was specified in an invalid format.
type InvalidActionDeclarationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidActionDeclarationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidActionDeclarationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidActionDeclarationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidActionDeclarationException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidActionDeclarationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The approval request already received a response or has expired.
type InvalidApprovalTokenException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidApprovalTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidApprovalTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidApprovalTokenException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidApprovalTokenException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidApprovalTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource ARN is invalid.
type InvalidArnException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidArnException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArnException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArnException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidArnException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidArnException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Reserved for future use.
type InvalidBlockerDeclarationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidBlockerDeclarationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidBlockerDeclarationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidBlockerDeclarationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidBlockerDeclarationException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidBlockerDeclarationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The client token was specified in an invalid format
type InvalidClientTokenException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidClientTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidClientTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidClientTokenException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidClientTokenException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidClientTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The job was specified in an invalid format or cannot be found.
type InvalidJobException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidJobException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidJobException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidJobException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidJobException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidJobException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The job state was specified in an invalid format.
type InvalidJobStateException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidJobStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidJobStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidJobStateException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidJobStateException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidJobStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The next token was specified in an invalid format. Make sure that the next
// token you provide is the token returned by a previous call.
type InvalidNextTokenException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidNextTokenException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The nonce was specified in an invalid format.
type InvalidNonceException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidNonceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNonceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNonceException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidNonceException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidNonceException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The stage declaration was specified in an invalid format.
type InvalidStageDeclarationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidStageDeclarationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidStageDeclarationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidStageDeclarationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidStageDeclarationException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidStageDeclarationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The structure was specified in an invalid format.
type InvalidStructureException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidStructureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidStructureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidStructureException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidStructureException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidStructureException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource tags are invalid.
type InvalidTagsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTagsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidTagsException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified authentication type is in an invalid format.
type InvalidWebhookAuthenticationParametersException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidWebhookAuthenticationParametersException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidWebhookAuthenticationParametersException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidWebhookAuthenticationParametersException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidWebhookAuthenticationParametersException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidWebhookAuthenticationParametersException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified event filter rule is in an invalid format.
type InvalidWebhookFilterPatternException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidWebhookFilterPatternException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidWebhookFilterPatternException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidWebhookFilterPatternException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidWebhookFilterPatternException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidWebhookFilterPatternException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The job was specified in an invalid format or cannot be found.
type JobNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *JobNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *JobNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *JobNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "JobNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *JobNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The number of pipelines associated with the AWS account has exceeded the limit
// allowed for the account.
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The stage has failed in a later run of the pipeline and the pipelineExecutionId
// associated with the request is out of date.
type NotLatestPipelineExecutionException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NotLatestPipelineExecutionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotLatestPipelineExecutionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotLatestPipelineExecutionException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NotLatestPipelineExecutionException"
	}
	return *e.ErrorCodeOverride
}
func (e *NotLatestPipelineExecutionException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Exceeded the total size limit for all variables in the pipeline.
type OutputVariablesSizeExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *OutputVariablesSizeExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OutputVariablesSizeExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OutputVariablesSizeExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "OutputVariablesSizeExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *OutputVariablesSizeExceededException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The pipeline execution was specified in an invalid format or cannot be found,
// or an execution ID does not belong to the specified pipeline.
type PipelineExecutionNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *PipelineExecutionNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PipelineExecutionNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PipelineExecutionNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PipelineExecutionNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *PipelineExecutionNotFoundException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Unable to stop the pipeline execution. The execution might already be in a
// Stopped state, or it might no longer be in progress.
type PipelineExecutionNotStoppableException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *PipelineExecutionNotStoppableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PipelineExecutionNotStoppableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PipelineExecutionNotStoppableException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PipelineExecutionNotStoppableException"
	}
	return *e.ErrorCodeOverride
}
func (e *PipelineExecutionNotStoppableException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified pipeline name is already in use.
type PipelineNameInUseException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *PipelineNameInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PipelineNameInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PipelineNameInUseException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PipelineNameInUseException"
	}
	return *e.ErrorCodeOverride
}
func (e *PipelineNameInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The pipeline was specified in an invalid format or cannot be found.
type PipelineNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *PipelineNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PipelineNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PipelineNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PipelineNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *PipelineNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The pipeline version was specified in an invalid format or cannot be found.
type PipelineVersionNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *PipelineVersionNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PipelineVersionNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PipelineVersionNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PipelineVersionNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *PipelineVersionNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request failed because of an unknown error, exception, or failure.
type RequestFailedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *RequestFailedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RequestFailedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RequestFailedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RequestFailedException"
	}
	return *e.ErrorCodeOverride
}
func (e *RequestFailedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource was specified in an invalid format.
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

// The stage was specified in an invalid format or cannot be found.
type StageNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *StageNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StageNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StageNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "StageNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *StageNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Unable to retry. The pipeline structure or stage state might have changed while
// actions awaited retry, or the stage contains no failed actions.
type StageNotRetryableException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *StageNotRetryableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StageNotRetryableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StageNotRetryableException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "StageNotRetryableException"
	}
	return *e.ErrorCodeOverride
}
func (e *StageNotRetryableException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The tags limit for a resource has been exceeded.
type TooManyTagsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyTagsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The validation was specified in an invalid format.
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

// The specified webhook was entered in an invalid format or cannot be found.
type WebhookNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WebhookNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WebhookNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WebhookNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WebhookNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *WebhookNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
