// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * The route table can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpc/routeTable:RouteTable foo vtb-abc123456
 * ```
 */
export class RouteTable extends pulumi.CustomResource {
    /**
     * Get an existing RouteTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteTableState, opts?: pulumi.CustomResourceOptions): RouteTable {
        return new RouteTable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/routeTable:RouteTable';

    /**
     * Returns true if the given object is an instance of RouteTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteTable.__pulumiType;
    }

    /**
     * The description of the route table instance.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Field `name` has been deprecated from provider version 1.119.1. New field `routeTableName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the route table.
     */
    public readonly routeTableName!: pulumi.Output<string>;
    /**
     * (Available in v1.119.1+) The status of the route table.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The vpcId of the route table, the field can't be changed.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a RouteTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteTableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteTableArgs | RouteTableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouteTableState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["routeTableName"] = state ? state.routeTableName : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as RouteTableArgs | undefined;
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["routeTableName"] = args ? args.routeTableName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RouteTable.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouteTable resources.
 */
export interface RouteTableState {
    /**
     * The description of the route table instance.
     */
    description?: pulumi.Input<string>;
    /**
     * Field `name` has been deprecated from provider version 1.119.1. New field `routeTableName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the route table.
     */
    routeTableName?: pulumi.Input<string>;
    /**
     * (Available in v1.119.1+) The status of the route table.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The vpcId of the route table, the field can't be changed.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouteTable resource.
 */
export interface RouteTableArgs {
    /**
     * The description of the route table instance.
     */
    description?: pulumi.Input<string>;
    /**
     * Field `name` has been deprecated from provider version 1.119.1. New field `routeTableName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the route table.
     */
    routeTableName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The vpcId of the route table, the field can't be changed.
     */
    vpcId: pulumi.Input<string>;
}
