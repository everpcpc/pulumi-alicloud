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
// Kubernetes cluster can be imported using the id, e.g. Then complete the main.tf accords to the result of `terraform plan`
//
// ```sh
//  $ pulumi import alicloud:cs/edgeKubernetes:EdgeKubernetes alicloud_cs_edge_kubernetes.main cluster-id
// ```
type EdgeKubernetes struct {
	pulumi.CustomResourceState

	Addons EdgeKubernetesAddonArrayOutput `pulumi:"addons"`
	// The ID of availability zone.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// (Available in 1.105.0+) Nested attribute containing certificate authority data for your cluster.
	CertificateAuthority EdgeKubernetesCertificateAuthorityOutput `pulumi:"certificateAuthority"`
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert pulumi.StringPtrOutput `pulumi:"clientCert"`
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey pulumi.StringPtrOutput `pulumi:"clientKey"`
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert pulumi.StringPtrOutput          `pulumi:"clusterCaCert"`
	Connections   EdgeKubernetesConnectionsOutput `pulumi:"connections"`
	// Whether to enable cluster deletion protection.
	DeletionProtection pulumi.BoolPtrOutput `pulumi:"deletionProtection"`
	// Default false, when you want to change `vpcId`, you have to set this field to true, then the cluster will be recreated.
	ForceUpdate pulumi.BoolPtrOutput `pulumi:"forceUpdate"`
	// Install cloud monitor agent on ECS. default: `true`.
	InstallCloudMonitor pulumi.BoolPtrOutput `pulumi:"installCloudMonitor"`
	// Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
	IsEnterpriseSecurityGroup pulumi.BoolOutput `pulumi:"isEnterpriseSecurityGroup"`
	// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	KeyName pulumi.StringPtrOutput `pulumi:"keyName"`
	// The path of kube config, like `~/.kube/config`.
	KubeConfig pulumi.StringPtrOutput           `pulumi:"kubeConfig"`
	LogConfig  EdgeKubernetesLogConfigPtrOutput `pulumi:"logConfig"`
	// The kubernetes cluster's name. It is unique in one Alicloud account.
	Name       pulumi.StringOutput    `pulumi:"name"`
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// The ID of nat gateway used to launch kubernetes cluster.
	NatGatewayId pulumi.StringOutput `pulumi:"natGatewayId"`
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
	NewNatGateway pulumi.BoolPtrOutput `pulumi:"newNatGateway"`
	// The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
	NodeCidrMask pulumi.IntPtrOutput `pulumi:"nodeCidrMask"`
	// The password of ssh login cluster node. You have to specify one of `password`, `keyName` `kmsEncryptedPassword` fields.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// [Flannel Specific] The CIDR block for the pod network when using Flannel.
	PodCidr pulumi.StringPtrOutput `pulumi:"podCidr"`
	// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
	ProxyMode    pulumi.StringPtrOutput   `pulumi:"proxyMode"`
	RdsInstances pulumi.StringArrayOutput `pulumi:"rdsInstances"`
	// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
	ResourceGroupId pulumi.StringOutput      `pulumi:"resourceGroupId"`
	RetainResources pulumi.StringArrayOutput `pulumi:"retainResources"`
	// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	ServiceCidr pulumi.StringPtrOutput `pulumi:"serviceCidr"`
	SlbInternet pulumi.StringOutput    `pulumi:"slbInternet"`
	// Whether to create internet load balancer for API Server. Default to true.
	SlbInternetEnabled pulumi.BoolPtrOutput `pulumi:"slbInternetEnabled"`
	// The ID of private load balancer where the current cluster master node is located.
	SlbIntranet pulumi.StringOutput `pulumi:"slbIntranet"`
	// Default nil, A map of tags assigned to the kubernetes cluster and work node.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
	Version pulumi.StringOutput `pulumi:"version"`
	// The ID of VPC where the current cluster is located.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The data disk configurations of worker nodes, such as the disk type and disk size.
	WorkerDataDisks EdgeKubernetesWorkerDataDiskArrayOutput `pulumi:"workerDataDisks"`
	// The system disk category of worker node. Its valid value are `cloudEfficiency`, `cloudSsd` and `cloudEssd` and . Default to `cloudEfficiency`.
	WorkerDiskCategory pulumi.StringPtrOutput `pulumi:"workerDiskCategory"`
	// Worker node system disk performance level, when `workerDiskCategory` values `cloudEssd`, the optional values are `PL0`, `PL1`, `PL2` or `PL3`, but the specific performance level is related to the disk capacity. For more information, see [Enhanced SSDs](https://www.alibabacloud.com/help/doc-detail/122389.htm). Default is `PL1`.
	WorkerDiskPerformanceLevel pulumi.StringPtrOutput `pulumi:"workerDiskPerformanceLevel"`
	// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
	WorkerDiskSize pulumi.IntPtrOutput `pulumi:"workerDiskSize"`
	// Worker node system disk auto snapshot policy.
	WorkerDiskSnapshotPolicyId pulumi.StringPtrOutput `pulumi:"workerDiskSnapshotPolicyId"`
	WorkerInstanceChargeType   pulumi.StringPtrOutput `pulumi:"workerInstanceChargeType"`
	// The instance types of worker node, you can set multiple types to avoid NoStock of a certain type
	WorkerInstanceTypes pulumi.StringArrayOutput `pulumi:"workerInstanceTypes"`
	// List of cluster worker nodes.
	WorkerNodes EdgeKubernetesWorkerNodeArrayOutput `pulumi:"workerNodes"`
	// The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumber     pulumi.IntOutput         `pulumi:"workerNumber"`
	WorkerVswitchIds pulumi.StringArrayOutput `pulumi:"workerVswitchIds"`
}

