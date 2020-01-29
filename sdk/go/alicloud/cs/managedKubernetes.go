// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cs

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cs_managed_kubernetes.html.markdown.
type ManagedKubernetes struct {
	pulumi.CustomResourceState

	// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitchIds` should be set, the value will be vswitch's zone.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert pulumi.StringPtrOutput `pulumi:"clientCert"`
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey pulumi.StringPtrOutput `pulumi:"clientKey"`
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert pulumi.StringPtrOutput `pulumi:"clusterCaCert"`
	// The network that cluster uses, use `flannel` or `terway`.
	ClusterNetworkType pulumi.StringPtrOutput `pulumi:"clusterNetworkType"`
	// Default false, when you want to change `workerInstanceTypes` and `vswitchIds`, you have to set this field to true, then the cluster will be recreated.
	ForceUpdate pulumi.BoolPtrOutput `pulumi:"forceUpdate"`
	// The ID of node image.
	ImageId pulumi.StringPtrOutput `pulumi:"imageId"`
	// Whether to install cloud monitor for the kubernetes' node.
	InstallCloudMonitor pulumi.BoolPtrOutput `pulumi:"installCloudMonitor"`
	// The keypair of ssh login cluster node, you have to create it first.
	KeyName pulumi.StringPtrOutput `pulumi:"keyName"`
	// An KMS encrypts password used to a cs managed kubernetes. It is conflicted with `password` and `keyName`.
	KmsEncryptedPassword pulumi.StringPtrOutput `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a cs managed kubernetes with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapOutput `pulumi:"kmsEncryptionContext"`
	// The path of kube config, like `~/.kube/config`.
	KubeConfig pulumi.StringPtrOutput `pulumi:"kubeConfig"`
	// A list of one element containing information about the associated log store. It contains the following attributes:
	LogConfig ManagedKubernetesLogConfigPtrOutput `pulumi:"logConfig"`
	// The kubernetes cluster's name. It is the only in one Alicloud account.
	Name pulumi.StringOutput `pulumi:"name"`
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true.
	NewNatGateway pulumi.BoolPtrOutput `pulumi:"newNatGateway"`
	// The password of ssh login cluster node. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The CIDR block for the pod network. When `clusterNetworkType` is  set to `flanne`, you must set value to this filed .
	// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	// Maximum number of hosts allowed in the cluster: 256. Refer to [Plan Kubernetes CIDR blocks under VPC](https://www.alibabacloud.com/help/doc-detail/64530.htm).
	PodCidr pulumi.StringPtrOutput `pulumi:"podCidr"`
	// The ID of security group where the current cluster worker node is located.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// The CIDR block for the service network.  
	// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	ServiceCidr pulumi.StringPtrOutput `pulumi:"serviceCidr"`
	// Whether to create internet load balancer for API Server. Default to false.
	SlbInternetEnabled pulumi.BoolPtrOutput `pulumi:"slbInternetEnabled"`
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
	Version pulumi.StringOutput `pulumi:"version"`
	// The ID of VPC where the current cluster is located.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The vswitch where new kubernetes cluster will be located. Specify one or more vswitch's id. It must be in the zone which `availabilityZone` specified.
	VswitchIds pulumi.StringArrayOutput `pulumi:"vswitchIds"`
	// Enable worker payment auto-renew, defaults to false.
	WorkerAutoRenew pulumi.BoolPtrOutput `pulumi:"workerAutoRenew"`
	// Worker payment auto-renew period. When period unit is `Month`, it can be one of {“1”, “2”, “3”, “6”, “12”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”}.
	WorkerAutoRenewPeriod pulumi.IntPtrOutput `pulumi:"workerAutoRenewPeriod"`
	// The data disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`, if not set, data disk will not be created.
	WorkerDataDiskCategory pulumi.StringPtrOutput `pulumi:"workerDataDiskCategory"`
	// The data disk size of worker node. Its valid value range [20~32768] in GB. When `workerDataDiskCategory` is presented, it defaults to 40.
	WorkerDataDiskSize pulumi.IntPtrOutput `pulumi:"workerDataDiskSize"`
	// The system disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`. Default to `cloudEfficiency`.
	WorkerDiskCategory pulumi.StringPtrOutput `pulumi:"workerDiskCategory"`
	// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 20.
	WorkerDiskSize pulumi.IntPtrOutput `pulumi:"workerDiskSize"`
	// Worker payment type. `PrePaid` or `PostPaid`, defaults to `PostPaid`.
	WorkerInstanceChargeType pulumi.StringPtrOutput `pulumi:"workerInstanceChargeType"`
	WorkerInstanceTypes pulumi.StringArrayOutput `pulumi:"workerInstanceTypes"`
	// List of cluster worker nodes. It contains several attributes to `Block Nodes`.
	WorkerNodes ManagedKubernetesWorkerNodeArrayOutput `pulumi:"workerNodes"`
	// The total worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumber pulumi.IntPtrOutput `pulumi:"workerNumber"`
	// The worker node number of the kubernetes cluster. Default to [3]. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumbers pulumi.IntPtrOutput `pulumi:"workerNumbers"`
	// Worker payment period. When period unit is `Month`, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”, “4”}.
	WorkerPeriod pulumi.IntPtrOutput `pulumi:"workerPeriod"`
	// Worker payment period unit. `Month` or `Week`, defaults to `Month`.
	WorkerPeriodUnit pulumi.StringPtrOutput `pulumi:"workerPeriodUnit"`
}

