// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Global Accelerator (GA) Forwarding Rule resource.
//
// For information about Global Accelerator (GA) Forwarding Rule and how to use it, see [What is Forwarding Rule](https://www.alibabacloud.com/help/zh/doc-detail/205815.htm).
//
// > **NOTE:** Available in v1.120.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ga"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAccelerator, err := ga.NewAccelerator(ctx, "exampleAccelerator", &ga.AcceleratorArgs{
// 			Duration:      pulumi.Int(1),
// 			AutoUseCoupon: pulumi.Bool(true),
// 			Spec:          pulumi.String("1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		deBandwidthPackage, err := ga.NewBandwidthPackage(ctx, "deBandwidthPackage", &ga.BandwidthPackageArgs{
// 			Bandwidth:     pulumi.Int(100),
// 			Type:          pulumi.String("Basic"),
// 			BandwidthType: pulumi.String("Basic"),
// 			PaymentType:   pulumi.String("PayAsYouGo"),
// 			BillingType:   pulumi.String("PayBy95"),
// 			Ratio:         pulumi.Int(30),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		deBandwidthPackageAttachment, err := ga.NewBandwidthPackageAttachment(ctx, "deBandwidthPackageAttachment", &ga.BandwidthPackageAttachmentArgs{
// 			AcceleratorId:      exampleAccelerator.ID(),
// 			BandwidthPackageId: deBandwidthPackage.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleListener, err := ga.NewListener(ctx, "exampleListener", &ga.ListenerArgs{
// 			AcceleratorId: exampleAccelerator.ID(),
// 			PortRanges: ga.ListenerPortRangeArray{
// 				&ga.ListenerPortRangeArgs{
// 					FromPort: pulumi.Int(70),
// 					ToPort:   pulumi.Int(70),
// 				},
// 			},
// 			Protocol: pulumi.String("HTTP"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			deBandwidthPackageAttachment,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		exampleEipAddress, err := ecs.NewEipAddress(ctx, "exampleEipAddress", &ecs.EipAddressArgs{
// 			Bandwidth:          pulumi.String("10"),
// 			InternetChargeType: pulumi.String("PayByBandwidth"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEndpointGroup, err := ga.NewEndpointGroup(ctx, "exampleEndpointGroup", &ga.EndpointGroupArgs{
// 			AcceleratorId: exampleAccelerator.ID(),
// 			EndpointConfigurations: ga.EndpointGroupEndpointConfigurationArray{
// 				&ga.EndpointGroupEndpointConfigurationArgs{
// 					Endpoint: exampleEipAddress.IpAddress,
// 					Type:     pulumi.String("PublicIp"),
// 					Weight:   pulumi.Int(20),
// 				},
// 			},
// 			EndpointGroupRegion: pulumi.String("cn-hangzhou"),
// 			ListenerId:          exampleListener.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ga.NewForwardingRule(ctx, "exampleForwardingRule", &ga.ForwardingRuleArgs{
// 			AcceleratorId: exampleAccelerator.ID(),
// 			ListenerId:    exampleListener.ID(),
// 			RuleConditions: ga.ForwardingRuleRuleConditionArray{
// 				&ga.ForwardingRuleRuleConditionArgs{
// 					RuleConditionType: pulumi.String("Path"),
// 					PathConfig: &ga.ForwardingRuleRuleConditionPathConfigArgs{
// 						Values: pulumi.StringArray{
// 							pulumi.String("/test"),
// 						},
// 					},
// 				},
// 			},
// 			RuleActions: ga.ForwardingRuleRuleActionArray{
// 				&ga.ForwardingRuleRuleActionArgs{
// 					Order:          pulumi.Int(30),
// 					RuleActionType: pulumi.String("ForwardGroup"),
// 					ForwardGroupConfig: &ga.ForwardingRuleRuleActionForwardGroupConfigArgs{
// 						ServerGroupTuples: ga.ForwardingRuleRuleActionForwardGroupConfigServerGroupTupleArray{
// 							&ga.ForwardingRuleRuleActionForwardGroupConfigServerGroupTupleArgs{
// 								EndpointGroupId: exampleEndpointGroup.ID(),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Ga Forwarding Rule can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:ga/forwardingRule:ForwardingRule example <id>
// ```
type ForwardingRule struct {
	pulumi.CustomResourceState

	// The ID of the Global Accelerator instance.
	AcceleratorId pulumi.StringOutput `pulumi:"acceleratorId"`
	// Forwarding Policy ID.
	ForwardingRuleId pulumi.StringOutput `pulumi:"forwardingRuleId"`
	// Forwarding policy name. The length of the name is 2-128 English or Chinese characters. It must start with uppercase and lowercase letters or Chinese characters. It can contain numbers, half width period (.), underscores (_) And dash (-).
	ForwardingRuleName pulumi.StringPtrOutput `pulumi:"forwardingRuleName"`
	// Forwarding Policy Status.
	ForwardingRuleStatus pulumi.StringOutput `pulumi:"forwardingRuleStatus"`
	// The ID of the listener.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// Forwarding policy priority.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Forward action.
	RuleActions ForwardingRuleRuleActionArrayOutput `pulumi:"ruleActions"`
	// Forwarding condition list.
	RuleConditions ForwardingRuleRuleConditionArrayOutput `pulumi:"ruleConditions"`
}

