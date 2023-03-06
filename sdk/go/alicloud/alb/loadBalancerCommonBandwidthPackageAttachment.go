// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Alb Load Balancer Common Bandwidth Package Attachment resource.
//
// For information about Alb Load Balancer Common Bandwidth Package Attachment and how to use it, see [What is Load Balancer Common Bandwidth Package Attachment](https://www.alibabacloud.com/help/en/server-load-balancer/latest/attachcommonbandwidthpackagetoloadbalancer).
//
// > **NOTE:** Available in v1.200.0+.
//
// ## Import
//
// Alb Load Balancer Common Bandwidth Package Attachment can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:alb/loadBalancerCommonBandwidthPackageAttachment:LoadBalancerCommonBandwidthPackageAttachment example <load_balancer_id>:<bandwidth_package_id>
//
// ```
type LoadBalancerCommonBandwidthPackageAttachment struct {
	pulumi.CustomResourceState

	// The ID of the bound shared bandwidth package.
	BandwidthPackageId pulumi.StringOutput `pulumi:"bandwidthPackageId"`
	// Whether to PreCheck this request only. Value:-**true**: sends a check request and does not bind the shared bandwidth package to the load balancing instance. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The ID of the applied server load balancer instance.
	LoadBalancerId pulumi.StringOutput `pulumi:"loadBalancerId"`
	// The status of the Application Load balancing instance. Value:-**Inactive**: Stopped, indicating that the instance listener will no longer forward traffic.-**Active**: running.-**Provisioning**: The project is being created.-**Configuring**: The configuration is being changed.-**CreateFailed**: The instance cannot be deleted without any charge.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewLoadBalancerCommonBandwidthPackageAttachment registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancerCommonBandwidthPackageAttachment(ctx *pulumi.Context,
	name string, args *LoadBalancerCommonBandwidthPackageAttachmentArgs, opts ...pulumi.ResourceOption) (*LoadBalancerCommonBandwidthPackageAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BandwidthPackageId == nil {
		return nil, errors.New("invalid value for required argument 'BandwidthPackageId'")
	}
	if args.LoadBalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerId'")
	}
	var resource LoadBalancerCommonBandwidthPackageAttachment
	err := ctx.RegisterResource("alicloud:alb/loadBalancerCommonBandwidthPackageAttachment:LoadBalancerCommonBandwidthPackageAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancerCommonBandwidthPackageAttachment gets an existing LoadBalancerCommonBandwidthPackageAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancerCommonBandwidthPackageAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerCommonBandwidthPackageAttachmentState, opts ...pulumi.ResourceOption) (*LoadBalancerCommonBandwidthPackageAttachment, error) {
	var resource LoadBalancerCommonBandwidthPackageAttachment
	err := ctx.ReadResource("alicloud:alb/loadBalancerCommonBandwidthPackageAttachment:LoadBalancerCommonBandwidthPackageAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancerCommonBandwidthPackageAttachment resources.
type loadBalancerCommonBandwidthPackageAttachmentState struct {
	// The ID of the bound shared bandwidth package.
	BandwidthPackageId *string `pulumi:"bandwidthPackageId"`
	// Whether to PreCheck this request only. Value:-**true**: sends a check request and does not bind the shared bandwidth package to the load balancing instance. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
	DryRun *bool `pulumi:"dryRun"`
	// The ID of the applied server load balancer instance.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// The status of the Application Load balancing instance. Value:-**Inactive**: Stopped, indicating that the instance listener will no longer forward traffic.-**Active**: running.-**Provisioning**: The project is being created.-**Configuring**: The configuration is being changed.-**CreateFailed**: The instance cannot be deleted without any charge.
	Status *string `pulumi:"status"`
}

type LoadBalancerCommonBandwidthPackageAttachmentState struct {
	// The ID of the bound shared bandwidth package.
	BandwidthPackageId pulumi.StringPtrInput
	// Whether to PreCheck this request only. Value:-**true**: sends a check request and does not bind the shared bandwidth package to the load balancing instance. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
	DryRun pulumi.BoolPtrInput
	// The ID of the applied server load balancer instance.
	LoadBalancerId pulumi.StringPtrInput
	// The status of the Application Load balancing instance. Value:-**Inactive**: Stopped, indicating that the instance listener will no longer forward traffic.-**Active**: running.-**Provisioning**: The project is being created.-**Configuring**: The configuration is being changed.-**CreateFailed**: The instance cannot be deleted without any charge.
	Status pulumi.StringPtrInput
}

func (LoadBalancerCommonBandwidthPackageAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerCommonBandwidthPackageAttachmentState)(nil)).Elem()
}

