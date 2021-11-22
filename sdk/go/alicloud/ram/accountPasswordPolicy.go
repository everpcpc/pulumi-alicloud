// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// RAM account password policy can be imported using the `id`, e.g. bash
//
// ```sh
//  $ pulumi import alicloud:ram/accountPasswordPolicy:AccountPasswordPolicy example ram-account-password-policy
// ```
type AccountPasswordPolicy struct {
	pulumi.CustomResourceState

	// Specifies if a password can expire in a hard way. Default to false.
	HardExpiry pulumi.BoolPtrOutput `pulumi:"hardExpiry"`
	// Maximum logon attempts with an incorrect password within an hour. Valid value range: [0-32]. Default to 5.
	MaxLoginAttempts pulumi.IntPtrOutput `pulumi:"maxLoginAttempts"`
	// The number of days after which password expires. A value of 0 indicates that the password never expires. Valid value range: [0-1095]. Default to 0.
	MaxPasswordAge pulumi.IntPtrOutput `pulumi:"maxPasswordAge"`
	// Minimal required length of password for a user. Valid value range: [8-32]. Default to 12.
	MinimumPasswordLength pulumi.IntPtrOutput `pulumi:"minimumPasswordLength"`
	// User is not allowed to use the latest number of passwords specified in this parameter. A value of 0 indicates the password history check policy is disabled. Valid value range: [0-24]. Default to 0.
	PasswordReusePrevention pulumi.IntPtrOutput `pulumi:"passwordReusePrevention"`
	// Specifies if the occurrence of a lowercase character in the password is mandatory. Default to true.
	RequireLowercaseCharacters pulumi.BoolPtrOutput `pulumi:"requireLowercaseCharacters"`
	// Specifies if the occurrence of a number in the password is mandatory. Default to true.
	RequireNumbers pulumi.BoolPtrOutput `pulumi:"requireNumbers"`
	// (Optional Specifies if the occurrence of a special character in the password is mandatory. Default to true.
	RequireSymbols pulumi.BoolPtrOutput `pulumi:"requireSymbols"`
	// Specifies if the occurrence of an uppercase character in the password is mandatory. Default to true.
	RequireUppercaseCharacters pulumi.BoolPtrOutput `pulumi:"requireUppercaseCharacters"`
}