// NewForwardingRule registers a new resource with the given unique name, arguments, and options.
func NewForwardingRule(ctx *pulumi.Context,
	name string, args *ForwardingRuleArgs, opts ...pulumi.ResourceOption) (*ForwardingRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AcceleratorId == nil {
		return nil, errors.New("invalid value for required argument 'AcceleratorId'")
	}
	if args.ListenerId == nil {
		return nil, errors.New("invalid value for required argument 'ListenerId'")
	}
	if args.RuleActions == nil {
		return nil, errors.New("invalid value for required argument 'RuleActions'")
	}
	if args.RuleConditions == nil {
		return nil, errors.New("invalid value for required argument 'RuleConditions'")
	}
	var resource ForwardingRule
	err := ctx.RegisterResource("alicloud:ga/forwardingRule:ForwardingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetForwardingRule gets an existing ForwardingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetForwardingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ForwardingRuleState, opts ...pulumi.ResourceOption) (*ForwardingRule, error) {
	var resource ForwardingRule
	err := ctx.ReadResource("alicloud:ga/forwardingRule:ForwardingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ForwardingRule resources.
type forwardingRuleState struct {
	// The ID of the Global Accelerator instance.
	AcceleratorId *string `pulumi:"acceleratorId"`
	// Forwarding Policy ID.
	ForwardingRuleId *string `pulumi:"forwardingRuleId"`
	// Forwarding policy name. The length of the name is 2-128 English or Chinese characters. It must start with uppercase and lowercase letters or Chinese characters. It can contain numbers, half width period (.), underscores (_) And dash (-).
	ForwardingRuleName *string `pulumi:"forwardingRuleName"`
	// Forwarding Policy Status.
	ForwardingRuleStatus *string `pulumi:"forwardingRuleStatus"`
	// The ID of the listener.
	ListenerId *string `pulumi:"listenerId"`
	// Forwarding policy priority.
	Priority *int `pulumi:"priority"`
	// Forward action.
	RuleActions []ForwardingRuleRuleAction `pulumi:"ruleActions"`
	// Forwarding condition list.
	RuleConditions []ForwardingRuleRuleCondition `pulumi:"ruleConditions"`
}

type ForwardingRuleState struct {
	// The ID of the Global Accelerator instance.
	AcceleratorId pulumi.StringPtrInput
	// Forwarding Policy ID.
	ForwardingRuleId pulumi.StringPtrInput
	// Forwarding policy name. The length of the name is 2-128 English or Chinese characters. It must start with uppercase and lowercase letters or Chinese characters. It can contain numbers, half width period (.), underscores (_) And dash (-).
	ForwardingRuleName pulumi.StringPtrInput
	// Forwarding Policy Status.
	ForwardingRuleStatus pulumi.StringPtrInput
	// The ID of the listener.
	ListenerId pulumi.StringPtrInput
	// Forwarding policy priority.
	Priority pulumi.IntPtrInput
	// Forward action.
	RuleActions ForwardingRuleRuleActionArrayInput
	// Forwarding condition list.
	RuleConditions ForwardingRuleRuleConditionArrayInput
}

func (ForwardingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardingRuleState)(nil)).Elem()
}

type forwardingRuleArgs struct {
	// The ID of the Global Accelerator instance.
	AcceleratorId string `pulumi:"acceleratorId"`
	// Forwarding policy name. The length of the name is 2-128 English or Chinese characters. It must start with uppercase and lowercase letters or Chinese characters. It can contain numbers, half width period (.), underscores (_) And dash (-).
	ForwardingRuleName *string `pulumi:"forwardingRuleName"`
	// The ID of the listener.
	ListenerId string `pulumi:"listenerId"`
	// Forwarding policy priority.
	Priority *int `pulumi:"priority"`
	// Forward action.
	RuleActions []ForwardingRuleRuleAction `pulumi:"ruleActions"`
	// Forwarding condition list.
	RuleConditions []ForwardingRuleRuleCondition `pulumi:"ruleConditions"`
}

// The set of arguments for constructing a ForwardingRule resource.
type ForwardingRuleArgs struct {
	// The ID of the Global Accelerator instance.
	AcceleratorId pulumi.StringInput
	// Forwarding policy name. The length of the name is 2-128 English or Chinese characters. It must start with uppercase and lowercase letters or Chinese characters. It can contain numbers, half width period (.), underscores (_) And dash (-).
	ForwardingRuleName pulumi.StringPtrInput
	// The ID of the listener.
	ListenerId pulumi.StringInput
	// Forwarding policy priority.
	Priority pulumi.IntPtrInput
	// Forward action.
	RuleActions ForwardingRuleRuleActionArrayInput
	// Forwarding condition list.
	RuleConditions ForwardingRuleRuleConditionArrayInput
}

func (ForwardingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardingRuleArgs)(nil)).Elem()
}

