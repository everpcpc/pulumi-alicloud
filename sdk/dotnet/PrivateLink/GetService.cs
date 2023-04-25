// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.PrivateLink
{
    public static class GetService
    {
        /// <summary>
        /// Using this data source can open Privatelink service automatically. If the service has been opened, it will return opened.
        /// 
        /// For information about Privatelink and how to use it, see [What is Privatelink](https://www.alibabacloud.com/help/en/product/120462.htm).
        /// 
        /// &gt; **NOTE:** Available in v1.113.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var open = AliCloud.PrivateLink.GetService.Invoke(new()
        ///     {
        ///         Enable = "On",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServiceResult> InvokeAsync(GetServiceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceResult>("alicloud:privatelink/getService:getService", args ?? new GetServiceArgs(), options.WithDefaults());

        /// <summary>
        /// Using this data source can open Privatelink service automatically. If the service has been opened, it will return opened.
        /// 
        /// For information about Privatelink and how to use it, see [What is Privatelink](https://www.alibabacloud.com/help/en/product/120462.htm).
        /// 
        /// &gt; **NOTE:** Available in v1.113.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var open = AliCloud.PrivateLink.GetService.Invoke(new()
        ///     {
        ///         Enable = "On",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServiceResult> Invoke(GetServiceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceResult>("alicloud:privatelink/getService:getService", args ?? new GetServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
        /// </summary>
        [Input("enable")]
        public string? Enable { get; set; }

        public GetServiceArgs()
        {
        }
        public static new GetServiceArgs Empty => new GetServiceArgs();
    }

    public sealed class GetServiceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
        /// </summary>
        [Input("enable")]
        public Input<string>? Enable { get; set; }

        public GetServiceInvokeArgs()
        {
        }
        public static new GetServiceInvokeArgs Empty => new GetServiceInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceResult
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
        private GetServiceResult(
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
