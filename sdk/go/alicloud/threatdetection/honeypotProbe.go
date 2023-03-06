// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package threatdetection

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Threat Detection Honeypot Probe resource.
//
// For information about Threat Detection Honeypot Probe and how to use it, see [What is Honeypot Probe](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-createhoneypotprobe).
//
// > **NOTE:** Available in v1.195.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/threatdetection"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := threatdetection.NewHoneypotProbe(ctx, "default", &threatdetection.HoneypotProbeArgs{
//				Arp:           pulumi.Bool(true),
//				ControlNodeId: pulumi.String("a44e1ab3-6945-444c-889d-5bacee7056e8"),
//				DisplayName:   pulumi.String("apispec"),
//				HoneypotBindLists: threatdetection.HoneypotProbeHoneypotBindListArray{
//					&threatdetection.HoneypotProbeHoneypotBindListArgs{
//						BindPortLists: threatdetection.HoneypotProbeHoneypotBindListBindPortListArray{
//							&threatdetection.HoneypotProbeHoneypotBindListBindPortListArgs{
//								EndPort:   pulumi.Int(80),
//								StartPort: pulumi.Int(80),
//							},
//						},
//						HoneypotId: pulumi.String("ede59ccdb1b7a2e21735d4593a6eb5ed31883af320c5ab63ab33818e94307be9"),
//					},
//				},
//				Ping:      pulumi.Bool(true),
//				ProbeType: pulumi.String("host_probe"),
//				Uuid:      pulumi.String("032b618f-b220-4a0d-bd37-fbdc6ef58b6a"),
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
// Threat Detection Honeypot Probe can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:threatdetection/honeypotProbe:HoneypotProbe example <id>
//
// ```
type HoneypotProbe struct {
	pulumi.CustomResourceState

	// ARP spoofing detection.**true**: Enable **false**: Disabled
	Arp pulumi.BoolPtrOutput `pulumi:"arp"`
	// The ID of the management node.
	ControlNodeId pulumi.StringOutput `pulumi:"controlNodeId"`
	// Probe display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Configure the service.See the following `Block HoneypotBindList`.
	HoneypotBindLists HoneypotProbeHoneypotBindListArrayOutput `pulumi:"honeypotBindLists"`
	// The first ID of the resource
	HoneypotProbeId pulumi.StringOutput `pulumi:"honeypotProbeId"`
	// Ping scan detection. Value: **true**: Enable **false**: Disabled
	Ping pulumi.BoolPtrOutput `pulumi:"ping"`
	// Probe type, support `hostProbe` and `vpcBlackHoleProbe`.
	ProbeType pulumi.StringOutput `pulumi:"probeType"`
	// The version of the probe.
	ProbeVersion pulumi.StringOutput `pulumi:"probeVersion"`
	// The IP address of the proxy.
	ProxyIp pulumi.StringPtrOutput `pulumi:"proxyIp"`
	// Listen to the IP address list.
	ServiceIpLists pulumi.StringArrayOutput `pulumi:"serviceIpLists"`
	// The status of the resource
	Status pulumi.StringOutput `pulumi:"status"`
	// Machine uuid, **probe_type** is `hostProbe`. This value cannot be empty.
	Uuid pulumi.StringPtrOutput `pulumi:"uuid"`
	// The ID of the VPC. **probe_type** is `vpcBlackHoleProbe`. This value cannot be empty.
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
}

