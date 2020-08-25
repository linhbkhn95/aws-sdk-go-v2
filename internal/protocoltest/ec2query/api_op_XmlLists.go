// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/ec2query/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This test case serializes XML lists for the following cases for both input and
// output:
//
//     * Normal XML lists.
//
//     * Normal XML sets.
//
//     * XML lists of
// lists.
//
//     * XML lists with @xmlName on its members
//
//     * Flattened XML
// lists.
//
//     * Flattened XML lists with @xmlName.
//
//     * Lists of structures.
func (c *Client) XmlLists(ctx context.Context, params *XmlListsInput, optFns ...func(*Options)) (*XmlListsOutput, error) {
	stack := middleware.NewStack("XmlLists", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpXmlListsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addServiceUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opXmlLists(options.Region), middleware.Before)

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
			OperationName: "XmlLists",
			Err:           err,
		}
	}
	out := result.(*XmlListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlListsInput struct {
}

type XmlListsOutput struct {
	StringList    []*string
	StringSet     []*string
	IntegerList   []*int32
	BooleanList   []*bool
	TimestampList []*time.Time
	EnumList      []types.FooEnum
	// A list of lists of strings.
	NestedStringList   [][]*string
	RenamedListMembers []*string
	FlattenedList      []*string
	FlattenedList2     []*string
	StructureList      []*types.StructureListMember

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpXmlListsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpXmlLists{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpXmlLists{}, middleware.After)
}

func newServiceMetadataMiddleware_opXmlLists(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceName:   "EC2 Protocol",
		ServiceID:     "EC2Protocol",
		OperationName: "XmlLists",
	}
}
