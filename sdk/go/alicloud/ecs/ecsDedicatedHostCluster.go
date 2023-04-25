// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ECS Dedicated Host Cluster resource.
//
// For information about ECS Dedicated Host Cluster and how to use it, see [What is Dedicated Host Cluster](https://www.alibabacloud.com/help/en/doc-detail/184667.html).
//
// > **NOTE:** Available in v1.146.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleZones, err := alicloud.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewEcsDedicatedHostCluster(ctx, "exampleEcsDedicatedHostCluster", &ecs.EcsDedicatedHostClusterArgs{
//				DedicatedHostClusterName: pulumi.String("example_value"),
//				Description:              pulumi.String("example_value"),
//				ZoneId:                   *pulumi.String(exampleZones.Zones[0].Id),
//				Tags: pulumi.AnyMap{
//					"Create": pulumi.Any("TF"),
//					"For":    pulumi.Any("DDH_Cluster_Test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ECS Dedicated Host Cluster can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ecs/ecsDedicatedHostCluster:EcsDedicatedHostCluster example <id>
//
// ```
type EcsDedicatedHostCluster struct {
	pulumi.CustomResourceState

	// The name of the dedicated host cluster. The name must be `2` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It cannot contain `http://` or `https://`.
	DedicatedHostClusterName pulumi.StringPtrOutput `pulumi:"dedicatedHostClusterName"`
	// The description of the dedicated host cluster. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The dry run.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the zone in which to create the dedicated host cluster.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewEcsDedicatedHostCluster registers a new resource with the given unique name, arguments, and options.
func NewEcsDedicatedHostCluster(ctx *pulumi.Context,
	name string, args *EcsDedicatedHostClusterArgs, opts ...pulumi.ResourceOption) (*EcsDedicatedHostCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	var resource EcsDedicatedHostCluster
	err := ctx.RegisterResource("alicloud:ecs/ecsDedicatedHostCluster:EcsDedicatedHostCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEcsDedicatedHostCluster gets an existing EcsDedicatedHostCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEcsDedicatedHostCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EcsDedicatedHostClusterState, opts ...pulumi.ResourceOption) (*EcsDedicatedHostCluster, error) {
	var resource EcsDedicatedHostCluster
	err := ctx.ReadResource("alicloud:ecs/ecsDedicatedHostCluster:EcsDedicatedHostCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EcsDedicatedHostCluster resources.
type ecsDedicatedHostClusterState struct {
	// The name of the dedicated host cluster. The name must be `2` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It cannot contain `http://` or `https://`.
	DedicatedHostClusterName *string `pulumi:"dedicatedHostClusterName"`
	// The description of the dedicated host cluster. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
	Description *string `pulumi:"description"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the zone in which to create the dedicated host cluster.
	ZoneId *string `pulumi:"zoneId"`
}

type EcsDedicatedHostClusterState struct {
	// The name of the dedicated host cluster. The name must be `2` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It cannot contain `http://` or `https://`.
	DedicatedHostClusterName pulumi.StringPtrInput
	// The description of the dedicated host cluster. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
	Description pulumi.StringPtrInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the zone in which to create the dedicated host cluster.
	ZoneId pulumi.StringPtrInput
}

func (EcsDedicatedHostClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsDedicatedHostClusterState)(nil)).Elem()
}

type ecsDedicatedHostClusterArgs struct {
	// The name of the dedicated host cluster. The name must be `2` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It cannot contain `http://` or `https://`.
	DedicatedHostClusterName *string `pulumi:"dedicatedHostClusterName"`
	// The description of the dedicated host cluster. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
	Description *string `pulumi:"description"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the zone in which to create the dedicated host cluster.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a EcsDedicatedHostCluster resource.
type EcsDedicatedHostClusterArgs struct {
	// The name of the dedicated host cluster. The name must be `2` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It cannot contain `http://` or `https://`.
	DedicatedHostClusterName pulumi.StringPtrInput
	// The description of the dedicated host cluster. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
	Description pulumi.StringPtrInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the zone in which to create the dedicated host cluster.
	ZoneId pulumi.StringInput
}

func (EcsDedicatedHostClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsDedicatedHostClusterArgs)(nil)).Elem()
}

type EcsDedicatedHostClusterInput interface {
	pulumi.Input

	ToEcsDedicatedHostClusterOutput() EcsDedicatedHostClusterOutput
	ToEcsDedicatedHostClusterOutputWithContext(ctx context.Context) EcsDedicatedHostClusterOutput
}

func (*EcsDedicatedHostCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsDedicatedHostCluster)(nil)).Elem()
}

func (i *EcsDedicatedHostCluster) ToEcsDedicatedHostClusterOutput() EcsDedicatedHostClusterOutput {
	return i.ToEcsDedicatedHostClusterOutputWithContext(context.Background())
}

func (i *EcsDedicatedHostCluster) ToEcsDedicatedHostClusterOutputWithContext(ctx context.Context) EcsDedicatedHostClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsDedicatedHostClusterOutput)
}

// EcsDedicatedHostClusterArrayInput is an input type that accepts EcsDedicatedHostClusterArray and EcsDedicatedHostClusterArrayOutput values.
// You can construct a concrete instance of `EcsDedicatedHostClusterArrayInput` via:
//
//	EcsDedicatedHostClusterArray{ EcsDedicatedHostClusterArgs{...} }
type EcsDedicatedHostClusterArrayInput interface {
	pulumi.Input

	ToEcsDedicatedHostClusterArrayOutput() EcsDedicatedHostClusterArrayOutput
	ToEcsDedicatedHostClusterArrayOutputWithContext(context.Context) EcsDedicatedHostClusterArrayOutput
}

