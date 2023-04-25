// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Global Accelerator (GA) Basic Accelerate Ip Endpoint Relation resource.
//
// For information about Global Accelerator (GA) Basic Accelerate Ip Endpoint Relation and how to use it, see [What is Basic Accelerate Ip Endpoint Relation](https://help.aliyun.com/document_detail/466842.html).
//
// > **NOTE:** Available in v1.194.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ga"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := alicloud.NewProvider(ctx, "sz", &alicloud.ProviderArgs{
//				Region: pulumi.String("cn-shenzhen"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = alicloud.NewProvider(ctx, "hz", &alicloud.ProviderArgs{
//				Region: pulumi.String("cn-hangzhou"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultNetworks, err := vpc.GetNetworks(ctx, &vpc.GetNetworksArgs{
//				NameRegex: pulumi.StringRef("your_vpc_name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultSwitches, err := vpc.GetSwitches(ctx, &vpc.GetSwitchesArgs{
//				VpcId: pulumi.StringRef(defaultNetworks.Ids[0]),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultSecurityGroup, err := ecs.NewSecurityGroup(ctx, "defaultSecurityGroup", &ecs.SecurityGroupArgs{
//				VpcId: *pulumi.String(defaultNetworks.Ids[0]),
//			}, pulumi.Provider("alicloud.sz"))
//			if err != nil {
//				return err
//			}
//			defaultEcsNetworkInterface, err := ecs.NewEcsNetworkInterface(ctx, "defaultEcsNetworkInterface", &ecs.EcsNetworkInterfaceArgs{
//				VswitchId: *pulumi.String(defaultSwitches.Ids[0]),
//				SecurityGroupIds: pulumi.StringArray{
//					defaultSecurityGroup.ID(),
//				},
//			}, pulumi.Provider("alicloud.sz"))
//			if err != nil {
//				return err
//			}
//			defaultBasicAccelerator, err := ga.NewBasicAccelerator(ctx, "defaultBasicAccelerator", &ga.BasicAcceleratorArgs{
//				Duration:             pulumi.Int(1),
//				PricingCycle:         pulumi.String("Month"),
//				BasicAcceleratorName: pulumi.Any(_var.Name),
//				Description:          pulumi.Any(_var.Name),
//				BandwidthBillingType: pulumi.String("CDT"),
//				AutoPay:              pulumi.Bool(true),
//				AutoUseCoupon:        pulumi.String("true"),
//				AutoRenew:            pulumi.Bool(false),
//				AutoRenewDuration:    pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBasicIpSet, err := ga.NewBasicIpSet(ctx, "defaultBasicIpSet", &ga.BasicIpSetArgs{
//				AcceleratorId:      defaultBasicAccelerator.ID(),
//				AccelerateRegionId: pulumi.String("cn-hangzhou"),
//				IspType:            pulumi.String("BGP"),
//				Bandwidth:          pulumi.Int(5),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBasicAccelerateIp, err := ga.NewBasicAccelerateIp(ctx, "defaultBasicAccelerateIp", &ga.BasicAccelerateIpArgs{
//				AcceleratorId: defaultBasicIpSet.AcceleratorId,
//				IpSetId:       defaultBasicIpSet.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBasicEndpointGroup, err := ga.NewBasicEndpointGroup(ctx, "defaultBasicEndpointGroup", &ga.BasicEndpointGroupArgs{
//				AcceleratorId:       defaultBasicAccelerator.ID(),
//				EndpointGroupRegion: pulumi.String("cn-shenzhen"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBasicEndpoint, err := ga.NewBasicEndpoint(ctx, "defaultBasicEndpoint", &ga.BasicEndpointArgs{
//				AcceleratorId:          defaultBasicAccelerator.ID(),
//				EndpointGroupId:        defaultBasicEndpointGroup.ID(),
//				EndpointType:           pulumi.String("ENI"),
//				EndpointAddress:        defaultEcsNetworkInterface.ID(),
//				EndpointSubAddressType: pulumi.String("primary"),
//				EndpointSubAddress:     pulumi.String("192.168.0.1"),
//				BasicEndpointName:      pulumi.Any(_var.Name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ga.NewBasicAccelerateIpEndpointRelation(ctx, "defaultBasicAccelerateIpEndpointRelation", &ga.BasicAccelerateIpEndpointRelationArgs{
//				AcceleratorId:  defaultBasicAccelerateIp.AcceleratorId,
//				AccelerateIpId: defaultBasicAccelerateIp.ID(),
//				EndpointId:     defaultBasicEndpoint.EndpointId,
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
// Global Accelerator (GA) Basic Accelerate Ip Endpoint Relation can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ga/basicAccelerateIpEndpointRelation:BasicAccelerateIpEndpointRelation example <accelerator_id>:<accelerate_ip_id>:<endpoint_id>
//
// ```
type BasicAccelerateIpEndpointRelation struct {
	pulumi.CustomResourceState

	// The ID of the Basic Accelerate IP.
	AccelerateIpId pulumi.StringOutput `pulumi:"accelerateIpId"`
	// The ID of the Basic GA instance.
	AcceleratorId pulumi.StringOutput `pulumi:"acceleratorId"`
	// The ID of the Basic Endpoint.
	EndpointId pulumi.StringOutput `pulumi:"endpointId"`
	// The status of the Basic Accelerate Ip Endpoint Relation.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewBasicAccelerateIpEndpointRelation registers a new resource with the given unique name, arguments, and options.
func NewBasicAccelerateIpEndpointRelation(ctx *pulumi.Context,
	name string, args *BasicAccelerateIpEndpointRelationArgs, opts ...pulumi.ResourceOption) (*BasicAccelerateIpEndpointRelation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccelerateIpId == nil {
		return nil, errors.New("invalid value for required argument 'AccelerateIpId'")
	}
	if args.AcceleratorId == nil {
		return nil, errors.New("invalid value for required argument 'AcceleratorId'")
	}
	if args.EndpointId == nil {
		return nil, errors.New("invalid value for required argument 'EndpointId'")
	}
	var resource BasicAccelerateIpEndpointRelation
	err := ctx.RegisterResource("alicloud:ga/basicAccelerateIpEndpointRelation:BasicAccelerateIpEndpointRelation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBasicAccelerateIpEndpointRelation gets an existing BasicAccelerateIpEndpointRelation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBasicAccelerateIpEndpointRelation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BasicAccelerateIpEndpointRelationState, opts ...pulumi.ResourceOption) (*BasicAccelerateIpEndpointRelation, error) {
	var resource BasicAccelerateIpEndpointRelation
	err := ctx.ReadResource("alicloud:ga/basicAccelerateIpEndpointRelation:BasicAccelerateIpEndpointRelation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BasicAccelerateIpEndpointRelation resources.
type basicAccelerateIpEndpointRelationState struct {
	// The ID of the Basic Accelerate IP.
	AccelerateIpId *string `pulumi:"accelerateIpId"`
	// The ID of the Basic GA instance.
	AcceleratorId *string `pulumi:"acceleratorId"`
	// The ID of the Basic Endpoint.
	EndpointId *string `pulumi:"endpointId"`
	// The status of the Basic Accelerate Ip Endpoint Relation.
	Status *string `pulumi:"status"`
}

type BasicAccelerateIpEndpointRelationState struct {
	// The ID of the Basic Accelerate IP.
	AccelerateIpId pulumi.StringPtrInput
	// The ID of the Basic GA instance.
	AcceleratorId pulumi.StringPtrInput
	// The ID of the Basic Endpoint.
	EndpointId pulumi.StringPtrInput
	// The status of the Basic Accelerate Ip Endpoint Relation.
	Status pulumi.StringPtrInput
}

func (BasicAccelerateIpEndpointRelationState) ElementType() reflect.Type {
	return reflect.TypeOf((*basicAccelerateIpEndpointRelationState)(nil)).Elem()
}

type basicAccelerateIpEndpointRelationArgs struct {
	// The ID of the Basic Accelerate IP.
	AccelerateIpId string `pulumi:"accelerateIpId"`
	// The ID of the Basic GA instance.
	AcceleratorId string `pulumi:"acceleratorId"`
	// The ID of the Basic Endpoint.
	EndpointId string `pulumi:"endpointId"`
}

// The set of arguments for constructing a BasicAccelerateIpEndpointRelation resource.
type BasicAccelerateIpEndpointRelationArgs struct {
	// The ID of the Basic Accelerate IP.
	AccelerateIpId pulumi.StringInput
	// The ID of the Basic GA instance.
	AcceleratorId pulumi.StringInput
	// The ID of the Basic Endpoint.
	EndpointId pulumi.StringInput
}

func (BasicAccelerateIpEndpointRelationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*basicAccelerateIpEndpointRelationArgs)(nil)).Elem()
}

type BasicAccelerateIpEndpointRelationInput interface {
	pulumi.Input

	ToBasicAccelerateIpEndpointRelationOutput() BasicAccelerateIpEndpointRelationOutput
	ToBasicAccelerateIpEndpointRelationOutputWithContext(ctx context.Context) BasicAccelerateIpEndpointRelationOutput
}

func (*BasicAccelerateIpEndpointRelation) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicAccelerateIpEndpointRelation)(nil)).Elem()
}

