// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudconnect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Connect Network Grant resource. If the CEN instance to be attached belongs to another account, authorization by the CEN instance is required.
//
// For information about Cloud Connect Network Grant and how to use it, see [What is Cloud Connect Network Grant](https://www.alibabacloud.com/help/doc-detail/94543.htm).
//
// > **NOTE:** Available in 1.63.0+
//
// > **NOTE:** Only the following regions support create Cloud Connect Network Grant. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudconnect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := alicloud.NewProvider(ctx, "ccnAccount", nil)
//			if err != nil {
//				return err
//			}
//			_, err = alicloud.NewProvider(ctx, "cenAccount", &alicloud.ProviderArgs{
//				Region:    pulumi.String("cn-hangzhou"),
//				AccessKey: pulumi.String("xxxxxx"),
//				SecretKey: pulumi.String("xxxxxx"),
//			})
//			if err != nil {
//				return err
//			}
//			cen, err := cen.NewInstance(ctx, "cen", nil, pulumi.Provider(alicloud.Cen_account))
//			if err != nil {
//				return err
//			}
//			ccn, err := cloudconnect.NewNetwork(ctx, "ccn", &cloudconnect.NetworkArgs{
//				IsDefault: pulumi.Bool(true),
//			}, pulumi.Provider(alicloud.Ccn_account))
//			if err != nil {
//				return err
//			}
//			_, err = cloudconnect.NewNetworkGrant(ctx, "default", &cloudconnect.NetworkGrantArgs{
//				CcnId:  ccn.ID(),
//				CenId:  cen.ID(),
//				CenUid: pulumi.String("xxxxxx"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				ccn,
//				cen,
//			}))
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
// The Cloud Connect Network Grant can be imported using the instance_id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cloudconnect/networkGrant:NetworkGrant example ccn-abc123456:cen-abc123456
//
// ```
type NetworkGrant struct {
	pulumi.CustomResourceState

	// The ID of the CCN instance.
	CcnId pulumi.StringOutput `pulumi:"ccnId"`
	// The ID of the CEN instance.
	CenId pulumi.StringOutput `pulumi:"cenId"`
	// The ID of the account to which the CEN instance belongs.
	CenUid pulumi.StringOutput `pulumi:"cenUid"`
}

// NewNetworkGrant registers a new resource with the given unique name, arguments, and options.
func NewNetworkGrant(ctx *pulumi.Context,
	name string, args *NetworkGrantArgs, opts ...pulumi.ResourceOption) (*NetworkGrant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CcnId == nil {
		return nil, errors.New("invalid value for required argument 'CcnId'")
	}
	if args.CenId == nil {
		return nil, errors.New("invalid value for required argument 'CenId'")
	}
	if args.CenUid == nil {
		return nil, errors.New("invalid value for required argument 'CenUid'")
	}
	var resource NetworkGrant
	err := ctx.RegisterResource("alicloud:cloudconnect/networkGrant:NetworkGrant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkGrant gets an existing NetworkGrant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkGrantState, opts ...pulumi.ResourceOption) (*NetworkGrant, error) {
	var resource NetworkGrant
	err := ctx.ReadResource("alicloud:cloudconnect/networkGrant:NetworkGrant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkGrant resources.
type networkGrantState struct {
	// The ID of the CCN instance.
	CcnId *string `pulumi:"ccnId"`
	// The ID of the CEN instance.
	CenId *string `pulumi:"cenId"`
	// The ID of the account to which the CEN instance belongs.
	CenUid *string `pulumi:"cenUid"`
}

type NetworkGrantState struct {
	// The ID of the CCN instance.
	CcnId pulumi.StringPtrInput
	// The ID of the CEN instance.
	CenId pulumi.StringPtrInput
	// The ID of the account to which the CEN instance belongs.
	CenUid pulumi.StringPtrInput
}

func (NetworkGrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkGrantState)(nil)).Elem()
}

type networkGrantArgs struct {
	// The ID of the CCN instance.
	CcnId string `pulumi:"ccnId"`
	// The ID of the CEN instance.
	CenId string `pulumi:"cenId"`
	// The ID of the account to which the CEN instance belongs.
	CenUid string `pulumi:"cenUid"`
}

// The set of arguments for constructing a NetworkGrant resource.
type NetworkGrantArgs struct {
	// The ID of the CCN instance.
	CcnId pulumi.StringInput
	// The ID of the CEN instance.
	CenId pulumi.StringInput
	// The ID of the account to which the CEN instance belongs.
	CenUid pulumi.StringInput
}

func (NetworkGrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkGrantArgs)(nil)).Elem()
}

type NetworkGrantInput interface {
	pulumi.Input

	ToNetworkGrantOutput() NetworkGrantOutput
	ToNetworkGrantOutputWithContext(ctx context.Context) NetworkGrantOutput
}

func (*NetworkGrant) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkGrant)(nil)).Elem()
}

