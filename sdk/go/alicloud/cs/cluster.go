// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
	// Deprecated: Field 'size' has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.
	Size      pulumi.IntPtrOutput `pulumi:"size"`
	SlbId     pulumi.StringOutput `pulumi:"slbId"`
	VpcId     pulumi.StringOutput `pulumi:"vpcId"`
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'CidrBlock'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
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
	// Deprecated: Field 'size' has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.
	Size      *int    `pulumi:"size"`
	SlbId     *string `pulumi:"slbId"`
	VpcId     *string `pulumi:"vpcId"`
	VswitchId *string `pulumi:"vswitchId"`
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
	// Deprecated: Field 'size' has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.
	Size      pulumi.IntPtrInput
	SlbId     pulumi.StringPtrInput
	VpcId     pulumi.StringPtrInput
	VswitchId pulumi.StringPtrInput
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
	// Deprecated: Field 'size' has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.
	Size      *int   `pulumi:"size"`
	VswitchId string `pulumi:"vswitchId"`
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
	// Deprecated: Field 'size' has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.
	Size      pulumi.IntPtrInput
	VswitchId pulumi.StringInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//	ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (i ClusterArray) ToClusterArrayOutput() ClusterArrayOutput {
	return i.ToClusterArrayOutputWithContext(context.Background())
}

func (i ClusterArray) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterArrayOutput)
}

// ClusterMapInput is an input type that accepts ClusterMap and ClusterMapOutput values.
// You can construct a concrete instance of `ClusterMapInput` via:
//
//	ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.AgentVersion }).(pulumi.StringOutput)
}

func (o ClusterOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.CidrBlock }).(pulumi.StringOutput)
}

func (o ClusterOutput) DiskCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.DiskCategory }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) DiskSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntPtrOutput { return v.DiskSize }).(pulumi.IntPtrOutput)
}

func (o ClusterOutput) ImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ImageId }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

func (o ClusterOutput) IsOutdated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.IsOutdated }).(pulumi.BoolPtrOutput)
}

func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClusterOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.NamePrefix }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) NeedSlb() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.NeedSlb }).(pulumi.BoolPtrOutput)
}

func (o ClusterOutput) NodeNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntPtrOutput { return v.NodeNumber }).(pulumi.IntPtrOutput)
}

func (o ClusterOutput) Nodes() ClusterNodeArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterNodeArrayOutput { return v.Nodes }).(ClusterNodeArrayOutput)
}

func (o ClusterOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

func (o ClusterOutput) ReleaseEip() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.ReleaseEip }).(pulumi.BoolPtrOutput)
}

func (o ClusterOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// Deprecated: Field 'size' has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.
func (o ClusterOutput) Size() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntPtrOutput { return v.Size }).(pulumi.IntPtrOutput)
}

func (o ClusterOutput) SlbId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.SlbId }).(pulumi.StringOutput)
}

func (o ClusterOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

func (o ClusterOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].([]*Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].(map[string]*Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterArrayInput)(nil)).Elem(), ClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMapInput)(nil)).Elem(), ClusterMap{})
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}
