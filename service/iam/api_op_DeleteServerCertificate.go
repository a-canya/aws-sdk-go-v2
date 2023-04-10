// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified server certificate. For more information about working
// with server certificates, see Working with server certificates (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html)
// in the IAM User Guide. This topic also includes a list of Amazon Web Services
// services that can use the server certificates that you manage with IAM. If you
// are using a server certificate with Elastic Load Balancing, deleting the
// certificate could have implications for your application. If Elastic Load
// Balancing doesn't detect the deletion of bound certificates, it may continue to
// use the certificates. This could cause Elastic Load Balancing to stop accepting
// traffic. We recommend that you remove the reference to the certificate from
// Elastic Load Balancing before using this command to delete the certificate. For
// more information, see DeleteLoadBalancerListeners (https://docs.aws.amazon.com/ElasticLoadBalancing/latest/APIReference/API_DeleteLoadBalancerListeners.html)
// in the Elastic Load Balancing API Reference.
func (c *Client) DeleteServerCertificate(ctx context.Context, params *DeleteServerCertificateInput, optFns ...func(*Options)) (*DeleteServerCertificateOutput, error) {
	if params == nil {
		params = &DeleteServerCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteServerCertificate", params, optFns, c.addOperationDeleteServerCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteServerCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteServerCertificateInput struct {

	// The name of the server certificate you want to delete. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex) ) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	ServerCertificateName *string

	noSmithyDocumentSerde
}

type DeleteServerCertificateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteServerCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteServerCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteServerCertificate{}, middleware.After)
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
	if err = addOpDeleteServerCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteServerCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteServerCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DeleteServerCertificate",
	}
}
