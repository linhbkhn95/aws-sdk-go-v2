// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeartifact

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetDomainPermissionsPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain to which the resource policy is attached.
	//
	// Domain is a required field
	Domain *string `location:"querystring" locationName:"domain" min:"2" type:"string" required:"true"`

	// The 12-digit account number of the AWS account that owns the domain. It does
	// not include dashes or spaces.
	DomainOwner *string `location:"querystring" locationName:"domain-owner" min:"12" type:"string"`
}

// String returns the string representation
func (s GetDomainPermissionsPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDomainPermissionsPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDomainPermissionsPolicyInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 2))
	}
	if s.DomainOwner != nil && len(*s.DomainOwner) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainOwner", 12))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDomainPermissionsPolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Domain != nil {
		v := *s.Domain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "domain", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainOwner != nil {
		v := *s.DomainOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "domain-owner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetDomainPermissionsPolicyOutput struct {
	_ struct{} `type:"structure"`

	// The returned resource policy.
	Policy *ResourcePolicy `locationName:"policy" type:"structure"`
}

// String returns the string representation
func (s GetDomainPermissionsPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDomainPermissionsPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Policy != nil {
		v := s.Policy

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "policy", v, metadata)
	}
	return nil
}

const opGetDomainPermissionsPolicy = "GetDomainPermissionsPolicy"

// GetDomainPermissionsPolicyRequest returns a request value for making API operation for
// CodeArtifact.
//
// Returns the resource policy attached to the specified domain.
//
// The policy is a resource-based policy, not an identity-based policy. For
// more information, see Identity-based policies and resource-based policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_identity-vs-resource.html)
// in the AWS Identity and Access Management User Guide.
//
//    // Example sending a request using GetDomainPermissionsPolicyRequest.
//    req := client.GetDomainPermissionsPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeartifact-2018-09-22/GetDomainPermissionsPolicy
func (c *Client) GetDomainPermissionsPolicyRequest(input *GetDomainPermissionsPolicyInput) GetDomainPermissionsPolicyRequest {
	op := &aws.Operation{
		Name:       opGetDomainPermissionsPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/domain/permissions/policy",
	}

	if input == nil {
		input = &GetDomainPermissionsPolicyInput{}
	}

	req := c.newRequest(op, input, &GetDomainPermissionsPolicyOutput{})

	return GetDomainPermissionsPolicyRequest{Request: req, Input: input, Copy: c.GetDomainPermissionsPolicyRequest}
}

// GetDomainPermissionsPolicyRequest is the request type for the
// GetDomainPermissionsPolicy API operation.
type GetDomainPermissionsPolicyRequest struct {
	*aws.Request
	Input *GetDomainPermissionsPolicyInput
	Copy  func(*GetDomainPermissionsPolicyInput) GetDomainPermissionsPolicyRequest
}

// Send marshals and sends the GetDomainPermissionsPolicy API request.
func (r GetDomainPermissionsPolicyRequest) Send(ctx context.Context) (*GetDomainPermissionsPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDomainPermissionsPolicyResponse{
		GetDomainPermissionsPolicyOutput: r.Request.Data.(*GetDomainPermissionsPolicyOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDomainPermissionsPolicyResponse is the response type for the
// GetDomainPermissionsPolicy API operation.
type GetDomainPermissionsPolicyResponse struct {
	*GetDomainPermissionsPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDomainPermissionsPolicy request.
func (r *GetDomainPermissionsPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
