// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPN Gateway Vco Route resource.
//
// For information about VPN Gateway Vco Route and how to use it, see [What is Vco Route](https://www.alibabacloud.com/help/zh/virtual-private-cloud/latest/createvcorouteentry).
//
// > **NOTE:** Available in v1.183.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultInstance, err := cen.NewInstance(ctx, "defaultInstance", &cen.InstanceArgs{
//				CenInstanceName: pulumi.Any(_var.Name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultTransitRouter, err := cen.NewTransitRouter(ctx, "defaultTransitRouter", &cen.TransitRouterArgs{
//				CenId:                    defaultInstance.ID(),
//				TransitRouterDescription: pulumi.String("desd"),
//				TransitRouterName:        pulumi.Any(_var.Name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultTransitRouterAvailableResources, err := cen.GetTransitRouterAvailableResources(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultCustomerGateway, err := vpn.NewCustomerGateway(ctx, "defaultCustomerGateway", &vpn.CustomerGatewayArgs{
//				IpAddress:   pulumi.String("42.104.22.210"),
//				Asn:         pulumi.String("45014"),
//				Description: pulumi.String("testAccVpnConnectionDesc"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultGatewayVpnAttachment, err := vpn.NewGatewayVpnAttachment(ctx, "defaultGatewayVpnAttachment", &vpn.GatewayVpnAttachmentArgs{
//				CustomerGatewayId: defaultCustomerGateway.ID(),
//				NetworkType:       pulumi.String("public"),
//				LocalSubnet:       pulumi.String("0.0.0.0/0"),
//				RemoteSubnet:      pulumi.String("0.0.0.0/0"),
//				EffectImmediately: pulumi.Bool(false),
//				IkeConfig: &vpn.GatewayVpnAttachmentIkeConfigArgs{
//					IkeAuthAlg:  pulumi.String("md5"),
//					IkeEncAlg:   pulumi.String("des"),
//					IkeVersion:  pulumi.String("ikev2"),
//					IkeMode:     pulumi.String("main"),
//					IkeLifetime: pulumi.Int(86400),
//					Psk:         pulumi.String("tf-testvpn2"),
//					IkePfs:      pulumi.String("group1"),
//					RemoteId:    pulumi.String("testbob2"),
//					LocalId:     pulumi.String("testalice2"),
//				},
//				IpsecConfig: &vpn.GatewayVpnAttachmentIpsecConfigArgs{
//					IpsecPfs:      pulumi.String("group5"),
//					IpsecEncAlg:   pulumi.String("des"),
//					IpsecAuthAlg:  pulumi.String("md5"),
//					IpsecLifetime: pulumi.Int(86400),
//				},
//				BgpConfig: &vpn.GatewayVpnAttachmentBgpConfigArgs{
//					Enable:     pulumi.Bool(true),
//					LocalAsn:   pulumi.Int(45014),
//					TunnelCidr: pulumi.String("169.254.11.0/30"),
//					LocalBgpIp: pulumi.String("169.254.11.1"),
//				},
//				HealthCheckConfig: &vpn.GatewayVpnAttachmentHealthCheckConfigArgs{
//					Enable:   pulumi.Bool(true),
//					Sip:      pulumi.String("192.168.1.1"),
//					Dip:      pulumi.String("10.0.0.1"),
//					Interval: pulumi.Int(10),
//					Retry:    pulumi.Int(10),
//					Policy:   pulumi.String("revoke_route"),
//				},
//				EnableDpd:          pulumi.Bool(true),
//				EnableNatTraversal: pulumi.Bool(true),
//				VpnAttachmentName:  pulumi.Any(_var.Name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultTransitRouterVpnAttachment, err := cen.NewTransitRouterVpnAttachment(ctx, "defaultTransitRouterVpnAttachment", &cen.TransitRouterVpnAttachmentArgs{
//				AutoPublishRouteEnabled:            pulumi.Bool(false),
//				TransitRouterAttachmentDescription: pulumi.Any(_var.Name),
//				TransitRouterAttachmentName:        pulumi.Any(_var.Name),
//				CenId:                              defaultTransitRouter.CenId,
//				TransitRouterId:                    defaultTransitRouter.TransitRouterId,
//				VpnId:                              defaultGatewayVpnAttachment.ID(),
//				Zones: cen.TransitRouterVpnAttachmentZoneArray{
//					&cen.TransitRouterVpnAttachmentZoneArgs{
//						ZoneId: *pulumi.String(defaultTransitRouterAvailableResources.Resources[0].MasterZones[0]),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpn.NewGatewayVcoRoute(ctx, "defaultGatewayVcoRoute", &vpn.GatewayVcoRouteArgs{
//				RouteDest:       pulumi.String("192.168.12.0/24"),
//				NextHop:         defaultTransitRouterVpnAttachment.VpnId,
//				VpnConnectionId: defaultTransitRouterVpnAttachment.VpnId,
//				Weight:          pulumi.Int(100),
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
// VPN Gateway Vco Route can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:vpn/gatewayVcoRoute:GatewayVcoRoute example <vpn_connection_id>:<route_dest>:<next_hop>:<weight>
//
// ```
type GatewayVcoRoute struct {
	pulumi.CustomResourceState

	// The next hop of the destination route.
	NextHop pulumi.StringOutput `pulumi:"nextHop"`
	// The destination network segment of the destination route.
	RouteDest pulumi.StringOutput `pulumi:"routeDest"`
	// The status of the vpn route entry.
	Status pulumi.StringOutput `pulumi:"status"`
	// The id of the vpn attachment.
	VpnConnectionId pulumi.StringOutput `pulumi:"vpnConnectionId"`
	// The weight value of the destination route. Valid values: `0`, `100`.
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewGatewayVcoRoute registers a new resource with the given unique name, arguments, and options.
func NewGatewayVcoRoute(ctx *pulumi.Context,
	name string, args *GatewayVcoRouteArgs, opts ...pulumi.ResourceOption) (*GatewayVcoRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NextHop == nil {
		return nil, errors.New("invalid value for required argument 'NextHop'")
	}
	if args.RouteDest == nil {
		return nil, errors.New("invalid value for required argument 'RouteDest'")
	}
	if args.VpnConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'VpnConnectionId'")
	}
	if args.Weight == nil {
		return nil, errors.New("invalid value for required argument 'Weight'")
	}
	var resource GatewayVcoRoute
	err := ctx.RegisterResource("alicloud:vpn/gatewayVcoRoute:GatewayVcoRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayVcoRoute gets an existing GatewayVcoRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayVcoRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayVcoRouteState, opts ...pulumi.ResourceOption) (*GatewayVcoRoute, error) {
	var resource GatewayVcoRoute
	err := ctx.ReadResource("alicloud:vpn/gatewayVcoRoute:GatewayVcoRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayVcoRoute resources.
type gatewayVcoRouteState struct {
	// The next hop of the destination route.
	NextHop *string `pulumi:"nextHop"`
	// The destination network segment of the destination route.
	RouteDest *string `pulumi:"routeDest"`
	// The status of the vpn route entry.
	Status *string `pulumi:"status"`
	// The id of the vpn attachment.
	VpnConnectionId *string `pulumi:"vpnConnectionId"`
	// The weight value of the destination route. Valid values: `0`, `100`.
	Weight *int `pulumi:"weight"`
}

type GatewayVcoRouteState struct {
	// The next hop of the destination route.
	NextHop pulumi.StringPtrInput
	// The destination network segment of the destination route.
	RouteDest pulumi.StringPtrInput
	// The status of the vpn route entry.
	Status pulumi.StringPtrInput
	// The id of the vpn attachment.
	VpnConnectionId pulumi.StringPtrInput
	// The weight value of the destination route. Valid values: `0`, `100`.
	Weight pulumi.IntPtrInput
}

func (GatewayVcoRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayVcoRouteState)(nil)).Elem()
}

type gatewayVcoRouteArgs struct {
	// The next hop of the destination route.
	NextHop string `pulumi:"nextHop"`
	// The destination network segment of the destination route.
	RouteDest string `pulumi:"routeDest"`
	// The id of the vpn attachment.
	VpnConnectionId string `pulumi:"vpnConnectionId"`
	// The weight value of the destination route. Valid values: `0`, `100`.
	Weight int `pulumi:"weight"`
}

// The set of arguments for constructing a GatewayVcoRoute resource.
type GatewayVcoRouteArgs struct {
	// The next hop of the destination route.
	NextHop pulumi.StringInput
	// The destination network segment of the destination route.
	RouteDest pulumi.StringInput
	// The id of the vpn attachment.
	VpnConnectionId pulumi.StringInput
	// The weight value of the destination route. Valid values: `0`, `100`.
	Weight pulumi.IntInput
}

func (GatewayVcoRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayVcoRouteArgs)(nil)).Elem()
}

