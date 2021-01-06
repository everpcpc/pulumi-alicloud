// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS
{
    /// <summary>
    /// ## Import
    /// 
    /// Kubernetes cluster can be imported using the id, e.g. Then complete the main.tf accords to the result of `terraform plan`
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cs/edgeKubernetes:EdgeKubernetes alicloud_cs_edge_kubernetes.main cluster-id
    /// ```
    /// </summary>
    public partial class EdgeKubernetes : Pulumi.CustomResource
    {
        [Output("addons")]
        public Output<ImmutableArray<Outputs.EdgeKubernetesAddon>> Addons { get; private set; } = null!;

        /// <summary>
        /// The ID of availability zone.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// (Available in 1.105.0+) Nested attribute containing certificate authority data for your cluster.
        /// </summary>
        [Output("certificateAuthority")]
        public Output<Outputs.EdgeKubernetesCertificateAuthority> CertificateAuthority { get; private set; } = null!;

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

        [Output("connections")]
        public Output<Outputs.EdgeKubernetesConnections> Connections { get; private set; } = null!;

        /// <summary>
        /// Whether to enable cluster deletion protection.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool?> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// Install cloud monitor agent on ECS. default: `true`.
        /// </summary>
        [Output("installCloudMonitor")]
        public Output<bool?> InstallCloudMonitor { get; private set; } = null!;

        /// <summary>
        /// Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
        /// </summary>
        [Output("isEnterpriseSecurityGroup")]
        public Output<bool> IsEnterpriseSecurityGroup { get; private set; } = null!;

        /// <summary>
        /// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Output("keyName")]
        public Output<string?> KeyName { get; private set; } = null!;

        /// <summary>
        /// The path of kube config, like `~/.kube/config`.
        /// </summary>
        [Output("kubeConfig")]
        public Output<string?> KubeConfig { get; private set; } = null!;

        [Output("logConfig")]
        public Output<Outputs.EdgeKubernetesLogConfig?> LogConfig { get; private set; } = null!;

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
        /// The password of ssh login cluster node. You have to specify one of `password`, `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// [Flannel Specific] The CIDR block for the pod network when using Flannel.
        /// </summary>
        [Output("podCidr")]
        public Output<string?> PodCidr { get; private set; } = null!;

        /// <summary>
        /// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
        /// </summary>
        [Output("proxyMode")]
        public Output<string?> ProxyMode { get; private set; } = null!;

        [Output("rdsInstances")]
        public Output<ImmutableArray<string>> RdsInstances { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string?> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
        /// </summary>
        [Output("serviceCidr")]
        public Output<string?> ServiceCidr { get; private set; } = null!;

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
        /// Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
        /// </summary>
        [Output("userData")]
        public Output<string?> UserData { get; private set; } = null!;

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
        /// The data disk configurations of worker nodes, such as the disk type and disk size.
        /// * `category`: the type of the data disks. Valid values:
        /// * cloud : basic disks.
        /// * cloud_efficiency : ultra disks.
        /// * cloud_ssd : SSDs.
        /// * cloud_essd : ESSDs.
        /// * `size`: the size of a data disk, at least 40. Unit: GiB.
        /// * `encrypted`: specifies whether to encrypt data disks. Valid values: true and false.
        /// </summary>
        [Output("workerDataDisks")]
        public Output<ImmutableArray<Outputs.EdgeKubernetesWorkerDataDisk>> WorkerDataDisks { get; private set; } = null!;

        /// <summary>
        /// The system disk category of worker node. Its valid value are `cloud_efficiency`, `cloud_ssd` and `cloud_essd` and . Default to `cloud_efficiency`.
        /// </summary>
        [Output("workerDiskCategory")]
        public Output<string?> WorkerDiskCategory { get; private set; } = null!;

        /// <summary>
        /// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
        /// </summary>
        [Output("workerDiskSize")]
        public Output<int?> WorkerDiskSize { get; private set; } = null!;

        [Output("workerInstanceChargeType")]
        public Output<string?> WorkerInstanceChargeType { get; private set; } = null!;

        /// <summary>
        /// The instance types of worker node, you can set multiple types to avoid NoStock of a certain type
        /// </summary>
        [Output("workerInstanceTypes")]
        public Output<ImmutableArray<string>> WorkerInstanceTypes { get; private set; } = null!;

        /// <summary>
        /// List of cluster worker nodes.
        /// </summary>
        [Output("workerNodes")]
        public Output<ImmutableArray<Outputs.EdgeKubernetesWorkerNode>> WorkerNodes { get; private set; } = null!;

        /// <summary>
        /// The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
        /// </summary>
        [Output("workerNumber")]
        public Output<int> WorkerNumber { get; private set; } = null!;

        [Output("workerVswitchIds")]
        public Output<ImmutableArray<string>> WorkerVswitchIds { get; private set; } = null!;


        /// <summary>
        /// Create a EdgeKubernetes resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EdgeKubernetes(string name, EdgeKubernetesArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cs/edgeKubernetes:EdgeKubernetes", name, args ?? new EdgeKubernetesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EdgeKubernetes(string name, Input<string> id, EdgeKubernetesState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cs/edgeKubernetes:EdgeKubernetes", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EdgeKubernetes resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EdgeKubernetes Get(string name, Input<string> id, EdgeKubernetesState? state = null, CustomResourceOptions? options = null)
        {
            return new EdgeKubernetes(name, id, state, options);
        }
    }

    public sealed class EdgeKubernetesArgs : Pulumi.ResourceArgs
    {
        [Input("addons")]
        private InputList<Inputs.EdgeKubernetesAddonArgs>? _addons;
        public InputList<Inputs.EdgeKubernetesAddonArgs> Addons
        {
            get => _addons ?? (_addons = new InputList<Inputs.EdgeKubernetesAddonArgs>());
            set => _addons = value;
        }

        /// <summary>
        /// The ID of availability zone.
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
        /// Whether to enable cluster deletion protection.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Install cloud monitor agent on ECS. default: `true`.
        /// </summary>
        [Input("installCloudMonitor")]
        public Input<bool>? InstallCloudMonitor { get; set; }

        /// <summary>
        /// Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
        /// </summary>
        [Input("isEnterpriseSecurityGroup")]
        public Input<bool>? IsEnterpriseSecurityGroup { get; set; }

        /// <summary>
        /// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// The path of kube config, like `~/.kube/config`.
        /// </summary>
        [Input("kubeConfig")]
        public Input<string>? KubeConfig { get; set; }

        [Input("logConfig")]
        public Input<Inputs.EdgeKubernetesLogConfigArgs>? LogConfig { get; set; }

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
        /// The password of ssh login cluster node. You have to specify one of `password`, `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// [Flannel Specific] The CIDR block for the pod network when using Flannel.
        /// </summary>
        [Input("podCidr")]
        public Input<string>? PodCidr { get; set; }

        /// <summary>
        /// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
        /// </summary>
        [Input("proxyMode")]
        public Input<string>? ProxyMode { get; set; }

        [Input("rdsInstances")]
        private InputList<string>? _rdsInstances;
        public InputList<string> RdsInstances
        {
            get => _rdsInstances ?? (_rdsInstances = new InputList<string>());
            set => _rdsInstances = value;
        }

        /// <summary>
        /// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

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
        /// Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        /// <summary>
        /// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        [Input("workerDataDisks")]
        private InputList<Inputs.EdgeKubernetesWorkerDataDiskArgs>? _workerDataDisks;

        /// <summary>
        /// The data disk configurations of worker nodes, such as the disk type and disk size.
        /// * `category`: the type of the data disks. Valid values:
        /// * cloud : basic disks.
        /// * cloud_efficiency : ultra disks.
        /// * cloud_ssd : SSDs.
        /// * cloud_essd : ESSDs.
        /// * `size`: the size of a data disk, at least 40. Unit: GiB.
        /// * `encrypted`: specifies whether to encrypt data disks. Valid values: true and false.
        /// </summary>
        public InputList<Inputs.EdgeKubernetesWorkerDataDiskArgs> WorkerDataDisks
        {
            get => _workerDataDisks ?? (_workerDataDisks = new InputList<Inputs.EdgeKubernetesWorkerDataDiskArgs>());
            set => _workerDataDisks = value;
        }

        /// <summary>
        /// The system disk category of worker node. Its valid value are `cloud_efficiency`, `cloud_ssd` and `cloud_essd` and . Default to `cloud_efficiency`.
        /// </summary>
        [Input("workerDiskCategory")]
        public Input<string>? WorkerDiskCategory { get; set; }

        /// <summary>
        /// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
        /// </summary>
        [Input("workerDiskSize")]
        public Input<int>? WorkerDiskSize { get; set; }

        [Input("workerInstanceChargeType")]
        public Input<string>? WorkerInstanceChargeType { get; set; }

        [Input("workerInstanceTypes", required: true)]
        private InputList<string>? _workerInstanceTypes;

        /// <summary>
        /// The instance types of worker node, you can set multiple types to avoid NoStock of a certain type
        /// </summary>
        public InputList<string> WorkerInstanceTypes
        {
            get => _workerInstanceTypes ?? (_workerInstanceTypes = new InputList<string>());
            set => _workerInstanceTypes = value;
        }

        /// <summary>
        /// The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
        /// </summary>
        [Input("workerNumber", required: true)]
        public Input<int> WorkerNumber { get; set; } = null!;

        [Input("workerVswitchIds", required: true)]
        private InputList<string>? _workerVswitchIds;
        public InputList<string> WorkerVswitchIds
        {
            get => _workerVswitchIds ?? (_workerVswitchIds = new InputList<string>());
            set => _workerVswitchIds = value;
        }

        public EdgeKubernetesArgs()
        {
        }
    }

    public sealed class EdgeKubernetesState : Pulumi.ResourceArgs
    {
        [Input("addons")]
        private InputList<Inputs.EdgeKubernetesAddonGetArgs>? _addons;
        public InputList<Inputs.EdgeKubernetesAddonGetArgs> Addons
        {
            get => _addons ?? (_addons = new InputList<Inputs.EdgeKubernetesAddonGetArgs>());
            set => _addons = value;
        }

        /// <summary>
        /// The ID of availability zone.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// (Available in 1.105.0+) Nested attribute containing certificate authority data for your cluster.
        /// </summary>
        [Input("certificateAuthority")]
        public Input<Inputs.EdgeKubernetesCertificateAuthorityGetArgs>? CertificateAuthority { get; set; }

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

        [Input("connections")]
        public Input<Inputs.EdgeKubernetesConnectionsGetArgs>? Connections { get; set; }

        /// <summary>
        /// Whether to enable cluster deletion protection.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Install cloud monitor agent on ECS. default: `true`.
        /// </summary>
        [Input("installCloudMonitor")]
        public Input<bool>? InstallCloudMonitor { get; set; }

        /// <summary>
        /// Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
        /// </summary>
        [Input("isEnterpriseSecurityGroup")]
        public Input<bool>? IsEnterpriseSecurityGroup { get; set; }

        /// <summary>
        /// The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// The path of kube config, like `~/.kube/config`.
        /// </summary>
        [Input("kubeConfig")]
        public Input<string>? KubeConfig { get; set; }

        [Input("logConfig")]
        public Input<Inputs.EdgeKubernetesLogConfigGetArgs>? LogConfig { get; set; }

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
        /// The password of ssh login cluster node. You have to specify one of `password`, `key_name` `kms_encrypted_password` fields.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// [Flannel Specific] The CIDR block for the pod network when using Flannel.
        /// </summary>
        [Input("podCidr")]
        public Input<string>? PodCidr { get; set; }

        /// <summary>
        /// Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
        /// </summary>
        [Input("proxyMode")]
        public Input<string>? ProxyMode { get; set; }

        [Input("rdsInstances")]
        private InputList<string>? _rdsInstances;
        public InputList<string> RdsInstances
        {
            get => _rdsInstances ?? (_rdsInstances = new InputList<string>());
            set => _rdsInstances = value;
        }

        /// <summary>
        /// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
        /// </summary>
        [Input("serviceCidr")]
        public Input<string>? ServiceCidr { get; set; }

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
        /// Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

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

        [Input("workerDataDisks")]
        private InputList<Inputs.EdgeKubernetesWorkerDataDiskGetArgs>? _workerDataDisks;

        /// <summary>
        /// The data disk configurations of worker nodes, such as the disk type and disk size.
        /// * `category`: the type of the data disks. Valid values:
        /// * cloud : basic disks.
        /// * cloud_efficiency : ultra disks.
        /// * cloud_ssd : SSDs.
        /// * cloud_essd : ESSDs.
        /// * `size`: the size of a data disk, at least 40. Unit: GiB.
        /// * `encrypted`: specifies whether to encrypt data disks. Valid values: true and false.
        /// </summary>
        public InputList<Inputs.EdgeKubernetesWorkerDataDiskGetArgs> WorkerDataDisks
        {
            get => _workerDataDisks ?? (_workerDataDisks = new InputList<Inputs.EdgeKubernetesWorkerDataDiskGetArgs>());
            set => _workerDataDisks = value;
        }

        /// <summary>
        /// The system disk category of worker node. Its valid value are `cloud_efficiency`, `cloud_ssd` and `cloud_essd` and . Default to `cloud_efficiency`.
        /// </summary>
        [Input("workerDiskCategory")]
        public Input<string>? WorkerDiskCategory { get; set; }

        /// <summary>
        /// The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
        /// </summary>
        [Input("workerDiskSize")]
        public Input<int>? WorkerDiskSize { get; set; }

        [Input("workerInstanceChargeType")]
        public Input<string>? WorkerInstanceChargeType { get; set; }

        [Input("workerInstanceTypes")]
        private InputList<string>? _workerInstanceTypes;

        /// <summary>
        /// The instance types of worker node, you can set multiple types to avoid NoStock of a certain type
        /// </summary>
        public InputList<string> WorkerInstanceTypes
        {
            get => _workerInstanceTypes ?? (_workerInstanceTypes = new InputList<string>());
            set => _workerInstanceTypes = value;
        }

        [Input("workerNodes")]
        private InputList<Inputs.EdgeKubernetesWorkerNodeGetArgs>? _workerNodes;

        /// <summary>
        /// List of cluster worker nodes.
        /// </summary>
        public InputList<Inputs.EdgeKubernetesWorkerNodeGetArgs> WorkerNodes
        {
            get => _workerNodes ?? (_workerNodes = new InputList<Inputs.EdgeKubernetesWorkerNodeGetArgs>());
            set => _workerNodes = value;
        }

        /// <summary>
        /// The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
        /// </summary>
        [Input("workerNumber")]
        public Input<int>? WorkerNumber { get; set; }

        [Input("workerVswitchIds")]
        private InputList<string>? _workerVswitchIds;
        public InputList<string> WorkerVswitchIds
        {
            get => _workerVswitchIds ?? (_workerVswitchIds = new InputList<string>());
            set => _workerVswitchIds = value;
        }

        public EdgeKubernetesState()
        {
        }
    }
}