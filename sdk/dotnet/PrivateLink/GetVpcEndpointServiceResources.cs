// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.PrivateLink
{
    public static class GetVpcEndpointServiceResources
    {
        /// <summary>
        /// This data source provides the Privatelink Vpc Endpoint Service Resources of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.110.0+.
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
        ///         var example = Output.Create(AliCloud.PrivateLink.GetVpcEndpointServiceResources.InvokeAsync(new AliCloud.PrivateLink.GetVpcEndpointServiceResourcesArgs
        ///         {
        ///             ServiceId = "epsrv-gw8ii1xxxx",
        ///         }));
        ///         this.FirstPrivatelinkVpcEndpointServiceResourceId = example.Apply(example =&gt; example.Resources[0].Id);
        ///     }
        /// 
        ///     [Output("firstPrivatelinkVpcEndpointServiceResourceId")]
        ///     public Output&lt;string&gt; FirstPrivatelinkVpcEndpointServiceResourceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVpcEndpointServiceResourcesResult> InvokeAsync(GetVpcEndpointServiceResourcesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcEndpointServiceResourcesResult>("alicloud:privatelink/getVpcEndpointServiceResources:getVpcEndpointServiceResources", args ?? new GetVpcEndpointServiceResourcesArgs(), options.WithVersion());
    }


    public sealed class GetVpcEndpointServiceResourcesArgs : Pulumi.InvokeArgs
    {
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ID of Vpc Endpoint Service.
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetVpcEndpointServiceResourcesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpcEndpointServiceResourcesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetVpcEndpointServiceResourcesResourceResult> Resources;
        public readonly string ServiceId;

        [OutputConstructor]
        private GetVpcEndpointServiceResourcesResult(
            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            ImmutableArray<Outputs.GetVpcEndpointServiceResourcesResourceResult> resources,

            string serviceId)
        {
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            Resources = resources;
            ServiceId = serviceId;
        }
    }
}