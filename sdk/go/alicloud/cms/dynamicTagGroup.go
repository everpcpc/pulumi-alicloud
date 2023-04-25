// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Monitor Service Dynamic Tag Group resource.
//
// For information about Cloud Monitor Service Dynamic Tag Group and how to use it, see [What is Dynamic Tag Group](https://www.alibabacloud.com/help/doc-detail/150123.html).
//
// > **NOTE:** Available in v1.142.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultAlarmContactGroup, err := cms.NewAlarmContactGroup(ctx, "defaultAlarmContactGroup", &cms.AlarmContactGroupArgs{
//				AlarmContactGroupName: pulumi.String("example_value"),
//				Describe:              pulumi.String("example_value"),
//				EnableSubscribed:      pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cms.NewDynamicTagGroup(ctx, "defaultDynamicTagGroup", &cms.DynamicTagGroupArgs{
//				ContactGroupLists: pulumi.StringArray{
//					defaultAlarmContactGroup.ID(),
//				},
//				TagKey: pulumi.String("your_tag_key"),
//				MatchExpresses: cms.DynamicTagGroupMatchExpressArray{
//					&cms.DynamicTagGroupMatchExpressArgs{
//						TagValue:              pulumi.String("your_tag_value"),
//						TagValueMatchFunction: pulumi.String("all"),
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
// Cloud Monitor Service Dynamic Tag Group can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cms/dynamicTagGroup:DynamicTagGroup example <id>
//
// ```
type DynamicTagGroup struct {
	pulumi.CustomResourceState

	// Alarm contact group. The value range of N is 1~100. The alarm notification of the application group is sent to the alarm contact in the alarm contact group.
	ContactGroupLists pulumi.StringArrayOutput `pulumi:"contactGroupLists"`
	// The relationship between conditional expressions. Valid values: `and`, `or`.
	MatchExpressFilterRelation pulumi.StringOutput `pulumi:"matchExpressFilterRelation"`
	// The label generates a matching expression that applies the grouping. See the following `Block matchExpress`.
	MatchExpresses DynamicTagGroupMatchExpressArrayOutput `pulumi:"matchExpresses"`
	// The status of the resource. Valid values: `RUNNING`, `FINISH`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tag key of the tag.
	TagKey pulumi.StringOutput `pulumi:"tagKey"`
	// Alarm template ID list.
	TemplateIdLists pulumi.StringArrayOutput `pulumi:"templateIdLists"`
}

