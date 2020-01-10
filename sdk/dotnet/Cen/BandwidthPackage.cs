// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    /// <summary>
    /// Provides a CEN bandwidth package resource. The CEN bandwidth package is an abstracted object that includes an interconnection bandwidth and interconnection areas. To buy a bandwidth package, you must specify the areas to connect. An area consists of one or more Alibaba Cloud regions. The areas in CEN include Mainland China, Asia Pacific, North America, and Europe.
    /// 
    /// For information about CEN and how to use it, see [Manage bandwidth packages](https://www.alibabacloud.com/help/doc-detail/65982.htm).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cen_bandwidth_package.html.markdown.
    /// </summary>
    public partial class BandwidthPackage : Pulumi.CustomResource
    {
        /// <summary>
        /// The bandwidth in Mbps of the bandwidth package. Cannot be less than 2Mbps.
        /// </summary>
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// The billing method. Valid value: PostPaid | PrePaid. Default to PostPaid. If set to PrePaid, the bandwidth package can't be deleted before expired time.
        /// </summary>
        [Output("chargeType")]
        public Output<string?> ChargeType { get; private set; } = null!;

        /// <summary>
        /// The description of the bandwidth package. Default to null.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The time of the bandwidth package to expire.
        /// </summary>
        [Output("expiredTime")]
        public Output<string> ExpiredTime { get; private set; } = null!;

        /// <summary>
        /// List of the two areas to connect. Valid value: China | North-America | Asia-Pacific | Europe | Middle-East.
        /// </summary>
        [Output("geographicRegionIds")]
        public Output<ImmutableArray<string>> GeographicRegionIds { get; private set; } = null!;

        /// <summary>
        /// The name of the bandwidth package. Defaults to null.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The purchase period in month. Valid value: 1, 2, 3, 6, 12. Default to 1.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// The status of the bandwidth, including "InUse" and "Idle".
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a BandwidthPackage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BandwidthPackage(string name, BandwidthPackageArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/bandwidthPackage:BandwidthPackage", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private BandwidthPackage(string name, Input<string> id, BandwidthPackageState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/bandwidthPackage:BandwidthPackage", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BandwidthPackage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BandwidthPackage Get(string name, Input<string> id, BandwidthPackageState? state = null, CustomResourceOptions? options = null)
        {
            return new BandwidthPackage(name, id, state, options);
        }
    }

    public sealed class BandwidthPackageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth in Mbps of the bandwidth package. Cannot be less than 2Mbps.
        /// </summary>
        [Input("bandwidth", required: true)]
        public Input<int> Bandwidth { get; set; } = null!;

        /// <summary>
        /// The billing method. Valid value: PostPaid | PrePaid. Default to PostPaid. If set to PrePaid, the bandwidth package can't be deleted before expired time.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The description of the bandwidth package. Default to null.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("geographicRegionIds", required: true)]
        private InputList<string>? _geographicRegionIds;

        /// <summary>
        /// List of the two areas to connect. Valid value: China | North-America | Asia-Pacific | Europe | Middle-East.
        /// </summary>
        public InputList<string> GeographicRegionIds
        {
            get => _geographicRegionIds ?? (_geographicRegionIds = new InputList<string>());
            set => _geographicRegionIds = value;
        }

        /// <summary>
        /// The name of the bandwidth package. Defaults to null.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The purchase period in month. Valid value: 1, 2, 3, 6, 12. Default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        public BandwidthPackageArgs()
        {
        }
    }

    public sealed class BandwidthPackageState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth in Mbps of the bandwidth package. Cannot be less than 2Mbps.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The billing method. Valid value: PostPaid | PrePaid. Default to PostPaid. If set to PrePaid, the bandwidth package can't be deleted before expired time.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The description of the bandwidth package. Default to null.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The time of the bandwidth package to expire.
        /// </summary>
        [Input("expiredTime")]
        public Input<string>? ExpiredTime { get; set; }

        [Input("geographicRegionIds")]
        private InputList<string>? _geographicRegionIds;

        /// <summary>
        /// List of the two areas to connect. Valid value: China | North-America | Asia-Pacific | Europe | Middle-East.
        /// </summary>
        public InputList<string> GeographicRegionIds
        {
            get => _geographicRegionIds ?? (_geographicRegionIds = new InputList<string>());
            set => _geographicRegionIds = value;
        }

        /// <summary>
        /// The name of the bandwidth package. Defaults to null.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The purchase period in month. Valid value: 1, 2, 3, 6, 12. Default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The status of the bandwidth, including "InUse" and "Idle".
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public BandwidthPackageState()
        {
        }
    }
}
