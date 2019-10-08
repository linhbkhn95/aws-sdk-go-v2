// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RestoreDBInstanceFromDBSnapshotInput struct {
	_ struct{} `type:"structure"`

	// A value that indicates whether minor version upgrades are applied automatically
	// to the DB instance during the maintenance window.
	AutoMinorVersionUpgrade *bool `type:"boolean"`

	// The Availability Zone (AZ) where the DB instance will be created.
	//
	// Default: A random, system-chosen Availability Zone.
	//
	// Constraint: You can't specify the AvailabilityZone parameter if the DB instance
	// is a Multi-AZ deployment.
	//
	// Example: us-east-1a
	AvailabilityZone *string `type:"string"`

	// A value that indicates whether to copy all tags from the restored DB instance
	// to snapshots of the DB instance. By default, tags are not copied.
	CopyTagsToSnapshot *bool `type:"boolean"`

	// The compute and memory capacity of the Amazon RDS DB instance, for example,
	// db.m4.large. Not all DB instance classes are available in all AWS Regions,
	// or for all database engines. For the full list of DB instance classes, and
	// availability for your engine, see DB Instance Class (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide.
	//
	// Default: The same DBInstanceClass as the original DB instance.
	DBInstanceClass *string `type:"string"`

	// Name of the DB instance to create from the DB snapshot. This parameter isn't
	// case-sensitive.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 numbers, letters, or hyphens
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	//
	// Example: my-snapshot-id
	//
	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `type:"string" required:"true"`

	// The database name for the restored DB instance.
	//
	// This parameter doesn't apply to the MySQL, PostgreSQL, or MariaDB engines.
	DBName *string `type:"string"`

	// The name of the DB parameter group to associate with this DB instance.
	//
	// If you do not specify a value for DBParameterGroupName, then the default
	// DBParameterGroup for the specified DB engine is used.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing DBParameterGroup.
	//
	//    * Must be 1 to 255 letters, numbers, or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens.
	DBParameterGroupName *string `type:"string"`

	// The identifier for the DB snapshot to restore from.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DBSnapshot.
	//
	//    * If you are restoring from a shared manual DB snapshot, the DBSnapshotIdentifier
	//    must be the ARN of the shared DB snapshot.
	//
	// DBSnapshotIdentifier is a required field
	DBSnapshotIdentifier *string `type:"string" required:"true"`

	// The DB subnet group name to use for the new instance.
	//
	// Constraints: If supplied, must match the name of an existing DBSubnetGroup.
	//
	// Example: mySubnetgroup
	DBSubnetGroupName *string `type:"string"`

	// A value that indicates whether the DB instance has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled. For more information, see Deleting a DB
	// Instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).
	DeletionProtection *bool `type:"boolean"`

	// Specify the Active Directory directory ID to restore the DB instance in.
	// The domain must be created prior to this operation. Currently, only Microsoft
	// SQL Server and Oracle DB instances can be created in an Active Directory
	// Domain.
	//
	// For Microsoft SQL Server DB instances, Amazon RDS can use Windows Authentication
	// to authenticate users that connect to the DB instance. For more information,
	// see Using Windows Authentication with an Amazon RDS DB Instance Running Microsoft
	// SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_SQLServerWinAuth.html)
	// in the Amazon RDS User Guide.
	//
	// For Oracle DB instances, Amazon RDS can use Kerberos Authentication to authenticate
	// users that connect to the DB instance. For more information, see Using Kerberos
	// Authentication with Amazon RDS for Oracle (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-kerberos.html)
	// in the Amazon RDS User Guide.
	Domain *string `type:"string"`

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service.
	DomainIAMRoleName *string `type:"string"`

	// The list of logs that the restored DB instance is to export to CloudWatch
	// Logs. The values in the list depend on the DB engine being used. For more
	// information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Aurora User Guide.
	EnableCloudwatchLogsExports []string `type:"list"`

	// A value that indicates whether to enable mapping of AWS Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	// For information about the supported DB engines, see CreateDBInstance.
	//
	// For more information about IAM database authentication, see IAM Database
	// Authentication for MySQL and PostgreSQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon RDS User Guide.
	EnableIAMDatabaseAuthentication *bool `type:"boolean"`

	// The database engine to use for the new instance.
	//
	// Default: The same as source
	//
	// Constraint: Must be compatible with the engine of the source. For example,
	// you can restore a MariaDB 10.1 DB instance from a MySQL 5.6 snapshot.
	//
	// Valid Values:
	//
	//    * mariadb
	//
	//    * mysql
	//
	//    * oracle-ee
	//
	//    * oracle-se2
	//
	//    * oracle-se1
	//
	//    * oracle-se
	//
	//    * postgres
	//
	//    * sqlserver-ee
	//
	//    * sqlserver-se
	//
	//    * sqlserver-ex
	//
	//    * sqlserver-web
	Engine *string `type:"string"`

	// Specifies the amount of provisioned IOPS for the DB instance, expressed in
	// I/O operations per second. If this parameter is not specified, the IOPS value
	// is taken from the backup. If this parameter is set to 0, the new instance
	// is converted to a non-PIOPS instance. The conversion takes additional time,
	// though your DB instance is available for connections before the conversion
	// starts.
	//
	// The provisioned IOPS value must follow the requirements for your database
	// engine. For more information, see Amazon RDS Provisioned IOPS Storage to
	// Improve Performance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS)
	// in the Amazon RDS User Guide.
	//
	// Constraints: Must be an integer greater than 1000.
	Iops *int64 `type:"integer"`

	// License model information for the restored DB instance.
	//
	// Default: Same as source.
	//
	// Valid values: license-included | bring-your-own-license | general-public-license
	LicenseModel *string `type:"string"`

	// A value that indicates whether the DB instance is a Multi-AZ deployment.
	//
	// Constraint: You can't specify the AvailabilityZone parameter if the DB instance
	// is a Multi-AZ deployment.
	MultiAZ *bool `type:"boolean"`

	// The name of the option group to be used for the restored DB instance.
	//
	// Permanent options, such as the TDE option for Oracle Advanced Security TDE,
	// can't be removed from an option group, and that option group can't be removed
	// from a DB instance once it is associated with a DB instance
	OptionGroupName *string `type:"string"`

	// The port number on which the database accepts connections.
	//
	// Default: The same port as the original DB instance
	//
	// Constraints: Value must be 1150-65535
	Port *int64 `type:"integer"`

	// The number of CPU cores and the number of threads per core for the DB instance
	// class of the DB instance.
	ProcessorFeatures []ProcessorFeature `locationNameList:"ProcessorFeature" type:"list"`

	// A value that indicates whether the DB instance is publicly accessible. When
	// the DB instance is publicly accessible, it is an Internet-facing instance
	// with a publicly resolvable DNS name, which resolves to a public IP address.
	// When the DB instance is not publicly accessible, it is an internal instance
	// with a DNS name that resolves to a private IP address. For more information,
	// see CreateDBInstance.
	PubliclyAccessible *bool `type:"boolean"`

	// Specifies the storage type to be associated with the DB instance.
	//
	// Valid values: standard | gp2 | io1
	//
	// If you specify io1, you must also include a value for the Iops parameter.
	//
	// Default: io1 if the Iops parameter is specified, otherwise gp2
	StorageType *string `type:"string"`

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// The ARN from the key store with which to associate the instance for TDE encryption.
	TdeCredentialArn *string `type:"string"`

	// The password for the given ARN from the key store in order to access the
	// device.
	TdeCredentialPassword *string `type:"string"`

	// A value that indicates whether the DB instance class of the DB instance uses
	// its default processor features.
	UseDefaultProcessorFeatures *bool `type:"boolean"`

	// A list of EC2 VPC security groups to associate with this DB instance.
	//
	// Default: The default EC2 VPC security group for the DB subnet group's VPC.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s RestoreDBInstanceFromDBSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreDBInstanceFromDBSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreDBInstanceFromDBSnapshotInput"}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}

	if s.DBSnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBSnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RestoreDBInstanceFromDBSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB instance.
	//
	// This data type is used as a response element in the DescribeDBInstances action.
	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s RestoreDBInstanceFromDBSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opRestoreDBInstanceFromDBSnapshot = "RestoreDBInstanceFromDBSnapshot"

// RestoreDBInstanceFromDBSnapshotRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Creates a new DB instance from a DB snapshot. The target database is created
// from the source database restore point with the most of original configuration
// with the default security group and the default DB parameter group. By default,
// the new DB instance is created as a single-AZ deployment except when the
// instance is a SQL Server instance that has an option group that is associated
// with mirroring; in this case, the instance becomes a mirrored AZ deployment
// and not a single-AZ deployment.
//
// If your intent is to replace your original DB instance with the new, restored
// DB instance, then rename your original DB instance before you call the RestoreDBInstanceFromDBSnapshot
// action. RDS doesn't allow two DB instances with the same name. Once you have
// renamed your original DB instance with a different identifier, then you can
// pass the original name of the DB instance as the DBInstanceIdentifier in
// the call to the RestoreDBInstanceFromDBSnapshot action. The result is that
// you will replace the original DB instance with the DB instance created from
// the snapshot.
//
// If you are restoring from a shared manual DB snapshot, the DBSnapshotIdentifier
// must be the ARN of the shared DB snapshot.
//
// This command doesn't apply to Aurora MySQL and Aurora PostgreSQL. For Aurora,
// use RestoreDBClusterFromSnapshot.
//
//    // Example sending a request using RestoreDBInstanceFromDBSnapshotRequest.
//    req := client.RestoreDBInstanceFromDBSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RestoreDBInstanceFromDBSnapshot
func (c *Client) RestoreDBInstanceFromDBSnapshotRequest(input *RestoreDBInstanceFromDBSnapshotInput) RestoreDBInstanceFromDBSnapshotRequest {
	op := &aws.Operation{
		Name:       opRestoreDBInstanceFromDBSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreDBInstanceFromDBSnapshotInput{}
	}

	req := c.newRequest(op, input, &RestoreDBInstanceFromDBSnapshotOutput{})
	return RestoreDBInstanceFromDBSnapshotRequest{Request: req, Input: input, Copy: c.RestoreDBInstanceFromDBSnapshotRequest}
}

// RestoreDBInstanceFromDBSnapshotRequest is the request type for the
// RestoreDBInstanceFromDBSnapshot API operation.
type RestoreDBInstanceFromDBSnapshotRequest struct {
	*aws.Request
	Input *RestoreDBInstanceFromDBSnapshotInput
	Copy  func(*RestoreDBInstanceFromDBSnapshotInput) RestoreDBInstanceFromDBSnapshotRequest
}

// Send marshals and sends the RestoreDBInstanceFromDBSnapshot API request.
func (r RestoreDBInstanceFromDBSnapshotRequest) Send(ctx context.Context) (*RestoreDBInstanceFromDBSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreDBInstanceFromDBSnapshotResponse{
		RestoreDBInstanceFromDBSnapshotOutput: r.Request.Data.(*RestoreDBInstanceFromDBSnapshotOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreDBInstanceFromDBSnapshotResponse is the response type for the
// RestoreDBInstanceFromDBSnapshot API operation.
type RestoreDBInstanceFromDBSnapshotResponse struct {
	*RestoreDBInstanceFromDBSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreDBInstanceFromDBSnapshot request.
func (r *RestoreDBInstanceFromDBSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}