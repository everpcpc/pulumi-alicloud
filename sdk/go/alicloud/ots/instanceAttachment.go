// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ots

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource will help you to bind a VPC to an OTS instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ots"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		fooInstance, err := ots.NewInstance(ctx, "fooInstance", &ots.InstanceArgs{
// 			Description: pulumi.String("for table"),
// 			AccessedBy:  pulumi.String("Vpc"),
// 			Tags: pulumi.AnyMap{
// 				"Created": pulumi.Any("TF"),
// 				"For":     pulumi.Any("Building table"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fooZones, err := alicloud.GetZones(ctx, &GetZonesArgs{
// 			AvailableResourceCreation: pulumi.StringRef("VSwitch"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		fooNetwork, err := vpc.NewNetwork(ctx, "fooNetwork", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fooSwitch, err := vpc.NewSwitch(ctx, "fooSwitch", &vpc.SwitchArgs{
// 			VpcId:       fooNetwork.ID(),
// 			VswitchName: pulumi.String("for-ots-instance"),
// 			CidrBlock:   pulumi.String("172.16.1.0/24"),
// 			ZoneId:      pulumi.String(fooZones.Zones[0].Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ots.NewInstanceAttachment(ctx, "fooInstanceAttachment", &ots.InstanceAttachmentArgs{
// 			InstanceName: fooInstance.Name,
// 			VpcName:      pulumi.String("attachment1"),
// 			VswitchId:    fooSwitch.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type InstanceAttachment struct {
	pulumi.CustomResourceState

	// The name of the OTS instance.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The ID of attaching VPC to instance.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The name of attaching VPC to instance.
	VpcName pulumi.StringOutput `pulumi:"vpcName"`
	// The ID of attaching VSwitch to instance.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewInstanceAttachment registers a new resource with the given unique name, arguments, and options.
func NewInstanceAttachment(ctx *pulumi.Context,
	name string, args *InstanceAttachmentArgs, opts ...pulumi.ResourceOption) (*InstanceAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.VpcName == nil {
		return nil, errors.New("invalid value for required argument 'VpcName'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	var resource InstanceAttachment
	err := ctx.RegisterResource("alicloud:ots/instanceAttachment:InstanceAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceAttachment gets an existing InstanceAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceAttachmentState, opts ...pulumi.ResourceOption) (*InstanceAttachment, error) {
	var resource InstanceAttachment
	err := ctx.ReadResource("alicloud:ots/instanceAttachment:InstanceAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceAttachment resources.
type instanceAttachmentState struct {
	// The name of the OTS instance.
	InstanceName *string `pulumi:"instanceName"`
	// The ID of attaching VPC to instance.
	VpcId *string `pulumi:"vpcId"`
	// The name of attaching VPC to instance.
	VpcName *string `pulumi:"vpcName"`
	// The ID of attaching VSwitch to instance.
	VswitchId *string `pulumi:"vswitchId"`
}

type InstanceAttachmentState struct {
	// The name of the OTS instance.
	InstanceName pulumi.StringPtrInput
	// The ID of attaching VPC to instance.
	VpcId pulumi.StringPtrInput
	// The name of attaching VPC to instance.
	VpcName pulumi.StringPtrInput
	// The ID of attaching VSwitch to instance.
	VswitchId pulumi.StringPtrInput
}

func (InstanceAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceAttachmentState)(nil)).Elem()
}

type instanceAttachmentArgs struct {
	// The name of the OTS instance.
	InstanceName string `pulumi:"instanceName"`
	// The name of attaching VPC to instance.
	VpcName string `pulumi:"vpcName"`
	// The ID of attaching VSwitch to instance.
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a InstanceAttachment resource.
type InstanceAttachmentArgs struct {
	// The name of the OTS instance.
	InstanceName pulumi.StringInput
	// The name of attaching VPC to instance.
	VpcName pulumi.StringInput
	// The ID of attaching VSwitch to instance.
	VswitchId pulumi.StringInput
}

func (InstanceAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceAttachmentArgs)(nil)).Elem()
}

type InstanceAttachmentInput interface {
	pulumi.Input

	ToInstanceAttachmentOutput() InstanceAttachmentOutput
	ToInstanceAttachmentOutputWithContext(ctx context.Context) InstanceAttachmentOutput
}

func (*InstanceAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceAttachment)(nil)).Elem()
}

func (i *InstanceAttachment) ToInstanceAttachmentOutput() InstanceAttachmentOutput {
	return i.ToInstanceAttachmentOutputWithContext(context.Background())
}

