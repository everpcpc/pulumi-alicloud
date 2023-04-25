// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Dcdn Waf Rule resource.
//
// For information about Dcdn Waf Rule and how to use it, see [What is Waf Rule](https://www.alibabacloud.com/help/en/dynamic-route-for-cdn/latest/configure-protection-rules).
//
// > **NOTE:** Available in v1.201.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dcdn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultWafPolicy, err := dcdn.NewWafPolicy(ctx, "defaultWafPolicy", &dcdn.WafPolicyArgs{
//				DefenseScene: pulumi.String("custom_acl"),
//				PolicyName:   pulumi.Any(_var.Name),
//				PolicyType:   pulumi.String("custom"),
//				Status:       pulumi.String("on"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dcdn.NewWafRule(ctx, "defaultWafRule", &dcdn.WafRuleArgs{
//				PolicyId: defaultWafPolicy.ID(),
//				RuleName: pulumi.Any(_var.Name),
//				Conditions: dcdn.WafRuleConditionArray{
//					&dcdn.WafRuleConditionArgs{
//						Key:     pulumi.String("URI"),
//						OpValue: pulumi.String("ne"),
//						Values:  pulumi.String("/login.php"),
//					},
//					&dcdn.WafRuleConditionArgs{
//						Key:     pulumi.String("Header"),
//						SubKey:  pulumi.String("a"),
//						OpValue: pulumi.String("eq"),
//						Values:  pulumi.String("b"),
//					},
//				},
//				Status:   pulumi.String("on"),
//				CcStatus: pulumi.String("on"),
//				Action:   pulumi.String("monitor"),
//				Effect:   pulumi.String("rule"),
//				RateLimit: &dcdn.WafRuleRateLimitArgs{
//					Target:    pulumi.String("IP"),
//					Interval:  pulumi.Int(5),
//					Threshold: pulumi.Int(5),
//					Ttl:       pulumi.Int(1800),
//					Status: &dcdn.WafRuleRateLimitStatusArgs{
//						Code:  pulumi.String("200"),
//						Ratio: pulumi.Int(60),
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
// Dcdn Waf Rule can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:dcdn/wafRule:WafRule example <id>
//
// ```
type WafRule struct {
	pulumi.CustomResourceState

	// Specifies the action of the rule. Valid values: `block`, `monitor`, `js`.
	Action pulumi.StringPtrOutput `pulumi:"action"`
	// Specifies whether to enable rate limiting. Valid values: `on` and `off`. **NOTE:** This parameter is required when policy is of type `customAcl`.
	CcStatus pulumi.StringOutput `pulumi:"ccStatus"`
	// The blocked regions in the Chinese mainland, separated by commas (,).
	CnRegionList pulumi.StringPtrOutput `pulumi:"cnRegionList"`
	// Conditions that trigger the rule. See the following `Block Conditions`. **NOTE:** This parameter is required when policy is of type `customAcl` or `whitelist`.
	Conditions WafRuleConditionArrayOutput `pulumi:"conditions"`
	// The type of protection policy. The following scenarios are supported:-waf_group:Web basic protection-custom_acl: Custom protection policy-whitelist: whitelist
	DefenseScene pulumi.StringOutput `pulumi:"defenseScene"`
	// The effective scope of the rate limiting blacklist. If you set ccStatus to on, you must configure this parameter. Valid values: `rule` (takes effect for the current rule) and `service` (takes effect globally).
	Effect pulumi.StringPtrOutput `pulumi:"effect"`
	// Revised the time. The date format is based on ISO8601 notation and uses UTC +0 time in the format of yyyy-MM-ddTHH:mm:ssZ.
	GmtModified pulumi.StringOutput `pulumi:"gmtModified"`
	// Blocked regions outside the Chinese mainland, separated by commas (,).
	OtherRegionList pulumi.StringPtrOutput `pulumi:"otherRegionList"`
	// The protection policy ID.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The rules of rate limiting. If you set `ccStatus` to on, you must configure this parameter. See the following `Block RateLimit`.
	RateLimit WafRuleRateLimitPtrOutput `pulumi:"rateLimit"`
	// The regular expression.e, when wafGroup appears in tags, this value can be filled in, and only one list of six digits in string format can appear with regultypes.
	RegularRules pulumi.StringArrayOutput `pulumi:"regularRules"`
	// Regular rule type, when wafGroup appears in tags, this value can be filled in, optional values:["sqli", "xss", "codeExec", "crlf", "lfileii", "rfileii", "webshell", "vvip", "other"]
	RegularTypes pulumi.StringArrayOutput `pulumi:"regularTypes"`
	// Filter by IP address.
	RemoteAddrs pulumi.StringArrayOutput `pulumi:"remoteAddrs"`
	// The name of the protection rule. The name can be up to 64 characters in length and can contain letters, digits, and underscores (_). **NOTE:** This parameter cannot be modified when policy is of type `regionBlock`.
	RuleName pulumi.StringOutput `pulumi:"ruleName"`
	// The types of the protection policies.
	Scenes pulumi.StringArrayOutput `pulumi:"scenes"`
	// The status of the waf rule. Valid values: `on` and `off`. Default value: on.
	Status pulumi.StringOutput `pulumi:"status"`
	// The id of the waf rule group. The default value is "1012". Multiple rules are separated by commas.
	WafGroupIds pulumi.StringPtrOutput `pulumi:"wafGroupIds"`
}

