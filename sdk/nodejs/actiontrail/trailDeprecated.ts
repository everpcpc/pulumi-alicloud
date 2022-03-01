// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Create a new action trail.
 * const foo = new alicloud.actiontrail.TrailDeprecated("foo", {
 *     eventRw: "Write-test",
 *     ossBucketName: alicloud_oss_bucket.bucket.id,
 *     roleName: alicloud_ram_role_policy_attachment.attach.role_name,
 *     ossKeyPrefix: "at-product-account-audit-B",
 * });
 * ```
 *
 * ## Import
 *
 * Action trail can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:actiontrail/trailDeprecated:TrailDeprecated foo abc12345678
 * ```
 *
 * @deprecated Resource renamed to `Trail`
 */
export class TrailDeprecated extends pulumi.CustomResource {
    /**
     * Get an existing TrailDeprecated resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TrailDeprecatedState, opts?: pulumi.CustomResourceOptions): TrailDeprecated {
        pulumi.log.warn("TrailDeprecated is deprecated: Resource renamed to `Trail`")
        return new TrailDeprecated(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:actiontrail/trailDeprecated:TrailDeprecated';

    /**
     * Returns true if the given object is an instance of TrailDeprecated.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TrailDeprecated {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TrailDeprecated.__pulumiType;
    }

    /**
     * Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
     */
    public readonly eventRw!: pulumi.Output<string | undefined>;
    public readonly isOrganizationTrail!: pulumi.Output<boolean | undefined>;
    /**
     * @deprecated Field 'mns_topic_arn' has been deprecated from version 1.118.0
     */
    public readonly mnsTopicArn!: pulumi.Output<string | undefined>;
    /**
     * The name of the trail to be created, which must be unique for an account.
     *
     * @deprecated Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
     */
    public readonly ossBucketName!: pulumi.Output<string | undefined>;
    /**
     * The prefix of the specified OSS bucket name. This parameter can be left empty.
     */
    public readonly ossKeyPrefix!: pulumi.Output<string | undefined>;
    public readonly ossWriteRoleArn!: pulumi.Output<string | undefined>;
    /**
     * The RAM role in ActionTrail permitted by the user.
     *
     * @deprecated Field 'role_name' has been deprecated from version 1.118.0
     */
    public readonly roleName!: pulumi.Output<string>;
    /**
     * The unique ARN of the Log Service project.
     */
    public readonly slsProjectArn!: pulumi.Output<string | undefined>;
    /**
     * The unique ARN of the Log Service role.
     */
    public readonly slsWriteRoleArn!: pulumi.Output<string | undefined>;
    public readonly status!: pulumi.Output<string | undefined>;
    public readonly trailName!: pulumi.Output<string>;
    public readonly trailRegion!: pulumi.Output<string | undefined>;

    /**
     * Create a TrailDeprecated resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Resource renamed to `Trail` */
    constructor(name: string, args?: TrailDeprecatedArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated Resource renamed to `Trail` */
    constructor(name: string, argsOrState?: TrailDeprecatedArgs | TrailDeprecatedState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("TrailDeprecated is deprecated: Resource renamed to `Trail`")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TrailDeprecatedState | undefined;
            resourceInputs["eventRw"] = state ? state.eventRw : undefined;
            resourceInputs["isOrganizationTrail"] = state ? state.isOrganizationTrail : undefined;
            resourceInputs["mnsTopicArn"] = state ? state.mnsTopicArn : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ossBucketName"] = state ? state.ossBucketName : undefined;
            resourceInputs["ossKeyPrefix"] = state ? state.ossKeyPrefix : undefined;
            resourceInputs["ossWriteRoleArn"] = state ? state.ossWriteRoleArn : undefined;
            resourceInputs["roleName"] = state ? state.roleName : undefined;
            resourceInputs["slsProjectArn"] = state ? state.slsProjectArn : undefined;
            resourceInputs["slsWriteRoleArn"] = state ? state.slsWriteRoleArn : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["trailName"] = state ? state.trailName : undefined;
            resourceInputs["trailRegion"] = state ? state.trailRegion : undefined;
        } else {
            const args = argsOrState as TrailDeprecatedArgs | undefined;
            resourceInputs["eventRw"] = args ? args.eventRw : undefined;
            resourceInputs["isOrganizationTrail"] = args ? args.isOrganizationTrail : undefined;
            resourceInputs["mnsTopicArn"] = args ? args.mnsTopicArn : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ossBucketName"] = args ? args.ossBucketName : undefined;
            resourceInputs["ossKeyPrefix"] = args ? args.ossKeyPrefix : undefined;
            resourceInputs["ossWriteRoleArn"] = args ? args.ossWriteRoleArn : undefined;
            resourceInputs["roleName"] = args ? args.roleName : undefined;
            resourceInputs["slsProjectArn"] = args ? args.slsProjectArn : undefined;
            resourceInputs["slsWriteRoleArn"] = args ? args.slsWriteRoleArn : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["trailName"] = args ? args.trailName : undefined;
            resourceInputs["trailRegion"] = args ? args.trailRegion : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TrailDeprecated.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TrailDeprecated resources.
 */
export interface TrailDeprecatedState {
    /**
     * Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
     */
    eventRw?: pulumi.Input<string>;
    isOrganizationTrail?: pulumi.Input<boolean>;
    /**
     * @deprecated Field 'mns_topic_arn' has been deprecated from version 1.118.0
     */
    mnsTopicArn?: pulumi.Input<string>;
    /**
     * The name of the trail to be created, which must be unique for an account.
     *
     * @deprecated Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
     */
    ossBucketName?: pulumi.Input<string>;
    /**
     * The prefix of the specified OSS bucket name. This parameter can be left empty.
     */
    ossKeyPrefix?: pulumi.Input<string>;
    ossWriteRoleArn?: pulumi.Input<string>;
    /**
     * The RAM role in ActionTrail permitted by the user.
     *
     * @deprecated Field 'role_name' has been deprecated from version 1.118.0
     */
    roleName?: pulumi.Input<string>;
    /**
     * The unique ARN of the Log Service project.
     */
    slsProjectArn?: pulumi.Input<string>;
    /**
     * The unique ARN of the Log Service role.
     */
    slsWriteRoleArn?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    trailName?: pulumi.Input<string>;
    trailRegion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TrailDeprecated resource.
 */
export interface TrailDeprecatedArgs {
    /**
     * Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
     */
    eventRw?: pulumi.Input<string>;
    isOrganizationTrail?: pulumi.Input<boolean>;
    /**
     * @deprecated Field 'mns_topic_arn' has been deprecated from version 1.118.0
     */
    mnsTopicArn?: pulumi.Input<string>;
    /**
     * The name of the trail to be created, which must be unique for an account.
     *
     * @deprecated Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
     */
    ossBucketName?: pulumi.Input<string>;
    /**
     * The prefix of the specified OSS bucket name. This parameter can be left empty.
     */
    ossKeyPrefix?: pulumi.Input<string>;
    ossWriteRoleArn?: pulumi.Input<string>;
    /**
     * The RAM role in ActionTrail permitted by the user.
     *
     * @deprecated Field 'role_name' has been deprecated from version 1.118.0
     */
    roleName?: pulumi.Input<string>;
    /**
     * The unique ARN of the Log Service project.
     */
    slsProjectArn?: pulumi.Input<string>;
    /**
     * The unique ARN of the Log Service role.
     */
    slsWriteRoleArn?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    trailName?: pulumi.Input<string>;
    trailRegion?: pulumi.Input<string>;
}
