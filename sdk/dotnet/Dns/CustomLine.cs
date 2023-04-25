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
    /// Provides a Alidns Custom Line resource.
    /// 
    /// For information about Alidns Custom Line and how to use it, see [What is Custom Line](https://www.alibabacloud.com/help/en/doc-detail/145059.html).
    /// 
    /// &gt; **NOTE:** Available in v1.151.0+.
    /// 
    /// ## Example Usage
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
    ///     var @default = new AliCloud.Dns.CustomLine("default", new()
    ///     {
    ///         CustomLineName = "tf-testacc",
    ///         DomainName = "your_domain_name",
    ///         IpSegmentLists = new[]
    ///         {
    ///             new AliCloud.Dns.Inputs.CustomLineIpSegmentListArgs
    ///             {
    ///                 EndIp = "192.0.2.125",
    ///                 StartIp = "192.0.2.123",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Alidns Custom Line can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:dns/customLine:CustomLine example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dns/customLine:CustomLine")]
    public partial class CustomLine : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Custom Line.
        /// </summary>
        [Output("customLineName")]
        public Output<string> CustomLineName { get; private set; } = null!;

        /// <summary>
        /// The Domain name.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// The IP segment list. See the following `Block ip_segment_list`.
        /// </summary>
        [Output("ipSegmentLists")]
        public Output<ImmutableArray<Outputs.CustomLineIpSegmentList>> IpSegmentLists { get; private set; } = null!;

        /// <summary>
        /// The lang.
        /// </summary>
        [Output("lang")]
        public Output<string?> Lang { get; private set; } = null!;


        /// <summary>
        /// Create a CustomLine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomLine(string name, CustomLineArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dns/customLine:CustomLine", name, args ?? new CustomLineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomLine(string name, Input<string> id, CustomLineState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dns/customLine:CustomLine", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomLine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomLine Get(string name, Input<string> id, CustomLineState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomLine(name, id, state, options);
        }
    }

    public sealed class CustomLineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Custom Line.
        /// </summary>
        [Input("customLineName", required: true)]
        public Input<string> CustomLineName { get; set; } = null!;

        /// <summary>
        /// The Domain name.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        [Input("ipSegmentLists", required: true)]
        private InputList<Inputs.CustomLineIpSegmentListArgs>? _ipSegmentLists;

        /// <summary>
        /// The IP segment list. See the following `Block ip_segment_list`.
        /// </summary>
        public InputList<Inputs.CustomLineIpSegmentListArgs> IpSegmentLists
        {
            get => _ipSegmentLists ?? (_ipSegmentLists = new InputList<Inputs.CustomLineIpSegmentListArgs>());
            set => _ipSegmentLists = value;
        }

        /// <summary>
        /// The lang.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        public CustomLineArgs()
        {
        }
        public static new CustomLineArgs Empty => new CustomLineArgs();
    }

    public sealed class CustomLineState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Custom Line.
        /// </summary>
        [Input("customLineName")]
        public Input<string>? CustomLineName { get; set; }

        /// <summary>
        /// The Domain name.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("ipSegmentLists")]
        private InputList<Inputs.CustomLineIpSegmentListGetArgs>? _ipSegmentLists;

        /// <summary>
        /// The IP segment list. See the following `Block ip_segment_list`.
        /// </summary>
        public InputList<Inputs.CustomLineIpSegmentListGetArgs> IpSegmentLists
        {
            get => _ipSegmentLists ?? (_ipSegmentLists = new InputList<Inputs.CustomLineIpSegmentListGetArgs>());
            set => _ipSegmentLists = value;
        }

        /// <summary>
        /// The lang.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        public CustomLineState()
        {
        }
        public static new CustomLineState Empty => new CustomLineState();
    }
}
