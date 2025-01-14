// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.FC
{
    /// <summary>
    /// Provides an Alicloud Function Compute Trigger resource. Based on trigger, execute your code in response to events in Alibaba Cloud.
    ///  For information about Service and how to use it, see [What is Function Compute](https://www.alibabacloud.com/help/doc-detail/52895.htm).
    /// 
    /// &gt; **NOTE:** The resource requires a provider field 'account_id'. See account_id.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var region = config.Get("region") ?? "cn-hangzhou";
    ///     var account = config.Get("account") ?? "12345";
    ///     var fooRole = new AliCloud.Ram.Role("fooRole", new()
    ///     {
    ///         Document = @"  {
    ///     ""Statement"": [
    ///       {
    ///         ""Action"": ""sts:AssumeRole"",
    ///         ""Effect"": ""Allow"",
    ///         ""Principal"": {
    ///           ""Service"": [
    ///             ""log.aliyuncs.com""
    ///           ]
    ///         }
    ///       }
    ///     ],
    ///     ""Version"": ""1""
    ///   }
    ///   
    /// ",
    ///         Description = "this is a test",
    ///         Force = true,
    ///     });
    /// 
    ///     var fooRolePolicyAttachment = new AliCloud.Ram.RolePolicyAttachment("fooRolePolicyAttachment", new()
    ///     {
    ///         RoleName = fooRole.Name,
    ///         PolicyName = "AliyunLogFullAccess",
    ///         PolicyType = "System",
    ///     });
    /// 
    ///     var fooTrigger = new AliCloud.FC.Trigger("fooTrigger", new()
    ///     {
    ///         Service = "my-fc-service",
    ///         Function = "hello-world",
    ///         Role = fooRole.Arn,
    ///         SourceArn = $"acs:log:{region}:{account}:project/{alicloud_log_project.Foo.Name}",
    ///         Type = "log",
    ///         Config = @"    {
    ///         ""sourceConfig"": {
    ///             ""project"": ""project-for-fc"",
    ///             ""logstore"": ""project-for-fc""
    ///         },
    ///         ""jobConfig"": {
    ///             ""maxRetryTime"": 3,
    ///             ""triggerInterval"": 60
    ///         },
    ///         ""functionParameter"": {
    ///             ""a"": ""b"",
    ///             ""c"": ""d""
    ///         },
    ///         ""logConfig"": {
    ///             ""project"": ""project-for-fc-log"",
    ///             ""logstore"": ""project-for-fc-log""
    ///         },
    ///         ""enable"": true
    ///     }
    ///   
    /// ",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             fooRolePolicyAttachment,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// MNS topic trigger:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "fctriggermnstopic";
    ///     var currentRegion = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var current = AliCloud.GetAccount.Invoke();
    /// 
    ///     var fooProject = new AliCloud.Log.Project("fooProject", new()
    ///     {
    ///         Description = "tf unit test",
    ///     });
    /// 
    ///     var bar = new AliCloud.Log.Store("bar", new()
    ///     {
    ///         Project = fooProject.Name,
    ///         RetentionPeriod = 3000,
    ///         ShardCount = 1,
    ///     });
    /// 
    ///     var fooStore = new AliCloud.Log.Store("fooStore", new()
    ///     {
    ///         Project = fooProject.Name,
    ///         RetentionPeriod = 3000,
    ///         ShardCount = 1,
    ///     });
    /// 
    ///     var fooTopic = new AliCloud.Mns.Topic("fooTopic");
    /// 
    ///     var fooService = new AliCloud.FC.Service("fooService", new()
    ///     {
    ///         InternetAccess = false,
    ///     });
    /// 
    ///     var fooBucket = new AliCloud.Oss.Bucket("fooBucket", new()
    ///     {
    ///         BucketName = name,
    ///     });
    /// 
    ///     // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
    ///     var fooBucketObject = new AliCloud.Oss.BucketObject("fooBucketObject", new()
    ///     {
    ///         Bucket = fooBucket.Id,
    ///         Key = "fc/hello.zip",
    ///         Source = "./hello.zip",
    ///     });
    /// 
    ///     var fooFunction = new AliCloud.FC.Function("fooFunction", new()
    ///     {
    ///         Handler = "hello.handler",
    ///         MemorySize = 512,
    ///         OssBucket = fooBucket.Id,
    ///         OssKey = fooBucketObject.Key,
    ///         Runtime = "python2.7",
    ///         Service = fooService.Name,
    ///     });
    /// 
    ///     var fooRole = new AliCloud.Ram.Role("fooRole", new()
    ///     {
    ///         Description = "this is a test",
    ///         Document = @"  {
    ///     ""Statement"": [
    ///       {
    ///         ""Action"": ""sts:AssumeRole"",
    ///         ""Effect"": ""Allow"",
    ///         ""Principal"": {
    ///           ""Service"": [
    ///             ""mns.aliyuncs.com""
    ///           ]
    ///         }
    ///       }
    ///     ],
    ///     ""Version"": ""1""
    ///   }
    ///   
    /// ",
    ///         Force = true,
    ///     });
    /// 
    ///     var fooRolePolicyAttachment = new AliCloud.Ram.RolePolicyAttachment("fooRolePolicyAttachment", new()
    ///     {
    ///         PolicyName = "AliyunMNSNotificationRolePolicy",
    ///         PolicyType = "System",
    ///         RoleName = fooRole.Name,
    ///     });
    /// 
    ///     var fooTrigger = new AliCloud.FC.Trigger("fooTrigger", new()
    ///     {
    ///         ConfigMns = @"  {
    ///     ""filterTag"":""testTag"",
    ///     ""notifyContentFormat"":""STREAM"",
    ///     ""notifyStrategy"":""BACKOFF_RETRY""
    ///   }
    ///   
    /// ",
    ///         Function = fooFunction.Name,
    ///         Role = fooRole.Arn,
    ///         Service = fooService.Name,
    ///         SourceArn = Output.Tuple(currentRegion, current, fooTopic.Name).Apply(values =&gt;
    ///         {
    ///             var currentRegion = values.Item1;
    ///             var current = values.Item2;
    ///             var name = values.Item3;
    ///             return $"acs:mns:{currentRegion.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id)}:{current.Apply(getAccountResult =&gt; getAccountResult.Id)}:/topics/{name}";
    ///         }),
    ///         Type = "mns_topic",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             "alicloud_ram_role_policy_attachment.foo",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// CDN events trigger:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "fctriggercdneventsconfig";
    ///     var current = AliCloud.GetAccount.Invoke();
    /// 
    ///     var domain = new AliCloud.Cdn.DomainNew("domain", new()
    ///     {
    ///         CdnType = "web",
    ///         DomainName = $"{name}.tf.com",
    ///         Scope = "overseas",
    ///         Sources = new[]
    ///         {
    ///             new AliCloud.Cdn.Inputs.DomainNewSourceArgs
    ///             {
    ///                 Content = "1.1.1.1",
    ///                 Port = 80,
    ///                 Priority = 20,
    ///                 Type = "ipaddr",
    ///                 Weight = 10,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var fooService = new AliCloud.FC.Service("fooService", new()
    ///     {
    ///         InternetAccess = false,
    ///     });
    /// 
    ///     var fooBucket = new AliCloud.Oss.Bucket("fooBucket", new()
    ///     {
    ///         BucketName = name,
    ///     });
    /// 
    ///     // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
    ///     var fooBucketObject = new AliCloud.Oss.BucketObject("fooBucketObject", new()
    ///     {
    ///         Bucket = fooBucket.Id,
    ///         Key = "fc/hello.zip",
    ///         Source = "./hello.zip",
    ///     });
    /// 
    ///     var fooFunction = new AliCloud.FC.Function("fooFunction", new()
    ///     {
    ///         Handler = "hello.handler",
    ///         MemorySize = 512,
    ///         OssBucket = fooBucket.Id,
    ///         OssKey = fooBucketObject.Key,
    ///         Runtime = "python2.7",
    ///         Service = fooService.Name,
    ///     });
    /// 
    ///     var fooRole = new AliCloud.Ram.Role("fooRole", new()
    ///     {
    ///         Description = "this is a test",
    ///         Document = @"    {
    ///         ""Version"": ""1"",
    ///         ""Statement"": [
    ///             {
    ///                 ""Action"": ""cdn:Describe*"",
    ///                 ""Resource"": ""*"",
    ///                 ""Effect"": ""Allow"",
    /// 		        ""Principal"": {
    ///                 ""Service"":
    ///                     [""log.aliyuncs.com""]
    ///                 }
    ///             }
    ///         ]
    ///     }
    ///     
    /// ",
    ///         Force = true,
    ///     });
    /// 
    ///     var fooPolicy = new AliCloud.Ram.Policy("fooPolicy", new()
    ///     {
    ///         Description = "this is a test",
    ///         Document = @"    {
    ///         ""Version"": ""1"",
    ///         ""Statement"": [
    ///         {
    ///             ""Action"": [
    ///             ""fc:InvokeFunction""
    ///             ],
    ///         ""Resource"": [
    ///             ""acs:fc:*:*:services/tf_cdnEvents/functions/*"",
    ///             ""acs:fc:*:*:services/tf_cdnEvents.*/functions/*""
    ///         ],
    ///         ""Effect"": ""Allow""
    ///         }
    ///         ]
    ///     }
    ///     
    /// ",
    ///         Force = true,
    ///     });
    /// 
    ///     var fooRolePolicyAttachment = new AliCloud.Ram.RolePolicyAttachment("fooRolePolicyAttachment", new()
    ///     {
    ///         PolicyName = fooPolicy.Name,
    ///         PolicyType = "Custom",
    ///         RoleName = fooRole.Name,
    ///     });
    /// 
    ///     var @default = new AliCloud.FC.Trigger("default", new()
    ///     {
    ///         Config = domain.DomainName.Apply(domainName =&gt; @$"      {{""eventName"":""LogFileCreated"",
    ///      ""eventVersion"":""1.0.0"",
    ///      ""notes"":""cdn events trigger"",
    ///      ""filter"":{{
    ///         ""domain"": [""{domainName}""]
    ///         }}
    ///     }}
    /// 
    /// "),
    ///         Function = fooFunction.Name,
    ///         Role = fooRole.Arn,
    ///         Service = fooService.Name,
    ///         SourceArn = $"acs:cdn:*:{current.Apply(getAccountResult =&gt; getAccountResult.Id)}",
    ///         Type = "cdn_events",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             "alicloud_ram_role_policy_attachment.foo",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// EventBridge trigger:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "fctriggereventbridgeconfig";
    ///     var current = AliCloud.GetAccount.Invoke();
    /// 
    ///     // Please make eventbridge available and then assume a specific service-linked role, which refers to https://registry.terraform.io/providers/aliyun/alicloud/latest/docs/resources/event_bridge_service_linked_role
    ///     var serviceLinkedRole = new AliCloud.EventBridge.ServiceLinkedRole("serviceLinkedRole", new()
    ///     {
    ///         ProductName = "AliyunServiceRoleForEventBridgeSendToFC",
    ///     });
    /// 
    ///     var fooService = new AliCloud.FC.Service("fooService", new()
    ///     {
    ///         InternetAccess = false,
    ///     });
    /// 
    ///     var fooBucket = new AliCloud.Oss.Bucket("fooBucket", new()
    ///     {
    ///         BucketName = name,
    ///     });
    /// 
    ///     // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
    ///     var fooBucketObject = new AliCloud.Oss.BucketObject("fooBucketObject", new()
    ///     {
    ///         Bucket = fooBucket.Id,
    ///         Key = "fc/hello.zip",
    ///         Source = "./hello.zip",
    ///     });
    /// 
    ///     var fooFunction = new AliCloud.FC.Function("fooFunction", new()
    ///     {
    ///         Handler = "hello.handler",
    ///         MemorySize = 512,
    ///         OssBucket = fooBucket.Id,
    ///         OssKey = fooBucketObject.Key,
    ///         Runtime = "python2.7",
    ///         Service = fooService.Name,
    ///     });
    /// 
    ///     var @default = new AliCloud.FC.Trigger("default", new()
    ///     {
    ///         Config = @"    {
    ///         ""triggerEnable"": false,
    ///         ""asyncInvocationType"": false,
    ///         ""eventRuleFilterPattern"": ""{\""source\"":[\""acs.oss\""],\""type\"":[\""oss:BucketCreated:PutBucket\""]}"",
    ///         ""eventSourceConfig"": {
    ///             ""eventSourceType"": ""Default""
    ///         }
    ///     }
    /// 
    /// ",
    ///         Function = fooFunction.Name,
    ///         Service = fooService.Name,
    ///         Type = "eventbridge",
    ///     });
    /// 
    ///     var mns = new AliCloud.FC.Trigger("mns", new()
    ///     {
    ///         Config = @"    {
    ///         ""triggerEnable"": false,
    ///         ""asyncInvocationType"": false,
    ///         ""eventRuleFilterPattern"": ""{}"",
    ///         ""eventSourceConfig"": {
    ///             ""eventSourceType"": ""MNS"",
    ///             ""eventSourceParameters"": {
    ///                 ""sourceMNSParameters"": {
    ///                     ""RegionId"": ""cn-hangzhou"",
    ///                     ""QueueName"": ""mns-queue"",
    ///                     ""IsBase64Decode"": true
    ///                 }
    ///             }
    ///         }
    ///     }
    /// 
    /// ",
    ///         Function = fooFunction.Name,
    ///         Service = fooService.Name,
    ///         Type = "eventbridge",
    ///     });
    /// 
    ///     var rocketmq = new AliCloud.FC.Trigger("rocketmq", new()
    ///     {
    ///         Config = @"    {
    ///         ""triggerEnable"": false,
    ///         ""asyncInvocationType"": false,
    ///         ""eventRuleFilterPattern"": ""{}"",
    ///         ""eventSourceConfig"": {
    ///             ""eventSourceType"": ""RocketMQ"",
    ///             ""eventSourceParameters"": {
    ///                 ""sourceRocketMQParameters"": {
    ///                     ""RegionId"": ""cn-hangzhou"",
    ///                     ""InstanceId"": ""MQ_INST_164901546557****_BAAN****"",
    ///                     ""GroupID"": ""GID_group1"",
    ///                     ""Topic"": ""mytopic"",
    ///                     ""Timestamp"": 1636597951984,
    ///                     ""Tag"": ""test-tag"",
    ///                     ""Offset"": ""CONSUME_FROM_LAST_OFFSET""
    ///                 }
    ///             }
    ///         }
    ///     }
    /// 
    /// ",
    ///         Function = fooFunction.Name,
    ///         Service = fooService.Name,
    ///         Type = "eventbridge",
    ///     });
    /// 
    ///     var rabbitmq = new AliCloud.FC.Trigger("rabbitmq", new()
    ///     {
    ///         Config = @"    {
    ///         ""triggerEnable"": false,
    ///         ""asyncInvocationType"": false,
    ///         ""eventRuleFilterPattern"": ""{}"",
    ///         ""eventSourceConfig"": {
    ///             ""eventSourceType"": ""RabbitMQ"",
    ///             ""eventSourceParameters"": {
    ///                 ""sourceRabbitMQParameters"": {
    ///                     ""RegionId"": ""cn-hangzhou"",
    ///                     ""InstanceId"": ""amqp-cn-****** "",
    ///                     ""VirtualHostName"": ""test-virtual"",
    ///                     ""QueueName"": ""test-queue""
    ///                 }
    ///             }
    ///         }
    ///     }
    /// 
    /// ",
    ///         Function = fooFunction.Name,
    ///         Service = fooService.Name,
    ///         Type = "eventbridge",
    ///     });
    /// 
    /// });
    /// ```
    /// ## Module Support
    /// 
    /// You can use to the existing fc module
    /// to create several triggers quickly.
    /// 
    /// ## Import
    /// 
    /// Function Compute trigger can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:fc/trigger:Trigger foo my-fc-service:hello-world:hello-trigger
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:fc/trigger:Trigger")]
    public partial class Trigger : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The config of Function Compute trigger.It is valid when `type` is not "mns_topic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
        /// </summary>
        [Output("config")]
        public Output<string?> Config { get; private set; } = null!;

        /// <summary>
        /// The config of Function Compute trigger when the type is "mns_topic".It is conflict with `config`.
        /// </summary>
        [Output("configMns")]
        public Output<string?> ConfigMns { get; private set; } = null!;

        /// <summary>
        /// The Function Compute function name.
        /// </summary>
        [Output("function")]
        public Output<string> Function { get; private set; } = null!;

        /// <summary>
        /// The date this resource was last modified.
        /// </summary>
        [Output("lastModified")]
        public Output<string> LastModified { get; private set; } = null!;

        /// <summary>
        /// The Function Compute trigger name. It is the only in one service and is conflict with "name_prefix".
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Setting a prefix to get a only trigger name. It is conflict with "name".
        /// </summary>
        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;

        /// <summary>
        /// The Function Compute service name.
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;

        /// <summary>
        /// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        /// </summary>
        [Output("sourceArn")]
        public Output<string?> SourceArn { get; private set; } = null!;

        /// <summary>
        /// The Function Compute trigger ID.
        /// </summary>
        [Output("triggerId")]
        public Output<string> TriggerId { get; private set; } = null!;

        /// <summary>
        /// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mns_topic", "cdn_events", "eventbridge"].
        /// 
        /// &gt; **NOTE:** Config does not support modification when type is mns_topic.
        /// &gt; **NOTE:** type = cdn_events, available in 1.47.0+.
        /// &gt; **NOTE:** type = eventbridge, available in 1.173.0+.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Trigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Trigger(string name, TriggerArgs args, CustomResourceOptions? options = null)
            : base("alicloud:fc/trigger:Trigger", name, args ?? new TriggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Trigger(string name, Input<string> id, TriggerState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:fc/trigger:Trigger", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Trigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Trigger Get(string name, Input<string> id, TriggerState? state = null, CustomResourceOptions? options = null)
        {
            return new Trigger(name, id, state, options);
        }
    }

    public sealed class TriggerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The config of Function Compute trigger.It is valid when `type` is not "mns_topic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// The config of Function Compute trigger when the type is "mns_topic".It is conflict with `config`.
        /// </summary>
        [Input("configMns")]
        public Input<string>? ConfigMns { get; set; }

        /// <summary>
        /// The Function Compute function name.
        /// </summary>
        [Input("function", required: true)]
        public Input<string> Function { get; set; } = null!;

        /// <summary>
        /// The Function Compute trigger name. It is the only in one service and is conflict with "name_prefix".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Setting a prefix to get a only trigger name. It is conflict with "name".
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The Function Compute service name.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        /// <summary>
        /// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        /// </summary>
        [Input("sourceArn")]
        public Input<string>? SourceArn { get; set; }

        /// <summary>
        /// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mns_topic", "cdn_events", "eventbridge"].
        /// 
        /// &gt; **NOTE:** Config does not support modification when type is mns_topic.
        /// &gt; **NOTE:** type = cdn_events, available in 1.47.0+.
        /// &gt; **NOTE:** type = eventbridge, available in 1.173.0+.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public TriggerArgs()
        {
        }
        public static new TriggerArgs Empty => new TriggerArgs();
    }

    public sealed class TriggerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The config of Function Compute trigger.It is valid when `type` is not "mns_topic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// The config of Function Compute trigger when the type is "mns_topic".It is conflict with `config`.
        /// </summary>
        [Input("configMns")]
        public Input<string>? ConfigMns { get; set; }

        /// <summary>
        /// The Function Compute function name.
        /// </summary>
        [Input("function")]
        public Input<string>? Function { get; set; }

        /// <summary>
        /// The date this resource was last modified.
        /// </summary>
        [Input("lastModified")]
        public Input<string>? LastModified { get; set; }

        /// <summary>
        /// The Function Compute trigger name. It is the only in one service and is conflict with "name_prefix".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Setting a prefix to get a only trigger name. It is conflict with "name".
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The Function Compute service name.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        /// </summary>
        [Input("sourceArn")]
        public Input<string>? SourceArn { get; set; }

        /// <summary>
        /// The Function Compute trigger ID.
        /// </summary>
        [Input("triggerId")]
        public Input<string>? TriggerId { get; set; }

        /// <summary>
        /// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mns_topic", "cdn_events", "eventbridge"].
        /// 
        /// &gt; **NOTE:** Config does not support modification when type is mns_topic.
        /// &gt; **NOTE:** type = cdn_events, available in 1.47.0+.
        /// &gt; **NOTE:** type = eventbridge, available in 1.173.0+.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public TriggerState()
        {
        }
        public static new TriggerState Empty => new TriggerState();
    }
}
