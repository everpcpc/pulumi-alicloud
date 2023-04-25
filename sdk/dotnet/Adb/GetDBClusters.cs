// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Adb
{
    public static class GetDBClusters
    {
        /// <summary>
        /// This data source provides the Adb DBClusters of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.121.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var example = AliCloud.Adb.GetDBClusters.Invoke(new()
        ///     {
        ///         DescriptionRegex = "example",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstAdbDbClusterId"] = example.Apply(getDBClustersResult =&gt; getDBClustersResult.Clusters[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDBClustersResult> InvokeAsync(GetDBClustersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDBClustersResult>("alicloud:adb/getDBClusters:getDBClusters", args ?? new GetDBClustersArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Adb DBClusters of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.121.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var example = AliCloud.Adb.GetDBClusters.Invoke(new()
        ///     {
        ///         DescriptionRegex = "example",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstAdbDbClusterId"] = example.Apply(getDBClustersResult =&gt; getDBClustersResult.Clusters[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDBClustersResult> Invoke(GetDBClustersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDBClustersResult>("alicloud:adb/getDBClusters:getDBClusters", args ?? new GetDBClustersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDBClustersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The description of DBCluster.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// A regex string to filter results by DBCluster description.
        /// </summary>
        [Input("descriptionRegex")]
        public string? DescriptionRegex { get; set; }

        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of DBCluster IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetDBClustersArgs()
        {
        }
        public static new GetDBClustersArgs Empty => new GetDBClustersArgs();
    }

    public sealed class GetDBClustersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The description of DBCluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A regex string to filter results by DBCluster description.
        /// </summary>
        [Input("descriptionRegex")]
        public Input<string>? DescriptionRegex { get; set; }

        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of DBCluster IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetDBClustersInvokeArgs()
        {
        }
        public static new GetDBClustersInvokeArgs Empty => new GetDBClustersInvokeArgs();
    }


    [OutputType]
    public sealed class GetDBClustersResult
    {
        public readonly ImmutableArray<Outputs.GetDBClustersClusterResult> Clusters;
        public readonly string? Description;
        public readonly string? DescriptionRegex;
        public readonly ImmutableArray<string> Descriptions;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;
        public readonly string? ResourceGroupId;
        public readonly string? Status;
        public readonly ImmutableDictionary<string, object>? Tags;
        public readonly int TotalCount;

        [OutputConstructor]
        private GetDBClustersResult(
            ImmutableArray<Outputs.GetDBClustersClusterResult> clusters,

            string? description,

            string? descriptionRegex,

            ImmutableArray<string> descriptions,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            int? pageNumber,

            int? pageSize,

            string? resourceGroupId,

            string? status,

            ImmutableDictionary<string, object>? tags,

            int totalCount)
        {
            Clusters = clusters;
            Description = description;
            DescriptionRegex = descriptionRegex;
            Descriptions = descriptions;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
            ResourceGroupId = resourceGroupId;
            Status = status;
            Tags = tags;
            TotalCount = totalCount;
        }
    }
}
