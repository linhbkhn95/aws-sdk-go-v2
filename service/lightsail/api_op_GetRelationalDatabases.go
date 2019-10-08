// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetRelationalDatabasesInput struct {
	_ struct{} `type:"structure"`

	// A token used for advancing to a specific page of results for your get relational
	// database request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetRelationalDatabasesInput) String() string {
	return awsutil.Prettify(s)
}

type GetRelationalDatabasesOutput struct {
	_ struct{} `type:"structure"`

	// A token used for advancing to the next page of results from your get relational
	// databases request.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`

	// An object describing the result of your get relational databases request.
	RelationalDatabases []RelationalDatabase `locationName:"relationalDatabases" type:"list"`
}

// String returns the string representation
func (s GetRelationalDatabasesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRelationalDatabases = "GetRelationalDatabases"

// GetRelationalDatabasesRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about all of your databases in Amazon Lightsail.
//
//    // Example sending a request using GetRelationalDatabasesRequest.
//    req := client.GetRelationalDatabasesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetRelationalDatabases
func (c *Client) GetRelationalDatabasesRequest(input *GetRelationalDatabasesInput) GetRelationalDatabasesRequest {
	op := &aws.Operation{
		Name:       opGetRelationalDatabases,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRelationalDatabasesInput{}
	}

	req := c.newRequest(op, input, &GetRelationalDatabasesOutput{})
	return GetRelationalDatabasesRequest{Request: req, Input: input, Copy: c.GetRelationalDatabasesRequest}
}

// GetRelationalDatabasesRequest is the request type for the
// GetRelationalDatabases API operation.
type GetRelationalDatabasesRequest struct {
	*aws.Request
	Input *GetRelationalDatabasesInput
	Copy  func(*GetRelationalDatabasesInput) GetRelationalDatabasesRequest
}

// Send marshals and sends the GetRelationalDatabases API request.
func (r GetRelationalDatabasesRequest) Send(ctx context.Context) (*GetRelationalDatabasesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRelationalDatabasesResponse{
		GetRelationalDatabasesOutput: r.Request.Data.(*GetRelationalDatabasesOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRelationalDatabasesResponse is the response type for the
// GetRelationalDatabases API operation.
type GetRelationalDatabasesResponse struct {
	*GetRelationalDatabasesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRelationalDatabases request.
func (r *GetRelationalDatabasesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}