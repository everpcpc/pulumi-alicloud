// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ECS Deployment Set resource.
 *
 * For information about ECS Deployment Set and how to use it, see [What is Deployment Set](https://www.alibabacloud.com/help/en/doc-detail/91269.htm).
 *
 * > **NOTE:** Available in v1.140.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultEcsDeploymentSet = new alicloud.ecs.EcsDeploymentSet("default", {
 *     deploymentSetName: "example_value",
 *     description: "example_value",
 *     domain: "Default",
 *     granularity: "Host",
 *     strategy: "Availability",
 * });
 * ```
 *
 * ## Import
 *
 * ECS Deployment Set can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ecs/ecsDeploymentSet:EcsDeploymentSet example <id>
 * ```
 */
export class EcsDeploymentSet extends pulumi.CustomResource {
    /**
     * Get an existing EcsDeploymentSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EcsDeploymentSetState, opts?: pulumi.CustomResourceOptions): EcsDeploymentSet {
        return new EcsDeploymentSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/ecsDeploymentSet:EcsDeploymentSet';

    /**
     * Returns true if the given object is an instance of EcsDeploymentSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EcsDeploymentSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EcsDeploymentSet.__pulumiType;
    }

    /**
     * The name of the deployment set. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
     */
    public readonly deploymentSetName!: pulumi.Output<string | undefined>;
    /**
     * The description of the deployment set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The deployment domain. Valid values: `Default`.
     */
    public readonly domain!: pulumi.Output<string | undefined>;
    /**
     * The deployment granularity. Valid values: `Host`.
     */
    public readonly granularity!: pulumi.Output<string | undefined>;
    /**
     * The on unable to redeploy failed instance. Valid values: `CancelMembershipAndStart`, `KeepStopped`.
     */
    public readonly onUnableToRedeployFailedInstance!: pulumi.Output<string | undefined>;
    /**
     * The deployment strategy. Valid values: `Availability`.
     */
    public readonly strategy!: pulumi.Output<string | undefined>;

    /**
     * Create a EcsDeploymentSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EcsDeploymentSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EcsDeploymentSetArgs | EcsDeploymentSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EcsDeploymentSetState | undefined;
            inputs["deploymentSetName"] = state ? state.deploymentSetName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["granularity"] = state ? state.granularity : undefined;
            inputs["onUnableToRedeployFailedInstance"] = state ? state.onUnableToRedeployFailedInstance : undefined;
            inputs["strategy"] = state ? state.strategy : undefined;
        } else {
            const args = argsOrState as EcsDeploymentSetArgs | undefined;
            inputs["deploymentSetName"] = args ? args.deploymentSetName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["granularity"] = args ? args.granularity : undefined;
            inputs["onUnableToRedeployFailedInstance"] = args ? args.onUnableToRedeployFailedInstance : undefined;
            inputs["strategy"] = args ? args.strategy : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EcsDeploymentSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EcsDeploymentSet resources.
 */
export interface EcsDeploymentSetState {
    /**
     * The name of the deployment set. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
     */
    deploymentSetName?: pulumi.Input<string>;
    /**
     * The description of the deployment set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * The deployment domain. Valid values: `Default`.
     */
    domain?: pulumi.Input<string>;
    /**
     * The deployment granularity. Valid values: `Host`.
     */
    granularity?: pulumi.Input<string>;
    /**
     * The on unable to redeploy failed instance. Valid values: `CancelMembershipAndStart`, `KeepStopped`.
     */
    onUnableToRedeployFailedInstance?: pulumi.Input<string>;
    /**
     * The deployment strategy. Valid values: `Availability`.
     */
    strategy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EcsDeploymentSet resource.
 */
export interface EcsDeploymentSetArgs {
    /**
     * The name of the deployment set. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
     */
    deploymentSetName?: pulumi.Input<string>;
    /**
     * The description of the deployment set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * The deployment domain. Valid values: `Default`.
     */
    domain?: pulumi.Input<string>;
    /**
     * The deployment granularity. Valid values: `Host`.
     */
    granularity?: pulumi.Input<string>;
    /**
     * The on unable to redeploy failed instance. Valid values: `CancelMembershipAndStart`, `KeepStopped`.
     */
    onUnableToRedeployFailedInstance?: pulumi.Input<string>;
    /**
     * The deployment strategy. Valid values: `Availability`.
     */
    strategy?: pulumi.Input<string>;
}
