// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultInstanceTypes, err := ecs.GetInstanceTypes(ctx, &ecs.GetInstanceTypesArgs{
//				AvailabilityZone: pulumi.StringRef(defaultZones.Zones[0].Id),
//				CpuCoreCount:     pulumi.IntRef(1),
//				MemorySize:       pulumi.Float64Ref(2),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultImages, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
//				NameRegex:  pulumi.StringRef("^ubuntu_18.*64"),
//				MostRecent: pulumi.BoolRef(true),
//				Owners:     pulumi.StringRef("system"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			cfg := config.New(ctx, "")
//			name := "test_havip_attachment"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			fooNetwork, err := vpc.NewNetwork(ctx, "fooNetwork", &vpc.NetworkArgs{
//				CidrBlock: pulumi.String("172.16.0.0/12"),
//			})
//			if err != nil {
//				return err
//			}
//			fooSwitch, err := vpc.NewSwitch(ctx, "fooSwitch", &vpc.SwitchArgs{
//				VpcId:     fooNetwork.ID(),
//				CidrBlock: pulumi.String("172.16.0.0/21"),
//				ZoneId:    *pulumi.String(defaultZones.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			fooHAVip, err := vpc.NewHAVip(ctx, "fooHAVip", &vpc.HAVipArgs{
//				VswitchId:   fooSwitch.ID(),
//				Description: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			tfTestFoo, err := ecs.NewSecurityGroup(ctx, "tfTestFoo", &ecs.SecurityGroupArgs{
//				Description: pulumi.String("foo"),
//				VpcId:       fooNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			fooInstance, err := ecs.NewInstance(ctx, "fooInstance", &ecs.InstanceArgs{
//				AvailabilityZone:        *pulumi.String(defaultZones.Zones[0].Id),
//				VswitchId:               fooSwitch.ID(),
//				ImageId:                 *pulumi.String(defaultImages.Images[0].Id),
//				InstanceType:            *pulumi.String(defaultInstanceTypes.InstanceTypes[0].Id),
//				SystemDiskCategory:      pulumi.String("cloud_efficiency"),
//				InternetChargeType:      pulumi.String("PayByTraffic"),
//				InternetMaxBandwidthOut: pulumi.Int(5),
//				SecurityGroups: pulumi.StringArray{
//					tfTestFoo.ID(),
//				},
//				InstanceName: pulumi.String(name),
//				UserData:     pulumi.String("echo 'net.ipv4.ip_forward=1'>> /etc/sysctl.conf"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewHAVipAttachment(ctx, "fooHAVipAttachment", &vpc.HAVipAttachmentArgs{
//				HavipId:    fooHAVip.ID(),
//				InstanceId: fooInstance.ID(),
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
// The havip attachment can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:vpc/hAVipAttachment:HAVipAttachment foo havip-abc123456:i-abc123456
//
// ```
type HAVipAttachment struct {
	pulumi.CustomResourceState

	// Specifies whether to forcefully disassociate the HAVIP from the ECS instance or ENI. Default value: `False`. Valid values: `True` and `False`.
	Force pulumi.StringPtrOutput `pulumi:"force"`
	// The havipId of the havip attachment, the field can't be changed.
	HavipId pulumi.StringOutput `pulumi:"havipId"`
	// The instanceId of the havip attachment, the field can't be changed.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The Type of instance to bind HaVip to. Valid values: `EcsInstance` and `NetworkInterface`. When the HaVip instance is bound to a resilient NIC, the resilient NIC instance must be filled in.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// (Available in v1.201.0+) The status of the HaVip instance.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewHAVipAttachment registers a new resource with the given unique name, arguments, and options.
func NewHAVipAttachment(ctx *pulumi.Context,
	name string, args *HAVipAttachmentArgs, opts ...pulumi.ResourceOption) (*HAVipAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HavipId == nil {
		return nil, errors.New("invalid value for required argument 'HavipId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource HAVipAttachment
	err := ctx.RegisterResource("alicloud:vpc/hAVipAttachment:HAVipAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHAVipAttachment gets an existing HAVipAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHAVipAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HAVipAttachmentState, opts ...pulumi.ResourceOption) (*HAVipAttachment, error) {
	var resource HAVipAttachment
	err := ctx.ReadResource("alicloud:vpc/hAVipAttachment:HAVipAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HAVipAttachment resources.
type havipAttachmentState struct {
	// Specifies whether to forcefully disassociate the HAVIP from the ECS instance or ENI. Default value: `False`. Valid values: `True` and `False`.
	Force *string `pulumi:"force"`
	// The havipId of the havip attachment, the field can't be changed.
	HavipId *string `pulumi:"havipId"`
	// The instanceId of the havip attachment, the field can't be changed.
	InstanceId *string `pulumi:"instanceId"`
	// The Type of instance to bind HaVip to. Valid values: `EcsInstance` and `NetworkInterface`. When the HaVip instance is bound to a resilient NIC, the resilient NIC instance must be filled in.
	InstanceType *string `pulumi:"instanceType"`
	// (Available in v1.201.0+) The status of the HaVip instance.
	Status *string `pulumi:"status"`
}

type HAVipAttachmentState struct {
	// Specifies whether to forcefully disassociate the HAVIP from the ECS instance or ENI. Default value: `False`. Valid values: `True` and `False`.
	Force pulumi.StringPtrInput
	// The havipId of the havip attachment, the field can't be changed.
	HavipId pulumi.StringPtrInput
	// The instanceId of the havip attachment, the field can't be changed.
	InstanceId pulumi.StringPtrInput
	// The Type of instance to bind HaVip to. Valid values: `EcsInstance` and `NetworkInterface`. When the HaVip instance is bound to a resilient NIC, the resilient NIC instance must be filled in.
	InstanceType pulumi.StringPtrInput
	// (Available in v1.201.0+) The status of the HaVip instance.
	Status pulumi.StringPtrInput
}

func (HAVipAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*havipAttachmentState)(nil)).Elem()
}

type havipAttachmentArgs struct {
	// Specifies whether to forcefully disassociate the HAVIP from the ECS instance or ENI. Default value: `False`. Valid values: `True` and `False`.
	Force *string `pulumi:"force"`
	// The havipId of the havip attachment, the field can't be changed.
	HavipId string `pulumi:"havipId"`
	// The instanceId of the havip attachment, the field can't be changed.
	InstanceId string `pulumi:"instanceId"`
	// The Type of instance to bind HaVip to. Valid values: `EcsInstance` and `NetworkInterface`. When the HaVip instance is bound to a resilient NIC, the resilient NIC instance must be filled in.
	InstanceType *string `pulumi:"instanceType"`
}

// The set of arguments for constructing a HAVipAttachment resource.
type HAVipAttachmentArgs struct {
	// Specifies whether to forcefully disassociate the HAVIP from the ECS instance or ENI. Default value: `False`. Valid values: `True` and `False`.
	Force pulumi.StringPtrInput
	// The havipId of the havip attachment, the field can't be changed.
	HavipId pulumi.StringInput
	// The instanceId of the havip attachment, the field can't be changed.
	InstanceId pulumi.StringInput
	// The Type of instance to bind HaVip to. Valid values: `EcsInstance` and `NetworkInterface`. When the HaVip instance is bound to a resilient NIC, the resilient NIC instance must be filled in.
	InstanceType pulumi.StringPtrInput
}

func (HAVipAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*havipAttachmentArgs)(nil)).Elem()
}

type HAVipAttachmentInput interface {
	pulumi.Input

	ToHAVipAttachmentOutput() HAVipAttachmentOutput
	ToHAVipAttachmentOutputWithContext(ctx context.Context) HAVipAttachmentOutput
}

func (*HAVipAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**HAVipAttachment)(nil)).Elem()
}

