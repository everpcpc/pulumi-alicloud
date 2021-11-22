// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.Cdn
{
    public static class GetRealTimeLogDeliveries
    {
        /// <summary>
        /// This data source provides the Cdn Real Time Log Deliveries of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.134.0+.
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
        ///         var example = Output.Create(AliCloud.Cdn.GetRealTimeLogDeliveries.InvokeAsync(new AliCloud.Cdn.GetRealTimeLogDeliveriesArgs
        ///         {
        ///             Domain = "example_value",
        ///         }));
        ///         this.CdnRealTimeLogDelivery1 = example.Apply(example =&gt; example.Deliveries?[0]?.Id);
        ///     }
        /// 
        ///     [Output("cdnRealTimeLogDelivery1")]
        ///     public Output&lt;string&gt; CdnRealTimeLogDelivery1 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRealTimeLogDeliveriesResult> InvokeAsync(GetRealTimeLogDeliveriesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRealTimeLogDeliveriesResult>("alicloud:cdn/getRealTimeLogDeliveries:getRealTimeLogDeliveries", args ?? new GetRealTimeLogDeliveriesArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the Cdn Real Time Log Deliveries of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.134.0+.
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
        ///         var example = Output.Create(AliCloud.Cdn.GetRealTimeLogDeliveries.InvokeAsync(new AliCloud.Cdn.GetRealTimeLogDeliveriesArgs
        ///         {
        ///             Domain = "example_value",
        ///         }));
        ///         this.CdnRealTimeLogDelivery1 = example.Apply(example =&gt; example.Deliveries?[0]?.Id);
        ///     }
        /// 
        ///     [Output("cdnRealTimeLogDelivery1")]
        ///     public Output&lt;string&gt; CdnRealTimeLogDelivery1 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRealTimeLogDeliveriesResult> Invoke(GetRealTimeLogDeliveriesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRealTimeLogDeliveriesResult>("alicloud:cdn/getRealTimeLogDeliveries:getRealTimeLogDeliveries", args ?? new GetRealTimeLogDeliveriesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetRealTimeLogDeliveriesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Real-Time Log Service Domain.
        /// </summary>
        [Input("domain", required: true)]
        public string Domain { get; set; } = null!;

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// -The status of the real-time log delivery feature. Valid Values: `online` and `offline`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetRealTimeLogDeliveriesArgs()
        {
        }
    }

    public sealed class GetRealTimeLogDeliveriesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Real-Time Log Service Domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// -The status of the real-time log delivery feature. Valid Values: `online` and `offline`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetRealTimeLogDeliveriesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRealTimeLogDeliveriesResult
    {
        public readonly ImmutableArray<Outputs.GetRealTimeLogDeliveriesDeliveryResult> Deliveries;
        public readonly string Domain;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;
        public readonly string? Status;

        [OutputConstructor]
        private GetRealTimeLogDeliveriesResult(
            ImmutableArray<Outputs.GetRealTimeLogDeliveriesDeliveryResult> deliveries,

            string domain,

            string id,

            string? outputFile,

            string? status)
        {
            Deliveries = deliveries;
            Domain = domain;
            Id = id;
            OutputFile = outputFile;
            Status = status;
        }
    }
}
