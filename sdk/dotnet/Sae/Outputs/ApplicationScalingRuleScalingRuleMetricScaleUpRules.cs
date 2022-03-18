// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Outputs
{

    [OutputType]
    public sealed class ApplicationScalingRuleScalingRuleMetricScaleUpRules
    {
        /// <summary>
        /// Whether shrinkage is prohibited.
        /// </summary>
        public readonly bool? Disabled;
        /// <summary>
        /// Cooling time for expansion or contraction. Valid values: `0` to `3600`. Unit: seconds. The default is `0` seconds.
        /// </summary>
        public readonly int? StabilizationWindowSeconds;
        /// <summary>
        /// Elastic expansion or contraction step size. the maximum number of instances to be scaled in per unit time.
        /// </summary>
        public readonly int? Step;

        [OutputConstructor]
        private ApplicationScalingRuleScalingRuleMetricScaleUpRules(
            bool? disabled,

            int? stabilizationWindowSeconds,

            int? step)
        {
            Disabled = disabled;
            StabilizationWindowSeconds = stabilizationWindowSeconds;
            Step = step;
        }
    }
}
