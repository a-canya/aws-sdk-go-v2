// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Add endpoints to an endpoint group. The AddEndpoints API operation is the
// recommended option for adding endpoints. The alternative options are to add
// endpoints when you create an endpoint group (with the CreateEndpointGroup (https://docs.aws.amazon.com/global-accelerator/latest/api/API_CreateEndpointGroup.html)
// API) or when you update an endpoint group (with the UpdateEndpointGroup (https://docs.aws.amazon.com/global-accelerator/latest/api/API_UpdateEndpointGroup.html)
// API). There are two advantages to using AddEndpoints to add endpoints:
//   - It's faster, because Global Accelerator only has to resolve the new
//     endpoints that you're adding.
//   - It's more convenient, because you don't need to specify all of the current
//     endpoints that are already in the endpoint group in addition to the new
//     endpoints that you want to add.
func (c *Client) AddEndpoints(ctx context.Context, params *AddEndpointsInput, optFns ...func(*Options)) (*AddEndpointsOutput, error) {
	if params == nil {
		params = &AddEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddEndpoints", params, optFns, c.addOperationAddEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddEndpointsInput struct {

	// The list of endpoint objects.
	//
	// This member is required.
	EndpointConfigurations []types.EndpointConfiguration

	// The Amazon Resource Name (ARN) of the endpoint group.
	//
	// This member is required.
	EndpointGroupArn *string

	noSmithyDocumentSerde
}

type AddEndpointsOutput struct {

	// The list of endpoint objects.
	EndpointDescriptions []types.EndpointDescription

	// The Amazon Resource Name (ARN) of the endpoint group.
	EndpointGroupArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddEndpoints{}, middleware.After)
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
	if err = addOpAddEndpointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddEndpoints(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "AddEndpoints",
	}
}
