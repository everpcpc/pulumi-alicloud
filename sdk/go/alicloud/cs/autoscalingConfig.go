// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource will help you configure auto scaling for the kubernetes cluster.
//
// > **NOTE:** Available in v1.127.0+.
// **NOTE:** From version 1.164.0, support for specifying whether to allow the scale-in of nodes by parameter `scaleDownEnabled`.
// **NOTE:** From version 1.164.0, support for selecting the policy for selecting which node pool to scale by parameter `expander`.
//
// ## Example Usage
//
// If you do not have an existing cluster, you need to create an ACK cluster through cs.ManagedKubernetes first, and then configure automatic scaling.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cs.NewAutoscalingConfig(ctx, "default", &cs.AutoscalingConfigArgs{
//				ClusterId:               pulumi.Any(alicloud_cs_managed_kubernetes.Default[0].Id),
//				CoolDownDuration:        pulumi.String("10m"),
//				UnneededDuration:        pulumi.String("10m"),
//				UtilizationThreshold:    pulumi.String("0.5"),
//				GpuUtilizationThreshold: pulumi.String("0.5"),
//				ScanInterval:            pulumi.String("30s"),
//				ScaleDownEnabled:        pulumi.Bool(true),
//				Expander:                pulumi.String("least-waste"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type AutoscalingConfig struct {
	pulumi.CustomResourceState

	// The id of kubernetes cluster.
	ClusterId pulumi.StringPtrOutput `pulumi:"clusterId"`
	// The cool down duration. Default is `10m`. If the delay (cooldown) value is set too long, there could be complaints that the Horizontal Pod Autoscaler is not responsive to workload changes. However, if the delay value is set too short, the scale of the replicas set may keep thrashing as usual.
	CoolDownDuration pulumi.StringPtrOutput `pulumi:"coolDownDuration"`
	// The policy for selecting which node pool to scale. Valid values: `least-waste`, `random`, `priority`. For more information on these policies, see [Configure auto scaling](https://www.alibabacloud.com/help/en/container-service-for-kubernetes/latest/auto-scaling-of-nodes#section-3bg-2ko-inl)
	Expander pulumi.StringPtrOutput `pulumi:"expander"`
	// The scale-in threshold for GPU instance. Default is `0.5`.
	GpuUtilizationThreshold pulumi.StringPtrOutput `pulumi:"gpuUtilizationThreshold"`
	// Specify whether to allow the scale-in of nodes. Default is `true`.
	ScaleDownEnabled pulumi.BoolPtrOutput `pulumi:"scaleDownEnabled"`
	// The interval at which the cluster is reevaluated for scaling. Default is `30s`.
	ScanInterval pulumi.StringPtrOutput `pulumi:"scanInterval"`
	// The unneeded duration. Default is `10m`.
	UnneededDuration pulumi.StringPtrOutput `pulumi:"unneededDuration"`
	// The scale-in threshold. Default is `0.5`.
	UtilizationThreshold pulumi.StringPtrOutput `pulumi:"utilizationThreshold"`
}

