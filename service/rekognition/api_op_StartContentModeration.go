// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts asynchronous detection of inappropriate, unwanted, or offensive content
// in a stored video. For a list of moderation labels in Amazon Rekognition, see
// Using the image and video moderation APIs (https://docs.aws.amazon.com/rekognition/latest/dg/moderation.html#moderation-api)
// . Amazon Rekognition Video can moderate content in a video stored in an Amazon
// S3 bucket. Use Video to specify the bucket name and the filename of the video.
// StartContentModeration returns a job identifier ( JobId ) which you use to get
// the results of the analysis. When content analysis is finished, Amazon
// Rekognition Video publishes a completion status to the Amazon Simple
// Notification Service topic that you specify in NotificationChannel . To get the
// results of the content analysis, first check that the status value published to
// the Amazon SNS topic is SUCCEEDED . If so, call GetContentModeration and pass
// the job identifier ( JobId ) from the initial call to StartContentModeration .
// For more information, see Moderating content in the Amazon Rekognition Developer
// Guide.
func (c *Client) StartContentModeration(ctx context.Context, params *StartContentModerationInput, optFns ...func(*Options)) (*StartContentModerationOutput, error) {
	if params == nil {
		params = &StartContentModerationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartContentModeration", params, optFns, c.addOperationStartContentModerationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartContentModerationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartContentModerationInput struct {

	// The video in which you want to detect inappropriate, unwanted, or offensive
	// content. The video must be stored in an Amazon S3 bucket.
	//
	// This member is required.
	Video *types.Video

	// Idempotent token used to identify the start request. If you use the same token
	// with multiple StartContentModeration requests, the same JobId is returned. Use
	// ClientRequestToken to prevent the same job from being accidently started more
	// than once.
	ClientRequestToken *string

	// An identifier you specify that's returned in the completion notification that's
	// published to your Amazon Simple Notification Service topic. For example, you can
	// use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string

	// Specifies the minimum confidence that Amazon Rekognition must have in order to
	// return a moderated content label. Confidence represents how certain Amazon
	// Rekognition is that the moderated content is correctly identified. 0 is the
	// lowest confidence. 100 is the highest confidence. Amazon Rekognition doesn't
	// return any moderated content labels with a confidence level lower than this
	// specified value. If you don't specify MinConfidence , GetContentModeration
	// returns labels with confidence values greater than or equal to 50 percent.
	MinConfidence *float32

	// The Amazon SNS topic ARN that you want Amazon Rekognition Video to publish the
	// completion status of the content analysis to. The Amazon SNS topic must have a
	// topic name that begins with AmazonRekognition if you are using the
	// AmazonRekognitionServiceRole permissions policy to access the topic.
	NotificationChannel *types.NotificationChannel

	noSmithyDocumentSerde
}

type StartContentModerationOutput struct {

	// The identifier for the content analysis job. Use JobId to identify the job in a
	// subsequent call to GetContentModeration .
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartContentModerationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartContentModeration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartContentModeration{}, middleware.After)
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
	if err = addOpStartContentModerationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartContentModeration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartContentModeration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "StartContentModeration",
	}
}
