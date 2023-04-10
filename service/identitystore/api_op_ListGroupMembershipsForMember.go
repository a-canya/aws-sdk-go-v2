// Code generated by smithy-go-codegen DO NOT EDIT.

package identitystore

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For the specified member in the specified identity store, returns the list of
// all GroupMembership objects and returns results in paginated form.
func (c *Client) ListGroupMembershipsForMember(ctx context.Context, params *ListGroupMembershipsForMemberInput, optFns ...func(*Options)) (*ListGroupMembershipsForMemberOutput, error) {
	if params == nil {
		params = &ListGroupMembershipsForMemberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGroupMembershipsForMember", params, optFns, c.addOperationListGroupMembershipsForMemberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGroupMembershipsForMemberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroupMembershipsForMemberInput struct {

	// The globally unique identifier for the identity store.
	//
	// This member is required.
	IdentityStoreId *string

	// An object that contains the identifier of a group member. Setting the UserID
	// field to the specific identifier for a user indicates that the user is a member
	// of the group.
	//
	// This member is required.
	MemberId types.MemberId

	// The maximum number of results to be returned per request. This parameter is
	// used in the ListUsers and ListGroups requests to specify how many results to
	// return in one page. The length limit is 50 characters.
	MaxResults *int32

	// The pagination token used for the ListUsers , ListGroups , and
	// ListGroupMemberships API operations. This value is generated by the identity
	// store service. It is returned in the API response if the total results are more
	// than the size of one page. This token is also returned when it is used in the
	// API request to search for the next page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGroupMembershipsForMemberOutput struct {

	// A list of GroupMembership objects in the group for a specified member.
	//
	// This member is required.
	GroupMemberships []types.GroupMembership

	// The pagination token used for the ListUsers , ListGroups , and
	// ListGroupMemberships API operations. This value is generated by the identity
	// store service. It is returned in the API response if the total results are more
	// than the size of one page. This token is also returned when it is used in the
	// API request to search for the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGroupMembershipsForMemberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListGroupMembershipsForMember{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListGroupMembershipsForMember{}, middleware.After)
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
	if err = addOpListGroupMembershipsForMemberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGroupMembershipsForMember(options.Region), middleware.Before); err != nil {
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

// ListGroupMembershipsForMemberAPIClient is a client that implements the
// ListGroupMembershipsForMember operation.
type ListGroupMembershipsForMemberAPIClient interface {
	ListGroupMembershipsForMember(context.Context, *ListGroupMembershipsForMemberInput, ...func(*Options)) (*ListGroupMembershipsForMemberOutput, error)
}

var _ ListGroupMembershipsForMemberAPIClient = (*Client)(nil)

// ListGroupMembershipsForMemberPaginatorOptions is the paginator options for
// ListGroupMembershipsForMember
type ListGroupMembershipsForMemberPaginatorOptions struct {
	// The maximum number of results to be returned per request. This parameter is
	// used in the ListUsers and ListGroups requests to specify how many results to
	// return in one page. The length limit is 50 characters.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGroupMembershipsForMemberPaginator is a paginator for
// ListGroupMembershipsForMember
type ListGroupMembershipsForMemberPaginator struct {
	options   ListGroupMembershipsForMemberPaginatorOptions
	client    ListGroupMembershipsForMemberAPIClient
	params    *ListGroupMembershipsForMemberInput
	nextToken *string
	firstPage bool
}

// NewListGroupMembershipsForMemberPaginator returns a new
// ListGroupMembershipsForMemberPaginator
func NewListGroupMembershipsForMemberPaginator(client ListGroupMembershipsForMemberAPIClient, params *ListGroupMembershipsForMemberInput, optFns ...func(*ListGroupMembershipsForMemberPaginatorOptions)) *ListGroupMembershipsForMemberPaginator {
	if params == nil {
		params = &ListGroupMembershipsForMemberInput{}
	}

	options := ListGroupMembershipsForMemberPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGroupMembershipsForMemberPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGroupMembershipsForMemberPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGroupMembershipsForMember page.
func (p *ListGroupMembershipsForMemberPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGroupMembershipsForMemberOutput, error) {
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

	result, err := p.client.ListGroupMembershipsForMember(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGroupMembershipsForMember(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "identitystore",
		OperationName: "ListGroupMembershipsForMember",
	}
}
