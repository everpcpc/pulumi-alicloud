// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Pvtz
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// Using `vpc_ids` to attach being in same region several vpc instances to a private zone
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var zone = new AliCloud.Pvtz.Zone("zone");
    /// 
    ///     var first = new AliCloud.Vpc.Network("first", new()
    ///     {
    ///         CidrBlock = "172.16.0.0/12",
    ///     });
    /// 
    ///     var second = new AliCloud.Vpc.Network("second", new()
    ///     {
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var zone_attachment = new AliCloud.Pvtz.ZoneAttachment("zone-attachment", new()
    ///     {
    ///         ZoneId = zone.Id,
    ///         VpcIds = new[]
    ///         {
    ///             first.Id,
    ///             second.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Using `vpcs` to attach being in same region several vpc instances to a private zone
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var zone = new AliCloud.Pvtz.Zone("zone");
    /// 
    ///     var first = new AliCloud.Vpc.Network("first", new()
    ///     {
    ///         CidrBlock = "172.16.0.0/12",
    ///     });
    /// 
    ///     var second = new AliCloud.Vpc.Network("second", new()
    ///     {
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var zone_attachment = new AliCloud.Pvtz.ZoneAttachment("zone-attachment", new()
    ///     {
    ///         ZoneId = zone.Id,
    ///         Vpcs = new[]
    ///         {
    ///             new AliCloud.Pvtz.Inputs.ZoneAttachmentVpcArgs
    ///             {
    ///                 VpcId = first.Id,
    ///             },
    ///             new AliCloud.Pvtz.Inputs.ZoneAttachmentVpcArgs
    ///             {
    ///                 VpcId = second.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Using `vpcs` to attach being in different regions several vpc instances to a private zone
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var zone = new AliCloud.Pvtz.Zone("zone");
    /// 
    ///     var first = new AliCloud.Vpc.Network("first", new()
    ///     {
    ///         CidrBlock = "172.16.0.0/12",
    ///     });
    /// 
    ///     var second = new AliCloud.Vpc.Network("second", new()
    ///     {
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var eu = new AliCloud.Provider("eu", new()
    ///     {
    ///         Region = "eu-central-1",
    ///     });
    /// 
    ///     var third = new AliCloud.Vpc.Network("third", new()
    ///     {
    ///         CidrBlock = "172.16.0.0/16",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Eu,
    ///     });
    /// 
    ///     var zone_attachment = new AliCloud.Pvtz.ZoneAttachment("zone-attachment", new()
    ///     {
    ///         ZoneId = zone.Id,
    ///         Vpcs = new[]
    ///         {
    ///             new AliCloud.Pvtz.Inputs.ZoneAttachmentVpcArgs
    ///             {
    ///                 VpcId = first.Id,
    ///             },
    ///             new AliCloud.Pvtz.Inputs.ZoneAttachmentVpcArgs
    ///             {
    ///                 VpcId = second.Id,
    ///             },
    ///             new AliCloud.Pvtz.Inputs.ZoneAttachmentVpcArgs
    ///             {
    ///                 RegionId = "eu-central-1",
    ///                 VpcId = third.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Private Zone attachment can be imported using the id(same with `zone_id`), e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:pvtz/zoneAttachment:ZoneAttachment example abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:pvtz/zoneAttachment:ZoneAttachment")]
    public partial class ZoneAttachment : global::Pulumi.CustomResource
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
        public Output<ImmutableArray<Outputs.ZoneAttachmentVpc>> Vpcs { get; private set; } = null!;

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
            : base("alicloud:pvtz/zoneAttachment:ZoneAttachment", name, args ?? new ZoneAttachmentArgs(), MakeResourceOptions(options, ""))
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

    public sealed class ZoneAttachmentArgs : global::Pulumi.ResourceArgs
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
        private InputList<Inputs.ZoneAttachmentVpcArgs>? _vpcs;

        /// <summary>
        /// The List of the VPC:
        /// </summary>
        public InputList<Inputs.ZoneAttachmentVpcArgs> Vpcs
        {
            get => _vpcs ?? (_vpcs = new InputList<Inputs.ZoneAttachmentVpcArgs>());
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
        public static new ZoneAttachmentArgs Empty => new ZoneAttachmentArgs();
    }

    public sealed class ZoneAttachmentState : global::Pulumi.ResourceArgs
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
        private InputList<Inputs.ZoneAttachmentVpcGetArgs>? _vpcs;

        /// <summary>
        /// The List of the VPC:
        /// </summary>
        public InputList<Inputs.ZoneAttachmentVpcGetArgs> Vpcs
        {
            get => _vpcs ?? (_vpcs = new InputList<Inputs.ZoneAttachmentVpcGetArgs>());
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
        public static new ZoneAttachmentState Empty => new ZoneAttachmentState();
    }
}
