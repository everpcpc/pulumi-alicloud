// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.EventBridge
{
    /// <summary>
    /// Using this data source can create Event Bridge service-linked roles(SLR). EventBridge may need to access another Alibaba Cloud service to implement a specific feature. In this case, EventBridge must assume a specific service-linked role, which is a Resource Access Management (RAM) role, to obtain permissions to access another Alibaba Cloud service.
    /// 
    /// For information about Event Bridge service-linked roles(SLR) and how to use it, see [What is service-linked roles](https://www.alibabacloud.com/help/doc-detail/181425.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.129.0+
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var slr = new AliCloud.EventBridge.Slr("slr", new AliCloud.EventBridge.SlrArgs
    ///         {
    ///             ProductName = "AliyunServiceRoleForEventBridgeSourceRocketMQ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:eventbridge/slr:Slr")]
    public partial class Slr : Pulumi.CustomResource
    {
        /// <summary>
        /// The product name for SLR. EventBridge can automatically create the following service-linked roles:
        /// Event source related: `AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
        /// Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
        /// </summary>
        [Output("productName")]
        public Output<string> ProductName { get; private set; } = null!;


        /// <summary>
        /// Create a Slr resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Slr(string name, SlrArgs args, CustomResourceOptions? options = null)
            : base("alicloud:eventbridge/slr:Slr", name, args ?? new SlrArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Slr(string name, Input<string> id, SlrState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:eventbridge/slr:Slr", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Slr resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Slr Get(string name, Input<string> id, SlrState? state = null, CustomResourceOptions? options = null)
        {
            return new Slr(name, id, state, options);
        }
    }

    public sealed class SlrArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The product name for SLR. EventBridge can automatically create the following service-linked roles:
        /// Event source related: `AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
        /// Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
        /// </summary>
        [Input("productName", required: true)]
        public Input<string> ProductName { get; set; } = null!;

        public SlrArgs()
        {
        }
    }

    public sealed class SlrState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The product name for SLR. EventBridge can automatically create the following service-linked roles:
        /// Event source related: `AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
        /// Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
        /// </summary>
        [Input("productName")]
        public Input<string>? ProductName { get; set; }

        public SlrState()
        {
        }
    }
}