// NewHoneypotProbe registers a new resource with the given unique name, arguments, and options.
func NewHoneypotProbe(ctx *pulumi.Context,
	name string, args *HoneypotProbeArgs, opts ...pulumi.ResourceOption) (*HoneypotProbe, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlNodeId == nil {
		return nil, errors.New("invalid value for required argument 'ControlNodeId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ProbeType == nil {
		return nil, errors.New("invalid value for required argument 'ProbeType'")
	}
	var resource HoneypotProbe
	err := ctx.RegisterResource("alicloud:threatdetection/honeypotProbe:HoneypotProbe", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHoneypotProbe gets an existing HoneypotProbe resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHoneypotProbe(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HoneypotProbeState, opts ...pulumi.ResourceOption) (*HoneypotProbe, error) {
	var resource HoneypotProbe
	err := ctx.ReadResource("alicloud:threatdetection/honeypotProbe:HoneypotProbe", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HoneypotProbe resources.
type honeypotProbeState struct {
	// ARP spoofing detection.**true**: Enable **false**: Disabled
	Arp *bool `pulumi:"arp"`
	// The ID of the management node.
	ControlNodeId *string `pulumi:"controlNodeId"`
	// Probe display name.
	DisplayName *string `pulumi:"displayName"`
	// Configure the service.See the following `Block HoneypotBindList`.
	HoneypotBindLists []HoneypotProbeHoneypotBindList `pulumi:"honeypotBindLists"`
	// The first ID of the resource
	HoneypotProbeId *string `pulumi:"honeypotProbeId"`
	// Ping scan detection. Value: **true**: Enable **false**: Disabled
	Ping *bool `pulumi:"ping"`
	// Probe type, support `hostProbe` and `vpcBlackHoleProbe`.
	ProbeType *string `pulumi:"probeType"`
	// The version of the probe.
	ProbeVersion *string `pulumi:"probeVersion"`
	// The IP address of the proxy.
	ProxyIp *string `pulumi:"proxyIp"`
	// Listen to the IP address list.
	ServiceIpLists []string `pulumi:"serviceIpLists"`
	// The status of the resource
	Status *string `pulumi:"status"`
	// Machine uuid, **probe_type** is `hostProbe`. This value cannot be empty.
	Uuid *string `pulumi:"uuid"`
	// The ID of the VPC. **probe_type** is `vpcBlackHoleProbe`. This value cannot be empty.
	VpcId *string `pulumi:"vpcId"`
}

type HoneypotProbeState struct {
	// ARP spoofing detection.**true**: Enable **false**: Disabled
	Arp pulumi.BoolPtrInput
	// The ID of the management node.
	ControlNodeId pulumi.StringPtrInput
	// Probe display name.
	DisplayName pulumi.StringPtrInput
	// Configure the service.See the following `Block HoneypotBindList`.
	HoneypotBindLists HoneypotProbeHoneypotBindListArrayInput
	// The first ID of the resource
	HoneypotProbeId pulumi.StringPtrInput
	// Ping scan detection. Value: **true**: Enable **false**: Disabled
	Ping pulumi.BoolPtrInput
	// Probe type, support `hostProbe` and `vpcBlackHoleProbe`.
	ProbeType pulumi.StringPtrInput
	// The version of the probe.
	ProbeVersion pulumi.StringPtrInput
	// The IP address of the proxy.
	ProxyIp pulumi.StringPtrInput
	// Listen to the IP address list.
	ServiceIpLists pulumi.StringArrayInput
	// The status of the resource
	Status pulumi.StringPtrInput
	// Machine uuid, **probe_type** is `hostProbe`. This value cannot be empty.
	Uuid pulumi.StringPtrInput
	// The ID of the VPC. **probe_type** is `vpcBlackHoleProbe`. This value cannot be empty.
	VpcId pulumi.StringPtrInput
}

func (HoneypotProbeState) ElementType() reflect.Type {
	return reflect.TypeOf((*honeypotProbeState)(nil)).Elem()
}

type honeypotProbeArgs struct {
	// ARP spoofing detection.**true**: Enable **false**: Disabled
	Arp *bool `pulumi:"arp"`
	// The ID of the management node.
	ControlNodeId string `pulumi:"controlNodeId"`
	// Probe display name.
	DisplayName string `pulumi:"displayName"`
	// Configure the service.See the following `Block HoneypotBindList`.
	HoneypotBindLists []HoneypotProbeHoneypotBindList `pulumi:"honeypotBindLists"`
	// Ping scan detection. Value: **true**: Enable **false**: Disabled
	Ping *bool `pulumi:"ping"`
	// Probe type, support `hostProbe` and `vpcBlackHoleProbe`.
	ProbeType string `pulumi:"probeType"`
	// The version of the probe.
	ProbeVersion *string `pulumi:"probeVersion"`
	// The IP address of the proxy.
	ProxyIp *string `pulumi:"proxyIp"`
	// Listen to the IP address list.
	ServiceIpLists []string `pulumi:"serviceIpLists"`
	// Machine uuid, **probe_type** is `hostProbe`. This value cannot be empty.
	Uuid *string `pulumi:"uuid"`
	// The ID of the VPC. **probe_type** is `vpcBlackHoleProbe`. This value cannot be empty.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a HoneypotProbe resource.
type HoneypotProbeArgs struct {
	// ARP spoofing detection.**true**: Enable **false**: Disabled
	Arp pulumi.BoolPtrInput
	// The ID of the management node.
	ControlNodeId pulumi.StringInput
	// Probe display name.
	DisplayName pulumi.StringInput
	// Configure the service.See the following `Block HoneypotBindList`.
	HoneypotBindLists HoneypotProbeHoneypotBindListArrayInput
	// Ping scan detection. Value: **true**: Enable **false**: Disabled
	Ping pulumi.BoolPtrInput
	// Probe type, support `hostProbe` and `vpcBlackHoleProbe`.
	ProbeType pulumi.StringInput
	// The version of the probe.
	ProbeVersion pulumi.StringPtrInput
	// The IP address of the proxy.
	ProxyIp pulumi.StringPtrInput
	// Listen to the IP address list.
	ServiceIpLists pulumi.StringArrayInput
	// Machine uuid, **probe_type** is `hostProbe`. This value cannot be empty.
	Uuid pulumi.StringPtrInput
	// The ID of the VPC. **probe_type** is `vpcBlackHoleProbe`. This value cannot be empty.
	VpcId pulumi.StringPtrInput
}

func (HoneypotProbeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*honeypotProbeArgs)(nil)).Elem()
}

type HoneypotProbeInput interface {
	pulumi.Input

	ToHoneypotProbeOutput() HoneypotProbeOutput
	ToHoneypotProbeOutputWithContext(ctx context.Context) HoneypotProbeOutput
}

func (*HoneypotProbe) ElementType() reflect.Type {
	return reflect.TypeOf((**HoneypotProbe)(nil)).Elem()
}

func (i *HoneypotProbe) ToHoneypotProbeOutput() HoneypotProbeOutput {
	return i.ToHoneypotProbeOutputWithContext(context.Background())
}

func (i *HoneypotProbe) ToHoneypotProbeOutputWithContext(ctx context.Context) HoneypotProbeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HoneypotProbeOutput)
}

// HoneypotProbeArrayInput is an input type that accepts HoneypotProbeArray and HoneypotProbeArrayOutput values.
// You can construct a concrete instance of `HoneypotProbeArrayInput` via:
//
//	HoneypotProbeArray{ HoneypotProbeArgs{...} }
type HoneypotProbeArrayInput interface {
	pulumi.Input

	ToHoneypotProbeArrayOutput() HoneypotProbeArrayOutput
	ToHoneypotProbeArrayOutputWithContext(context.Context) HoneypotProbeArrayOutput
}

type HoneypotProbeArray []HoneypotProbeInput

func (HoneypotProbeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HoneypotProbe)(nil)).Elem()
}

