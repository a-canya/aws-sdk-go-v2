// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a share request for a custom framework in Audit Manager. The share
// request specifies a recipient and notifies them that a custom framework is
// available. Recipients have 120 days to accept or decline the request. If no
// action is taken, the share request expires. When you create a share request,
// Audit Manager stores a snapshot of your custom framework in the US East (N.
// Virginia) Amazon Web Services Region. Audit Manager also stores a backup of the
// same snapshot in the US West (Oregon) Amazon Web Services Region. Audit Manager
// deletes the snapshot and the backup snapshot when one of the following events
// occurs:
//   - The sender revokes the share request.
//   - The recipient declines the share request.
//   - The recipient encounters an error and doesn't successfully accept the share
//     request.
//   - The share request expires before the recipient responds to the request.
//
// When a sender resends a share request (https://docs.aws.amazon.com/audit-manager/latest/userguide/framework-sharing.html#framework-sharing-resend)
// , the snapshot is replaced with an updated version that corresponds with the
// latest version of the custom framework. When a recipient accepts a share
// request, the snapshot is replicated into their Amazon Web Services account under
// the Amazon Web Services Region that was specified in the share request. When you
// invoke the StartAssessmentFrameworkShare API, you are about to share a custom
// framework with another Amazon Web Services account. You may not share a custom
// framework that is derived from a standard framework if the standard framework is
// designated as not eligible for sharing by Amazon Web Services, unless you have
// obtained permission to do so from the owner of the standard framework. To learn
// more about which standard frameworks are eligible for sharing, see Framework
// sharing eligibility (https://docs.aws.amazon.com/audit-manager/latest/userguide/share-custom-framework-concepts-and-terminology.html#eligibility)
// in the Audit Manager User Guide.
func (c *Client) StartAssessmentFrameworkShare(ctx context.Context, params *StartAssessmentFrameworkShareInput, optFns ...func(*Options)) (*StartAssessmentFrameworkShareOutput, error) {
	if params == nil {
		params = &StartAssessmentFrameworkShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartAssessmentFrameworkShare", params, optFns, c.addOperationStartAssessmentFrameworkShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartAssessmentFrameworkShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartAssessmentFrameworkShareInput struct {

	// The Amazon Web Services account of the recipient.
	//
	// This member is required.
	DestinationAccount *string

	// The Amazon Web Services Region of the recipient.
	//
	// This member is required.
	DestinationRegion *string

	// The unique identifier for the custom framework to be shared.
	//
	// This member is required.
	FrameworkId *string

	// An optional comment from the sender about the share request.
	Comment *string

	noSmithyDocumentSerde
}

type StartAssessmentFrameworkShareOutput struct {

	// The share request that's created by the StartAssessmentFrameworkShare API.
	AssessmentFrameworkShareRequest *types.AssessmentFrameworkShareRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartAssessmentFrameworkShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartAssessmentFrameworkShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartAssessmentFrameworkShare{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpStartAssessmentFrameworkShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartAssessmentFrameworkShare(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opStartAssessmentFrameworkShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "StartAssessmentFrameworkShare",
	}
}
