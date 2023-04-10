// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Exports a running or stopped instance to an Amazon S3 bucket. For information
// about the supported operating systems, image formats, and known limitations for
// the types of instances you can export, see Exporting an instance as a VM Using
// VM Import/Export (https://docs.aws.amazon.com/vm-import/latest/userguide/vmexport.html)
// in the VM Import/Export User Guide.
func (c *Client) CreateInstanceExportTask(ctx context.Context, params *CreateInstanceExportTaskInput, optFns ...func(*Options)) (*CreateInstanceExportTaskOutput, error) {
	if params == nil {
		params = &CreateInstanceExportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInstanceExportTask", params, optFns, c.addOperationCreateInstanceExportTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInstanceExportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInstanceExportTaskInput struct {

	// The format and location for an export instance task.
	//
	// This member is required.
	ExportToS3Task *types.ExportToS3TaskSpecification

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// The target virtualization environment.
	//
	// This member is required.
	TargetEnvironment types.ExportEnvironment

	// A description for the conversion task or the resource being exported. The
	// maximum length is 255 characters.
	Description *string

	// The tags to apply to the export instance task during creation.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateInstanceExportTaskOutput struct {

	// Information about the export instance task.
	ExportTask *types.ExportTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateInstanceExportTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateInstanceExportTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateInstanceExportTask{}, middleware.After)
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
	if err = addOpCreateInstanceExportTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInstanceExportTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateInstanceExportTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateInstanceExportTask",
	}
}
