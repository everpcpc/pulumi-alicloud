// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Global Accelerator (GA) Basic Endpoint resource.
//
// For information about Global Accelerator (GA) Basic Endpoint and how to use it, see [What is Basic Endpoint](https://help.aliyun.com/document_detail/466839.html).
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
//			defaultEcsNetworkInterface, err := ecs.NewEcsNetworkInterface(ctx, "defaultEcsNetworkInterface", &ecs.EcsNetworkInterfaceArgs{
//				VswitchId: pulumi.String("your_vswitch_id"),
//				SecurityGroupIds: pulumi.StringArray{
//					pulumi.String("your_security_group_id"),
//				},
//			}, pulumi.Provider("alicloud.sz"))
//			if err != nil {
//				return err
//			}
//			_, err = ga.NewBasicEndpoint(ctx, "defaultBasicEndpoint", &ga.BasicEndpointArgs{
//				AcceleratorId:          pulumi.String("your_accelerator_id"),
//				EndpointGroupId:        pulumi.String("your_endpoint_group_id"),
//				EndpointType:           pulumi.String("ENI"),
//				EndpointAddress:        defaultEcsNetworkInterface.ID(),
//				EndpointSubAddressType: pulumi.String("secondary"),
//				EndpointSubAddress:     pulumi.String("192.168.0.1"),
//				BasicEndpointName:      pulumi.String("example_value"),
//			}, pulumi.Provider("alicloud.hz"))
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
// Global Accelerator (GA) Basic Endpoint can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ga/basicEndpoint:BasicEndpoint example <endpoint_group_id>:<endpoint_id>
//
// ```
type BasicEndpoint struct {
	pulumi.CustomResourceState

	// The ID of the Basic GA instance.
	AcceleratorId pulumi.StringOutput `pulumi:"acceleratorId"`
	// The name of the Basic Endpoint.
	BasicEndpointName pulumi.StringPtrOutput `pulumi:"basicEndpointName"`
	// The address of the Basic Endpoint.
	EndpointAddress pulumi.StringOutput `pulumi:"endpointAddress"`
	// The ID of the Basic Endpoint Group.
	EndpointGroupId pulumi.StringOutput `pulumi:"endpointGroupId"`
	// The ID of the Basic Endpoint.
	EndpointId pulumi.StringOutput `pulumi:"endpointId"`
	// The sub address of the Basic Endpoint.
	EndpointSubAddress pulumi.StringPtrOutput `pulumi:"endpointSubAddress"`
	// The sub address type of the Basic Endpoint. Valid values: `primary`, `secondary`.
	EndpointSubAddressType pulumi.StringPtrOutput `pulumi:"endpointSubAddressType"`
	// The type of the Basic Endpoint. Valid values: `ENI`, `SLB`, `ECS` and `NLB`.
	EndpointType pulumi.StringOutput `pulumi:"endpointType"`
	// The zone id of the Basic Endpoint.
	EndpointZoneId pulumi.StringPtrOutput `pulumi:"endpointZoneId"`
	// The status of the Basic Endpoint.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewBasicEndpoint registers a new resource with the given unique name, arguments, and options.
func NewBasicEndpoint(ctx *pulumi.Context,
	name string, args *BasicEndpointArgs, opts ...pulumi.ResourceOption) (*BasicEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AcceleratorId == nil {
		return nil, errors.New("invalid value for required argument 'AcceleratorId'")
	}
	if args.EndpointAddress == nil {
		return nil, errors.New("invalid value for required argument 'EndpointAddress'")
	}
	if args.EndpointGroupId == nil {
		return nil, errors.New("invalid value for required argument 'EndpointGroupId'")
	}
	if args.EndpointType == nil {
		return nil, errors.New("invalid value for required argument 'EndpointType'")
	}
	var resource BasicEndpoint
	err := ctx.RegisterResource("alicloud:ga/basicEndpoint:BasicEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBasicEndpoint gets an existing BasicEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBasicEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BasicEndpointState, opts ...pulumi.ResourceOption) (*BasicEndpoint, error) {
	var resource BasicEndpoint
	err := ctx.ReadResource("alicloud:ga/basicEndpoint:BasicEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BasicEndpoint resources.
type basicEndpointState struct {
	// The ID of the Basic GA instance.
	AcceleratorId *string `pulumi:"acceleratorId"`
	// The name of the Basic Endpoint.
	BasicEndpointName *string `pulumi:"basicEndpointName"`
	// The address of the Basic Endpoint.
	EndpointAddress *string `pulumi:"endpointAddress"`
	// The ID of the Basic Endpoint Group.
	EndpointGroupId *string `pulumi:"endpointGroupId"`
	// The ID of the Basic Endpoint.
	EndpointId *string `pulumi:"endpointId"`
	// The sub address of the Basic Endpoint.
	EndpointSubAddress *string `pulumi:"endpointSubAddress"`
	// The sub address type of the Basic Endpoint. Valid values: `primary`, `secondary`.
	EndpointSubAddressType *string `pulumi:"endpointSubAddressType"`
	// The type of the Basic Endpoint. Valid values: `ENI`, `SLB`, `ECS` and `NLB`.
	EndpointType *string `pulumi:"endpointType"`
	// The zone id of the Basic Endpoint.
	EndpointZoneId *string `pulumi:"endpointZoneId"`
	// The status of the Basic Endpoint.
	Status *string `pulumi:"status"`
}

type BasicEndpointState struct {
	// The ID of the Basic GA instance.
	AcceleratorId pulumi.StringPtrInput
	// The name of the Basic Endpoint.
	BasicEndpointName pulumi.StringPtrInput
	// The address of the Basic Endpoint.
	EndpointAddress pulumi.StringPtrInput
	// The ID of the Basic Endpoint Group.
	EndpointGroupId pulumi.StringPtrInput
	// The ID of the Basic Endpoint.
	EndpointId pulumi.StringPtrInput
	// The sub address of the Basic Endpoint.
	EndpointSubAddress pulumi.StringPtrInput
	// The sub address type of the Basic Endpoint. Valid values: `primary`, `secondary`.
	EndpointSubAddressType pulumi.StringPtrInput
	// The type of the Basic Endpoint. Valid values: `ENI`, `SLB`, `ECS` and `NLB`.
	EndpointType pulumi.StringPtrInput
	// The zone id of the Basic Endpoint.
	EndpointZoneId pulumi.StringPtrInput
	// The status of the Basic Endpoint.
	Status pulumi.StringPtrInput
}

func (BasicEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*basicEndpointState)(nil)).Elem()
}

type basicEndpointArgs struct {
	// The ID of the Basic GA instance.
	AcceleratorId string `pulumi:"acceleratorId"`
	// The name of the Basic Endpoint.
	BasicEndpointName *string `pulumi:"basicEndpointName"`
	// The address of the Basic Endpoint.
	EndpointAddress string `pulumi:"endpointAddress"`
	// The ID of the Basic Endpoint Group.
	EndpointGroupId string `pulumi:"endpointGroupId"`
	// The sub address of the Basic Endpoint.
	EndpointSubAddress *string `pulumi:"endpointSubAddress"`
	// The sub address type of the Basic Endpoint. Valid values: `primary`, `secondary`.
	EndpointSubAddressType *string `pulumi:"endpointSubAddressType"`
	// The type of the Basic Endpoint. Valid values: `ENI`, `SLB`, `ECS` and `NLB`.
	EndpointType string `pulumi:"endpointType"`
	// The zone id of the Basic Endpoint.
	EndpointZoneId *string `pulumi:"endpointZoneId"`
}

// The set of arguments for constructing a BasicEndpoint resource.
type BasicEndpointArgs struct {
	// The ID of the Basic GA instance.
	AcceleratorId pulumi.StringInput
	// The name of the Basic Endpoint.
	BasicEndpointName pulumi.StringPtrInput
	// The address of the Basic Endpoint.
	EndpointAddress pulumi.StringInput
	// The ID of the Basic Endpoint Group.
	EndpointGroupId pulumi.StringInput
	// The sub address of the Basic Endpoint.
	EndpointSubAddress pulumi.StringPtrInput
	// The sub address type of the Basic Endpoint. Valid values: `primary`, `secondary`.
	EndpointSubAddressType pulumi.StringPtrInput
	// The type of the Basic Endpoint. Valid values: `ENI`, `SLB`, `ECS` and `NLB`.
	EndpointType pulumi.StringInput
	// The zone id of the Basic Endpoint.
	EndpointZoneId pulumi.StringPtrInput
}

func (BasicEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*basicEndpointArgs)(nil)).Elem()
}

