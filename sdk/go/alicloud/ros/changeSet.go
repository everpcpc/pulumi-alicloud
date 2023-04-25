// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ros

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ROS Change Set resource.
//
// For information about ROS Change Set and how to use it, see [What is Change Set](https://www.alibabacloud.com/help/doc-detail/131051.htm).
//
// > **NOTE:** Available in v1.105.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ros"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ros.NewChangeSet(ctx, "example", &ros.ChangeSetArgs{
//				ChangeSetName: pulumi.String("example_value"),
//				ChangeSetType: pulumi.String("CREATE"),
//				Description:   pulumi.String("Test From Terraform"),
//				StackName:     pulumi.String("tf-testacc"),
//				TemplateBody:  pulumi.String("{\"ROSTemplateFormatVersion\":\"2015-09-01\"}"),
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
// ROS Change Set can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ros/changeSet:ChangeSet example <change_set_id>
//
// ```
type ChangeSet struct {
	pulumi.CustomResourceState

	// The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	ChangeSetName pulumi.StringOutput `pulumi:"changeSetName"`
	// The type of the change set. Valid values:  CREATE: creates a change set for a new stack. UPDATE: creates a change set for an existing stack. IMPORT: creates a change set for a new stack or an existing stack to import non-ROS-managed resources. If you create a change set for a new stack, ROS creates a stack that has a unique stack ID. The stack is in the REVIEW_IN_PROGRESS state until you execute the change set.  You cannot use the UPDATE type to create a change set for a new stack or the CREATE type to create a change set for an existing stack.
	ChangeSetType pulumi.StringPtrOutput `pulumi:"changeSetType"`
	// The description of the change set. The description can be up to 1,024 bytes in length.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether to disable rollback on stack creation failure. Default value: false.  Valid values:  true: disables rollback on stack creation failure. false: enables rollback on stack creation failure. Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
	DisableRollback pulumi.BoolPtrOutput `pulumi:"disableRollback"`
	// The notification urls.
	NotificationUrls pulumi.StringArrayOutput `pulumi:"notificationUrls"`
	// Parameters.
	Parameters ChangeSetParameterArrayOutput `pulumi:"parameters"`
	// The ram role name.
	RamRoleName pulumi.StringPtrOutput `pulumi:"ramRoleName"`
	// The replacement option.
	ReplacementOption pulumi.StringPtrOutput `pulumi:"replacementOption"`
	// The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
	StackId pulumi.StringOutput `pulumi:"stackId"`
	// The name of the stack for which you want to create the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.  Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
	StackName pulumi.StringPtrOutput `pulumi:"stackName"`
	// The stack policy body.
	StackPolicyBody pulumi.StringPtrOutput `pulumi:"stackPolicyBody"`
	// The stack policy during update body.
	StackPolicyDuringUpdateBody pulumi.StringPtrOutput `pulumi:"stackPolicyDuringUpdateBody"`
	// The stack policy during update url.
	StackPolicyDuringUpdateUrl pulumi.StringPtrOutput `pulumi:"stackPolicyDuringUpdateUrl"`
	// The stack policy url.
	StackPolicyUrl pulumi.StringPtrOutput `pulumi:"stackPolicyUrl"`
	// The status of the change set.
	Status pulumi.StringOutput `pulumi:"status"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You can specify one of TemplateBody or TemplateURL parameters, but you cannot specify both of them.
	TemplateBody pulumi.StringPtrOutput `pulumi:"templateBody"`
	// The template url.
	TemplateUrl pulumi.StringPtrOutput `pulumi:"templateUrl"`
	// Timeout In Minutes.
	TimeoutInMinutes pulumi.IntOutput `pulumi:"timeoutInMinutes"`
	// The use previous parameters.
	UsePreviousParameters pulumi.BoolPtrOutput `pulumi:"usePreviousParameters"`
}

// NewChangeSet registers a new resource with the given unique name, arguments, and options.
func NewChangeSet(ctx *pulumi.Context,
	name string, args *ChangeSetArgs, opts ...pulumi.ResourceOption) (*ChangeSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ChangeSetName == nil {
		return nil, errors.New("invalid value for required argument 'ChangeSetName'")
	}
	var resource ChangeSet
	err := ctx.RegisterResource("alicloud:ros/changeSet:ChangeSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChangeSet gets an existing ChangeSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChangeSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChangeSetState, opts ...pulumi.ResourceOption) (*ChangeSet, error) {
	var resource ChangeSet
	err := ctx.ReadResource("alicloud:ros/changeSet:ChangeSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChangeSet resources.
type changeSetState struct {
	// The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	ChangeSetName *string `pulumi:"changeSetName"`
	// The type of the change set. Valid values:  CREATE: creates a change set for a new stack. UPDATE: creates a change set for an existing stack. IMPORT: creates a change set for a new stack or an existing stack to import non-ROS-managed resources. If you create a change set for a new stack, ROS creates a stack that has a unique stack ID. The stack is in the REVIEW_IN_PROGRESS state until you execute the change set.  You cannot use the UPDATE type to create a change set for a new stack or the CREATE type to create a change set for an existing stack.
	ChangeSetType *string `pulumi:"changeSetType"`
	// The description of the change set. The description can be up to 1,024 bytes in length.
	Description *string `pulumi:"description"`
	// Specifies whether to disable rollback on stack creation failure. Default value: false.  Valid values:  true: disables rollback on stack creation failure. false: enables rollback on stack creation failure. Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
	DisableRollback *bool `pulumi:"disableRollback"`
	// The notification urls.
	NotificationUrls []string `pulumi:"notificationUrls"`
	// Parameters.
	Parameters []ChangeSetParameter `pulumi:"parameters"`
	// The ram role name.
	RamRoleName *string `pulumi:"ramRoleName"`
	// The replacement option.
	ReplacementOption *string `pulumi:"replacementOption"`
	// The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
	StackId *string `pulumi:"stackId"`
	// The name of the stack for which you want to create the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.  Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
	StackName *string `pulumi:"stackName"`
	// The stack policy body.
	StackPolicyBody *string `pulumi:"stackPolicyBody"`
	// The stack policy during update body.
	StackPolicyDuringUpdateBody *string `pulumi:"stackPolicyDuringUpdateBody"`
	// The stack policy during update url.
	StackPolicyDuringUpdateUrl *string `pulumi:"stackPolicyDuringUpdateUrl"`
	// The stack policy url.
	StackPolicyUrl *string `pulumi:"stackPolicyUrl"`
	// The status of the change set.
	Status *string `pulumi:"status"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You can specify one of TemplateBody or TemplateURL parameters, but you cannot specify both of them.
	TemplateBody *string `pulumi:"templateBody"`
	// The template url.
	TemplateUrl *string `pulumi:"templateUrl"`
	// Timeout In Minutes.
	TimeoutInMinutes *int `pulumi:"timeoutInMinutes"`
	// The use previous parameters.
	UsePreviousParameters *bool `pulumi:"usePreviousParameters"`
}