func (i *InstanceAttachment) ToInstanceAttachmentOutputWithContext(ctx context.Context) InstanceAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceAttachmentOutput)
}

// InstanceAttachmentArrayInput is an input type that accepts InstanceAttachmentArray and InstanceAttachmentArrayOutput values.
// You can construct a concrete instance of `InstanceAttachmentArrayInput` via:
//
//          InstanceAttachmentArray{ InstanceAttachmentArgs{...} }
type InstanceAttachmentArrayInput interface {
	pulumi.Input

	ToInstanceAttachmentArrayOutput() InstanceAttachmentArrayOutput
	ToInstanceAttachmentArrayOutputWithContext(context.Context) InstanceAttachmentArrayOutput
}

type InstanceAttachmentArray []InstanceAttachmentInput

func (InstanceAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceAttachment)(nil)).Elem()
}

func (i InstanceAttachmentArray) ToInstanceAttachmentArrayOutput() InstanceAttachmentArrayOutput {
	return i.ToInstanceAttachmentArrayOutputWithContext(context.Background())
}

func (i InstanceAttachmentArray) ToInstanceAttachmentArrayOutputWithContext(ctx context.Context) InstanceAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceAttachmentArrayOutput)
}

// InstanceAttachmentMapInput is an input type that accepts InstanceAttachmentMap and InstanceAttachmentMapOutput values.
// You can construct a concrete instance of `InstanceAttachmentMapInput` via:
//
//          InstanceAttachmentMap{ "key": InstanceAttachmentArgs{...} }
type InstanceAttachmentMapInput interface {
	pulumi.Input

	ToInstanceAttachmentMapOutput() InstanceAttachmentMapOutput
	ToInstanceAttachmentMapOutputWithContext(context.Context) InstanceAttachmentMapOutput
}

type InstanceAttachmentMap map[string]InstanceAttachmentInput

func (InstanceAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceAttachment)(nil)).Elem()
}

func (i InstanceAttachmentMap) ToInstanceAttachmentMapOutput() InstanceAttachmentMapOutput {
	return i.ToInstanceAttachmentMapOutputWithContext(context.Background())
}

func (i InstanceAttachmentMap) ToInstanceAttachmentMapOutputWithContext(ctx context.Context) InstanceAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceAttachmentMapOutput)
}

type InstanceAttachmentOutput struct{ *pulumi.OutputState }

func (InstanceAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceAttachment)(nil)).Elem()
}

func (o InstanceAttachmentOutput) ToInstanceAttachmentOutput() InstanceAttachmentOutput {
	return o
}

func (o InstanceAttachmentOutput) ToInstanceAttachmentOutputWithContext(ctx context.Context) InstanceAttachmentOutput {
	return o
}

type InstanceAttachmentArrayOutput struct{ *pulumi.OutputState }

func (InstanceAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceAttachment)(nil)).Elem()
}

func (o InstanceAttachmentArrayOutput) ToInstanceAttachmentArrayOutput() InstanceAttachmentArrayOutput {
	return o
}

func (o InstanceAttachmentArrayOutput) ToInstanceAttachmentArrayOutputWithContext(ctx context.Context) InstanceAttachmentArrayOutput {
	return o
}

func (o InstanceAttachmentArrayOutput) Index(i pulumi.IntInput) InstanceAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceAttachment {
		return vs[0].([]*InstanceAttachment)[vs[1].(int)]
	}).(InstanceAttachmentOutput)
}

type InstanceAttachmentMapOutput struct{ *pulumi.OutputState }

func (InstanceAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceAttachment)(nil)).Elem()
}

func (o InstanceAttachmentMapOutput) ToInstanceAttachmentMapOutput() InstanceAttachmentMapOutput {
	return o
}

func (o InstanceAttachmentMapOutput) ToInstanceAttachmentMapOutputWithContext(ctx context.Context) InstanceAttachmentMapOutput {
	return o
}

func (o InstanceAttachmentMapOutput) MapIndex(k pulumi.StringInput) InstanceAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceAttachment {
		return vs[0].(map[string]*InstanceAttachment)[vs[1].(string)]
	}).(InstanceAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceAttachmentInput)(nil)).Elem(), &InstanceAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceAttachmentArrayInput)(nil)).Elem(), InstanceAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceAttachmentMapInput)(nil)).Elem(), InstanceAttachmentMap{})
	pulumi.RegisterOutputType(InstanceAttachmentOutput{})
	pulumi.RegisterOutputType(InstanceAttachmentArrayOutput{})
	pulumi.RegisterOutputType(InstanceAttachmentMapOutput{})
}
