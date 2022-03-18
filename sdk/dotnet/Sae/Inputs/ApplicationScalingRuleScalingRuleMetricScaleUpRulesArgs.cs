// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Inputs
{

    public sealed class ApplicationScalingRuleScalingRuleMetricScaleUpRulesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether shrinkage is prohibited.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Cooling time for expansion or contraction. Valid values: `0` to `3600`. Unit: seconds. The default is `0` seconds.
        /// </summary>
        [Input("stabilizationWindowSeconds")]
        public Input<int>? StabilizationWindowSeconds { get; set; }

        /// <summary>
        /// Elastic expansion or contraction step size. the maximum number of instances to be scaled in per unit time.
        /// </summary>
        [Input("step")]
        public Input<int>? Step { get; set; }

        public ApplicationScalingRuleScalingRuleMetricScaleUpRulesArgs()
        {
        }
    }
}
