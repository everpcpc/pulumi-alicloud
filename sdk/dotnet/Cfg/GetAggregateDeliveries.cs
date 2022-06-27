// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cfg
{
    public static class GetAggregateDeliveries
    {
        /// <summary>
        /// This data source provides the Config Aggregate Deliveries of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.172.0+.
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
        ///         var ids = Output.Create(AliCloud.Cfg.GetAggregateDeliveries.InvokeAsync(new AliCloud.Cfg.GetAggregateDeliveriesArgs
        ///         {
        ///             AggregatorId = "example_value",
        ///             Ids = 
        ///             {
        ///                 "example_value-1",
        ///                 "example_value-2",
        ///             },
        ///         }));
        ///         this.ConfigAggregateDeliveryId1 = ids.Apply(ids =&gt; ids.Deliveries?[0]?.Id);
        ///     }
        /// 
        ///     [Output("configAggregateDeliveryId1")]
        ///     public Output&lt;string&gt; ConfigAggregateDeliveryId1 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAggregateDeliveriesResult> InvokeAsync(GetAggregateDeliveriesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAggregateDeliveriesResult>("alicloud:cfg/getAggregateDeliveries:getAggregateDeliveries", args ?? new GetAggregateDeliveriesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Config Aggregate Deliveries of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.172.0+.
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
        ///         var ids = Output.Create(AliCloud.Cfg.GetAggregateDeliveries.InvokeAsync(new AliCloud.Cfg.GetAggregateDeliveriesArgs
        ///         {
        ///             AggregatorId = "example_value",
        ///             Ids = 
        ///             {
        ///                 "example_value-1",
        ///                 "example_value-2",
        ///             },
        ///         }));
        ///         this.ConfigAggregateDeliveryId1 = ids.Apply(ids =&gt; ids.Deliveries?[0]?.Id);
        ///     }
        /// 
        ///     [Output("configAggregateDeliveryId1")]
        ///     public Output&lt;string&gt; ConfigAggregateDeliveryId1 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAggregateDeliveriesResult> Invoke(GetAggregateDeliveriesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAggregateDeliveriesResult>("alicloud:cfg/getAggregateDeliveries:getAggregateDeliveries", args ?? new GetAggregateDeliveriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAggregateDeliveriesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Aggregator.
        /// </summary>
        [Input("aggregatorId", required: true)]
        public string AggregatorId { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Aggregate Delivery IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled.
        /// </summary>
        [Input("status")]
        public int? Status { get; set; }

        public GetAggregateDeliveriesArgs()
        {
        }
    }

    public sealed class GetAggregateDeliveriesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Aggregator.
        /// </summary>
        [Input("aggregatorId", required: true)]
        public Input<string> AggregatorId { get; set; } = null!;

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Aggregate Delivery IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        public GetAggregateDeliveriesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAggregateDeliveriesResult
    {
        public readonly string AggregatorId;
        public readonly ImmutableArray<Outputs.GetAggregateDeliveriesDeliveryResult> Deliveries;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly int? Status;

        [OutputConstructor]
        private GetAggregateDeliveriesResult(
            string aggregatorId,

            ImmutableArray<Outputs.GetAggregateDeliveriesDeliveryResult> deliveries,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            int? status)
        {
            AggregatorId = aggregatorId;
            Deliveries = deliveries;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
        }
    }
}
