// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eci

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides ECI Container Group resource.
//
// For information about ECI Container Group and how to use it, see [What is Container Group](https://www.alibabacloud.com/help/en/doc-detail/90341.htm).
//
// > **NOTE:** Available in v1.111.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eci"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := eci.NewContainerGroup(ctx, "example", &eci.ContainerGroupArgs{
// 			ContainerGroupName: pulumi.String("tf-testacc-eci-gruop"),
// 			Cpu:                pulumi.Float64(8),
// 			Memory:             pulumi.Float64(16),
// 			RestartPolicy:      pulumi.String("OnFailure"),
// 			SecurityGroupId:    pulumi.Any(alicloud_security_group.Group.Id),
// 			VswitchId:          pulumi.Any(data.Alicloud_vpcs.Default.Vpcs[0].Vswitch_ids[0]),
// 			Tags: pulumi.AnyMap{
// 				"TF": pulumi.Any("create"),
// 			},
// 			Containers: eci.ContainerGroupContainerArray{
// 				&eci.ContainerGroupContainerArgs{
// 					Image:           pulumi.String("registry-vpc.cn-beijing.aliyuncs.com/eci_open/nginx:alpine"),
// 					Name:            pulumi.String("nginx"),
// 					WorkingDir:      pulumi.String("/tmp/nginx"),
// 					ImagePullPolicy: pulumi.String("IfNotPresent"),
// 					Commands: pulumi.StringArray{
// 						pulumi.String("/bin/sh"),
// 						pulumi.String("-c"),
// 						pulumi.String("sleep 9999"),
// 					},
// 					VolumeMounts: eci.ContainerGroupContainerVolumeMountArray{
// 						&eci.ContainerGroupContainerVolumeMountArgs{
// 							MountPath: pulumi.String("/tmp/test"),
// 							ReadOnly:  pulumi.Bool(false),
// 							Name:      pulumi.String("empty1"),
// 						},
// 					},
// 					Ports: eci.ContainerGroupContainerPortArray{
// 						&eci.ContainerGroupContainerPortArgs{
// 							Port:     pulumi.Int(80),
// 							Protocol: pulumi.String("TCP"),
// 						},
// 					},
// 					EnvironmentVars: eci.ContainerGroupContainerEnvironmentVarArray{
// 						&eci.ContainerGroupContainerEnvironmentVarArgs{
// 							Key:   pulumi.String("test"),
// 							Value: pulumi.String("nginx"),
// 						},
// 					},
// 				},
// 				&eci.ContainerGroupContainerArgs{
// 					Image: pulumi.String("registry-vpc.cn-beijing.aliyuncs.com/eci_open/centos:7"),
// 					Name:  pulumi.String("centos"),
// 					Commands: pulumi.StringArray{
// 						pulumi.String("/bin/sh"),
// 						pulumi.String("-c"),
// 						pulumi.String("sleep 9999"),
// 					},
// 				},
// 			},
// 			InitContainers: eci.ContainerGroupInitContainerArray{
// 				&eci.ContainerGroupInitContainerArgs{
// 					Name:            pulumi.String("init-busybox"),
// 					Image:           pulumi.String("registry-vpc.cn-beijing.aliyuncs.com/eci_open/busybox:1.30"),
// 					ImagePullPolicy: pulumi.String("IfNotPresent"),
// 					Commands: pulumi.StringArray{
// 						pulumi.String("echo"),
// 					},
// 					Args: pulumi.StringArray{
// 						pulumi.String("hello initcontainer"),
// 					},
// 				},
// 			},
// 			Volumes: eci.ContainerGroupVolumeArray{
// 				&eci.ContainerGroupVolumeArgs{
// 					Name: pulumi.String("empty1"),
// 					Type: pulumi.String("EmptyDirVolume"),
// 				},
// 				&eci.ContainerGroupVolumeArgs{
// 					Name: pulumi.String("empty2"),
// 					Type: pulumi.String("EmptyDirVolume"),
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
// ECI Container Group can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:eci/containerGroup:ContainerGroup example <container_group_id>
// ```
type ContainerGroup struct {
	pulumi.CustomResourceState

	// The name of the container group.
	ContainerGroupName pulumi.StringOutput `pulumi:"containerGroupName"`
	// The list of containers.
	Containers ContainerGroupContainerArrayOutput `pulumi:"containers"`
	// The amount of CPU resources allocated to the container.
	Cpu pulumi.Float64PtrOutput `pulumi:"cpu"`
	// The structure of dnsConfig.
	DnsConfig          ContainerGroupDnsConfigPtrOutput          `pulumi:"dnsConfig"`
	EciSecurityContext ContainerGroupEciSecurityContextPtrOutput `pulumi:"eciSecurityContext"`
	// HostAliases.
	HostAliases ContainerGroupHostAliasArrayOutput `pulumi:"hostAliases"`
	// The image registry credential. The details see Block `imageRegistryCredential`.
	ImageRegistryCredentials ContainerGroupImageRegistryCredentialArrayOutput `pulumi:"imageRegistryCredentials"`
	// The list of initContainers.
	InitContainers ContainerGroupInitContainerArrayOutput `pulumi:"initContainers"`
	// The type of the ECS instance.
	InstanceType pulumi.StringPtrOutput `pulumi:"instanceType"`
	// The amount of memory resources allocated to the container.
	Memory pulumi.Float64PtrOutput `pulumi:"memory"`
	// The RAM role that the container group assumes. ECI and ECS share the same RAM role.
	RamRoleName pulumi.StringPtrOutput `pulumi:"ramRoleName"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The restart policy of the container group. Default to `Always`.
	RestartPolicy pulumi.StringPtrOutput `pulumi:"restartPolicy"`
	// The ID of the security group to which the container group belongs. Container groups within the same security group can access each other.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// The status of container group.
	Status pulumi.StringOutput `pulumi:"status"`
	Tags   pulumi.MapOutput    `pulumi:"tags"`
	// The list of volumes.
	Volumes ContainerGroupVolumeArrayOutput `pulumi:"volumes"`
	// The ID of the VSwitch. Currently, container groups can only be deployed in VPC networks. The number of IP addresses in the VSwitch CIDR block determines the maximum number of container groups that can be created in the VSwitch. Before you can create an ECI instance, plan the CIDR block of the VSwitch.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
	// The ID of the zone where you want to deploy the container group. If no value is specified, the system assigns a zone to the container group. By default, no value is specified.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewContainerGroup registers a new resource with the given unique name, arguments, and options.
func NewContainerGroup(ctx *pulumi.Context,
	name string, args *ContainerGroupArgs, opts ...pulumi.ResourceOption) (*ContainerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerGroupName'")
	}
	if args.Containers == nil {
		return nil, errors.New("invalid value for required argument 'Containers'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	var resource ContainerGroup
	err := ctx.RegisterResource("alicloud:eci/containerGroup:ContainerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerGroup gets an existing ContainerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerGroupState, opts ...pulumi.ResourceOption) (*ContainerGroup, error) {
	var resource ContainerGroup
	err := ctx.ReadResource("alicloud:eci/containerGroup:ContainerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerGroup resources.
type containerGroupState struct {
	// The name of the container group.
	ContainerGroupName *string `pulumi:"containerGroupName"`
	// The list of containers.
	Containers []ContainerGroupContainer `pulumi:"containers"`
	// The amount of CPU resources allocated to the container.
	Cpu *float64 `pulumi:"cpu"`
	// The structure of dnsConfig.
	DnsConfig          *ContainerGroupDnsConfig          `pulumi:"dnsConfig"`
	EciSecurityContext *ContainerGroupEciSecurityContext `pulumi:"eciSecurityContext"`
	// HostAliases.
	HostAliases []ContainerGroupHostAlias `pulumi:"hostAliases"`
	// The image registry credential. The details see Block `imageRegistryCredential`.
	ImageRegistryCredentials []ContainerGroupImageRegistryCredential `pulumi:"imageRegistryCredentials"`
	// The list of initContainers.
	InitContainers []ContainerGroupInitContainer `pulumi:"initContainers"`
	// The type of the ECS instance.
	InstanceType *string `pulumi:"instanceType"`
	// The amount of memory resources allocated to the container.
	Memory *float64 `pulumi:"memory"`
	// The RAM role that the container group assumes. ECI and ECS share the same RAM role.
	RamRoleName *string `pulumi:"ramRoleName"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The restart policy of the container group. Default to `Always`.
	RestartPolicy *string `pulumi:"restartPolicy"`
	// The ID of the security group to which the container group belongs. Container groups within the same security group can access each other.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The status of container group.
	Status *string                `pulumi:"status"`
	Tags   map[string]interface{} `pulumi:"tags"`
	// The list of volumes.
	Volumes []ContainerGroupVolume `pulumi:"volumes"`
	// The ID of the VSwitch. Currently, container groups can only be deployed in VPC networks. The number of IP addresses in the VSwitch CIDR block determines the maximum number of container groups that can be created in the VSwitch. Before you can create an ECI instance, plan the CIDR block of the VSwitch.
	VswitchId *string `pulumi:"vswitchId"`
	// The ID of the zone where you want to deploy the container group. If no value is specified, the system assigns a zone to the container group. By default, no value is specified.
	ZoneId *string `pulumi:"zoneId"`
}

