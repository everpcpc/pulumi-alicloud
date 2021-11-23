// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cdn
{
    /// <summary>
    /// Provides a CDN Real Time Log Delivery resource.
    /// 
    /// For information about CDN Real Time Log Delivery and how to use it, see [What is Real Time Log Delivery](https://www.alibabacloud.com/help/doc-detail/100456.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.134.0+.
    /// 
    /// ## Example Usage
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
    ///         var example = new AliCloud.Cdn.RealTimeLogDelivery("example", new AliCloud.Cdn.RealTimeLogDeliveryArgs
    ///         {
    ///             Domain = "example_value",
    ///             Logstore = "example_value",
    ///             Project = "example_value",
    ///             SlsRegion = "cn-hanghzou",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CDN Real Time Log Delivery can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cdn/realTimeLogDelivery:RealTimeLogDelivery example &lt;domain&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cdn/realTimeLogDelivery:RealTimeLogDelivery")]
    public partial class RealTimeLogDelivery : Pulumi.CustomResource
    {
        /// <summary>
        /// The accelerated domain name for which you want to configure real-time log delivery. You can specify multiple domain names and separate them with commas (,).
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The name of the Logstore that collects log data from Alibaba Cloud Content Delivery Network (CDN) in real time.
        /// </summary>
        [Output("logstore")]
        public Output<string> Logstore { get; private set; } = null!;

        /// <summary>
        /// The name of the Log Service project that is used for real-time log delivery.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The region where the Log Service project is deployed.
        /// </summary>
        [Output("slsRegion")]
        public Output<string> SlsRegion { get; private set; } = null!;

        /// <summary>
        /// The status of the real-time log delivery feature. Valid Values: `online` and `offline`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a RealTimeLogDelivery resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RealTimeLogDelivery(string name, RealTimeLogDeliveryArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cdn/realTimeLogDelivery:RealTimeLogDelivery", name, args ?? new RealTimeLogDeliveryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RealTimeLogDelivery(string name, Input<string> id, RealTimeLogDeliveryState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cdn/realTimeLogDelivery:RealTimeLogDelivery", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RealTimeLogDelivery resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RealTimeLogDelivery Get(string name, Input<string> id, RealTimeLogDeliveryState? state = null, CustomResourceOptions? options = null)
        {
            return new RealTimeLogDelivery(name, id, state, options);
        }
    }

    public sealed class RealTimeLogDeliveryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The accelerated domain name for which you want to configure real-time log delivery. You can specify multiple domain names and separate them with commas (,).
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// The name of the Logstore that collects log data from Alibaba Cloud Content Delivery Network (CDN) in real time.
        /// </summary>
        [Input("logstore", required: true)]
        public Input<string> Logstore { get; set; } = null!;

        /// <summary>
        /// The name of the Log Service project that is used for real-time log delivery.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The region where the Log Service project is deployed.
        /// </summary>
        [Input("slsRegion", required: true)]
        public Input<string> SlsRegion { get; set; } = null!;

        public RealTimeLogDeliveryArgs()
        {
        }
    }

    public sealed class RealTimeLogDeliveryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The accelerated domain name for which you want to configure real-time log delivery. You can specify multiple domain names and separate them with commas (,).
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The name of the Logstore that collects log data from Alibaba Cloud Content Delivery Network (CDN) in real time.
        /// </summary>
        [Input("logstore")]
        public Input<string>? Logstore { get; set; }

        /// <summary>
        /// The name of the Log Service project that is used for real-time log delivery.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region where the Log Service project is deployed.
        /// </summary>
        [Input("slsRegion")]
        public Input<string>? SlsRegion { get; set; }

        /// <summary>
        /// The status of the real-time log delivery feature. Valid Values: `online` and `offline`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public RealTimeLogDeliveryState()
        {
        }
    }
}