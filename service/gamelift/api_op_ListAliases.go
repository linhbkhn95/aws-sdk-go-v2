// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type ListAliasesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return. Use this parameter with NextToken
	// to get results as a set of sequential pages.
	Limit *int64 `min:"1" type:"integer"`

	// A descriptive label that is associated with an alias. Alias names do not
	// need to be unique.
	Name *string `min:"1" type:"string"`

	// A token that indicates the start of the next sequential page of results.
	// Use the token that is returned with a previous call to this action. To start
	// at the beginning of the result set, do not specify a value.
	NextToken *string `min:"1" type:"string"`

	// The routing type to filter results on. Use this parameter to retrieve only
	// aliases with a certain routing type. To retrieve all aliases, leave this
	// parameter empty.
	//
	// Possible routing types include the following:
	//
	//    * SIMPLE -- The alias resolves to one specific fleet. Use this type when
	//    routing to active fleets.
	//
	//    * TERMINAL -- The alias does not resolve to a fleet but instead can be
	//    used to display a message to the user. A terminal alias throws a TerminalRoutingStrategyException
	//    with the RoutingStrategy message embedded.
	RoutingStrategyType RoutingStrategyType `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListAliasesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAliasesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAliasesInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type ListAliasesOutput struct {
	_ struct{} `type:"structure"`

	// A collection of alias resources that match the request parameters.
	Aliases []Alias `type:"list"`

	// A token that indicates where to resume retrieving results on the next call
	// to this action. If no token is returned, these results represent the end
	// of the list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListAliasesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAliases = "ListAliases"

// ListAliasesRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves all aliases for this AWS account. You can filter the result set
// by alias name and/or routing strategy type. Use the pagination parameters
// to retrieve results in sequential pages.
//
// Returned aliases are not listed in any particular order.
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
//    // Example sending a request using ListAliasesRequest.
//    req := client.ListAliasesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/ListAliases
func (c *Client) ListAliasesRequest(input *ListAliasesInput) ListAliasesRequest {
	op := &aws.Operation{
		Name:       opListAliases,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAliasesInput{}
	}

	req := c.newRequest(op, input, &ListAliasesOutput{})

	return ListAliasesRequest{Request: req, Input: input, Copy: c.ListAliasesRequest}
}

// ListAliasesRequest is the request type for the
// ListAliases API operation.
type ListAliasesRequest struct {
	*aws.Request
	Input *ListAliasesInput
	Copy  func(*ListAliasesInput) ListAliasesRequest
}

// Send marshals and sends the ListAliases API request.
func (r ListAliasesRequest) Send(ctx context.Context) (*ListAliasesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAliasesResponse{
		ListAliasesOutput: r.Request.Data.(*ListAliasesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAliasesResponse is the response type for the
// ListAliases API operation.
type ListAliasesResponse struct {
	*ListAliasesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAliases request.
func (r *ListAliasesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
