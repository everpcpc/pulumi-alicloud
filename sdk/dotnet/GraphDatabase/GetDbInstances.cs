// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.GraphDatabase
{
    public static class GetDbInstances
    {
        /// <summary>
        /// This data source provides the Graph Database Db Instances of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.136.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var ids = Output.Create(AliCloud.GraphDatabase.GetDbInstances.InvokeAsync(new AliCloud.GraphDatabase.GetDbInstancesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_id",
        ///             },
        ///         }));
        ///         this.GraphDatabaseDbInstanceId1 = ids.Apply(ids =&gt; ids.Instances?[0]?.Id);
        ///         var status = Output.Create(AliCloud.GraphDatabase.GetDbInstances.InvokeAsync(new AliCloud.GraphDatabase.GetDbInstancesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_id",
        ///             },
        ///             Status = "Running",
        ///         }));
        ///         this.GraphDatabaseDbInstanceId2 = status.Apply(status =&gt; status.Instances?[0]?.Id);
        ///         var description = Output.Create(AliCloud.GraphDatabase.GetDbInstances.InvokeAsync(new AliCloud.GraphDatabase.GetDbInstancesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_id",
        ///             },
        ///             DbInstanceDescription = "example_value",
        ///         }));
        ///         this.GraphDatabaseDbInstanceId3 = description.Apply(description =&gt; description.Instances?[0]?.Id);
        ///     }
        /// 
        ///     [Output("graphDatabaseDbInstanceId1")]
        ///     public Output&lt;string&gt; GraphDatabaseDbInstanceId1 { get; set; }
        ///     [Output("graphDatabaseDbInstanceId2")]
        ///     public Output&lt;string&gt; GraphDatabaseDbInstanceId2 { get; set; }
        ///     [Output("graphDatabaseDbInstanceId3")]
        ///     public Output&lt;string&gt; GraphDatabaseDbInstanceId3 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDbInstancesResult> InvokeAsync(GetDbInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDbInstancesResult>("alicloud:graphdatabase/getDbInstances:getDbInstances", args ?? new GetDbInstancesArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the Graph Database Db Instances of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.136.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var ids = Output.Create(AliCloud.GraphDatabase.GetDbInstances.InvokeAsync(new AliCloud.GraphDatabase.GetDbInstancesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_id",
        ///             },
        ///         }));
        ///         this.GraphDatabaseDbInstanceId1 = ids.Apply(ids =&gt; ids.Instances?[0]?.Id);
        ///         var status = Output.Create(AliCloud.GraphDatabase.GetDbInstances.InvokeAsync(new AliCloud.GraphDatabase.GetDbInstancesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_id",
        ///             },
        ///             Status = "Running",
        ///         }));
        ///         this.GraphDatabaseDbInstanceId2 = status.Apply(status =&gt; status.Instances?[0]?.Id);
        ///         var description = Output.Create(AliCloud.GraphDatabase.GetDbInstances.InvokeAsync(new AliCloud.GraphDatabase.GetDbInstancesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_id",
        ///             },
        ///             DbInstanceDescription = "example_value",
        ///         }));
        ///         this.GraphDatabaseDbInstanceId3 = description.Apply(description =&gt; description.Instances?[0]?.Id);
        ///     }
        /// 
        ///     [Output("graphDatabaseDbInstanceId1")]
        ///     public Output&lt;string&gt; GraphDatabaseDbInstanceId1 { get; set; }
        ///     [Output("graphDatabaseDbInstanceId2")]
        ///     public Output&lt;string&gt; GraphDatabaseDbInstanceId2 { get; set; }
        ///     [Output("graphDatabaseDbInstanceId3")]
        ///     public Output&lt;string&gt; GraphDatabaseDbInstanceId3 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDbInstancesResult> Invoke(GetDbInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDbInstancesResult>("alicloud:graphdatabase/getDbInstances:getDbInstances", args ?? new GetDbInstancesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetDbInstancesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// According to the practical example or notes.
        /// </summary>
        [Input("dbInstanceDescription")]
        public string? DbInstanceDescription { get; set; }

        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Db Instance IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Instance status. Value range: `Creating`, `Running`, `Deleting`, `Rebooting`, `DBInstanceClassChanging`, `NetAddressCreating` and `NetAddressDeleting`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetDbInstancesArgs()
        {
        }
    }

    public sealed class GetDbInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// According to the practical example or notes.
        /// </summary>
        [Input("dbInstanceDescription")]
        public Input<string>? DbInstanceDescription { get; set; }

        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Db Instance IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Instance status. Value range: `Creating`, `Running`, `Deleting`, `Rebooting`, `DBInstanceClassChanging`, `NetAddressCreating` and `NetAddressDeleting`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetDbInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDbInstancesResult
    {
        public readonly string? DbInstanceDescription;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<Outputs.GetDbInstancesInstanceResult> Instances;
        public readonly string? OutputFile;
        public readonly string? Status;

        [OutputConstructor]
        private GetDbInstancesResult(
            string? dbInstanceDescription,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetDbInstancesInstanceResult> instances,

            string? outputFile,

            string? status)
        {
            DbInstanceDescription = dbInstanceDescription;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            Instances = instances;
            OutputFile = outputFile;
            Status = status;
        }
    }
}
