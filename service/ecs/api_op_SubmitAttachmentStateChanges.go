// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action is only used by the Amazon ECS agent, and it is not intended for
// use outside of the agent. Sent to acknowledge that an attachment changed states.
func (c *Client) SubmitAttachmentStateChanges(ctx context.Context, params *SubmitAttachmentStateChangesInput, optFns ...func(*Options)) (*SubmitAttachmentStateChangesOutput, error) {
	if params == nil {
		params = &SubmitAttachmentStateChangesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SubmitAttachmentStateChanges", params, optFns, c.addOperationSubmitAttachmentStateChangesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SubmitAttachmentStateChangesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SubmitAttachmentStateChangesInput struct {

	// Any attachments associated with the state change request.
	//
	// This member is required.
	Attachments []types.AttachmentStateChange

	// The short name or full ARN of the cluster that hosts the container instance the
	// attachment belongs to.
	Cluster *string

	noSmithyDocumentSerde
}

type SubmitAttachmentStateChangesOutput struct {

	// Acknowledgement of the state change.
	Acknowledgment *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSubmitAttachmentStateChangesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSubmitAttachmentStateChanges{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSubmitAttachmentStateChanges{}, middleware.After)
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
	if err = addOpSubmitAttachmentStateChangesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSubmitAttachmentStateChanges(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSubmitAttachmentStateChanges(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "SubmitAttachmentStateChanges",
	}
}
