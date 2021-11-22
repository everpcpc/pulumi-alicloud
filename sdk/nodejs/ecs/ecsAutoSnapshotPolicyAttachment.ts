// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ECS Auto Snapshot Policy Attachment resource.
 *
 * For information about ECS Auto Snapshot Policy Attachment and how to use it, see [What is Auto Snapshot Policy Attachment](https://www.alibabacloud.com/help/en/doc-detail/25531.htm).
 *
 * > **NOTE:** Available in v1.122.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.ecs.EcsAutoSnapshotPolicyAttachment("example", {
 *     autoSnapshotPolicyId: "s-ge465xxxx",
 *     diskId: "d-gw835xxxx",
 * });
 * ```
 *
 * ## Import
 *
 * ECS Auto Snapshot Policy Attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ecs/ecsAutoSnapshotPolicyAttachment:EcsAutoSnapshotPolicyAttachment example s-abcd12345:d-abcd12345
 * ```
 */
export class EcsAutoSnapshotPolicyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing EcsAutoSnapshotPolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EcsAutoSnapshotPolicyAttachmentState, opts?: pulumi.CustomResourceOptions): EcsAutoSnapshotPolicyAttachment {
        return new EcsAutoSnapshotPolicyAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/ecsAutoSnapshotPolicyAttachment:EcsAutoSnapshotPolicyAttachment';

    /**
     * Returns true if the given object is an instance of EcsAutoSnapshotPolicyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EcsAutoSnapshotPolicyAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EcsAutoSnapshotPolicyAttachment.__pulumiType;
    }

    /**
     * The auto snapshot policy id.
     */
    public readonly autoSnapshotPolicyId!: pulumi.Output<string>;
    /**
     * The disk id.
     */
    public readonly diskId!: pulumi.Output<string>;

    /**
     * Create a EcsAutoSnapshotPolicyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EcsAutoSnapshotPolicyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EcsAutoSnapshotPolicyAttachmentArgs | EcsAutoSnapshotPolicyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EcsAutoSnapshotPolicyAttachmentState | undefined;
            inputs["autoSnapshotPolicyId"] = state ? state.autoSnapshotPolicyId : undefined;
            inputs["diskId"] = state ? state.diskId : undefined;
        } else {
            const args = argsOrState as EcsAutoSnapshotPolicyAttachmentArgs | undefined;
            if ((!args || args.autoSnapshotPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoSnapshotPolicyId'");
            }
            if ((!args || args.diskId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diskId'");
            }
            inputs["autoSnapshotPolicyId"] = args ? args.autoSnapshotPolicyId : undefined;
            inputs["diskId"] = args ? args.diskId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EcsAutoSnapshotPolicyAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EcsAutoSnapshotPolicyAttachment resources.
 */
export interface EcsAutoSnapshotPolicyAttachmentState {
    /**
     * The auto snapshot policy id.
     */
    autoSnapshotPolicyId?: pulumi.Input<string>;
    /**
     * The disk id.
     */
    diskId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EcsAutoSnapshotPolicyAttachment resource.
 */
export interface EcsAutoSnapshotPolicyAttachmentArgs {
    /**
     * The auto snapshot policy id.
     */
    autoSnapshotPolicyId: pulumi.Input<string>;
    /**
     * The disk id.
     */
    diskId: pulumi.Input<string>;
}
