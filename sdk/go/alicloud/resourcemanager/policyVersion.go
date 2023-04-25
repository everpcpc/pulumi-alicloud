// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			examplePolicy, err := resourcemanager.NewPolicy(ctx, "examplePolicy", &resourcemanager.PolicyArgs{
//				PolicyName:     pulumi.String("tftest"),
//				PolicyDocument: pulumi.String("		{\n			\"Statement\": [{\n				\"Action\": [\"oss:*\"],\n				\"Effect\": \"Allow\",\n				\"Resource\": [\"acs:oss:*:*:*\"]\n			}],\n			\"Version\": \"1\"\n		}\n"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = resourcemanager.NewPolicyVersion(ctx, "examplePolicyVersion", &resourcemanager.PolicyVersionArgs{
//				PolicyName:     examplePolicy.PolicyName,
//				PolicyDocument: pulumi.String("		{\n			\"Statement\": [{\n				\"Action\": [\"oss:*\"],\n				\"Effect\": \"Allow\",\n				\"Resource\": [\"acs:oss:*:*:myphotos\"]\n			}],\n			\"Version\": \"1\"\n		}\n"),
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
// Resource Manager Policy Version can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:resourcemanager/policyVersion:PolicyVersion example tftest:v2
//
// ```
type PolicyVersion struct {
	pulumi.CustomResourceState

	// Specifies whether to set the policy version as the default version. Default to `false`.
	//
	// Deprecated: Field 'is_default_version' has been deprecated from provider version 1.90.0
	IsDefaultVersion pulumi.BoolPtrOutput `pulumi:"isDefaultVersion"`
	// The content of the policy. The content must be 1 to 2,048 characters in length.
	PolicyDocument pulumi.StringOutput `pulumi:"policyDocument"`
	// The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
}

// NewPolicyVersion registers a new resource with the given unique name, arguments, and options.
func NewPolicyVersion(ctx *pulumi.Context,
	name string, args *PolicyVersionArgs, opts ...pulumi.ResourceOption) (*PolicyVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDocument'")
	}
	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
	}
	var resource PolicyVersion
	err := ctx.RegisterResource("alicloud:resourcemanager/policyVersion:PolicyVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyVersion gets an existing PolicyVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyVersionState, opts ...pulumi.ResourceOption) (*PolicyVersion, error) {
	var resource PolicyVersion
	err := ctx.ReadResource("alicloud:resourcemanager/policyVersion:PolicyVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyVersion resources.
type policyVersionState struct {
	// Specifies whether to set the policy version as the default version. Default to `false`.
	//
	// Deprecated: Field 'is_default_version' has been deprecated from provider version 1.90.0
	IsDefaultVersion *bool `pulumi:"isDefaultVersion"`
	// The content of the policy. The content must be 1 to 2,048 characters in length.
	PolicyDocument *string `pulumi:"policyDocument"`
	// The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName *string `pulumi:"policyName"`
}

type PolicyVersionState struct {
	// Specifies whether to set the policy version as the default version. Default to `false`.
	//
	// Deprecated: Field 'is_default_version' has been deprecated from provider version 1.90.0
	IsDefaultVersion pulumi.BoolPtrInput
	// The content of the policy. The content must be 1 to 2,048 characters in length.
	PolicyDocument pulumi.StringPtrInput
	// The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName pulumi.StringPtrInput
}

func (PolicyVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyVersionState)(nil)).Elem()
}

type policyVersionArgs struct {
	// Specifies whether to set the policy version as the default version. Default to `false`.
	//
	// Deprecated: Field 'is_default_version' has been deprecated from provider version 1.90.0
	IsDefaultVersion *bool `pulumi:"isDefaultVersion"`
	// The content of the policy. The content must be 1 to 2,048 characters in length.
	PolicyDocument string `pulumi:"policyDocument"`
	// The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName string `pulumi:"policyName"`
}

// The set of arguments for constructing a PolicyVersion resource.
type PolicyVersionArgs struct {
	// Specifies whether to set the policy version as the default version. Default to `false`.
	//
	// Deprecated: Field 'is_default_version' has been deprecated from provider version 1.90.0
	IsDefaultVersion pulumi.BoolPtrInput
	// The content of the policy. The content must be 1 to 2,048 characters in length.
	PolicyDocument pulumi.StringInput
	// The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName pulumi.StringInput
}

func (PolicyVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyVersionArgs)(nil)).Elem()
}

type PolicyVersionInput interface {
	pulumi.Input

	ToPolicyVersionOutput() PolicyVersionOutput
	ToPolicyVersionOutputWithContext(ctx context.Context) PolicyVersionOutput
}

func (*PolicyVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyVersion)(nil)).Elem()
}

func (i *PolicyVersion) ToPolicyVersionOutput() PolicyVersionOutput {
	return i.ToPolicyVersionOutputWithContext(context.Background())
}

