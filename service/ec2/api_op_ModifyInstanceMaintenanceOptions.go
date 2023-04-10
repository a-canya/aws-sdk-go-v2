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

// Modifies the recovery behavior of your instance to disable simplified automatic
// recovery or set the recovery behavior to default. The default configuration will
// not enable simplified automatic recovery for an unsupported instance type. For
// more information, see Simplified automatic recovery (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-recover.html#instance-configuration-recovery)
// .
func (c *Client) ModifyInstanceMaintenanceOptions(ctx context.Context, params *ModifyInstanceMaintenanceOptionsInput, optFns ...func(*Options)) (*ModifyInstanceMaintenanceOptionsOutput, error) {
	if params == nil {
		params = &ModifyInstanceMaintenanceOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyInstanceMaintenanceOptions", params, optFns, c.addOperationModifyInstanceMaintenanceOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyInstanceMaintenanceOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyInstanceMaintenanceOptionsInput struct {

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// Disables the automatic recovery behavior of your instance or sets it to default.
	AutoRecovery types.InstanceAutoRecoveryState

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type ModifyInstanceMaintenanceOptionsOutput struct {

	// Provides information on the current automatic recovery behavior of your
	// instance.
	AutoRecovery types.InstanceAutoRecoveryState

	// The ID of the instance.
	InstanceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyInstanceMaintenanceOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyInstanceMaintenanceOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyInstanceMaintenanceOptions{}, middleware.After)
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
	if err = addOpModifyInstanceMaintenanceOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyInstanceMaintenanceOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyInstanceMaintenanceOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyInstanceMaintenanceOptions",
	}
}
