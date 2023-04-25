// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cassandra
{
    public static class GetBackupPlans
    {
        /// <summary>
        /// This data source provides the Cassandra Backup Plans of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.128.0+.
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
        ///     var example = AliCloud.Cassandra.GetBackupPlans.Invoke(new()
        ///     {
        ///         ClusterId = "example_value",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstCassandraBackupPlanId"] = example.Apply(getBackupPlansResult =&gt; getBackupPlansResult.Plans[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBackupPlansResult> InvokeAsync(GetBackupPlansArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBackupPlansResult>("alicloud:cassandra/getBackupPlans:getBackupPlans", args ?? new GetBackupPlansArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Cassandra Backup Plans of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.128.0+.
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
        ///     var example = AliCloud.Cassandra.GetBackupPlans.Invoke(new()
        ///     {
        ///         ClusterId = "example_value",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstCassandraBackupPlanId"] = example.Apply(getBackupPlansResult =&gt; getBackupPlansResult.Plans[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBackupPlansResult> Invoke(GetBackupPlansInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBackupPlansResult>("alicloud:cassandra/getBackupPlans:getBackupPlans", args ?? new GetBackupPlansInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackupPlansArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the cluster for the backup.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetBackupPlansArgs()
        {
        }
        public static new GetBackupPlansArgs Empty => new GetBackupPlansArgs();
    }

    public sealed class GetBackupPlansInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the cluster for the backup.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetBackupPlansInvokeArgs()
        {
        }
        public static new GetBackupPlansInvokeArgs Empty => new GetBackupPlansInvokeArgs();
    }


    [OutputType]
    public sealed class GetBackupPlansResult
    {
        public readonly string ClusterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetBackupPlansPlanResult> Plans;

        [OutputConstructor]
        private GetBackupPlansResult(
            string clusterId,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            ImmutableArray<Outputs.GetBackupPlansPlanResult> plans)
        {
            ClusterId = clusterId;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            Plans = plans;
        }
    }
}