type loadBalancerCommonBandwidthPackageAttachmentArgs struct {
	// The ID of the bound shared bandwidth package.
	BandwidthPackageId string `pulumi:"bandwidthPackageId"`
	// Whether to PreCheck this request only. Value:-**true**: sends a check request and does not bind the shared bandwidth package to the load balancing instance. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
	DryRun *bool `pulumi:"dryRun"`
	// The ID of the applied server load balancer instance.
	LoadBalancerId string `pulumi:"loadBalancerId"`
}

// The set of arguments for constructing a LoadBalancerCommonBandwidthPackageAttachment resource.
type LoadBalancerCommonBandwidthPackageAttachmentArgs struct {
	// The ID of the bound shared bandwidth package.
	BandwidthPackageId pulumi.StringInput
	// Whether to PreCheck this request only. Value:-**true**: sends a check request and does not bind the shared bandwidth package to the load balancing instance. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
	DryRun pulumi.BoolPtrInput
	// The ID of the applied server load balancer instance.
	LoadBalancerId pulumi.StringInput
}

func (LoadBalancerCommonBandwidthPackageAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerCommonBandwidthPackageAttachmentArgs)(nil)).Elem()
}

type LoadBalancerCommonBandwidthPackageAttachmentInput interface {
	pulumi.Input

	ToLoadBalancerCommonBandwidthPackageAttachmentOutput() LoadBalancerCommonBandwidthPackageAttachmentOutput
	ToLoadBalancerCommonBandwidthPackageAttachmentOutputWithContext(ctx context.Context) LoadBalancerCommonBandwidthPackageAttachmentOutput
}

func (*LoadBalancerCommonBandwidthPackageAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerCommonBandwidthPackageAttachment)(nil)).Elem()
}

func (i *LoadBalancerCommonBandwidthPackageAttachment) ToLoadBalancerCommonBandwidthPackageAttachmentOutput() LoadBalancerCommonBandwidthPackageAttachmentOutput {
	return i.ToLoadBalancerCommonBandwidthPackageAttachmentOutputWithContext(context.Background())
}

func (i *LoadBalancerCommonBandwidthPackageAttachment) ToLoadBalancerCommonBandwidthPackageAttachmentOutputWithContext(ctx context.Context) LoadBalancerCommonBandwidthPackageAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerCommonBandwidthPackageAttachmentOutput)
}

// LoadBalancerCommonBandwidthPackageAttachmentArrayInput is an input type that accepts LoadBalancerCommonBandwidthPackageAttachmentArray and LoadBalancerCommonBandwidthPackageAttachmentArrayOutput values.
// You can construct a concrete instance of `LoadBalancerCommonBandwidthPackageAttachmentArrayInput` via:
//
//	LoadBalancerCommonBandwidthPackageAttachmentArray{ LoadBalancerCommonBandwidthPackageAttachmentArgs{...} }
type LoadBalancerCommonBandwidthPackageAttachmentArrayInput interface {
	pulumi.Input

	ToLoadBalancerCommonBandwidthPackageAttachmentArrayOutput() LoadBalancerCommonBandwidthPackageAttachmentArrayOutput
	ToLoadBalancerCommonBandwidthPackageAttachmentArrayOutputWithContext(context.Context) LoadBalancerCommonBandwidthPackageAttachmentArrayOutput
}

