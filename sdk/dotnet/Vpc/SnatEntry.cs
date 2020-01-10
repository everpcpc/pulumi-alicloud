// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    /// <summary>
    /// Provides a snat resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/snat_entry.html.markdown.
    /// </summary>
    public partial class SnatEntry : Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the snat entry on the server.
        /// </summary>
        [Output("snatEntryId")]
        public Output<string> SnatEntryId { get; private set; } = null!;

        /// <summary>
        /// The SNAT ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidth_packages`.
        /// </summary>
        [Output("snatIp")]
        public Output<string> SnatIp { get; private set; } = null!;

        /// <summary>
        /// The value can get from `alicloud.vpc.NatGateway` Attributes "snat_table_ids".
        /// </summary>
        [Output("snatTableId")]
        public Output<string> SnatTableId { get; private set; } = null!;

        /// <summary>
        /// The vswitch ID.
        /// </summary>
        [Output("sourceVswitchId")]
        public Output<string> SourceVswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a SnatEntry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SnatEntry(string name, SnatEntryArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/snatEntry:SnatEntry", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SnatEntry(string name, Input<string> id, SnatEntryState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/snatEntry:SnatEntry", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SnatEntry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SnatEntry Get(string name, Input<string> id, SnatEntryState? state = null, CustomResourceOptions? options = null)
        {
            return new SnatEntry(name, id, state, options);
        }
    }

    public sealed class SnatEntryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The SNAT ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidth_packages`.
        /// </summary>
        [Input("snatIp", required: true)]
        public Input<string> SnatIp { get; set; } = null!;

        /// <summary>
        /// The value can get from `alicloud.vpc.NatGateway` Attributes "snat_table_ids".
        /// </summary>
        [Input("snatTableId", required: true)]
        public Input<string> SnatTableId { get; set; } = null!;

        /// <summary>
        /// The vswitch ID.
        /// </summary>
        [Input("sourceVswitchId", required: true)]
        public Input<string> SourceVswitchId { get; set; } = null!;

        public SnatEntryArgs()
        {
        }
    }

    public sealed class SnatEntryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the snat entry on the server.
        /// </summary>
        [Input("snatEntryId")]
        public Input<string>? SnatEntryId { get; set; }

        /// <summary>
        /// The SNAT ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidth_packages`.
        /// </summary>
        [Input("snatIp")]
        public Input<string>? SnatIp { get; set; }

        /// <summary>
        /// The value can get from `alicloud.vpc.NatGateway` Attributes "snat_table_ids".
        /// </summary>
        [Input("snatTableId")]
        public Input<string>? SnatTableId { get; set; }

        /// <summary>
        /// The vswitch ID.
        /// </summary>
        [Input("sourceVswitchId")]
        public Input<string>? SourceVswitchId { get; set; }

        public SnatEntryState()
        {
        }
    }
}
