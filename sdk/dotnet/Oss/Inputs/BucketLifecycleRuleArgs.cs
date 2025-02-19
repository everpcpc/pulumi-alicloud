// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oss.Inputs
{

    public sealed class BucketLifecycleRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("abortMultipartUploads")]
        private InputList<Inputs.BucketLifecycleRuleAbortMultipartUploadArgs>? _abortMultipartUploads;

        /// <summary>
        /// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed (documented below).
        /// </summary>
        public InputList<Inputs.BucketLifecycleRuleAbortMultipartUploadArgs> AbortMultipartUploads
        {
            get => _abortMultipartUploads ?? (_abortMultipartUploads = new InputList<Inputs.BucketLifecycleRuleAbortMultipartUploadArgs>());
            set => _abortMultipartUploads = value;
        }

        /// <summary>
        /// Specifies lifecycle rule status.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("expirations")]
        private InputList<Inputs.BucketLifecycleRuleExpirationArgs>? _expirations;

        /// <summary>
        /// Specifies a period in the object's expire (documented below).
        /// </summary>
        public InputList<Inputs.BucketLifecycleRuleExpirationArgs> Expirations
        {
            get => _expirations ?? (_expirations = new InputList<Inputs.BucketLifecycleRuleExpirationArgs>());
            set => _expirations = value;
        }

        /// <summary>
        /// Unique identifier for the rule. If omitted, OSS bucket will assign a unique name.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("noncurrentVersionExpirations")]
        private InputList<Inputs.BucketLifecycleRuleNoncurrentVersionExpirationArgs>? _noncurrentVersionExpirations;

        /// <summary>
        /// Specifies when noncurrent object versions expire (documented below).
        /// </summary>
        public InputList<Inputs.BucketLifecycleRuleNoncurrentVersionExpirationArgs> NoncurrentVersionExpirations
        {
            get => _noncurrentVersionExpirations ?? (_noncurrentVersionExpirations = new InputList<Inputs.BucketLifecycleRuleNoncurrentVersionExpirationArgs>());
            set => _noncurrentVersionExpirations = value;
        }

        [Input("noncurrentVersionTransitions")]
        private InputList<Inputs.BucketLifecycleRuleNoncurrentVersionTransitionArgs>? _noncurrentVersionTransitions;

        /// <summary>
        /// Specifies when noncurrent object versions transitions (documented below).
        /// 
        /// `NOTE`: At least one of expiration, transitions, abort_multipart_upload, noncurrent_version_expiration and noncurrent_version_transition should be configured.
        /// </summary>
        public InputList<Inputs.BucketLifecycleRuleNoncurrentVersionTransitionArgs> NoncurrentVersionTransitions
        {
            get => _noncurrentVersionTransitions ?? (_noncurrentVersionTransitions = new InputList<Inputs.BucketLifecycleRuleNoncurrentVersionTransitionArgs>());
            set => _noncurrentVersionTransitions = value;
        }

        /// <summary>
        /// Object key prefix identifying one or more objects to which the rule applies. Default value is null, the rule applies to all objects in a bucket.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("transitions")]
        private InputList<Inputs.BucketLifecycleRuleTransitionArgs>? _transitions;

        /// <summary>
        /// Specifies the time when an object is converted to the IA or archive storage class during a valid life cycle. (documented below).
        /// </summary>
        public InputList<Inputs.BucketLifecycleRuleTransitionArgs> Transitions
        {
            get => _transitions ?? (_transitions = new InputList<Inputs.BucketLifecycleRuleTransitionArgs>());
            set => _transitions = value;
        }

        public BucketLifecycleRuleArgs()
        {
        }
        public static new BucketLifecycleRuleArgs Empty => new BucketLifecycleRuleArgs();
    }
}
