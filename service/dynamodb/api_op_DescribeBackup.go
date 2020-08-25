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

// Describes an existing backup of a table. You can call DescribeBackup at a
// maximum rate of 10 times per second.
func (c *Client) DescribeBackup(ctx context.Context, params *DescribeBackupInput, optFns ...func(*Options)) (*DescribeBackupOutput, error) {
	stack := middleware.NewStack("DescribeBackup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpDescribeBackupMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addServiceUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribeBackupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBackup(options.Region), middleware.Before)

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
			OperationName: "DescribeBackup",
			Err:           err,
		}
	}
	out := result.(*DescribeBackupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBackupInput struct {
	// The Amazon Resource Name (ARN) associated with the backup.
	BackupArn *string
}

type DescribeBackupOutput struct {
	// Contains the description of the backup created for the table.
	BackupDescription *types.BackupDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpDescribeBackupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeBackup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeBackup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeBackup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceName:   "DynamoDB",
		ServiceID:     "DynamoDB",
		SigningName:   "dynamodb",
		OperationName: "DescribeBackup",
	}
}