type ContainerGroupState struct {
	// The name of the container group.
	ContainerGroupName pulumi.StringPtrInput
	// The list of containers.
	Containers ContainerGroupContainerArrayInput
	// The amount of CPU resources allocated to the container.
	Cpu pulumi.Float64PtrInput
	// The structure of dnsConfig.
	DnsConfig          ContainerGroupDnsConfigPtrInput
	EciSecurityContext ContainerGroupEciSecurityContextPtrInput
	// HostAliases.
	HostAliases ContainerGroupHostAliasArrayInput
	// The image registry credential. The details see Block `imageRegistryCredential`.
	ImageRegistryCredentials ContainerGroupImageRegistryCredentialArrayInput
	// The list of initContainers.
	InitContainers ContainerGroupInitContainerArrayInput
	// The type of the ECS instance.
	InstanceType pulumi.StringPtrInput
	// The amount of memory resources allocated to the container.
	Memory pulumi.Float64PtrInput
	// The RAM role that the container group assumes. ECI and ECS share the same RAM role.
	RamRoleName pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The restart policy of the container group. Default to `Always`.
	RestartPolicy pulumi.StringPtrInput
	// The ID of the security group to which the container group belongs. Container groups within the same security group can access each other.
	SecurityGroupId pulumi.StringPtrInput
	// The status of container group.
	Status pulumi.StringPtrInput
	Tags   pulumi.MapInput
	// The list of volumes.
	Volumes ContainerGroupVolumeArrayInput
	// The ID of the VSwitch. Currently, container groups can only be deployed in VPC networks. The number of IP addresses in the VSwitch CIDR block determines the maximum number of container groups that can be created in the VSwitch. Before you can create an ECI instance, plan the CIDR block of the VSwitch.
	VswitchId pulumi.StringPtrInput
	// The ID of the zone where you want to deploy the container group. If no value is specified, the system assigns a zone to the container group. By default, no value is specified.
	ZoneId pulumi.StringPtrInput
}

