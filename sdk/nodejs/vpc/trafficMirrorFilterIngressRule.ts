// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPC Traffic Mirror Filter Ingress Rule resource.
 *
 * For information about VPC Traffic Mirror Filter Ingress Rule and how to use it, see [What is Traffic Mirror Filter Ingress Rule](https://www.alibabacloud.com/help/doc-detail/261357.htm).
 *
 * > **NOTE:** Available in v1.141.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const exampleTrafficMirrorFilter = new alicloud.vpc.TrafficMirrorFilter("exampleTrafficMirrorFilter", {trafficMirrorFilterName: "example_value"});
 * const exampleTrafficMirrorFilterIngressRule = new alicloud.vpc.TrafficMirrorFilterIngressRule("exampleTrafficMirrorFilterIngressRule", {
 *     trafficMirrorFilterId: exampleTrafficMirrorFilter.id,
 *     priority: "1",
 *     ruleAction: "accept",
 *     protocol: "UDP",
 *     destinationCidrBlock: "10.0.0.0/24",
 *     sourceCidrBlock: "10.0.0.0/24",
 *     destinationPortRange: "1/120",
 *     sourcePortRange: "1/120",
 * });
 * ```
 *
 * ## Import
 *
 * VPC Traffic Mirror Filter Ingress Rule can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpc/trafficMirrorFilterIngressRule:TrafficMirrorFilterIngressRule example <traffic_mirror_filter_id>:<traffic_mirror_filter_ingress_rule_id>
 * ```
 */
export class TrafficMirrorFilterIngressRule extends pulumi.CustomResource {
    /**
     * Get an existing TrafficMirrorFilterIngressRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TrafficMirrorFilterIngressRuleState, opts?: pulumi.CustomResourceOptions): TrafficMirrorFilterIngressRule {
        return new TrafficMirrorFilterIngressRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/trafficMirrorFilterIngressRule:TrafficMirrorFilterIngressRule';

    /**
     * Returns true if the given object is an instance of TrafficMirrorFilterIngressRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TrafficMirrorFilterIngressRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TrafficMirrorFilterIngressRule.__pulumiType;
    }

    /**
     * The destination CIDR block of the inbound traffic.
     */
    public readonly destinationCidrBlock!: pulumi.Output<string>;
    /**
     * The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     */
    public readonly destinationPortRange!: pulumi.Output<string>;
    /**
     * Whether to pre-check this request only. Default to: `false`
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
     */
    public readonly ruleAction!: pulumi.Output<string>;
    /**
     * The source CIDR block of the inbound traffic.
     */
    public readonly sourceCidrBlock!: pulumi.Output<string>;
    /**
     * The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     */
    public readonly sourcePortRange!: pulumi.Output<string>;
    /**
     * The state of the inbound rule. Valid values:`Creating`, `Created`, `Modifying` and `Deleting`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The ID of the filter.
     */
    public readonly trafficMirrorFilterId!: pulumi.Output<string>;
    /**
     * The ID of the inbound rule.
     */
    public /*out*/ readonly trafficMirrorFilterIngressRuleId!: pulumi.Output<string>;

    /**
     * Create a TrafficMirrorFilterIngressRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TrafficMirrorFilterIngressRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TrafficMirrorFilterIngressRuleArgs | TrafficMirrorFilterIngressRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TrafficMirrorFilterIngressRuleState | undefined;
            inputs["destinationCidrBlock"] = state ? state.destinationCidrBlock : undefined;
            inputs["destinationPortRange"] = state ? state.destinationPortRange : undefined;
            inputs["dryRun"] = state ? state.dryRun : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["ruleAction"] = state ? state.ruleAction : undefined;
            inputs["sourceCidrBlock"] = state ? state.sourceCidrBlock : undefined;
            inputs["sourcePortRange"] = state ? state.sourcePortRange : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["trafficMirrorFilterId"] = state ? state.trafficMirrorFilterId : undefined;
            inputs["trafficMirrorFilterIngressRuleId"] = state ? state.trafficMirrorFilterIngressRuleId : undefined;
        } else {
            const args = argsOrState as TrafficMirrorFilterIngressRuleArgs | undefined;
            if ((!args || args.destinationCidrBlock === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationCidrBlock'");
            }
            if ((!args || args.priority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'priority'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.ruleAction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleAction'");
            }
            if ((!args || args.sourceCidrBlock === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceCidrBlock'");
            }
            if ((!args || args.trafficMirrorFilterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trafficMirrorFilterId'");
            }
            inputs["destinationCidrBlock"] = args ? args.destinationCidrBlock : undefined;
            inputs["destinationPortRange"] = args ? args.destinationPortRange : undefined;
            inputs["dryRun"] = args ? args.dryRun : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["ruleAction"] = args ? args.ruleAction : undefined;
            inputs["sourceCidrBlock"] = args ? args.sourceCidrBlock : undefined;
            inputs["sourcePortRange"] = args ? args.sourcePortRange : undefined;
            inputs["trafficMirrorFilterId"] = args ? args.trafficMirrorFilterId : undefined;
            inputs["status"] = undefined /*out*/;
            inputs["trafficMirrorFilterIngressRuleId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TrafficMirrorFilterIngressRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TrafficMirrorFilterIngressRule resources.
 */
export interface TrafficMirrorFilterIngressRuleState {
    /**
     * The destination CIDR block of the inbound traffic.
     */
    destinationCidrBlock?: pulumi.Input<string>;
    /**
     * The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     */
    destinationPortRange?: pulumi.Input<string>;
    /**
     * Whether to pre-check this request only. Default to: `false`
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
     */
    priority?: pulumi.Input<number>;
    /**
     * The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
     */
    ruleAction?: pulumi.Input<string>;
    /**
     * The source CIDR block of the inbound traffic.
     */
    sourceCidrBlock?: pulumi.Input<string>;
    /**
     * The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     */
    sourcePortRange?: pulumi.Input<string>;
    /**
     * The state of the inbound rule. Valid values:`Creating`, `Created`, `Modifying` and `Deleting`.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the filter.
     */
    trafficMirrorFilterId?: pulumi.Input<string>;
    /**
     * The ID of the inbound rule.
     */
    trafficMirrorFilterIngressRuleId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TrafficMirrorFilterIngressRule resource.
 */
export interface TrafficMirrorFilterIngressRuleArgs {
    /**
     * The destination CIDR block of the inbound traffic.
     */
    destinationCidrBlock: pulumi.Input<string>;
    /**
     * The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     */
    destinationPortRange?: pulumi.Input<string>;
    /**
     * Whether to pre-check this request only. Default to: `false`
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
     */
    priority: pulumi.Input<number>;
    /**
     * The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
     */
    protocol: pulumi.Input<string>;
    /**
     * The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
     */
    ruleAction: pulumi.Input<string>;
    /**
     * The source CIDR block of the inbound traffic.
     */
    sourceCidrBlock: pulumi.Input<string>;
    /**
     * The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     */
    sourcePortRange?: pulumi.Input<string>;
    /**
     * The ID of the filter.
     */
    trafficMirrorFilterId: pulumi.Input<string>;
}
