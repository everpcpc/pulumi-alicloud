// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cdn
{
    /// <summary>
    /// Provides a CDN Accelerated Domain resource. This resource is based on CDN's new version OpenAPI.
    /// 
    /// For information about Cdn Domain New and how to use it, see [Add a domain](https://www.alibabacloud.com/help/doc-detail/91176.html).
    /// 
    /// &gt; **NOTE:** Available in v1.34.0+.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cdn_domain_new.html.markdown.
    /// </summary>
    public partial class DomainNew : Pulumi.CustomResource
    {
        /// <summary>
        /// Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
        /// </summary>
        [Output("cdnType")]
        public Output<string> CdnType { get; private set; } = null!;

        /// <summary>
        /// Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
        /// </summary>
        [Output("certificateConfig")]
        public Output<Outputs.DomainNewCertificateConfig> CertificateConfig { get; private set; } = null!;

        /// <summary>
        /// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// Resource group ID.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;

        /// <summary>
        /// The source address list of the accelerated domain. Defaults to null. See Block Sources.
        /// </summary>
        [Output("sources")]
        public Output<Outputs.DomainNewSources> Sources { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DomainNew resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainNew(string name, DomainNewArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cdn/domainNew:DomainNew", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private DomainNew(string name, Input<string> id, DomainNewState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cdn/domainNew:DomainNew", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DomainNew resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainNew Get(string name, Input<string> id, DomainNewState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainNew(name, id, state, options);
        }
    }

    public sealed class DomainNewArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
        /// </summary>
        [Input("cdnType", required: true)]
        public Input<string> CdnType { get; set; } = null!;

        /// <summary>
        /// Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
        /// </summary>
        [Input("certificateConfig")]
        public Input<Inputs.DomainNewCertificateConfigArgs>? CertificateConfig { get; set; }

        /// <summary>
        /// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// Resource group ID.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The source address list of the accelerated domain. Defaults to null. See Block Sources.
        /// </summary>
        [Input("sources", required: true)]
        public Input<Inputs.DomainNewSourcesArgs> Sources { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public DomainNewArgs()
        {
        }
    }

    public sealed class DomainNewState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
        /// </summary>
        [Input("cdnType")]
        public Input<string>? CdnType { get; set; }

        /// <summary>
        /// Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
        /// </summary>
        [Input("certificateConfig")]
        public Input<Inputs.DomainNewCertificateConfigGetArgs>? CertificateConfig { get; set; }

        /// <summary>
        /// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Resource group ID.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The source address list of the accelerated domain. Defaults to null. See Block Sources.
        /// </summary>
        [Input("sources")]
        public Input<Inputs.DomainNewSourcesGetArgs>? Sources { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public DomainNewState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class DomainNewCertificateConfigArgs : Pulumi.ResourceArgs
    {
        [Input("certName")]
        public Input<string>? CertName { get; set; }

        [Input("certType")]
        public Input<string>? CertType { get; set; }

        [Input("forceSet")]
        public Input<string>? ForceSet { get; set; }

        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        [Input("serverCertificate")]
        public Input<string>? ServerCertificate { get; set; }

        [Input("serverCertificateStatus")]
        public Input<string>? ServerCertificateStatus { get; set; }

        public DomainNewCertificateConfigArgs()
        {
        }
    }

    public sealed class DomainNewCertificateConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("certName")]
        public Input<string>? CertName { get; set; }

        [Input("certType")]
        public Input<string>? CertType { get; set; }

        [Input("forceSet")]
        public Input<string>? ForceSet { get; set; }

        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        [Input("serverCertificate")]
        public Input<string>? ServerCertificate { get; set; }

        [Input("serverCertificateStatus")]
        public Input<string>? ServerCertificateStatus { get; set; }

        public DomainNewCertificateConfigGetArgs()
        {
        }
    }

    public sealed class DomainNewSourcesArgs : Pulumi.ResourceArgs
    {
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public DomainNewSourcesArgs()
        {
        }
    }

    public sealed class DomainNewSourcesGetArgs : Pulumi.ResourceArgs
    {
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public DomainNewSourcesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class DomainNewCertificateConfig
    {
        public readonly string? CertName;
        public readonly string? CertType;
        public readonly string? ForceSet;
        public readonly string? PrivateKey;
        public readonly string? ServerCertificate;
        public readonly string? ServerCertificateStatus;

        [OutputConstructor]
        private DomainNewCertificateConfig(
            string? certName,
            string? certType,
            string? forceSet,
            string? privateKey,
            string? serverCertificate,
            string? serverCertificateStatus)
        {
            CertName = certName;
            CertType = certType;
            ForceSet = forceSet;
            PrivateKey = privateKey;
            ServerCertificate = serverCertificate;
            ServerCertificateStatus = serverCertificateStatus;
        }
    }

    [OutputType]
    public sealed class DomainNewSources
    {
        public readonly string Content;
        public readonly int? Port;
        public readonly int? Priority;
        public readonly string Type;
        public readonly int? Weight;

        [OutputConstructor]
        private DomainNewSources(
            string content,
            int? port,
            int? priority,
            string type,
            int? weight)
        {
            Content = content;
            Port = port;
            Priority = priority;
            Type = type;
            Weight = weight;
        }
    }
    }
}
