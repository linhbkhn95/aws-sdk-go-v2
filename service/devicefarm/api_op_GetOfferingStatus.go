// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to retrieve the offering status for the specified
// customer or account.
type GetOfferingStatusInput struct {
	_ struct{} `type:"structure"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s GetOfferingStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetOfferingStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetOfferingStatusInput"}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Returns the status result for a device offering.
type GetOfferingStatusOutput struct {
	_ struct{} `type:"structure"`

	// When specified, gets the offering status for the current period.
	Current map[string]OfferingStatus `locationName:"current" type:"map"`

	// When specified, gets the offering status for the next period.
	NextPeriod map[string]OfferingStatus `locationName:"nextPeriod" type:"map"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s GetOfferingStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetOfferingStatus = "GetOfferingStatus"

// GetOfferingStatusRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets the current status and future status of all offerings purchased by an
// AWS account. The response indicates how many offerings are currently available
// and the offerings that will be available in the next period. The API returns
// a NotEligible error if the user is not permitted to invoke the operation.
// Please contact aws-devicefarm-support@amazon.com (mailto:aws-devicefarm-support@amazon.com)
// if you believe that you should be able to invoke this operation.
//
//    // Example sending a request using GetOfferingStatusRequest.
//    req := client.GetOfferingStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetOfferingStatus
func (c *Client) GetOfferingStatusRequest(input *GetOfferingStatusInput) GetOfferingStatusRequest {
	op := &aws.Operation{
		Name:       opGetOfferingStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetOfferingStatusInput{}
	}

	req := c.newRequest(op, input, &GetOfferingStatusOutput{})
	return GetOfferingStatusRequest{Request: req, Input: input, Copy: c.GetOfferingStatusRequest}
}

// GetOfferingStatusRequest is the request type for the
// GetOfferingStatus API operation.
type GetOfferingStatusRequest struct {
	*aws.Request
	Input *GetOfferingStatusInput
	Copy  func(*GetOfferingStatusInput) GetOfferingStatusRequest
}

// Send marshals and sends the GetOfferingStatus API request.
func (r GetOfferingStatusRequest) Send(ctx context.Context) (*GetOfferingStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetOfferingStatusResponse{
		GetOfferingStatusOutput: r.Request.Data.(*GetOfferingStatusOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetOfferingStatusRequestPaginator returns a paginator for GetOfferingStatus.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetOfferingStatusRequest(input)
//   p := devicefarm.NewGetOfferingStatusRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetOfferingStatusPaginator(req GetOfferingStatusRequest) GetOfferingStatusPaginator {
	return GetOfferingStatusPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetOfferingStatusInput
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

// GetOfferingStatusPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetOfferingStatusPaginator struct {
	aws.Pager
}

func (p *GetOfferingStatusPaginator) CurrentPage() *GetOfferingStatusOutput {
	return p.Pager.CurrentPage().(*GetOfferingStatusOutput)
}

// GetOfferingStatusResponse is the response type for the
// GetOfferingStatus API operation.
type GetOfferingStatusResponse struct {
	*GetOfferingStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetOfferingStatus request.
func (r *GetOfferingStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}