// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    /// <summary>
    /// Provides a CEN child instance attachment resource.
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
    ///         var config = new Config();
    ///         var name = config.Get("name") ?? "tf-testAccCenInstanceAttachmentBasic";
    ///         var cen = new AliCloud.Cen.Instance("cen", new AliCloud.Cen.InstanceArgs
    ///         {
    ///             Description = "terraform01",
    ///         });
    ///         var vpc = new AliCloud.Vpc.Network("vpc", new AliCloud.Vpc.NetworkArgs
    ///         {
    ///             CidrBlock = "192.168.0.0/16",
    ///         });
    ///         var foo = new AliCloud.Cen.InstanceAttachment("foo", new AliCloud.Cen.InstanceAttachmentArgs
    ///         {
    ///             InstanceId = cen.Id,
    ///             ChildInstanceId = vpc.Id,
    ///             ChildInstanceRegionId = "cn-beijing",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class InstanceAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the child instance to attach.
        /// </summary>
        [Output("childInstanceId")]
        public Output<string> ChildInstanceId { get; private set; } = null!;

        /// <summary>
        /// The uid of the child instance. Only used when attach a child instance of other account.
        /// </summary>
        [Output("childInstanceOwnerId")]
        public Output<string> ChildInstanceOwnerId { get; private set; } = null!;

        /// <summary>
        /// The region ID of the child instance to attach.
        /// </summary>
        [Output("childInstanceRegionId")]
        public Output<string> ChildInstanceRegionId { get; private set; } = null!;

        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceAttachment(string name, InstanceAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/instanceAttachment:InstanceAttachment", name, args ?? new InstanceAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceAttachment(string name, Input<string> id, InstanceAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/instanceAttachment:InstanceAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceAttachment Get(string name, Input<string> id, InstanceAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceAttachment(name, id, state, options);
        }
    }

    public sealed class InstanceAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the child instance to attach.
        /// </summary>
        [Input("childInstanceId", required: true)]
        public Input<string> ChildInstanceId { get; set; } = null!;

        /// <summary>
        /// The uid of the child instance. Only used when attach a child instance of other account.
        /// </summary>
        [Input("childInstanceOwnerId")]
        public Input<string>? ChildInstanceOwnerId { get; set; }

        /// <summary>
        /// The region ID of the child instance to attach.
        /// </summary>
        [Input("childInstanceRegionId", required: true)]
        public Input<string> ChildInstanceRegionId { get; set; } = null!;

        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public InstanceAttachmentArgs()
        {
        }
    }

    public sealed class InstanceAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the child instance to attach.
        /// </summary>
        [Input("childInstanceId")]
        public Input<string>? ChildInstanceId { get; set; }

        /// <summary>
        /// The uid of the child instance. Only used when attach a child instance of other account.
        /// </summary>
        [Input("childInstanceOwnerId")]
        public Input<string>? ChildInstanceOwnerId { get; set; }

        /// <summary>
        /// The region ID of the child instance to attach.
        /// </summary>
        [Input("childInstanceRegionId")]
        public Input<string>? ChildInstanceRegionId { get; set; }

        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public InstanceAttachmentState()
        {
        }
    }
}