func (i *PolicyVersion) ToPolicyVersionOutputWithContext(ctx context.Context) PolicyVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyVersionOutput)
}

// PolicyVersionArrayInput is an input type that accepts PolicyVersionArray and PolicyVersionArrayOutput values.
// You can construct a concrete instance of `PolicyVersionArrayInput` via:
//
//	PolicyVersionArray{ PolicyVersionArgs{...} }
type PolicyVersionArrayInput interface {
	pulumi.Input

	ToPolicyVersionArrayOutput() PolicyVersionArrayOutput
	ToPolicyVersionArrayOutputWithContext(context.Context) PolicyVersionArrayOutput
}

type PolicyVersionArray []PolicyVersionInput

func (PolicyVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyVersion)(nil)).Elem()
}

func (i PolicyVersionArray) ToPolicyVersionArrayOutput() PolicyVersionArrayOutput {
	return i.ToPolicyVersionArrayOutputWithContext(context.Background())
}

func (i PolicyVersionArray) ToPolicyVersionArrayOutputWithContext(ctx context.Context) PolicyVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyVersionArrayOutput)
}

// PolicyVersionMapInput is an input type that accepts PolicyVersionMap and PolicyVersionMapOutput values.
// You can construct a concrete instance of `PolicyVersionMapInput` via:
//
//	PolicyVersionMap{ "key": PolicyVersionArgs{...} }
type PolicyVersionMapInput interface {
	pulumi.Input

	ToPolicyVersionMapOutput() PolicyVersionMapOutput
	ToPolicyVersionMapOutputWithContext(context.Context) PolicyVersionMapOutput
}

type PolicyVersionMap map[string]PolicyVersionInput

func (PolicyVersionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyVersion)(nil)).Elem()
}

func (i PolicyVersionMap) ToPolicyVersionMapOutput() PolicyVersionMapOutput {
	return i.ToPolicyVersionMapOutputWithContext(context.Background())
}

func (i PolicyVersionMap) ToPolicyVersionMapOutputWithContext(ctx context.Context) PolicyVersionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyVersionMapOutput)
}

type PolicyVersionOutput struct{ *pulumi.OutputState }

func (PolicyVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyVersion)(nil)).Elem()
}

func (o PolicyVersionOutput) ToPolicyVersionOutput() PolicyVersionOutput {
	return o
}

func (o PolicyVersionOutput) ToPolicyVersionOutputWithContext(ctx context.Context) PolicyVersionOutput {
	return o
}

// Specifies whether to set the policy version as the default version. Default to `false`.
//
// Deprecated: Field 'is_default_version' has been deprecated from provider version 1.90.0
func (o PolicyVersionOutput) IsDefaultVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyVersion) pulumi.BoolPtrOutput { return v.IsDefaultVersion }).(pulumi.BoolPtrOutput)
}

// The content of the policy. The content must be 1 to 2,048 characters in length.
func (o PolicyVersionOutput) PolicyDocument() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyVersion) pulumi.StringOutput { return v.PolicyDocument }).(pulumi.StringOutput)
}

// The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
func (o PolicyVersionOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyVersion) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

type PolicyVersionArrayOutput struct{ *pulumi.OutputState }

func (PolicyVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyVersion)(nil)).Elem()
}

func (o PolicyVersionArrayOutput) ToPolicyVersionArrayOutput() PolicyVersionArrayOutput {
	return o
}

func (o PolicyVersionArrayOutput) ToPolicyVersionArrayOutputWithContext(ctx context.Context) PolicyVersionArrayOutput {
	return o
}

func (o PolicyVersionArrayOutput) Index(i pulumi.IntInput) PolicyVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyVersion {
		return vs[0].([]*PolicyVersion)[vs[1].(int)]
	}).(PolicyVersionOutput)
}

type PolicyVersionMapOutput struct{ *pulumi.OutputState }

func (PolicyVersionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyVersion)(nil)).Elem()
}

func (o PolicyVersionMapOutput) ToPolicyVersionMapOutput() PolicyVersionMapOutput {
	return o
}

func (o PolicyVersionMapOutput) ToPolicyVersionMapOutputWithContext(ctx context.Context) PolicyVersionMapOutput {
	return o
}

func (o PolicyVersionMapOutput) MapIndex(k pulumi.StringInput) PolicyVersionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyVersion {
		return vs[0].(map[string]*PolicyVersion)[vs[1].(string)]
	}).(PolicyVersionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyVersionInput)(nil)).Elem(), &PolicyVersion{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyVersionArrayInput)(nil)).Elem(), PolicyVersionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyVersionMapInput)(nil)).Elem(), PolicyVersionMap{})
	pulumi.RegisterOutputType(PolicyVersionOutput{})
	pulumi.RegisterOutputType(PolicyVersionArrayOutput{})
	pulumi.RegisterOutputType(PolicyVersionMapOutput{})
}
