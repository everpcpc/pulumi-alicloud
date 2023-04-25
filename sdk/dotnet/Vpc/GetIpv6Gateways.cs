// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    public static class GetIpv6Gateways
    {
        /// <summary>
        /// This data source provides the Vpc Ipv6 Gateways of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.142.0+.
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
        ///     var ids = AliCloud.Vpc.GetIpv6Gateways.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Vpc.GetIpv6Gateways.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Ipv6Gateway",
        ///     });
        /// 
        ///     var vpcId = AliCloud.Vpc.GetIpv6Gateways.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///         VpcId = "example_value",
        ///     });
        /// 
        ///     var status = AliCloud.Vpc.GetIpv6Gateways.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///         Status = "Available",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["vpcIpv6GatewayId1"] = ids.Apply(getIpv6GatewaysResult =&gt; getIpv6GatewaysResult.Gateways[0]?.Id),
        ///         ["vpcIpv6GatewayId2"] = nameRegex.Apply(getIpv6GatewaysResult =&gt; getIpv6GatewaysResult.Gateways[0]?.Id),
        ///         ["vpcIpv6GatewayId3"] = vpcId.Apply(getIpv6GatewaysResult =&gt; getIpv6GatewaysResult.Gateways[0]?.Id),
        ///         ["vpcIpv6GatewayId4"] = status.Apply(getIpv6GatewaysResult =&gt; getIpv6GatewaysResult.Gateways[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIpv6GatewaysResult> InvokeAsync(GetIpv6GatewaysArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpv6GatewaysResult>("alicloud:vpc/getIpv6Gateways:getIpv6Gateways", args ?? new GetIpv6GatewaysArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Vpc Ipv6 Gateways of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.142.0+.
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
        ///     var ids = AliCloud.Vpc.GetIpv6Gateways.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Vpc.GetIpv6Gateways.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Ipv6Gateway",
        ///     });
        /// 
        ///     var vpcId = AliCloud.Vpc.GetIpv6Gateways.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///         VpcId = "example_value",
        ///     });
        /// 
        ///     var status = AliCloud.Vpc.GetIpv6Gateways.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///         Status = "Available",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["vpcIpv6GatewayId1"] = ids.Apply(getIpv6GatewaysResult =&gt; getIpv6GatewaysResult.Gateways[0]?.Id),
        ///         ["vpcIpv6GatewayId2"] = nameRegex.Apply(getIpv6GatewaysResult =&gt; getIpv6GatewaysResult.Gateways[0]?.Id),
        ///         ["vpcIpv6GatewayId3"] = vpcId.Apply(getIpv6GatewaysResult =&gt; getIpv6GatewaysResult.Gateways[0]?.Id),
        ///         ["vpcIpv6GatewayId4"] = status.Apply(getIpv6GatewaysResult =&gt; getIpv6GatewaysResult.Gateways[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIpv6GatewaysResult> Invoke(GetIpv6GatewaysInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpv6GatewaysResult>("alicloud:vpc/getIpv6Gateways:getIpv6Gateways", args ?? new GetIpv6GatewaysInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpv6GatewaysArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Ipv6 Gateway IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of the IPv6 gateway.
        /// </summary>
        [Input("ipv6GatewayName")]
        public string? Ipv6GatewayName { get; set; }

        /// <summary>
        /// A regex string to filter results by Ipv6 Gateway name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the IPv6 gateway. Valid values: `Available`, `Deleting`, `Pending`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The ID of the virtual private cloud (VPC) to which the IPv6 gateway belongs.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetIpv6GatewaysArgs()
        {
        }
        public static new GetIpv6GatewaysArgs Empty => new GetIpv6GatewaysArgs();
    }

    public sealed class GetIpv6GatewaysInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Ipv6 Gateway IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of the IPv6 gateway.
        /// </summary>
        [Input("ipv6GatewayName")]
        public Input<string>? Ipv6GatewayName { get; set; }

        /// <summary>
        /// A regex string to filter results by Ipv6 Gateway name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of the IPv6 gateway. Valid values: `Available`, `Deleting`, `Pending`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the virtual private cloud (VPC) to which the IPv6 gateway belongs.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetIpv6GatewaysInvokeArgs()
        {
        }
        public static new GetIpv6GatewaysInvokeArgs Empty => new GetIpv6GatewaysInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpv6GatewaysResult
    {
        public readonly ImmutableArray<Outputs.GetIpv6GatewaysGatewayResult> Gateways;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? Ipv6GatewayName;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? Status;
        public readonly string? VpcId;

        [OutputConstructor]
        private GetIpv6GatewaysResult(
            ImmutableArray<Outputs.GetIpv6GatewaysGatewayResult> gateways,

            string id,

            ImmutableArray<string> ids,

            string? ipv6GatewayName,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? status,

            string? vpcId)
        {
            Gateways = gateways;
            Id = id;
            Ids = ids;
            Ipv6GatewayName = ipv6GatewayName;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
            VpcId = vpcId;
        }
    }
}
