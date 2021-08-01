// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Log
{
    /// <summary>
    /// ## Import
    /// 
    /// Log store can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:log/store:Store example tf-log:tf-log-store
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:log/store:Store")]
    public partial class Store : Pulumi.CustomResource
    {
        /// <summary>
        /// Determines whether to append log meta automatically. The meta includes log receive time and client IP address. Default to `true`.
        /// </summary>
        [Output("appendMeta")]
        public Output<bool?> AppendMeta { get; private set; } = null!;

        /// <summary>
        /// Determines whether to automatically split a shard. Default to `false`.
        /// </summary>
        [Output("autoSplit")]
        public Output<bool?> AutoSplit { get; private set; } = null!;

        /// <summary>
        /// Determines whether to enable Web Tracking. Default `false`.
        /// </summary>
        [Output("enableWebTracking")]
        public Output<bool?> EnableWebTracking { get; private set; } = null!;

        /// <summary>
        /// Encrypted storage of data, providing data static protection capability, only supported at creation time.
        /// </summary>
        [Output("encryptConf")]
        public Output<Outputs.StoreEncryptConf?> EncryptConf { get; private set; } = null!;

        /// <summary>
        /// The maximum number of shards for automatic split, which is in the range of 1 to 64. You must specify this parameter when autoSplit is true.
        /// </summary>
        [Output("maxSplitShardCount")]
        public Output<int?> MaxSplitShardCount { get; private set; } = null!;

        /// <summary>
        /// The log store, which is unique in the same project.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project name to the log store belongs.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The data retention time (in days). Valid values: [1-3650]. Default to `30`. Log store data will be stored permanently when the value is `3650`.
        /// </summary>
        [Output("retentionPeriod")]
        public Output<int?> RetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// The number of shards in this log store. Default to 2. You can modify it by "Split" or "Merge" operations. [Refer to details](https://www.alibabacloud.com/help/doc-detail/28976.htm)
        /// </summary>
        [Output("shardCount")]
        public Output<int?> ShardCount { get; private set; } = null!;

        [Output("shards")]
        public Output<ImmutableArray<Outputs.StoreShard>> Shards { get; private set; } = null!;


        /// <summary>
        /// Create a Store resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Store(string name, StoreArgs args, CustomResourceOptions? options = null)
            : base("alicloud:log/store:Store", name, args ?? new StoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Store(string name, Input<string> id, StoreState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:log/store:Store", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Store resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Store Get(string name, Input<string> id, StoreState? state = null, CustomResourceOptions? options = null)
        {
            return new Store(name, id, state, options);
        }
    }

    public sealed class StoreArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines whether to append log meta automatically. The meta includes log receive time and client IP address. Default to `true`.
        /// </summary>
        [Input("appendMeta")]
        public Input<bool>? AppendMeta { get; set; }

        /// <summary>
        /// Determines whether to automatically split a shard. Default to `false`.
        /// </summary>
        [Input("autoSplit")]
        public Input<bool>? AutoSplit { get; set; }

        /// <summary>
        /// Determines whether to enable Web Tracking. Default `false`.
        /// </summary>
        [Input("enableWebTracking")]
        public Input<bool>? EnableWebTracking { get; set; }

        /// <summary>
        /// Encrypted storage of data, providing data static protection capability, only supported at creation time.
        /// </summary>
        [Input("encryptConf")]
        public Input<Inputs.StoreEncryptConfArgs>? EncryptConf { get; set; }

        /// <summary>
        /// The maximum number of shards for automatic split, which is in the range of 1 to 64. You must specify this parameter when autoSplit is true.
        /// </summary>
        [Input("maxSplitShardCount")]
        public Input<int>? MaxSplitShardCount { get; set; }

        /// <summary>
        /// The log store, which is unique in the same project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project name to the log store belongs.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The data retention time (in days). Valid values: [1-3650]. Default to `30`. Log store data will be stored permanently when the value is `3650`.
        /// </summary>
        [Input("retentionPeriod")]
        public Input<int>? RetentionPeriod { get; set; }

        /// <summary>
        /// The number of shards in this log store. Default to 2. You can modify it by "Split" or "Merge" operations. [Refer to details](https://www.alibabacloud.com/help/doc-detail/28976.htm)
        /// </summary>
        [Input("shardCount")]
        public Input<int>? ShardCount { get; set; }

        public StoreArgs()
        {
        }
    }

    public sealed class StoreState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines whether to append log meta automatically. The meta includes log receive time and client IP address. Default to `true`.
        /// </summary>
        [Input("appendMeta")]
        public Input<bool>? AppendMeta { get; set; }

        /// <summary>
        /// Determines whether to automatically split a shard. Default to `false`.
        /// </summary>
        [Input("autoSplit")]
        public Input<bool>? AutoSplit { get; set; }

        /// <summary>
        /// Determines whether to enable Web Tracking. Default `false`.
        /// </summary>
        [Input("enableWebTracking")]
        public Input<bool>? EnableWebTracking { get; set; }

        /// <summary>
        /// Encrypted storage of data, providing data static protection capability, only supported at creation time.
        /// </summary>
        [Input("encryptConf")]
        public Input<Inputs.StoreEncryptConfGetArgs>? EncryptConf { get; set; }

        /// <summary>
        /// The maximum number of shards for automatic split, which is in the range of 1 to 64. You must specify this parameter when autoSplit is true.
        /// </summary>
        [Input("maxSplitShardCount")]
        public Input<int>? MaxSplitShardCount { get; set; }

        /// <summary>
        /// The log store, which is unique in the same project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project name to the log store belongs.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The data retention time (in days). Valid values: [1-3650]. Default to `30`. Log store data will be stored permanently when the value is `3650`.
        /// </summary>
        [Input("retentionPeriod")]
        public Input<int>? RetentionPeriod { get; set; }

        /// <summary>
        /// The number of shards in this log store. Default to 2. You can modify it by "Split" or "Merge" operations. [Refer to details](https://www.alibabacloud.com/help/doc-detail/28976.htm)
        /// </summary>
        [Input("shardCount")]
        public Input<int>? ShardCount { get; set; }

        [Input("shards")]
        private InputList<Inputs.StoreShardGetArgs>? _shards;
        public InputList<Inputs.StoreShardGetArgs> Shards
        {
            get => _shards ?? (_shards = new InputList<Inputs.StoreShardGetArgs>());
            set => _shards = value;
        }

        public StoreState()
        {
        }
    }
}
