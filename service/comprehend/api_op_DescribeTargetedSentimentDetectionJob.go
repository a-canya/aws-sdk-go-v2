// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the properties associated with a targeted sentiment detection job. Use
// this operation to get the status of the job.
func (c *Client) DescribeTargetedSentimentDetectionJob(ctx context.Context, params *DescribeTargetedSentimentDetectionJobInput, optFns ...func(*Options)) (*DescribeTargetedSentimentDetectionJobOutput, error) {
	if params == nil {
		params = &DescribeTargetedSentimentDetectionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTargetedSentimentDetectionJob", params, optFns, c.addOperationDescribeTargetedSentimentDetectionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTargetedSentimentDetectionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTargetedSentimentDetectionJobInput struct {

	// The identifier that Amazon Comprehend generated for the job. The
	// StartTargetedSentimentDetectionJob operation returns this identifier in its
	// response.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type DescribeTargetedSentimentDetectionJobOutput struct {

	// An object that contains the properties associated with a targeted sentiment
	// detection job.
	TargetedSentimentDetectionJobProperties *types.TargetedSentimentDetectionJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTargetedSentimentDetectionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTargetedSentimentDetectionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTargetedSentimentDetectionJob{}, middleware.After)
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
	if err = addOpDescribeTargetedSentimentDetectionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTargetedSentimentDetectionJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTargetedSentimentDetectionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DescribeTargetedSentimentDetectionJob",
	}
}
