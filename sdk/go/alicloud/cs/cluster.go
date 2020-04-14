// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Cluster struct {
	pulumi.CustomResourceState

	AgentVersion    pulumi.StringOutput    `pulumi:"agentVersion"`
	CidrBlock       pulumi.StringOutput    `pulumi:"cidrBlock"`
	DiskCategory    pulumi.StringPtrOutput `pulumi:"diskCategory"`
	DiskSize        pulumi.IntPtrOutput    `pulumi:"diskSize"`
	ImageId         pulumi.StringPtrOutput `pulumi:"imageId"`
	InstanceType    pulumi.StringOutput    `pulumi:"instanceType"`
	IsOutdated      pulumi.BoolPtrOutput   `pulumi:"isOutdated"`
	Name            pulumi.StringOutput    `pulumi:"name"`
	NamePrefix      pulumi.StringPtrOutput `pulumi:"namePrefix"`
	NeedSlb         pulumi.BoolPtrOutput   `pulumi:"needSlb"`
	NodeNumber      pulumi.IntPtrOutput    `pulumi:"nodeNumber"`
	Nodes           ClusterNodeArrayOutput `pulumi:"nodes"`
	Password        pulumi.StringOutput    `pulumi:"password"`
	ReleaseEip      pulumi.BoolPtrOutput   `pulumi:"releaseEip"`
	SecurityGroupId pulumi.StringOutput    `pulumi:"securityGroupId"`
	Size            pulumi.IntPtrOutput    `pulumi:"size"`
	SlbId           pulumi.StringOutput    `pulumi:"slbId"`
	VpcId           pulumi.StringOutput    `pulumi:"vpcId"`
	VswitchId       pulumi.StringOutput    `pulumi:"vswitchId"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.CidrBlock == nil {
		return nil, errors.New("missing required argument 'CidrBlock'")
	}
	if args == nil || args.InstanceType == nil {
		return nil, errors.New("missing required argument 'InstanceType'")
	}
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil || args.VswitchId == nil {
		return nil, errors.New("missing required argument 'VswitchId'")
	}
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("alicloud:cs/cluster:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:cs/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	AgentVersion    *string       `pulumi:"agentVersion"`
	CidrBlock       *string       `pulumi:"cidrBlock"`
	DiskCategory    *string       `pulumi:"diskCategory"`
	DiskSize        *int          `pulumi:"diskSize"`
	ImageId         *string       `pulumi:"imageId"`
	InstanceType    *string       `pulumi:"instanceType"`
	IsOutdated      *bool         `pulumi:"isOutdated"`
	Name            *string       `pulumi:"name"`
	NamePrefix      *string       `pulumi:"namePrefix"`
	NeedSlb         *bool         `pulumi:"needSlb"`
	NodeNumber      *int          `pulumi:"nodeNumber"`
	Nodes           []ClusterNode `pulumi:"nodes"`
	Password        *string       `pulumi:"password"`
	ReleaseEip      *bool         `pulumi:"releaseEip"`
	SecurityGroupId *string       `pulumi:"securityGroupId"`
	Size            *int          `pulumi:"size"`
	SlbId           *string       `pulumi:"slbId"`
	VpcId           *string       `pulumi:"vpcId"`
	VswitchId       *string       `pulumi:"vswitchId"`
}

type ClusterState struct {
	AgentVersion    pulumi.StringPtrInput
	CidrBlock       pulumi.StringPtrInput
	DiskCategory    pulumi.StringPtrInput
	DiskSize        pulumi.IntPtrInput
	ImageId         pulumi.StringPtrInput
	InstanceType    pulumi.StringPtrInput
	IsOutdated      pulumi.BoolPtrInput
	Name            pulumi.StringPtrInput
	NamePrefix      pulumi.StringPtrInput
	NeedSlb         pulumi.BoolPtrInput
	NodeNumber      pulumi.IntPtrInput
	Nodes           ClusterNodeArrayInput
	Password        pulumi.StringPtrInput
	ReleaseEip      pulumi.BoolPtrInput
	SecurityGroupId pulumi.StringPtrInput
	Size            pulumi.IntPtrInput
	SlbId           pulumi.StringPtrInput
	VpcId           pulumi.StringPtrInput
	VswitchId       pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	CidrBlock    string  `pulumi:"cidrBlock"`
	DiskCategory *string `pulumi:"diskCategory"`
	DiskSize     *int    `pulumi:"diskSize"`
	ImageId      *string `pulumi:"imageId"`
	InstanceType string  `pulumi:"instanceType"`
	IsOutdated   *bool   `pulumi:"isOutdated"`
	Name         *string `pulumi:"name"`
	NamePrefix   *string `pulumi:"namePrefix"`
	NeedSlb      *bool   `pulumi:"needSlb"`
	NodeNumber   *int    `pulumi:"nodeNumber"`
	Password     string  `pulumi:"password"`
	ReleaseEip   *bool   `pulumi:"releaseEip"`
	Size         *int    `pulumi:"size"`
	VswitchId    string  `pulumi:"vswitchId"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	CidrBlock    pulumi.StringInput
	DiskCategory pulumi.StringPtrInput
	DiskSize     pulumi.IntPtrInput
	ImageId      pulumi.StringPtrInput
	InstanceType pulumi.StringInput
	IsOutdated   pulumi.BoolPtrInput
	Name         pulumi.StringPtrInput
	NamePrefix   pulumi.StringPtrInput
	NeedSlb      pulumi.BoolPtrInput
	NodeNumber   pulumi.IntPtrInput
	Password     pulumi.StringInput
	ReleaseEip   pulumi.BoolPtrInput
	Size         pulumi.IntPtrInput
	VswitchId    pulumi.StringInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
