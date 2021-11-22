// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Resource Manager handshake resource. You can invite accounts to join a resource directory for unified management.
 * For information about Resource Manager handshake and how to use it, see [What is Resource Manager handshake](https://www.alibabacloud.com/help/en/doc-detail/135287.htm).
 *
 * > **NOTE:** Available in v1.82.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Add a Resource Manager handshake.
 * const example = new alicloud.resourcemanager.Handshake("example", {
 *     note: "test resource manager handshake",
 *     targetEntity: "1182775234******",
 *     targetType: "Account",
 * });
 * ```
 *
 * ## Import
 *
 * Resource Manager handshake can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:resourcemanager/handshake:Handshake example h-QmdexeFm1kE*****
 * ```
 */
export class Handshake extends pulumi.CustomResource {
    /**
     * Get an existing Handshake resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HandshakeState, opts?: pulumi.CustomResourceOptions): Handshake {
        return new Handshake(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:resourcemanager/handshake:Handshake';

    /**
     * Returns true if the given object is an instance of Handshake.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Handshake {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Handshake.__pulumiType;
    }

    /**
     * The expiration time of the invitation.
     */
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    /**
     * Resource account master account ID.
     */
    public /*out*/ readonly masterAccountId!: pulumi.Output<string>;
    /**
     * The name of the main account of the resource directory.
     */
    public /*out*/ readonly masterAccountName!: pulumi.Output<string>;
    /**
     * The modification time of the invitation.
     */
    public /*out*/ readonly modifyTime!: pulumi.Output<string>;
    /**
     * Remarks. The maximum length is 1024 characters.
     */
    public readonly note!: pulumi.Output<string | undefined>;
    /**
     * Resource directory ID.
     */
    public /*out*/ readonly resourceDirectoryId!: pulumi.Output<string>;
    /**
     * Invitation status. Valid values: `Pending` waiting for confirmation, `Accepted`, `Cancelled`, `Declined`, `Expired`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Invited account ID or login email.
     */
    public readonly targetEntity!: pulumi.Output<string>;
    /**
     * Type of account being invited. Valid values: `Account`, `Email`.
     */
    public readonly targetType!: pulumi.Output<string>;

    /**
     * Create a Handshake resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HandshakeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HandshakeArgs | HandshakeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HandshakeState | undefined;
            inputs["expireTime"] = state ? state.expireTime : undefined;
            inputs["masterAccountId"] = state ? state.masterAccountId : undefined;
            inputs["masterAccountName"] = state ? state.masterAccountName : undefined;
            inputs["modifyTime"] = state ? state.modifyTime : undefined;
            inputs["note"] = state ? state.note : undefined;
            inputs["resourceDirectoryId"] = state ? state.resourceDirectoryId : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["targetEntity"] = state ? state.targetEntity : undefined;
            inputs["targetType"] = state ? state.targetType : undefined;
        } else {
            const args = argsOrState as HandshakeArgs | undefined;
            if ((!args || args.targetEntity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetEntity'");
            }
            if ((!args || args.targetType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetType'");
            }
            inputs["note"] = args ? args.note : undefined;
            inputs["targetEntity"] = args ? args.targetEntity : undefined;
            inputs["targetType"] = args ? args.targetType : undefined;
            inputs["expireTime"] = undefined /*out*/;
            inputs["masterAccountId"] = undefined /*out*/;
            inputs["masterAccountName"] = undefined /*out*/;
            inputs["modifyTime"] = undefined /*out*/;
            inputs["resourceDirectoryId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Handshake.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Handshake resources.
 */
export interface HandshakeState {
    /**
     * The expiration time of the invitation.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * Resource account master account ID.
     */
    masterAccountId?: pulumi.Input<string>;
    /**
     * The name of the main account of the resource directory.
     */
    masterAccountName?: pulumi.Input<string>;
    /**
     * The modification time of the invitation.
     */
    modifyTime?: pulumi.Input<string>;
    /**
     * Remarks. The maximum length is 1024 characters.
     */
    note?: pulumi.Input<string>;
    /**
     * Resource directory ID.
     */
    resourceDirectoryId?: pulumi.Input<string>;
    /**
     * Invitation status. Valid values: `Pending` waiting for confirmation, `Accepted`, `Cancelled`, `Declined`, `Expired`.
     */
    status?: pulumi.Input<string>;
    /**
     * Invited account ID or login email.
     */
    targetEntity?: pulumi.Input<string>;
    /**
     * Type of account being invited. Valid values: `Account`, `Email`.
     */
    targetType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Handshake resource.
 */
export interface HandshakeArgs {
    /**
     * Remarks. The maximum length is 1024 characters.
     */
    note?: pulumi.Input<string>;
    /**
     * Invited account ID or login email.
     */
    targetEntity: pulumi.Input<string>;
    /**
     * Type of account being invited. Valid values: `Account`, `Email`.
     */
    targetType: pulumi.Input<string>;
}
