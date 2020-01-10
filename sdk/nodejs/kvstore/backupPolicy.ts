// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a backup policy for ApsaraDB Redis / Memcache instance resource. 
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
 * const creation = config.get("creation") || "KVStore";
 * const multiAz = config.get("multiAz") || "false";
 * const name = config.get("name") || "kvstorebackuppolicyvpc";
 * 
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: creation,
 * });
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     availabilityZone: defaultZones.zones[0].id,
 *     cidrBlock: "172.16.0.0/24",
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultInstance = new alicloud.kvstore.Instance("default", {
 *     engineVersion: "2.8",
 *     instanceClass: "Memcache",
 *     instanceName: name,
 *     instanceType: "memcache.master.small.default",
 *     privateIp: "172.16.0.10",
 *     securityIps: ["10.0.0.1"],
 *     vswitchId: defaultSwitch.id,
 * });
 * const defaultBackupPolicy = new alicloud.kvstore.BackupPolicy("default", {
 *     backupPeriods: [
 *         "Tuesday",
 *         "Wednesday",
 *     ],
 *     backupTime: "10:00Z-11:00Z",
 *     instanceId: defaultInstance.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/kvstore_backup_policy.html.markdown.
 */
export class BackupPolicy extends pulumi.CustomResource {
    /**
     * Get an existing BackupPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackupPolicyState, opts?: pulumi.CustomResourceOptions): BackupPolicy {
        return new BackupPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:kvstore/backupPolicy:BackupPolicy';

    /**
     * Returns true if the given object is an instance of BackupPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackupPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackupPolicy.__pulumiType;
    }

    /**
     * Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
     */
    public readonly backupPeriods!: pulumi.Output<string[]>;
    /**
     * Backup time, in the format of HH:mmZ- HH:mm Z
     */
    public readonly backupTime!: pulumi.Output<string | undefined>;
    /**
     * The id of ApsaraDB for Redis or Memcache intance.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a BackupPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackupPolicyArgs | BackupPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BackupPolicyState | undefined;
            inputs["backupPeriods"] = state ? state.backupPeriods : undefined;
            inputs["backupTime"] = state ? state.backupTime : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as BackupPolicyArgs | undefined;
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            inputs["backupPeriods"] = args ? args.backupPeriods : undefined;
            inputs["backupTime"] = args ? args.backupTime : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BackupPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackupPolicy resources.
 */
export interface BackupPolicyState {
    /**
     * Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
     */
    readonly backupPeriods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Backup time, in the format of HH:mmZ- HH:mm Z
     */
    readonly backupTime?: pulumi.Input<string>;
    /**
     * The id of ApsaraDB for Redis or Memcache intance.
     */
    readonly instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackupPolicy resource.
 */
export interface BackupPolicyArgs {
    /**
     * Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
     */
    readonly backupPeriods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Backup time, in the format of HH:mmZ- HH:mm Z
     */
    readonly backupTime?: pulumi.Input<string>;
    /**
     * The id of ApsaraDB for Redis or Memcache intance.
     */
    readonly instanceId: pulumi.Input<string>;
}
