// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Uses the join token and call metadata in a meeting request (From number, To
// number, and so forth) to initiate an outbound call to a public switched
// telephone network (PSTN) and join them into a Chime meeting. Also ensures that
// the From number belongs to the customer. To play welcome audio or implement an
// interactive voice response (IVR), use the CreateSipMediaApplicationCall action
// with the corresponding SIP media application ID.
func (c *Client) CreateMeetingDialOut(ctx context.Context, params *CreateMeetingDialOutInput, optFns ...func(*Options)) (*CreateMeetingDialOutOutput, error) {
	if params == nil {
		params = &CreateMeetingDialOutInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMeetingDialOut", params, optFns, c.addOperationCreateMeetingDialOutMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMeetingDialOutOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMeetingDialOutInput struct {

	// Phone number used as the caller ID when the remote party receives a call.
	//
	// This member is required.
	FromPhoneNumber *string

	// Token used by the Amazon Chime SDK attendee. Call the CreateAttendee (https://docs.aws.amazon.com/chime/latest/APIReference/API_CreateAttendee.html)
	// action to get a join token.
	//
	// This member is required.
	JoinToken *string

	// The Amazon Chime SDK meeting ID.
	//
	// This member is required.
	MeetingId *string

	// Phone number called when inviting someone to a meeting.
	//
	// This member is required.
	ToPhoneNumber *string

	noSmithyDocumentSerde
}

type CreateMeetingDialOutOutput struct {

	// Unique ID that tracks API calls.
	TransactionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMeetingDialOutMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMeetingDialOut{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMeetingDialOut{}, middleware.After)
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
	if err = addOpCreateMeetingDialOutValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMeetingDialOut(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMeetingDialOut(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateMeetingDialOut",
	}
}