// NewWafRule registers a new resource with the given unique name, arguments, and options.
func NewWafRule(ctx *pulumi.Context,
	name string, args *WafRuleArgs, opts ...pulumi.ResourceOption) (*WafRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	if args.RuleName == nil {
		return nil, errors.New("invalid value for required argument 'RuleName'")
	}
	var resource WafRule
	err := ctx.RegisterResource("alicloud:dcdn/wafRule:WafRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWafRule gets an existing WafRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWafRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WafRuleState, opts ...pulumi.ResourceOption) (*WafRule, error) {
	var resource WafRule
	err := ctx.ReadResource("alicloud:dcdn/wafRule:WafRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WafRule resources.
type wafRuleState struct {
	// Specifies the action of the rule. Valid values: `block`, `monitor`, `js`.
	Action *string `pulumi:"action"`
	// Specifies whether to enable rate limiting. Valid values: `on` and `off`. **NOTE:** This parameter is required when policy is of type `customAcl`.
	CcStatus *string `pulumi:"ccStatus"`
	// The blocked regions in the Chinese mainland, separated by commas (,).
	CnRegionList *string `pulumi:"cnRegionList"`
	// Conditions that trigger the rule. See the following `Block Conditions`. **NOTE:** This parameter is required when policy is of type `customAcl` or `whitelist`.
	Conditions []WafRuleCondition `pulumi:"conditions"`
	// The type of protection policy. The following scenarios are supported:-waf_group:Web basic protection-custom_acl: Custom protection policy-whitelist: whitelist
	DefenseScene *string `pulumi:"defenseScene"`
	// The effective scope of the rate limiting blacklist. If you set ccStatus to on, you must configure this parameter. Valid values: `rule` (takes effect for the current rule) and `service` (takes effect globally).
	Effect *string `pulumi:"effect"`
	// Revised the time. The date format is based on ISO8601 notation and uses UTC +0 time in the format of yyyy-MM-ddTHH:mm:ssZ.
	GmtModified *string `pulumi:"gmtModified"`
	// Blocked regions outside the Chinese mainland, separated by commas (,).
	OtherRegionList *string `pulumi:"otherRegionList"`
	// The protection policy ID.
	PolicyId *string `pulumi:"policyId"`
	// The rules of rate limiting. If you set `ccStatus` to on, you must configure this parameter. See the following `Block RateLimit`.
	RateLimit *WafRuleRateLimit `pulumi:"rateLimit"`
	// The regular expression.e, when wafGroup appears in tags, this value can be filled in, and only one list of six digits in string format can appear with regultypes.
	RegularRules []string `pulumi:"regularRules"`
	// Regular rule type, when wafGroup appears in tags, this value can be filled in, optional values:["sqli", "xss", "codeExec", "crlf", "lfileii", "rfileii", "webshell", "vvip", "other"]
	RegularTypes []string `pulumi:"regularTypes"`
	// Filter by IP address.
	RemoteAddrs []string `pulumi:"remoteAddrs"`
	// The name of the protection rule. The name can be up to 64 characters in length and can contain letters, digits, and underscores (_). **NOTE:** This parameter cannot be modified when policy is of type `regionBlock`.
	RuleName *string `pulumi:"ruleName"`
	// The types of the protection policies.
	Scenes []string `pulumi:"scenes"`
	// The status of the waf rule. Valid values: `on` and `off`. Default value: on.
	Status *string `pulumi:"status"`
	// The id of the waf rule group. The default value is "1012". Multiple rules are separated by commas.
	WafGroupIds *string `pulumi:"wafGroupIds"`
}

type WafRuleState struct {
	// Specifies the action of the rule. Valid values: `block`, `monitor`, `js`.
	Action pulumi.StringPtrInput
	// Specifies whether to enable rate limiting. Valid values: `on` and `off`. **NOTE:** This parameter is required when policy is of type `customAcl`.
	CcStatus pulumi.StringPtrInput
	// The blocked regions in the Chinese mainland, separated by commas (,).
	CnRegionList pulumi.StringPtrInput
	// Conditions that trigger the rule. See the following `Block Conditions`. **NOTE:** This parameter is required when policy is of type `customAcl` or `whitelist`.
	Conditions WafRuleConditionArrayInput
	// The type of protection policy. The following scenarios are supported:-waf_group:Web basic protection-custom_acl: Custom protection policy-whitelist: whitelist
	DefenseScene pulumi.StringPtrInput
	// The effective scope of the rate limiting blacklist. If you set ccStatus to on, you must configure this parameter. Valid values: `rule` (takes effect for the current rule) and `service` (takes effect globally).
	Effect pulumi.StringPtrInput
	// Revised the time. The date format is based on ISO8601 notation and uses UTC +0 time in the format of yyyy-MM-ddTHH:mm:ssZ.
	GmtModified pulumi.StringPtrInput
	// Blocked regions outside the Chinese mainland, separated by commas (,).
	OtherRegionList pulumi.StringPtrInput
	// The protection policy ID.
	PolicyId pulumi.StringPtrInput
	// The rules of rate limiting. If you set `ccStatus` to on, you must configure this parameter. See the following `Block RateLimit`.
	RateLimit WafRuleRateLimitPtrInput
	// The regular expression.e, when wafGroup appears in tags, this value can be filled in, and only one list of six digits in string format can appear with regultypes.
	RegularRules pulumi.StringArrayInput
	// Regular rule type, when wafGroup appears in tags, this value can be filled in, optional values:["sqli", "xss", "codeExec", "crlf", "lfileii", "rfileii", "webshell", "vvip", "other"]
	RegularTypes pulumi.StringArrayInput
	// Filter by IP address.
	RemoteAddrs pulumi.StringArrayInput
	// The name of the protection rule. The name can be up to 64 characters in length and can contain letters, digits, and underscores (_). **NOTE:** This parameter cannot be modified when policy is of type `regionBlock`.
	RuleName pulumi.StringPtrInput
	// The types of the protection policies.
	Scenes pulumi.StringArrayInput
	// The status of the waf rule. Valid values: `on` and `off`. Default value: on.
	Status pulumi.StringPtrInput
	// The id of the waf rule group. The default value is "1012". Multiple rules are separated by commas.
	WafGroupIds pulumi.StringPtrInput
}

func (WafRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*wafRuleState)(nil)).Elem()
}

type wafRuleArgs struct {
	// Specifies the action of the rule. Valid values: `block`, `monitor`, `js`.
	Action *string `pulumi:"action"`
	// Specifies whether to enable rate limiting. Valid values: `on` and `off`. **NOTE:** This parameter is required when policy is of type `customAcl`.
	CcStatus *string `pulumi:"ccStatus"`
	// The blocked regions in the Chinese mainland, separated by commas (,).
	CnRegionList *string `pulumi:"cnRegionList"`
	// Conditions that trigger the rule. See the following `Block Conditions`. **NOTE:** This parameter is required when policy is of type `customAcl` or `whitelist`.
	Conditions []WafRuleCondition `pulumi:"conditions"`
	// The effective scope of the rate limiting blacklist. If you set ccStatus to on, you must configure this parameter. Valid values: `rule` (takes effect for the current rule) and `service` (takes effect globally).
	Effect *string `pulumi:"effect"`
	// Blocked regions outside the Chinese mainland, separated by commas (,).
	OtherRegionList *string `pulumi:"otherRegionList"`
	// The protection policy ID.
	PolicyId string `pulumi:"policyId"`
	// The rules of rate limiting. If you set `ccStatus` to on, you must configure this parameter. See the following `Block RateLimit`.
	RateLimit *WafRuleRateLimit `pulumi:"rateLimit"`
	// The regular expression.e, when wafGroup appears in tags, this value can be filled in, and only one list of six digits in string format can appear with regultypes.
	RegularRules []string `pulumi:"regularRules"`
	// Regular rule type, when wafGroup appears in tags, this value can be filled in, optional values:["sqli", "xss", "codeExec", "crlf", "lfileii", "rfileii", "webshell", "vvip", "other"]
	RegularTypes []string `pulumi:"regularTypes"`
	// Filter by IP address.
	RemoteAddrs []string `pulumi:"remoteAddrs"`
	// The name of the protection rule. The name can be up to 64 characters in length and can contain letters, digits, and underscores (_). **NOTE:** This parameter cannot be modified when policy is of type `regionBlock`.
	RuleName string `pulumi:"ruleName"`
	// The types of the protection policies.
	Scenes []string `pulumi:"scenes"`
	// The status of the waf rule. Valid values: `on` and `off`. Default value: on.
	Status *string `pulumi:"status"`
	// The id of the waf rule group. The default value is "1012". Multiple rules are separated by commas.
	WafGroupIds *string `pulumi:"wafGroupIds"`
}

// The set of arguments for constructing a WafRule resource.
type WafRuleArgs struct {
	// Specifies the action of the rule. Valid values: `block`, `monitor`, `js`.
	Action pulumi.StringPtrInput
	// Specifies whether to enable rate limiting. Valid values: `on` and `off`. **NOTE:** This parameter is required when policy is of type `customAcl`.
	CcStatus pulumi.StringPtrInput
	// The blocked regions in the Chinese mainland, separated by commas (,).
	CnRegionList pulumi.StringPtrInput
	// Conditions that trigger the rule. See the following `Block Conditions`. **NOTE:** This parameter is required when policy is of type `customAcl` or `whitelist`.
	Conditions WafRuleConditionArrayInput
	// The effective scope of the rate limiting blacklist. If you set ccStatus to on, you must configure this parameter. Valid values: `rule` (takes effect for the current rule) and `service` (takes effect globally).
	Effect pulumi.StringPtrInput
	// Blocked regions outside the Chinese mainland, separated by commas (,).
	OtherRegionList pulumi.StringPtrInput
	// The protection policy ID.
	PolicyId pulumi.StringInput
	// The rules of rate limiting. If you set `ccStatus` to on, you must configure this parameter. See the following `Block RateLimit`.
	RateLimit WafRuleRateLimitPtrInput
	// The regular expression.e, when wafGroup appears in tags, this value can be filled in, and only one list of six digits in string format can appear with regultypes.
	RegularRules pulumi.StringArrayInput
	// Regular rule type, when wafGroup appears in tags, this value can be filled in, optional values:["sqli", "xss", "codeExec", "crlf", "lfileii", "rfileii", "webshell", "vvip", "other"]
	RegularTypes pulumi.StringArrayInput
	// Filter by IP address.
	RemoteAddrs pulumi.StringArrayInput
	// The name of the protection rule. The name can be up to 64 characters in length and can contain letters, digits, and underscores (_). **NOTE:** This parameter cannot be modified when policy is of type `regionBlock`.
	RuleName pulumi.StringInput
	// The types of the protection policies.
	Scenes pulumi.StringArrayInput
	// The status of the waf rule. Valid values: `on` and `off`. Default value: on.
	Status pulumi.StringPtrInput
	// The id of the waf rule group. The default value is "1012". Multiple rules are separated by commas.
	WafGroupIds pulumi.StringPtrInput
}

func (WafRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wafRuleArgs)(nil)).Elem()
}

