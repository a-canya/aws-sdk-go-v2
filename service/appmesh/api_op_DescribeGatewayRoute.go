// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes an existing gateway route.
func (c *Client) DescribeGatewayRoute(ctx context.Context, params *DescribeGatewayRouteInput, optFns ...func(*Options)) (*DescribeGatewayRouteOutput, error) {
	if params == nil {
		params = &DescribeGatewayRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGatewayRoute", params, optFns, c.addOperationDescribeGatewayRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGatewayRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGatewayRouteInput struct {

	// The name of the gateway route to describe.
	//
	// This member is required.
	GatewayRouteName *string

	// The name of the service mesh that the gateway route resides in.
	//
	// This member is required.
	MeshName *string

	// The name of the virtual gateway that the gateway route is associated with.
	//
	// This member is required.
	VirtualGatewayName *string

	// The Amazon Web Services IAM account ID of the service mesh owner. If the
	// account ID is not your own, then it's the ID of the account that shared the mesh
	// with your account. For more information about mesh sharing, see Working with
	// shared meshes (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html)
	// .
	MeshOwner *string

	noSmithyDocumentSerde
}

type DescribeGatewayRouteOutput struct {

	// The full description of your gateway route.
	//
	// This member is required.
	GatewayRoute *types.GatewayRouteData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeGatewayRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeGatewayRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeGatewayRoute{}, middleware.After)
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
	if err = addOpDescribeGatewayRouteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGatewayRoute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeGatewayRoute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appmesh",
		OperationName: "DescribeGatewayRoute",
	}
}
