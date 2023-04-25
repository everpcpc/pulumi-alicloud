// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Ipv6 Internet Bandwidth resource.
//
// For information about VPC Ipv6 Internet Bandwidth and how to use it, see [What is Ipv6 Internet Bandwidth](https://www.alibabacloud.com/help/doc-detail/102213.htm).
//
// > **NOTE:** Available in v1.143.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleInstances, err := ecs.GetInstances(ctx, &ecs.GetInstancesArgs{
//				NameRegex: pulumi.StringRef("ecs_with_ipv6_address"),
//				Status:    pulumi.StringRef("Running"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleIpv6Addresses, err := vpc.GetIpv6Addresses(ctx, &vpc.GetIpv6AddressesArgs{
//				AssociatedInstanceId: pulumi.StringRef(exampleInstances.Instances[0].Id),
//				Status:               pulumi.StringRef("Available"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewIpv6InternetBandwidth(ctx, "exampleIpv6InternetBandwidth", &vpc.Ipv6InternetBandwidthArgs{
//				Ipv6AddressId:      *pulumi.String(exampleIpv6Addresses.Addresses[0].Id),
//				Ipv6GatewayId:      *pulumi.String(exampleIpv6Addresses.Addresses[0].Ipv6GatewayId),
//				InternetChargeType: pulumi.String("PayByBandwidth"),
//				Bandwidth:          pulumi.Int(20),
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
// VPC Ipv6 Internet Bandwidth can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:vpc/ipv6InternetBandwidth:Ipv6InternetBandwidth example <id>
//
// ```
type Ipv6InternetBandwidth struct {
	pulumi.CustomResourceState

	// The amount of Internet bandwidth resources of the IPv6 address, Unit: `Mbit/s`. Valid values: `1` to `5000`. **NOTE:** If `internetChargeType` is set to `PayByTraffic`, the amount of Internet bandwidth resources of the IPv6 address is limited by the specification of the IPv6 gateway. `Small` (default): specifies the Free edition and the Internet bandwidth is from `1` to `500` Mbit/s. `Medium`: specifies the Medium edition and the Internet bandwidth is from `1` to `1000` Mbit/s. `Large`: specifies the Large edition and the Internet bandwidth is from `1` to `2000` Mbit/s.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// The metering method of the Internet bandwidth resources of the IPv6 gateway. Valid values: `PayByBandwidth`, `PayByTraffic`.
	InternetChargeType pulumi.StringOutput `pulumi:"internetChargeType"`
	// The ID of the IPv6 address.
	Ipv6AddressId pulumi.StringOutput `pulumi:"ipv6AddressId"`
	// The ID of the IPv6 gateway.
	Ipv6GatewayId pulumi.StringOutput `pulumi:"ipv6GatewayId"`
	// The status of the resource.Valid values:`Normal`, `FinancialLocked` and `SecurityLocked`.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewIpv6InternetBandwidth registers a new resource with the given unique name, arguments, and options.
func NewIpv6InternetBandwidth(ctx *pulumi.Context,
	name string, args *Ipv6InternetBandwidthArgs, opts ...pulumi.ResourceOption) (*Ipv6InternetBandwidth, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bandwidth == nil {
		return nil, errors.New("invalid value for required argument 'Bandwidth'")
	}
	if args.Ipv6AddressId == nil {
		return nil, errors.New("invalid value for required argument 'Ipv6AddressId'")
	}
	if args.Ipv6GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'Ipv6GatewayId'")
	}
	var resource Ipv6InternetBandwidth
	err := ctx.RegisterResource("alicloud:vpc/ipv6InternetBandwidth:Ipv6InternetBandwidth", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpv6InternetBandwidth gets an existing Ipv6InternetBandwidth resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpv6InternetBandwidth(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ipv6InternetBandwidthState, opts ...pulumi.ResourceOption) (*Ipv6InternetBandwidth, error) {
	var resource Ipv6InternetBandwidth
	err := ctx.ReadResource("alicloud:vpc/ipv6InternetBandwidth:Ipv6InternetBandwidth", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ipv6InternetBandwidth resources.
type ipv6InternetBandwidthState struct {
	// The amount of Internet bandwidth resources of the IPv6 address, Unit: `Mbit/s`. Valid values: `1` to `5000`. **NOTE:** If `internetChargeType` is set to `PayByTraffic`, the amount of Internet bandwidth resources of the IPv6 address is limited by the specification of the IPv6 gateway. `Small` (default): specifies the Free edition and the Internet bandwidth is from `1` to `500` Mbit/s. `Medium`: specifies the Medium edition and the Internet bandwidth is from `1` to `1000` Mbit/s. `Large`: specifies the Large edition and the Internet bandwidth is from `1` to `2000` Mbit/s.
	Bandwidth *int `pulumi:"bandwidth"`
	// The metering method of the Internet bandwidth resources of the IPv6 gateway. Valid values: `PayByBandwidth`, `PayByTraffic`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The ID of the IPv6 address.
	Ipv6AddressId *string `pulumi:"ipv6AddressId"`
	// The ID of the IPv6 gateway.
	Ipv6GatewayId *string `pulumi:"ipv6GatewayId"`
	// The status of the resource.Valid values:`Normal`, `FinancialLocked` and `SecurityLocked`.
	Status *string `pulumi:"status"`
}

type Ipv6InternetBandwidthState struct {
	// The amount of Internet bandwidth resources of the IPv6 address, Unit: `Mbit/s`. Valid values: `1` to `5000`. **NOTE:** If `internetChargeType` is set to `PayByTraffic`, the amount of Internet bandwidth resources of the IPv6 address is limited by the specification of the IPv6 gateway. `Small` (default): specifies the Free edition and the Internet bandwidth is from `1` to `500` Mbit/s. `Medium`: specifies the Medium edition and the Internet bandwidth is from `1` to `1000` Mbit/s. `Large`: specifies the Large edition and the Internet bandwidth is from `1` to `2000` Mbit/s.
	Bandwidth pulumi.IntPtrInput
	// The metering method of the Internet bandwidth resources of the IPv6 gateway. Valid values: `PayByBandwidth`, `PayByTraffic`.
	InternetChargeType pulumi.StringPtrInput
	// The ID of the IPv6 address.
	Ipv6AddressId pulumi.StringPtrInput
	// The ID of the IPv6 gateway.
	Ipv6GatewayId pulumi.StringPtrInput
	// The status of the resource.Valid values:`Normal`, `FinancialLocked` and `SecurityLocked`.
	Status pulumi.StringPtrInput
}

func (Ipv6InternetBandwidthState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv6InternetBandwidthState)(nil)).Elem()
}

type ipv6InternetBandwidthArgs struct {
	// The amount of Internet bandwidth resources of the IPv6 address, Unit: `Mbit/s`. Valid values: `1` to `5000`. **NOTE:** If `internetChargeType` is set to `PayByTraffic`, the amount of Internet bandwidth resources of the IPv6 address is limited by the specification of the IPv6 gateway. `Small` (default): specifies the Free edition and the Internet bandwidth is from `1` to `500` Mbit/s. `Medium`: specifies the Medium edition and the Internet bandwidth is from `1` to `1000` Mbit/s. `Large`: specifies the Large edition and the Internet bandwidth is from `1` to `2000` Mbit/s.
	Bandwidth int `pulumi:"bandwidth"`
	// The metering method of the Internet bandwidth resources of the IPv6 gateway. Valid values: `PayByBandwidth`, `PayByTraffic`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The ID of the IPv6 address.
	Ipv6AddressId string `pulumi:"ipv6AddressId"`
	// The ID of the IPv6 gateway.
	Ipv6GatewayId string `pulumi:"ipv6GatewayId"`
}

// The set of arguments for constructing a Ipv6InternetBandwidth resource.
type Ipv6InternetBandwidthArgs struct {
	// The amount of Internet bandwidth resources of the IPv6 address, Unit: `Mbit/s`. Valid values: `1` to `5000`. **NOTE:** If `internetChargeType` is set to `PayByTraffic`, the amount of Internet bandwidth resources of the IPv6 address is limited by the specification of the IPv6 gateway. `Small` (default): specifies the Free edition and the Internet bandwidth is from `1` to `500` Mbit/s. `Medium`: specifies the Medium edition and the Internet bandwidth is from `1` to `1000` Mbit/s. `Large`: specifies the Large edition and the Internet bandwidth is from `1` to `2000` Mbit/s.
	Bandwidth pulumi.IntInput
	// The metering method of the Internet bandwidth resources of the IPv6 gateway. Valid values: `PayByBandwidth`, `PayByTraffic`.
	InternetChargeType pulumi.StringPtrInput
	// The ID of the IPv6 address.
	Ipv6AddressId pulumi.StringInput
	// The ID of the IPv6 gateway.
	Ipv6GatewayId pulumi.StringInput
}

func (Ipv6InternetBandwidthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv6InternetBandwidthArgs)(nil)).Elem()
}

