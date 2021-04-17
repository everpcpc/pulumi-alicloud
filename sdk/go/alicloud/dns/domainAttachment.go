// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// DNS domain attachment can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:dns/domainAttachment:DomainAttachment example dns-cn-v0h1ldjhxxx
// ```
type DomainAttachment struct {
	pulumi.CustomResourceState

	// The domain names bound to the DNS instance.
	DomainNames pulumi.StringArrayOutput `pulumi:"domainNames"`
	// The id of the DNS instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewDomainAttachment registers a new resource with the given unique name, arguments, and options.
func NewDomainAttachment(ctx *pulumi.Context,
	name string, args *DomainAttachmentArgs, opts ...pulumi.ResourceOption) (*DomainAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainNames == nil {
		return nil, errors.New("invalid value for required argument 'DomainNames'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource DomainAttachment
	err := ctx.RegisterResource("alicloud:dns/domainAttachment:DomainAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainAttachment gets an existing DomainAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainAttachmentState, opts ...pulumi.ResourceOption) (*DomainAttachment, error) {
	var resource DomainAttachment
	err := ctx.ReadResource("alicloud:dns/domainAttachment:DomainAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainAttachment resources.
type domainAttachmentState struct {
	// The domain names bound to the DNS instance.
	DomainNames []string `pulumi:"domainNames"`
	// The id of the DNS instance.
	InstanceId *string `pulumi:"instanceId"`
}

type DomainAttachmentState struct {
	// The domain names bound to the DNS instance.
	DomainNames pulumi.StringArrayInput
	// The id of the DNS instance.
	InstanceId pulumi.StringPtrInput
}

func (DomainAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainAttachmentState)(nil)).Elem()
}

type domainAttachmentArgs struct {
	// The domain names bound to the DNS instance.
	DomainNames []string `pulumi:"domainNames"`
	// The id of the DNS instance.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a DomainAttachment resource.
type DomainAttachmentArgs struct {
	// The domain names bound to the DNS instance.
	DomainNames pulumi.StringArrayInput
	// The id of the DNS instance.
	InstanceId pulumi.StringInput
}

func (DomainAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainAttachmentArgs)(nil)).Elem()
}

type DomainAttachmentInput interface {
	pulumi.Input

	ToDomainAttachmentOutput() DomainAttachmentOutput
	ToDomainAttachmentOutputWithContext(ctx context.Context) DomainAttachmentOutput
}

func (*DomainAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainAttachment)(nil))
}

func (i *DomainAttachment) ToDomainAttachmentOutput() DomainAttachmentOutput {
	return i.ToDomainAttachmentOutputWithContext(context.Background())
}

func (i *DomainAttachment) ToDomainAttachmentOutputWithContext(ctx context.Context) DomainAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAttachmentOutput)
}

func (i *DomainAttachment) ToDomainAttachmentPtrOutput() DomainAttachmentPtrOutput {
	return i.ToDomainAttachmentPtrOutputWithContext(context.Background())
}

func (i *DomainAttachment) ToDomainAttachmentPtrOutputWithContext(ctx context.Context) DomainAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAttachmentPtrOutput)
}

type DomainAttachmentPtrInput interface {
	pulumi.Input

	ToDomainAttachmentPtrOutput() DomainAttachmentPtrOutput
	ToDomainAttachmentPtrOutputWithContext(ctx context.Context) DomainAttachmentPtrOutput
}

type domainAttachmentPtrType DomainAttachmentArgs

func (*domainAttachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainAttachment)(nil))
}

func (i *domainAttachmentPtrType) ToDomainAttachmentPtrOutput() DomainAttachmentPtrOutput {
	return i.ToDomainAttachmentPtrOutputWithContext(context.Background())
}

func (i *domainAttachmentPtrType) ToDomainAttachmentPtrOutputWithContext(ctx context.Context) DomainAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAttachmentPtrOutput)
}

// DomainAttachmentArrayInput is an input type that accepts DomainAttachmentArray and DomainAttachmentArrayOutput values.
// You can construct a concrete instance of `DomainAttachmentArrayInput` via:
//
//          DomainAttachmentArray{ DomainAttachmentArgs{...} }
type DomainAttachmentArrayInput interface {
	pulumi.Input

	ToDomainAttachmentArrayOutput() DomainAttachmentArrayOutput
	ToDomainAttachmentArrayOutputWithContext(context.Context) DomainAttachmentArrayOutput
}

