// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an ALIKAFKA topic resource.
 *
 * > **NOTE:** Available in 1.56.0+
 *
 * > **NOTE:**  Only the following regions support create alikafka topic.
 * [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`cn-chengdu`,`cn-heyuan`,`ap-southeast-1`,`ap-southeast-3`,`ap-southeast-5`,`ap-south-1`,`ap-northeast-1`,`eu-central-1`,`eu-west-1`,`us-west-1`,`us-east-1`]
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {cidrBlock: "172.16.0.0/12"});
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?[0]?.id),
 * });
 * const defaultInstance = new alicloud.alikafka.Instance("defaultInstance", {
 *     topicQuota: "50",
 *     diskType: "1",
 *     diskSize: "500",
 *     deployType: "5",
 *     ioMax: "20",
 *     vswitchId: defaultSwitch.id,
 * });
 * const config = new pulumi.Config();
 * const topic = config.get("topic") || "alikafkaTopicName";
 * const defaultTopic = new alicloud.alikafka.Topic("defaultTopic", {
 *     instanceId: defaultInstance.id,
 *     topic: topic,
 *     localTopic: "false",
 *     compactTopic: "false",
 *     partitionNum: "12",
 *     remark: "dafault_kafka_topic_remark",
 * });
 * ```
 *
 * ## Import
 *
 * ### Timeouts The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions* `create` - (Defaults to 10 mins) Used when creating the topic (until it reaches the initial `Running` status).
 */
export class Topic extends pulumi.CustomResource {
    /**
     * Get an existing Topic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TopicState, opts?: pulumi.CustomResourceOptions): Topic {
        return new Topic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:alikafka/topic:Topic';

    /**
     * Returns true if the given object is an instance of Topic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Topic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Topic.__pulumiType;
    }

    /**
     * Whether the topic is compactTopic or not. Compact topic must be a localTopic.
     */
    public readonly compactTopic!: pulumi.Output<boolean | undefined>;
    /**
     * InstanceId of your Kafka resource, the topic will create in this instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Whether the topic is localTopic or not.
     */
    public readonly localTopic!: pulumi.Output<boolean | undefined>;
    /**
     * The number of partitions of the topic. The number should between 1 and 48.
     */
    public readonly partitionNum!: pulumi.Output<number | undefined>;
    /**
     * This attribute is a concise description of topic. The length cannot exceed 64.
     */
    public readonly remark!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Name of the topic. Two topics on a single instance cannot have the same name. The length cannot exceed 64 characters.
     */
    public readonly topic!: pulumi.Output<string>;

    /**
     * Create a Topic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicArgs | TopicState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TopicState | undefined;
            inputs["compactTopic"] = state ? state.compactTopic : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["localTopic"] = state ? state.localTopic : undefined;
            inputs["partitionNum"] = state ? state.partitionNum : undefined;
            inputs["remark"] = state ? state.remark : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["topic"] = state ? state.topic : undefined;
        } else {
            const args = argsOrState as TopicArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.remark === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remark'");
            }
            if ((!args || args.topic === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topic'");
            }
            inputs["compactTopic"] = args ? args.compactTopic : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["localTopic"] = args ? args.localTopic : undefined;
            inputs["partitionNum"] = args ? args.partitionNum : undefined;
            inputs["remark"] = args ? args.remark : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["topic"] = args ? args.topic : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Topic.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Topic resources.
 */
export interface TopicState {
    /**
     * Whether the topic is compactTopic or not. Compact topic must be a localTopic.
     */
    compactTopic?: pulumi.Input<boolean>;
    /**
     * InstanceId of your Kafka resource, the topic will create in this instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Whether the topic is localTopic or not.
     */
    localTopic?: pulumi.Input<boolean>;
    /**
     * The number of partitions of the topic. The number should between 1 and 48.
     */
    partitionNum?: pulumi.Input<number>;
    /**
     * This attribute is a concise description of topic. The length cannot exceed 64.
     */
    remark?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Name of the topic. Two topics on a single instance cannot have the same name. The length cannot exceed 64 characters.
     */
    topic?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    /**
     * Whether the topic is compactTopic or not. Compact topic must be a localTopic.
     */
    compactTopic?: pulumi.Input<boolean>;
    /**
     * InstanceId of your Kafka resource, the topic will create in this instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Whether the topic is localTopic or not.
     */
    localTopic?: pulumi.Input<boolean>;
    /**
     * The number of partitions of the topic. The number should between 1 and 48.
     */
    partitionNum?: pulumi.Input<number>;
    /**
     * This attribute is a concise description of topic. The length cannot exceed 64.
     */
    remark: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Name of the topic. Two topics on a single instance cannot have the same name. The length cannot exceed 64 characters.
     */
    topic: pulumi.Input<string>;
}
