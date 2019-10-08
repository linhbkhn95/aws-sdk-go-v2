// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeExportConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// A list of continuous export ids to search for.
	ExportIds []string `locationName:"exportIds" type:"list"`

	// A number between 1 and 100 specifying the maximum number of continuous export
	// descriptions returned.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token from the previous call to describe-export-tasks.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeExportConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeExportConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	ExportsInfo []ExportInfo `locationName:"exportsInfo" type:"list"`

	// The token from the previous call to describe-export-tasks.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeExportConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeExportConfigurations = "DescribeExportConfigurations"

// DescribeExportConfigurationsRequest returns a request value for making API operation for
// AWS Application Discovery Service.
//
// DescribeExportConfigurations is deprecated. Use DescribeImportTasks (https://docs.aws.amazon.com/application-discovery/latest/APIReference/API_DescribeExportTasks.html),
// instead.
//
//    // Example sending a request using DescribeExportConfigurationsRequest.
//    req := client.DescribeExportConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/DescribeExportConfigurations
func (c *Client) DescribeExportConfigurationsRequest(input *DescribeExportConfigurationsInput) DescribeExportConfigurationsRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, DescribeExportConfigurations, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opDescribeExportConfigurations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeExportConfigurationsInput{}
	}

	req := c.newRequest(op, input, &DescribeExportConfigurationsOutput{})
	return DescribeExportConfigurationsRequest{Request: req, Input: input, Copy: c.DescribeExportConfigurationsRequest}
}

// DescribeExportConfigurationsRequest is the request type for the
// DescribeExportConfigurations API operation.
type DescribeExportConfigurationsRequest struct {
	*aws.Request
	Input *DescribeExportConfigurationsInput
	Copy  func(*DescribeExportConfigurationsInput) DescribeExportConfigurationsRequest
}

// Send marshals and sends the DescribeExportConfigurations API request.
func (r DescribeExportConfigurationsRequest) Send(ctx context.Context) (*DescribeExportConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeExportConfigurationsResponse{
		DescribeExportConfigurationsOutput: r.Request.Data.(*DescribeExportConfigurationsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeExportConfigurationsResponse is the response type for the
// DescribeExportConfigurations API operation.
type DescribeExportConfigurationsResponse struct {
	*DescribeExportConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeExportConfigurations request.
func (r *DescribeExportConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}