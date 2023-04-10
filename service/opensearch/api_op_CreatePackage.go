// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a package for use with Amazon OpenSearch Service domains. For more
// information, see Custom packages for Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/custom-packages.html)
// .
func (c *Client) CreatePackage(ctx context.Context, params *CreatePackageInput, optFns ...func(*Options)) (*CreatePackageOutput, error) {
	if params == nil {
		params = &CreatePackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePackage", params, optFns, c.addOperationCreatePackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to the CreatePackage operation.
type CreatePackageInput struct {

	// Unique name for the package.
	//
	// This member is required.
	PackageName *string

	// The Amazon S3 location from which to import the package.
	//
	// This member is required.
	PackageSource *types.PackageSource

	// Type of package.
	//
	// This member is required.
	PackageType types.PackageType

	// Description of the package.
	PackageDescription *string

	noSmithyDocumentSerde
}

// Container for the response returned by the CreatePackage operation.
type CreatePackageOutput struct {

	// Basic information about an OpenSearch Service package.
	PackageDetails *types.PackageDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePackage{}, middleware.After)
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
	if err = addOpCreatePackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "CreatePackage",
	}
}
