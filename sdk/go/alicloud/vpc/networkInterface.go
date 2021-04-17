// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an ECS Elastic Network Interface resource.
//
// For information about Elastic Network Interface and how to use it, see [Elastic Network Interface](https://www.alibabacloud.com/help/doc-detail/58496.html).
//
// > **NOTE** Only one of privateIps or privateIpsCount can be specified when assign private IPs.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		name := "networkInterfaceName"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		vpc, err := vpc.NewNetwork(ctx, "vpc", &vpc.NetworkArgs{
// 			VpcName:   pulumi.String(name),
// 			CidrBlock: pulumi.String("192.168.0.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		opt0 := "VSwitch"
// 		defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
// 			AvailableResourceCreation: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		vswitch, err := vpc.NewSwitch(ctx, "vswitch", &vpc.SwitchArgs{
// 			CidrBlock:        pulumi.String("192.168.0.0/24"),
// 			AvailabilityZone: pulumi.String(defaultZones.Zones[0].Id),
// 			VpcId:            vpc.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		group, err := ecs.NewSecurityGroup(ctx, "group", &ecs.SecurityGroupArgs{
// 			VpcId: vpc.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vpc.NewNetworkInterface(ctx, "defaultNetworkInterface", &vpc.NetworkInterfaceArgs{
// 			VswitchId: vswitch.ID(),
// 			SecurityGroups: pulumi.StringArray{
// 				group.ID(),
// 			},
// 			PrivateIp:       pulumi.String("192.168.0.2"),
// 			PrivateIpsCount: pulumi.Int(3),
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
// ENI can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:vpc/networkInterface:NetworkInterface eni eni-abc1234567890000
// ```
type NetworkInterface struct {
	pulumi.CustomResourceState

	// Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// (Available in 1.54.0+) The MAC address of an ENI.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	Name pulumi.StringOutput `pulumi:"name"`
	// The primary private IP of the ENI.
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	PrivateIps pulumi.StringArrayOutput `pulumi:"privateIps"`
	// Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	PrivateIpsCount pulumi.IntOutput `pulumi:"privateIpsCount"`
	// The Id of resource group which the network interface belongs.
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	// A list of security group ids to associate with.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The VSwitch to create the ENI in.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewNetworkInterface registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterface(ctx *pulumi.Context,
	name string, args *NetworkInterfaceArgs, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SecurityGroups == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroups'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	var resource NetworkInterface
	err := ctx.RegisterResource("alicloud:vpc/networkInterface:NetworkInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterface gets an existing NetworkInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceState, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	var resource NetworkInterface
	err := ctx.ReadResource("alicloud:vpc/networkInterface:NetworkInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterface resources.
type networkInterfaceState struct {
	// Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description *string `pulumi:"description"`
	// (Available in 1.54.0+) The MAC address of an ENI.
	Mac *string `pulumi:"mac"`
	// Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	Name *string `pulumi:"name"`
	// The primary private IP of the ENI.
	PrivateIp *string `pulumi:"privateIp"`
	// List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	PrivateIps []string `pulumi:"privateIps"`
	// Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	PrivateIpsCount *int `pulumi:"privateIpsCount"`
	// The Id of resource group which the network interface belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A list of security group ids to associate with.
	SecurityGroups []string `pulumi:"securityGroups"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VSwitch to create the ENI in.
	VswitchId *string `pulumi:"vswitchId"`
}

type NetworkInterfaceState struct {
	// Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description pulumi.StringPtrInput
	// (Available in 1.54.0+) The MAC address of an ENI.
	Mac pulumi.StringPtrInput
	// Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	Name pulumi.StringPtrInput
	// The primary private IP of the ENI.
	PrivateIp pulumi.StringPtrInput
	// List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	PrivateIps pulumi.StringArrayInput
	// Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	PrivateIpsCount pulumi.IntPtrInput
	// The Id of resource group which the network interface belongs.
	ResourceGroupId pulumi.StringPtrInput
	// A list of security group ids to associate with.
	SecurityGroups pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The VSwitch to create the ENI in.
	VswitchId pulumi.StringPtrInput
}

func (NetworkInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceState)(nil)).Elem()
}

type networkInterfaceArgs struct {
	// Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description *string `pulumi:"description"`
	// Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	Name *string `pulumi:"name"`
	// The primary private IP of the ENI.
	PrivateIp *string `pulumi:"privateIp"`
	// List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	PrivateIps []string `pulumi:"privateIps"`
	// Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	PrivateIpsCount *int `pulumi:"privateIpsCount"`
	// The Id of resource group which the network interface belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A list of security group ids to associate with.
	SecurityGroups []string `pulumi:"securityGroups"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VSwitch to create the ENI in.
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a NetworkInterface resource.
type NetworkInterfaceArgs struct {
	// Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description pulumi.StringPtrInput
	// Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	Name pulumi.StringPtrInput
	// The primary private IP of the ENI.
	PrivateIp pulumi.StringPtrInput
	// List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	PrivateIps pulumi.StringArrayInput
	// Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	PrivateIpsCount pulumi.IntPtrInput
	// The Id of resource group which the network interface belongs.
	ResourceGroupId pulumi.StringPtrInput
	// A list of security group ids to associate with.
	SecurityGroups pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The VSwitch to create the ENI in.
	VswitchId pulumi.StringInput
}

func (NetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceArgs)(nil)).Elem()
}

type NetworkInterfaceInput interface {
	pulumi.Input

	ToNetworkInterfaceOutput() NetworkInterfaceOutput
	ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput
}

func (*NetworkInterface) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterface)(nil))
}

