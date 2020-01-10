// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cen_instance_attachment.html.markdown.
 */
export class InstanceAttachment extends pulumi.CustomResource {
    /**
     * Get an existing InstanceAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceAttachmentState, opts?: pulumi.CustomResourceOptions): InstanceAttachment {
        return new InstanceAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cen/instanceAttachment:InstanceAttachment';

    /**
     * Returns true if the given object is an instance of InstanceAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceAttachment.__pulumiType;
    }

    /**
     * The ID of the child instance to attach.
     */
    public readonly childInstanceId!: pulumi.Output<string>;
    /**
     * The uid of the child instance. Only used when attach a child instance of other account.
     */
    public readonly childInstanceOwnerId!: pulumi.Output<string>;
    /**
     * The region ID of the child instance to attach.
     */
    public readonly childInstanceRegionId!: pulumi.Output<string>;
    /**
     * The ID of the CEN.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a InstanceAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceAttachmentArgs | InstanceAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceAttachmentState | undefined;
            inputs["childInstanceId"] = state ? state.childInstanceId : undefined;
            inputs["childInstanceOwnerId"] = state ? state.childInstanceOwnerId : undefined;
            inputs["childInstanceRegionId"] = state ? state.childInstanceRegionId : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as InstanceAttachmentArgs | undefined;
            if (!args || args.childInstanceId === undefined) {
                throw new Error("Missing required property 'childInstanceId'");
            }
            if (!args || args.childInstanceRegionId === undefined) {
                throw new Error("Missing required property 'childInstanceRegionId'");
            }
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            inputs["childInstanceId"] = args ? args.childInstanceId : undefined;
            inputs["childInstanceOwnerId"] = args ? args.childInstanceOwnerId : undefined;
            inputs["childInstanceRegionId"] = args ? args.childInstanceRegionId : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InstanceAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceAttachment resources.
 */
export interface InstanceAttachmentState {
    /**
     * The ID of the child instance to attach.
     */
    readonly childInstanceId?: pulumi.Input<string>;
    /**
     * The uid of the child instance. Only used when attach a child instance of other account.
     */
    readonly childInstanceOwnerId?: pulumi.Input<string>;
    /**
     * The region ID of the child instance to attach.
     */
    readonly childInstanceRegionId?: pulumi.Input<string>;
    /**
     * The ID of the CEN.
     */
    readonly instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceAttachment resource.
 */
export interface InstanceAttachmentArgs {
    /**
     * The ID of the child instance to attach.
     */
    readonly childInstanceId: pulumi.Input<string>;
    /**
     * The uid of the child instance. Only used when attach a child instance of other account.
     */
    readonly childInstanceOwnerId?: pulumi.Input<string>;
    /**
     * The region ID of the child instance to attach.
     */
    readonly childInstanceRegionId: pulumi.Input<string>;
    /**
     * The ID of the CEN.
     */
    readonly instanceId: pulumi.Input<string>;
}