func (i *NetworkGrant) ToNetworkGrantOutput() NetworkGrantOutput {
	return i.ToNetworkGrantOutputWithContext(context.Background())
}

func (i *NetworkGrant) ToNetworkGrantOutputWithContext(ctx context.Context) NetworkGrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkGrantOutput)
}

// NetworkGrantArrayInput is an input type that accepts NetworkGrantArray and NetworkGrantArrayOutput values.
// You can construct a concrete instance of `NetworkGrantArrayInput` via:
//
//	NetworkGrantArray{ NetworkGrantArgs{...} }
type NetworkGrantArrayInput interface {
	pulumi.Input

	ToNetworkGrantArrayOutput() NetworkGrantArrayOutput
	ToNetworkGrantArrayOutputWithContext(context.Context) NetworkGrantArrayOutput
}

type NetworkGrantArray []NetworkGrantInput

func (NetworkGrantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkGrant)(nil)).Elem()
}

func (i NetworkGrantArray) ToNetworkGrantArrayOutput() NetworkGrantArrayOutput {
	return i.ToNetworkGrantArrayOutputWithContext(context.Background())
}

func (i NetworkGrantArray) ToNetworkGrantArrayOutputWithContext(ctx context.Context) NetworkGrantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkGrantArrayOutput)
}

// NetworkGrantMapInput is an input type that accepts NetworkGrantMap and NetworkGrantMapOutput values.
// You can construct a concrete instance of `NetworkGrantMapInput` via:
//
//	NetworkGrantMap{ "key": NetworkGrantArgs{...} }
type NetworkGrantMapInput interface {
	pulumi.Input

	ToNetworkGrantMapOutput() NetworkGrantMapOutput
	ToNetworkGrantMapOutputWithContext(context.Context) NetworkGrantMapOutput
}

type NetworkGrantMap map[string]NetworkGrantInput

func (NetworkGrantMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkGrant)(nil)).Elem()
}

func (i NetworkGrantMap) ToNetworkGrantMapOutput() NetworkGrantMapOutput {
	return i.ToNetworkGrantMapOutputWithContext(context.Background())
}

func (i NetworkGrantMap) ToNetworkGrantMapOutputWithContext(ctx context.Context) NetworkGrantMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkGrantMapOutput)
}

type NetworkGrantOutput struct{ *pulumi.OutputState }

func (NetworkGrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkGrant)(nil)).Elem()
}

func (o NetworkGrantOutput) ToNetworkGrantOutput() NetworkGrantOutput {
	return o
}

func (o NetworkGrantOutput) ToNetworkGrantOutputWithContext(ctx context.Context) NetworkGrantOutput {
	return o
}

// The ID of the CCN instance.
func (o NetworkGrantOutput) CcnId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkGrant) pulumi.StringOutput { return v.CcnId }).(pulumi.StringOutput)
}

// The ID of the CEN instance.
func (o NetworkGrantOutput) CenId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkGrant) pulumi.StringOutput { return v.CenId }).(pulumi.StringOutput)
}

// The ID of the account to which the CEN instance belongs.
func (o NetworkGrantOutput) CenUid() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkGrant) pulumi.StringOutput { return v.CenUid }).(pulumi.StringOutput)
}

type NetworkGrantArrayOutput struct{ *pulumi.OutputState }

func (NetworkGrantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkGrant)(nil)).Elem()
}

func (o NetworkGrantArrayOutput) ToNetworkGrantArrayOutput() NetworkGrantArrayOutput {
	return o
}

func (o NetworkGrantArrayOutput) ToNetworkGrantArrayOutputWithContext(ctx context.Context) NetworkGrantArrayOutput {
	return o
}

func (o NetworkGrantArrayOutput) Index(i pulumi.IntInput) NetworkGrantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkGrant {
		return vs[0].([]*NetworkGrant)[vs[1].(int)]
	}).(NetworkGrantOutput)
}

type NetworkGrantMapOutput struct{ *pulumi.OutputState }

func (NetworkGrantMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkGrant)(nil)).Elem()
}

func (o NetworkGrantMapOutput) ToNetworkGrantMapOutput() NetworkGrantMapOutput {
	return o
}

func (o NetworkGrantMapOutput) ToNetworkGrantMapOutputWithContext(ctx context.Context) NetworkGrantMapOutput {
	return o
}

func (o NetworkGrantMapOutput) MapIndex(k pulumi.StringInput) NetworkGrantOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkGrant {
		return vs[0].(map[string]*NetworkGrant)[vs[1].(string)]
	}).(NetworkGrantOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkGrantInput)(nil)).Elem(), &NetworkGrant{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkGrantArrayInput)(nil)).Elem(), NetworkGrantArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkGrantMapInput)(nil)).Elem(), NetworkGrantMap{})
	pulumi.RegisterOutputType(NetworkGrantOutput{})
	pulumi.RegisterOutputType(NetworkGrantArrayOutput{})
	pulumi.RegisterOutputType(NetworkGrantMapOutput{})
}
