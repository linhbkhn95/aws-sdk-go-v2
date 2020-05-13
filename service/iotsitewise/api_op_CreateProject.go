// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateProjectInput struct {
	_ struct{} `type:"structure"`

	// A unique case-sensitive identifier that you can provide to ensure the idempotency
	// of the request. Don't reuse this client token if a new idempotent request
	// is required.
	ClientToken *string `locationName:"clientToken" min:"36" type:"string" idempotencyToken:"true"`

	// The ID of the portal in which to create the project.
	//
	// PortalId is a required field
	PortalId *string `locationName:"portalId" min:"36" type:"string" required:"true"`

	// A description for the project.
	ProjectDescription *string `locationName:"projectDescription" min:"1" type:"string"`

	// A friendly name for the project.
	//
	// ProjectName is a required field
	ProjectName *string `locationName:"projectName" min:"1" type:"string" required:"true"`

	// A list of key-value pairs that contain metadata for the project. For more
	// information, see Tagging your AWS IoT SiteWise resources (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html)
	// in the AWS IoT SiteWise User Guide.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s CreateProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateProjectInput"}
	if s.ClientToken != nil && len(*s.ClientToken) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 36))
	}

	if s.PortalId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortalId"))
	}
	if s.PortalId != nil && len(*s.PortalId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("PortalId", 36))
	}
	if s.ProjectDescription != nil && len(*s.ProjectDescription) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectDescription", 1))
	}

	if s.ProjectName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectName"))
	}
	if s.ProjectName != nil && len(*s.ProjectName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectName", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateProjectInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PortalId != nil {
		v := *s.PortalId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "portalId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProjectDescription != nil {
		v := *s.ProjectDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "projectDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProjectName != nil {
		v := *s.ProjectName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "projectName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

type CreateProjectOutput struct {
	_ struct{} `type:"structure"`

	// The ARN (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the project, which has the following format.
	//
	// arn:${Partition}:iotsitewise:${Region}:${Account}:project/${ProjectId}
	//
	// ProjectArn is a required field
	ProjectArn *string `locationName:"projectArn" min:"1" type:"string" required:"true"`

	// The ID of the project.
	//
	// ProjectId is a required field
	ProjectId *string `locationName:"projectId" min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateProjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateProjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ProjectArn != nil {
		v := *s.ProjectArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "projectArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProjectId != nil {
		v := *s.ProjectId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "projectId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateProject = "CreateProject"

// CreateProjectRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Creates a project in the specified portal.
//
//    // Example sending a request using CreateProjectRequest.
//    req := client.CreateProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/CreateProject
func (c *Client) CreateProjectRequest(input *CreateProjectInput) CreateProjectRequest {
	op := &aws.Operation{
		Name:       opCreateProject,
		HTTPMethod: "POST",
		HTTPPath:   "/projects",
	}

	if input == nil {
		input = &CreateProjectInput{}
	}

	req := c.newRequest(op, input, &CreateProjectOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("monitor.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return CreateProjectRequest{Request: req, Input: input, Copy: c.CreateProjectRequest}
}

// CreateProjectRequest is the request type for the
// CreateProject API operation.
type CreateProjectRequest struct {
	*aws.Request
	Input *CreateProjectInput
	Copy  func(*CreateProjectInput) CreateProjectRequest
}

// Send marshals and sends the CreateProject API request.
func (r CreateProjectRequest) Send(ctx context.Context) (*CreateProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProjectResponse{
		CreateProjectOutput: r.Request.Data.(*CreateProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProjectResponse is the response type for the
// CreateProject API operation.
type CreateProjectResponse struct {
	*CreateProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProject request.
func (r *CreateProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}