type GatewayVcoRouteInput interface {
	pulumi.Input

	ToGatewayVcoRouteOutput() GatewayVcoRouteOutput
	ToGatewayVcoRouteOutputWithContext(ctx context.Context) GatewayVcoRouteOutput
}

func (*GatewayVcoRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayVcoRoute)(nil)).Elem()
}

func (i *GatewayVcoRoute) ToGatewayVcoRouteOutput() GatewayVcoRouteOutput {
	return i.ToGatewayVcoRouteOutputWithContext(context.Background())
}

func (i *GatewayVcoRoute) ToGatewayVcoRouteOutputWithContext(ctx context.Context) GatewayVcoRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayVcoRouteOutput)
}

// GatewayVcoRouteArrayInput is an input type that accepts GatewayVcoRouteArray and GatewayVcoRouteArrayOutput values.
// You can construct a concrete instance of `GatewayVcoRouteArrayInput` via:
//
//	GatewayVcoRouteArray{ GatewayVcoRouteArgs{...} }
type GatewayVcoRouteArrayInput interface {
	pulumi.Input

	ToGatewayVcoRouteArrayOutput() GatewayVcoRouteArrayOutput
	ToGatewayVcoRouteArrayOutputWithContext(context.Context) GatewayVcoRouteArrayOutput
}