type WafRuleInput interface {
	pulumi.Input

	ToWafRuleOutput() WafRuleOutput
	ToWafRuleOutputWithContext(ctx context.Context) WafRuleOutput
}

func (*WafRule) ElementType() reflect.Type {
	return reflect.TypeOf((**WafRule)(nil)).Elem()
}

func (i *WafRule) ToWafRuleOutput() WafRuleOutput {
	return i.ToWafRuleOutputWithContext(context.Background())
}

func (i *WafRule) ToWafRuleOutputWithContext(ctx context.Context) WafRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafRuleOutput)
}

// WafRuleArrayInput is an input type that accepts WafRuleArray and WafRuleArrayOutput values.
// You can construct a concrete instance of `WafRuleArrayInput` via:
//
//	WafRuleArray{ WafRuleArgs{...} }
type WafRuleArrayInput interface {
	pulumi.Input

	ToWafRuleArrayOutput() WafRuleArrayOutput
	ToWafRuleArrayOutputWithContext(context.Context) WafRuleArrayOutput
}

type WafRuleArray []WafRuleInput

func (WafRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WafRule)(nil)).Elem()
}

func (i WafRuleArray) ToWafRuleArrayOutput() WafRuleArrayOutput {
	return i.ToWafRuleArrayOutputWithContext(context.Background())
}