type ForwardingRuleInput interface {
	pulumi.Input

	ToForwardingRuleOutput() ForwardingRuleOutput
	ToForwardingRuleOutputWithContext(ctx context.Context) ForwardingRuleOutput
}

func (*ForwardingRule) ElementType() reflect.Type {
	return reflect.TypeOf((*ForwardingRule)(nil))
}

func (i *ForwardingRule) ToForwardingRuleOutput() ForwardingRuleOutput {
	return i.ToForwardingRuleOutputWithContext(context.Background())
}

func (i *ForwardingRule) ToForwardingRuleOutputWithContext(ctx context.Context) ForwardingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardingRuleOutput)
}

func (i *ForwardingRule) ToForwardingRulePtrOutput() ForwardingRulePtrOutput {
	return i.ToForwardingRulePtrOutputWithContext(context.Background())
}

func (i *ForwardingRule) ToForwardingRulePtrOutputWithContext(ctx context.Context) ForwardingRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardingRulePtrOutput)
}

type ForwardingRulePtrInput interface {
	pulumi.Input

	ToForwardingRulePtrOutput() ForwardingRulePtrOutput
	ToForwardingRulePtrOutputWithContext(ctx context.Context) ForwardingRulePtrOutput
}

type forwardingRulePtrType ForwardingRuleArgs

func (*forwardingRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardingRule)(nil))
}

func (i *forwardingRulePtrType) ToForwardingRulePtrOutput() ForwardingRulePtrOutput {
	return i.ToForwardingRulePtrOutputWithContext(context.Background())
}

func (i *forwardingRulePtrType) ToForwardingRulePtrOutputWithContext(ctx context.Context) ForwardingRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardingRulePtrOutput)
}

// ForwardingRuleArrayInput is an input type that accepts ForwardingRuleArray and ForwardingRuleArrayOutput values.
// You can construct a concrete instance of `ForwardingRuleArrayInput` via:
//
//          ForwardingRuleArray{ ForwardingRuleArgs{...} }
type ForwardingRuleArrayInput interface {
	pulumi.Input

	ToForwardingRuleArrayOutput() ForwardingRuleArrayOutput
	ToForwardingRuleArrayOutputWithContext(context.Context) ForwardingRuleArrayOutput
}

type ForwardingRuleArray []ForwardingRuleInput

func (ForwardingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ForwardingRule)(nil)).Elem()
}

func (i ForwardingRuleArray) ToForwardingRuleArrayOutput() ForwardingRuleArrayOutput {
	return i.ToForwardingRuleArrayOutputWithContext(context.Background())
}

func (i ForwardingRuleArray) ToForwardingRuleArrayOutputWithContext(ctx context.Context) ForwardingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardingRuleArrayOutput)
}

