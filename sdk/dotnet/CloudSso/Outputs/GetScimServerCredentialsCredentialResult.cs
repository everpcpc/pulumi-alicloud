// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CloudSso.Outputs
{

    [OutputType]
    public sealed class GetScimServerCredentialsCredentialResult
    {
        /// <summary>
        /// The CreateTime of the resource.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The CredentialId of the resource.
        /// </summary>
        public readonly string CredentialId;
        /// <summary>
        /// The CredentialSecret of the resource.
        /// </summary>
        public readonly string CredentialSecret;
        /// <summary>
        /// The CredentialType of the resource.
        /// </summary>
        public readonly string CredentialType;
        /// <summary>
        /// The ID of the Directory.
        /// </summary>
        public readonly string DirectoryId;
        /// <summary>
        /// The ExpireTime of the resource.
        /// </summary>
        public readonly string ExpireTime;
        /// <summary>
        /// The ID of the SCIM Server Credential.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Status of the resource. Valid values: `Disabled`, `Enabled`.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetScimServerCredentialsCredentialResult(
            string createTime,

            string credentialId,

            string credentialSecret,

            string credentialType,

            string directoryId,

            string expireTime,

            string id,

            string status)
        {
            CreateTime = createTime;
            CredentialId = credentialId;
            CredentialSecret = credentialSecret;
            CredentialType = credentialType;
            DirectoryId = directoryId;
            ExpireTime = expireTime;
            Id = id;
            Status = status;
        }
    }
}