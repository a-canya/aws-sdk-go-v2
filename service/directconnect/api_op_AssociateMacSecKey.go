// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a MAC Security (MACsec) Connection Key Name (CKN)/ Connectivity
// Association Key (CAK) pair with an Direct Connect dedicated connection. You must
// supply either the secretARN, or the CKN/CAK ( ckn and cak ) pair in the request.
// For information about MAC Security (MACsec) key considerations, see MACsec
// pre-shared CKN/CAK key considerations  (https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-key-consideration)
// in the Direct Connect User Guide.
func (c *Client) AssociateMacSecKey(ctx context.Context, params *AssociateMacSecKeyInput, optFns ...func(*Options)) (*AssociateMacSecKeyOutput, error) {
	if params == nil {
		params = &AssociateMacSecKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateMacSecKey", params, optFns, c.addOperationAssociateMacSecKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateMacSecKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateMacSecKeyInput struct {

	// The ID of the dedicated connection (dxcon-xxxx), or the ID of the LAG
	// (dxlag-xxxx). You can use DescribeConnections or DescribeLags to retrieve
	// connection ID.
	//
	// This member is required.
	ConnectionId *string

	// The MAC Security (MACsec) CAK to associate with the dedicated connection. You
	// can create the CKN/CAK pair using an industry standard tool. The valid values
	// are 64 hexadecimal characters (0-9, A-E). If you use this request parameter, you
	// must use the ckn request parameter and not use the secretARN request parameter.
	Cak *string

	// The MAC Security (MACsec) CKN to associate with the dedicated connection. You
	// can create the CKN/CAK pair using an industry standard tool. The valid values
	// are 64 hexadecimal characters (0-9, A-E). If you use this request parameter, you
	// must use the cak request parameter and not use the secretARN request parameter.
	Ckn *string

	// The Amazon Resource Name (ARN) of the MAC Security (MACsec) secret key to
	// associate with the dedicated connection. You can use DescribeConnections or
	// DescribeLags to retrieve the MAC Security (MACsec) secret key. If you use this
	// request parameter, you do not use the ckn and cak request parameters.
	SecretARN *string

	noSmithyDocumentSerde
}

type AssociateMacSecKeyOutput struct {

	// The ID of the dedicated connection (dxcon-xxxx), or the ID of the LAG
	// (dxlag-xxxx).
	ConnectionId *string

	// The MAC Security (MACsec) security keys associated with the dedicated
	// connection.
	MacSecKeys []types.MacSecKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateMacSecKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateMacSecKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateMacSecKey{}, middleware.After)
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
	if err = addOpAssociateMacSecKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateMacSecKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateMacSecKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "AssociateMacSecKey",
	}
}
