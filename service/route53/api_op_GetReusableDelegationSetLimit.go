// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the maximum number of hosted zones that you can associate with the
// specified reusable delegation set. For the default limit, see Limits (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
// in the Amazon Route 53 Developer Guide. To request a higher limit, open a case (https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-route53)
// .
func (c *Client) GetReusableDelegationSetLimit(ctx context.Context, params *GetReusableDelegationSetLimitInput, optFns ...func(*Options)) (*GetReusableDelegationSetLimitOutput, error) {
	if params == nil {
		params = &GetReusableDelegationSetLimitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReusableDelegationSetLimit", params, optFns, c.addOperationGetReusableDelegationSetLimitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReusableDelegationSetLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the request to create a hosted
// zone.
type GetReusableDelegationSetLimitInput struct {

	// The ID of the delegation set that you want to get the limit for.
	//
	// This member is required.
	DelegationSetId *string

	// Specify MAX_ZONES_BY_REUSABLE_DELEGATION_SET to get the maximum number of
	// hosted zones that you can associate with the specified reusable delegation set.
	//
	// This member is required.
	Type types.ReusableDelegationSetLimitType

	noSmithyDocumentSerde
}

// A complex type that contains the requested limit.
type GetReusableDelegationSetLimitOutput struct {

	// The current number of hosted zones that you can associate with the specified
	// reusable delegation set.
	//
	// This member is required.
	Count int64

	// The current setting for the limit on hosted zones that you can associate with
	// the specified reusable delegation set.
	//
	// This member is required.
	Limit *types.ReusableDelegationSetLimit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReusableDelegationSetLimitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetReusableDelegationSetLimit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetReusableDelegationSetLimit{}, middleware.After)
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
	if err = addOpGetReusableDelegationSetLimitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReusableDelegationSetLimit(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetReusableDelegationSetLimit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetReusableDelegationSetLimit",
	}
}
