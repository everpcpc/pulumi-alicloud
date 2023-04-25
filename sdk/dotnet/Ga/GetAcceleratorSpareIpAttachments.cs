// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga
{
    public static class GetAcceleratorSpareIpAttachments
    {
        /// <summary>
        /// This data source provides the Ga Accelerator Spare Ip Attachments of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.167.0+.
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
        ///     var ids = AliCloud.Ga.GetAcceleratorSpareIpAttachments.Invoke(new()
        ///     {
        ///         AcceleratorId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["gaAcceleratorSpareIpAttachmentId1"] = ids.Apply(getAcceleratorSpareIpAttachmentsResult =&gt; getAcceleratorSpareIpAttachmentsResult.Attachments[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAcceleratorSpareIpAttachmentsResult> InvokeAsync(GetAcceleratorSpareIpAttachmentsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAcceleratorSpareIpAttachmentsResult>("alicloud:ga/getAcceleratorSpareIpAttachments:getAcceleratorSpareIpAttachments", args ?? new GetAcceleratorSpareIpAttachmentsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Ga Accelerator Spare Ip Attachments of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.167.0+.
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
        ///     var ids = AliCloud.Ga.GetAcceleratorSpareIpAttachments.Invoke(new()
        ///     {
        ///         AcceleratorId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["gaAcceleratorSpareIpAttachmentId1"] = ids.Apply(getAcceleratorSpareIpAttachmentsResult =&gt; getAcceleratorSpareIpAttachmentsResult.Attachments[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAcceleratorSpareIpAttachmentsResult> Invoke(GetAcceleratorSpareIpAttachmentsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAcceleratorSpareIpAttachmentsResult>("alicloud:ga/getAcceleratorSpareIpAttachments:getAcceleratorSpareIpAttachments", args ?? new GetAcceleratorSpareIpAttachmentsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAcceleratorSpareIpAttachmentsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the global acceleration instance.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public string AcceleratorId { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Accelerator Spare Ip Attachment IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the standby CNAME IP address. Valid values: `active`, `inuse`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetAcceleratorSpareIpAttachmentsArgs()
        {
        }
        public static new GetAcceleratorSpareIpAttachmentsArgs Empty => new GetAcceleratorSpareIpAttachmentsArgs();
    }

    public sealed class GetAcceleratorSpareIpAttachmentsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the global acceleration instance.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public Input<string> AcceleratorId { get; set; } = null!;

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Accelerator Spare Ip Attachment IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of the standby CNAME IP address. Valid values: `active`, `inuse`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetAcceleratorSpareIpAttachmentsInvokeArgs()
        {
        }
        public static new GetAcceleratorSpareIpAttachmentsInvokeArgs Empty => new GetAcceleratorSpareIpAttachmentsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAcceleratorSpareIpAttachmentsResult
    {
        public readonly string AcceleratorId;
        public readonly ImmutableArray<Outputs.GetAcceleratorSpareIpAttachmentsAttachmentResult> Attachments;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly string? Status;

        [OutputConstructor]
        private GetAcceleratorSpareIpAttachmentsResult(
            string acceleratorId,

            ImmutableArray<Outputs.GetAcceleratorSpareIpAttachmentsAttachmentResult> attachments,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            string? status)
        {
            AcceleratorId = acceleratorId;
            Attachments = attachments;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            Status = status;
        }
    }
}
