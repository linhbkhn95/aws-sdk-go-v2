// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListBackupSelectionsInput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies a backup plan.
	//
	// BackupPlanId is a required field
	BackupPlanId *string `location:"uri" locationName:"backupPlanId" type:"string" required:"true"`

	// The maximum number of items to be returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListBackupSelectionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBackupSelectionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBackupSelectionsInput"}

	if s.BackupPlanId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupPlanId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBackupSelectionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BackupPlanId != nil {
		v := *s.BackupPlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "backupPlanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListBackupSelectionsOutput struct {
	_ struct{} `type:"structure"`

	// An array of backup selection list items containing metadata about each resource
	// in the list.
	BackupSelectionsList []BackupSelectionsListMember `type:"list"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListBackupSelectionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBackupSelectionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackupSelectionsList != nil {
		v := s.BackupSelectionsList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "BackupSelectionsList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListBackupSelections = "ListBackupSelections"

// ListBackupSelectionsRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns an array containing metadata of the resources associated with the
// target backup plan.
//
//    // Example sending a request using ListBackupSelectionsRequest.
//    req := client.ListBackupSelectionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListBackupSelections
func (c *Client) ListBackupSelectionsRequest(input *ListBackupSelectionsInput) ListBackupSelectionsRequest {
	op := &aws.Operation{
		Name:       opListBackupSelections,
		HTTPMethod: "GET",
		HTTPPath:   "/backup/plans/{backupPlanId}/selections/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListBackupSelectionsInput{}
	}

	req := c.newRequest(op, input, &ListBackupSelectionsOutput{})
	return ListBackupSelectionsRequest{Request: req, Input: input, Copy: c.ListBackupSelectionsRequest}
}

// ListBackupSelectionsRequest is the request type for the
// ListBackupSelections API operation.
type ListBackupSelectionsRequest struct {
	*aws.Request
	Input *ListBackupSelectionsInput
	Copy  func(*ListBackupSelectionsInput) ListBackupSelectionsRequest
}

// Send marshals and sends the ListBackupSelections API request.
func (r ListBackupSelectionsRequest) Send(ctx context.Context) (*ListBackupSelectionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBackupSelectionsResponse{
		ListBackupSelectionsOutput: r.Request.Data.(*ListBackupSelectionsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListBackupSelectionsRequestPaginator returns a paginator for ListBackupSelections.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListBackupSelectionsRequest(input)
//   p := backup.NewListBackupSelectionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListBackupSelectionsPaginator(req ListBackupSelectionsRequest) ListBackupSelectionsPaginator {
	return ListBackupSelectionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListBackupSelectionsInput
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

// ListBackupSelectionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListBackupSelectionsPaginator struct {
	aws.Pager
}

func (p *ListBackupSelectionsPaginator) CurrentPage() *ListBackupSelectionsOutput {
	return p.Pager.CurrentPage().(*ListBackupSelectionsOutput)
}

// ListBackupSelectionsResponse is the response type for the
// ListBackupSelections API operation.
type ListBackupSelectionsResponse struct {
	*ListBackupSelectionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBackupSelections request.
func (r *ListBackupSelectionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}