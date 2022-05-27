// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CDN Fc Trigger resource.
 *
 * For information about CDN Fc Trigger and how to use it, see [What is Fc Trigger](https://www.alibabacloud.com/help/zh/alibaba-cloud-cdn/latest/add-function-calculation-trigger).
 *
 * > **NOTE:** Available in v1.165.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultAccount = pulumi.output(alicloud.getAccount());
 * const defaultRegions = pulumi.output(alicloud.getRegions({
 *     current: true,
 * }));
 * const example = new alicloud.cdn.FcTrigger("example", {
 *     eventMetaName: "LogFileCreated",
 *     eventMetaVersion: "1.0.0",
 *     notes: "example_value",
 *     roleArn: pulumi.interpolate`acs:ram::${defaultAccount.id}:role/aliyuncdneventnotificationrole`,
 *     sourceArn: pulumi.interpolate`acs:cdn:*:${defaultAccount.id}:domain/example.com`,
 *     triggerArn: pulumi.interpolate`acs:fc:${defaultRegions.regions[0].id}:${defaultAccount.id}:services/FCTestService/functions/printEvent/triggers/testtrigger`,
 * });
 * ```
 *
 * ## Import
 *
 * CDN Fc Trigger can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cdn/fcTrigger:FcTrigger example <trigger_arn>
 * ```
 */
export class FcTrigger extends pulumi.CustomResource {
    /**
     * Get an existing FcTrigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FcTriggerState, opts?: pulumi.CustomResourceOptions): FcTrigger {
        return new FcTrigger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cdn/fcTrigger:FcTrigger';

    /**
     * Returns true if the given object is an instance of FcTrigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FcTrigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FcTrigger.__pulumiType;
    }

    /**
     * The name of the Event.
     */
    public readonly eventMetaName!: pulumi.Output<string>;
    /**
     * The version of the Event.
     */
    public readonly eventMetaVersion!: pulumi.Output<string>;
    /**
     * The function arn. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`.
     */
    public readonly functionArn!: pulumi.Output<string | undefined>;
    /**
     * The Note information.
     */
    public readonly notes!: pulumi.Output<string>;
    /**
     * The role authorized by RAM. The value formats as `acs:ram::{AccountID}:role/{RoleName}`.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * Resources and filters for event listening. The value formats as `acs:cdn:{RegionID}:{AccountID}:{Filter}`.
     */
    public readonly sourceArn!: pulumi.Output<string>;
    /**
     * The trigger corresponding to the function Compute Service. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`. See [Create a CDN Fc Trigger](https://www.alibabacloud.com/help/zh/alibaba-cloud-cdn/latest/add-function-calculation-trigger) for more details.
     */
    public readonly triggerArn!: pulumi.Output<string>;

    /**
     * Create a FcTrigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FcTriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FcTriggerArgs | FcTriggerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FcTriggerState | undefined;
            resourceInputs["eventMetaName"] = state ? state.eventMetaName : undefined;
            resourceInputs["eventMetaVersion"] = state ? state.eventMetaVersion : undefined;
            resourceInputs["functionArn"] = state ? state.functionArn : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
            resourceInputs["roleArn"] = state ? state.roleArn : undefined;
            resourceInputs["sourceArn"] = state ? state.sourceArn : undefined;
            resourceInputs["triggerArn"] = state ? state.triggerArn : undefined;
        } else {
            const args = argsOrState as FcTriggerArgs | undefined;
            if ((!args || args.eventMetaName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventMetaName'");
            }
            if ((!args || args.eventMetaVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventMetaVersion'");
            }
            if ((!args || args.notes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notes'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            if ((!args || args.sourceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceArn'");
            }
            if ((!args || args.triggerArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'triggerArn'");
            }
            resourceInputs["eventMetaName"] = args ? args.eventMetaName : undefined;
            resourceInputs["eventMetaVersion"] = args ? args.eventMetaVersion : undefined;
            resourceInputs["functionArn"] = args ? args.functionArn : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["sourceArn"] = args ? args.sourceArn : undefined;
            resourceInputs["triggerArn"] = args ? args.triggerArn : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FcTrigger.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FcTrigger resources.
 */
export interface FcTriggerState {
    /**
     * The name of the Event.
     */
    eventMetaName?: pulumi.Input<string>;
    /**
     * The version of the Event.
     */
    eventMetaVersion?: pulumi.Input<string>;
    /**
     * The function arn. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`.
     */
    functionArn?: pulumi.Input<string>;
    /**
     * The Note information.
     */
    notes?: pulumi.Input<string>;
    /**
     * The role authorized by RAM. The value formats as `acs:ram::{AccountID}:role/{RoleName}`.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * Resources and filters for event listening. The value formats as `acs:cdn:{RegionID}:{AccountID}:{Filter}`.
     */
    sourceArn?: pulumi.Input<string>;
    /**
     * The trigger corresponding to the function Compute Service. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`. See [Create a CDN Fc Trigger](https://www.alibabacloud.com/help/zh/alibaba-cloud-cdn/latest/add-function-calculation-trigger) for more details.
     */
    triggerArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FcTrigger resource.
 */
export interface FcTriggerArgs {
    /**
     * The name of the Event.
     */
    eventMetaName: pulumi.Input<string>;
    /**
     * The version of the Event.
     */
    eventMetaVersion: pulumi.Input<string>;
    /**
     * The function arn. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`.
     */
    functionArn?: pulumi.Input<string>;
    /**
     * The Note information.
     */
    notes: pulumi.Input<string>;
    /**
     * The role authorized by RAM. The value formats as `acs:ram::{AccountID}:role/{RoleName}`.
     */
    roleArn: pulumi.Input<string>;
    /**
     * Resources and filters for event listening. The value formats as `acs:cdn:{RegionID}:{AccountID}:{Filter}`.
     */
    sourceArn: pulumi.Input<string>;
    /**
     * The trigger corresponding to the function Compute Service. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`. See [Create a CDN Fc Trigger](https://www.alibabacloud.com/help/zh/alibaba-cloud-cdn/latest/add-function-calculation-trigger) for more details.
     */
    triggerArn: pulumi.Input<string>;
}
