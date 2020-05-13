// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The structure representing the getProfileRequest.
type GetProfileInput struct {
	_ struct{} `type:"structure"`

	// The format of the returned profiling data. The format maps to the Accept
	// and Content-Type headers of the HTTP request. You can specify one of the
	// following: or the default .
	//
	//    <ul> <li> <p> <code>application/json</code> — standard JSON format </p>
	//    </li> <li> <p> <code>application/x-amzn-ion</code> — the Amazon Ion
	//    data format. For more information, see <a href="http://amzn.github.io/ion-docs/">Amazon
	//    Ion</a>. </p> </li> </ul>
	Accept *string `location:"header" locationName:"Accept" type:"string"`

	// The end time of the requested profile. Specify using the ISO 8601 format.
	// For example, 2020-06-01T13:15:02.001Z represents 1 millisecond past June
	// 1, 2020 1:15:02 PM UTC.
	//
	// If you specify endTime, then you must also specify period or startTime, but
	// not both.
	EndTime *time.Time `location:"querystring" locationName:"endTime" type:"timestamp" timestampFormat:"iso8601"`

	// The maximum depth of the stacks in the code that is represented in the aggregated
	// profile. For example, if CodeGuru Profiler finds a method A, which calls
	// method B, which calls method C, which calls method D, then the depth is 4.
	// If the maxDepth is set to 2, then the aggregated profile contains representations
	// of methods A and B.
	MaxDepth *int64 `location:"querystring" locationName:"maxDepth" min:"1" type:"integer"`

	// Used with startTime or endTime to specify the time range for the returned
	// aggregated profile. Specify using the ISO 8601 format. For example, P1DT1H1M1S.
	//
	//    <p> To get the latest aggregated profile, specify only <code>period</code>.
	//    </p>
	Period *string `location:"querystring" locationName:"period" min:"1" type:"string"`

	// The name of the profiling group to get.
	//
	// ProfilingGroupName is a required field
	ProfilingGroupName *string `location:"uri" locationName:"profilingGroupName" min:"1" type:"string" required:"true"`

	// The start time of the profile to get. Specify using the ISO 8601 format.
	// For example, 2020-06-01T13:15:02.001Z represents 1 millisecond past June
	// 1, 2020 1:15:02 PM UTC.
	//
	//    <p> If you specify <code>startTime</code>, then you must also specify
	//    <code>period</code> or <code>endTime</code>, but not both. </p>
	StartTime *time.Time `location:"querystring" locationName:"startTime" type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s GetProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetProfileInput"}
	if s.MaxDepth != nil && *s.MaxDepth < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxDepth", 1))
	}
	if s.Period != nil && len(*s.Period) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Period", 1))
	}

	if s.ProfilingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProfilingGroupName"))
	}
	if s.ProfilingGroupName != nil && len(*s.ProfilingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProfilingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetProfileInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Accept != nil {
		v := *s.Accept

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Accept", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProfilingGroupName != nil {
		v := *s.ProfilingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "profilingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndTime != nil {
		v := *s.EndTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "endTime",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: false}, metadata)
	}
	if s.MaxDepth != nil {
		v := *s.MaxDepth

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxDepth", protocol.Int64Value(v), metadata)
	}
	if s.Period != nil {
		v := *s.Period

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "period", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StartTime != nil {
		v := *s.StartTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "startTime",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: false}, metadata)
	}
	return nil
}

// The structure representing the getProfileResponse.
type GetProfileOutput struct {
	_ struct{} `type:"structure" payload:"Profile"`

	// The content encoding of the profile.
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`

	// The content type of the profile in the payload. It is either application/json
	// or the default application/x-amzn-ion.
	//
	// ContentType is a required field
	ContentType *string `location:"header" locationName:"Content-Type" type:"string" required:"true"`

	// Information about the profile.
	//
	// Profile is a required field
	Profile []byte `locationName:"profile" type:"blob" required:"true"`
}

