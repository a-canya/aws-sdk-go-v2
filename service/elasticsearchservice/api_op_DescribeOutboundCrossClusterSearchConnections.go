// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the outbound cross-cluster search connections for a source domain.
func (c *Client) DescribeOutboundCrossClusterSearchConnections(ctx context.Context, params *DescribeOutboundCrossClusterSearchConnectionsInput, optFns ...func(*Options)) (*DescribeOutboundCrossClusterSearchConnectionsOutput, error) {
	if params == nil {
		params = &DescribeOutboundCrossClusterSearchConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOutboundCrossClusterSearchConnections", params, optFns, c.addOperationDescribeOutboundCrossClusterSearchConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOutboundCrossClusterSearchConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the
// DescribeOutboundCrossClusterSearchConnections operation.
type DescribeOutboundCrossClusterSearchConnectionsInput struct {

	// A list of filters used to match properties for outbound cross-cluster search
	// connection. Available Filter names for this operation are:
	//   - cross-cluster-search-connection-id
	//   - destination-domain-info.domain-name
	//   - destination-domain-info.owner-id
	//   - destination-domain-info.region
	//   - source-domain-info.domain-name
	Filters []types.Filter

	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	MaxResults int32

	// NextToken is sent in case the earlier API call results contain the NextToken.
	// It is used for pagination.
	NextToken *string

	noSmithyDocumentSerde
}

// The result of a DescribeOutboundCrossClusterSearchConnections request. Contains
// the list of connections matching the filter criteria.
type DescribeOutboundCrossClusterSearchConnectionsOutput struct {

	// Consists of list of OutboundCrossClusterSearchConnection matching the specified
	// filter criteria.
	CrossClusterSearchConnections []types.OutboundCrossClusterSearchConnection

	// If more results are available and NextToken is present, make the next request
	// to the same API with the received NextToken to paginate the remaining results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOutboundCrossClusterSearchConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeOutboundCrossClusterSearchConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeOutboundCrossClusterSearchConnections{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOutboundCrossClusterSearchConnections(options.Region), middleware.Before); err != nil {
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

// DescribeOutboundCrossClusterSearchConnectionsAPIClient is a client that
// implements the DescribeOutboundCrossClusterSearchConnections operation.
type DescribeOutboundCrossClusterSearchConnectionsAPIClient interface {
	DescribeOutboundCrossClusterSearchConnections(context.Context, *DescribeOutboundCrossClusterSearchConnectionsInput, ...func(*Options)) (*DescribeOutboundCrossClusterSearchConnectionsOutput, error)
}

var _ DescribeOutboundCrossClusterSearchConnectionsAPIClient = (*Client)(nil)

// DescribeOutboundCrossClusterSearchConnectionsPaginatorOptions is the paginator
// options for DescribeOutboundCrossClusterSearchConnections
type DescribeOutboundCrossClusterSearchConnectionsPaginatorOptions struct {
	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeOutboundCrossClusterSearchConnectionsPaginator is a paginator for
// DescribeOutboundCrossClusterSearchConnections
type DescribeOutboundCrossClusterSearchConnectionsPaginator struct {
	options   DescribeOutboundCrossClusterSearchConnectionsPaginatorOptions
	client    DescribeOutboundCrossClusterSearchConnectionsAPIClient
	params    *DescribeOutboundCrossClusterSearchConnectionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeOutboundCrossClusterSearchConnectionsPaginator returns a new
// DescribeOutboundCrossClusterSearchConnectionsPaginator
func NewDescribeOutboundCrossClusterSearchConnectionsPaginator(client DescribeOutboundCrossClusterSearchConnectionsAPIClient, params *DescribeOutboundCrossClusterSearchConnectionsInput, optFns ...func(*DescribeOutboundCrossClusterSearchConnectionsPaginatorOptions)) *DescribeOutboundCrossClusterSearchConnectionsPaginator {
	if params == nil {
		params = &DescribeOutboundCrossClusterSearchConnectionsInput{}
	}

	options := DescribeOutboundCrossClusterSearchConnectionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeOutboundCrossClusterSearchConnectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOutboundCrossClusterSearchConnectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeOutboundCrossClusterSearchConnections page.
func (p *DescribeOutboundCrossClusterSearchConnectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOutboundCrossClusterSearchConnectionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeOutboundCrossClusterSearchConnections(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeOutboundCrossClusterSearchConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeOutboundCrossClusterSearchConnections",
	}
}