// NewAutoscalingConfig registers a new resource with the given unique name, arguments, and options.
func NewAutoscalingConfig(ctx *pulumi.Context,
	name string, args *AutoscalingConfigArgs, opts ...pulumi.ResourceOption) (*AutoscalingConfig, error) {
	if args == nil {
		args = &AutoscalingConfigArgs{}
	}

	var resource AutoscalingConfig
	err := ctx.RegisterResource("alicloud:cs/autoscalingConfig:AutoscalingConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoscalingConfig gets an existing AutoscalingConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoscalingConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoscalingConfigState, opts ...pulumi.ResourceOption) (*AutoscalingConfig, error) {
	var resource AutoscalingConfig
	err := ctx.ReadResource("alicloud:cs/autoscalingConfig:AutoscalingConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoscalingConfig resources.
type autoscalingConfigState struct {
	// The id of kubernetes cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The cool down duration. Default is `10m`. If the delay (cooldown) value is set too long, there could be complaints that the Horizontal Pod Autoscaler is not responsive to workload changes. However, if the delay value is set too short, the scale of the replicas set may keep thrashing as usual.
	CoolDownDuration *string `pulumi:"coolDownDuration"`
	// The policy for selecting which node pool to scale. Valid values: `least-waste`, `random`, `priority`. For more information on these policies, see [Configure auto scaling](https://www.alibabacloud.com/help/en/container-service-for-kubernetes/latest/auto-scaling-of-nodes#section-3bg-2ko-inl)
	Expander *string `pulumi:"expander"`
	// The scale-in threshold for GPU instance. Default is `0.5`.
	GpuUtilizationThreshold *string `pulumi:"gpuUtilizationThreshold"`
	// Specify whether to allow the scale-in of nodes. Default is `true`.
	ScaleDownEnabled *bool `pulumi:"scaleDownEnabled"`
	// The interval at which the cluster is reevaluated for scaling. Default is `30s`.
	ScanInterval *string `pulumi:"scanInterval"`
	// The unneeded duration. Default is `10m`.
	UnneededDuration *string `pulumi:"unneededDuration"`
	// The scale-in threshold. Default is `0.5`.
	UtilizationThreshold *string `pulumi:"utilizationThreshold"`
}

type AutoscalingConfigState struct {
	// The id of kubernetes cluster.
	ClusterId pulumi.StringPtrInput
	// The cool down duration. Default is `10m`. If the delay (cooldown) value is set too long, there could be complaints that the Horizontal Pod Autoscaler is not responsive to workload changes. However, if the delay value is set too short, the scale of the replicas set may keep thrashing as usual.
	CoolDownDuration pulumi.StringPtrInput
	// The policy for selecting which node pool to scale. Valid values: `least-waste`, `random`, `priority`. For more information on these policies, see [Configure auto scaling](https://www.alibabacloud.com/help/en/container-service-for-kubernetes/latest/auto-scaling-of-nodes#section-3bg-2ko-inl)
	Expander pulumi.StringPtrInput
	// The scale-in threshold for GPU instance. Default is `0.5`.
	GpuUtilizationThreshold pulumi.StringPtrInput
	// Specify whether to allow the scale-in of nodes. Default is `true`.
	ScaleDownEnabled pulumi.BoolPtrInput
	// The interval at which the cluster is reevaluated for scaling. Default is `30s`.
	ScanInterval pulumi.StringPtrInput
	// The unneeded duration. Default is `10m`.
	UnneededDuration pulumi.StringPtrInput
	// The scale-in threshold. Default is `0.5`.
	UtilizationThreshold pulumi.StringPtrInput
}

func (AutoscalingConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalingConfigState)(nil)).Elem()
}

type autoscalingConfigArgs struct {
	// The id of kubernetes cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The cool down duration. Default is `10m`. If the delay (cooldown) value is set too long, there could be complaints that the Horizontal Pod Autoscaler is not responsive to workload changes. However, if the delay value is set too short, the scale of the replicas set may keep thrashing as usual.
	CoolDownDuration *string `pulumi:"coolDownDuration"`
	// The policy for selecting which node pool to scale. Valid values: `least-waste`, `random`, `priority`. For more information on these policies, see [Configure auto scaling](https://www.alibabacloud.com/help/en/container-service-for-kubernetes/latest/auto-scaling-of-nodes#section-3bg-2ko-inl)
	Expander *string `pulumi:"expander"`
	// The scale-in threshold for GPU instance. Default is `0.5`.
	GpuUtilizationThreshold *string `pulumi:"gpuUtilizationThreshold"`
	// Specify whether to allow the scale-in of nodes. Default is `true`.
	ScaleDownEnabled *bool `pulumi:"scaleDownEnabled"`
	// The interval at which the cluster is reevaluated for scaling. Default is `30s`.
	ScanInterval *string `pulumi:"scanInterval"`
	// The unneeded duration. Default is `10m`.
	UnneededDuration *string `pulumi:"unneededDuration"`
	// The scale-in threshold. Default is `0.5`.
	UtilizationThreshold *string `pulumi:"utilizationThreshold"`
}

// The set of arguments for constructing a AutoscalingConfig resource.
type AutoscalingConfigArgs struct {
	// The id of kubernetes cluster.
	ClusterId pulumi.StringPtrInput
	// The cool down duration. Default is `10m`. If the delay (cooldown) value is set too long, there could be complaints that the Horizontal Pod Autoscaler is not responsive to workload changes. However, if the delay value is set too short, the scale of the replicas set may keep thrashing as usual.
	CoolDownDuration pulumi.StringPtrInput
	// The policy for selecting which node pool to scale. Valid values: `least-waste`, `random`, `priority`. For more information on these policies, see [Configure auto scaling](https://www.alibabacloud.com/help/en/container-service-for-kubernetes/latest/auto-scaling-of-nodes#section-3bg-2ko-inl)
	Expander pulumi.StringPtrInput
	// The scale-in threshold for GPU instance. Default is `0.5`.
	GpuUtilizationThreshold pulumi.StringPtrInput
	// Specify whether to allow the scale-in of nodes. Default is `true`.
	ScaleDownEnabled pulumi.BoolPtrInput
	// The interval at which the cluster is reevaluated for scaling. Default is `30s`.
	ScanInterval pulumi.StringPtrInput
	// The unneeded duration. Default is `10m`.
	UnneededDuration pulumi.StringPtrInput
	// The scale-in threshold. Default is `0.5`.
	UtilizationThreshold pulumi.StringPtrInput
}

func (AutoscalingConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalingConfigArgs)(nil)).Elem()
}

