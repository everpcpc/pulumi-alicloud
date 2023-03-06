// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Alb Load Balancer Common Bandwidth Package Attachment resource.
 *
 * For information about Alb Load Balancer Common Bandwidth Package Attachment and how to use it, see [What is Load Balancer Common Bandwidth Package Attachment](https://www.alibabacloud.com/help/en/server-load-balancer/latest/attachcommonbandwidthpackagetoloadbalancer).
 *
 * > **NOTE:** Available in v1.200.0+.
 *
 * ## Import
 *
 * Alb Load Balancer Common Bandwidth Package Attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:alb/loadBalancerCommonBandwidthPackageAttachment:LoadBalancerCommonBandwidthPackageAttachment example <load_balancer_id>:<bandwidth_package_id>
 * ```
 */
export class LoadBalancerCommonBandwidthPackageAttachment extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancerCommonBandwidthPackageAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerCommonBandwidthPackageAttachmentState, opts?: pulumi.CustomResourceOptions): LoadBalancerCommonBandwidthPackageAttachment {
        return new LoadBalancerCommonBandwidthPackageAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:alb/loadBalancerCommonBandwidthPackageAttachment:LoadBalancerCommonBandwidthPackageAttachment';

    /**
     * Returns true if the given object is an instance of LoadBalancerCommonBandwidthPackageAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancerCommonBandwidthPackageAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancerCommonBandwidthPackageAttachment.__pulumiType;
    }

    /**
     * The ID of the bound shared bandwidth package.
     */
    public readonly bandwidthPackageId!: pulumi.Output<string>;
    /**
     * Whether to PreCheck this request only. Value:-**true**: sends a check request and does not bind the shared bandwidth package to the load balancing instance. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the applied server load balancer instance.
     */
    public readonly loadBalancerId!: pulumi.Output<string>;
    /**
     * The status of the Application Load balancing instance. Value:-**Inactive**: Stopped, indicating that the instance listener will no longer forward traffic.-**Active**: running.-**Provisioning**: The project is being created.-**Configuring**: The configuration is being changed.-**CreateFailed**: The instance cannot be deleted without any charge.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a LoadBalancerCommonBandwidthPackageAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerCommonBandwidthPackageAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadBalancerCommonBandwidthPackageAttachmentArgs | LoadBalancerCommonBandwidthPackageAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadBalancerCommonBandwidthPackageAttachmentState | undefined;
            resourceInputs["bandwidthPackageId"] = state ? state.bandwidthPackageId : undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["loadBalancerId"] = state ? state.loadBalancerId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as LoadBalancerCommonBandwidthPackageAttachmentArgs | undefined;
            if ((!args || args.bandwidthPackageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidthPackageId'");
            }
            if ((!args || args.loadBalancerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerId'");
            }
            resourceInputs["bandwidthPackageId"] = args ? args.bandwidthPackageId : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["loadBalancerId"] = args ? args.loadBalancerId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadBalancerCommonBandwidthPackageAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancerCommonBandwidthPackageAttachment resources.
 */
export interface LoadBalancerCommonBandwidthPackageAttachmentState {
    /**
     * The ID of the bound shared bandwidth package.
     */
    bandwidthPackageId?: pulumi.Input<string>;
    /**
     * Whether to PreCheck this request only. Value:-**true**: sends a check request and does not bind the shared bandwidth package to the load balancing instance. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the applied server load balancer instance.
     */
    loadBalancerId?: pulumi.Input<string>;
    /**
     * The status of the Application Load balancing instance. Value:-**Inactive**: Stopped, indicating that the instance listener will no longer forward traffic.-**Active**: running.-**Provisioning**: The project is being created.-**Configuring**: The configuration is being changed.-**CreateFailed**: The instance cannot be deleted without any charge.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadBalancerCommonBandwidthPackageAttachment resource.
 */
export interface LoadBalancerCommonBandwidthPackageAttachmentArgs {
    /**
     * The ID of the bound shared bandwidth package.
     */
    bandwidthPackageId: pulumi.Input<string>;
    /**
     * Whether to PreCheck this request only. Value:-**true**: sends a check request and does not bind the shared bandwidth package to the load balancing instance. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the applied server load balancer instance.
     */
    loadBalancerId: pulumi.Input<string>;
}
