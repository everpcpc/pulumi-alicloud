// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    public static class GetPeerConnections
    {
        /// <summary>
        /// This data source provides the Vpc Peer Connections of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.186.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.Vpc.GetPeerConnections.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Vpc.GetPeerConnections.Invoke(new()
        ///     {
        ///         NameRegex = "^my-PeerConnection",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["vpcPeerConnectionId1"] = ids.Apply(getPeerConnectionsResult =&gt; getPeerConnectionsResult.Connections[0]?.Id),
        ///         ["vpcPeerConnectionId2"] = nameRegex.Apply(getPeerConnectionsResult =&gt; getPeerConnectionsResult.Connections[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPeerConnectionsResult> InvokeAsync(GetPeerConnectionsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPeerConnectionsResult>("alicloud:vpc/getPeerConnections:getPeerConnections", args ?? new GetPeerConnectionsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Vpc Peer Connections of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.186.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.Vpc.GetPeerConnections.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Vpc.GetPeerConnections.Invoke(new()
        ///     {
        ///         NameRegex = "^my-PeerConnection",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["vpcPeerConnectionId1"] = ids.Apply(getPeerConnectionsResult =&gt; getPeerConnectionsResult.Connections[0]?.Id),
        ///         ["vpcPeerConnectionId2"] = nameRegex.Apply(getPeerConnectionsResult =&gt; getPeerConnectionsResult.Connections[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPeerConnectionsResult> Invoke(GetPeerConnectionsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPeerConnectionsResult>("alicloud:vpc/getPeerConnections:getPeerConnections", args ?? new GetPeerConnectionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPeerConnectionsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of PeerConnection IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by PeerConnection name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("peerConnectionName")]
        public string? PeerConnectionName { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The ID of the requester VPC.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetPeerConnectionsArgs()
        {
        }
        public static new GetPeerConnectionsArgs Empty => new GetPeerConnectionsArgs();
    }

    public sealed class GetPeerConnectionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of PeerConnection IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by PeerConnection name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("peerConnectionName")]
        public Input<string>? PeerConnectionName { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the requester VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetPeerConnectionsInvokeArgs()
        {
        }
        public static new GetPeerConnectionsInvokeArgs Empty => new GetPeerConnectionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetPeerConnectionsResult
    {
        public readonly ImmutableArray<Outputs.GetPeerConnectionsConnectionResult> Connections;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? PeerConnectionName;
        public readonly string? Status;
        public readonly string? VpcId;

        [OutputConstructor]
        private GetPeerConnectionsResult(
            ImmutableArray<Outputs.GetPeerConnectionsConnectionResult> connections,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? peerConnectionName,

            string? status,

            string? vpcId)
        {
            Connections = connections;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PeerConnectionName = peerConnectionName;
            Status = status;
            VpcId = vpcId;
        }
    }
}