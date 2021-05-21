// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Log.Outputs
{

    [OutputType]
    public sealed class EtlEtlSink
    {
        /// <summary>
        /// Dekms_encryption_access_key_id_contextlivery target logstore access key id.
        /// </summary>
        public readonly string? AccessKeyId;
        /// <summary>
        /// Delivery target logstore access key secret.
        /// </summary>
        public readonly string? AccessKeySecret;
        /// <summary>
        /// Delivery target logstore region.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// An KMS encrypts access key id used to a log etl job. If the `access_key_id` is filled in, this field will be ignored.
        /// </summary>
        public readonly string? KmsEncryptedAccessKeyId;
        /// <summary>
        /// An KMS encrypts access key secret used to a log etl job. If the `access_key_secret` is filled in, this field will be ignored.
        /// </summary>
        public readonly string? KmsEncryptedAccessKeySecret;
        /// <summary>
        /// Delivery target logstore.
        /// </summary>
        public readonly string Logstore;
        /// <summary>
        /// Delivery target name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The project where the target logstore is delivered.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// Sts role info.
        /// </summary>
        public readonly string? RoleArn;
        /// <summary>
        /// ETL sinks type, the default value is AliyunLOG.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private EtlEtlSink(
            string? accessKeyId,

            string? accessKeySecret,

            string endpoint,

            string? kmsEncryptedAccessKeyId,

            string? kmsEncryptedAccessKeySecret,

            string logstore,

            string name,

            string project,

            string? roleArn,

            string? type)
        {
            AccessKeyId = accessKeyId;
            AccessKeySecret = accessKeySecret;
            Endpoint = endpoint;
            KmsEncryptedAccessKeyId = kmsEncryptedAccessKeyId;
            KmsEncryptedAccessKeySecret = kmsEncryptedAccessKeySecret;
            Logstore = logstore;
            Name = name;
            Project = project;
            RoleArn = roleArn;
            Type = type;
        }
    }
}