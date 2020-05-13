// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateDashboardPermissionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the AWS account that contains the dashboard whose permissions you're
	// updating.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the dashboard.
	//
	// DashboardId is a required field
	DashboardId *string `location:"uri" locationName:"DashboardId" min:"1" type:"string" required:"true"`

	// The permissions that you want to grant on this resource.
	GrantPermissions []ResourcePermission `type:"list"`

	// The permissions that you want to revoke from this resource.
	RevokePermissions []ResourcePermission `type:"list"`
}

// String returns the string representation
func (s UpdateDashboardPermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDashboardPermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDashboardPermissionsInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DashboardId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashboardId"))
	}
	if s.DashboardId != nil && len(*s.DashboardId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DashboardId", 1))
	}
	if s.GrantPermissions != nil {
		for i, v := range s.GrantPermissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "GrantPermissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.RevokePermissions != nil {
		for i, v := range s.RevokePermissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RevokePermissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDashboardPermissionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GrantPermissions != nil {
		v := s.GrantPermissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "GrantPermissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RevokePermissions != nil {
		v := s.RevokePermissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "RevokePermissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DashboardId != nil {
		v := *s.DashboardId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DashboardId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateDashboardPermissionsOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dashboard.
	DashboardArn *string `type:"string"`

	// The ID for the dashboard.
	DashboardId *string `min:"1" type:"string"`

	// Information about the permissions on the dashboard.
	Permissions []ResourcePermission `min:"1" type:"list"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s UpdateDashboardPermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDashboardPermissionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DashboardArn != nil {
		v := *s.DashboardArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DashboardArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DashboardId != nil {
		v := *s.DashboardId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DashboardId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Permissions != nil {
		v := s.Permissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Permissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opUpdateDashboardPermissions = "UpdateDashboardPermissions"

// UpdateDashboardPermissionsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Updates read and write permissions on a dashboard.
//
//    // Example sending a request using UpdateDashboardPermissionsRequest.
//    req := client.UpdateDashboardPermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateDashboardPermissions
func (c *Client) UpdateDashboardPermissionsRequest(input *UpdateDashboardPermissionsInput) UpdateDashboardPermissionsRequest {
	op := &aws.Operation{
		Name:       opUpdateDashboardPermissions,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/dashboards/{DashboardId}/permissions",
	}

	if input == nil {
		input = &UpdateDashboardPermissionsInput{}
	}

	req := c.newRequest(op, input, &UpdateDashboardPermissionsOutput{})

	return UpdateDashboardPermissionsRequest{Request: req, Input: input, Copy: c.UpdateDashboardPermissionsRequest}
}

// UpdateDashboardPermissionsRequest is the request type for the
// UpdateDashboardPermissions API operation.
type UpdateDashboardPermissionsRequest struct {
	*aws.Request
	Input *UpdateDashboardPermissionsInput
	Copy  func(*UpdateDashboardPermissionsInput) UpdateDashboardPermissionsRequest
}

// Send marshals and sends the UpdateDashboardPermissions API request.
func (r UpdateDashboardPermissionsRequest) Send(ctx context.Context) (*UpdateDashboardPermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDashboardPermissionsResponse{
		UpdateDashboardPermissionsOutput: r.Request.Data.(*UpdateDashboardPermissionsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDashboardPermissionsResponse is the response type for the
// UpdateDashboardPermissions API operation.
type UpdateDashboardPermissionsResponse struct {
	*UpdateDashboardPermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDashboardPermissions request.
func (r *UpdateDashboardPermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}