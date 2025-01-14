// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oss.Inputs
{

    public sealed class BucketVersioningGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the versioning state of a bucket. Valid values: `Enabled` and `Suspended`.
        /// 
        /// `NOTE`: Currently, the `versioning` feature is only available in ap-south-1 and with white list. If you want to use it, please contact us.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        public BucketVersioningGetArgs()
        {
        }
        public static new BucketVersioningGetArgs Empty => new BucketVersioningGetArgs();
    }
}
