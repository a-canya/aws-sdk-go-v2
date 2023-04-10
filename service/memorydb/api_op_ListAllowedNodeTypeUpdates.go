// Code generated by smithy-go-codegen DO NOT EDIT.

package memorydb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all available node types that you can scale to from your cluster's
// current node type. When you use the UpdateCluster operation to scale your
// cluster, the value of the NodeType parameter must be one of the node types
// returned by this operation.
func (c *Client) ListAllowedNodeTypeUpdates(ctx context.Context, params *ListAllowedNodeTypeUpdatesInput, optFns ...func(*Options)) (*ListAllowedNodeTypeUpdatesOutput, error) {
	if params == nil {
		params = &ListAllowedNodeTypeUpdatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAllowedNodeTypeUpdates", params, optFns, c.addOperationListAllowedNodeTypeUpdatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAllowedNodeTypeUpdatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAllowedNodeTypeUpdatesInput struct {

	// The name of the cluster you want to scale. MemoryDB uses the cluster name to
	// identify the current node type being used by this cluster, and from that to
	// create a list of node types you can scale up to.
	//
	// This member is required.
	ClusterName *string

	noSmithyDocumentSerde
}

type ListAllowedNodeTypeUpdatesOutput struct {

	// A list node types which you can use to scale down your cluster.
	ScaleDownNodeTypes []string

	// A list node types which you can use to scale up your cluster.
	ScaleUpNodeTypes []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAllowedNodeTypeUpdatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAllowedNodeTypeUpdates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAllowedNodeTypeUpdates{}, middleware.After)
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
	if err = addOpListAllowedNodeTypeUpdatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAllowedNodeTypeUpdates(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListAllowedNodeTypeUpdates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "memorydb",
		OperationName: "ListAllowedNodeTypeUpdates",
	}
}
