// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Schedules a service software update for an Amazon OpenSearch Service domain.
// For more information, see Service software updates in Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/service-software.html)
// .
func (c *Client) StartServiceSoftwareUpdate(ctx context.Context, params *StartServiceSoftwareUpdateInput, optFns ...func(*Options)) (*StartServiceSoftwareUpdateOutput, error) {
	if params == nil {
		params = &StartServiceSoftwareUpdateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartServiceSoftwareUpdate", params, optFns, c.addOperationStartServiceSoftwareUpdateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartServiceSoftwareUpdateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the request parameters to the StartServiceSoftwareUpdate
// operation.
type StartServiceSoftwareUpdateInput struct {

	// The name of the domain that you want to update to the latest service software.
	//
	// This member is required.
	DomainName *string

	// The Epoch timestamp when you want the service software update to start. You
	// only need to specify this parameter if you set ScheduleAt to TIMESTAMP .
	DesiredStartTime *int64

	// When to start the service software update.
	//   - NOW - Immediately schedules the update to happen in the current hour if
	//   there's capacity available.
	//   - TIMESTAMP - Lets you specify a custom date and time to apply the update. If
	//   you specify this value, you must also provide a value for DesiredStartTime .
	//   - OFF_PEAK_WINDOW - Marks the update to be picked up during an upcoming
	//   off-peak window. There's no guarantee that the update will happen during the
	//   next immediate window. Depending on capacity, it might happen in subsequent
	//   days.
	// Default: NOW if you don't specify a value for DesiredStartTime , and TIMESTAMP
	// if you do.
	ScheduleAt types.ScheduleAt

	noSmithyDocumentSerde
}

// Represents the output of a StartServiceSoftwareUpdate operation. Contains the
// status of the update.
type StartServiceSoftwareUpdateOutput struct {

	// The current status of the OpenSearch Service software update.
	ServiceSoftwareOptions *types.ServiceSoftwareOptions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartServiceSoftwareUpdateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartServiceSoftwareUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartServiceSoftwareUpdate{}, middleware.After)
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
	if err = addOpStartServiceSoftwareUpdateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartServiceSoftwareUpdate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartServiceSoftwareUpdate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "StartServiceSoftwareUpdate",
	}
}
