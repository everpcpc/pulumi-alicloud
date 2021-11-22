// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.ResourceManager
{
    public static class GetHandshakes
    {
        /// <summary>
        /// This data source provides the Resource Manager Handshakes of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:**  Available in 1.86.0+.
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
        ///         var example = Output.Create(AliCloud.ResourceManager.GetHandshakes.InvokeAsync());
        ///         this.FirstHandshakeId = example.Apply(example =&gt; example.Handshakes?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstHandshakeId")]
        ///     public Output&lt;string&gt; FirstHandshakeId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetHandshakesResult> InvokeAsync(GetHandshakesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHandshakesResult>("alicloud:resourcemanager/getHandshakes:getHandshakes", args ?? new GetHandshakesArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the Resource Manager Handshakes of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:**  Available in 1.86.0+.
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
        ///         var example = Output.Create(AliCloud.ResourceManager.GetHandshakes.InvokeAsync());
        ///         this.FirstHandshakeId = example.Apply(example =&gt; example.Handshakes?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstHandshakeId")]
        ///     public Output&lt;string&gt; FirstHandshakeId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetHandshakesResult> Invoke(GetHandshakesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetHandshakesResult>("alicloud:resourcemanager/getHandshakes:getHandshakes", args ?? new GetHandshakesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetHandshakesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// -(Optional, Available in v1.114.0+) Default to `false`. Set it to true can output more details.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Resource Manager Handshake IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of handshake, valid values: `Accepted`, `Cancelled`, `Declined`, `Deleted`, `Expired` and `Pending`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetHandshakesArgs()
        {
        }
    }

    public sealed class GetHandshakesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// -(Optional, Available in v1.114.0+) Default to `false`. Set it to true can output more details.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Resource Manager Handshake IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of handshake, valid values: `Accepted`, `Cancelled`, `Declined`, `Deleted`, `Expired` and `Pending`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetHandshakesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetHandshakesResult
    {
        public readonly bool? EnableDetails;
        /// <summary>
        /// A list of Resource Manager Handshakes. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetHandshakesHandshakeResult> Handshakes;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Resource Manager Handshake IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        /// <summary>
        /// The status of the invitation.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private GetHandshakesResult(
            bool? enableDetails,

            ImmutableArray<Outputs.GetHandshakesHandshakeResult> handshakes,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            string? status)
        {
            EnableDetails = enableDetails;
            Handshakes = handshakes;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            Status = status;
        }
    }
}
