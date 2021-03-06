// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListSchemaExtensionsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the directory from which to retrieve the schema extension
	// information.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`

	// The maximum number of items to return.
	Limit *int64 `type:"integer"`

	// The ListSchemaExtensions.NextToken value from a previous call to ListSchemaExtensions.
	// Pass null if this is the first call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListSchemaExtensionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSchemaExtensionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSchemaExtensionsInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListSchemaExtensionsOutput struct {
	_ struct{} `type:"structure"`

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to ListSchemaExtensions to retrieve the next
	// set of items.
	NextToken *string `type:"string"`

	// Information about the schema extensions applied to the directory.
	SchemaExtensionsInfo []SchemaExtensionInfo `type:"list"`
}

// String returns the string representation
func (s ListSchemaExtensionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListSchemaExtensions = "ListSchemaExtensions"

// ListSchemaExtensionsRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Lists all schema extensions applied to a Microsoft AD Directory.
//
//    // Example sending a request using ListSchemaExtensionsRequest.
//    req := client.ListSchemaExtensionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/ListSchemaExtensions
func (c *Client) ListSchemaExtensionsRequest(input *ListSchemaExtensionsInput) ListSchemaExtensionsRequest {
	op := &aws.Operation{
		Name:       opListSchemaExtensions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListSchemaExtensionsInput{}
	}

	req := c.newRequest(op, input, &ListSchemaExtensionsOutput{})

	return ListSchemaExtensionsRequest{Request: req, Input: input, Copy: c.ListSchemaExtensionsRequest}
}

// ListSchemaExtensionsRequest is the request type for the
// ListSchemaExtensions API operation.
type ListSchemaExtensionsRequest struct {
	*aws.Request
	Input *ListSchemaExtensionsInput
	Copy  func(*ListSchemaExtensionsInput) ListSchemaExtensionsRequest
}

// Send marshals and sends the ListSchemaExtensions API request.
func (r ListSchemaExtensionsRequest) Send(ctx context.Context) (*ListSchemaExtensionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSchemaExtensionsResponse{
		ListSchemaExtensionsOutput: r.Request.Data.(*ListSchemaExtensionsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListSchemaExtensionsResponse is the response type for the
// ListSchemaExtensions API operation.
type ListSchemaExtensionsResponse struct {
	*ListSchemaExtensionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSchemaExtensions request.
func (r *ListSchemaExtensionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