type Ipv6InternetBandwidthInput interface {
	pulumi.Input

	ToIpv6InternetBandwidthOutput() Ipv6InternetBandwidthOutput
	ToIpv6InternetBandwidthOutputWithContext(ctx context.Context) Ipv6InternetBandwidthOutput
}

func (*Ipv6InternetBandwidth) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6InternetBandwidth)(nil)).Elem()
}

func (i *Ipv6InternetBandwidth) ToIpv6InternetBandwidthOutput() Ipv6InternetBandwidthOutput {
	return i.ToIpv6InternetBandwidthOutputWithContext(context.Background())
}

func (i *Ipv6InternetBandwidth) ToIpv6InternetBandwidthOutputWithContext(ctx context.Context) Ipv6InternetBandwidthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6InternetBandwidthOutput)
}

// Ipv6InternetBandwidthArrayInput is an input type that accepts Ipv6InternetBandwidthArray and Ipv6InternetBandwidthArrayOutput values.
// You can construct a concrete instance of `Ipv6InternetBandwidthArrayInput` via:
//
//	Ipv6InternetBandwidthArray{ Ipv6InternetBandwidthArgs{...} }
type Ipv6InternetBandwidthArrayInput interface {
	pulumi.Input

	ToIpv6InternetBandwidthArrayOutput() Ipv6InternetBandwidthArrayOutput
	ToIpv6InternetBandwidthArrayOutputWithContext(context.Context) Ipv6InternetBandwidthArrayOutput
}

