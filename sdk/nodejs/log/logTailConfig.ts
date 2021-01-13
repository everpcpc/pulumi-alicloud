// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Logtial config can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:log/logTailConfig:LogTailConfig example tf-log:tf-log-store:tf-log-config
 * ```
 */
export class LogTailConfig extends pulumi.CustomResource {
    /**
     * Get an existing LogTailConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogTailConfigState, opts?: pulumi.CustomResourceOptions): LogTailConfig {
        return new LogTailConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:log/logTailConfig:LogTailConfig';

    /**
     * Returns true if the given object is an instance of LogTailConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogTailConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogTailConfig.__pulumiType;
    }

    /**
     * The logtail configure the required JSON files. ([Refer to details](https://www.alibabacloud.com/help/doc-detail/29058.htm))
     */
    public readonly inputDetail!: pulumi.Output<string>;
    /**
     * The input type. Currently only two types of files and plugin are supported.
     */
    public readonly inputType!: pulumi.Output<string>;
    /**
     * （Optional）The log sample of the Logtail configuration. The log size cannot exceed 1,000 bytes.
     */
    public readonly logSample!: pulumi.Output<string | undefined>;
    /**
     * The log store name to the query index belongs.
     */
    public readonly logstore!: pulumi.Output<string>;
    /**
     * The Logtail configuration name, which is unique in the same project.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The output type. Currently, only LogService is supported.
     */
    public readonly outputType!: pulumi.Output<string>;
    /**
     * The project name to the log store belongs.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a LogTailConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogTailConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogTailConfigArgs | LogTailConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LogTailConfigState | undefined;
            inputs["inputDetail"] = state ? state.inputDetail : undefined;
            inputs["inputType"] = state ? state.inputType : undefined;
            inputs["logSample"] = state ? state.logSample : undefined;
            inputs["logstore"] = state ? state.logstore : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["outputType"] = state ? state.outputType : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as LogTailConfigArgs | undefined;
            if ((!args || args.inputDetail === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'inputDetail'");
            }
            if ((!args || args.inputType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'inputType'");
            }
            if ((!args || args.logstore === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'logstore'");
            }
            if ((!args || args.outputType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'outputType'");
            }
            if ((!args || args.project === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'project'");
            }
            inputs["inputDetail"] = args ? args.inputDetail : undefined;
            inputs["inputType"] = args ? args.inputType : undefined;
            inputs["logSample"] = args ? args.logSample : undefined;
            inputs["logstore"] = args ? args.logstore : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["outputType"] = args ? args.outputType : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LogTailConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogTailConfig resources.
 */
export interface LogTailConfigState {
    /**
     * The logtail configure the required JSON files. ([Refer to details](https://www.alibabacloud.com/help/doc-detail/29058.htm))
     */
    readonly inputDetail?: pulumi.Input<string>;
    /**
     * The input type. Currently only two types of files and plugin are supported.
     */
    readonly inputType?: pulumi.Input<string>;
    /**
     * （Optional）The log sample of the Logtail configuration. The log size cannot exceed 1,000 bytes.
     */
    readonly logSample?: pulumi.Input<string>;
    /**
     * The log store name to the query index belongs.
     */
    readonly logstore?: pulumi.Input<string>;
    /**
     * The Logtail configuration name, which is unique in the same project.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The output type. Currently, only LogService is supported.
     */
    readonly outputType?: pulumi.Input<string>;
    /**
     * The project name to the log store belongs.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogTailConfig resource.
 */
export interface LogTailConfigArgs {
    /**
     * The logtail configure the required JSON files. ([Refer to details](https://www.alibabacloud.com/help/doc-detail/29058.htm))
     */
    readonly inputDetail: pulumi.Input<string>;
    /**
     * The input type. Currently only two types of files and plugin are supported.
     */
    readonly inputType: pulumi.Input<string>;
    /**
     * （Optional）The log sample of the Logtail configuration. The log size cannot exceed 1,000 bytes.
     */
    readonly logSample?: pulumi.Input<string>;
    /**
     * The log store name to the query index belongs.
     */
    readonly logstore: pulumi.Input<string>;
    /**
     * The Logtail configuration name, which is unique in the same project.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The output type. Currently, only LogService is supported.
     */
    readonly outputType: pulumi.Input<string>;
    /**
     * The project name to the log store belongs.
     */
    readonly project: pulumi.Input<string>;
}
