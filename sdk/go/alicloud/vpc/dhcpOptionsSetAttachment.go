// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Dhcp Options Set Attachment resource.
//
// For information about VPC Dhcp Options Set and how to use it, see [What is Dhcp Options Set](https://www.alibabacloud.com/help/doc-detail/174112.htm).
//
// > **NOTE:** Available in v1.153.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleNetwork, err := vpc.NewNetwork(ctx, "exampleNetwork", &vpc.NetworkArgs{
// 			VpcName:   pulumi.String("test"),
// 			CidrBlock: pulumi.String("172.16.0.0/12"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDhcpOptionsSet, err := vpc.NewDhcpOptionsSet(ctx, "exampleDhcpOptionsSet", &vpc.DhcpOptionsSetArgs{
// 			DhcpOptionsSetName:        pulumi.String("example_value"),
// 			DhcpOptionsSetDescription: pulumi.String("example_value"),
// 			DomainName:                pulumi.String("example.com"),
// 			DomainNameServers:         pulumi.String("100.100.2.136"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vpc.NewDhcpOptionsSetAttachment(ctx, "exampleDhcpOptionsSetAttachment", &vpc.DhcpOptionsSetAttachmentArgs{
// 			VpcId:            exampleNetwork.ID(),
// 			DhcpOptionsSetId: exampleDhcpOptionsSet.ID(),
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
// VPC Dhcp Options Set Attachment can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:vpc/dhcpOptionsSetAttachment:DhcpOptionsSetAttachment example <id>
// ```
type DhcpOptionsSetAttachment struct {
	pulumi.CustomResourceState

	// The ID of the DHCP options set.
	DhcpOptionsSetId pulumi.StringOutput `pulumi:"dhcpOptionsSetId"`
	// Specifies whether to precheck this request only. Default values: `false`. Valid values:
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	Status pulumi.StringOutput  `pulumi:"status"`
	// The ID of the VPC network that is to be associated with the DHCP options set..
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewDhcpOptionsSetAttachment registers a new resource with the given unique name, arguments, and options.
func NewDhcpOptionsSetAttachment(ctx *pulumi.Context,
	name string, args *DhcpOptionsSetAttachmentArgs, opts ...pulumi.ResourceOption) (*DhcpOptionsSetAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DhcpOptionsSetId == nil {
		return nil, errors.New("invalid value for required argument 'DhcpOptionsSetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource DhcpOptionsSetAttachment
	err := ctx.RegisterResource("alicloud:vpc/dhcpOptionsSetAttachment:DhcpOptionsSetAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDhcpOptionsSetAttachment gets an existing DhcpOptionsSetAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDhcpOptionsSetAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DhcpOptionsSetAttachmentState, opts ...pulumi.ResourceOption) (*DhcpOptionsSetAttachment, error) {
	var resource DhcpOptionsSetAttachment
	err := ctx.ReadResource("alicloud:vpc/dhcpOptionsSetAttachment:DhcpOptionsSetAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DhcpOptionsSetAttachment resources.
type dhcpOptionsSetAttachmentState struct {
	// The ID of the DHCP options set.
	DhcpOptionsSetId *string `pulumi:"dhcpOptionsSetId"`
	// Specifies whether to precheck this request only. Default values: `false`. Valid values:
	DryRun *bool   `pulumi:"dryRun"`
	Status *string `pulumi:"status"`
	// The ID of the VPC network that is to be associated with the DHCP options set..
	VpcId *string `pulumi:"vpcId"`
}

type DhcpOptionsSetAttachmentState struct {
	// The ID of the DHCP options set.
	DhcpOptionsSetId pulumi.StringPtrInput
	// Specifies whether to precheck this request only. Default values: `false`. Valid values:
	DryRun pulumi.BoolPtrInput
	Status pulumi.StringPtrInput
	// The ID of the VPC network that is to be associated with the DHCP options set..
	VpcId pulumi.StringPtrInput
}

func (DhcpOptionsSetAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpOptionsSetAttachmentState)(nil)).Elem()
}

type dhcpOptionsSetAttachmentArgs struct {
	// The ID of the DHCP options set.
	DhcpOptionsSetId string `pulumi:"dhcpOptionsSetId"`
	// Specifies whether to precheck this request only. Default values: `false`. Valid values:
	DryRun *bool `pulumi:"dryRun"`
	// The ID of the VPC network that is to be associated with the DHCP options set..
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a DhcpOptionsSetAttachment resource.
type DhcpOptionsSetAttachmentArgs struct {
	// The ID of the DHCP options set.
	DhcpOptionsSetId pulumi.StringInput
	// Specifies whether to precheck this request only. Default values: `false`. Valid values:
	DryRun pulumi.BoolPtrInput
	// The ID of the VPC network that is to be associated with the DHCP options set..
	VpcId pulumi.StringInput
}

func (DhcpOptionsSetAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpOptionsSetAttachmentArgs)(nil)).Elem()
}

type DhcpOptionsSetAttachmentInput interface {
	pulumi.Input

	ToDhcpOptionsSetAttachmentOutput() DhcpOptionsSetAttachmentOutput
	ToDhcpOptionsSetAttachmentOutputWithContext(ctx context.Context) DhcpOptionsSetAttachmentOutput
}

func (*DhcpOptionsSetAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpOptionsSetAttachment)(nil)).Elem()
}

func (i *DhcpOptionsSetAttachment) ToDhcpOptionsSetAttachmentOutput() DhcpOptionsSetAttachmentOutput {
	return i.ToDhcpOptionsSetAttachmentOutputWithContext(context.Background())
}

func (i *DhcpOptionsSetAttachment) ToDhcpOptionsSetAttachmentOutputWithContext(ctx context.Context) DhcpOptionsSetAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpOptionsSetAttachmentOutput)
}

