// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Mns
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides a list of MNS queues in an Alibaba Cloud account according to the specified parameters.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/mns_queues.html.markdown.
        /// </summary>
        [Obsolete("Use GetQueues.InvokeAsync() instead")]
        public static Task<GetQueuesResult> GetQueues(GetQueuesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetQueuesResult>("alicloud:mns/getQueues:getQueues", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetQueues
    {
        /// <summary>
        /// This data source provides a list of MNS queues in an Alibaba Cloud account according to the specified parameters.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/mns_queues.html.markdown.
        /// </summary>
        public static Task<GetQueuesResult> InvokeAsync(GetQueuesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetQueuesResult>("alicloud:mns/getQueues:getQueues", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetQueuesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A string to filter resulting queues by their name prefixs.
        /// </summary>
        [Input("namePrefix")]
        public string? NamePrefix { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetQueuesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetQueuesResult
    {
        public readonly string? NamePrefix;
        /// <summary>
        /// A list of queue names. 
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of queues. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetQueuesQueuesResult> Queues;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetQueuesResult(
            string? namePrefix,
            ImmutableArray<string> names,
            string? outputFile,
            ImmutableArray<Outputs.GetQueuesQueuesResult> queues,
            string id)
        {
            NamePrefix = namePrefix;
            Names = names;
            OutputFile = outputFile;
            Queues = queues;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetQueuesQueuesResult
    {
        /// <summary>
        /// This attribute defines the length of time, in seconds, after which every message sent to the queue is dequeued.
        /// </summary>
        public readonly int DelaySeconds;
        /// <summary>
        /// The id of the queue, The value is set to `name`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// This indicates the maximum length, in bytes, of any message body sent to the queue.
        /// </summary>
        public readonly int MaximumMessageSize;
        /// <summary>
        /// Messages are deleted from the queue after a specified length of time, whether they have been activated or not. This attribute defines the viability period, in seconds, for every message in the queue.
        /// </summary>
        public readonly int MessageRetentionPeriod;
        /// <summary>
        /// The name of the queue
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Long polling is measured in seconds. When this attribute is set to 0, long polling is disabled. When it is not set to 0, long polling is enabled and message dequeue requests will be processed only when valid messages are received or when long polling times out.
        /// </summary>
        public readonly int PollingWaitSeconds;
        /// <summary>
        /// Dequeued messages change from active (visible) status to inactive (invisible) status. This attribute defines the length of time, in seconds, that messages remain invisible. Messages return to active status after the set period.
        /// </summary>
        public readonly int VisibilityTimeouts;

        [OutputConstructor]
        private GetQueuesQueuesResult(
            int delaySeconds,
            string id,
            int maximumMessageSize,
            int messageRetentionPeriod,
            string name,
            int pollingWaitSeconds,
            int visibilityTimeouts)
        {
            DelaySeconds = delaySeconds;
            Id = id;
            MaximumMessageSize = maximumMessageSize;
            MessageRetentionPeriod = messageRetentionPeriod;
            Name = name;
            PollingWaitSeconds = pollingWaitSeconds;
            VisibilityTimeouts = visibilityTimeouts;
        }
    }
    }
}
