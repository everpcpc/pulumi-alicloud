// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dns
{
    /// <summary>
    /// Provides bind the domain name to the Alidns instance resource.
    /// 
    /// &gt; **NOTE:** Available in v1.99.0+.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var dns = new AliCloud.Dns.AlidnsDomainAttachment("dns", new()
    ///     {
    ///         DomainNames = new[]
    ///         {
    ///             "test111.abc",
    ///             "test222.abc",
    ///         },
    ///         InstanceId = "dns-cn-mp91lyq9xxxx",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DNS domain attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:dns/alidnsDomainAttachment:AlidnsDomainAttachment example dns-cn-v0h1ldjhxxx
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dns/alidnsDomainAttachment:AlidnsDomainAttachment")]
    public partial class AlidnsDomainAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The domain names bound to the DNS instance.
        /// </summary>
        [Output("domainNames")]
        public Output<ImmutableArray<string>> DomainNames { get; private set; } = null!;

        /// <summary>
        /// The id of the DNS instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a AlidnsDomainAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AlidnsDomainAttachment(string name, AlidnsDomainAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dns/alidnsDomainAttachment:AlidnsDomainAttachment", name, args ?? new AlidnsDomainAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AlidnsDomainAttachment(string name, Input<string> id, AlidnsDomainAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dns/alidnsDomainAttachment:AlidnsDomainAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AlidnsDomainAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AlidnsDomainAttachment Get(string name, Input<string> id, AlidnsDomainAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AlidnsDomainAttachment(name, id, state, options);
        }
    }

    public sealed class AlidnsDomainAttachmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("domainNames", required: true)]
        private InputList<string>? _domainNames;

        /// <summary>
        /// The domain names bound to the DNS instance.
        /// </summary>
        public InputList<string> DomainNames
        {
            get => _domainNames ?? (_domainNames = new InputList<string>());
            set => _domainNames = value;
        }

        /// <summary>
        /// The id of the DNS instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public AlidnsDomainAttachmentArgs()
        {
        }
        public static new AlidnsDomainAttachmentArgs Empty => new AlidnsDomainAttachmentArgs();
    }

    public sealed class AlidnsDomainAttachmentState : global::Pulumi.ResourceArgs
    {
        [Input("domainNames")]
        private InputList<string>? _domainNames;

        /// <summary>
        /// The domain names bound to the DNS instance.
        /// </summary>
        public InputList<string> DomainNames
        {
            get => _domainNames ?? (_domainNames = new InputList<string>());
            set => _domainNames = value;
        }

        /// <summary>
        /// The id of the DNS instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public AlidnsDomainAttachmentState()
        {
        }
        public static new AlidnsDomainAttachmentState Empty => new AlidnsDomainAttachmentState();
    }
}
