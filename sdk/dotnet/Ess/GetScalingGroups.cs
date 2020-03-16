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
        /// This data source provides available scaling group resources. 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ess_scaling_groups.html.markdown.
        /// </summary>
        public static Task<GetScalingGroupsResult> GetScalingGroups(GetScalingGroupsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetScalingGroupsResult>("alicloud:ess/getScalingGroups:getScalingGroups", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetScalingGroupsArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of scaling group IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter resulting scaling groups by name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetScalingGroupsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetScalingGroupsResult
    {
        /// <summary>
        /// A list of scaling groups. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetScalingGroupsGroupsResult> Groups;
        /// <summary>
        /// A list of scaling group ids.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of scaling group names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetScalingGroupsResult(
            ImmutableArray<Outputs.GetScalingGroupsGroupsResult> groups,
            ImmutableArray<string> ids,
            string? nameRegex,
            ImmutableArray<string> names,
            string? outputFile,
            string id)
        {
            Groups = groups;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetScalingGroupsGroupsResult
    {
        /// <summary>
        /// Number of active instances in scaling group.
        /// </summary>
        public readonly int ActiveCapacity;
        public readonly string ActiveScalingConfiguration;
        /// <summary>
        /// Default cooldown time of scaling group.
        /// </summary>
        public readonly int CooldownTime;
        /// <summary>
        /// Creation time of scaling group.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// Db instances id which the ECS instance attached to.
        /// </summary>
        public readonly ImmutableArray<string> DbInstanceIds;
        /// <summary>
        /// ID of the scaling group.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Active launch template ID for scaling group.
        /// </summary>
        public readonly string LaunchTemplateId;
        /// <summary>
        /// Version of active launch template.
        /// </summary>
        public readonly string LaunchTemplateVersion;
        /// <summary>
        /// Lifecycle state of scaling group.
        /// </summary>
        public readonly string LifecycleState;
        /// <summary>
        /// Slb instances id which the ECS instance attached to.
        /// </summary>
        public readonly ImmutableArray<string> LoadBalancerIds;
        /// <summary>
        /// The maximum number of ECS instances.
        /// </summary>
        public readonly int MaxSize;
        /// <summary>
        /// The minimum number of ECS instances.
        /// </summary>
        public readonly int MinSize;
        /// <summary>
        /// Name of the scaling group.
        /// * `active_scaling_configuration` -Active scaling configuration for scaling group.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Number of pending instances in scaling group.
        /// </summary>
        public readonly int PendingCapacity;
        /// <summary>
        /// Region ID the scaling group belongs to.
        /// </summary>
        public readonly string RegionId;
        /// <summary>
        /// Removal policy used to select the ECS instance to remove from the scaling group.
        /// </summary>
        public readonly ImmutableArray<string> RemovalPolicies;
        /// <summary>
        /// Number of removing instances in scaling group.
        /// </summary>
        public readonly int RemovingCapacity;
        /// <summary>
        /// Number of instances in scaling group.
        /// </summary>
        public readonly int TotalCapacity;
        /// <summary>
        /// Vswitches id in which the ECS instance launched.
        /// </summary>
        public readonly ImmutableArray<string> VswitchIds;

        [OutputConstructor]
        private GetScalingGroupsGroupsResult(
            int activeCapacity,
            string activeScalingConfiguration,
            int cooldownTime,
            string creationTime,
            ImmutableArray<string> dbInstanceIds,
            string id,
            string launchTemplateId,
            string launchTemplateVersion,
            string lifecycleState,
            ImmutableArray<string> loadBalancerIds,
            int maxSize,
            int minSize,
            string name,
            int pendingCapacity,
            string regionId,
            ImmutableArray<string> removalPolicies,
            int removingCapacity,
            int totalCapacity,
            ImmutableArray<string> vswitchIds)
        {
            ActiveCapacity = activeCapacity;
            ActiveScalingConfiguration = activeScalingConfiguration;
            CooldownTime = cooldownTime;
            CreationTime = creationTime;
            DbInstanceIds = dbInstanceIds;
            Id = id;
            LaunchTemplateId = launchTemplateId;
            LaunchTemplateVersion = launchTemplateVersion;
            LifecycleState = lifecycleState;
            LoadBalancerIds = loadBalancerIds;
            MaxSize = maxSize;
            MinSize = minSize;
            Name = name;
            PendingCapacity = pendingCapacity;
            RegionId = regionId;
            RemovalPolicies = removalPolicies;
            RemovingCapacity = removingCapacity;
            TotalCapacity = totalCapacity;
            VswitchIds = vswitchIds;
        }
    }
    }
}
