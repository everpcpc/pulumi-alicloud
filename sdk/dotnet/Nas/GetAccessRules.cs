// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nas
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides AccessRule available to the user.
        /// 
        /// &gt; NOTE: Available in 1.35.0+
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/nas_access_rules.html.markdown.
        /// </summary>
        public static Task<GetAccessRulesResult> GetAccessRules(GetAccessRulesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccessRulesResult>("alicloud:nas/getAccessRules:getAccessRules", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetAccessRulesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter results by a specific AccessGroupName.
        /// </summary>
        [Input("accessGroupName", required: true)]
        public string AccessGroupName { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of rule IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Filter results by a specific RWAccess. 
        /// </summary>
        [Input("rwAccess")]
        public string? RwAccess { get; set; }

        /// <summary>
        /// Filter results by a specific SourceCidrIp. 
        /// </summary>
        [Input("sourceCidrIp")]
        public string? SourceCidrIp { get; set; }

        /// <summary>
        /// Filter results by a specific UserAccess. 
        /// </summary>
        [Input("userAccess")]
        public string? UserAccess { get; set; }

        public GetAccessRulesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetAccessRulesResult
    {
        public readonly string AccessGroupName;
        /// <summary>
        /// A list of rule IDs, Each element set to `access_rule_id` (Each element formats as `&lt;access_group_name&gt;:&lt;access rule id&gt;` before 1.53.0).
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of AccessRules. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccessRulesRulesResult> Rules;
        /// <summary>
        /// RWAccess of the AccessRule.
        /// </summary>
        public readonly string? RwAccess;
        /// <summary>
        /// SourceCidrIp of the AccessRule.
        /// </summary>
        public readonly string? SourceCidrIp;
        /// <summary>
        /// UserAccess of the AccessRule
        /// </summary>
        public readonly string? UserAccess;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAccessRulesResult(
            string accessGroupName,
            ImmutableArray<string> ids,
            string? outputFile,
            ImmutableArray<Outputs.GetAccessRulesRulesResult> rules,
            string? rwAccess,
            string? sourceCidrIp,
            string? userAccess,
            string id)
        {
            AccessGroupName = accessGroupName;
            Ids = ids;
            OutputFile = outputFile;
            Rules = rules;
            RwAccess = rwAccess;
            SourceCidrIp = sourceCidrIp;
            UserAccess = userAccess;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetAccessRulesRulesResult
    {
        /// <summary>
        /// AccessRuleId of the AccessRule.
        /// </summary>
        public readonly string AccessRuleId;
        /// <summary>
        /// Priority of the AccessRule.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// Filter results by a specific RWAccess. 
        /// </summary>
        public readonly string RwAccess;
        /// <summary>
        /// Filter results by a specific SourceCidrIp. 
        /// </summary>
        public readonly string SourceCidrIp;
        /// <summary>
        /// Filter results by a specific UserAccess. 
        /// </summary>
        public readonly string UserAccess;

        [OutputConstructor]
        private GetAccessRulesRulesResult(
            string accessRuleId,
            int priority,
            string rwAccess,
            string sourceCidrIp,
            string userAccess)
        {
            AccessRuleId = accessRuleId;
            Priority = priority;
            RwAccess = rwAccess;
            SourceCidrIp = sourceCidrIp;
            UserAccess = userAccess;
        }
    }
    }
}
