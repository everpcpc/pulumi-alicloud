// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an Alicloud Function Compute Trigger resource. Based on trigger, execute your code in response to events in Alibaba Cloud.
 *  For information about Service and how to use it, see [What is Function Compute](https://www.alibabacloud.com/help/doc-detail/52895.htm).
 *
 * > **NOTE:** The resource requires a provider field 'account_id'. See account_id.
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
 * const region = config.get("region") || "cn-hangzhou";
 * const account = config.get("account") || "12345";
 * const fooRole = new alicloud.ram.Role("fooRole", {
 *     document: `  {
 *     "Statement": [
 *       {
 *         "Action": "sts:AssumeRole",
 *         "Effect": "Allow",
 *         "Principal": {
 *           "Service": [
 *             "log.aliyuncs.com"
 *           ]
 *         }
 *       }
 *     ],
 *     "Version": "1"
 *   }
 *   
 * `,
 *     description: "this is a test",
 *     force: true,
 * });
 * const fooRolePolicyAttachment = new alicloud.ram.RolePolicyAttachment("fooRolePolicyAttachment", {
 *     roleName: fooRole.name,
 *     policyName: "AliyunLogFullAccess",
 *     policyType: "System",
 * });
 * const fooTrigger = new alicloud.fc.Trigger("fooTrigger", {
 *     service: "my-fc-service",
 *     "function": "hello-world",
 *     role: fooRole.arn,
 *     sourceArn: `acs:log:${region}:${account}:project/${alicloud_log_project.foo.name}`,
 *     type: "log",
 *     config: `    {
 *         "sourceConfig": {
 *             "project": "project-for-fc",
 *             "logstore": "project-for-fc"
 *         },
 *         "jobConfig": {
 *             "maxRetryTime": 3,
 *             "triggerInterval": 60
 *         },
 *         "functionParameter": {
 *             "a": "b",
 *             "c": "d"
 *         },
 *         "logConfig": {
 *             "project": "project-for-fc-log",
 *             "logstore": "project-for-fc-log"
 *         },
 *         "enable": true
 *     }
 *   
 * `,
 * }, {
 *     dependsOn: [fooRolePolicyAttachment],
 * });
 * ```
 *
 * MNS topic trigger:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "fctriggermnstopic";
 * const currentRegion = alicloud.getRegions({
 *     current: true,
 * });
 * const current = alicloud.getAccount({});
 * const fooProject = new alicloud.log.Project("fooProject", {description: "tf unit test"});
 * const bar = new alicloud.log.Store("bar", {
 *     project: fooProject.name,
 *     retentionPeriod: 3000,
 *     shardCount: 1,
 * });
 * const fooStore = new alicloud.log.Store("fooStore", {
 *     project: fooProject.name,
 *     retentionPeriod: 3000,
 *     shardCount: 1,
 * });
 * const fooTopic = new alicloud.mns.Topic("fooTopic", {});
 * const fooService = new alicloud.fc.Service("fooService", {internetAccess: false});
 * const fooBucket = new alicloud.oss.Bucket("fooBucket", {bucket: name});
 * // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
 * const fooBucketObject = new alicloud.oss.BucketObject("fooBucketObject", {
 *     bucket: fooBucket.id,
 *     key: "fc/hello.zip",
 *     source: "./hello.zip",
 * });
 * const fooFunction = new alicloud.fc.Function("fooFunction", {
 *     handler: "hello.handler",
 *     memorySize: 512,
 *     ossBucket: fooBucket.id,
 *     ossKey: fooBucketObject.key,
 *     runtime: "python2.7",
 *     service: fooService.name,
 * });
 * const fooRole = new alicloud.ram.Role("fooRole", {
 *     description: "this is a test",
 *     document: `  {
 *     "Statement": [
 *       {
 *         "Action": "sts:AssumeRole",
 *         "Effect": "Allow",
 *         "Principal": {
 *           "Service": [
 *             "mns.aliyuncs.com"
 *           ]
 *         }
 *       }
 *     ],
 *     "Version": "1"
 *   }
 *   
 * `,
 *     force: true,
 * });
 * const fooRolePolicyAttachment = new alicloud.ram.RolePolicyAttachment("fooRolePolicyAttachment", {
 *     policyName: "AliyunMNSNotificationRolePolicy",
 *     policyType: "System",
 *     roleName: fooRole.name,
 * });
 * const fooTrigger = new alicloud.fc.Trigger("fooTrigger", {
 *     configMns: `  {
 *     "filterTag":"testTag",
 *     "notifyContentFormat":"STREAM",
 *     "notifyStrategy":"BACKOFF_RETRY"
 *   }
 *   
 * `,
 *     "function": fooFunction.name,
 *     role: fooRole.arn,
 *     service: fooService.name,
 *     sourceArn: pulumi.all([currentRegion, current, fooTopic.name]).apply(([currentRegion, current, name]) => `acs:mns:${currentRegion.regions?.[0]?.id}:${current.id}:/topics/${name}`),
 *     type: "mns_topic",
 * }, {
 *     dependsOn: ["alicloud_ram_role_policy_attachment.foo"],
 * });
 * ```
 *
 * CDN events trigger:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "fctriggercdneventsconfig";
 * const current = alicloud.getAccount({});
 * const domain = new alicloud.cdn.DomainNew("domain", {
 *     cdnType: "web",
 *     domainName: `${name}.tf.com`,
 *     scope: "overseas",
 *     sources: [{
 *         content: "1.1.1.1",
 *         port: 80,
 *         priority: 20,
 *         type: "ipaddr",
 *         weight: 10,
 *     }],
 * });
 * const fooService = new alicloud.fc.Service("fooService", {internetAccess: false});
 * const fooBucket = new alicloud.oss.Bucket("fooBucket", {bucket: name});
 * // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
 * const fooBucketObject = new alicloud.oss.BucketObject("fooBucketObject", {
 *     bucket: fooBucket.id,
 *     key: "fc/hello.zip",
 *     source: "./hello.zip",
 * });
 * const fooFunction = new alicloud.fc.Function("fooFunction", {
 *     handler: "hello.handler",
 *     memorySize: 512,
 *     ossBucket: fooBucket.id,
 *     ossKey: fooBucketObject.key,
 *     runtime: "python2.7",
 *     service: fooService.name,
 * });
 * const fooRole = new alicloud.ram.Role("fooRole", {
 *     description: "this is a test",
 *     document: `    {
 *         "Version": "1",
 *         "Statement": [
 *             {
 *                 "Action": "cdn:Describe*",
 *                 "Resource": "*",
 *                 "Effect": "Allow",
 * 		        "Principal": {
 *                 "Service":
 *                     ["log.aliyuncs.com"]
 *                 }
 *             }
 *         ]
 *     }
 *     
 * `,
 *     force: true,
 * });
 * const fooPolicy = new alicloud.ram.Policy("fooPolicy", {
 *     description: "this is a test",
 *     document: `    {
 *         "Version": "1",
 *         "Statement": [
 *         {
 *             "Action": [
 *             "fc:InvokeFunction"
 *             ],
 *         "Resource": [
 *             "acs:fc:*:*:services/tf_cdnEvents/functions/*",
 *             "acs:fc:*:*:services/tf_cdnEvents.*&#47;functions/*"
 *         ],
 *         "Effect": "Allow"
 *         }
 *         ]
 *     }
 *     
 * `,
 *     force: true,
 * });
 * const fooRolePolicyAttachment = new alicloud.ram.RolePolicyAttachment("fooRolePolicyAttachment", {
 *     policyName: fooPolicy.name,
 *     policyType: "Custom",
 *     roleName: fooRole.name,
 * });
 * const _default = new alicloud.fc.Trigger("default", {
 *     config: pulumi.interpolate`      {"eventName":"LogFileCreated",
 *      "eventVersion":"1.0.0",
 *      "notes":"cdn events trigger",
 *      "filter":{
 *         "domain": ["${domain.domainName}"]
 *         }
 *     }
 *
 * `,
 *     "function": fooFunction.name,
 *     role: fooRole.arn,
 *     service: fooService.name,
 *     sourceArn: current.then(current => `acs:cdn:*:${current.id}`),
 *     type: "cdn_events",
 * }, {
 *     dependsOn: ["alicloud_ram_role_policy_attachment.foo"],
 * });
 * ```
 *
 * EventBridge trigger:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "fctriggereventbridgeconfig";
 * const current = alicloud.getAccount({});
 * // Please make eventbridge available and then assume a specific service-linked role, which refers to https://registry.terraform.io/providers/aliyun/alicloud/latest/docs/resources/event_bridge_service_linked_role
 * const serviceLinkedRole = new alicloud.eventbridge.ServiceLinkedRole("serviceLinkedRole", {productName: "AliyunServiceRoleForEventBridgeSendToFC"});
 * const fooService = new alicloud.fc.Service("fooService", {internetAccess: false});
 * const fooBucket = new alicloud.oss.Bucket("fooBucket", {bucket: name});
 * // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
 * const fooBucketObject = new alicloud.oss.BucketObject("fooBucketObject", {
 *     bucket: fooBucket.id,
 *     key: "fc/hello.zip",
 *     source: "./hello.zip",
 * });
 * const fooFunction = new alicloud.fc.Function("fooFunction", {
 *     handler: "hello.handler",
 *     memorySize: 512,
 *     ossBucket: fooBucket.id,
 *     ossKey: fooBucketObject.key,
 *     runtime: "python2.7",
 *     service: fooService.name,
 * });
 * const _default = new alicloud.fc.Trigger("default", {
 *     config: `    {
 *         "triggerEnable": false,
 *         "asyncInvocationType": false,
 *         "eventRuleFilterPattern": "{\\"source\\":[\\"acs.oss\\"],\\"type\\":[\\"oss:BucketCreated:PutBucket\\"]}",
 *         "eventSourceConfig": {
 *             "eventSourceType": "Default"
 *         }
 *     }
 *
 * `,
 *     "function": fooFunction.name,
 *     service: fooService.name,
 *     type: "eventbridge",
 * });
 * const mns = new alicloud.fc.Trigger("mns", {
 *     config: `    {
 *         "triggerEnable": false,
 *         "asyncInvocationType": false,
 *         "eventRuleFilterPattern": "{}",
 *         "eventSourceConfig": {
 *             "eventSourceType": "MNS",
 *             "eventSourceParameters": {
 *                 "sourceMNSParameters": {
 *                     "RegionId": "cn-hangzhou",
 *                     "QueueName": "mns-queue",
 *                     "IsBase64Decode": true
 *                 }
 *             }
 *         }
 *     }
 *
 * `,
 *     "function": fooFunction.name,
 *     service: fooService.name,
 *     type: "eventbridge",
 * });
 * const rocketmq = new alicloud.fc.Trigger("rocketmq", {
 *     config: `    {
 *         "triggerEnable": false,
 *         "asyncInvocationType": false,
 *         "eventRuleFilterPattern": "{}",
 *         "eventSourceConfig": {
 *             "eventSourceType": "RocketMQ",
 *             "eventSourceParameters": {
 *                 "sourceRocketMQParameters": {
 *                     "RegionId": "cn-hangzhou",
 *                     "InstanceId": "MQ_INST_164901546557****_BAAN****",
 *                     "GroupID": "GID_group1",
 *                     "Topic": "mytopic",
 *                     "Timestamp": 1636597951984,
 *                     "Tag": "test-tag",
 *                     "Offset": "CONSUME_FROM_LAST_OFFSET"
 *                 }
 *             }
 *         }
 *     }
 *
 * `,
 *     "function": fooFunction.name,
 *     service: fooService.name,
 *     type: "eventbridge",
 * });
 * const rabbitmq = new alicloud.fc.Trigger("rabbitmq", {
 *     config: `    {
 *         "triggerEnable": false,
 *         "asyncInvocationType": false,
 *         "eventRuleFilterPattern": "{}",
 *         "eventSourceConfig": {
 *             "eventSourceType": "RabbitMQ",
 *             "eventSourceParameters": {
 *                 "sourceRabbitMQParameters": {
 *                     "RegionId": "cn-hangzhou",
 *                     "InstanceId": "amqp-cn-****** ",
 *                     "VirtualHostName": "test-virtual",
 *                     "QueueName": "test-queue"
 *                 }
 *             }
 *         }
 *     }
 *
 * `,
 *     "function": fooFunction.name,
 *     service: fooService.name,
 *     type: "eventbridge",
 * });
 * ```
 * ## Module Support
 *
 * You can use to the existing fc module
 * to create several triggers quickly.
 *
 * ## Import
 *
 * Function Compute trigger can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:fc/trigger:Trigger foo my-fc-service:hello-world:hello-trigger
 * ```
 */
export class Trigger extends pulumi.CustomResource {
    /**
     * Get an existing Trigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TriggerState, opts?: pulumi.CustomResourceOptions): Trigger {
        return new Trigger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:fc/trigger:Trigger';

    /**
     * Returns true if the given object is an instance of Trigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Trigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Trigger.__pulumiType;
    }

    /**
     * The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
     */
    public readonly config!: pulumi.Output<string | undefined>;
    /**
     * The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
     */
    public readonly configMns!: pulumi.Output<string | undefined>;
    /**
     * The Function Compute function name.
     */
    public readonly function!: pulumi.Output<string>;
    /**
     * The date this resource was last modified.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Setting a prefix to get a only trigger name. It is conflict with "name".
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    public readonly role!: pulumi.Output<string | undefined>;
    /**
     * The Function Compute service name.
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    public readonly sourceArn!: pulumi.Output<string | undefined>;
    /**
     * The Function Compute trigger ID.
     */
    public /*out*/ readonly triggerId!: pulumi.Output<string>;
    /**
     * The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents", "eventbridge"].
     *
     * > **NOTE:** Config does not support modification when type is mns_topic.
     * > **NOTE:** type = cdn_events, available in 1.47.0+.
     * > **NOTE:** type = eventbridge, available in 1.173.0+.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Trigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TriggerArgs | TriggerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TriggerState | undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["configMns"] = state ? state.configMns : undefined;
            resourceInputs["function"] = state ? state.function : undefined;
            resourceInputs["lastModified"] = state ? state.lastModified : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["service"] = state ? state.service : undefined;
            resourceInputs["sourceArn"] = state ? state.sourceArn : undefined;
            resourceInputs["triggerId"] = state ? state.triggerId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as TriggerArgs | undefined;
            if ((!args || args.function === undefined) && !opts.urn) {
                throw new Error("Missing required property 'function'");
            }
            if ((!args || args.service === undefined) && !opts.urn) {
                throw new Error("Missing required property 'service'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["configMns"] = args ? args.configMns : undefined;
            resourceInputs["function"] = args ? args.function : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["service"] = args ? args.service : undefined;
            resourceInputs["sourceArn"] = args ? args.sourceArn : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["lastModified"] = undefined /*out*/;
            resourceInputs["triggerId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Trigger.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Trigger resources.
 */
export interface TriggerState {
    /**
     * The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
     */
    config?: pulumi.Input<string>;
    /**
     * The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
     */
    configMns?: pulumi.Input<string>;
    /**
     * The Function Compute function name.
     */
    function?: pulumi.Input<string>;
    /**
     * The date this resource was last modified.
     */
    lastModified?: pulumi.Input<string>;
    /**
     * The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
     */
    name?: pulumi.Input<string>;
    /**
     * Setting a prefix to get a only trigger name. It is conflict with "name".
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    role?: pulumi.Input<string>;
    /**
     * The Function Compute service name.
     */
    service?: pulumi.Input<string>;
    /**
     * Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    sourceArn?: pulumi.Input<string>;
    /**
     * The Function Compute trigger ID.
     */
    triggerId?: pulumi.Input<string>;
    /**
     * The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents", "eventbridge"].
     *
     * > **NOTE:** Config does not support modification when type is mns_topic.
     * > **NOTE:** type = cdn_events, available in 1.47.0+.
     * > **NOTE:** type = eventbridge, available in 1.173.0+.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Trigger resource.
 */
export interface TriggerArgs {
    /**
     * The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
     */
    config?: pulumi.Input<string>;
    /**
     * The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
     */
    configMns?: pulumi.Input<string>;
    /**
     * The Function Compute function name.
     */
    function: pulumi.Input<string>;
    /**
     * The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
     */
    name?: pulumi.Input<string>;
    /**
     * Setting a prefix to get a only trigger name. It is conflict with "name".
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    role?: pulumi.Input<string>;
    /**
     * The Function Compute service name.
     */
    service: pulumi.Input<string>;
    /**
     * Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    sourceArn?: pulumi.Input<string>;
    /**
     * The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents", "eventbridge"].
     *
     * > **NOTE:** Config does not support modification when type is mns_topic.
     * > **NOTE:** type = cdn_events, available in 1.47.0+.
     * > **NOTE:** type = eventbridge, available in 1.173.0+.
     */
    type: pulumi.Input<string>;
}
