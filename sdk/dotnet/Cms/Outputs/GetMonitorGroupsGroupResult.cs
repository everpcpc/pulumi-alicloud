// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms.Outputs
{

    [OutputType]
    public sealed class GetMonitorGroupsGroupResult
    {
        /// <summary>
        /// The URL of the Kubernetes cluster from which the application group is synchronized.
        /// </summary>
        public readonly string BindUrl;
        /// <summary>
        /// The list of  alert groups that receive alert notifications for the application group.
        /// </summary>
        public readonly ImmutableArray<string> ContactGroups;
        /// <summary>
        /// The ID of the tag rule.
        /// </summary>
        public readonly string DynamicTagRuleId;
        /// <summary>
        /// The time when the application group was created.
        /// </summary>
        public readonly int GmtCreate;
        /// <summary>
        /// The time when the application group was modified.
        /// </summary>
        public readonly int GmtModified;
        /// <summary>
        /// The ID of the application group.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// The ID of the Monitor Group.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the application group.
        /// </summary>
        public readonly string MonitorGroupName;
        /// <summary>
        /// The ID of the Alibaba Cloud service.
        /// </summary>
        public readonly string ServiceId;
        /// <summary>
        /// A map of tags assigned to the Cms Monitor Group.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// The alert templates applied to the application group.
        /// </summary>
        public readonly ImmutableArray<string> TemplateIds;
        /// <summary>
        /// The type of the application group.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMonitorGroupsGroupResult(
            string bindUrl,

            ImmutableArray<string> contactGroups,

            string dynamicTagRuleId,

            int gmtCreate,

            int gmtModified,

            string groupId,

            string id,

            string monitorGroupName,

            string serviceId,

            ImmutableDictionary<string, object> tags,

            ImmutableArray<string> templateIds,

            string type)
        {
            BindUrl = bindUrl;
            ContactGroups = contactGroups;
            DynamicTagRuleId = dynamicTagRuleId;
            GmtCreate = gmtCreate;
            GmtModified = gmtModified;
            GroupId = groupId;
            Id = id;
            MonitorGroupName = monitorGroupName;
            ServiceId = serviceId;
            Tags = tags;
            TemplateIds = templateIds;
            Type = type;
        }
    }
}
