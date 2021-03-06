// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery_test

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
)

var _ aws.Config

// CreateHttpNamespace example
//
// This example creates an HTTP namespace.
func ExampleClient_CreateHttpNamespaceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.CreateHttpNamespaceInput{
		CreatorRequestId: aws.String("example-creator-request-id-0001"),
		Description:      aws.String("Example.com AWS Cloud Map HTTP Namespace"),
		Name:             aws.String("example-http.com"),
	}

	req := svc.CreateHttpNamespaceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeNamespaceAlreadyExists:
				fmt.Println(servicediscovery.ErrCodeNamespaceAlreadyExists, aerr.Error())
			case servicediscovery.ErrCodeResourceLimitExceeded:
				fmt.Println(servicediscovery.ErrCodeResourceLimitExceeded, aerr.Error())
			case servicediscovery.ErrCodeDuplicateRequest:
				fmt.Println(servicediscovery.ErrCodeDuplicateRequest, aerr.Error())
			case servicediscovery.ErrCodeTooManyTagsException:
				fmt.Println(servicediscovery.ErrCodeTooManyTagsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Example: Create private DNS namespace
//
// Example: Create private DNS namespace
func ExampleClient_CreatePrivateDnsNamespaceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.CreatePrivateDnsNamespaceInput{
		CreatorRequestId: aws.String("eedd6892-50f3-41b2-8af9-611d6e1d1a8c"),
		Name:             aws.String("example.com"),
		Vpc:              aws.String("vpc-1c56417b"),
	}

	req := svc.CreatePrivateDnsNamespaceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeNamespaceAlreadyExists:
				fmt.Println(servicediscovery.ErrCodeNamespaceAlreadyExists, aerr.Error())
			case servicediscovery.ErrCodeResourceLimitExceeded:
				fmt.Println(servicediscovery.ErrCodeResourceLimitExceeded, aerr.Error())
			case servicediscovery.ErrCodeDuplicateRequest:
				fmt.Println(servicediscovery.ErrCodeDuplicateRequest, aerr.Error())
			case servicediscovery.ErrCodeTooManyTagsException:
				fmt.Println(servicediscovery.ErrCodeTooManyTagsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// CreatePublicDnsNamespace example
//
// This example creates a public namespace based on DNS.
func ExampleClient_CreatePublicDnsNamespaceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.CreatePublicDnsNamespaceInput{
		CreatorRequestId: aws.String("example-creator-request-id-0003"),
		Description:      aws.String("Example.com AWS Cloud Map Public DNS Namespace"),
		Name:             aws.String("example-public-dns.com"),
	}

	req := svc.CreatePublicDnsNamespaceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeNamespaceAlreadyExists:
				fmt.Println(servicediscovery.ErrCodeNamespaceAlreadyExists, aerr.Error())
			case servicediscovery.ErrCodeResourceLimitExceeded:
				fmt.Println(servicediscovery.ErrCodeResourceLimitExceeded, aerr.Error())
			case servicediscovery.ErrCodeDuplicateRequest:
				fmt.Println(servicediscovery.ErrCodeDuplicateRequest, aerr.Error())
			case servicediscovery.ErrCodeTooManyTagsException:
				fmt.Println(servicediscovery.ErrCodeTooManyTagsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Example: Create service
//
// Example: Create service
func ExampleClient_CreateServiceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.CreateServiceInput{
		CreatorRequestId: aws.String("567c1193-6b00-4308-bd57-ad38a8822d25"),
		DnsConfig: &servicediscovery.DnsConfig{
			DnsRecords: []servicediscovery.DnsRecord{
				{
					TTL:  aws.Int64(60),
					Type: servicediscovery.RecordTypeA,
				},
			},
			NamespaceId:   aws.String("ns-ylexjili4cdxy3xm"),
			RoutingPolicy: servicediscovery.RoutingPolicyMultivalue,
		},
		Name:        aws.String("myservice"),
		NamespaceId: aws.String("ns-ylexjili4cdxy3xm"),
	}

	req := svc.CreateServiceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeResourceLimitExceeded:
				fmt.Println(servicediscovery.ErrCodeResourceLimitExceeded, aerr.Error())
			case servicediscovery.ErrCodeNamespaceNotFound:
				fmt.Println(servicediscovery.ErrCodeNamespaceNotFound, aerr.Error())
			case servicediscovery.ErrCodeServiceAlreadyExists:
				fmt.Println(servicediscovery.ErrCodeServiceAlreadyExists, aerr.Error())
			case servicediscovery.ErrCodeTooManyTagsException:
				fmt.Println(servicediscovery.ErrCodeTooManyTagsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Example: Delete namespace
//
// Example: Delete namespace
func ExampleClient_DeleteNamespaceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.DeleteNamespaceInput{
		Id: aws.String("ns-ylexjili4cdxy3xm"),
	}

	req := svc.DeleteNamespaceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeNamespaceNotFound:
				fmt.Println(servicediscovery.ErrCodeNamespaceNotFound, aerr.Error())
			case servicediscovery.ErrCodeResourceInUse:
				fmt.Println(servicediscovery.ErrCodeResourceInUse, aerr.Error())
			case servicediscovery.ErrCodeDuplicateRequest:
				fmt.Println(servicediscovery.ErrCodeDuplicateRequest, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Example: Delete service
//
// Example: Delete service
func ExampleClient_DeleteServiceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.DeleteServiceInput{
		Id: aws.String("srv-p5zdwlg5uvvzjita"),
	}

	req := svc.DeleteServiceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
			case servicediscovery.ErrCodeResourceInUse:
				fmt.Println(servicediscovery.ErrCodeResourceInUse, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Example: Deregister a service instance
//
// Example: Deregister a service instance
func ExampleClient_DeregisterInstanceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.DeregisterInstanceInput{
		InstanceId: aws.String("myservice-53"),
		ServiceId:  aws.String("srv-p5zdwlg5uvvzjita"),
	}

	req := svc.DeregisterInstanceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeDuplicateRequest:
				fmt.Println(servicediscovery.ErrCodeDuplicateRequest, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeInstanceNotFound:
				fmt.Println(servicediscovery.ErrCodeInstanceNotFound, aerr.Error())
			case servicediscovery.ErrCodeResourceInUse:
				fmt.Println(servicediscovery.ErrCodeResourceInUse, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Example: Discover registered instances
//
// Example: Discover registered instances
func ExampleClient_DiscoverInstancesRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.DiscoverInstancesInput{
		HealthStatus:  servicediscovery.HealthStatusFilterAll,
		MaxResults:    aws.Int64(10),
		NamespaceName: aws.String("example.com"),
		ServiceName:   aws.String("myservice"),
	}

	req := svc.DiscoverInstancesRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
			case servicediscovery.ErrCodeNamespaceNotFound:
				fmt.Println(servicediscovery.ErrCodeNamespaceNotFound, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeRequestLimitExceeded:
				fmt.Println(servicediscovery.ErrCodeRequestLimitExceeded, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// GetInstance example
//
// This example gets information about a specified instance.
func ExampleClient_GetInstanceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.GetInstanceInput{
		InstanceId: aws.String("i-abcd1234"),
		ServiceId:  aws.String("srv-e4anhexample0004"),
	}

	req := svc.GetInstanceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInstanceNotFound:
				fmt.Println(servicediscovery.ErrCodeInstanceNotFound, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// GetInstancesHealthStatus example
//
// This example gets the current health status of one or more instances that are associate
// with a specified service.
func ExampleClient_GetInstancesHealthStatusRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.GetInstancesHealthStatusInput{
		ServiceId: aws.String("srv-e4anhexample0004"),
	}

	req := svc.GetInstancesHealthStatusRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInstanceNotFound:
				fmt.Println(servicediscovery.ErrCodeInstanceNotFound, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// GetNamespace example
//
// This example gets information about a specified namespace.
func ExampleClient_GetNamespaceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.GetNamespaceInput{
		Id: aws.String("ns-e4anhexample0004"),
	}

	req := svc.GetNamespaceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeNamespaceNotFound:
				fmt.Println(servicediscovery.ErrCodeNamespaceNotFound, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Example: Get operation result
//
// Example: Get operation result
func ExampleClient_GetOperationRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.GetOperationInput{
		OperationId: aws.String("gv4g5meo7ndmeh4fqskygvk23d2fijwa-k9302yzd"),
	}

	req := svc.GetOperationRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeOperationNotFound:
				fmt.Println(servicediscovery.ErrCodeOperationNotFound, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// GetService Example
//
// This example gets the settings for a specified service.
func ExampleClient_GetServiceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.GetServiceInput{
		Id: aws.String("srv-e4anhexample0004"),
	}

	req := svc.GetServiceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Example: List service instances
//
// Example: List service instances
func ExampleClient_ListInstancesRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.ListInstancesInput{
		ServiceId: aws.String("srv-qzpwvt2tfqcegapy"),
	}

	req := svc.ListInstancesRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Example: List namespaces
//
// Example: List namespaces
func ExampleClient_ListNamespacesRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.ListNamespacesInput{}

	req := svc.ListNamespacesRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// ListOperations Example
//
// This example gets the operations that have a STATUS of either PENDING or SUCCESS.
func ExampleClient_ListOperationsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.ListOperationsInput{
		Filters: []servicediscovery.OperationFilter{
			{
				Condition: servicediscovery.FilterConditionIn,
				Name:      servicediscovery.OperationFilterNameStatus,
				Values: []string{
					"PENDING",
					"SUCCESS",
				},
			},
		},
	}

	req := svc.ListOperationsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Example: List services
//
// Example: List services
func ExampleClient_ListServicesRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.ListServicesInput{}

	req := svc.ListServicesRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// ListTagsForResource example
//
// This example lists the tags of a resource.
func ExampleClient_ListTagsForResourceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.ListTagsForResourceInput{
		ResourceARN: aws.String("arn:aws:servicediscovery:us-east-1:123456789012:namespace/ns-ylexjili4cdxy3xm"),
	}

	req := svc.ListTagsForResourceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeResourceNotFoundException:
				fmt.Println(servicediscovery.ErrCodeResourceNotFoundException, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Example: Register Instance
//
// Example: Register Instance
func ExampleClient_RegisterInstanceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.RegisterInstanceInput{
		Attributes: map[string]string{
			"AWS_INSTANCE_IPV4": "172.2.1.3",
			"AWS_INSTANCE_PORT": "808",
		},
		CreatorRequestId: aws.String("7a48a98a-72e6-4849-bfa7-1a458e030d7b"),
		InstanceId:       aws.String("myservice-53"),
		ServiceId:        aws.String("srv-p5zdwlg5uvvzjita"),
	}

	req := svc.RegisterInstanceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeDuplicateRequest:
				fmt.Println(servicediscovery.ErrCodeDuplicateRequest, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeResourceInUse:
				fmt.Println(servicediscovery.ErrCodeResourceInUse, aerr.Error())
			case servicediscovery.ErrCodeResourceLimitExceeded:
				fmt.Println(servicediscovery.ErrCodeResourceLimitExceeded, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// TagResource example
//
// This example adds "Department" and "Project" tags to a resource.
func ExampleClient_TagResourceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.TagResourceInput{
		ResourceARN: aws.String("arn:aws:servicediscovery:us-east-1:123456789012:namespace/ns-ylexjili4cdxy3xm"),
		Tags: []servicediscovery.Tag{
			{
				Key:   aws.String("Department"),
				Value: aws.String("Engineering"),
			},
			{
				Key:   aws.String("Project"),
				Value: aws.String("Zeta"),
			},
		},
	}

	req := svc.TagResourceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeResourceNotFoundException:
				fmt.Println(servicediscovery.ErrCodeResourceNotFoundException, aerr.Error())
			case servicediscovery.ErrCodeTooManyTagsException:
				fmt.Println(servicediscovery.ErrCodeTooManyTagsException, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// UntagResource example
//
// This example removes the "Department" and "Project" tags from a resource.
func ExampleClient_UntagResourceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.UntagResourceInput{
		ResourceARN: aws.String("arn:aws:servicediscovery:us-east-1:123456789012:namespace/ns-ylexjili4cdxy3xm"),
		TagKeys: []string{
			"Project",
			"Department",
		},
	}

	req := svc.UntagResourceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeResourceNotFoundException:
				fmt.Println(servicediscovery.ErrCodeResourceNotFoundException, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// UpdateInstanceCustomHealthStatus Example
//
// This example submits a request to change the health status of an instance associated
// with a service with a custom health check to HEALTHY.
func ExampleClient_UpdateInstanceCustomHealthStatusRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.UpdateInstanceCustomHealthStatusInput{
		InstanceId: aws.String("i-abcd1234"),
		ServiceId:  aws.String("srv-e4anhexample0004"),
		Status:     servicediscovery.CustomHealthStatusHealthy,
	}

	req := svc.UpdateInstanceCustomHealthStatusRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInstanceNotFound:
				fmt.Println(servicediscovery.ErrCodeInstanceNotFound, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
			case servicediscovery.ErrCodeCustomHealthNotFound:
				fmt.Println(servicediscovery.ErrCodeCustomHealthNotFound, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// UpdateService Example
//
// This example submits a request to replace the DnsConfig and HealthCheckConfig settings
// of a specified service.
func ExampleClient_UpdateServiceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := servicediscovery.New(cfg)
	input := &servicediscovery.UpdateServiceInput{
		Id: aws.String("srv-e4anhexample0004"),
		Service: &servicediscovery.ServiceChange{
			DnsConfig: &servicediscovery.DnsConfigChange{
				DnsRecords: []servicediscovery.DnsRecord{
					{
						TTL:  aws.Int64(60),
						Type: servicediscovery.RecordTypeA,
					},
				},
			},
			HealthCheckConfig: &servicediscovery.HealthCheckConfig{
				FailureThreshold: aws.Int64(2),
				ResourcePath:     aws.String("/"),
				Type:             servicediscovery.HealthCheckTypeHttp,
			},
		},
	}

	req := svc.UpdateServiceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeDuplicateRequest:
				fmt.Println(servicediscovery.ErrCodeDuplicateRequest, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
