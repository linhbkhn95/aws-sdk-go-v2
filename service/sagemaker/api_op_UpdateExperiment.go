// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateExperimentInput struct {
	_ struct{} `type:"structure"`

	// The description of the experiment.
	Description *string `type:"string"`

	// The name of the experiment as displayed. The name doesn't need to be unique.
	// If DisplayName isn't specified, ExperimentName is displayed.
	DisplayName *string `min:"1" type:"string"`

	// The name of the experiment to update.
	//
	// ExperimentName is a required field
	ExperimentName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateExperimentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateExperimentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateExperimentInput"}
	if s.DisplayName != nil && len(*s.DisplayName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DisplayName", 1))
	}

	if s.ExperimentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExperimentName"))
	}
	if s.ExperimentName != nil && len(*s.ExperimentName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ExperimentName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateExperimentOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the experiment.
	ExperimentArn *string `type:"string"`
}

// String returns the string representation
func (s UpdateExperimentOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateExperiment = "UpdateExperiment"

// UpdateExperimentRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Adds, updates, or removes the description of an experiment. Updates the display
// name of an experiment.
//
//    // Example sending a request using UpdateExperimentRequest.
//    req := client.UpdateExperimentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/UpdateExperiment
func (c *Client) UpdateExperimentRequest(input *UpdateExperimentInput) UpdateExperimentRequest {
	op := &aws.Operation{
		Name:       opUpdateExperiment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateExperimentInput{}
	}

	req := c.newRequest(op, input, &UpdateExperimentOutput{})

	return UpdateExperimentRequest{Request: req, Input: input, Copy: c.UpdateExperimentRequest}
}

// UpdateExperimentRequest is the request type for the
// UpdateExperiment API operation.
type UpdateExperimentRequest struct {
	*aws.Request
	Input *UpdateExperimentInput
	Copy  func(*UpdateExperimentInput) UpdateExperimentRequest
}

// Send marshals and sends the UpdateExperiment API request.
func (r UpdateExperimentRequest) Send(ctx context.Context) (*UpdateExperimentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateExperimentResponse{
		UpdateExperimentOutput: r.Request.Data.(*UpdateExperimentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateExperimentResponse is the response type for the
// UpdateExperiment API operation.
type UpdateExperimentResponse struct {
	*UpdateExperimentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateExperiment request.
func (r *UpdateExperimentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
