// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CMS Alarm Contact Group resource.
 *
 * For information about CMS Alarm Contact Group and how to use it, see [What is Alarm Contact Group](https://www.alibabacloud.com/help/en/doc-detail/114929.htm).
 *
 * > **NOTE:** Available in v1.101.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.cms.AlarmContactGroup("example", {
 *     alarmContactGroupName: "tf-test",
 * });
 * ```
 *
 * ## Import
 *
 * CMS Alarm Contact Group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cms/alarmContactGroup:AlarmContactGroup example tf-testacc123
 * ```
 */
export class AlarmContactGroup extends pulumi.CustomResource {
    /**
     * Get an existing AlarmContactGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlarmContactGroupState, opts?: pulumi.CustomResourceOptions): AlarmContactGroup {
        return new AlarmContactGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cms/alarmContactGroup:AlarmContactGroup';

    /**
     * Returns true if the given object is an instance of AlarmContactGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlarmContactGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlarmContactGroup.__pulumiType;
    }

    /**
     * The name of the alarm group.
     */
    public readonly alarmContactGroupName!: pulumi.Output<string>;
    /**
     * The name of the alert contact.
     */
    public readonly contacts!: pulumi.Output<string[] | undefined>;
    /**
     * The description of the alert group.
     */
    public readonly describe!: pulumi.Output<string | undefined>;
    /**
     * Whether to open weekly subscription.
     */
    public readonly enableSubscribed!: pulumi.Output<boolean>;

    /**
     * Create a AlarmContactGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlarmContactGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlarmContactGroupArgs | AlarmContactGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlarmContactGroupState | undefined;
            inputs["alarmContactGroupName"] = state ? state.alarmContactGroupName : undefined;
            inputs["contacts"] = state ? state.contacts : undefined;
            inputs["describe"] = state ? state.describe : undefined;
            inputs["enableSubscribed"] = state ? state.enableSubscribed : undefined;
        } else {
            const args = argsOrState as AlarmContactGroupArgs | undefined;
            if ((!args || args.alarmContactGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alarmContactGroupName'");
            }
            inputs["alarmContactGroupName"] = args ? args.alarmContactGroupName : undefined;
            inputs["contacts"] = args ? args.contacts : undefined;
            inputs["describe"] = args ? args.describe : undefined;
            inputs["enableSubscribed"] = args ? args.enableSubscribed : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AlarmContactGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlarmContactGroup resources.
 */
export interface AlarmContactGroupState {
    /**
     * The name of the alarm group.
     */
    alarmContactGroupName?: pulumi.Input<string>;
    /**
     * The name of the alert contact.
     */
    contacts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the alert group.
     */
    describe?: pulumi.Input<string>;
    /**
     * Whether to open weekly subscription.
     */
    enableSubscribed?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AlarmContactGroup resource.
 */
export interface AlarmContactGroupArgs {
    /**
     * The name of the alarm group.
     */
    alarmContactGroupName: pulumi.Input<string>;
    /**
     * The name of the alert contact.
     */
    contacts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the alert group.
     */
    describe?: pulumi.Input<string>;
    /**
     * Whether to open weekly subscription.
     */
    enableSubscribed?: pulumi.Input<boolean>;
}
