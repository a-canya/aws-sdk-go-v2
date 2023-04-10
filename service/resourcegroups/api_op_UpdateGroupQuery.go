// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the resource query of a group. For more information about resource
// queries, see Create a tag-based group in Resource Groups (https://docs.aws.amazon.com/ARG/latest/userguide/gettingstarted-query.html#gettingstarted-query-cli-tag)
// . Minimum permissions To run this command, you must have the following
// permissions:
//   - resource-groups:UpdateGroupQuery
func (c *Client) UpdateGroupQuery(ctx context.Context, params *UpdateGroupQueryInput, optFns ...func(*Options)) (*UpdateGroupQueryOutput, error) {
	if params == nil {
		params = &UpdateGroupQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGroupQuery", params, optFns, c.addOperationUpdateGroupQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGroupQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGroupQueryInput struct {

	// The resource query to determine which Amazon Web Services resources are members
	// of this resource group. A resource group can contain either a Configuration or
	// a ResourceQuery , but not both.
	//
	// This member is required.
	ResourceQuery *types.ResourceQuery

	// The name or the ARN of the resource group to query.
	Group *string

	// Don't use this parameter. Use Group instead.
	//
	// Deprecated: This field is deprecated, use Group instead.
	GroupName *string

	noSmithyDocumentSerde
}

type UpdateGroupQueryOutput struct {

	// The updated resource query associated with the resource group after the update.
	GroupQuery *types.GroupQuery

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGroupQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGroupQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGroupQuery{}, middleware.After)
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
	if err = addOpUpdateGroupQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGroupQuery(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateGroupQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-groups",
		OperationName: "UpdateGroupQuery",
	}
}
