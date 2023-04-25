// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ResourceManager
{
    /// <summary>
    /// Provides a Resource Manager Control Policy Attachment resource.
    /// 
    /// For information about Resource Manager Control Policy Attachment and how to use it, see [What is Control Policy Attachment](https://help.aliyun.com/document_detail/208330.html).
    /// 
    /// &gt; **NOTE:** Available in v1.120.0+.
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
    ///     // Enable the control policy
    ///     var exampleResourceDirectory = new AliCloud.ResourceManager.ResourceDirectory("exampleResourceDirectory", new()
    ///     {
    ///         Status = "Enabled",
    ///     });
    /// 
    ///     var exampleControlPolicy = new AliCloud.ResourceManager.ControlPolicy("exampleControlPolicy", new()
    ///     {
    ///         ControlPolicyName = "tf-testAccName",
    ///         Description = "tf-testAccRDControlPolicy",
    ///         EffectScope = "RAM",
    ///         PolicyDocument = @"  {
    ///     ""Version"": ""1"",
    ///     ""Statement"": [
    ///       {
    ///         ""Effect"": ""Deny"",
    ///         ""Action"": [
    ///           ""ram:UpdateRole"",
    ///           ""ram:DeleteRole"",
    ///           ""ram:AttachPolicyToRole"",
    ///           ""ram:DetachPolicyFromRole""
    ///         ],
    ///         ""Resource"": ""acs:ram:*:*:role/ResourceDirectoryAccountAccessRole""
    ///       }
    ///     ]
    ///   }
    /// ",
    ///     });
    /// 
    ///     var exampleFolder = new AliCloud.ResourceManager.Folder("exampleFolder", new()
    ///     {
    ///         FolderName = "tf-testAccName",
    ///     });
    /// 
    ///     var exampleControlPolicyAttachment = new AliCloud.ResourceManager.ControlPolicyAttachment("exampleControlPolicyAttachment", new()
    ///     {
    ///         PolicyId = exampleControlPolicy.Id,
    ///         TargetId = exampleFolder.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             exampleResourceDirectory,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Resource Manager Control Policy Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:resourcemanager/controlPolicyAttachment:ControlPolicyAttachment example &lt;policy_id&gt;:&lt;target_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:resourcemanager/controlPolicyAttachment:ControlPolicyAttachment")]
    public partial class ControlPolicyAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of control policy.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        /// <summary>
        /// The ID of target.
        /// </summary>
        [Output("targetId")]
        public Output<string> TargetId { get; private set; } = null!;


        /// <summary>
        /// Create a ControlPolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ControlPolicyAttachment(string name, ControlPolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/controlPolicyAttachment:ControlPolicyAttachment", name, args ?? new ControlPolicyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ControlPolicyAttachment(string name, Input<string> id, ControlPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/controlPolicyAttachment:ControlPolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ControlPolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ControlPolicyAttachment Get(string name, Input<string> id, ControlPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ControlPolicyAttachment(name, id, state, options);
        }
    }

    public sealed class ControlPolicyAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of control policy.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        /// <summary>
        /// The ID of target.
        /// </summary>
        [Input("targetId", required: true)]
        public Input<string> TargetId { get; set; } = null!;

        public ControlPolicyAttachmentArgs()
        {
        }
        public static new ControlPolicyAttachmentArgs Empty => new ControlPolicyAttachmentArgs();
    }

    public sealed class ControlPolicyAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of control policy.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// The ID of target.
        /// </summary>
        [Input("targetId")]
        public Input<string>? TargetId { get; set; }

        public ControlPolicyAttachmentState()
        {
        }
        public static new ControlPolicyAttachmentState Empty => new ControlPolicyAttachmentState();
    }
}
