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

// Detects Personal Protective Equipment (PPE) worn by people detected in an
// image. Amazon Rekognition can detect the following types of PPE.
//   - Face cover
//   - Hand cover
//   - Head cover
//
// You pass the input image as base64-encoded image bytes or as a reference to an
// image in an Amazon S3 bucket. The image must be either a PNG or JPG formatted
// file. DetectProtectiveEquipment detects PPE worn by up to 15 persons detected
// in an image. For each person detected in the image the API returns an array of
// body parts (face, head, left-hand, right-hand). For each body part, an array of
// detected items of PPE is returned, including an indicator of whether or not the
// PPE covers the body part. The API returns the confidence it has in each
// detection (person, PPE, body part and body part coverage). It also returns a
// bounding box ( BoundingBox ) for each detected person and each detected item of
// PPE. You can optionally request a summary of detected PPE items with the
// SummarizationAttributes input parameter. The summary provides the following
// information.
//   - The persons detected as wearing all of the types of PPE that you specify.
//   - The persons detected as not wearing all of the types PPE that you specify.
//   - The persons detected where PPE adornment could not be determined.
//
// This is a stateless API operation. That is, the operation does not persist any
// data. This operation requires permissions to perform the
// rekognition:DetectProtectiveEquipment action.
func (c *Client) DetectProtectiveEquipment(ctx context.Context, params *DetectProtectiveEquipmentInput, optFns ...func(*Options)) (*DetectProtectiveEquipmentOutput, error) {
	if params == nil {
		params = &DetectProtectiveEquipmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectProtectiveEquipment", params, optFns, c.addOperationDetectProtectiveEquipmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectProtectiveEquipmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectProtectiveEquipmentInput struct {

	// The image in which you want to detect PPE on detected persons. The image can be
	// passed as image bytes or you can reference an image stored in an Amazon S3
	// bucket.
	//
	// This member is required.
	Image *types.Image

	// An array of PPE types that you want to summarize.
	SummarizationAttributes *types.ProtectiveEquipmentSummarizationAttributes

	noSmithyDocumentSerde
}

type DetectProtectiveEquipmentOutput struct {

	// An array of persons detected in the image (including persons not wearing PPE).
	Persons []types.ProtectiveEquipmentPerson

	// The version number of the PPE detection model used to detect PPE in the image.
	ProtectiveEquipmentModelVersion *string

	// Summary information for the types of PPE specified in the
	// SummarizationAttributes input parameter.
	Summary *types.ProtectiveEquipmentSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectProtectiveEquipmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectProtectiveEquipment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectProtectiveEquipment{}, middleware.After)
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
	if err = addOpDetectProtectiveEquipmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectProtectiveEquipment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectProtectiveEquipment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DetectProtectiveEquipment",
	}
}
