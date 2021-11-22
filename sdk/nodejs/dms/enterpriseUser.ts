// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a DMS Enterprise User resource. For information about Alidms Enterprise User and how to use it, see [What is Resource Alidms Enterprise User](https://www.alibabacloud.com/help/doc-detail/98001.htm).
 *
 * > **NOTE:** Available in 1.90.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.dms.EnterpriseUser("example", {
 *     mobile: "1591066xxxx",
 *     roleNames: ["DBA"],
 *     uid: "uid",
 *     userName: "tf-test",
 * });
 * ```
 *
 * ## Import
 *
 * DMS Enterprise User can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:dms/enterpriseUser:EnterpriseUser example 24356xxx
 * ```
 */
export class EnterpriseUser extends pulumi.CustomResource {
    /**
     * Get an existing EnterpriseUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnterpriseUserState, opts?: pulumi.CustomResourceOptions): EnterpriseUser {
        return new EnterpriseUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dms/enterpriseUser:EnterpriseUser';

    /**
     * Returns true if the given object is an instance of EnterpriseUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EnterpriseUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EnterpriseUser.__pulumiType;
    }

    /**
     * Maximum number of inquiries on the day.
     */
    public readonly maxExecuteCount!: pulumi.Output<number | undefined>;
    /**
     * Query the maximum number of rows on the day.
     */
    public readonly maxResultCount!: pulumi.Output<number | undefined>;
    /**
     * The DingTalk number or mobile number of the user.
     */
    public readonly mobile!: pulumi.Output<string | undefined>;
    /**
     * It has been deprecated from 1.100.0 and use `userName` instead.
     *
     * @deprecated Field 'nick_name' has been deprecated from version 1.100.0. Use 'user_name' instead.
     */
    public readonly nickName!: pulumi.Output<string>;
    /**
     * The roles that the user plays.
     */
    public readonly roleNames!: pulumi.Output<string[] | undefined>;
    /**
     * The state of DMS Enterprise User. Valid values: `NORMAL`, `DISABLE`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * The tenant ID.
     */
    public readonly tid!: pulumi.Output<number | undefined>;
    /**
     * The Alibaba Cloud unique ID (UID) of the user to add.
     */
    public readonly uid!: pulumi.Output<string>;
    /**
     * The nickname of the user.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a EnterpriseUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnterpriseUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnterpriseUserArgs | EnterpriseUserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EnterpriseUserState | undefined;
            inputs["maxExecuteCount"] = state ? state.maxExecuteCount : undefined;
            inputs["maxResultCount"] = state ? state.maxResultCount : undefined;
            inputs["mobile"] = state ? state.mobile : undefined;
            inputs["nickName"] = state ? state.nickName : undefined;
            inputs["roleNames"] = state ? state.roleNames : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tid"] = state ? state.tid : undefined;
            inputs["uid"] = state ? state.uid : undefined;
            inputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as EnterpriseUserArgs | undefined;
            if ((!args || args.uid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'uid'");
            }
            inputs["maxExecuteCount"] = args ? args.maxExecuteCount : undefined;
            inputs["maxResultCount"] = args ? args.maxResultCount : undefined;
            inputs["mobile"] = args ? args.mobile : undefined;
            inputs["nickName"] = args ? args.nickName : undefined;
            inputs["roleNames"] = args ? args.roleNames : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tid"] = args ? args.tid : undefined;
            inputs["uid"] = args ? args.uid : undefined;
            inputs["userName"] = args ? args.userName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EnterpriseUser.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EnterpriseUser resources.
 */
export interface EnterpriseUserState {
    /**
     * Maximum number of inquiries on the day.
     */
    maxExecuteCount?: pulumi.Input<number>;
    /**
     * Query the maximum number of rows on the day.
     */
    maxResultCount?: pulumi.Input<number>;
    /**
     * The DingTalk number or mobile number of the user.
     */
    mobile?: pulumi.Input<string>;
    /**
     * It has been deprecated from 1.100.0 and use `userName` instead.
     *
     * @deprecated Field 'nick_name' has been deprecated from version 1.100.0. Use 'user_name' instead.
     */
    nickName?: pulumi.Input<string>;
    /**
     * The roles that the user plays.
     */
    roleNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The state of DMS Enterprise User. Valid values: `NORMAL`, `DISABLE`.
     */
    status?: pulumi.Input<string>;
    /**
     * The tenant ID.
     */
    tid?: pulumi.Input<number>;
    /**
     * The Alibaba Cloud unique ID (UID) of the user to add.
     */
    uid?: pulumi.Input<string>;
    /**
     * The nickname of the user.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EnterpriseUser resource.
 */
export interface EnterpriseUserArgs {
    /**
     * Maximum number of inquiries on the day.
     */
    maxExecuteCount?: pulumi.Input<number>;
    /**
     * Query the maximum number of rows on the day.
     */
    maxResultCount?: pulumi.Input<number>;
    /**
     * The DingTalk number or mobile number of the user.
     */
    mobile?: pulumi.Input<string>;
    /**
     * It has been deprecated from 1.100.0 and use `userName` instead.
     *
     * @deprecated Field 'nick_name' has been deprecated from version 1.100.0. Use 'user_name' instead.
     */
    nickName?: pulumi.Input<string>;
    /**
     * The roles that the user plays.
     */
    roleNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The state of DMS Enterprise User. Valid values: `NORMAL`, `DISABLE`.
     */
    status?: pulumi.Input<string>;
    /**
     * The tenant ID.
     */
    tid?: pulumi.Input<number>;
    /**
     * The Alibaba Cloud unique ID (UID) of the user to add.
     */
    uid: pulumi.Input<string>;
    /**
     * The nickname of the user.
     */
    userName?: pulumi.Input<string>;
}