type BasicEndpointInput interface {
	pulumi.Input

	ToBasicEndpointOutput() BasicEndpointOutput
	ToBasicEndpointOutputWithContext(ctx context.Context) BasicEndpointOutput
}

func (*BasicEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicEndpoint)(nil)).Elem()
}

func (i *BasicEndpoint) ToBasicEndpointOutput() BasicEndpointOutput {
	return i.ToBasicEndpointOutputWithContext(context.Background())
}

func (i *BasicEndpoint) ToBasicEndpointOutputWithContext(ctx context.Context) BasicEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicEndpointOutput)
}

// BasicEndpointArrayInput is an input type that accepts BasicEndpointArray and BasicEndpointArrayOutput values.
// You can construct a concrete instance of `BasicEndpointArrayInput` via:
//
//	BasicEndpointArray{ BasicEndpointArgs{...} }
type BasicEndpointArrayInput interface {
	pulumi.Input

	ToBasicEndpointArrayOutput() BasicEndpointArrayOutput
	ToBasicEndpointArrayOutputWithContext(context.Context) BasicEndpointArrayOutput
}

type BasicEndpointArray []BasicEndpointInput

func (BasicEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BasicEndpoint)(nil)).Elem()
}

func (i BasicEndpointArray) ToBasicEndpointArrayOutput() BasicEndpointArrayOutput {
	return i.ToBasicEndpointArrayOutputWithContext(context.Background())
}