// NewManagedKubernetes registers a new resource with the given unique name, arguments, and options.
func NewManagedKubernetes(ctx *pulumi.Context,
	name string, args *ManagedKubernetesArgs, opts ...pulumi.ResourceOption) (*ManagedKubernetes, error) {
	if args == nil || args.WorkerInstanceTypes == nil {
		return nil, errors.New("missing required argument 'WorkerInstanceTypes'")
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
	// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitchIds` should be set, the value will be vswitch's zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert *string `pulumi:"clientCert"`
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey *string `pulumi:"clientKey"`
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert *string `pulumi:"clusterCaCert"`
	// The network that cluster uses, use `flannel` or `terway`.
	ClusterNetworkType *string `pulumi:"clusterNetworkType"`
	// Default false, when you want to change `workerInstanceTypes` and `vswitchIds`, you have to set this field to true, then the cluster will be recreated.
	ForceUpdate *bool `pulumi:"forceUpdate"`
	// The ID of node image.
	ImageId *string `pulumi:"imageId"`
	// Whether to install cloud monitor for the kubernetes' node.
	InstallCloudMonitor *bool `pulumi:"installCloudMonitor"`
	// The keypair of ssh login cluster node, you have to create it first.
	KeyName *string `pulumi:"keyName"`
	// An KMS encrypts password used to a cs managed kubernetes. It is conflicted with `password` and `keyName`.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a cs managed kubernetes with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The path of kube config, like `~/.kube/config`.
	KubeConfig *string `pulumi:"kubeConfig"`
	// A list of one element containing information about the associated log store. It contains the following attributes:
	LogConfig *ManagedKubernetesLogConfig `pulumi:"logConfig"`
	// The kubernetes cluster's name. It is the only in one Alicloud account.
	Name *string `pulumi:"name"`
	NamePrefix *string `pulumi:"namePrefix"`
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true.
	NewNatGateway *bool `pulumi:"newNatGateway"`
	// The password of ssh login cluster node. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	Password *string `pulumi:"password"`
	// The CIDR block for the pod network. When `clusterNetworkType` is  set to `flanne`, you must set value to this filed .
	// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	// Maximum number of hosts allowed in the cluster: 256. Refer to [Plan Kubernetes CIDR blocks under VPC](https://www.alibabacloud.com/help/doc-detail/64530.htm).
	PodCidr *string `pulumi:"podCidr"`
	// The ID of security group where the current cluster worker node is located.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The CIDR block for the service network.  
	// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	ServiceCidr *string `pulumi:"serviceCidr"`
	// Whether to create internet load balancer for API Server. Default to false.
	SlbInternetEnabled *bool `pulumi:"slbInternetEnabled"`
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
	Version *string `pulumi:"version"`
	// The ID of VPC where the current cluster is located.
	VpcId *string `pulumi:"vpcId"`
	// The vswitch where new kubernetes cluster will be located. Specify one or more vswitch's id. It must be in the zone which `availabilityZone` specified.
	VswitchIds []string `pulumi:"vswitchIds"`
	// Enable worker payment auto-renew, defaults to false.
	WorkerAutoRenew *bool `pulumi:"workerAutoRenew"`
	// Worker payment auto-renew period. When period unit is `Month`, it can be one of {“1”, “2”, “3”, “6”, “12”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”}.
	WorkerAutoRenewPeriod *int `pulumi:"workerAutoRenewPeriod"`
	// The data disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`, if not set, data disk will not be created.
	WorkerDataDiskCategory *string `pulumi:"workerDataDiskCategory"`
	// The data disk size of worker node. Its valid value range [20~32768] in GB. When `workerDataDiskCategory` is presented, it defaults to 40.
	WorkerDataDiskSize *int `pulumi:"workerDataDiskSize"`
	// The system disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`. Default to `cloudEfficiency`.
	WorkerDiskCategory *string `pulumi:"workerDiskCategory"`
	// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 20.
	WorkerDiskSize *int `pulumi:"workerDiskSize"`
	// Worker payment type. `PrePaid` or `PostPaid`, defaults to `PostPaid`.
	WorkerInstanceChargeType *string `pulumi:"workerInstanceChargeType"`
	WorkerInstanceTypes []string `pulumi:"workerInstanceTypes"`
	// List of cluster worker nodes. It contains several attributes to `Block Nodes`.
	WorkerNodes []ManagedKubernetesWorkerNode `pulumi:"workerNodes"`
	// The total worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumber *int `pulumi:"workerNumber"`
	// The worker node number of the kubernetes cluster. Default to [3]. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumbers *int `pulumi:"workerNumbers"`
	// Worker payment period. When period unit is `Month`, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”, “4”}.
	WorkerPeriod *int `pulumi:"workerPeriod"`
	// Worker payment period unit. `Month` or `Week`, defaults to `Month`.
	WorkerPeriodUnit *string `pulumi:"workerPeriodUnit"`
}

type ManagedKubernetesState struct {
	// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitchIds` should be set, the value will be vswitch's zone.
	AvailabilityZone pulumi.StringPtrInput
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert pulumi.StringPtrInput
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey pulumi.StringPtrInput
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert pulumi.StringPtrInput
	// The network that cluster uses, use `flannel` or `terway`.
	ClusterNetworkType pulumi.StringPtrInput
	// Default false, when you want to change `workerInstanceTypes` and `vswitchIds`, you have to set this field to true, then the cluster will be recreated.
	ForceUpdate pulumi.BoolPtrInput
	// The ID of node image.
	ImageId pulumi.StringPtrInput
	// Whether to install cloud monitor for the kubernetes' node.
	InstallCloudMonitor pulumi.BoolPtrInput
	// The keypair of ssh login cluster node, you have to create it first.
	KeyName pulumi.StringPtrInput
	// An KMS encrypts password used to a cs managed kubernetes. It is conflicted with `password` and `keyName`.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a cs managed kubernetes with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The path of kube config, like `~/.kube/config`.
	KubeConfig pulumi.StringPtrInput
	// A list of one element containing information about the associated log store. It contains the following attributes:
	LogConfig ManagedKubernetesLogConfigPtrInput
	// The kubernetes cluster's name. It is the only in one Alicloud account.
	Name pulumi.StringPtrInput
	NamePrefix pulumi.StringPtrInput
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true.
	NewNatGateway pulumi.BoolPtrInput
	// The password of ssh login cluster node. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	Password pulumi.StringPtrInput
	// The CIDR block for the pod network. When `clusterNetworkType` is  set to `flanne`, you must set value to this filed .
	// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	// Maximum number of hosts allowed in the cluster: 256. Refer to [Plan Kubernetes CIDR blocks under VPC](https://www.alibabacloud.com/help/doc-detail/64530.htm).
	PodCidr pulumi.StringPtrInput
	// The ID of security group where the current cluster worker node is located.
	SecurityGroupId pulumi.StringPtrInput
	// The CIDR block for the service network.  
	// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	ServiceCidr pulumi.StringPtrInput
	// Whether to create internet load balancer for API Server. Default to false.
	SlbInternetEnabled pulumi.BoolPtrInput
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
	Version pulumi.StringPtrInput
	// The ID of VPC where the current cluster is located.
	VpcId pulumi.StringPtrInput
	// The vswitch where new kubernetes cluster will be located. Specify one or more vswitch's id. It must be in the zone which `availabilityZone` specified.
	VswitchIds pulumi.StringArrayInput
	// Enable worker payment auto-renew, defaults to false.
	WorkerAutoRenew pulumi.BoolPtrInput
	// Worker payment auto-renew period. When period unit is `Month`, it can be one of {“1”, “2”, “3”, “6”, “12”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”}.
	WorkerAutoRenewPeriod pulumi.IntPtrInput
	// The data disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`, if not set, data disk will not be created.
	WorkerDataDiskCategory pulumi.StringPtrInput
	// The data disk size of worker node. Its valid value range [20~32768] in GB. When `workerDataDiskCategory` is presented, it defaults to 40.
	WorkerDataDiskSize pulumi.IntPtrInput
	// The system disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`. Default to `cloudEfficiency`.
	WorkerDiskCategory pulumi.StringPtrInput
	// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 20.
	WorkerDiskSize pulumi.IntPtrInput
	// Worker payment type. `PrePaid` or `PostPaid`, defaults to `PostPaid`.
	WorkerInstanceChargeType pulumi.StringPtrInput
	WorkerInstanceTypes pulumi.StringArrayInput
	// List of cluster worker nodes. It contains several attributes to `Block Nodes`.
	WorkerNodes ManagedKubernetesWorkerNodeArrayInput
	// The total worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumber pulumi.IntPtrInput
	// The worker node number of the kubernetes cluster. Default to [3]. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumbers pulumi.IntPtrInput
	// Worker payment period. When period unit is `Month`, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”, “4”}.
	WorkerPeriod pulumi.IntPtrInput
	// Worker payment period unit. `Month` or `Week`, defaults to `Month`.
	WorkerPeriodUnit pulumi.StringPtrInput
}

func (ManagedKubernetesState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedKubernetesState)(nil)).Elem()
}