func (i HoneypotProbeArray) ToHoneypotProbeArrayOutput() HoneypotProbeArrayOutput {
	return i.ToHoneypotProbeArrayOutputWithContext(context.Background())
}

func (i HoneypotProbeArray) ToHoneypotProbeArrayOutputWithContext(ctx context.Context) HoneypotProbeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HoneypotProbeArrayOutput)
}

// HoneypotProbeMapInput is an input type that accepts HoneypotProbeMap and HoneypotProbeMapOutput values.
// You can construct a concrete instance of `HoneypotProbeMapInput` via:
//
//	HoneypotProbeMap{ "key": HoneypotProbeArgs{...} }
type HoneypotProbeMapInput interface {
	pulumi.Input

	ToHoneypotProbeMapOutput() HoneypotProbeMapOutput
	ToHoneypotProbeMapOutputWithContext(context.Context) HoneypotProbeMapOutput
}

type HoneypotProbeMap map[string]HoneypotProbeInput

func (HoneypotProbeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HoneypotProbe)(nil)).Elem()
}

func (i HoneypotProbeMap) ToHoneypotProbeMapOutput() HoneypotProbeMapOutput {
	return i.ToHoneypotProbeMapOutputWithContext(context.Background())
}

func (i HoneypotProbeMap) ToHoneypotProbeMapOutputWithContext(ctx context.Context) HoneypotProbeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HoneypotProbeMapOutput)
}

type HoneypotProbeOutput struct{ *pulumi.OutputState }

func (HoneypotProbeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HoneypotProbe)(nil)).Elem()
}

func (o HoneypotProbeOutput) ToHoneypotProbeOutput() HoneypotProbeOutput {
	return o
}

func (o HoneypotProbeOutput) ToHoneypotProbeOutputWithContext(ctx context.Context) HoneypotProbeOutput {
	return o
}

// ARP spoofing detection.**true**: Enable **false**: Disabled
func (o HoneypotProbeOutput) Arp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HoneypotProbe) pulumi.BoolPtrOutput { return v.Arp }).(pulumi.BoolPtrOutput)
}

// The ID of the management node.
func (o HoneypotProbeOutput) ControlNodeId() pulumi.StringOutput {
	return o.ApplyT(func(v *HoneypotProbe) pulumi.StringOutput { return v.ControlNodeId }).(pulumi.StringOutput)
}

