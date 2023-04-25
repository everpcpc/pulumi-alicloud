// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.PrivateLink
{
    public static class GetVpcEndpointServices
    {
        /// <summary>
        /// This data source provides the Privatelink Vpc Endpoint Services of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.109.0+.
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
        ///     var example = AliCloud.PrivateLink.GetVpcEndpointServices.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_value",
        ///         },
        ///         NameRegex = "the_resource_name",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstPrivatelinkVpcEndpointServiceId"] = example.Apply(getVpcEndpointServicesResult =&gt; getVpcEndpointServicesResult.Services[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVpcEndpointServicesResult> InvokeAsync(GetVpcEndpointServicesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcEndpointServicesResult>("alicloud:privatelink/getVpcEndpointServices:getVpcEndpointServices", args ?? new GetVpcEndpointServicesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Privatelink Vpc Endpoint Services of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.109.0+.
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
        ///     var example = AliCloud.PrivateLink.GetVpcEndpointServices.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_value",
        ///         },
        ///         NameRegex = "the_resource_name",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstPrivatelinkVpcEndpointServiceId"] = example.Apply(getVpcEndpointServicesResult =&gt; getVpcEndpointServicesResult.Services[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVpcEndpointServicesResult> Invoke(GetVpcEndpointServicesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcEndpointServicesResult>("alicloud:privatelink/getVpcEndpointServices:getVpcEndpointServices", args ?? new GetVpcEndpointServicesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcEndpointServicesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether to automatically accept terminal node connections..
        /// </summary>
        [Input("autoAcceptConnection")]
        public bool? AutoAcceptConnection { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Vpc Endpoint Service IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Vpc Endpoint Service name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The business status of the terminal node service..
        /// </summary>
        [Input("serviceBusinessStatus")]
        public string? ServiceBusinessStatus { get; set; }

        /// <summary>
        /// The Status of Vpc Endpoint Service.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The name of Vpc Endpoint Service.
        /// </summary>
        [Input("vpcEndpointServiceName")]
        public string? VpcEndpointServiceName { get; set; }

        public GetVpcEndpointServicesArgs()
        {
        }
        public static new GetVpcEndpointServicesArgs Empty => new GetVpcEndpointServicesArgs();
    }

    public sealed class GetVpcEndpointServicesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether to automatically accept terminal node connections..
        /// </summary>
        [Input("autoAcceptConnection")]
        public Input<bool>? AutoAcceptConnection { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Vpc Endpoint Service IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Vpc Endpoint Service name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The business status of the terminal node service..
        /// </summary>
        [Input("serviceBusinessStatus")]
        public Input<string>? ServiceBusinessStatus { get; set; }

        /// <summary>
        /// The Status of Vpc Endpoint Service.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The name of Vpc Endpoint Service.
        /// </summary>
        [Input("vpcEndpointServiceName")]
        public Input<string>? VpcEndpointServiceName { get; set; }

        public GetVpcEndpointServicesInvokeArgs()
        {
        }
        public static new GetVpcEndpointServicesInvokeArgs Empty => new GetVpcEndpointServicesInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcEndpointServicesResult
    {
        public readonly bool? AutoAcceptConnection;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? ServiceBusinessStatus;
        public readonly ImmutableArray<Outputs.GetVpcEndpointServicesServiceResult> Services;
        public readonly string? Status;
        public readonly string? VpcEndpointServiceName;

        [OutputConstructor]
        private GetVpcEndpointServicesResult(
            bool? autoAcceptConnection,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? serviceBusinessStatus,

            ImmutableArray<Outputs.GetVpcEndpointServicesServiceResult> services,

            string? status,

            string? vpcEndpointServiceName)
        {
            AutoAcceptConnection = autoAcceptConnection;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ServiceBusinessStatus = serviceBusinessStatus;
            Services = services;
            Status = status;
            VpcEndpointServiceName = vpcEndpointServiceName;
        }
    }
}