func (i *HAVipAttachment) ToHAVipAttachmentOutput() HAVipAttachmentOutput {
	return i.ToHAVipAttachmentOutputWithContext(context.Background())
}

func (i *HAVipAttachment) ToHAVipAttachmentOutputWithContext(ctx context.Context) HAVipAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HAVipAttachmentOutput)
}

// HAVipAttachmentArrayInput is an input type that accepts HAVipAttachmentArray and HAVipAttachmentArrayOutput values.
// You can construct a concrete instance of `HAVipAttachmentArrayInput` via:
//
//	HAVipAttachmentArray{ HAVipAttachmentArgs{...} }
type HAVipAttachmentArrayInput interface {
	pulumi.Input

	ToHAVipAttachmentArrayOutput() HAVipAttachmentArrayOutput
	ToHAVipAttachmentArrayOutputWithContext(context.Context) HAVipAttachmentArrayOutput
}

type HAVipAttachmentArray []HAVipAttachmentInput

func (HAVipAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HAVipAttachment)(nil)).Elem()
}

func (i HAVipAttachmentArray) ToHAVipAttachmentArrayOutput() HAVipAttachmentArrayOutput {
	return i.ToHAVipAttachmentArrayOutputWithContext(context.Background())
}

func (i HAVipAttachmentArray) ToHAVipAttachmentArrayOutputWithContext(ctx context.Context) HAVipAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HAVipAttachmentArrayOutput)
}

