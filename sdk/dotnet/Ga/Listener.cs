// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga
{
    /// <summary>
    /// Provides a Global Accelerator (GA) Listener resource.
    /// 
    /// For information about Global Accelerator (GA) Listener and how to use it, see [What is Listener](https://help.aliyun.com/document_detail/153253.html).
    /// 
    /// &gt; **NOTE:** Available in v1.111.0+.
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
    ///     var exampleAccelerator = new AliCloud.Ga.Accelerator("exampleAccelerator", new()
    ///     {
    ///         Duration = 1,
    ///         AutoUseCoupon = true,
    ///         Spec = "1",
    ///     });
    /// 
    ///     var deBandwidthPackage = new AliCloud.Ga.BandwidthPackage("deBandwidthPackage", new()
    ///     {
    ///         Bandwidth = 100,
    ///         Type = "Basic",
    ///         BandwidthType = "Basic",
    ///         PaymentType = "PayAsYouGo",
    ///         BillingType = "PayBy95",
    ///         Ratio = 30,
    ///     });
    /// 
    ///     var deBandwidthPackageAttachment = new AliCloud.Ga.BandwidthPackageAttachment("deBandwidthPackageAttachment", new()
    ///     {
    ///         AcceleratorId = exampleAccelerator.Id,
    ///         BandwidthPackageId = deBandwidthPackage.Id,
    ///     });
    /// 
    ///     var exampleListener = new AliCloud.Ga.Listener("exampleListener", new()
    ///     {
    ///         AcceleratorId = exampleAccelerator.Id,
    ///         PortRanges = new[]
    ///         {
    ///             new AliCloud.Ga.Inputs.ListenerPortRangeArgs
    ///             {
    ///                 FromPort = 60,
    ///                 ToPort = 70,
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             deBandwidthPackageAttachment,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Ga Listener can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ga/listener:Listener example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ga/listener:Listener")]
    public partial class Listener : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The accelerator id.
        /// </summary>
        [Output("acceleratorId")]
        public Output<string> AcceleratorId { get; private set; } = null!;

        /// <summary>
        /// The certificates of the listener.
        /// 
        /// &gt; **NOTE:** This parameter needs to be configured only for monitoring of the HTTPS protocol.
        /// </summary>
        [Output("certificates")]
        public Output<ImmutableArray<Outputs.ListenerCertificate>> Certificates { get; private set; } = null!;

        /// <summary>
        /// The clientAffinity of the listener. Default value is `NONE`. Valid values:
        /// `NONE`: client affinity is not maintained, that is, connection requests from the same client cannot always be directed to the same terminal node.
        /// `SOURCE_IP`: maintain client affinity. When a client accesses a stateful application, all requests from the same client can be directed to the same terminal node, regardless of the source port and protocol.
        /// </summary>
        [Output("clientAffinity")]
        public Output<string?> ClientAffinity { get; private set; } = null!;

        /// <summary>
        /// The description of the listener.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The routing type of the listener. Default Value: `Standard`. Valid values:
        /// </summary>
        [Output("listenerType")]
        public Output<string> ListenerType { get; private set; } = null!;

        /// <summary>
        /// The name of the listener. The length of the name is 2-128 characters. It starts with uppercase and lowercase letters or Chinese characters. It can contain numbers and underscores and dashes.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The portRanges of the listener.
        /// 
        /// &gt; **NOTE:** For HTTP or HTTPS protocol monitoring, only one monitoring port can be configured, that is, the start monitoring port and end monitoring port should be the same.
        /// </summary>
        [Output("portRanges")]
        public Output<ImmutableArray<Outputs.ListenerPortRange>> PortRanges { get; private set; } = null!;

        /// <summary>
        /// Type of network transport protocol monitored. Default value is `TCP`. Valid values: `TCP`, `UDP`, `HTTP`, `HTTPS`.
        /// 
        /// &gt; **NOTE:** At present, the white list of HTTP and HTTPS monitoring protocols is open. If you need to use it, please submit a work order.
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        /// <summary>
        /// The proxy protocol of the listener. Default value is `false`. Valid value:
        /// `true`: Turn on the keep client source IP function. After it is turned on, the back-end service is supported to view the original IP address of the client.
        /// `false`: keep client source IP function is not turned on.
        /// </summary>
        [Output("proxyProtocol")]
        public Output<bool?> ProxyProtocol { get; private set; } = null!;

        /// <summary>
        /// The ID of the security policy. **NOTE:** Only HTTPS listeners support this parameter. Valid values:
        /// </summary>
        [Output("securityPolicyId")]
        public Output<string> SecurityPolicyId { get; private set; } = null!;

        /// <summary>
        /// The status of the listener.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Listener resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Listener(string name, ListenerArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ga/listener:Listener", name, args ?? new ListenerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Listener(string name, Input<string> id, ListenerState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ga/listener:Listener", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Listener resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Listener Get(string name, Input<string> id, ListenerState? state = null, CustomResourceOptions? options = null)
        {
            return new Listener(name, id, state, options);
        }
    }

    public sealed class ListenerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The accelerator id.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public Input<string> AcceleratorId { get; set; } = null!;

        [Input("certificates")]
        private InputList<Inputs.ListenerCertificateArgs>? _certificates;

        /// <summary>
        /// The certificates of the listener.
        /// 
        /// &gt; **NOTE:** This parameter needs to be configured only for monitoring of the HTTPS protocol.
        /// </summary>
        public InputList<Inputs.ListenerCertificateArgs> Certificates
        {
            get => _certificates ?? (_certificates = new InputList<Inputs.ListenerCertificateArgs>());
            set => _certificates = value;
        }

        /// <summary>
        /// The clientAffinity of the listener. Default value is `NONE`. Valid values:
        /// `NONE`: client affinity is not maintained, that is, connection requests from the same client cannot always be directed to the same terminal node.
        /// `SOURCE_IP`: maintain client affinity. When a client accesses a stateful application, all requests from the same client can be directed to the same terminal node, regardless of the source port and protocol.
        /// </summary>
        [Input("clientAffinity")]
        public Input<string>? ClientAffinity { get; set; }

        /// <summary>
        /// The description of the listener.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The routing type of the listener. Default Value: `Standard`. Valid values:
        /// </summary>
        [Input("listenerType")]
        public Input<string>? ListenerType { get; set; }

        /// <summary>
        /// The name of the listener. The length of the name is 2-128 characters. It starts with uppercase and lowercase letters or Chinese characters. It can contain numbers and underscores and dashes.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("portRanges", required: true)]
        private InputList<Inputs.ListenerPortRangeArgs>? _portRanges;

        /// <summary>
        /// The portRanges of the listener.
        /// 
        /// &gt; **NOTE:** For HTTP or HTTPS protocol monitoring, only one monitoring port can be configured, that is, the start monitoring port and end monitoring port should be the same.
        /// </summary>
        public InputList<Inputs.ListenerPortRangeArgs> PortRanges
        {
            get => _portRanges ?? (_portRanges = new InputList<Inputs.ListenerPortRangeArgs>());
            set => _portRanges = value;
        }

        /// <summary>
        /// Type of network transport protocol monitored. Default value is `TCP`. Valid values: `TCP`, `UDP`, `HTTP`, `HTTPS`.
        /// 
        /// &gt; **NOTE:** At present, the white list of HTTP and HTTPS monitoring protocols is open. If you need to use it, please submit a work order.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The proxy protocol of the listener. Default value is `false`. Valid value:
        /// `true`: Turn on the keep client source IP function. After it is turned on, the back-end service is supported to view the original IP address of the client.
        /// `false`: keep client source IP function is not turned on.
        /// </summary>
        [Input("proxyProtocol")]
        public Input<bool>? ProxyProtocol { get; set; }

        /// <summary>
        /// The ID of the security policy. **NOTE:** Only HTTPS listeners support this parameter. Valid values:
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        public ListenerArgs()
        {
        }
        public static new ListenerArgs Empty => new ListenerArgs();
    }

    public sealed class ListenerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The accelerator id.
        /// </summary>
        [Input("acceleratorId")]
        public Input<string>? AcceleratorId { get; set; }

        [Input("certificates")]
        private InputList<Inputs.ListenerCertificateGetArgs>? _certificates;

        /// <summary>
        /// The certificates of the listener.
        /// 
        /// &gt; **NOTE:** This parameter needs to be configured only for monitoring of the HTTPS protocol.
        /// </summary>
        public InputList<Inputs.ListenerCertificateGetArgs> Certificates
        {
            get => _certificates ?? (_certificates = new InputList<Inputs.ListenerCertificateGetArgs>());
            set => _certificates = value;
        }

        /// <summary>
        /// The clientAffinity of the listener. Default value is `NONE`. Valid values:
        /// `NONE`: client affinity is not maintained, that is, connection requests from the same client cannot always be directed to the same terminal node.
        /// `SOURCE_IP`: maintain client affinity. When a client accesses a stateful application, all requests from the same client can be directed to the same terminal node, regardless of the source port and protocol.
        /// </summary>
        [Input("clientAffinity")]
        public Input<string>? ClientAffinity { get; set; }

        /// <summary>
        /// The description of the listener.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The routing type of the listener. Default Value: `Standard`. Valid values:
        /// </summary>
        [Input("listenerType")]
        public Input<string>? ListenerType { get; set; }

        /// <summary>
        /// The name of the listener. The length of the name is 2-128 characters. It starts with uppercase and lowercase letters or Chinese characters. It can contain numbers and underscores and dashes.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("portRanges")]
        private InputList<Inputs.ListenerPortRangeGetArgs>? _portRanges;

        /// <summary>
        /// The portRanges of the listener.
        /// 
        /// &gt; **NOTE:** For HTTP or HTTPS protocol monitoring, only one monitoring port can be configured, that is, the start monitoring port and end monitoring port should be the same.
        /// </summary>
        public InputList<Inputs.ListenerPortRangeGetArgs> PortRanges
        {
            get => _portRanges ?? (_portRanges = new InputList<Inputs.ListenerPortRangeGetArgs>());
            set => _portRanges = value;
        }

        /// <summary>
        /// Type of network transport protocol monitored. Default value is `TCP`. Valid values: `TCP`, `UDP`, `HTTP`, `HTTPS`.
        /// 
        /// &gt; **NOTE:** At present, the white list of HTTP and HTTPS monitoring protocols is open. If you need to use it, please submit a work order.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The proxy protocol of the listener. Default value is `false`. Valid value:
        /// `true`: Turn on the keep client source IP function. After it is turned on, the back-end service is supported to view the original IP address of the client.
        /// `false`: keep client source IP function is not turned on.
        /// </summary>
        [Input("proxyProtocol")]
        public Input<bool>? ProxyProtocol { get; set; }

        /// <summary>
        /// The ID of the security policy. **NOTE:** Only HTTPS listeners support this parameter. Valid values:
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        /// <summary>
        /// The status of the listener.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ListenerState()
        {
        }
        public static new ListenerState Empty => new ListenerState();
    }
}
