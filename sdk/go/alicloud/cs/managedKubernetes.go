// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ManagedKubernetes struct {
	pulumi.CustomResourceState

	Addons ManagedKubernetesAddonArrayOutput `pulumi:"addons"`
	// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitchIds` should be set, its value will be vswitch's zone.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert pulumi.StringPtrOutput `pulumi:"clientCert"`
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey pulumi.StringPtrOutput `pulumi:"clientKey"`
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert pulumi.StringPtrOutput `pulumi:"clusterCaCert"`
	// Map of kubernetes cluster connection information. It contains several attributes to `Block Connections`.
	Connections ManagedKubernetesConnectionsOutput `pulumi:"connections"`
	// kubelet cpu policy. options: static|none. default: none.
	CpuPolicy pulumi.StringPtrOutput `pulumi:"cpuPolicy"`
	// Enable login to the node through SSH. default: false
	EnableSsh pulumi.BoolPtrOutput `pulumi:"enableSsh"`
	// Custom Image support. Must based on CentOS7 or AliyunLinux2.
	ImageId pulumi.StringPtrOutput `pulumi:"imageId"`
	// Install cloud monitor agent on ECS. default: true
	InstallCloudMonitor pulumi.BoolPtrOutput `pulumi:"installCloudMonitor"`
	// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	KeyName pulumi.StringPtrOutput `pulumi:"keyName"`
	// An KMS encrypts password used to a cs kubernetes. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	KmsEncryptedPassword pulumi.StringPtrOutput `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a cs kubernetes with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapOutput `pulumi:"kmsEncryptionContext"`
	// The path of kube config, like `~/.kube/config`.
	KubeConfig pulumi.StringPtrOutput `pulumi:"kubeConfig"`
	// The kubernetes cluster's name. It is unique in one Alicloud account.
	Name       pulumi.StringOutput    `pulumi:"name"`
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// The ID of nat gateway used to launch kubernetes cluster.
	NatGatewayId pulumi.StringOutput `pulumi:"natGatewayId"`
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
	NewNatGateway pulumi.BoolPtrOutput `pulumi:"newNatGateway"`
	// The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
	NodeCidrMask pulumi.IntPtrOutput `pulumi:"nodeCidrMask"`
	// The password of ssh login cluster node. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// [Flannel Specific] The CIDR block for the pod network when using Flannel.
	PodCidr pulumi.StringPtrOutput `pulumi:"podCidr"`
	// [Terway Specific] The vswitches for the pod network when using Terway.Be careful the `podVswitchIds` can not equal to `workerVswtichIds`.but must be in same availability zones.
	PodVswitchIds pulumi.StringArrayOutput `pulumi:"podVswitchIds"`
	// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
	ProxyMode pulumi.StringPtrOutput `pulumi:"proxyMode"`
	// The ID of security group where the current cluster worker node is located.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	ServiceCidr pulumi.StringPtrOutput `pulumi:"serviceCidr"`
	SlbId       pulumi.StringOutput    `pulumi:"slbId"`
	SlbInternet pulumi.StringOutput    `pulumi:"slbInternet"`
	// Whether to create internet load balancer for API Server. Default to true.
	SlbInternetEnabled pulumi.BoolPtrOutput `pulumi:"slbInternetEnabled"`
	// The ID of private load balancer where the current cluster master node is located.
	SlbIntranet pulumi.StringOutput `pulumi:"slbIntranet"`
	// The path of customized CA cert, you can use this CA to sign client certs to connect your cluster.
	UserCa pulumi.StringPtrOutput `pulumi:"userCa"`
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
	Version pulumi.StringPtrOutput `pulumi:"version"`
	// The ID of VPC where the current cluster is located.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// Enable worker payment auto-renew, defaults to false.
	WorkerAutoRenew pulumi.BoolPtrOutput `pulumi:"workerAutoRenew"`
	// Worker payment auto-renew period. When period unit is `Month`, it can be one of {“1”, “2”, “3”, “6”, “12”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”}.
	WorkerAutoRenewPeriod  pulumi.IntPtrOutput    `pulumi:"workerAutoRenewPeriod"`
	WorkerDataDiskCategory pulumi.StringPtrOutput `pulumi:"workerDataDiskCategory"`
	WorkerDataDiskSize     pulumi.IntPtrOutput    `pulumi:"workerDataDiskSize"`
	// The system disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`. Default to `cloudEfficiency`.
	WorkerDiskCategory pulumi.StringPtrOutput `pulumi:"workerDiskCategory"`
	// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 20.
	WorkerDiskSize pulumi.IntPtrOutput `pulumi:"workerDiskSize"`
	// Worker payment type. `PrePaid` or `PostPaid`, defaults to `PostPaid`.
	WorkerInstanceChargeType pulumi.StringPtrOutput `pulumi:"workerInstanceChargeType"`
	// The instance type of worker node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster.
	WorkerInstanceTypes pulumi.StringArrayOutput `pulumi:"workerInstanceTypes"`
	// List of cluster worker nodes. It contains several attributes to `Block Nodes`.
	WorkerNodes ManagedKubernetesWorkerNodeArrayOutput `pulumi:"workerNodes"`
	// The worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumber pulumi.IntOutput `pulumi:"workerNumber"`
	// Worker payment period. When period unit is `Month`, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”, “4”}.
	WorkerPeriod pulumi.IntPtrOutput `pulumi:"workerPeriod"`
	// Worker payment period unit. `Month` or `Week`, defaults to `Month`.
	WorkerPeriodUnit pulumi.StringPtrOutput   `pulumi:"workerPeriodUnit"`
	WorkerVswitchIds pulumi.StringArrayOutput `pulumi:"workerVswitchIds"`
}

