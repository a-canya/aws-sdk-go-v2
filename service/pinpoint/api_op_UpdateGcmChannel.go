// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables the GCM channel for an application or updates the status and settings
// of the GCM channel for an application.
func (c *Client) UpdateGcmChannel(ctx context.Context, params *UpdateGcmChannelInput, optFns ...func(*Options)) (*UpdateGcmChannelOutput, error) {
	if params == nil {
		params = &UpdateGcmChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGcmChannel", params, optFns, c.addOperationUpdateGcmChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGcmChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGcmChannelInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// Specifies the status and settings of the GCM channel for an application. This
	// channel enables Amazon Pinpoint to send push notifications through the Firebase
	// Cloud Messaging (FCM), formerly Google Cloud Messaging (GCM), service.
	//
	// This member is required.
	GCMChannelRequest *types.GCMChannelRequest

	noSmithyDocumentSerde
}

type UpdateGcmChannelOutput struct {

	// Provides information about the status and settings of the GCM channel for an
	// application. The GCM channel enables Amazon Pinpoint to send push notifications
	// through the Firebase Cloud Messaging (FCM), formerly Google Cloud Messaging
	// (GCM), service.
	//
	// This member is required.
	GCMChannelResponse *types.GCMChannelResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGcmChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGcmChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGcmChannel{}, middleware.After)
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
	if err = addOpUpdateGcmChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGcmChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateGcmChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "UpdateGcmChannel",
	}
}
