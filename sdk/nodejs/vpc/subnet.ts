// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * @deprecated This resource has been deprecated and replaced by the Switch resource.
 */
export class Subnet extends pulumi.CustomResource {
    /**
     * Get an existing Subnet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubnetState, opts?: pulumi.CustomResourceOptions): Subnet {
        pulumi.log.warn("Subnet is deprecated: This resource has been deprecated and replaced by the Switch resource.")
        return new Subnet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/subnet:Subnet';

    /**
     * Returns true if the given object is an instance of Subnet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subnet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Subnet.__pulumiType;
    }

    /**
     * @deprecated Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    public readonly cidrBlock!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * @deprecated Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly vpcId!: pulumi.Output<string>;
    public readonly vswitchName!: pulumi.Output<string>;
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a Subnet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated This resource has been deprecated and replaced by the Switch resource. */
    constructor(name: string, args: SubnetArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated This resource has been deprecated and replaced by the Switch resource. */
    constructor(name: string, argsOrState?: SubnetArgs | SubnetState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Subnet is deprecated: This resource has been deprecated and replaced by the Switch resource.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubnetState | undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vswitchName"] = state ? state.vswitchName : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as SubnetArgs | undefined;
            if ((!args || args.cidrBlock === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cidrBlock'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vswitchName"] = args ? args.vswitchName : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Subnet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subnet resources.
 */
export interface SubnetState {
    /**
     * @deprecated Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.
     */
    availabilityZone?: pulumi.Input<string>;
    cidrBlock?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    /**
     * @deprecated Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.
     */
    name?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: any}>;
    vpcId?: pulumi.Input<string>;
    vswitchName?: pulumi.Input<string>;
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Subnet resource.
 */
export interface SubnetArgs {
    /**
     * @deprecated Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.
     */
    availabilityZone?: pulumi.Input<string>;
    cidrBlock: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    /**
     * @deprecated Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.
     */
    name?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: any}>;
    vpcId: pulumi.Input<string>;
    vswitchName?: pulumi.Input<string>;
    zoneId?: pulumi.Input<string>;
}