// Probe display name.
func (o HoneypotProbeOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *HoneypotProbe) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Configure the service.See the following `Block HoneypotBindList`.
func (o HoneypotProbeOutput) HoneypotBindLists() HoneypotProbeHoneypotBindListArrayOutput {
	return o.ApplyT(func(v *HoneypotProbe) HoneypotProbeHoneypotBindListArrayOutput { return v.HoneypotBindLists }).(HoneypotProbeHoneypotBindListArrayOutput)
}

// The first ID of the resource
func (o HoneypotProbeOutput) HoneypotProbeId() pulumi.StringOutput {
	return o.ApplyT(func(v *HoneypotProbe) pulumi.StringOutput { return v.HoneypotProbeId }).(pulumi.StringOutput)
}

// Ping scan detection. Value: **true**: Enable **false**: Disabled
func (o HoneypotProbeOutput) Ping() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HoneypotProbe) pulumi.BoolPtrOutput { return v.Ping }).(pulumi.BoolPtrOutput)
}

// Probe type, support `hostProbe` and `vpcBlackHoleProbe`.
func (o HoneypotProbeOutput) ProbeType() pulumi.StringOutput {
	return o.ApplyT(func(v *HoneypotProbe) pulumi.StringOutput { return v.ProbeType }).(pulumi.StringOutput)
}

// The version of the probe.
func (o HoneypotProbeOutput) ProbeVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *HoneypotProbe) pulumi.StringOutput { return v.ProbeVersion }).(pulumi.StringOutput)
}

// The IP address of the proxy.
func (o HoneypotProbeOutput) ProxyIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HoneypotProbe) pulumi.StringPtrOutput { return v.ProxyIp }).(pulumi.StringPtrOutput)
}

// Listen to the IP address list.
func (o HoneypotProbeOutput) ServiceIpLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HoneypotProbe) pulumi.StringArrayOutput { return v.ServiceIpLists }).(pulumi.StringArrayOutput)
}

// The status of the resource
func (o HoneypotProbeOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *HoneypotProbe) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Machine uuid, **probe_type** is `hostProbe`. This value cannot be empty.
func (o HoneypotProbeOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HoneypotProbe) pulumi.StringPtrOutput { return v.Uuid }).(pulumi.StringPtrOutput)
}

// The ID of the VPC. **probe_type** is `vpcBlackHoleProbe`. This value cannot be empty.
func (o HoneypotProbeOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HoneypotProbe) pulumi.StringPtrOutput { return v.VpcId }).(pulumi.StringPtrOutput)
}

type HoneypotProbeArrayOutput struct{ *pulumi.OutputState }

func (HoneypotProbeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HoneypotProbe)(nil)).Elem()
}

func (o HoneypotProbeArrayOutput) ToHoneypotProbeArrayOutput() HoneypotProbeArrayOutput {
	return o
}

func (o HoneypotProbeArrayOutput) ToHoneypotProbeArrayOutputWithContext(ctx context.Context) HoneypotProbeArrayOutput {
	return o
}

func (o HoneypotProbeArrayOutput) Index(i pulumi.IntInput) HoneypotProbeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HoneypotProbe {
		return vs[0].([]*HoneypotProbe)[vs[1].(int)]
	}).(HoneypotProbeOutput)
}

type HoneypotProbeMapOutput struct{ *pulumi.OutputState }

func (HoneypotProbeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HoneypotProbe)(nil)).Elem()
}

func (o HoneypotProbeMapOutput) ToHoneypotProbeMapOutput() HoneypotProbeMapOutput {
	return o
}

func (o HoneypotProbeMapOutput) ToHoneypotProbeMapOutputWithContext(ctx context.Context) HoneypotProbeMapOutput {
	return o
}

func (o HoneypotProbeMapOutput) MapIndex(k pulumi.StringInput) HoneypotProbeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HoneypotProbe {
		return vs[0].(map[string]*HoneypotProbe)[vs[1].(string)]
	}).(HoneypotProbeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HoneypotProbeInput)(nil)).Elem(), &HoneypotProbe{})
	pulumi.RegisterInputType(reflect.TypeOf((*HoneypotProbeArrayInput)(nil)).Elem(), HoneypotProbeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HoneypotProbeMapInput)(nil)).Elem(), HoneypotProbeMap{})
	pulumi.RegisterOutputType(HoneypotProbeOutput{})
	pulumi.RegisterOutputType(HoneypotProbeArrayOutput{})
	pulumi.RegisterOutputType(HoneypotProbeMapOutput{})
}