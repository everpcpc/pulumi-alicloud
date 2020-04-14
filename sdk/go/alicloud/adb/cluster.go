// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package adb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a ADB cluster resource. A ADB cluster is an isolated database
// environment in the cloud. A ADB cluster can contain multiple user-created
// databases.
//
// > **NOTE:** Available in v1.71.0+.
type Cluster struct {
	pulumi.CustomResourceState

	// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
	AutoRenewPeriod pulumi.IntPtrOutput `pulumi:"autoRenewPeriod"`
	// Cluster category. Value options: `Basic`, `Cluster`.
	DbClusterCategory pulumi.StringOutput `pulumi:"dbClusterCategory"`
	// Cluster version. Value options: `3.0`, Default to `3.0`.
	DbClusterVersion pulumi.StringPtrOutput `pulumi:"dbClusterVersion"`
	// The dbNodeClass of cluster node.
	DbNodeClass pulumi.StringOutput `pulumi:"dbNodeClass"`
	// The dbNodeCount of cluster node.
	DbNodeCount pulumi.IntOutput `pulumi:"dbNodeCount"`
	// The dbNodeStorage of cluster node.
	DbNodeStorage pulumi.IntOutput `pulumi:"dbNodeStorage"`
	// The description of cluster.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime pulumi.StringOutput `pulumi:"maintainTime"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
	PayType pulumi.StringPtrOutput `pulumi:"payType"`
	// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
	RenewalStatus pulumi.StringPtrOutput `pulumi:"renewalStatus"`
	// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps pulumi.StringArrayOutput `pulumi:"securityIps"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
	// The Zone to launch the DB cluster.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.DbClusterCategory == nil {
		return nil, errors.New("missing required argument 'DbClusterCategory'")
	}
	if args == nil || args.DbNodeClass == nil {
		return nil, errors.New("missing required argument 'DbNodeClass'")
	}
	if args == nil || args.DbNodeCount == nil {
		return nil, errors.New("missing required argument 'DbNodeCount'")
	}
	if args == nil || args.DbNodeStorage == nil {
		return nil, errors.New("missing required argument 'DbNodeStorage'")
	}
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("alicloud:adb/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("alicloud:adb/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
	AutoRenewPeriod *int `pulumi:"autoRenewPeriod"`
	// Cluster category. Value options: `Basic`, `Cluster`.
	DbClusterCategory *string `pulumi:"dbClusterCategory"`
	// Cluster version. Value options: `3.0`, Default to `3.0`.
	DbClusterVersion *string `pulumi:"dbClusterVersion"`
	// The dbNodeClass of cluster node.
	DbNodeClass *string `pulumi:"dbNodeClass"`
	// The dbNodeCount of cluster node.
	DbNodeCount *int `pulumi:"dbNodeCount"`
	// The dbNodeStorage of cluster node.
	DbNodeStorage *int `pulumi:"dbNodeStorage"`
	// The description of cluster.
	Description *string `pulumi:"description"`
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime *string `pulumi:"maintainTime"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
	PayType *string `pulumi:"payType"`
	// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period *int `pulumi:"period"`
	// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
	RenewalStatus *string `pulumi:"renewalStatus"`
	// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps []string `pulumi:"securityIps"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId *string `pulumi:"vswitchId"`
	// The Zone to launch the DB cluster.
	ZoneId *string `pulumi:"zoneId"`
}

type ClusterState struct {
	// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
	AutoRenewPeriod pulumi.IntPtrInput
	// Cluster category. Value options: `Basic`, `Cluster`.
	DbClusterCategory pulumi.StringPtrInput
	// Cluster version. Value options: `3.0`, Default to `3.0`.
	DbClusterVersion pulumi.StringPtrInput
	// The dbNodeClass of cluster node.
	DbNodeClass pulumi.StringPtrInput
	// The dbNodeCount of cluster node.
	DbNodeCount pulumi.IntPtrInput
	// The dbNodeStorage of cluster node.
	DbNodeStorage pulumi.IntPtrInput
	// The description of cluster.
	Description pulumi.StringPtrInput
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime pulumi.StringPtrInput
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
	PayType pulumi.StringPtrInput
	// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period pulumi.IntPtrInput
	// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
	RenewalStatus pulumi.StringPtrInput
	// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrInput
	// The Zone to launch the DB cluster.
	ZoneId pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
	AutoRenewPeriod *int `pulumi:"autoRenewPeriod"`
	// Cluster category. Value options: `Basic`, `Cluster`.
	DbClusterCategory string `pulumi:"dbClusterCategory"`
	// Cluster version. Value options: `3.0`, Default to `3.0`.
	DbClusterVersion *string `pulumi:"dbClusterVersion"`
	// The dbNodeClass of cluster node.
	DbNodeClass string `pulumi:"dbNodeClass"`
	// The dbNodeCount of cluster node.
	DbNodeCount int `pulumi:"dbNodeCount"`
	// The dbNodeStorage of cluster node.
	DbNodeStorage int `pulumi:"dbNodeStorage"`
	// The description of cluster.
	Description *string `pulumi:"description"`
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime *string `pulumi:"maintainTime"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
	PayType *string `pulumi:"payType"`
	// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period *int `pulumi:"period"`
	// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
	RenewalStatus *string `pulumi:"renewalStatus"`
	// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps []string `pulumi:"securityIps"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId *string `pulumi:"vswitchId"`
	// The Zone to launch the DB cluster.
	ZoneId *string `pulumi:"zoneId"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
	AutoRenewPeriod pulumi.IntPtrInput
	// Cluster category. Value options: `Basic`, `Cluster`.
	DbClusterCategory pulumi.StringInput
	// Cluster version. Value options: `3.0`, Default to `3.0`.
	DbClusterVersion pulumi.StringPtrInput
	// The dbNodeClass of cluster node.
	DbNodeClass pulumi.StringInput
	// The dbNodeCount of cluster node.
	DbNodeCount pulumi.IntInput
	// The dbNodeStorage of cluster node.
	DbNodeStorage pulumi.IntInput
	// The description of cluster.
	Description pulumi.StringPtrInput
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime pulumi.StringPtrInput
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
	PayType pulumi.StringPtrInput
	// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period pulumi.IntPtrInput
	// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
	RenewalStatus pulumi.StringPtrInput
	// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrInput
	// The Zone to launch the DB cluster.
	ZoneId pulumi.StringPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
