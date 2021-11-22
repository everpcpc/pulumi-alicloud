// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.Brain
{
    public static class GetIndustrialSerice
    {
        /// <summary>
        /// Using this data source can open Brain Industrial service automatically. If the service has been opened, it will return opened.
        /// 
        /// &gt; **NOTE:** Available in v1.115.0+
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
        ///         var open = Output.Create(AliCloud.Brain.GetIndustrialSerice.InvokeAsync(new AliCloud.Brain.GetIndustrialSericeArgs
        ///         {
        ///             Enable = "On",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIndustrialSericeResult> InvokeAsync(GetIndustrialSericeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIndustrialSericeResult>("alicloud:brain/getIndustrialSerice:getIndustrialSerice", args ?? new GetIndustrialSericeArgs(), options.WithVersion());

        /// <summary>
        /// Using this data source can open Brain Industrial service automatically. If the service has been opened, it will return opened.
        /// 
        /// &gt; **NOTE:** Available in v1.115.0+
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
        ///         var open = Output.Create(AliCloud.Brain.GetIndustrialSerice.InvokeAsync(new AliCloud.Brain.GetIndustrialSericeArgs
        ///         {
        ///             Enable = "On",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIndustrialSericeResult> Invoke(GetIndustrialSericeInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetIndustrialSericeResult>("alicloud:brain/getIndustrialSerice:getIndustrialSerice", args ?? new GetIndustrialSericeInvokeArgs(), options.WithVersion());
    }


    public sealed class GetIndustrialSericeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
        /// </summary>
        [Input("enable")]
        public string? Enable { get; set; }

        public GetIndustrialSericeArgs()
        {
        }
    }

    public sealed class GetIndustrialSericeInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
        /// </summary>
        [Input("enable")]
        public Input<string>? Enable { get; set; }

        public GetIndustrialSericeInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIndustrialSericeResult
    {
        public readonly string? Enable;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The current service enable status.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetIndustrialSericeResult(
            string? enable,

            string id,

            string status)
        {
            Enable = enable;
            Id = id;
            Status = status;
        }
    }
}
