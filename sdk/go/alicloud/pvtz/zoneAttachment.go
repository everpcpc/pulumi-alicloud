// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pvtz

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// Using `vpcIds` to attach being in same region several vpc instances to a private zone
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/pvtz"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		zone, err := pvtz.NewZone(ctx, "zone", nil)
// 		if err != nil {
// 			return err
// 		}
// 		first, err := vpc.NewNetwork(ctx, "first", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/12"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		second, err := vpc.NewNetwork(ctx, "second", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pvtz.NewZoneAttachment(ctx, "zone-attachment", &pvtz.ZoneAttachmentArgs{
// 			ZoneId: zone.ID(),
// 			VpcIds: pulumi.StringArray{
// 				first.ID(),
// 				second.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Using `vpcs` to attach being in same region several vpc instances to a private zone
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/pvtz"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		zone, err := pvtz.NewZone(ctx, "zone", nil)
// 		if err != nil {
// 			return err
// 		}
// 		first, err := vpc.NewNetwork(ctx, "first", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/12"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		second, err := vpc.NewNetwork(ctx, "second", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pvtz.NewZoneAttachment(ctx, "zone-attachment", &pvtz.ZoneAttachmentArgs{
// 			ZoneId: zone.ID(),
// 			Vpcs: pvtz.ZoneAttachmentVpcArray{
// 				&pvtz.ZoneAttachmentVpcArgs{
// 					VpcId: first.ID(),
// 				},
// 				&pvtz.ZoneAttachmentVpcArgs{
// 					VpcId: second.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Using `vpcs` to attach being in different regions several vpc instances to a private zone
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/providers"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/pvtz"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		zone, err := pvtz.NewZone(ctx, "zone", nil)
// 		if err != nil {
// 			return err
// 		}
// 		first, err := vpc.NewNetwork(ctx, "first", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/12"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		second, err := vpc.NewNetwork(ctx, "second", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = providers.Newalicloud(ctx, "eu", &providers.alicloudArgs{
// 			Region: "eu-central-1",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		third, err := vpc.NewNetwork(ctx, "third", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/16"),
// 		}, pulumi.Provider(alicloud.Eu))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pvtz.NewZoneAttachment(ctx, "zone-attachment", &pvtz.ZoneAttachmentArgs{
// 			ZoneId: zone.ID(),
// 			Vpcs: pvtz.ZoneAttachmentVpcArray{
// 				&pvtz.ZoneAttachmentVpcArgs{
// 					VpcId: first.ID(),
// 				},
// 				&pvtz.ZoneAttachmentVpcArgs{
// 					VpcId: second.ID(),
// 				},
// 				&pvtz.ZoneAttachmentVpcArgs{
// 					RegionId: pulumi.String("eu-central-1"),
// 					VpcId:    third.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Private Zone attachment can be imported using the id(same with `zone_id`), e.g.
//
// ```sh
//  $ pulumi import alicloud:pvtz/zoneAttachment:ZoneAttachment example abc123456
// ```
type ZoneAttachment struct {
	pulumi.CustomResourceState

	// The language of code.
	Lang pulumi.StringPtrOutput `pulumi:"lang"`
	// The user custom IP address.
	UserClientIp pulumi.StringPtrOutput `pulumi:"userClientIp"`
	// The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].
	VpcIds pulumi.StringArrayOutput `pulumi:"vpcIds"`
	// The List of the VPC:
	Vpcs ZoneAttachmentVpcArrayOutput `pulumi:"vpcs"`
	// The name of the Private Zone Record.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewZoneAttachment registers a new resource with the given unique name, arguments, and options.
func NewZoneAttachment(ctx *pulumi.Context,
	name string, args *ZoneAttachmentArgs, opts ...pulumi.ResourceOption) (*ZoneAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	var resource ZoneAttachment
	err := ctx.RegisterResource("alicloud:pvtz/zoneAttachment:ZoneAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZoneAttachment gets an existing ZoneAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZoneAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneAttachmentState, opts ...pulumi.ResourceOption) (*ZoneAttachment, error) {
	var resource ZoneAttachment
	err := ctx.ReadResource("alicloud:pvtz/zoneAttachment:ZoneAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZoneAttachment resources.
type zoneAttachmentState struct {
	// The language of code.
	Lang *string `pulumi:"lang"`
	// The user custom IP address.
	UserClientIp *string `pulumi:"userClientIp"`
	// The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].
	VpcIds []string `pulumi:"vpcIds"`
	// The List of the VPC:
	Vpcs []ZoneAttachmentVpc `pulumi:"vpcs"`
	// The name of the Private Zone Record.
	ZoneId *string `pulumi:"zoneId"`
}

type ZoneAttachmentState struct {
	// The language of code.
	Lang pulumi.StringPtrInput
	// The user custom IP address.
	UserClientIp pulumi.StringPtrInput
	// The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].
	VpcIds pulumi.StringArrayInput
	// The List of the VPC:
	Vpcs ZoneAttachmentVpcArrayInput
	// The name of the Private Zone Record.
	ZoneId pulumi.StringPtrInput
}

func (ZoneAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneAttachmentState)(nil)).Elem()
}

type zoneAttachmentArgs struct {
	// The language of code.
	Lang *string `pulumi:"lang"`
	// The user custom IP address.
	UserClientIp *string `pulumi:"userClientIp"`
	// The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].
	VpcIds []string `pulumi:"vpcIds"`
	// The List of the VPC:
	Vpcs []ZoneAttachmentVpc `pulumi:"vpcs"`
	// The name of the Private Zone Record.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a ZoneAttachment resource.
type ZoneAttachmentArgs struct {
	// The language of code.
	Lang pulumi.StringPtrInput
	// The user custom IP address.
	UserClientIp pulumi.StringPtrInput
	// The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].
	VpcIds pulumi.StringArrayInput
	// The List of the VPC:
	Vpcs ZoneAttachmentVpcArrayInput
	// The name of the Private Zone Record.
	ZoneId pulumi.StringInput
}

func (ZoneAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneAttachmentArgs)(nil)).Elem()
}

type ZoneAttachmentInput interface {
	pulumi.Input

	ToZoneAttachmentOutput() ZoneAttachmentOutput
	ToZoneAttachmentOutputWithContext(ctx context.Context) ZoneAttachmentOutput
}

func (*ZoneAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneAttachment)(nil)).Elem()
}

