// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Outputs
{

    [OutputType]
    public sealed class NodePoolManagement
    {
        /// <summary>
        /// Whether automatic repair, Default to `false`.
        /// </summary>
        public readonly bool? AutoRepair;
        /// <summary>
        /// Whether auto upgrade, Default to `false`.
        /// </summary>
        public readonly bool? AutoUpgrade;
        /// <summary>
        /// Max number of unavailable nodes. Default to `1`.
        /// </summary>
        public readonly int MaxUnavailable;
        /// <summary>
        /// Number of additional nodes. You have to specify one of surge, surge_percentage.
        /// </summary>
        public readonly int? Surge;
        /// <summary>
        /// Proportion of additional nodes. You have to specify one of surge, surge_percentage.
        /// </summary>
        public readonly int? SurgePercentage;

        [OutputConstructor]
        private NodePoolManagement(
            bool? autoRepair,

            bool? autoUpgrade,

            int maxUnavailable,

            int? surge,

            int? surgePercentage)
        {
            AutoRepair = autoRepair;
            AutoUpgrade = autoUpgrade;
            MaxUnavailable = maxUnavailable;
            Surge = surge;
            SurgePercentage = surgePercentage;
        }
    }
}