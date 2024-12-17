// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified message from the specified queue. To select the message
// to delete, use the ReceiptHandle of the message (not the MessageId which you
// receive when you send the message). Amazon SQS can delete a message from a queue
// even if a visibility timeout setting causes the message to be locked by another
// consumer. Amazon SQS automatically deletes messages left in a queue longer than
// the retention period configured for the queue.
//
// The ReceiptHandle is associated with a specific instance of receiving a
// message. If you receive a message more than once, the ReceiptHandle is
// different each time you receive a message. When you use the DeleteMessage
// action, you must provide the most recently received ReceiptHandle for the
// message (otherwise, the request succeeds, but the message will not be deleted).
//
// For standard queues, it is possible to receive a message even after you delete
// it. This might happen on rare occasions if one of the servers which stores a
// copy of the message is unavailable when you send the request to delete the
// message. The copy remains on the server and might be returned to you during a
// subsequent receive request. You should ensure that your application is
// idempotent, so that receiving a message more than once does not cause issues.
func (c *Client) DeleteMessage(ctx context.Context, params *DeleteMessageInput, optFns ...func(*Options)) (*DeleteMessageOutput, error) {
	if params == nil {
		params = &DeleteMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMessage", params, optFns, c.addOperationDeleteMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteMessageInput struct {

	// The URL of the Amazon SQS queue from which messages are deleted.
	//
	// Queue URLs and names are case-sensitive.
	//
	// This member is required.
	QueueUrl *string

	// The receipt handle associated with the message to delete.
	//
	// This member is required.
	ReceiptHandle *string

	noSmithyDocumentSerde
}

type DeleteMessageOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteMessage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteMessage"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMessage(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeleteMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteMessage",
	}
}