// NewEdgeKubernetes registers a new resource with the given unique name, arguments, and options.
func NewEdgeKubernetes(ctx *pulumi.Context,
	name string, args *EdgeKubernetesArgs, opts ...pulumi.ResourceOption) (*EdgeKubernetes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.WorkerInstanceTypes == nil {
		return nil, errors.New("invalid value for required argument 'WorkerInstanceTypes'")
	}
	if args.WorkerNumber == nil {
		return nil, errors.New("invalid value for required argument 'WorkerNumber'")
	}
	if args.WorkerVswitchIds == nil {
		return nil, errors.New("invalid value for required argument 'WorkerVswitchIds'")
	}
	var resource EdgeKubernetes
	err := ctx.RegisterResource("alicloud:cs/edgeKubernetes:EdgeKubernetes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEdgeKubernetes gets an existing EdgeKubernetes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEdgeKubernetes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgeKubernetesState, opts ...pulumi.ResourceOption) (*EdgeKubernetes, error) {
	var resource EdgeKubernetes
	err := ctx.ReadResource("alicloud:cs/edgeKubernetes:EdgeKubernetes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EdgeKubernetes resources.
type edgeKubernetesState struct {
	Addons []EdgeKubernetesAddon `pulumi:"addons"`
	// The ID of availability zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// (Available in 1.105.0+) Nested attribute containing certificate authority data for your cluster.
	CertificateAuthority *EdgeKubernetesCertificateAuthority `pulumi:"certificateAuthority"`
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert *string `pulumi:"clientCert"`
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey *string `pulumi:"clientKey"`
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert *string                    `pulumi:"clusterCaCert"`
	Connections   *EdgeKubernetesConnections `pulumi:"connections"`
	// Whether to enable cluster deletion protection.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Default false, when you want to change `vpcId`, you have to set this field to true, then the cluster will be recreated.
	ForceUpdate *bool `pulumi:"forceUpdate"`
	// Install cloud monitor agent on ECS. default: `true`.
	InstallCloudMonitor *bool `pulumi:"installCloudMonitor"`
	// Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
	IsEnterpriseSecurityGroup *bool `pulumi:"isEnterpriseSecurityGroup"`
	// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	KeyName *string `pulumi:"keyName"`
	// The path of kube config, like `~/.kube/config`.
	KubeConfig *string                  `pulumi:"kubeConfig"`
	LogConfig  *EdgeKubernetesLogConfig `pulumi:"logConfig"`
	// The kubernetes cluster's name. It is unique in one Alicloud account.
	Name       *string `pulumi:"name"`
	NamePrefix *string `pulumi:"namePrefix"`
	// The ID of nat gateway used to launch kubernetes cluster.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
	NewNatGateway *bool `pulumi:"newNatGateway"`
	// The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
	NodeCidrMask *int `pulumi:"nodeCidrMask"`
	// The password of ssh login cluster node. You have to specify one of `password`, `keyName` `kmsEncryptedPassword` fields.
	Password *string `pulumi:"password"`
	// [Flannel Specific] The CIDR block for the pod network when using Flannel.
	PodCidr *string `pulumi:"podCidr"`
	// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
	ProxyMode    *string  `pulumi:"proxyMode"`
	RdsInstances []string `pulumi:"rdsInstances"`
	// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
	ResourceGroupId *string  `pulumi:"resourceGroupId"`
	RetainResources []string `pulumi:"retainResources"`
	// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	ServiceCidr *string `pulumi:"serviceCidr"`
	SlbInternet *string `pulumi:"slbInternet"`
	// Whether to create internet load balancer for API Server. Default to true.
	SlbInternetEnabled *bool `pulumi:"slbInternetEnabled"`
	// The ID of private load balancer where the current cluster master node is located.
	SlbIntranet *string `pulumi:"slbIntranet"`
	// Default nil, A map of tags assigned to the kubernetes cluster and work node.
	Tags map[string]interface{} `pulumi:"tags"`
	// Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
	UserData *string `pulumi:"userData"`
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
	Version *string `pulumi:"version"`
	// The ID of VPC where the current cluster is located.
	VpcId *string `pulumi:"vpcId"`
	// The data disk configurations of worker nodes, such as the disk type and disk size.
	WorkerDataDisks []EdgeKubernetesWorkerDataDisk `pulumi:"workerDataDisks"`
	// The system disk category of worker node. Its valid value are `cloudEfficiency`, `cloudSsd` and `cloudEssd` and . Default to `cloudEfficiency`.
	WorkerDiskCategory *string `pulumi:"workerDiskCategory"`
	// Worker node system disk performance level, when `workerDiskCategory` values `cloudEssd`, the optional values are `PL0`, `PL1`, `PL2` or `PL3`, but the specific performance level is related to the disk capacity. For more information, see [Enhanced SSDs](https://www.alibabacloud.com/help/doc-detail/122389.htm). Default is `PL1`.
	WorkerDiskPerformanceLevel *string `pulumi:"workerDiskPerformanceLevel"`
	// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
	WorkerDiskSize *int `pulumi:"workerDiskSize"`
	// Worker node system disk auto snapshot policy.
	WorkerDiskSnapshotPolicyId *string `pulumi:"workerDiskSnapshotPolicyId"`
	WorkerInstanceChargeType   *string `pulumi:"workerInstanceChargeType"`
	// The instance types of worker node, you can set multiple types to avoid NoStock of a certain type
	WorkerInstanceTypes []string `pulumi:"workerInstanceTypes"`
	// List of cluster worker nodes.
	WorkerNodes []EdgeKubernetesWorkerNode `pulumi:"workerNodes"`
	// The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumber     *int     `pulumi:"workerNumber"`
	WorkerVswitchIds []string `pulumi:"workerVswitchIds"`
}

type EdgeKubernetesState struct {
	Addons EdgeKubernetesAddonArrayInput
	// The ID of availability zone.
	AvailabilityZone pulumi.StringPtrInput
	// (Available in 1.105.0+) Nested attribute containing certificate authority data for your cluster.
	CertificateAuthority EdgeKubernetesCertificateAuthorityPtrInput
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert pulumi.StringPtrInput
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey pulumi.StringPtrInput
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert pulumi.StringPtrInput
	Connections   EdgeKubernetesConnectionsPtrInput
	// Whether to enable cluster deletion protection.
	DeletionProtection pulumi.BoolPtrInput
	// Default false, when you want to change `vpcId`, you have to set this field to true, then the cluster will be recreated.
	ForceUpdate pulumi.BoolPtrInput
	// Install cloud monitor agent on ECS. default: `true`.
	InstallCloudMonitor pulumi.BoolPtrInput
	// Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
	IsEnterpriseSecurityGroup pulumi.BoolPtrInput
	// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	KeyName pulumi.StringPtrInput
	// The path of kube config, like `~/.kube/config`.
	KubeConfig pulumi.StringPtrInput
	LogConfig  EdgeKubernetesLogConfigPtrInput
	// The kubernetes cluster's name. It is unique in one Alicloud account.
	Name       pulumi.StringPtrInput
	NamePrefix pulumi.StringPtrInput
	// The ID of nat gateway used to launch kubernetes cluster.
	NatGatewayId pulumi.StringPtrInput
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
	NewNatGateway pulumi.BoolPtrInput
	// The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
	NodeCidrMask pulumi.IntPtrInput
	// The password of ssh login cluster node. You have to specify one of `password`, `keyName` `kmsEncryptedPassword` fields.
	Password pulumi.StringPtrInput
	// [Flannel Specific] The CIDR block for the pod network when using Flannel.
	PodCidr pulumi.StringPtrInput
	// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
	ProxyMode    pulumi.StringPtrInput
	RdsInstances pulumi.StringArrayInput
	// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
	ResourceGroupId pulumi.StringPtrInput
	RetainResources pulumi.StringArrayInput
	// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
	SecurityGroupId pulumi.StringPtrInput
	// The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	ServiceCidr pulumi.StringPtrInput
	SlbInternet pulumi.StringPtrInput
	// Whether to create internet load balancer for API Server. Default to true.
	SlbInternetEnabled pulumi.BoolPtrInput
	// The ID of private load balancer where the current cluster master node is located.
	SlbIntranet pulumi.StringPtrInput
	// Default nil, A map of tags assigned to the kubernetes cluster and work node.
	Tags pulumi.MapInput
	// Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
	UserData pulumi.StringPtrInput
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
	Version pulumi.StringPtrInput
	// The ID of VPC where the current cluster is located.
	VpcId pulumi.StringPtrInput
	// The data disk configurations of worker nodes, such as the disk type and disk size.
	WorkerDataDisks EdgeKubernetesWorkerDataDiskArrayInput
	// The system disk category of worker node. Its valid value are `cloudEfficiency`, `cloudSsd` and `cloudEssd` and . Default to `cloudEfficiency`.
	WorkerDiskCategory pulumi.StringPtrInput
	// Worker node system disk performance level, when `workerDiskCategory` values `cloudEssd`, the optional values are `PL0`, `PL1`, `PL2` or `PL3`, but the specific performance level is related to the disk capacity. For more information, see [Enhanced SSDs](https://www.alibabacloud.com/help/doc-detail/122389.htm). Default is `PL1`.
	WorkerDiskPerformanceLevel pulumi.StringPtrInput
	// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
	WorkerDiskSize pulumi.IntPtrInput
	// Worker node system disk auto snapshot policy.
	WorkerDiskSnapshotPolicyId pulumi.StringPtrInput
	WorkerInstanceChargeType   pulumi.StringPtrInput
	// The instance types of worker node, you can set multiple types to avoid NoStock of a certain type
	WorkerInstanceTypes pulumi.StringArrayInput
	// List of cluster worker nodes.
	WorkerNodes EdgeKubernetesWorkerNodeArrayInput
	// The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumber     pulumi.IntPtrInput
	WorkerVswitchIds pulumi.StringArrayInput
}

func (EdgeKubernetesState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeKubernetesState)(nil)).Elem()
}

type edgeKubernetesArgs struct {
	Addons []EdgeKubernetesAddon `pulumi:"addons"`
	// The ID of availability zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert *string `pulumi:"clientCert"`
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey *string `pulumi:"clientKey"`
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert *string `pulumi:"clusterCaCert"`
	// Whether to enable cluster deletion protection.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Default false, when you want to change `vpcId`, you have to set this field to true, then the cluster will be recreated.
	ForceUpdate *bool `pulumi:"forceUpdate"`
	// Install cloud monitor agent on ECS. default: `true`.
	InstallCloudMonitor *bool `pulumi:"installCloudMonitor"`
	// Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
	IsEnterpriseSecurityGroup *bool `pulumi:"isEnterpriseSecurityGroup"`
	// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	KeyName *string `pulumi:"keyName"`
	// The path of kube config, like `~/.kube/config`.
	KubeConfig *string                  `pulumi:"kubeConfig"`
	LogConfig  *EdgeKubernetesLogConfig `pulumi:"logConfig"`
	// The kubernetes cluster's name. It is unique in one Alicloud account.
	Name       *string `pulumi:"name"`
	NamePrefix *string `pulumi:"namePrefix"`
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
	NewNatGateway *bool `pulumi:"newNatGateway"`
	// The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
	NodeCidrMask *int `pulumi:"nodeCidrMask"`
	// The password of ssh login cluster node. You have to specify one of `password`, `keyName` `kmsEncryptedPassword` fields.
	Password *string `pulumi:"password"`
	// [Flannel Specific] The CIDR block for the pod network when using Flannel.
	PodCidr *string `pulumi:"podCidr"`
	// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
	ProxyMode    *string  `pulumi:"proxyMode"`
	RdsInstances []string `pulumi:"rdsInstances"`
	// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
	ResourceGroupId *string  `pulumi:"resourceGroupId"`
	RetainResources []string `pulumi:"retainResources"`
	// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	ServiceCidr *string `pulumi:"serviceCidr"`
	// Whether to create internet load balancer for API Server. Default to true.
	SlbInternetEnabled *bool `pulumi:"slbInternetEnabled"`
	// Default nil, A map of tags assigned to the kubernetes cluster and work node.
	Tags map[string]interface{} `pulumi:"tags"`
	// Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
	UserData *string `pulumi:"userData"`
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
	Version *string `pulumi:"version"`
	// The data disk configurations of worker nodes, such as the disk type and disk size.
	WorkerDataDisks []EdgeKubernetesWorkerDataDisk `pulumi:"workerDataDisks"`
	// The system disk category of worker node. Its valid value are `cloudEfficiency`, `cloudSsd` and `cloudEssd` and . Default to `cloudEfficiency`.
	WorkerDiskCategory *string `pulumi:"workerDiskCategory"`
	// Worker node system disk performance level, when `workerDiskCategory` values `cloudEssd`, the optional values are `PL0`, `PL1`, `PL2` or `PL3`, but the specific performance level is related to the disk capacity. For more information, see [Enhanced SSDs](https://www.alibabacloud.com/help/doc-detail/122389.htm). Default is `PL1`.
	WorkerDiskPerformanceLevel *string `pulumi:"workerDiskPerformanceLevel"`
	// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
	WorkerDiskSize *int `pulumi:"workerDiskSize"`
	// Worker node system disk auto snapshot policy.
	WorkerDiskSnapshotPolicyId *string `pulumi:"workerDiskSnapshotPolicyId"`
	WorkerInstanceChargeType   *string `pulumi:"workerInstanceChargeType"`
	// The instance types of worker node, you can set multiple types to avoid NoStock of a certain type
	WorkerInstanceTypes []string `pulumi:"workerInstanceTypes"`
	// The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumber     int      `pulumi:"workerNumber"`
	WorkerVswitchIds []string `pulumi:"workerVswitchIds"`
}

// The set of arguments for constructing a EdgeKubernetes resource.
type EdgeKubernetesArgs struct {
	Addons EdgeKubernetesAddonArrayInput
	// The ID of availability zone.
	AvailabilityZone pulumi.StringPtrInput
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert pulumi.StringPtrInput
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey pulumi.StringPtrInput
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert pulumi.StringPtrInput
	// Whether to enable cluster deletion protection.
	DeletionProtection pulumi.BoolPtrInput
	// Default false, when you want to change `vpcId`, you have to set this field to true, then the cluster will be recreated.
	ForceUpdate pulumi.BoolPtrInput
	// Install cloud monitor agent on ECS. default: `true`.
	InstallCloudMonitor pulumi.BoolPtrInput
	// Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
	IsEnterpriseSecurityGroup pulumi.BoolPtrInput
	// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	KeyName pulumi.StringPtrInput
	// The path of kube config, like `~/.kube/config`.
	KubeConfig pulumi.StringPtrInput
	LogConfig  EdgeKubernetesLogConfigPtrInput
	// The kubernetes cluster's name. It is unique in one Alicloud account.
	Name       pulumi.StringPtrInput
	NamePrefix pulumi.StringPtrInput
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
	NewNatGateway pulumi.BoolPtrInput
	// The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
	NodeCidrMask pulumi.IntPtrInput
	// The password of ssh login cluster node. You have to specify one of `password`, `keyName` `kmsEncryptedPassword` fields.
	Password pulumi.StringPtrInput
	// [Flannel Specific] The CIDR block for the pod network when using Flannel.
	PodCidr pulumi.StringPtrInput
	// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
	ProxyMode    pulumi.StringPtrInput
	RdsInstances pulumi.StringArrayInput
	// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
	ResourceGroupId pulumi.StringPtrInput
	RetainResources pulumi.StringArrayInput
	// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
	SecurityGroupId pulumi.StringPtrInput
	// The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	ServiceCidr pulumi.StringPtrInput
	// Whether to create internet load balancer for API Server. Default to true.
	SlbInternetEnabled pulumi.BoolPtrInput
	// Default nil, A map of tags assigned to the kubernetes cluster and work node.
	Tags pulumi.MapInput
	// Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
	UserData pulumi.StringPtrInput
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
	Version pulumi.StringPtrInput
	// The data disk configurations of worker nodes, such as the disk type and disk size.
	WorkerDataDisks EdgeKubernetesWorkerDataDiskArrayInput
	// The system disk category of worker node. Its valid value are `cloudEfficiency`, `cloudSsd` and `cloudEssd` and . Default to `cloudEfficiency`.
	WorkerDiskCategory pulumi.StringPtrInput
	// Worker node system disk performance level, when `workerDiskCategory` values `cloudEssd`, the optional values are `PL0`, `PL1`, `PL2` or `PL3`, but the specific performance level is related to the disk capacity. For more information, see [Enhanced SSDs](https://www.alibabacloud.com/help/doc-detail/122389.htm). Default is `PL1`.
	WorkerDiskPerformanceLevel pulumi.StringPtrInput
	// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
	WorkerDiskSize pulumi.IntPtrInput
	// Worker node system disk auto snapshot policy.
	WorkerDiskSnapshotPolicyId pulumi.StringPtrInput
	WorkerInstanceChargeType   pulumi.StringPtrInput
	// The instance types of worker node, you can set multiple types to avoid NoStock of a certain type
	WorkerInstanceTypes pulumi.StringArrayInput
	// The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumber     pulumi.IntInput
	WorkerVswitchIds pulumi.StringArrayInput
}

func (EdgeKubernetesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeKubernetesArgs)(nil)).Elem()
}

