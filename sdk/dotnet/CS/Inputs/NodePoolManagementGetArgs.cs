// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Inputs
{

    public sealed class NodePoolManagementGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether automatic repair, Default to `false`.
        /// </summary>
        [Input("autoRepair")]
        public Input<bool>? AutoRepair { get; set; }

        /// <summary>
        /// Whether auto upgrade, Default to `false`.
        /// </summary>
        [Input("autoUpgrade")]
        public Input<bool>? AutoUpgrade { get; set; }

        /// <summary>
        /// Max number of unavailable nodes. Default to `1`.
        /// </summary>
        [Input("maxUnavailable", required: true)]
        public Input<int> MaxUnavailable { get; set; } = null!;

        /// <summary>
        /// Number of additional nodes. You have to specify one of surge, surge_percentage.
        /// </summary>
        [Input("surge")]
        public Input<int>? Surge { get; set; }

        /// <summary>
        /// Proportion of additional nodes. You have to specify one of surge, surge_percentage.
        /// </summary>
        [Input("surgePercentage")]
        public Input<int>? SurgePercentage { get; set; }

        public NodePoolManagementGetArgs()
        {
        }
    }
}
