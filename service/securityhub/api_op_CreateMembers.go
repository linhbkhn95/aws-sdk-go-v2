// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateMembersInput struct {
	_ struct{} `type:"structure"`

	// The list of accounts to associate with the Security Hub master account. For
	// each account, the list includes the account ID and the email address.
	AccountDetails []AccountDetails `type:"list"`
}

// String returns the string representation
func (s CreateMembersInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMembersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountDetails != nil {
		v := s.AccountDetails

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "AccountDetails", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type CreateMembersOutput struct {
	_ struct{} `type:"structure"`

	// The list of AWS accounts that were not processed. For each account, the list
	// includes the account ID and the email address.
	UnprocessedAccounts []Result `type:"list"`
}

// String returns the string representation
func (s CreateMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.UnprocessedAccounts != nil {
		v := s.UnprocessedAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UnprocessedAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opCreateMembers = "CreateMembers"

// CreateMembersRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Creates a member association in Security Hub between the specified accounts
// and the account used to make the request, which is the master account. To
// successfully create a member, you must use this action from an account that
// already has Security Hub enabled. To enable Security Hub, you can use the
// EnableSecurityHub operation.
//
// After you use CreateMembers to create member account associations in Security
// Hub, you must use the InviteMembers operation to invite the accounts to enable
// Security Hub and become member accounts in Security Hub.
//
// If the account owner accepts the invitation, the account becomes a member
// account in Security Hub. A permissions policy is added that permits the master
// account to view the findings generated in the member account. When Security
// Hub is enabled in the invited account, findings start to be sent to both
// the member and master accounts.
//
// To remove the association between the master and member accounts, use the
// DisassociateFromMasterAccount or DisassociateMembers operation.
//
//    // Example sending a request using CreateMembersRequest.
//    req := client.CreateMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/CreateMembers
func (c *Client) CreateMembersRequest(input *CreateMembersInput) CreateMembersRequest {
	op := &aws.Operation{
		Name:       opCreateMembers,
		HTTPMethod: "POST",
		HTTPPath:   "/members",
	}

	if input == nil {
		input = &CreateMembersInput{}
	}

	req := c.newRequest(op, input, &CreateMembersOutput{})

	return CreateMembersRequest{Request: req, Input: input, Copy: c.CreateMembersRequest}
}

// CreateMembersRequest is the request type for the
// CreateMembers API operation.
type CreateMembersRequest struct {
	*aws.Request
	Input *CreateMembersInput
	Copy  func(*CreateMembersInput) CreateMembersRequest
}

// Send marshals and sends the CreateMembers API request.
func (r CreateMembersRequest) Send(ctx context.Context) (*CreateMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMembersResponse{
		CreateMembersOutput: r.Request.Data.(*CreateMembersOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMembersResponse is the response type for the
// CreateMembers API operation.
type CreateMembersResponse struct {
	*CreateMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMembers request.
func (r *CreateMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}