type EdgeKubernetesInput interface {
	pulumi.Input

	ToEdgeKubernetesOutput() EdgeKubernetesOutput
	ToEdgeKubernetesOutputWithContext(ctx context.Context) EdgeKubernetesOutput
}

func (*EdgeKubernetes) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeKubernetes)(nil))
}

func (i *EdgeKubernetes) ToEdgeKubernetesOutput() EdgeKubernetesOutput {
	return i.ToEdgeKubernetesOutputWithContext(context.Background())
}

func (i *EdgeKubernetes) ToEdgeKubernetesOutputWithContext(ctx context.Context) EdgeKubernetesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeKubernetesOutput)
}

func (i *EdgeKubernetes) ToEdgeKubernetesPtrOutput() EdgeKubernetesPtrOutput {
	return i.ToEdgeKubernetesPtrOutputWithContext(context.Background())
}

func (i *EdgeKubernetes) ToEdgeKubernetesPtrOutputWithContext(ctx context.Context) EdgeKubernetesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeKubernetesPtrOutput)
}

type EdgeKubernetesPtrInput interface {
	pulumi.Input

	ToEdgeKubernetesPtrOutput() EdgeKubernetesPtrOutput
	ToEdgeKubernetesPtrOutputWithContext(ctx context.Context) EdgeKubernetesPtrOutput
}

