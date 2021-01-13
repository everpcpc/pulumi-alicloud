// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Private Zone Record can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:pvtz/zoneRecord:ZoneRecord example abc123456
 * ```
 */
export class ZoneRecord extends pulumi.CustomResource {
    /**
     * Get an existing ZoneRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneRecordState, opts?: pulumi.CustomResourceOptions): ZoneRecord {
        return new ZoneRecord(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:pvtz/zoneRecord:ZoneRecord';

    /**
     * Returns true if the given object is an instance of ZoneRecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ZoneRecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ZoneRecord.__pulumiType;
    }

    /**
     * User language.
     */
    public readonly lang!: pulumi.Output<string | undefined>;
    /**
     * The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-99]. Default to 1.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The Private Zone Record ID.
     */
    public /*out*/ readonly recordId!: pulumi.Output<number>;
    /**
     * The remark of the Private Zone Record.
     */
    public readonly remark!: pulumi.Output<string | undefined>;
    /**
     * The resource record of the Private Zone Record.
     *
     * @deprecated Field 'resource_record' has been deprecated from version 1.109.0. Use 'rr' instead.
     */
    public readonly resourceRecord!: pulumi.Output<string>;
    /**
     * The rr of the Private Zone Record.
     */
    public readonly rr!: pulumi.Output<string>;
    /**
     * Resolve record status. Value:
     * - ENABLE: enable resolution.
     * - DISABLE: pause parsing.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * The ttl of the Private Zone Record. Default to `60`.
     */
    public readonly ttl!: pulumi.Output<number | undefined>;
    /**
     * The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
     */
    public readonly type!: pulumi.Output<string>;
    public readonly userClientIp!: pulumi.Output<string | undefined>;
    /**
     * The value of the Private Zone Record.
     */
    public readonly value!: pulumi.Output<string>;
    /**
     * The name of the Private Zone Record.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a ZoneRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZoneRecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZoneRecordArgs | ZoneRecordState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ZoneRecordState | undefined;
            inputs["lang"] = state ? state.lang : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["recordId"] = state ? state.recordId : undefined;
            inputs["remark"] = state ? state.remark : undefined;
            inputs["resourceRecord"] = state ? state.resourceRecord : undefined;
            inputs["rr"] = state ? state.rr : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["userClientIp"] = state ? state.userClientIp : undefined;
            inputs["value"] = state ? state.value : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ZoneRecordArgs | undefined;
            if ((!args || args.type === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.value === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'value'");
            }
            if ((!args || args.zoneId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'zoneId'");
            }
            inputs["lang"] = args ? args.lang : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["remark"] = args ? args.remark : undefined;
            inputs["resourceRecord"] = args ? args.resourceRecord : undefined;
            inputs["rr"] = args ? args.rr : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["userClientIp"] = args ? args.userClientIp : undefined;
            inputs["value"] = args ? args.value : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
            inputs["recordId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ZoneRecord.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ZoneRecord resources.
 */
export interface ZoneRecordState {
    /**
     * User language.
     */
    readonly lang?: pulumi.Input<string>;
    /**
     * The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-99]. Default to 1.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The Private Zone Record ID.
     */
    readonly recordId?: pulumi.Input<number>;
    /**
     * The remark of the Private Zone Record.
     */
    readonly remark?: pulumi.Input<string>;
    /**
     * The resource record of the Private Zone Record.
     *
     * @deprecated Field 'resource_record' has been deprecated from version 1.109.0. Use 'rr' instead.
     */
    readonly resourceRecord?: pulumi.Input<string>;
    /**
     * The rr of the Private Zone Record.
     */
    readonly rr?: pulumi.Input<string>;
    /**
     * Resolve record status. Value:
     * - ENABLE: enable resolution.
     * - DISABLE: pause parsing.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The ttl of the Private Zone Record. Default to `60`.
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
     */
    readonly type?: pulumi.Input<string>;
    readonly userClientIp?: pulumi.Input<string>;
    /**
     * The value of the Private Zone Record.
     */
    readonly value?: pulumi.Input<string>;
    /**
     * The name of the Private Zone Record.
     */
    readonly zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ZoneRecord resource.
 */
export interface ZoneRecordArgs {
    /**
     * User language.
     */
    readonly lang?: pulumi.Input<string>;
    /**
     * The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-99]. Default to 1.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The remark of the Private Zone Record.
     */
    readonly remark?: pulumi.Input<string>;
    /**
     * The resource record of the Private Zone Record.
     *
     * @deprecated Field 'resource_record' has been deprecated from version 1.109.0. Use 'rr' instead.
     */
    readonly resourceRecord?: pulumi.Input<string>;
    /**
     * The rr of the Private Zone Record.
     */
    readonly rr?: pulumi.Input<string>;
    /**
     * Resolve record status. Value:
     * - ENABLE: enable resolution.
     * - DISABLE: pause parsing.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The ttl of the Private Zone Record. Default to `60`.
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
     */
    readonly type: pulumi.Input<string>;
    readonly userClientIp?: pulumi.Input<string>;
    /**
     * The value of the Private Zone Record.
     */
    readonly value: pulumi.Input<string>;
    /**
     * The name of the Private Zone Record.
     */
    readonly zoneId: pulumi.Input<string>;
}
