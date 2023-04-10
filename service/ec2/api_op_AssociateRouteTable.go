// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a subnet in your VPC or an internet gateway or virtual private
// gateway attached to your VPC with a route table in your VPC. This association
// causes traffic from the subnet or gateway to be routed according to the routes
// in the route table. The action returns an association ID, which you need in
// order to disassociate the route table later. A route table can be associated
// with multiple subnets. For more information, see Route tables (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html)
// in the Amazon Virtual Private Cloud User Guide.
func (c *Client) AssociateRouteTable(ctx context.Context, params *AssociateRouteTableInput, optFns ...func(*Options)) (*AssociateRouteTableOutput, error) {
	if params == nil {
		params = &AssociateRouteTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateRouteTable", params, optFns, c.addOperationAssociateRouteTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateRouteTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateRouteTableInput struct {

	// The ID of the route table.
	//
	// This member is required.
	RouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The ID of the internet gateway or virtual private gateway.
	GatewayId *string

	// The ID of the subnet.
	SubnetId *string

	noSmithyDocumentSerde
}

type AssociateRouteTableOutput struct {

	// The route table association ID. This ID is required for disassociating the
	// route table.
	AssociationId *string

	// The state of the association.
	AssociationState *types.RouteTableAssociationState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateRouteTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAssociateRouteTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAssociateRouteTable{}, middleware.After)
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
	if err = addOpAssociateRouteTableValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateRouteTable(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateRouteTable(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AssociateRouteTable",
	}
}
