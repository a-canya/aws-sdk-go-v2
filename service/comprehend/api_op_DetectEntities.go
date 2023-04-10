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

// Detects named entities in input text when you use the pre-trained model.
// Detects custom entities if you have a custom entity recognition model. When
// detecting named entities using the pre-trained model, use plain text as the
// input. For more information about named entities, see Entities (https://docs.aws.amazon.com/comprehend/latest/dg/how-entities.html)
// in the Comprehend Developer Guide. When you use a custom entity recognition
// model, you can input plain text or you can upload a single-page input document
// (text, PDF, Word, or image). If the system detects errors while processing a
// page in the input document, the API response includes an entry in Errors for
// each error. If the system detects a document-level error in your input document,
// the API returns an InvalidRequestException error response. For details about
// this exception, see Errors in semi-structured documents (https://docs.aws.amazon.com/comprehend/latest/dg/idp-inputs-sync-err.html)
// in the Comprehend Developer Guide.
func (c *Client) DetectEntities(ctx context.Context, params *DetectEntitiesInput, optFns ...func(*Options)) (*DetectEntitiesOutput, error) {
	if params == nil {
		params = &DetectEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectEntities", params, optFns, c.addOperationDetectEntitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectEntitiesInput struct {

	// This field applies only when you use a custom entity recognition model that was
	// trained with PDF annotations. For other cases, enter your text input in the Text
	// field. Use the Bytes parameter to input a text, PDF, Word or image file. Using
	// a plain-text file in the Bytes parameter is equivelent to using the Text
	// parameter (the Entities field in the response is identical). You can also use
	// the Bytes parameter to input an Amazon Textract DetectDocumentText or
	// AnalyzeDocument output file. Provide the input document as a sequence of
	// base64-encoded bytes. If your code uses an Amazon Web Services SDK to detect
	// entities, the SDK may encode the document file bytes for you. The maximum length
	// of this field depends on the input document type. For details, see Inputs for
	// real-time custom analysis (https://docs.aws.amazon.com/comprehend/latest/dg/idp-inputs-sync.html)
	// in the Comprehend Developer Guide. If you use the Bytes parameter, do not use
	// the Text parameter.
	Bytes []byte

	// Provides configuration parameters to override the default actions for
	// extracting text from PDF documents and image files.
	DocumentReaderConfig *types.DocumentReaderConfig

	// The Amazon Resource Name of an endpoint that is associated with a custom entity
	// recognition model. Provide an endpoint if you want to detect entities by using
	// your own custom model instead of the default model that is used by Amazon
	// Comprehend. If you specify an endpoint, Amazon Comprehend uses the language of
	// your custom model, and it ignores any language code that you provide in your
	// request. For information about endpoints, see Managing endpoints (https://docs.aws.amazon.com/comprehend/latest/dg/manage-endpoints.html)
	// .
	EndpointArn *string

	// The language of the input documents. You can specify any of the primary
	// languages supported by Amazon Comprehend. If your request includes the endpoint
	// for a custom entity recognition model, Amazon Comprehend uses the language of
	// your custom model, and it ignores any language code that you specify here. All
	// input documents must be in the same language.
	LanguageCode types.LanguageCode

	// A UTF-8 text string. The maximum string size is 100 KB. If you enter text using
	// this parameter, do not use the Bytes parameter.
	Text *string

	noSmithyDocumentSerde
}

type DetectEntitiesOutput struct {

	// Information about each block of text in the input document. Blocks are nested.
	// A page block contains a block for each line of text, which contains a block for
	// each word. The Block content for a Word input document does not include a
	// Geometry field. The Block field is not present in the response for plain-text
	// inputs.
	Blocks []types.Block

	// Information about the document, discovered during text extraction. This field
	// is present in the response only if your request used the Byte parameter.
	DocumentMetadata *types.DocumentMetadata

	// The document type for each page in the input document. This field is present in
	// the response only if your request used the Byte parameter.
	DocumentType []types.DocumentTypeListItem

	// A collection of entities identified in the input text. For each entity, the
	// response provides the entity text, entity type, where the entity text begins and
	// ends, and the level of confidence that Amazon Comprehend has in the detection.
	// If your request uses a custom entity recognition model, Amazon Comprehend
	// detects the entities that the model is trained to recognize. Otherwise, it
	// detects the default entity types. For a list of default entity types, see
	// Entities (https://docs.aws.amazon.com/comprehend/latest/dg/how-entities.html) in
	// the Comprehend Developer Guide.
	Entities []types.Entity

	// Page-level errors that the system detected while processing the input document.
	// The field is empty if the system encountered no errors.
	Errors []types.ErrorsListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectEntitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectEntities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectEntities{}, middleware.After)
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
	if err = addOpDetectEntitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectEntities(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectEntities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DetectEntities",
	}
}
