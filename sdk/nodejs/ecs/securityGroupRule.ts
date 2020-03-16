// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class SecurityGroupRule extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroupRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupRuleState, opts?: pulumi.CustomResourceOptions): SecurityGroupRule {
        return new SecurityGroupRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/securityGroupRule:SecurityGroupRule';

    /**
     * Returns true if the given object is an instance of SecurityGroupRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroupRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroupRule.__pulumiType;
    }

    /**
     * The target IP address range. The default value is 0.0.0.0/0 (which means no restriction will be applied). Other supported formats include 10.159.6.18/12. Only IPv4 is supported.
     */
    public readonly cidrIp!: pulumi.Output<string | undefined>;
    /**
     * The description of the security group rule. The description can be up to 1 to 512 characters in length. Defaults to null.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The protocol. Can be `tcp`, `udp`, `icmp`, `gre` or `all`.
     */
    public readonly ipProtocol!: pulumi.Output<string>;
    /**
     * Network type, can be either `internet` or `intranet`, the default value is `internet`.
     */
    public readonly nicType!: pulumi.Output<string>;
    /**
     * Authorization policy, can be either `accept` or `drop`, the default value is `accept`.
     */
    public readonly policy!: pulumi.Output<string | undefined>;
    /**
     * The range of port numbers relevant to the IP protocol. Default to "-1/-1". When the protocol is tcp or udp, each side port number range from 1 to 65535 and '-1/-1' will be invalid.
     * For example, `1/200` means that the range of the port numbers is 1-200. Other protocols' 'port_range' can only be "-1/-1", and other values will be invalid.
     */
    public readonly portRange!: pulumi.Output<string | undefined>;
    /**
     * Authorization policy priority, with parameter values: `1-100`, default value: 1.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The security group to apply this rule to.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * The Alibaba Cloud user account Id of the target security group when security groups are authorized across accounts.  This parameter is invalid if `cidrIp` has already been set.
     */
    public readonly sourceGroupOwnerAccount!: pulumi.Output<string | undefined>;
    /**
     * The target security group ID within the same region. If this field is specified, the `nicType` can only select `intranet`.
     */
    public readonly sourceSecurityGroupId!: pulumi.Output<string | undefined>;
    /**
     * The type of rule being created. Valid options are `ingress` (inbound) or `egress` (outbound).
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a SecurityGroupRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityGroupRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupRuleArgs | SecurityGroupRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecurityGroupRuleState | undefined;
            inputs["cidrIp"] = state ? state.cidrIp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["ipProtocol"] = state ? state.ipProtocol : undefined;
            inputs["nicType"] = state ? state.nicType : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["portRange"] = state ? state.portRange : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            inputs["sourceGroupOwnerAccount"] = state ? state.sourceGroupOwnerAccount : undefined;
            inputs["sourceSecurityGroupId"] = state ? state.sourceSecurityGroupId : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as SecurityGroupRuleArgs | undefined;
            if (!args || args.ipProtocol === undefined) {
                throw new Error("Missing required property 'ipProtocol'");
            }
            if (!args || args.securityGroupId === undefined) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["cidrIp"] = args ? args.cidrIp : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["ipProtocol"] = args ? args.ipProtocol : undefined;
            inputs["nicType"] = args ? args.nicType : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["portRange"] = args ? args.portRange : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            inputs["sourceGroupOwnerAccount"] = args ? args.sourceGroupOwnerAccount : undefined;
            inputs["sourceSecurityGroupId"] = args ? args.sourceSecurityGroupId : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecurityGroupRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroupRule resources.
 */
export interface SecurityGroupRuleState {
    /**
     * The target IP address range. The default value is 0.0.0.0/0 (which means no restriction will be applied). Other supported formats include 10.159.6.18/12. Only IPv4 is supported.
     */
    readonly cidrIp?: pulumi.Input<string>;
    /**
     * The description of the security group rule. The description can be up to 1 to 512 characters in length. Defaults to null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The protocol. Can be `tcp`, `udp`, `icmp`, `gre` or `all`.
     */
    readonly ipProtocol?: pulumi.Input<string>;
    /**
     * Network type, can be either `internet` or `intranet`, the default value is `internet`.
     */
    readonly nicType?: pulumi.Input<string>;
    /**
     * Authorization policy, can be either `accept` or `drop`, the default value is `accept`.
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * The range of port numbers relevant to the IP protocol. Default to "-1/-1". When the protocol is tcp or udp, each side port number range from 1 to 65535 and '-1/-1' will be invalid.
     * For example, `1/200` means that the range of the port numbers is 1-200. Other protocols' 'port_range' can only be "-1/-1", and other values will be invalid.
     */
    readonly portRange?: pulumi.Input<string>;
    /**
     * Authorization policy priority, with parameter values: `1-100`, default value: 1.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The security group to apply this rule to.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * The Alibaba Cloud user account Id of the target security group when security groups are authorized across accounts.  This parameter is invalid if `cidrIp` has already been set.
     */
    readonly sourceGroupOwnerAccount?: pulumi.Input<string>;
    /**
     * The target security group ID within the same region. If this field is specified, the `nicType` can only select `intranet`.
     */
    readonly sourceSecurityGroupId?: pulumi.Input<string>;
    /**
     * The type of rule being created. Valid options are `ingress` (inbound) or `egress` (outbound).
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityGroupRule resource.
 */
export interface SecurityGroupRuleArgs {
    /**
     * The target IP address range. The default value is 0.0.0.0/0 (which means no restriction will be applied). Other supported formats include 10.159.6.18/12. Only IPv4 is supported.
     */
    readonly cidrIp?: pulumi.Input<string>;
    /**
     * The description of the security group rule. The description can be up to 1 to 512 characters in length. Defaults to null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The protocol. Can be `tcp`, `udp`, `icmp`, `gre` or `all`.
     */
    readonly ipProtocol: pulumi.Input<string>;
    /**
     * Network type, can be either `internet` or `intranet`, the default value is `internet`.
     */
    readonly nicType?: pulumi.Input<string>;
    /**
     * Authorization policy, can be either `accept` or `drop`, the default value is `accept`.
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * The range of port numbers relevant to the IP protocol. Default to "-1/-1". When the protocol is tcp or udp, each side port number range from 1 to 65535 and '-1/-1' will be invalid.
     * For example, `1/200` means that the range of the port numbers is 1-200. Other protocols' 'port_range' can only be "-1/-1", and other values will be invalid.
     */
    readonly portRange?: pulumi.Input<string>;
    /**
     * Authorization policy priority, with parameter values: `1-100`, default value: 1.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The security group to apply this rule to.
     */
    readonly securityGroupId: pulumi.Input<string>;
    /**
     * The Alibaba Cloud user account Id of the target security group when security groups are authorized across accounts.  This parameter is invalid if `cidrIp` has already been set.
     */
    readonly sourceGroupOwnerAccount?: pulumi.Input<string>;
    /**
     * The target security group ID within the same region. If this field is specified, the `nicType` can only select `intranet`.
     */
    readonly sourceSecurityGroupId?: pulumi.Input<string>;
    /**
     * The type of rule being created. Valid options are `ingress` (inbound) or `egress` (outbound).
     */
    readonly type: pulumi.Input<string>;
}
