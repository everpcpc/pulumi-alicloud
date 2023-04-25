// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ESS notification resource. More about Ess notification, see [Autoscaling Notification](https://www.alibabacloud.com/help/doc-detail/71114.htm).
 *
 * > **NOTE:** Available in 1.55.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-testAccEssNotification-%d";
 * const defaultRegions = alicloud.getRegions({
 *     current: true,
 * });
 * const defaultAccount = alicloud.getAccount({});
 * const defaultZones = alicloud.getZones({
 *     availableDiskCategory: "cloud_efficiency",
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     vswitchName: name,
 * });
 * const defaultScalingGroup = new alicloud.ess.ScalingGroup("defaultScalingGroup", {
 *     minSize: 1,
 *     maxSize: 1,
 *     scalingGroupName: name,
 *     removalPolicies: [
 *         "OldestInstance",
 *         "NewestInstance",
 *     ],
 *     vswitchIds: [defaultSwitch.id],
 * });
 * const defaultQueue = new alicloud.mns.Queue("defaultQueue", {});
 * const defaultNotification = new alicloud.ess.Notification("defaultNotification", {
 *     scalingGroupId: defaultScalingGroup.id,
 *     notificationTypes: [
 *         "AUTOSCALING:SCALE_OUT_SUCCESS",
 *         "AUTOSCALING:SCALE_OUT_ERROR",
 *     ],
 *     notificationArn: pulumi.all([defaultRegions, defaultAccount, defaultQueue.name]).apply(([defaultRegions, defaultAccount, name]) => `acs:ess:${defaultRegions.regions?.[0]?.id}:${defaultAccount.id}:queue/${name}`),
 * });
 * ```
 *
 * ## Import
 *
 * Ess notification can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ess/notification:Notification example 'scaling_group_id:notification_arn'
 * ```
 */
export class Notification extends pulumi.CustomResource {
    /**
     * Get an existing Notification resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotificationState, opts?: pulumi.CustomResourceOptions): Notification {
        return new Notification(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ess/notification:Notification';

    /**
     * Returns true if the given object is an instance of Notification.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Notification {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Notification.__pulumiType;
    }

    /**
     * The Alibaba Cloud Resource Name (ARN) of the notification object, The value must be in `acs:ess:{region}:{account-id}:{resource-relative-id}` format.
     * * region: the region ID of the scaling group. For more information, see `Regions and zones`
     * * account-id: the ID of your account.
     * * resource-relative-id: the notification method. Valid values : `cloudmonitor`, MNS queue: `queue/{queuename}`, Replace the queuename with the specific MNS queue name, MNS topic: `topic/{topicname}`, Replace the topicname with the specific MNS topic name.
     */
    public readonly notificationArn!: pulumi.Output<string>;
    /**
     * The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'.
     */
    public readonly notificationTypes!: pulumi.Output<string[]>;
    /**
     * The ID of the Auto Scaling group.
     */
    public readonly scalingGroupId!: pulumi.Output<string>;

    /**
     * Create a Notification resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotificationArgs | NotificationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NotificationState | undefined;
            resourceInputs["notificationArn"] = state ? state.notificationArn : undefined;
            resourceInputs["notificationTypes"] = state ? state.notificationTypes : undefined;
            resourceInputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
        } else {
            const args = argsOrState as NotificationArgs | undefined;
            if ((!args || args.notificationArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notificationArn'");
            }
            if ((!args || args.notificationTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notificationTypes'");
            }
            if ((!args || args.scalingGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            resourceInputs["notificationArn"] = args ? args.notificationArn : undefined;
            resourceInputs["notificationTypes"] = args ? args.notificationTypes : undefined;
            resourceInputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Notification.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Notification resources.
 */
export interface NotificationState {
    /**
     * The Alibaba Cloud Resource Name (ARN) of the notification object, The value must be in `acs:ess:{region}:{account-id}:{resource-relative-id}` format.
     * * region: the region ID of the scaling group. For more information, see `Regions and zones`
     * * account-id: the ID of your account.
     * * resource-relative-id: the notification method. Valid values : `cloudmonitor`, MNS queue: `queue/{queuename}`, Replace the queuename with the specific MNS queue name, MNS topic: `topic/{topicname}`, Replace the topicname with the specific MNS topic name.
     */
    notificationArn?: pulumi.Input<string>;
    /**
     * The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'.
     */
    notificationTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the Auto Scaling group.
     */
    scalingGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Notification resource.
 */
export interface NotificationArgs {
    /**
     * The Alibaba Cloud Resource Name (ARN) of the notification object, The value must be in `acs:ess:{region}:{account-id}:{resource-relative-id}` format.
     * * region: the region ID of the scaling group. For more information, see `Regions and zones`
     * * account-id: the ID of your account.
     * * resource-relative-id: the notification method. Valid values : `cloudmonitor`, MNS queue: `queue/{queuename}`, Replace the queuename with the specific MNS queue name, MNS topic: `topic/{topicname}`, Replace the topicname with the specific MNS topic name.
     */
    notificationArn: pulumi.Input<string>;
    /**
     * The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'.
     */
    notificationTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the Auto Scaling group.
     */
    scalingGroupId: pulumi.Input<string>;
}
