// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Serverless Kubernetes cluster can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:cs/serverlessKubernetes:ServerlessKubernetes main ce4273f9156874b46bb
// ```
type ServerlessKubernetes struct {
	pulumi.CustomResourceState

	Addons ServerlessKubernetesAddonArrayOutput `pulumi:"addons"`
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert pulumi.StringPtrOutput `pulumi:"clientCert"`
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey pulumi.StringPtrOutput `pulumi:"clientKey"`
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert pulumi.StringPtrOutput `pulumi:"clusterCaCert"`
	// Whether enable the deletion protection or not.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection pulumi.BoolPtrOutput `pulumi:"deletionProtection"`
	// Whether to create internet  eip for API Server. Default to false.
	EndpointPublicAccessEnabled pulumi.BoolPtrOutput `pulumi:"endpointPublicAccessEnabled"`
	// Default false, when you want to change `vpcId` and `vswitchId`, you have to set this field to true, then the cluster will be recreated.
	ForceUpdate pulumi.BoolPtrOutput `pulumi:"forceUpdate"`
	// The path of kube config, like `~/.kube/config`.
	KubeConfig pulumi.StringPtrOutput `pulumi:"kubeConfig"`
	// The cluster api server load balance instance specification, default `slb.s1.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
	LoadBalancerSpec pulumi.StringPtrOutput `pulumi:"loadBalancerSpec"`
	// The kubernetes cluster's name. It is the only in one Alicloud account.
	Name       pulumi.StringOutput    `pulumi:"name"`
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true.
	NewNatGateway pulumi.BoolPtrOutput `pulumi:"newNatGateway"`
	// Enable Privatezone if you need to use the service discovery feature within the serverless cluster. Default to false.
	PrivateZone pulumi.BoolPtrOutput `pulumi:"privateZone"`
	// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// Default nil, A map of tags assigned to the kubernetes cluster and work nodes.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used.
	Version pulumi.StringOutput `pulumi:"version"`
	// The vpc where new kubernetes cluster will be located. Specify one vpc's id, if it is not specified, a new VPC  will be built.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// (Required, ForceNew) The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availabilityZone` specified.
	//
	// Deprecated: Field 'vswitch_id' has been deprecated from provider version 1.91.0. New field 'vswitch_ids' replace it.
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
	// The vswitches where new kubernetes cluster will be located.
	VswitchIds pulumi.StringArrayOutput `pulumi:"vswitchIds"`
}