type managedKubernetesArgs struct {
	// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitchIds` should be set, the value will be vswitch's zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert *string `pulumi:"clientCert"`
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey *string `pulumi:"clientKey"`
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert *string `pulumi:"clusterCaCert"`
	// The network that cluster uses, use `flannel` or `terway`.
	ClusterNetworkType *string `pulumi:"clusterNetworkType"`
	// Default false, when you want to change `workerInstanceTypes` and `vswitchIds`, you have to set this field to true, then the cluster will be recreated.
	ForceUpdate *bool `pulumi:"forceUpdate"`
	// The ID of node image.
	ImageId *string `pulumi:"imageId"`
	// Whether to install cloud monitor for the kubernetes' node.
	InstallCloudMonitor *bool `pulumi:"installCloudMonitor"`
	// The keypair of ssh login cluster node, you have to create it first.
	KeyName *string `pulumi:"keyName"`
	// An KMS encrypts password used to a cs managed kubernetes. It is conflicted with `password` and `keyName`.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a cs managed kubernetes with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The path of kube config, like `~/.kube/config`.
	KubeConfig *string `pulumi:"kubeConfig"`
	// A list of one element containing information about the associated log store. It contains the following attributes:
	LogConfig *ManagedKubernetesLogConfig `pulumi:"logConfig"`
	// The kubernetes cluster's name. It is the only in one Alicloud account.
	Name *string `pulumi:"name"`
	NamePrefix *string `pulumi:"namePrefix"`
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true.
	NewNatGateway *bool `pulumi:"newNatGateway"`
	// The password of ssh login cluster node. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	Password *string `pulumi:"password"`
	// The CIDR block for the pod network. When `clusterNetworkType` is  set to `flanne`, you must set value to this filed .
	// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	// Maximum number of hosts allowed in the cluster: 256. Refer to [Plan Kubernetes CIDR blocks under VPC](https://www.alibabacloud.com/help/doc-detail/64530.htm).
	PodCidr *string `pulumi:"podCidr"`
	// The CIDR block for the service network.  
	// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	ServiceCidr *string `pulumi:"serviceCidr"`
	// Whether to create internet load balancer for API Server. Default to false.
	SlbInternetEnabled *bool `pulumi:"slbInternetEnabled"`
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
	Version *string `pulumi:"version"`
	// The vswitch where new kubernetes cluster will be located. Specify one or more vswitch's id. It must be in the zone which `availabilityZone` specified.
	VswitchIds []string `pulumi:"vswitchIds"`
	// Enable worker payment auto-renew, defaults to false.
	WorkerAutoRenew *bool `pulumi:"workerAutoRenew"`
	// Worker payment auto-renew period. When period unit is `Month`, it can be one of {“1”, “2”, “3”, “6”, “12”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”}.
	WorkerAutoRenewPeriod *int `pulumi:"workerAutoRenewPeriod"`
	// The data disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`, if not set, data disk will not be created.
	WorkerDataDiskCategory *string `pulumi:"workerDataDiskCategory"`
	// The data disk size of worker node. Its valid value range [20~32768] in GB. When `workerDataDiskCategory` is presented, it defaults to 40.
	WorkerDataDiskSize *int `pulumi:"workerDataDiskSize"`
	// The system disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`. Default to `cloudEfficiency`.
	WorkerDiskCategory *string `pulumi:"workerDiskCategory"`
	// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 20.
	WorkerDiskSize *int `pulumi:"workerDiskSize"`
	// Worker payment type. `PrePaid` or `PostPaid`, defaults to `PostPaid`.
	WorkerInstanceChargeType *string `pulumi:"workerInstanceChargeType"`
	WorkerInstanceTypes []string `pulumi:"workerInstanceTypes"`
	// The total worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumber *int `pulumi:"workerNumber"`
	// The worker node number of the kubernetes cluster. Default to [3]. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumbers *int `pulumi:"workerNumbers"`
	// Worker payment period. When period unit is `Month`, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”, “4”}.
	WorkerPeriod *int `pulumi:"workerPeriod"`
	// Worker payment period unit. `Month` or `Week`, defaults to `Month`.
	WorkerPeriodUnit *string `pulumi:"workerPeriodUnit"`
}

