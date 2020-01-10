// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Pvtz
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/pvtz_zone_attachment.html.markdown.
    /// </summary>
    public partial class ZoneAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The language of code.
        /// </summary>
        [Output("lang")]
        public Output<string?> Lang { get; private set; } = null!;

        /// <summary>
        /// The user custom IP address.
        /// </summary>
        [Output("userClientIp")]
        public Output<string?> UserClientIp { get; private set; } = null!;

        /// <summary>
        /// The id List of the VPC with the same region, for example:["vpc-1","vpc-2"]. 
        /// </summary>
        [Output("vpcIds")]
        public Output<ImmutableArray<string>> VpcIds { get; private set; } = null!;

        /// <summary>
        /// The List of the VPC:
        /// </summary>
        [Output("vpcs")]
        public Output<ImmutableArray<Outputs.ZoneAttachmentVpcs>> Vpcs { get; private set; } = null!;

        /// <summary>
        /// The name of the Private Zone Record.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a ZoneAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ZoneAttachment(string name, ZoneAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:pvtz/zoneAttachment:ZoneAttachment", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ZoneAttachment(string name, Input<string> id, ZoneAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:pvtz/zoneAttachment:ZoneAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ZoneAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ZoneAttachment Get(string name, Input<string> id, ZoneAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ZoneAttachment(name, id, state, options);
        }
    }

    public sealed class ZoneAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The language of code.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// The user custom IP address.
        /// </summary>
        [Input("userClientIp")]
        public Input<string>? UserClientIp { get; set; }

        [Input("vpcIds")]
        private InputList<string>? _vpcIds;

        /// <summary>
        /// The id List of the VPC with the same region, for example:["vpc-1","vpc-2"]. 
        /// </summary>
        public InputList<string> VpcIds
        {
            get => _vpcIds ?? (_vpcIds = new InputList<string>());
            set => _vpcIds = value;
        }

        [Input("vpcs")]
        private InputList<Inputs.ZoneAttachmentVpcsArgs>? _vpcs;

        /// <summary>
        /// The List of the VPC:
        /// </summary>
        public InputList<Inputs.ZoneAttachmentVpcsArgs> Vpcs
        {
            get => _vpcs ?? (_vpcs = new InputList<Inputs.ZoneAttachmentVpcsArgs>());
            set => _vpcs = value;
        }

        /// <summary>
        /// The name of the Private Zone Record.
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public ZoneAttachmentArgs()
        {
        }
    }

    public sealed class ZoneAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The language of code.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// The user custom IP address.
        /// </summary>
        [Input("userClientIp")]
        public Input<string>? UserClientIp { get; set; }

        [Input("vpcIds")]
        private InputList<string>? _vpcIds;

        /// <summary>
        /// The id List of the VPC with the same region, for example:["vpc-1","vpc-2"]. 
        /// </summary>
        public InputList<string> VpcIds
        {
            get => _vpcIds ?? (_vpcIds = new InputList<string>());
            set => _vpcIds = value;
        }

        [Input("vpcs")]
        private InputList<Inputs.ZoneAttachmentVpcsGetArgs>? _vpcs;

        /// <summary>
        /// The List of the VPC:
        /// </summary>
        public InputList<Inputs.ZoneAttachmentVpcsGetArgs> Vpcs
        {
            get => _vpcs ?? (_vpcs = new InputList<Inputs.ZoneAttachmentVpcsGetArgs>());
            set => _vpcs = value;
        }

        /// <summary>
        /// The name of the Private Zone Record.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ZoneAttachmentState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ZoneAttachmentVpcsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The region of the vpc. If not set, the current region will instead of.
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// The Id of the vpc.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public ZoneAttachmentVpcsArgs()
        {
        }
    }

    public sealed class ZoneAttachmentVpcsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The region of the vpc. If not set, the current region will instead of.
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// The Id of the vpc.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public ZoneAttachmentVpcsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ZoneAttachmentVpcs
    {
        /// <summary>
        /// The region of the vpc. If not set, the current region will instead of.
        /// </summary>
        public readonly string RegionId;
        /// <summary>
        /// The Id of the vpc.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private ZoneAttachmentVpcs(
            string regionId,
            string vpcId)
        {
            RegionId = regionId;
            VpcId = vpcId;
        }
    }
    }
}
