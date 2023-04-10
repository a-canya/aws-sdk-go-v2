// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves all aliases for this Amazon Web Services account. You can filter the
// result set by alias name and/or routing strategy type. Use the pagination
// parameters to retrieve results in sequential pages. Returned aliases are not
// listed in any particular order. Related actions All APIs by task (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) ListAliases(ctx context.Context, params *ListAliasesInput, optFns ...func(*Options)) (*ListAliasesOutput, error) {
	if params == nil {
		params = &ListAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAliases", params, optFns, c.addOperationListAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAliasesInput struct {

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// A descriptive label that is associated with an alias. Alias names do not need
	// to be unique.
	Name *string

	// A token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value.
	NextToken *string

	// The routing type to filter results on. Use this parameter to retrieve only
	// aliases with a certain routing type. To retrieve all aliases, leave this
	// parameter empty. Possible routing types include the following:
	//   - SIMPLE -- The alias resolves to one specific fleet. Use this type when
	//   routing to active fleets.
	//   - TERMINAL -- The alias does not resolve to a fleet but instead can be used
	//   to display a message to the user. A terminal alias throws a
	//   TerminalRoutingStrategyException with the RoutingStrategy (https://docs.aws.amazon.com/gamelift/latest/apireference/API_RoutingStrategy.html)
	//   message embedded.
	RoutingStrategyType types.RoutingStrategyType

	noSmithyDocumentSerde
}

type ListAliasesOutput struct {

	// A collection of alias resources that match the request parameters.
	Aliases []types.Alias

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAliases{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAliases(options.Region), middleware.Before); err != nil {
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

// ListAliasesAPIClient is a client that implements the ListAliases operation.
type ListAliasesAPIClient interface {
	ListAliases(context.Context, *ListAliasesInput, ...func(*Options)) (*ListAliasesOutput, error)
}

var _ ListAliasesAPIClient = (*Client)(nil)

// ListAliasesPaginatorOptions is the paginator options for ListAliases
type ListAliasesPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAliasesPaginator is a paginator for ListAliases
type ListAliasesPaginator struct {
	options   ListAliasesPaginatorOptions
	client    ListAliasesAPIClient
	params    *ListAliasesInput
	nextToken *string
	firstPage bool
}

// NewListAliasesPaginator returns a new ListAliasesPaginator
func NewListAliasesPaginator(client ListAliasesAPIClient, params *ListAliasesInput, optFns ...func(*ListAliasesPaginatorOptions)) *ListAliasesPaginator {
	if params == nil {
		params = &ListAliasesInput{}
	}

	options := ListAliasesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAliasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAliasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAliases page.
func (p *ListAliasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAliasesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListAliases(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ListAliases",
	}
}