// NewDynamicTagGroup registers a new resource with the given unique name, arguments, and options.
func NewDynamicTagGroup(ctx *pulumi.Context,
	name string, args *DynamicTagGroupArgs, opts ...pulumi.ResourceOption) (*DynamicTagGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContactGroupLists == nil {
		return nil, errors.New("invalid value for required argument 'ContactGroupLists'")
	}
	if args.MatchExpresses == nil {
		return nil, errors.New("invalid value for required argument 'MatchExpresses'")
	}
	if args.TagKey == nil {
		return nil, errors.New("invalid value for required argument 'TagKey'")
	}
	var resource DynamicTagGroup
	err := ctx.RegisterResource("alicloud:cms/dynamicTagGroup:DynamicTagGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDynamicTagGroup gets an existing DynamicTagGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDynamicTagGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DynamicTagGroupState, opts ...pulumi.ResourceOption) (*DynamicTagGroup, error) {
	var resource DynamicTagGroup
	err := ctx.ReadResource("alicloud:cms/dynamicTagGroup:DynamicTagGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DynamicTagGroup resources.
type dynamicTagGroupState struct {
	// Alarm contact group. The value range of N is 1~100. The alarm notification of the application group is sent to the alarm contact in the alarm contact group.
	ContactGroupLists []string `pulumi:"contactGroupLists"`
	// The relationship between conditional expressions. Valid values: `and`, `or`.
	MatchExpressFilterRelation *string `pulumi:"matchExpressFilterRelation"`
	// The label generates a matching expression that applies the grouping. See the following `Block matchExpress`.
	MatchExpresses []DynamicTagGroupMatchExpress `pulumi:"matchExpresses"`
	// The status of the resource. Valid values: `RUNNING`, `FINISH`.
	Status *string `pulumi:"status"`
	// The tag key of the tag.
	TagKey *string `pulumi:"tagKey"`
	// Alarm template ID list.
	TemplateIdLists []string `pulumi:"templateIdLists"`
}

type DynamicTagGroupState struct {
	// Alarm contact group. The value range of N is 1~100. The alarm notification of the application group is sent to the alarm contact in the alarm contact group.
	ContactGroupLists pulumi.StringArrayInput
	// The relationship between conditional expressions. Valid values: `and`, `or`.
	MatchExpressFilterRelation pulumi.StringPtrInput
	// The label generates a matching expression that applies the grouping. See the following `Block matchExpress`.
	MatchExpresses DynamicTagGroupMatchExpressArrayInput
	// The status of the resource. Valid values: `RUNNING`, `FINISH`.
	Status pulumi.StringPtrInput
	// The tag key of the tag.
	TagKey pulumi.StringPtrInput
	// Alarm template ID list.
	TemplateIdLists pulumi.StringArrayInput
}

func (DynamicTagGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dynamicTagGroupState)(nil)).Elem()
}

type dynamicTagGroupArgs struct {
	// Alarm contact group. The value range of N is 1~100. The alarm notification of the application group is sent to the alarm contact in the alarm contact group.
	ContactGroupLists []string `pulumi:"contactGroupLists"`
	// The relationship between conditional expressions. Valid values: `and`, `or`.
	MatchExpressFilterRelation *string `pulumi:"matchExpressFilterRelation"`
	// The label generates a matching expression that applies the grouping. See the following `Block matchExpress`.
	MatchExpresses []DynamicTagGroupMatchExpress `pulumi:"matchExpresses"`
	// The tag key of the tag.
	TagKey string `pulumi:"tagKey"`
	// Alarm template ID list.
	TemplateIdLists []string `pulumi:"templateIdLists"`
}

// The set of arguments for constructing a DynamicTagGroup resource.
type DynamicTagGroupArgs struct {
	// Alarm contact group. The value range of N is 1~100. The alarm notification of the application group is sent to the alarm contact in the alarm contact group.
	ContactGroupLists pulumi.StringArrayInput
	// The relationship between conditional expressions. Valid values: `and`, `or`.
	MatchExpressFilterRelation pulumi.StringPtrInput
	// The label generates a matching expression that applies the grouping. See the following `Block matchExpress`.
	MatchExpresses DynamicTagGroupMatchExpressArrayInput
	// The tag key of the tag.
	TagKey pulumi.StringInput
	// Alarm template ID list.
	TemplateIdLists pulumi.StringArrayInput
}

func (DynamicTagGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dynamicTagGroupArgs)(nil)).Elem()
}

type DynamicTagGroupInput interface {
	pulumi.Input

	ToDynamicTagGroupOutput() DynamicTagGroupOutput
	ToDynamicTagGroupOutputWithContext(ctx context.Context) DynamicTagGroupOutput
}

func (*DynamicTagGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DynamicTagGroup)(nil)).Elem()
}

func (i *DynamicTagGroup) ToDynamicTagGroupOutput() DynamicTagGroupOutput {
	return i.ToDynamicTagGroupOutputWithContext(context.Background())
}

func (i *DynamicTagGroup) ToDynamicTagGroupOutputWithContext(ctx context.Context) DynamicTagGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamicTagGroupOutput)
}

// DynamicTagGroupArrayInput is an input type that accepts DynamicTagGroupArray and DynamicTagGroupArrayOutput values.
// You can construct a concrete instance of `DynamicTagGroupArrayInput` via:
//
//	DynamicTagGroupArray{ DynamicTagGroupArgs{...} }
type DynamicTagGroupArrayInput interface {
	pulumi.Input

	ToDynamicTagGroupArrayOutput() DynamicTagGroupArrayOutput
	ToDynamicTagGroupArrayOutputWithContext(context.Context) DynamicTagGroupArrayOutput
}

type DynamicTagGroupArray []DynamicTagGroupInput

func (DynamicTagGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DynamicTagGroup)(nil)).Elem()
}

func (i DynamicTagGroupArray) ToDynamicTagGroupArrayOutput() DynamicTagGroupArrayOutput {
	return i.ToDynamicTagGroupArrayOutputWithContext(context.Background())
}

func (i DynamicTagGroupArray) ToDynamicTagGroupArrayOutputWithContext(ctx context.Context) DynamicTagGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamicTagGroupArrayOutput)
}

// DynamicTagGroupMapInput is an input type that accepts DynamicTagGroupMap and DynamicTagGroupMapOutput values.
// You can construct a concrete instance of `DynamicTagGroupMapInput` via:
//
//	DynamicTagGroupMap{ "key": DynamicTagGroupArgs{...} }
type DynamicTagGroupMapInput interface {
	pulumi.Input

	ToDynamicTagGroupMapOutput() DynamicTagGroupMapOutput
	ToDynamicTagGroupMapOutputWithContext(context.Context) DynamicTagGroupMapOutput
}

