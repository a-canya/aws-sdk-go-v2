// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List service instances with summary data. This action lists service instances
// of all services in the Amazon Web Services account.
func (c *Client) ListServiceInstances(ctx context.Context, params *ListServiceInstancesInput, optFns ...func(*Options)) (*ListServiceInstancesOutput, error) {
	if params == nil {
		params = &ListServiceInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceInstances", params, optFns, c.addOperationListServiceInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceInstancesInput struct {

	// An array of filtering criteria that scope down the result list. By default, all
	// service instances in the Amazon Web Services account are returned.
	Filters []types.ListServiceInstancesFilter

	// The maximum number of service instances to list.
	MaxResults *int32

	// A token that indicates the location of the next service in the array of service
	// instances, after the list of service instances that was previously requested.
	NextToken *string

	// The name of the service that the service instance belongs to.
	ServiceName *string

	// The field that the result list is sorted by. When you choose to sort by
	// serviceName , service instances within each service are sorted by service
	// instance name. Default: serviceName
	SortBy types.ListServiceInstancesSortBy

	// Result list sort order. Default: ASCENDING
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListServiceInstancesOutput struct {

	// An array of service instances with summary data.
	//
	// This member is required.
	ServiceInstances []types.ServiceInstanceSummary

	// A token that indicates the location of the next service instance in the array
	// of service instances, after the current requested list of service instances.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServiceInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListServiceInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListServiceInstances{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceInstances(options.Region), middleware.Before); err != nil {
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

// ListServiceInstancesAPIClient is a client that implements the
// ListServiceInstances operation.
type ListServiceInstancesAPIClient interface {
	ListServiceInstances(context.Context, *ListServiceInstancesInput, ...func(*Options)) (*ListServiceInstancesOutput, error)
}

var _ ListServiceInstancesAPIClient = (*Client)(nil)

// ListServiceInstancesPaginatorOptions is the paginator options for
// ListServiceInstances
type ListServiceInstancesPaginatorOptions struct {
	// The maximum number of service instances to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServiceInstancesPaginator is a paginator for ListServiceInstances
type ListServiceInstancesPaginator struct {
	options   ListServiceInstancesPaginatorOptions
	client    ListServiceInstancesAPIClient
	params    *ListServiceInstancesInput
	nextToken *string
	firstPage bool
}

// NewListServiceInstancesPaginator returns a new ListServiceInstancesPaginator
func NewListServiceInstancesPaginator(client ListServiceInstancesAPIClient, params *ListServiceInstancesInput, optFns ...func(*ListServiceInstancesPaginatorOptions)) *ListServiceInstancesPaginator {
	if params == nil {
		params = &ListServiceInstancesInput{}
	}

	options := ListServiceInstancesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServiceInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServiceInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListServiceInstances page.
func (p *ListServiceInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServiceInstancesOutput, error) {
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

	result, err := p.client.ListServiceInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListServiceInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "ListServiceInstances",
	}
}
