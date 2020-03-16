// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class TopicSubscription extends pulumi.CustomResource {
    /**
     * Get an existing TopicSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TopicSubscriptionState, opts?: pulumi.CustomResourceOptions): TopicSubscription {
        return new TopicSubscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:mns/topicSubscription:TopicSubscription';

    /**
     * Returns true if the given object is an instance of TopicSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TopicSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TopicSubscription.__pulumiType;
    }

    /**
     * The endpoint has three format. Available values format:
     * - HTTP Format: http://xxx.com/xxx
     * - Queue Format: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
     * - Email Format: mail:directmail:{MailAddress}
     */
    public readonly endpoint!: pulumi.Output<string>;
    /**
     * The length should be shorter than 16.
     */
    public readonly filterTag!: pulumi.Output<string | undefined>;
    /**
     * Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: 'SIMPLIFIED', 'XML' and 'JSON'. Default to 'SIMPLIFIED'.
     */
    public readonly notifyContentFormat!: pulumi.Output<string | undefined>;
    /**
     * The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. the attribute has two value EXPONENTIAL_DECAY_RETR or BACKOFF_RETRY. Default value to BACKOFF_RETRY .
     */
    public readonly notifyStrategy!: pulumi.Output<string | undefined>;
    public readonly topicName!: pulumi.Output<string>;

    /**
     * Create a TopicSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicSubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicSubscriptionArgs | TopicSubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TopicSubscriptionState | undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["filterTag"] = state ? state.filterTag : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notifyContentFormat"] = state ? state.notifyContentFormat : undefined;
            inputs["notifyStrategy"] = state ? state.notifyStrategy : undefined;
            inputs["topicName"] = state ? state.topicName : undefined;
        } else {
            const args = argsOrState as TopicSubscriptionArgs | undefined;
            if (!args || args.endpoint === undefined) {
                throw new Error("Missing required property 'endpoint'");
            }
            if (!args || args.topicName === undefined) {
                throw new Error("Missing required property 'topicName'");
            }
            inputs["endpoint"] = args ? args.endpoint : undefined;
            inputs["filterTag"] = args ? args.filterTag : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notifyContentFormat"] = args ? args.notifyContentFormat : undefined;
            inputs["notifyStrategy"] = args ? args.notifyStrategy : undefined;
            inputs["topicName"] = args ? args.topicName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TopicSubscription.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TopicSubscription resources.
 */
export interface TopicSubscriptionState {
    /**
     * The endpoint has three format. Available values format:
     * - HTTP Format: http://xxx.com/xxx
     * - Queue Format: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
     * - Email Format: mail:directmail:{MailAddress}
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * The length should be shorter than 16.
     */
    readonly filterTag?: pulumi.Input<string>;
    /**
     * Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: 'SIMPLIFIED', 'XML' and 'JSON'. Default to 'SIMPLIFIED'.
     */
    readonly notifyContentFormat?: pulumi.Input<string>;
    /**
     * The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. the attribute has two value EXPONENTIAL_DECAY_RETR or BACKOFF_RETRY. Default value to BACKOFF_RETRY .
     */
    readonly notifyStrategy?: pulumi.Input<string>;
    readonly topicName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TopicSubscription resource.
 */
export interface TopicSubscriptionArgs {
    /**
     * The endpoint has three format. Available values format:
     * - HTTP Format: http://xxx.com/xxx
     * - Queue Format: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
     * - Email Format: mail:directmail:{MailAddress}
     */
    readonly endpoint: pulumi.Input<string>;
    /**
     * The length should be shorter than 16.
     */
    readonly filterTag?: pulumi.Input<string>;
    /**
     * Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: 'SIMPLIFIED', 'XML' and 'JSON'. Default to 'SIMPLIFIED'.
     */
    readonly notifyContentFormat?: pulumi.Input<string>;
    /**
     * The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. the attribute has two value EXPONENTIAL_DECAY_RETR or BACKOFF_RETRY. Default value to BACKOFF_RETRY .
     */
    readonly notifyStrategy?: pulumi.Input<string>;
    readonly topicName: pulumi.Input<string>;
}
