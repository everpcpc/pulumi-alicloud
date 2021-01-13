// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an ALIKAFKA sasl acl resource.
 *
 * > **NOTE:** Available in 1.66.0+
 *
 * > **NOTE:**  Only the following regions support create alikafka sasl user.
 * [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`ap-southeast-1`,`ap-south-1`,`ap-southeast-5`]
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const username = config.get("username") || "testusername";
 * const password = config.get("password") || "testpassword";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {cidrBlock: "172.16.0.0/12"});
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     availabilityZone: defaultZones.then(defaultZones => defaultZones.zones[0].id),
 * });
 * const defaultInstance = new alicloud.alikafka.Instance("defaultInstance", {
 *     topicQuota: "50",
 *     diskType: "1",
 *     diskSize: "500",
 *     deployType: "5",
 *     ioMax: "20",
 *     vswitchId: defaultSwitch.id,
 * });
 * const defaultTopic = new alicloud.alikafka.Topic("defaultTopic", {
 *     instanceId: defaultInstance.id,
 *     topic: "test-topic",
 *     remark: "topic-remark",
 * });
 * const defaultSaslUser = new alicloud.alikafka.SaslUser("defaultSaslUser", {
 *     instanceId: defaultInstance.id,
 *     username: username,
 *     password: password,
 * });
 * const defaultSaslAcl = new alicloud.alikafka.SaslAcl("defaultSaslAcl", {
 *     instanceId: defaultInstance.id,
 *     username: defaultSaslUser.username,
 *     aclResourceType: "Topic",
 *     aclResourceName: defaultTopic.topic,
 *     aclResourcePatternType: "LITERAL",
 *     aclOperationType: "Write",
 * });
 * ```
 *
 * ## Import
 *
 * ALIKAFKA GROUP can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:alikafka/saslAcl:SaslAcl acl alikafka_post-cn-123455abc:username:Topic:test-topic:LITERAL:Write
 * ```
 */
export class SaslAcl extends pulumi.CustomResource {
    /**
     * Get an existing SaslAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SaslAclState, opts?: pulumi.CustomResourceOptions): SaslAcl {
        return new SaslAcl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:alikafka/saslAcl:SaslAcl';

    /**
     * Returns true if the given object is an instance of SaslAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SaslAcl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SaslAcl.__pulumiType;
    }

    /**
     * Operation type for this acl. The operation type can only be "Write" and "Read".
     */
    public readonly aclOperationType!: pulumi.Output<string>;
    /**
     * Resource name for this acl. The resource name should be a topic or consumer group name.
     */
    public readonly aclResourceName!: pulumi.Output<string>;
    /**
     * Resource pattern type for this acl. The resource pattern support two types "LITERAL" and "PREFIXED". "LITERAL": A literal name defines the full name of a resource. The special wildcard character "*" can be used to represent a resource with any name. "PREFIXED": A prefixed name defines a prefix for a resource.
     */
    public readonly aclResourcePatternType!: pulumi.Output<string>;
    /**
     * Resource type for this acl. The resource type can only be "Topic" and "Group".
     */
    public readonly aclResourceType!: pulumi.Output<string>;
    /**
     * The host of the acl.
     */
    public /*out*/ readonly host!: pulumi.Output<string>;
    /**
     * ID of the ALIKAFKA Instance that owns the groups.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a SaslAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SaslAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SaslAclArgs | SaslAclState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SaslAclState | undefined;
            inputs["aclOperationType"] = state ? state.aclOperationType : undefined;
            inputs["aclResourceName"] = state ? state.aclResourceName : undefined;
            inputs["aclResourcePatternType"] = state ? state.aclResourcePatternType : undefined;
            inputs["aclResourceType"] = state ? state.aclResourceType : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as SaslAclArgs | undefined;
            if ((!args || args.aclOperationType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'aclOperationType'");
            }
            if ((!args || args.aclResourceName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'aclResourceName'");
            }
            if ((!args || args.aclResourcePatternType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'aclResourcePatternType'");
            }
            if ((!args || args.aclResourceType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'aclResourceType'");
            }
            if ((!args || args.instanceId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.username === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'username'");
            }
            inputs["aclOperationType"] = args ? args.aclOperationType : undefined;
            inputs["aclResourceName"] = args ? args.aclResourceName : undefined;
            inputs["aclResourcePatternType"] = args ? args.aclResourcePatternType : undefined;
            inputs["aclResourceType"] = args ? args.aclResourceType : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["username"] = args ? args.username : undefined;
            inputs["host"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SaslAcl.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SaslAcl resources.
 */
export interface SaslAclState {
    /**
     * Operation type for this acl. The operation type can only be "Write" and "Read".
     */
    readonly aclOperationType?: pulumi.Input<string>;
    /**
     * Resource name for this acl. The resource name should be a topic or consumer group name.
     */
    readonly aclResourceName?: pulumi.Input<string>;
    /**
     * Resource pattern type for this acl. The resource pattern support two types "LITERAL" and "PREFIXED". "LITERAL": A literal name defines the full name of a resource. The special wildcard character "*" can be used to represent a resource with any name. "PREFIXED": A prefixed name defines a prefix for a resource.
     */
    readonly aclResourcePatternType?: pulumi.Input<string>;
    /**
     * Resource type for this acl. The resource type can only be "Topic" and "Group".
     */
    readonly aclResourceType?: pulumi.Input<string>;
    /**
     * The host of the acl.
     */
    readonly host?: pulumi.Input<string>;
    /**
     * ID of the ALIKAFKA Instance that owns the groups.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.
     */
    readonly username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SaslAcl resource.
 */
export interface SaslAclArgs {
    /**
     * Operation type for this acl. The operation type can only be "Write" and "Read".
     */
    readonly aclOperationType: pulumi.Input<string>;
    /**
     * Resource name for this acl. The resource name should be a topic or consumer group name.
     */
    readonly aclResourceName: pulumi.Input<string>;
    /**
     * Resource pattern type for this acl. The resource pattern support two types "LITERAL" and "PREFIXED". "LITERAL": A literal name defines the full name of a resource. The special wildcard character "*" can be used to represent a resource with any name. "PREFIXED": A prefixed name defines a prefix for a resource.
     */
    readonly aclResourcePatternType: pulumi.Input<string>;
    /**
     * Resource type for this acl. The resource type can only be "Topic" and "Group".
     */
    readonly aclResourceType: pulumi.Input<string>;
    /**
     * ID of the ALIKAFKA Instance that owns the groups.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.
     */
    readonly username: pulumi.Input<string>;
}