type edgeKubernetesPtrType EdgeKubernetesArgs

func (*edgeKubernetesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeKubernetes)(nil))
}

func (i *edgeKubernetesPtrType) ToEdgeKubernetesPtrOutput() EdgeKubernetesPtrOutput {
	return i.ToEdgeKubernetesPtrOutputWithContext(context.Background())
}

func (i *edgeKubernetesPtrType) ToEdgeKubernetesPtrOutputWithContext(ctx context.Context) EdgeKubernetesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeKubernetesPtrOutput)
}

// EdgeKubernetesArrayInput is an input type that accepts EdgeKubernetesArray and EdgeKubernetesArrayOutput values.
// You can construct a concrete instance of `EdgeKubernetesArrayInput` via:
//
//          EdgeKubernetesArray{ EdgeKubernetesArgs{...} }
type EdgeKubernetesArrayInput interface {
	pulumi.Input

	ToEdgeKubernetesArrayOutput() EdgeKubernetesArrayOutput
	ToEdgeKubernetesArrayOutputWithContext(context.Context) EdgeKubernetesArrayOutput
}

type EdgeKubernetesArray []EdgeKubernetesInput

func (EdgeKubernetesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgeKubernetes)(nil)).Elem()
}

func (i EdgeKubernetesArray) ToEdgeKubernetesArrayOutput() EdgeKubernetesArrayOutput {
	return i.ToEdgeKubernetesArrayOutputWithContext(context.Background())
}

