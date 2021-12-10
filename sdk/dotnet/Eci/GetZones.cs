// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.Eci
{
    public static class GetZones
    {
        /// <summary>
        /// This data source provides the available zones with the Application Load Balancer (ALB) Instance of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.145.0+.
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
        ///         var @default = Output.Create(AliCloud.Eci.GetZones.InvokeAsync());
        ///         this.FirstEciZonesId = @default.Apply(@default =&gt; @default.Zones?[0]?.ZoneIds?[0]);
        ///     }
        /// 
        ///     [Output("firstEciZonesId")]
        ///     public Output&lt;string&gt; FirstEciZonesId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetZonesResult> InvokeAsync(GetZonesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetZonesResult>("alicloud:eci/getZones:getZones", args ?? new GetZonesArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the available zones with the Application Load Balancer (ALB) Instance of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.145.0+.
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
        ///         var @default = Output.Create(AliCloud.Eci.GetZones.InvokeAsync());
        ///         this.FirstEciZonesId = @default.Apply(@default =&gt; @default.Zones?[0]?.ZoneIds?[0]);
        ///     }
        /// 
        ///     [Output("firstEciZonesId")]
        ///     public Output&lt;string&gt; FirstEciZonesId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetZonesResult> Invoke(GetZonesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetZonesResult>("alicloud:eci/getZones:getZones", args ?? new GetZonesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetZonesArgs : Pulumi.InvokeArgs
    {
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetZonesArgs()
        {
        }
    }

    public sealed class GetZonesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetZonesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetZonesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetZonesZoneResult> Zones;

        [OutputConstructor]
        private GetZonesResult(
            string id,

            string? outputFile,

            ImmutableArray<Outputs.GetZonesZoneResult> zones)
        {
            Id = id;
            OutputFile = outputFile;
            Zones = zones;
        }
    }
}