func (i BasicEndpointArray) ToBasicEndpointArrayOutputWithContext(ctx context.Context) BasicEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicEndpointArrayOutput)
}

// BasicEndpointMapInput is an input type that accepts BasicEndpointMap and BasicEndpointMapOutput values.
// You can construct a concrete instance of `BasicEndpointMapInput` via:
//
//	BasicEndpointMap{ "key": BasicEndpointArgs{...} }
type BasicEndpointMapInput interface {
	pulumi.Input

	ToBasicEndpointMapOutput() BasicEndpointMapOutput
	ToBasicEndpointMapOutputWithContext(context.Context) BasicEndpointMapOutput
}

type BasicEndpointMap map[string]BasicEndpointInput

func (BasicEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BasicEndpoint)(nil)).Elem()
}

func (i BasicEndpointMap) ToBasicEndpointMapOutput() BasicEndpointMapOutput {
	return i.ToBasicEndpointMapOutputWithContext(context.Background())
}

func (i BasicEndpointMap) ToBasicEndpointMapOutputWithContext(ctx context.Context) BasicEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicEndpointMapOutput)
}

type BasicEndpointOutput struct{ *pulumi.OutputState }

func (BasicEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicEndpoint)(nil)).Elem()
}

func (o BasicEndpointOutput) ToBasicEndpointOutput() BasicEndpointOutput {
	return o
}

func (o BasicEndpointOutput) ToBasicEndpointOutputWithContext(ctx context.Context) BasicEndpointOutput {
	return o
}

