// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ActionTrail Trail resource. For information about alicloud actiontrail trail and how to use it, see [What is Resource Alicloud ActionTrail Trail](https://www.alibabacloud.com/help/doc-detail/28804.htm).
 *
 * > **NOTE:** Available in 1.95.0+
 *
 * > **NOTE:** You can create a trail to deliver events to Log Service, Object Storage Service (OSS), or both. Before you call this operation to create a trail, make sure that the following requirements are met.
 * - Deliver events to Log Service: A project is created in Log Service.
 * - Deliver events to OSS: A bucket is created in OSS.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Create a new actiontrail trail.
 * const defaultTrail = new alicloud.actiontrail.Trail("default", {
 *     eventRw: "All",
 *     ossBucketName: "bucket_name",
 *     ossWriteRoleArn: "acs:ram::1182725xxxxxxxxxxx",
 *     trailName: "action-trail",
 *     trailRegion: "cn-hangzhou",
 * });
 * ```
 *
 * ## Import
 *
 * Action trail can be imported using the id or trail_name, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:actiontrail/trail:Trail default abc12345678
 * ```
 */
export class Trail extends pulumi.CustomResource {
    /**
     * Get an existing Trail resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TrailState, opts?: pulumi.CustomResourceOptions): Trail {
        return new Trail(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:actiontrail/trail:Trail';

    /**
     * Returns true if the given object is an instance of Trail.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Trail {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Trail.__pulumiType;
    }

    /**
     * Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
     */
    public readonly eventRw!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to create a multi-account trail. Valid values:`true`: Create a multi-account trail.`false`: Create a single-account trail. It is the default value.
     */
    public readonly isOrganizationTrail!: pulumi.Output<boolean | undefined>;
    /**
     * Field `mnsTopicArn` has been deprecated from version 1.118.0.
     *
     * @deprecated Field 'mns_topic_arn' has been deprecated from version 1.118.0
     */
    public readonly mnsTopicArn!: pulumi.Output<string | undefined>;
    /**
     * Field `name` has been deprecated from version 1.95.0. Use `trailName` instead.
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
     * The unique ARN of the Oss role.
     */
    public readonly ossWriteRoleArn!: pulumi.Output<string | undefined>;
    /**
     * Field `name` has been deprecated from version 1.118.0.
     *
     * @deprecated Field 'role_name' has been deprecated from version 1.118.0
     */
    public readonly roleName!: pulumi.Output<string>;
    /**
     * The unique ARN of the Log Service project. Ensure that `slsProjectArn` is valid .
     */
    public readonly slsProjectArn!: pulumi.Output<string | undefined>;
    /**
     * The unique ARN of the Log Service role.
     */
    public readonly slsWriteRoleArn!: pulumi.Output<string>;
    /**
     * The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * The name of the trail to be created, which must be unique for an account.
     */
    public readonly trailName!: pulumi.Output<string>;
    /**
     * The regions to which the trail is applied. Valid values: `cn-beijing`, `cn-hangzhou`, and `All`. Default to `All`.
     */
    public readonly trailRegion!: pulumi.Output<string | undefined>;

    /**
     * Create a Trail resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TrailArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TrailArgs | TrailState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TrailState | undefined;
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
            const args = argsOrState as TrailArgs | undefined;
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
        super(Trail.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Trail resources.
 */
export interface TrailState {
    /**
     * Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
     */
    eventRw?: pulumi.Input<string>;
    /**
     * Specifies whether to create a multi-account trail. Valid values:`true`: Create a multi-account trail.`false`: Create a single-account trail. It is the default value.
     */
    isOrganizationTrail?: pulumi.Input<boolean>;
    /**
     * Field `mnsTopicArn` has been deprecated from version 1.118.0.
     *
     * @deprecated Field 'mns_topic_arn' has been deprecated from version 1.118.0
     */
    mnsTopicArn?: pulumi.Input<string>;
    /**
     * Field `name` has been deprecated from version 1.95.0. Use `trailName` instead.
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
    /**
     * The unique ARN of the Oss role.
     */
    ossWriteRoleArn?: pulumi.Input<string>;
    /**
     * Field `name` has been deprecated from version 1.118.0.
     *
     * @deprecated Field 'role_name' has been deprecated from version 1.118.0
     */
    roleName?: pulumi.Input<string>;
    /**
     * The unique ARN of the Log Service project. Ensure that `slsProjectArn` is valid .
     */
    slsProjectArn?: pulumi.Input<string>;
    /**
     * The unique ARN of the Log Service role.
     */
    slsWriteRoleArn?: pulumi.Input<string>;
    /**
     * The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * The name of the trail to be created, which must be unique for an account.
     */
    trailName?: pulumi.Input<string>;
    /**
     * The regions to which the trail is applied. Valid values: `cn-beijing`, `cn-hangzhou`, and `All`. Default to `All`.
     */
    trailRegion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Trail resource.
 */
export interface TrailArgs {
    /**
     * Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
     */
    eventRw?: pulumi.Input<string>;
    /**
     * Specifies whether to create a multi-account trail. Valid values:`true`: Create a multi-account trail.`false`: Create a single-account trail. It is the default value.
     */
    isOrganizationTrail?: pulumi.Input<boolean>;
    /**
     * Field `mnsTopicArn` has been deprecated from version 1.118.0.
     *
     * @deprecated Field 'mns_topic_arn' has been deprecated from version 1.118.0
     */
    mnsTopicArn?: pulumi.Input<string>;
    /**
     * Field `name` has been deprecated from version 1.95.0. Use `trailName` instead.
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
    /**
     * The unique ARN of the Oss role.
     */
    ossWriteRoleArn?: pulumi.Input<string>;
    /**
     * Field `name` has been deprecated from version 1.118.0.
     *
     * @deprecated Field 'role_name' has been deprecated from version 1.118.0
     */
    roleName?: pulumi.Input<string>;
    /**
     * The unique ARN of the Log Service project. Ensure that `slsProjectArn` is valid .
     */
    slsProjectArn?: pulumi.Input<string>;
    /**
     * The unique ARN of the Log Service role.
     */
    slsWriteRoleArn?: pulumi.Input<string>;
    /**
     * The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * The name of the trail to be created, which must be unique for an account.
     */
    trailName?: pulumi.Input<string>;
    /**
     * The regions to which the trail is applied. Valid values: `cn-beijing`, `cn-hangzhou`, and `All`. Default to `All`.
     */
    trailRegion?: pulumi.Input<string>;
}
