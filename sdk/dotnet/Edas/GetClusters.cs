// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Edas
{
    public static class GetClusters
    {
        /// <summary>
        /// This data source provides a list of EDAS clusters in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.82.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var clusters = AliCloud.Edas.GetClusters.Invoke(new()
        ///     {
        ///         LogicalRegionId = "cn-shenzhen:xxx",
        ///         Ids = new[]
        ///         {
        ///             "addfs-dfsasd",
        ///         },
        ///         OutputFile = "clusters.txt",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstClusterName"] = data.Alicloud_alikafka_consumer_groups.Clusters.Clusters[0].Cluster_name,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClustersResult> InvokeAsync(GetClustersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClustersResult>("alicloud:edas/getClusters:getClusters", args ?? new GetClustersArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides a list of EDAS clusters in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.82.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var clusters = AliCloud.Edas.GetClusters.Invoke(new()
        ///     {
        ///         LogicalRegionId = "cn-shenzhen:xxx",
        ///         Ids = new[]
        ///         {
        ///             "addfs-dfsasd",
        ///         },
        ///         OutputFile = "clusters.txt",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstClusterName"] = data.Alicloud_alikafka_consumer_groups.Clusters.Clusters[0].Cluster_name,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetClustersResult> Invoke(GetClustersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClustersResult>("alicloud:edas/getClusters:getClusters", args ?? new GetClustersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClustersArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// An ids string to filter results by the cluster id.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// ID of the namespace in EDAS.
        /// </summary>
        [Input("logicalRegionId", required: true)]
        public string LogicalRegionId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by the cluster name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetClustersArgs()
        {
        }
        public static new GetClustersArgs Empty => new GetClustersArgs();
    }

    public sealed class GetClustersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// An ids string to filter results by the cluster id.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// ID of the namespace in EDAS.
        /// </summary>
        [Input("logicalRegionId", required: true)]
        public Input<string> LogicalRegionId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by the cluster name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetClustersInvokeArgs()
        {
        }
        public static new GetClustersInvokeArgs Empty => new GetClustersInvokeArgs();
    }


    [OutputType]
    public sealed class GetClustersResult
    {
        /// <summary>
        /// A list of clusters.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClustersClusterResult> Clusters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of cluster IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string LogicalRegionId;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of cluster names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetClustersResult(
            ImmutableArray<Outputs.GetClustersClusterResult> clusters,

            string id,

            ImmutableArray<string> ids,

            string logicalRegionId,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile)
        {
            Clusters = clusters;
            Id = id;
            Ids = ids;
            LogicalRegionId = logicalRegionId;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
        }
    }
}
