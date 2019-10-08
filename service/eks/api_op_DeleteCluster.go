// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteClusterInput struct {
	_ struct{} `type:"structure"`

	// The name of the cluster to delete.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteClusterInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteClusterInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteClusterOutput struct {
	_ struct{} `type:"structure"`

	// The full description of the cluster to delete.
	Cluster *Cluster `locationName:"cluster" type:"structure"`
}

// String returns the string representation
func (s DeleteClusterOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteClusterOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Cluster != nil {
		v := s.Cluster

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "cluster", v, metadata)
	}
	return nil
}

const opDeleteCluster = "DeleteCluster"

// DeleteClusterRequest returns a request value for making API operation for
// Amazon Elastic Kubernetes Service.
//
// Deletes the Amazon EKS cluster control plane.
//
// If you have active services in your cluster that are associated with a load
// balancer, you must delete those services before deleting the cluster so that
// the load balancers are deleted properly. Otherwise, you can have orphaned
// resources in your VPC that prevent you from being able to delete the VPC.
// For more information, see Deleting a Cluster (https://docs.aws.amazon.com/eks/latest/userguide/delete-cluster.html)
// in the Amazon EKS User Guide.
//
//    // Example sending a request using DeleteClusterRequest.
//    req := client.DeleteClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/DeleteCluster
func (c *Client) DeleteClusterRequest(input *DeleteClusterInput) DeleteClusterRequest {
	op := &aws.Operation{
		Name:       opDeleteCluster,
		HTTPMethod: "DELETE",
		HTTPPath:   "/clusters/{name}",
	}

	if input == nil {
		input = &DeleteClusterInput{}
	}

	req := c.newRequest(op, input, &DeleteClusterOutput{})
	return DeleteClusterRequest{Request: req, Input: input, Copy: c.DeleteClusterRequest}
}

// DeleteClusterRequest is the request type for the
// DeleteCluster API operation.
type DeleteClusterRequest struct {
	*aws.Request
	Input *DeleteClusterInput
	Copy  func(*DeleteClusterInput) DeleteClusterRequest
}

// Send marshals and sends the DeleteCluster API request.
func (r DeleteClusterRequest) Send(ctx context.Context) (*DeleteClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteClusterResponse{
		DeleteClusterOutput: r.Request.Data.(*DeleteClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteClusterResponse is the response type for the
// DeleteCluster API operation.
type DeleteClusterResponse struct {
	*DeleteClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCluster request.
func (r *DeleteClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}