func (i *NetworkInterface) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return i.ToNetworkInterfaceOutputWithContext(context.Background())
}

func (i *NetworkInterface) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceOutput)
}

func (i *NetworkInterface) ToNetworkInterfacePtrOutput() NetworkInterfacePtrOutput {
	return i.ToNetworkInterfacePtrOutputWithContext(context.Background())
}

func (i *NetworkInterface) ToNetworkInterfacePtrOutputWithContext(ctx context.Context) NetworkInterfacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfacePtrOutput)
}

type NetworkInterfacePtrInput interface {
	pulumi.Input

	ToNetworkInterfacePtrOutput() NetworkInterfacePtrOutput
	ToNetworkInterfacePtrOutputWithContext(ctx context.Context) NetworkInterfacePtrOutput
}

type networkInterfacePtrType NetworkInterfaceArgs

func (*networkInterfacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterface)(nil))
}

func (i *networkInterfacePtrType) ToNetworkInterfacePtrOutput() NetworkInterfacePtrOutput {
	return i.ToNetworkInterfacePtrOutputWithContext(context.Background())
}

func (i *networkInterfacePtrType) ToNetworkInterfacePtrOutputWithContext(ctx context.Context) NetworkInterfacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfacePtrOutput)
}

// NetworkInterfaceArrayInput is an input type that accepts NetworkInterfaceArray and NetworkInterfaceArrayOutput values.
// You can construct a concrete instance of `NetworkInterfaceArrayInput` via:
//
//          NetworkInterfaceArray{ NetworkInterfaceArgs{...} }
type NetworkInterfaceArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput
	ToNetworkInterfaceArrayOutputWithContext(context.Context) NetworkInterfaceArrayOutput
}

type NetworkInterfaceArray []NetworkInterfaceInput

func (NetworkInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*NetworkInterface)(nil))
}

func (i NetworkInterfaceArray) ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput {
	return i.ToNetworkInterfaceArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceArray) ToNetworkInterfaceArrayOutputWithContext(ctx context.Context) NetworkInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceArrayOutput)
}

// NetworkInterfaceMapInput is an input type that accepts NetworkInterfaceMap and NetworkInterfaceMapOutput values.
// You can construct a concrete instance of `NetworkInterfaceMapInput` via:
//
//          NetworkInterfaceMap{ "key": NetworkInterfaceArgs{...} }
type NetworkInterfaceMapInput interface {
	pulumi.Input

	ToNetworkInterfaceMapOutput() NetworkInterfaceMapOutput
	ToNetworkInterfaceMapOutputWithContext(context.Context) NetworkInterfaceMapOutput
}

type NetworkInterfaceMap map[string]NetworkInterfaceInput

func (NetworkInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*NetworkInterface)(nil))
}

func (i NetworkInterfaceMap) ToNetworkInterfaceMapOutput() NetworkInterfaceMapOutput {
	return i.ToNetworkInterfaceMapOutputWithContext(context.Background())
}

func (i NetworkInterfaceMap) ToNetworkInterfaceMapOutputWithContext(ctx context.Context) NetworkInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceMapOutput)
}

type NetworkInterfaceOutput struct {
	*pulumi.OutputState
}

func (NetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterface)(nil))
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) ToNetworkInterfacePtrOutput() NetworkInterfacePtrOutput {
	return o.ToNetworkInterfacePtrOutputWithContext(context.Background())
}

func (o NetworkInterfaceOutput) ToNetworkInterfacePtrOutputWithContext(ctx context.Context) NetworkInterfacePtrOutput {
	return o.ApplyT(func(v NetworkInterface) *NetworkInterface {
		return &v
	}).(NetworkInterfacePtrOutput)
}

type NetworkInterfacePtrOutput struct {
	*pulumi.OutputState
}

func (NetworkInterfacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterface)(nil))
}

func (o NetworkInterfacePtrOutput) ToNetworkInterfacePtrOutput() NetworkInterfacePtrOutput {
	return o
}

func (o NetworkInterfacePtrOutput) ToNetworkInterfacePtrOutputWithContext(ctx context.Context) NetworkInterfacePtrOutput {
	return o
}

type NetworkInterfaceArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterface)(nil))
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutputWithContext(ctx context.Context) NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterface {
		return vs[0].([]NetworkInterface)[vs[1].(int)]
	}).(NetworkInterfaceOutput)
}

type NetworkInterfaceMapOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkInterface)(nil))
}

func (o NetworkInterfaceMapOutput) ToNetworkInterfaceMapOutput() NetworkInterfaceMapOutput {
	return o
}

func (o NetworkInterfaceMapOutput) ToNetworkInterfaceMapOutputWithContext(ctx context.Context) NetworkInterfaceMapOutput {
	return o
}

func (o NetworkInterfaceMapOutput) MapIndex(k pulumi.StringInput) NetworkInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkInterface {
		return vs[0].(map[string]NetworkInterface)[vs[1].(string)]
	}).(NetworkInterfaceOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
	pulumi.RegisterOutputType(NetworkInterfacePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceMapOutput{})
}
