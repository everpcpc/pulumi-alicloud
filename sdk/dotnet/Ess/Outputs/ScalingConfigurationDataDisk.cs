// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ess.Outputs
{

    [OutputType]
    public sealed class ScalingConfigurationDataDisk
    {
        public readonly string? AutoSnapshotPolicyId;
        public readonly string? Category;
        public readonly bool? DeleteWithInstance;
        public readonly string? Description;
        public readonly string? Device;
        public readonly bool? Encrypted;
        public readonly string? KmsKeyId;
        public readonly string? Name;
        public readonly string? PerformanceLevel;
        public readonly int? Size;
        public readonly string? SnapshotId;

        [OutputConstructor]
        private ScalingConfigurationDataDisk(
            string? autoSnapshotPolicyId,

            string? category,

            bool? deleteWithInstance,

            string? description,

            string? device,

            bool? encrypted,

            string? kmsKeyId,

            string? name,

            string? performanceLevel,

            int? size,

            string? snapshotId)
        {
            AutoSnapshotPolicyId = autoSnapshotPolicyId;
            Category = category;
            DeleteWithInstance = deleteWithInstance;
            Description = description;
            Device = device;
            Encrypted = encrypted;
            KmsKeyId = kmsKeyId;
            Name = name;
            PerformanceLevel = performanceLevel;
            Size = size;
            SnapshotId = snapshotId;
        }
    }
}
