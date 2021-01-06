// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Private Link Vpc Endpoint Service Resource resource.
 *
 * For information about Private Link Vpc Endpoint Service Resource and how to use it, see [What is Vpc Endpoint Service Resource](https://help.aliyun.com/document_detail/183548.html).
 *
 * > **NOTE:** Available in v1.110.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.privatelink.VpcEndpointServiceResource("example", {
 *     resourceId: "lb-gw8nuym5xxxxx",
 *     resourceType: "slb",
 *     serviceId: "epsrv-gw8ii1xxxx",
 * });
 * ```
 *
 * ## Import
 *
 * Private Link Vpc Endpoint Service Resource can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource example <service_id>:<resource_id>
 * ```
 */
export class VpcEndpointServiceResource extends pulumi.CustomResource {
    /**
     * Get an existing VpcEndpointServiceResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcEndpointServiceResourceState, opts?: pulumi.CustomResourceOptions): VpcEndpointServiceResource {
        return new VpcEndpointServiceResource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource';

    /**
     * Returns true if the given object is an instance of VpcEndpointServiceResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcEndpointServiceResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcEndpointServiceResource.__pulumiType;
    }

    /**
     * The dry run.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of Resource.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * The Type of Resource.
     */
    public readonly resourceType!: pulumi.Output<string>;
    /**
     * The ID of Vpc Endpoint Service.
     */
    public readonly serviceId!: pulumi.Output<string>;

    /**
     * Create a VpcEndpointServiceResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcEndpointServiceResourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcEndpointServiceResourceArgs | VpcEndpointServiceResourceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpcEndpointServiceResourceState | undefined;
            inputs["dryRun"] = state ? state.dryRun : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["resourceType"] = state ? state.resourceType : undefined;
            inputs["serviceId"] = state ? state.serviceId : undefined;
        } else {
            const args = argsOrState as VpcEndpointServiceResourceArgs | undefined;
            if (!args || args.resourceId === undefined) {
                throw new Error("Missing required property 'resourceId'");
            }
            if (!args || args.resourceType === undefined) {
                throw new Error("Missing required property 'resourceType'");
            }
            if (!args || args.serviceId === undefined) {
                throw new Error("Missing required property 'serviceId'");
            }
            inputs["dryRun"] = args ? args.dryRun : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["resourceType"] = args ? args.resourceType : undefined;
            inputs["serviceId"] = args ? args.serviceId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpcEndpointServiceResource.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcEndpointServiceResource resources.
 */
export interface VpcEndpointServiceResourceState {
    /**
     * The dry run.
     */
    readonly dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of Resource.
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * The Type of Resource.
     */
    readonly resourceType?: pulumi.Input<string>;
    /**
     * The ID of Vpc Endpoint Service.
     */
    readonly serviceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcEndpointServiceResource resource.
 */
export interface VpcEndpointServiceResourceArgs {
    /**
     * The dry run.
     */
    readonly dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of Resource.
     */
    readonly resourceId: pulumi.Input<string>;
    /**
     * The Type of Resource.
     */
    readonly resourceType: pulumi.Input<string>;
    /**
     * The ID of Vpc Endpoint Service.
     */
    readonly serviceId: pulumi.Input<string>;
}