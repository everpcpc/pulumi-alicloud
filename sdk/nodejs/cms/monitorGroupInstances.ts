// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Monitor Service Monitor Group Instances resource.
 *
 * For information about Cloud Monitor Service Monitor Group Instances and how to use it, see [What is Monitor Group Instances](https://www.alibabacloud.com/help/en/doc-detail/115031.htm).
 *
 * > **NOTE:** Available in v1.115.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {cidrBlock: "192.168.0.0/16"});
 * const defaultMonitorGroup = new alicloud.cms.MonitorGroup("defaultMonitorGroup", {monitorGroupName: "tf-testaccmonitorgroup"});
 * const example = new alicloud.cms.MonitorGroupInstances("example", {
 *     groupId: defaultMonitorGroup.id,
 *     instances: [{
 *         instanceId: defaultNetwork.id,
 *         instanceName: "tf-testacc-vpcname",
 *         regionId: "cn-hangzhou",
 *         category: "vpc",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Cloud Monitor Service Monitor Group Instances can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cms/monitorGroupInstances:MonitorGroupInstances example <group_id>
 * ```
 */
export class MonitorGroupInstances extends pulumi.CustomResource {
    /**
     * Get an existing MonitorGroupInstances resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MonitorGroupInstancesState, opts?: pulumi.CustomResourceOptions): MonitorGroupInstances {
        return new MonitorGroupInstances(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cms/monitorGroupInstances:MonitorGroupInstances';

    /**
     * Returns true if the given object is an instance of MonitorGroupInstances.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MonitorGroupInstances {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MonitorGroupInstances.__pulumiType;
    }

    /**
     * The id of Cms Group.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * Instance information added to the Cms Group.
     */
    public readonly instances!: pulumi.Output<outputs.cms.MonitorGroupInstancesInstance[]>;

    /**
     * Create a MonitorGroupInstances resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MonitorGroupInstancesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MonitorGroupInstancesArgs | MonitorGroupInstancesState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MonitorGroupInstancesState | undefined;
            inputs["groupId"] = state ? state.groupId : undefined;
            inputs["instances"] = state ? state.instances : undefined;
        } else {
            const args = argsOrState as MonitorGroupInstancesArgs | undefined;
            if ((!args || args.groupId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.instances === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'instances'");
            }
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["instances"] = args ? args.instances : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(MonitorGroupInstances.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MonitorGroupInstances resources.
 */
export interface MonitorGroupInstancesState {
    /**
     * The id of Cms Group.
     */
    readonly groupId?: pulumi.Input<string>;
    /**
     * Instance information added to the Cms Group.
     */
    readonly instances?: pulumi.Input<pulumi.Input<inputs.cms.MonitorGroupInstancesInstance>[]>;
}

/**
 * The set of arguments for constructing a MonitorGroupInstances resource.
 */
export interface MonitorGroupInstancesArgs {
    /**
     * The id of Cms Group.
     */
    readonly groupId: pulumi.Input<string>;
    /**
     * Instance information added to the Cms Group.
     */
    readonly instances: pulumi.Input<pulumi.Input<inputs.cms.MonitorGroupInstancesInstance>[]>;
}
