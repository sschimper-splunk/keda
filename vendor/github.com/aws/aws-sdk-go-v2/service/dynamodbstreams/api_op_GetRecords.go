// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodbstreams

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the stream records from a given shard.
//
// Specify a shard iterator using the ShardIterator parameter. The shard iterator
// specifies the position in the shard from which you want to start reading stream
// records sequentially. If there are no stream records available in the portion of
// the shard that the iterator points to, GetRecords returns an empty list. Note
// that it might take multiple calls to get to a portion of the shard that contains
// stream records.
//
// GetRecords can retrieve a maximum of 1 MB of data or 1000 stream records,
// whichever comes first.
func (c *Client) GetRecords(ctx context.Context, params *GetRecordsInput, optFns ...func(*Options)) (*GetRecordsOutput, error) {
	if params == nil {
		params = &GetRecordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRecords", params, optFns, c.addOperationGetRecordsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetRecords operation.
type GetRecordsInput struct {

	// A shard iterator that was retrieved from a previous GetShardIterator operation.
	// This iterator can be used to access the stream records in this shard.
	//
	// This member is required.
	ShardIterator *string

	// The maximum number of records to return from the shard. The upper limit is 1000.
	Limit *int32

	noSmithyDocumentSerde
}

// Represents the output of a GetRecords operation.
type GetRecordsOutput struct {

	// The next position in the shard from which to start sequentially reading stream
	// records. If set to null , the shard has been closed and the requested iterator
	// will not return any more data.
	NextShardIterator *string

	// The stream records from the shard, which were retrieved using the shard
	// iterator.
	Records []types.Record

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRecordsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetRecords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetRecords{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRecords"); err != nil {
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
	if err = addOpGetRecordsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRecords(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRecords(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRecords",
	}
}
