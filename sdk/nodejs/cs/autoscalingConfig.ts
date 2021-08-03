// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class AutoscalingConfig extends pulumi.CustomResource {
    /**
     * Get an existing AutoscalingConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutoscalingConfigState, opts?: pulumi.CustomResourceOptions): AutoscalingConfig {
        return new AutoscalingConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cs/autoscalingConfig:AutoscalingConfig';

    /**
     * Returns true if the given object is an instance of AutoscalingConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutoscalingConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutoscalingConfig.__pulumiType;
    }

    /**
     * The id of kubernetes cluster.
     */
    public readonly clusterId!: pulumi.Output<string | undefined>;
    /**
     * The cool down duration. Default is `10m`. If the delay (cooldown) value is set too long, there could be complaints that the Horizontal Pod Autoscaler is not responsive to workload changes. However, if the delay value is set too short, the scale of the replicas set may keep thrashing as usual.
     */
    public readonly coolDownDuration!: pulumi.Output<string | undefined>;
    /**
     * The scale-in threshold for GPU instance. Default is `0.5`.
     */
    public readonly gpuUtilizationThreshold!: pulumi.Output<string | undefined>;
    /**
     * The scan interval. Default is `30s`
     */
    public readonly scanInterval!: pulumi.Output<string | undefined>;
    /**
     * The unneeded duration. Default is `10m`.
     */
    public readonly unneededDuration!: pulumi.Output<string | undefined>;
    /**
     * The scale-in threshold. Default is `0.5`.
     */
    public readonly utilizationThreshold!: pulumi.Output<string | undefined>;

    /**
     * Create a AutoscalingConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AutoscalingConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutoscalingConfigArgs | AutoscalingConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutoscalingConfigState | undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["coolDownDuration"] = state ? state.coolDownDuration : undefined;
            inputs["gpuUtilizationThreshold"] = state ? state.gpuUtilizationThreshold : undefined;
            inputs["scanInterval"] = state ? state.scanInterval : undefined;
            inputs["unneededDuration"] = state ? state.unneededDuration : undefined;
            inputs["utilizationThreshold"] = state ? state.utilizationThreshold : undefined;
        } else {
            const args = argsOrState as AutoscalingConfigArgs | undefined;
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["coolDownDuration"] = args ? args.coolDownDuration : undefined;
            inputs["gpuUtilizationThreshold"] = args ? args.gpuUtilizationThreshold : undefined;
            inputs["scanInterval"] = args ? args.scanInterval : undefined;
            inputs["unneededDuration"] = args ? args.unneededDuration : undefined;
            inputs["utilizationThreshold"] = args ? args.utilizationThreshold : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AutoscalingConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AutoscalingConfig resources.
 */
export interface AutoscalingConfigState {
    /**
     * The id of kubernetes cluster.
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * The cool down duration. Default is `10m`. If the delay (cooldown) value is set too long, there could be complaints that the Horizontal Pod Autoscaler is not responsive to workload changes. However, if the delay value is set too short, the scale of the replicas set may keep thrashing as usual.
     */
    readonly coolDownDuration?: pulumi.Input<string>;
    /**
     * The scale-in threshold for GPU instance. Default is `0.5`.
     */
    readonly gpuUtilizationThreshold?: pulumi.Input<string>;
    /**
     * The scan interval. Default is `30s`
     */
    readonly scanInterval?: pulumi.Input<string>;
    /**
     * The unneeded duration. Default is `10m`.
     */
    readonly unneededDuration?: pulumi.Input<string>;
    /**
     * The scale-in threshold. Default is `0.5`.
     */
    readonly utilizationThreshold?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AutoscalingConfig resource.
 */
export interface AutoscalingConfigArgs {
    /**
     * The id of kubernetes cluster.
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * The cool down duration. Default is `10m`. If the delay (cooldown) value is set too long, there could be complaints that the Horizontal Pod Autoscaler is not responsive to workload changes. However, if the delay value is set too short, the scale of the replicas set may keep thrashing as usual.
     */
    readonly coolDownDuration?: pulumi.Input<string>;
    /**
     * The scale-in threshold for GPU instance. Default is `0.5`.
     */
    readonly gpuUtilizationThreshold?: pulumi.Input<string>;
    /**
     * The scan interval. Default is `30s`
     */
    readonly scanInterval?: pulumi.Input<string>;
    /**
     * The unneeded duration. Default is `10m`.
     */
    readonly unneededDuration?: pulumi.Input<string>;
    /**
     * The scale-in threshold. Default is `0.5`.
     */
    readonly utilizationThreshold?: pulumi.Input<string>;
}