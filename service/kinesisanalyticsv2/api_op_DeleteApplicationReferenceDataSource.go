// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a reference data source configuration from the specified SQL-based
// Kinesis Data Analytics application's configuration. If the application is
// running, Kinesis Data Analytics immediately removes the in-application table
// that you created using the AddApplicationReferenceDataSource operation.
func (c *Client) DeleteApplicationReferenceDataSource(ctx context.Context, params *DeleteApplicationReferenceDataSourceInput, optFns ...func(*Options)) (*DeleteApplicationReferenceDataSourceOutput, error) {
	if params == nil {
		params = &DeleteApplicationReferenceDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteApplicationReferenceDataSource", params, optFns, c.addOperationDeleteApplicationReferenceDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteApplicationReferenceDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteApplicationReferenceDataSourceInput struct {

	// The name of an existing application.
	//
	// This member is required.
	ApplicationName *string

	// The current application version. You can use the DescribeApplication operation
	// to get the current application version. If the version specified is not the
	// current version, the ConcurrentModificationException is returned.
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// The ID of the reference data source. When you add a reference data source to
	// your application using the AddApplicationReferenceDataSource , Kinesis Data
	// Analytics assigns an ID. You can use the DescribeApplication operation to get
	// the reference ID.
	//
	// This member is required.
	ReferenceId *string

	noSmithyDocumentSerde
}

type DeleteApplicationReferenceDataSourceOutput struct {

	// The application Amazon Resource Name (ARN).
	ApplicationARN *string

	// The updated version ID of the application.
	ApplicationVersionId *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteApplicationReferenceDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteApplicationReferenceDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteApplicationReferenceDataSource{}, middleware.After)
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
	if err = addOpDeleteApplicationReferenceDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplicationReferenceDataSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteApplicationReferenceDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "DeleteApplicationReferenceDataSource",
	}
}
