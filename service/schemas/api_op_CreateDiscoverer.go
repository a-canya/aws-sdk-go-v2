// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/schemas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a discoverer.
func (c *Client) CreateDiscoverer(ctx context.Context, params *CreateDiscovererInput, optFns ...func(*Options)) (*CreateDiscovererOutput, error) {
	if params == nil {
		params = &CreateDiscovererInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDiscoverer", params, optFns, c.addOperationCreateDiscovererMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDiscovererOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDiscovererInput struct {

	// The ARN of the event bus.
	//
	// This member is required.
	SourceArn *string

	// Support discovery of schemas in events sent to the bus from another account.
	// (default: true).
	CrossAccount bool

	// A description for the discoverer.
	Description *string

	// Tags associated with the resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateDiscovererOutput struct {

	// The Status if the discoverer will discover schemas from events sent from
	// another account.
	CrossAccount bool

	// The description of the discoverer.
	Description *string

	// The ARN of the discoverer.
	DiscovererArn *string

	// The ID of the discoverer.
	DiscovererId *string

	// The ARN of the event bus.
	SourceArn *string

	// The state of the discoverer.
	State types.DiscovererState

	// Tags associated with the resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDiscovererMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDiscoverer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDiscoverer{}, middleware.After)
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
	if err = addOpCreateDiscovererValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDiscoverer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDiscoverer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "CreateDiscoverer",
	}
}
