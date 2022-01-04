// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.Ga
{
    public static class GetAdditionalCertificates
    {
        /// <summary>
        /// This data source provides the Ga Additional Certificates of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.150.0+.
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
        ///         var ids = Output.Create(AliCloud.Ga.GetAdditionalCertificates.InvokeAsync(new AliCloud.Ga.GetAdditionalCertificatesArgs
        ///         {
        ///             AcceleratorId = "example_value",
        ///             ListenerId = "example_value",
        ///             Ids = 
        ///             {
        ///                 "example_value-1",
        ///                 "example_value-2",
        ///             },
        ///         }));
        ///         this.GaAdditionalCertificateId1 = ids.Apply(ids =&gt; ids.Certificates?[0]?.Id);
        ///     }
        /// 
        ///     [Output("gaAdditionalCertificateId1")]
        ///     public Output&lt;string&gt; GaAdditionalCertificateId1 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAdditionalCertificatesResult> InvokeAsync(GetAdditionalCertificatesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAdditionalCertificatesResult>("alicloud:ga/getAdditionalCertificates:getAdditionalCertificates", args ?? new GetAdditionalCertificatesArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the Ga Additional Certificates of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.150.0+.
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
        ///         var ids = Output.Create(AliCloud.Ga.GetAdditionalCertificates.InvokeAsync(new AliCloud.Ga.GetAdditionalCertificatesArgs
        ///         {
        ///             AcceleratorId = "example_value",
        ///             ListenerId = "example_value",
        ///             Ids = 
        ///             {
        ///                 "example_value-1",
        ///                 "example_value-2",
        ///             },
        ///         }));
        ///         this.GaAdditionalCertificateId1 = ids.Apply(ids =&gt; ids.Certificates?[0]?.Id);
        ///     }
        /// 
        ///     [Output("gaAdditionalCertificateId1")]
        ///     public Output&lt;string&gt; GaAdditionalCertificateId1 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAdditionalCertificatesResult> Invoke(GetAdditionalCertificatesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAdditionalCertificatesResult>("alicloud:ga/getAdditionalCertificates:getAdditionalCertificates", args ?? new GetAdditionalCertificatesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetAdditionalCertificatesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the GA instance.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public string AcceleratorId { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Additional Certificate IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The ID of the listener. Only HTTPS listeners support this parameter.
        /// </summary>
        [Input("listenerId", required: true)]
        public string ListenerId { get; set; } = null!;

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetAdditionalCertificatesArgs()
        {
        }
    }

    public sealed class GetAdditionalCertificatesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the GA instance.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public Input<string> AcceleratorId { get; set; } = null!;

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Additional Certificate IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The ID of the listener. Only HTTPS listeners support this parameter.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetAdditionalCertificatesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAdditionalCertificatesResult
    {
        public readonly string AcceleratorId;
        public readonly ImmutableArray<Outputs.GetAdditionalCertificatesCertificateResult> Certificates;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string ListenerId;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetAdditionalCertificatesResult(
            string acceleratorId,

            ImmutableArray<Outputs.GetAdditionalCertificatesCertificateResult> certificates,

            string id,

            ImmutableArray<string> ids,

            string listenerId,

            string? outputFile)
        {
            AcceleratorId = acceleratorId;
            Certificates = certificates;
            Id = id;
            Ids = ids;
            ListenerId = listenerId;
            OutputFile = outputFile;
        }
    }
}