type DomainAttachmentArray []DomainAttachmentInput

func (DomainAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DomainAttachment)(nil))
}

func (i DomainAttachmentArray) ToDomainAttachmentArrayOutput() DomainAttachmentArrayOutput {
	return i.ToDomainAttachmentArrayOutputWithContext(context.Background())
}

func (i DomainAttachmentArray) ToDomainAttachmentArrayOutputWithContext(ctx context.Context) DomainAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAttachmentArrayOutput)
}

// DomainAttachmentMapInput is an input type that accepts DomainAttachmentMap and DomainAttachmentMapOutput values.
// You can construct a concrete instance of `DomainAttachmentMapInput` via:
//
//          DomainAttachmentMap{ "key": DomainAttachmentArgs{...} }
type DomainAttachmentMapInput interface {
	pulumi.Input

	ToDomainAttachmentMapOutput() DomainAttachmentMapOutput
	ToDomainAttachmentMapOutputWithContext(context.Context) DomainAttachmentMapOutput
}

type DomainAttachmentMap map[string]DomainAttachmentInput

func (DomainAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DomainAttachment)(nil))
}

func (i DomainAttachmentMap) ToDomainAttachmentMapOutput() DomainAttachmentMapOutput {
	return i.ToDomainAttachmentMapOutputWithContext(context.Background())
}

func (i DomainAttachmentMap) ToDomainAttachmentMapOutputWithContext(ctx context.Context) DomainAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAttachmentMapOutput)
}

type DomainAttachmentOutput struct {
	*pulumi.OutputState
}

func (DomainAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainAttachment)(nil))
}

func (o DomainAttachmentOutput) ToDomainAttachmentOutput() DomainAttachmentOutput {
	return o
}

func (o DomainAttachmentOutput) ToDomainAttachmentOutputWithContext(ctx context.Context) DomainAttachmentOutput {
	return o
}

func (o DomainAttachmentOutput) ToDomainAttachmentPtrOutput() DomainAttachmentPtrOutput {
	return o.ToDomainAttachmentPtrOutputWithContext(context.Background())
}

func (o DomainAttachmentOutput) ToDomainAttachmentPtrOutputWithContext(ctx context.Context) DomainAttachmentPtrOutput {
	return o.ApplyT(func(v DomainAttachment) *DomainAttachment {
		return &v
	}).(DomainAttachmentPtrOutput)
}

type DomainAttachmentPtrOutput struct {
	*pulumi.OutputState
}

func (DomainAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainAttachment)(nil))
}

func (o DomainAttachmentPtrOutput) ToDomainAttachmentPtrOutput() DomainAttachmentPtrOutput {
	return o
}

func (o DomainAttachmentPtrOutput) ToDomainAttachmentPtrOutputWithContext(ctx context.Context) DomainAttachmentPtrOutput {
	return o
}

type DomainAttachmentArrayOutput struct{ *pulumi.OutputState }

func (DomainAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainAttachment)(nil))
}

func (o DomainAttachmentArrayOutput) ToDomainAttachmentArrayOutput() DomainAttachmentArrayOutput {
	return o
}

func (o DomainAttachmentArrayOutput) ToDomainAttachmentArrayOutputWithContext(ctx context.Context) DomainAttachmentArrayOutput {
	return o
}

func (o DomainAttachmentArrayOutput) Index(i pulumi.IntInput) DomainAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainAttachment {
		return vs[0].([]DomainAttachment)[vs[1].(int)]
	}).(DomainAttachmentOutput)
}

type DomainAttachmentMapOutput struct{ *pulumi.OutputState }

func (DomainAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DomainAttachment)(nil))
}

func (o DomainAttachmentMapOutput) ToDomainAttachmentMapOutput() DomainAttachmentMapOutput {
	return o
}

func (o DomainAttachmentMapOutput) ToDomainAttachmentMapOutputWithContext(ctx context.Context) DomainAttachmentMapOutput {
	return o
}

func (o DomainAttachmentMapOutput) MapIndex(k pulumi.StringInput) DomainAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DomainAttachment {
		return vs[0].(map[string]DomainAttachment)[vs[1].(string)]
	}).(DomainAttachmentOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainAttachmentOutput{})
	pulumi.RegisterOutputType(DomainAttachmentPtrOutput{})
	pulumi.RegisterOutputType(DomainAttachmentArrayOutput{})
	pulumi.RegisterOutputType(DomainAttachmentMapOutput{})
}
