// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oss.Inputs
{

    public sealed class BucketLifecycleRuleExpirationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the time before which the rules take effect. The date must conform to the ISO8601 format and always be UTC 00:00. For example: 2002-10-11T00:00:00.000Z indicates that objects updated before 2002-10-11T00:00:00.000Z are deleted or converted to another storage class, and objects updated after this time (including this time) are not deleted or converted.
        /// </summary>
        [Input("createdBeforeDate")]
        public Input<string>? CreatedBeforeDate { get; set; }

        /// <summary>
        /// Specifies the date after which you want the corresponding action to take effect. The value obeys ISO8601 format like `2017-03-09`.
        /// </summary>
        [Input("date")]
        public Input<string>? Date { get; set; }

        /// <summary>
        /// Specifies the number of days after object creation when the specific rule action takes effect.
        /// 
        /// `NOTE`: One and only one of "created_before_date" and "days" can be specified in one abort_multipart_upload configuration.
        /// </summary>
        [Input("days")]
        public Input<int>? Days { get; set; }

        /// <summary>
        /// On a versioned bucket (versioning-enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct OSS to delete expired object delete markers. This cannot be specified with Days, Date or CreatedBeforeDate in a Lifecycle Expiration Policy.
        /// 
        /// `NOTE`: One and only one of "date", "days", "created_before_date" and "expired_object_delete_marker" can be specified in one expiration configuration.
        /// </summary>
        [Input("expiredObjectDeleteMarker")]
        public Input<bool>? ExpiredObjectDeleteMarker { get; set; }

        public BucketLifecycleRuleExpirationGetArgs()
        {
        }
        public static new BucketLifecycleRuleExpirationGetArgs Empty => new BucketLifecycleRuleExpirationGetArgs();
    }
}
