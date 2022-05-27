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
    /// Provides a Bastion Host Account Share Key Attachment resource.
    /// 
    /// For information about Bastion Host Host Account Share Key Attachment and how to use it, see [What is Host Account Share Key Attachment](https://www.alibabacloud.com/help/en/bastion-host/latest/attachhostaccountstohostsharekey).
    /// 
    /// &gt; **NOTE:** Available in v1.165.0+.
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
    ///         var name = config.Get("name") ?? "tfacc_host_account_share_key_attachment";
    ///         var defaultInstances = Output.Create(AliCloud.BastionHost.GetInstances.InvokeAsync());
    ///         var defaultHostShareKey = new AliCloud.BastionHost.HostShareKey("defaultHostShareKey", new AliCloud.BastionHost.HostShareKeyArgs
    ///         {
    ///             HostShareKeyName = "example_name",
    ///             InstanceId = defaultInstances.Apply(defaultInstances =&gt; defaultInstances.Instances?[0]?.Id),
    ///             PassPhrase = "example_value",
    ///             PrivateKey = "example_value",
    ///         });
    ///         var defaultHost = new AliCloud.BastionHost.Host("defaultHost", new AliCloud.BastionHost.HostArgs
    ///         {
    ///             InstanceId = defaultInstances.Apply(defaultInstances =&gt; defaultInstances.Ids?[0]),
    ///             HostName = name,
    ///             ActiveAddressType = "Private",
    ///             HostPrivateAddress = "172.16.0.10",
    ///             OsType = "Linux",
    ///             Source = "Local",
    ///         });
    ///         var defaultHostAccount = new AliCloud.BastionHost.HostAccount("defaultHostAccount", new AliCloud.BastionHost.HostAccountArgs
    ///         {
    ///             InstanceId = defaultInstances.Apply(defaultInstances =&gt; defaultInstances.Ids?[0]),
    ///             HostAccountName = name,
    ///             HostId = defaultHost.HostId,
    ///             ProtocolName = "SSH",
    ///             Password = "YourPassword12345",
    ///         });
    ///         var defaultHostAccountShareKeyAttachment = new AliCloud.BastionHost.HostAccountShareKeyAttachment("defaultHostAccountShareKeyAttachment", new AliCloud.BastionHost.HostAccountShareKeyAttachmentArgs
    ///         {
    ///             InstanceId = defaultInstances.Apply(defaultInstances =&gt; defaultInstances.Instances?[0]?.Id),
    ///             HostShareKeyId = defaultHostShareKey.HostShareKeyId,
    ///             HostAccountId = defaultHostAccount.HostAccountId,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Bastion Host Account Share Key Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:bastionhost/hostAccountShareKeyAttachment:HostAccountShareKeyAttachment example &lt;instance_id&gt;:&lt;host_share_key_id&gt;:&lt;host_account_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:bastionhost/hostAccountShareKeyAttachment:HostAccountShareKeyAttachment")]
    public partial class HostAccountShareKeyAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID list of the host account.
        /// </summary>
        [Output("hostAccountId")]
        public Output<string> HostAccountId { get; private set; } = null!;

        /// <summary>
        /// The ID of the host shared key.
        /// </summary>
        [Output("hostShareKeyId")]
        public Output<string> HostShareKeyId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Bastion machine instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a HostAccountShareKeyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HostAccountShareKeyAttachment(string name, HostAccountShareKeyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:bastionhost/hostAccountShareKeyAttachment:HostAccountShareKeyAttachment", name, args ?? new HostAccountShareKeyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HostAccountShareKeyAttachment(string name, Input<string> id, HostAccountShareKeyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:bastionhost/hostAccountShareKeyAttachment:HostAccountShareKeyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HostAccountShareKeyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HostAccountShareKeyAttachment Get(string name, Input<string> id, HostAccountShareKeyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new HostAccountShareKeyAttachment(name, id, state, options);
        }
    }

    public sealed class HostAccountShareKeyAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID list of the host account.
        /// </summary>
        [Input("hostAccountId", required: true)]
        public Input<string> HostAccountId { get; set; } = null!;

        /// <summary>
        /// The ID of the host shared key.
        /// </summary>
        [Input("hostShareKeyId", required: true)]
        public Input<string> HostShareKeyId { get; set; } = null!;

        /// <summary>
        /// The ID of the Bastion machine instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public HostAccountShareKeyAttachmentArgs()
        {
        }
    }

    public sealed class HostAccountShareKeyAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID list of the host account.
        /// </summary>
        [Input("hostAccountId")]
        public Input<string>? HostAccountId { get; set; }

        /// <summary>
        /// The ID of the host shared key.
        /// </summary>
        [Input("hostShareKeyId")]
        public Input<string>? HostShareKeyId { get; set; }

        /// <summary>
        /// The ID of the Bastion machine instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public HostAccountShareKeyAttachmentState()
        {
        }
    }
}