type EcsDedicatedHostClusterArray []EcsDedicatedHostClusterInput

func (EcsDedicatedHostClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsDedicatedHostCluster)(nil)).Elem()
}

func (i EcsDedicatedHostClusterArray) ToEcsDedicatedHostClusterArrayOutput() EcsDedicatedHostClusterArrayOutput {
	return i.ToEcsDedicatedHostClusterArrayOutputWithContext(context.Background())
}

func (i EcsDedicatedHostClusterArray) ToEcsDedicatedHostClusterArrayOutputWithContext(ctx context.Context) EcsDedicatedHostClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsDedicatedHostClusterArrayOutput)
}

// EcsDedicatedHostClusterMapInput is an input type that accepts EcsDedicatedHostClusterMap and EcsDedicatedHostClusterMapOutput values.
// You can construct a concrete instance of `EcsDedicatedHostClusterMapInput` via:
//
//	EcsDedicatedHostClusterMap{ "key": EcsDedicatedHostClusterArgs{...} }
type EcsDedicatedHostClusterMapInput interface {
	pulumi.Input

	ToEcsDedicatedHostClusterMapOutput() EcsDedicatedHostClusterMapOutput
	ToEcsDedicatedHostClusterMapOutputWithContext(context.Context) EcsDedicatedHostClusterMapOutput
}

type EcsDedicatedHostClusterMap map[string]EcsDedicatedHostClusterInput

func (EcsDedicatedHostClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsDedicatedHostCluster)(nil)).Elem()
}

func (i EcsDedicatedHostClusterMap) ToEcsDedicatedHostClusterMapOutput() EcsDedicatedHostClusterMapOutput {
	return i.ToEcsDedicatedHostClusterMapOutputWithContext(context.Background())
}

func (i EcsDedicatedHostClusterMap) ToEcsDedicatedHostClusterMapOutputWithContext(ctx context.Context) EcsDedicatedHostClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsDedicatedHostClusterMapOutput)
}

type EcsDedicatedHostClusterOutput struct{ *pulumi.OutputState }

func (EcsDedicatedHostClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsDedicatedHostCluster)(nil)).Elem()
}

func (o EcsDedicatedHostClusterOutput) ToEcsDedicatedHostClusterOutput() EcsDedicatedHostClusterOutput {
	return o
}

func (o EcsDedicatedHostClusterOutput) ToEcsDedicatedHostClusterOutputWithContext(ctx context.Context) EcsDedicatedHostClusterOutput {
	return o
}

// The name of the dedicated host cluster. The name must be `2` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It cannot contain `http://` or `https://`.
func (o EcsDedicatedHostClusterOutput) DedicatedHostClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsDedicatedHostCluster) pulumi.StringPtrOutput { return v.DedicatedHostClusterName }).(pulumi.StringPtrOutput)
}

// The description of the dedicated host cluster. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
func (o EcsDedicatedHostClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsDedicatedHostCluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The dry run.
func (o EcsDedicatedHostClusterOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EcsDedicatedHostCluster) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// A mapping of tags to assign to the resource.
func (o EcsDedicatedHostClusterOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *EcsDedicatedHostCluster) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The ID of the zone in which to create the dedicated host cluster.
func (o EcsDedicatedHostClusterOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsDedicatedHostCluster) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type EcsDedicatedHostClusterArrayOutput struct{ *pulumi.OutputState }

func (EcsDedicatedHostClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsDedicatedHostCluster)(nil)).Elem()
}

func (o EcsDedicatedHostClusterArrayOutput) ToEcsDedicatedHostClusterArrayOutput() EcsDedicatedHostClusterArrayOutput {
	return o
}

func (o EcsDedicatedHostClusterArrayOutput) ToEcsDedicatedHostClusterArrayOutputWithContext(ctx context.Context) EcsDedicatedHostClusterArrayOutput {
	return o
}

func (o EcsDedicatedHostClusterArrayOutput) Index(i pulumi.IntInput) EcsDedicatedHostClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EcsDedicatedHostCluster {
		return vs[0].([]*EcsDedicatedHostCluster)[vs[1].(int)]
	}).(EcsDedicatedHostClusterOutput)
}

type EcsDedicatedHostClusterMapOutput struct{ *pulumi.OutputState }

func (EcsDedicatedHostClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsDedicatedHostCluster)(nil)).Elem()
}

func (o EcsDedicatedHostClusterMapOutput) ToEcsDedicatedHostClusterMapOutput() EcsDedicatedHostClusterMapOutput {
	return o
}

func (o EcsDedicatedHostClusterMapOutput) ToEcsDedicatedHostClusterMapOutputWithContext(ctx context.Context) EcsDedicatedHostClusterMapOutput {
	return o
}

func (o EcsDedicatedHostClusterMapOutput) MapIndex(k pulumi.StringInput) EcsDedicatedHostClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EcsDedicatedHostCluster {
		return vs[0].(map[string]*EcsDedicatedHostCluster)[vs[1].(string)]
	}).(EcsDedicatedHostClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EcsDedicatedHostClusterInput)(nil)).Elem(), &EcsDedicatedHostCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsDedicatedHostClusterArrayInput)(nil)).Elem(), EcsDedicatedHostClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsDedicatedHostClusterMapInput)(nil)).Elem(), EcsDedicatedHostClusterMap{})
	pulumi.RegisterOutputType(EcsDedicatedHostClusterOutput{})
	pulumi.RegisterOutputType(EcsDedicatedHostClusterArrayOutput{})
	pulumi.RegisterOutputType(EcsDedicatedHostClusterMapOutput{})
}
