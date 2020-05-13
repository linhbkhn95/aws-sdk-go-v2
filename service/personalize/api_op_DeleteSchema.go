// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteSchemaInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the schema to delete.
	//
	// SchemaArn is a required field
	SchemaArn *string `locationName:"schemaArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSchemaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSchemaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSchemaInput"}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteSchemaOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSchemaOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteSchema = "DeleteSchema"

// DeleteSchemaRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Deletes a schema. Before deleting a schema, you must delete all datasets
// referencing the schema. For more information on schemas, see CreateSchema.
//
//    // Example sending a request using DeleteSchemaRequest.
//    req := client.DeleteSchemaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DeleteSchema
func (c *Client) DeleteSchemaRequest(input *DeleteSchemaInput) DeleteSchemaRequest {
	op := &aws.Operation{
		Name:       opDeleteSchema,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteSchemaInput{}
	}

	req := c.newRequest(op, input, &DeleteSchemaOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteSchemaRequest{Request: req, Input: input, Copy: c.DeleteSchemaRequest}
}

// DeleteSchemaRequest is the request type for the
// DeleteSchema API operation.
type DeleteSchemaRequest struct {
	*aws.Request
	Input *DeleteSchemaInput
	Copy  func(*DeleteSchemaInput) DeleteSchemaRequest
}

// Send marshals and sends the DeleteSchema API request.
func (r DeleteSchemaRequest) Send(ctx context.Context) (*DeleteSchemaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSchemaResponse{
		DeleteSchemaOutput: r.Request.Data.(*DeleteSchemaOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSchemaResponse is the response type for the
// DeleteSchema API operation.
type DeleteSchemaResponse struct {
	*DeleteSchemaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSchema request.
func (r *DeleteSchemaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}