// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a job that uses workers to label the data objects in your input
// dataset. You can use the labeled data to train machine learning models. You can
// select your workforce from one of three providers:
//   - A private workforce that you create. It can include employees, contractors,
//     and outside experts. Use a private workforce when want the data to stay within
//     your organization or when a specific set of skills is required.
//   - One or more vendors that you select from the Amazon Web Services
//     Marketplace. Vendors provide expertise in specific areas.
//   - The Amazon Mechanical Turk workforce. This is the largest workforce, but it
//     should only be used for public data or data that has been stripped of any
//     personally identifiable information.
//
// You can also use automated data labeling to reduce the number of data objects
// that need to be labeled by a human. Automated data labeling uses active learning
// to determine if a data object can be labeled by machine or if it needs to be
// sent to a human worker. For more information, see Using Automated Data Labeling (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-automated-labeling.html)
// . The data objects to be labeled are contained in an Amazon S3 bucket. You
// create a manifest file that describes the location of each object. For more
// information, see Using Input and Output Data (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-data.html)
// . The output can be used as the manifest file for another labeling job or as
// training data for your machine learning models. You can use this operation to
// create a static labeling job or a streaming labeling job. A static labeling job
// stops if all data objects in the input manifest file identified in ManifestS3Uri
// have been labeled. A streaming labeling job runs perpetually until it is
// manually stopped, or remains idle for 10 days. You can send new data objects to
// an active ( InProgress ) streaming labeling job in real time. To learn how to
// create a static labeling job, see Create a Labeling Job (API)  (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-create-labeling-job-api.html)
// in the Amazon SageMaker Developer Guide. To learn how to create a streaming
// labeling job, see Create a Streaming Labeling Job (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-streaming-create-job.html)
// .
func (c *Client) CreateLabelingJob(ctx context.Context, params *CreateLabelingJobInput, optFns ...func(*Options)) (*CreateLabelingJobOutput, error) {
	if params == nil {
		params = &CreateLabelingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLabelingJob", params, optFns, c.addOperationCreateLabelingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLabelingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLabelingJobInput struct {

	// Configures the labeling task and how it is presented to workers; including, but
	// not limited to price, keywords, and batch size (task count).
	//
	// This member is required.
	HumanTaskConfig *types.HumanTaskConfig

	// Input data for the labeling job, such as the Amazon S3 location of the data
	// objects and the location of the manifest file that describes the data objects.
	// You must specify at least one of the following: S3DataSource or SnsDataSource .
	//   - Use SnsDataSource to specify an SNS input topic for a streaming labeling
	//   job. If you do not specify and SNS input topic ARN, Ground Truth will create a
	//   one-time labeling job that stops after all data objects in the input manifest
	//   file have been labeled.
	//   - Use S3DataSource to specify an input manifest file for both streaming and
	//   one-time labeling jobs. Adding an S3DataSource is optional if you use
	//   SnsDataSource to create a streaming labeling job.
	// If you use the Amazon Mechanical Turk workforce, your input data should not
	// include confidential information, personal information or protected health
	// information. Use ContentClassifiers to specify that your data is free of
	// personally identifiable information and adult content.
	//
	// This member is required.
	InputConfig *types.LabelingJobInputConfig

	// The attribute name to use for the label in the output manifest file. This is
	// the key for the key/value pair formed with the label that a worker assigns to
	// the object. The LabelAttributeName must meet the following requirements.
	//   - The name can't end with "-metadata".
	//   - If you are using one of the following built-in task types (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-task-types.html)
	//   , the attribute name must end with "-ref". If the task type you are using is not
	//   listed below, the attribute name must not end with "-ref".
	//   - Image semantic segmentation ( SemanticSegmentation) , and adjustment (
	//   AdjustmentSemanticSegmentation ) and verification (
	//   VerificationSemanticSegmentation ) labeling jobs for this task type.
	//   - Video frame object detection ( VideoObjectDetection ), and adjustment and
	//   verification ( AdjustmentVideoObjectDetection ) labeling jobs for this task
	//   type.
	//   - Video frame object tracking ( VideoObjectTracking ), and adjustment and
	//   verification ( AdjustmentVideoObjectTracking ) labeling jobs for this task
	//   type.
	//   - 3D point cloud semantic segmentation ( 3DPointCloudSemanticSegmentation ),
	//   and adjustment and verification ( Adjustment3DPointCloudSemanticSegmentation )
	//   labeling jobs for this task type.
	//   - 3D point cloud object tracking ( 3DPointCloudObjectTracking ), and
	//   adjustment and verification ( Adjustment3DPointCloudObjectTracking ) labeling
	//   jobs for this task type.
	// If you are creating an adjustment or verification labeling job, you must use a
	// different LabelAttributeName than the one used in the original labeling job.
	// The original labeling job is the Ground Truth labeling job that produced the
	// labels that you want verified or adjusted. To learn more about adjustment and
	// verification labeling jobs, see Verify and Adjust Labels (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-verification-data.html)
	// .
	//
	// This member is required.
	LabelAttributeName *string

	// The name of the labeling job. This name is used to identify the job in a list
	// of labeling jobs. Labeling job names must be unique within an Amazon Web
	// Services account and region. LabelingJobName is not case sensitive. For
	// example, Example-job and example-job are considered the same labeling job name
	// by Ground Truth.
	//
	// This member is required.
	LabelingJobName *string

	// The location of the output data and the Amazon Web Services Key Management
	// Service key ID for the key used to encrypt the output data, if any.
	//
	// This member is required.
	OutputConfig *types.LabelingJobOutputConfig

	// The Amazon Resource Number (ARN) that Amazon SageMaker assumes to perform tasks
	// on your behalf during data labeling. You must grant this role the necessary
	// permissions so that Amazon SageMaker can successfully complete data labeling.
	//
	// This member is required.
	RoleArn *string

	// The S3 URI of the file, referred to as a label category configuration file,
	// that defines the categories used to label the data objects. For 3D point cloud
	// and video frame task types, you can add label category attributes and frame
	// attributes to your label category configuration file. To learn how, see Create
	// a Labeling Category Configuration File for 3D Point Cloud Labeling Jobs (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-point-cloud-label-category-config.html)
	// . For named entity recognition jobs, in addition to "labels" , you must provide
	// worker instructions in the label category configuration file using the
	// "instructions" parameter: "instructions": {"shortInstruction":"
	// Add header
	//
	//     Add Instructions
	// ", "fullInstruction":"Add additional instructions."} . For details and an
	// example, see Create a Named Entity Recognition Labeling Job (API)  (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-named-entity-recg.html#sms-creating-ner-api)
	// . For all other built-in task types (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-task-types.html)
	// and custom tasks (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-custom-templates.html)
	// , your label category configuration file must be a JSON file in the following
	// format. Identify the labels you want to use by replacing label_1 , label_2 , ...
	// , label_n with your label categories. {
	//     "document-version": "2018-11-28",
	//
	//     "labels": [{"label": "label_1"},{"label": "label_2"},...{"label": "label_n"}]
	// } Note the following about the label category configuration file:
	//   - For image classification and text classification (single and multi-label)
	//   you must specify at least two label categories. For all other task types, the
	//   minimum number of label categories required is one.
	//   - Each label category must be unique, you cannot specify duplicate label
	//   categories.
	//   - If you create a 3D point cloud or video frame adjustment or verification
	//   labeling job, you must include auditLabelAttributeName in the label category
	//   configuration. Use this parameter to enter the LabelAttributeName (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateLabelingJob.html#sagemaker-CreateLabelingJob-request-LabelAttributeName)
	//   of the labeling job you want to adjust or verify annotations of.
	LabelCategoryConfigS3Uri *string

	// Configures the information required to perform automated data labeling.
	LabelingJobAlgorithmsConfig *types.LabelingJobAlgorithmsConfig

	// A set of conditions for stopping the labeling job. If any of the conditions are
	// met, the job is automatically stopped. You can use these conditions to control
	// the cost of data labeling.
	StoppingConditions *types.LabelingJobStoppingConditions

	// An array of key/value pairs. For more information, see Using Cost Allocation
	// Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what)
	// in the Amazon Web Services Billing and Cost Management User Guide.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateLabelingJobOutput struct {

	// The Amazon Resource Name (ARN) of the labeling job. You use this ARN to
	// identify the labeling job.
	//
	// This member is required.
	LabelingJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLabelingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLabelingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLabelingJob{}, middleware.After)
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
	if err = addOpCreateLabelingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLabelingJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLabelingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateLabelingJob",
	}
}
