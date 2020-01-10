// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cs_managed_kubernetes.html.markdown.
    /// </summary>
    public partial class ManagedKubernetes : Pulumi.CustomResource
    {
        /// <summary>
        /// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitch_ids` should be set, the value will be vswitch's zone.
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
        /// The network that cluster uses, use `flannel` or `terway`.
        /// </summary>
        [Output("clusterNetworkType")]
        public Output<string?> ClusterNetworkType { get; private set; } = null!;

        /// <summary>
        /// Default false, when you want to change `worker_instance_types` and `vswitch_ids`, you have to set this field to true, then the cluster will be recreated.
        /// </summary>
        [Output("forceUpdate")]
        public Output<bool?> ForceUpdate { get; private set; } = null!;

        /// <summary>
        /// The ID of node image.
        /// </summary>
        [Output("imageId")]
        public Output<string?> ImageId { get; private set; } = null!;

        /// <summary>
        /// Whether to install cloud monitor for the kubernetes' node.
        /// </summary>
        [Output("installCloudMonitor")]
        public Output<bool?> InstallCloudMonitor { get; private set; } = null!;

        /// <summary>
        /// The keypair of ssh login cluster node, you have to create it first.
        /// </summary>
        [Output("keyName")]
        public Output<string?> KeyName { get; private set; } = null!;

        /// <summary>
        /// An KMS encrypts password used to a cs managed kubernetes. It is conflicted with `password` and `key_name`.
        /// </summary>
        [Output("kmsEncryptedPassword")]
        public Output<string?> KmsEncryptedPassword { get; private set; } = null!;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a cs managed kubernetes with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        [Output("kmsEncryptionContext")]
        public Output<ImmutableDictionary<string, object>?> KmsEncryptionContext { get; private set; } = null!;

        /// <summary>
        /// The path of kube config, like `~/.kube/config`.
        /// </summary>
        [Output("kubeConfig")]
        public Output<string?> KubeConfig { get; private set; } = null!;

        /// <summary>
        /// A list of one element containing information about the associated log store. It contains the following attributes:
        /// </summary>
        [Output("logConfig")]
        public Output<Outputs.ManagedKubernetesLogConfig?> LogConfig { get; private set; } = null!;

        /// <summary>
        /// The kubernetes cluster's name. It is the only in one Alicloud account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// Whether to create a new nat gateway while creating kubernetes cluster. Default to true.
        /// </summary>
        [Output("newNatGateway")]
        public Output<bool?> NewNatGateway { get; private set; } = null!;

        /// <summary>
        /// The password of ssh login cluster node. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The CIDR block for the pod network. When `cluster_network_type` is  set to `flanne`, you must set value to this filed .
        /// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
        /// Maximum number of hosts allowed in the cluster: 256. Refer to [Plan Kubernetes CIDR blocks under VPC](https://www.alibabacloud.com/help/doc-detail/64530.htm).
        /// </summary>
        [Output("podCidr")]
        public Output<string?> PodCidr { get; private set; } = null!;

        /// <summary>
        /// The ID of security group where the current cluster worker node is located.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// The CIDR block for the service network.  
        /// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
        /// </summary>
        [Output("serviceCidr")]
        public Output<string?> ServiceCidr { get; private set; } = null!;

        /// <summary>
        /// Whether to create internet load balancer for API Server. Default to false.
        /// </summary>
        [Output("slbInternetEnabled")]
        public Output<bool?> SlbInternetEnabled { get; private set; } = null!;

        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;

        /// <summary>
        /// The ID of VPC where the current cluster is located.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The vswitch where new kubernetes cluster will be located. Specify one or more vswitch's id. It must be in the zone which `availability_zone` specified.
        /// </summary>
        [Output("vswitchIds")]
        public Output<ImmutableArray<string>> VswitchIds { get; private set; } = null!;

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

        /// <summary>
        /// The data disk category of worker node. Its valid value are `cloud_ssd` and `cloud_efficiency`, if not set, data disk will not be created.
        /// </summary>
        [Output("workerDataDiskCategory")]
        public Output<string?> WorkerDataDiskCategory { get; private set; } = null!;

        /// <summary>
        /// The data disk size of worker node. Its valid value range [20~32768] in GB. When `worker_data_disk_category` is presented, it defaults to 40.
        /// </summary>
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

        [Output("workerInstanceTypes")]
        public Output<ImmutableArray<string>> WorkerInstanceTypes { get; private set; } = null!;

        /// <summary>
        /// List of cluster worker nodes. It contains several attributes to `Block Nodes`.
        /// </summary>
        [Output("workerNodes")]
        public Output<ImmutableArray<Outputs.ManagedKubernetesWorkerNodes>> WorkerNodes { get; private set; } = null!;

        /// <summary>
        /// The total worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
        /// </summary>
        [Output("workerNumber")]
        public Output<int?> WorkerNumber { get; private set; } = null!;

        /// <summary>
        /// The worker node number of the kubernetes cluster. Default to [3]. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
        /// </summary>
        [Output("workerNumbers")]
        public Output<int?> WorkerNumbers { get; private set; } = null!;

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


        /// <summary>
        /// Create a ManagedKubernetes resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedKubernetes(string name, ManagedKubernetesArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cs/managedKubernetes:ManagedKubernetes", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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
        /// <summary>
        /// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitch_ids` should be set, the value will be vswitch's zone.
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
        /// The network that cluster uses, use `flannel` or `terway`.
        /// </summary>
        [Input("clusterNetworkType")]
        public Input<string>? ClusterNetworkType { get; set; }

        /// <summary>
        /// Default false, when you want to change `worker_instance_types` and `vswitch_ids`, you have to set this field to true, then the cluster will be recreated.
        /// </summary>
        [Input("forceUpdate")]
        public Input<bool>? ForceUpdate { get; set; }

        /// <summary>
        /// The ID of node image.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// Whether to install cloud monitor for the kubernetes' node.
        /// </summary>
        [Input("installCloudMonitor")]
        public Input<bool>? InstallCloudMonitor { get; set; }

        /// <summary>
        /// The keypair of ssh login cluster node, you have to create it first.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// An KMS encrypts password used to a cs managed kubernetes. It is conflicted with `password` and `key_name`.
        /// </summary>
        [Input("kmsEncryptedPassword")]
        public Input<string>? KmsEncryptedPassword { get; set; }

        [Input("kmsEncryptionContext")]
        private InputMap<object>? _kmsEncryptionContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a cs managed kubernetes with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
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
        /// A list of one element containing information about the associated log store. It contains the following attributes:
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.ManagedKubernetesLogConfigArgs>? LogConfig { get; set; }

        /// <summary>
        /// The kubernetes cluster's name. It is the only in one Alicloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Whether to create a new nat gateway while creating kubernetes cluster. Default to true.
        /// </summary>
        [Input("newNatGateway")]
        public Input<bool>? NewNatGateway { get; set; }

        /// <summary>
        /// The password of ssh login cluster node. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The CIDR block for the pod network. When `cluster_network_type` is  set to `flanne`, you must set value to this filed .
        /// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
        /// Maximum number of hosts allowed in the cluster: 256. Refer to [Plan Kubernetes CIDR blocks under VPC](https://www.alibabacloud.com/help/doc-detail/64530.htm).
        /// </summary>
        [Input("podCidr")]
        public Input<string>? PodCidr { get; set; }

        /// <summary>
        /// The CIDR block for the service network.  
        /// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
        /// </summary>
        [Input("serviceCidr")]
        public Input<string>? ServiceCidr { get; set; }

        /// <summary>
        /// Whether to create internet load balancer for API Server. Default to false.
        /// </summary>
        [Input("slbInternetEnabled")]
        public Input<bool>? SlbInternetEnabled { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        [Input("vswitchIds")]
        private InputList<string>? _vswitchIds;

        /// <summary>
        /// The vswitch where new kubernetes cluster will be located. Specify one or more vswitch's id. It must be in the zone which `availability_zone` specified.
        /// </summary>
        public InputList<string> VswitchIds
        {
            get => _vswitchIds ?? (_vswitchIds = new InputList<string>());
            set => _vswitchIds = value;
        }

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

        /// <summary>
        /// The data disk category of worker node. Its valid value are `cloud_ssd` and `cloud_efficiency`, if not set, data disk will not be created.
        /// </summary>
        [Input("workerDataDiskCategory")]
        public Input<string>? WorkerDataDiskCategory { get; set; }

        /// <summary>
        /// The data disk size of worker node. Its valid value range [20~32768] in GB. When `worker_data_disk_category` is presented, it defaults to 40.
        /// </summary>
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
        public InputList<string> WorkerInstanceTypes
        {
            get => _workerInstanceTypes ?? (_workerInstanceTypes = new InputList<string>());
            set => _workerInstanceTypes = value;
        }

        /// <summary>
        /// The total worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
        /// </summary>
        [Input("workerNumber")]
        public Input<int>? WorkerNumber { get; set; }

        /// <summary>
        /// The worker node number of the kubernetes cluster. Default to [3]. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
        /// </summary>
        [Input("workerNumbers")]
        public Input<int>? WorkerNumbers { get; set; }

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

        public ManagedKubernetesArgs()
        {
        }
    }

    public sealed class ManagedKubernetesState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Zone where new kubernetes cluster will be located. If it is not be specified, the `vswitch_ids` should be set, the value will be vswitch's zone.
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
        /// The network that cluster uses, use `flannel` or `terway`.
        /// </summary>
        [Input("clusterNetworkType")]
        public Input<string>? ClusterNetworkType { get; set; }

        /// <summary>
        /// Default false, when you want to change `worker_instance_types` and `vswitch_ids`, you have to set this field to true, then the cluster will be recreated.
        /// </summary>
        [Input("forceUpdate")]
        public Input<bool>? ForceUpdate { get; set; }

        /// <summary>
        /// The ID of node image.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// Whether to install cloud monitor for the kubernetes' node.
        /// </summary>
        [Input("installCloudMonitor")]
        public Input<bool>? InstallCloudMonitor { get; set; }

        /// <summary>
        /// The keypair of ssh login cluster node, you have to create it first.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// An KMS encrypts password used to a cs managed kubernetes. It is conflicted with `password` and `key_name`.
        /// </summary>
        [Input("kmsEncryptedPassword")]
        public Input<string>? KmsEncryptedPassword { get; set; }

        [Input("kmsEncryptionContext")]
        private InputMap<object>? _kmsEncryptionContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a cs managed kubernetes with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
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
        /// A list of one element containing information about the associated log store. It contains the following attributes:
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.ManagedKubernetesLogConfigGetArgs>? LogConfig { get; set; }

        /// <summary>
        /// The kubernetes cluster's name. It is the only in one Alicloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Whether to create a new nat gateway while creating kubernetes cluster. Default to true.
        /// </summary>
        [Input("newNatGateway")]
        public Input<bool>? NewNatGateway { get; set; }

        /// <summary>
        /// The password of ssh login cluster node. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The CIDR block for the pod network. When `cluster_network_type` is  set to `flanne`, you must set value to this filed .
        /// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
        /// Maximum number of hosts allowed in the cluster: 256. Refer to [Plan Kubernetes CIDR blocks under VPC](https://www.alibabacloud.com/help/doc-detail/64530.htm).
        /// </summary>
        [Input("podCidr")]
        public Input<string>? PodCidr { get; set; }

        /// <summary>
        /// The ID of security group where the current cluster worker node is located.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// The CIDR block for the service network.  
        /// It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
        /// </summary>
        [Input("serviceCidr")]
        public Input<string>? ServiceCidr { get; set; }

        /// <summary>
        /// Whether to create internet load balancer for API Server. Default to false.
        /// </summary>
        [Input("slbInternetEnabled")]
        public Input<bool>? SlbInternetEnabled { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The ID of VPC where the current cluster is located.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        [Input("vswitchIds")]
        private InputList<string>? _vswitchIds;

        /// <summary>
        /// The vswitch where new kubernetes cluster will be located. Specify one or more vswitch's id. It must be in the zone which `availability_zone` specified.
        /// </summary>
        public InputList<string> VswitchIds
        {
            get => _vswitchIds ?? (_vswitchIds = new InputList<string>());
            set => _vswitchIds = value;
        }

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

        /// <summary>
        /// The data disk category of worker node. Its valid value are `cloud_ssd` and `cloud_efficiency`, if not set, data disk will not be created.
        /// </summary>
        [Input("workerDataDiskCategory")]
        public Input<string>? WorkerDataDiskCategory { get; set; }

        /// <summary>
        /// The data disk size of worker node. Its valid value range [20~32768] in GB. When `worker_data_disk_category` is presented, it defaults to 40.
        /// </summary>
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
        public InputList<string> WorkerInstanceTypes
        {
            get => _workerInstanceTypes ?? (_workerInstanceTypes = new InputList<string>());
            set => _workerInstanceTypes = value;
        }

        [Input("workerNodes")]
        private InputList<Inputs.ManagedKubernetesWorkerNodesGetArgs>? _workerNodes;

        /// <summary>
        /// List of cluster worker nodes. It contains several attributes to `Block Nodes`.
        /// </summary>
        public InputList<Inputs.ManagedKubernetesWorkerNodesGetArgs> WorkerNodes
        {
            get => _workerNodes ?? (_workerNodes = new InputList<Inputs.ManagedKubernetesWorkerNodesGetArgs>());
            set => _workerNodes = value;
        }

        /// <summary>
        /// The total worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
        /// </summary>
        [Input("workerNumber")]
        public Input<int>? WorkerNumber { get; set; }

        /// <summary>
        /// The worker node number of the kubernetes cluster. Default to [3]. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
        /// </summary>
        [Input("workerNumbers")]
        public Input<int>? WorkerNumbers { get; set; }

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

        public ManagedKubernetesState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ManagedKubernetesLogConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Log Service project name, cluster logs will output to this project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Type of collecting logs, only `SLS` are supported currently.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ManagedKubernetesLogConfigArgs()
        {
        }
    }

    public sealed class ManagedKubernetesLogConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Log Service project name, cluster logs will output to this project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Type of collecting logs, only `SLS` are supported currently.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ManagedKubernetesLogConfigGetArgs()
        {
        }
    }

    public sealed class ManagedKubernetesWorkerNodesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the node.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The kubernetes cluster's name. It is the only in one Alicloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The private IP address of node.
        /// </summary>
        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        public ManagedKubernetesWorkerNodesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ManagedKubernetesLogConfig
    {
        /// <summary>
        /// Log Service project name, cluster logs will output to this project.
        /// </summary>
        public readonly string? Project;
        /// <summary>
        /// Type of collecting logs, only `SLS` are supported currently.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ManagedKubernetesLogConfig(
            string? project,
            string type)
        {
            Project = project;
            Type = type;
        }
    }

    [OutputType]
    public sealed class ManagedKubernetesWorkerNodes
    {
        /// <summary>
        /// ID of the node.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The kubernetes cluster's name. It is the only in one Alicloud account.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The private IP address of node.
        /// </summary>
        public readonly string PrivateIp;

        [OutputConstructor]
        private ManagedKubernetesWorkerNodes(
            string id,
            string name,
            string privateIp)
        {
            Id = id;
            Name = name;
            PrivateIp = privateIp;
        }
    }
    }
}