type LoadBalancerCommonBandwidthPackageAttachmentArray []LoadBalancerCommonBandwidthPackageAttachmentInput

func (LoadBalancerCommonBandwidthPackageAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancerCommonBandwidthPackageAttachment)(nil)).Elem()
}

func (i LoadBalancerCommonBandwidthPackageAttachmentArray) ToLoadBalancerCommonBandwidthPackageAttachmentArrayOutput() LoadBalancerCommonBandwidthPackageAttachmentArrayOutput {
	return i.ToLoadBalancerCommonBandwidthPackageAttachmentArrayOutputWithContext(context.Background())
}

func (i LoadBalancerCommonBandwidthPackageAttachmentArray) ToLoadBalancerCommonBandwidthPackageAttachmentArrayOutputWithContext(ctx context.Context) LoadBalancerCommonBandwidthPackageAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerCommonBandwidthPackageAttachmentArrayOutput)
}

// LoadBalancerCommonBandwidthPackageAttachmentMapInput is an input type that accepts LoadBalancerCommonBandwidthPackageAttachmentMap and LoadBalancerCommonBandwidthPackageAttachmentMapOutput values.
// You can construct a concrete instance of `LoadBalancerCommonBandwidthPackageAttachmentMapInput` via:
//
//	LoadBalancerCommonBandwidthPackageAttachmentMap{ "key": LoadBalancerCommonBandwidthPackageAttachmentArgs{...} }
type LoadBalancerCommonBandwidthPackageAttachmentMapInput interface {
	pulumi.Input

	ToLoadBalancerCommonBandwidthPackageAttachmentMapOutput() LoadBalancerCommonBandwidthPackageAttachmentMapOutput
	ToLoadBalancerCommonBandwidthPackageAttachmentMapOutputWithContext(context.Context) LoadBalancerCommonBandwidthPackageAttachmentMapOutput
}

type LoadBalancerCommonBandwidthPackageAttachmentMap map[string]LoadBalancerCommonBandwidthPackageAttachmentInput

func (LoadBalancerCommonBandwidthPackageAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancerCommonBandwidthPackageAttachment)(nil)).Elem()
}

func (i LoadBalancerCommonBandwidthPackageAttachmentMap) ToLoadBalancerCommonBandwidthPackageAttachmentMapOutput() LoadBalancerCommonBandwidthPackageAttachmentMapOutput {
	return i.ToLoadBalancerCommonBandwidthPackageAttachmentMapOutputWithContext(context.Background())
}

func (i LoadBalancerCommonBandwidthPackageAttachmentMap) ToLoadBalancerCommonBandwidthPackageAttachmentMapOutputWithContext(ctx context.Context) LoadBalancerCommonBandwidthPackageAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerCommonBandwidthPackageAttachmentMapOutput)
}

type LoadBalancerCommonBandwidthPackageAttachmentOutput struct{ *pulumi.OutputState }

func (LoadBalancerCommonBandwidthPackageAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerCommonBandwidthPackageAttachment)(nil)).Elem()
}

func (o LoadBalancerCommonBandwidthPackageAttachmentOutput) ToLoadBalancerCommonBandwidthPackageAttachmentOutput() LoadBalancerCommonBandwidthPackageAttachmentOutput {
	return o
}

func (o LoadBalancerCommonBandwidthPackageAttachmentOutput) ToLoadBalancerCommonBandwidthPackageAttachmentOutputWithContext(ctx context.Context) LoadBalancerCommonBandwidthPackageAttachmentOutput {
	return o
}

// The ID of the bound shared bandwidth package.
func (o LoadBalancerCommonBandwidthPackageAttachmentOutput) BandwidthPackageId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancerCommonBandwidthPackageAttachment) pulumi.StringOutput { return v.BandwidthPackageId }).(pulumi.StringOutput)
}

