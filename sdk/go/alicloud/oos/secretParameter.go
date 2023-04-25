// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oos

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a OOS Secret Parameter resource.
//
// For information about OOS Secret Parameter and how to use it, see [What is Secret Parameter](https://www.alibabacloud.com/help/en/doc-detail/183418.html).
//
// > **NOTE:** Available in v1.147.0+.
//
// ## Import
//
// OOS Secret Parameter can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:oos/secretParameter:SecretParameter example <secret_parameter_name>
//
// ```
type SecretParameter struct {
	pulumi.CustomResourceState

	// The constraints of the encryption parameter. By default, this parameter is null. Valid values:
	Constraints pulumi.StringPtrOutput `pulumi:"constraints"`
	// The description of the encryption parameter. The description must be `1` to `200` characters in length.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Customer Master Key (CMK) of Key Management Service (KMS) that is used to encrypt the parameter.
	KeyId pulumi.StringPtrOutput `pulumi:"keyId"`
	// The ID of the Resource Group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The name of the encryption parameter.  The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
	SecretParameterName pulumi.StringOutput `pulumi:"secretParameterName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The data type of the encryption parameter. Valid values: `Secret`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The value of the encryption parameter. The value must be `1` to `4096` characters in length.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewSecretParameter registers a new resource with the given unique name, arguments, and options.
func NewSecretParameter(ctx *pulumi.Context,
	name string, args *SecretParameterArgs, opts ...pulumi.ResourceOption) (*SecretParameter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SecretParameterName == nil {
		return nil, errors.New("invalid value for required argument 'SecretParameterName'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	if args.Value != nil {
		args.Value = pulumi.ToSecret(args.Value).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"value",
	})
	opts = append(opts, secrets)
	var resource SecretParameter
	err := ctx.RegisterResource("alicloud:oos/secretParameter:SecretParameter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretParameter gets an existing SecretParameter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretParameter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretParameterState, opts ...pulumi.ResourceOption) (*SecretParameter, error) {
	var resource SecretParameter
	err := ctx.ReadResource("alicloud:oos/secretParameter:SecretParameter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretParameter resources.
type secretParameterState struct {
	// The constraints of the encryption parameter. By default, this parameter is null. Valid values:
	Constraints *string `pulumi:"constraints"`
	// The description of the encryption parameter. The description must be `1` to `200` characters in length.
	Description *string `pulumi:"description"`
	// The Customer Master Key (CMK) of Key Management Service (KMS) that is used to encrypt the parameter.
	KeyId *string `pulumi:"keyId"`
	// The ID of the Resource Group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The name of the encryption parameter.  The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
	SecretParameterName *string `pulumi:"secretParameterName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The data type of the encryption parameter. Valid values: `Secret`.
	Type *string `pulumi:"type"`
	// The value of the encryption parameter. The value must be `1` to `4096` characters in length.
	Value *string `pulumi:"value"`
}

type SecretParameterState struct {
	// The constraints of the encryption parameter. By default, this parameter is null. Valid values:
	Constraints pulumi.StringPtrInput
	// The description of the encryption parameter. The description must be `1` to `200` characters in length.
	Description pulumi.StringPtrInput
	// The Customer Master Key (CMK) of Key Management Service (KMS) that is used to encrypt the parameter.
	KeyId pulumi.StringPtrInput
	// The ID of the Resource Group.
	ResourceGroupId pulumi.StringPtrInput
	// The name of the encryption parameter.  The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
	SecretParameterName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The data type of the encryption parameter. Valid values: `Secret`.
	Type pulumi.StringPtrInput
	// The value of the encryption parameter. The value must be `1` to `4096` characters in length.
	Value pulumi.StringPtrInput
}

func (SecretParameterState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretParameterState)(nil)).Elem()
}

