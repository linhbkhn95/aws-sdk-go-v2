// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeApplicationStateInput struct {
	_ struct{} `type:"structure"`

	// The configurationId in ADS that uniquely identifies the grouped application.
	//
	// ApplicationId is a required field
	ApplicationId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeApplicationStateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeApplicationStateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeApplicationStateInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}
	if s.ApplicationId != nil && len(*s.ApplicationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeApplicationStateOutput struct {
	_ struct{} `type:"structure"`

	// Status of the application - Not Started, In-Progress, Complete.
	ApplicationStatus ApplicationStatus `type:"string" enum:"true"`

	// The timestamp when the application status was last updated.
	LastUpdatedTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s DescribeApplicationStateOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeApplicationState = "DescribeApplicationState"

// DescribeApplicationStateRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Gets the migration status of an application.
//
//    // Example sending a request using DescribeApplicationStateRequest.
//    req := client.DescribeApplicationStateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/DescribeApplicationState
func (c *Client) DescribeApplicationStateRequest(input *DescribeApplicationStateInput) DescribeApplicationStateRequest {
	op := &aws.Operation{
		Name:       opDescribeApplicationState,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeApplicationStateInput{}
	}

	req := c.newRequest(op, input, &DescribeApplicationStateOutput{})
	return DescribeApplicationStateRequest{Request: req, Input: input, Copy: c.DescribeApplicationStateRequest}
}

// DescribeApplicationStateRequest is the request type for the
// DescribeApplicationState API operation.
type DescribeApplicationStateRequest struct {
	*aws.Request
	Input *DescribeApplicationStateInput
	Copy  func(*DescribeApplicationStateInput) DescribeApplicationStateRequest
}

// Send marshals and sends the DescribeApplicationState API request.
func (r DescribeApplicationStateRequest) Send(ctx context.Context) (*DescribeApplicationStateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeApplicationStateResponse{
		DescribeApplicationStateOutput: r.Request.Data.(*DescribeApplicationStateOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeApplicationStateResponse is the response type for the
// DescribeApplicationState API operation.
type DescribeApplicationStateResponse struct {
	*DescribeApplicationStateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeApplicationState request.
func (r *DescribeApplicationStateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}