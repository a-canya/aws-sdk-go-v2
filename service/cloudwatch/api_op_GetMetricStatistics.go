// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets statistics for the specified metric. The maximum number of data points
// returned from a single call is 1,440. If you request more than 1,440 data
// points, CloudWatch returns an error. To reduce the number of data points, you
// can narrow the specified time range and make multiple requests across adjacent
// time ranges, or you can increase the specified period. Data points are not
// returned in chronological order. CloudWatch aggregates data points based on the
// length of the period that you specify. For example, if you request statistics
// with a one-hour period, CloudWatch aggregates all data points with time stamps
// that fall within each one-hour period. Therefore, the number of values
// aggregated by CloudWatch is larger than the number of data points returned.
// CloudWatch needs raw data points to calculate percentile statistics. If you
// publish data using a statistic set instead, you can only retrieve percentile
// statistics for this data if one of the following conditions is true:
//   - The SampleCount value of the statistic set is 1.
//   - The Min and the Max values of the statistic set are equal.
//
// Percentile statistics are not available for metrics when any of the metric
// values are negative numbers. Amazon CloudWatch retains metric data as follows:
//   - Data points with a period of less than 60 seconds are available for 3
//     hours. These data points are high-resolution metrics and are available only for
//     custom metrics that have been defined with a StorageResolution of 1.
//   - Data points with a period of 60 seconds (1-minute) are available for 15
//     days.
//   - Data points with a period of 300 seconds (5-minute) are available for 63
//     days.
//   - Data points with a period of 3600 seconds (1 hour) are available for 455
//     days (15 months).
//
// Data points that are initially published with a shorter period are aggregated
// together for long-term storage. For example, if you collect data using a period
// of 1 minute, the data remains available for 15 days with 1-minute resolution.
// After 15 days, this data is still available, but is aggregated and retrievable
// only with a resolution of 5 minutes. After 63 days, the data is further
// aggregated and is available with a resolution of 1 hour. CloudWatch started
// retaining 5-minute and 1-hour metric data as of July 9, 2016. For information
// about metrics and dimensions supported by Amazon Web Services services, see the
// Amazon CloudWatch Metrics and Dimensions Reference (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CW_Support_For_AWS.html)
// in the Amazon CloudWatch User Guide.
func (c *Client) GetMetricStatistics(ctx context.Context, params *GetMetricStatisticsInput, optFns ...func(*Options)) (*GetMetricStatisticsOutput, error) {
	if params == nil {
		params = &GetMetricStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMetricStatistics", params, optFns, c.addOperationGetMetricStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMetricStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMetricStatisticsInput struct {

	// The time stamp that determines the last data point to return. The value
	// specified is exclusive; results include data points up to the specified time
	// stamp. In a raw HTTP query, the time stamp must be in ISO 8601 UTC format (for
	// example, 2016-10-10T23:00:00Z).
	//
	// This member is required.
	EndTime *time.Time

	// The name of the metric, with or without spaces.
	//
	// This member is required.
	MetricName *string

	// The namespace of the metric, with or without spaces.
	//
	// This member is required.
	Namespace *string

	// The granularity, in seconds, of the returned data points. For metrics with
	// regular resolution, a period can be as short as one minute (60 seconds) and must
	// be a multiple of 60. For high-resolution metrics that are collected at intervals
	// of less than one minute, the period can be 1, 5, 10, 30, 60, or any multiple of
	// 60. High-resolution metrics are those metrics stored by a PutMetricData call
	// that includes a StorageResolution of 1 second. If the StartTime parameter
	// specifies a time stamp that is greater than 3 hours ago, you must specify the
	// period as follows or no data points in that time range is returned:
	//   - Start time between 3 hours and 15 days ago - Use a multiple of 60 seconds
	//   (1 minute).
	//   - Start time between 15 and 63 days ago - Use a multiple of 300 seconds (5
	//   minutes).
	//   - Start time greater than 63 days ago - Use a multiple of 3600 seconds (1
	//   hour).
	//
	// This member is required.
	Period *int32

	// The time stamp that determines the first data point to return. Start times are
	// evaluated relative to the time that CloudWatch receives the request. The value
	// specified is inclusive; results include data points with the specified time
	// stamp. In a raw HTTP query, the time stamp must be in ISO 8601 UTC format (for
	// example, 2016-10-03T23:00:00Z). CloudWatch rounds the specified time stamp as
	// follows:
	//   - Start time less than 15 days ago - Round down to the nearest whole minute.
	//   For example, 12:32:34 is rounded down to 12:32:00.
	//   - Start time between 15 and 63 days ago - Round down to the nearest 5-minute
	//   clock interval. For example, 12:32:34 is rounded down to 12:30:00.
	//   - Start time greater than 63 days ago - Round down to the nearest 1-hour
	//   clock interval. For example, 12:32:34 is rounded down to 12:00:00.
	// If you set Period to 5, 10, or 30, the start time of your request is rounded
	// down to the nearest time that corresponds to even 5-, 10-, or 30-second
	// divisions of a minute. For example, if you make a query at (HH:mm:ss) 01:05:23
	// for the previous 10-second period, the start time of your request is rounded
	// down and you receive data from 01:05:10 to 01:05:20. If you make a query at
	// 15:07:17 for the previous 5 minutes of data, using a period of 5 seconds, you
	// receive data timestamped between 15:02:15 and 15:07:15.
	//
	// This member is required.
	StartTime *time.Time

	// The dimensions. If the metric contains multiple dimensions, you must include a
	// value for each dimension. CloudWatch treats each unique combination of
	// dimensions as a separate metric. If a specific combination of dimensions was not
	// published, you can't retrieve statistics for it. You must specify the same
	// dimensions that were used when the metrics were created. For an example, see
	// Dimension Combinations (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#dimension-combinations)
	// in the Amazon CloudWatch User Guide. For more information about specifying
	// dimensions, see Publishing Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html)
	// in the Amazon CloudWatch User Guide.
	Dimensions []types.Dimension

	// The percentile statistics. Specify values between p0.0 and p100. When calling
	// GetMetricStatistics , you must specify either Statistics or ExtendedStatistics ,
	// but not both. Percentile statistics are not available for metrics when any of
	// the metric values are negative numbers.
	ExtendedStatistics []string

	// The metric statistics, other than percentile. For percentile statistics, use
	// ExtendedStatistics . When calling GetMetricStatistics , you must specify either
	// Statistics or ExtendedStatistics , but not both.
	Statistics []types.Statistic

	// The unit for a given metric. If you omit Unit , all data that was collected with
	// any unit is returned, along with the corresponding units that were specified
	// when the data was reported to CloudWatch. If you specify a unit, the operation
	// returns only data that was collected with that unit specified. If you specify a
	// unit that does not match the data collected, the results of the operation are
	// null. CloudWatch does not perform unit conversions.
	Unit types.StandardUnit

	noSmithyDocumentSerde
}

type GetMetricStatisticsOutput struct {

	// The data points for the specified metric.
	Datapoints []types.Datapoint

	// A label for the specified metric.
	Label *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMetricStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetMetricStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetMetricStatistics{}, middleware.After)
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
	if err = addOpGetMetricStatisticsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMetricStatistics(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMetricStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "GetMetricStatistics",
	}
}
