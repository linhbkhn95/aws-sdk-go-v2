// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateMemberToGroupInput struct {
	_ struct{} `type:"structure"`

	// The group to which the member (user or group) is associated.
	//
	// GroupId is a required field
	GroupId *string `min:"12" type:"string" required:"true"`

	// The member (user or group) to associate to the group.
	//
	// MemberId is a required field
	MemberId *string `min:"12" type:"string" required:"true"`

	// The organization under which the group exists.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateMemberToGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateMemberToGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateMemberToGroupInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}
	if s.GroupId != nil && len(*s.GroupId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupId", 12))
	}

	if s.MemberId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MemberId"))
	}
	if s.MemberId != nil && len(*s.MemberId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("MemberId", 12))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateMemberToGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateMemberToGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateMemberToGroup = "AssociateMemberToGroup"

// AssociateMemberToGroupRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Adds a member (user or group) to the group's set.
//
//    // Example sending a request using AssociateMemberToGroupRequest.
//    req := client.AssociateMemberToGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/AssociateMemberToGroup
func (c *Client) AssociateMemberToGroupRequest(input *AssociateMemberToGroupInput) AssociateMemberToGroupRequest {
	op := &aws.Operation{
		Name:       opAssociateMemberToGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateMemberToGroupInput{}
	}

	req := c.newRequest(op, input, &AssociateMemberToGroupOutput{})
	return AssociateMemberToGroupRequest{Request: req, Input: input, Copy: c.AssociateMemberToGroupRequest}
}

// AssociateMemberToGroupRequest is the request type for the
// AssociateMemberToGroup API operation.
type AssociateMemberToGroupRequest struct {
	*aws.Request
	Input *AssociateMemberToGroupInput
	Copy  func(*AssociateMemberToGroupInput) AssociateMemberToGroupRequest
}

// Send marshals and sends the AssociateMemberToGroup API request.
func (r AssociateMemberToGroupRequest) Send(ctx context.Context) (*AssociateMemberToGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateMemberToGroupResponse{
		AssociateMemberToGroupOutput: r.Request.Data.(*AssociateMemberToGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateMemberToGroupResponse is the response type for the
// AssociateMemberToGroup API operation.
type AssociateMemberToGroupResponse struct {
	*AssociateMemberToGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateMemberToGroup request.
func (r *AssociateMemberToGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}