type Ipv6InternetBandwidthArray []Ipv6InternetBandwidthInput

func (Ipv6InternetBandwidthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv6InternetBandwidth)(nil)).Elem()
}

func (i Ipv6InternetBandwidthArray) ToIpv6InternetBandwidthArrayOutput() Ipv6InternetBandwidthArrayOutput {
	return i.ToIpv6InternetBandwidthArrayOutputWithContext(context.Background())
}

func (i Ipv6InternetBandwidthArray) ToIpv6InternetBandwidthArrayOutputWithContext(ctx context.Context) Ipv6InternetBandwidthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6InternetBandwidthArrayOutput)
}

// Ipv6InternetBandwidthMapInput is an input type that accepts Ipv6InternetBandwidthMap and Ipv6InternetBandwidthMapOutput values.
// You can construct a concrete instance of `Ipv6InternetBandwidthMapInput` via:
//
//	Ipv6InternetBandwidthMap{ "key": Ipv6InternetBandwidthArgs{...} }
type Ipv6InternetBandwidthMapInput interface {
	pulumi.Input

	ToIpv6InternetBandwidthMapOutput() Ipv6InternetBandwidthMapOutput
	ToIpv6InternetBandwidthMapOutputWithContext(context.Context) Ipv6InternetBandwidthMapOutput
}

type Ipv6InternetBandwidthMap map[string]Ipv6InternetBandwidthInput

func (Ipv6InternetBandwidthMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv6InternetBandwidth)(nil)).Elem()
}

func (i Ipv6InternetBandwidthMap) ToIpv6InternetBandwidthMapOutput() Ipv6InternetBandwidthMapOutput {
	return i.ToIpv6InternetBandwidthMapOutputWithContext(context.Background())
}

func (i Ipv6InternetBandwidthMap) ToIpv6InternetBandwidthMapOutputWithContext(ctx context.Context) Ipv6InternetBandwidthMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6InternetBandwidthMapOutput)
}

type Ipv6InternetBandwidthOutput struct{ *pulumi.OutputState }

func (Ipv6InternetBandwidthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6InternetBandwidth)(nil)).Elem()
}

func (o Ipv6InternetBandwidthOutput) ToIpv6InternetBandwidthOutput() Ipv6InternetBandwidthOutput {
	return o
}