// NewAccountPasswordPolicy registers a new resource with the given unique name, arguments, and options.
func NewAccountPasswordPolicy(ctx *pulumi.Context,
	name string, args *AccountPasswordPolicyArgs, opts ...pulumi.ResourceOption) (*AccountPasswordPolicy, error) {
	if args == nil {
		args = &AccountPasswordPolicyArgs{}
	}

	var resource AccountPasswordPolicy
	err := ctx.RegisterResource("alicloud:ram/accountPasswordPolicy:AccountPasswordPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountPasswordPolicy gets an existing AccountPasswordPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountPasswordPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountPasswordPolicyState, opts ...pulumi.ResourceOption) (*AccountPasswordPolicy, error) {
	var resource AccountPasswordPolicy
	err := ctx.ReadResource("alicloud:ram/accountPasswordPolicy:AccountPasswordPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountPasswordPolicy resources.
type accountPasswordPolicyState struct {
	// Specifies if a password can expire in a hard way. Default to false.
	HardExpiry *bool `pulumi:"hardExpiry"`
	// Maximum logon attempts with an incorrect password within an hour. Valid value range: [0-32]. Default to 5.
	MaxLoginAttempts *int `pulumi:"maxLoginAttempts"`
	// The number of days after which password expires. A value of 0 indicates that the password never expires. Valid value range: [0-1095]. Default to 0.
	MaxPasswordAge *int `pulumi:"maxPasswordAge"`
	// Minimal required length of password for a user. Valid value range: [8-32]. Default to 12.
	MinimumPasswordLength *int `pulumi:"minimumPasswordLength"`
	// User is not allowed to use the latest number of passwords specified in this parameter. A value of 0 indicates the password history check policy is disabled. Valid value range: [0-24]. Default to 0.
	PasswordReusePrevention *int `pulumi:"passwordReusePrevention"`
	// Specifies if the occurrence of a lowercase character in the password is mandatory. Default to true.
	RequireLowercaseCharacters *bool `pulumi:"requireLowercaseCharacters"`
	// Specifies if the occurrence of a number in the password is mandatory. Default to true.
	RequireNumbers *bool `pulumi:"requireNumbers"`
	// (Optional Specifies if the occurrence of a special character in the password is mandatory. Default to true.
	RequireSymbols *bool `pulumi:"requireSymbols"`
	// Specifies if the occurrence of an uppercase character in the password is mandatory. Default to true.
	RequireUppercaseCharacters *bool `pulumi:"requireUppercaseCharacters"`
}

type AccountPasswordPolicyState struct {
	// Specifies if a password can expire in a hard way. Default to false.
	HardExpiry pulumi.BoolPtrInput
	// Maximum logon attempts with an incorrect password within an hour. Valid value range: [0-32]. Default to 5.
	MaxLoginAttempts pulumi.IntPtrInput
	// The number of days after which password expires. A value of 0 indicates that the password never expires. Valid value range: [0-1095]. Default to 0.
	MaxPasswordAge pulumi.IntPtrInput
	// Minimal required length of password for a user. Valid value range: [8-32]. Default to 12.
	MinimumPasswordLength pulumi.IntPtrInput
	// User is not allowed to use the latest number of passwords specified in this parameter. A value of 0 indicates the password history check policy is disabled. Valid value range: [0-24]. Default to 0.
	PasswordReusePrevention pulumi.IntPtrInput
	// Specifies if the occurrence of a lowercase character in the password is mandatory. Default to true.
	RequireLowercaseCharacters pulumi.BoolPtrInput
	// Specifies if the occurrence of a number in the password is mandatory. Default to true.
	RequireNumbers pulumi.BoolPtrInput
	// (Optional Specifies if the occurrence of a special character in the password is mandatory. Default to true.
	RequireSymbols pulumi.BoolPtrInput
	// Specifies if the occurrence of an uppercase character in the password is mandatory. Default to true.
	RequireUppercaseCharacters pulumi.BoolPtrInput
}

func (AccountPasswordPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPasswordPolicyState)(nil)).Elem()
}

type accountPasswordPolicyArgs struct {
	// Specifies if a password can expire in a hard way. Default to false.
	HardExpiry *bool `pulumi:"hardExpiry"`
	// Maximum logon attempts with an incorrect password within an hour. Valid value range: [0-32]. Default to 5.
	MaxLoginAttempts *int `pulumi:"maxLoginAttempts"`
	// The number of days after which password expires. A value of 0 indicates that the password never expires. Valid value range: [0-1095]. Default to 0.
	MaxPasswordAge *int `pulumi:"maxPasswordAge"`
	// Minimal required length of password for a user. Valid value range: [8-32]. Default to 12.
	MinimumPasswordLength *int `pulumi:"minimumPasswordLength"`
	// User is not allowed to use the latest number of passwords specified in this parameter. A value of 0 indicates the password history check policy is disabled. Valid value range: [0-24]. Default to 0.
	PasswordReusePrevention *int `pulumi:"passwordReusePrevention"`
	// Specifies if the occurrence of a lowercase character in the password is mandatory. Default to true.
	RequireLowercaseCharacters *bool `pulumi:"requireLowercaseCharacters"`
	// Specifies if the occurrence of a number in the password is mandatory. Default to true.
	RequireNumbers *bool `pulumi:"requireNumbers"`
	// (Optional Specifies if the occurrence of a special character in the password is mandatory. Default to true.
	RequireSymbols *bool `pulumi:"requireSymbols"`
	// Specifies if the occurrence of an uppercase character in the password is mandatory. Default to true.
	RequireUppercaseCharacters *bool `pulumi:"requireUppercaseCharacters"`
}

// The set of arguments for constructing a AccountPasswordPolicy resource.
type AccountPasswordPolicyArgs struct {
	// Specifies if a password can expire in a hard way. Default to false.
	HardExpiry pulumi.BoolPtrInput
	// Maximum logon attempts with an incorrect password within an hour. Valid value range: [0-32]. Default to 5.
	MaxLoginAttempts pulumi.IntPtrInput
	// The number of days after which password expires. A value of 0 indicates that the password never expires. Valid value range: [0-1095]. Default to 0.
	MaxPasswordAge pulumi.IntPtrInput
	// Minimal required length of password for a user. Valid value range: [8-32]. Default to 12.
	MinimumPasswordLength pulumi.IntPtrInput
	// User is not allowed to use the latest number of passwords specified in this parameter. A value of 0 indicates the password history check policy is disabled. Valid value range: [0-24]. Default to 0.
	PasswordReusePrevention pulumi.IntPtrInput
	// Specifies if the occurrence of a lowercase character in the password is mandatory. Default to true.
	RequireLowercaseCharacters pulumi.BoolPtrInput
	// Specifies if the occurrence of a number in the password is mandatory. Default to true.
	RequireNumbers pulumi.BoolPtrInput
	// (Optional Specifies if the occurrence of a special character in the password is mandatory. Default to true.
	RequireSymbols pulumi.BoolPtrInput
	// Specifies if the occurrence of an uppercase character in the password is mandatory. Default to true.
	RequireUppercaseCharacters pulumi.BoolPtrInput
}

func (AccountPasswordPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPasswordPolicyArgs)(nil)).Elem()
}