func (i *BasicAccelerateIpEndpointRelation) ToBasicAccelerateIpEndpointRelationOutput() BasicAccelerateIpEndpointRelationOutput {
	return i.ToBasicAccelerateIpEndpointRelationOutputWithContext(context.Background())
}

func (i *BasicAccelerateIpEndpointRelation) ToBasicAccelerateIpEndpointRelationOutputWithContext(ctx context.Context) BasicAccelerateIpEndpointRelationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicAccelerateIpEndpointRelationOutput)
}

// BasicAccelerateIpEndpointRelationArrayInput is an input type that accepts BasicAccelerateIpEndpointRelationArray and BasicAccelerateIpEndpointRelationArrayOutput values.
// You can construct a concrete instance of `BasicAccelerateIpEndpointRelationArrayInput` via:
//
//	BasicAccelerateIpEndpointRelationArray{ BasicAccelerateIpEndpointRelationArgs{...} }
type BasicAccelerateIpEndpointRelationArrayInput interface {
	pulumi.Input

	ToBasicAccelerateIpEndpointRelationArrayOutput() BasicAccelerateIpEndpointRelationArrayOutput
	ToBasicAccelerateIpEndpointRelationArrayOutputWithContext(context.Context) BasicAccelerateIpEndpointRelationArrayOutput
}

