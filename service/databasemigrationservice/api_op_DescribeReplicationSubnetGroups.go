// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeReplicationSubnetGroupsInput struct {
	_ struct{} `type:"structure"`

	// Filters applied to the describe action.
	Filters []Filter `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeReplicationSubnetGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeReplicationSubnetGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeReplicationSubnetGroupsInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeReplicationSubnetGroupsOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// A description of the replication subnet groups.
	ReplicationSubnetGroups []ReplicationSubnetGroup `type:"list"`
}

// String returns the string representation
func (s DescribeReplicationSubnetGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeReplicationSubnetGroups = "DescribeReplicationSubnetGroups"

// DescribeReplicationSubnetGroupsRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Returns information about the replication subnet groups.
//
//    // Example sending a request using DescribeReplicationSubnetGroupsRequest.
//    req := client.DescribeReplicationSubnetGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeReplicationSubnetGroups
func (c *Client) DescribeReplicationSubnetGroupsRequest(input *DescribeReplicationSubnetGroupsInput) DescribeReplicationSubnetGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeReplicationSubnetGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeReplicationSubnetGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeReplicationSubnetGroupsOutput{})
	return DescribeReplicationSubnetGroupsRequest{Request: req, Input: input, Copy: c.DescribeReplicationSubnetGroupsRequest}
}

// DescribeReplicationSubnetGroupsRequest is the request type for the
// DescribeReplicationSubnetGroups API operation.
type DescribeReplicationSubnetGroupsRequest struct {
	*aws.Request
	Input *DescribeReplicationSubnetGroupsInput
	Copy  func(*DescribeReplicationSubnetGroupsInput) DescribeReplicationSubnetGroupsRequest
}

// Send marshals and sends the DescribeReplicationSubnetGroups API request.
func (r DescribeReplicationSubnetGroupsRequest) Send(ctx context.Context) (*DescribeReplicationSubnetGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeReplicationSubnetGroupsResponse{
		DescribeReplicationSubnetGroupsOutput: r.Request.Data.(*DescribeReplicationSubnetGroupsOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeReplicationSubnetGroupsRequestPaginator returns a paginator for DescribeReplicationSubnetGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeReplicationSubnetGroupsRequest(input)
//   p := databasemigrationservice.NewDescribeReplicationSubnetGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeReplicationSubnetGroupsPaginator(req DescribeReplicationSubnetGroupsRequest) DescribeReplicationSubnetGroupsPaginator {
	return DescribeReplicationSubnetGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeReplicationSubnetGroupsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeReplicationSubnetGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeReplicationSubnetGroupsPaginator struct {
	aws.Pager
}

func (p *DescribeReplicationSubnetGroupsPaginator) CurrentPage() *DescribeReplicationSubnetGroupsOutput {
	return p.Pager.CurrentPage().(*DescribeReplicationSubnetGroupsOutput)
}

// DescribeReplicationSubnetGroupsResponse is the response type for the
// DescribeReplicationSubnetGroups API operation.
type DescribeReplicationSubnetGroupsResponse struct {
	*DescribeReplicationSubnetGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeReplicationSubnetGroups request.
func (r *DescribeReplicationSubnetGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}