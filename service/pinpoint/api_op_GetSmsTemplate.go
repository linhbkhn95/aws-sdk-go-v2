// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetSmsTemplateInput struct {
	_ struct{} `type:"structure"`

	// TemplateName is a required field
	TemplateName *string `location:"uri" locationName:"template-name" type:"string" required:"true"`

	Version *string `location:"querystring" locationName:"version" type:"string"`
}

// String returns the string representation
func (s GetSmsTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSmsTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSmsTemplateInput"}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSmsTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TemplateName != nil {
		v := *s.TemplateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "template-name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetSmsTemplateOutput struct {
	_ struct{} `type:"structure" payload:"SMSTemplateResponse"`

	// Provides information about the content and settings for a message template
	// that can be used in text messages that are sent through the SMS channel.
	//
	// SMSTemplateResponse is a required field
	SMSTemplateResponse *SMSTemplateResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetSmsTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSmsTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SMSTemplateResponse != nil {
		v := s.SMSTemplateResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "SMSTemplateResponse", v, metadata)
	}
	return nil
}

const opGetSmsTemplate = "GetSmsTemplate"

// GetSmsTemplateRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves the content and settings of a message template for messages that
// are sent through the SMS channel.
//
//    // Example sending a request using GetSmsTemplateRequest.
//    req := client.GetSmsTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetSmsTemplate
func (c *Client) GetSmsTemplateRequest(input *GetSmsTemplateInput) GetSmsTemplateRequest {
	op := &aws.Operation{
		Name:       opGetSmsTemplate,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/templates/{template-name}/sms",
	}

	if input == nil {
		input = &GetSmsTemplateInput{}
	}

	req := c.newRequest(op, input, &GetSmsTemplateOutput{})

	return GetSmsTemplateRequest{Request: req, Input: input, Copy: c.GetSmsTemplateRequest}
}

// GetSmsTemplateRequest is the request type for the
// GetSmsTemplate API operation.
type GetSmsTemplateRequest struct {
	*aws.Request
	Input *GetSmsTemplateInput
	Copy  func(*GetSmsTemplateInput) GetSmsTemplateRequest
}

// Send marshals and sends the GetSmsTemplate API request.
func (r GetSmsTemplateRequest) Send(ctx context.Context) (*GetSmsTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSmsTemplateResponse{
		GetSmsTemplateOutput: r.Request.Data.(*GetSmsTemplateOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSmsTemplateResponse is the response type for the
// GetSmsTemplate API operation.
type GetSmsTemplateResponse struct {
	*GetSmsTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSmsTemplate request.
func (r *GetSmsTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