func (i EdgeKubernetesArray) ToEdgeKubernetesArrayOutputWithContext(ctx context.Context) EdgeKubernetesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeKubernetesArrayOutput)
}

// EdgeKubernetesMapInput is an input type that accepts EdgeKubernetesMap and EdgeKubernetesMapOutput values.
// You can construct a concrete instance of `EdgeKubernetesMapInput` via:
//
//          EdgeKubernetesMap{ "key": EdgeKubernetesArgs{...} }
type EdgeKubernetesMapInput interface {
	pulumi.Input

	ToEdgeKubernetesMapOutput() EdgeKubernetesMapOutput
	ToEdgeKubernetesMapOutputWithContext(context.Context) EdgeKubernetesMapOutput
}

type EdgeKubernetesMap map[string]EdgeKubernetesInput

func (EdgeKubernetesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgeKubernetes)(nil)).Elem()
}

func (i EdgeKubernetesMap) ToEdgeKubernetesMapOutput() EdgeKubernetesMapOutput {
	return i.ToEdgeKubernetesMapOutputWithContext(context.Background())
}

func (i EdgeKubernetesMap) ToEdgeKubernetesMapOutputWithContext(ctx context.Context) EdgeKubernetesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeKubernetesMapOutput)
}

type EdgeKubernetesOutput struct{ *pulumi.OutputState }

