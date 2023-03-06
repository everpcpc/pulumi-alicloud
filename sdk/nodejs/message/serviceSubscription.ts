// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Message Notification Service Subscription resource.
 *
 * For information about Message Notification Service Subscription and how to use it, see [What is Subscription](https://www.alibabacloud.com/help/en/message-service/latest/subscribe-1).
 *
 * > **NOTE:** Available in v1.188.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultServiceTopic = new alicloud.message.ServiceTopic("defaultServiceTopic", {
 *     topicName: "tf-example-value",
 *     maxMessageSize: 12357,
 *     loggingEnabled: true,
 * });
 * const defaultServiceSubscription = new alicloud.message.ServiceSubscription("defaultServiceSubscription", {
 *     topicName: defaultServiceTopic.topicName,
 *     subscriptionName: "tf-example-value",
 *     endpoint: "http://www.test.com/test",
 *     pushType: "http",
 *     filterTag: "tf-test",
 *     notifyContentFormat: "XML",
 *     notifyStrategy: "BACKOFF_RETRY",
 * });
 * ```
 *
 * ## Import
 *
 * Message Notification Service Subscription can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:message/serviceSubscription:ServiceSubscription example <topic_name>:<subscription_name>
 * ```
 */
export class ServiceSubscription extends pulumi.CustomResource {
    /**
     * Get an existing ServiceSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceSubscriptionState, opts?: pulumi.CustomResourceOptions): ServiceSubscription {
        return new ServiceSubscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:message/serviceSubscription:ServiceSubscription';

    /**
     * Returns true if the given object is an instance of ServiceSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceSubscription.__pulumiType;
    }

    /**
     * The endpoint has three format. Available values format:
     * - `HTTP Format`: http://xxx.com/xxx
     * - `Queue Format`: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
     * - `Email Format`: mail:directmail:{MailAddress}
     */
    public readonly endpoint!: pulumi.Output<string>;
    /**
     * The tag that is used to filter messages. Only the messages that have the same tag can be pushed. A tag is a string that can be up to 16 characters in length. By default, no tag is specified to filter messages.
     */
    public readonly filterTag!: pulumi.Output<string | undefined>;
    /**
     * The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. Valid values: `XML`, `JSON` and `SIMPLIFIED`. Default value: `XML`.
     */
    public readonly notifyContentFormat!: pulumi.Output<string>;
    /**
     * The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. Default value: `BACKOFF_RETRY`. Valid values:
     */
    public readonly notifyStrategy!: pulumi.Output<string>;
    /**
     * The Push type of Subscription. The Valid values: `http`, `queue`, `mpush`, `alisms` and `email`.
     */
    public readonly pushType!: pulumi.Output<string>;
    /**
     * Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
     */
    public readonly subscriptionName!: pulumi.Output<string>;
    /**
     * The topic which The subscription belongs to was named with the name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
     */
    public readonly topicName!: pulumi.Output<string>;

    /**
     * Create a ServiceSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceSubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceSubscriptionArgs | ServiceSubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceSubscriptionState | undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["filterTag"] = state ? state.filterTag : undefined;
            resourceInputs["notifyContentFormat"] = state ? state.notifyContentFormat : undefined;
            resourceInputs["notifyStrategy"] = state ? state.notifyStrategy : undefined;
            resourceInputs["pushType"] = state ? state.pushType : undefined;
            resourceInputs["subscriptionName"] = state ? state.subscriptionName : undefined;
            resourceInputs["topicName"] = state ? state.topicName : undefined;
        } else {
            const args = argsOrState as ServiceSubscriptionArgs | undefined;
            if ((!args || args.endpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpoint'");
            }
            if ((!args || args.pushType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pushType'");
            }
            if ((!args || args.subscriptionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subscriptionName'");
            }
            if ((!args || args.topicName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicName'");
            }
            resourceInputs["endpoint"] = args ? args.endpoint : undefined;
            resourceInputs["filterTag"] = args ? args.filterTag : undefined;
            resourceInputs["notifyContentFormat"] = args ? args.notifyContentFormat : undefined;
            resourceInputs["notifyStrategy"] = args ? args.notifyStrategy : undefined;
            resourceInputs["pushType"] = args ? args.pushType : undefined;
            resourceInputs["subscriptionName"] = args ? args.subscriptionName : undefined;
            resourceInputs["topicName"] = args ? args.topicName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceSubscription.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceSubscription resources.
 */
export interface ServiceSubscriptionState {
    /**
     * The endpoint has three format. Available values format:
     * - `HTTP Format`: http://xxx.com/xxx
     * - `Queue Format`: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
     * - `Email Format`: mail:directmail:{MailAddress}
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The tag that is used to filter messages. Only the messages that have the same tag can be pushed. A tag is a string that can be up to 16 characters in length. By default, no tag is specified to filter messages.
     */
    filterTag?: pulumi.Input<string>;
    /**
     * The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. Valid values: `XML`, `JSON` and `SIMPLIFIED`. Default value: `XML`.
     */
    notifyContentFormat?: pulumi.Input<string>;
    /**
     * The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. Default value: `BACKOFF_RETRY`. Valid values:
     */
    notifyStrategy?: pulumi.Input<string>;
    /**
     * The Push type of Subscription. The Valid values: `http`, `queue`, `mpush`, `alisms` and `email`.
     */
    pushType?: pulumi.Input<string>;
    /**
     * Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
     */
    subscriptionName?: pulumi.Input<string>;
    /**
     * The topic which The subscription belongs to was named with the name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
     */
    topicName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceSubscription resource.
 */
export interface ServiceSubscriptionArgs {
    /**
     * The endpoint has three format. Available values format:
     * - `HTTP Format`: http://xxx.com/xxx
     * - `Queue Format`: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
     * - `Email Format`: mail:directmail:{MailAddress}
     */
    endpoint: pulumi.Input<string>;
    /**
     * The tag that is used to filter messages. Only the messages that have the same tag can be pushed. A tag is a string that can be up to 16 characters in length. By default, no tag is specified to filter messages.
     */
    filterTag?: pulumi.Input<string>;
    /**
     * The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. Valid values: `XML`, `JSON` and `SIMPLIFIED`. Default value: `XML`.
     */
    notifyContentFormat?: pulumi.Input<string>;
    /**
     * The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. Default value: `BACKOFF_RETRY`. Valid values:
     */
    notifyStrategy?: pulumi.Input<string>;
    /**
     * The Push type of Subscription. The Valid values: `http`, `queue`, `mpush`, `alisms` and `email`.
     */
    pushType: pulumi.Input<string>;
    /**
     * Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
     */
    subscriptionName: pulumi.Input<string>;
    /**
     * The topic which The subscription belongs to was named with the name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
     */
    topicName: pulumi.Input<string>;
}