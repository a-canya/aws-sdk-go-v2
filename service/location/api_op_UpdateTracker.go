// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified properties of a given tracker resource.
func (c *Client) UpdateTracker(ctx context.Context, params *UpdateTrackerInput, optFns ...func(*Options)) (*UpdateTrackerOutput, error) {
	if params == nil {
		params = &UpdateTrackerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTracker", params, optFns, c.addOperationUpdateTrackerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTrackerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTrackerInput struct {

	// The name of the tracker resource to update.
	//
	// This member is required.
	TrackerName *string

	// Updates the description for the tracker resource.
	Description *string

	// Updates the position filtering for the tracker resource. Valid values:
	//   - TimeBased - Location updates are evaluated against linked geofence
	//   collections, but not every location update is stored. If your update frequency
	//   is more often than 30 seconds, only one update per 30 seconds is stored for each
	//   unique device ID.
	//   - DistanceBased - If the device has moved less than 30 m (98.4 ft), location
	//   updates are ignored. Location updates within this distance are neither evaluated
	//   against linked geofence collections, nor stored. This helps control costs by
	//   reducing the number of geofence evaluations and historical device positions to
	//   paginate through. Distance-based filtering can also reduce the effects of GPS
	//   noise when displaying device trajectories on a map.
	//   - AccuracyBased - If the device has moved less than the measured accuracy,
	//   location updates are ignored. For example, if two consecutive updates from a
	//   device have a horizontal accuracy of 5 m and 10 m, the second update is ignored
	//   if the device has moved less than 15 m. Ignored location updates are neither
	//   evaluated against linked geofence collections, nor stored. This helps educe the
	//   effects of GPS noise when displaying device trajectories on a map, and can help
	//   control costs by reducing the number of geofence evaluations.
	PositionFiltering types.PositionFiltering

	// No longer used. If included, the only allowed value is RequestBasedUsage .
	//
	// Deprecated: Deprecated. If included, the only allowed value is
	// RequestBasedUsage.
	PricingPlan types.PricingPlan

	// This parameter is no longer used.
	//
	// Deprecated: Deprecated. No longer allowed.
	PricingPlanDataSource *string

	noSmithyDocumentSerde
}

type UpdateTrackerOutput struct {

	// The Amazon Resource Name (ARN) of the updated tracker resource. Used to specify
	// a resource across AWS.
	//   - Format example: arn:aws:geo:region:account-id:tracker/ExampleTracker
	//
	// This member is required.
	TrackerArn *string

	// The name of the updated tracker resource.
	//
	// This member is required.
	TrackerName *string

	// The timestamp for when the tracker resource was last updated in  ISO 8601 (https://www.iso.org/iso-8601-date-and-time-format.html)
	// format: YYYY-MM-DDThh:mm:ss.sssZ .
	//
	// This member is required.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTrackerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTracker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTracker{}, middleware.After)
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
	if err = addEndpointPrefix_opUpdateTrackerMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateTrackerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTracker(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opUpdateTrackerMiddleware struct {
}

func (*endpointPrefix_opUpdateTrackerMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateTrackerMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "tracking." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opUpdateTrackerMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opUpdateTrackerMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateTracker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "UpdateTracker",
	}
}
