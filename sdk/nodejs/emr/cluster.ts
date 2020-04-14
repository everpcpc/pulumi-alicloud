// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a EMR Cluster resource. With this you can create, read, and release  EMR Cluster. 
 * 
 * > **NOTE:** Available in 1.57.0+.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/emr_cluster.html.markdown.
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:emr/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    public readonly bootstrapActions!: pulumi.Output<outputs.emr.ClusterBootstrapAction[] | undefined>;
    /**
     * Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
     */
    public readonly chargeType!: pulumi.Output<string | undefined>;
    /**
     * EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
     */
    public readonly clusterType!: pulumi.Output<string>;
    public readonly depositType!: pulumi.Output<string | undefined>;
    public readonly easEnable!: pulumi.Output<boolean | undefined>;
    /**
     * EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
     */
    public readonly emrVer!: pulumi.Output<string>;
    public readonly highAvailabilityEnable!: pulumi.Output<boolean | undefined>;
    /**
     * Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
     */
    public readonly hostGroups!: pulumi.Output<outputs.emr.ClusterHostGroup[] | undefined>;
    public readonly isOpenPublicIp!: pulumi.Output<boolean | undefined>;
    public readonly keyPairName!: pulumi.Output<string | undefined>;
    public readonly masterPwd!: pulumi.Output<string | undefined>;
    /**
     * bootstrap action name.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly optionSoftwareLists!: pulumi.Output<string[] | undefined>;
    public readonly relatedClusterId!: pulumi.Output<string | undefined>;
    public readonly securityGroupId!: pulumi.Output<string | undefined>;
    public readonly sshEnable!: pulumi.Output<boolean | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly useLocalMetadb!: pulumi.Output<boolean | undefined>;
    public readonly userDefinedEmrEcsRole!: pulumi.Output<string | undefined>;
    public readonly vswitchId!: pulumi.Output<string | undefined>;
    /**
     * Zone ID, e.g. cn-huhehaote-a
     * * `securityGroupId` (Optional, ForceNew) Security Group ID for Cluster, you can also specify this key for each host group.
     * * `vswitchId` (Optional, ForceNew) Global vswitch id, you can also specify it in host group.
     * * `optionSoftwareList` (Optional, ForceNew) Optional software list.
     * * `highAvailabilityEnable` (Optional, ForceNew) High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
     * * `useLocalMetadb` (Optional, ForceNew) Use local metadb. Default is false.
     * * `sshEnable` (Optional, ForceNew) If this is set true, we can ssh into cluster. Default value is false.
     * * `masterPwd` (Optional, ForceNew) Master ssh password.
     * * `easEnable` (Optional, ForceNew) High security cluster (true) or not. Default value is false.
     * * `userDefinedEmrEcsRole` (Optional, ForceNew) Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
     * * `keyPairName` (Optional, ForceNew) Ssh key pair.
     * * `depositType` (Optional, ForceNew) Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
     * * `relatedClusterId` (Optional, ForceNew) This specify the related cluster id, if this cluster is a Gateway.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterState | undefined;
            inputs["bootstrapActions"] = state ? state.bootstrapActions : undefined;
            inputs["chargeType"] = state ? state.chargeType : undefined;
            inputs["clusterType"] = state ? state.clusterType : undefined;
            inputs["depositType"] = state ? state.depositType : undefined;
            inputs["easEnable"] = state ? state.easEnable : undefined;
            inputs["emrVer"] = state ? state.emrVer : undefined;
            inputs["highAvailabilityEnable"] = state ? state.highAvailabilityEnable : undefined;
            inputs["hostGroups"] = state ? state.hostGroups : undefined;
            inputs["isOpenPublicIp"] = state ? state.isOpenPublicIp : undefined;
            inputs["keyPairName"] = state ? state.keyPairName : undefined;
            inputs["masterPwd"] = state ? state.masterPwd : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["optionSoftwareLists"] = state ? state.optionSoftwareLists : undefined;
            inputs["relatedClusterId"] = state ? state.relatedClusterId : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            inputs["sshEnable"] = state ? state.sshEnable : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["useLocalMetadb"] = state ? state.useLocalMetadb : undefined;
            inputs["userDefinedEmrEcsRole"] = state ? state.userDefinedEmrEcsRole : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if (!args || args.clusterType === undefined) {
                throw new Error("Missing required property 'clusterType'");
            }
            if (!args || args.emrVer === undefined) {
                throw new Error("Missing required property 'emrVer'");
            }
            if (!args || args.zoneId === undefined) {
                throw new Error("Missing required property 'zoneId'");
            }
            inputs["bootstrapActions"] = args ? args.bootstrapActions : undefined;
            inputs["chargeType"] = args ? args.chargeType : undefined;
            inputs["clusterType"] = args ? args.clusterType : undefined;
            inputs["depositType"] = args ? args.depositType : undefined;
            inputs["easEnable"] = args ? args.easEnable : undefined;
            inputs["emrVer"] = args ? args.emrVer : undefined;
            inputs["highAvailabilityEnable"] = args ? args.highAvailabilityEnable : undefined;
            inputs["hostGroups"] = args ? args.hostGroups : undefined;
            inputs["isOpenPublicIp"] = args ? args.isOpenPublicIp : undefined;
            inputs["keyPairName"] = args ? args.keyPairName : undefined;
            inputs["masterPwd"] = args ? args.masterPwd : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["optionSoftwareLists"] = args ? args.optionSoftwareLists : undefined;
            inputs["relatedClusterId"] = args ? args.relatedClusterId : undefined;
            inputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            inputs["sshEnable"] = args ? args.sshEnable : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["useLocalMetadb"] = args ? args.useLocalMetadb : undefined;
            inputs["userDefinedEmrEcsRole"] = args ? args.userDefinedEmrEcsRole : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    readonly bootstrapActions?: pulumi.Input<pulumi.Input<inputs.emr.ClusterBootstrapAction>[]>;
    /**
     * Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
     */
    readonly chargeType?: pulumi.Input<string>;
    /**
     * EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
     */
    readonly clusterType?: pulumi.Input<string>;
    readonly depositType?: pulumi.Input<string>;
    readonly easEnable?: pulumi.Input<boolean>;
    /**
     * EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
     */
    readonly emrVer?: pulumi.Input<string>;
    readonly highAvailabilityEnable?: pulumi.Input<boolean>;
    /**
     * Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
     */
    readonly hostGroups?: pulumi.Input<pulumi.Input<inputs.emr.ClusterHostGroup>[]>;
    readonly isOpenPublicIp?: pulumi.Input<boolean>;
    readonly keyPairName?: pulumi.Input<string>;
    readonly masterPwd?: pulumi.Input<string>;
    /**
     * bootstrap action name.
     */
    readonly name?: pulumi.Input<string>;
    readonly optionSoftwareLists?: pulumi.Input<pulumi.Input<string>[]>;
    readonly relatedClusterId?: pulumi.Input<string>;
    readonly securityGroupId?: pulumi.Input<string>;
    readonly sshEnable?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly useLocalMetadb?: pulumi.Input<boolean>;
    readonly userDefinedEmrEcsRole?: pulumi.Input<string>;
    readonly vswitchId?: pulumi.Input<string>;
    /**
     * Zone ID, e.g. cn-huhehaote-a
     * * `securityGroupId` (Optional, ForceNew) Security Group ID for Cluster, you can also specify this key for each host group.
     * * `vswitchId` (Optional, ForceNew) Global vswitch id, you can also specify it in host group.
     * * `optionSoftwareList` (Optional, ForceNew) Optional software list.
     * * `highAvailabilityEnable` (Optional, ForceNew) High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
     * * `useLocalMetadb` (Optional, ForceNew) Use local metadb. Default is false.
     * * `sshEnable` (Optional, ForceNew) If this is set true, we can ssh into cluster. Default value is false.
     * * `masterPwd` (Optional, ForceNew) Master ssh password.
     * * `easEnable` (Optional, ForceNew) High security cluster (true) or not. Default value is false.
     * * `userDefinedEmrEcsRole` (Optional, ForceNew) Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
     * * `keyPairName` (Optional, ForceNew) Ssh key pair.
     * * `depositType` (Optional, ForceNew) Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
     * * `relatedClusterId` (Optional, ForceNew) This specify the related cluster id, if this cluster is a Gateway.
     */
    readonly zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    readonly bootstrapActions?: pulumi.Input<pulumi.Input<inputs.emr.ClusterBootstrapAction>[]>;
    /**
     * Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
     */
    readonly chargeType?: pulumi.Input<string>;
    /**
     * EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
     */
    readonly clusterType: pulumi.Input<string>;
    readonly depositType?: pulumi.Input<string>;
    readonly easEnable?: pulumi.Input<boolean>;
    /**
     * EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
     */
    readonly emrVer: pulumi.Input<string>;
    readonly highAvailabilityEnable?: pulumi.Input<boolean>;
    /**
     * Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
     */
    readonly hostGroups?: pulumi.Input<pulumi.Input<inputs.emr.ClusterHostGroup>[]>;
    readonly isOpenPublicIp?: pulumi.Input<boolean>;
    readonly keyPairName?: pulumi.Input<string>;
    readonly masterPwd?: pulumi.Input<string>;
    /**
     * bootstrap action name.
     */
    readonly name?: pulumi.Input<string>;
    readonly optionSoftwareLists?: pulumi.Input<pulumi.Input<string>[]>;
    readonly relatedClusterId?: pulumi.Input<string>;
    readonly securityGroupId?: pulumi.Input<string>;
    readonly sshEnable?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly useLocalMetadb?: pulumi.Input<boolean>;
    readonly userDefinedEmrEcsRole?: pulumi.Input<string>;
    readonly vswitchId?: pulumi.Input<string>;
    /**
     * Zone ID, e.g. cn-huhehaote-a
     * * `securityGroupId` (Optional, ForceNew) Security Group ID for Cluster, you can also specify this key for each host group.
     * * `vswitchId` (Optional, ForceNew) Global vswitch id, you can also specify it in host group.
     * * `optionSoftwareList` (Optional, ForceNew) Optional software list.
     * * `highAvailabilityEnable` (Optional, ForceNew) High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
     * * `useLocalMetadb` (Optional, ForceNew) Use local metadb. Default is false.
     * * `sshEnable` (Optional, ForceNew) If this is set true, we can ssh into cluster. Default value is false.
     * * `masterPwd` (Optional, ForceNew) Master ssh password.
     * * `easEnable` (Optional, ForceNew) High security cluster (true) or not. Default value is false.
     * * `userDefinedEmrEcsRole` (Optional, ForceNew) Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
     * * `keyPairName` (Optional, ForceNew) Ssh key pair.
     * * `depositType` (Optional, ForceNew) Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
     * * `relatedClusterId` (Optional, ForceNew) This specify the related cluster id, if this cluster is a Gateway.
     */
    readonly zoneId: pulumi.Input<string>;
}
