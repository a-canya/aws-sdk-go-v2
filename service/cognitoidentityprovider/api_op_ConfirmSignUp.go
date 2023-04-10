// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Confirms registration of a new user.
func (c *Client) ConfirmSignUp(ctx context.Context, params *ConfirmSignUpInput, optFns ...func(*Options)) (*ConfirmSignUpOutput, error) {
	if params == nil {
		params = &ConfirmSignUpInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConfirmSignUp", params, optFns, c.addOperationConfirmSignUpMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConfirmSignUpOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to confirm registration of a user.
type ConfirmSignUpInput struct {

	// The ID of the app client associated with the user pool.
	//
	// This member is required.
	ClientId *string

	// The confirmation code sent by a user's request to confirm registration.
	//
	// This member is required.
	ConfirmationCode *string

	// The user name of the user whose registration you want to confirm.
	//
	// This member is required.
	Username *string

	// The Amazon Pinpoint analytics metadata for collecting metrics for ConfirmSignUp
	// calls.
	AnalyticsMetadata *types.AnalyticsMetadataType

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// Lambda functions to user pool triggers. When you use the ConfirmSignUp API
	// action, Amazon Cognito invokes the function that is assigned to the post
	// confirmation trigger. When Amazon Cognito invokes this function, it passes a
	// JSON payload, which the function receives as input. This payload contains a
	// clientMetadata attribute, which provides the data that you assigned to the
	// ClientMetadata parameter in your ConfirmSignUp request. In your function code in
	// Lambda, you can process the clientMetadata value to enhance your workflow for
	// your specific needs. For more information, see Customizing user pool Workflows
	// with Lambda Triggers (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. When you use the ClientMetadata
	// parameter, remember that Amazon Cognito won't do the following:
	//   - Store the ClientMetadata value. This data is available only to Lambda
	//   triggers that are assigned to a user pool to support custom workflows. If your
	//   user pool configuration doesn't include triggers, the ClientMetadata parameter
	//   serves no purpose.
	//   - Validate the ClientMetadata value.
	//   - Encrypt the ClientMetadata value. Don't use Amazon Cognito to provide
	//   sensitive information.
	ClientMetadata map[string]string

	// Boolean to be specified to force user confirmation irrespective of existing
	// alias. By default set to False . If this parameter is set to True and the phone
	// number/email used for sign up confirmation already exists as an alias with a
	// different user, the API call will migrate the alias from the previous user to
	// the newly created user being confirmed. If set to False , the API will throw an
	// AliasExistsException error.
	ForceAliasCreation bool

	// A keyed-hash message authentication code (HMAC) calculated using the secret key
	// of a user pool client and username plus the client ID in the message.
	SecretHash *string

	// Contextual data about your user session, such as the device fingerprint, IP
	// address, or location. Amazon Cognito advanced security evaluates the risk of an
	// authentication event based on the context that your app generates and passes to
	// Amazon Cognito when it makes API requests.
	UserContextData *types.UserContextDataType

	noSmithyDocumentSerde
}

// Represents the response from the server for the registration confirmation.
type ConfirmSignUpOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationConfirmSignUpMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpConfirmSignUp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpConfirmSignUp{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addOpConfirmSignUpValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opConfirmSignUp(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opConfirmSignUp(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ConfirmSignUp",
	}
}
