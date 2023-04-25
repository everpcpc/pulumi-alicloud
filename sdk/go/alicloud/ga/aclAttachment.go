// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Global Accelerator (GA) Acl Attachment resource.
//
// For information about Global Accelerator (GA) Acl Attachment and how to use it, see [What is Acl Attachment](https://www.alibabacloud.com/help/en/doc-detail/258295.html).
//
// > **NOTE:** Available in v1.150.0+.
//
// ## Import
//
// Global Accelerator (GA) Acl Attachment can be imported using the id. Format to `<listener_id>:<acl_id>`, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ga/aclAttachment:AclAttachment example your_listener_id:your_acl_id
//
// ```
type AclAttachment struct {
	pulumi.CustomResourceState

	// The ID of an ACL.
	AclId pulumi.StringOutput `pulumi:"aclId"`
	// The type of the ACL. Valid values: `white`, `black`.
	AclType pulumi.StringOutput `pulumi:"aclType"`
	// The dry run.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The ID of the listener.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAclAttachment registers a new resource with the given unique name, arguments, and options.
func NewAclAttachment(ctx *pulumi.Context,
	name string, args *AclAttachmentArgs, opts ...pulumi.ResourceOption) (*AclAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AclId == nil {
		return nil, errors.New("invalid value for required argument 'AclId'")
	}
	if args.AclType == nil {
		return nil, errors.New("invalid value for required argument 'AclType'")
	}
	if args.ListenerId == nil {
		return nil, errors.New("invalid value for required argument 'ListenerId'")
	}
	var resource AclAttachment
	err := ctx.RegisterResource("alicloud:ga/aclAttachment:AclAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAclAttachment gets an existing AclAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAclAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AclAttachmentState, opts ...pulumi.ResourceOption) (*AclAttachment, error) {
	var resource AclAttachment
	err := ctx.ReadResource("alicloud:ga/aclAttachment:AclAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AclAttachment resources.
type aclAttachmentState struct {
	// The ID of an ACL.
	AclId *string `pulumi:"aclId"`
	// The type of the ACL. Valid values: `white`, `black`.
	AclType *string `pulumi:"aclType"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The ID of the listener.
	ListenerId *string `pulumi:"listenerId"`
	// The status of the resource.
	Status *string `pulumi:"status"`
}

type AclAttachmentState struct {
	// The ID of an ACL.
	AclId pulumi.StringPtrInput
	// The type of the ACL. Valid values: `white`, `black`.
	AclType pulumi.StringPtrInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The ID of the listener.
	ListenerId pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
}

func (AclAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclAttachmentState)(nil)).Elem()
}

type aclAttachmentArgs struct {
	// The ID of an ACL.
	AclId string `pulumi:"aclId"`
	// The type of the ACL. Valid values: `white`, `black`.
	AclType string `pulumi:"aclType"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The ID of the listener.
	ListenerId string `pulumi:"listenerId"`
}

// The set of arguments for constructing a AclAttachment resource.
type AclAttachmentArgs struct {
	// The ID of an ACL.
	AclId pulumi.StringInput
	// The type of the ACL. Valid values: `white`, `black`.
	AclType pulumi.StringInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The ID of the listener.
	ListenerId pulumi.StringInput
}

func (AclAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclAttachmentArgs)(nil)).Elem()
}

type AclAttachmentInput interface {
	pulumi.Input

	ToAclAttachmentOutput() AclAttachmentOutput
	ToAclAttachmentOutputWithContext(ctx context.Context) AclAttachmentOutput
}

func (*AclAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**AclAttachment)(nil)).Elem()
}

func (i *AclAttachment) ToAclAttachmentOutput() AclAttachmentOutput {
	return i.ToAclAttachmentOutputWithContext(context.Background())
}

func (i *AclAttachment) ToAclAttachmentOutputWithContext(ctx context.Context) AclAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclAttachmentOutput)
}

// AclAttachmentArrayInput is an input type that accepts AclAttachmentArray and AclAttachmentArrayOutput values.
// You can construct a concrete instance of `AclAttachmentArrayInput` via:
//
//	AclAttachmentArray{ AclAttachmentArgs{...} }
type AclAttachmentArrayInput interface {
	pulumi.Input

	ToAclAttachmentArrayOutput() AclAttachmentArrayOutput
	ToAclAttachmentArrayOutputWithContext(context.Context) AclAttachmentArrayOutput
}

