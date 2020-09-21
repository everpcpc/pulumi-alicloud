// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.RocketMQ
{
    public static class GetInstances
    {
        /// <summary>
        /// This data source provides a list of ONS Instances in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.52.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var config = new Config();
        ///         var name = config.Get("name") ?? "onsInstanceDatasourceName";
        ///         var @default = new AliCloud.RocketMQ.Instance("default", new AliCloud.RocketMQ.InstanceArgs
        ///         {
        ///             Remark = "default_ons_instance_remark",
        ///         });
        ///         var instancesDs = Output.Tuple(@default.Id, @default.Name).Apply(values =&gt;
        ///         {
        ///             var id = values.Item1;
        ///             var name = values.Item2;
        ///             return AliCloud.RocketMQ.GetInstances.InvokeAsync(new AliCloud.RocketMQ.GetInstancesArgs
        ///             {
        ///                 Ids = 
        ///                 {
        ///                     id,
        ///                 },
        ///                 NameRegex = name,
        ///                 OutputFile = "instances.txt",
        ///             });
        ///         });
        ///         this.FirstInstanceId = instancesDs.Apply(instancesDs =&gt; instancesDs.Instances[0].InstanceId);
        ///     }
        /// 
        ///     [Output("firstInstanceId")]
        ///     public Output&lt;string&gt; FirstInstanceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("alicloud:rocketmq/getInstances:getInstances", args ?? new GetInstancesArgs(), options.WithVersion());
    }


    public sealed class GetInstancesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to true can output more details.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of instance IDs to filter results.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by the instance name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of Ons instance. Valid values: `0` deploying, `2` arrears, `5` running, `7` upgrading.
        /// </summary>
        [Input("status")]
        public int? Status { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A map of tags assigned to the Ons instance.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetInstancesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstancesResult
    {
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of instance IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of instances. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstancesInstanceResult> Instances;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of instance names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// The status of the instance. Read [Fields in InstanceVO](https://www.alibabacloud.com/help/doc-detail/106351.html) for further details.
        /// </summary>
        public readonly int? Status;
        /// <summary>
        /// A map of tags assigned to the Ons instance.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetInstancesResult(
            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetInstancesInstanceResult> instances,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            int? status,

            ImmutableDictionary<string, object>? tags)
        {
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            Instances = instances;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
            Tags = tags;
        }
    }
}
