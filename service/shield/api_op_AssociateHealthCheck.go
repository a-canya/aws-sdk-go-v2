// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds health-based detection to the Shield Advanced protection for a resource.
// Shield Advanced health-based detection uses the health of your Amazon Web
// Services resource to improve responsiveness and accuracy in attack detection and
// response. You define the health check in Route 53 and then associate it with
// your Shield Advanced protection. For more information, see Shield Advanced
// Health-Based Detection (https://docs.aws.amazon.com/waf/latest/developerguide/ddos-overview.html#ddos-advanced-health-check-option)
// in the WAF Developer Guide.
func (c *Client) AssociateHealthCheck(ctx context.Context, params *AssociateHealthCheckInput, optFns ...func(*Options)) (*AssociateHealthCheckOutput, error) {
	if params == nil {
		params = &AssociateHealthCheckInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateHealthCheck", params, optFns, c.addOperationAssociateHealthCheckMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateHealthCheckOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateHealthCheckInput struct {

	// The Amazon Resource Name (ARN) of the health check to associate with the
	// protection.
	//
	// This member is required.
	HealthCheckArn *string

	// The unique identifier (ID) for the Protection object to add the health check
	// association to.
	//
	// This member is required.
	ProtectionId *string

	noSmithyDocumentSerde
}

type AssociateHealthCheckOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateHealthCheckMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateHealthCheck{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateHealthCheck{}, middleware.After)
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
	if err = addOpAssociateHealthCheckValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateHealthCheck(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateHealthCheck(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "AssociateHealthCheck",
	}
}