// ForwardingRuleMapInput is an input type that accepts ForwardingRuleMap and ForwardingRuleMapOutput values.
// You can construct a concrete instance of `ForwardingRuleMapInput` via:
//
//          ForwardingRuleMap{ "key": ForwardingRuleArgs{...} }
type ForwardingRuleMapInput interface {
	pulumi.Input

	ToForwardingRuleMapOutput() ForwardingRuleMapOutput
	ToForwardingRuleMapOutputWithContext(context.Context) ForwardingRuleMapOutput
}

type ForwardingRuleMap map[string]ForwardingRuleInput

func (ForwardingRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ForwardingRule)(nil)).Elem()
}

func (i ForwardingRuleMap) ToForwardingRuleMapOutput() ForwardingRuleMapOutput {
	return i.ToForwardingRuleMapOutputWithContext(context.Background())
}

func (i ForwardingRuleMap) ToForwardingRuleMapOutputWithContext(ctx context.Context) ForwardingRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardingRuleMapOutput)
}

type ForwardingRuleOutput struct{ *pulumi.OutputState }

func (ForwardingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ForwardingRule)(nil))
}

func (o ForwardingRuleOutput) ToForwardingRuleOutput() ForwardingRuleOutput {
	return o
}

func (o ForwardingRuleOutput) ToForwardingRuleOutputWithContext(ctx context.Context) ForwardingRuleOutput {
	return o
}

func (o ForwardingRuleOutput) ToForwardingRulePtrOutput() ForwardingRulePtrOutput {
	return o.ToForwardingRulePtrOutputWithContext(context.Background())
}

func (o ForwardingRuleOutput) ToForwardingRulePtrOutputWithContext(ctx context.Context) ForwardingRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ForwardingRule) *ForwardingRule {
		return &v
	}).(ForwardingRulePtrOutput)
}

type ForwardingRulePtrOutput struct{ *pulumi.OutputState }

func (ForwardingRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardingRule)(nil))
}

func (o ForwardingRulePtrOutput) ToForwardingRulePtrOutput() ForwardingRulePtrOutput {
	return o
}

func (o ForwardingRulePtrOutput) ToForwardingRulePtrOutputWithContext(ctx context.Context) ForwardingRulePtrOutput {
	return o
}

func (o ForwardingRulePtrOutput) Elem() ForwardingRuleOutput {
	return o.ApplyT(func(v *ForwardingRule) ForwardingRule {
		if v != nil {
			return *v
		}
		var ret ForwardingRule
		return ret
	}).(ForwardingRuleOutput)
}

type ForwardingRuleArrayOutput struct{ *pulumi.OutputState }

func (ForwardingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ForwardingRule)(nil))
}

func (o ForwardingRuleArrayOutput) ToForwardingRuleArrayOutput() ForwardingRuleArrayOutput {
	return o
}

func (o ForwardingRuleArrayOutput) ToForwardingRuleArrayOutputWithContext(ctx context.Context) ForwardingRuleArrayOutput {
	return o
}

func (o ForwardingRuleArrayOutput) Index(i pulumi.IntInput) ForwardingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ForwardingRule {
		return vs[0].([]ForwardingRule)[vs[1].(int)]
	}).(ForwardingRuleOutput)
}

type ForwardingRuleMapOutput struct{ *pulumi.OutputState }

func (ForwardingRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ForwardingRule)(nil))
}

func (o ForwardingRuleMapOutput) ToForwardingRuleMapOutput() ForwardingRuleMapOutput {
	return o
}

func (o ForwardingRuleMapOutput) ToForwardingRuleMapOutputWithContext(ctx context.Context) ForwardingRuleMapOutput {
	return o
}

func (o ForwardingRuleMapOutput) MapIndex(k pulumi.StringInput) ForwardingRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ForwardingRule {
		return vs[0].(map[string]ForwardingRule)[vs[1].(string)]
	}).(ForwardingRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardingRuleInput)(nil)).Elem(), &ForwardingRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardingRulePtrInput)(nil)).Elem(), &ForwardingRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardingRuleArrayInput)(nil)).Elem(), ForwardingRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardingRuleMapInput)(nil)).Elem(), ForwardingRuleMap{})
	pulumi.RegisterOutputType(ForwardingRuleOutput{})
	pulumi.RegisterOutputType(ForwardingRulePtrOutput{})
	pulumi.RegisterOutputType(ForwardingRuleArrayOutput{})
	pulumi.RegisterOutputType(ForwardingRuleMapOutput{})
}
