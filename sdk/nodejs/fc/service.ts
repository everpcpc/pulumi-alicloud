// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
     * The function compute service description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether to allow the service to access Internet. Default to "true".
     */
    public readonly internetAccess!: pulumi.Output<boolean | undefined>;
    /**
     * The date this resource was last modified.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * Provide this to store your FC service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm).
     */
    public readonly logConfig!: pulumi.Output<outputs.fc.ServiceLogConfig | undefined>;
    /**
     * The Function Compute service name. It is the only in one Alicloud account and is conflict with "namePrefix".
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Setting a prefix to get a only name. It is conflict with "name".
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * RAM role arn attached to the Function Compute service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
     */
    public readonly role!: pulumi.Output<string | undefined>;
    /**
     * The Function Compute service ID.
     */
    public /*out*/ readonly serviceId!: pulumi.Output<string>;
    /**
     * Provide this to allow your FC service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm).
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
        if (opts && opts.id) {
            const state = argsOrState as ServiceState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["internetAccess"] = state ? state.internetAccess : undefined;
            inputs["lastModified"] = state ? state.lastModified : undefined;
            inputs["logConfig"] = state ? state.logConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["serviceId"] = state ? state.serviceId : undefined;
            inputs["vpcConfig"] = state ? state.vpcConfig : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["internetAccess"] = args ? args.internetAccess : undefined;
            inputs["logConfig"] = args ? args.logConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["vpcConfig"] = args ? args.vpcConfig : undefined;
            inputs["lastModified"] = undefined /*out*/;
            inputs["serviceId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Service.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * The function compute service description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether to allow the service to access Internet. Default to "true".
     */
    readonly internetAccess?: pulumi.Input<boolean>;
    /**
     * The date this resource was last modified.
     */
    readonly lastModified?: pulumi.Input<string>;
    /**
     * Provide this to store your FC service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm).
     */
    readonly logConfig?: pulumi.Input<inputs.fc.ServiceLogConfig>;
    /**
     * The Function Compute service name. It is the only in one Alicloud account and is conflict with "namePrefix".
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Setting a prefix to get a only name. It is conflict with "name".
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * RAM role arn attached to the Function Compute service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * The Function Compute service ID.
     */
    readonly serviceId?: pulumi.Input<string>;
    /**
     * Provide this to allow your FC service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm).
     */
    readonly vpcConfig?: pulumi.Input<inputs.fc.ServiceVpcConfig>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * The function compute service description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether to allow the service to access Internet. Default to "true".
     */
    readonly internetAccess?: pulumi.Input<boolean>;
    /**
     * Provide this to store your FC service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm).
     */
    readonly logConfig?: pulumi.Input<inputs.fc.ServiceLogConfig>;
    /**
     * The Function Compute service name. It is the only in one Alicloud account and is conflict with "namePrefix".
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Setting a prefix to get a only name. It is conflict with "name".
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * RAM role arn attached to the Function Compute service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * Provide this to allow your FC service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm).
     */
    readonly vpcConfig?: pulumi.Input<inputs.fc.ServiceVpcConfig>;
}
