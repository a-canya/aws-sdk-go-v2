// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the refresh properties of a dataset.
func (c *Client) DescribeDataSetRefreshProperties(ctx context.Context, params *DescribeDataSetRefreshPropertiesInput, optFns ...func(*Options)) (*DescribeDataSetRefreshPropertiesOutput, error) {
	if params == nil {
		params = &DescribeDataSetRefreshPropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDataSetRefreshProperties", params, optFns, c.addOperationDescribeDataSetRefreshPropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDataSetRefreshPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDataSetRefreshPropertiesInput struct {

	// The Amazon Web Services account ID.
	//
	// This member is required.
	AwsAccountId *string

	// The ID of the dataset.
	//
	// This member is required.
	DataSetId *string

	noSmithyDocumentSerde
}

type DescribeDataSetRefreshPropertiesOutput struct {

	// The dataset refresh properties.
	DataSetRefreshProperties *types.DataSetRefreshProperties

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDataSetRefreshPropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDataSetRefreshProperties{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDataSetRefreshProperties{}, middleware.After)
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
	if err = addOpDescribeDataSetRefreshPropertiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDataSetRefreshProperties(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDataSetRefreshProperties(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeDataSetRefreshProperties",
	}
}
