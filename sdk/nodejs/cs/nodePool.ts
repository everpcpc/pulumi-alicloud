// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Cluster nodepool can be imported using the id, e.g. Then complete the nodepool.tf accords to the result of `terraform plan`.
 *
 * ```sh
 *  $ pulumi import alicloud:cs/nodePool:NodePool alicloud_cs_node_pool.custom_nodepool cluster_id:nodepool_id
 * ```
 */
export class NodePool extends pulumi.CustomResource {
    /**
     * Get an existing NodePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodePoolState, opts?: pulumi.CustomResourceOptions): NodePool {
        return new NodePool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cs/nodePool:NodePool';

    /**
     * Returns true if the given object is an instance of NodePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodePool.__pulumiType;
    }

    /**
     * Enable Node payment auto-renew, default is `false`.
     */
    public readonly autoRenew!: pulumi.Output<boolean | undefined>;
    /**
     * Node payment auto-renew period, one of `1`, `2`, `3`,`6`, `12`.
     */
    public readonly autoRenewPeriod!: pulumi.Output<number | undefined>;
    /**
     * The id of kubernetes cluster.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The data disk configurations of worker nodes, such as the disk type and disk size.
     */
    public readonly dataDisks!: pulumi.Output<outputs.cs.NodePoolDataDisk[] | undefined>;
    /**
     * After you select this check box, if data disks have been attached to the specified ECS instances and the file system of the last data disk is uninitialized, the system automatically formats the last data disk to ext4 and mounts the data disk to /var/lib/docker and /var/lib/kubelet. The original data on the disk will be cleared. Make sure that you back up data in advance. If no data disk is mounted on the ECS instance, no new data disk will be purchased. Default is `false`.
     */
    public readonly formatDisk!: pulumi.Output<boolean>;
    /**
     * Custom Image support. Must based on CentOS7 or AliyunLinux2.
     */
    public readonly imageId!: pulumi.Output<string>;
    /**
     * The image type, instead of `platform`. This field cannot be modified. One of `AliyunLinux`, `AliyunLinux3`, `AliyunLinux3Arm64`, `AliyunLinuxUEFI`, `CentOS`, `Windows`,`WindowsCore`,`AliyunLinux Qboot`,`ContainerOS`. If you select `Windows` or `WindowsCore`, the `passord` is required.
     */
    public readonly imageType!: pulumi.Output<string>;
    /**
     * Install the cloud monitoring plug-in on the node, and you can view the monitoring information of the instance through the cloud monitoring console. Default is `true`.
     */
    public readonly installCloudMonitor!: pulumi.Output<boolean | undefined>;
    /**
     * Node payment type. Valid values: `PostPaid`, `PrePaid`, default is `PostPaid`. If value is `PrePaid`, the arguments `period`, `periodUnit`, `autoRenew` and `autoRenewPeriod` are required.
     */
    public readonly instanceChargeType!: pulumi.Output<string | undefined>;
    /**
     * The instance type of worker node.
     */
    public readonly instanceTypes!: pulumi.Output<string[]>;
    /**
     * The instance list. Add existing nodes under the same cluster VPC to the node pool.
     */
    public readonly instances!: pulumi.Output<string[] | undefined>;
    /**
     * The billing method for network usage. Valid values `PayByBandwidth` and `PayByTraffic`. Conflict with `eipInternetChargeType`, EIP and public network IP can only choose one.
     */
    public readonly internetChargeType!: pulumi.Output<string>;
    /**
     * The maximum outbound bandwidth for the public network. Unit: Mbit/s. Valid values: 0 to 100.
     */
    public readonly internetMaxBandwidthOut!: pulumi.Output<number>;
    /**
     * Add an existing instance to the node pool, whether to keep the original instance name. It is recommended to set to `true`.
     */
    public readonly keepInstanceName!: pulumi.Output<boolean>;
    /**
     * The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields. Only `keyName` is supported in the management node pool.
     */
    public readonly keyName!: pulumi.Output<string | undefined>;
    /**
     * An KMS encrypts password used to a cs kubernetes. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
     */
    public readonly kmsEncryptedPassword!: pulumi.Output<string | undefined>;
    /**
     * A List of Kubernetes labels to assign to the nodes . Only labels that are applied with the ACK API are managed by this argument.
     */
    public readonly labels!: pulumi.Output<outputs.cs.NodePoolLabel[] | undefined>;
    /**
     * Managed node pool configuration. When using a managed node pool, the node key must use `keyName`. Detailed below.
     */
    public readonly management!: pulumi.Output<outputs.cs.NodePoolManagement>;
    /**
     * The name of node pool.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The worker node number of the node pool. From version 1.111.0, `nodeCount` is not required.
     */
    public readonly nodeCount!: pulumi.Output<number>;
    /**
     * Each node name consists of a prefix, an IP substring, and a suffix. For example "customized,aliyun.com,5,test", if the node IP address is 192.168.0.55, the prefix is aliyun.com, IP substring length is 5, and the suffix is test, the node name will be aliyun.com00055test.
     */
    public readonly nodeNameMode!: pulumi.Output<string>;
    /**
     * The password of ssh login cluster node. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Node payment period. Its valid value is one of {1, 2, 3, 6, 12, 24, 36, 48, 60}.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Node payment period unit, valid value: `Month`. Default is `Month`.
     */
    public readonly periodUnit!: pulumi.Output<string | undefined>;
    /**
     * The platform. One of `AliyunLinux`, `Windows`, `CentOS`, `WindowsCore`. If you select `Windows` or `WindowsCore`, the `passord` is required. Field `platform` has been deprecated from provider version 1.145.0. New field `imageType` instead.
     *
     * @deprecated Field 'platform' has been deprecated from provider version 1.145.0. New field 'image_type' instead
     */
    public readonly platform!: pulumi.Output<string>;
    /**
     * The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The runtime name of containers. If not set, the cluster runtime will be used as the node pool runtime. If you select another container runtime, see [Comparison of Docker, containerd, and Sandboxed-Container](https://www.alibabacloud.com/help/doc-detail/160313.htm).
     */
    public readonly runtimeName!: pulumi.Output<string>;
    /**
     * The runtime version of containers. If not set, the cluster runtime will be used as the node pool runtime.
     */
    public readonly runtimeVersion!: pulumi.Output<string>;
    /**
     * Auto scaling node pool configuration. For more details, see `scalingConfig`. With auto-scaling is enabled, the nodes in the node pool will be labeled with `k8s.aliyun.com=true` to prevent system pods such as coredns, metrics-servers from being scheduled to elastic nodes, and to prevent node shrinkage from causing business abnormalities.
     */
    public readonly scalingConfig!: pulumi.Output<outputs.cs.NodePoolScalingConfig>;
    /**
     * (Available in 1.105.0+) Id of the Scaling Group.
     */
    public /*out*/ readonly scalingGroupId!: pulumi.Output<string>;
    /**
     * The scaling mode. Valid values: `release`, `recycle`, default is `release`. Standard mode(release): Create and release ECS instances based on requests.Swift mode(recycle): Create, stop, and restart ECS instances based on needs. New ECS instances are only created when no stopped ECS instance is avalible. This mode further accelerates the scaling process. Apart from ECS instances that use local storage, when an ECS instance is stopped, you are only chatged for storage space.
     */
    public readonly scalingPolicy!: pulumi.Output<string>;
    /**
     * The security group id for worker node. Field `securityGroupId` has been deprecated from provider version 1.145.0. New field `securityGroupIds` instead.
     *
     * @deprecated Field 'security_group_id' has been deprecated from provider version 1.145.0. New field 'security_group_ids' instead
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * Multiple security groups can be configured for a node pool. If both `securityGroupIds` and `securityGroupId` are configured, `securityGroupIds` takes effect. This field cannot be modified.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * The maximum hourly price of the instance. This parameter takes effect only when `spotStrategy` is set to `SpotWithPriceLimit`. A maximum of three decimal places are allowed.
     */
    public readonly spotPriceLimits!: pulumi.Output<outputs.cs.NodePoolSpotPriceLimit[]>;
    /**
     * The preemption policy for the pay-as-you-go instance. This parameter takes effect only when `instanceChargeType` is set to `PostPaid`. Valid value `SpotWithPriceLimit`.
     */
    public readonly spotStrategy!: pulumi.Output<string>;
    /**
     * The system disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`. Default to `cloudEfficiency`.
     */
    public readonly systemDiskCategory!: pulumi.Output<string | undefined>;
    public readonly systemDiskPerformanceLevel!: pulumi.Output<string | undefined>;
    /**
     * The system disk category of worker node. Its valid value range [40~500] in GB. Default to `120`.
     */
    public readonly systemDiskSize!: pulumi.Output<number | undefined>;
    /**
     * A Map of tags to assign to the resource. It will be applied for ECS instances finally.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * A List of Kubernetes taints to assign to the nodes.
     */
    public readonly taints!: pulumi.Output<outputs.cs.NodePoolTaint[] | undefined>;
    /**
     * Set the newly added node as unschedulable. If you want to open the scheduling option, you can open it in the node list of the console. If you are using an auto-scaling node pool, the setting will not take effect. Default is `false`.
     */
    public readonly unschedulable!: pulumi.Output<boolean | undefined>;
    /**
     * Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
     */
    public readonly userData!: pulumi.Output<string | undefined>;
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * The vswitches used by node pool workers.
     */
    public readonly vswitchIds!: pulumi.Output<string[]>;

