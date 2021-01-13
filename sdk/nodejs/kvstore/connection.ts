// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Operate the public network ip of the specified resource. How to use it, see [What is Resource Alicloud KVStore Connection](https://www.alibabacloud.com/help/doc-detail/125795.htm).
 *
 * > **NOTE:** Available in v1.101.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultConnection = new alicloud.kvstore.Connection("default", {
 *     connectionStringPrefix: "allocatetestupdate",
 *     instanceId: "r-abc123456",
 *     port: "6370",
 * });
 * ```
 *
 * ## Import
 *
 * KVStore connection can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:kvstore/connection:Connection example r-abc12345678
 * ```
 */
export class Connection extends pulumi.CustomResource {
    /**
     * Get an existing Connection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectionState, opts?: pulumi.CustomResourceOptions): Connection {
        return new Connection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:kvstore/connection:Connection';

    /**
     * Returns true if the given object is an instance of Connection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connection.__pulumiType;
    }

    /**
     * The public connection string of KVStore DBInstance.
     */
    public /*out*/ readonly connectionString!: pulumi.Output<string>;
    /**
     * The prefix of the public endpoint. The prefix can be 8 to 64 characters in length, and can contain lowercase letters and digits. It must start with a lowercase letter.
     */
    public readonly connectionStringPrefix!: pulumi.Output<string>;
    /**
     * The ID of the instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The service port number of the instance.
     */
    public readonly port!: pulumi.Output<string>;

    /**
     * Create a Connection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectionArgs | ConnectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ConnectionState | undefined;
            inputs["connectionString"] = state ? state.connectionString : undefined;
            inputs["connectionStringPrefix"] = state ? state.connectionStringPrefix : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["port"] = state ? state.port : undefined;
        } else {
            const args = argsOrState as ConnectionArgs | undefined;
            if ((!args || args.connectionStringPrefix === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'connectionStringPrefix'");
            }
            if ((!args || args.instanceId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.port === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'port'");
            }
            inputs["connectionStringPrefix"] = args ? args.connectionStringPrefix : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["connectionString"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Connection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Connection resources.
 */
export interface ConnectionState {
    /**
     * The public connection string of KVStore DBInstance.
     */
    readonly connectionString?: pulumi.Input<string>;
    /**
     * The prefix of the public endpoint. The prefix can be 8 to 64 characters in length, and can contain lowercase letters and digits. It must start with a lowercase letter.
     */
    readonly connectionStringPrefix?: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * The service port number of the instance.
     */
    readonly port?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Connection resource.
 */
export interface ConnectionArgs {
    /**
     * The prefix of the public endpoint. The prefix can be 8 to 64 characters in length, and can contain lowercase letters and digits. It must start with a lowercase letter.
     */
    readonly connectionStringPrefix: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * The service port number of the instance.
     */
    readonly port: pulumi.Input<string>;
}