func (i *ZoneAttachment) ToZoneAttachmentOutput() ZoneAttachmentOutput {
	return i.ToZoneAttachmentOutputWithContext(context.Background())
}

func (i *ZoneAttachment) ToZoneAttachmentOutputWithContext(ctx context.Context) ZoneAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneAttachmentOutput)
}

// ZoneAttachmentArrayInput is an input type that accepts ZoneAttachmentArray and ZoneAttachmentArrayOutput values.
// You can construct a concrete instance of `ZoneAttachmentArrayInput` via:
//
//          ZoneAttachmentArray{ ZoneAttachmentArgs{...} }
type ZoneAttachmentArrayInput interface {
	pulumi.Input

	ToZoneAttachmentArrayOutput() ZoneAttachmentArrayOutput
	ToZoneAttachmentArrayOutputWithContext(context.Context) ZoneAttachmentArrayOutput
}

type ZoneAttachmentArray []ZoneAttachmentInput

func (ZoneAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZoneAttachment)(nil)).Elem()
}

func (i ZoneAttachmentArray) ToZoneAttachmentArrayOutput() ZoneAttachmentArrayOutput {
	return i.ToZoneAttachmentArrayOutputWithContext(context.Background())
}

func (i ZoneAttachmentArray) ToZoneAttachmentArrayOutputWithContext(ctx context.Context) ZoneAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneAttachmentArrayOutput)
}

// ZoneAttachmentMapInput is an input type that accepts ZoneAttachmentMap and ZoneAttachmentMapOutput values.
// You can construct a concrete instance of `ZoneAttachmentMapInput` via:
//
//          ZoneAttachmentMap{ "key": ZoneAttachmentArgs{...} }
type ZoneAttachmentMapInput interface {
	pulumi.Input

	ToZoneAttachmentMapOutput() ZoneAttachmentMapOutput
	ToZoneAttachmentMapOutputWithContext(context.Context) ZoneAttachmentMapOutput
}

type ZoneAttachmentMap map[string]ZoneAttachmentInput

func (ZoneAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZoneAttachment)(nil)).Elem()
}

func (i ZoneAttachmentMap) ToZoneAttachmentMapOutput() ZoneAttachmentMapOutput {
	return i.ToZoneAttachmentMapOutputWithContext(context.Background())
}

func (i ZoneAttachmentMap) ToZoneAttachmentMapOutputWithContext(ctx context.Context) ZoneAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneAttachmentMapOutput)
}

type ZoneAttachmentOutput struct{ *pulumi.OutputState }

func (ZoneAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneAttachment)(nil)).Elem()
}

func (o ZoneAttachmentOutput) ToZoneAttachmentOutput() ZoneAttachmentOutput {
	return o
}

func (o ZoneAttachmentOutput) ToZoneAttachmentOutputWithContext(ctx context.Context) ZoneAttachmentOutput {
	return o
}

type ZoneAttachmentArrayOutput struct{ *pulumi.OutputState }

func (ZoneAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZoneAttachment)(nil)).Elem()
}

func (o ZoneAttachmentArrayOutput) ToZoneAttachmentArrayOutput() ZoneAttachmentArrayOutput {
	return o
}

func (o ZoneAttachmentArrayOutput) ToZoneAttachmentArrayOutputWithContext(ctx context.Context) ZoneAttachmentArrayOutput {
	return o
}

func (o ZoneAttachmentArrayOutput) Index(i pulumi.IntInput) ZoneAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ZoneAttachment {
		return vs[0].([]*ZoneAttachment)[vs[1].(int)]
	}).(ZoneAttachmentOutput)
}

type ZoneAttachmentMapOutput struct{ *pulumi.OutputState }

func (ZoneAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZoneAttachment)(nil)).Elem()
}

func (o ZoneAttachmentMapOutput) ToZoneAttachmentMapOutput() ZoneAttachmentMapOutput {
	return o
}

func (o ZoneAttachmentMapOutput) ToZoneAttachmentMapOutputWithContext(ctx context.Context) ZoneAttachmentMapOutput {
	return o
}

func (o ZoneAttachmentMapOutput) MapIndex(k pulumi.StringInput) ZoneAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ZoneAttachment {
		return vs[0].(map[string]*ZoneAttachment)[vs[1].(string)]
	}).(ZoneAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneAttachmentInput)(nil)).Elem(), &ZoneAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneAttachmentArrayInput)(nil)).Elem(), ZoneAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneAttachmentMapInput)(nil)).Elem(), ZoneAttachmentMap{})
	pulumi.RegisterOutputType(ZoneAttachmentOutput{})
	pulumi.RegisterOutputType(ZoneAttachmentArrayOutput{})
	pulumi.RegisterOutputType(ZoneAttachmentMapOutput{})
}