func (i WafRuleArray) ToWafRuleArrayOutputWithContext(ctx context.Context) WafRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafRuleArrayOutput)
}

// WafRuleMapInput is an input type that accepts WafRuleMap and WafRuleMapOutput values.
// You can construct a concrete instance of `WafRuleMapInput` via:
//
//	WafRuleMap{ "key": WafRuleArgs{...} }
type WafRuleMapInput interface {
	pulumi.Input

	ToWafRuleMapOutput() WafRuleMapOutput
	ToWafRuleMapOutputWithContext(context.Context) WafRuleMapOutput
}

type WafRuleMap map[string]WafRuleInput

func (WafRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WafRule)(nil)).Elem()
}

func (i WafRuleMap) ToWafRuleMapOutput() WafRuleMapOutput {
	return i.ToWafRuleMapOutputWithContext(context.Background())
}

func (i WafRuleMap) ToWafRuleMapOutputWithContext(ctx context.Context) WafRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafRuleMapOutput)
}

type WafRuleOutput struct{ *pulumi.OutputState }

func (WafRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WafRule)(nil)).Elem()
}

func (o WafRuleOutput) ToWafRuleOutput() WafRuleOutput {
	return o
}

func (o WafRuleOutput) ToWafRuleOutputWithContext(ctx context.Context) WafRuleOutput {
	return o
}

