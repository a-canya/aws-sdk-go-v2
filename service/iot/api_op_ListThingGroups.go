// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the thing groups in your account. Requires permission to access the
// ListThingGroups (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListThingGroups(ctx context.Context, params *ListThingGroupsInput, optFns ...func(*Options)) (*ListThingGroupsOutput, error) {
	if params == nil {
		params = &ListThingGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThingGroups", params, optFns, c.addOperationListThingGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThingGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThingGroupsInput struct {

	// The maximum number of results to return at one time.
	MaxResults *int32

	// A filter that limits the results to those with the specified name prefix.
	NamePrefixFilter *string

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	// A filter that limits the results to those with the specified parent group.
	ParentGroup *string

	// If true, return child groups as well.
	Recursive *bool

	noSmithyDocumentSerde
}

type ListThingGroupsOutput struct {

	// The token to use to get the next set of results. Will not be returned if
	// operation has returned all results.
	NextToken *string

	// The thing groups.
	ThingGroups []types.GroupNameAndArn

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListThingGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThingGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThingGroups(options.Region), middleware.Before); err != nil {
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

// ListThingGroupsAPIClient is a client that implements the ListThingGroups
// operation.
type ListThingGroupsAPIClient interface {
	ListThingGroups(context.Context, *ListThingGroupsInput, ...func(*Options)) (*ListThingGroupsOutput, error)
}

var _ ListThingGroupsAPIClient = (*Client)(nil)

// ListThingGroupsPaginatorOptions is the paginator options for ListThingGroups
type ListThingGroupsPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListThingGroupsPaginator is a paginator for ListThingGroups
type ListThingGroupsPaginator struct {
	options   ListThingGroupsPaginatorOptions
	client    ListThingGroupsAPIClient
	params    *ListThingGroupsInput
	nextToken *string
	firstPage bool
}

// NewListThingGroupsPaginator returns a new ListThingGroupsPaginator
func NewListThingGroupsPaginator(client ListThingGroupsAPIClient, params *ListThingGroupsInput, optFns ...func(*ListThingGroupsPaginatorOptions)) *ListThingGroupsPaginator {
	if params == nil {
		params = &ListThingGroupsInput{}
	}

	options := ListThingGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListThingGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListThingGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListThingGroups page.
func (p *ListThingGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListThingGroupsOutput, error) {
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

	result, err := p.client.ListThingGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListThingGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThingGroups",
	}
}
