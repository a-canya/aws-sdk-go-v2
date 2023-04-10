// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearchserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an OpenSearch Serverless security policy. For more information, see
// Network access for Amazon OpenSearch Serverless (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-network.html)
// and Encryption at rest for Amazon OpenSearch Serverless (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-encryption.html)
// .
func (c *Client) UpdateSecurityPolicy(ctx context.Context, params *UpdateSecurityPolicyInput, optFns ...func(*Options)) (*UpdateSecurityPolicyOutput, error) {
	if params == nil {
		params = &UpdateSecurityPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSecurityPolicy", params, optFns, c.addOperationUpdateSecurityPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSecurityPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSecurityPolicyInput struct {

	// The name of the policy.
	//
	// This member is required.
	Name *string

	// The version of the policy being updated.
	//
	// This member is required.
	PolicyVersion *string

	// The type of access policy.
	//
	// This member is required.
	Type types.SecurityPolicyType

	// Unique, case-sensitive identifier to ensure idempotency of the request.
	ClientToken *string

	// A description of the policy. Typically used to store information about the
	// permissions defined in the policy.
	Description *string

	// The JSON policy document to use as the content for the new policy.
	Policy *string

	noSmithyDocumentSerde
}

type UpdateSecurityPolicyOutput struct {

	// Details about the updated security policy.
	SecurityPolicyDetail *types.SecurityPolicyDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSecurityPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateSecurityPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateSecurityPolicy{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateSecurityPolicyMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateSecurityPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSecurityPolicy(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateSecurityPolicy struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateSecurityPolicy) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateSecurityPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateSecurityPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateSecurityPolicyInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateSecurityPolicyMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateSecurityPolicy{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateSecurityPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aoss",
		OperationName: "UpdateSecurityPolicy",
	}
}
