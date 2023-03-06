// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a DAS Switch Das Pro resource.
 *
 * For information about DAS Switch Das Pro and how to use it, see [What is Switch Das Pro](https://www.alibabacloud.com/help/en/database-autonomy-service/latest/enabledaspro).
 *
 * > **NOTE:** Available in v1.193.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _default = new alicloud.das.SwitchDasPro("default", {
 *     instanceId: "your_sql_id",
 *     sqlRetention: 30,
 *     userId: "your_account_id",
 * });
 * ```
 *
 * ## Import
 *
 * DAS Switch Das Pro can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:das/switchDasPro:SwitchDasPro example <id>
 * ```
 */
export class SwitchDasPro extends pulumi.CustomResource {
    /**
     * Get an existing SwitchDasPro resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchDasProState, opts?: pulumi.CustomResourceOptions): SwitchDasPro {
        return new SwitchDasPro(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:das/switchDasPro:SwitchDasPro';

    /**
     * Returns true if the given object is an instance of SwitchDasPro.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchDasPro {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchDasPro.__pulumiType;
    }

    /**
     * The ID of the database instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The storage duration of SQL Explorer data. Valid values: `30`, `180`, `365`, `1095`, `1825`. Unit: days. Default value: `30`.
     */
    public readonly sqlRetention!: pulumi.Output<number>;
    /**
     * Whether the database instance has DAS professional.
     */
    public /*out*/ readonly status!: pulumi.Output<boolean>;
    /**
     * The ID of the Alibaba Cloud account that is used to create the database instance.
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a SwitchDasPro resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SwitchDasProArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchDasProArgs | SwitchDasProState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchDasProState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["sqlRetention"] = state ? state.sqlRetention : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as SwitchDasProArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["sqlRetention"] = args ? args.sqlRetention : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwitchDasPro.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchDasPro resources.
 */
export interface SwitchDasProState {
    /**
     * The ID of the database instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The storage duration of SQL Explorer data. Valid values: `30`, `180`, `365`, `1095`, `1825`. Unit: days. Default value: `30`.
     */
    sqlRetention?: pulumi.Input<number>;
    /**
     * Whether the database instance has DAS professional.
     */
    status?: pulumi.Input<boolean>;
    /**
     * The ID of the Alibaba Cloud account that is used to create the database instance.
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchDasPro resource.
 */
export interface SwitchDasProArgs {
    /**
     * The ID of the database instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The storage duration of SQL Explorer data. Valid values: `30`, `180`, `365`, `1095`, `1825`. Unit: days. Default value: `30`.
     */
    sqlRetention?: pulumi.Input<number>;
    /**
     * The ID of the Alibaba Cloud account that is used to create the database instance.
     */
    userId?: pulumi.Input<string>;
}