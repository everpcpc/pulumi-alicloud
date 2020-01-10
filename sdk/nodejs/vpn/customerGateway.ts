// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/vpn_customer_gateway.html.markdown.
 */
export class CustomerGateway extends pulumi.CustomResource {
    /**
     * Get an existing CustomerGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomerGatewayState, opts?: pulumi.CustomResourceOptions): CustomerGateway {
        return new CustomerGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpn/customerGateway:CustomerGateway';

    /**
     * Returns true if the given object is an instance of CustomerGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomerGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomerGateway.__pulumiType;
    }

    /**
     * The description of the VPN customer gateway instance.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The IP address of the customer gateway.
     */
    public readonly ipAddress!: pulumi.Output<string>;
    /**
     * The name of the VPN customer gateway. Defaults to null.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a CustomerGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomerGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomerGatewayArgs | CustomerGatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CustomerGatewayState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as CustomerGatewayArgs | undefined;
            if (!args || args.ipAddress === undefined) {
                throw new Error("Missing required property 'ipAddress'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["ipAddress"] = args ? args.ipAddress : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CustomerGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomerGateway resources.
 */
export interface CustomerGatewayState {
    /**
     * The description of the VPN customer gateway instance.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The IP address of the customer gateway.
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * The name of the VPN customer gateway. Defaults to null.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomerGateway resource.
 */
export interface CustomerGatewayArgs {
    /**
     * The description of the VPN customer gateway instance.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The IP address of the customer gateway.
     */
    readonly ipAddress: pulumi.Input<string>;
    /**
     * The name of the VPN customer gateway. Defaults to null.
     */
    readonly name?: pulumi.Input<string>;
}