type AutoscalingConfigInput interface {
	pulumi.Input

	ToAutoscalingConfigOutput() AutoscalingConfigOutput
	ToAutoscalingConfigOutputWithContext(ctx context.Context) AutoscalingConfigOutput
}

func (*AutoscalingConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscalingConfig)(nil)).Elem()
}

func (i *AutoscalingConfig) ToAutoscalingConfigOutput() AutoscalingConfigOutput {
	return i.ToAutoscalingConfigOutputWithContext(context.Background())
}

func (i *AutoscalingConfig) ToAutoscalingConfigOutputWithContext(ctx context.Context) AutoscalingConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalingConfigOutput)
}

// AutoscalingConfigArrayInput is an input type that accepts AutoscalingConfigArray and AutoscalingConfigArrayOutput values.
// You can construct a concrete instance of `AutoscalingConfigArrayInput` via:
//
//	AutoscalingConfigArray{ AutoscalingConfigArgs{...} }
type AutoscalingConfigArrayInput interface {
	pulumi.Input

	ToAutoscalingConfigArrayOutput() AutoscalingConfigArrayOutput
	ToAutoscalingConfigArrayOutputWithContext(context.Context) AutoscalingConfigArrayOutput
}

type AutoscalingConfigArray []AutoscalingConfigInput

func (AutoscalingConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoscalingConfig)(nil)).Elem()
}

func (i AutoscalingConfigArray) ToAutoscalingConfigArrayOutput() AutoscalingConfigArrayOutput {
	return i.ToAutoscalingConfigArrayOutputWithContext(context.Background())
}

func (i AutoscalingConfigArray) ToAutoscalingConfigArrayOutputWithContext(ctx context.Context) AutoscalingConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalingConfigArrayOutput)
}

// AutoscalingConfigMapInput is an input type that accepts AutoscalingConfigMap and AutoscalingConfigMapOutput values.
// You can construct a concrete instance of `AutoscalingConfigMapInput` via:
//
//	AutoscalingConfigMap{ "key": AutoscalingConfigArgs{...} }
type AutoscalingConfigMapInput interface {
	pulumi.Input

	ToAutoscalingConfigMapOutput() AutoscalingConfigMapOutput
	ToAutoscalingConfigMapOutputWithContext(context.Context) AutoscalingConfigMapOutput
}

type AutoscalingConfigMap map[string]AutoscalingConfigInput

func (AutoscalingConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoscalingConfig)(nil)).Elem()
}

func (i AutoscalingConfigMap) ToAutoscalingConfigMapOutput() AutoscalingConfigMapOutput {
	return i.ToAutoscalingConfigMapOutputWithContext(context.Background())
}

func (i AutoscalingConfigMap) ToAutoscalingConfigMapOutputWithContext(ctx context.Context) AutoscalingConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalingConfigMapOutput)
}

type AutoscalingConfigOutput struct{ *pulumi.OutputState }

func (AutoscalingConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscalingConfig)(nil)).Elem()
}

func (o AutoscalingConfigOutput) ToAutoscalingConfigOutput() AutoscalingConfigOutput {
	return o
}

func (o AutoscalingConfigOutput) ToAutoscalingConfigOutputWithContext(ctx context.Context) AutoscalingConfigOutput {
	return o
}