func (EdgeKubernetesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeKubernetes)(nil))
}

func (o EdgeKubernetesOutput) ToEdgeKubernetesOutput() EdgeKubernetesOutput {
	return o
}

func (o EdgeKubernetesOutput) ToEdgeKubernetesOutputWithContext(ctx context.Context) EdgeKubernetesOutput {
	return o
}

func (o EdgeKubernetesOutput) ToEdgeKubernetesPtrOutput() EdgeKubernetesPtrOutput {
	return o.ToEdgeKubernetesPtrOutputWithContext(context.Background())
}

func (o EdgeKubernetesOutput) ToEdgeKubernetesPtrOutputWithContext(ctx context.Context) EdgeKubernetesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdgeKubernetes) *EdgeKubernetes {
		return &v
	}).(EdgeKubernetesPtrOutput)
}

type EdgeKubernetesPtrOutput struct{ *pulumi.OutputState }

func (EdgeKubernetesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeKubernetes)(nil))
}

func (o EdgeKubernetesPtrOutput) ToEdgeKubernetesPtrOutput() EdgeKubernetesPtrOutput {
	return o
}

func (o EdgeKubernetesPtrOutput) ToEdgeKubernetesPtrOutputWithContext(ctx context.Context) EdgeKubernetesPtrOutput {
	return o
}

