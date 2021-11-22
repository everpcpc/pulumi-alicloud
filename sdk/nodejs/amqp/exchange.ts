// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a RabbitMQ (AMQP) Exchange resource.
 *
 * For information about RabbitMQ (AMQP) Exchange and how to use it, see [What is Exchange](https://help.aliyun.com/).
 *
 * > **NOTE:** Available in v1.128.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const exampleVirtualHost = new alicloud.amqp.VirtualHost("exampleVirtualHost", {
 *     instanceId: "amqp-abc12345",
 *     virtualHostName: "my-VirtualHost",
 * });
 * const exampleExchange = new alicloud.amqp.Exchange("exampleExchange", {
 *     autoDeleteState: false,
 *     exchangeName: "my-Exchange",
 *     exchangeType: "DIRECT",
 *     instanceId: exampleVirtualHost.instanceId,
 *     internal: false,
 *     virtualHostName: exampleVirtualHost.virtualHostName,
 * });
 * ```
 *
 * ## Import
 *
 * RabbitMQ (AMQP) Exchange can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:amqp/exchange:Exchange example <instance_id>:<virtual_host_name>:<exchange_name>
 * ```
 */
export class Exchange extends pulumi.CustomResource {
    /**
     * Get an existing Exchange resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExchangeState, opts?: pulumi.CustomResourceOptions): Exchange {
        return new Exchange(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:amqp/exchange:Exchange';

    /**
     * Returns true if the given object is an instance of Exchange.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Exchange {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Exchange.__pulumiType;
    }

    /**
     * The alternate exchange. An alternate exchange is configured for an existing exchange. It is used to receive messages that fail to be routed to queues from the existing exchange.
     */
    public readonly alternateExchange!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the Auto Delete attribute is configured. Valid values:
     * * true: The Auto Delete attribute is configured. If the last queue that is bound to an exchange is unbound, the exchange is automatically deleted.
     * * false: The Auto Delete attribute is not configured. If the last queue that is bound to an exchange is unbound, the exchange is not automatically deleted.
     */
    public readonly autoDeleteState!: pulumi.Output<boolean>;
    /**
     * The name of the exchange. It must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
     */
    public readonly exchangeName!: pulumi.Output<string>;
    /**
     * The type of the exchange. Valid values:
     * * FANOUT: An exchange of this type routes all the received messages to all the queues bound to this exchange. You can use a fanout exchange to broadcast messages.
     * * DIRECT: An exchange of this type routes a message to the queue whose binding key is exactly the same as the routing key of the message.
     * * TOPIC: This type is similar to the direct exchange type. An exchange of this type routes a message to one or more queues based on the fuzzy match or multi-condition match result between the routing key of the message and the binding keys of the current exchange.
     * * HEADERS: Headers Exchange uses the Headers property instead of Routing Key for routing matching.
     * When binding Headers Exchange and Queue, set the key-value pair of the binding property;
     * when sending a message to the Headers Exchange, set the message's Headers property key-value pair and use the message Headers
     * The message is routed to the bound Queue by comparing the attribute key-value pair and the bound attribute key-value pair.
     */
    public readonly exchangeType!: pulumi.Output<string>;
    /**
     * The ID of the instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Specifies whether an exchange is an internal exchange. Valid values:
     * * false: The exchange is not an internal exchange.
     * * true: The exchange is an internal exchange.
     */
    public readonly internal!: pulumi.Output<boolean>;
    /**
     * The name of virtual host where an exchange resides.
     */
    public readonly virtualHostName!: pulumi.Output<string>;

    /**
     * Create a Exchange resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExchangeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExchangeArgs | ExchangeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExchangeState | undefined;
            inputs["alternateExchange"] = state ? state.alternateExchange : undefined;
            inputs["autoDeleteState"] = state ? state.autoDeleteState : undefined;
            inputs["exchangeName"] = state ? state.exchangeName : undefined;
            inputs["exchangeType"] = state ? state.exchangeType : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["internal"] = state ? state.internal : undefined;
            inputs["virtualHostName"] = state ? state.virtualHostName : undefined;
        } else {
            const args = argsOrState as ExchangeArgs | undefined;
            if ((!args || args.autoDeleteState === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoDeleteState'");
            }
            if ((!args || args.exchangeName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'exchangeName'");
            }
            if ((!args || args.exchangeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'exchangeType'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.internal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'internal'");
            }
            if ((!args || args.virtualHostName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualHostName'");
            }
            inputs["alternateExchange"] = args ? args.alternateExchange : undefined;
            inputs["autoDeleteState"] = args ? args.autoDeleteState : undefined;
            inputs["exchangeName"] = args ? args.exchangeName : undefined;
            inputs["exchangeType"] = args ? args.exchangeType : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["internal"] = args ? args.internal : undefined;
            inputs["virtualHostName"] = args ? args.virtualHostName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Exchange.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Exchange resources.
 */
export interface ExchangeState {
    /**
     * The alternate exchange. An alternate exchange is configured for an existing exchange. It is used to receive messages that fail to be routed to queues from the existing exchange.
     */
    alternateExchange?: pulumi.Input<string>;
    /**
     * Specifies whether the Auto Delete attribute is configured. Valid values:
     * * true: The Auto Delete attribute is configured. If the last queue that is bound to an exchange is unbound, the exchange is automatically deleted.
     * * false: The Auto Delete attribute is not configured. If the last queue that is bound to an exchange is unbound, the exchange is not automatically deleted.
     */
    autoDeleteState?: pulumi.Input<boolean>;
    /**
     * The name of the exchange. It must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
     */
    exchangeName?: pulumi.Input<string>;
    /**
     * The type of the exchange. Valid values:
     * * FANOUT: An exchange of this type routes all the received messages to all the queues bound to this exchange. You can use a fanout exchange to broadcast messages.
     * * DIRECT: An exchange of this type routes a message to the queue whose binding key is exactly the same as the routing key of the message.
     * * TOPIC: This type is similar to the direct exchange type. An exchange of this type routes a message to one or more queues based on the fuzzy match or multi-condition match result between the routing key of the message and the binding keys of the current exchange.
     * * HEADERS: Headers Exchange uses the Headers property instead of Routing Key for routing matching.
     * When binding Headers Exchange and Queue, set the key-value pair of the binding property;
     * when sending a message to the Headers Exchange, set the message's Headers property key-value pair and use the message Headers
     * The message is routed to the bound Queue by comparing the attribute key-value pair and the bound attribute key-value pair.
     */
    exchangeType?: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies whether an exchange is an internal exchange. Valid values:
     * * false: The exchange is not an internal exchange.
     * * true: The exchange is an internal exchange.
     */
    internal?: pulumi.Input<boolean>;
    /**
     * The name of virtual host where an exchange resides.
     */
    virtualHostName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Exchange resource.
 */
export interface ExchangeArgs {
    /**
     * The alternate exchange. An alternate exchange is configured for an existing exchange. It is used to receive messages that fail to be routed to queues from the existing exchange.
     */
    alternateExchange?: pulumi.Input<string>;
    /**
     * Specifies whether the Auto Delete attribute is configured. Valid values:
     * * true: The Auto Delete attribute is configured. If the last queue that is bound to an exchange is unbound, the exchange is automatically deleted.
     * * false: The Auto Delete attribute is not configured. If the last queue that is bound to an exchange is unbound, the exchange is not automatically deleted.
     */
    autoDeleteState: pulumi.Input<boolean>;
    /**
     * The name of the exchange. It must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
     */
    exchangeName: pulumi.Input<string>;
    /**
     * The type of the exchange. Valid values:
     * * FANOUT: An exchange of this type routes all the received messages to all the queues bound to this exchange. You can use a fanout exchange to broadcast messages.
     * * DIRECT: An exchange of this type routes a message to the queue whose binding key is exactly the same as the routing key of the message.
     * * TOPIC: This type is similar to the direct exchange type. An exchange of this type routes a message to one or more queues based on the fuzzy match or multi-condition match result between the routing key of the message and the binding keys of the current exchange.
     * * HEADERS: Headers Exchange uses the Headers property instead of Routing Key for routing matching.
     * When binding Headers Exchange and Queue, set the key-value pair of the binding property;
     * when sending a message to the Headers Exchange, set the message's Headers property key-value pair and use the message Headers
     * The message is routed to the bound Queue by comparing the attribute key-value pair and the bound attribute key-value pair.
     */
    exchangeType: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies whether an exchange is an internal exchange. Valid values:
     * * false: The exchange is not an internal exchange.
     * * true: The exchange is an internal exchange.
     */
    internal: pulumi.Input<boolean>;
    /**
     * The name of virtual host where an exchange resides.
     */
    virtualHostName: pulumi.Input<string>;
}
