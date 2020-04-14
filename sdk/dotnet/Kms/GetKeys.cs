// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Kms
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides a list of KMS keys in an Alibaba Cloud account according to the specified filters.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/kms_keys.html.markdown.
        /// </summary>
        [Obsolete("Use GetKeys.InvokeAsync() instead")]
        public static Task<GetKeysResult> GetKeys(GetKeysArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeysResult>("alicloud:kms/getKeys:getKeys", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetKeys
    {
        /// <summary>
        /// This data source provides a list of KMS keys in an Alibaba Cloud account according to the specified filters.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/kms_keys.html.markdown.
        /// </summary>
        public static Task<GetKeysResult> InvokeAsync(GetKeysArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeysResult>("alicloud:kms/getKeys:getKeys", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A regex string to filter the results by the KMS key description.
        /// </summary>
        [Input("descriptionRegex")]
        public string? DescriptionRegex { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of KMS key IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Filter the results by status of the KMS keys. Valid values: `Enabled`, `Disabled`, `PendingDeletion`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetKeysArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetKeysResult
    {
        public readonly string? DescriptionRegex;
        /// <summary>
        /// A list of KMS key IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of KMS keys. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKeysKeysResult> Keys;
        public readonly string? OutputFile;
        /// <summary>
        /// Status of the key. Possible values: `Enabled`, `Disabled` and `PendingDeletion`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetKeysResult(
            string? descriptionRegex,
            ImmutableArray<string> ids,
            ImmutableArray<Outputs.GetKeysKeysResult> keys,
            string? outputFile,
            string? status,
            string id)
        {
            DescriptionRegex = descriptionRegex;
            Ids = ids;
            Keys = keys;
            OutputFile = outputFile;
            Status = status;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetKeysKeysResult
    {
        /// <summary>
        /// The Alibaba Cloud Resource Name (ARN) of the key.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Creation date of key.
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// The owner of the key.
        /// </summary>
        public readonly string Creator;
        /// <summary>
        /// Deletion date of key.
        /// </summary>
        public readonly string DeleteDate;
        /// <summary>
        /// Description of the key.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// ID of the key.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Filter the results by status of the KMS keys. Valid values: `Enabled`, `Disabled`, `PendingDeletion`.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetKeysKeysResult(
            string arn,
            string creationDate,
            string creator,
            string deleteDate,
            string description,
            string id,
            string status)
        {
            Arn = arn;
            CreationDate = creationDate;
            Creator = creator;
            DeleteDate = deleteDate;
            Description = description;
            Id = id;
            Status = status;
        }
    }
    }
}
