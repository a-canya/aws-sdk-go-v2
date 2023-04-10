// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Get detailed data for an environment.
func (c *Client) GetEnvironment(ctx context.Context, params *GetEnvironmentInput, optFns ...func(*Options)) (*GetEnvironmentOutput, error) {
	if params == nil {
		params = &GetEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEnvironment", params, optFns, c.addOperationGetEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEnvironmentInput struct {

	// The name of the environment that you want to get the detailed data for.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type GetEnvironmentOutput struct {

	// The detailed data of the requested environment.
	//
	// This member is required.
	Environment *types.Environment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetEnvironment{}, middleware.After)
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
	if err = addOpGetEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEnvironment(options.Region), middleware.Before); err != nil {
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

// GetEnvironmentAPIClient is a client that implements the GetEnvironment
// operation.
type GetEnvironmentAPIClient interface {
	GetEnvironment(context.Context, *GetEnvironmentInput, ...func(*Options)) (*GetEnvironmentOutput, error)
}

var _ GetEnvironmentAPIClient = (*Client)(nil)

// EnvironmentDeployedWaiterOptions are waiter options for
// EnvironmentDeployedWaiter
type EnvironmentDeployedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// EnvironmentDeployedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, EnvironmentDeployedWaiter will use default max delay of 4999
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetEnvironmentInput, *GetEnvironmentOutput, error) (bool, error)
}

// EnvironmentDeployedWaiter defines the waiters for EnvironmentDeployed
type EnvironmentDeployedWaiter struct {
	client GetEnvironmentAPIClient

	options EnvironmentDeployedWaiterOptions
}

// NewEnvironmentDeployedWaiter constructs a EnvironmentDeployedWaiter.
func NewEnvironmentDeployedWaiter(client GetEnvironmentAPIClient, optFns ...func(*EnvironmentDeployedWaiterOptions)) *EnvironmentDeployedWaiter {
	options := EnvironmentDeployedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 4999 * time.Second
	options.Retryable = environmentDeployedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &EnvironmentDeployedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for EnvironmentDeployed waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *EnvironmentDeployedWaiter) Wait(ctx context.Context, params *GetEnvironmentInput, maxWaitDur time.Duration, optFns ...func(*EnvironmentDeployedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for EnvironmentDeployed waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *EnvironmentDeployedWaiter) WaitForOutput(ctx context.Context, params *GetEnvironmentInput, maxWaitDur time.Duration, optFns ...func(*EnvironmentDeployedWaiterOptions)) (*GetEnvironmentOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 4999 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetEnvironment(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for EnvironmentDeployed waiter")
}

func environmentDeployedStateRetryable(ctx context.Context, input *GetEnvironmentInput, output *GetEnvironmentOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("environment.deploymentStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "SUCCEEDED"
		value, ok := pathValue.(types.DeploymentStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.DeploymentStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("environment.deploymentStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.DeploymentStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.DeploymentStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "GetEnvironment",
	}
}
