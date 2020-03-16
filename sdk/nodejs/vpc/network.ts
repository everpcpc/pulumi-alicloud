// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Network extends pulumi.CustomResource {
    /**
     * Get an existing Network resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
     * The CIDR block for the VPC.
     */
    public readonly cidrBlock!: pulumi.Output<string>;
    /**
     * The VPC description. Defaults to null.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the VPC. Defaults to null.
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
    public /*out*/ readonly routerTableId!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Network resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkArgs | NetworkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NetworkState | undefined;
            inputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["routeTableId"] = state ? state.routeTableId : undefined;
            inputs["routerId"] = state ? state.routerId : undefined;
            inputs["routerTableId"] = state ? state.routerTableId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as NetworkArgs | undefined;
            if (!args || args.cidrBlock === undefined) {
                throw new Error("Missing required property 'cidrBlock'");
            }
            inputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["routeTableId"] = undefined /*out*/;
            inputs["routerId"] = undefined /*out*/;
            inputs["routerTableId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Network.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Network resources.
 */
export interface NetworkState {
    /**
     * The CIDR block for the VPC.
     */
    readonly cidrBlock?: pulumi.Input<string>;
    /**
     * The VPC description. Defaults to null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the VPC. Defaults to null.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Id of resource group which the VPC belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The route table ID of the router created by default on VPC creation.
     */
    readonly routeTableId?: pulumi.Input<string>;
    /**
     * The ID of the router created by default on VPC creation.
     */
    readonly routerId?: pulumi.Input<string>;
    readonly routerTableId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Network resource.
 */
export interface NetworkArgs {
    /**
     * The CIDR block for the VPC.
     */
    readonly cidrBlock: pulumi.Input<string>;
    /**
     * The VPC description. Defaults to null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the VPC. Defaults to null.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Id of resource group which the VPC belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