// Specifies the action of the rule. Valid values: `block`, `monitor`, `js`.
func (o WafRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WafRule) pulumi.StringPtrOutput { return v.Action }).(pulumi.StringPtrOutput)
}

// Specifies whether to enable rate limiting. Valid values: `on` and `off`. **NOTE:** This parameter is required when policy is of type `customAcl`.
func (o WafRuleOutput) CcStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *WafRule) pulumi.StringOutput { return v.CcStatus }).(pulumi.StringOutput)
}

// The blocked regions in the Chinese mainland, separated by commas (,).
func (o WafRuleOutput) CnRegionList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WafRule) pulumi.StringPtrOutput { return v.CnRegionList }).(pulumi.StringPtrOutput)
}

// Conditions that trigger the rule. See the following `Block Conditions`. **NOTE:** This parameter is required when policy is of type `customAcl` or `whitelist`.
func (o WafRuleOutput) Conditions() WafRuleConditionArrayOutput {
	return o.ApplyT(func(v *WafRule) WafRuleConditionArrayOutput { return v.Conditions }).(WafRuleConditionArrayOutput)
}

// The type of protection policy. The following scenarios are supported:-waf_group:Web basic protection-custom_acl: Custom protection policy-whitelist: whitelist
func (o WafRuleOutput) DefenseScene() pulumi.StringOutput {
	return o.ApplyT(func(v *WafRule) pulumi.StringOutput { return v.DefenseScene }).(pulumi.StringOutput)
}

// The effective scope of the rate limiting blacklist. If you set ccStatus to on, you must configure this parameter. Valid values: `rule` (takes effect for the current rule) and `service` (takes effect globally).
func (o WafRuleOutput) Effect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WafRule) pulumi.StringPtrOutput { return v.Effect }).(pulumi.StringPtrOutput)
}

