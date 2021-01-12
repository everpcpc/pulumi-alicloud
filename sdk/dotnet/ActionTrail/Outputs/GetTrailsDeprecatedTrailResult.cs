// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ActionTrail.Outputs
{

    [OutputType]
    public sealed class GetTrailsDeprecatedTrailResult
    {
        /// <summary>
        /// Indicates whether the event is a read or a write event.
        /// </summary>
        public readonly string EventRw;
        public readonly string Id;
        public readonly string MnsTopicArn;
        /// <summary>
        /// The name of the specified OSS bucket.
        /// </summary>
        public readonly string OssBucketName;
        /// <summary>
        /// The prefix of the specified OSS bucket name.
        /// </summary>
        public readonly string OssKeyPrefix;
        /// <summary>
        /// The role in ActionTrail.
        /// </summary>
        public readonly string RoleName;
        /// <summary>
        /// The unique ARN of the Log Service project.
        /// </summary>
        public readonly string SlsProjectArn;
        /// <summary>
        /// The unique ARN of the Log Service role.
        /// </summary>
        public readonly string SlsWriteRoleArn;
        public readonly string Status;
        public readonly string TrailName;
        public readonly string TrailRegion;

        [OutputConstructor]
        private GetTrailsDeprecatedTrailResult(
            string eventRw,

            string id,

            string mnsTopicArn,

            string ossBucketName,

            string ossKeyPrefix,

            string roleName,

            string slsProjectArn,

            string slsWriteRoleArn,

            string status,

            string trailName,

            string trailRegion)
        {
            EventRw = eventRw;
            Id = id;
            MnsTopicArn = mnsTopicArn;
            OssBucketName = ossBucketName;
            OssKeyPrefix = ossKeyPrefix;
            RoleName = roleName;
            SlsProjectArn = slsProjectArn;
            SlsWriteRoleArn = slsWriteRoleArn;
            Status = status;
            TrailName = trailName;
            TrailRegion = trailRegion;
        }
    }
}
