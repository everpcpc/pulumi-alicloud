// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ResourceManager.Outputs
{

    [OutputType]
    public sealed class GetAccountsAccountResult
    {
        public readonly string AccountId;
        public readonly string DisplayName;
        /// <summary>
        /// The ID of the folder.
        /// </summary>
        public readonly string FolderId;
        /// <summary>
        /// The ID of the resource.
        /// * `account_id`- The ID of the account.
        /// * `display_name`- The name of the member account.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The way in which the member account joined the resource directory. 
        /// </summary>
        public readonly string JoinMethod;
        /// <summary>
        /// The time when the member account joined the resource directory.
        /// </summary>
        public readonly string JoinTime;
        /// <summary>
        /// The time when the member account was modified.
        /// </summary>
        public readonly string ModifyTime;
        /// <summary>
        /// The ID of the resource directory.
        /// </summary>
        public readonly string ResourceDirectoryId;
        /// <summary>
        /// The status of the member account. 
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The type of the member account. 
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAccountsAccountResult(
            string accountId,

            string displayName,

            string folderId,

            string id,

            string joinMethod,

            string joinTime,

            string modifyTime,

            string resourceDirectoryId,

            string status,

            string type)
        {
            AccountId = accountId;
            DisplayName = displayName;
            FolderId = folderId;
            Id = id;
            JoinMethod = joinMethod;
            JoinTime = joinTime;
            ModifyTime = modifyTime;
            ResourceDirectoryId = resourceDirectoryId;
            Status = status;
            Type = type;
        }
    }
}
