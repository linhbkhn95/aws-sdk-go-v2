// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type UpdateAliasInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a fleet alias. Specify the alias you want to update.
	//
	// AliasId is a required field
	AliasId *string `type:"string" required:"true"`

	// Human-readable description of an alias.
	Description *string `min:"1" type:"string"`

	// Descriptive label that is associated with an alias. Alias names do not need
	// to be unique.
	Name *string `min:"1" type:"string"`

	// Object that specifies the fleet and routing type to use for the alias.
	RoutingStrategy *RoutingStrategy `type:"structure"`
}

// String returns the string representation
func (s UpdateAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAliasInput"}

	if s.AliasId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AliasId"))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type UpdateAliasOutput struct {
	_ struct{} `type:"structure"`

	// Object that contains the updated alias configuration.
	Alias *Alias `type:"structure"`
}

// String returns the string representation
func (s UpdateAliasOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateAlias = "UpdateAlias"

// UpdateAliasRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Updates properties for an alias. To update properties, specify the alias
// ID to be updated and provide the information to be changed. To reassign an
// alias to another fleet, provide an updated routing strategy. If successful,
// the updated alias record is returned.
//
//    * CreateAlias
//
//    * ListAliases
//
//    * DescribeAlias
//
//    * UpdateAlias
//
//    * DeleteAlias
//
//    * ResolveAlias
//
//    // Example sending a request using UpdateAliasRequest.
//    req := client.UpdateAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/UpdateAlias
func (c *Client) UpdateAliasRequest(input *UpdateAliasInput) UpdateAliasRequest {
	op := &aws.Operation{
		Name:       opUpdateAlias,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAliasInput{}
	}

	req := c.newRequest(op, input, &UpdateAliasOutput{})
	return UpdateAliasRequest{Request: req, Input: input, Copy: c.UpdateAliasRequest}
}

// UpdateAliasRequest is the request type for the
// UpdateAlias API operation.
type UpdateAliasRequest struct {
	*aws.Request
	Input *UpdateAliasInput
	Copy  func(*UpdateAliasInput) UpdateAliasRequest
}

// Send marshals and sends the UpdateAlias API request.
func (r UpdateAliasRequest) Send(ctx context.Context) (*UpdateAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAliasResponse{
		UpdateAliasOutput: r.Request.Data.(*UpdateAliasOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAliasResponse is the response type for the
// UpdateAlias API operation.
type UpdateAliasResponse struct {
	*UpdateAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAlias request.
func (r *UpdateAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}