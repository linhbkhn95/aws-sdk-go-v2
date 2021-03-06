// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateConferenceProviderInput struct {
	_ struct{} `type:"structure"`

	// The request token of the client.
	ClientRequestToken *string `min:"10" type:"string" idempotencyToken:"true"`

	// The name of the conference provider.
	//
	// ConferenceProviderName is a required field
	ConferenceProviderName *string `min:"1" type:"string" required:"true"`

	// Represents a type within a list of predefined types.
	//
	// ConferenceProviderType is a required field
	ConferenceProviderType ConferenceProviderType `type:"string" required:"true" enum:"true"`

	// The IP endpoint and protocol for calling.
	IPDialIn *IPDialIn `type:"structure"`

	// The meeting settings for the conference provider.
	//
	// MeetingSetting is a required field
	MeetingSetting *MeetingSetting `type:"structure" required:"true"`

	// The information for PSTN conferencing.
	PSTNDialIn *PSTNDialIn `type:"structure"`
}

// String returns the string representation
func (s CreateConferenceProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateConferenceProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateConferenceProviderInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 10))
	}

	if s.ConferenceProviderName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConferenceProviderName"))
	}
	if s.ConferenceProviderName != nil && len(*s.ConferenceProviderName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConferenceProviderName", 1))
	}
	if len(s.ConferenceProviderType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ConferenceProviderType"))
	}

	if s.MeetingSetting == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeetingSetting"))
	}
	if s.IPDialIn != nil {
		if err := s.IPDialIn.Validate(); err != nil {
			invalidParams.AddNested("IPDialIn", err.(aws.ErrInvalidParams))
		}
	}
	if s.MeetingSetting != nil {
		if err := s.MeetingSetting.Validate(); err != nil {
			invalidParams.AddNested("MeetingSetting", err.(aws.ErrInvalidParams))
		}
	}
	if s.PSTNDialIn != nil {
		if err := s.PSTNDialIn.Validate(); err != nil {
			invalidParams.AddNested("PSTNDialIn", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateConferenceProviderOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the newly-created conference provider.
	ConferenceProviderArn *string `type:"string"`
}

// String returns the string representation
func (s CreateConferenceProviderOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateConferenceProvider = "CreateConferenceProvider"

// CreateConferenceProviderRequest returns a request value for making API operation for
// Alexa For Business.
//
// Adds a new conference provider under the user's AWS account.
//
//    // Example sending a request using CreateConferenceProviderRequest.
//    req := client.CreateConferenceProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/CreateConferenceProvider
func (c *Client) CreateConferenceProviderRequest(input *CreateConferenceProviderInput) CreateConferenceProviderRequest {
	op := &aws.Operation{
		Name:       opCreateConferenceProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateConferenceProviderInput{}
	}

	req := c.newRequest(op, input, &CreateConferenceProviderOutput{})

	return CreateConferenceProviderRequest{Request: req, Input: input, Copy: c.CreateConferenceProviderRequest}
}

// CreateConferenceProviderRequest is the request type for the
// CreateConferenceProvider API operation.
type CreateConferenceProviderRequest struct {
	*aws.Request
	Input *CreateConferenceProviderInput
	Copy  func(*CreateConferenceProviderInput) CreateConferenceProviderRequest
}

// Send marshals and sends the CreateConferenceProvider API request.
func (r CreateConferenceProviderRequest) Send(ctx context.Context) (*CreateConferenceProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateConferenceProviderResponse{
		CreateConferenceProviderOutput: r.Request.Data.(*CreateConferenceProviderOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateConferenceProviderResponse is the response type for the
// CreateConferenceProvider API operation.
type CreateConferenceProviderResponse struct {
	*CreateConferenceProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateConferenceProvider request.
func (r *CreateConferenceProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