type AccountPasswordPolicyInput interface {
	pulumi.Input

	ToAccountPasswordPolicyOutput() AccountPasswordPolicyOutput
	ToAccountPasswordPolicyOutputWithContext(ctx context.Context) AccountPasswordPolicyOutput
}

func (*AccountPasswordPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountPasswordPolicy)(nil))
}

func (i *AccountPasswordPolicy) ToAccountPasswordPolicyOutput() AccountPasswordPolicyOutput {
	return i.ToAccountPasswordPolicyOutputWithContext(context.Background())
}

func (i *AccountPasswordPolicy) ToAccountPasswordPolicyOutputWithContext(ctx context.Context) AccountPasswordPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPasswordPolicyOutput)
}

func (i *AccountPasswordPolicy) ToAccountPasswordPolicyPtrOutput() AccountPasswordPolicyPtrOutput {
	return i.ToAccountPasswordPolicyPtrOutputWithContext(context.Background())
}

func (i *AccountPasswordPolicy) ToAccountPasswordPolicyPtrOutputWithContext(ctx context.Context) AccountPasswordPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPasswordPolicyPtrOutput)
}

type AccountPasswordPolicyPtrInput interface {
	pulumi.Input

	ToAccountPasswordPolicyPtrOutput() AccountPasswordPolicyPtrOutput
	ToAccountPasswordPolicyPtrOutputWithContext(ctx context.Context) AccountPasswordPolicyPtrOutput
}

type accountPasswordPolicyPtrType AccountPasswordPolicyArgs

func (*accountPasswordPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountPasswordPolicy)(nil))
}

func (i *accountPasswordPolicyPtrType) ToAccountPasswordPolicyPtrOutput() AccountPasswordPolicyPtrOutput {
	return i.ToAccountPasswordPolicyPtrOutputWithContext(context.Background())
}

func (i *accountPasswordPolicyPtrType) ToAccountPasswordPolicyPtrOutputWithContext(ctx context.Context) AccountPasswordPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPasswordPolicyPtrOutput)
}

// AccountPasswordPolicyArrayInput is an input type that accepts AccountPasswordPolicyArray and AccountPasswordPolicyArrayOutput values.
// You can construct a concrete instance of `AccountPasswordPolicyArrayInput` via:
//
//          AccountPasswordPolicyArray{ AccountPasswordPolicyArgs{...} }
type AccountPasswordPolicyArrayInput interface {
	pulumi.Input

	ToAccountPasswordPolicyArrayOutput() AccountPasswordPolicyArrayOutput
	ToAccountPasswordPolicyArrayOutputWithContext(context.Context) AccountPasswordPolicyArrayOutput
}

type AccountPasswordPolicyArray []AccountPasswordPolicyInput

func (AccountPasswordPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountPasswordPolicy)(nil)).Elem()
}

func (i AccountPasswordPolicyArray) ToAccountPasswordPolicyArrayOutput() AccountPasswordPolicyArrayOutput {
	return i.ToAccountPasswordPolicyArrayOutputWithContext(context.Background())
}

func (i AccountPasswordPolicyArray) ToAccountPasswordPolicyArrayOutputWithContext(ctx context.Context) AccountPasswordPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPasswordPolicyArrayOutput)
}

// AccountPasswordPolicyMapInput is an input type that accepts AccountPasswordPolicyMap and AccountPasswordPolicyMapOutput values.
// You can construct a concrete instance of `AccountPasswordPolicyMapInput` via:
//
//          AccountPasswordPolicyMap{ "key": AccountPasswordPolicyArgs{...} }
type AccountPasswordPolicyMapInput interface {
	pulumi.Input

	ToAccountPasswordPolicyMapOutput() AccountPasswordPolicyMapOutput
	ToAccountPasswordPolicyMapOutputWithContext(context.Context) AccountPasswordPolicyMapOutput
}

type AccountPasswordPolicyMap map[string]AccountPasswordPolicyInput

func (AccountPasswordPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountPasswordPolicy)(nil)).Elem()
}

