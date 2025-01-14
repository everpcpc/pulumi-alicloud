// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dts
{
    /// <summary>
    /// Provides a DTS Consumer Channel resource.
    /// 
    /// For information about DTS Consumer Channel and how to use it, see [What is Consumer Channel](https://www.alibabacloud.com/help/en/doc-detail/264593.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.146.0+.
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
    ///     var name = config.Get("name") ?? "tftestdts";
    ///     var creation = config.Get("creation") ?? "Rds";
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = creation,
    ///     });
    /// 
    ///     var defaultNetworks = AliCloud.Vpc.GetNetworks.Invoke(new()
    ///     {
    ///         NameRegex = "default-NODELETING",
    ///     });
    /// 
    ///     var defaultSwitches = AliCloud.Vpc.GetSwitches.Invoke(new()
    ///     {
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var instance = new AliCloud.Rds.Instance("instance", new()
    ///     {
    ///         Engine = "MySQL",
    ///         EngineVersion = "5.6",
    ///         InstanceType = "rds.mysql.s1.small",
    ///         InstanceStorage = 10,
    ///         VswitchId = defaultSwitches.Apply(getSwitchesResult =&gt; getSwitchesResult.Ids[0]),
    ///         InstanceName = name,
    ///     });
    /// 
    ///     var db = new List&lt;AliCloud.Rds.Database&gt;();
    ///     for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
    ///     {
    ///         var range = new { Value = rangeIndex };
    ///         db.Add(new AliCloud.Rds.Database($"db-{range.Value}", new()
    ///         {
    ///             InstanceId = instance.Id,
    ///             Description = "from terraform",
    ///         }));
    ///     }
    ///     var account = new AliCloud.Rds.Account("account", new()
    ///     {
    ///         DbInstanceId = instance.Id,
    ///         AccountName = "tftestprivilege",
    ///         AccountPassword = "Test12345",
    ///         AccountDescription = "from terraform",
    ///     });
    /// 
    ///     var privilege = new AliCloud.Rds.AccountPrivilege("privilege", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         AccountName = account.Name,
    ///         Privilege = "ReadWrite",
    ///         DbNames = db.Select(__item =&gt; __item.Name).ToList(),
    ///     });
    /// 
    ///     var defaultSubscriptionJob = new AliCloud.Dts.SubscriptionJob("defaultSubscriptionJob", new()
    ///     {
    ///         DtsJobName = name,
    ///         PaymentType = "PayAsYouGo",
    ///         SourceEndpointEngineName = "MySQL",
    ///         SourceEndpointRegion = "cn-hangzhou",
    ///         SourceEndpointInstanceType = "RDS",
    ///         SourceEndpointInstanceId = instance.Id,
    ///         SourceEndpointDatabaseName = "tfaccountpri_0",
    ///         SourceEndpointUserName = "tftestprivilege",
    ///         SourceEndpointPassword = "Test12345",
    ///         SubscriptionInstanceNetworkType = "vpc",
    ///         DbList = @"        {""dtstestdata"": {""name"": ""tfaccountpri_0"", ""all"": true}}
    /// ",
    ///         SubscriptionInstanceVpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         SubscriptionInstanceVswitchId = defaultSwitches.Apply(getSwitchesResult =&gt; getSwitchesResult.Ids[0]),
    ///         Status = "Normal",
    ///     });
    /// 
    ///     var defaultConsumerChannel = new AliCloud.Dts.ConsumerChannel("defaultConsumerChannel", new()
    ///     {
    ///         DtsInstanceId = defaultSubscriptionJob.DtsInstanceId,
    ///         ConsumerGroupName = name,
    ///         ConsumerGroupUserName = name,
    ///         ConsumerGroupPassword = "tftestAcc123",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DTS Consumer Channel can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:dts/consumerChannel:ConsumerChannel example &lt;dts_instance_id&gt;:&lt;consumer_group_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dts/consumerChannel:ConsumerChannel")]
    public partial class ConsumerChannel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the consumer group.
        /// </summary>
        [Output("consumerGroupId")]
        public Output<string> ConsumerGroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the consumer group.
        /// </summary>
        [Output("consumerGroupName")]
        public Output<string> ConsumerGroupName { get; private set; } = null!;

        /// <summary>
        /// The password of the consumer group account. The length of the `consumer_group_password` is limited to `8` to `32` characters. It can contain two or more of the following characters: uppercase letters, lowercase letters, digits, and special characters.
        /// </summary>
        [Output("consumerGroupPassword")]
        public Output<string> ConsumerGroupPassword { get; private set; } = null!;

        /// <summary>
        /// The username of the consumer group. The length of the `consumer_group_user_name` is limited to `1` to `16` characters. It can contain one or more of the following characters: uppercase letters, lowercase letters, digits, and underscores (_).
        /// </summary>
        [Output("consumerGroupUserName")]
        public Output<string> ConsumerGroupUserName { get; private set; } = null!;

        /// <summary>
        /// The ID of the subscription instance.
        /// </summary>
        [Output("dtsInstanceId")]
        public Output<string> DtsInstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a ConsumerChannel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConsumerChannel(string name, ConsumerChannelArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dts/consumerChannel:ConsumerChannel", name, args ?? new ConsumerChannelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConsumerChannel(string name, Input<string> id, ConsumerChannelState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dts/consumerChannel:ConsumerChannel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConsumerChannel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConsumerChannel Get(string name, Input<string> id, ConsumerChannelState? state = null, CustomResourceOptions? options = null)
        {
            return new ConsumerChannel(name, id, state, options);
        }
    }

    public sealed class ConsumerChannelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the consumer group.
        /// </summary>
        [Input("consumerGroupName", required: true)]
        public Input<string> ConsumerGroupName { get; set; } = null!;

        /// <summary>
        /// The password of the consumer group account. The length of the `consumer_group_password` is limited to `8` to `32` characters. It can contain two or more of the following characters: uppercase letters, lowercase letters, digits, and special characters.
        /// </summary>
        [Input("consumerGroupPassword", required: true)]
        public Input<string> ConsumerGroupPassword { get; set; } = null!;

        /// <summary>
        /// The username of the consumer group. The length of the `consumer_group_user_name` is limited to `1` to `16` characters. It can contain one or more of the following characters: uppercase letters, lowercase letters, digits, and underscores (_).
        /// </summary>
        [Input("consumerGroupUserName", required: true)]
        public Input<string> ConsumerGroupUserName { get; set; } = null!;

        /// <summary>
        /// The ID of the subscription instance.
        /// </summary>
        [Input("dtsInstanceId", required: true)]
        public Input<string> DtsInstanceId { get; set; } = null!;

        public ConsumerChannelArgs()
        {
        }
        public static new ConsumerChannelArgs Empty => new ConsumerChannelArgs();
    }

    public sealed class ConsumerChannelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the consumer group.
        /// </summary>
        [Input("consumerGroupId")]
        public Input<string>? ConsumerGroupId { get; set; }

        /// <summary>
        /// The name of the consumer group.
        /// </summary>
        [Input("consumerGroupName")]
        public Input<string>? ConsumerGroupName { get; set; }

        /// <summary>
        /// The password of the consumer group account. The length of the `consumer_group_password` is limited to `8` to `32` characters. It can contain two or more of the following characters: uppercase letters, lowercase letters, digits, and special characters.
        /// </summary>
        [Input("consumerGroupPassword")]
        public Input<string>? ConsumerGroupPassword { get; set; }

        /// <summary>
        /// The username of the consumer group. The length of the `consumer_group_user_name` is limited to `1` to `16` characters. It can contain one or more of the following characters: uppercase letters, lowercase letters, digits, and underscores (_).
        /// </summary>
        [Input("consumerGroupUserName")]
        public Input<string>? ConsumerGroupUserName { get; set; }

        /// <summary>
        /// The ID of the subscription instance.
        /// </summary>
        [Input("dtsInstanceId")]
        public Input<string>? DtsInstanceId { get; set; }

        public ConsumerChannelState()
        {
        }
        public static new ConsumerChannelState Empty => new ConsumerChannelState();
    }
}
