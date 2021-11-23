// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Arms.Outputs
{

    [OutputType]
    public sealed class GetDispatchRulesRuleResult
    {
        /// <summary>
        /// Dispatch rule ID.
        /// </summary>
        public readonly string DispatchRuleId;
        /// <summary>
        /// The name of the dispatch rule.
        /// </summary>
        public readonly string DispatchRuleName;
        public readonly string DispatchType;
        /// <summary>
        /// Sets the event group.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDispatchRulesRuleGroupRuleResult> GroupRules;
        /// <summary>
        /// The ID of the Dispatch Rule.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Sets the dispatch rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDispatchRulesRuleLabelMatchExpressionGridResult> LabelMatchExpressionGrids;
        /// <summary>
        /// Sets the notification rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDispatchRulesRuleNotifyRuleResult> NotifyRules;
        /// <summary>
        /// The resource status of Alert Dispatch Rule.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetDispatchRulesRuleResult(
            string dispatchRuleId,

            string dispatchRuleName,

            string dispatchType,

            ImmutableArray<Outputs.GetDispatchRulesRuleGroupRuleResult> groupRules,

            string id,

            ImmutableArray<Outputs.GetDispatchRulesRuleLabelMatchExpressionGridResult> labelMatchExpressionGrids,

            ImmutableArray<Outputs.GetDispatchRulesRuleNotifyRuleResult> notifyRules,

            string status)
        {
            DispatchRuleId = dispatchRuleId;
            DispatchRuleName = dispatchRuleName;
            DispatchType = dispatchType;
            GroupRules = groupRules;
            Id = id;
            LabelMatchExpressionGrids = labelMatchExpressionGrids;
            NotifyRules = notifyRules;
            Status = status;
        }
    }
}