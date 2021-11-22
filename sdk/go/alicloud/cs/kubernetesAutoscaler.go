// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KubernetesAutoscaler struct {
	pulumi.CustomResourceState

	// The id of kubernetes cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The coolDownDuration option of cluster-autoscaler.
	CoolDownDuration pulumi.StringOutput `pulumi:"coolDownDuration"`
	// The deferScaleInDuration option of cluster-autoscaler.
	DeferScaleInDuration pulumi.StringOutput `pulumi:"deferScaleInDuration"`
	// * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
	// * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
	// * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
	Nodepools KubernetesAutoscalerNodepoolArrayOutput `pulumi:"nodepools"`
	// Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
	UseEcsRamRoleToken pulumi.BoolPtrOutput `pulumi:"useEcsRamRoleToken"`
	// The utilization option of cluster-autoscaler.
	Utilization pulumi.StringOutput `pulumi:"utilization"`
}

// NewKubernetesAutoscaler registers a new resource with the given unique name, arguments, and options.
func NewKubernetesAutoscaler(ctx *pulumi.Context,
	name string, args *KubernetesAutoscalerArgs, opts ...pulumi.ResourceOption) (*KubernetesAutoscaler, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.CoolDownDuration == nil {
		return nil, errors.New("invalid value for required argument 'CoolDownDuration'")
	}
	if args.DeferScaleInDuration == nil {
		return nil, errors.New("invalid value for required argument 'DeferScaleInDuration'")
	}
	if args.Utilization == nil {
		return nil, errors.New("invalid value for required argument 'Utilization'")
	}
	var resource KubernetesAutoscaler
	err := ctx.RegisterResource("alicloud:cs/kubernetesAutoscaler:KubernetesAutoscaler", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesAutoscaler gets an existing KubernetesAutoscaler resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesAutoscaler(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesAutoscalerState, opts ...pulumi.ResourceOption) (*KubernetesAutoscaler, error) {
	var resource KubernetesAutoscaler
	err := ctx.ReadResource("alicloud:cs/kubernetesAutoscaler:KubernetesAutoscaler", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesAutoscaler resources.
type kubernetesAutoscalerState struct {
	// The id of kubernetes cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The coolDownDuration option of cluster-autoscaler.
	CoolDownDuration *string `pulumi:"coolDownDuration"`
	// The deferScaleInDuration option of cluster-autoscaler.
	DeferScaleInDuration *string `pulumi:"deferScaleInDuration"`
	// * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
	// * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
	// * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
	Nodepools []KubernetesAutoscalerNodepool `pulumi:"nodepools"`
	// Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
	UseEcsRamRoleToken *bool `pulumi:"useEcsRamRoleToken"`
	// The utilization option of cluster-autoscaler.
	Utilization *string `pulumi:"utilization"`
}

type KubernetesAutoscalerState struct {
	// The id of kubernetes cluster.
	ClusterId pulumi.StringPtrInput
	// The coolDownDuration option of cluster-autoscaler.
	CoolDownDuration pulumi.StringPtrInput
	// The deferScaleInDuration option of cluster-autoscaler.
	DeferScaleInDuration pulumi.StringPtrInput
	// * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
	// * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
	// * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
	Nodepools KubernetesAutoscalerNodepoolArrayInput
	// Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
	UseEcsRamRoleToken pulumi.BoolPtrInput
	// The utilization option of cluster-autoscaler.
	Utilization pulumi.StringPtrInput
}

func (KubernetesAutoscalerState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesAutoscalerState)(nil)).Elem()
}

type kubernetesAutoscalerArgs struct {
	// The id of kubernetes cluster.
	ClusterId string `pulumi:"clusterId"`
	// The coolDownDuration option of cluster-autoscaler.
	CoolDownDuration string `pulumi:"coolDownDuration"`
	// The deferScaleInDuration option of cluster-autoscaler.
	DeferScaleInDuration string `pulumi:"deferScaleInDuration"`
	// * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
	// * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
	// * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
	Nodepools []KubernetesAutoscalerNodepool `pulumi:"nodepools"`
	// Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
	UseEcsRamRoleToken *bool `pulumi:"useEcsRamRoleToken"`
	// The utilization option of cluster-autoscaler.
	Utilization string `pulumi:"utilization"`
}

// The set of arguments for constructing a KubernetesAutoscaler resource.
type KubernetesAutoscalerArgs struct {
	// The id of kubernetes cluster.
	ClusterId pulumi.StringInput
	// The coolDownDuration option of cluster-autoscaler.
	CoolDownDuration pulumi.StringInput
	// The deferScaleInDuration option of cluster-autoscaler.
	DeferScaleInDuration pulumi.StringInput
	// * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
	// * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
	// * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
	Nodepools KubernetesAutoscalerNodepoolArrayInput
	// Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
	UseEcsRamRoleToken pulumi.BoolPtrInput
	// The utilization option of cluster-autoscaler.
	Utilization pulumi.StringInput
}

func (KubernetesAutoscalerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesAutoscalerArgs)(nil)).Elem()
}