type BasicAccelerateIpEndpointRelationArray []BasicAccelerateIpEndpointRelationInput

func (BasicAccelerateIpEndpointRelationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BasicAccelerateIpEndpointRelation)(nil)).Elem()
}

func (i BasicAccelerateIpEndpointRelationArray) ToBasicAccelerateIpEndpointRelationArrayOutput() BasicAccelerateIpEndpointRelationArrayOutput {
	return i.ToBasicAccelerateIpEndpointRelationArrayOutputWithContext(context.Background())
}

func (i BasicAccelerateIpEndpointRelationArray) ToBasicAccelerateIpEndpointRelationArrayOutputWithContext(ctx context.Context) BasicAccelerateIpEndpointRelationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicAccelerateIpEndpointRelationArrayOutput)
}

// BasicAccelerateIpEndpointRelationMapInput is an input type that accepts BasicAccelerateIpEndpointRelationMap and BasicAccelerateIpEndpointRelationMapOutput values.
// You can construct a concrete instance of `BasicAccelerateIpEndpointRelationMapInput` via:
//
//	BasicAccelerateIpEndpointRelationMap{ "key": BasicAccelerateIpEndpointRelationArgs{...} }
type BasicAccelerateIpEndpointRelationMapInput interface {
	pulumi.Input

	ToBasicAccelerateIpEndpointRelationMapOutput() BasicAccelerateIpEndpointRelationMapOutput
	ToBasicAccelerateIpEndpointRelationMapOutputWithContext(context.Context) BasicAccelerateIpEndpointRelationMapOutput
}

type BasicAccelerateIpEndpointRelationMap map[string]BasicAccelerateIpEndpointRelationInput

func (BasicAccelerateIpEndpointRelationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BasicAccelerateIpEndpointRelation)(nil)).Elem()
}

func (i BasicAccelerateIpEndpointRelationMap) ToBasicAccelerateIpEndpointRelationMapOutput() BasicAccelerateIpEndpointRelationMapOutput {
	return i.ToBasicAccelerateIpEndpointRelationMapOutputWithContext(context.Background())
}

func (i BasicAccelerateIpEndpointRelationMap) ToBasicAccelerateIpEndpointRelationMapOutputWithContext(ctx context.Context) BasicAccelerateIpEndpointRelationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicAccelerateIpEndpointRelationMapOutput)
}

