// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a dashboard in an Amazon Web Services account. Updating a Dashboard
// creates a new dashboard version but does not immediately publish the new
// version. You can update the published version of a dashboard by using the
// UpdateDashboardPublishedVersion (https://docs.aws.amazon.com/quicksight/latest/APIReference/API_UpdateDashboardPublishedVersion.html)
// API operation.
func (c *Client) UpdateDashboard(ctx context.Context, params *UpdateDashboardInput, optFns ...func(*Options)) (*UpdateDashboardOutput, error) {
	if params == nil {
		params = &UpdateDashboardInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDashboard", params, optFns, c.addOperationUpdateDashboardMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDashboardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDashboardInput struct {

	// The ID of the Amazon Web Services account that contains the dashboard that
	// you're updating.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the dashboard.
	//
	// This member is required.
	DashboardId *string

	// The display name of the dashboard.
	//
	// This member is required.
	Name *string

	// Options for publishing the dashboard when you create it:
	//   - AvailabilityStatus for AdHocFilteringOption - This status can be either
	//   ENABLED or DISABLED . When this is set to DISABLED , Amazon QuickSight
	//   disables the left filter pane on the published dashboard, which can be used for
	//   ad hoc (one-time) filtering. This option is ENABLED by default.
	//   - AvailabilityStatus for ExportToCSVOption - This status can be either ENABLED
	//   or DISABLED . The visual option to export data to .CSV format isn't enabled
	//   when this is set to DISABLED . This option is ENABLED by default.
	//   - VisibilityState for SheetControlsOption - This visibility state can be
	//   either COLLAPSED or EXPANDED . This option is COLLAPSED by default.
	DashboardPublishOptions *types.DashboardPublishOptions

	// The definition of a dashboard. A definition is the data model of all features
	// in a Dashboard, Template, or Analysis.
	Definition *types.DashboardVersionDefinition

	// A structure that contains the parameters of the dashboard. These are parameter
	// overrides for a dashboard. A dashboard can have any type of parameters, and some
	// parameters might accept multiple values.
	Parameters *types.Parameters

	// The entity that you are using as a source when you update the dashboard. In
	// SourceEntity , you specify the type of object you're using as source. You can
	// only update a dashboard from a template, so you use a SourceTemplate entity. If
	// you need to update a dashboard from an analysis, first convert the analysis to a
	// template by using the CreateTemplate (https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateTemplate.html)
	// API operation. For SourceTemplate , specify the Amazon Resource Name (ARN) of
	// the source template. The SourceTemplate ARN can contain any Amazon Web Services
	// account and any Amazon QuickSight-supported Amazon Web Services Region. Use the
	// DataSetReferences entity within SourceTemplate to list the replacement datasets
	// for the placeholders listed in the original. The schema in each dataset must
	// match its placeholder.
	SourceEntity *types.DashboardSourceEntity

	// The Amazon Resource Name (ARN) of the theme that is being used for this
	// dashboard. If you add a value for this field, it overrides the value that was
	// originally associated with the entity. The theme ARN must exist in the same
	// Amazon Web Services account where you create the dashboard.
	ThemeArn *string

	// A description for the first version of the dashboard being created.
	VersionDescription *string

	noSmithyDocumentSerde
}

type UpdateDashboardOutput struct {

	// The Amazon Resource Name (ARN) of the resource.
	Arn *string

	// The creation status of the request.
	CreationStatus types.ResourceStatus

	// The ID for the dashboard.
	DashboardId *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// The ARN of the dashboard, including the version number.
	VersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDashboardMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDashboard{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDashboard{}, middleware.After)
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
	if err = addOpUpdateDashboardValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDashboard(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDashboard(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateDashboard",
	}
}