type KubernetesAutoscalerInput interface {
	pulumi.Input

	ToKubernetesAutoscalerOutput() KubernetesAutoscalerOutput
	ToKubernetesAutoscalerOutputWithContext(ctx context.Context) KubernetesAutoscalerOutput
}

func (*KubernetesAutoscaler) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesAutoscaler)(nil))
}

func (i *KubernetesAutoscaler) ToKubernetesAutoscalerOutput() KubernetesAutoscalerOutput {
	return i.ToKubernetesAutoscalerOutputWithContext(context.Background())
}

func (i *KubernetesAutoscaler) ToKubernetesAutoscalerOutputWithContext(ctx context.Context) KubernetesAutoscalerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesAutoscalerOutput)
}

func (i *KubernetesAutoscaler) ToKubernetesAutoscalerPtrOutput() KubernetesAutoscalerPtrOutput {
	return i.ToKubernetesAutoscalerPtrOutputWithContext(context.Background())
}

func (i *KubernetesAutoscaler) ToKubernetesAutoscalerPtrOutputWithContext(ctx context.Context) KubernetesAutoscalerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesAutoscalerPtrOutput)
}

type KubernetesAutoscalerPtrInput interface {
	pulumi.Input

	ToKubernetesAutoscalerPtrOutput() KubernetesAutoscalerPtrOutput
	ToKubernetesAutoscalerPtrOutputWithContext(ctx context.Context) KubernetesAutoscalerPtrOutput
}

type kubernetesAutoscalerPtrType KubernetesAutoscalerArgs

func (*kubernetesAutoscalerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesAutoscaler)(nil))
}

func (i *kubernetesAutoscalerPtrType) ToKubernetesAutoscalerPtrOutput() KubernetesAutoscalerPtrOutput {
	return i.ToKubernetesAutoscalerPtrOutputWithContext(context.Background())
}

func (i *kubernetesAutoscalerPtrType) ToKubernetesAutoscalerPtrOutputWithContext(ctx context.Context) KubernetesAutoscalerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesAutoscalerPtrOutput)
}

// KubernetesAutoscalerArrayInput is an input type that accepts KubernetesAutoscalerArray and KubernetesAutoscalerArrayOutput values.
// You can construct a concrete instance of `KubernetesAutoscalerArrayInput` via:
//
//          KubernetesAutoscalerArray{ KubernetesAutoscalerArgs{...} }
type KubernetesAutoscalerArrayInput interface {
	pulumi.Input

	ToKubernetesAutoscalerArrayOutput() KubernetesAutoscalerArrayOutput
	ToKubernetesAutoscalerArrayOutputWithContext(context.Context) KubernetesAutoscalerArrayOutput
}

type KubernetesAutoscalerArray []KubernetesAutoscalerInput

func (KubernetesAutoscalerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubernetesAutoscaler)(nil)).Elem()
}

func (i KubernetesAutoscalerArray) ToKubernetesAutoscalerArrayOutput() KubernetesAutoscalerArrayOutput {
	return i.ToKubernetesAutoscalerArrayOutputWithContext(context.Background())
}

func (i KubernetesAutoscalerArray) ToKubernetesAutoscalerArrayOutputWithContext(ctx context.Context) KubernetesAutoscalerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesAutoscalerArrayOutput)
}

// KubernetesAutoscalerMapInput is an input type that accepts KubernetesAutoscalerMap and KubernetesAutoscalerMapOutput values.
// You can construct a concrete instance of `KubernetesAutoscalerMapInput` via:
//
//          KubernetesAutoscalerMap{ "key": KubernetesAutoscalerArgs{...} }
type KubernetesAutoscalerMapInput interface {
	pulumi.Input

	ToKubernetesAutoscalerMapOutput() KubernetesAutoscalerMapOutput
	ToKubernetesAutoscalerMapOutputWithContext(context.Context) KubernetesAutoscalerMapOutput
}

type KubernetesAutoscalerMap map[string]KubernetesAutoscalerInput

func (KubernetesAutoscalerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubernetesAutoscaler)(nil)).Elem()
}

