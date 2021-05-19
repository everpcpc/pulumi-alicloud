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
     * Custom Image support. Must based on CentOS7 or AliyunLinux2.
     */
    public readonly imageId!: pulumi.Output<string>;
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
    public readonly management!: pulumi.Output<outputs.cs.NodePoolManagement | undefined>;
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
     * Auto scaling node pool configuration. For more details, see `scalingConfig`.
     */
    public readonly scalingConfig!: pulumi.Output<outputs.cs.NodePoolScalingConfig>;
    /**
     * (Available in 1.105.0+) Id of the Scaling Group.
     */
    public /*out*/ readonly scalingGroupId!: pulumi.Output<string>;
    /**
     * The security group id for worker node.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
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
            inputs["imageId"] = state ? state.imageId : undefined;
            inputs["installCloudMonitor"] = state ? state.installCloudMonitor : undefined;
            inputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            inputs["instanceTypes"] = state ? state.instanceTypes : undefined;
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
            inputs["scalingConfig"] = state ? state.scalingConfig : undefined;
            inputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
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
            inputs["imageId"] = args ? args.imageId : undefined;
            inputs["installCloudMonitor"] = args ? args.installCloudMonitor : undefined;
            inputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            inputs["instanceTypes"] = args ? args.instanceTypes : undefined;
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
            inputs["scalingConfig"] = args ? args.scalingConfig : undefined;
            inputs["securityGroupId"] = args ? args.securityGroupId : undefined;
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
    readonly autoRenew?: pulumi.Input<boolean>;
    /**
     * Node payment auto-renew period, one of `1`, `2`, `3`,`6`, `12`.
     */
    readonly autoRenewPeriod?: pulumi.Input<number>;
    /**
     * The id of kubernetes cluster.
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * The data disk configurations of worker nodes, such as the disk type and disk size.
     */
    readonly dataDisks?: pulumi.Input<pulumi.Input<inputs.cs.NodePoolDataDisk>[]>;
    /**
     * Custom Image support. Must based on CentOS7 or AliyunLinux2.
     */
    readonly imageId?: pulumi.Input<string>;
    /**
     * Install the cloud monitoring plug-in on the node, and you can view the monitoring information of the instance through the cloud monitoring console. Default is `true`.
     */
    readonly installCloudMonitor?: pulumi.Input<boolean>;
    /**
     * Node payment type. Valid values: `PostPaid`, `PrePaid`, default is `PostPaid`. If value is `PrePaid`, the arguments `period`, `periodUnit`, `autoRenew` and `autoRenewPeriod` are required.
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * The instance type of worker node.
     */
    readonly instanceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields. Only `keyName` is supported in the management node pool.
     */
    readonly keyName?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a cs kubernetes. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
     */
    readonly kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * A List of Kubernetes labels to assign to the nodes . Only labels that are applied with the ACK API are managed by this argument.
     */
    readonly labels?: pulumi.Input<pulumi.Input<inputs.cs.NodePoolLabel>[]>;
    /**
     * Managed node pool configuration. When using a managed node pool, the node key must use `keyName`. Detailed below.
     */
    readonly management?: pulumi.Input<inputs.cs.NodePoolManagement>;
    /**
     * The name of node pool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The worker node number of the node pool. From version 1.111.0, `nodeCount` is not required.
     */
    readonly nodeCount?: pulumi.Input<number>;
    /**
     * Each node name consists of a prefix, an IP substring, and a suffix. For example "customized,aliyun.com,5,test", if the node IP address is 192.168.0.55, the prefix is aliyun.com, IP substring length is 5, and the suffix is test, the node name will be aliyun.com00055test.
     */
    readonly nodeNameMode?: pulumi.Input<string>;
    /**
     * The password of ssh login cluster node. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Node payment period. Its valid value is one of {1, 2, 3, 6, 12, 24, 36, 48, 60}.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * Node payment period unit, valid value: `Month`. Default is `Month`.
     */
    readonly periodUnit?: pulumi.Input<string>;
    /**
     * Auto scaling node pool configuration. For more details, see `scalingConfig`.
     */
    readonly scalingConfig?: pulumi.Input<inputs.cs.NodePoolScalingConfig>;
    /**
     * (Available in 1.105.0+) Id of the Scaling Group.
     */
    readonly scalingGroupId?: pulumi.Input<string>;
    /**
     * The security group id for worker node.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * The system disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`. Default to `cloudEfficiency`.
     */
    readonly systemDiskCategory?: pulumi.Input<string>;
    readonly systemDiskPerformanceLevel?: pulumi.Input<string>;
    /**
     * The system disk category of worker node. Its valid value range [40~500] in GB. Default to `120`.
     */
    readonly systemDiskSize?: pulumi.Input<number>;
    /**
     * A Map of tags to assign to the resource. It will be applied for ECS instances finally.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * A List of Kubernetes taints to assign to the nodes.
     */
    readonly taints?: pulumi.Input<pulumi.Input<inputs.cs.NodePoolTaint>[]>;
    /**
     * Set the newly added node as unschedulable. If you want to open the scheduling option, you can open it in the node list of the console. If you are using an auto-scaling node pool, the setting will not take effect. Default is `false`.
     */
    readonly unschedulable?: pulumi.Input<boolean>;
    /**
     * Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
     */
    readonly userData?: pulumi.Input<string>;
    readonly vpcId?: pulumi.Input<string>;
    /**
     * The vswitches used by node pool workers.
     */
    readonly vswitchIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a NodePool resource.
 */