func (i AccountPasswordPolicyMap) ToAccountPasswordPolicyMapOutput() AccountPasswordPolicyMapOutput {
	return i.ToAccountPasswordPolicyMapOutputWithContext(context.Background())
}

func (i AccountPasswordPolicyMap) ToAccountPasswordPolicyMapOutputWithContext(ctx context.Context) AccountPasswordPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPasswordPolicyMapOutput)
}

type AccountPasswordPolicyOutput struct{ *pulumi.OutputState }

func (AccountPasswordPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountPasswordPolicy)(nil))
}

func (o AccountPasswordPolicyOutput) ToAccountPasswordPolicyOutput() AccountPasswordPolicyOutput {
	return o
}

func (o AccountPasswordPolicyOutput) ToAccountPasswordPolicyOutputWithContext(ctx context.Context) AccountPasswordPolicyOutput {
	return o
}

func (o AccountPasswordPolicyOutput) ToAccountPasswordPolicyPtrOutput() AccountPasswordPolicyPtrOutput {
	return o.ToAccountPasswordPolicyPtrOutputWithContext(context.Background())
}

func (o AccountPasswordPolicyOutput) ToAccountPasswordPolicyPtrOutputWithContext(ctx context.Context) AccountPasswordPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountPasswordPolicy) *AccountPasswordPolicy {
		return &v
	}).(AccountPasswordPolicyPtrOutput)
}

type AccountPasswordPolicyPtrOutput struct{ *pulumi.OutputState }

func (AccountPasswordPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountPasswordPolicy)(nil))
}

func (o AccountPasswordPolicyPtrOutput) ToAccountPasswordPolicyPtrOutput() AccountPasswordPolicyPtrOutput {
	return o
}

func (o AccountPasswordPolicyPtrOutput) ToAccountPasswordPolicyPtrOutputWithContext(ctx context.Context) AccountPasswordPolicyPtrOutput {
	return o
}

func (o AccountPasswordPolicyPtrOutput) Elem() AccountPasswordPolicyOutput {
	return o.ApplyT(func(v *AccountPasswordPolicy) AccountPasswordPolicy {
		if v != nil {
			return *v
		}
		var ret AccountPasswordPolicy
		return ret
	}).(AccountPasswordPolicyOutput)
}

type AccountPasswordPolicyArrayOutput struct{ *pulumi.OutputState }

func (AccountPasswordPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccountPasswordPolicy)(nil))
}

func (o AccountPasswordPolicyArrayOutput) ToAccountPasswordPolicyArrayOutput() AccountPasswordPolicyArrayOutput {
	return o
}

func (o AccountPasswordPolicyArrayOutput) ToAccountPasswordPolicyArrayOutputWithContext(ctx context.Context) AccountPasswordPolicyArrayOutput {
	return o
}

func (o AccountPasswordPolicyArrayOutput) Index(i pulumi.IntInput) AccountPasswordPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccountPasswordPolicy {
		return vs[0].([]AccountPasswordPolicy)[vs[1].(int)]
	}).(AccountPasswordPolicyOutput)
}

type AccountPasswordPolicyMapOutput struct{ *pulumi.OutputState }

func (AccountPasswordPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AccountPasswordPolicy)(nil))
}

func (o AccountPasswordPolicyMapOutput) ToAccountPasswordPolicyMapOutput() AccountPasswordPolicyMapOutput {
	return o
}

func (o AccountPasswordPolicyMapOutput) ToAccountPasswordPolicyMapOutputWithContext(ctx context.Context) AccountPasswordPolicyMapOutput {
	return o
}

func (o AccountPasswordPolicyMapOutput) MapIndex(k pulumi.StringInput) AccountPasswordPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AccountPasswordPolicy {
		return vs[0].(map[string]AccountPasswordPolicy)[vs[1].(string)]
	}).(AccountPasswordPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPasswordPolicyInput)(nil)).Elem(), &AccountPasswordPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPasswordPolicyPtrInput)(nil)).Elem(), &AccountPasswordPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPasswordPolicyArrayInput)(nil)).Elem(), AccountPasswordPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPasswordPolicyMapInput)(nil)).Elem(), AccountPasswordPolicyMap{})
	pulumi.RegisterOutputType(AccountPasswordPolicyOutput{})
	pulumi.RegisterOutputType(AccountPasswordPolicyPtrOutput{})
	pulumi.RegisterOutputType(AccountPasswordPolicyArrayOutput{})
	pulumi.RegisterOutputType(AccountPasswordPolicyMapOutput{})
}
