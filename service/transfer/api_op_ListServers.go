// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transfer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListServersInput struct {
	_ struct{} `type:"structure"`

	// Specifies the number of servers to return as a response to the ListServers
	// query.
	MaxResults *int64 `min:"1" type:"integer"`

	// When additional results are obtained from the ListServers command, a NextToken
	// parameter is returned in the output. You can then pass the NextToken parameter
	// in a subsequent command to continue listing additional servers.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListServersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListServersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListServersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListServersOutput struct {
	_ struct{} `type:"structure"`

	// When you can get additional results from the ListServers operation, a NextToken
	// parameter is returned in the output. In a following command, you can pass
	// in the NextToken parameter to continue listing additional servers.
	NextToken *string `min:"1" type:"string"`

	// An array of servers that were listed.
	//
	// Servers is a required field
	Servers []ListedServer `type:"list" required:"true"`
}

// String returns the string representation
func (s ListServersOutput) String() string {
	return awsutil.Prettify(s)
}

const opListServers = "ListServers"

// ListServersRequest returns a request value for making API operation for
// AWS Transfer for SFTP.
//
// Lists the Secure File Transfer Protocol (SFTP) servers that are associated
// with your AWS account.
//
//    // Example sending a request using ListServersRequest.
//    req := client.ListServersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/ListServers
func (c *Client) ListServersRequest(input *ListServersInput) ListServersRequest {
	op := &aws.Operation{
		Name:       opListServers,
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
		input = &ListServersInput{}
	}

	req := c.newRequest(op, input, &ListServersOutput{})
	return ListServersRequest{Request: req, Input: input, Copy: c.ListServersRequest}
}

// ListServersRequest is the request type for the
// ListServers API operation.
type ListServersRequest struct {
	*aws.Request
	Input *ListServersInput
	Copy  func(*ListServersInput) ListServersRequest
}

// Send marshals and sends the ListServers API request.
func (r ListServersRequest) Send(ctx context.Context) (*ListServersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListServersResponse{
		ListServersOutput: r.Request.Data.(*ListServersOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListServersRequestPaginator returns a paginator for ListServers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListServersRequest(input)
//   p := transfer.NewListServersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListServersPaginator(req ListServersRequest) ListServersPaginator {
	return ListServersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListServersInput
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

// ListServersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListServersPaginator struct {
	aws.Pager
}

func (p *ListServersPaginator) CurrentPage() *ListServersOutput {
	return p.Pager.CurrentPage().(*ListServersOutput)
}

// ListServersResponse is the response type for the
// ListServers API operation.
type ListServersResponse struct {
	*ListServersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListServers request.
func (r *ListServersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}