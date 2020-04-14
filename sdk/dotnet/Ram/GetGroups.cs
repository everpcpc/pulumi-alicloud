// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ram
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides a list of RAM Groups in an Alibaba Cloud account according to the specified filters.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ram_groups.html.markdown.
        /// </summary>
        [Obsolete("Use GetGroups.InvokeAsync() instead")]
        public static Task<GetGroupsResult> GetGroups(GetGroupsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupsResult>("alicloud:ram/getGroups:getGroups", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetGroups
    {
        /// <summary>
        /// This data source provides a list of RAM Groups in an Alibaba Cloud account according to the specified filters.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ram_groups.html.markdown.
        /// </summary>
        public static Task<GetGroupsResult> InvokeAsync(GetGroupsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupsResult>("alicloud:ram/getGroups:getGroups", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetGroupsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A regex string to filter the returned groups by their names.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Filter the results by a specific policy name. If you set this parameter without setting `policy_type`, it will be automatically set to `System`.
        /// </summary>
        [Input("policyName")]
        public string? PolicyName { get; set; }

        /// <summary>
        /// Filter the results by a specific policy type. Valid items are `Custom` and `System`. If you set this parameter, you must set `policy_name` as well.
        /// </summary>
        [Input("policyType")]
        public string? PolicyType { get; set; }

        /// <summary>
        /// Filter the results by a specific the user name.
        /// </summary>
        [Input("userName")]
        public string? UserName { get; set; }

        public GetGroupsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetGroupsResult
    {
        /// <summary>
        /// A list of groups. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupsGroupsResult> Groups;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of ram group names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? PolicyName;
        public readonly string? PolicyType;
        public readonly string? UserName;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetGroupsResult(
            ImmutableArray<Outputs.GetGroupsGroupsResult> groups,
            string? nameRegex,
            ImmutableArray<string> names,
            string? outputFile,
            string? policyName,
            string? policyType,
            string? userName,
            string id)
        {
            Groups = groups;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PolicyName = policyName;
            PolicyType = policyType;
            UserName = userName;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetGroupsGroupsResult
    {
        /// <summary>
        /// Comments of the group.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// Name of the group.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetGroupsGroupsResult(
            string comments,
            string name)
        {
            Comments = comments;
            Name = name;
        }
    }
    }
}