// The ID of the Basic GA instance.
func (o BasicEndpointOutput) AcceleratorId() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicEndpoint) pulumi.StringOutput { return v.AcceleratorId }).(pulumi.StringOutput)
}

// The name of the Basic Endpoint.
func (o BasicEndpointOutput) BasicEndpointName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicEndpoint) pulumi.StringPtrOutput { return v.BasicEndpointName }).(pulumi.StringPtrOutput)
}

// The address of the Basic Endpoint.
func (o BasicEndpointOutput) EndpointAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicEndpoint) pulumi.StringOutput { return v.EndpointAddress }).(pulumi.StringOutput)
}

// The ID of the Basic Endpoint Group.
func (o BasicEndpointOutput) EndpointGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicEndpoint) pulumi.StringOutput { return v.EndpointGroupId }).(pulumi.StringOutput)
}

// The ID of the Basic Endpoint.
func (o BasicEndpointOutput) EndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicEndpoint) pulumi.StringOutput { return v.EndpointId }).(pulumi.StringOutput)
}

// The sub address of the Basic Endpoint.
func (o BasicEndpointOutput) EndpointSubAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicEndpoint) pulumi.StringPtrOutput { return v.EndpointSubAddress }).(pulumi.StringPtrOutput)
}

// The sub address type of the Basic Endpoint. Valid values: `primary`, `secondary`.
func (o BasicEndpointOutput) EndpointSubAddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicEndpoint) pulumi.StringPtrOutput { return v.EndpointSubAddressType }).(pulumi.StringPtrOutput)
}

// The type of the Basic Endpoint. Valid values: `ENI`, `SLB`, `ECS` and `NLB`.
func (o BasicEndpointOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicEndpoint) pulumi.StringOutput { return v.EndpointType }).(pulumi.StringOutput)
}

// The zone id of the Basic Endpoint.
func (o BasicEndpointOutput) EndpointZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicEndpoint) pulumi.StringPtrOutput { return v.EndpointZoneId }).(pulumi.StringPtrOutput)
}

// The status of the Basic Endpoint.
func (o BasicEndpointOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicEndpoint) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type BasicEndpointArrayOutput struct{ *pulumi.OutputState }

func (BasicEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BasicEndpoint)(nil)).Elem()
}

func (o BasicEndpointArrayOutput) ToBasicEndpointArrayOutput() BasicEndpointArrayOutput {
	return o
}

func (o BasicEndpointArrayOutput) ToBasicEndpointArrayOutputWithContext(ctx context.Context) BasicEndpointArrayOutput {
	return o
}

func (o BasicEndpointArrayOutput) Index(i pulumi.IntInput) BasicEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BasicEndpoint {
		return vs[0].([]*BasicEndpoint)[vs[1].(int)]
	}).(BasicEndpointOutput)
}

type BasicEndpointMapOutput struct{ *pulumi.OutputState }

func (BasicEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BasicEndpoint)(nil)).Elem()
}

func (o BasicEndpointMapOutput) ToBasicEndpointMapOutput() BasicEndpointMapOutput {
	return o
}

func (o BasicEndpointMapOutput) ToBasicEndpointMapOutputWithContext(ctx context.Context) BasicEndpointMapOutput {
	return o
}

func (o BasicEndpointMapOutput) MapIndex(k pulumi.StringInput) BasicEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BasicEndpoint {
		return vs[0].(map[string]*BasicEndpoint)[vs[1].(string)]
	}).(BasicEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BasicEndpointInput)(nil)).Elem(), &BasicEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicEndpointArrayInput)(nil)).Elem(), BasicEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicEndpointMapInput)(nil)).Elem(), BasicEndpointMap{})
	pulumi.RegisterOutputType(BasicEndpointOutput{})
	pulumi.RegisterOutputType(BasicEndpointArrayOutput{})
	pulumi.RegisterOutputType(BasicEndpointMapOutput{})
}