func (o Ipv6InternetBandwidthOutput) ToIpv6InternetBandwidthOutputWithContext(ctx context.Context) Ipv6InternetBandwidthOutput {
	return o
}

// The amount of Internet bandwidth resources of the IPv6 address, Unit: `Mbit/s`. Valid values: `1` to `5000`. **NOTE:** If `internetChargeType` is set to `PayByTraffic`, the amount of Internet bandwidth resources of the IPv6 address is limited by the specification of the IPv6 gateway. `Small` (default): specifies the Free edition and the Internet bandwidth is from `1` to `500` Mbit/s. `Medium`: specifies the Medium edition and the Internet bandwidth is from `1` to `1000` Mbit/s. `Large`: specifies the Large edition and the Internet bandwidth is from `1` to `2000` Mbit/s.
func (o Ipv6InternetBandwidthOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *Ipv6InternetBandwidth) pulumi.IntOutput { return v.Bandwidth }).(pulumi.IntOutput)
}

// The metering method of the Internet bandwidth resources of the IPv6 gateway. Valid values: `PayByBandwidth`, `PayByTraffic`.
func (o Ipv6InternetBandwidthOutput) InternetChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6InternetBandwidth) pulumi.StringOutput { return v.InternetChargeType }).(pulumi.StringOutput)
}

// The ID of the IPv6 address.
func (o Ipv6InternetBandwidthOutput) Ipv6AddressId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6InternetBandwidth) pulumi.StringOutput { return v.Ipv6AddressId }).(pulumi.StringOutput)
}

// The ID of the IPv6 gateway.
func (o Ipv6InternetBandwidthOutput) Ipv6GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6InternetBandwidth) pulumi.StringOutput { return v.Ipv6GatewayId }).(pulumi.StringOutput)
}

// The status of the resource.Valid values:`Normal`, `FinancialLocked` and `SecurityLocked`.
func (o Ipv6InternetBandwidthOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6InternetBandwidth) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type Ipv6InternetBandwidthArrayOutput struct{ *pulumi.OutputState }

func (Ipv6InternetBandwidthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv6InternetBandwidth)(nil)).Elem()
}

func (o Ipv6InternetBandwidthArrayOutput) ToIpv6InternetBandwidthArrayOutput() Ipv6InternetBandwidthArrayOutput {
	return o
}

func (o Ipv6InternetBandwidthArrayOutput) ToIpv6InternetBandwidthArrayOutputWithContext(ctx context.Context) Ipv6InternetBandwidthArrayOutput {
	return o
}

func (o Ipv6InternetBandwidthArrayOutput) Index(i pulumi.IntInput) Ipv6InternetBandwidthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ipv6InternetBandwidth {
		return vs[0].([]*Ipv6InternetBandwidth)[vs[1].(int)]
	}).(Ipv6InternetBandwidthOutput)
}

type Ipv6InternetBandwidthMapOutput struct{ *pulumi.OutputState }

func (Ipv6InternetBandwidthMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv6InternetBandwidth)(nil)).Elem()
}

func (o Ipv6InternetBandwidthMapOutput) ToIpv6InternetBandwidthMapOutput() Ipv6InternetBandwidthMapOutput {
	return o
}

func (o Ipv6InternetBandwidthMapOutput) ToIpv6InternetBandwidthMapOutputWithContext(ctx context.Context) Ipv6InternetBandwidthMapOutput {
	return o
}

func (o Ipv6InternetBandwidthMapOutput) MapIndex(k pulumi.StringInput) Ipv6InternetBandwidthOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ipv6InternetBandwidth {
		return vs[0].(map[string]*Ipv6InternetBandwidth)[vs[1].(string)]
	}).(Ipv6InternetBandwidthOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6InternetBandwidthInput)(nil)).Elem(), &Ipv6InternetBandwidth{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6InternetBandwidthArrayInput)(nil)).Elem(), Ipv6InternetBandwidthArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6InternetBandwidthMapInput)(nil)).Elem(), Ipv6InternetBandwidthMap{})
	pulumi.RegisterOutputType(Ipv6InternetBandwidthOutput{})
	pulumi.RegisterOutputType(Ipv6InternetBandwidthArrayOutput{})
	pulumi.RegisterOutputType(Ipv6InternetBandwidthMapOutput{})
}
