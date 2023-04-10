// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides information about an import job.
func (c *Client) GetImportJob(ctx context.Context, params *GetImportJobInput, optFns ...func(*Options)) (*GetImportJobOutput, error) {
	if params == nil {
		params = &GetImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetImportJob", params, optFns, c.addOperationGetImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request for information about an import job using the import job
// ID.
type GetImportJobInput struct {

	// The ID of the import job.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type GetImportJobOutput struct {

	// The time stamp of when the import job was completed.
	CompletedTimestamp *time.Time

	// The time stamp of when the import job was created.
	CreatedTimestamp *time.Time

	// The number of records that failed processing because of invalid input or other
	// reasons.
	FailedRecordsCount *int32

	// The failure details about an import job.
	FailureInfo *types.FailureInfo

	// The data source of the import job.
	ImportDataSource *types.ImportDataSource

	// The destination of the import job.
	ImportDestination *types.ImportDestination

	// A string that represents the import job ID.
	JobId *string

	// The status of the import job.
	JobStatus types.JobStatus

	// The current number of records processed.
	ProcessedRecordsCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetImportJob{}, middleware.After)
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
	if err = addOpGetImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetImportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetImportJob",
	}
}
