// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class KubernetesAutoscaler extends pulumi.CustomResource {
    /**
     * Get an existing KubernetesAutoscaler resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KubernetesAutoscalerState, opts?: pulumi.CustomResourceOptions): KubernetesAutoscaler {
        return new KubernetesAutoscaler(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cs/kubernetesAutoscaler:KubernetesAutoscaler';

    /**
     * Returns true if the given object is an instance of KubernetesAutoscaler.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KubernetesAutoscaler {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KubernetesAutoscaler.__pulumiType;
    }

    /**
     * The id of kubernetes cluster.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The coolDownDuration option of cluster-autoscaler.
     */
    public readonly coolDownDuration!: pulumi.Output<string>;
    /**
     * The deferScaleInDuration option of cluster-autoscaler.
     */
    public readonly deferScaleInDuration!: pulumi.Output<string>;
    /**
     * * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
     * * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
     * * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
     */
    public readonly nodepools!: pulumi.Output<outputs.cs.KubernetesAutoscalerNodepool[] | undefined>;
    /**
     * Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
     */
    public readonly useEcsRamRoleToken!: pulumi.Output<boolean | undefined>;
    /**
     * The utilization option of cluster-autoscaler.
     */
    public readonly utilization!: pulumi.Output<string>;

    /**
     * Create a KubernetesAutoscaler resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KubernetesAutoscalerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KubernetesAutoscalerArgs | KubernetesAutoscalerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KubernetesAutoscalerState | undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["coolDownDuration"] = state ? state.coolDownDuration : undefined;
            inputs["deferScaleInDuration"] = state ? state.deferScaleInDuration : undefined;
            inputs["nodepools"] = state ? state.nodepools : undefined;
            inputs["useEcsRamRoleToken"] = state ? state.useEcsRamRoleToken : undefined;
            inputs["utilization"] = state ? state.utilization : undefined;
        } else {
            const args = argsOrState as KubernetesAutoscalerArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.coolDownDuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'coolDownDuration'");
            }
            if ((!args || args.deferScaleInDuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deferScaleInDuration'");
            }
            if ((!args || args.utilization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'utilization'");
            }
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["coolDownDuration"] = args ? args.coolDownDuration : undefined;
            inputs["deferScaleInDuration"] = args ? args.deferScaleInDuration : undefined;
            inputs["nodepools"] = args ? args.nodepools : undefined;
            inputs["useEcsRamRoleToken"] = args ? args.useEcsRamRoleToken : undefined;
            inputs["utilization"] = args ? args.utilization : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(KubernetesAutoscaler.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KubernetesAutoscaler resources.
 */
export interface KubernetesAutoscalerState {
    /**
     * The id of kubernetes cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The coolDownDuration option of cluster-autoscaler.
     */
    coolDownDuration?: pulumi.Input<string>;
    /**
     * The deferScaleInDuration option of cluster-autoscaler.
     */
    deferScaleInDuration?: pulumi.Input<string>;
    /**
     * * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
     * * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
     * * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
     */
    nodepools?: pulumi.Input<pulumi.Input<inputs.cs.KubernetesAutoscalerNodepool>[]>;
    /**
     * Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
     */
    useEcsRamRoleToken?: pulumi.Input<boolean>;
    /**
     * The utilization option of cluster-autoscaler.
     */
    utilization?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KubernetesAutoscaler resource.
 */
export interface KubernetesAutoscalerArgs {
    /**
     * The id of kubernetes cluster.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The coolDownDuration option of cluster-autoscaler.
     */
    coolDownDuration: pulumi.Input<string>;
    /**
     * The deferScaleInDuration option of cluster-autoscaler.
     */
    deferScaleInDuration: pulumi.Input<string>;
    /**
     * * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
     * * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
     * * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
     */
    nodepools?: pulumi.Input<pulumi.Input<inputs.cs.KubernetesAutoscalerNodepool>[]>;
    /**
     * Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
     */
    useEcsRamRoleToken?: pulumi.Input<boolean>;
    /**
     * The utilization option of cluster-autoscaler.
     */
    utilization: pulumi.Input<string>;
}
