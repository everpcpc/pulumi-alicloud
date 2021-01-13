// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The Logtail access service is a log collection agent provided by Log Service.
 * You can use Logtail to collect logs from servers such as Alibaba Cloud Elastic
 * Compute Service (ECS) instances in real time in the Log Service console. [Refer to details](https://www.alibabacloud.com/help/doc-detail/29058.htm)
 *
 * This resource amis to attach one logtail configure to a machine group.
 *
 * > **NOTE:** One logtail configure can be attached to multiple machine groups and one machine group can attach several logtail configures.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const testProject = new alicloud.log.Project("testProject", {description: "create by terraform"});
 * const testStore = new alicloud.log.Store("testStore", {
 *     project: testProject.name,
 *     retentionPeriod: 3650,
 *     shardCount: 3,
 *     autoSplit: true,
 *     maxSplitShardCount: 60,
 *     appendMeta: true,
 * });
 * const testMachineGroup = new alicloud.log.MachineGroup("testMachineGroup", {
 *     project: testProject.name,
 *     topic: "terraform",
 *     identifyLists: [
 *         "10.0.0.1",
 *         "10.0.0.3",
 *         "10.0.0.2",
 *     ],
 * });
 * const testLogTailConfig = new alicloud.log.LogTailConfig("testLogTailConfig", {
 *     project: testProject.name,
 *     logstore: testStore.name,
 *     inputType: "file",
 *     logSample: "test",
 *     outputType: "LogService",
 *     inputDetail: `  	{
 * 		"logPath": "/logPath",
 * 		"filePattern": "access.log",
 * 		"logType": "json_log",
 * 		"topicFormat": "default",
 * 		"discardUnmatch": false,
 * 		"enableRawLog": true,
 * 		"fileEncoding": "gbk",
 * 		"maxDepth": 10
 * 	}
 * 	
 * `,
 * });
 * const testLogTailAttachment = new alicloud.log.LogTailAttachment("testLogTailAttachment", {
 *     project: testProject.name,
 *     logtailConfigName: testLogTailConfig.name,
 *     machineGroupName: testMachineGroup.name,
 * });
 * ```
 *
 * ## Import
 *
 * Logtial to machine group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:log/logTailAttachment:LogTailAttachment example tf-log:tf-log-config:tf-log-machine-group
 * ```
 */
export class LogTailAttachment extends pulumi.CustomResource {
    /**
     * Get an existing LogTailAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogTailAttachmentState, opts?: pulumi.CustomResourceOptions): LogTailAttachment {
        return new LogTailAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:log/logTailAttachment:LogTailAttachment';

    /**
     * Returns true if the given object is an instance of LogTailAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogTailAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogTailAttachment.__pulumiType;
    }

    /**
     * The Logtail configuration name, which is unique in the same project.
     */
    public readonly logtailConfigName!: pulumi.Output<string>;
    /**
     * The machine group name, which is unique in the same project.
     */
    public readonly machineGroupName!: pulumi.Output<string>;
    /**
     * The project name to the log store belongs.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a LogTailAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogTailAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogTailAttachmentArgs | LogTailAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LogTailAttachmentState | undefined;
            inputs["logtailConfigName"] = state ? state.logtailConfigName : undefined;
            inputs["machineGroupName"] = state ? state.machineGroupName : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as LogTailAttachmentArgs | undefined;
            if ((!args || args.logtailConfigName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'logtailConfigName'");
            }
            if ((!args || args.machineGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'machineGroupName'");
            }
            if ((!args || args.project === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'project'");
            }
            inputs["logtailConfigName"] = args ? args.logtailConfigName : undefined;
            inputs["machineGroupName"] = args ? args.machineGroupName : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LogTailAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogTailAttachment resources.
 */
export interface LogTailAttachmentState {
    /**
     * The Logtail configuration name, which is unique in the same project.
     */
    readonly logtailConfigName?: pulumi.Input<string>;
    /**
     * The machine group name, which is unique in the same project.
     */
    readonly machineGroupName?: pulumi.Input<string>;
    /**
     * The project name to the log store belongs.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogTailAttachment resource.
 */
export interface LogTailAttachmentArgs {
    /**
     * The Logtail configuration name, which is unique in the same project.
     */
    readonly logtailConfigName: pulumi.Input<string>;
    /**
     * The machine group name, which is unique in the same project.
     */
    readonly machineGroupName: pulumi.Input<string>;
    /**
     * The project name to the log store belongs.
     */
    readonly project: pulumi.Input<string>;
}