func (o EdgeKubernetesPtrOutput) Elem() EdgeKubernetesOutput {
	return o.ApplyT(func(v *EdgeKubernetes) EdgeKubernetes {
		if v != nil {
			return *v
		}
		var ret EdgeKubernetes
		return ret
	}).(EdgeKubernetesOutput)
}

type EdgeKubernetesArrayOutput struct{ *pulumi.OutputState }

func (EdgeKubernetesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdgeKubernetes)(nil))
}

func (o EdgeKubernetesArrayOutput) ToEdgeKubernetesArrayOutput() EdgeKubernetesArrayOutput {
	return o
}

func (o EdgeKubernetesArrayOutput) ToEdgeKubernetesArrayOutputWithContext(ctx context.Context) EdgeKubernetesArrayOutput {
	return o
}

func (o EdgeKubernetesArrayOutput) Index(i pulumi.IntInput) EdgeKubernetesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdgeKubernetes {
		return vs[0].([]EdgeKubernetes)[vs[1].(int)]
	}).(EdgeKubernetesOutput)
}

type EdgeKubernetesMapOutput struct{ *pulumi.OutputState }

func (EdgeKubernetesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EdgeKubernetes)(nil))
}

func (o EdgeKubernetesMapOutput) ToEdgeKubernetesMapOutput() EdgeKubernetesMapOutput {
	return o
}

func (o EdgeKubernetesMapOutput) ToEdgeKubernetesMapOutputWithContext(ctx context.Context) EdgeKubernetesMapOutput {
	return o
}

func (o EdgeKubernetesMapOutput) MapIndex(k pulumi.StringInput) EdgeKubernetesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EdgeKubernetes {
		return vs[0].(map[string]EdgeKubernetes)[vs[1].(string)]
	}).(EdgeKubernetesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeKubernetesInput)(nil)).Elem(), &EdgeKubernetes{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeKubernetesPtrInput)(nil)).Elem(), &EdgeKubernetes{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeKubernetesArrayInput)(nil)).Elem(), EdgeKubernetesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeKubernetesMapInput)(nil)).Elem(), EdgeKubernetesMap{})
	pulumi.RegisterOutputType(EdgeKubernetesOutput{})
	pulumi.RegisterOutputType(EdgeKubernetesPtrOutput{})
	pulumi.RegisterOutputType(EdgeKubernetesArrayOutput{})
	pulumi.RegisterOutputType(EdgeKubernetesMapOutput{})
}
