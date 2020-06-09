// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Alidns Record resource. For information about Alidns Domain Record and how to use it, see [What is Resource Alidns Record](https://www.alibabacloud.com/help/en/doc-detail/29772.htm).
 *
 * > **NOTE:** Available in v1.85.0+.
 *
 * > **NOTE:** When the site is an international site, the `type` neither supports `REDIRECT_URL` nor `REDIRECT_URL`
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Create a new Domain Record
 * const record = new alicloud.dns.AlidnsRecord("record", {
 *     domainName: "domainname",
 *     remark: "Test new alidns record.",
 *     rr: "@",
 *     status: "ENABLE",
 *     type: "A",
 *     value: "192.168.99.99",
 * });
 * ```
 */
export class AlidnsRecord extends pulumi.CustomResource {
    /**
     * Get an existing AlidnsRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlidnsRecordState, opts?: pulumi.CustomResourceOptions): AlidnsRecord {
        return new AlidnsRecord(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dns/alidnsRecord:AlidnsRecord';

    /**
     * Returns true if the given object is an instance of AlidnsRecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlidnsRecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlidnsRecord.__pulumiType;
    }

    /**
     * Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * User language. 
     */
    public readonly lang!: pulumi.Output<string | undefined>;
    /**
     * The resolution line of domain record. Valid values: `default`, `telecom`, `unicom`, `mobile`, `oversea`, `edu`, `drpeng`, `btvn`. When the `type` is `FORWORD_URL`, this parameter must be `default`. Default value is `default`. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm) or using alicloud.dns.getResolutionLines in data source to get the value. 
     */
    public readonly line!: pulumi.Output<string | undefined>;
    /**
     * The priority of domain record. Valid values: `[1-10]`. When the `type` is `MX`, this parameter is required.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The remark of the domain record. 
     */
    public readonly remark!: pulumi.Output<string | undefined>;
    /**
     * Host record for the domain record. This hostRecord can have at most 253 characters, and each part split with `.` can have at most 63 characters, and must contain only alphanumeric characters or hyphens, such as `-`, `.`, `*`, `@`, and must not begin or end with `-`.
     */
    public readonly rr!: pulumi.Output<string>;
    /**
     * The status of the domain record. Valid values: `ENABLE`,`DISABLE`. 
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * The effective time of domain record. Its scope depends on the edition of the cloud resolution. Free is `[600, 86400]`, Basic is `[120, 86400]`, Standard is `[60, 86400]`, Ultimate is `[10, 86400]`, Exclusive is `[1, 86400]`. Default value is `600`.
     */
    public readonly ttl!: pulumi.Output<number | undefined>;
    /**
     * The type of domain record. Valid values: `A`,`NS`,`MX`,`TXT`,`CNAME`,`SRV`,`AAAA`,`CAA`, `REDIRECT_URL` and `FORWORD_URL`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The IP address of the client. 
     */
    public readonly userClientIp!: pulumi.Output<string | undefined>;
    /**
     * The value of domain record, When the `type` is `MX`,`NS`,`CNAME`,`SRV`, the server will treat the `value` as a fully qualified domain name, so it's no need to add a `.` at the end.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a AlidnsRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlidnsRecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlidnsRecordArgs | AlidnsRecordState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AlidnsRecordState | undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["lang"] = state ? state.lang : undefined;
            inputs["line"] = state ? state.line : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["remark"] = state ? state.remark : undefined;
            inputs["rr"] = state ? state.rr : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["userClientIp"] = state ? state.userClientIp : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as AlidnsRecordArgs | undefined;
            if (!args || args.domainName === undefined) {
                throw new Error("Missing required property 'domainName'");
            }
            if (!args || args.rr === undefined) {
                throw new Error("Missing required property 'rr'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            if (!args || args.value === undefined) {
                throw new Error("Missing required property 'value'");
            }
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["lang"] = args ? args.lang : undefined;
            inputs["line"] = args ? args.line : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["remark"] = args ? args.remark : undefined;
            inputs["rr"] = args ? args.rr : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["userClientIp"] = args ? args.userClientIp : undefined;
            inputs["value"] = args ? args.value : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AlidnsRecord.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlidnsRecord resources.
 */
export interface AlidnsRecordState {
    /**
     * Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    readonly domainName?: pulumi.Input<string>;
    /**
     * User language. 
     */
    readonly lang?: pulumi.Input<string>;
    /**
     * The resolution line of domain record. Valid values: `default`, `telecom`, `unicom`, `mobile`, `oversea`, `edu`, `drpeng`, `btvn`. When the `type` is `FORWORD_URL`, this parameter must be `default`. Default value is `default`. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm) or using alicloud.dns.getResolutionLines in data source to get the value. 
     */
    readonly line?: pulumi.Input<string>;
    /**
     * The priority of domain record. Valid values: `[1-10]`. When the `type` is `MX`, this parameter is required.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The remark of the domain record. 
     */
    readonly remark?: pulumi.Input<string>;
    /**
     * Host record for the domain record. This hostRecord can have at most 253 characters, and each part split with `.` can have at most 63 characters, and must contain only alphanumeric characters or hyphens, such as `-`, `.`, `*`, `@`, and must not begin or end with `-`.
     */
    readonly rr?: pulumi.Input<string>;
    /**
     * The status of the domain record. Valid values: `ENABLE`,`DISABLE`. 
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The effective time of domain record. Its scope depends on the edition of the cloud resolution. Free is `[600, 86400]`, Basic is `[120, 86400]`, Standard is `[60, 86400]`, Ultimate is `[10, 86400]`, Exclusive is `[1, 86400]`. Default value is `600`.
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * The type of domain record. Valid values: `A`,`NS`,`MX`,`TXT`,`CNAME`,`SRV`,`AAAA`,`CAA`, `REDIRECT_URL` and `FORWORD_URL`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The IP address of the client. 
     */
    readonly userClientIp?: pulumi.Input<string>;
    /**
     * The value of domain record, When the `type` is `MX`,`NS`,`CNAME`,`SRV`, the server will treat the `value` as a fully qualified domain name, so it's no need to add a `.` at the end.
     */
    readonly value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AlidnsRecord resource.
 */
export interface AlidnsRecordArgs {
    /**
     * Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    readonly domainName: pulumi.Input<string>;
    /**
     * User language. 
     */
    readonly lang?: pulumi.Input<string>;
    /**
     * The resolution line of domain record. Valid values: `default`, `telecom`, `unicom`, `mobile`, `oversea`, `edu`, `drpeng`, `btvn`. When the `type` is `FORWORD_URL`, this parameter must be `default`. Default value is `default`. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm) or using alicloud.dns.getResolutionLines in data source to get the value. 
     */
    readonly line?: pulumi.Input<string>;
    /**
     * The priority of domain record. Valid values: `[1-10]`. When the `type` is `MX`, this parameter is required.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The remark of the domain record. 
     */
    readonly remark?: pulumi.Input<string>;
    /**
     * Host record for the domain record. This hostRecord can have at most 253 characters, and each part split with `.` can have at most 63 characters, and must contain only alphanumeric characters or hyphens, such as `-`, `.`, `*`, `@`, and must not begin or end with `-`.
     */
    readonly rr: pulumi.Input<string>;
    /**
     * The status of the domain record. Valid values: `ENABLE`,`DISABLE`. 
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The effective time of domain record. Its scope depends on the edition of the cloud resolution. Free is `[600, 86400]`, Basic is `[120, 86400]`, Standard is `[60, 86400]`, Ultimate is `[10, 86400]`, Exclusive is `[1, 86400]`. Default value is `600`.
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * The type of domain record. Valid values: `A`,`NS`,`MX`,`TXT`,`CNAME`,`SRV`,`AAAA`,`CAA`, `REDIRECT_URL` and `FORWORD_URL`.
     */
    readonly type: pulumi.Input<string>;
    /**
     * The IP address of the client. 
     */
    readonly userClientIp?: pulumi.Input<string>;
    /**
     * The value of domain record, When the `type` is `MX`,`NS`,`CNAME`,`SRV`, the server will treat the `value` as a fully qualified domain name, so it's no need to add a `.` at the end.
     */
    readonly value: pulumi.Input<string>;
}
