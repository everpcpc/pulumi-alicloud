// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Peer Connection resource.
//
// For information about VPC Peer Connection and how to use it, see [What is Peer Connection](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createvpcpeer).
//
// > **NOTE:** Available in v1.186.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultAccount, err := alicloud.GetAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetworks, err := vpc.GetNetworks(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewPeerConnection(ctx, "defaultPeerConnection", &vpc.PeerConnectionArgs{
//				PeerConnectionName: pulumi.Any(_var.Name),
//				VpcId:              *pulumi.String(defaultNetworks.Ids[0]),
//				AcceptingAliUid:    *pulumi.String(defaultAccount.Id),
//				AcceptingRegionId:  pulumi.String("cn-hangzhou"),
//				AcceptingVpcId:     *pulumi.String(defaultNetworks.Ids[1]),
//				Description:        pulumi.Any(_var.Name),
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
// VPC Peer Connection can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:vpc/peerConnection:PeerConnection example <id>
//
// ```
type PeerConnection struct {
	pulumi.CustomResourceState

	// The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.
	// - Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.
	// - Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.
	// - If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
	AcceptingAliUid pulumi.IntOutput `pulumi:"acceptingAliUid"`
	// The region ID of the recipient of the VPC peering connection to be created.
	// - When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.
	// - When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
	AcceptingRegionId pulumi.StringOutput `pulumi:"acceptingRegionId"`
	// The VPC ID of the receiving end of the VPC peer connection.
	AcceptingVpcId pulumi.StringOutput `pulumi:"acceptingVpcId"`
	// The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// The description of the VPC peer connection to be created. It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The dry run.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The name of the resource. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
	PeerConnectionName pulumi.StringPtrOutput `pulumi:"peerConnectionName"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the requester VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewPeerConnection registers a new resource with the given unique name, arguments, and options.
func NewPeerConnection(ctx *pulumi.Context,
	name string, args *PeerConnectionArgs, opts ...pulumi.ResourceOption) (*PeerConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AcceptingAliUid == nil {
		return nil, errors.New("invalid value for required argument 'AcceptingAliUid'")
	}
	if args.AcceptingRegionId == nil {
		return nil, errors.New("invalid value for required argument 'AcceptingRegionId'")
	}
	if args.AcceptingVpcId == nil {
		return nil, errors.New("invalid value for required argument 'AcceptingVpcId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource PeerConnection
	err := ctx.RegisterResource("alicloud:vpc/peerConnection:PeerConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeerConnection gets an existing PeerConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeerConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeerConnectionState, opts ...pulumi.ResourceOption) (*PeerConnection, error) {
	var resource PeerConnection
	err := ctx.ReadResource("alicloud:vpc/peerConnection:PeerConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PeerConnection resources.
type peerConnectionState struct {
	// The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.
	// - Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.
	// - Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.
	// - If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
	AcceptingAliUid *int `pulumi:"acceptingAliUid"`
	// The region ID of the recipient of the VPC peering connection to be created.
	// - When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.
	// - When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
	AcceptingRegionId *string `pulumi:"acceptingRegionId"`
	// The VPC ID of the receiving end of the VPC peer connection.
	AcceptingVpcId *string `pulumi:"acceptingVpcId"`
	// The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
	Bandwidth *int `pulumi:"bandwidth"`
	// The description of the VPC peer connection to be created. It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
	Description *string `pulumi:"description"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The name of the resource. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
	PeerConnectionName *string `pulumi:"peerConnectionName"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// The ID of the requester VPC.
	VpcId *string `pulumi:"vpcId"`
}

type PeerConnectionState struct {
	// The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.
	// - Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.
	// - Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.
	// - If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
	AcceptingAliUid pulumi.IntPtrInput
	// The region ID of the recipient of the VPC peering connection to be created.
	// - When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.
	// - When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
	AcceptingRegionId pulumi.StringPtrInput
	// The VPC ID of the receiving end of the VPC peer connection.
	AcceptingVpcId pulumi.StringPtrInput
	// The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
	Bandwidth pulumi.IntPtrInput
	// The description of the VPC peer connection to be created. It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
	Description pulumi.StringPtrInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The name of the resource. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
	PeerConnectionName pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// The ID of the requester VPC.
	VpcId pulumi.StringPtrInput
}

func (PeerConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*peerConnectionState)(nil)).Elem()
}

type peerConnectionArgs struct {
	// The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.
	// - Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.
	// - Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.
	// - If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
	AcceptingAliUid int `pulumi:"acceptingAliUid"`
	// The region ID of the recipient of the VPC peering connection to be created.
	// - When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.
	// - When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
	AcceptingRegionId string `pulumi:"acceptingRegionId"`
	// The VPC ID of the receiving end of the VPC peer connection.
	AcceptingVpcId string `pulumi:"acceptingVpcId"`
	// The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
	Bandwidth *int `pulumi:"bandwidth"`
	// The description of the VPC peer connection to be created. It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
	Description *string `pulumi:"description"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The name of the resource. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
	PeerConnectionName *string `pulumi:"peerConnectionName"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// The ID of the requester VPC.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a PeerConnection resource.
type PeerConnectionArgs struct {
	// The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.
	// - Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.
	// - Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.
	// - If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
	AcceptingAliUid pulumi.IntInput
	// The region ID of the recipient of the VPC peering connection to be created.
	// - When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.
	// - When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
	AcceptingRegionId pulumi.StringInput
	// The VPC ID of the receiving end of the VPC peer connection.
	AcceptingVpcId pulumi.StringInput
	// The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
	Bandwidth pulumi.IntPtrInput
	// The description of the VPC peer connection to be created. It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
	Description pulumi.StringPtrInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The name of the resource. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
	PeerConnectionName pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// The ID of the requester VPC.
	VpcId pulumi.StringInput
}

func (PeerConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peerConnectionArgs)(nil)).Elem()
}

type PeerConnectionInput interface {
	pulumi.Input

	ToPeerConnectionOutput() PeerConnectionOutput
	ToPeerConnectionOutputWithContext(ctx context.Context) PeerConnectionOutput
}

func (*PeerConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**PeerConnection)(nil)).Elem()
}

func (i *PeerConnection) ToPeerConnectionOutput() PeerConnectionOutput {
	return i.ToPeerConnectionOutputWithContext(context.Background())
}

func (i *PeerConnection) ToPeerConnectionOutputWithContext(ctx context.Context) PeerConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeerConnectionOutput)
}

// PeerConnectionArrayInput is an input type that accepts PeerConnectionArray and PeerConnectionArrayOutput values.
// You can construct a concrete instance of `PeerConnectionArrayInput` via:
//
//	PeerConnectionArray{ PeerConnectionArgs{...} }
type PeerConnectionArrayInput interface {
	pulumi.Input

	ToPeerConnectionArrayOutput() PeerConnectionArrayOutput
	ToPeerConnectionArrayOutputWithContext(context.Context) PeerConnectionArrayOutput
}

type PeerConnectionArray []PeerConnectionInput

func (PeerConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PeerConnection)(nil)).Elem()
}

func (i PeerConnectionArray) ToPeerConnectionArrayOutput() PeerConnectionArrayOutput {
	return i.ToPeerConnectionArrayOutputWithContext(context.Background())
}

func (i PeerConnectionArray) ToPeerConnectionArrayOutputWithContext(ctx context.Context) PeerConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeerConnectionArrayOutput)
}

// PeerConnectionMapInput is an input type that accepts PeerConnectionMap and PeerConnectionMapOutput values.
// You can construct a concrete instance of `PeerConnectionMapInput` via:
//
//	PeerConnectionMap{ "key": PeerConnectionArgs{...} }
type PeerConnectionMapInput interface {
	pulumi.Input

	ToPeerConnectionMapOutput() PeerConnectionMapOutput
	ToPeerConnectionMapOutputWithContext(context.Context) PeerConnectionMapOutput
}

type PeerConnectionMap map[string]PeerConnectionInput

func (PeerConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PeerConnection)(nil)).Elem()
}

func (i PeerConnectionMap) ToPeerConnectionMapOutput() PeerConnectionMapOutput {
	return i.ToPeerConnectionMapOutputWithContext(context.Background())
}

func (i PeerConnectionMap) ToPeerConnectionMapOutputWithContext(ctx context.Context) PeerConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeerConnectionMapOutput)
}

type PeerConnectionOutput struct{ *pulumi.OutputState }

func (PeerConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeerConnection)(nil)).Elem()
}

func (o PeerConnectionOutput) ToPeerConnectionOutput() PeerConnectionOutput {
	return o
}

func (o PeerConnectionOutput) ToPeerConnectionOutputWithContext(ctx context.Context) PeerConnectionOutput {
	return o
}

// The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.
// - Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.
// - Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.
// - If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
func (o PeerConnectionOutput) AcceptingAliUid() pulumi.IntOutput {
	return o.ApplyT(func(v *PeerConnection) pulumi.IntOutput { return v.AcceptingAliUid }).(pulumi.IntOutput)
}

// The region ID of the recipient of the VPC peering connection to be created.
// - When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.
// - When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
func (o PeerConnectionOutput) AcceptingRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerConnection) pulumi.StringOutput { return v.AcceptingRegionId }).(pulumi.StringOutput)
}

// The VPC ID of the receiving end of the VPC peer connection.
func (o PeerConnectionOutput) AcceptingVpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerConnection) pulumi.StringOutput { return v.AcceptingVpcId }).(pulumi.StringOutput)
}

// The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
func (o PeerConnectionOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *PeerConnection) pulumi.IntOutput { return v.Bandwidth }).(pulumi.IntOutput)
}

// The description of the VPC peer connection to be created. It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
func (o PeerConnectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeerConnection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The dry run.
func (o PeerConnectionOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PeerConnection) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The name of the resource. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
func (o PeerConnectionOutput) PeerConnectionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeerConnection) pulumi.StringPtrOutput { return v.PeerConnectionName }).(pulumi.StringPtrOutput)
}

// The status of the resource.
func (o PeerConnectionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerConnection) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the requester VPC.
func (o PeerConnectionOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerConnection) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type PeerConnectionArrayOutput struct{ *pulumi.OutputState }

func (PeerConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PeerConnection)(nil)).Elem()
}

func (o PeerConnectionArrayOutput) ToPeerConnectionArrayOutput() PeerConnectionArrayOutput {
	return o
}

func (o PeerConnectionArrayOutput) ToPeerConnectionArrayOutputWithContext(ctx context.Context) PeerConnectionArrayOutput {
	return o
}

func (o PeerConnectionArrayOutput) Index(i pulumi.IntInput) PeerConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PeerConnection {
		return vs[0].([]*PeerConnection)[vs[1].(int)]
	}).(PeerConnectionOutput)
}

type PeerConnectionMapOutput struct{ *pulumi.OutputState }

func (PeerConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PeerConnection)(nil)).Elem()
}

func (o PeerConnectionMapOutput) ToPeerConnectionMapOutput() PeerConnectionMapOutput {
	return o
}

func (o PeerConnectionMapOutput) ToPeerConnectionMapOutputWithContext(ctx context.Context) PeerConnectionMapOutput {
	return o
}

func (o PeerConnectionMapOutput) MapIndex(k pulumi.StringInput) PeerConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PeerConnection {
		return vs[0].(map[string]*PeerConnection)[vs[1].(string)]
	}).(PeerConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PeerConnectionInput)(nil)).Elem(), &PeerConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeerConnectionArrayInput)(nil)).Elem(), PeerConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeerConnectionMapInput)(nil)).Elem(), PeerConnectionMap{})
	pulumi.RegisterOutputType(PeerConnectionOutput{})
	pulumi.RegisterOutputType(PeerConnectionArrayOutput{})
	pulumi.RegisterOutputType(PeerConnectionMapOutput{})
}
