// Code generated by smithy-go-codegen DO NOT EDIT.

package appintegrations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appintegrations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of DataIntegration associations in the account. You
// cannot create a DataIntegration association for a DataIntegration that has been
// previously associated. Use a different DataIntegration, or recreate the
// DataIntegration using the CreateDataIntegration (https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html)
// API.
func (c *Client) ListDataIntegrationAssociations(ctx context.Context, params *ListDataIntegrationAssociationsInput, optFns ...func(*Options)) (*ListDataIntegrationAssociationsOutput, error) {
	if params == nil {
		params = &ListDataIntegrationAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataIntegrationAssociations", params, optFns, c.addOperationListDataIntegrationAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataIntegrationAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataIntegrationAssociationsInput struct {

	// A unique identifier for the DataIntegration.
	//
	// This member is required.
	DataIntegrationIdentifier *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDataIntegrationAssociationsOutput struct {

	// The Amazon Resource Name (ARN) and unique ID of the DataIntegration association.
	DataIntegrationAssociations []types.DataIntegrationAssociationSummary

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataIntegrationAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDataIntegrationAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDataIntegrationAssociations{}, middleware.After)
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
	if err = addOpListDataIntegrationAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataIntegrationAssociations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDataIntegrationAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "app-integrations",
		OperationName: "ListDataIntegrationAssociations",
	}
}