// NewManagedKubernetes registers a new resource with the given unique name, arguments, and options.
func NewManagedKubernetes(ctx *pulumi.Context,
	name string, args *ManagedKubernetesArgs, opts ...pulumi.ResourceOption) (*ManagedKubernetes, error) {
	if args == nil || args.WorkerInstanceTypes == nil {
		return nil, errors.New("missing required argument 'WorkerInstanceTypes'")
	}
	if args == nil || args.WorkerNumber == nil {
		return nil, errors.New("missing required argument 'WorkerNumber'")
	}
	if args == nil || args.WorkerVswitchIds == nil {
		return nil, errors.New("missing required argument 'WorkerVswitchIds'")
	}
	if args == nil {
		args = &ManagedKubernetesArgs{}
	}
	var resource ManagedKubernetes
	err := ctx.RegisterResource("alicloud:cs/managedKubernetes:ManagedKubernetes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedKubernetes gets an existing ManagedKubernetes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedKubernetes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedKubernetesState, opts ...pulumi.ResourceOption) (*ManagedKubernetes, error) {
	var resource ManagedKubernetes
	err := ctx.ReadResource("alicloud:cs/managedKubernetes:ManagedKubernetes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedKubernetes resources.
type managedKubernetesState struct {
	Addons []ManagedKubernetesAddon `pulumi:"addons"`
	// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitchIds` should be set, its value will be vswitch's zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert *string `pulumi:"clientCert"`
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey *string `pulumi:"clientKey"`
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert *string `pulumi:"clusterCaCert"`
	// Map of kubernetes cluster connection information. It contains several attributes to `Block Connections`.
	Connections *ManagedKubernetesConnections `pulumi:"connections"`
	// kubelet cpu policy. options: static|none. default: none.
	CpuPolicy *string `pulumi:"cpuPolicy"`
	// Enable login to the node through SSH. default: false
	EnableSsh *bool `pulumi:"enableSsh"`
	// Custom Image support. Must based on CentOS7 or AliyunLinux2.
	ImageId *string `pulumi:"imageId"`
	// Install cloud monitor agent on ECS. default: true
	InstallCloudMonitor *bool `pulumi:"installCloudMonitor"`
	// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	KeyName *string `pulumi:"keyName"`
	// An KMS encrypts password used to a cs kubernetes. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a cs kubernetes with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The path of kube config, like `~/.kube/config`.
	KubeConfig *string `pulumi:"kubeConfig"`
	// The kubernetes cluster's name. It is unique in one Alicloud account.
	Name       *string `pulumi:"name"`
	NamePrefix *string `pulumi:"namePrefix"`
	// The ID of nat gateway used to launch kubernetes cluster.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
	NewNatGateway *bool `pulumi:"newNatGateway"`
	// The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
	NodeCidrMask *int `pulumi:"nodeCidrMask"`
	// The password of ssh login cluster node. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	Password *string `pulumi:"password"`
	// [Flannel Specific] The CIDR block for the pod network when using Flannel.
	PodCidr *string `pulumi:"podCidr"`
	// [Terway Specific] The vswitches for the pod network when using Terway.Be careful the `podVswitchIds` can not equal to `workerVswtichIds`.but must be in same availability zones.
	PodVswitchIds []string `pulumi:"podVswitchIds"`
	// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
	ProxyMode *string `pulumi:"proxyMode"`
	// The ID of security group where the current cluster worker node is located.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	ServiceCidr *string `pulumi:"serviceCidr"`
	SlbId       *string `pulumi:"slbId"`
	SlbInternet *string `pulumi:"slbInternet"`
	// Whether to create internet load balancer for API Server. Default to true.
	SlbInternetEnabled *bool `pulumi:"slbInternetEnabled"`
	// The ID of private load balancer where the current cluster master node is located.
	SlbIntranet *string `pulumi:"slbIntranet"`
	// The path of customized CA cert, you can use this CA to sign client certs to connect your cluster.
	UserCa *string `pulumi:"userCa"`
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
	Version *string `pulumi:"version"`
	// The ID of VPC where the current cluster is located.
	VpcId *string `pulumi:"vpcId"`
	// Enable worker payment auto-renew, defaults to false.
	WorkerAutoRenew *bool `pulumi:"workerAutoRenew"`
	// Worker payment auto-renew period. When period unit is `Month`, it can be one of {“1”, “2”, “3”, “6”, “12”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”}.
	WorkerAutoRenewPeriod  *int    `pulumi:"workerAutoRenewPeriod"`
	WorkerDataDiskCategory *string `pulumi:"workerDataDiskCategory"`
	WorkerDataDiskSize     *int    `pulumi:"workerDataDiskSize"`
	// The system disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`. Default to `cloudEfficiency`.
	WorkerDiskCategory *string `pulumi:"workerDiskCategory"`
	// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 20.
	WorkerDiskSize *int `pulumi:"workerDiskSize"`
	// Worker payment type. `PrePaid` or `PostPaid`, defaults to `PostPaid`.
	WorkerInstanceChargeType *string `pulumi:"workerInstanceChargeType"`
	// The instance type of worker node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster.
	WorkerInstanceTypes []string `pulumi:"workerInstanceTypes"`
	// List of cluster worker nodes. It contains several attributes to `Block Nodes`.
	WorkerNodes []ManagedKubernetesWorkerNode `pulumi:"workerNodes"`
	// The worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumber *int `pulumi:"workerNumber"`
	// Worker payment period. When period unit is `Month`, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”, “4”}.
	WorkerPeriod *int `pulumi:"workerPeriod"`
	// Worker payment period unit. `Month` or `Week`, defaults to `Month`.
	WorkerPeriodUnit *string  `pulumi:"workerPeriodUnit"`
	WorkerVswitchIds []string `pulumi:"workerVswitchIds"`
}

type ManagedKubernetesState struct {
	Addons ManagedKubernetesAddonArrayInput
	// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitchIds` should be set, its value will be vswitch's zone.
	AvailabilityZone pulumi.StringPtrInput
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert pulumi.StringPtrInput
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey pulumi.StringPtrInput
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert pulumi.StringPtrInput
	// Map of kubernetes cluster connection information. It contains several attributes to `Block Connections`.
	Connections ManagedKubernetesConnectionsPtrInput
	// kubelet cpu policy. options: static|none. default: none.
	CpuPolicy pulumi.StringPtrInput
	// Enable login to the node through SSH. default: false
	EnableSsh pulumi.BoolPtrInput
	// Custom Image support. Must based on CentOS7 or AliyunLinux2.
	ImageId pulumi.StringPtrInput
	// Install cloud monitor agent on ECS. default: true
	InstallCloudMonitor pulumi.BoolPtrInput
	// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	KeyName pulumi.StringPtrInput
	// An KMS encrypts password used to a cs kubernetes. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a cs kubernetes with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The path of kube config, like `~/.kube/config`.
	KubeConfig pulumi.StringPtrInput
	// The kubernetes cluster's name. It is unique in one Alicloud account.
	Name       pulumi.StringPtrInput
	NamePrefix pulumi.StringPtrInput
	// The ID of nat gateway used to launch kubernetes cluster.
	NatGatewayId pulumi.StringPtrInput
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
	NewNatGateway pulumi.BoolPtrInput
	// The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
	NodeCidrMask pulumi.IntPtrInput
	// The password of ssh login cluster node. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	Password pulumi.StringPtrInput
	// [Flannel Specific] The CIDR block for the pod network when using Flannel.
	PodCidr pulumi.StringPtrInput
	// [Terway Specific] The vswitches for the pod network when using Terway.Be careful the `podVswitchIds` can not equal to `workerVswtichIds`.but must be in same availability zones.
	PodVswitchIds pulumi.StringArrayInput
	// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
	ProxyMode pulumi.StringPtrInput
	// The ID of security group where the current cluster worker node is located.
	SecurityGroupId pulumi.StringPtrInput
	// The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	ServiceCidr pulumi.StringPtrInput
	SlbId       pulumi.StringPtrInput
	SlbInternet pulumi.StringPtrInput
	// Whether to create internet load balancer for API Server. Default to true.
	SlbInternetEnabled pulumi.BoolPtrInput
	// The ID of private load balancer where the current cluster master node is located.
	SlbIntranet pulumi.StringPtrInput
	// The path of customized CA cert, you can use this CA to sign client certs to connect your cluster.
	UserCa pulumi.StringPtrInput
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
	Version pulumi.StringPtrInput
	// The ID of VPC where the current cluster is located.
	VpcId pulumi.StringPtrInput
	// Enable worker payment auto-renew, defaults to false.
	WorkerAutoRenew pulumi.BoolPtrInput
	// Worker payment auto-renew period. When period unit is `Month`, it can be one of {“1”, “2”, “3”, “6”, “12”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”}.
	WorkerAutoRenewPeriod  pulumi.IntPtrInput
	WorkerDataDiskCategory pulumi.StringPtrInput
	WorkerDataDiskSize     pulumi.IntPtrInput
	// The system disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`. Default to `cloudEfficiency`.
	WorkerDiskCategory pulumi.StringPtrInput
	// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 20.
	WorkerDiskSize pulumi.IntPtrInput
	// Worker payment type. `PrePaid` or `PostPaid`, defaults to `PostPaid`.
	WorkerInstanceChargeType pulumi.StringPtrInput
	// The instance type of worker node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster.
	WorkerInstanceTypes pulumi.StringArrayInput
	// List of cluster worker nodes. It contains several attributes to `Block Nodes`.
	WorkerNodes ManagedKubernetesWorkerNodeArrayInput
	// The worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumber pulumi.IntPtrInput
	// Worker payment period. When period unit is `Month`, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”, “4”}.
	WorkerPeriod pulumi.IntPtrInput
	// Worker payment period unit. `Month` or `Week`, defaults to `Month`.
	WorkerPeriodUnit pulumi.StringPtrInput
	WorkerVswitchIds pulumi.StringArrayInput
}

func (ManagedKubernetesState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedKubernetesState)(nil)).Elem()
}

type managedKubernetesArgs struct {
	Addons []ManagedKubernetesAddon `pulumi:"addons"`
	// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitchIds` should be set, its value will be vswitch's zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert *string `pulumi:"clientCert"`
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey *string `pulumi:"clientKey"`
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert *string `pulumi:"clusterCaCert"`
	// kubelet cpu policy. options: static|none. default: none.
	CpuPolicy *string `pulumi:"cpuPolicy"`
	// Enable login to the node through SSH. default: false
	EnableSsh *bool `pulumi:"enableSsh"`
	// Custom Image support. Must based on CentOS7 or AliyunLinux2.
	ImageId *string `pulumi:"imageId"`
	// Install cloud monitor agent on ECS. default: true
	InstallCloudMonitor *bool `pulumi:"installCloudMonitor"`
	// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	KeyName *string `pulumi:"keyName"`
	// An KMS encrypts password used to a cs kubernetes. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a cs kubernetes with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The path of kube config, like `~/.kube/config`.
	KubeConfig *string `pulumi:"kubeConfig"`
	// The kubernetes cluster's name. It is unique in one Alicloud account.
	Name       *string `pulumi:"name"`
	NamePrefix *string `pulumi:"namePrefix"`
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
	NewNatGateway *bool `pulumi:"newNatGateway"`
	// The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
	NodeCidrMask *int `pulumi:"nodeCidrMask"`
	// The password of ssh login cluster node. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	Password *string `pulumi:"password"`
	// [Flannel Specific] The CIDR block for the pod network when using Flannel.
	PodCidr *string `pulumi:"podCidr"`
	// [Terway Specific] The vswitches for the pod network when using Terway.Be careful the `podVswitchIds` can not equal to `workerVswtichIds`.but must be in same availability zones.
	PodVswitchIds []string `pulumi:"podVswitchIds"`
	// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
	ProxyMode *string `pulumi:"proxyMode"`
	// The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	ServiceCidr *string `pulumi:"serviceCidr"`
	// Whether to create internet load balancer for API Server. Default to true.
	SlbInternetEnabled *bool `pulumi:"slbInternetEnabled"`
	// The path of customized CA cert, you can use this CA to sign client certs to connect your cluster.
	UserCa *string `pulumi:"userCa"`
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
	Version *string `pulumi:"version"`
	// Enable worker payment auto-renew, defaults to false.
	WorkerAutoRenew *bool `pulumi:"workerAutoRenew"`
	// Worker payment auto-renew period. When period unit is `Month`, it can be one of {“1”, “2”, “3”, “6”, “12”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”}.
	WorkerAutoRenewPeriod  *int    `pulumi:"workerAutoRenewPeriod"`
	WorkerDataDiskCategory *string `pulumi:"workerDataDiskCategory"`
	WorkerDataDiskSize     *int    `pulumi:"workerDataDiskSize"`
	// The system disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`. Default to `cloudEfficiency`.
	WorkerDiskCategory *string `pulumi:"workerDiskCategory"`
	// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 20.
	WorkerDiskSize *int `pulumi:"workerDiskSize"`
	// Worker payment type. `PrePaid` or `PostPaid`, defaults to `PostPaid`.
	WorkerInstanceChargeType *string `pulumi:"workerInstanceChargeType"`
	// The instance type of worker node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster.
	WorkerInstanceTypes []string `pulumi:"workerInstanceTypes"`
	// The worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumber int `pulumi:"workerNumber"`
	// Worker payment period. When period unit is `Month`, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”, “4”}.
	WorkerPeriod *int `pulumi:"workerPeriod"`
	// Worker payment period unit. `Month` or `Week`, defaults to `Month`.
	WorkerPeriodUnit *string  `pulumi:"workerPeriodUnit"`
	WorkerVswitchIds []string `pulumi:"workerVswitchIds"`
}

// The set of arguments for constructing a ManagedKubernetes resource.
type ManagedKubernetesArgs struct {
	Addons ManagedKubernetesAddonArrayInput
	// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitchIds` should be set, its value will be vswitch's zone.
	AvailabilityZone pulumi.StringPtrInput
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert pulumi.StringPtrInput
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey pulumi.StringPtrInput
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert pulumi.StringPtrInput
	// kubelet cpu policy. options: static|none. default: none.
	CpuPolicy pulumi.StringPtrInput
	// Enable login to the node through SSH. default: false
	EnableSsh pulumi.BoolPtrInput
	// Custom Image support. Must based on CentOS7 or AliyunLinux2.
	ImageId pulumi.StringPtrInput
	// Install cloud monitor agent on ECS. default: true
	InstallCloudMonitor pulumi.BoolPtrInput
	// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	KeyName pulumi.StringPtrInput
	// An KMS encrypts password used to a cs kubernetes. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a cs kubernetes with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The path of kube config, like `~/.kube/config`.
	KubeConfig pulumi.StringPtrInput
	// The kubernetes cluster's name. It is unique in one Alicloud account.
	Name       pulumi.StringPtrInput
	NamePrefix pulumi.StringPtrInput
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
	NewNatGateway pulumi.BoolPtrInput
	// The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
	NodeCidrMask pulumi.IntPtrInput
	// The password of ssh login cluster node. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	Password pulumi.StringPtrInput
	// [Flannel Specific] The CIDR block for the pod network when using Flannel.
	PodCidr pulumi.StringPtrInput
	// [Terway Specific] The vswitches for the pod network when using Terway.Be careful the `podVswitchIds` can not equal to `workerVswtichIds`.but must be in same availability zones.
	PodVswitchIds pulumi.StringArrayInput
	// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
	ProxyMode pulumi.StringPtrInput
	// The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	ServiceCidr pulumi.StringPtrInput
	// Whether to create internet load balancer for API Server. Default to true.
	SlbInternetEnabled pulumi.BoolPtrInput
	// The path of customized CA cert, you can use this CA to sign client certs to connect your cluster.
	UserCa pulumi.StringPtrInput
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
	Version pulumi.StringPtrInput
	// Enable worker payment auto-renew, defaults to false.
	WorkerAutoRenew pulumi.BoolPtrInput
	// Worker payment auto-renew period. When period unit is `Month`, it can be one of {“1”, “2”, “3”, “6”, “12”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”}.
	WorkerAutoRenewPeriod  pulumi.IntPtrInput
	WorkerDataDiskCategory pulumi.StringPtrInput
	WorkerDataDiskSize     pulumi.IntPtrInput
	// The system disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`. Default to `cloudEfficiency`.
	WorkerDiskCategory pulumi.StringPtrInput
	// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 20.
	WorkerDiskSize pulumi.IntPtrInput
	// Worker payment type. `PrePaid` or `PostPaid`, defaults to `PostPaid`.
	WorkerInstanceChargeType pulumi.StringPtrInput
	// The instance type of worker node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster.
	WorkerInstanceTypes pulumi.StringArrayInput
	// The worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumber pulumi.IntInput
	// Worker payment period. When period unit is `Month`, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”, “4”}.
	WorkerPeriod pulumi.IntPtrInput
	// Worker payment period unit. `Month` or `Week`, defaults to `Month`.
	WorkerPeriodUnit pulumi.StringPtrInput
	WorkerVswitchIds pulumi.StringArrayInput
}

func (ManagedKubernetesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedKubernetesArgs)(nil)).Elem()
}