// The set of arguments for constructing a ManagedKubernetes resource.
type ManagedKubernetesArgs struct {
	// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitchIds` should be set, the value will be vswitch's zone.
	AvailabilityZone pulumi.StringPtrInput
	// The path of client certificate, like `~/.kube/client-cert.pem`.
	ClientCert pulumi.StringPtrInput
	// The path of client key, like `~/.kube/client-key.pem`.
	ClientKey pulumi.StringPtrInput
	// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
	ClusterCaCert pulumi.StringPtrInput
	// The network that cluster uses, use `flannel` or `terway`.
	ClusterNetworkType pulumi.StringPtrInput
	// Default false, when you want to change `workerInstanceTypes` and `vswitchIds`, you have to set this field to true, then the cluster will be recreated.
	ForceUpdate pulumi.BoolPtrInput
	// The ID of node image.
	ImageId pulumi.StringPtrInput
	// Whether to install cloud monitor for the kubernetes' node.
	InstallCloudMonitor pulumi.BoolPtrInput
	// The keypair of ssh login cluster node, you have to create it first.
	KeyName pulumi.StringPtrInput
	// An KMS encrypts password used to a cs managed kubernetes. It is conflicted with `password` and `keyName`.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a cs managed kubernetes with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The path of kube config, like `~/.kube/config`.
	KubeConfig pulumi.StringPtrInput
	// A list of one element containing information about the associated log store. It contains the following attributes:
	LogConfig ManagedKubernetesLogConfigPtrInput
	// The kubernetes cluster's name. It is the only in one Alicloud account.
	Name pulumi.StringPtrInput
	NamePrefix pulumi.StringPtrInput
	// Whether to create a new nat gateway while creating kubernetes cluster. Default to true.
	NewNatGateway pulumi.BoolPtrInput
	// The password of ssh login cluster node. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
	Password pulumi.StringPtrInput
	// The CIDR block for the pod network. When `clusterNetworkType` is  set to `flanne`, you must set value to this filed .
	// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	// Maximum number of hosts allowed in the cluster: 256. Refer to [Plan Kubernetes CIDR blocks under VPC](https://www.alibabacloud.com/help/doc-detail/64530.htm).
	PodCidr pulumi.StringPtrInput
	// The CIDR block for the service network.  
	// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
	ServiceCidr pulumi.StringPtrInput
	// Whether to create internet load balancer for API Server. Default to false.
	SlbInternetEnabled pulumi.BoolPtrInput
	// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
	Version pulumi.StringPtrInput
	// The vswitch where new kubernetes cluster will be located. Specify one or more vswitch's id. It must be in the zone which `availabilityZone` specified.
	VswitchIds pulumi.StringArrayInput
	// Enable worker payment auto-renew, defaults to false.
	WorkerAutoRenew pulumi.BoolPtrInput
	// Worker payment auto-renew period. When period unit is `Month`, it can be one of {“1”, “2”, “3”, “6”, “12”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”}.
	WorkerAutoRenewPeriod pulumi.IntPtrInput
	// The data disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`, if not set, data disk will not be created.
	WorkerDataDiskCategory pulumi.StringPtrInput
	// The data disk size of worker node. Its valid value range [20~32768] in GB. When `workerDataDiskCategory` is presented, it defaults to 40.
	WorkerDataDiskSize pulumi.IntPtrInput
	// The system disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`. Default to `cloudEfficiency`.
	WorkerDiskCategory pulumi.StringPtrInput
	// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 20.
	WorkerDiskSize pulumi.IntPtrInput
	// Worker payment type. `PrePaid` or `PostPaid`, defaults to `PostPaid`.
	WorkerInstanceChargeType pulumi.StringPtrInput
	WorkerInstanceTypes pulumi.StringArrayInput
	// The total worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumber pulumi.IntPtrInput
	// The worker node number of the kubernetes cluster. Default to [3]. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
	WorkerNumbers pulumi.IntPtrInput
	// Worker payment period. When period unit is `Month`, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”, “4”}.
	WorkerPeriod pulumi.IntPtrInput
	// Worker payment period unit. `Month` or `Week`, defaults to `Month`.
	WorkerPeriodUnit pulumi.StringPtrInput
}

func (ManagedKubernetesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedKubernetesArgs)(nil)).Elem()
}

