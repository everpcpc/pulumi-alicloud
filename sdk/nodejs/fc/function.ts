// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Function Compute function can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:fc/function:Function foo my-fc-service:hello-world
 * ```
 */
export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionState, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:fc/function:Function';

    /**
     * Returns true if the given object is an instance of Function.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Function {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Function.__pulumiType;
    }

    /**
     * The port that the function listen to, only valid for [custom runtime](https://www.alibabacloud.com/help/doc-detail/132044.htm) and [custom container runtime](https://www.alibabacloud.com/help/doc-detail/179368.htm).
     */
    public readonly caPort!: pulumi.Output<number | undefined>;
    /**
     * The checksum (crc64) of the function code.
     */
    public readonly codeChecksum!: pulumi.Output<string>;
    /**
     * The configuration for custom container runtime.
     */
    public readonly customContainerConfig!: pulumi.Output<outputs.fc.FunctionCustomContainerConfig | undefined>;
    /**
     * The Function Compute function description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A map that defines environment variables for the function.
     */
    public readonly environmentVariables!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The path to the function's deployment package within the local filesystem. It is conflict with the `oss_`-prefixed options.
     */
    public readonly filename!: pulumi.Output<string | undefined>;
    /**
     * The Function Compute service ID.
     */
    public /*out*/ readonly functionId!: pulumi.Output<string>;
    /**
     * The function [entry point](https://www.alibabacloud.com/help/doc-detail/157704.htm) in your code.
     */
    public readonly handler!: pulumi.Output<string>;
    /**
     * The maximum length of time, in seconds, that the function's initialization should be run for.
     */
    public readonly initializationTimeout!: pulumi.Output<number | undefined>;
    /**
     * The entry point of the function's [initialization](https://www.alibabacloud.com/help/doc-detail/157704.htm).
     */
    public readonly initializer!: pulumi.Output<string | undefined>;
    /**
     * The maximum number of requests can be executed concurrently within the single function instance.
     */
    public readonly instanceConcurrency!: pulumi.Output<number | undefined>;
    /**
     * The instance type of the function.
     */
    public readonly instanceType!: pulumi.Output<string | undefined>;
    /**
     * The date this resource was last modified.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * Amount of memory in MB your function can use at runtime. Defaults to `128`. Limits to [128, 3072].
     */
    public readonly memorySize!: pulumi.Output<number | undefined>;
    /**
     * The Function Compute function name. It is the only in one service and is conflict with "namePrefix".
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Setting a prefix to get a only function name. It is conflict with "name".
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * The OSS bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same Alibaba Cloud region where you are creating the function.
     */
    public readonly ossBucket!: pulumi.Output<string | undefined>;
    /**
     * The OSS key of an object containing the function's deployment package. Conflicts with `filename`.
     */
    public readonly ossKey!: pulumi.Output<string | undefined>;
    /**
     * See [Runtimes][https://www.alibabacloud.com/help/doc-detail/52077.htm] for valid values.
     */
    public readonly runtime!: pulumi.Output<string>;
    /**
     * The Function Compute service name.
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * The amount of time your function has to run in seconds.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionArgs | FunctionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionState | undefined;
            inputs["caPort"] = state ? state.caPort : undefined;
            inputs["codeChecksum"] = state ? state.codeChecksum : undefined;
            inputs["customContainerConfig"] = state ? state.customContainerConfig : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["environmentVariables"] = state ? state.environmentVariables : undefined;
            inputs["filename"] = state ? state.filename : undefined;
            inputs["functionId"] = state ? state.functionId : undefined;
            inputs["handler"] = state ? state.handler : undefined;
            inputs["initializationTimeout"] = state ? state.initializationTimeout : undefined;
            inputs["initializer"] = state ? state.initializer : undefined;
            inputs["instanceConcurrency"] = state ? state.instanceConcurrency : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["lastModified"] = state ? state.lastModified : undefined;
            inputs["memorySize"] = state ? state.memorySize : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["ossBucket"] = state ? state.ossBucket : undefined;
            inputs["ossKey"] = state ? state.ossKey : undefined;
            inputs["runtime"] = state ? state.runtime : undefined;
            inputs["service"] = state ? state.service : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as FunctionArgs | undefined;
            if ((!args || args.handler === undefined) && !opts.urn) {
                throw new Error("Missing required property 'handler'");
            }
            if ((!args || args.runtime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runtime'");
            }
            if ((!args || args.service === undefined) && !opts.urn) {
                throw new Error("Missing required property 'service'");
            }
            inputs["caPort"] = args ? args.caPort : undefined;
            inputs["codeChecksum"] = args ? args.codeChecksum : undefined;
            inputs["customContainerConfig"] = args ? args.customContainerConfig : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            inputs["filename"] = args ? args.filename : undefined;
            inputs["handler"] = args ? args.handler : undefined;
            inputs["initializationTimeout"] = args ? args.initializationTimeout : undefined;
            inputs["initializer"] = args ? args.initializer : undefined;
            inputs["instanceConcurrency"] = args ? args.instanceConcurrency : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["memorySize"] = args ? args.memorySize : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["ossBucket"] = args ? args.ossBucket : undefined;
            inputs["ossKey"] = args ? args.ossKey : undefined;
            inputs["runtime"] = args ? args.runtime : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
            inputs["functionId"] = undefined /*out*/;
            inputs["lastModified"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Function.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Function resources.
 */
export interface FunctionState {
    /**
     * The port that the function listen to, only valid for [custom runtime](https://www.alibabacloud.com/help/doc-detail/132044.htm) and [custom container runtime](https://www.alibabacloud.com/help/doc-detail/179368.htm).
     */
    readonly caPort?: pulumi.Input<number>;
    /**
     * The checksum (crc64) of the function code.
     */
    readonly codeChecksum?: pulumi.Input<string>;
    /**
     * The configuration for custom container runtime.
     */
    readonly customContainerConfig?: pulumi.Input<inputs.fc.FunctionCustomContainerConfig>;
    /**
     * The Function Compute function description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A map that defines environment variables for the function.
     */
    readonly environmentVariables?: pulumi.Input<{[key: string]: any}>;
    /**
     * The path to the function's deployment package within the local filesystem. It is conflict with the `oss_`-prefixed options.
     */
    readonly filename?: pulumi.Input<string>;
    /**
     * The Function Compute service ID.
     */
    readonly functionId?: pulumi.Input<string>;
    /**
     * The function [entry point](https://www.alibabacloud.com/help/doc-detail/157704.htm) in your code.
     */
    readonly handler?: pulumi.Input<string>;
    /**
     * The maximum length of time, in seconds, that the function's initialization should be run for.
     */
    readonly initializationTimeout?: pulumi.Input<number>;
    /**
     * The entry point of the function's [initialization](https://www.alibabacloud.com/help/doc-detail/157704.htm).
     */
    readonly initializer?: pulumi.Input<string>;
    /**
     * The maximum number of requests can be executed concurrently within the single function instance.
     */
    readonly instanceConcurrency?: pulumi.Input<number>;
    /**
     * The instance type of the function.
     */
    readonly instanceType?: pulumi.Input<string>;
    /**
     * The date this resource was last modified.
     */
    readonly lastModified?: pulumi.Input<string>;
    /**
     * Amount of memory in MB your function can use at runtime. Defaults to `128`. Limits to [128, 3072].
     */
    readonly memorySize?: pulumi.Input<number>;
    /**
     * The Function Compute function name. It is the only in one service and is conflict with "namePrefix".
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Setting a prefix to get a only function name. It is conflict with "name".
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * The OSS bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same Alibaba Cloud region where you are creating the function.
     */
    readonly ossBucket?: pulumi.Input<string>;
    /**
     * The OSS key of an object containing the function's deployment package. Conflicts with `filename`.
     */
    readonly ossKey?: pulumi.Input<string>;
    /**
     * See [Runtimes][https://www.alibabacloud.com/help/doc-detail/52077.htm] for valid values.
     */
    readonly runtime?: pulumi.Input<string>;
    /**
     * The Function Compute service name.
     */
    readonly service?: pulumi.Input<string>;
    /**
     * The amount of time your function has to run in seconds.
     */
    readonly timeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    /**
     * The port that the function listen to, only valid for [custom runtime](https://www.alibabacloud.com/help/doc-detail/132044.htm) and [custom container runtime](https://www.alibabacloud.com/help/doc-detail/179368.htm).
     */
    readonly caPort?: pulumi.Input<number>;
    /**
     * The checksum (crc64) of the function code.
     */
    readonly codeChecksum?: pulumi.Input<string>;
    /**
     * The configuration for custom container runtime.
     */
    readonly customContainerConfig?: pulumi.Input<inputs.fc.FunctionCustomContainerConfig>;
    /**
     * The Function Compute function description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A map that defines environment variables for the function.
     */
    readonly environmentVariables?: pulumi.Input<{[key: string]: any}>;
    /**
     * The path to the function's deployment package within the local filesystem. It is conflict with the `oss_`-prefixed options.
     */
    readonly filename?: pulumi.Input<string>;
    /**
     * The function [entry point](https://www.alibabacloud.com/help/doc-detail/157704.htm) in your code.
     */
    readonly handler: pulumi.Input<string>;
    /**
     * The maximum length of time, in seconds, that the function's initialization should be run for.
     */
    readonly initializationTimeout?: pulumi.Input<number>;
    /**
     * The entry point of the function's [initialization](https://www.alibabacloud.com/help/doc-detail/157704.htm).
     */
    readonly initializer?: pulumi.Input<string>;
    /**
     * The maximum number of requests can be executed concurrently within the single function instance.
     */
    readonly instanceConcurrency?: pulumi.Input<number>;
    /**
     * The instance type of the function.
     */
    readonly instanceType?: pulumi.Input<string>;
    /**
     * Amount of memory in MB your function can use at runtime. Defaults to `128`. Limits to [128, 3072].
     */
    readonly memorySize?: pulumi.Input<number>;
    /**
     * The Function Compute function name. It is the only in one service and is conflict with "namePrefix".
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Setting a prefix to get a only function name. It is conflict with "name".
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * The OSS bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same Alibaba Cloud region where you are creating the function.
     */
    readonly ossBucket?: pulumi.Input<string>;
    /**
     * The OSS key of an object containing the function's deployment package. Conflicts with `filename`.
     */
    readonly ossKey?: pulumi.Input<string>;
    /**
     * See [Runtimes][https://www.alibabacloud.com/help/doc-detail/52077.htm] for valid values.
     */
    readonly runtime: pulumi.Input<string>;
    /**
     * The Function Compute service name.
     */
    readonly service: pulumi.Input<string>;
    /**
     * The amount of time your function has to run in seconds.
     */
    readonly timeout?: pulumi.Input<number>;
}
