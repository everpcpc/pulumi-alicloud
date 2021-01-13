// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ROS Template resource.
 *
 * For information about ROS Template and how to use it, see [What is Template](https://www.alibabacloud.com/help/en/doc-detail/141851.htm).
 *
 * > **NOTE:** Available in v1.108.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.ros.Template("example", {
 *     templateBody: `    {
 *     	"ROSTemplateFormatVersion": "2015-09-01"
 *     }
 *     `,
 *     templateName: "example_value",
 * });
 * ```
 *
 * ## Import
 *
 * ROS Template can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ros/template:Template example <template_id>
 * ```
 */
export class Template extends pulumi.CustomResource {
    /**
     * Get an existing Template resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TemplateState, opts?: pulumi.CustomResourceOptions): Template {
        return new Template(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ros/template:Template';

    /**
     * Returns true if the given object is an instance of Template.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Template {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Template.__pulumiType;
    }

    /**
     * The description of the template. The description can be up to 256 characters in length.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You must specify one of the TemplateBody and TemplateURL parameters, but you cannot specify both of them.
     */
    public readonly templateBody!: pulumi.Output<string | undefined>;
    /**
     * The name of the template. The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
     */
    public readonly templateName!: pulumi.Output<string>;
    /**
     * The template url.
     */
    public readonly templateUrl!: pulumi.Output<string | undefined>;

    /**
     * Create a Template resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TemplateArgs | TemplateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TemplateState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["templateBody"] = state ? state.templateBody : undefined;
            inputs["templateName"] = state ? state.templateName : undefined;
            inputs["templateUrl"] = state ? state.templateUrl : undefined;
        } else {
            const args = argsOrState as TemplateArgs | undefined;
            if ((!args || args.templateName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'templateName'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["templateBody"] = args ? args.templateBody : undefined;
            inputs["templateName"] = args ? args.templateName : undefined;
            inputs["templateUrl"] = args ? args.templateUrl : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Template.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Template resources.
 */
export interface TemplateState {
    /**
     * The description of the template. The description can be up to 256 characters in length.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You must specify one of the TemplateBody and TemplateURL parameters, but you cannot specify both of them.
     */
    readonly templateBody?: pulumi.Input<string>;
    /**
     * The name of the template. The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
     */
    readonly templateName?: pulumi.Input<string>;
    /**
     * The template url.
     */
    readonly templateUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Template resource.
 */
export interface TemplateArgs {
    /**
     * The description of the template. The description can be up to 256 characters in length.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You must specify one of the TemplateBody and TemplateURL parameters, but you cannot specify both of them.
     */
    readonly templateBody?: pulumi.Input<string>;
    /**
     * The name of the template. The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
     */
    readonly templateName: pulumi.Input<string>;
    /**
     * The template url.
     */
    readonly templateUrl?: pulumi.Input<string>;
}
