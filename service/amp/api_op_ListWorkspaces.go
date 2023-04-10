// Code generated by smithy-go-codegen DO NOT EDIT.

package amp

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all AMP workspaces, including workspaces being created or deleted.
func (c *Client) ListWorkspaces(ctx context.Context, params *ListWorkspacesInput, optFns ...func(*Options)) (*ListWorkspacesOutput, error) {
	if params == nil {
		params = &ListWorkspacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkspaces", params, optFns, c.addOperationListWorkspacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkspacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListWorkspaces operation.
type ListWorkspacesInput struct {

	// Optional filter for workspace alias. Only the workspaces with aliases that
	// begin with this value will be returned.
	Alias *string

	// Maximum results to return in response (default=100, maximum=1000).
	MaxResults *int32

	// Pagination token to request the next page in a paginated list. This token is
	// obtained from the output of the previous ListWorkspaces request.
	NextToken *string

	noSmithyDocumentSerde
}

// Represents the output of a ListWorkspaces operation.
type ListWorkspacesOutput struct {

	// The list of existing workspaces, including those undergoing creation or
	// deletion.
	//
	// This member is required.
	Workspaces []types.WorkspaceSummary

	// Pagination token to use when requesting the next page in this list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkspacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorkspaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorkspaces{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkspaces(options.Region), middleware.Before); err != nil {
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

// ListWorkspacesAPIClient is a client that implements the ListWorkspaces
// operation.
type ListWorkspacesAPIClient interface {
	ListWorkspaces(context.Context, *ListWorkspacesInput, ...func(*Options)) (*ListWorkspacesOutput, error)
}

var _ ListWorkspacesAPIClient = (*Client)(nil)

// ListWorkspacesPaginatorOptions is the paginator options for ListWorkspaces
type ListWorkspacesPaginatorOptions struct {
	// Maximum results to return in response (default=100, maximum=1000).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkspacesPaginator is a paginator for ListWorkspaces
type ListWorkspacesPaginator struct {
	options   ListWorkspacesPaginatorOptions
	client    ListWorkspacesAPIClient
	params    *ListWorkspacesInput
	nextToken *string
	firstPage bool
}

// NewListWorkspacesPaginator returns a new ListWorkspacesPaginator
func NewListWorkspacesPaginator(client ListWorkspacesAPIClient, params *ListWorkspacesInput, optFns ...func(*ListWorkspacesPaginatorOptions)) *ListWorkspacesPaginator {
	if params == nil {
		params = &ListWorkspacesInput{}
	}

	options := ListWorkspacesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkspacesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkspacesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkspaces page.
func (p *ListWorkspacesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkspacesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListWorkspaces(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorkspaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aps",
		OperationName: "ListWorkspaces",
	}
}