// NewServerlessKubernetes registers a new resource with the given unique name, arguments, and options.
func NewServerlessKubernetes(ctx *pulumi.Context,
	name string, args *ServerlessKubernetesArgs, opts ...pulumi.ResourceOption) (*ServerlessKubernetes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource ServerlessKubernetes
	err := ctx.RegisterResource("alicloud:cs/serverlessKubernetes:ServerlessKubernetes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerlessKubernetes gets an existing ServerlessKubernetes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerlessKubernetes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerlessKubernetesState, opts ...pulumi.ResourceOption) (*ServerlessKubernetes, error) {
	var resource ServerlessKubernetes
	err := ctx.ReadResource("alicloud:cs/serverlessKubernetes:ServerlessKubernetes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerlessKubernetes resources.
type serverlessKubernetesState struct {
	Addons []ServerlessKubernetesAddon `pulumi:"addons"`
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert *string `pulumi:"clientCert"`
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey *string `pulumi:"clientKey"`
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert *string `pulumi:"clusterCaCert"`
	// Whether enable the deletion protection or not.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Whether to create internet  eip for API Server. Default to false.
	EndpointPublicAccessEnabled *bool `pulumi:"endpointPublicAccessEnabled"`
	// Default false, when you want to change `vpcId` and `vswitchId`, you have to set this field to true, then the cluster will be recreated.
	ForceUpdate *bool `pulumi:"forceUpdate"`
	// The path of kube config, like `~/.kube/config`.
	KubeConfig *string `pulumi:"kubeConfig"`
	// The cluster api server load balance instance specification, default `slb.s1.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
	LoadBalancerSpec *string `pulumi:"loadBalancerSpec"`
	// The kubernetes cluster's name. It is the only in one Alicloud account.
	Name       *string `pulumi:"name"`
	NamePrefix *string `pulumi:"namePrefix"`
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true.
	NewNatGateway *bool `pulumi:"newNatGateway"`
	// Enable Privatezone if you need to use the service discovery feature within the serverless cluster. Default to false.
	PrivateZone *bool `pulumi:"privateZone"`
	// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// Default nil, A map of tags assigned to the kubernetes cluster and work nodes.
	Tags map[string]interface{} `pulumi:"tags"`
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used.
	Version *string `pulumi:"version"`
	// The vpc where new kubernetes cluster will be located. Specify one vpc's id, if it is not specified, a new VPC  will be built.
	VpcId *string `pulumi:"vpcId"`
	// (Required, ForceNew) The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availabilityZone` specified.
	//
	// Deprecated: Field 'vswitch_id' has been deprecated from provider version 1.91.0. New field 'vswitch_ids' replace it.
	VswitchId *string `pulumi:"vswitchId"`
	// The vswitches where new kubernetes cluster will be located.
	VswitchIds []string `pulumi:"vswitchIds"`
}

type ServerlessKubernetesState struct {
	Addons ServerlessKubernetesAddonArrayInput
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert pulumi.StringPtrInput
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey pulumi.StringPtrInput
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert pulumi.StringPtrInput
	// Whether enable the deletion protection or not.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection pulumi.BoolPtrInput
	// Whether to create internet  eip for API Server. Default to false.
	EndpointPublicAccessEnabled pulumi.BoolPtrInput
	// Default false, when you want to change `vpcId` and `vswitchId`, you have to set this field to true, then the cluster will be recreated.
	ForceUpdate pulumi.BoolPtrInput
	// The path of kube config, like `~/.kube/config`.
	KubeConfig pulumi.StringPtrInput
	// The cluster api server load balance instance specification, default `slb.s1.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
	LoadBalancerSpec pulumi.StringPtrInput
	// The kubernetes cluster's name. It is the only in one Alicloud account.
	Name       pulumi.StringPtrInput
	NamePrefix pulumi.StringPtrInput
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true.
	NewNatGateway pulumi.BoolPtrInput
	// Enable Privatezone if you need to use the service discovery feature within the serverless cluster. Default to false.
	PrivateZone pulumi.BoolPtrInput
	// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
	SecurityGroupId pulumi.StringPtrInput
	// Default nil, A map of tags assigned to the kubernetes cluster and work nodes.
	Tags pulumi.MapInput
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used.
	Version pulumi.StringPtrInput
	// The vpc where new kubernetes cluster will be located. Specify one vpc's id, if it is not specified, a new VPC  will be built.
	VpcId pulumi.StringPtrInput
	// (Required, ForceNew) The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availabilityZone` specified.
	//
	// Deprecated: Field 'vswitch_id' has been deprecated from provider version 1.91.0. New field 'vswitch_ids' replace it.
	VswitchId pulumi.StringPtrInput
	// The vswitches where new kubernetes cluster will be located.
	VswitchIds pulumi.StringArrayInput
}

func (ServerlessKubernetesState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverlessKubernetesState)(nil)).Elem()
}

type serverlessKubernetesArgs struct {
	Addons []ServerlessKubernetesAddon `pulumi:"addons"`
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert *string `pulumi:"clientCert"`
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey *string `pulumi:"clientKey"`
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert *string `pulumi:"clusterCaCert"`
	// Whether enable the deletion protection or not.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Whether to create internet  eip for API Server. Default to false.
	EndpointPublicAccessEnabled *bool `pulumi:"endpointPublicAccessEnabled"`
	// Default false, when you want to change `vpcId` and `vswitchId`, you have to set this field to true, then the cluster will be recreated.
	ForceUpdate *bool `pulumi:"forceUpdate"`
	// The path of kube config, like `~/.kube/config`.
	KubeConfig *string `pulumi:"kubeConfig"`
	// The cluster api server load balance instance specification, default `slb.s1.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
	LoadBalancerSpec *string `pulumi:"loadBalancerSpec"`
	// The kubernetes cluster's name. It is the only in one Alicloud account.
	Name       *string `pulumi:"name"`
	NamePrefix *string `pulumi:"namePrefix"`
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true.
	NewNatGateway *bool `pulumi:"newNatGateway"`
	// Enable Privatezone if you need to use the service discovery feature within the serverless cluster. Default to false.
	PrivateZone *bool `pulumi:"privateZone"`
	// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// Default nil, A map of tags assigned to the kubernetes cluster and work nodes.
	Tags map[string]interface{} `pulumi:"tags"`
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used.
	Version *string `pulumi:"version"`
	// The vpc where new kubernetes cluster will be located. Specify one vpc's id, if it is not specified, a new VPC  will be built.
	VpcId string `pulumi:"vpcId"`
	// (Required, ForceNew) The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availabilityZone` specified.
	//
	// Deprecated: Field 'vswitch_id' has been deprecated from provider version 1.91.0. New field 'vswitch_ids' replace it.
	VswitchId *string `pulumi:"vswitchId"`
	// The vswitches where new kubernetes cluster will be located.
	VswitchIds []string `pulumi:"vswitchIds"`
}

// The set of arguments for constructing a ServerlessKubernetes resource.
type ServerlessKubernetesArgs struct {
	Addons ServerlessKubernetesAddonArrayInput
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert pulumi.StringPtrInput
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey pulumi.StringPtrInput
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert pulumi.StringPtrInput
	// Whether enable the deletion protection or not.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection pulumi.BoolPtrInput
	// Whether to create internet  eip for API Server. Default to false.
	EndpointPublicAccessEnabled pulumi.BoolPtrInput
	// Default false, when you want to change `vpcId` and `vswitchId`, you have to set this field to true, then the cluster will be recreated.
	ForceUpdate pulumi.BoolPtrInput
	// The path of kube config, like `~/.kube/config`.
	KubeConfig pulumi.StringPtrInput
	// The cluster api server load balance instance specification, default `slb.s1.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
	LoadBalancerSpec pulumi.StringPtrInput
	// The kubernetes cluster's name. It is the only in one Alicloud account.
	Name       pulumi.StringPtrInput
	NamePrefix pulumi.StringPtrInput
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true.
	NewNatGateway pulumi.BoolPtrInput
	// Enable Privatezone if you need to use the service discovery feature within the serverless cluster. Default to false.
	PrivateZone pulumi.BoolPtrInput
	// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
	SecurityGroupId pulumi.StringPtrInput
	// Default nil, A map of tags assigned to the kubernetes cluster and work nodes.
	Tags pulumi.MapInput
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used.
	Version pulumi.StringPtrInput
	// The vpc where new kubernetes cluster will be located. Specify one vpc's id, if it is not specified, a new VPC  will be built.
	VpcId pulumi.StringInput
	// (Required, ForceNew) The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availabilityZone` specified.
	//
	// Deprecated: Field 'vswitch_id' has been deprecated from provider version 1.91.0. New field 'vswitch_ids' replace it.
	VswitchId pulumi.StringPtrInput
	// The vswitches where new kubernetes cluster will be located.
	VswitchIds pulumi.StringArrayInput
}

func (ServerlessKubernetesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverlessKubernetesArgs)(nil)).Elem()
}