// HAVipAttachmentMapInput is an input type that accepts HAVipAttachmentMap and HAVipAttachmentMapOutput values.
// You can construct a concrete instance of `HAVipAttachmentMapInput` via:
//
//	HAVipAttachmentMap{ "key": HAVipAttachmentArgs{...} }
type HAVipAttachmentMapInput interface {
	pulumi.Input

	ToHAVipAttachmentMapOutput() HAVipAttachmentMapOutput
	ToHAVipAttachmentMapOutputWithContext(context.Context) HAVipAttachmentMapOutput
}

type HAVipAttachmentMap map[string]HAVipAttachmentInput

func (HAVipAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HAVipAttachment)(nil)).Elem()
}

func (i HAVipAttachmentMap) ToHAVipAttachmentMapOutput() HAVipAttachmentMapOutput {
	return i.ToHAVipAttachmentMapOutputWithContext(context.Background())
}

func (i HAVipAttachmentMap) ToHAVipAttachmentMapOutputWithContext(ctx context.Context) HAVipAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HAVipAttachmentMapOutput)
}

type HAVipAttachmentOutput struct{ *pulumi.OutputState }

func (HAVipAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HAVipAttachment)(nil)).Elem()
}

func (o HAVipAttachmentOutput) ToHAVipAttachmentOutput() HAVipAttachmentOutput {
	return o
}

func (o HAVipAttachmentOutput) ToHAVipAttachmentOutputWithContext(ctx context.Context) HAVipAttachmentOutput {
	return o
}

// Specifies whether to forcefully disassociate the HAVIP from the ECS instance or ENI. Default value: `False`. Valid values: `True` and `False`.
func (o HAVipAttachmentOutput) Force() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HAVipAttachment) pulumi.StringPtrOutput { return v.Force }).(pulumi.StringPtrOutput)
}

// The havipId of the havip attachment, the field can't be changed.
func (o HAVipAttachmentOutput) HavipId() pulumi.StringOutput {
	return o.ApplyT(func(v *HAVipAttachment) pulumi.StringOutput { return v.HavipId }).(pulumi.StringOutput)
}

// The instanceId of the havip attachment, the field can't be changed.
func (o HAVipAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *HAVipAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The Type of instance to bind HaVip to. Valid values: `EcsInstance` and `NetworkInterface`. When the HaVip instance is bound to a resilient NIC, the resilient NIC instance must be filled in.
func (o HAVipAttachmentOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *HAVipAttachment) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// (Available in v1.201.0+) The status of the HaVip instance.
func (o HAVipAttachmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *HAVipAttachment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type HAVipAttachmentArrayOutput struct{ *pulumi.OutputState }

func (HAVipAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HAVipAttachment)(nil)).Elem()
}

func (o HAVipAttachmentArrayOutput) ToHAVipAttachmentArrayOutput() HAVipAttachmentArrayOutput {
	return o
}

func (o HAVipAttachmentArrayOutput) ToHAVipAttachmentArrayOutputWithContext(ctx context.Context) HAVipAttachmentArrayOutput {
	return o
}

func (o HAVipAttachmentArrayOutput) Index(i pulumi.IntInput) HAVipAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HAVipAttachment {
		return vs[0].([]*HAVipAttachment)[vs[1].(int)]
	}).(HAVipAttachmentOutput)
}

type HAVipAttachmentMapOutput struct{ *pulumi.OutputState }

func (HAVipAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HAVipAttachment)(nil)).Elem()
}

func (o HAVipAttachmentMapOutput) ToHAVipAttachmentMapOutput() HAVipAttachmentMapOutput {
	return o
}

func (o HAVipAttachmentMapOutput) ToHAVipAttachmentMapOutputWithContext(ctx context.Context) HAVipAttachmentMapOutput {
	return o
}

func (o HAVipAttachmentMapOutput) MapIndex(k pulumi.StringInput) HAVipAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HAVipAttachment {
		return vs[0].(map[string]*HAVipAttachment)[vs[1].(string)]
	}).(HAVipAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HAVipAttachmentInput)(nil)).Elem(), &HAVipAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*HAVipAttachmentArrayInput)(nil)).Elem(), HAVipAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HAVipAttachmentMapInput)(nil)).Elem(), HAVipAttachmentMap{})
	pulumi.RegisterOutputType(HAVipAttachmentOutput{})
	pulumi.RegisterOutputType(HAVipAttachmentArrayOutput{})
	pulumi.RegisterOutputType(HAVipAttachmentMapOutput{})
}
