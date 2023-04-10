// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about cluster snapshots. This API operation supports
// pagination.
func (c *Client) DescribeDBClusterSnapshots(ctx context.Context, params *DescribeDBClusterSnapshotsInput, optFns ...func(*Options)) (*DescribeDBClusterSnapshotsOutput, error) {
	if params == nil {
		params = &DescribeDBClusterSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusterSnapshots", params, optFns, c.addOperationDescribeDBClusterSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClusterSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to DescribeDBClusterSnapshots .
type DescribeDBClusterSnapshotsInput struct {

	// The ID of the cluster to retrieve the list of cluster snapshots for. This
	// parameter can't be used with the DBClusterSnapshotIdentifier parameter. This
	// parameter is not case sensitive. Constraints:
	//   - If provided, must match the identifier of an existing DBCluster .
	DBClusterIdentifier *string

	// A specific cluster snapshot identifier to describe. This parameter can't be
	// used with the DBClusterIdentifier parameter. This value is stored as a
	// lowercase string. Constraints:
	//   - If provided, must match the identifier of an existing DBClusterSnapshot .
	//   - If this identifier is for an automated snapshot, the SnapshotType parameter
	//   must also be specified.
	DBClusterSnapshotIdentifier *string

	// This parameter is not currently supported.
	Filters []types.Filter

	// Set to true to include manual cluster snapshots that are public and can be
	// copied or restored by any Amazon Web Services account, and otherwise false . The
	// default is false .
	IncludePublic bool

	// Set to true to include shared manual cluster snapshots from other Amazon Web
	// Services accounts that this Amazon Web Services account has been given
	// permission to copy or restore, and otherwise false . The default is false .
	IncludeShared bool

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token (marker) is included in
	// the response so that the remaining results can be retrieved. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The type of cluster snapshots to be returned. You can specify one of the
	// following values:
	//   - automated - Return all cluster snapshots that Amazon DocumentDB has
	//   automatically created for your Amazon Web Services account.
	//   - manual - Return all cluster snapshots that you have manually created for
	//   your Amazon Web Services account.
	//   - shared - Return all manual cluster snapshots that have been shared to your
	//   Amazon Web Services account.
	//   - public - Return all cluster snapshots that have been marked as public.
	// If you don't specify a SnapshotType value, then both automated and manual
	// cluster snapshots are returned. You can include shared cluster snapshots with
	// these results by setting the IncludeShared parameter to true . You can include
	// public cluster snapshots with these results by setting the IncludePublic
	// parameter to true . The IncludeShared and IncludePublic parameters don't apply
	// for SnapshotType values of manual or automated . The IncludePublic parameter
	// doesn't apply when SnapshotType is set to shared . The IncludeShared parameter
	// doesn't apply when SnapshotType is set to public .
	SnapshotType *string

	noSmithyDocumentSerde
}

// Represents the output of DescribeDBClusterSnapshots .
type DescribeDBClusterSnapshotsOutput struct {

	// Provides a list of cluster snapshots.
	DBClusterSnapshots []types.DBClusterSnapshot

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBClusterSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterSnapshots{}, middleware.After)
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
	if err = addOpDescribeDBClusterSnapshotsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterSnapshots(options.Region), middleware.Before); err != nil {
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

// DescribeDBClusterSnapshotsAPIClient is a client that implements the
// DescribeDBClusterSnapshots operation.
type DescribeDBClusterSnapshotsAPIClient interface {
	DescribeDBClusterSnapshots(context.Context, *DescribeDBClusterSnapshotsInput, ...func(*Options)) (*DescribeDBClusterSnapshotsOutput, error)
}

var _ DescribeDBClusterSnapshotsAPIClient = (*Client)(nil)

// DescribeDBClusterSnapshotsPaginatorOptions is the paginator options for
// DescribeDBClusterSnapshots
type DescribeDBClusterSnapshotsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token (marker) is included in
	// the response so that the remaining results can be retrieved. Default: 100
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBClusterSnapshotsPaginator is a paginator for
// DescribeDBClusterSnapshots
type DescribeDBClusterSnapshotsPaginator struct {
	options   DescribeDBClusterSnapshotsPaginatorOptions
	client    DescribeDBClusterSnapshotsAPIClient
	params    *DescribeDBClusterSnapshotsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBClusterSnapshotsPaginator returns a new
// DescribeDBClusterSnapshotsPaginator
func NewDescribeDBClusterSnapshotsPaginator(client DescribeDBClusterSnapshotsAPIClient, params *DescribeDBClusterSnapshotsInput, optFns ...func(*DescribeDBClusterSnapshotsPaginatorOptions)) *DescribeDBClusterSnapshotsPaginator {
	if params == nil {
		params = &DescribeDBClusterSnapshotsInput{}
	}

	options := DescribeDBClusterSnapshotsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBClusterSnapshotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBClusterSnapshotsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBClusterSnapshots page.
func (p *DescribeDBClusterSnapshotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBClusterSnapshotsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeDBClusterSnapshots(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeDBClusterSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusterSnapshots",
	}
}