// The id of kubernetes cluster.
func (o AutoscalingConfigOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoscalingConfig) pulumi.StringPtrOutput { return v.ClusterId }).(pulumi.StringPtrOutput)
}

// The cool down duration. Default is `10m`. If the delay (cooldown) value is set too long, there could be complaints that the Horizontal Pod Autoscaler is not responsive to workload changes. However, if the delay value is set too short, the scale of the replicas set may keep thrashing as usual.
func (o AutoscalingConfigOutput) CoolDownDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoscalingConfig) pulumi.StringPtrOutput { return v.CoolDownDuration }).(pulumi.StringPtrOutput)
}

// The policy for selecting which node pool to scale. Valid values: `least-waste`, `random`, `priority`. For more information on these policies, see [Configure auto scaling](https://www.alibabacloud.com/help/en/container-service-for-kubernetes/latest/auto-scaling-of-nodes#section-3bg-2ko-inl)
func (o AutoscalingConfigOutput) Expander() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoscalingConfig) pulumi.StringPtrOutput { return v.Expander }).(pulumi.StringPtrOutput)
}

// The scale-in threshold for GPU instance. Default is `0.5`.
func (o AutoscalingConfigOutput) GpuUtilizationThreshold() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoscalingConfig) pulumi.StringPtrOutput { return v.GpuUtilizationThreshold }).(pulumi.StringPtrOutput)
}

// Specify whether to allow the scale-in of nodes. Default is `true`.
func (o AutoscalingConfigOutput) ScaleDownEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoscalingConfig) pulumi.BoolPtrOutput { return v.ScaleDownEnabled }).(pulumi.BoolPtrOutput)
}

// The interval at which the cluster is reevaluated for scaling. Default is `30s`.
func (o AutoscalingConfigOutput) ScanInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoscalingConfig) pulumi.StringPtrOutput { return v.ScanInterval }).(pulumi.StringPtrOutput)
}

// The unneeded duration. Default is `10m`.
func (o AutoscalingConfigOutput) UnneededDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoscalingConfig) pulumi.StringPtrOutput { return v.UnneededDuration }).(pulumi.StringPtrOutput)
}

// The scale-in threshold. Default is `0.5`.
func (o AutoscalingConfigOutput) UtilizationThreshold() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoscalingConfig) pulumi.StringPtrOutput { return v.UtilizationThreshold }).(pulumi.StringPtrOutput)
}

type AutoscalingConfigArrayOutput struct{ *pulumi.OutputState }

func (AutoscalingConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoscalingConfig)(nil)).Elem()
}

func (o AutoscalingConfigArrayOutput) ToAutoscalingConfigArrayOutput() AutoscalingConfigArrayOutput {
	return o
}

func (o AutoscalingConfigArrayOutput) ToAutoscalingConfigArrayOutputWithContext(ctx context.Context) AutoscalingConfigArrayOutput {
	return o
}

func (o AutoscalingConfigArrayOutput) Index(i pulumi.IntInput) AutoscalingConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AutoscalingConfig {
		return vs[0].([]*AutoscalingConfig)[vs[1].(int)]
	}).(AutoscalingConfigOutput)
}

type AutoscalingConfigMapOutput struct{ *pulumi.OutputState }

func (AutoscalingConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoscalingConfig)(nil)).Elem()
}

func (o AutoscalingConfigMapOutput) ToAutoscalingConfigMapOutput() AutoscalingConfigMapOutput {
	return o
}

func (o AutoscalingConfigMapOutput) ToAutoscalingConfigMapOutputWithContext(ctx context.Context) AutoscalingConfigMapOutput {
	return o
}

func (o AutoscalingConfigMapOutput) MapIndex(k pulumi.StringInput) AutoscalingConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AutoscalingConfig {
		return vs[0].(map[string]*AutoscalingConfig)[vs[1].(string)]
	}).(AutoscalingConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoscalingConfigInput)(nil)).Elem(), &AutoscalingConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoscalingConfigArrayInput)(nil)).Elem(), AutoscalingConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoscalingConfigMapInput)(nil)).Elem(), AutoscalingConfigMap{})
	pulumi.RegisterOutputType(AutoscalingConfigOutput{})
	pulumi.RegisterOutputType(AutoscalingConfigArrayOutput{})
	pulumi.RegisterOutputType(AutoscalingConfigMapOutput{})
}
