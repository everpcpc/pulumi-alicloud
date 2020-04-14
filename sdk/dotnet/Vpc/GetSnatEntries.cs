// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides a list of Snat Entries owned by an Alibaba Cloud account.
        /// 
        /// &gt; **NOTE:** Available in 1.37.0+.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/snat_entries.html.markdown.
        /// </summary>
        [Obsolete("Use GetSnatEntries.InvokeAsync() instead")]
        public static Task<GetSnatEntriesResult> GetSnatEntries(GetSnatEntriesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSnatEntriesResult>("alicloud:vpc/getSnatEntries:getSnatEntries", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetSnatEntries
    {
        /// <summary>
        /// This data source provides a list of Snat Entries owned by an Alibaba Cloud account.
        /// 
        /// &gt; **NOTE:** Available in 1.37.0+.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/snat_entries.html.markdown.
        /// </summary>
        public static Task<GetSnatEntriesResult> InvokeAsync(GetSnatEntriesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSnatEntriesResult>("alicloud:vpc/getSnatEntries:getSnatEntries", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetSnatEntriesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Snat Entries IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The public IP of the Snat Entry.
        /// </summary>
        [Input("snatIp")]
        public string? SnatIp { get; set; }

        /// <summary>
        /// The ID of the Snat table.
        /// </summary>
        [Input("snatTableId", required: true)]
        public string SnatTableId { get; set; } = null!;

        /// <summary>
        /// The source CIDR block of the Snat Entry.
        /// </summary>
        [Input("sourceCidr")]
        public string? SourceCidr { get; set; }

        public GetSnatEntriesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSnatEntriesResult
    {
        /// <summary>
        /// A list of Snat Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSnatEntriesEntriesResult> Entries;
        /// <summary>
        /// (Optional) A list of Snat Entries IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        /// <summary>
        /// The public IP of the Snat Entry.
        /// </summary>
        public readonly string? SnatIp;
        public readonly string SnatTableId;
        /// <summary>
        /// The source CIDR block of the Snat Entry.
        /// </summary>
        public readonly string? SourceCidr;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSnatEntriesResult(
            ImmutableArray<Outputs.GetSnatEntriesEntriesResult> entries,
            ImmutableArray<string> ids,
            string? outputFile,
            string? snatIp,
            string snatTableId,
            string? sourceCidr,
            string id)
        {
            Entries = entries;
            Ids = ids;
            OutputFile = outputFile;
            SnatIp = snatIp;
            SnatTableId = snatTableId;
            SourceCidr = sourceCidr;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetSnatEntriesEntriesResult
    {
        /// <summary>
        /// The ID of the Snat Entry.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The public IP of the Snat Entry.
        /// </summary>
        public readonly string SnatIp;
        /// <summary>
        /// The source CIDR block of the Snat Entry.
        /// </summary>
        public readonly string SourceCidr;
        /// <summary>
        /// The status of the Snat Entry.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetSnatEntriesEntriesResult(
            string id,
            string snatIp,
            string sourceCidr,
            string status)
        {
            Id = id;
            SnatIp = snatIp;
            SourceCidr = sourceCidr;
            Status = status;
        }
    }
    }
}
