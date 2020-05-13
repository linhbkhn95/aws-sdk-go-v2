// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateConstraintInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The description of the constraint.
	Description *string `type:"string"`

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	//
	// IdempotencyToken is a required field
	IdempotencyToken *string `min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The constraint parameters, in JSON format. The syntax depends on the constraint
	// type as follows:
	//
	// LAUNCH
	//
	// You are required to specify either the RoleArn or the LocalRoleName but can't
	// use both.
	//
	// Specify the RoleArn property as follows:
	//
	// {"RoleArn" : "arn:aws:iam::123456789012:role/LaunchRole"}
	//
	// Specify the LocalRoleName property as follows:
	//
	// {"LocalRoleName": "SCBasicLaunchRole"}
	//
	// If you specify the LocalRoleName property, when an account uses the launch
	// constraint, the IAM role with that name in the account will be used. This
	// allows launch-role constraints to be account-agnostic so the administrator
	// can create fewer resources per shared account.
	//
	// The given role name must exist in the account used to create the launch constraint
	// and the account of the user who launches a product with this launch constraint.
	//
	// You cannot have both a LAUNCH and a STACKSET constraint.
	//
	// You also cannot have more than one LAUNCH constraint on a product and portfolio.
	//
	// NOTIFICATION
	//
	// Specify the NotificationArns property as follows:
	//
	// {"NotificationArns" : ["arn:aws:sns:us-east-1:123456789012:Topic"]}
	//
	// RESOURCE_UPDATE
	//
	// Specify the TagUpdatesOnProvisionedProduct property as follows:
	//
	// {"Version":"2.0","Properties":{"TagUpdateOnProvisionedProduct":"String"}}
	//
	// The TagUpdatesOnProvisionedProduct property accepts a string value of ALLOWED
	// or NOT_ALLOWED.
	//
	// STACKSET
	//
	// Specify the Parameters property as follows:
	//
	// {"Version": "String", "Properties": {"AccountList": [ "String" ], "RegionList":
	// [ "String" ], "AdminRole": "String", "ExecutionRole": "String"}}
	//
	// You cannot have both a LAUNCH and a STACKSET constraint.
	//
	// You also cannot have more than one STACKSET constraint on a product and portfolio.
	//
	// Products with a STACKSET constraint will launch an AWS CloudFormation stack
	// set.
	//
	// TEMPLATE
	//
	// Specify the Rules property. For more information, see Template Constraint
	// Rules (http://docs.aws.amazon.com/servicecatalog/latest/adminguide/reference-template_constraint_rules.html).
	//
	// Parameters is a required field
	Parameters *string `type:"string" required:"true"`

	// The portfolio identifier.
	//
	// PortfolioId is a required field
	PortfolioId *string `min:"1" type:"string" required:"true"`

	// The product identifier.
	//
	// ProductId is a required field
	ProductId *string `min:"1" type:"string" required:"true"`

	// The type of constraint.
	//
	//    * LAUNCH
	//
	//    * NOTIFICATION
	//
	//    * RESOURCE_UPDATE
	//
	//    * STACKSET
	//
	//    * TEMPLATE
	//
	// Type is a required field
	Type *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateConstraintInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateConstraintInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateConstraintInput"}

	if s.IdempotencyToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdempotencyToken"))
	}
	if s.IdempotencyToken != nil && len(*s.IdempotencyToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdempotencyToken", 1))
	}

	if s.Parameters == nil {
		invalidParams.Add(aws.NewErrParamRequired("Parameters"))
	}

	if s.PortfolioId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortfolioId"))
	}
	if s.PortfolioId != nil && len(*s.PortfolioId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortfolioId", 1))
	}

	if s.ProductId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductId"))
	}
	if s.ProductId != nil && len(*s.ProductId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductId", 1))
	}

	if s.Type == nil {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}
	if s.Type != nil && len(*s.Type) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Type", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateConstraintOutput struct {
	_ struct{} `type:"structure"`

	// Information about the constraint.
	ConstraintDetail *ConstraintDetail `type:"structure"`

	// The constraint parameters.
	ConstraintParameters *string `type:"string"`

	// The status of the current request.
	Status Status `type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateConstraintOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateConstraint = "CreateConstraint"

// CreateConstraintRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Creates a constraint.
//
// A delegated admin is authorized to invoke this command.
//
//    // Example sending a request using CreateConstraintRequest.
//    req := client.CreateConstraintRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreateConstraint
func (c *Client) CreateConstraintRequest(input *CreateConstraintInput) CreateConstraintRequest {
	op := &aws.Operation{
		Name:       opCreateConstraint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateConstraintInput{}
	}

	req := c.newRequest(op, input, &CreateConstraintOutput{})

	return CreateConstraintRequest{Request: req, Input: input, Copy: c.CreateConstraintRequest}
}

// CreateConstraintRequest is the request type for the
// CreateConstraint API operation.
type CreateConstraintRequest struct {
	*aws.Request
	Input *CreateConstraintInput
	Copy  func(*CreateConstraintInput) CreateConstraintRequest
}

// Send marshals and sends the CreateConstraint API request.
func (r CreateConstraintRequest) Send(ctx context.Context) (*CreateConstraintResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateConstraintResponse{
		CreateConstraintOutput: r.Request.Data.(*CreateConstraintOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateConstraintResponse is the response type for the
// CreateConstraint API operation.
type CreateConstraintResponse struct {
	*CreateConstraintOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateConstraint request.
func (r *CreateConstraintResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}