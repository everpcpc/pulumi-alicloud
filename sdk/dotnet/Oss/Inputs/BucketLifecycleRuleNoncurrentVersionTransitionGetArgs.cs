// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oss.Inputs
{

    public sealed class BucketLifecycleRuleNoncurrentVersionTransitionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the number of days after object creation when the specific rule action takes effect.
        /// </summary>
        [Input("days", required: true)]
        public Input<int> Days { get; set; } = null!;

        /// <summary>
        /// The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be "Standard", "IA", "Archive" and "ColdArchive". Defaults to "Standard". "ColdArchive" is available in 1.203.0+.
        /// </summary>
        [Input("storageClass", required: true)]
        public Input<string> StorageClass { get; set; } = null!;

        public BucketLifecycleRuleNoncurrentVersionTransitionGetArgs()
        {
        }
        public static new BucketLifecycleRuleNoncurrentVersionTransitionGetArgs Empty => new BucketLifecycleRuleNoncurrentVersionTransitionGetArgs();
    }
}