// DhcpOptionsSetAttachmentArrayInput is an input type that accepts DhcpOptionsSetAttachmentArray and DhcpOptionsSetAttachmentArrayOutput values.
// You can construct a concrete instance of `DhcpOptionsSetAttachmentArrayInput` via:
//
//          DhcpOptionsSetAttachmentArray{ DhcpOptionsSetAttachmentArgs{...} }
type DhcpOptionsSetAttachmentArrayInput interface {
	pulumi.Input

	ToDhcpOptionsSetAttachmentArrayOutput() DhcpOptionsSetAttachmentArrayOutput
	ToDhcpOptionsSetAttachmentArrayOutputWithContext(context.Context) DhcpOptionsSetAttachmentArrayOutput
}

type DhcpOptionsSetAttachmentArray []DhcpOptionsSetAttachmentInput

func (DhcpOptionsSetAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpOptionsSetAttachment)(nil)).Elem()
}

func (i DhcpOptionsSetAttachmentArray) ToDhcpOptionsSetAttachmentArrayOutput() DhcpOptionsSetAttachmentArrayOutput {
	return i.ToDhcpOptionsSetAttachmentArrayOutputWithContext(context.Background())
}

func (i DhcpOptionsSetAttachmentArray) ToDhcpOptionsSetAttachmentArrayOutputWithContext(ctx context.Context) DhcpOptionsSetAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpOptionsSetAttachmentArrayOutput)
}

// DhcpOptionsSetAttachmentMapInput is an input type that accepts DhcpOptionsSetAttachmentMap and DhcpOptionsSetAttachmentMapOutput values.
// You can construct a concrete instance of `DhcpOptionsSetAttachmentMapInput` via:
//
//          DhcpOptionsSetAttachmentMap{ "key": DhcpOptionsSetAttachmentArgs{...} }
type DhcpOptionsSetAttachmentMapInput interface {
	pulumi.Input

	ToDhcpOptionsSetAttachmentMapOutput() DhcpOptionsSetAttachmentMapOutput
	ToDhcpOptionsSetAttachmentMapOutputWithContext(context.Context) DhcpOptionsSetAttachmentMapOutput
}

