// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/pvtz_zone.html.markdown.
 */
export class Zone extends pulumi.CustomResource {
    /**
     * Get an existing Zone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneState, opts?: pulumi.CustomResourceOptions): Zone {
        return new Zone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:pvtz/zone:Zone';

    /**
     * Returns true if the given object is an instance of Zone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Zone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Zone.__pulumiType;
    }

    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    public /*out*/ readonly isPtr!: pulumi.Output<boolean>;
    /**
     * The language. Valid values: "zh", "en", "jp".
     */
    public readonly lang!: pulumi.Output<string | undefined>;
    /**
     * The name of the Private Zone.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The recursive DNS proxy. Valid values:
     * - ZONE: indicates that the recursive DNS proxy is disabled.
     * - RECORD: indicates that the recursive DNS proxy is enabled.
     */
    public readonly proxyPattern!: pulumi.Output<string | undefined>;
    /**
     * The count of the Private Zone Record.
     */
    public /*out*/ readonly recordCount!: pulumi.Output<number>;
    /**
     * The remark of the Private Zone.
     */
    public readonly remark!: pulumi.Output<string | undefined>;
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The IP address of the client.
     */
    public readonly userClientIp!: pulumi.Output<string | undefined>;

    /**
     * Create a Zone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ZoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZoneArgs | ZoneState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ZoneState | undefined;
            inputs["creationTime"] = state ? state.creationTime : undefined;
            inputs["isPtr"] = state ? state.isPtr : undefined;
            inputs["lang"] = state ? state.lang : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["proxyPattern"] = state ? state.proxyPattern : undefined;
            inputs["recordCount"] = state ? state.recordCount : undefined;
            inputs["remark"] = state ? state.remark : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
            inputs["userClientIp"] = state ? state.userClientIp : undefined;
        } else {
            const args = argsOrState as ZoneArgs | undefined;
            inputs["lang"] = args ? args.lang : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["proxyPattern"] = args ? args.proxyPattern : undefined;
            inputs["remark"] = args ? args.remark : undefined;
            inputs["userClientIp"] = args ? args.userClientIp : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["isPtr"] = undefined /*out*/;
            inputs["recordCount"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Zone.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Zone resources.
 */
export interface ZoneState {
    readonly creationTime?: pulumi.Input<string>;
    readonly isPtr?: pulumi.Input<boolean>;
    /**
     * The language. Valid values: "zh", "en", "jp".
     */
    readonly lang?: pulumi.Input<string>;
    /**
     * The name of the Private Zone.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The recursive DNS proxy. Valid values:
     * - ZONE: indicates that the recursive DNS proxy is disabled.
     * - RECORD: indicates that the recursive DNS proxy is enabled.
     */
    readonly proxyPattern?: pulumi.Input<string>;
    /**
     * The count of the Private Zone Record.
     */
    readonly recordCount?: pulumi.Input<number>;
    /**
     * The remark of the Private Zone.
     */
    readonly remark?: pulumi.Input<string>;
    readonly updateTime?: pulumi.Input<string>;
    /**
     * The IP address of the client.
     */
    readonly userClientIp?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Zone resource.
 */
export interface ZoneArgs {
    /**
     * The language. Valid values: "zh", "en", "jp".
     */
    readonly lang?: pulumi.Input<string>;
    /**
     * The name of the Private Zone.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The recursive DNS proxy. Valid values:
     * - ZONE: indicates that the recursive DNS proxy is disabled.
     * - RECORD: indicates that the recursive DNS proxy is enabled.
     */
    readonly proxyPattern?: pulumi.Input<string>;
    /**
     * The remark of the Private Zone.
     */
    readonly remark?: pulumi.Input<string>;
    /**
     * The IP address of the client.
     */
    readonly userClientIp?: pulumi.Input<string>;
}
