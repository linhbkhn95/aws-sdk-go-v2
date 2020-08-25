// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation has three possible return values:
//
//     * A successful response in
// the form of GreetingWithErrorsOutput
//
//     * An InvalidGreeting error.
//
//     * A
// BadRequest error.
func (c *Client) GreetingWithErrors(ctx context.Context, params *GreetingWithErrorsInput, optFns ...func(*Options)) (*GreetingWithErrorsOutput, error) {
	stack := middleware.NewStack("GreetingWithErrors", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGreetingWithErrorsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addServiceUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGreetingWithErrors(options.Region), middleware.Before)

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
			OperationName: "GreetingWithErrors",
			Err:           err,
		}
	}
	out := result.(*GreetingWithErrorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GreetingWithErrorsInput struct {
}

type GreetingWithErrorsOutput struct {
	Greeting *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGreetingWithErrorsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGreetingWithErrors{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGreetingWithErrors{}, middleware.After)
}

func newServiceMetadataMiddleware_opGreetingWithErrors(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceName:   "Query Protocol",
		ServiceID:     "QueryProtocol",
		OperationName: "GreetingWithErrors",
	}
}
