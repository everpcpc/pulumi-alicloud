// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Application Load Balancer (ALB) Rule resource.
//
// For information about Application Load Balancer (ALB) Rule and how to use it, see [What is Rule](https://www.alibabacloud.com/help/doc-detail/214375.htm).
//
// > **NOTE:** Available in v1.133.0+.
//
// > **NOTE:** This version only supports forwarding rules in the request direction.
//
// ## Import
//
// Application Load Balancer (ALB) Rule can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:alb/rule:Rule example <id>
// ```
type Rule struct {
	pulumi.CustomResourceState

	// Specifies whether to precheck this request.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The ID of the listener to which the forwarding rule belongs.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// The priority of the rule. Valid values: 1 to 10000. A smaller value indicates a higher priority. **Note*:* The priority of each rule within the same listener must be unique.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The actions of the forwarding rules. See the following `Block ruleActions`.
	RuleActions RuleRuleActionArrayOutput `pulumi:"ruleActions"`
	// The conditions of the forwarding rule. See the following `Block ruleConditions`.
	RuleConditions RuleRuleConditionArrayOutput `pulumi:"ruleConditions"`
	// The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	RuleName pulumi.StringOutput `pulumi:"ruleName"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ListenerId == nil {
		return nil, errors.New("invalid value for required argument 'ListenerId'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.RuleActions == nil {
		return nil, errors.New("invalid value for required argument 'RuleActions'")
	}
	if args.RuleConditions == nil {
		return nil, errors.New("invalid value for required argument 'RuleConditions'")
	}
	if args.RuleName == nil {
		return nil, errors.New("invalid value for required argument 'RuleName'")
	}
	var resource Rule
	err := ctx.RegisterResource("alicloud:alb/rule:Rule", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:alb/rule:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	// Specifies whether to precheck this request.
	DryRun *bool `pulumi:"dryRun"`
	// The ID of the listener to which the forwarding rule belongs.
	ListenerId *string `pulumi:"listenerId"`
	// The priority of the rule. Valid values: 1 to 10000. A smaller value indicates a higher priority. **Note*:* The priority of each rule within the same listener must be unique.
	Priority *int `pulumi:"priority"`
	// The actions of the forwarding rules. See the following `Block ruleActions`.
	RuleActions []RuleRuleAction `pulumi:"ruleActions"`
	// The conditions of the forwarding rule. See the following `Block ruleConditions`.
	RuleConditions []RuleRuleCondition `pulumi:"ruleConditions"`
	// The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	RuleName *string `pulumi:"ruleName"`
	// The status of the resource.
	Status *string `pulumi:"status"`
}

type RuleState struct {
	// Specifies whether to precheck this request.
	DryRun pulumi.BoolPtrInput
	// The ID of the listener to which the forwarding rule belongs.
	ListenerId pulumi.StringPtrInput
	// The priority of the rule. Valid values: 1 to 10000. A smaller value indicates a higher priority. **Note*:* The priority of each rule within the same listener must be unique.
	Priority pulumi.IntPtrInput
	// The actions of the forwarding rules. See the following `Block ruleActions`.
	RuleActions RuleRuleActionArrayInput
	// The conditions of the forwarding rule. See the following `Block ruleConditions`.
	RuleConditions RuleRuleConditionArrayInput
	// The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	RuleName pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// Specifies whether to precheck this request.
	DryRun *bool `pulumi:"dryRun"`
	// The ID of the listener to which the forwarding rule belongs.
	ListenerId string `pulumi:"listenerId"`
	// The priority of the rule. Valid values: 1 to 10000. A smaller value indicates a higher priority. **Note*:* The priority of each rule within the same listener must be unique.
	Priority int `pulumi:"priority"`
	// The actions of the forwarding rules. See the following `Block ruleActions`.
	RuleActions []RuleRuleAction `pulumi:"ruleActions"`
	// The conditions of the forwarding rule. See the following `Block ruleConditions`.
	RuleConditions []RuleRuleCondition `pulumi:"ruleConditions"`
	// The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	RuleName string `pulumi:"ruleName"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// Specifies whether to precheck this request.
	DryRun pulumi.BoolPtrInput
	// The ID of the listener to which the forwarding rule belongs.
	ListenerId pulumi.StringInput
	// The priority of the rule. Valid values: 1 to 10000. A smaller value indicates a higher priority. **Note*:* The priority of each rule within the same listener must be unique.
	Priority pulumi.IntInput
	// The actions of the forwarding rules. See the following `Block ruleActions`.
	RuleActions RuleRuleActionArrayInput
	// The conditions of the forwarding rule. See the following `Block ruleConditions`.
	RuleConditions RuleRuleConditionArrayInput
	// The name of the forwarding rule. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	RuleName pulumi.StringInput
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
	return reflect.TypeOf((*Rule)(nil))
}

func (i *Rule) ToRuleOutput() RuleOutput {
	return i.ToRuleOutputWithContext(context.Background())
}

func (i *Rule) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleOutput)
}

