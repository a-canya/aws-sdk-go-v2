// Code generated by smithy-go-codegen DO NOT EDIT.

package workspacesweb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspacesweb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a web portal.
func (c *Client) CreatePortal(ctx context.Context, params *CreatePortalInput, optFns ...func(*Options)) (*CreatePortalOutput, error) {
	if params == nil {
		params = &CreatePortalInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePortal", params, optFns, c.addOperationCreatePortalMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePortalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePortalInput struct {

	// The additional encryption context of the portal.
	AdditionalEncryptionContext map[string]string

	// The type of authentication integration points used when signing into the web
	// portal. Defaults to Standard . Standard web portals are authenticated directly
	// through your identity provider. You need to call CreateIdentityProvider to
	// integrate your identity provider with your web portal. User and group access to
	// your web portal is controlled through your identity provider.
	// IAM_Identity_Center web portals are authenticated through AWS IAM Identity
	// Center (successor to AWS Single Sign-On). They provide additional features, such
	// as IdP-initiated authentication. Identity sources (including external identity
	// provider integration), plus user and group access to your web portal, can be
	// configured in the IAM Identity Center.
	AuthenticationType types.AuthenticationType

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. Idempotency ensures that an API request completes only once.
	// With an idempotent request, if the original request completes successfully,
	// subsequent retries with the same client token returns the result from the
	// original successful request. If you do not specify a client token, one is
	// automatically generated by the AWS SDK.
	ClientToken *string

	// The customer managed key of the web portal.
	CustomerManagedKey *string

	// The name of the web portal. This is not visible to users who log into the web
	// portal.
	DisplayName *string

	// The tags to add to the web portal. A tag is a key-value pair.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreatePortalOutput struct {

	// The ARN of the web portal.
	//
	// This member is required.
	PortalArn *string

	// The endpoint URL of the web portal that users access in order to start
	// streaming sessions.
	//
	// This member is required.
	PortalEndpoint *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePortalMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePortal{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePortal{}, middleware.After)
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
	if err = addIdempotencyToken_opCreatePortalMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePortalValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePortal(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreatePortal struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePortal) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePortal) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePortalInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePortalInput ")
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
func addIdempotencyToken_opCreatePortalMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreatePortal{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePortal(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces-web",
		OperationName: "CreatePortal",
	}
}
