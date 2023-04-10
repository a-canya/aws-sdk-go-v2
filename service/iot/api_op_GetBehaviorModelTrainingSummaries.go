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

// Returns a Device Defender's ML Detect Security Profile training model's status.
// Requires permission to access the GetBehaviorModelTrainingSummaries (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) GetBehaviorModelTrainingSummaries(ctx context.Context, params *GetBehaviorModelTrainingSummariesInput, optFns ...func(*Options)) (*GetBehaviorModelTrainingSummariesOutput, error) {
	if params == nil {
		params = &GetBehaviorModelTrainingSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBehaviorModelTrainingSummaries", params, optFns, c.addOperationGetBehaviorModelTrainingSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBehaviorModelTrainingSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBehaviorModelTrainingSummariesInput struct {

	// The maximum number of results to return at one time. The default is 10.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// The name of the security profile.
	SecurityProfileName *string

	noSmithyDocumentSerde
}

type GetBehaviorModelTrainingSummariesOutput struct {

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// A list of all ML Detect behaviors and their model status for a given Security
	// Profile.
	Summaries []types.BehaviorModelTrainingSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBehaviorModelTrainingSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBehaviorModelTrainingSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBehaviorModelTrainingSummaries{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBehaviorModelTrainingSummaries(options.Region), middleware.Before); err != nil {
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

// GetBehaviorModelTrainingSummariesAPIClient is a client that implements the
// GetBehaviorModelTrainingSummaries operation.
type GetBehaviorModelTrainingSummariesAPIClient interface {
	GetBehaviorModelTrainingSummaries(context.Context, *GetBehaviorModelTrainingSummariesInput, ...func(*Options)) (*GetBehaviorModelTrainingSummariesOutput, error)
}

var _ GetBehaviorModelTrainingSummariesAPIClient = (*Client)(nil)

// GetBehaviorModelTrainingSummariesPaginatorOptions is the paginator options for
// GetBehaviorModelTrainingSummaries
type GetBehaviorModelTrainingSummariesPaginatorOptions struct {
	// The maximum number of results to return at one time. The default is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetBehaviorModelTrainingSummariesPaginator is a paginator for
// GetBehaviorModelTrainingSummaries
type GetBehaviorModelTrainingSummariesPaginator struct {
	options   GetBehaviorModelTrainingSummariesPaginatorOptions
	client    GetBehaviorModelTrainingSummariesAPIClient
	params    *GetBehaviorModelTrainingSummariesInput
	nextToken *string
	firstPage bool
}

// NewGetBehaviorModelTrainingSummariesPaginator returns a new
// GetBehaviorModelTrainingSummariesPaginator
func NewGetBehaviorModelTrainingSummariesPaginator(client GetBehaviorModelTrainingSummariesAPIClient, params *GetBehaviorModelTrainingSummariesInput, optFns ...func(*GetBehaviorModelTrainingSummariesPaginatorOptions)) *GetBehaviorModelTrainingSummariesPaginator {
	if params == nil {
		params = &GetBehaviorModelTrainingSummariesInput{}
	}

	options := GetBehaviorModelTrainingSummariesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetBehaviorModelTrainingSummariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetBehaviorModelTrainingSummariesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetBehaviorModelTrainingSummaries page.
func (p *GetBehaviorModelTrainingSummariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetBehaviorModelTrainingSummariesOutput, error) {
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

	result, err := p.client.GetBehaviorModelTrainingSummaries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetBehaviorModelTrainingSummaries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "GetBehaviorModelTrainingSummaries",
	}
}
