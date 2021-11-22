// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.Slb
{
    public static class GetAttachments
    {
        /// <summary>
        /// This data source provides the server load balancer attachments of the current Alibaba Cloud user.
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
        ///         var sampleDs = Output.Create(AliCloud.Slb.GetAttachments.InvokeAsync(new AliCloud.Slb.GetAttachmentsArgs
        ///         {
        ///             LoadBalancerId = alicloud_slb_load_balancer.Sample_slb.Id,
        ///         }));
        ///         this.FirstSlbAttachmentInstanceId = sampleDs.Apply(sampleDs =&gt; sampleDs.SlbAttachments?[0]?.InstanceId);
        ///     }
        /// 
        ///     [Output("firstSlbAttachmentInstanceId")]
        ///     public Output&lt;string&gt; FirstSlbAttachmentInstanceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAttachmentsResult> InvokeAsync(GetAttachmentsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAttachmentsResult>("alicloud:slb/getAttachments:getAttachments", args ?? new GetAttachmentsArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the server load balancer attachments of the current Alibaba Cloud user.
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
        ///         var sampleDs = Output.Create(AliCloud.Slb.GetAttachments.InvokeAsync(new AliCloud.Slb.GetAttachmentsArgs
        ///         {
        ///             LoadBalancerId = alicloud_slb_load_balancer.Sample_slb.Id,
        ///         }));
        ///         this.FirstSlbAttachmentInstanceId = sampleDs.Apply(sampleDs =&gt; sampleDs.SlbAttachments?[0]?.InstanceId);
        ///     }
        /// 
        ///     [Output("firstSlbAttachmentInstanceId")]
        ///     public Output&lt;string&gt; FirstSlbAttachmentInstanceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAttachmentsResult> Invoke(GetAttachmentsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAttachmentsResult>("alicloud:slb/getAttachments:getAttachments", args ?? new GetAttachmentsInvokeArgs(), options.WithVersion());
    }


    public sealed class GetAttachmentsArgs : Pulumi.InvokeArgs
    {
        [Input("instanceIds")]
        private List<string>? _instanceIds;

        /// <summary>
        /// List of attached ECS instance IDs.
        /// </summary>
        public List<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new List<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// ID of the SLB with attachments.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public string LoadBalancerId { get; set; } = null!;

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetAttachmentsArgs()
        {
        }
    }

    public sealed class GetAttachmentsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("instanceIds")]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// List of attached ECS instance IDs.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// ID of the SLB with attachments.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public Input<string> LoadBalancerId { get; set; } = null!;

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetAttachmentsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAttachmentsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> InstanceIds;
        public readonly string LoadBalancerId;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of SLB attachments. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAttachmentsSlbAttachmentResult> SlbAttachments;

        [OutputConstructor]
        private GetAttachmentsResult(
            string id,

            ImmutableArray<string> instanceIds,

            string loadBalancerId,

            string? outputFile,

            ImmutableArray<Outputs.GetAttachmentsSlbAttachmentResult> slbAttachments)
        {
            Id = id;
            InstanceIds = instanceIds;
            LoadBalancerId = loadBalancerId;
            OutputFile = outputFile;
            SlbAttachments = slbAttachments;
        }
    }
}
