// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ess
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides available scheduled task resources. 
        /// 
        /// &gt; **NOTE:** Available in 1.72.0+
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ess_scheduled_tasks.html.markdown.
        /// </summary>
        [Obsolete("Use GetScheduledTasks.InvokeAsync() instead")]
        public static Task<GetScheduledTasksResult> GetScheduledTasks(GetScheduledTasksArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetScheduledTasksResult>("alicloud:ess/getScheduledTasks:getScheduledTasks", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetScheduledTasks
    {
        /// <summary>
        /// This data source provides available scheduled task resources. 
        /// 
        /// &gt; **NOTE:** Available in 1.72.0+
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ess_scheduled_tasks.html.markdown.
        /// </summary>
        public static Task<GetScheduledTasksResult> InvokeAsync(GetScheduledTasksArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetScheduledTasksResult>("alicloud:ess/getScheduledTasks:getScheduledTasks", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetScheduledTasksArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of scheduled task IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter resulting scheduled tasks by name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The operation to be performed when a scheduled task is triggered.
        /// </summary>
        [Input("scheduledAction")]
        public string? ScheduledAction { get; set; }

        /// <summary>
        /// The id of the scheduled task.
        /// </summary>
        [Input("scheduledTaskId")]
        public string? ScheduledTaskId { get; set; }

        public GetScheduledTasksArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetScheduledTasksResult
    {
        /// <summary>
        /// A list of scheduled task ids.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of scheduled task names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// The operation to be performed when a scheduled task is triggered.
        /// </summary>
        public readonly string? ScheduledAction;
        public readonly string? ScheduledTaskId;
        /// <summary>
        /// A list of scheduled tasks. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetScheduledTasksTasksResult> Tasks;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetScheduledTasksResult(
            ImmutableArray<string> ids,
            string? nameRegex,
            ImmutableArray<string> names,
            string? outputFile,
            string? scheduledAction,
            string? scheduledTaskId,
            ImmutableArray<Outputs.GetScheduledTasksTasksResult> tasks,
            string id)
        {
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ScheduledAction = scheduledAction;
            ScheduledTaskId = scheduledTaskId;
            Tasks = tasks;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetScheduledTasksTasksResult
    {
        /// <summary>
        /// Description of the scheduled task.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// ID of the scheduled task id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The time period during which a failed scheduled task is retried.
        /// </summary>
        public readonly int LaunchExpirationTime;
        /// <summary>
        /// The time at which the scheduled task is triggered.
        /// </summary>
        public readonly string LaunchTime;
        public readonly int MaxValue;
        public readonly int MinValue;
        /// <summary>
        /// Name of the scheduled task name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the end time after which the scheduled task is no longer repeated.
        /// </summary>
        public readonly string RecurrenceEndTime;
        /// <summary>
        /// Specifies the recurrence type of the scheduled task. 
        /// </summary>
        public readonly string RecurrenceType;
        /// <summary>
        /// Specifies how often a scheduled task recurs. 
        /// </summary>
        public readonly string RecurrenceValue;
        /// <summary>
        /// The operation to be performed when a scheduled task is triggered.
        /// </summary>
        public readonly string ScheduledAction;
        public readonly bool TaskEnabled;

        [OutputConstructor]
        private GetScheduledTasksTasksResult(
            string description,
            string id,
            int launchExpirationTime,
            string launchTime,
            int maxValue,
            int minValue,
            string name,
            string recurrenceEndTime,
            string recurrenceType,
            string recurrenceValue,
            string scheduledAction,
            bool taskEnabled)
        {
            Description = description;
            Id = id;
            LaunchExpirationTime = launchExpirationTime;
            LaunchTime = launchTime;
            MaxValue = maxValue;
            MinValue = minValue;
            Name = name;
            RecurrenceEndTime = recurrenceEndTime;
            RecurrenceType = recurrenceType;
            RecurrenceValue = recurrenceValue;
            ScheduledAction = scheduledAction;
            TaskEnabled = taskEnabled;
        }
    }
    }
}