type secretParameterArgs struct {
	// The constraints of the encryption parameter. By default, this parameter is null. Valid values:
	Constraints *string `pulumi:"constraints"`
	// The description of the encryption parameter. The description must be `1` to `200` characters in length.
	Description *string `pulumi:"description"`
	// The Customer Master Key (CMK) of Key Management Service (KMS) that is used to encrypt the parameter.
	KeyId *string `pulumi:"keyId"`
	// The ID of the Resource Group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The name of the encryption parameter.  The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
	SecretParameterName string `pulumi:"secretParameterName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The data type of the encryption parameter. Valid values: `Secret`.
	Type *string `pulumi:"type"`
	// The value of the encryption parameter. The value must be `1` to `4096` characters in length.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a SecretParameter resource.
type SecretParameterArgs struct {
	// The constraints of the encryption parameter. By default, this parameter is null. Valid values:
	Constraints pulumi.StringPtrInput
	// The description of the encryption parameter. The description must be `1` to `200` characters in length.
	Description pulumi.StringPtrInput
	// The Customer Master Key (CMK) of Key Management Service (KMS) that is used to encrypt the parameter.
	KeyId pulumi.StringPtrInput
	// The ID of the Resource Group.
	ResourceGroupId pulumi.StringPtrInput
	// The name of the encryption parameter.  The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
	SecretParameterName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The data type of the encryption parameter. Valid values: `Secret`.
	Type pulumi.StringPtrInput
	// The value of the encryption parameter. The value must be `1` to `4096` characters in length.
	Value pulumi.StringInput
}

func (SecretParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretParameterArgs)(nil)).Elem()
}

type SecretParameterInput interface {
	pulumi.Input

	ToSecretParameterOutput() SecretParameterOutput
	ToSecretParameterOutputWithContext(ctx context.Context) SecretParameterOutput
}

func (*SecretParameter) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretParameter)(nil)).Elem()
}

func (i *SecretParameter) ToSecretParameterOutput() SecretParameterOutput {
	return i.ToSecretParameterOutputWithContext(context.Background())
}

func (i *SecretParameter) ToSecretParameterOutputWithContext(ctx context.Context) SecretParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretParameterOutput)
}

// SecretParameterArrayInput is an input type that accepts SecretParameterArray and SecretParameterArrayOutput values.
// You can construct a concrete instance of `SecretParameterArrayInput` via:
//
//	SecretParameterArray{ SecretParameterArgs{...} }
type SecretParameterArrayInput interface {
	pulumi.Input

	ToSecretParameterArrayOutput() SecretParameterArrayOutput
	ToSecretParameterArrayOutputWithContext(context.Context) SecretParameterArrayOutput
}

type SecretParameterArray []SecretParameterInput

func (SecretParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretParameter)(nil)).Elem()
}

func (i SecretParameterArray) ToSecretParameterArrayOutput() SecretParameterArrayOutput {
	return i.ToSecretParameterArrayOutputWithContext(context.Background())
}

func (i SecretParameterArray) ToSecretParameterArrayOutputWithContext(ctx context.Context) SecretParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretParameterArrayOutput)
}

// SecretParameterMapInput is an input type that accepts SecretParameterMap and SecretParameterMapOutput values.
// You can construct a concrete instance of `SecretParameterMapInput` via:
//
//	SecretParameterMap{ "key": SecretParameterArgs{...} }
type SecretParameterMapInput interface {
	pulumi.Input

	ToSecretParameterMapOutput() SecretParameterMapOutput
	ToSecretParameterMapOutputWithContext(context.Context) SecretParameterMapOutput
}

type SecretParameterMap map[string]SecretParameterInput

func (SecretParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretParameter)(nil)).Elem()
}

func (i SecretParameterMap) ToSecretParameterMapOutput() SecretParameterMapOutput {
	return i.ToSecretParameterMapOutputWithContext(context.Background())
}

func (i SecretParameterMap) ToSecretParameterMapOutputWithContext(ctx context.Context) SecretParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretParameterMapOutput)
}

type SecretParameterOutput struct{ *pulumi.OutputState }

func (SecretParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretParameter)(nil)).Elem()
}

func (o SecretParameterOutput) ToSecretParameterOutput() SecretParameterOutput {
	return o
}

func (o SecretParameterOutput) ToSecretParameterOutputWithContext(ctx context.Context) SecretParameterOutput {
	return o
}

// The constraints of the encryption parameter. By default, this parameter is null. Valid values:
func (o SecretParameterOutput) Constraints() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretParameter) pulumi.StringPtrOutput { return v.Constraints }).(pulumi.StringPtrOutput)
}

// The description of the encryption parameter. The description must be `1` to `200` characters in length.
func (o SecretParameterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretParameter) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Customer Master Key (CMK) of Key Management Service (KMS) that is used to encrypt the parameter.
func (o SecretParameterOutput) KeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretParameter) pulumi.StringPtrOutput { return v.KeyId }).(pulumi.StringPtrOutput)
}

// The ID of the Resource Group.
func (o SecretParameterOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretParameter) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The name of the encryption parameter.  The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
func (o SecretParameterOutput) SecretParameterName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretParameter) pulumi.StringOutput { return v.SecretParameterName }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o SecretParameterOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *SecretParameter) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The data type of the encryption parameter. Valid values: `Secret`.
func (o SecretParameterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretParameter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The value of the encryption parameter. The value must be `1` to `4096` characters in length.
func (o SecretParameterOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretParameter) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type SecretParameterArrayOutput struct{ *pulumi.OutputState }

func (SecretParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretParameter)(nil)).Elem()
}

func (o SecretParameterArrayOutput) ToSecretParameterArrayOutput() SecretParameterArrayOutput {
	return o
}

func (o SecretParameterArrayOutput) ToSecretParameterArrayOutputWithContext(ctx context.Context) SecretParameterArrayOutput {
	return o
}

func (o SecretParameterArrayOutput) Index(i pulumi.IntInput) SecretParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretParameter {
		return vs[0].([]*SecretParameter)[vs[1].(int)]
	}).(SecretParameterOutput)
}

type SecretParameterMapOutput struct{ *pulumi.OutputState }

func (SecretParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretParameter)(nil)).Elem()
}

func (o SecretParameterMapOutput) ToSecretParameterMapOutput() SecretParameterMapOutput {
	return o
}

func (o SecretParameterMapOutput) ToSecretParameterMapOutputWithContext(ctx context.Context) SecretParameterMapOutput {
	return o
}

func (o SecretParameterMapOutput) MapIndex(k pulumi.StringInput) SecretParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretParameter {
		return vs[0].(map[string]*SecretParameter)[vs[1].(string)]
	}).(SecretParameterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretParameterInput)(nil)).Elem(), &SecretParameter{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretParameterArrayInput)(nil)).Elem(), SecretParameterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretParameterMapInput)(nil)).Elem(), SecretParameterMap{})
	pulumi.RegisterOutputType(SecretParameterOutput{})
	pulumi.RegisterOutputType(SecretParameterArrayOutput{})
	pulumi.RegisterOutputType(SecretParameterMapOutput{})
}