type ChangeSetState struct {
	// The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	ChangeSetName pulumi.StringPtrInput
	// The type of the change set. Valid values:  CREATE: creates a change set for a new stack. UPDATE: creates a change set for an existing stack. IMPORT: creates a change set for a new stack or an existing stack to import non-ROS-managed resources. If you create a change set for a new stack, ROS creates a stack that has a unique stack ID. The stack is in the REVIEW_IN_PROGRESS state until you execute the change set.  You cannot use the UPDATE type to create a change set for a new stack or the CREATE type to create a change set for an existing stack.
	ChangeSetType pulumi.StringPtrInput
	// The description of the change set. The description can be up to 1,024 bytes in length.
	Description pulumi.StringPtrInput
	// Specifies whether to disable rollback on stack creation failure. Default value: false.  Valid values:  true: disables rollback on stack creation failure. false: enables rollback on stack creation failure. Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
	DisableRollback pulumi.BoolPtrInput
	// The notification urls.
	NotificationUrls pulumi.StringArrayInput
	// Parameters.
	Parameters ChangeSetParameterArrayInput
	// The ram role name.
	RamRoleName pulumi.StringPtrInput
	// The replacement option.
	ReplacementOption pulumi.StringPtrInput
	// The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
	StackId pulumi.StringPtrInput
	// The name of the stack for which you want to create the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.  Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
	StackName pulumi.StringPtrInput
	// The stack policy body.
	StackPolicyBody pulumi.StringPtrInput
	// The stack policy during update body.
	StackPolicyDuringUpdateBody pulumi.StringPtrInput
	// The stack policy during update url.
	StackPolicyDuringUpdateUrl pulumi.StringPtrInput
	// The stack policy url.
	StackPolicyUrl pulumi.StringPtrInput
	// The status of the change set.
	Status pulumi.StringPtrInput
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You can specify one of TemplateBody or TemplateURL parameters, but you cannot specify both of them.
	TemplateBody pulumi.StringPtrInput
	// The template url.
	TemplateUrl pulumi.StringPtrInput
	// Timeout In Minutes.
	TimeoutInMinutes pulumi.IntPtrInput
	// The use previous parameters.
	UsePreviousParameters pulumi.BoolPtrInput
}