type DynamicTagGroupMap map[string]DynamicTagGroupInput

func (DynamicTagGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DynamicTagGroup)(nil)).Elem()
}

func (i DynamicTagGroupMap) ToDynamicTagGroupMapOutput() DynamicTagGroupMapOutput {
	return i.ToDynamicTagGroupMapOutputWithContext(context.Background())
}

func (i DynamicTagGroupMap) ToDynamicTagGroupMapOutputWithContext(ctx context.Context) DynamicTagGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamicTagGroupMapOutput)
}

type DynamicTagGroupOutput struct{ *pulumi.OutputState }

func (DynamicTagGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DynamicTagGroup)(nil)).Elem()
}

func (o DynamicTagGroupOutput) ToDynamicTagGroupOutput() DynamicTagGroupOutput {
	return o
}

func (o DynamicTagGroupOutput) ToDynamicTagGroupOutputWithContext(ctx context.Context) DynamicTagGroupOutput {
	return o
}

// Alarm contact group. The value range of N is 1~100. The alarm notification of the application group is sent to the alarm contact in the alarm contact group.
func (o DynamicTagGroupOutput) ContactGroupLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DynamicTagGroup) pulumi.StringArrayOutput { return v.ContactGroupLists }).(pulumi.StringArrayOutput)
}

// The relationship between conditional expressions. Valid values: `and`, `or`.
func (o DynamicTagGroupOutput) MatchExpressFilterRelation() pulumi.StringOutput {
	return o.ApplyT(func(v *DynamicTagGroup) pulumi.StringOutput { return v.MatchExpressFilterRelation }).(pulumi.StringOutput)
}

// The label generates a matching expression that applies the grouping. See the following `Block matchExpress`.
func (o DynamicTagGroupOutput) MatchExpresses() DynamicTagGroupMatchExpressArrayOutput {
	return o.ApplyT(func(v *DynamicTagGroup) DynamicTagGroupMatchExpressArrayOutput { return v.MatchExpresses }).(DynamicTagGroupMatchExpressArrayOutput)
}

// The status of the resource. Valid values: `RUNNING`, `FINISH`.
func (o DynamicTagGroupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DynamicTagGroup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tag key of the tag.
func (o DynamicTagGroupOutput) TagKey() pulumi.StringOutput {
	return o.ApplyT(func(v *DynamicTagGroup) pulumi.StringOutput { return v.TagKey }).(pulumi.StringOutput)
}

// Alarm template ID list.
func (o DynamicTagGroupOutput) TemplateIdLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DynamicTagGroup) pulumi.StringArrayOutput { return v.TemplateIdLists }).(pulumi.StringArrayOutput)
}

type DynamicTagGroupArrayOutput struct{ *pulumi.OutputState }

func (DynamicTagGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DynamicTagGroup)(nil)).Elem()
}

func (o DynamicTagGroupArrayOutput) ToDynamicTagGroupArrayOutput() DynamicTagGroupArrayOutput {
	return o
}

func (o DynamicTagGroupArrayOutput) ToDynamicTagGroupArrayOutputWithContext(ctx context.Context) DynamicTagGroupArrayOutput {
	return o
}

func (o DynamicTagGroupArrayOutput) Index(i pulumi.IntInput) DynamicTagGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DynamicTagGroup {
		return vs[0].([]*DynamicTagGroup)[vs[1].(int)]
	}).(DynamicTagGroupOutput)
}

type DynamicTagGroupMapOutput struct{ *pulumi.OutputState }

func (DynamicTagGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DynamicTagGroup)(nil)).Elem()
}

func (o DynamicTagGroupMapOutput) ToDynamicTagGroupMapOutput() DynamicTagGroupMapOutput {
	return o
}

func (o DynamicTagGroupMapOutput) ToDynamicTagGroupMapOutputWithContext(ctx context.Context) DynamicTagGroupMapOutput {
	return o
}

func (o DynamicTagGroupMapOutput) MapIndex(k pulumi.StringInput) DynamicTagGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DynamicTagGroup {
		return vs[0].(map[string]*DynamicTagGroup)[vs[1].(string)]
	}).(DynamicTagGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DynamicTagGroupInput)(nil)).Elem(), &DynamicTagGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*DynamicTagGroupArrayInput)(nil)).Elem(), DynamicTagGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DynamicTagGroupMapInput)(nil)).Elem(), DynamicTagGroupMap{})
	pulumi.RegisterOutputType(DynamicTagGroupOutput{})
	pulumi.RegisterOutputType(DynamicTagGroupArrayOutput{})
	pulumi.RegisterOutputType(DynamicTagGroupMapOutput{})
}