    /**
     * Create a NodePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodePoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodePoolArgs | NodePoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NodePoolState | undefined;
            inputs["autoRenew"] = state ? state.autoRenew : undefined;
            inputs["autoRenewPeriod"] = state ? state.autoRenewPeriod : undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["dataDisks"] = state ? state.dataDisks : undefined;
            inputs["formatDisk"] = state ? state.formatDisk : undefined;
            inputs["imageId"] = state ? state.imageId : undefined;
            inputs["imageType"] = state ? state.imageType : undefined;
            inputs["installCloudMonitor"] = state ? state.installCloudMonitor : undefined;
            inputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            inputs["instanceTypes"] = state ? state.instanceTypes : undefined;
            inputs["instances"] = state ? state.instances : undefined;
            inputs["internetChargeType"] = state ? state.internetChargeType : undefined;
            inputs["internetMaxBandwidthOut"] = state ? state.internetMaxBandwidthOut : undefined;
            inputs["keepInstanceName"] = state ? state.keepInstanceName : undefined;
            inputs["keyName"] = state ? state.keyName : undefined;
            inputs["kmsEncryptedPassword"] = state ? state.kmsEncryptedPassword : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["management"] = state ? state.management : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nodeCount"] = state ? state.nodeCount : undefined;
            inputs["nodeNameMode"] = state ? state.nodeNameMode : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["periodUnit"] = state ? state.periodUnit : undefined;
            inputs["platform"] = state ? state.platform : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["runtimeName"] = state ? state.runtimeName : undefined;
            inputs["runtimeVersion"] = state ? state.runtimeVersion : undefined;
            inputs["scalingConfig"] = state ? state.scalingConfig : undefined;
            inputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
            inputs["scalingPolicy"] = state ? state.scalingPolicy : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["spotPriceLimits"] = state ? state.spotPriceLimits : undefined;
            inputs["spotStrategy"] = state ? state.spotStrategy : undefined;
            inputs["systemDiskCategory"] = state ? state.systemDiskCategory : undefined;
            inputs["systemDiskPerformanceLevel"] = state ? state.systemDiskPerformanceLevel : undefined;
            inputs["systemDiskSize"] = state ? state.systemDiskSize : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["taints"] = state ? state.taints : undefined;
            inputs["unschedulable"] = state ? state.unschedulable : undefined;
            inputs["userData"] = state ? state.userData : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
            inputs["vswitchIds"] = state ? state.vswitchIds : undefined;
        } else {
            const args = argsOrState as NodePoolArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.instanceTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceTypes'");
            }
            if ((!args || args.vswitchIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchIds'");
            }
            inputs["autoRenew"] = args ? args.autoRenew : undefined;
            inputs["autoRenewPeriod"] = args ? args.autoRenewPeriod : undefined;
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["dataDisks"] = args ? args.dataDisks : undefined;
            inputs["formatDisk"] = args ? args.formatDisk : undefined;
            inputs["imageId"] = args ? args.imageId : undefined;
            inputs["imageType"] = args ? args.imageType : undefined;
            inputs["installCloudMonitor"] = args ? args.installCloudMonitor : undefined;
            inputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            inputs["instanceTypes"] = args ? args.instanceTypes : undefined;
            inputs["instances"] = args ? args.instances : undefined;
            inputs["internetChargeType"] = args ? args.internetChargeType : undefined;
            inputs["internetMaxBandwidthOut"] = args ? args.internetMaxBandwidthOut : undefined;
            inputs["keepInstanceName"] = args ? args.keepInstanceName : undefined;
            inputs["keyName"] = args ? args.keyName : undefined;
            inputs["kmsEncryptedPassword"] = args ? args.kmsEncryptedPassword : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["management"] = args ? args.management : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeCount"] = args ? args.nodeCount : undefined;
            inputs["nodeNameMode"] = args ? args.nodeNameMode : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["periodUnit"] = args ? args.periodUnit : undefined;
            inputs["platform"] = args ? args.platform : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["runtimeName"] = args ? args.runtimeName : undefined;
            inputs["runtimeVersion"] = args ? args.runtimeVersion : undefined;
            inputs["scalingConfig"] = args ? args.scalingConfig : undefined;
            inputs["scalingPolicy"] = args ? args.scalingPolicy : undefined;
            inputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["spotPriceLimits"] = args ? args.spotPriceLimits : undefined;
            inputs["spotStrategy"] = args ? args.spotStrategy : undefined;
            inputs["systemDiskCategory"] = args ? args.systemDiskCategory : undefined;
            inputs["systemDiskPerformanceLevel"] = args ? args.systemDiskPerformanceLevel : undefined;
            inputs["systemDiskSize"] = args ? args.systemDiskSize : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["taints"] = args ? args.taints : undefined;
            inputs["unschedulable"] = args ? args.unschedulable : undefined;
            inputs["userData"] = args ? args.userData : undefined;
            inputs["vswitchIds"] = args ? args.vswitchIds : undefined;
            inputs["scalingGroupId"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NodePool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NodePool resources.
 */
export interface NodePoolState {
    /**
     * Enable Node payment auto-renew, default is `false`.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * Node payment auto-renew period, one of `1`, `2`, `3`,`6`, `12`.
     */
    autoRenewPeriod?: pulumi.Input<number>;
    /**
     * The id of kubernetes cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The data disk configurations of worker nodes, such as the disk type and disk size.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.cs.NodePoolDataDisk>[]>;
    /**
     * After you select this check box, if data disks have been attached to the specified ECS instances and the file system of the last data disk is uninitialized, the system automatically formats the last data disk to ext4 and mounts the data disk to /var/lib/docker and /var/lib/kubelet. The original data on the disk will be cleared. Make sure that you back up data in advance. If no data disk is mounted on the ECS instance, no new data disk will be purchased. Default is `false`.
     */
    formatDisk?: pulumi.Input<boolean>;
    /**
     * Custom Image support. Must based on CentOS7 or AliyunLinux2.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The image type, instead of `platform`. This field cannot be modified. One of `AliyunLinux`, `AliyunLinux3`, `AliyunLinux3Arm64`, `AliyunLinuxUEFI`, `CentOS`, `Windows`,`WindowsCore`,`AliyunLinux Qboot`,`ContainerOS`. If you select `Windows` or `WindowsCore`, the `passord` is required.
     */
    imageType?: pulumi.Input<string>;
    /**
     * Install the cloud monitoring plug-in on the node, and you can view the monitoring information of the instance through the cloud monitoring console. Default is `true`.
     */
    installCloudMonitor?: pulumi.Input<boolean>;
    /**
     * Node payment type. Valid values: `PostPaid`, `PrePaid`, default is `PostPaid`. If value is `PrePaid`, the arguments `period`, `periodUnit`, `autoRenew` and `autoRenewPeriod` are required.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The instance type of worker node.
     */
    instanceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The instance list. Add existing nodes under the same cluster VPC to the node pool.
     */
    instances?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The billing method for network usage. Valid values `PayByBandwidth` and `PayByTraffic`. Conflict with `eipInternetChargeType`, EIP and public network IP can only choose one.
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * The maximum outbound bandwidth for the public network. Unit: Mbit/s. Valid values: 0 to 100.
     */
    internetMaxBandwidthOut?: pulumi.Input<number>;
    /**
     * Add an existing instance to the node pool, whether to keep the original instance name. It is recommended to set to `true`.
     */
    keepInstanceName?: pulumi.Input<boolean>;
    /**
     * The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields. Only `keyName` is supported in the management node pool.
     */
    keyName?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a cs kubernetes. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
     */
    kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * A List of Kubernetes labels to assign to the nodes . Only labels that are applied with the ACK API are managed by this argument.
     */
    labels?: pulumi.Input<pulumi.Input<inputs.cs.NodePoolLabel>[]>;
    /**
     * Managed node pool configuration. When using a managed node pool, the node key must use `keyName`. Detailed below.
     */
    management?: pulumi.Input<inputs.cs.NodePoolManagement>;
    /**
     * The name of node pool.
     */
    name?: pulumi.Input<string>;
    /**
     * The worker node number of the node pool. From version 1.111.0, `nodeCount` is not required.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * Each node name consists of a prefix, an IP substring, and a suffix. For example "customized,aliyun.com,5,test", if the node IP address is 192.168.0.55, the prefix is aliyun.com, IP substring length is 5, and the suffix is test, the node name will be aliyun.com00055test.
     */
    nodeNameMode?: pulumi.Input<string>;
    /**
     * The password of ssh login cluster node. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
     */
    password?: pulumi.Input<string>;
    /**
     * Node payment period. Its valid value is one of {1, 2, 3, 6, 12, 24, 36, 48, 60}.
     */
    period?: pulumi.Input<number>;
    /**
     * Node payment period unit, valid value: `Month`. Default is `Month`.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The platform. One of `AliyunLinux`, `Windows`, `CentOS`, `WindowsCore`. If you select `Windows` or `WindowsCore`, the `passord` is required. Field `platform` has been deprecated from provider version 1.145.0. New field `imageType` instead.
     *
     * @deprecated Field 'platform' has been deprecated from provider version 1.145.0. New field 'image_type' instead
     */
    platform?: pulumi.Input<string>;
    /**
     * The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The runtime name of containers. If not set, the cluster runtime will be used as the node pool runtime. If you select another container runtime, see [Comparison of Docker, containerd, and Sandboxed-Container](https://www.alibabacloud.com/help/doc-detail/160313.htm).
     */
    runtimeName?: pulumi.Input<string>;
    /**
     * The runtime version of containers. If not set, the cluster runtime will be used as the node pool runtime.
     */
    runtimeVersion?: pulumi.Input<string>;
    /**
     * Auto scaling node pool configuration. For more details, see `scalingConfig`. With auto-scaling is enabled, the nodes in the node pool will be labeled with `k8s.aliyun.com=true` to prevent system pods such as coredns, metrics-servers from being scheduled to elastic nodes, and to prevent node shrinkage from causing business abnormalities.
     */
    scalingConfig?: pulumi.Input<inputs.cs.NodePoolScalingConfig>;
    /**
     * (Available in 1.105.0+) Id of the Scaling Group.
     */
    scalingGroupId?: pulumi.Input<string>;
    /**
     * The scaling mode. Valid values: `release`, `recycle`, default is `release`. Standard mode(release): Create and release ECS instances based on requests.Swift mode(recycle): Create, stop, and restart ECS instances based on needs. New ECS instances are only created when no stopped ECS instance is avalible. This mode further accelerates the scaling process. Apart from ECS instances that use local storage, when an ECS instance is stopped, you are only chatged for storage space.
     */
    scalingPolicy?: pulumi.Input<string>;
    /**
     * The security group id for worker node. Field `securityGroupId` has been deprecated from provider version 1.145.0. New field `securityGroupIds` instead.
     *
     * @deprecated Field 'security_group_id' has been deprecated from provider version 1.145.0. New field 'security_group_ids' instead
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * Multiple security groups can be configured for a node pool. If both `securityGroupIds` and `securityGroupId` are configured, `securityGroupIds` takes effect. This field cannot be modified.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum hourly price of the instance. This parameter takes effect only when `spotStrategy` is set to `SpotWithPriceLimit`. A maximum of three decimal places are allowed.
     */
    spotPriceLimits?: pulumi.Input<pulumi.Input<inputs.cs.NodePoolSpotPriceLimit>[]>;
    /**
     * The preemption policy for the pay-as-you-go instance. This parameter takes effect only when `instanceChargeType` is set to `PostPaid`. Valid value `SpotWithPriceLimit`.
     */
    spotStrategy?: pulumi.Input<string>;
    /**
     * The system disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`. Default to `cloudEfficiency`.
     */
    systemDiskCategory?: pulumi.Input<string>;
    systemDiskPerformanceLevel?: pulumi.Input<string>;
    /**
     * The system disk category of worker node. Its valid value range [40~500] in GB. Default to `120`.
     */
    systemDiskSize?: pulumi.Input<number>;
    /**
     * A Map of tags to assign to the resource. It will be applied for ECS instances finally.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * A List of Kubernetes taints to assign to the nodes.
     */
    taints?: pulumi.Input<pulumi.Input<inputs.cs.NodePoolTaint>[]>;
    /**
     * Set the newly added node as unschedulable. If you want to open the scheduling option, you can open it in the node list of the console. If you are using an auto-scaling node pool, the setting will not take effect. Default is `false`.
     */
    unschedulable?: pulumi.Input<boolean>;
    /**
     * Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
     */
    userData?: pulumi.Input<string>;
    vpcId?: pulumi.Input<string>;
    /**
     * The vswitches used by node pool workers.
     */
    vswitchIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a NodePool resource.
 */
export interface NodePoolArgs {
    /**
     * Enable Node payment auto-renew, default is `false`.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * Node payment auto-renew period, one of `1`, `2`, `3`,`6`, `12`.
     */
    autoRenewPeriod?: pulumi.Input<number>;
    /**
     * The id of kubernetes cluster.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The data disk configurations of worker nodes, such as the disk type and disk size.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.cs.NodePoolDataDisk>[]>;
    /**
     * After you select this check box, if data disks have been attached to the specified ECS instances and the file system of the last data disk is uninitialized, the system automatically formats the last data disk to ext4 and mounts the data disk to /var/lib/docker and /var/lib/kubelet. The original data on the disk will be cleared. Make sure that you back up data in advance. If no data disk is mounted on the ECS instance, no new data disk will be purchased. Default is `false`.
     */
    formatDisk?: pulumi.Input<boolean>;
    /**
     * Custom Image support. Must based on CentOS7 or AliyunLinux2.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The image type, instead of `platform`. This field cannot be modified. One of `AliyunLinux`, `AliyunLinux3`, `AliyunLinux3Arm64`, `AliyunLinuxUEFI`, `CentOS`, `Windows`,`WindowsCore`,`AliyunLinux Qboot`,`ContainerOS`. If you select `Windows` or `WindowsCore`, the `passord` is required.
     */
    imageType?: pulumi.Input<string>;
    /**
     * Install the cloud monitoring plug-in on the node, and you can view the monitoring information of the instance through the cloud monitoring console. Default is `true`.
     */
    installCloudMonitor?: pulumi.Input<boolean>;
    /**
     * Node payment type. Valid values: `PostPaid`, `PrePaid`, default is `PostPaid`. If value is `PrePaid`, the arguments `period`, `periodUnit`, `autoRenew` and `autoRenewPeriod` are required.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The instance type of worker node.
     */
    instanceTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The instance list. Add existing nodes under the same cluster VPC to the node pool.
     */
    instances?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The billing method for network usage. Valid values `PayByBandwidth` and `PayByTraffic`. Conflict with `eipInternetChargeType`, EIP and public network IP can only choose one.
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * The maximum outbound bandwidth for the public network. Unit: Mbit/s. Valid values: 0 to 100.
     */
    internetMaxBandwidthOut?: pulumi.Input<number>;
    /**
     * Add an existing instance to the node pool, whether to keep the original instance name. It is recommended to set to `true`.
     */
    keepInstanceName?: pulumi.Input<boolean>;
    /**
     * The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields. Only `keyName` is supported in the management node pool.
     */
    keyName?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a cs kubernetes. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
     */
    kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * A List of Kubernetes labels to assign to the nodes . Only labels that are applied with the ACK API are managed by this argument.
     */
    labels?: pulumi.Input<pulumi.Input<inputs.cs.NodePoolLabel>[]>;
    /**
     * Managed node pool configuration. When using a managed node pool, the node key must use `keyName`. Detailed below.
     */
    management?: pulumi.Input<inputs.cs.NodePoolManagement>;
    /**
     * The name of node pool.
     */
    name?: pulumi.Input<string>;
    /**
     * The worker node number of the node pool. From version 1.111.0, `nodeCount` is not required.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * Each node name consists of a prefix, an IP substring, and a suffix. For example "customized,aliyun.com,5,test", if the node IP address is 192.168.0.55, the prefix is aliyun.com, IP substring length is 5, and the suffix is test, the node name will be aliyun.com00055test.
     */
    nodeNameMode?: pulumi.Input<string>;
    /**
     * The password of ssh login cluster node. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
     */
    password?: pulumi.Input<string>;
    /**
     * Node payment period. Its valid value is one of {1, 2, 3, 6, 12, 24, 36, 48, 60}.
     */
    period?: pulumi.Input<number>;
    /**
     * Node payment period unit, valid value: `Month`. Default is `Month`.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The platform. One of `AliyunLinux`, `Windows`, `CentOS`, `WindowsCore`. If you select `Windows` or `WindowsCore`, the `passord` is required. Field `platform` has been deprecated from provider version 1.145.0. New field `imageType` instead.
     *
     * @deprecated Field 'platform' has been deprecated from provider version 1.145.0. New field 'image_type' instead
     */
    platform?: pulumi.Input<string>;
    /**
     * The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The runtime name of containers. If not set, the cluster runtime will be used as the node pool runtime. If you select another container runtime, see [Comparison of Docker, containerd, and Sandboxed-Container](https://www.alibabacloud.com/help/doc-detail/160313.htm).
     */
    runtimeName?: pulumi.Input<string>;
    /**
     * The runtime version of containers. If not set, the cluster runtime will be used as the node pool runtime.
     */
    runtimeVersion?: pulumi.Input<string>;
    /**
     * Auto scaling node pool configuration. For more details, see `scalingConfig`. With auto-scaling is enabled, the nodes in the node pool will be labeled with `k8s.aliyun.com=true` to prevent system pods such as coredns, metrics-servers from being scheduled to elastic nodes, and to prevent node shrinkage from causing business abnormalities.
     */
    scalingConfig?: pulumi.Input<inputs.cs.NodePoolScalingConfig>;
    /**
     * The scaling mode. Valid values: `release`, `recycle`, default is `release`. Standard mode(release): Create and release ECS instances based on requests.Swift mode(recycle): Create, stop, and restart ECS instances based on needs. New ECS instances are only created when no stopped ECS instance is avalible. This mode further accelerates the scaling process. Apart from ECS instances that use local storage, when an ECS instance is stopped, you are only chatged for storage space.
     */
    scalingPolicy?: pulumi.Input<string>;
    /**
     * The security group id for worker node. Field `securityGroupId` has been deprecated from provider version 1.145.0. New field `securityGroupIds` instead.
     *
     * @deprecated Field 'security_group_id' has been deprecated from provider version 1.145.0. New field 'security_group_ids' instead
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * Multiple security groups can be configured for a node pool. If both `securityGroupIds` and `securityGroupId` are configured, `securityGroupIds` takes effect. This field cannot be modified.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum hourly price of the instance. This parameter takes effect only when `spotStrategy` is set to `SpotWithPriceLimit`. A maximum of three decimal places are allowed.
     */
    spotPriceLimits?: pulumi.Input<pulumi.Input<inputs.cs.NodePoolSpotPriceLimit>[]>;
    /**
     * The preemption policy for the pay-as-you-go instance. This parameter takes effect only when `instanceChargeType` is set to `PostPaid`. Valid value `SpotWithPriceLimit`.
     */
    spotStrategy?: pulumi.Input<string>;
    /**
     * The system disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`. Default to `cloudEfficiency`.
     */
    systemDiskCategory?: pulumi.Input<string>;
    systemDiskPerformanceLevel?: pulumi.Input<string>;
    /**
     * The system disk category of worker node. Its valid value range [40~500] in GB. Default to `120`.
     */
    systemDiskSize?: pulumi.Input<number>;
    /**
     * A Map of tags to assign to the resource. It will be applied for ECS instances finally.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * A List of Kubernetes taints to assign to the nodes.
     */
    taints?: pulumi.Input<pulumi.Input<inputs.cs.NodePoolTaint>[]>;
    /**
     * Set the newly added node as unschedulable. If you want to open the scheduling option, you can open it in the node list of the console. If you are using an auto-scaling node pool, the setting will not take effect. Default is `false`.
     */
    unschedulable?: pulumi.Input<boolean>;
    /**
     * Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
     */
    userData?: pulumi.Input<string>;
    /**
     * The vswitches used by node pool workers.
     */
    vswitchIds: pulumi.Input<pulumi.Input<string>[]>;
}