// Revised the time. The date format is based on ISO8601 notation and uses UTC +0 time in the format of yyyy-MM-ddTHH:mm:ssZ.
func (o WafRuleOutput) GmtModified() pulumi.StringOutput {
	return o.ApplyT(func(v *WafRule) pulumi.StringOutput { return v.GmtModified }).(pulumi.StringOutput)
}

// Blocked regions outside the Chinese mainland, separated by commas (,).
func (o WafRuleOutput) OtherRegionList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WafRule) pulumi.StringPtrOutput { return v.OtherRegionList }).(pulumi.StringPtrOutput)
}

// The protection policy ID.
func (o WafRuleOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *WafRule) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// The rules of rate limiting. If you set `ccStatus` to on, you must configure this parameter. See the following `Block RateLimit`.
func (o WafRuleOutput) RateLimit() WafRuleRateLimitPtrOutput {
	return o.ApplyT(func(v *WafRule) WafRuleRateLimitPtrOutput { return v.RateLimit }).(WafRuleRateLimitPtrOutput)
}

// The regular expression.e, when wafGroup appears in tags, this value can be filled in, and only one list of six digits in string format can appear with regultypes.
func (o WafRuleOutput) RegularRules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WafRule) pulumi.StringArrayOutput { return v.RegularRules }).(pulumi.StringArrayOutput)
}

// Regular rule type, when wafGroup appears in tags, this value can be filled in, optional values:["sqli", "xss", "codeExec", "crlf", "lfileii", "rfileii", "webshell", "vvip", "other"]
func (o WafRuleOutput) RegularTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WafRule) pulumi.StringArrayOutput { return v.RegularTypes }).(pulumi.StringArrayOutput)
}

// Filter by IP address.
func (o WafRuleOutput) RemoteAddrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WafRule) pulumi.StringArrayOutput { return v.RemoteAddrs }).(pulumi.StringArrayOutput)
}

// The name of the protection rule. The name can be up to 64 characters in length and can contain letters, digits, and underscores (_). **NOTE:** This parameter cannot be modified when policy is of type `regionBlock`.
func (o WafRuleOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *WafRule) pulumi.StringOutput { return v.RuleName }).(pulumi.StringOutput)
}

// The types of the protection policies.
func (o WafRuleOutput) Scenes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WafRule) pulumi.StringArrayOutput { return v.Scenes }).(pulumi.StringArrayOutput)
}

// The status of the waf rule. Valid values: `on` and `off`. Default value: on.
func (o WafRuleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WafRule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The id of the waf rule group. The default value is "1012". Multiple rules are separated by commas.
func (o WafRuleOutput) WafGroupIds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WafRule) pulumi.StringPtrOutput { return v.WafGroupIds }).(pulumi.StringPtrOutput)
}

type WafRuleArrayOutput struct{ *pulumi.OutputState }

func (WafRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WafRule)(nil)).Elem()
}

func (o WafRuleArrayOutput) ToWafRuleArrayOutput() WafRuleArrayOutput {
	return o
}

func (o WafRuleArrayOutput) ToWafRuleArrayOutputWithContext(ctx context.Context) WafRuleArrayOutput {
	return o
}

func (o WafRuleArrayOutput) Index(i pulumi.IntInput) WafRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WafRule {
		return vs[0].([]*WafRule)[vs[1].(int)]
	}).(WafRuleOutput)
}

type WafRuleMapOutput struct{ *pulumi.OutputState }

func (WafRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WafRule)(nil)).Elem()
}

func (o WafRuleMapOutput) ToWafRuleMapOutput() WafRuleMapOutput {
	return o
}

func (o WafRuleMapOutput) ToWafRuleMapOutputWithContext(ctx context.Context) WafRuleMapOutput {
	return o
}

func (o WafRuleMapOutput) MapIndex(k pulumi.StringInput) WafRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WafRule {
		return vs[0].(map[string]*WafRule)[vs[1].(string)]
	}).(WafRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WafRuleInput)(nil)).Elem(), &WafRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*WafRuleArrayInput)(nil)).Elem(), WafRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WafRuleMapInput)(nil)).Elem(), WafRuleMap{})
	pulumi.RegisterOutputType(WafRuleOutput{})
	pulumi.RegisterOutputType(WafRuleArrayOutput{})
	pulumi.RegisterOutputType(WafRuleMapOutput{})
}