export interface NodePoolArgs {
    /**
     * Enable Node payment auto-renew, default is `false`.
     */
    readonly autoRenew?: pulumi.Input<boolean>;
    /**
     * Node payment auto-renew period, one of `1`, `2`, `3`,`6`, `12`.
     */
    readonly autoRenewPeriod?: pulumi.Input<number>;
    /**
     * The id of kubernetes cluster.
     */
    readonly clusterId: pulumi.Input<string>;
    /**
     * The data disk configurations of worker nodes, such as the disk type and disk size.
     */
    readonly dataDisks?: pulumi.Input<pulumi.Input<inputs.cs.NodePoolDataDisk>[]>;
    /**
     * Custom Image support. Must based on CentOS7 or AliyunLinux2.
     */
    readonly imageId?: pulumi.Input<string>;
    /**
     * Install the cloud monitoring plug-in on the node, and you can view the monitoring information of the instance through the cloud monitoring console. Default is `true`.
     */
    readonly installCloudMonitor?: pulumi.Input<boolean>;
    /**
     * Node payment type. Valid values: `PostPaid`, `PrePaid`, default is `PostPaid`. If value is `PrePaid`, the arguments `period`, `periodUnit`, `autoRenew` and `autoRenewPeriod` are required.
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * The instance type of worker node.
     */
    readonly instanceTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields. Only `keyName` is supported in the management node pool.
     */
    readonly keyName?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a cs kubernetes. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
     */
    readonly kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * A List of Kubernetes labels to assign to the nodes . Only labels that are applied with the ACK API are managed by this argument.
     */
    readonly labels?: pulumi.Input<pulumi.Input<inputs.cs.NodePoolLabel>[]>;
    /**
     * Managed node pool configuration. When using a managed node pool, the node key must use `keyName`. Detailed below.
     */
    readonly management?: pulumi.Input<inputs.cs.NodePoolManagement>;
    /**
     * The name of node pool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The worker node number of the node pool. From version 1.111.0, `nodeCount` is not required.
     */
    readonly nodeCount?: pulumi.Input<number>;
    /**
     * Each node name consists of a prefix, an IP substring, and a suffix. For example "customized,aliyun.com,5,test", if the node IP address is 192.168.0.55, the prefix is aliyun.com, IP substring length is 5, and the suffix is test, the node name will be aliyun.com00055test.
     */
    readonly nodeNameMode?: pulumi.Input<string>;
    /**
     * The password of ssh login cluster node. You have to specify one of `password` `keyName` `kmsEncryptedPassword` fields.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Node payment period. Its valid value is one of {1, 2, 3, 6, 12, 24, 36, 48, 60}.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * Node payment period unit, valid value: `Month`. Default is `Month`.
     */
    readonly periodUnit?: pulumi.Input<string>;
    /**
     * Auto scaling node pool configuration. For more details, see `scalingConfig`.
     */
    readonly scalingConfig?: pulumi.Input<inputs.cs.NodePoolScalingConfig>;
    /**
     * The security group id for worker node.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * The system disk category of worker node. Its valid value are `cloudSsd` and `cloudEfficiency`. Default to `cloudEfficiency`.
     */
    readonly systemDiskCategory?: pulumi.Input<string>;
    readonly systemDiskPerformanceLevel?: pulumi.Input<string>;
    /**
     * The system disk category of worker node. Its valid value range [40~500] in GB. Default to `120`.
     */
    readonly systemDiskSize?: pulumi.Input<number>;
    /**
     * A Map of tags to assign to the resource. It will be applied for ECS instances finally.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * A List of Kubernetes taints to assign to the nodes.
     */
    readonly taints?: pulumi.Input<pulumi.Input<inputs.cs.NodePoolTaint>[]>;
    /**
     * Set the newly added node as unschedulable. If you want to open the scheduling option, you can open it in the node list of the console. If you are using an auto-scaling node pool, the setting will not take effect. Default is `false`.
     */
    readonly unschedulable?: pulumi.Input<boolean>;
    /**
     * Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
     */
    readonly userData?: pulumi.Input<string>;
    /**
     * The vswitches used by node pool workers.
     */
    readonly vswitchIds: pulumi.Input<pulumi.Input<string>[]>;
}
