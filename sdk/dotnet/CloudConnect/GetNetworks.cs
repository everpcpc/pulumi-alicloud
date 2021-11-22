// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.CloudConnect
{
    public static class GetNetworks
    {
        /// <summary>
        /// This data source provides Cloud Connect Networks available to the user.
        /// 
        /// &gt; **NOTE:** Available in 1.59.0+
        /// 
        /// &gt; **NOTE:** Only the following regions support create Cloud Connect Network. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
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
        ///         var defaultNetworks = Output.Create(AliCloud.CloudConnect.GetNetworks.InvokeAsync(new AliCloud.CloudConnect.GetNetworksArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 alicloud_cloud_connect_networks.Default.Id,
        ///             },
        ///             NameRegex = "^tf-testAcc.*",
        ///         }));
        ///         var defaultNetwork = new AliCloud.CloudConnect.Network("defaultNetwork", new AliCloud.CloudConnect.NetworkArgs
        ///         {
        ///             CidrBlock = "192.168.0.0/24",
        ///             Description = "tf-testAccCloudConnectNetworkDescription",
        ///             IsDefault = true,
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNetworksResult> InvokeAsync(GetNetworksArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworksResult>("alicloud:cloudconnect/getNetworks:getNetworks", args ?? new GetNetworksArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides Cloud Connect Networks available to the user.
        /// 
        /// &gt; **NOTE:** Available in 1.59.0+
        /// 
        /// &gt; **NOTE:** Only the following regions support create Cloud Connect Network. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
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
        ///         var defaultNetworks = Output.Create(AliCloud.CloudConnect.GetNetworks.InvokeAsync(new AliCloud.CloudConnect.GetNetworksArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 alicloud_cloud_connect_networks.Default.Id,
        ///             },
        ///             NameRegex = "^tf-testAcc.*",
        ///         }));
        ///         var defaultNetwork = new AliCloud.CloudConnect.Network("defaultNetwork", new AliCloud.CloudConnect.NetworkArgs
        ///         {
        ///             CidrBlock = "192.168.0.0/24",
        ///             Description = "tf-testAccCloudConnectNetworkDescription",
        ///             IsDefault = true,
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNetworksResult> Invoke(GetNetworksInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNetworksResult>("alicloud:cloudconnect/getNetworks:getNetworks", args ?? new GetNetworksInvokeArgs(), options.WithVersion());
    }


    public sealed class GetNetworksArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of CCN instances IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter CCN instances by name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetNetworksArgs()
        {
        }
    }

    public sealed class GetNetworksInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of CCN instances IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter CCN instances by name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetNetworksInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNetworksResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of CCN instances IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of CCN instances names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        /// <summary>
        /// A list of CCN instances. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNetworksNetworkResult> Networks;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetNetworksResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            ImmutableArray<Outputs.GetNetworksNetworkResult> networks,

            string? outputFile)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            Networks = networks;
            OutputFile = outputFile;
        }
    }
}
