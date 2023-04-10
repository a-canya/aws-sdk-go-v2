// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List template post migration custom actions.
func (c *Client) ListTemplateActions(ctx context.Context, params *ListTemplateActionsInput, optFns ...func(*Options)) (*ListTemplateActionsOutput, error) {
	if params == nil {
		params = &ListTemplateActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTemplateActions", params, optFns, c.addOperationListTemplateActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTemplateActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTemplateActionsInput struct {

	// Launch configuration template ID.
	//
	// This member is required.
	LaunchConfigurationTemplateID *string

	// Filters to apply when listing template post migration custom actions.
	Filters *types.TemplateActionsRequestFilters

	// Maximum amount of items to return when listing template post migration custom
	// actions.
	MaxResults int32

	// Next token to use when listing template post migration custom actions.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTemplateActionsOutput struct {

	// List of template post migration custom actions.
	Items []types.TemplateActionDocument

	// Next token returned when listing template post migration custom actions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTemplateActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTemplateActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTemplateActions{}, middleware.After)
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
	if err = addOpListTemplateActionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTemplateActions(options.Region), middleware.Before); err != nil {
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

// ListTemplateActionsAPIClient is a client that implements the
// ListTemplateActions operation.
type ListTemplateActionsAPIClient interface {
	ListTemplateActions(context.Context, *ListTemplateActionsInput, ...func(*Options)) (*ListTemplateActionsOutput, error)
}

var _ ListTemplateActionsAPIClient = (*Client)(nil)

// ListTemplateActionsPaginatorOptions is the paginator options for
// ListTemplateActions
type ListTemplateActionsPaginatorOptions struct {
	// Maximum amount of items to return when listing template post migration custom
	// actions.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTemplateActionsPaginator is a paginator for ListTemplateActions
type ListTemplateActionsPaginator struct {
	options   ListTemplateActionsPaginatorOptions
	client    ListTemplateActionsAPIClient
	params    *ListTemplateActionsInput
	nextToken *string
	firstPage bool
}

// NewListTemplateActionsPaginator returns a new ListTemplateActionsPaginator
func NewListTemplateActionsPaginator(client ListTemplateActionsAPIClient, params *ListTemplateActionsInput, optFns ...func(*ListTemplateActionsPaginatorOptions)) *ListTemplateActionsPaginator {
	if params == nil {
		params = &ListTemplateActionsInput{}
	}

	options := ListTemplateActionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTemplateActionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTemplateActionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTemplateActions page.
func (p *ListTemplateActionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTemplateActionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListTemplateActions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListTemplateActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "ListTemplateActions",
	}
}
