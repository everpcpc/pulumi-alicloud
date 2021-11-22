// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Application Real-Time Monitoring Service (ARMS) Alert Contact resource.
 *
 * For information about Application Real-Time Monitoring Service (ARMS) Alert Contact and how to use it, see [What is Alert Contact](https://www.alibabacloud.com/help/en/doc-detail/42953.htm).
 *
 * > **NOTE:** Available in v1.129.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.arms.AlertContact("example", {
 *     alertContactName: "example_value",
 *     dingRobotWebhookUrl: "https://oapi.dingtalk.com/robot/send?access_token=91f2f6****",
 *     email: "someone@example.com",
 *     phoneNum: "1381111****",
 * });
 * ```
 *
 * ## Import
 *
 * Application Real-Time Monitoring Service (ARMS) Alert Contact can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:arms/alertContact:AlertContact example <id>
 * ```
 */
export class AlertContact extends pulumi.CustomResource {
    /**
     * Get an existing AlertContact resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlertContactState, opts?: pulumi.CustomResourceOptions): AlertContact {
        return new AlertContact(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:arms/alertContact:AlertContact';

    /**
     * Returns true if the given object is an instance of AlertContact.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlertContact {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlertContact.__pulumiType;
    }

    /**
     * The name of the alert contact.
     */
    public readonly alertContactName!: pulumi.Output<string | undefined>;
    /**
     * The webhook URL of the DingTalk chatbot. For more information about how to obtain the URL, see Configure a DingTalk chatbot to send alert notifications: https://www.alibabacloud.com/help/en/doc-detail/106247.htm. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
     */
    public readonly dingRobotWebhookUrl!: pulumi.Output<string | undefined>;
    /**
     * The email address of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
     */
    public readonly email!: pulumi.Output<string | undefined>;
    /**
     * The mobile number of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
     */
    public readonly phoneNum!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the alert contact receives system notifications. Valid values:  true: receives system notifications. false: does not receive system notifications.
     */
    public readonly systemNoc!: pulumi.Output<boolean | undefined>;

    /**
     * Create a AlertContact resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AlertContactArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlertContactArgs | AlertContactState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlertContactState | undefined;
            inputs["alertContactName"] = state ? state.alertContactName : undefined;
            inputs["dingRobotWebhookUrl"] = state ? state.dingRobotWebhookUrl : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["phoneNum"] = state ? state.phoneNum : undefined;
            inputs["systemNoc"] = state ? state.systemNoc : undefined;
        } else {
            const args = argsOrState as AlertContactArgs | undefined;
            inputs["alertContactName"] = args ? args.alertContactName : undefined;
            inputs["dingRobotWebhookUrl"] = args ? args.dingRobotWebhookUrl : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["phoneNum"] = args ? args.phoneNum : undefined;
            inputs["systemNoc"] = args ? args.systemNoc : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AlertContact.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlertContact resources.
 */
export interface AlertContactState {
    /**
     * The name of the alert contact.
     */
    alertContactName?: pulumi.Input<string>;
    /**
     * The webhook URL of the DingTalk chatbot. For more information about how to obtain the URL, see Configure a DingTalk chatbot to send alert notifications: https://www.alibabacloud.com/help/en/doc-detail/106247.htm. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
     */
    dingRobotWebhookUrl?: pulumi.Input<string>;
    /**
     * The email address of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
     */
    email?: pulumi.Input<string>;
    /**
     * The mobile number of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
     */
    phoneNum?: pulumi.Input<string>;
    /**
     * Specifies whether the alert contact receives system notifications. Valid values:  true: receives system notifications. false: does not receive system notifications.
     */
    systemNoc?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AlertContact resource.
 */
export interface AlertContactArgs {
    /**
     * The name of the alert contact.
     */
    alertContactName?: pulumi.Input<string>;
    /**
     * The webhook URL of the DingTalk chatbot. For more information about how to obtain the URL, see Configure a DingTalk chatbot to send alert notifications: https://www.alibabacloud.com/help/en/doc-detail/106247.htm. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
     */
    dingRobotWebhookUrl?: pulumi.Input<string>;
    /**
     * The email address of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
     */
    email?: pulumi.Input<string>;
    /**
     * The mobile number of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
     */
    phoneNum?: pulumi.Input<string>;
    /**
     * Specifies whether the alert contact receives system notifications. Valid values:  true: receives system notifications. false: does not receive system notifications.
     */
    systemNoc?: pulumi.Input<boolean>;
}
