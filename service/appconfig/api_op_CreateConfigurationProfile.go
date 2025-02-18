// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a configuration profile, which is information that enables AppConfig to
// access the configuration source. Valid configuration sources include the
// following:
//   - Configuration data in YAML, JSON, and other formats stored in the AppConfig
//     hosted configuration store
//   - Configuration data stored as objects in an Amazon Simple Storage Service
//     (Amazon S3) bucket
//   - Pipelines stored in CodePipeline
//   - Secrets stored in Secrets Manager
//   - Standard and secure string parameters stored in Amazon Web Services Systems
//     Manager Parameter Store
//   - Configuration data in SSM documents stored in the Systems Manager document
//     store
//
// A configuration profile includes the following information:
//   - The URI location of the configuration data.
//   - The Identity and Access Management (IAM) role that provides access to the
//     configuration data.
//   - A validator for the configuration data. Available validators include either
//     a JSON Schema or an Amazon Web Services Lambda function.
//
// For more information, see Create a Configuration and a Configuration Profile (http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-configuration-and-profile.html)
// in the AppConfig User Guide.
func (c *Client) CreateConfigurationProfile(ctx context.Context, params *CreateConfigurationProfileInput, optFns ...func(*Options)) (*CreateConfigurationProfileOutput, error) {
	if params == nil {
		params = &CreateConfigurationProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConfigurationProfile", params, optFns, c.addOperationCreateConfigurationProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConfigurationProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConfigurationProfileInput struct {

	// The application ID.
	//
	// This member is required.
	ApplicationId *string

	// A URI to locate the configuration. You can specify the following:
	//   - For the AppConfig hosted configuration store and for feature flags, specify
	//   hosted .
	//   - For an Amazon Web Services Systems Manager Parameter Store parameter,
	//   specify either the parameter name in the format ssm-parameter:// or the ARN.
	//   - For an Amazon Web Services CodePipeline pipeline, specify the URI in the
	//   following format: codepipeline ://.
	//   - For an Secrets Manager secret, specify the URI in the following format:
	//   secretsmanager ://.
	//   - For an Amazon S3 object, specify the URI in the following format: s3:/// .
	//   Here is an example: s3://my-bucket/my-app/us-east-1/my-config.json
	//   - For an SSM document, specify either the document name in the format
	//   ssm-document:// or the Amazon Resource Name (ARN).
	//
	// This member is required.
	LocationUri *string

	// A name for the configuration profile.
	//
	// This member is required.
	Name *string

	// A description of the configuration profile.
	Description *string

	// The ARN of an IAM role with permission to access the configuration at the
	// specified LocationUri . A retrieval role ARN is not required for configurations
	// stored in the AppConfig hosted configuration store. It is required for all other
	// sources that store your configuration.
	RetrievalRoleArn *string

	// Metadata to assign to the configuration profile. Tags help organize and
	// categorize your AppConfig resources. Each tag consists of a key and an optional
	// value, both of which you define.
	Tags map[string]string

	// The type of configurations contained in the profile. AppConfig supports feature
	// flags and freeform configurations. We recommend you create feature flag
	// configurations to enable or disable new features and freeform configurations to
	// distribute configurations to an application. When calling this API, enter one of
	// the following values for Type : AWS.AppConfig.FeatureFlags
	//     AWS.Freeform
	Type *string

	// A list of methods for validating the configuration.
	Validators []types.Validator

	noSmithyDocumentSerde
}

type CreateConfigurationProfileOutput struct {

	// The application ID.
	ApplicationId *string

	// The configuration profile description.
	Description *string

	// The configuration profile ID.
	Id *string

	// The URI location of the configuration.
	LocationUri *string

	// The name of the configuration profile.
	Name *string

	// The ARN of an IAM role with permission to access the configuration at the
	// specified LocationUri .
	RetrievalRoleArn *string

	// The type of configurations contained in the profile. AppConfig supports feature
	// flags and freeform configurations. We recommend you create feature flag
	// configurations to enable or disable new features and freeform configurations to
	// distribute configurations to an application. When calling this API, enter one of
	// the following values for Type : AWS.AppConfig.FeatureFlags
	//     AWS.Freeform
	Type *string

	// A list of methods for validating the configuration.
	Validators []types.Validator

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConfigurationProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateConfigurationProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConfigurationProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addCreateConfigurationProfileResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateConfigurationProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConfigurationProfile(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateConfigurationProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "CreateConfigurationProfile",
	}
}

type opCreateConfigurationProfileResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateConfigurationProfileResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateConfigurationProfileResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "appconfig"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "appconfig"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("appconfig")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addCreateConfigurationProfileResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateConfigurationProfileResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
