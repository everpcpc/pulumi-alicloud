// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Alikms Key Version resource. For information about Alikms Key Version and how to use it, see [What is Resource Alikms Key Version](https://www.alibabacloud.com/help/doc-detail/133838.htm).
 *
 * > **NOTE:** Available in v1.85.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _this = new alicloud.kms.Key("this", {});
 * const keyversion = new alicloud.kms.KeyVersion("keyversion", {keyId: _this.id});
 * ```
 *
 * ## Import
 *
 * Alikms key version can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:kms/keyVersion:KeyVersion example 72da539a-2fa8-4f2d-b854-*****	
 * ```
 */
export class KeyVersion extends pulumi.CustomResource {
    /**
     * Get an existing KeyVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyVersionState, opts?: pulumi.CustomResourceOptions): KeyVersion {
        return new KeyVersion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:kms/keyVersion:KeyVersion';

    /**
     * Returns true if the given object is an instance of KeyVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeyVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeyVersion.__pulumiType;
    }

    /**
     * The date and time (UTC time) when the Alikms key version was created.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * The id of the master key (CMK).
     */
    public readonly keyId!: pulumi.Output<string>;
    /**
     * The id of the Alikms key version.
     */
    public /*out*/ readonly keyVersionId!: pulumi.Output<string>;

    /**
     * Create a KeyVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyVersionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyVersionArgs | KeyVersionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as KeyVersionState | undefined;
            inputs["creationDate"] = state ? state.creationDate : undefined;
            inputs["keyId"] = state ? state.keyId : undefined;
            inputs["keyVersionId"] = state ? state.keyVersionId : undefined;
        } else {
            const args = argsOrState as KeyVersionArgs | undefined;
            if ((!args || args.keyId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'keyId'");
            }
            inputs["keyId"] = args ? args.keyId : undefined;
            inputs["creationDate"] = undefined /*out*/;
            inputs["keyVersionId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(KeyVersion.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeyVersion resources.
 */
export interface KeyVersionState {
    /**
     * The date and time (UTC time) when the Alikms key version was created.
     */
    readonly creationDate?: pulumi.Input<string>;
    /**
     * The id of the master key (CMK).
     */
    readonly keyId?: pulumi.Input<string>;
    /**
     * The id of the Alikms key version.
     */
    readonly keyVersionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KeyVersion resource.
 */
export interface KeyVersionArgs {
    /**
     * The id of the master key (CMK).
     */
    readonly keyId: pulumi.Input<string>;
}