// String returns the string representation
func (s GetProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContentEncoding != nil {
		v := *s.ContentEncoding

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Encoding", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Profile != nil {
		v := s.Profile

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "profile", protocol.BytesStream(v), metadata)
	}
	return nil
}

const opGetProfile = "GetProfile"

// GetProfileRequest returns a request value for making API operation for
// Amazon CodeGuru Profiler.
//
// Gets the aggregated profile of a profiling group for a specified time range.
// Amazon CodeGuru Profiler collects posted agent profiles for a profiling group
// into aggregated profiles.
//
//    <note> <p> Because aggregated profiles expire over time <code>GetProfile</code>
//    is not idempotent. </p> </note> <p> Specify the time range for the requested
//    aggregated profile using 1 or 2 of the following parameters: <code>startTime</code>,
//    <code>endTime</code>, <code>period</code>. The maximum time range allowed
//    is 7 days. If you specify all 3 parameters, an exception is thrown. If
//    you specify only <code>period</code>, the latest aggregated profile is
//    returned. </p> <p> Aggregated profiles are available with aggregation
//    periods of 5 minutes, 1 hour, and 1 day, aligned to UTC. The aggregation
//    period of an aggregated profile determines how long it is retained. For
//    more information, see <a href="https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_AggregatedProfileTime.html">
//    <code>AggregatedProfileTime</code> </a>. The aggregated profile's aggregation
//    period determines how long it is retained by CodeGuru Profiler. </p> <ul>
//    <li> <p> If the aggregation period is 5 minutes, the aggregated profile
//    is retained for 15 days. </p> </li> <li> <p> If the aggregation period
//    is 1 hour, the aggregated profile is retained for 60 days. </p> </li>
//    <li> <p> If the aggregation period is 1 day, the aggregated profile is
//    retained for 3 years. </p> </li> </ul> <p>There are two use cases for
//    calling <code>GetProfile</code>.</p> <ol> <li> <p> If you want to return
//    an aggregated profile that already exists, use <a href="https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ListProfileTimes.html">
//    <code>ListProfileTimes</code> </a> to view the time ranges of existing
//    aggregated profiles. Use them in a <code>GetProfile</code> request to
//    return a specific, existing aggregated profile. </p> </li> <li> <p> If
//    you want to return an aggregated profile for a time range that doesn't
//    align with an existing aggregated profile, then CodeGuru Profiler makes
//    a best effort to combine existing aggregated profiles from the requested
//    time range and return them as one aggregated profile. </p> <p> If aggregated
//    profiles do not exist for the full time range requested, then aggregated
//    profiles for a smaller time range are returned. For example, if the requested
//    time range is from 00:00 to 00:20, and the existing aggregated profiles
//    are from 00:15 and 00:25, then the aggregated profiles from 00:15 to 00:20
//    are returned. </p> </li> </ol>
//
//    // Example sending a request using GetProfileRequest.
//    req := client.GetProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeguruprofiler-2019-07-18/GetProfile
func (c *Client) GetProfileRequest(input *GetProfileInput) GetProfileRequest {
	op := &aws.Operation{
		Name:       opGetProfile,
		HTTPMethod: "GET",
		HTTPPath:   "/profilingGroups/{profilingGroupName}/profile",
	}

	if input == nil {
		input = &GetProfileInput{}
	}

	req := c.newRequest(op, input, &GetProfileOutput{})

	return GetProfileRequest{Request: req, Input: input, Copy: c.GetProfileRequest}
}

// GetProfileRequest is the request type for the
// GetProfile API operation.
type GetProfileRequest struct {
	*aws.Request
	Input *GetProfileInput
	Copy  func(*GetProfileInput) GetProfileRequest
}

// Send marshals and sends the GetProfile API request.
func (r GetProfileRequest) Send(ctx context.Context) (*GetProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetProfileResponse{
		GetProfileOutput: r.Request.Data.(*GetProfileOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetProfileResponse is the response type for the
// GetProfile API operation.
type GetProfileResponse struct {
	*GetProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetProfile request.
func (r *GetProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}