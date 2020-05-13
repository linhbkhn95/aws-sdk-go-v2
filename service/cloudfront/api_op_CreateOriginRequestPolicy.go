// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateOriginRequestPolicyInput struct {
	_ struct{} `type:"structure" payload:"OriginRequestPolicyConfig"`

	// An origin request policy configuration.
	//
	// OriginRequestPolicyConfig is a required field
	OriginRequestPolicyConfig *OriginRequestPolicyConfig `locationName:"OriginRequestPolicyConfig" type:"structure" required:"true" xmlURI:"http://cloudfront.amazonaws.com/doc/2020-05-31/"`
}

// String returns the string representation
func (s CreateOriginRequestPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateOriginRequestPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateOriginRequestPolicyInput"}

	if s.OriginRequestPolicyConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("OriginRequestPolicyConfig"))
	}
	if s.OriginRequestPolicyConfig != nil {
		if err := s.OriginRequestPolicyConfig.Validate(); err != nil {
			invalidParams.AddNested("OriginRequestPolicyConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateOriginRequestPolicyInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.OriginRequestPolicyConfig != nil {
		v := s.OriginRequestPolicyConfig

		metadata := protocol.Metadata{XMLNamespaceURI: "http://cloudfront.amazonaws.com/doc/2020-05-31/"}
		e.SetFields(protocol.PayloadTarget, "OriginRequestPolicyConfig", v, metadata)
	}
	return nil
}

type CreateOriginRequestPolicyOutput struct {
	_ struct{} `type:"structure" payload:"OriginRequestPolicy"`

	// The current version of the origin request policy.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// The fully qualified URI of the origin request policy just created.
	Location *string `location:"header" locationName:"Location" type:"string"`

	// An origin request policy.
	OriginRequestPolicy *OriginRequestPolicy `type:"structure"`
}

// String returns the string representation
func (s CreateOriginRequestPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateOriginRequestPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Location", protocol.StringValue(v), metadata)
	}
	if s.OriginRequestPolicy != nil {
		v := s.OriginRequestPolicy

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "OriginRequestPolicy", v, metadata)
	}
	return nil
}

const opCreateOriginRequestPolicy = "CreateOriginRequestPolicy2020_05_31"

// CreateOriginRequestPolicyRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Creates an origin request policy.
//
// After you create an origin request policy, you can attach it to one or more
// cache behaviors. When it’s attached to a cache behavior, the origin request
// policy determines the values that CloudFront includes in requests that it
// sends to the origin. Each request that CloudFront sends to the origin includes
// the following:
//
//    * The request body and the URL path (without the domain name) from the
//    viewer request.
//
//    * The headers that CloudFront automatically includes in every origin request,
//    including Host, User-Agent, and X-Amz-Cf-Id.
//
//    * All HTTP headers, cookies, and URL query strings that are specified
//    in the cache policy or the origin request policy. These can include items
//    from the viewer request and, in the case of headers, additional ones that
//    are added by CloudFront.
//
// CloudFront sends a request when it can’t find a valid object in its cache
// that matches the request. If you want to send values to the origin and also
// include them in the cache key, use CreateCachePolicy.
//
// For more information about origin request policies, see Controlling origin
// requests (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html)
// in the Amazon CloudFront Developer Guide.
//
//    // Example sending a request using CreateOriginRequestPolicyRequest.
//    req := client.CreateOriginRequestPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2020-05-31/CreateOriginRequestPolicy
func (c *Client) CreateOriginRequestPolicyRequest(input *CreateOriginRequestPolicyInput) CreateOriginRequestPolicyRequest {
	op := &aws.Operation{
		Name:       opCreateOriginRequestPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/2020-05-31/origin-request-policy",
	}

	if input == nil {
		input = &CreateOriginRequestPolicyInput{}
	}

	req := c.newRequest(op, input, &CreateOriginRequestPolicyOutput{})

	return CreateOriginRequestPolicyRequest{Request: req, Input: input, Copy: c.CreateOriginRequestPolicyRequest}
}

// CreateOriginRequestPolicyRequest is the request type for the
// CreateOriginRequestPolicy API operation.
type CreateOriginRequestPolicyRequest struct {
	*aws.Request
	Input *CreateOriginRequestPolicyInput
	Copy  func(*CreateOriginRequestPolicyInput) CreateOriginRequestPolicyRequest
}

// Send marshals and sends the CreateOriginRequestPolicy API request.
func (r CreateOriginRequestPolicyRequest) Send(ctx context.Context) (*CreateOriginRequestPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateOriginRequestPolicyResponse{
		CreateOriginRequestPolicyOutput: r.Request.Data.(*CreateOriginRequestPolicyOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateOriginRequestPolicyResponse is the response type for the
// CreateOriginRequestPolicy API operation.
type CreateOriginRequestPolicyResponse struct {
	*CreateOriginRequestPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateOriginRequestPolicy request.
func (r *CreateOriginRequestPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}