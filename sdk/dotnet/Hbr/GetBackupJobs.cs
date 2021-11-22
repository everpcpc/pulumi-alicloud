// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.Hbr
{
    public static class GetBackupJobs
    {
        /// <summary>
        /// This data source provides the Hbr Backup Jobs of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.138.0+.
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
        ///         var defaultEcsBackupPlans = Output.Create(AliCloud.Hbr.GetEcsBackupPlans.InvokeAsync(new AliCloud.Hbr.GetEcsBackupPlansArgs
        ///         {
        ///             NameRegex = "plan-name",
        ///         }));
        ///         var defaultBackupJobs = Output.Tuple(defaultEcsBackupPlans, defaultEcsBackupPlans).Apply(values =&gt;
        ///         {
        ///             var defaultEcsBackupPlans = values.Item1;
        ///             var defaultEcsBackupPlans1 = values.Item2;
        ///             return Output.Create(AliCloud.Hbr.GetBackupJobs.InvokeAsync(new AliCloud.Hbr.GetBackupJobsArgs
        ///             {
        ///                 SourceType = "ECS_FILE",
        ///                 Filters = 
        ///                 {
        ///                     new AliCloud.Hbr.Inputs.GetBackupJobsFilterArgs
        ///                     {
        ///                         Key = "VaultId",
        ///                         Operator = "IN",
        ///                         Values = 
        ///                         {
        ///                             defaultEcsBackupPlans.Plans?[0]?.VaultId,
        ///                         },
        ///                     },
        ///                     new AliCloud.Hbr.Inputs.GetBackupJobsFilterArgs
        ///                     {
        ///                         Key = "InstanceId",
        ///                         Operator = "IN",
        ///                         Values = 
        ///                         {
        ///                             defaultEcsBackupPlans1.Plans?[0]?.InstanceId,
        ///                         },
        ///                     },
        ///                     new AliCloud.Hbr.Inputs.GetBackupJobsFilterArgs
        ///                     {
        ///                         Key = "CompleteTime",
        ///                         Operator = "BETWEEN",
        ///                         Values = 
        ///                         {
        ///                             "2021-08-23T14:17:15CST",
        ///                             "2021-08-24T14:17:15CST",
        ///                         },
        ///                     },
        ///                 },
        ///             }));
        ///         });
        ///         var example = Output.Tuple(defaultEcsBackupPlans, defaultEcsBackupPlans).Apply(values =&gt;
        ///         {
        ///             var defaultEcsBackupPlans = values.Item1;
        ///             var defaultEcsBackupPlans1 = values.Item2;
        ///             return Output.Create(AliCloud.Hbr.GetBackupJobs.InvokeAsync(new AliCloud.Hbr.GetBackupJobsArgs
        ///             {
        ///                 SourceType = "ECS_FILE",
        ///                 Status = "COMPLETE",
        ///                 Filters = 
        ///                 {
        ///                     new AliCloud.Hbr.Inputs.GetBackupJobsFilterArgs
        ///                     {
        ///                         Key = "VaultId",
        ///                         Operator = "IN",
        ///                         Values = 
        ///                         {
        ///                             defaultEcsBackupPlans.Plans?[0]?.VaultId,
        ///                         },
        ///                     },
        ///                     new AliCloud.Hbr.Inputs.GetBackupJobsFilterArgs
        ///                     {
        ///                         Key = "InstanceId",
        ///                         Operator = "IN",
        ///                         Values = 
        ///                         {
        ///                             defaultEcsBackupPlans1.Plans?[0]?.InstanceId,
        ///                         },
        ///                     },
        ///                     new AliCloud.Hbr.Inputs.GetBackupJobsFilterArgs
        ///                     {
        ///                         Key = "CompleteTime",
        ///                         Operator = "LESS_THAN",
        ///                         Values = 
        ///                         {
        ///                             "2021-10-20T20:20:20CST",
        ///                         },
        ///                     },
        ///                 },
        ///             }));
        ///         });
        ///         this.AlicloudHbrBackupJobsDefault1 = defaultBackupJobs.Apply(defaultBackupJobs =&gt; defaultBackupJobs.Jobs?[0]?.Id);
        ///         this.AlicloudHbrBackupJobsExample1 = example.Apply(example =&gt; example.Jobs?[0]?.Id);
        ///     }
        /// 
        ///     [Output("alicloudHbrBackupJobsDefault1")]
        ///     public Output&lt;string&gt; AlicloudHbrBackupJobsDefault1 { get; set; }
        ///     [Output("alicloudHbrBackupJobsExample1")]
        ///     public Output&lt;string&gt; AlicloudHbrBackupJobsExample1 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBackupJobsResult> InvokeAsync(GetBackupJobsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackupJobsResult>("alicloud:hbr/getBackupJobs:getBackupJobs", args ?? new GetBackupJobsArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the Hbr Backup Jobs of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.138.0+.
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
        ///         var defaultEcsBackupPlans = Output.Create(AliCloud.Hbr.GetEcsBackupPlans.InvokeAsync(new AliCloud.Hbr.GetEcsBackupPlansArgs
        ///         {
        ///             NameRegex = "plan-name",
        ///         }));
        ///         var defaultBackupJobs = Output.Tuple(defaultEcsBackupPlans, defaultEcsBackupPlans).Apply(values =&gt;
        ///         {
        ///             var defaultEcsBackupPlans = values.Item1;
        ///             var defaultEcsBackupPlans1 = values.Item2;
        ///             return Output.Create(AliCloud.Hbr.GetBackupJobs.InvokeAsync(new AliCloud.Hbr.GetBackupJobsArgs
        ///             {
        ///                 SourceType = "ECS_FILE",
        ///                 Filters = 
        ///                 {
        ///                     new AliCloud.Hbr.Inputs.GetBackupJobsFilterArgs
        ///                     {
        ///                         Key = "VaultId",
        ///                         Operator = "IN",
        ///                         Values = 
        ///                         {
        ///                             defaultEcsBackupPlans.Plans?[0]?.VaultId,
        ///                         },
        ///                     },
        ///                     new AliCloud.Hbr.Inputs.GetBackupJobsFilterArgs
        ///                     {
        ///                         Key = "InstanceId",
        ///                         Operator = "IN",
        ///                         Values = 
        ///                         {
        ///                             defaultEcsBackupPlans1.Plans?[0]?.InstanceId,
        ///                         },
        ///                     },
        ///                     new AliCloud.Hbr.Inputs.GetBackupJobsFilterArgs
        ///                     {
        ///                         Key = "CompleteTime",
        ///                         Operator = "BETWEEN",
        ///                         Values = 
        ///                         {
        ///                             "2021-08-23T14:17:15CST",
        ///                             "2021-08-24T14:17:15CST",
        ///                         },
        ///                     },
        ///                 },
        ///             }));
        ///         });
        ///         var example = Output.Tuple(defaultEcsBackupPlans, defaultEcsBackupPlans).Apply(values =&gt;
        ///         {
        ///             var defaultEcsBackupPlans = values.Item1;
        ///             var defaultEcsBackupPlans1 = values.Item2;
        ///             return Output.Create(AliCloud.Hbr.GetBackupJobs.InvokeAsync(new AliCloud.Hbr.GetBackupJobsArgs
        ///             {
        ///                 SourceType = "ECS_FILE",
        ///                 Status = "COMPLETE",
        ///                 Filters = 
        ///                 {
        ///                     new AliCloud.Hbr.Inputs.GetBackupJobsFilterArgs
        ///                     {
        ///                         Key = "VaultId",
        ///                         Operator = "IN",
        ///                         Values = 
        ///                         {
        ///                             defaultEcsBackupPlans.Plans?[0]?.VaultId,
        ///                         },
        ///                     },
        ///                     new AliCloud.Hbr.Inputs.GetBackupJobsFilterArgs
        ///                     {
        ///                         Key = "InstanceId",
        ///                         Operator = "IN",
        ///                         Values = 
        ///                         {
        ///                             defaultEcsBackupPlans1.Plans?[0]?.InstanceId,
        ///                         },
        ///                     },
        ///                     new AliCloud.Hbr.Inputs.GetBackupJobsFilterArgs
        ///                     {
        ///                         Key = "CompleteTime",
        ///                         Operator = "LESS_THAN",
        ///                         Values = 
        ///                         {
        ///                             "2021-10-20T20:20:20CST",
        ///                         },
        ///                     },
        ///                 },
        ///             }));
        ///         });
        ///         this.AlicloudHbrBackupJobsDefault1 = defaultBackupJobs.Apply(defaultBackupJobs =&gt; defaultBackupJobs.Jobs?[0]?.Id);
        ///         this.AlicloudHbrBackupJobsExample1 = example.Apply(example =&gt; example.Jobs?[0]?.Id);
        ///     }
        /// 
        ///     [Output("alicloudHbrBackupJobsDefault1")]
        ///     public Output&lt;string&gt; AlicloudHbrBackupJobsDefault1 { get; set; }
        ///     [Output("alicloudHbrBackupJobsExample1")]
        ///     public Output&lt;string&gt; AlicloudHbrBackupJobsExample1 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBackupJobsResult> Invoke(GetBackupJobsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBackupJobsResult>("alicloud:hbr/getBackupJobs:getBackupJobs", args ?? new GetBackupJobsInvokeArgs(), options.WithVersion());
    }


    public sealed class GetBackupJobsArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetBackupJobsFilterArgs>? _filters;
        public List<Inputs.GetBackupJobsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetBackupJobsFilterArgs>());
            set => _filters = value;
        }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Backup Job IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The sort direction. Valid values: `ASCEND`, `DESCEND`.
        /// </summary>
        [Input("sortDirection")]
        public string? SortDirection { get; set; }

        /// <summary>
        /// The type of data source. Valid Values: `ECS_FILE`, `OSS`, `NAS`, `UDM_DISK`.
        /// </summary>
        [Input("sourceType", required: true)]
        public string SourceType { get; set; } = null!;

        /// <summary>
        /// The status of restore job. Valid values: `COMPLETE` , `PARTIAL_COMPLETE`, `FAILED`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetBackupJobsArgs()
        {
        }
    }

    public sealed class GetBackupJobsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetBackupJobsFilterInputArgs>? _filters;
        public InputList<Inputs.GetBackupJobsFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetBackupJobsFilterInputArgs>());
            set => _filters = value;
        }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Backup Job IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The sort direction. Valid values: `ASCEND`, `DESCEND`.
        /// </summary>
        [Input("sortDirection")]
        public Input<string>? SortDirection { get; set; }

        /// <summary>
        /// The type of data source. Valid Values: `ECS_FILE`, `OSS`, `NAS`, `UDM_DISK`.
        /// </summary>
        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        /// <summary>
        /// The status of restore job. Valid values: `COMPLETE` , `PARTIAL_COMPLETE`, `FAILED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetBackupJobsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBackupJobsResult
    {
        public readonly ImmutableArray<Outputs.GetBackupJobsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<Outputs.GetBackupJobsJobResult> Jobs;
        public readonly string? OutputFile;
        public readonly string? SortDirection;
        public readonly string SourceType;
        public readonly string? Status;

        [OutputConstructor]
        private GetBackupJobsResult(
            ImmutableArray<Outputs.GetBackupJobsFilterResult> filters,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetBackupJobsJobResult> jobs,

            string? outputFile,

            string? sortDirection,

            string sourceType,

            string? status)
        {
            Filters = filters;
            Id = id;
            Ids = ids;
            Jobs = jobs;
            OutputFile = outputFile;
            SortDirection = sortDirection;
            SourceType = sourceType;
            Status = status;
        }
    }
}
