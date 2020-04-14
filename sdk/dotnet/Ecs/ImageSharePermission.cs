// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    /// <summary>
    /// Manage image sharing permissions. You can share your custom image to other Alibaba Cloud users. The user can use the shared custom image to create ECS instances or replace the system disk of the instance.
    /// 
    /// &gt; **NOTE:** You can only share your own custom images to other Alibaba Cloud users.
    /// 
    /// &gt; **NOTE:** Each custom image can be shared with up to 50 Alibaba Cloud accounts. You can submit a ticket to share with more users.
    /// 
    /// &gt; **NOTE:** After creating an ECS instance using a shared image, once the custom image owner releases the image sharing relationship or deletes the custom image, the instance cannot initialize the system disk.
    /// 
    /// &gt; **NOTE:** Available in 1.68.0+.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/image_share_permission.html.markdown.
    /// </summary>
    public partial class ImageSharePermission : Pulumi.CustomResource
    {
        /// <summary>
        /// Alibaba Cloud Account ID. It is used to share images.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The source image ID.
        /// </summary>
        [Output("imageId")]
        public Output<string> ImageId { get; private set; } = null!;


        /// <summary>
        /// Create a ImageSharePermission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImageSharePermission(string name, ImageSharePermissionArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ecs/imageSharePermission:ImageSharePermission", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ImageSharePermission(string name, Input<string> id, ImageSharePermissionState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/imageSharePermission:ImageSharePermission", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ImageSharePermission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImageSharePermission Get(string name, Input<string> id, ImageSharePermissionState? state = null, CustomResourceOptions? options = null)
        {
            return new ImageSharePermission(name, id, state, options);
        }
    }

    public sealed class ImageSharePermissionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alibaba Cloud Account ID. It is used to share images.
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// The source image ID.
        /// </summary>
        [Input("imageId", required: true)]
        public Input<string> ImageId { get; set; } = null!;

        public ImageSharePermissionArgs()
        {
        }
    }

    public sealed class ImageSharePermissionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alibaba Cloud Account ID. It is used to share images.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The source image ID.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        public ImageSharePermissionState()
        {
        }
    }
}
