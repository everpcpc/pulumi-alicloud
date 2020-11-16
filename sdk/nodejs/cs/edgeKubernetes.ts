// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class EdgeKubernetes extends pulumi.CustomResource {
    /**
     * Get an existing EdgeKubernetes resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EdgeKubernetesState, opts?: pulumi.CustomResourceOptions): EdgeKubernetes {
        return new EdgeKubernetes(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cs/edgeKubernetes:EdgeKubernetes';

    /**
     * Returns true if the given object is an instance of EdgeKubernetes.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EdgeKubernetes {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EdgeKubernetes.__pulumiType;
    }

    public readonly addons!: pulumi.Output<outputs.cs.EdgeKubernetesAddon[] | undefined>;
    /**
     * The ID of availability zone.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * The path of client certificate, like `~/.kube/client-cert.pem`.
     */
    public readonly clientCert!: pulumi.Output<string | undefined>;
    /**
     * The path of client key, like `~/.kube/client-key.pem`.
     */
    public readonly clientKey!: pulumi.Output<string | undefined>;
    /**
     * The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
     */
    public readonly clusterCaCert!: pulumi.Output<string | undefined>;
    public /*out*/ readonly connections!: pulumi.Output<outputs.cs.EdgeKubernetesConnections>;
    /**
     * Whether to enable cluster deletion protection.
     */
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    /**
     * Install cloud monitor agent on ECS. default: `true`.
     */
    public readonly installCloudMonitor!: pulumi.Output<boolean | undefined>;
    /**
     * Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
     */
    public readonly isEnterpriseSecurityGroup!: pulumi.Output<boolean>;
    /**
     * The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
     */
    public readonly keyName!: pulumi.Output<string | undefined>;
    /**
     * The path of kube config, like `~/.kube/config`.
     */
    public readonly kubeConfig!: pulumi.Output<string | undefined>;
    public readonly logConfig!: pulumi.Output<outputs.cs.EdgeKubernetesLogConfig | undefined>;
    /**
     * The kubernetes cluster's name. It is unique in one Alicloud account.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * The ID of nat gateway used to launch kubernetes cluster.
     */
    public /*out*/ readonly natGatewayId!: pulumi.Output<string>;
    /**
     * Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
     */
    public readonly newNatGateway!: pulumi.Output<boolean | undefined>;
    /**
     * The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
     */
    public readonly nodeCidrMask!: pulumi.Output<number | undefined>;
    /**
     * The password of ssh login cluster node. You have to specify one of `password`, `keyName` `kmsEncryptedPassword` fields.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * [Flannel Specific] The CIDR block for the pod network when using Flannel.
     */
    public readonly podCidr!: pulumi.Output<string | undefined>;
    /**
     * Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
     */
    public readonly proxyMode!: pulumi.Output<string | undefined>;
    public readonly rdsInstances!: pulumi.Output<string[] | undefined>;
    /**
     * The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
     */
    public readonly serviceCidr!: pulumi.Output<string | undefined>;
    public /*out*/ readonly slbInternet!: pulumi.Output<string>;
    /**
     * Whether to create internet load balancer for API Server. Default to true.
     */
    public readonly slbInternetEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of private load balancer where the current cluster master node is located.
     */
    public /*out*/ readonly slbIntranet!: pulumi.Output<string>;
    /**
     * Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
     */
    public readonly userData!: pulumi.Output<string | undefined>;
    /**
     * Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * The ID of VPC where the current cluster is located.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * The data disk configurations of worker nodes, such as the disk type and disk size.
     * * `category`: the type of the data disks. Valid values:
     * * cloud: basic disks.
     * * cloud_efficiency: ultra disks.
     * * cloud_ssd: SSDs.
     * * `size`: the size of a data disk. Unit: GiB.
     * * `encrypted`: specifies whether to encrypt data disks. Valid values: true and false.
     */
    public readonly workerDataDisks!: pulumi.Output<outputs.cs.EdgeKubernetesWorkerDataDisk[] | undefined>;
    /**
     * The system disk category of worker node. Its valid value are `cloudEfficiency`, `cloudSsd` and `cloudEssd` and . Default to `cloudEfficiency`.
     */
    public readonly workerDiskCategory!: pulumi.Output<string | undefined>;
    /**
     * The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
     */
    public readonly workerDiskSize!: pulumi.Output<number | undefined>;
    public readonly workerInstanceChargeType!: pulumi.Output<string | undefined>;
    /**
     * The instance types of worker node, you can set multiple types to avoid NoStock of a certain type
     */
    public readonly workerInstanceTypes!: pulumi.Output<string[]>;
    /**
     * List of cluster worker nodes. It contains several attributes to `Block Nodes`.
     */
    public /*out*/ readonly workerNodes!: pulumi.Output<outputs.cs.EdgeKubernetesWorkerNode[]>;
    /**
     * The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
     */
    public readonly workerNumber!: pulumi.Output<number>;
    public readonly workerVswitchIds!: pulumi.Output<string[]>;

    /**
     * Create a EdgeKubernetes resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EdgeKubernetesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EdgeKubernetesArgs | EdgeKubernetesState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EdgeKubernetesState | undefined;
            inputs["addons"] = state ? state.addons : undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["clientCert"] = state ? state.clientCert : undefined;
            inputs["clientKey"] = state ? state.clientKey : undefined;
            inputs["clusterCaCert"] = state ? state.clusterCaCert : undefined;
            inputs["connections"] = state ? state.connections : undefined;
            inputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            inputs["installCloudMonitor"] = state ? state.installCloudMonitor : undefined;
            inputs["isEnterpriseSecurityGroup"] = state ? state.isEnterpriseSecurityGroup : undefined;
            inputs["keyName"] = state ? state.keyName : undefined;
            inputs["kubeConfig"] = state ? state.kubeConfig : undefined;
            inputs["logConfig"] = state ? state.logConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["natGatewayId"] = state ? state.natGatewayId : undefined;
            inputs["newNatGateway"] = state ? state.newNatGateway : undefined;
            inputs["nodeCidrMask"] = state ? state.nodeCidrMask : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["podCidr"] = state ? state.podCidr : undefined;
            inputs["proxyMode"] = state ? state.proxyMode : undefined;
            inputs["rdsInstances"] = state ? state.rdsInstances : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            inputs["serviceCidr"] = state ? state.serviceCidr : undefined;
            inputs["slbInternet"] = state ? state.slbInternet : undefined;
            inputs["slbInternetEnabled"] = state ? state.slbInternetEnabled : undefined;
            inputs["slbIntranet"] = state ? state.slbIntranet : undefined;
            inputs["userData"] = state ? state.userData : undefined;
            inputs["version"] = state ? state.version : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
            inputs["workerDataDisks"] = state ? state.workerDataDisks : undefined;
            inputs["workerDiskCategory"] = state ? state.workerDiskCategory : undefined;
            inputs["workerDiskSize"] = state ? state.workerDiskSize : undefined;
            inputs["workerInstanceChargeType"] = state ? state.workerInstanceChargeType : undefined;
            inputs["workerInstanceTypes"] = state ? state.workerInstanceTypes : undefined;
            inputs["workerNodes"] = state ? state.workerNodes : undefined;
            inputs["workerNumber"] = state ? state.workerNumber : undefined;
            inputs["workerVswitchIds"] = state ? state.workerVswitchIds : undefined;
        } else {
            const args = argsOrState as EdgeKubernetesArgs | undefined;
            if (!args || args.workerInstanceTypes === undefined) {
                throw new Error("Missing required property 'workerInstanceTypes'");
            }
            if (!args || args.workerNumber === undefined) {
                throw new Error("Missing required property 'workerNumber'");
            }
            if (!args || args.workerVswitchIds === undefined) {
                throw new Error("Missing required property 'workerVswitchIds'");
            }
            inputs["addons"] = args ? args.addons : undefined;
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["clientCert"] = args ? args.clientCert : undefined;
            inputs["clientKey"] = args ? args.clientKey : undefined;
            inputs["clusterCaCert"] = args ? args.clusterCaCert : undefined;
            inputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            inputs["installCloudMonitor"] = args ? args.installCloudMonitor : undefined;
            inputs["isEnterpriseSecurityGroup"] = args ? args.isEnterpriseSecurityGroup : undefined;
            inputs["keyName"] = args ? args.keyName : undefined;
            inputs["kubeConfig"] = args ? args.kubeConfig : undefined;
            inputs["logConfig"] = args ? args.logConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["newNatGateway"] = args ? args.newNatGateway : undefined;
            inputs["nodeCidrMask"] = args ? args.nodeCidrMask : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["podCidr"] = args ? args.podCidr : undefined;
            inputs["proxyMode"] = args ? args.proxyMode : undefined;
            inputs["rdsInstances"] = args ? args.rdsInstances : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            inputs["serviceCidr"] = args ? args.serviceCidr : undefined;
            inputs["slbInternetEnabled"] = args ? args.slbInternetEnabled : undefined;
            inputs["userData"] = args ? args.userData : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["workerDataDisks"] = args ? args.workerDataDisks : undefined;
            inputs["workerDiskCategory"] = args ? args.workerDiskCategory : undefined;
            inputs["workerDiskSize"] = args ? args.workerDiskSize : undefined;
            inputs["workerInstanceChargeType"] = args ? args.workerInstanceChargeType : undefined;
            inputs["workerInstanceTypes"] = args ? args.workerInstanceTypes : undefined;
            inputs["workerNumber"] = args ? args.workerNumber : undefined;
            inputs["workerVswitchIds"] = args ? args.workerVswitchIds : undefined;
            inputs["connections"] = undefined /*out*/;
            inputs["natGatewayId"] = undefined /*out*/;
            inputs["slbInternet"] = undefined /*out*/;
            inputs["slbIntranet"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
            inputs["workerNodes"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EdgeKubernetes.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EdgeKubernetes resources.
 */
export interface EdgeKubernetesState {
    readonly addons?: pulumi.Input<pulumi.Input<inputs.cs.EdgeKubernetesAddon>[]>;
    /**
     * The ID of availability zone.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * The path of client certificate, like `~/.kube/client-cert.pem`.
     */
    readonly clientCert?: pulumi.Input<string>;
    /**
     * The path of client key, like `~/.kube/client-key.pem`.
     */
    readonly clientKey?: pulumi.Input<string>;
    /**
     * The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
     */
    readonly clusterCaCert?: pulumi.Input<string>;
    readonly connections?: pulumi.Input<inputs.cs.EdgeKubernetesConnections>;
    /**
     * Whether to enable cluster deletion protection.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * Install cloud monitor agent on ECS. default: `true`.
     */
    readonly installCloudMonitor?: pulumi.Input<boolean>;
    /**
     * Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
     */
    readonly isEnterpriseSecurityGroup?: pulumi.Input<boolean>;
    /**
     * The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
     */
    readonly keyName?: pulumi.Input<string>;
    /**
     * The path of kube config, like `~/.kube/config`.
     */
    readonly kubeConfig?: pulumi.Input<string>;
    readonly logConfig?: pulumi.Input<inputs.cs.EdgeKubernetesLogConfig>;
    /**
     * The kubernetes cluster's name. It is unique in one Alicloud account.
     */
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * The ID of nat gateway used to launch kubernetes cluster.
     */
    readonly natGatewayId?: pulumi.Input<string>;
    /**
     * Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
     */
    readonly newNatGateway?: pulumi.Input<boolean>;
    /**
     * The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
     */
    readonly nodeCidrMask?: pulumi.Input<number>;
    /**
     * The password of ssh login cluster node. You have to specify one of `password`, `keyName` `kmsEncryptedPassword` fields.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * [Flannel Specific] The CIDR block for the pod network when using Flannel.
     */
    readonly podCidr?: pulumi.Input<string>;
    /**
     * Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
     */
    readonly proxyMode?: pulumi.Input<string>;
    readonly rdsInstances?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
     */
    readonly serviceCidr?: pulumi.Input<string>;
    readonly slbInternet?: pulumi.Input<string>;
    /**
     * Whether to create internet load balancer for API Server. Default to true.
     */
    readonly slbInternetEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of private load balancer where the current cluster master node is located.
     */
    readonly slbIntranet?: pulumi.Input<string>;
    /**
     * Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
     */
    readonly userData?: pulumi.Input<string>;
    /**
     * Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * The ID of VPC where the current cluster is located.
     */
    readonly vpcId?: pulumi.Input<string>;
    /**
     * The data disk configurations of worker nodes, such as the disk type and disk size.
     * * `category`: the type of the data disks. Valid values:
     * * cloud: basic disks.
     * * cloud_efficiency: ultra disks.
     * * cloud_ssd: SSDs.
     * * `size`: the size of a data disk. Unit: GiB.
     * * `encrypted`: specifies whether to encrypt data disks. Valid values: true and false.
     */
    readonly workerDataDisks?: pulumi.Input<pulumi.Input<inputs.cs.EdgeKubernetesWorkerDataDisk>[]>;
    /**
     * The system disk category of worker node. Its valid value are `cloudEfficiency`, `cloudSsd` and `cloudEssd` and . Default to `cloudEfficiency`.
     */
    readonly workerDiskCategory?: pulumi.Input<string>;
    /**
     * The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
     */
    readonly workerDiskSize?: pulumi.Input<number>;
    readonly workerInstanceChargeType?: pulumi.Input<string>;
    /**
     * The instance types of worker node, you can set multiple types to avoid NoStock of a certain type
     */
    readonly workerInstanceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of cluster worker nodes. It contains several attributes to `Block Nodes`.
     */
    readonly workerNodes?: pulumi.Input<pulumi.Input<inputs.cs.EdgeKubernetesWorkerNode>[]>;
    /**
     * The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
     */
    readonly workerNumber?: pulumi.Input<number>;
    readonly workerVswitchIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a EdgeKubernetes resource.
 */
export interface EdgeKubernetesArgs {
    readonly addons?: pulumi.Input<pulumi.Input<inputs.cs.EdgeKubernetesAddon>[]>;
    /**
     * The ID of availability zone.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * The path of client certificate, like `~/.kube/client-cert.pem`.
     */
    readonly clientCert?: pulumi.Input<string>;
    /**
     * The path of client key, like `~/.kube/client-key.pem`.
     */
    readonly clientKey?: pulumi.Input<string>;
    /**
     * The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
     */
    readonly clusterCaCert?: pulumi.Input<string>;
    /**
     * Whether to enable cluster deletion protection.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * Install cloud monitor agent on ECS. default: `true`.
     */
    readonly installCloudMonitor?: pulumi.Input<boolean>;
    /**
     * Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
     */
    readonly isEnterpriseSecurityGroup?: pulumi.Input<boolean>;
    /**
     * The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
     */
    readonly keyName?: pulumi.Input<string>;
    /**
     * The path of kube config, like `~/.kube/config`.
     */
    readonly kubeConfig?: pulumi.Input<string>;
    readonly logConfig?: pulumi.Input<inputs.cs.EdgeKubernetesLogConfig>;
    /**
     * The kubernetes cluster's name. It is unique in one Alicloud account.
     */
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
     */
    readonly newNatGateway?: pulumi.Input<boolean>;
    /**
     * The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
     */
    readonly nodeCidrMask?: pulumi.Input<number>;
    /**
     * The password of ssh login cluster node. You have to specify one of `password`, `keyName` `kmsEncryptedPassword` fields.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * [Flannel Specific] The CIDR block for the pod network when using Flannel.
     */
    readonly podCidr?: pulumi.Input<string>;
    /**
     * Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
     */
    readonly proxyMode?: pulumi.Input<string>;
    readonly rdsInstances?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
     */
    readonly serviceCidr?: pulumi.Input<string>;
    /**
     * Whether to create internet load balancer for API Server. Default to true.
     */
    readonly slbInternetEnabled?: pulumi.Input<boolean>;
    /**
     * Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
     */
    readonly userData?: pulumi.Input<string>;
    /**
     * Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * The data disk configurations of worker nodes, such as the disk type and disk size.
     * * `category`: the type of the data disks. Valid values:
     * * cloud: basic disks.
     * * cloud_efficiency: ultra disks.
     * * cloud_ssd: SSDs.
     * * `size`: the size of a data disk. Unit: GiB.
     * * `encrypted`: specifies whether to encrypt data disks. Valid values: true and false.
     */
    readonly workerDataDisks?: pulumi.Input<pulumi.Input<inputs.cs.EdgeKubernetesWorkerDataDisk>[]>;
    /**
     * The system disk category of worker node. Its valid value are `cloudEfficiency`, `cloudSsd` and `cloudEssd` and . Default to `cloudEfficiency`.
     */
    readonly workerDiskCategory?: pulumi.Input<string>;
    /**
     * The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
     */
    readonly workerDiskSize?: pulumi.Input<number>;
    readonly workerInstanceChargeType?: pulumi.Input<string>;
    /**
     * The instance types of worker node, you can set multiple types to avoid NoStock of a certain type
     */
    readonly workerInstanceTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
     */
    readonly workerNumber: pulumi.Input<number>;
    readonly workerVswitchIds: pulumi.Input<pulumi.Input<string>[]>;
}