// Whether to PreCheck this request only. Value:-**true**: sends a check request and does not bind the shared bandwidth package to the load balancing instance. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
func (o LoadBalancerCommonBandwidthPackageAttachmentOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LoadBalancerCommonBandwidthPackageAttachment) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The ID of the applied server load balancer instance.
func (o LoadBalancerCommonBandwidthPackageAttachmentOutput) LoadBalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancerCommonBandwidthPackageAttachment) pulumi.StringOutput { return v.LoadBalancerId }).(pulumi.StringOutput)
}

// The status of the Application Load balancing instance. Value:-**Inactive**: Stopped, indicating that the instance listener will no longer forward traffic.-**Active**: running.-**Provisioning**: The project is being created.-**Configuring**: The configuration is being changed.-**CreateFailed**: The instance cannot be deleted without any charge.
func (o LoadBalancerCommonBandwidthPackageAttachmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancerCommonBandwidthPackageAttachment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type LoadBalancerCommonBandwidthPackageAttachmentArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerCommonBandwidthPackageAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancerCommonBandwidthPackageAttachment)(nil)).Elem()
}

func (o LoadBalancerCommonBandwidthPackageAttachmentArrayOutput) ToLoadBalancerCommonBandwidthPackageAttachmentArrayOutput() LoadBalancerCommonBandwidthPackageAttachmentArrayOutput {
	return o
}

func (o LoadBalancerCommonBandwidthPackageAttachmentArrayOutput) ToLoadBalancerCommonBandwidthPackageAttachmentArrayOutputWithContext(ctx context.Context) LoadBalancerCommonBandwidthPackageAttachmentArrayOutput {
	return o
}

func (o LoadBalancerCommonBandwidthPackageAttachmentArrayOutput) Index(i pulumi.IntInput) LoadBalancerCommonBandwidthPackageAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LoadBalancerCommonBandwidthPackageAttachment {
		return vs[0].([]*LoadBalancerCommonBandwidthPackageAttachment)[vs[1].(int)]
	}).(LoadBalancerCommonBandwidthPackageAttachmentOutput)
}

type LoadBalancerCommonBandwidthPackageAttachmentMapOutput struct{ *pulumi.OutputState }

func (LoadBalancerCommonBandwidthPackageAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancerCommonBandwidthPackageAttachment)(nil)).Elem()
}

func (o LoadBalancerCommonBandwidthPackageAttachmentMapOutput) ToLoadBalancerCommonBandwidthPackageAttachmentMapOutput() LoadBalancerCommonBandwidthPackageAttachmentMapOutput {
	return o
}

func (o LoadBalancerCommonBandwidthPackageAttachmentMapOutput) ToLoadBalancerCommonBandwidthPackageAttachmentMapOutputWithContext(ctx context.Context) LoadBalancerCommonBandwidthPackageAttachmentMapOutput {
	return o
}

func (o LoadBalancerCommonBandwidthPackageAttachmentMapOutput) MapIndex(k pulumi.StringInput) LoadBalancerCommonBandwidthPackageAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LoadBalancerCommonBandwidthPackageAttachment {
		return vs[0].(map[string]*LoadBalancerCommonBandwidthPackageAttachment)[vs[1].(string)]
	}).(LoadBalancerCommonBandwidthPackageAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerCommonBandwidthPackageAttachmentInput)(nil)).Elem(), &LoadBalancerCommonBandwidthPackageAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerCommonBandwidthPackageAttachmentArrayInput)(nil)).Elem(), LoadBalancerCommonBandwidthPackageAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerCommonBandwidthPackageAttachmentMapInput)(nil)).Elem(), LoadBalancerCommonBandwidthPackageAttachmentMap{})
	pulumi.RegisterOutputType(LoadBalancerCommonBandwidthPackageAttachmentOutput{})
	pulumi.RegisterOutputType(LoadBalancerCommonBandwidthPackageAttachmentArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerCommonBandwidthPackageAttachmentMapOutput{})
}