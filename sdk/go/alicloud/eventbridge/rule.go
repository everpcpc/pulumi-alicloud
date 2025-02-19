// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventbridge

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Event Bridge Rule resource.
//
// For information about Event Bridge Rule and how to use it, see [What is Rule](https://help.aliyun.com/document_detail/167854.html).
//
// > **NOTE:** Available in v1.129.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eventbridge"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleEventBus, err := eventbridge.NewEventBus(ctx, "exampleEventBus", &eventbridge.EventBusArgs{
//				EventBusName: pulumi.String("example_value"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = eventbridge.NewRule(ctx, "exampleRule", &eventbridge.RuleArgs{
//				EventBusName:  exampleEventBus.ID(),
//				RuleName:      pulumi.Any(_var.Name),
//				Description:   pulumi.String("test"),
//				FilterPattern: pulumi.String("{\"source\":[\"crmabc.newsletter\"],\"type\":[\"UserSignUp\", \"UserLogin\"]}"),
//				Targets: eventbridge.RuleTargetArray{
//					&eventbridge.RuleTargetArgs{
//						TargetId: pulumi.String("tf-test"),
//						Endpoint: pulumi.String("acs:mns:cn-hangzhou:118938335****:queues/tf-test"),
//						Type:     pulumi.String("acs.mns.queue"),
//						ParamLists: eventbridge.RuleTargetParamListArray{
//							&eventbridge.RuleTargetParamListArgs{
//								ResourceKey: pulumi.String("queue"),
//								Form:        pulumi.String("CONSTANT"),
//								Value:       pulumi.String("tf-testaccEbRule"),
//							},
//							&eventbridge.RuleTargetParamListArgs{
//								ResourceKey: pulumi.String("Body"),
//								Form:        pulumi.String("ORIGINAL"),
//							},
//						},
//					},
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
// Event Bridge Rule can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:eventbridge/rule:Rule example <event_bus_name>:<rule_name>
//
// ```
type Rule struct {
	pulumi.CustomResourceState

	// The description of rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of event bus.
	EventBusName pulumi.StringOutput `pulumi:"eventBusName"`
	// The pattern to match interested events. Event mode, JSON format. The value description is as follows: `stringEqual` mode. `stringExpression` mode. Each field has up to 5 expressions (map structure).
	FilterPattern pulumi.StringOutput `pulumi:"filterPattern"`
	// The name of rule.
	RuleName pulumi.StringOutput `pulumi:"ruleName"`
	// Rule status, either Enable or Disable. Valid values: `DISABLE`, `ENABLE`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The target of rule.
	Targets RuleTargetArrayOutput `pulumi:"targets"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventBusName == nil {
		return nil, errors.New("invalid value for required argument 'EventBusName'")
	}
	if args.FilterPattern == nil {
		return nil, errors.New("invalid value for required argument 'FilterPattern'")
	}
	if args.RuleName == nil {
		return nil, errors.New("invalid value for required argument 'RuleName'")
	}
	if args.Targets == nil {
		return nil, errors.New("invalid value for required argument 'Targets'")
	}
	var resource Rule
	err := ctx.RegisterResource("alicloud:eventbridge/rule:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("alicloud:eventbridge/rule:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	// The description of rule.
	Description *string `pulumi:"description"`
	// The name of event bus.
	EventBusName *string `pulumi:"eventBusName"`
	// The pattern to match interested events. Event mode, JSON format. The value description is as follows: `stringEqual` mode. `stringExpression` mode. Each field has up to 5 expressions (map structure).
	FilterPattern *string `pulumi:"filterPattern"`
	// The name of rule.
	RuleName *string `pulumi:"ruleName"`
	// Rule status, either Enable or Disable. Valid values: `DISABLE`, `ENABLE`.
	Status *string `pulumi:"status"`
	// The target of rule.
	Targets []RuleTarget `pulumi:"targets"`
}

type RuleState struct {
	// The description of rule.
	Description pulumi.StringPtrInput
	// The name of event bus.
	EventBusName pulumi.StringPtrInput
	// The pattern to match interested events. Event mode, JSON format. The value description is as follows: `stringEqual` mode. `stringExpression` mode. Each field has up to 5 expressions (map structure).
	FilterPattern pulumi.StringPtrInput
	// The name of rule.
	RuleName pulumi.StringPtrInput
	// Rule status, either Enable or Disable. Valid values: `DISABLE`, `ENABLE`.
	Status pulumi.StringPtrInput
	// The target of rule.
	Targets RuleTargetArrayInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// The description of rule.
	Description *string `pulumi:"description"`
	// The name of event bus.
	EventBusName string `pulumi:"eventBusName"`
	// The pattern to match interested events. Event mode, JSON format. The value description is as follows: `stringEqual` mode. `stringExpression` mode. Each field has up to 5 expressions (map structure).
	FilterPattern string `pulumi:"filterPattern"`
	// The name of rule.
	RuleName string `pulumi:"ruleName"`
	// Rule status, either Enable or Disable. Valid values: `DISABLE`, `ENABLE`.
	Status *string `pulumi:"status"`
	// The target of rule.
	Targets []RuleTarget `pulumi:"targets"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// The description of rule.
	Description pulumi.StringPtrInput
	// The name of event bus.
	EventBusName pulumi.StringInput
	// The pattern to match interested events. Event mode, JSON format. The value description is as follows: `stringEqual` mode. `stringExpression` mode. Each field has up to 5 expressions (map structure).
	FilterPattern pulumi.StringInput
	// The name of rule.
	RuleName pulumi.StringInput
	// Rule status, either Enable or Disable. Valid values: `DISABLE`, `ENABLE`.
	Status pulumi.StringPtrInput
	// The target of rule.
	Targets RuleTargetArrayInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}

