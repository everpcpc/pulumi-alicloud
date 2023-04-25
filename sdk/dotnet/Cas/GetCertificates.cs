// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cas
{
    [Obsolete(@"This resource has been deprecated in favour of getServiceCertificates")]
    public static class GetCertificates
    {
        /// <summary>
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
        ///     var certs = AliCloud.Cas.GetCertificates.Invoke(new()
        ///     {
        ///         NameRegex = "^cas",
        ///         OutputFile = $"{path.Module}/cas_certificates.json",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cert"] = certs.Apply(getCertificatesResult =&gt; getCertificatesResult.Certificates[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCertificatesResult> InvokeAsync(GetCertificatesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCertificatesResult>("alicloud:cas/getCertificates:getCertificates", args ?? new GetCertificatesArgs(), options.WithDefaults());

        /// <summary>
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
        ///     var certs = AliCloud.Cas.GetCertificates.Invoke(new()
        ///     {
        ///         NameRegex = "^cas",
        ///         OutputFile = $"{path.Module}/cas_certificates.json",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cert"] = certs.Apply(getCertificatesResult =&gt; getCertificatesResult.Certificates[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCertificatesResult> Invoke(GetCertificatesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCertificatesResult>("alicloud:cas/getCertificates:getCertificates", args ?? new GetCertificatesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificatesArgs : global::Pulumi.InvokeArgs
    {
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of cert IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("lang")]
        public string? Lang { get; set; }

        /// <summary>
        /// A regex string to filter results by the certificate name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetCertificatesArgs()
        {
        }
        public static new GetCertificatesArgs Empty => new GetCertificatesArgs();
    }

    public sealed class GetCertificatesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of cert IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// A regex string to filter results by the certificate name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetCertificatesInvokeArgs()
        {
        }
        public static new GetCertificatesInvokeArgs Empty => new GetCertificatesInvokeArgs();
    }


    [OutputType]
    public sealed class GetCertificatesResult
    {
        /// <summary>
        /// A list of apis. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCertificatesCertificateResult> Certificates;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of cert IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? Lang;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of cert names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetCertificatesResult(
            ImmutableArray<Outputs.GetCertificatesCertificateResult> certificates,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string? lang,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile)
        {
            Certificates = certificates;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            Lang = lang;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
        }
    }
}