func (ContainerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerGroupState)(nil)).Elem()
}

type containerGroupArgs struct {
	// The name of the container group.
	ContainerGroupName string `pulumi:"containerGroupName"`
	// The list of containers.
	Containers []ContainerGroupContainer `pulumi:"containers"`
	// The amount of CPU resources allocated to the container.
	Cpu *float64 `pulumi:"cpu"`
	// The structure of dnsConfig.
	DnsConfig          *ContainerGroupDnsConfig          `pulumi:"dnsConfig"`
	EciSecurityContext *ContainerGroupEciSecurityContext `pulumi:"eciSecurityContext"`
	// HostAliases.
	HostAliases []ContainerGroupHostAlias `pulumi:"hostAliases"`
	// The image registry credential. The details see Block `imageRegistryCredential`.
	ImageRegistryCredentials []ContainerGroupImageRegistryCredential `pulumi:"imageRegistryCredentials"`
	// The list of initContainers.
	InitContainers []ContainerGroupInitContainer `pulumi:"initContainers"`
	// The type of the ECS instance.
	InstanceType *string `pulumi:"instanceType"`
	// The amount of memory resources allocated to the container.
	Memory *float64 `pulumi:"memory"`
	// The RAM role that the container group assumes. ECI and ECS share the same RAM role.
	RamRoleName *string `pulumi:"ramRoleName"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The restart policy of the container group. Default to `Always`.
	RestartPolicy *string `pulumi:"restartPolicy"`
	// The ID of the security group to which the container group belongs. Container groups within the same security group can access each other.
	SecurityGroupId string                 `pulumi:"securityGroupId"`
	Tags            map[string]interface{} `pulumi:"tags"`
	// The list of volumes.
	Volumes []ContainerGroupVolume `pulumi:"volumes"`
	// The ID of the VSwitch. Currently, container groups can only be deployed in VPC networks. The number of IP addresses in the VSwitch CIDR block determines the maximum number of container groups that can be created in the VSwitch. Before you can create an ECI instance, plan the CIDR block of the VSwitch.
	VswitchId string `pulumi:"vswitchId"`
	// The ID of the zone where you want to deploy the container group. If no value is specified, the system assigns a zone to the container group. By default, no value is specified.
	ZoneId *string `pulumi:"zoneId"`
}

// The set of arguments for constructing a ContainerGroup resource.
type ContainerGroupArgs struct {
	// The name of the container group.
	ContainerGroupName pulumi.StringInput
	// The list of containers.
	Containers ContainerGroupContainerArrayInput
	// The amount of CPU resources allocated to the container.
	Cpu pulumi.Float64PtrInput
	// The structure of dnsConfig.
	DnsConfig          ContainerGroupDnsConfigPtrInput
	EciSecurityContext ContainerGroupEciSecurityContextPtrInput
	// HostAliases.
	HostAliases ContainerGroupHostAliasArrayInput
	// The image registry credential. The details see Block `imageRegistryCredential`.
	ImageRegistryCredentials ContainerGroupImageRegistryCredentialArrayInput
	// The list of initContainers.
	InitContainers ContainerGroupInitContainerArrayInput
	// The type of the ECS instance.
	InstanceType pulumi.StringPtrInput
	// The amount of memory resources allocated to the container.
	Memory pulumi.Float64PtrInput
	// The RAM role that the container group assumes. ECI and ECS share the same RAM role.
	RamRoleName pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The restart policy of the container group. Default to `Always`.
	RestartPolicy pulumi.StringPtrInput
	// The ID of the security group to which the container group belongs. Container groups within the same security group can access each other.
	SecurityGroupId pulumi.StringInput
	Tags            pulumi.MapInput
	// The list of volumes.
	Volumes ContainerGroupVolumeArrayInput
	// The ID of the VSwitch. Currently, container groups can only be deployed in VPC networks. The number of IP addresses in the VSwitch CIDR block determines the maximum number of container groups that can be created in the VSwitch. Before you can create an ECI instance, plan the CIDR block of the VSwitch.
	VswitchId pulumi.StringInput
	// The ID of the zone where you want to deploy the container group. If no value is specified, the system assigns a zone to the container group. By default, no value is specified.
	ZoneId pulumi.StringPtrInput
}

func (ContainerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerGroupArgs)(nil)).Elem()
}

type ContainerGroupInput interface {
	pulumi.Input

	ToContainerGroupOutput() ContainerGroupOutput
	ToContainerGroupOutputWithContext(ctx context.Context) ContainerGroupOutput
}

func (*ContainerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroup)(nil))
}

func (i *ContainerGroup) ToContainerGroupOutput() ContainerGroupOutput {
	return i.ToContainerGroupOutputWithContext(context.Background())
}

func (i *ContainerGroup) ToContainerGroupOutputWithContext(ctx context.Context) ContainerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupOutput)
}

func (i *ContainerGroup) ToContainerGroupPtrOutput() ContainerGroupPtrOutput {
	return i.ToContainerGroupPtrOutputWithContext(context.Background())
}

func (i *ContainerGroup) ToContainerGroupPtrOutputWithContext(ctx context.Context) ContainerGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupPtrOutput)
}

type ContainerGroupPtrInput interface {
	pulumi.Input

	ToContainerGroupPtrOutput() ContainerGroupPtrOutput
	ToContainerGroupPtrOutputWithContext(ctx context.Context) ContainerGroupPtrOutput
}

type containerGroupPtrType ContainerGroupArgs

func (*containerGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroup)(nil))
}

func (i *containerGroupPtrType) ToContainerGroupPtrOutput() ContainerGroupPtrOutput {
	return i.ToContainerGroupPtrOutputWithContext(context.Background())
}

func (i *containerGroupPtrType) ToContainerGroupPtrOutputWithContext(ctx context.Context) ContainerGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupPtrOutput)
}

// ContainerGroupArrayInput is an input type that accepts ContainerGroupArray and ContainerGroupArrayOutput values.
// You can construct a concrete instance of `ContainerGroupArrayInput` via:
//
//          ContainerGroupArray{ ContainerGroupArgs{...} }
type ContainerGroupArrayInput interface {
	pulumi.Input

	ToContainerGroupArrayOutput() ContainerGroupArrayOutput
	ToContainerGroupArrayOutputWithContext(context.Context) ContainerGroupArrayOutput
}

type ContainerGroupArray []ContainerGroupInput

func (ContainerGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerGroup)(nil)).Elem()
}

func (i ContainerGroupArray) ToContainerGroupArrayOutput() ContainerGroupArrayOutput {
	return i.ToContainerGroupArrayOutputWithContext(context.Background())
}

func (i ContainerGroupArray) ToContainerGroupArrayOutputWithContext(ctx context.Context) ContainerGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupArrayOutput)
}

// ContainerGroupMapInput is an input type that accepts ContainerGroupMap and ContainerGroupMapOutput values.
// You can construct a concrete instance of `ContainerGroupMapInput` via:
//
//          ContainerGroupMap{ "key": ContainerGroupArgs{...} }
type ContainerGroupMapInput interface {
	pulumi.Input

	ToContainerGroupMapOutput() ContainerGroupMapOutput
	ToContainerGroupMapOutputWithContext(context.Context) ContainerGroupMapOutput
}

type ContainerGroupMap map[string]ContainerGroupInput

func (ContainerGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerGroup)(nil)).Elem()
}

func (i ContainerGroupMap) ToContainerGroupMapOutput() ContainerGroupMapOutput {
	return i.ToContainerGroupMapOutputWithContext(context.Background())
}

func (i ContainerGroupMap) ToContainerGroupMapOutputWithContext(ctx context.Context) ContainerGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupMapOutput)
}

type ContainerGroupOutput struct{ *pulumi.OutputState }

func (ContainerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroup)(nil))
}

func (o ContainerGroupOutput) ToContainerGroupOutput() ContainerGroupOutput {
	return o
}

func (o ContainerGroupOutput) ToContainerGroupOutputWithContext(ctx context.Context) ContainerGroupOutput {
	return o
}

func (o ContainerGroupOutput) ToContainerGroupPtrOutput() ContainerGroupPtrOutput {
	return o.ToContainerGroupPtrOutputWithContext(context.Background())
}

func (o ContainerGroupOutput) ToContainerGroupPtrOutputWithContext(ctx context.Context) ContainerGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerGroup) *ContainerGroup {
		return &v
	}).(ContainerGroupPtrOutput)
}

type ContainerGroupPtrOutput struct{ *pulumi.OutputState }

func (ContainerGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroup)(nil))
}

func (o ContainerGroupPtrOutput) ToContainerGroupPtrOutput() ContainerGroupPtrOutput {
	return o
}

func (o ContainerGroupPtrOutput) ToContainerGroupPtrOutputWithContext(ctx context.Context) ContainerGroupPtrOutput {
	return o
}

func (o ContainerGroupPtrOutput) Elem() ContainerGroupOutput {
	return o.ApplyT(func(v *ContainerGroup) ContainerGroup {
		if v != nil {
			return *v
		}
		var ret ContainerGroup
		return ret
	}).(ContainerGroupOutput)
}

type ContainerGroupArrayOutput struct{ *pulumi.OutputState }

func (ContainerGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerGroup)(nil))
}

func (o ContainerGroupArrayOutput) ToContainerGroupArrayOutput() ContainerGroupArrayOutput {
	return o
}

func (o ContainerGroupArrayOutput) ToContainerGroupArrayOutputWithContext(ctx context.Context) ContainerGroupArrayOutput {
	return o
}

func (o ContainerGroupArrayOutput) Index(i pulumi.IntInput) ContainerGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerGroup {
		return vs[0].([]ContainerGroup)[vs[1].(int)]
	}).(ContainerGroupOutput)
}

type ContainerGroupMapOutput struct{ *pulumi.OutputState }

func (ContainerGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ContainerGroup)(nil))
}

func (o ContainerGroupMapOutput) ToContainerGroupMapOutput() ContainerGroupMapOutput {
	return o
}

func (o ContainerGroupMapOutput) ToContainerGroupMapOutputWithContext(ctx context.Context) ContainerGroupMapOutput {
	return o
}

func (o ContainerGroupMapOutput) MapIndex(k pulumi.StringInput) ContainerGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ContainerGroup {
		return vs[0].(map[string]ContainerGroup)[vs[1].(string)]
	}).(ContainerGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupInput)(nil)).Elem(), &ContainerGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupPtrInput)(nil)).Elem(), &ContainerGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupArrayInput)(nil)).Elem(), ContainerGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupMapInput)(nil)).Elem(), ContainerGroupMap{})
	pulumi.RegisterOutputType(ContainerGroupOutput{})
	pulumi.RegisterOutputType(ContainerGroupPtrOutput{})
	pulumi.RegisterOutputType(ContainerGroupArrayOutput{})
	pulumi.RegisterOutputType(ContainerGroupMapOutput{})
}
