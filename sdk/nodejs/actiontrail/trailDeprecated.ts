// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
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
    /**
     * The RAM role in ActionTrail permitted by the user.
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TrailDeprecatedState | undefined;
            inputs["eventRw"] = state ? state.eventRw : undefined;
            inputs["isOrganizationTrail"] = state ? state.isOrganizationTrail : undefined;
            inputs["mnsTopicArn"] = state ? state.mnsTopicArn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ossBucketName"] = state ? state.ossBucketName : undefined;
            inputs["ossKeyPrefix"] = state ? state.ossKeyPrefix : undefined;
            inputs["roleName"] = state ? state.roleName : undefined;
            inputs["slsProjectArn"] = state ? state.slsProjectArn : undefined;
            inputs["slsWriteRoleArn"] = state ? state.slsWriteRoleArn : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["trailName"] = state ? state.trailName : undefined;
            inputs["trailRegion"] = state ? state.trailRegion : undefined;
        } else {
            const args = argsOrState as TrailDeprecatedArgs | undefined;
            inputs["eventRw"] = args ? args.eventRw : undefined;
            inputs["isOrganizationTrail"] = args ? args.isOrganizationTrail : undefined;
            inputs["mnsTopicArn"] = args ? args.mnsTopicArn : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["ossBucketName"] = args ? args.ossBucketName : undefined;
            inputs["ossKeyPrefix"] = args ? args.ossKeyPrefix : undefined;
            inputs["roleName"] = args ? args.roleName : undefined;
            inputs["slsProjectArn"] = args ? args.slsProjectArn : undefined;
            inputs["slsWriteRoleArn"] = args ? args.slsWriteRoleArn : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["trailName"] = args ? args.trailName : undefined;
            inputs["trailRegion"] = args ? args.trailRegion : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TrailDeprecated.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TrailDeprecated resources.
 */
export interface TrailDeprecatedState {
    /**
     * Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
     */
    readonly eventRw?: pulumi.Input<string>;
    readonly isOrganizationTrail?: pulumi.Input<boolean>;
    readonly mnsTopicArn?: pulumi.Input<string>;
    /**
     * The name of the trail to be created, which must be unique for an account.
     *
     * @deprecated Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
     */
    readonly ossBucketName?: pulumi.Input<string>;
    /**
     * The prefix of the specified OSS bucket name. This parameter can be left empty.
     */
    readonly ossKeyPrefix?: pulumi.Input<string>;
    /**
     * The RAM role in ActionTrail permitted by the user.
     */
    readonly roleName?: pulumi.Input<string>;
    /**
     * The unique ARN of the Log Service project.
     */
    readonly slsProjectArn?: pulumi.Input<string>;
    /**
     * The unique ARN of the Log Service role.
     */
    readonly slsWriteRoleArn?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    readonly trailName?: pulumi.Input<string>;
    readonly trailRegion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TrailDeprecated resource.
 */
export interface TrailDeprecatedArgs {
    /**
     * Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
     */
    readonly eventRw?: pulumi.Input<string>;
    readonly isOrganizationTrail?: pulumi.Input<boolean>;
    readonly mnsTopicArn?: pulumi.Input<string>;
    /**
     * The name of the trail to be created, which must be unique for an account.
     *
     * @deprecated Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
     */
    readonly ossBucketName?: pulumi.Input<string>;
    /**
     * The prefix of the specified OSS bucket name. This parameter can be left empty.
     */
    readonly ossKeyPrefix?: pulumi.Input<string>;
    /**
     * The RAM role in ActionTrail permitted by the user.
     */
    readonly roleName?: pulumi.Input<string>;
    /**
     * The unique ARN of the Log Service project.
     */
    readonly slsProjectArn?: pulumi.Input<string>;
    /**
     * The unique ARN of the Log Service role.
     */
    readonly slsWriteRoleArn?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    readonly trailName?: pulumi.Input<string>;
    readonly trailRegion?: pulumi.Input<string>;
}
