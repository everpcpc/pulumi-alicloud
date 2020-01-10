// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Datahub
{
    /// <summary>
    /// The subscription is the basic unit of resource usage in Datahub Service under Publish/Subscribe model. You can manage the relationships between user and topics by using subscriptions. [Refer to details](https://help.aliyun.com/document_detail/47440.html).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/datahub_subscription.html.markdown.
    /// </summary>
    public partial class Subscription : Pulumi.CustomResource
    {
        /// <summary>
        /// Comment of the datahub subscription. It cannot be longer than 255 characters.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Create time of the datahub subscription. It is a human-readable string rather than 64-bits UTC.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Last modify time of the datahub subscription. It is the same as *create_time* at the beginning. It is also a human-readable string rather than 64-bits UTC.
        /// </summary>
        [Output("lastModifyTime")]
        public Output<string> LastModifyTime { get; private set; } = null!;

        /// <summary>
        /// The name of the datahub project that the subscription belongs to. Its length is limited to 3-32 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The identidy of the subscritpion, generate from server side.
        /// </summary>
        [Output("subId")]
        public Output<string> SubId { get; private set; } = null!;

        /// <summary>
        /// The name of the datahub topic that the subscription belongs to. Its length is limited to 1-128 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        /// </summary>
        [Output("topicName")]
        public Output<string> TopicName { get; private set; } = null!;


        /// <summary>
        /// Create a Subscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subscription(string name, SubscriptionArgs args, CustomResourceOptions? options = null)
            : base("alicloud:datahub/subscription:Subscription", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Subscription(string name, Input<string> id, SubscriptionState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:datahub/subscription:Subscription", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Subscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subscription Get(string name, Input<string> id, SubscriptionState? state = null, CustomResourceOptions? options = null)
        {
            return new Subscription(name, id, state, options);
        }
    }

    public sealed class SubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment of the datahub subscription. It cannot be longer than 255 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The name of the datahub project that the subscription belongs to. Its length is limited to 3-32 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        /// <summary>
        /// The name of the datahub topic that the subscription belongs to. Its length is limited to 1-128 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public SubscriptionArgs()
        {
        }
    }

    public sealed class SubscriptionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment of the datahub subscription. It cannot be longer than 255 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Create time of the datahub subscription. It is a human-readable string rather than 64-bits UTC.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Last modify time of the datahub subscription. It is the same as *create_time* at the beginning. It is also a human-readable string rather than 64-bits UTC.
        /// </summary>
        [Input("lastModifyTime")]
        public Input<string>? LastModifyTime { get; set; }

        /// <summary>
        /// The name of the datahub project that the subscription belongs to. Its length is limited to 3-32 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The identidy of the subscritpion, generate from server side.
        /// </summary>
        [Input("subId")]
        public Input<string>? SubId { get; set; }

        /// <summary>
        /// The name of the datahub topic that the subscription belongs to. Its length is limited to 1-128 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        public SubscriptionState()
        {
        }
    }
}
