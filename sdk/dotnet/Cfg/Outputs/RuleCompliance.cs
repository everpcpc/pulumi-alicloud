// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cfg.Outputs
{

    [OutputType]
    public sealed class RuleCompliance
    {
        /// <summary>
        /// The type of compliance. Valid values: `COMPLIANT`, `NON_COMPLIANT`, `NOT_APPLICABLE`, `INSUFFICIENT_DATA`.
        /// </summary>
        public readonly string? ComplianceType;
        /// <summary>
        /// The count of compliance.
        /// </summary>
        public readonly int? Count;

        [OutputConstructor]
        private RuleCompliance(
            string? complianceType,

            int? count)
        {
            ComplianceType = complianceType;
            Count = count;
        }
    }
}