type GatewayVcoRouteArray []GatewayVcoRouteInput

func (GatewayVcoRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayVcoRoute)(nil)).Elem()
}

func (i GatewayVcoRouteArray) ToGatewayVcoRouteArrayOutput() GatewayVcoRouteArrayOutput {
	return i.ToGatewayVcoRouteArrayOutputWithContext(context.Background())
}

func (i GatewayVcoRouteArray) ToGatewayVcoRouteArrayOutputWithContext(ctx context.Context) GatewayVcoRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayVcoRouteArrayOutput)
}

// GatewayVcoRouteMapInput is an input type that accepts GatewayVcoRouteMap and GatewayVcoRouteMapOutput values.
// You can construct a concrete instance of `GatewayVcoRouteMapInput` via:
//
//	GatewayVcoRouteMap{ "key": GatewayVcoRouteArgs{...} }
type GatewayVcoRouteMapInput interface {
	pulumi.Input

	ToGatewayVcoRouteMapOutput() GatewayVcoRouteMapOutput
	ToGatewayVcoRouteMapOutputWithContext(context.Context) GatewayVcoRouteMapOutput
}

type GatewayVcoRouteMap map[string]GatewayVcoRouteInput

func (GatewayVcoRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayVcoRoute)(nil)).Elem()
}

func (i GatewayVcoRouteMap) ToGatewayVcoRouteMapOutput() GatewayVcoRouteMapOutput {
	return i.ToGatewayVcoRouteMapOutputWithContext(context.Background())
}

func (i GatewayVcoRouteMap) ToGatewayVcoRouteMapOutputWithContext(ctx context.Context) GatewayVcoRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayVcoRouteMapOutput)
}

type GatewayVcoRouteOutput struct{ *pulumi.OutputState }

func (GatewayVcoRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayVcoRoute)(nil)).Elem()
}

func (o GatewayVcoRouteOutput) ToGatewayVcoRouteOutput() GatewayVcoRouteOutput {
	return o
}

func (o GatewayVcoRouteOutput) ToGatewayVcoRouteOutputWithContext(ctx context.Context) GatewayVcoRouteOutput {
	return o
}

// The next hop of the destination route.
func (o GatewayVcoRouteOutput) NextHop() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayVcoRoute) pulumi.StringOutput { return v.NextHop }).(pulumi.StringOutput)
}

// The destination network segment of the destination route.
func (o GatewayVcoRouteOutput) RouteDest() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayVcoRoute) pulumi.StringOutput { return v.RouteDest }).(pulumi.StringOutput)
}

// The status of the vpn route entry.
func (o GatewayVcoRouteOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayVcoRoute) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The id of the vpn attachment.
func (o GatewayVcoRouteOutput) VpnConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayVcoRoute) pulumi.StringOutput { return v.VpnConnectionId }).(pulumi.StringOutput)
}

// The weight value of the destination route. Valid values: `0`, `100`.
func (o GatewayVcoRouteOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayVcoRoute) pulumi.IntOutput { return v.Weight }).(pulumi.IntOutput)
}

type GatewayVcoRouteArrayOutput struct{ *pulumi.OutputState }

func (GatewayVcoRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayVcoRoute)(nil)).Elem()
}

func (o GatewayVcoRouteArrayOutput) ToGatewayVcoRouteArrayOutput() GatewayVcoRouteArrayOutput {
	return o
}

func (o GatewayVcoRouteArrayOutput) ToGatewayVcoRouteArrayOutputWithContext(ctx context.Context) GatewayVcoRouteArrayOutput {
	return o
}

func (o GatewayVcoRouteArrayOutput) Index(i pulumi.IntInput) GatewayVcoRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayVcoRoute {
		return vs[0].([]*GatewayVcoRoute)[vs[1].(int)]
	}).(GatewayVcoRouteOutput)
}

type GatewayVcoRouteMapOutput struct{ *pulumi.OutputState }

func (GatewayVcoRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayVcoRoute)(nil)).Elem()
}

func (o GatewayVcoRouteMapOutput) ToGatewayVcoRouteMapOutput() GatewayVcoRouteMapOutput {
	return o
}

func (o GatewayVcoRouteMapOutput) ToGatewayVcoRouteMapOutputWithContext(ctx context.Context) GatewayVcoRouteMapOutput {
	return o
}

func (o GatewayVcoRouteMapOutput) MapIndex(k pulumi.StringInput) GatewayVcoRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayVcoRoute {
		return vs[0].(map[string]*GatewayVcoRoute)[vs[1].(string)]
	}).(GatewayVcoRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayVcoRouteInput)(nil)).Elem(), &GatewayVcoRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayVcoRouteArrayInput)(nil)).Elem(), GatewayVcoRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayVcoRouteMapInput)(nil)).Elem(), GatewayVcoRouteMap{})
	pulumi.RegisterOutputType(GatewayVcoRouteOutput{})
	pulumi.RegisterOutputType(GatewayVcoRouteArrayOutput{})
	pulumi.RegisterOutputType(GatewayVcoRouteMapOutput{})
}
