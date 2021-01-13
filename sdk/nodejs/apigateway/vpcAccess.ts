// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Api gateway app can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:apigateway/vpcAccess:VpcAccess example "APiGatewayVpc:vpc-aswcj19ajsz:i-ajdjfsdlf:8080"
 * ```
 */
export class VpcAccess extends pulumi.CustomResource {
    /**
     * Get an existing VpcAccess resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcAccessState, opts?: pulumi.CustomResourceOptions): VpcAccess {
        return new VpcAccess(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:apigateway/vpcAccess:VpcAccess';

    /**
     * Returns true if the given object is an instance of VpcAccess.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcAccess {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcAccess.__pulumiType;
    }

    /**
     * ID of the instance in VPC (ECS/Server Load Balance).
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The name of the vpc authorization.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the port corresponding to the instance.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * The vpc id of the vpc authorization.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a VpcAccess resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcAccessArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcAccessArgs | VpcAccessState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpcAccessState | undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as VpcAccessArgs | undefined;
            if ((!args || args.instanceId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.port === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.vpcId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpcAccess.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcAccess resources.
 */
export interface VpcAccessState {
    /**
     * ID of the instance in VPC (ECS/Server Load Balance).
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * The name of the vpc authorization.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * ID of the port corresponding to the instance.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The vpc id of the vpc authorization.
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcAccess resource.
 */
export interface VpcAccessArgs {
    /**
     * ID of the instance in VPC (ECS/Server Load Balance).
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * The name of the vpc authorization.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * ID of the port corresponding to the instance.
     */
    readonly port: pulumi.Input<number>;
    /**
     * The vpc id of the vpc authorization.
     */
    readonly vpcId: pulumi.Input<string>;
}
