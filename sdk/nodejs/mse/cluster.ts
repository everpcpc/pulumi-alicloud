// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a MSE Cluster resource. It is a one-stop microservice platform for the industry's mainstream open source microservice frameworks Spring Cloud and Dubbo, providing governance center, managed registry and managed configuration center.
 *
 * > **NOTE:** Available in 1.94.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.mse.Cluster("example", {
 *     aclEntryLists: ["127.0.0.1/32"],
 *     clusterAliasName: "tf-testAccMseCluster",
 *     clusterSpecification: "MSE_SC_1_2_200_c",
 *     clusterType: "Eureka",
 *     clusterVersion: "EUREKA_1_9_3",
 *     instanceCount: 1,
 *     netType: "privatenet",
 *     pubNetworkFlow: "1",
 *     vswitchId: "vsw-123456",
 * });
 * ```
 *
 * ## Import
 *
 * MSE Cluster can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:mse/cluster:Cluster example mse-cn-0d9xxxx
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:mse/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * The whitelist.
     */
    public readonly aclEntryLists!: pulumi.Output<string[] | undefined>;
    /**
     * The alias of MSE Cluster.
     */
    public readonly clusterAliasName!: pulumi.Output<string | undefined>;
    /**
     * The engine specification of MSE Cluster. Valid values: `MSE_SC_1_2_200_c`, `MSE_SC_2`, `MSE_SC_4_8_200_c_4_200_c`, `MSE_SC_8_16_200_c`.
     */
    public readonly clusterSpecification!: pulumi.Output<string>;
    /**
     * The type of MSE Cluster.
     */
    public readonly clusterType!: pulumi.Output<string>;
    /**
     * The version of MSE Cluster.
     */
    public readonly clusterVersion!: pulumi.Output<string>;
    /**
     * The type of Disk.
     */
    public readonly diskType!: pulumi.Output<string | undefined>;
    /**
     * The count of instance.
     */
    public readonly instanceCount!: pulumi.Output<number>;
    /**
     * The type of network. Range limit: 1~5.
     */
    public readonly netType!: pulumi.Output<string>;
    /**
     * The specification of private network SLB.
     */
    public readonly privateSlbSpecification!: pulumi.Output<string | undefined>;
    /**
     * The public network bandwidth. `0` means no access to the public network.
     */
    public readonly pubNetworkFlow!: pulumi.Output<string | undefined>;
    /**
     * The specification of public network SLB.
     */
    public readonly pubSlbSpecification!: pulumi.Output<string | undefined>;
    /**
     * The status of MSE Cluster.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The id of VSwitch.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterState | undefined;
            inputs["aclEntryLists"] = state ? state.aclEntryLists : undefined;
            inputs["clusterAliasName"] = state ? state.clusterAliasName : undefined;
            inputs["clusterSpecification"] = state ? state.clusterSpecification : undefined;
            inputs["clusterType"] = state ? state.clusterType : undefined;
            inputs["clusterVersion"] = state ? state.clusterVersion : undefined;
            inputs["diskType"] = state ? state.diskType : undefined;
            inputs["instanceCount"] = state ? state.instanceCount : undefined;
            inputs["netType"] = state ? state.netType : undefined;
            inputs["privateSlbSpecification"] = state ? state.privateSlbSpecification : undefined;
            inputs["pubNetworkFlow"] = state ? state.pubNetworkFlow : undefined;
            inputs["pubSlbSpecification"] = state ? state.pubSlbSpecification : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.clusterSpecification === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'clusterSpecification'");
            }
            if ((!args || args.clusterType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'clusterType'");
            }
            if ((!args || args.clusterVersion === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'clusterVersion'");
            }
            if ((!args || args.instanceCount === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'instanceCount'");
            }
            if ((!args || args.netType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'netType'");
            }
            inputs["aclEntryLists"] = args ? args.aclEntryLists : undefined;
            inputs["clusterAliasName"] = args ? args.clusterAliasName : undefined;
            inputs["clusterSpecification"] = args ? args.clusterSpecification : undefined;
            inputs["clusterType"] = args ? args.clusterType : undefined;
            inputs["clusterVersion"] = args ? args.clusterVersion : undefined;
            inputs["diskType"] = args ? args.diskType : undefined;
            inputs["instanceCount"] = args ? args.instanceCount : undefined;
            inputs["netType"] = args ? args.netType : undefined;
            inputs["privateSlbSpecification"] = args ? args.privateSlbSpecification : undefined;
            inputs["pubNetworkFlow"] = args ? args.pubNetworkFlow : undefined;
            inputs["pubSlbSpecification"] = args ? args.pubSlbSpecification : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * The whitelist.
     */
    readonly aclEntryLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The alias of MSE Cluster.
     */
    readonly clusterAliasName?: pulumi.Input<string>;
    /**
     * The engine specification of MSE Cluster. Valid values: `MSE_SC_1_2_200_c`, `MSE_SC_2`, `MSE_SC_4_8_200_c_4_200_c`, `MSE_SC_8_16_200_c`.
     */
    readonly clusterSpecification?: pulumi.Input<string>;
    /**
     * The type of MSE Cluster.
     */
    readonly clusterType?: pulumi.Input<string>;
    /**
     * The version of MSE Cluster.
     */
    readonly clusterVersion?: pulumi.Input<string>;
    /**
     * The type of Disk.
     */
    readonly diskType?: pulumi.Input<string>;
    /**
     * The count of instance.
     */
    readonly instanceCount?: pulumi.Input<number>;
    /**
     * The type of network. Range limit: 1~5.
     */
    readonly netType?: pulumi.Input<string>;
    /**
     * The specification of private network SLB.
     */
    readonly privateSlbSpecification?: pulumi.Input<string>;
    /**
     * The public network bandwidth. `0` means no access to the public network.
     */
    readonly pubNetworkFlow?: pulumi.Input<string>;
    /**
     * The specification of public network SLB.
     */
    readonly pubSlbSpecification?: pulumi.Input<string>;
    /**
     * The status of MSE Cluster.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The id of VSwitch.
     */
    readonly vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The whitelist.
     */
    readonly aclEntryLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The alias of MSE Cluster.
     */
    readonly clusterAliasName?: pulumi.Input<string>;
    /**
     * The engine specification of MSE Cluster. Valid values: `MSE_SC_1_2_200_c`, `MSE_SC_2`, `MSE_SC_4_8_200_c_4_200_c`, `MSE_SC_8_16_200_c`.
     */
    readonly clusterSpecification: pulumi.Input<string>;
    /**
     * The type of MSE Cluster.
     */
    readonly clusterType: pulumi.Input<string>;
    /**
     * The version of MSE Cluster.
     */
    readonly clusterVersion: pulumi.Input<string>;
    /**
     * The type of Disk.
     */
    readonly diskType?: pulumi.Input<string>;
    /**
     * The count of instance.
     */
    readonly instanceCount: pulumi.Input<number>;
    /**
     * The type of network. Range limit: 1~5.
     */
    readonly netType: pulumi.Input<string>;
    /**
     * The specification of private network SLB.
     */
    readonly privateSlbSpecification?: pulumi.Input<string>;
    /**
     * The public network bandwidth. `0` means no access to the public network.
     */
    readonly pubNetworkFlow?: pulumi.Input<string>;
    /**
     * The specification of public network SLB.
     */
    readonly pubSlbSpecification?: pulumi.Input<string>;
    /**
     * The id of VSwitch.
     */
    readonly vswitchId?: pulumi.Input<string>;
}
