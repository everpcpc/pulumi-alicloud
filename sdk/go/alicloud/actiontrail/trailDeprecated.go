// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package actiontrail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Action trail can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:actiontrail/trailDeprecated:TrailDeprecated foo abc12345678
// ```
//
// Deprecated: Resource renamed to `Trail`
type TrailDeprecated struct {
	pulumi.CustomResourceState

	// Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
	EventRw             pulumi.StringPtrOutput `pulumi:"eventRw"`
	IsOrganizationTrail pulumi.BoolPtrOutput   `pulumi:"isOrganizationTrail"`
	MnsTopicArn         pulumi.StringPtrOutput `pulumi:"mnsTopicArn"`
	// The name of the trail to be created, which must be unique for an account.
	//
	// Deprecated: Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
	OssBucketName pulumi.StringPtrOutput `pulumi:"ossBucketName"`
	// The prefix of the specified OSS bucket name. This parameter can be left empty.
	OssKeyPrefix pulumi.StringPtrOutput `pulumi:"ossKeyPrefix"`
	// The RAM role in ActionTrail permitted by the user.
	RoleName pulumi.StringOutput `pulumi:"roleName"`
	// The unique ARN of the Log Service project.
	SlsProjectArn pulumi.StringPtrOutput `pulumi:"slsProjectArn"`
	// The unique ARN of the Log Service role.
	SlsWriteRoleArn pulumi.StringPtrOutput `pulumi:"slsWriteRoleArn"`
	Status          pulumi.StringPtrOutput `pulumi:"status"`
	TrailName       pulumi.StringOutput    `pulumi:"trailName"`
	TrailRegion     pulumi.StringPtrOutput `pulumi:"trailRegion"`
}

