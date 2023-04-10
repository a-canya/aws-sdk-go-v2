// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the current service opt-in settings for the Region. If service-opt-in
// is enabled for a service, Backup tries to protect that service's resources in
// this Region, when the resource is included in an on-demand backup or scheduled
// backup plan. Otherwise, Backup does not try to protect that service's resources
// in this Region. Use the DescribeRegionSettings API to determine the resource
// types that are supported.
func (c *Client) UpdateRegionSettings(ctx context.Context, params *UpdateRegionSettingsInput, optFns ...func(*Options)) (*UpdateRegionSettingsOutput, error) {
	if params == nil {
		params = &UpdateRegionSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRegionSettings", params, optFns, c.addOperationUpdateRegionSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRegionSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRegionSettingsInput struct {

	// Enables or disables full Backup management of backups for a resource type. To
	// enable full Backup management for DynamoDB along with Backup's advanced
	// DynamoDB backup features (https://docs.aws.amazon.com/aws-backup/latest/devguide/advanced-ddb-backup.html)
	// , follow the procedure to enable advanced DynamoDB backup programmatically (https://docs.aws.amazon.com/aws-backup/latest/devguide/advanced-ddb-backup.html#advanced-ddb-backup-enable-cli)
	// .
	ResourceTypeManagementPreference map[string]bool

	// Updates the list of services along with the opt-in preferences for the Region.
	ResourceTypeOptInPreference map[string]bool

	noSmithyDocumentSerde
}

type UpdateRegionSettingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRegionSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRegionSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRegionSettings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRegionSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRegionSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "UpdateRegionSettings",
	}
}