type ServerlessKubernetesInput interface {
	pulumi.Input

	ToServerlessKubernetesOutput() ServerlessKubernetesOutput
	ToServerlessKubernetesOutputWithContext(ctx context.Context) ServerlessKubernetesOutput
}

func (*ServerlessKubernetes) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerlessKubernetes)(nil))
}

func (i *ServerlessKubernetes) ToServerlessKubernetesOutput() ServerlessKubernetesOutput {
	return i.ToServerlessKubernetesOutputWithContext(context.Background())
}

func (i *ServerlessKubernetes) ToServerlessKubernetesOutputWithContext(ctx context.Context) ServerlessKubernetesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessKubernetesOutput)
}

func (i *ServerlessKubernetes) ToServerlessKubernetesPtrOutput() ServerlessKubernetesPtrOutput {
	return i.ToServerlessKubernetesPtrOutputWithContext(context.Background())
}

func (i *ServerlessKubernetes) ToServerlessKubernetesPtrOutputWithContext(ctx context.Context) ServerlessKubernetesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessKubernetesPtrOutput)
}

type ServerlessKubernetesPtrInput interface {
	pulumi.Input

	ToServerlessKubernetesPtrOutput() ServerlessKubernetesPtrOutput
	ToServerlessKubernetesPtrOutputWithContext(ctx context.Context) ServerlessKubernetesPtrOutput
}

type serverlessKubernetesPtrType ServerlessKubernetesArgs

func (*serverlessKubernetesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerlessKubernetes)(nil))
}

func (i *serverlessKubernetesPtrType) ToServerlessKubernetesPtrOutput() ServerlessKubernetesPtrOutput {
	return i.ToServerlessKubernetesPtrOutputWithContext(context.Background())
}

func (i *serverlessKubernetesPtrType) ToServerlessKubernetesPtrOutputWithContext(ctx context.Context) ServerlessKubernetesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessKubernetesPtrOutput)
}

// ServerlessKubernetesArrayInput is an input type that accepts ServerlessKubernetesArray and ServerlessKubernetesArrayOutput values.
// You can construct a concrete instance of `ServerlessKubernetesArrayInput` via:
//
//          ServerlessKubernetesArray{ ServerlessKubernetesArgs{...} }
type ServerlessKubernetesArrayInput interface {
	pulumi.Input

	ToServerlessKubernetesArrayOutput() ServerlessKubernetesArrayOutput
	ToServerlessKubernetesArrayOutputWithContext(context.Context) ServerlessKubernetesArrayOutput
}

type ServerlessKubernetesArray []ServerlessKubernetesInput

func (ServerlessKubernetesArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ServerlessKubernetes)(nil))
}

