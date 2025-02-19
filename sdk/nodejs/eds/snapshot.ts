// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ECD Snapshot resource.
 *
 * For information about ECD Snapshot and how to use it, see [What is Snapshot](https://www.alibabacloud.com/help/en/elastic-desktop-service/latest/createsnapshot).
 *
 * > **NOTE:** Available in v1.169.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "example_value";
 * const defaultSimpleOfficeSite = new alicloud.eds.SimpleOfficeSite("defaultSimpleOfficeSite", {
 *     cidrBlock: "172.16.0.0/12",
 *     desktopAccessType: "Internet",
 *     officeSiteName: name,
 *     enableInternetAccess: false,
 * });
 * const defaultBundles = alicloud.eds.getBundles({
 *     bundleType: "SYSTEM",
 * });
 * const defaultEcdPolicyGroup = new alicloud.eds.EcdPolicyGroup("defaultEcdPolicyGroup", {
 *     policyGroupName: name,
 *     clipboard: "readwrite",
 *     localDrive: "read",
 *     authorizeAccessPolicyRules: [{
 *         description: "example_value",
 *         cidrIp: "1.2.3.4/24",
 *     }],
 *     authorizeSecurityPolicyRules: [{
 *         type: "inflow",
 *         policy: "accept",
 *         description: "example_value",
 *         portRange: "80/80",
 *         ipProtocol: "TCP",
 *         priority: "1",
 *         cidrIp: "0.0.0.0/0",
 *     }],
 * });
 * const defaultDesktop = new alicloud.eds.Desktop("defaultDesktop", {
 *     officeSiteId: defaultSimpleOfficeSite.id,
 *     policyGroupId: defaultEcdPolicyGroup.id,
 *     bundleId: defaultBundles.then(defaultBundles => defaultBundles.bundles?.[0]?.id),
 *     desktopName: name,
 * });
 * const defaultSnapshot = new alicloud.eds.Snapshot("defaultSnapshot", {
 *     description: name,
 *     desktopId: defaultDesktop.id,
 *     snapshotName: name,
 *     sourceDiskType: "SYSTEM",
 * });
 * ```
 *
 * ## Import
 *
 * ECD Snapshot can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:eds/snapshot:Snapshot example <id>
 * ```
 */
export class Snapshot extends pulumi.CustomResource {
    /**
     * Get an existing Snapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnapshotState, opts?: pulumi.CustomResourceOptions): Snapshot {
        return new Snapshot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eds/snapshot:Snapshot';

    /**
     * Returns true if the given object is an instance of Snapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Snapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Snapshot.__pulumiType;
    }

    /**
     * The description of the Snapshot.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Desktop.
     */
    public readonly desktopId!: pulumi.Output<string>;
    /**
     * The name of the Snapshot.
     */
    public readonly snapshotName!: pulumi.Output<string>;
    /**
     * The type of the disk for which to create a snapshot. Valid values: `SYSTEM`, `DATA`.
     */
    public readonly sourceDiskType!: pulumi.Output<string>;
    /**
     * The status of the snapshot.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Snapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnapshotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnapshotArgs | SnapshotState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnapshotState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["desktopId"] = state ? state.desktopId : undefined;
            resourceInputs["snapshotName"] = state ? state.snapshotName : undefined;
            resourceInputs["sourceDiskType"] = state ? state.sourceDiskType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as SnapshotArgs | undefined;
            if ((!args || args.desktopId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'desktopId'");
            }
            if ((!args || args.snapshotName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'snapshotName'");
            }
            if ((!args || args.sourceDiskType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceDiskType'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["desktopId"] = args ? args.desktopId : undefined;
            resourceInputs["snapshotName"] = args ? args.snapshotName : undefined;
            resourceInputs["sourceDiskType"] = args ? args.sourceDiskType : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Snapshot.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Snapshot resources.
 */
export interface SnapshotState {
    /**
     * The description of the Snapshot.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the Desktop.
     */
    desktopId?: pulumi.Input<string>;
    /**
     * The name of the Snapshot.
     */
    snapshotName?: pulumi.Input<string>;
    /**
     * The type of the disk for which to create a snapshot. Valid values: `SYSTEM`, `DATA`.
     */
    sourceDiskType?: pulumi.Input<string>;
    /**
     * The status of the snapshot.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Snapshot resource.
 */
export interface SnapshotArgs {
    /**
     * The description of the Snapshot.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the Desktop.
     */
    desktopId: pulumi.Input<string>;
    /**
     * The name of the Snapshot.
     */
    snapshotName: pulumi.Input<string>;
    /**
     * The type of the disk for which to create a snapshot. Valid values: `SYSTEM`, `DATA`.
     */
    sourceDiskType: pulumi.Input<string>;
}
