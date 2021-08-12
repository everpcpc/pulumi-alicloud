// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Elastic Desktop Service(EDS) Policy Group resource.
 *
 * For information about Elastic Desktop Service(EDS) Policy Group and how to use it, see [What is Policy Group](https://help.aliyun.com/document_detail/188382.html).
 *
 * > **NOTE:** Available in v1.130.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultEcdPolicyGroup = new alicloud.eds.EcdPolicyGroup("default", {
 *     authorizeAccessPolicyRules: [{
 *         cidrIp: "1.2.3.45/24",
 *         description: "my-description1",
 *     }],
 *     authorizeSecurityPolicyRules: [{
 *         cidrIp: "1.2.3.4/24",
 *         description: "my-description",
 *         ipProtocol: "TCP",
 *         policy: "accept",
 *         portRange: "80/80",
 *         priority: "1",
 *         type: "inflow",
 *     }],
 *     clipboard: "read",
 *     localDrive: "read",
 *     policyGroupName: "my-policy-group",
 *     usbRedirect: "off",
 *     watermark: "off",
 * });
 * ```
 *
 * ## Import
 *
 * Elastic Desktop Service(EDS) Policy Group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:eds/ecdPolicyGroup:EcdPolicyGroup example <id>
 * ```
 */
export class EcdPolicyGroup extends pulumi.CustomResource {
    /**
     * Get an existing EcdPolicyGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EcdPolicyGroupState, opts?: pulumi.CustomResourceOptions): EcdPolicyGroup {
        return new EcdPolicyGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eds/ecdPolicyGroup:EcdPolicyGroup';

    /**
     * Returns true if the given object is an instance of EcdPolicyGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EcdPolicyGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EcdPolicyGroup.__pulumiType;
    }

    /**
     * The rule of authorize access rule.
     */
    public readonly authorizeAccessPolicyRules!: pulumi.Output<outputs.eds.EcdPolicyGroupAuthorizeAccessPolicyRule[] | undefined>;
    /**
     * The policy rule.
     */
    public readonly authorizeSecurityPolicyRules!: pulumi.Output<outputs.eds.EcdPolicyGroupAuthorizeSecurityPolicyRule[] | undefined>;
    /**
     * The clipboard policy. Valid values: `off`, `read`, `readwrite`.
     */
    public readonly clipboard!: pulumi.Output<string>;
    /**
     * The list of domain.
     */
    public readonly domainList!: pulumi.Output<string | undefined>;
    /**
     * The access of html5. Valid values: `off`, `on`.
     */
    public readonly htmlAccess!: pulumi.Output<string>;
    /**
     * The html5 file transfer. Valid values: `all`, `download`, `off`, `upload`.
     */
    public readonly htmlFileTransfer!: pulumi.Output<string>;
    /**
     * Local drive redirect policy. Valid values: ` readwrite`, `off`, `read`.
     */
    public readonly localDrive!: pulumi.Output<string>;
    /**
     * The name of policy group.
     */
    public readonly policyGroupName!: pulumi.Output<string | undefined>;
    /**
     * The status of policy.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The usb redirect policy. Valid values: `off`, `on`.
     */
    public readonly usbRedirect!: pulumi.Output<string>;
    /**
     * The quality of visual. Valid values: `high`, `lossless`, `low`, `medium`.
     */
    public readonly visualQuality!: pulumi.Output<string>;
    /**
     * The watermark policy. Valid values: `off`, `on`.
     */
    public readonly watermark!: pulumi.Output<string>;
    /**
     * The watermark transparency. Valid values: `DARK`, `LIGHT`, `MIDDLE`.
     */
    public readonly watermarkTransparency!: pulumi.Output<string>;
    /**
     * The type of watemark. Valid values: `EndUserId`, `HostName`.
     */
    public readonly watermarkType!: pulumi.Output<string>;

    /**
     * Create a EcdPolicyGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EcdPolicyGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EcdPolicyGroupArgs | EcdPolicyGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EcdPolicyGroupState | undefined;
            inputs["authorizeAccessPolicyRules"] = state ? state.authorizeAccessPolicyRules : undefined;
            inputs["authorizeSecurityPolicyRules"] = state ? state.authorizeSecurityPolicyRules : undefined;
            inputs["clipboard"] = state ? state.clipboard : undefined;
            inputs["domainList"] = state ? state.domainList : undefined;
            inputs["htmlAccess"] = state ? state.htmlAccess : undefined;
            inputs["htmlFileTransfer"] = state ? state.htmlFileTransfer : undefined;
            inputs["localDrive"] = state ? state.localDrive : undefined;
            inputs["policyGroupName"] = state ? state.policyGroupName : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["usbRedirect"] = state ? state.usbRedirect : undefined;
            inputs["visualQuality"] = state ? state.visualQuality : undefined;
            inputs["watermark"] = state ? state.watermark : undefined;
            inputs["watermarkTransparency"] = state ? state.watermarkTransparency : undefined;
            inputs["watermarkType"] = state ? state.watermarkType : undefined;
        } else {
            const args = argsOrState as EcdPolicyGroupArgs | undefined;
            inputs["authorizeAccessPolicyRules"] = args ? args.authorizeAccessPolicyRules : undefined;
            inputs["authorizeSecurityPolicyRules"] = args ? args.authorizeSecurityPolicyRules : undefined;
            inputs["clipboard"] = args ? args.clipboard : undefined;
            inputs["domainList"] = args ? args.domainList : undefined;
            inputs["htmlAccess"] = args ? args.htmlAccess : undefined;
            inputs["htmlFileTransfer"] = args ? args.htmlFileTransfer : undefined;
            inputs["localDrive"] = args ? args.localDrive : undefined;
            inputs["policyGroupName"] = args ? args.policyGroupName : undefined;
            inputs["usbRedirect"] = args ? args.usbRedirect : undefined;
            inputs["visualQuality"] = args ? args.visualQuality : undefined;
            inputs["watermark"] = args ? args.watermark : undefined;
            inputs["watermarkTransparency"] = args ? args.watermarkTransparency : undefined;
            inputs["watermarkType"] = args ? args.watermarkType : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EcdPolicyGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EcdPolicyGroup resources.
 */
export interface EcdPolicyGroupState {
    /**
     * The rule of authorize access rule.
     */
    readonly authorizeAccessPolicyRules?: pulumi.Input<pulumi.Input<inputs.eds.EcdPolicyGroupAuthorizeAccessPolicyRule>[]>;
    /**
     * The policy rule.
     */
    readonly authorizeSecurityPolicyRules?: pulumi.Input<pulumi.Input<inputs.eds.EcdPolicyGroupAuthorizeSecurityPolicyRule>[]>;
    /**
     * The clipboard policy. Valid values: `off`, `read`, `readwrite`.
     */
    readonly clipboard?: pulumi.Input<string>;
    /**
     * The list of domain.
     */
    readonly domainList?: pulumi.Input<string>;
    /**
     * The access of html5. Valid values: `off`, `on`.
     */
    readonly htmlAccess?: pulumi.Input<string>;
    /**
     * The html5 file transfer. Valid values: `all`, `download`, `off`, `upload`.
     */
    readonly htmlFileTransfer?: pulumi.Input<string>;
    /**
     * Local drive redirect policy. Valid values: ` readwrite`, `off`, `read`.
     */
    readonly localDrive?: pulumi.Input<string>;
    /**
     * The name of policy group.
     */
    readonly policyGroupName?: pulumi.Input<string>;
    /**
     * The status of policy.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The usb redirect policy. Valid values: `off`, `on`.
     */
    readonly usbRedirect?: pulumi.Input<string>;
    /**
     * The quality of visual. Valid values: `high`, `lossless`, `low`, `medium`.
     */
    readonly visualQuality?: pulumi.Input<string>;
    /**
     * The watermark policy. Valid values: `off`, `on`.
     */
    readonly watermark?: pulumi.Input<string>;
    /**
     * The watermark transparency. Valid values: `DARK`, `LIGHT`, `MIDDLE`.
     */
    readonly watermarkTransparency?: pulumi.Input<string>;
    /**
     * The type of watemark. Valid values: `EndUserId`, `HostName`.
     */
    readonly watermarkType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EcdPolicyGroup resource.
 */
export interface EcdPolicyGroupArgs {
    /**
     * The rule of authorize access rule.
     */
    readonly authorizeAccessPolicyRules?: pulumi.Input<pulumi.Input<inputs.eds.EcdPolicyGroupAuthorizeAccessPolicyRule>[]>;
    /**
     * The policy rule.
     */
    readonly authorizeSecurityPolicyRules?: pulumi.Input<pulumi.Input<inputs.eds.EcdPolicyGroupAuthorizeSecurityPolicyRule>[]>;
    /**
     * The clipboard policy. Valid values: `off`, `read`, `readwrite`.
     */
    readonly clipboard?: pulumi.Input<string>;
    /**
     * The list of domain.
     */
    readonly domainList?: pulumi.Input<string>;
    /**
     * The access of html5. Valid values: `off`, `on`.
     */
    readonly htmlAccess?: pulumi.Input<string>;
    /**
     * The html5 file transfer. Valid values: `all`, `download`, `off`, `upload`.
     */
    readonly htmlFileTransfer?: pulumi.Input<string>;
    /**
     * Local drive redirect policy. Valid values: ` readwrite`, `off`, `read`.
     */
    readonly localDrive?: pulumi.Input<string>;
    /**
     * The name of policy group.
     */
    readonly policyGroupName?: pulumi.Input<string>;
    /**
     * The usb redirect policy. Valid values: `off`, `on`.
     */
    readonly usbRedirect?: pulumi.Input<string>;
    /**
     * The quality of visual. Valid values: `high`, `lossless`, `low`, `medium`.
     */
    readonly visualQuality?: pulumi.Input<string>;
    /**
     * The watermark policy. Valid values: `off`, `on`.
     */
    readonly watermark?: pulumi.Input<string>;
    /**
     * The watermark transparency. Valid values: `DARK`, `LIGHT`, `MIDDLE`.
     */
    readonly watermarkTransparency?: pulumi.Input<string>;
    /**
     * The type of watemark. Valid values: `EndUserId`, `HostName`.
     */
    readonly watermarkType?: pulumi.Input<string>;
}