func (i ServerlessKubernetesArray) ToServerlessKubernetesArrayOutput() ServerlessKubernetesArrayOutput {
	return i.ToServerlessKubernetesArrayOutputWithContext(context.Background())
}

func (i ServerlessKubernetesArray) ToServerlessKubernetesArrayOutputWithContext(ctx context.Context) ServerlessKubernetesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessKubernetesArrayOutput)
}

// ServerlessKubernetesMapInput is an input type that accepts ServerlessKubernetesMap and ServerlessKubernetesMapOutput values.
// You can construct a concrete instance of `ServerlessKubernetesMapInput` via:
//
//          ServerlessKubernetesMap{ "key": ServerlessKubernetesArgs{...} }
type ServerlessKubernetesMapInput interface {
	pulumi.Input

	ToServerlessKubernetesMapOutput() ServerlessKubernetesMapOutput
	ToServerlessKubernetesMapOutputWithContext(context.Context) ServerlessKubernetesMapOutput
}

type ServerlessKubernetesMap map[string]ServerlessKubernetesInput

func (ServerlessKubernetesMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ServerlessKubernetes)(nil))
}

func (i ServerlessKubernetesMap) ToServerlessKubernetesMapOutput() ServerlessKubernetesMapOutput {
	return i.ToServerlessKubernetesMapOutputWithContext(context.Background())
}

func (i ServerlessKubernetesMap) ToServerlessKubernetesMapOutputWithContext(ctx context.Context) ServerlessKubernetesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessKubernetesMapOutput)
}

type ServerlessKubernetesOutput struct {
	*pulumi.OutputState
}

func (ServerlessKubernetesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerlessKubernetes)(nil))
}

func (o ServerlessKubernetesOutput) ToServerlessKubernetesOutput() ServerlessKubernetesOutput {
	return o
}

func (o ServerlessKubernetesOutput) ToServerlessKubernetesOutputWithContext(ctx context.Context) ServerlessKubernetesOutput {
	return o
}

func (o ServerlessKubernetesOutput) ToServerlessKubernetesPtrOutput() ServerlessKubernetesPtrOutput {
	return o.ToServerlessKubernetesPtrOutputWithContext(context.Background())
}

func (o ServerlessKubernetesOutput) ToServerlessKubernetesPtrOutputWithContext(ctx context.Context) ServerlessKubernetesPtrOutput {
	return o.ApplyT(func(v ServerlessKubernetes) *ServerlessKubernetes {
		return &v
	}).(ServerlessKubernetesPtrOutput)
}

type ServerlessKubernetesPtrOutput struct {
	*pulumi.OutputState
}

func (ServerlessKubernetesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerlessKubernetes)(nil))
}

func (o ServerlessKubernetesPtrOutput) ToServerlessKubernetesPtrOutput() ServerlessKubernetesPtrOutput {
	return o
}

func (o ServerlessKubernetesPtrOutput) ToServerlessKubernetesPtrOutputWithContext(ctx context.Context) ServerlessKubernetesPtrOutput {
	return o
}

type ServerlessKubernetesArrayOutput struct{ *pulumi.OutputState }

func (ServerlessKubernetesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerlessKubernetes)(nil))
}

func (o ServerlessKubernetesArrayOutput) ToServerlessKubernetesArrayOutput() ServerlessKubernetesArrayOutput {
	return o
}

func (o ServerlessKubernetesArrayOutput) ToServerlessKubernetesArrayOutputWithContext(ctx context.Context) ServerlessKubernetesArrayOutput {
	return o
}

func (o ServerlessKubernetesArrayOutput) Index(i pulumi.IntInput) ServerlessKubernetesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerlessKubernetes {
		return vs[0].([]ServerlessKubernetes)[vs[1].(int)]
	}).(ServerlessKubernetesOutput)
}

type ServerlessKubernetesMapOutput struct{ *pulumi.OutputState }

func (ServerlessKubernetesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServerlessKubernetes)(nil))
}

func (o ServerlessKubernetesMapOutput) ToServerlessKubernetesMapOutput() ServerlessKubernetesMapOutput {
	return o
}

func (o ServerlessKubernetesMapOutput) ToServerlessKubernetesMapOutputWithContext(ctx context.Context) ServerlessKubernetesMapOutput {
	return o
}

func (o ServerlessKubernetesMapOutput) MapIndex(k pulumi.StringInput) ServerlessKubernetesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServerlessKubernetes {
		return vs[0].(map[string]ServerlessKubernetes)[vs[1].(string)]
	}).(ServerlessKubernetesOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerlessKubernetesOutput{})
	pulumi.RegisterOutputType(ServerlessKubernetesPtrOutput{})
	pulumi.RegisterOutputType(ServerlessKubernetesArrayOutput{})
	pulumi.RegisterOutputType(ServerlessKubernetesMapOutput{})
}
