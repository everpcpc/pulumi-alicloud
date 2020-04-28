// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS
{
    public partial class ManagedKubernetes : Pulumi.CustomResource
    {
        [Output("addons")]
        public Output<ImmutableArray<Outputs.ManagedKubernetesAddon>> Addons { get; private set; } = null!;

        /// <summary>
        /// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitch_ids` should be set, its value will be vswitch's zone.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// The path of client certificate, like `~/.kube/client-cert.pem`.
        /// </summary>
        [Output("clientCert")]
        public Output<string?> ClientCert { get; private set; } = null!;

        /// <summary>
        /// The path of client key, like `~/.kube/client-key.pem`.
        /// </summary>
        [Output("clientKey")]
        public Output<string?> ClientKey { get; private set; } = null!;

        /// <summary>
        /// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
        /// </summary>
        [Output("clusterCaCert")]
        public Output<string?> ClusterCaCert { get; private set; } = null!;

        /// <summary>
        /// Map of kubernetes cluster connection information. It contains several attributes to `Block Connections`.
        /// </summary>
        [Output("connections")]
        public Output<Outputs.ManagedKubernetesConnections> Connections { get; private set; } = null!;

        /// <summary>
        /// kubelet cpu policy. options: static|none. default: none.
        /// </summary>
        [Output("cpuPolicy")]
        public Output<string?> CpuPolicy { get; private set; } = null!;

        /// <summary>
        /// Enable login to the node through SSH. default: false 
        /// </summary>
        [Output("enableSsh")]
        public Output<bool?> EnableSsh { get; private set; } = null!;

        /// <summary>
        /// Custom Image support. Must based on CentOS7 or AliyunLinux2.
        /// </summary>
        [Output("imageId")]
        public Output<string?> ImageId { get; private set; } = null!;

        /// <summary>
        /// Install cloud monitor agent on ECS. default: true 
        /// </summary>
        [Output("installCloudMonitor")]
        public Output<bool?> InstallCloudMonitor { get; private set; } = null!;

        /// <summary>
        /// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Output("keyName")]
        public Output<string?> KeyName { get; private set; } = null!;

        /// <summary>
        /// An KMS encrypts password used to a cs kubernetes. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Output("kmsEncryptedPassword")]
        public Output<string?> KmsEncryptedPassword { get; private set; } = null!;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a cs kubernetes with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        [Output("kmsEncryptionContext")]
        public Output<ImmutableDictionary<string, object>?> KmsEncryptionContext { get; private set; } = null!;

        /// <summary>
        /// The path of kube config, like `~/.kube/config`.
        /// </summary>
        [Output("kubeConfig")]
        public Output<string?> KubeConfig { get; private set; } = null!;

        /// <summary>
        /// The kubernetes cluster's name. It is unique in one Alicloud account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// The ID of nat gateway used to launch kubernetes cluster.
        /// </summary>
        [Output("natGatewayId")]
        public Output<string> NatGatewayId { get; private set; } = null!;

        /// <summary>
        /// Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
        /// </summary>
        [Output("newNatGateway")]
        public Output<bool?> NewNatGateway { get; private set; } = null!;

        /// <summary>
        /// The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
        /// </summary>
        [Output("nodeCidrMask")]
        public Output<int?> NodeCidrMask { get; private set; } = null!;

        /// <summary>
        /// The password of ssh login cluster node. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// [Flannel Specific] The CIDR block for the pod network when using Flannel. 
        /// </summary>
        [Output("podCidr")]
        public Output<string?> PodCidr { get; private set; } = null!;

        /// <summary>
        /// [Terway Specific] The vswitches for the pod network when using Terway.Be careful the `pod_vswitch_ids` can not equal to `worker_vswtich_ids`.but must be in same availability zones.
        /// </summary>
        [Output("podVswitchIds")]
        public Output<ImmutableArray<string>> PodVswitchIds { get; private set; } = null!;

        /// <summary>
        /// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
        /// </summary>
        [Output("proxyMode")]
        public Output<string?> ProxyMode { get; private set; } = null!;

        /// <summary>
        /// The ID of security group where the current cluster worker node is located.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
        /// </summary>
        [Output("serviceCidr")]
        public Output<string?> ServiceCidr { get; private set; } = null!;

        [Output("slbId")]
        public Output<string> SlbId { get; private set; } = null!;

        [Output("slbInternet")]
        public Output<string> SlbInternet { get; private set; } = null!;

        /// <summary>
        /// Whether to create internet load balancer for API Server. Default to true.
        /// </summary>
        [Output("slbInternetEnabled")]
        public Output<bool?> SlbInternetEnabled { get; private set; } = null!;

        /// <summary>
        /// The ID of private load balancer where the current cluster master node is located.
        /// </summary>
        [Output("slbIntranet")]
        public Output<string> SlbIntranet { get; private set; } = null!;

        /// <summary>
        /// The path of customized CA cert, you can use this CA to sign client certs to connect your cluster.
        /// </summary>
        [Output("userCa")]
        public Output<string?> UserCa { get; private set; } = null!;

        /// <summary>
        /// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The ID of VPC where the current cluster is located.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// Enable worker payment auto-renew, defaults to false.
        /// </summary>
        [Output("workerAutoRenew")]
        public Output<bool?> WorkerAutoRenew { get; private set; } = null!;

        /// <summary>
        /// Worker payment auto-renew period. When period unit is `Month`, it can be one of {“1”, “2”, “3”, “6”, “12”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”}.
        /// </summary>
        [Output("workerAutoRenewPeriod")]
        public Output<int?> WorkerAutoRenewPeriod { get; private set; } = null!;

        [Output("workerDataDiskCategory")]
        public Output<string?> WorkerDataDiskCategory { get; private set; } = null!;

        [Output("workerDataDiskSize")]
        public Output<int?> WorkerDataDiskSize { get; private set; } = null!;

        /// <summary>
        /// The system disk category of worker node. Its valid value are `cloud_ssd` and `cloud_efficiency`. Default to `cloud_efficiency`.
        /// </summary>
        [Output("workerDiskCategory")]
        public Output<string?> WorkerDiskCategory { get; private set; } = null!;

        /// <summary>
        /// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 20.
        /// </summary>
        [Output("workerDiskSize")]
        public Output<int?> WorkerDiskSize { get; private set; } = null!;

        /// <summary>
        /// Worker payment type. `PrePaid` or `PostPaid`, defaults to `PostPaid`.
        /// </summary>
        [Output("workerInstanceChargeType")]
        public Output<string?> WorkerInstanceChargeType { get; private set; } = null!;

        /// <summary>
        /// The instance type of worker node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster.
        /// </summary>
        [Output("workerInstanceTypes")]
        public Output<ImmutableArray<string>> WorkerInstanceTypes { get; private set; } = null!;

        /// <summary>
        /// List of cluster worker nodes. It contains several attributes to `Block Nodes`.
        /// </summary>
        [Output("workerNodes")]
        public Output<ImmutableArray<Outputs.ManagedKubernetesWorkerNode>> WorkerNodes { get; private set; } = null!;

        /// <summary>
        /// The worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
        /// </summary>
        [Output("workerNumber")]
        public Output<int> WorkerNumber { get; private set; } = null!;

        /// <summary>
        /// Worker payment period. When period unit is `Month`, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”, “4”}.
        /// </summary>
        [Output("workerPeriod")]
        public Output<int?> WorkerPeriod { get; private set; } = null!;

        /// <summary>
        /// Worker payment period unit. `Month` or `Week`, defaults to `Month`.
        /// </summary>
        [Output("workerPeriodUnit")]
        public Output<string?> WorkerPeriodUnit { get; private set; } = null!;

        [Output("workerVswitchIds")]
        public Output<ImmutableArray<string>> WorkerVswitchIds { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedKubernetes resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedKubernetes(string name, ManagedKubernetesArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cs/managedKubernetes:ManagedKubernetes", name, args ?? new ManagedKubernetesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedKubernetes(string name, Input<string> id, ManagedKubernetesState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cs/managedKubernetes:ManagedKubernetes", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedKubernetes resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedKubernetes Get(string name, Input<string> id, ManagedKubernetesState? state = null, CustomResourceOptions? options = null)
        {
            return new ManagedKubernetes(name, id, state, options);
        }
    }

    public sealed class ManagedKubernetesArgs : Pulumi.ResourceArgs
    {
        [Input("addons")]
        private InputList<Inputs.ManagedKubernetesAddonArgs>? _addons;
        public InputList<Inputs.ManagedKubernetesAddonArgs> Addons
        {
            get => _addons ?? (_addons = new InputList<Inputs.ManagedKubernetesAddonArgs>());
            set => _addons = value;
        }

        /// <summary>
        /// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitch_ids` should be set, its value will be vswitch's zone.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The path of client certificate, like `~/.kube/client-cert.pem`.
        /// </summary>
        [Input("clientCert")]
        public Input<string>? ClientCert { get; set; }

        /// <summary>
        /// The path of client key, like `~/.kube/client-key.pem`.
        /// </summary>
        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        /// <summary>
        /// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
        /// </summary>
        [Input("clusterCaCert")]
        public Input<string>? ClusterCaCert { get; set; }

        /// <summary>
        /// kubelet cpu policy. options: static|none. default: none.
        /// </summary>
        [Input("cpuPolicy")]
        public Input<string>? CpuPolicy { get; set; }

        /// <summary>
        /// Enable login to the node through SSH. default: false 
        /// </summary>
        [Input("enableSsh")]
        public Input<bool>? EnableSsh { get; set; }

        /// <summary>
        /// Custom Image support. Must based on CentOS7 or AliyunLinux2.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// Install cloud monitor agent on ECS. default: true 
        /// </summary>
        [Input("installCloudMonitor")]
        public Input<bool>? InstallCloudMonitor { get; set; }

        /// <summary>
        /// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// An KMS encrypts password used to a cs kubernetes. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Input("kmsEncryptedPassword")]
        public Input<string>? KmsEncryptedPassword { get; set; }

        [Input("kmsEncryptionContext")]
        private InputMap<object>? _kmsEncryptionContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a cs kubernetes with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        public InputMap<object> KmsEncryptionContext
        {
            get => _kmsEncryptionContext ?? (_kmsEncryptionContext = new InputMap<object>());
            set => _kmsEncryptionContext = value;
        }

        /// <summary>
        /// The path of kube config, like `~/.kube/config`.
        /// </summary>
        [Input("kubeConfig")]
        public Input<string>? KubeConfig { get; set; }

        /// <summary>
        /// The kubernetes cluster's name. It is unique in one Alicloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
        /// </summary>
        [Input("newNatGateway")]
        public Input<bool>? NewNatGateway { get; set; }

        /// <summary>
        /// The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
        /// </summary>
        [Input("nodeCidrMask")]
        public Input<int>? NodeCidrMask { get; set; }

        /// <summary>
        /// The password of ssh login cluster node. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// [Flannel Specific] The CIDR block for the pod network when using Flannel. 
        /// </summary>
        [Input("podCidr")]
        public Input<string>? PodCidr { get; set; }

        [Input("podVswitchIds")]
        private InputList<string>? _podVswitchIds;

        /// <summary>
        /// [Terway Specific] The vswitches for the pod network when using Terway.Be careful the `pod_vswitch_ids` can not equal to `worker_vswtich_ids`.but must be in same availability zones.
        /// </summary>
        public InputList<string> PodVswitchIds
        {
            get => _podVswitchIds ?? (_podVswitchIds = new InputList<string>());
            set => _podVswitchIds = value;
        }

        /// <summary>
        /// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
        /// </summary>
        [Input("proxyMode")]
        public Input<string>? ProxyMode { get; set; }

        /// <summary>
        /// The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
        /// </summary>
        [Input("serviceCidr")]
        public Input<string>? ServiceCidr { get; set; }

        /// <summary>
        /// Whether to create internet load balancer for API Server. Default to true.
        /// </summary>
        [Input("slbInternetEnabled")]
        public Input<bool>? SlbInternetEnabled { get; set; }

        /// <summary>
        /// The path of customized CA cert, you can use this CA to sign client certs to connect your cluster.
        /// </summary>
        [Input("userCa")]
        public Input<string>? UserCa { get; set; }

        /// <summary>
        /// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// Enable worker payment auto-renew, defaults to false.
        /// </summary>
        [Input("workerAutoRenew")]
        public Input<bool>? WorkerAutoRenew { get; set; }

        /// <summary>
        /// Worker payment auto-renew period. When period unit is `Month`, it can be one of {“1”, “2”, “3”, “6”, “12”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”}.
        /// </summary>
        [Input("workerAutoRenewPeriod")]
        public Input<int>? WorkerAutoRenewPeriod { get; set; }

        [Input("workerDataDiskCategory")]
        public Input<string>? WorkerDataDiskCategory { get; set; }

        [Input("workerDataDiskSize")]
        public Input<int>? WorkerDataDiskSize { get; set; }

        /// <summary>
        /// The system disk category of worker node. Its valid value are `cloud_ssd` and `cloud_efficiency`. Default to `cloud_efficiency`.
        /// </summary>
        [Input("workerDiskCategory")]
        public Input<string>? WorkerDiskCategory { get; set; }

        /// <summary>
        /// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 20.
        /// </summary>
        [Input("workerDiskSize")]
        public Input<int>? WorkerDiskSize { get; set; }

        /// <summary>
        /// Worker payment type. `PrePaid` or `PostPaid`, defaults to `PostPaid`.
        /// </summary>
        [Input("workerInstanceChargeType")]
        public Input<string>? WorkerInstanceChargeType { get; set; }

        [Input("workerInstanceTypes", required: true)]
        private InputList<string>? _workerInstanceTypes;

        /// <summary>
        /// The instance type of worker node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster.
        /// </summary>
        public InputList<string> WorkerInstanceTypes
        {
            get => _workerInstanceTypes ?? (_workerInstanceTypes = new InputList<string>());
            set => _workerInstanceTypes = value;
        }

        /// <summary>
        /// The worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
        /// </summary>
        [Input("workerNumber", required: true)]
        public Input<int> WorkerNumber { get; set; } = null!;

        /// <summary>
        /// Worker payment period. When period unit is `Month`, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”, “4”}.
        /// </summary>
        [Input("workerPeriod")]
        public Input<int>? WorkerPeriod { get; set; }

        /// <summary>
        /// Worker payment period unit. `Month` or `Week`, defaults to `Month`.
        /// </summary>
        [Input("workerPeriodUnit")]
        public Input<string>? WorkerPeriodUnit { get; set; }

        [Input("workerVswitchIds", required: true)]
        private InputList<string>? _workerVswitchIds;
        public InputList<string> WorkerVswitchIds
        {
            get => _workerVswitchIds ?? (_workerVswitchIds = new InputList<string>());
            set => _workerVswitchIds = value;
        }

        public ManagedKubernetesArgs()
        {
        }
    }

    public sealed class ManagedKubernetesState : Pulumi.ResourceArgs
    {
        [Input("addons")]
        private InputList<Inputs.ManagedKubernetesAddonGetArgs>? _addons;
        public InputList<Inputs.ManagedKubernetesAddonGetArgs> Addons
        {
            get => _addons ?? (_addons = new InputList<Inputs.ManagedKubernetesAddonGetArgs>());
            set => _addons = value;
        }

        /// <summary>
        /// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitch_ids` should be set, its value will be vswitch's zone.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The path of client certificate, like `~/.kube/client-cert.pem`.
        /// </summary>
        [Input("clientCert")]
        public Input<string>? ClientCert { get; set; }

        /// <summary>
        /// The path of client key, like `~/.kube/client-key.pem`.
        /// </summary>
        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        /// <summary>
        /// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
        /// </summary>
        [Input("clusterCaCert")]
        public Input<string>? ClusterCaCert { get; set; }

        /// <summary>
        /// Map of kubernetes cluster connection information. It contains several attributes to `Block Connections`.
        /// </summary>
        [Input("connections")]
        public Input<Inputs.ManagedKubernetesConnectionsGetArgs>? Connections { get; set; }

        /// <summary>
        /// kubelet cpu policy. options: static|none. default: none.
        /// </summary>
        [Input("cpuPolicy")]
        public Input<string>? CpuPolicy { get; set; }

        /// <summary>
        /// Enable login to the node through SSH. default: false 
        /// </summary>
        [Input("enableSsh")]
        public Input<bool>? EnableSsh { get; set; }

        /// <summary>
        /// Custom Image support. Must based on CentOS7 or AliyunLinux2.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// Install cloud monitor agent on ECS. default: true 
        /// </summary>
        [Input("installCloudMonitor")]
        public Input<bool>? InstallCloudMonitor { get; set; }

        /// <summary>
        /// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// An KMS encrypts password used to a cs kubernetes. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Input("kmsEncryptedPassword")]
        public Input<string>? KmsEncryptedPassword { get; set; }

        [Input("kmsEncryptionContext")]
        private InputMap<object>? _kmsEncryptionContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a cs kubernetes with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        public InputMap<object> KmsEncryptionContext
        {
            get => _kmsEncryptionContext ?? (_kmsEncryptionContext = new InputMap<object>());
            set => _kmsEncryptionContext = value;
        }

        /// <summary>
        /// The path of kube config, like `~/.kube/config`.
        /// </summary>
        [Input("kubeConfig")]
        public Input<string>? KubeConfig { get; set; }

        /// <summary>
        /// The kubernetes cluster's name. It is unique in one Alicloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The ID of nat gateway used to launch kubernetes cluster.
        /// </summary>
        [Input("natGatewayId")]
        public Input<string>? NatGatewayId { get; set; }

        /// <summary>
        /// Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
        /// </summary>
        [Input("newNatGateway")]
        public Input<bool>? NewNatGateway { get; set; }

        /// <summary>
        /// The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
        /// </summary>
        [Input("nodeCidrMask")]
        public Input<int>? NodeCidrMask { get; set; }

        /// <summary>
        /// The password of ssh login cluster node. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// [Flannel Specific] The CIDR block for the pod network when using Flannel. 
        /// </summary>
        [Input("podCidr")]
        public Input<string>? PodCidr { get; set; }

        [Input("podVswitchIds")]
        private InputList<string>? _podVswitchIds;

        /// <summary>
        /// [Terway Specific] The vswitches for the pod network when using Terway.Be careful the `pod_vswitch_ids` can not equal to `worker_vswtich_ids`.but must be in same availability zones.
        /// </summary>
        public InputList<string> PodVswitchIds
        {
            get => _podVswitchIds ?? (_podVswitchIds = new InputList<string>());
            set => _podVswitchIds = value;
        }

        /// <summary>
        /// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
        /// </summary>
        [Input("proxyMode")]
        public Input<string>? ProxyMode { get; set; }

        /// <summary>
        /// The ID of security group where the current cluster worker node is located.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
        /// </summary>
        [Input("serviceCidr")]
        public Input<string>? ServiceCidr { get; set; }

        [Input("slbId")]
        public Input<string>? SlbId { get; set; }

        [Input("slbInternet")]
        public Input<string>? SlbInternet { get; set; }

        /// <summary>
        /// Whether to create internet load balancer for API Server. Default to true.
        /// </summary>
        [Input("slbInternetEnabled")]
        public Input<bool>? SlbInternetEnabled { get; set; }

        /// <summary>
        /// The ID of private load balancer where the current cluster master node is located.
        /// </summary>
        [Input("slbIntranet")]
        public Input<string>? SlbIntranet { get; set; }

        /// <summary>
        /// The path of customized CA cert, you can use this CA to sign client certs to connect your cluster.
        /// </summary>
        [Input("userCa")]
        public Input<string>? UserCa { get; set; }

        /// <summary>
        /// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The ID of VPC where the current cluster is located.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Enable worker payment auto-renew, defaults to false.
        /// </summary>
        [Input("workerAutoRenew")]
        public Input<bool>? WorkerAutoRenew { get; set; }

        /// <summary>
        /// Worker payment auto-renew period. When period unit is `Month`, it can be one of {“1”, “2”, “3”, “6”, “12”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”}.
        /// </summary>
        [Input("workerAutoRenewPeriod")]
        public Input<int>? WorkerAutoRenewPeriod { get; set; }

        [Input("workerDataDiskCategory")]
        public Input<string>? WorkerDataDiskCategory { get; set; }

        [Input("workerDataDiskSize")]
        public Input<int>? WorkerDataDiskSize { get; set; }

        /// <summary>
        /// The system disk category of worker node. Its valid value are `cloud_ssd` and `cloud_efficiency`. Default to `cloud_efficiency`.
        /// </summary>
        [Input("workerDiskCategory")]
        public Input<string>? WorkerDiskCategory { get; set; }

        /// <summary>
        /// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 20.
        /// </summary>
        [Input("workerDiskSize")]
        public Input<int>? WorkerDiskSize { get; set; }

        /// <summary>
        /// Worker payment type. `PrePaid` or `PostPaid`, defaults to `PostPaid`.
        /// </summary>
        [Input("workerInstanceChargeType")]
        public Input<string>? WorkerInstanceChargeType { get; set; }

        [Input("workerInstanceTypes")]
        private InputList<string>? _workerInstanceTypes;

        /// <summary>
        /// The instance type of worker node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster.
        /// </summary>
        public InputList<string> WorkerInstanceTypes
        {
            get => _workerInstanceTypes ?? (_workerInstanceTypes = new InputList<string>());
            set => _workerInstanceTypes = value;
        }

        [Input("workerNodes")]
        private InputList<Inputs.ManagedKubernetesWorkerNodeGetArgs>? _workerNodes;

        /// <summary>
        /// List of cluster worker nodes. It contains several attributes to `Block Nodes`.
        /// </summary>
        public InputList<Inputs.ManagedKubernetesWorkerNodeGetArgs> WorkerNodes
        {
            get => _workerNodes ?? (_workerNodes = new InputList<Inputs.ManagedKubernetesWorkerNodeGetArgs>());
            set => _workerNodes = value;
        }

        /// <summary>
        /// The worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
        /// </summary>
        [Input("workerNumber")]
        public Input<int>? WorkerNumber { get; set; }

        /// <summary>
        /// Worker payment period. When period unit is `Month`, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}.  When period unit is `Week`, it can be one of {“1”, “2”, “3”, “4”}.
        /// </summary>
        [Input("workerPeriod")]
        public Input<int>? WorkerPeriod { get; set; }

        /// <summary>
        /// Worker payment period unit. `Month` or `Week`, defaults to `Month`.
        /// </summary>
        [Input("workerPeriodUnit")]
        public Input<string>? WorkerPeriodUnit { get; set; }

        [Input("workerVswitchIds")]
        private InputList<string>? _workerVswitchIds;
        public InputList<string> WorkerVswitchIds
        {
            get => _workerVswitchIds ?? (_workerVswitchIds = new InputList<string>());
            set => _workerVswitchIds = value;
        }

        public ManagedKubernetesState()
        {
        }
    }
}