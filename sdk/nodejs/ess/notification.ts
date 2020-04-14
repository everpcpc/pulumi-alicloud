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
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-testAccEssNotification-%d";
 * 
 * const defaultRegions = alicloud.getRegions({
 *     current: true,
 * });
 * const defaultAccount = alicloud.getAccount();
 * const defaultZones = alicloud.getZones({
 *     availableDiskCategory: "cloudEfficiency",
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     availabilityZone: defaultZones.zones[0].id,
 *     cidrBlock: "172.16.0.0/24",
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultScalingGroup = new alicloud.ess.ScalingGroup("default", {
 *     maxSize: 1,
 *     minSize: 1,
 *     removalPolicies: [
 *         "OldestInstance",
 *         "NewestInstance",
 *     ],
 *     scalingGroupName: name,
 *     vswitchIds: [defaultSwitch.id],
 * });
 * const defaultQueue = new alicloud.mns.Queue("default", {});
 * const defaultNotification = new alicloud.ess.Notification("default", {
 *     notificationArn: pulumi.interpolate`acs:ess:${defaultRegions.regions[0].id}:${defaultAccount.id}:queue/${defaultQueue.name}`,
 *     notificationTypes: [
 *         "AUTOSCALING:SCALE_OUT_SUCCESS",
 *         "AUTOSCALING:SCALE_OUT_ERROR",
 *     ],
 *     scalingGroupId: defaultScalingGroup.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/ess_notification.html.markdown.
 */
export class Notification extends pulumi.CustomResource {
    /**
     * Get an existing Notification resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
     * The Alibaba Cloud Resource Name (ARN) for the notification object. The format of `notificationArn` is acs:ess:{region}:{account-id}:{resource-relative-id}. Valid values for `resource-relative-id`: 'cloudmonitor', 'queue/', 'topic/'.
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NotificationState | undefined;
            inputs["notificationArn"] = state ? state.notificationArn : undefined;
            inputs["notificationTypes"] = state ? state.notificationTypes : undefined;
            inputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
        } else {
            const args = argsOrState as NotificationArgs | undefined;
            if (!args || args.notificationArn === undefined) {
                throw new Error("Missing required property 'notificationArn'");
            }
            if (!args || args.notificationTypes === undefined) {
                throw new Error("Missing required property 'notificationTypes'");
            }
            if (!args || args.scalingGroupId === undefined) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            inputs["notificationArn"] = args ? args.notificationArn : undefined;
            inputs["notificationTypes"] = args ? args.notificationTypes : undefined;
            inputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Notification.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Notification resources.
 */
export interface NotificationState {
    /**
     * The Alibaba Cloud Resource Name (ARN) for the notification object. The format of `notificationArn` is acs:ess:{region}:{account-id}:{resource-relative-id}. Valid values for `resource-relative-id`: 'cloudmonitor', 'queue/', 'topic/'.
     */
    readonly notificationArn?: pulumi.Input<string>;
    /**
     * The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'.
     */
    readonly notificationTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the Auto Scaling group.
     */
    readonly scalingGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Notification resource.
 */
export interface NotificationArgs {
    /**
     * The Alibaba Cloud Resource Name (ARN) for the notification object. The format of `notificationArn` is acs:ess:{region}:{account-id}:{resource-relative-id}. Valid values for `resource-relative-id`: 'cloudmonitor', 'queue/', 'topic/'.
     */
    readonly notificationArn: pulumi.Input<string>;
    /**
     * The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'.
     */
    readonly notificationTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the Auto Scaling group.
     */
    readonly scalingGroupId: pulumi.Input<string>;
}
