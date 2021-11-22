// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.Cen
{
    public static class GetRouteServices
    {
        /// <summary>
        /// This data source provides CEN Route Service available to the user.
        /// 
        /// &gt; **NOTE:** Available in v1.102.0+
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
        ///         var example = Output.Create(AliCloud.Cen.GetRouteServices.InvokeAsync(new AliCloud.Cen.GetRouteServicesArgs
        ///         {
        ///             CenId = "cen-7qthudw0ll6jmc****",
        ///         }));
        ///         this.FirstCenRouteServiceId = example.Apply(example =&gt; example.Services?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstCenRouteServiceId")]
        ///     public Output&lt;string&gt; FirstCenRouteServiceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRouteServicesResult> InvokeAsync(GetRouteServicesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRouteServicesResult>("alicloud:cen/getRouteServices:getRouteServices", args ?? new GetRouteServicesArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides CEN Route Service available to the user.
        /// 
        /// &gt; **NOTE:** Available in v1.102.0+
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
        ///         var example = Output.Create(AliCloud.Cen.GetRouteServices.InvokeAsync(new AliCloud.Cen.GetRouteServicesArgs
        ///         {
        ///             CenId = "cen-7qthudw0ll6jmc****",
        ///         }));
        ///         this.FirstCenRouteServiceId = example.Apply(example =&gt; example.Services?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstCenRouteServiceId")]
        ///     public Output&lt;string&gt; FirstCenRouteServiceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRouteServicesResult> Invoke(GetRouteServicesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRouteServicesResult>("alicloud:cen/getRouteServices:getRouteServices", args ?? new GetRouteServicesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetRouteServicesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The region of the network instances that access the cloud services.
        /// </summary>
        [Input("accessRegionId")]
        public string? AccessRegionId { get; set; }

        /// <summary>
        /// -(Required, ForceNew) The ID of the CEN instance.
        /// </summary>
        [Input("cenId", required: true)]
        public string CenId { get; set; } = null!;

        /// <summary>
        /// -(Optional, ForceNew) The domain name or IP address of the cloud service.
        /// </summary>
        [Input("host")]
        public string? Host { get; set; }

        /// <summary>
        /// The region of the cloud service.
        /// </summary>
        [Input("hostRegionId")]
        public string? HostRegionId { get; set; }

        /// <summary>
        /// The VPC associated with the cloud service.
        /// </summary>
        [Input("hostVpcId")]
        public string? HostVpcId { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the cloud service. Valid values: `Active`, `Creating` and `Deleting`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetRouteServicesArgs()
        {
        }
    }

    public sealed class GetRouteServicesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The region of the network instances that access the cloud services.
        /// </summary>
        [Input("accessRegionId")]
        public Input<string>? AccessRegionId { get; set; }

        /// <summary>
        /// -(Required, ForceNew) The ID of the CEN instance.
        /// </summary>
        [Input("cenId", required: true)]
        public Input<string> CenId { get; set; } = null!;

        /// <summary>
        /// -(Optional, ForceNew) The domain name or IP address of the cloud service.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The region of the cloud service.
        /// </summary>
        [Input("hostRegionId")]
        public Input<string>? HostRegionId { get; set; }

        /// <summary>
        /// The VPC associated with the cloud service.
        /// </summary>
        [Input("hostVpcId")]
        public Input<string>? HostVpcId { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of the cloud service. Valid values: `Active`, `Creating` and `Deleting`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetRouteServicesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRouteServicesResult
    {
        /// <summary>
        /// The region of the network instances that access the cloud services.
        /// </summary>
        public readonly string? AccessRegionId;
        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        public readonly string CenId;
        /// <summary>
        /// The domain name or IP address of the cloud service.
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// The region of the cloud service.
        /// </summary>
        public readonly string? HostRegionId;
        /// <summary>
        /// The VPC associated with the cloud service.
        /// </summary>
        public readonly string? HostVpcId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of CEN Route Service IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of CEN Route Services. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouteServicesServiceResult> Services;
        /// <summary>
        /// The status of the cloud service.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private GetRouteServicesResult(
            string? accessRegionId,

            string cenId,

            string? host,

            string? hostRegionId,

            string? hostVpcId,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            ImmutableArray<Outputs.GetRouteServicesServiceResult> services,

            string? status)
        {
            AccessRegionId = accessRegionId;
            CenId = cenId;
            Host = host;
            HostRegionId = hostRegionId;
            HostVpcId = hostVpcId;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            Services = services;
            Status = status;
        }
    }
}
