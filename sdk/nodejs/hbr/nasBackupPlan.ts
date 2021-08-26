// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a HBR Nas Backup Plan resource.
 *
 * For information about HBR Nas Backup Plan and how to use it, see [What is Nas Backup Plan](https://www.alibabacloud.com/help/doc-detail/132248.htm).
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
 * const example = new alicloud.hbr.NasBackupPlan("example", {
 *     backupType: "COMPLETE",
 *     createTime: "1603163444",
 *     exclude: `  ["/home/exclude"]
 *   `,
 *     fileSystemId: "031cf4964f",
 *     include: `  ["/home/include"]
 *   `,
 *     nasBackupPlanName: "example_value",
 *     paths: [
 *         "/home",
 *         "/var",
 *     ],
 *     retention: "1",
 *     schedule: "I|1602673264|PT2H",
 *     speedLimit: "I|1602673264|PT2H",
 *     vaultId: "v-0003gxoksflhu46w185s",
 * });
 * ```
 *
 * ## Import
 *
 * HBR Nas Backup Plan can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:hbr/nasBackupPlan:NasBackupPlan example <id>
 * ```
 */
export class NasBackupPlan extends pulumi.CustomResource {
    /**
     * Get an existing NasBackupPlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NasBackupPlanState, opts?: pulumi.CustomResourceOptions): NasBackupPlan {
        return new NasBackupPlan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:hbr/nasBackupPlan:NasBackupPlan';

    /**
     * Returns true if the given object is an instance of NasBackupPlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NasBackupPlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NasBackupPlan.__pulumiType;
    }

    /**
     * Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
     */
    public readonly backupType!: pulumi.Output<string>;
    /**
     * File System Creation Time. Unix Time Seconds.
     */
    public readonly createTime!: pulumi.Output<string>;
    public readonly detail!: pulumi.Output<string | undefined>;
    /**
     * Whether to Disable the Backup Task. Valid Values: true, false.
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * The exclude path. String of Json List, most 255 Characters. e.g. `"[\"/var\"]"`
     */
    public readonly exclude!: pulumi.Output<string | undefined>;
    /**
     * The File System ID.
     */
    public readonly fileSystemId!: pulumi.Output<string | undefined>;
    /**
     * The include path. String of Json List, most 255 Characters. e.g. `"[\"/home/work\"]"`
     */
    public readonly include!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public readonly nasBackupPlanName!: pulumi.Output<string>;
    /**
     * Options. NAS Backup Plan Does Not Support Yet.
     */
    public readonly options!: pulumi.Output<string | undefined>;
    /**
     * Backup Path. Up to 65536 Characters. e.g.`["/home", "/var"]`
     */
    public readonly paths!: pulumi.Output<string[] | undefined>;
    /**
     * Backup Retention Period, the Minimum Value of 1.
     */
    public readonly retention!: pulumi.Output<string>;
    /**
     * The Backup Policy. Formats: I | {Range Specified by the StartTime }|{ Interval}\n* The Time Range Specified by the StartTime Backup Start Time in Unix Time Seconds.\n* Interval ISO8601 Time Intervals. For Example:\n**PT1H Interval for an Hour.\n**P1D Interval Day.\nMeaning from {Range Specified by the Starttime} Every {Interval} of the Time Where We Took Backups Once a Task. Does Not Compensate the Has Elapsed Time the Backup Task. If the Last Backup Has Not Been Completed without Triggering the next Backup.
     */
    public readonly schedule!: pulumi.Output<string>;
    /**
     * flow control. The format is: {start}|{end}|{bandwidth} * start starting hour * end end hour * bandwidth limit rate, in KiB ** Use | to separate multiple flow control configurations; ** Multiple flow control configurations are not allowed to have overlapping times.
     */
    public readonly speedLimit!: pulumi.Output<string | undefined>;
    public readonly updatePaths!: pulumi.Output<boolean | undefined>;
    public readonly vaultId!: pulumi.Output<string | undefined>;

    /**
     * Create a NasBackupPlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NasBackupPlanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NasBackupPlanArgs | NasBackupPlanState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NasBackupPlanState | undefined;
            inputs["backupType"] = state ? state.backupType : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["detail"] = state ? state.detail : undefined;
            inputs["disabled"] = state ? state.disabled : undefined;
            inputs["exclude"] = state ? state.exclude : undefined;
            inputs["fileSystemId"] = state ? state.fileSystemId : undefined;
            inputs["include"] = state ? state.include : undefined;
            inputs["nasBackupPlanName"] = state ? state.nasBackupPlanName : undefined;
            inputs["options"] = state ? state.options : undefined;
            inputs["paths"] = state ? state.paths : undefined;
            inputs["retention"] = state ? state.retention : undefined;
            inputs["schedule"] = state ? state.schedule : undefined;
            inputs["speedLimit"] = state ? state.speedLimit : undefined;
            inputs["updatePaths"] = state ? state.updatePaths : undefined;
            inputs["vaultId"] = state ? state.vaultId : undefined;
        } else {
            const args = argsOrState as NasBackupPlanArgs | undefined;
            if ((!args || args.createTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'createTime'");
            }
            if ((!args || args.nasBackupPlanName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nasBackupPlanName'");
            }
            if ((!args || args.retention === undefined) && !opts.urn) {
                throw new Error("Missing required property 'retention'");
            }
            if ((!args || args.schedule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schedule'");
            }
            inputs["backupType"] = args ? args.backupType : undefined;
            inputs["createTime"] = args ? args.createTime : undefined;
            inputs["detail"] = args ? args.detail : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["exclude"] = args ? args.exclude : undefined;
            inputs["fileSystemId"] = args ? args.fileSystemId : undefined;
            inputs["include"] = args ? args.include : undefined;
            inputs["nasBackupPlanName"] = args ? args.nasBackupPlanName : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["paths"] = args ? args.paths : undefined;
            inputs["retention"] = args ? args.retention : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["speedLimit"] = args ? args.speedLimit : undefined;
            inputs["updatePaths"] = args ? args.updatePaths : undefined;
            inputs["vaultId"] = args ? args.vaultId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NasBackupPlan.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NasBackupPlan resources.
 */
export interface NasBackupPlanState {
    /**
     * Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
     */
    readonly backupType?: pulumi.Input<string>;
    /**
     * File System Creation Time. Unix Time Seconds.
     */
    readonly createTime?: pulumi.Input<string>;
    readonly detail?: pulumi.Input<string>;
    /**
     * Whether to Disable the Backup Task. Valid Values: true, false.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * The exclude path. String of Json List, most 255 Characters. e.g. `"[\"/var\"]"`
     */
    readonly exclude?: pulumi.Input<string>;
    /**
     * The File System ID.
     */
    readonly fileSystemId?: pulumi.Input<string>;
    /**
     * The include path. String of Json List, most 255 Characters. e.g. `"[\"/home/work\"]"`
     */
    readonly include?: pulumi.Input<string>;
    /**
     * The name of the resource.
     */
    readonly nasBackupPlanName?: pulumi.Input<string>;
    /**
     * Options. NAS Backup Plan Does Not Support Yet.
     */
    readonly options?: pulumi.Input<string>;
    /**
     * Backup Path. Up to 65536 Characters. e.g.`["/home", "/var"]`
     */
    readonly paths?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Backup Retention Period, the Minimum Value of 1.
     */
    readonly retention?: pulumi.Input<string>;
    /**
     * The Backup Policy. Formats: I | {Range Specified by the StartTime }|{ Interval}\n* The Time Range Specified by the StartTime Backup Start Time in Unix Time Seconds.\n* Interval ISO8601 Time Intervals. For Example:\n**PT1H Interval for an Hour.\n**P1D Interval Day.\nMeaning from {Range Specified by the Starttime} Every {Interval} of the Time Where We Took Backups Once a Task. Does Not Compensate the Has Elapsed Time the Backup Task. If the Last Backup Has Not Been Completed without Triggering the next Backup.
     */
    readonly schedule?: pulumi.Input<string>;
    /**
     * flow control. The format is: {start}|{end}|{bandwidth} * start starting hour * end end hour * bandwidth limit rate, in KiB ** Use | to separate multiple flow control configurations; ** Multiple flow control configurations are not allowed to have overlapping times.
     */
    readonly speedLimit?: pulumi.Input<string>;
    readonly updatePaths?: pulumi.Input<boolean>;
    readonly vaultId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NasBackupPlan resource.
 */
export interface NasBackupPlanArgs {
    /**
     * Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
     */
    readonly backupType?: pulumi.Input<string>;
    /**
     * File System Creation Time. Unix Time Seconds.
     */
    readonly createTime: pulumi.Input<string>;
    readonly detail?: pulumi.Input<string>;
    /**
     * Whether to Disable the Backup Task. Valid Values: true, false.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * The exclude path. String of Json List, most 255 Characters. e.g. `"[\"/var\"]"`
     */
    readonly exclude?: pulumi.Input<string>;
    /**
     * The File System ID.
     */
    readonly fileSystemId?: pulumi.Input<string>;
    /**
     * The include path. String of Json List, most 255 Characters. e.g. `"[\"/home/work\"]"`
     */
    readonly include?: pulumi.Input<string>;
    /**
     * The name of the resource.
     */
    readonly nasBackupPlanName: pulumi.Input<string>;
    /**
     * Options. NAS Backup Plan Does Not Support Yet.
     */
    readonly options?: pulumi.Input<string>;
    /**
     * Backup Path. Up to 65536 Characters. e.g.`["/home", "/var"]`
     */
    readonly paths?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Backup Retention Period, the Minimum Value of 1.
     */
    readonly retention: pulumi.Input<string>;
    /**
     * The Backup Policy. Formats: I | {Range Specified by the StartTime }|{ Interval}\n* The Time Range Specified by the StartTime Backup Start Time in Unix Time Seconds.\n* Interval ISO8601 Time Intervals. For Example:\n**PT1H Interval for an Hour.\n**P1D Interval Day.\nMeaning from {Range Specified by the Starttime} Every {Interval} of the Time Where We Took Backups Once a Task. Does Not Compensate the Has Elapsed Time the Backup Task. If the Last Backup Has Not Been Completed without Triggering the next Backup.
     */
    readonly schedule: pulumi.Input<string>;
    /**
     * flow control. The format is: {start}|{end}|{bandwidth} * start starting hour * end end hour * bandwidth limit rate, in KiB ** Use | to separate multiple flow control configurations; ** Multiple flow control configurations are not allowed to have overlapping times.
     */
    readonly speedLimit?: pulumi.Input<string>;
    readonly updatePaths?: pulumi.Input<boolean>;
    readonly vaultId?: pulumi.Input<string>;
}