type BasicAccelerateIpEndpointRelationOutput struct{ *pulumi.OutputState }

func (BasicAccelerateIpEndpointRelationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicAccelerateIpEndpointRelation)(nil)).Elem()
}

func (o BasicAccelerateIpEndpointRelationOutput) ToBasicAccelerateIpEndpointRelationOutput() BasicAccelerateIpEndpointRelationOutput {
	return o
}

func (o BasicAccelerateIpEndpointRelationOutput) ToBasicAccelerateIpEndpointRelationOutputWithContext(ctx context.Context) BasicAccelerateIpEndpointRelationOutput {
	return o
}

// The ID of the Basic Accelerate IP.
func (o BasicAccelerateIpEndpointRelationOutput) AccelerateIpId() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicAccelerateIpEndpointRelation) pulumi.StringOutput { return v.AccelerateIpId }).(pulumi.StringOutput)
}

// The ID of the Basic GA instance.
func (o BasicAccelerateIpEndpointRelationOutput) AcceleratorId() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicAccelerateIpEndpointRelation) pulumi.StringOutput { return v.AcceleratorId }).(pulumi.StringOutput)
}

// The ID of the Basic Endpoint.
func (o BasicAccelerateIpEndpointRelationOutput) EndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicAccelerateIpEndpointRelation) pulumi.StringOutput { return v.EndpointId }).(pulumi.StringOutput)
}

// The status of the Basic Accelerate Ip Endpoint Relation.
func (o BasicAccelerateIpEndpointRelationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicAccelerateIpEndpointRelation) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type BasicAccelerateIpEndpointRelationArrayOutput struct{ *pulumi.OutputState }

func (BasicAccelerateIpEndpointRelationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BasicAccelerateIpEndpointRelation)(nil)).Elem()
}

func (o BasicAccelerateIpEndpointRelationArrayOutput) ToBasicAccelerateIpEndpointRelationArrayOutput() BasicAccelerateIpEndpointRelationArrayOutput {
	return o
}

func (o BasicAccelerateIpEndpointRelationArrayOutput) ToBasicAccelerateIpEndpointRelationArrayOutputWithContext(ctx context.Context) BasicAccelerateIpEndpointRelationArrayOutput {
	return o
}

func (o BasicAccelerateIpEndpointRelationArrayOutput) Index(i pulumi.IntInput) BasicAccelerateIpEndpointRelationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BasicAccelerateIpEndpointRelation {
		return vs[0].([]*BasicAccelerateIpEndpointRelation)[vs[1].(int)]
	}).(BasicAccelerateIpEndpointRelationOutput)
}

type BasicAccelerateIpEndpointRelationMapOutput struct{ *pulumi.OutputState }

func (BasicAccelerateIpEndpointRelationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BasicAccelerateIpEndpointRelation)(nil)).Elem()
}

func (o BasicAccelerateIpEndpointRelationMapOutput) ToBasicAccelerateIpEndpointRelationMapOutput() BasicAccelerateIpEndpointRelationMapOutput {
	return o
}

func (o BasicAccelerateIpEndpointRelationMapOutput) ToBasicAccelerateIpEndpointRelationMapOutputWithContext(ctx context.Context) BasicAccelerateIpEndpointRelationMapOutput {
	return o
}

func (o BasicAccelerateIpEndpointRelationMapOutput) MapIndex(k pulumi.StringInput) BasicAccelerateIpEndpointRelationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BasicAccelerateIpEndpointRelation {
		return vs[0].(map[string]*BasicAccelerateIpEndpointRelation)[vs[1].(string)]
	}).(BasicAccelerateIpEndpointRelationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BasicAccelerateIpEndpointRelationInput)(nil)).Elem(), &BasicAccelerateIpEndpointRelation{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicAccelerateIpEndpointRelationArrayInput)(nil)).Elem(), BasicAccelerateIpEndpointRelationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicAccelerateIpEndpointRelationMapInput)(nil)).Elem(), BasicAccelerateIpEndpointRelationMap{})
	pulumi.RegisterOutputType(BasicAccelerateIpEndpointRelationOutput{})
	pulumi.RegisterOutputType(BasicAccelerateIpEndpointRelationArrayOutput{})
	pulumi.RegisterOutputType(BasicAccelerateIpEndpointRelationMapOutput{})
}
