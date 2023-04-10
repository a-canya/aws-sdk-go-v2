// Code generated by smithy-go-codegen DO NOT EDIT.

package pi

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve the metadata for different features. For example, the metadata might
// indicate that a feature is turned on or off on a specific DB instance.
func (c *Client) GetResourceMetadata(ctx context.Context, params *GetResourceMetadataInput, optFns ...func(*Options)) (*GetResourceMetadataOutput, error) {
	if params == nil {
		params = &GetResourceMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourceMetadata", params, optFns, c.addOperationGetResourceMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourceMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourceMetadataInput struct {

	// An immutable identifier for a data source that is unique for an Amazon Web
	// Services Region. Performance Insights gathers metrics from this data source. To
	// use a DB instance as a data source, specify its DbiResourceId value. For
	// example, specify db-ABCDEFGHIJKLMNOPQRSTU1VW2X .
	//
	// This member is required.
	Identifier *string

	// The Amazon Web Services service for which Performance Insights returns metrics.
	//
	// This member is required.
	ServiceType types.ServiceType

	noSmithyDocumentSerde
}

type GetResourceMetadataOutput struct {

	// The metadata for different features. For example, the metadata might indicate
	// that a feature is turned on or off on a specific DB instance.
	Features map[string]types.FeatureMetadata

	// An immutable identifier for a data source that is unique for an Amazon Web
	// Services Region. Performance Insights gathers metrics from this data source. To
	// use a DB instance as a data source, specify its DbiResourceId value. For
	// example, specify db-ABCDEFGHIJKLMNOPQRSTU1VW2X .
	Identifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourceMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetResourceMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetResourceMetadata{}, middleware.After)
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
	if err = addOpGetResourceMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourceMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetResourceMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "pi",
		OperationName: "GetResourceMetadata",
	}
}