func (i KubernetesAutoscalerMap) ToKubernetesAutoscalerMapOutput() KubernetesAutoscalerMapOutput {
	return i.ToKubernetesAutoscalerMapOutputWithContext(context.Background())
}

func (i KubernetesAutoscalerMap) ToKubernetesAutoscalerMapOutputWithContext(ctx context.Context) KubernetesAutoscalerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesAutoscalerMapOutput)
}

type KubernetesAutoscalerOutput struct{ *pulumi.OutputState }

func (KubernetesAutoscalerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesAutoscaler)(nil))
}

func (o KubernetesAutoscalerOutput) ToKubernetesAutoscalerOutput() KubernetesAutoscalerOutput {
	return o
}

func (o KubernetesAutoscalerOutput) ToKubernetesAutoscalerOutputWithContext(ctx context.Context) KubernetesAutoscalerOutput {
	return o
}

func (o KubernetesAutoscalerOutput) ToKubernetesAutoscalerPtrOutput() KubernetesAutoscalerPtrOutput {
	return o.ToKubernetesAutoscalerPtrOutputWithContext(context.Background())
}

func (o KubernetesAutoscalerOutput) ToKubernetesAutoscalerPtrOutputWithContext(ctx context.Context) KubernetesAutoscalerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesAutoscaler) *KubernetesAutoscaler {
		return &v
	}).(KubernetesAutoscalerPtrOutput)
}

type KubernetesAutoscalerPtrOutput struct{ *pulumi.OutputState }

func (KubernetesAutoscalerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesAutoscaler)(nil))
}

func (o KubernetesAutoscalerPtrOutput) ToKubernetesAutoscalerPtrOutput() KubernetesAutoscalerPtrOutput {
	return o
}

func (o KubernetesAutoscalerPtrOutput) ToKubernetesAutoscalerPtrOutputWithContext(ctx context.Context) KubernetesAutoscalerPtrOutput {
	return o
}

func (o KubernetesAutoscalerPtrOutput) Elem() KubernetesAutoscalerOutput {
	return o.ApplyT(func(v *KubernetesAutoscaler) KubernetesAutoscaler {
		if v != nil {
			return *v
		}
		var ret KubernetesAutoscaler
		return ret
	}).(KubernetesAutoscalerOutput)
}

type KubernetesAutoscalerArrayOutput struct{ *pulumi.OutputState }

func (KubernetesAutoscalerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesAutoscaler)(nil))
}

func (o KubernetesAutoscalerArrayOutput) ToKubernetesAutoscalerArrayOutput() KubernetesAutoscalerArrayOutput {
	return o
}

func (o KubernetesAutoscalerArrayOutput) ToKubernetesAutoscalerArrayOutputWithContext(ctx context.Context) KubernetesAutoscalerArrayOutput {
	return o
}

func (o KubernetesAutoscalerArrayOutput) Index(i pulumi.IntInput) KubernetesAutoscalerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KubernetesAutoscaler {
		return vs[0].([]KubernetesAutoscaler)[vs[1].(int)]
	}).(KubernetesAutoscalerOutput)
}

type KubernetesAutoscalerMapOutput struct{ *pulumi.OutputState }

func (KubernetesAutoscalerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KubernetesAutoscaler)(nil))
}

func (o KubernetesAutoscalerMapOutput) ToKubernetesAutoscalerMapOutput() KubernetesAutoscalerMapOutput {
	return o
}

func (o KubernetesAutoscalerMapOutput) ToKubernetesAutoscalerMapOutputWithContext(ctx context.Context) KubernetesAutoscalerMapOutput {
	return o
}

func (o KubernetesAutoscalerMapOutput) MapIndex(k pulumi.StringInput) KubernetesAutoscalerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KubernetesAutoscaler {
		return vs[0].(map[string]KubernetesAutoscaler)[vs[1].(string)]
	}).(KubernetesAutoscalerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesAutoscalerInput)(nil)).Elem(), &KubernetesAutoscaler{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesAutoscalerPtrInput)(nil)).Elem(), &KubernetesAutoscaler{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesAutoscalerArrayInput)(nil)).Elem(), KubernetesAutoscalerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesAutoscalerMapInput)(nil)).Elem(), KubernetesAutoscalerMap{})
	pulumi.RegisterOutputType(KubernetesAutoscalerOutput{})
	pulumi.RegisterOutputType(KubernetesAutoscalerPtrOutput{})
	pulumi.RegisterOutputType(KubernetesAutoscalerArrayOutput{})
	pulumi.RegisterOutputType(KubernetesAutoscalerMapOutput{})
}
