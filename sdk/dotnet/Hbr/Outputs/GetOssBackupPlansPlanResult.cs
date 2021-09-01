// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hbr.Outputs
{

    [OutputType]
    public sealed class GetOssBackupPlansPlanResult
    {
        /// <summary>
        /// Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
        /// </summary>
        public readonly string BackupType;
        /// <summary>
        /// The OSS Bucket Name.
        /// </summary>
        public readonly string Bucket;
        public readonly bool Disabled;
        public readonly string Id;
        public readonly string OssBackupPlanId;
        /// <summary>
        /// The Configuration Page of a Backup Plan Name. 1-64 Characters, requiring a Single Warehouse under Each of the Data Source Type Drop-down List of the Configuration Page of a Backup Plan Name Is Unique.
        /// </summary>
        public readonly string OssBackupPlanName;
        public readonly string Prefix;
        /// <summary>
        /// Backup Retention Period, the Minimum Value of 1.
        /// </summary>
        public readonly string Retention;
        /// <summary>
        /// Backup strategy. Optional format: I|{startTime}|{interval} * startTime Backup start time, UNIX time, in seconds. * interval ISO8601 time interval. E.g: ** PT1H, one hour apart. ** P1D, one day apart. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed, the next backup task will not be triggered.
        /// </summary>
        public readonly string Schedule;
        /// <summary>
        /// Vault ID.
        /// </summary>
        public readonly string VaultId;

        [OutputConstructor]
        private GetOssBackupPlansPlanResult(
            string backupType,

            string bucket,

            bool disabled,

            string id,

            string ossBackupPlanId,

            string ossBackupPlanName,

            string prefix,

            string retention,

            string schedule,

            string vaultId)
        {
            BackupType = backupType;
            Bucket = bucket;
            Disabled = disabled;
            Id = id;
            OssBackupPlanId = ossBackupPlanId;
            OssBackupPlanName = ossBackupPlanName;
            Prefix = prefix;
            Retention = retention;
            Schedule = schedule;
            VaultId = vaultId;
        }
    }
}