func (i *Rule) ToRulePtrOutput() RulePtrOutput {
	return i.ToRulePtrOutputWithContext(context.Background())
}

func (i *Rule) ToRulePtrOutputWithContext(ctx context.Context) RulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulePtrOutput)
}

type RulePtrInput interface {
	pulumi.Input

	ToRulePtrOutput() RulePtrOutput
	ToRulePtrOutputWithContext(ctx context.Context) RulePtrOutput
}

type rulePtrType RuleArgs

func (*rulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil))
}

func (i *rulePtrType) ToRulePtrOutput() RulePtrOutput {
	return i.ToRulePtrOutputWithContext(context.Background())
}

func (i *rulePtrType) ToRulePtrOutputWithContext(ctx context.Context) RulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulePtrOutput)
}

// RuleArrayInput is an input type that accepts RuleArray and RuleArrayOutput values.
// You can construct a concrete instance of `RuleArrayInput` via:
//
//          RuleArray{ RuleArgs{...} }
type RuleArrayInput interface {
	pulumi.Input

	ToRuleArrayOutput() RuleArrayOutput
	ToRuleArrayOutputWithContext(context.Context) RuleArrayOutput
}

type RuleArray []RuleInput

func (RuleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Rule)(nil))
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
//          RuleMap{ "key": RuleArgs{...} }
type RuleMapInput interface {
	pulumi.Input

	ToRuleMapOutput() RuleMapOutput
	ToRuleMapOutputWithContext(context.Context) RuleMapOutput
}

type RuleMap map[string]RuleInput

func (RuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Rule)(nil))
}

func (i RuleMap) ToRuleMapOutput() RuleMapOutput {
	return i.ToRuleMapOutputWithContext(context.Background())
}

func (i RuleMap) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleMapOutput)
}

type RuleOutput struct {
	*pulumi.OutputState
}

func (RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Rule)(nil))
}

func (o RuleOutput) ToRuleOutput() RuleOutput {
	return o
}

func (o RuleOutput) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return o
}

func (o RuleOutput) ToRulePtrOutput() RulePtrOutput {
	return o.ToRulePtrOutputWithContext(context.Background())
}

func (o RuleOutput) ToRulePtrOutputWithContext(ctx context.Context) RulePtrOutput {
	return o.ApplyT(func(v Rule) *Rule {
		return &v
	}).(RulePtrOutput)
}

type RulePtrOutput struct {
	*pulumi.OutputState
}

func (RulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil))
}

func (o RulePtrOutput) ToRulePtrOutput() RulePtrOutput {
	return o
}

func (o RulePtrOutput) ToRulePtrOutputWithContext(ctx context.Context) RulePtrOutput {
	return o
}

type RuleArrayOutput struct{ *pulumi.OutputState }

func (RuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Rule)(nil))
}

func (o RuleArrayOutput) ToRuleArrayOutput() RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) Index(i pulumi.IntInput) RuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Rule {
		return vs[0].([]Rule)[vs[1].(int)]
	}).(RuleOutput)
}

type RuleMapOutput struct{ *pulumi.OutputState }

func (RuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Rule)(nil))
}

func (o RuleMapOutput) ToRuleMapOutput() RuleMapOutput {
	return o
}

func (o RuleMapOutput) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return o
}

func (o RuleMapOutput) MapIndex(k pulumi.StringInput) RuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Rule {
		return vs[0].(map[string]Rule)[vs[1].(string)]
	}).(RuleOutput)
}

func init() {
	pulumi.RegisterOutputType(RuleOutput{})
	pulumi.RegisterOutputType(RulePtrOutput{})
	pulumi.RegisterOutputType(RuleArrayOutput{})
	pulumi.RegisterOutputType(RuleMapOutput{})
}