type RuleInput interface {
	pulumi.Input

	ToRuleOutput() RuleOutput
	ToRuleOutputWithContext(ctx context.Context) RuleOutput
}

func (*Rule) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (i *Rule) ToRuleOutput() RuleOutput {
	return i.ToRuleOutputWithContext(context.Background())
}

func (i *Rule) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleOutput)
}

// RuleArrayInput is an input type that accepts RuleArray and RuleArrayOutput values.
// You can construct a concrete instance of `RuleArrayInput` via:
//
//	RuleArray{ RuleArgs{...} }
type RuleArrayInput interface {
	pulumi.Input

	ToRuleArrayOutput() RuleArrayOutput
	ToRuleArrayOutputWithContext(context.Context) RuleArrayOutput
}

type RuleArray []RuleInput

func (RuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rule)(nil)).Elem()
}

func (i RuleArray) ToRuleArrayOutput() RuleArrayOutput {
	return i.ToRuleArrayOutputWithContext(context.Background())
}

func (i RuleArray) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleArrayOutput)
}

// RuleMapInput is an input type that accepts RuleMap and RuleMapOutput values.
// You can construct a concrete instance of `RuleMapInput` via:
//
//	RuleMap{ "key": RuleArgs{...} }
type RuleMapInput interface {
	pulumi.Input

	ToRuleMapOutput() RuleMapOutput
	ToRuleMapOutputWithContext(context.Context) RuleMapOutput
}

type RuleMap map[string]RuleInput

func (RuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rule)(nil)).Elem()
}

func (i RuleMap) ToRuleMapOutput() RuleMapOutput {
	return i.ToRuleMapOutputWithContext(context.Background())
}

func (i RuleMap) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleMapOutput)
}

type RuleOutput struct{ *pulumi.OutputState }

func (RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (o RuleOutput) ToRuleOutput() RuleOutput {
	return o
}

func (o RuleOutput) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return o
}

// The description of rule.
func (o RuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of event bus.
func (o RuleOutput) EventBusName() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.EventBusName }).(pulumi.StringOutput)
}

// The pattern to match interested events. Event mode, JSON format. The value description is as follows: `stringEqual` mode. `stringExpression` mode. Each field has up to 5 expressions (map structure).
func (o RuleOutput) FilterPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.FilterPattern }).(pulumi.StringOutput)
}

// The name of rule.
func (o RuleOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.RuleName }).(pulumi.StringOutput)
}

// Rule status, either Enable or Disable. Valid values: `DISABLE`, `ENABLE`.
func (o RuleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The target of rule.
func (o RuleOutput) Targets() RuleTargetArrayOutput {
	return o.ApplyT(func(v *Rule) RuleTargetArrayOutput { return v.Targets }).(RuleTargetArrayOutput)
}

type RuleArrayOutput struct{ *pulumi.OutputState }

func (RuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rule)(nil)).Elem()
}

func (o RuleArrayOutput) ToRuleArrayOutput() RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) Index(i pulumi.IntInput) RuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Rule {
		return vs[0].([]*Rule)[vs[1].(int)]
	}).(RuleOutput)
}

type RuleMapOutput struct{ *pulumi.OutputState }

func (RuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rule)(nil)).Elem()
}

func (o RuleMapOutput) ToRuleMapOutput() RuleMapOutput {
	return o
}

func (o RuleMapOutput) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return o
}

func (o RuleMapOutput) MapIndex(k pulumi.StringInput) RuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Rule {
		return vs[0].(map[string]*Rule)[vs[1].(string)]
	}).(RuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleInput)(nil)).Elem(), &Rule{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleArrayInput)(nil)).Elem(), RuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleMapInput)(nil)).Elem(), RuleMap{})
	pulumi.RegisterOutputType(RuleOutput{})
	pulumi.RegisterOutputType(RuleArrayOutput{})
	pulumi.RegisterOutputType(RuleMapOutput{})
}
