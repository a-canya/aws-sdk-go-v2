// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copies the input published schema, at the specified version, into the Directory
// with the same name and version as that of the published schema.
func (c *Client) ApplySchema(ctx context.Context, params *ApplySchemaInput, optFns ...func(*Options)) (*ApplySchemaOutput, error) {
	if params == nil {
		params = &ApplySchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ApplySchema", params, optFns, c.addOperationApplySchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ApplySchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ApplySchemaInput struct {

	// The Amazon Resource Name (ARN) that is associated with the Directory into which
	// the schema is copied. For more information, see arns .
	//
	// This member is required.
	DirectoryArn *string

	// Published schema Amazon Resource Name (ARN) that needs to be copied. For more
	// information, see arns .
	//
	// This member is required.
	PublishedSchemaArn *string

	noSmithyDocumentSerde
}

type ApplySchemaOutput struct {

	// The applied schema ARN that is associated with the copied schema in the
	// Directory . You can use this ARN to describe the schema information applied on
	// this directory. For more information, see arns .
	AppliedSchemaArn *string

	// The ARN that is associated with the Directory . For more information, see arns .
	DirectoryArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationApplySchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpApplySchema{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpApplySchema{}, middleware.After)
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
	if err = addOpApplySchemaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opApplySchema(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opApplySchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ApplySchema",
	}
}
