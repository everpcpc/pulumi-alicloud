// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.SimpleApplicationServer
{
    public static class GetServerPlans
    {
        /// <summary>
        /// This data source provides the Simple Application Server Plans of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.135.0+.
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
        ///         var example = Output.Create(AliCloud.SimpleApplicationServer.GetServerPlans.InvokeAsync(new AliCloud.SimpleApplicationServer.GetServerPlansArgs
        ///         {
        ///             Memory = 1,
        ///             Bandwidth = 3,
        ///             DiskSize = 40,
        ///             Flow = 6,
        ///             Core = 2,
        ///         }));
        ///         this.SimpleApplicationServerPlanId1 = data.Alicloud_simple_application_server_plans.Ids.Plans[0].Id;
        ///     }
        /// 
        ///     [Output("simpleApplicationServerPlanId1")]
        ///     public Output&lt;string&gt; SimpleApplicationServerPlanId1 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServerPlansResult> InvokeAsync(GetServerPlansArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerPlansResult>("alicloud:simpleapplicationserver/getServerPlans:getServerPlans", args ?? new GetServerPlansArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the Simple Application Server Plans of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.135.0+.
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
        ///         var example = Output.Create(AliCloud.SimpleApplicationServer.GetServerPlans.InvokeAsync(new AliCloud.SimpleApplicationServer.GetServerPlansArgs
        ///         {
        ///             Memory = 1,
        ///             Bandwidth = 3,
        ///             DiskSize = 40,
        ///             Flow = 6,
        ///             Core = 2,
        ///         }));
        ///         this.SimpleApplicationServerPlanId1 = data.Alicloud_simple_application_server_plans.Ids.Plans[0].Id;
        ///     }
        /// 
        ///     [Output("simpleApplicationServerPlanId1")]
        ///     public Output&lt;string&gt; SimpleApplicationServerPlanId1 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServerPlansResult> Invoke(GetServerPlansInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetServerPlansResult>("alicloud:simpleapplicationserver/getServerPlans:getServerPlans", args ?? new GetServerPlansInvokeArgs(), options.WithVersion());
    }


    public sealed class GetServerPlansArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The peak bandwidth. Unit: Mbit/s.
        /// </summary>
        [Input("bandwidth")]
        public int? Bandwidth { get; set; }

        /// <summary>
        /// The number of CPU cores.
        /// </summary>
        [Input("core")]
        public int? Core { get; set; }

        /// <summary>
        /// The size of the enhanced SSD (ESSD). Unit: GB.
        /// </summary>
        [Input("diskSize")]
        public int? DiskSize { get; set; }

        /// <summary>
        /// The monthly data transfer quota. Unit: GB.
        /// </summary>
        [Input("flow")]
        public int? Flow { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Instance Plan IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The memory size. Unit: GB.
        /// </summary>
        [Input("memory")]
        public int? Memory { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetServerPlansArgs()
        {
        }
    }

    public sealed class GetServerPlansInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The peak bandwidth. Unit: Mbit/s.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The number of CPU cores.
        /// </summary>
        [Input("core")]
        public Input<int>? Core { get; set; }

        /// <summary>
        /// The size of the enhanced SSD (ESSD). Unit: GB.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// The monthly data transfer quota. Unit: GB.
        /// </summary>
        [Input("flow")]
        public Input<int>? Flow { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Instance Plan IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The memory size. Unit: GB.
        /// </summary>
        [Input("memory")]
        public Input<int>? Memory { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetServerPlansInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServerPlansResult
    {
        public readonly int? Bandwidth;
        public readonly int? Core;
        public readonly int? DiskSize;
        public readonly int? Flow;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly int? Memory;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetServerPlansPlanResult> Plans;

        [OutputConstructor]
        private GetServerPlansResult(
            int? bandwidth,

            int? core,

            int? diskSize,

            int? flow,

            string id,

            ImmutableArray<string> ids,

            int? memory,

            string? outputFile,

            ImmutableArray<Outputs.GetServerPlansPlanResult> plans)
        {
            Bandwidth = bandwidth;
            Core = core;
            DiskSize = diskSize;
            Flow = flow;
            Id = id;
            Ids = ids;
            Memory = memory;
            OutputFile = outputFile;
            Plans = plans;
        }
    }
}
