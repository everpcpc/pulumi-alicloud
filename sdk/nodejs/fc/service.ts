// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Function Compute Service can be imported using the id or name, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:fc/service:Service foo my-fc-service
 * ```
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:fc/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * The Function Compute Service description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether to allow the Service to access Internet. Default to "true".
     */
    public readonly internetAccess!: pulumi.Output<boolean | undefined>;
    /**
     * The date this resource was last modified.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * Provide this to store your Function Compute Service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm).
     */
    public readonly logConfig!: pulumi.Output<outputs.fc.ServiceLogConfig | undefined>;
    /**
     * The Function Compute Service name. It is the only in one Alicloud account and is conflict with `namePrefix`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Setting a prefix to get a only name. It is conflict with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * Provide [NAS configuration](https://www.alibabacloud.com/help/doc-detail/87401.htm) to allow Function Compute Service to access your NAS resources.
     */
    public readonly nasConfig!: pulumi.Output<outputs.fc.ServiceNasConfig | undefined>;
    /**
     * Whether to publish creation/change as new Function Compute Service Version. Defaults to `false`.
     */
    public readonly publish!: pulumi.Output<boolean | undefined>;
    /**
     * RAM role arn attached to the Function Compute Service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
     */
    public readonly role!: pulumi.Output<string | undefined>;
    /**
     * The Function Compute Service ID.
     */
    public /*out*/ readonly serviceId!: pulumi.Output<string>;
    /**
     * The latest published version of your Function Compute Service.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;
    /**
     * Provide this to allow your Function Compute Service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm).
     */
    public readonly vpcConfig!: pulumi.Output<outputs.fc.ServiceVpcConfig | undefined>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["internetAccess"] = state ? state.internetAccess : undefined;
            inputs["lastModified"] = state ? state.lastModified : undefined;
            inputs["logConfig"] = state ? state.logConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["nasConfig"] = state ? state.nasConfig : undefined;
            inputs["publish"] = state ? state.publish : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["serviceId"] = state ? state.serviceId : undefined;
            inputs["version"] = state ? state.version : undefined;
            inputs["vpcConfig"] = state ? state.vpcConfig : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["internetAccess"] = args ? args.internetAccess : undefined;
            inputs["logConfig"] = args ? args.logConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["nasConfig"] = args ? args.nasConfig : undefined;
            inputs["publish"] = args ? args.publish : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["vpcConfig"] = args ? args.vpcConfig : undefined;
            inputs["lastModified"] = undefined /*out*/;
            inputs["serviceId"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Service.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * The Function Compute Service description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether to allow the Service to access Internet. Default to "true".
     */
    readonly internetAccess?: pulumi.Input<boolean>;
    /**
     * The date this resource was last modified.
     */
    readonly lastModified?: pulumi.Input<string>;
    /**
     * Provide this to store your Function Compute Service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm).
     */
    readonly logConfig?: pulumi.Input<inputs.fc.ServiceLogConfig>;
    /**
     * The Function Compute Service name. It is the only in one Alicloud account and is conflict with `namePrefix`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Setting a prefix to get a only name. It is conflict with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * Provide [NAS configuration](https://www.alibabacloud.com/help/doc-detail/87401.htm) to allow Function Compute Service to access your NAS resources.
     */
    readonly nasConfig?: pulumi.Input<inputs.fc.ServiceNasConfig>;
    /**
     * Whether to publish creation/change as new Function Compute Service Version. Defaults to `false`.
     */
    readonly publish?: pulumi.Input<boolean>;
    /**
     * RAM role arn attached to the Function Compute Service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * The Function Compute Service ID.
     */
    readonly serviceId?: pulumi.Input<string>;
    /**
     * The latest published version of your Function Compute Service.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * Provide this to allow your Function Compute Service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm).
     */
    readonly vpcConfig?: pulumi.Input<inputs.fc.ServiceVpcConfig>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * The Function Compute Service description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether to allow the Service to access Internet. Default to "true".
     */
    readonly internetAccess?: pulumi.Input<boolean>;
    /**
     * Provide this to store your Function Compute Service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm).
     */
    readonly logConfig?: pulumi.Input<inputs.fc.ServiceLogConfig>;
    /**
     * The Function Compute Service name. It is the only in one Alicloud account and is conflict with `namePrefix`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Setting a prefix to get a only name. It is conflict with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * Provide [NAS configuration](https://www.alibabacloud.com/help/doc-detail/87401.htm) to allow Function Compute Service to access your NAS resources.
     */
    readonly nasConfig?: pulumi.Input<inputs.fc.ServiceNasConfig>;
    /**
     * Whether to publish creation/change as new Function Compute Service Version. Defaults to `false`.
     */
    readonly publish?: pulumi.Input<boolean>;
    /**
     * RAM role arn attached to the Function Compute Service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * Provide this to allow your Function Compute Service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm).
     */
    readonly vpcConfig?: pulumi.Input<inputs.fc.ServiceVpcConfig>;
}