type DhcpOptionsSetAttachmentMap map[string]DhcpOptionsSetAttachmentInput

func (DhcpOptionsSetAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpOptionsSetAttachment)(nil)).Elem()
}

func (i DhcpOptionsSetAttachmentMap) ToDhcpOptionsSetAttachmentMapOutput() DhcpOptionsSetAttachmentMapOutput {
	return i.ToDhcpOptionsSetAttachmentMapOutputWithContext(context.Background())
}

func (i DhcpOptionsSetAttachmentMap) ToDhcpOptionsSetAttachmentMapOutputWithContext(ctx context.Context) DhcpOptionsSetAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpOptionsSetAttachmentMapOutput)
}

type DhcpOptionsSetAttachmentOutput struct{ *pulumi.OutputState }

func (DhcpOptionsSetAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpOptionsSetAttachment)(nil)).Elem()
}

func (o DhcpOptionsSetAttachmentOutput) ToDhcpOptionsSetAttachmentOutput() DhcpOptionsSetAttachmentOutput {
	return o
}

func (o DhcpOptionsSetAttachmentOutput) ToDhcpOptionsSetAttachmentOutputWithContext(ctx context.Context) DhcpOptionsSetAttachmentOutput {
	return o
}

type DhcpOptionsSetAttachmentArrayOutput struct{ *pulumi.OutputState }

func (DhcpOptionsSetAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpOptionsSetAttachment)(nil)).Elem()
}

func (o DhcpOptionsSetAttachmentArrayOutput) ToDhcpOptionsSetAttachmentArrayOutput() DhcpOptionsSetAttachmentArrayOutput {
	return o
}

func (o DhcpOptionsSetAttachmentArrayOutput) ToDhcpOptionsSetAttachmentArrayOutputWithContext(ctx context.Context) DhcpOptionsSetAttachmentArrayOutput {
	return o
}

func (o DhcpOptionsSetAttachmentArrayOutput) Index(i pulumi.IntInput) DhcpOptionsSetAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DhcpOptionsSetAttachment {
		return vs[0].([]*DhcpOptionsSetAttachment)[vs[1].(int)]
	}).(DhcpOptionsSetAttachmentOutput)
}

type DhcpOptionsSetAttachmentMapOutput struct{ *pulumi.OutputState }

func (DhcpOptionsSetAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpOptionsSetAttachment)(nil)).Elem()
}

func (o DhcpOptionsSetAttachmentMapOutput) ToDhcpOptionsSetAttachmentMapOutput() DhcpOptionsSetAttachmentMapOutput {
	return o
}

func (o DhcpOptionsSetAttachmentMapOutput) ToDhcpOptionsSetAttachmentMapOutputWithContext(ctx context.Context) DhcpOptionsSetAttachmentMapOutput {
	return o
}

func (o DhcpOptionsSetAttachmentMapOutput) MapIndex(k pulumi.StringInput) DhcpOptionsSetAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DhcpOptionsSetAttachment {
		return vs[0].(map[string]*DhcpOptionsSetAttachment)[vs[1].(string)]
	}).(DhcpOptionsSetAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpOptionsSetAttachmentInput)(nil)).Elem(), &DhcpOptionsSetAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpOptionsSetAttachmentArrayInput)(nil)).Elem(), DhcpOptionsSetAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpOptionsSetAttachmentMapInput)(nil)).Elem(), DhcpOptionsSetAttachmentMap{})
	pulumi.RegisterOutputType(DhcpOptionsSetAttachmentOutput{})
	pulumi.RegisterOutputType(DhcpOptionsSetAttachmentArrayOutput{})
	pulumi.RegisterOutputType(DhcpOptionsSetAttachmentMapOutput{})
}