// NewTrailDeprecated registers a new resource with the given unique name, arguments, and options.
func NewTrailDeprecated(ctx *pulumi.Context,
	name string, args *TrailDeprecatedArgs, opts ...pulumi.ResourceOption) (*TrailDeprecated, error) {
	if args == nil {
		args = &TrailDeprecatedArgs{}
	}

	var resource TrailDeprecated
	err := ctx.RegisterResource("alicloud:actiontrail/trailDeprecated:TrailDeprecated", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrailDeprecated gets an existing TrailDeprecated resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrailDeprecated(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrailDeprecatedState, opts ...pulumi.ResourceOption) (*TrailDeprecated, error) {
	var resource TrailDeprecated
	err := ctx.ReadResource("alicloud:actiontrail/trailDeprecated:TrailDeprecated", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrailDeprecated resources.
type trailDeprecatedState struct {
	// Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
	EventRw             *string `pulumi:"eventRw"`
	IsOrganizationTrail *bool   `pulumi:"isOrganizationTrail"`
	MnsTopicArn         *string `pulumi:"mnsTopicArn"`
	// The name of the trail to be created, which must be unique for an account.
	//
	// Deprecated: Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
	Name *string `pulumi:"name"`
	// The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
	OssBucketName *string `pulumi:"ossBucketName"`
	// The prefix of the specified OSS bucket name. This parameter can be left empty.
	OssKeyPrefix *string `pulumi:"ossKeyPrefix"`
	// The RAM role in ActionTrail permitted by the user.
	RoleName *string `pulumi:"roleName"`
	// The unique ARN of the Log Service project.
	SlsProjectArn *string `pulumi:"slsProjectArn"`
	// The unique ARN of the Log Service role.
	SlsWriteRoleArn *string `pulumi:"slsWriteRoleArn"`
	Status          *string `pulumi:"status"`
	TrailName       *string `pulumi:"trailName"`
	TrailRegion     *string `pulumi:"trailRegion"`
}

type TrailDeprecatedState struct {
	// Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
	EventRw             pulumi.StringPtrInput
	IsOrganizationTrail pulumi.BoolPtrInput
	MnsTopicArn         pulumi.StringPtrInput
	// The name of the trail to be created, which must be unique for an account.
	//
	// Deprecated: Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
	Name pulumi.StringPtrInput
	// The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
	OssBucketName pulumi.StringPtrInput
	// The prefix of the specified OSS bucket name. This parameter can be left empty.
	OssKeyPrefix pulumi.StringPtrInput
	// The RAM role in ActionTrail permitted by the user.
	RoleName pulumi.StringPtrInput
	// The unique ARN of the Log Service project.
	SlsProjectArn pulumi.StringPtrInput
	// The unique ARN of the Log Service role.
	SlsWriteRoleArn pulumi.StringPtrInput
	Status          pulumi.StringPtrInput
	TrailName       pulumi.StringPtrInput
	TrailRegion     pulumi.StringPtrInput
}

func (TrailDeprecatedState) ElementType() reflect.Type {
	return reflect.TypeOf((*trailDeprecatedState)(nil)).Elem()
}

type trailDeprecatedArgs struct {
	// Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
	EventRw             *string `pulumi:"eventRw"`
	IsOrganizationTrail *bool   `pulumi:"isOrganizationTrail"`
	MnsTopicArn         *string `pulumi:"mnsTopicArn"`
	// The name of the trail to be created, which must be unique for an account.
	//
	// Deprecated: Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
	Name *string `pulumi:"name"`
	// The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
	OssBucketName *string `pulumi:"ossBucketName"`
	// The prefix of the specified OSS bucket name. This parameter can be left empty.
	OssKeyPrefix *string `pulumi:"ossKeyPrefix"`
	// The RAM role in ActionTrail permitted by the user.
	RoleName *string `pulumi:"roleName"`
	// The unique ARN of the Log Service project.
	SlsProjectArn *string `pulumi:"slsProjectArn"`
	// The unique ARN of the Log Service role.
	SlsWriteRoleArn *string `pulumi:"slsWriteRoleArn"`
	Status          *string `pulumi:"status"`
	TrailName       *string `pulumi:"trailName"`
	TrailRegion     *string `pulumi:"trailRegion"`
}

// The set of arguments for constructing a TrailDeprecated resource.
type TrailDeprecatedArgs struct {
	// Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
	EventRw             pulumi.StringPtrInput
	IsOrganizationTrail pulumi.BoolPtrInput
	MnsTopicArn         pulumi.StringPtrInput
	// The name of the trail to be created, which must be unique for an account.
	//
	// Deprecated: Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
	Name pulumi.StringPtrInput
	// The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
	OssBucketName pulumi.StringPtrInput
	// The prefix of the specified OSS bucket name. This parameter can be left empty.
	OssKeyPrefix pulumi.StringPtrInput
	// The RAM role in ActionTrail permitted by the user.
	RoleName pulumi.StringPtrInput
	// The unique ARN of the Log Service project.
	SlsProjectArn pulumi.StringPtrInput
	// The unique ARN of the Log Service role.
	SlsWriteRoleArn pulumi.StringPtrInput
	Status          pulumi.StringPtrInput
	TrailName       pulumi.StringPtrInput
	TrailRegion     pulumi.StringPtrInput
}

func (TrailDeprecatedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trailDeprecatedArgs)(nil)).Elem()
}

type TrailDeprecatedInput interface {
	pulumi.Input

	ToTrailDeprecatedOutput() TrailDeprecatedOutput
	ToTrailDeprecatedOutputWithContext(ctx context.Context) TrailDeprecatedOutput
}

func (TrailDeprecated) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailDeprecated)(nil)).Elem()
}

func (i TrailDeprecated) ToTrailDeprecatedOutput() TrailDeprecatedOutput {
	return i.ToTrailDeprecatedOutputWithContext(context.Background())
}

func (i TrailDeprecated) ToTrailDeprecatedOutputWithContext(ctx context.Context) TrailDeprecatedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailDeprecatedOutput)
}

type TrailDeprecatedOutput struct {
	*pulumi.OutputState
}

func (TrailDeprecatedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailDeprecatedOutput)(nil)).Elem()
}

func (o TrailDeprecatedOutput) ToTrailDeprecatedOutput() TrailDeprecatedOutput {
	return o
}

func (o TrailDeprecatedOutput) ToTrailDeprecatedOutputWithContext(ctx context.Context) TrailDeprecatedOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TrailDeprecatedOutput{})
}
