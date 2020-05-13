// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetInstancesInput struct {
	_ struct{} `type:"structure"`

	// The token to advance to the next page of results from your request.
	//
	// To get a page token, perform an initial GetInstances request. If your results
	// are paginated, the response will return a next page token that you can specify
	// as the page token in a subsequent request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetInstancesInput) String() string {
	return awsutil.Prettify(s)
}

type GetInstancesOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about your instances.
	Instances []Instance `locationName:"instances" type:"list"`

	// The token to advance to the next page of resutls from your request.
	//
	// A next page token is not returned if there are no more results to display.
	//
	// To get the next page of results, perform another GetInstances request and
	// specify the next page token using the pageToken parameter.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetInstances = "GetInstances"

// GetInstancesRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about all Amazon Lightsail virtual private servers, or
// instances.
//
//    // Example sending a request using GetInstancesRequest.
//    req := client.GetInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetInstances
func (c *Client) GetInstancesRequest(input *GetInstancesInput) GetInstancesRequest {
	op := &aws.Operation{
		Name:       opGetInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetInstancesInput{}
	}

	req := c.newRequest(op, input, &GetInstancesOutput{})

	return GetInstancesRequest{Request: req, Input: input, Copy: c.GetInstancesRequest}
}

// GetInstancesRequest is the request type for the
// GetInstances API operation.
type GetInstancesRequest struct {
	*aws.Request
	Input *GetInstancesInput
	Copy  func(*GetInstancesInput) GetInstancesRequest
}

// Send marshals and sends the GetInstances API request.
func (r GetInstancesRequest) Send(ctx context.Context) (*GetInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInstancesResponse{
		GetInstancesOutput: r.Request.Data.(*GetInstancesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInstancesResponse is the response type for the
// GetInstances API operation.
type GetInstancesResponse struct {
	*GetInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInstances request.
func (r *GetInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}