type AclAttachmentArray []AclAttachmentInput

func (AclAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AclAttachment)(nil)).Elem()
}

func (i AclAttachmentArray) ToAclAttachmentArrayOutput() AclAttachmentArrayOutput {
	return i.ToAclAttachmentArrayOutputWithContext(context.Background())
}

func (i AclAttachmentArray) ToAclAttachmentArrayOutputWithContext(ctx context.Context) AclAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclAttachmentArrayOutput)
}

// AclAttachmentMapInput is an input type that accepts AclAttachmentMap and AclAttachmentMapOutput values.
// You can construct a concrete instance of `AclAttachmentMapInput` via:
//
//	AclAttachmentMap{ "key": AclAttachmentArgs{...} }
type AclAttachmentMapInput interface {
	pulumi.Input

	ToAclAttachmentMapOutput() AclAttachmentMapOutput
	ToAclAttachmentMapOutputWithContext(context.Context) AclAttachmentMapOutput
}

type AclAttachmentMap map[string]AclAttachmentInput

func (AclAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AclAttachment)(nil)).Elem()
}

func (i AclAttachmentMap) ToAclAttachmentMapOutput() AclAttachmentMapOutput {
	return i.ToAclAttachmentMapOutputWithContext(context.Background())
}

func (i AclAttachmentMap) ToAclAttachmentMapOutputWithContext(ctx context.Context) AclAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclAttachmentMapOutput)
}

type AclAttachmentOutput struct{ *pulumi.OutputState }

func (AclAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AclAttachment)(nil)).Elem()
}

func (o AclAttachmentOutput) ToAclAttachmentOutput() AclAttachmentOutput {
	return o
}

func (o AclAttachmentOutput) ToAclAttachmentOutputWithContext(ctx context.Context) AclAttachmentOutput {
	return o
}

// The ID of an ACL.
func (o AclAttachmentOutput) AclId() pulumi.StringOutput {
	return o.ApplyT(func(v *AclAttachment) pulumi.StringOutput { return v.AclId }).(pulumi.StringOutput)
}

// The type of the ACL. Valid values: `white`, `black`.
func (o AclAttachmentOutput) AclType() pulumi.StringOutput {
	return o.ApplyT(func(v *AclAttachment) pulumi.StringOutput { return v.AclType }).(pulumi.StringOutput)
}

// The dry run.
func (o AclAttachmentOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AclAttachment) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The ID of the listener.
func (o AclAttachmentOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *AclAttachment) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

// The status of the resource.
func (o AclAttachmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AclAttachment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AclAttachmentArrayOutput struct{ *pulumi.OutputState }

func (AclAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AclAttachment)(nil)).Elem()
}

func (o AclAttachmentArrayOutput) ToAclAttachmentArrayOutput() AclAttachmentArrayOutput {
	return o
}

func (o AclAttachmentArrayOutput) ToAclAttachmentArrayOutputWithContext(ctx context.Context) AclAttachmentArrayOutput {
	return o
}

func (o AclAttachmentArrayOutput) Index(i pulumi.IntInput) AclAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AclAttachment {
		return vs[0].([]*AclAttachment)[vs[1].(int)]
	}).(AclAttachmentOutput)
}

type AclAttachmentMapOutput struct{ *pulumi.OutputState }

func (AclAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AclAttachment)(nil)).Elem()
}

func (o AclAttachmentMapOutput) ToAclAttachmentMapOutput() AclAttachmentMapOutput {
	return o
}

func (o AclAttachmentMapOutput) ToAclAttachmentMapOutputWithContext(ctx context.Context) AclAttachmentMapOutput {
	return o
}

func (o AclAttachmentMapOutput) MapIndex(k pulumi.StringInput) AclAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AclAttachment {
		return vs[0].(map[string]*AclAttachment)[vs[1].(string)]
	}).(AclAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AclAttachmentInput)(nil)).Elem(), &AclAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclAttachmentArrayInput)(nil)).Elem(), AclAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclAttachmentMapInput)(nil)).Elem(), AclAttachmentMap{})
	pulumi.RegisterOutputType(AclAttachmentOutput{})
	pulumi.RegisterOutputType(AclAttachmentArrayOutput{})
	pulumi.RegisterOutputType(AclAttachmentMapOutput{})
}
