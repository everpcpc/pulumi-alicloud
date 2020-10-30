// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Inputs
{

    public sealed class EdgeKubernetesWorkerDataDiskArgs : Pulumi.ResourceArgs
    {
        [Input("autoSnapshotPolicyId")]
        public Input<string>? AutoSnapshotPolicyId { get; set; }

        [Input("category")]
        public Input<string>? Category { get; set; }

        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("encrypted")]
        public Input<string>? Encrypted { get; set; }

        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// The kubernetes cluster's name. It is unique in one Alicloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("size")]
        public Input<string>? Size { get; set; }

        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        public EdgeKubernetesWorkerDataDiskArgs()
        {
        }
    }
}
