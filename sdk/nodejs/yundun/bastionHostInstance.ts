// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class BastionHostInstance extends pulumi.CustomResource {
    /**
     * Get an existing BastionHostInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BastionHostInstanceState, opts?: pulumi.CustomResourceOptions): BastionHostInstance {
        return new BastionHostInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:yundun/bastionHostInstance:BastionHostInstance';

    /**
     * Returns true if the given object is an instance of BastionHostInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BastionHostInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BastionHostInstance.__pulumiType;
    }

    public readonly description!: pulumi.Output<string>;
    public readonly licenseCode!: pulumi.Output<string>;
    public readonly period!: pulumi.Output<number | undefined>;
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly vswitchId!: pulumi.Output<string>;

    /**
     * Create a BastionHostInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BastionHostInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BastionHostInstanceArgs | BastionHostInstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BastionHostInstanceState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["licenseCode"] = state ? state.licenseCode : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as BastionHostInstanceArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.licenseCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'licenseCode'");
            }
            if ((!args || args.securityGroupIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupIds'");
            }
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["licenseCode"] = args ? args.licenseCode : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BastionHostInstance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BastionHostInstance resources.
 */
export interface BastionHostInstanceState {
    description?: pulumi.Input<string>;
    licenseCode?: pulumi.Input<string>;
    period?: pulumi.Input<number>;
    resourceGroupId?: pulumi.Input<string>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    tags?: pulumi.Input<{[key: string]: any}>;
    vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BastionHostInstance resource.
 */
export interface BastionHostInstanceArgs {
    description: pulumi.Input<string>;
    licenseCode: pulumi.Input<string>;
    period?: pulumi.Input<number>;
    resourceGroupId?: pulumi.Input<string>;
    securityGroupIds: pulumi.Input<pulumi.Input<string>[]>;
    tags?: pulumi.Input<{[key: string]: any}>;
    vswitchId: pulumi.Input<string>;
}
