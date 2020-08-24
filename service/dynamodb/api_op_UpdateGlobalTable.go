// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or removes replicas in the specified global table. The global table must
// already exist to be able to use this operation. Any replica to be added must be
// empty, have the same name as the global table, have the same key schema, have
// DynamoDB Streams enabled, and have the same provisioned and maximum write
// capacity units. Although you can use UpdateGlobalTable to add replicas and
// remove replicas in a single request, for simplicity we recommend that you issue
// separate requests for adding or removing replicas. If global secondary indexes
// are specified, then the following conditions must also be met:
//
//     * The global
// secondary indexes must have the same name.
//
//     * The global secondary indexes
// must have the same hash key and sort key (if present).
//
//     * The global
// secondary indexes must have the same provisioned and maximum write capacity
// units.
func (c *Client) UpdateGlobalTable(ctx context.Context, params *UpdateGlobalTableInput, optFns ...func(*Options)) (*UpdateGlobalTableOutput, error) {
	stack := middleware.NewStack("UpdateGlobalTable", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpUpdateGlobalTableMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	registerHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUpdateGlobalTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGlobalTable(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "UpdateGlobalTable",
			Err:           err,
		}
	}
	out := result.(*UpdateGlobalTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGlobalTableInput struct {
	// The global table name.
	GlobalTableName *string
	// A list of Regions that should be added or removed from the global table.
	ReplicaUpdates []*types.ReplicaUpdate
}

type UpdateGlobalTableOutput struct {
	// Contains the details of the global table.
	GlobalTableDescription *types.GlobalTableDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpUpdateGlobalTableMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateGlobalTable{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateGlobalTable{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateGlobalTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "DynamoDB",
		ServiceID:      "dynamodb",
		EndpointPrefix: "dynamodb",
		SigningName:    "dynamodb",
		OperationName:  "UpdateGlobalTable",
	}
}
