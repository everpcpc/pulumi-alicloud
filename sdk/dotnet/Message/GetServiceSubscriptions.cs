// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Message
{
    public static class GetServiceSubscriptions
    {
        /// <summary>
        /// This data source provides the Message Notification Service Subscriptions of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.188.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.Message.GetServiceSubscriptions.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///         TopicName = "tf-example",
        ///     });
        /// 
        ///     var name = AliCloud.Message.GetServiceSubscriptions.Invoke(new()
        ///     {
        ///         TopicName = "tf-example",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["subscriptionId1"] = ids.Apply(getServiceSubscriptionsResult =&gt; getServiceSubscriptionsResult.Subscriptions[0]?.Id),
        ///         ["subscriptionId2"] = name.Apply(getServiceSubscriptionsResult =&gt; getServiceSubscriptionsResult.Subscriptions[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServiceSubscriptionsResult> InvokeAsync(GetServiceSubscriptionsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceSubscriptionsResult>("alicloud:message/getServiceSubscriptions:getServiceSubscriptions", args ?? new GetServiceSubscriptionsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Message Notification Service Subscriptions of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.188.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.Message.GetServiceSubscriptions.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///         TopicName = "tf-example",
        ///     });
        /// 
        ///     var name = AliCloud.Message.GetServiceSubscriptions.Invoke(new()
        ///     {
        ///         TopicName = "tf-example",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["subscriptionId1"] = ids.Apply(getServiceSubscriptionsResult =&gt; getServiceSubscriptionsResult.Subscriptions[0]?.Id),
        ///         ["subscriptionId2"] = name.Apply(getServiceSubscriptionsResult =&gt; getServiceSubscriptionsResult.Subscriptions[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServiceSubscriptionsResult> Invoke(GetServiceSubscriptionsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceSubscriptionsResult>("alicloud:message/getServiceSubscriptions:getServiceSubscriptions", args ?? new GetServiceSubscriptionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceSubscriptionsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Subscription IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Subscription name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        /// <summary>
        /// The name of the subscription.
        /// </summary>
        [Input("subscriptionName")]
        public string? SubscriptionName { get; set; }

        /// <summary>
        /// The name of the topic.
        /// </summary>
        [Input("topicName", required: true)]
        public string TopicName { get; set; } = null!;

        public GetServiceSubscriptionsArgs()
        {
        }
        public static new GetServiceSubscriptionsArgs Empty => new GetServiceSubscriptionsArgs();
    }

    public sealed class GetServiceSubscriptionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Subscription IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Subscription name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        /// <summary>
        /// The name of the subscription.
        /// </summary>
        [Input("subscriptionName")]
        public Input<string>? SubscriptionName { get; set; }

        /// <summary>
        /// The name of the topic.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public GetServiceSubscriptionsInvokeArgs()
        {
        }
        public static new GetServiceSubscriptionsInvokeArgs Empty => new GetServiceSubscriptionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceSubscriptionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of Subscription names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;
        /// <summary>
        /// The name of the subscription.
        /// </summary>
        public readonly string? SubscriptionName;
        /// <summary>
        /// A list of Subscriptions. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServiceSubscriptionsSubscriptionResult> Subscriptions;
        /// <summary>
        /// The name of the topic.
        /// </summary>
        public readonly string TopicName;

        [OutputConstructor]
        private GetServiceSubscriptionsResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            int? pageNumber,

            int? pageSize,

            string? subscriptionName,

            ImmutableArray<Outputs.GetServiceSubscriptionsSubscriptionResult> subscriptions,

            string topicName)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
            SubscriptionName = subscriptionName;
            Subscriptions = subscriptions;
            TopicName = topicName;
        }
    }
}
