// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * VPC can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpc/network:Network example vpc-abc123456
 * ```
 */
export class Network extends pulumi.CustomResource {
    /**
     * Get an existing Network resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkState, opts?: pulumi.CustomResourceOptions): Network {
        return new Network(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/network:Network';

    /**
     * Returns true if the given object is an instance of Network.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Network {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Network.__pulumiType;
    }

    /**
     * The CIDR block for the VPC. The `cidrBlock` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
     */
    public readonly cidrBlock!: pulumi.Output<string | undefined>;
    /**
     * The VPC description. Defaults to null.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to precheck this request only. Valid values: `true` and `false`.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether to enable the IPv6 CIDR block. Valid values: `false` (Default): disables IPv6 CIDR blocks. `true`: enables IPv6 CIDR blocks. If the `enableIpv6` is `true`, the system will automatically create a free version of an IPv6 gateway for your private network and assign an IPv6 network segment assigned as /56.
     */
    public readonly enableIpv6!: pulumi.Output<boolean | undefined>;
    /**
     * (Available in v1.119.0+) ) The ipv6 cidr block of VPC.
     */
    public /*out*/ readonly ipv6CidrBlock!: pulumi.Output<string>;
    /**
     * Field `name` has been deprecated from provider version 1.119.0. New field `vpcName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Id of resource group which the VPC belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The route table ID of the router created by default on VPC creation.
     */
    public /*out*/ readonly routeTableId!: pulumi.Output<string>;
    /**
     * The ID of the router created by default on VPC creation.
     */
    public /*out*/ readonly routerId!: pulumi.Output<string>;
    /**
     * @deprecated Attribute router_table_id has been deprecated and replaced with route_table_id.
     */
    public /*out*/ readonly routerTableId!: pulumi.Output<string>;
    /**
     * The secondary CIDR blocks for the VPC.
     */
    public readonly secondaryCidrBlocks!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The user cidrs of the VPC.
     */
    public readonly userCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the VPC. Defaults to null.
     */
    public readonly vpcName!: pulumi.Output<string>;

    /**
     * Create a Network resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkArgs | NetworkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkState | undefined;
            inputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["dryRun"] = state ? state.dryRun : undefined;
            inputs["enableIpv6"] = state ? state.enableIpv6 : undefined;
            inputs["ipv6CidrBlock"] = state ? state.ipv6CidrBlock : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["routeTableId"] = state ? state.routeTableId : undefined;
            inputs["routerId"] = state ? state.routerId : undefined;
            inputs["routerTableId"] = state ? state.routerTableId : undefined;
            inputs["secondaryCidrBlocks"] = state ? state.secondaryCidrBlocks : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["userCidrs"] = state ? state.userCidrs : undefined;
            inputs["vpcName"] = state ? state.vpcName : undefined;
        } else {
            const args = argsOrState as NetworkArgs | undefined;
            inputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["dryRun"] = args ? args.dryRun : undefined;
            inputs["enableIpv6"] = args ? args.enableIpv6 : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["secondaryCidrBlocks"] = args ? args.secondaryCidrBlocks : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["userCidrs"] = args ? args.userCidrs : undefined;
            inputs["vpcName"] = args ? args.vpcName : undefined;
            inputs["ipv6CidrBlock"] = undefined /*out*/;
            inputs["routeTableId"] = undefined /*out*/;
            inputs["routerId"] = undefined /*out*/;
            inputs["routerTableId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Network.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Network resources.
 */
export interface NetworkState {
    /**
     * The CIDR block for the VPC. The `cidrBlock` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * The VPC description. Defaults to null.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies whether to precheck this request only. Valid values: `true` and `false`.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable the IPv6 CIDR block. Valid values: `false` (Default): disables IPv6 CIDR blocks. `true`: enables IPv6 CIDR blocks. If the `enableIpv6` is `true`, the system will automatically create a free version of an IPv6 gateway for your private network and assign an IPv6 network segment assigned as /56.
     */
    enableIpv6?: pulumi.Input<boolean>;
    /**
     * (Available in v1.119.0+) ) The ipv6 cidr block of VPC.
     */
    ipv6CidrBlock?: pulumi.Input<string>;
    /**
     * Field `name` has been deprecated from provider version 1.119.0. New field `vpcName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The Id of resource group which the VPC belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The route table ID of the router created by default on VPC creation.
     */
    routeTableId?: pulumi.Input<string>;
    /**
     * The ID of the router created by default on VPC creation.
     */
    routerId?: pulumi.Input<string>;
    /**
     * @deprecated Attribute router_table_id has been deprecated and replaced with route_table_id.
     */
    routerTableId?: pulumi.Input<string>;
    /**
     * The secondary CIDR blocks for the VPC.
     */
    secondaryCidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The user cidrs of the VPC.
     */
    userCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the VPC. Defaults to null.
     */
    vpcName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Network resource.
 */
export interface NetworkArgs {
    /**
     * The CIDR block for the VPC. The `cidrBlock` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * The VPC description. Defaults to null.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies whether to precheck this request only. Valid values: `true` and `false`.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable the IPv6 CIDR block. Valid values: `false` (Default): disables IPv6 CIDR blocks. `true`: enables IPv6 CIDR blocks. If the `enableIpv6` is `true`, the system will automatically create a free version of an IPv6 gateway for your private network and assign an IPv6 network segment assigned as /56.
     */
    enableIpv6?: pulumi.Input<boolean>;
    /**
     * Field `name` has been deprecated from provider version 1.119.0. New field `vpcName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The Id of resource group which the VPC belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The secondary CIDR blocks for the VPC.
     */
    secondaryCidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The user cidrs of the VPC.
     */
    userCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the VPC. Defaults to null.
     */
    vpcName?: pulumi.Input<string>;
}
