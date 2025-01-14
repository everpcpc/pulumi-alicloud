// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.BastionHost
{
    /// <summary>
    /// Provides a Bastion Host User Attachment resource to add user to one user group.
    /// 
    /// &gt; **NOTE:** Available in v1.134.0+.
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
    ///     var example = new AliCloud.BastionHost.UserAttachment("example", new()
    ///     {
    ///         InstanceId = "bastionhost-cn-tl3xxxxxxx",
    ///         UserGroupId = "10",
    ///         UserId = "100",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Bastion Host User Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:bastionhost/userAttachment:UserAttachment example &lt;instance_id&gt;:&lt;user_group_id&gt;:&lt;user_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:bastionhost/userAttachment:UserAttachment")]
    public partial class UserAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the user group to add the user's bastion host ID of.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Specifies the user group to which you want to add the user ID.
        /// </summary>
        [Output("userGroupId")]
        public Output<string> UserGroupId { get; private set; } = null!;

        /// <summary>
        /// Specify that you want to add to the policy attached to the user group ID. This includes response parameters in a Json-formatted string supports up to set up 100 USER ID.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a UserAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserAttachment(string name, UserAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:bastionhost/userAttachment:UserAttachment", name, args ?? new UserAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserAttachment(string name, Input<string> id, UserAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:bastionhost/userAttachment:UserAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserAttachment Get(string name, Input<string> id, UserAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new UserAttachment(name, id, state, options);
        }
    }

    public sealed class UserAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the user group to add the user's bastion host ID of.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Specifies the user group to which you want to add the user ID.
        /// </summary>
        [Input("userGroupId", required: true)]
        public Input<string> UserGroupId { get; set; } = null!;

        /// <summary>
        /// Specify that you want to add to the policy attached to the user group ID. This includes response parameters in a Json-formatted string supports up to set up 100 USER ID.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public UserAttachmentArgs()
        {
        }
        public static new UserAttachmentArgs Empty => new UserAttachmentArgs();
    }

    public sealed class UserAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the user group to add the user's bastion host ID of.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Specifies the user group to which you want to add the user ID.
        /// </summary>
        [Input("userGroupId")]
        public Input<string>? UserGroupId { get; set; }

        /// <summary>
        /// Specify that you want to add to the policy attached to the user group ID. This includes response parameters in a Json-formatted string supports up to set up 100 USER ID.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public UserAttachmentState()
        {
        }
        public static new UserAttachmentState Empty => new UserAttachmentState();
    }
}