func (ChangeSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*changeSetState)(nil)).Elem()
}

type changeSetArgs struct {
	// The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	ChangeSetName string `pulumi:"changeSetName"`
	// The type of the change set. Valid values:  CREATE: creates a change set for a new stack. UPDATE: creates a change set for an existing stack. IMPORT: creates a change set for a new stack or an existing stack to import non-ROS-managed resources. If you create a change set for a new stack, ROS creates a stack that has a unique stack ID. The stack is in the REVIEW_IN_PROGRESS state until you execute the change set.  You cannot use the UPDATE type to create a change set for a new stack or the CREATE type to create a change set for an existing stack.
	ChangeSetType *string `pulumi:"changeSetType"`
	// The description of the change set. The description can be up to 1,024 bytes in length.
	Description *string `pulumi:"description"`
	// Specifies whether to disable rollback on stack creation failure. Default value: false.  Valid values:  true: disables rollback on stack creation failure. false: enables rollback on stack creation failure. Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
	DisableRollback *bool `pulumi:"disableRollback"`
	// The notification urls.
	NotificationUrls []string `pulumi:"notificationUrls"`
	// Parameters.
	Parameters []ChangeSetParameter `pulumi:"parameters"`
	// The ram role name.
	RamRoleName *string `pulumi:"ramRoleName"`
	// The replacement option.
	ReplacementOption *string `pulumi:"replacementOption"`
	// The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
	StackId *string `pulumi:"stackId"`
	// The name of the stack for which you want to create the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.  Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
	StackName *string `pulumi:"stackName"`
	// The stack policy body.
	StackPolicyBody *string `pulumi:"stackPolicyBody"`
	// The stack policy during update body.
	StackPolicyDuringUpdateBody *string `pulumi:"stackPolicyDuringUpdateBody"`
	// The stack policy during update url.
	StackPolicyDuringUpdateUrl *string `pulumi:"stackPolicyDuringUpdateUrl"`
	// The stack policy url.
	StackPolicyUrl *string `pulumi:"stackPolicyUrl"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You can specify one of TemplateBody or TemplateURL parameters, but you cannot specify both of them.
	TemplateBody *string `pulumi:"templateBody"`
	// The template url.
	TemplateUrl *string `pulumi:"templateUrl"`
	// Timeout In Minutes.
	TimeoutInMinutes *int `pulumi:"timeoutInMinutes"`
	// The use previous parameters.
	UsePreviousParameters *bool `pulumi:"usePreviousParameters"`
}

// The set of arguments for constructing a ChangeSet resource.
type ChangeSetArgs struct {
	// The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	ChangeSetName pulumi.StringInput
	// The type of the change set. Valid values:  CREATE: creates a change set for a new stack. UPDATE: creates a change set for an existing stack. IMPORT: creates a change set for a new stack or an existing stack to import non-ROS-managed resources. If you create a change set for a new stack, ROS creates a stack that has a unique stack ID. The stack is in the REVIEW_IN_PROGRESS state until you execute the change set.  You cannot use the UPDATE type to create a change set for a new stack or the CREATE type to create a change set for an existing stack.
	ChangeSetType pulumi.StringPtrInput
	// The description of the change set. The description can be up to 1,024 bytes in length.
	Description pulumi.StringPtrInput
	// Specifies whether to disable rollback on stack creation failure. Default value: false.  Valid values:  true: disables rollback on stack creation failure. false: enables rollback on stack creation failure. Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
	DisableRollback pulumi.BoolPtrInput
	// The notification urls.
	NotificationUrls pulumi.StringArrayInput
	// Parameters.
	Parameters ChangeSetParameterArrayInput
	// The ram role name.
	RamRoleName pulumi.StringPtrInput
	// The replacement option.
	ReplacementOption pulumi.StringPtrInput
	// The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
	StackId pulumi.StringPtrInput
	// The name of the stack for which you want to create the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.  Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
	StackName pulumi.StringPtrInput
	// The stack policy body.
	StackPolicyBody pulumi.StringPtrInput
	// The stack policy during update body.
	StackPolicyDuringUpdateBody pulumi.StringPtrInput
	// The stack policy during update url.
	StackPolicyDuringUpdateUrl pulumi.StringPtrInput
	// The stack policy url.
	StackPolicyUrl pulumi.StringPtrInput
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You can specify one of TemplateBody or TemplateURL parameters, but you cannot specify both of them.
	TemplateBody pulumi.StringPtrInput
	// The template url.
	TemplateUrl pulumi.StringPtrInput
	// Timeout In Minutes.
	TimeoutInMinutes pulumi.IntPtrInput
	// The use previous parameters.
	UsePreviousParameters pulumi.BoolPtrInput
}

func (ChangeSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*changeSetArgs)(nil)).Elem()
}

type ChangeSetInput interface {
	pulumi.Input

	ToChangeSetOutput() ChangeSetOutput
	ToChangeSetOutputWithContext(ctx context.Context) ChangeSetOutput
}

func (*ChangeSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ChangeSet)(nil)).Elem()
}

func (i *ChangeSet) ToChangeSetOutput() ChangeSetOutput {
	return i.ToChangeSetOutputWithContext(context.Background())
}

func (i *ChangeSet) ToChangeSetOutputWithContext(ctx context.Context) ChangeSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChangeSetOutput)
}

// ChangeSetArrayInput is an input type that accepts ChangeSetArray and ChangeSetArrayOutput values.
// You can construct a concrete instance of `ChangeSetArrayInput` via:
//
//	ChangeSetArray{ ChangeSetArgs{...} }
type ChangeSetArrayInput interface {
	pulumi.Input

	ToChangeSetArrayOutput() ChangeSetArrayOutput
	ToChangeSetArrayOutputWithContext(context.Context) ChangeSetArrayOutput
}

type ChangeSetArray []ChangeSetInput

func (ChangeSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ChangeSet)(nil)).Elem()
}

func (i ChangeSetArray) ToChangeSetArrayOutput() ChangeSetArrayOutput {
	return i.ToChangeSetArrayOutputWithContext(context.Background())
}

func (i ChangeSetArray) ToChangeSetArrayOutputWithContext(ctx context.Context) ChangeSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChangeSetArrayOutput)
}

// ChangeSetMapInput is an input type that accepts ChangeSetMap and ChangeSetMapOutput values.
// You can construct a concrete instance of `ChangeSetMapInput` via:
//
//	ChangeSetMap{ "key": ChangeSetArgs{...} }
type ChangeSetMapInput interface {
	pulumi.Input

	ToChangeSetMapOutput() ChangeSetMapOutput
	ToChangeSetMapOutputWithContext(context.Context) ChangeSetMapOutput
}

type ChangeSetMap map[string]ChangeSetInput

func (ChangeSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ChangeSet)(nil)).Elem()
}

func (i ChangeSetMap) ToChangeSetMapOutput() ChangeSetMapOutput {
	return i.ToChangeSetMapOutputWithContext(context.Background())
}

func (i ChangeSetMap) ToChangeSetMapOutputWithContext(ctx context.Context) ChangeSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChangeSetMapOutput)
}

type ChangeSetOutput struct{ *pulumi.OutputState }

func (ChangeSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChangeSet)(nil)).Elem()
}

func (o ChangeSetOutput) ToChangeSetOutput() ChangeSetOutput {
	return o
}

func (o ChangeSetOutput) ToChangeSetOutputWithContext(ctx context.Context) ChangeSetOutput {
	return o
}

// The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
func (o ChangeSetOutput) ChangeSetName() pulumi.StringOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.StringOutput { return v.ChangeSetName }).(pulumi.StringOutput)
}

// The type of the change set. Valid values:  CREATE: creates a change set for a new stack. UPDATE: creates a change set for an existing stack. IMPORT: creates a change set for a new stack or an existing stack to import non-ROS-managed resources. If you create a change set for a new stack, ROS creates a stack that has a unique stack ID. The stack is in the REVIEW_IN_PROGRESS state until you execute the change set.  You cannot use the UPDATE type to create a change set for a new stack or the CREATE type to create a change set for an existing stack.
func (o ChangeSetOutput) ChangeSetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.StringPtrOutput { return v.ChangeSetType }).(pulumi.StringPtrOutput)
}

// The description of the change set. The description can be up to 1,024 bytes in length.
func (o ChangeSetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies whether to disable rollback on stack creation failure. Default value: false.  Valid values:  true: disables rollback on stack creation failure. false: enables rollback on stack creation failure. Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
func (o ChangeSetOutput) DisableRollback() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.BoolPtrOutput { return v.DisableRollback }).(pulumi.BoolPtrOutput)
}

// The notification urls.
func (o ChangeSetOutput) NotificationUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.StringArrayOutput { return v.NotificationUrls }).(pulumi.StringArrayOutput)
}

// Parameters.
func (o ChangeSetOutput) Parameters() ChangeSetParameterArrayOutput {
	return o.ApplyT(func(v *ChangeSet) ChangeSetParameterArrayOutput { return v.Parameters }).(ChangeSetParameterArrayOutput)
}

// The ram role name.
func (o ChangeSetOutput) RamRoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.StringPtrOutput { return v.RamRoleName }).(pulumi.StringPtrOutput)
}

// The replacement option.
func (o ChangeSetOutput) ReplacementOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.StringPtrOutput { return v.ReplacementOption }).(pulumi.StringPtrOutput)
}

// The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
func (o ChangeSetOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.StringOutput { return v.StackId }).(pulumi.StringOutput)
}

// The name of the stack for which you want to create the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.  Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
func (o ChangeSetOutput) StackName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.StringPtrOutput { return v.StackName }).(pulumi.StringPtrOutput)
}

// The stack policy body.
func (o ChangeSetOutput) StackPolicyBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.StringPtrOutput { return v.StackPolicyBody }).(pulumi.StringPtrOutput)
}

// The stack policy during update body.
func (o ChangeSetOutput) StackPolicyDuringUpdateBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.StringPtrOutput { return v.StackPolicyDuringUpdateBody }).(pulumi.StringPtrOutput)
}

// The stack policy during update url.
func (o ChangeSetOutput) StackPolicyDuringUpdateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.StringPtrOutput { return v.StackPolicyDuringUpdateUrl }).(pulumi.StringPtrOutput)
}

// The stack policy url.
func (o ChangeSetOutput) StackPolicyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.StringPtrOutput { return v.StackPolicyUrl }).(pulumi.StringPtrOutput)
}

// The status of the change set.
func (o ChangeSetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You can specify one of TemplateBody or TemplateURL parameters, but you cannot specify both of them.
func (o ChangeSetOutput) TemplateBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.StringPtrOutput { return v.TemplateBody }).(pulumi.StringPtrOutput)
}

// The template url.
func (o ChangeSetOutput) TemplateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.StringPtrOutput { return v.TemplateUrl }).(pulumi.StringPtrOutput)
}

// Timeout In Minutes.
func (o ChangeSetOutput) TimeoutInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.IntOutput { return v.TimeoutInMinutes }).(pulumi.IntOutput)
}

// The use previous parameters.
func (o ChangeSetOutput) UsePreviousParameters() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ChangeSet) pulumi.BoolPtrOutput { return v.UsePreviousParameters }).(pulumi.BoolPtrOutput)
}

type ChangeSetArrayOutput struct{ *pulumi.OutputState }

func (ChangeSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ChangeSet)(nil)).Elem()
}

func (o ChangeSetArrayOutput) ToChangeSetArrayOutput() ChangeSetArrayOutput {
	return o
}

func (o ChangeSetArrayOutput) ToChangeSetArrayOutputWithContext(ctx context.Context) ChangeSetArrayOutput {
	return o
}

func (o ChangeSetArrayOutput) Index(i pulumi.IntInput) ChangeSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ChangeSet {
		return vs[0].([]*ChangeSet)[vs[1].(int)]
	}).(ChangeSetOutput)
}

type ChangeSetMapOutput struct{ *pulumi.OutputState }

func (ChangeSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ChangeSet)(nil)).Elem()
}

func (o ChangeSetMapOutput) ToChangeSetMapOutput() ChangeSetMapOutput {
	return o
}

func (o ChangeSetMapOutput) ToChangeSetMapOutputWithContext(ctx context.Context) ChangeSetMapOutput {
	return o
}

func (o ChangeSetMapOutput) MapIndex(k pulumi.StringInput) ChangeSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ChangeSet {
		return vs[0].(map[string]*ChangeSet)[vs[1].(string)]
	}).(ChangeSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChangeSetInput)(nil)).Elem(), &ChangeSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChangeSetArrayInput)(nil)).Elem(), ChangeSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChangeSetMapInput)(nil)).Elem(), ChangeSetMap{})
	pulumi.RegisterOutputType(ChangeSetOutput{})
	pulumi.RegisterOutputType(ChangeSetArrayOutput{})
	pulumi.RegisterOutputType(ChangeSetMapOutput{})
}
