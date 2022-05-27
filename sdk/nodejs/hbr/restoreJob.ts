// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Hybrid Backup Recovery (HBR) Restore Job resource.
 *
 * For information about Hybrid Backup Recovery (HBR) Restore Job and how to use it, see [What is Restore Job](https://www.alibabacloud.com/help/doc-detail/186575.htm).
 *
 * > **NOTE:** Available in v1.133.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultEcsBackupPlans = alicloud.hbr.getEcsBackupPlans({
 *     nameRegex: "plan-tf-used-dont-delete",
 * });
 * const defaultOssBackupPlans = alicloud.hbr.getOssBackupPlans({
 *     nameRegex: "plan-tf-used-dont-delete",
 * });
 * const defaultNasBackupPlans = alicloud.hbr.getNasBackupPlans({
 *     nameRegex: "plan-tf-used-dont-delete",
 * });
 * const ecsSnapshots = Promise.all([defaultEcsBackupPlans, defaultEcsBackupPlans]).then(([defaultEcsBackupPlans, defaultEcsBackupPlans1]) => alicloud.hbr.getSnapshots({
 *     sourceType: "ECS_FILE",
 *     vaultId: defaultEcsBackupPlans.plans?[0]?.vaultId,
 *     instanceId: defaultEcsBackupPlans1.plans?[0]?.instanceId,
 * }));
 * const ossSnapshots = Promise.all([defaultOssBackupPlans, defaultOssBackupPlans]).then(([defaultOssBackupPlans, defaultOssBackupPlans1]) => alicloud.hbr.getSnapshots({
 *     sourceType: "OSS",
 *     vaultId: defaultOssBackupPlans.plans?[0]?.vaultId,
 *     bucket: defaultOssBackupPlans1.plans?[0]?.bucket,
 * }));
 * const nasSnapshots = Promise.all([defaultNasBackupPlans, defaultNasBackupPlans, defaultNasBackupPlans]).then(([defaultNasBackupPlans, defaultNasBackupPlans1, defaultNasBackupPlans2]) => alicloud.hbr.getSnapshots({
 *     sourceType: "NAS",
 *     vaultId: defaultNasBackupPlans.plans?[0]?.vaultId,
 *     fileSystemId: defaultNasBackupPlans1.plans?[0]?.fileSystemId,
 *     createTime: defaultNasBackupPlans2.plans?[0]?.createTime,
 * }));
 * const nasJob = new alicloud.hbr.RestoreJob("nasJob", {
 *     snapshotHash: nasSnapshots.then(nasSnapshots => nasSnapshots.snapshots?[0]?.snapshotHash),
 *     vaultId: defaultNasBackupPlans.then(defaultNasBackupPlans => defaultNasBackupPlans.plans?[0]?.vaultId),
 *     sourceType: "NAS",
 *     restoreType: "NAS",
 *     snapshotId: nasSnapshots.then(nasSnapshots => nasSnapshots.snapshots?[0]?.snapshotId),
 *     targetFileSystemId: defaultNasBackupPlans.then(defaultNasBackupPlans => defaultNasBackupPlans.plans?[0]?.fileSystemId),
 *     targetCreateTime: defaultNasBackupPlans.then(defaultNasBackupPlans => defaultNasBackupPlans.plans?[0]?.createTime),
 *     targetPath: "/",
 *     options: "    {\"includes\":[], \"excludes\":[]}\n",
 * });
 * const ossJob = new alicloud.hbr.RestoreJob("ossJob", {
 *     snapshotHash: ossSnapshots.then(ossSnapshots => ossSnapshots.snapshots?[0]?.snapshotHash),
 *     vaultId: defaultOssBackupPlans.then(defaultOssBackupPlans => defaultOssBackupPlans.plans?[0]?.vaultId),
 *     sourceType: "OSS",
 *     restoreType: "OSS",
 *     snapshotId: ossSnapshots.then(ossSnapshots => ossSnapshots.snapshots?[0]?.snapshotId),
 *     targetBucket: defaultOssBackupPlans.then(defaultOssBackupPlans => defaultOssBackupPlans.plans?[0]?.bucket),
 *     targetPrefix: "",
 *     options: "    {\"includes\":[], \"excludes\":[]}\n",
 * });
 * const ecsJob = new alicloud.hbr.RestoreJob("ecsJob", {
 *     snapshotHash: ecsSnapshots.then(ecsSnapshots => ecsSnapshots.snapshots?[0]?.snapshotHash),
 *     vaultId: defaultEcsBackupPlans.then(defaultEcsBackupPlans => defaultEcsBackupPlans.plans?[0]?.vaultId),
 *     sourceType: "ECS_FILE",
 *     restoreType: "ECS_FILE",
 *     snapshotId: ecsSnapshots.then(ecsSnapshots => ecsSnapshots.snapshots?[0]?.snapshotId),
 *     targetInstanceId: defaultEcsBackupPlans.then(defaultEcsBackupPlans => defaultEcsBackupPlans.plans?[0]?.instanceId),
 *     targetPath: "/",
 * });
 * ```
 *
 * > **NOTE:** This resource can only be created, cannot be modified or deleted. Therefore, any modification of the resource attribute will not affect exist resource.
 *
 * ## Import
 *
 * Hybrid Backup Recovery (HBR) Restore Job can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:hbr/restoreJob:RestoreJob example <restore_job_id>:<restore_type>
 * ```
 */
export class RestoreJob extends pulumi.CustomResource {
    /**
     * Get an existing RestoreJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RestoreJobState, opts?: pulumi.CustomResourceOptions): RestoreJob {
        return new RestoreJob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:hbr/restoreJob:RestoreJob';

    /**
     * Returns true if the given object is an instance of RestoreJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RestoreJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RestoreJob.__pulumiType;
    }

    /**
     * The exclude path. **NOTE:** Invalid while sourceType equals `OSS` or `NAS`. It's a json string with format:`["/excludePath]`, up to 255 characters. **WARNING:** If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
     */
    public readonly exclude!: pulumi.Output<string | undefined>;
    /**
     * The include path. **NOTE:** Invalid while sourceType equals `OSS` or `NAS`. It's a json string with format:`["/includePath"]`, Up to 255 characters. **WARNING:** The field is required while sourceType equals `OTS_TABLE` which means source table name. If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
     */
    public readonly include!: pulumi.Output<string | undefined>;
    /**
     * Recovery options. **NOTE:** Required while sourceType equals `OSS` or `NAS`, invalid while sourceType equals `ECS_FILE`. It's a json string with format:`"{"includes":[],"excludes":[]}",`. Recovery options. When restores OTS_TABLE and real target time is the rangEnd time of the snapshot, it should be a string with format: `{"UI_TargetTime":1650032529018`}`
     */
    public readonly options!: pulumi.Output<string | undefined>;
    /**
     * Restore Job ID. It's the unique key of this resource, if you want to set this argument by yourself, you must specify a unique keyword that never appears.
     */
    public readonly restoreJobId!: pulumi.Output<string>;
    /**
     * The type of recovery destination. Valid values: `ECS_FILE`, `NAS`, `OSS`,`OTS_TABLE`,`UDM_ECS_ROLLBACK`. **Note**: Currently, there is a one-to-one correspondence between the data source type with the recovery destination type.
     */
    public readonly restoreType!: pulumi.Output<string>;
    /**
     * The hashcode of Snapshot.
     */
    public readonly snapshotHash!: pulumi.Output<string>;
    /**
     * The ID of Snapshot.
     */
    public readonly snapshotId!: pulumi.Output<string>;
    /**
     * The type of data source. Valid values: `ECS_FILE`, `NAS`, `OSS`,`OTS_TABLE`,`UDM_ECS`.
     */
    public readonly sourceType!: pulumi.Output<string>;
    /**
     * The Restore Job Status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The target name of OSS bucket. **NOTE:** Required while sourceType equals `OSS`,
     */
    public readonly targetBucket!: pulumi.Output<string | undefined>;
    /**
     * The target client ID.
     */
    public readonly targetClientId!: pulumi.Output<string | undefined>;
    /**
     * The creation time of destination File System. **NOTE:** While sourceType equals `NAS`, this parameter must be set. **Note:** The time format of the API adopts the ISO 8601 format, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
     */
    public readonly targetCreateTime!: pulumi.Output<string | undefined>;
    /**
     * The target data source ID.
     */
    public readonly targetDataSourceId!: pulumi.Output<string | undefined>;
    /**
     * The ID of destination File System. **NOTE:** Required while sourceType equals `NAS`
     */
    public readonly targetFileSystemId!: pulumi.Output<string | undefined>;
    /**
     * The target ID of ECS instance. **NOTE:** Required while sourceType equals `ECS_FILE`
     */
    public readonly targetInstanceId!: pulumi.Output<string | undefined>;
    /**
     * The name of the Table store instance to which you want to restore data.**WARNING:** Required while sourceType equals `OTS_TABLE`.
     */
    public readonly targetInstanceName!: pulumi.Output<string | undefined>;
    /**
     * The target file path of (ECS) instance. **WARNING:** Required while sourceType equals `NAS` or `ECS_FILE`, If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
     */
    public readonly targetPath!: pulumi.Output<string | undefined>;
    /**
     * The target prefix of the OSS object. **WARNING:** Required while sourceType equals `OSS`. If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
     */
    public readonly targetPrefix!: pulumi.Output<string | undefined>;
    /**
     * The name of the table that stores the restored data. **WARNING:** Required while sourceType equals `OTS_TABLE`.
     */
    public readonly targetTableName!: pulumi.Output<string | undefined>;
    /**
     * The time when data is restored to the Table store instance. This value is a UNIX timestamp. Unit: seconds. **WARNING:** Required while sourceType equals `OTS_TABLE`. **Note:** The time when data is restored to the Tablestore instance. It should be 0 if restores data at the rangEnd time of the snapshot.
     */
    public readonly targetTime!: pulumi.Output<string | undefined>;
    /**
     * The full machine backup details.
     */
    public readonly udmDetail!: pulumi.Output<string | undefined>;
    /**
     * The ID of backup vault.
     */
    public readonly vaultId!: pulumi.Output<string>;

    /**
     * Create a RestoreJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RestoreJobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RestoreJobArgs | RestoreJobState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RestoreJobState | undefined;
            resourceInputs["exclude"] = state ? state.exclude : undefined;
            resourceInputs["include"] = state ? state.include : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["restoreJobId"] = state ? state.restoreJobId : undefined;
            resourceInputs["restoreType"] = state ? state.restoreType : undefined;
            resourceInputs["snapshotHash"] = state ? state.snapshotHash : undefined;
            resourceInputs["snapshotId"] = state ? state.snapshotId : undefined;
            resourceInputs["sourceType"] = state ? state.sourceType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["targetBucket"] = state ? state.targetBucket : undefined;
            resourceInputs["targetClientId"] = state ? state.targetClientId : undefined;
            resourceInputs["targetCreateTime"] = state ? state.targetCreateTime : undefined;
            resourceInputs["targetDataSourceId"] = state ? state.targetDataSourceId : undefined;
            resourceInputs["targetFileSystemId"] = state ? state.targetFileSystemId : undefined;
            resourceInputs["targetInstanceId"] = state ? state.targetInstanceId : undefined;
            resourceInputs["targetInstanceName"] = state ? state.targetInstanceName : undefined;
            resourceInputs["targetPath"] = state ? state.targetPath : undefined;
            resourceInputs["targetPrefix"] = state ? state.targetPrefix : undefined;
            resourceInputs["targetTableName"] = state ? state.targetTableName : undefined;
            resourceInputs["targetTime"] = state ? state.targetTime : undefined;
            resourceInputs["udmDetail"] = state ? state.udmDetail : undefined;
            resourceInputs["vaultId"] = state ? state.vaultId : undefined;
        } else {
            const args = argsOrState as RestoreJobArgs | undefined;
            if ((!args || args.restoreType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restoreType'");
            }
            if ((!args || args.snapshotHash === undefined) && !opts.urn) {
                throw new Error("Missing required property 'snapshotHash'");
            }
            if ((!args || args.snapshotId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'snapshotId'");
            }
            if ((!args || args.sourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceType'");
            }
            if ((!args || args.vaultId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vaultId'");
            }
            resourceInputs["exclude"] = args ? args.exclude : undefined;
            resourceInputs["include"] = args ? args.include : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["restoreJobId"] = args ? args.restoreJobId : undefined;
            resourceInputs["restoreType"] = args ? args.restoreType : undefined;
            resourceInputs["snapshotHash"] = args ? args.snapshotHash : undefined;
            resourceInputs["snapshotId"] = args ? args.snapshotId : undefined;
            resourceInputs["sourceType"] = args ? args.sourceType : undefined;
            resourceInputs["targetBucket"] = args ? args.targetBucket : undefined;
            resourceInputs["targetClientId"] = args ? args.targetClientId : undefined;
            resourceInputs["targetCreateTime"] = args ? args.targetCreateTime : undefined;
            resourceInputs["targetDataSourceId"] = args ? args.targetDataSourceId : undefined;
            resourceInputs["targetFileSystemId"] = args ? args.targetFileSystemId : undefined;
            resourceInputs["targetInstanceId"] = args ? args.targetInstanceId : undefined;
            resourceInputs["targetInstanceName"] = args ? args.targetInstanceName : undefined;
            resourceInputs["targetPath"] = args ? args.targetPath : undefined;
            resourceInputs["targetPrefix"] = args ? args.targetPrefix : undefined;
            resourceInputs["targetTableName"] = args ? args.targetTableName : undefined;
            resourceInputs["targetTime"] = args ? args.targetTime : undefined;
            resourceInputs["udmDetail"] = args ? args.udmDetail : undefined;
            resourceInputs["vaultId"] = args ? args.vaultId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RestoreJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RestoreJob resources.
 */
export interface RestoreJobState {
    /**
     * The exclude path. **NOTE:** Invalid while sourceType equals `OSS` or `NAS`. It's a json string with format:`["/excludePath]`, up to 255 characters. **WARNING:** If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
     */
    exclude?: pulumi.Input<string>;
    /**
     * The include path. **NOTE:** Invalid while sourceType equals `OSS` or `NAS`. It's a json string with format:`["/includePath"]`, Up to 255 characters. **WARNING:** The field is required while sourceType equals `OTS_TABLE` which means source table name. If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
     */
    include?: pulumi.Input<string>;
    /**
     * Recovery options. **NOTE:** Required while sourceType equals `OSS` or `NAS`, invalid while sourceType equals `ECS_FILE`. It's a json string with format:`"{"includes":[],"excludes":[]}",`. Recovery options. When restores OTS_TABLE and real target time is the rangEnd time of the snapshot, it should be a string with format: `{"UI_TargetTime":1650032529018`}`
     */
    options?: pulumi.Input<string>;
    /**
     * Restore Job ID. It's the unique key of this resource, if you want to set this argument by yourself, you must specify a unique keyword that never appears.
     */
    restoreJobId?: pulumi.Input<string>;
    /**
     * The type of recovery destination. Valid values: `ECS_FILE`, `NAS`, `OSS`,`OTS_TABLE`,`UDM_ECS_ROLLBACK`. **Note**: Currently, there is a one-to-one correspondence between the data source type with the recovery destination type.
     */
    restoreType?: pulumi.Input<string>;
    /**
     * The hashcode of Snapshot.
     */
    snapshotHash?: pulumi.Input<string>;
    /**
     * The ID of Snapshot.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * The type of data source. Valid values: `ECS_FILE`, `NAS`, `OSS`,`OTS_TABLE`,`UDM_ECS`.
     */
    sourceType?: pulumi.Input<string>;
    /**
     * The Restore Job Status.
     */
    status?: pulumi.Input<string>;
    /**
     * The target name of OSS bucket. **NOTE:** Required while sourceType equals `OSS`,
     */
    targetBucket?: pulumi.Input<string>;
    /**
     * The target client ID.
     */
    targetClientId?: pulumi.Input<string>;
    /**
     * The creation time of destination File System. **NOTE:** While sourceType equals `NAS`, this parameter must be set. **Note:** The time format of the API adopts the ISO 8601 format, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
     */
    targetCreateTime?: pulumi.Input<string>;
    /**
     * The target data source ID.
     */
    targetDataSourceId?: pulumi.Input<string>;
    /**
     * The ID of destination File System. **NOTE:** Required while sourceType equals `NAS`
     */
    targetFileSystemId?: pulumi.Input<string>;
    /**
     * The target ID of ECS instance. **NOTE:** Required while sourceType equals `ECS_FILE`
     */
    targetInstanceId?: pulumi.Input<string>;
    /**
     * The name of the Table store instance to which you want to restore data.**WARNING:** Required while sourceType equals `OTS_TABLE`.
     */
    targetInstanceName?: pulumi.Input<string>;
    /**
     * The target file path of (ECS) instance. **WARNING:** Required while sourceType equals `NAS` or `ECS_FILE`, If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
     */
    targetPath?: pulumi.Input<string>;
    /**
     * The target prefix of the OSS object. **WARNING:** Required while sourceType equals `OSS`. If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
     */
    targetPrefix?: pulumi.Input<string>;
    /**
     * The name of the table that stores the restored data. **WARNING:** Required while sourceType equals `OTS_TABLE`.
     */
    targetTableName?: pulumi.Input<string>;
    /**
     * The time when data is restored to the Table store instance. This value is a UNIX timestamp. Unit: seconds. **WARNING:** Required while sourceType equals `OTS_TABLE`. **Note:** The time when data is restored to the Tablestore instance. It should be 0 if restores data at the rangEnd time of the snapshot.
     */
    targetTime?: pulumi.Input<string>;
    /**
     * The full machine backup details.
     */
    udmDetail?: pulumi.Input<string>;
    /**
     * The ID of backup vault.
     */
    vaultId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RestoreJob resource.
 */
export interface RestoreJobArgs {
    /**
     * The exclude path. **NOTE:** Invalid while sourceType equals `OSS` or `NAS`. It's a json string with format:`["/excludePath]`, up to 255 characters. **WARNING:** If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
     */
    exclude?: pulumi.Input<string>;
    /**
     * The include path. **NOTE:** Invalid while sourceType equals `OSS` or `NAS`. It's a json string with format:`["/includePath"]`, Up to 255 characters. **WARNING:** The field is required while sourceType equals `OTS_TABLE` which means source table name. If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
     */
    include?: pulumi.Input<string>;
    /**
     * Recovery options. **NOTE:** Required while sourceType equals `OSS` or `NAS`, invalid while sourceType equals `ECS_FILE`. It's a json string with format:`"{"includes":[],"excludes":[]}",`. Recovery options. When restores OTS_TABLE and real target time is the rangEnd time of the snapshot, it should be a string with format: `{"UI_TargetTime":1650032529018`}`
     */
    options?: pulumi.Input<string>;
    /**
     * Restore Job ID. It's the unique key of this resource, if you want to set this argument by yourself, you must specify a unique keyword that never appears.
     */
    restoreJobId?: pulumi.Input<string>;
    /**
     * The type of recovery destination. Valid values: `ECS_FILE`, `NAS`, `OSS`,`OTS_TABLE`,`UDM_ECS_ROLLBACK`. **Note**: Currently, there is a one-to-one correspondence between the data source type with the recovery destination type.
     */
    restoreType: pulumi.Input<string>;
    /**
     * The hashcode of Snapshot.
     */
    snapshotHash: pulumi.Input<string>;
    /**
     * The ID of Snapshot.
     */
    snapshotId: pulumi.Input<string>;
    /**
     * The type of data source. Valid values: `ECS_FILE`, `NAS`, `OSS`,`OTS_TABLE`,`UDM_ECS`.
     */
    sourceType: pulumi.Input<string>;
    /**
     * The target name of OSS bucket. **NOTE:** Required while sourceType equals `OSS`,
     */
    targetBucket?: pulumi.Input<string>;
    /**
     * The target client ID.
     */
    targetClientId?: pulumi.Input<string>;
    /**
     * The creation time of destination File System. **NOTE:** While sourceType equals `NAS`, this parameter must be set. **Note:** The time format of the API adopts the ISO 8601 format, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
     */
    targetCreateTime?: pulumi.Input<string>;
    /**
     * The target data source ID.
     */
    targetDataSourceId?: pulumi.Input<string>;
    /**
     * The ID of destination File System. **NOTE:** Required while sourceType equals `NAS`
     */
    targetFileSystemId?: pulumi.Input<string>;
    /**
     * The target ID of ECS instance. **NOTE:** Required while sourceType equals `ECS_FILE`
     */
    targetInstanceId?: pulumi.Input<string>;
    /**
     * The name of the Table store instance to which you want to restore data.**WARNING:** Required while sourceType equals `OTS_TABLE`.
     */
    targetInstanceName?: pulumi.Input<string>;
    /**
     * The target file path of (ECS) instance. **WARNING:** Required while sourceType equals `NAS` or `ECS_FILE`, If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
     */
    targetPath?: pulumi.Input<string>;
    /**
     * The target prefix of the OSS object. **WARNING:** Required while sourceType equals `OSS`. If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
     */
    targetPrefix?: pulumi.Input<string>;
    /**
     * The name of the table that stores the restored data. **WARNING:** Required while sourceType equals `OTS_TABLE`.
     */
    targetTableName?: pulumi.Input<string>;
    /**
     * The time when data is restored to the Table store instance. This value is a UNIX timestamp. Unit: seconds. **WARNING:** Required while sourceType equals `OTS_TABLE`. **Note:** The time when data is restored to the Tablestore instance. It should be 0 if restores data at the rangEnd time of the snapshot.
     */
    targetTime?: pulumi.Input<string>;
    /**
     * The full machine backup details.
     */
    udmDetail?: pulumi.Input<string>;
    /**
     * The ID of backup vault.
     */
    vaultId: pulumi.Input<string>;
}
