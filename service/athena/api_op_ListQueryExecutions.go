// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListQueryExecutionsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of query executions to return in this request.
	MaxResults *int64 `type:"integer"`

	// The token that specifies where to start pagination if a previous request
	// was truncated.
	NextToken *string `min:"1" type:"string"`

	// The name of the workgroup from which queries are being returned.
	WorkGroup *string `type:"string"`
}

// String returns the string representation
func (s ListQueryExecutionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListQueryExecutionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListQueryExecutionsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListQueryExecutionsOutput struct {
	_ struct{} `type:"structure"`

	// A token to be used by the next request if this request is truncated.
	NextToken *string `min:"1" type:"string"`

	// The unique IDs of each query execution as an array of strings.
	QueryExecutionIds []string `min:"1" type:"list"`
}

// String returns the string representation
func (s ListQueryExecutionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListQueryExecutions = "ListQueryExecutions"

// ListQueryExecutionsRequest returns a request value for making API operation for
// Amazon Athena.
//
// Provides a list of available query execution IDs for the queries in the specified
// workgroup. Requires you to have access to the workgroup in which the queries
// ran.
//
// For code samples using the AWS SDK for Java, see Examples and Code Samples
// (http://docs.aws.amazon.com/athena/latest/ug/code-samples.html) in the Amazon
// Athena User Guide.
//
//    // Example sending a request using ListQueryExecutionsRequest.
//    req := client.ListQueryExecutionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/ListQueryExecutions
func (c *Client) ListQueryExecutionsRequest(input *ListQueryExecutionsInput) ListQueryExecutionsRequest {
	op := &aws.Operation{
		Name:       opListQueryExecutions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListQueryExecutionsInput{}
	}

	req := c.newRequest(op, input, &ListQueryExecutionsOutput{})
	return ListQueryExecutionsRequest{Request: req, Input: input, Copy: c.ListQueryExecutionsRequest}
}

// ListQueryExecutionsRequest is the request type for the
// ListQueryExecutions API operation.
type ListQueryExecutionsRequest struct {
	*aws.Request
	Input *ListQueryExecutionsInput
	Copy  func(*ListQueryExecutionsInput) ListQueryExecutionsRequest
}

// Send marshals and sends the ListQueryExecutions API request.
func (r ListQueryExecutionsRequest) Send(ctx context.Context) (*ListQueryExecutionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListQueryExecutionsResponse{
		ListQueryExecutionsOutput: r.Request.Data.(*ListQueryExecutionsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListQueryExecutionsRequestPaginator returns a paginator for ListQueryExecutions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListQueryExecutionsRequest(input)
//   p := athena.NewListQueryExecutionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListQueryExecutionsPaginator(req ListQueryExecutionsRequest) ListQueryExecutionsPaginator {
	return ListQueryExecutionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListQueryExecutionsInput
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

// ListQueryExecutionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListQueryExecutionsPaginator struct {
	aws.Pager
}

func (p *ListQueryExecutionsPaginator) CurrentPage() *ListQueryExecutionsOutput {
	return p.Pager.CurrentPage().(*ListQueryExecutionsOutput)
}

// ListQueryExecutionsResponse is the response type for the
// ListQueryExecutions API operation.
type ListQueryExecutionsResponse struct {
	*ListQueryExecutionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListQueryExecutions request.
func (r *ListQueryExecutionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}