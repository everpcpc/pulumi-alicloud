// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Connect Network Grant resource. If the CEN instance to be attached belongs to another account, authorization by the CEN instance is required.
 *
 * For information about Cloud Connect Network Grant and how to use it, see [What is Cloud Connect Network Grant](https://www.alibabacloud.com/help/doc-detail/94543.htm).
 *
 * > **NOTE:** Available in 1.63.0+
 *
 * > **NOTE:** Only the following regions support create Cloud Connect Network Grant. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ccnAccount = new alicloud.Provider("ccnAccount", {});
 * const cenAccount = new alicloud.Provider("cenAccount", {
 *     region: "cn-hangzhou",
 *     accessKey: "xxxxxx",
 *     secretKey: "xxxxxx",
 * });
 * const cen = new alicloud.cen.Instance("cen", {}, {
 *     provider: alicloud.cen_account,
 * });
 * const ccn = new alicloud.cloudconnect.Network("ccn", {isDefault: "true"}, {
 *     provider: alicloud.ccn_account,
 * });
 * const _default = new alicloud.cloudconnect.NetworkGrant("default", {
 *     ccnId: ccn.id,
 *     cenId: cen.id,
 *     cenUid: "xxxxxx",
 * }, {
 *     dependsOn: [
 *         ccn,
 *         cen,
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * The Cloud Connect Network Grant can be imported using the instance_id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cloudconnect/networkGrant:NetworkGrant example ccn-abc123456:cen-abc123456
 * ```
 */
export class NetworkGrant extends pulumi.CustomResource {
    /**
     * Get an existing NetworkGrant resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkGrantState, opts?: pulumi.CustomResourceOptions): NetworkGrant {
        return new NetworkGrant(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cloudconnect/networkGrant:NetworkGrant';

    /**
     * Returns true if the given object is an instance of NetworkGrant.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkGrant {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkGrant.__pulumiType;
    }

    /**
     * The ID of the CCN instance.
     */
    public readonly ccnId!: pulumi.Output<string>;
    /**
     * The ID of the CEN instance.
     */
    public readonly cenId!: pulumi.Output<string>;
    /**
     * The ID of the account to which the CEN instance belongs.
     */
    public readonly cenUid!: pulumi.Output<string>;

    /**
     * Create a NetworkGrant resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkGrantArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkGrantArgs | NetworkGrantState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NetworkGrantState | undefined;
            inputs["ccnId"] = state ? state.ccnId : undefined;
            inputs["cenId"] = state ? state.cenId : undefined;
            inputs["cenUid"] = state ? state.cenUid : undefined;
        } else {
            const args = argsOrState as NetworkGrantArgs | undefined;
            if ((!args || args.ccnId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'ccnId'");
            }
            if ((!args || args.cenId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'cenId'");
            }
            if ((!args || args.cenUid === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'cenUid'");
            }
            inputs["ccnId"] = args ? args.ccnId : undefined;
            inputs["cenId"] = args ? args.cenId : undefined;
            inputs["cenUid"] = args ? args.cenUid : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NetworkGrant.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkGrant resources.
 */
export interface NetworkGrantState {
    /**
     * The ID of the CCN instance.
     */
    readonly ccnId?: pulumi.Input<string>;
    /**
     * The ID of the CEN instance.
     */
    readonly cenId?: pulumi.Input<string>;
    /**
     * The ID of the account to which the CEN instance belongs.
     */
    readonly cenUid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkGrant resource.
 */
export interface NetworkGrantArgs {
    /**
     * The ID of the CCN instance.
     */
    readonly ccnId: pulumi.Input<string>;
    /**
     * The ID of the CEN instance.
     */
    readonly cenId: pulumi.Input<string>;
    /**
     * The ID of the account to which the CEN instance belongs.
     */
    readonly cenUid: pulumi.Input<string>;
}
