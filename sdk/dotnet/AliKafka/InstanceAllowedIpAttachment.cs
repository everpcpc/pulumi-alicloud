// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.AliKafka
{
    /// <summary>
    /// Provides a AliKafka Instance Allowed Ip Attachment resource.
    /// 
    /// For information about Ali Kafka Instance Allowed Ip Attachment and how to use it, see [What is Instance Allowed Ip Attachment](https://www.alibabacloud.com/help/en/doc-detail/68151.html).
    /// 
    /// &gt; **NOTE:** Available in v1.163.0+.
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
    ///         var name = config.Get("name") ?? "tftest";
    ///         var defaultNetworks = Output.Create(AliCloud.Vpc.GetNetworks.InvokeAsync(new AliCloud.Vpc.GetNetworksArgs
    ///         {
    ///             NameRegex = "^default-NODELETING",
    ///         }));
    ///         var defaultSwitches = defaultNetworks.Apply(defaultNetworks =&gt; Output.Create(AliCloud.Vpc.GetSwitches.InvokeAsync(new AliCloud.Vpc.GetSwitchesArgs
    ///         {
    ///             VpcId = defaultNetworks.Ids?[0],
    ///         })));
    ///         var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("defaultSecurityGroup", new AliCloud.Ecs.SecurityGroupArgs
    ///         {
    ///             VpcId = defaultNetworks.Apply(defaultNetworks =&gt; defaultNetworks.Ids?[0]),
    ///         });
    ///         var defaultInstance = new AliCloud.AliKafka.Instance("defaultInstance", new AliCloud.AliKafka.InstanceArgs
    ///         {
    ///             TopicQuota = 50,
    ///             DiskType = 1,
    ///             DiskSize = 500,
    ///             DeployType = 5,
    ///             IoMax = 20,
    ///             VswitchId = defaultSwitches.Apply(defaultSwitches =&gt; defaultSwitches.Ids?[0]),
    ///             SecurityGroup = defaultSecurityGroup.Id,
    ///         });
    ///         var defaultInstanceAllowedIpAttachment = new AliCloud.AliKafka.InstanceAllowedIpAttachment("defaultInstanceAllowedIpAttachment", new AliCloud.AliKafka.InstanceAllowedIpAttachmentArgs
    ///         {
    ///             AllowedIp = "114.237.9.78/32",
    ///             AllowedType = "vpc",
    ///             InstanceId = defaultInstance.Id,
    ///             PortRange = "9092/9092",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// AliKafka Instance Allowed Ip Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:alikafka/instanceAllowedIpAttachment:InstanceAllowedIpAttachment example &lt;instance_id&gt;:&lt;allowed_type&gt;:&lt;port_range&gt;:&lt;allowed_ip&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:alikafka/instanceAllowedIpAttachment:InstanceAllowedIpAttachment")]
    public partial class InstanceAllowedIpAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The allowed ip. It can be a CIDR block.
        /// </summary>
        [Output("allowedIp")]
        public Output<string> AllowedIp { get; private set; } = null!;

        /// <summary>
        /// The type of whitelist. Valid Value: `vpc`.
        /// </summary>
        [Output("allowedType")]
        public Output<string> AllowedType { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The Port range.  Valid Value: `9092/9092`.
        /// </summary>
        [Output("portRange")]
        public Output<string> PortRange { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceAllowedIpAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceAllowedIpAttachment(string name, InstanceAllowedIpAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:alikafka/instanceAllowedIpAttachment:InstanceAllowedIpAttachment", name, args ?? new InstanceAllowedIpAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceAllowedIpAttachment(string name, Input<string> id, InstanceAllowedIpAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:alikafka/instanceAllowedIpAttachment:InstanceAllowedIpAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceAllowedIpAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceAllowedIpAttachment Get(string name, Input<string> id, InstanceAllowedIpAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceAllowedIpAttachment(name, id, state, options);
        }
    }

    public sealed class InstanceAllowedIpAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The allowed ip. It can be a CIDR block.
        /// </summary>
        [Input("allowedIp", required: true)]
        public Input<string> AllowedIp { get; set; } = null!;

        /// <summary>
        /// The type of whitelist. Valid Value: `vpc`.
        /// </summary>
        [Input("allowedType", required: true)]
        public Input<string> AllowedType { get; set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The Port range.  Valid Value: `9092/9092`.
        /// </summary>
        [Input("portRange", required: true)]
        public Input<string> PortRange { get; set; } = null!;

        public InstanceAllowedIpAttachmentArgs()
        {
        }
    }

    public sealed class InstanceAllowedIpAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The allowed ip. It can be a CIDR block.
        /// </summary>
        [Input("allowedIp")]
        public Input<string>? AllowedIp { get; set; }

        /// <summary>
        /// The type of whitelist. Valid Value: `vpc`.
        /// </summary>
        [Input("allowedType")]
        public Input<string>? AllowedType { get; set; }

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The Port range.  Valid Value: `9092/9092`.
        /// </summary>
        [Input("portRange")]
        public Input<string>? PortRange { get; set; }

        public InstanceAllowedIpAttachmentState()
        {
        }
    }
}