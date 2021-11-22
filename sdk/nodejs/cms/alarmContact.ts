// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates or modifies an alarm contact. For information about alarm contact and how to use it, see [What is alarm contact](https://www.alibabacloud.com/help/en/doc-detail/114923.htm).
 *
 * > **NOTE:** Available in v1.99.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // If you use this template, you need to activate the link before you can return to the alarm contact information, otherwise diff will appear in terraform. So please confirm the activation link as soon as possible.
 * const example = new alicloud.cms.AlarmContact("example", {
 *     alarmContactName: "zhangsan",
 *     channelsMail: "terraform.test.com",
 *     describe: "For Test",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // If you use this template, you can ignore the diff of the alarm contact information by `lifestyle`. We recommend the above usage and activate the link in time.
 * const example = new alicloud.cms.AlarmContact("example", {
 *     alarmContactName: "zhangsan",
 *     describe: "For Test",
 *     channelsMail: "terraform.test.com",
 * });
 * ```
 *
 * ## Import
 *
 * Alarm contact can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cms/alarmContact:AlarmContact example abc12345
 * ```
 */
export class AlarmContact extends pulumi.CustomResource {
    /**
     * Get an existing AlarmContact resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlarmContactState, opts?: pulumi.CustomResourceOptions): AlarmContact {
        return new AlarmContact(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cms/alarmContact:AlarmContact';

    /**
     * Returns true if the given object is an instance of AlarmContact.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlarmContact {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlarmContact.__pulumiType;
    }

    /**
     * The name of the alarm contact. The length should between 2 and 40 characters.
     */
    public readonly alarmContactName!: pulumi.Output<string>;
    /**
     * The TradeManager ID of the alarm contact.
     */
    public readonly channelsAliim!: pulumi.Output<string | undefined>;
    /**
     * The webhook URL of the DingTalk chatbot.
     */
    public readonly channelsDingWebHook!: pulumi.Output<string | undefined>;
    /**
     * The email address of the alarm contact. After you add or modify an email address, the recipient receives an email that contains an activation link. The system adds the recipient to the list of alarm contacts only after the recipient activates the email address.
     */
    public readonly channelsMail!: pulumi.Output<string | undefined>;
    /**
     * The phone number of the alarm contact. After you add or modify an email address, the recipient receives an email that contains an activation link. The system adds the recipient to the list of alarm contacts only after the recipient activates the email address.
     */
    public readonly channelsSms!: pulumi.Output<string | undefined>;
    /**
     * The description of the alarm contact.
     */
    public readonly describe!: pulumi.Output<string>;
    /**
     * The language type of the alarm. Valid values: `en`, `zh-cn`.
     */
    public readonly lang!: pulumi.Output<string | undefined>;

    /**
     * Create a AlarmContact resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlarmContactArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlarmContactArgs | AlarmContactState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlarmContactState | undefined;
            inputs["alarmContactName"] = state ? state.alarmContactName : undefined;
            inputs["channelsAliim"] = state ? state.channelsAliim : undefined;
            inputs["channelsDingWebHook"] = state ? state.channelsDingWebHook : undefined;
            inputs["channelsMail"] = state ? state.channelsMail : undefined;
            inputs["channelsSms"] = state ? state.channelsSms : undefined;
            inputs["describe"] = state ? state.describe : undefined;
            inputs["lang"] = state ? state.lang : undefined;
        } else {
            const args = argsOrState as AlarmContactArgs | undefined;
            if ((!args || args.alarmContactName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alarmContactName'");
            }
            if ((!args || args.describe === undefined) && !opts.urn) {
                throw new Error("Missing required property 'describe'");
            }
            inputs["alarmContactName"] = args ? args.alarmContactName : undefined;
            inputs["channelsAliim"] = args ? args.channelsAliim : undefined;
            inputs["channelsDingWebHook"] = args ? args.channelsDingWebHook : undefined;
            inputs["channelsMail"] = args ? args.channelsMail : undefined;
            inputs["channelsSms"] = args ? args.channelsSms : undefined;
            inputs["describe"] = args ? args.describe : undefined;
            inputs["lang"] = args ? args.lang : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AlarmContact.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlarmContact resources.
 */
export interface AlarmContactState {
    /**
     * The name of the alarm contact. The length should between 2 and 40 characters.
     */
    alarmContactName?: pulumi.Input<string>;
    /**
     * The TradeManager ID of the alarm contact.
     */
    channelsAliim?: pulumi.Input<string>;
    /**
     * The webhook URL of the DingTalk chatbot.
     */
    channelsDingWebHook?: pulumi.Input<string>;
    /**
     * The email address of the alarm contact. After you add or modify an email address, the recipient receives an email that contains an activation link. The system adds the recipient to the list of alarm contacts only after the recipient activates the email address.
     */
    channelsMail?: pulumi.Input<string>;
    /**
     * The phone number of the alarm contact. After you add or modify an email address, the recipient receives an email that contains an activation link. The system adds the recipient to the list of alarm contacts only after the recipient activates the email address.
     */
    channelsSms?: pulumi.Input<string>;
    /**
     * The description of the alarm contact.
     */
    describe?: pulumi.Input<string>;
    /**
     * The language type of the alarm. Valid values: `en`, `zh-cn`.
     */
    lang?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AlarmContact resource.
 */
export interface AlarmContactArgs {
    /**
     * The name of the alarm contact. The length should between 2 and 40 characters.
     */
    alarmContactName: pulumi.Input<string>;
    /**
     * The TradeManager ID of the alarm contact.
     */
    channelsAliim?: pulumi.Input<string>;
    /**
     * The webhook URL of the DingTalk chatbot.
     */
    channelsDingWebHook?: pulumi.Input<string>;
    /**
     * The email address of the alarm contact. After you add or modify an email address, the recipient receives an email that contains an activation link. The system adds the recipient to the list of alarm contacts only after the recipient activates the email address.
     */
    channelsMail?: pulumi.Input<string>;
    /**
     * The phone number of the alarm contact. After you add or modify an email address, the recipient receives an email that contains an activation link. The system adds the recipient to the list of alarm contacts only after the recipient activates the email address.
     */
    channelsSms?: pulumi.Input<string>;
    /**
     * The description of the alarm contact.
     */
    describe: pulumi.Input<string>;
    /**
     * The language type of the alarm. Valid values: `en`, `zh-cn`.
     */
    lang?: pulumi.Input<string>;
}
