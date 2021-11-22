// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.PolarDB
{
    public static class GetDatabases
    {
        /// <summary>
        /// The `alicloud.polardb.getDatabases` data source provides a collection of PolarDB cluster database available in Alibaba Cloud account.
        /// Filters support regular expression for the database name, searches by clusterId.
        /// 
        /// &gt; **NOTE:** Available in v1.70.0+.
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
        ///         var polardbClustersDs = Output.Create(AliCloud.PolarDB.GetClusters.InvokeAsync(new AliCloud.PolarDB.GetClustersArgs
        ///         {
        ///             DescriptionRegex = "pc-\\w+",
        ///             Status = "Running",
        ///         }));
        ///         var @default = polardbClustersDs.Apply(polardbClustersDs =&gt; Output.Create(AliCloud.PolarDB.GetDatabases.InvokeAsync(new AliCloud.PolarDB.GetDatabasesArgs
        ///         {
        ///             DbClusterId = polardbClustersDs.Clusters?[0]?.Id,
        ///         })));
        ///         this.Database = @default.Apply(@default =&gt; @default.Databases?[0]?.DbName);
        ///     }
        /// 
        ///     [Output("database")]
        ///     public Output&lt;string&gt; Database { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatabasesResult> InvokeAsync(GetDatabasesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabasesResult>("alicloud:polardb/getDatabases:getDatabases", args ?? new GetDatabasesArgs(), options.WithVersion());

        /// <summary>
        /// The `alicloud.polardb.getDatabases` data source provides a collection of PolarDB cluster database available in Alibaba Cloud account.
        /// Filters support regular expression for the database name, searches by clusterId.
        /// 
        /// &gt; **NOTE:** Available in v1.70.0+.
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
        ///         var polardbClustersDs = Output.Create(AliCloud.PolarDB.GetClusters.InvokeAsync(new AliCloud.PolarDB.GetClustersArgs
        ///         {
        ///             DescriptionRegex = "pc-\\w+",
        ///             Status = "Running",
        ///         }));
        ///         var @default = polardbClustersDs.Apply(polardbClustersDs =&gt; Output.Create(AliCloud.PolarDB.GetDatabases.InvokeAsync(new AliCloud.PolarDB.GetDatabasesArgs
        ///         {
        ///             DbClusterId = polardbClustersDs.Clusters?[0]?.Id,
        ///         })));
        ///         this.Database = @default.Apply(@default =&gt; @default.Databases?[0]?.DbName);
        ///     }
        /// 
        ///     [Output("database")]
        ///     public Output&lt;string&gt; Database { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDatabasesResult> Invoke(GetDatabasesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDatabasesResult>("alicloud:polardb/getDatabases:getDatabases", args ?? new GetDatabasesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetDatabasesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The polarDB cluster ID.
        /// </summary>
        [Input("dbClusterId", required: true)]
        public string DbClusterId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by database name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        public GetDatabasesArgs()
        {
        }
    }

    public sealed class GetDatabasesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The polarDB cluster ID.
        /// </summary>
        [Input("dbClusterId", required: true)]
        public Input<string> DbClusterId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by database name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        public GetDatabasesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabasesResult
    {
        /// <summary>
        /// A list of PolarDB cluster databases. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabasesDatabaseResult> Databases;
        public readonly string DbClusterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? NameRegex;
        /// <summary>
        /// database name of the cluster.
        /// </summary>
        public readonly ImmutableArray<string> Names;

        [OutputConstructor]
        private GetDatabasesResult(
            ImmutableArray<Outputs.GetDatabasesDatabaseResult> databases,

            string dbClusterId,

            string id,

            string? nameRegex,

            ImmutableArray<string> names)
        {
            Databases = databases;
            DbClusterId = dbClusterId;
            Id = id;
            NameRegex = nameRegex;
            Names = names;
        }
    }
}
