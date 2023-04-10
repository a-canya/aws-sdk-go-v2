// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information on VPC peering connections. Use this operation to get
// peering information for all fleets or for one specific fleet ID. To retrieve
// connection information, call this operation from the Amazon Web Services account
// that is used to manage the Amazon GameLift fleets. Specify a fleet ID or leave
// the parameter empty to retrieve all connection records. If successful, the
// retrieved information includes both active and pending connections. Active
// connections identify the IpV4 CIDR block that the VPC uses to connect. Related
// actions All APIs by task (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) DescribeVpcPeeringConnections(ctx context.Context, params *DescribeVpcPeeringConnectionsInput, optFns ...func(*Options)) (*DescribeVpcPeeringConnectionsOutput, error) {
	if params == nil {
		params = &DescribeVpcPeeringConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcPeeringConnections", params, optFns, c.addOperationDescribeVpcPeeringConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcPeeringConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcPeeringConnectionsInput struct {

	// A unique identifier for the fleet. You can use either the fleet ID or ARN value.
	FleetId *string

	noSmithyDocumentSerde
}

type DescribeVpcPeeringConnectionsOutput struct {

	// A collection of VPC peering connection records that match the request.
	VpcPeeringConnections []types.VpcPeeringConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVpcPeeringConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeVpcPeeringConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeVpcPeeringConnections{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcPeeringConnections(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeVpcPeeringConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeVpcPeeringConnections",
	}
}
