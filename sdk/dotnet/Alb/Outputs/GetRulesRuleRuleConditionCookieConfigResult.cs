// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Outputs
{

    [OutputType]
    public sealed class GetRulesRuleRuleConditionCookieConfigResult
    {
        /// <summary>
        /// Add one or more IP addresses or IP address segments.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRulesRuleRuleConditionCookieConfigValueResult> Values;

        [OutputConstructor]
        private GetRulesRuleRuleConditionCookieConfigResult(ImmutableArray<Outputs.GetRulesRuleRuleConditionCookieConfigValueResult> values)
        {
            Values = values;
        }
    }
}
