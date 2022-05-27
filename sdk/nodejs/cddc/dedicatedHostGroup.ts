// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ApsaraDB for MyBase Dedicated Host Group resource.
 *
 * For information about ApsaraDB for MyBase Dedicated Host Group and how to use it, see [What is Dedicated Host Group](https://www.alibabacloud.com/help/doc-detail/141455.htm).
 *
 * > **NOTE:** Available in v1.132.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const vpc = new alicloud.vpc.Network("vpc", {
 *     vpcName: "tf_test_foo",
 *     cidrBlock: "172.16.0.0/12",
 * });
 * const _default = new alicloud.cddc.DedicatedHostGroup("default", {
 *     engine: "MongoDB",
 *     vpcId: vpc.id,
 *     cpuAllocationRatio: 101,
 *     memAllocationRatio: 50,
 *     diskAllocationRatio: 200,
 *     allocationPolicy: "Evenly",
 *     hostReplacePolicy: "Manual",
 *     dedicatedHostGroupDesc: "tf-testaccDesc",
 * });
 * ```
 *
 * ## Import
 *
 * ApsaraDB for MyBase Dedicated Host Group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cddc/dedicatedHostGroup:DedicatedHostGroup example <id>
 * ```
 */
export class DedicatedHostGroup extends pulumi.CustomResource {
    /**
     * Get an existing DedicatedHostGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DedicatedHostGroupState, opts?: pulumi.CustomResourceOptions): DedicatedHostGroup {
        return new DedicatedHostGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cddc/dedicatedHostGroup:DedicatedHostGroup';

    /**
     * Returns true if the given object is an instance of DedicatedHostGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DedicatedHostGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DedicatedHostGroup.__pulumiType;
    }

    /**
     * AThe policy that is used to allocate resources in the dedicated cluster. Valid values:`Evenly`,`Intensively`
     */
    public readonly allocationPolicy!: pulumi.Output<string>;
    /**
     * The CPU overcommitment ratio of the dedicated cluster.Valid values: 100 to 300. Default value: 200.
     */
    public readonly cpuAllocationRatio!: pulumi.Output<number>;
    /**
     * The name of the dedicated cluster. The name must be 1 to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
     */
    public readonly dedicatedHostGroupDesc!: pulumi.Output<string | undefined>;
    /**
     * The Disk Allocation Ratio of the Dedicated Host Group. **NOTE:** When `engine = SQLServer`, this attribute does not support to set.
     */
    public readonly diskAllocationRatio!: pulumi.Output<number>;
    /**
     * Database Engine Type.The database engine of the dedicated cluster. Valid values:`Redis`, `SQLServer`, `MySQL`, `PostgreSQL`, `MongoDB`
     */
    public readonly engine!: pulumi.Output<string>;
    /**
     * The policy based on which the system handles host failures. Valid values:`Auto`,`Manual`
     */
    public readonly hostReplacePolicy!: pulumi.Output<string>;
    /**
     * The Memory Allocation Ratio of the Dedicated Host Group.
     */
    public readonly memAllocationRatio!: pulumi.Output<number>;
    /**
     * Whether to enable the feature that allows you to have OS permissions on the hosts in the dedicated cluster. Valid values: `true` and `false`.
     * **NOTE:** The `openPermission` should be `true` when `engine = "SQLServer"`
     */
    public readonly openPermission!: pulumi.Output<boolean>;
    /**
     * The virtual private cloud (VPC) ID of the dedicated cluster.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a DedicatedHostGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DedicatedHostGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DedicatedHostGroupArgs | DedicatedHostGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DedicatedHostGroupState | undefined;
            resourceInputs["allocationPolicy"] = state ? state.allocationPolicy : undefined;
            resourceInputs["cpuAllocationRatio"] = state ? state.cpuAllocationRatio : undefined;
            resourceInputs["dedicatedHostGroupDesc"] = state ? state.dedicatedHostGroupDesc : undefined;
            resourceInputs["diskAllocationRatio"] = state ? state.diskAllocationRatio : undefined;
            resourceInputs["engine"] = state ? state.engine : undefined;
            resourceInputs["hostReplacePolicy"] = state ? state.hostReplacePolicy : undefined;
            resourceInputs["memAllocationRatio"] = state ? state.memAllocationRatio : undefined;
            resourceInputs["openPermission"] = state ? state.openPermission : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as DedicatedHostGroupArgs | undefined;
            if ((!args || args.engine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engine'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["allocationPolicy"] = args ? args.allocationPolicy : undefined;
            resourceInputs["cpuAllocationRatio"] = args ? args.cpuAllocationRatio : undefined;
            resourceInputs["dedicatedHostGroupDesc"] = args ? args.dedicatedHostGroupDesc : undefined;
            resourceInputs["diskAllocationRatio"] = args ? args.diskAllocationRatio : undefined;
            resourceInputs["engine"] = args ? args.engine : undefined;
            resourceInputs["hostReplacePolicy"] = args ? args.hostReplacePolicy : undefined;
            resourceInputs["memAllocationRatio"] = args ? args.memAllocationRatio : undefined;
            resourceInputs["openPermission"] = args ? args.openPermission : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DedicatedHostGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DedicatedHostGroup resources.
 */
export interface DedicatedHostGroupState {
    /**
     * AThe policy that is used to allocate resources in the dedicated cluster. Valid values:`Evenly`,`Intensively`
     */
    allocationPolicy?: pulumi.Input<string>;
    /**
     * The CPU overcommitment ratio of the dedicated cluster.Valid values: 100 to 300. Default value: 200.
     */
    cpuAllocationRatio?: pulumi.Input<number>;
    /**
     * The name of the dedicated cluster. The name must be 1 to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
     */
    dedicatedHostGroupDesc?: pulumi.Input<string>;
    /**
     * The Disk Allocation Ratio of the Dedicated Host Group. **NOTE:** When `engine = SQLServer`, this attribute does not support to set.
     */
    diskAllocationRatio?: pulumi.Input<number>;
    /**
     * Database Engine Type.The database engine of the dedicated cluster. Valid values:`Redis`, `SQLServer`, `MySQL`, `PostgreSQL`, `MongoDB`
     */
    engine?: pulumi.Input<string>;
    /**
     * The policy based on which the system handles host failures. Valid values:`Auto`,`Manual`
     */
    hostReplacePolicy?: pulumi.Input<string>;
    /**
     * The Memory Allocation Ratio of the Dedicated Host Group.
     */
    memAllocationRatio?: pulumi.Input<number>;
    /**
     * Whether to enable the feature that allows you to have OS permissions on the hosts in the dedicated cluster. Valid values: `true` and `false`.
     * **NOTE:** The `openPermission` should be `true` when `engine = "SQLServer"`
     */
    openPermission?: pulumi.Input<boolean>;
    /**
     * The virtual private cloud (VPC) ID of the dedicated cluster.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DedicatedHostGroup resource.
 */
export interface DedicatedHostGroupArgs {
    /**
     * AThe policy that is used to allocate resources in the dedicated cluster. Valid values:`Evenly`,`Intensively`
     */
    allocationPolicy?: pulumi.Input<string>;
    /**
     * The CPU overcommitment ratio of the dedicated cluster.Valid values: 100 to 300. Default value: 200.
     */
    cpuAllocationRatio?: pulumi.Input<number>;
    /**
     * The name of the dedicated cluster. The name must be 1 to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
     */
    dedicatedHostGroupDesc?: pulumi.Input<string>;
    /**
     * The Disk Allocation Ratio of the Dedicated Host Group. **NOTE:** When `engine = SQLServer`, this attribute does not support to set.
     */
    diskAllocationRatio?: pulumi.Input<number>;
    /**
     * Database Engine Type.The database engine of the dedicated cluster. Valid values:`Redis`, `SQLServer`, `MySQL`, `PostgreSQL`, `MongoDB`
     */
    engine: pulumi.Input<string>;
    /**
     * The policy based on which the system handles host failures. Valid values:`Auto`,`Manual`
     */
    hostReplacePolicy?: pulumi.Input<string>;
    /**
     * The Memory Allocation Ratio of the Dedicated Host Group.
     */
    memAllocationRatio?: pulumi.Input<number>;
    /**
     * Whether to enable the feature that allows you to have OS permissions on the hosts in the dedicated cluster. Valid values: `true` and `false`.
     * **NOTE:** The `openPermission` should be `true` when `engine = "SQLServer"`
     */
    openPermission?: pulumi.Input<boolean>;
    /**
     * The virtual private cloud (VPC) ID of the dedicated cluster.
     */
    vpcId: pulumi.Input<string>;
}
