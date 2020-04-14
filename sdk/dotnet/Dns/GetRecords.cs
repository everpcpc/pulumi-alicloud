// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dns
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides a list of DNS Domain Records in an Alibaba Cloud account according to the specified filters.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/dns_records.html.markdown.
        /// </summary>
        [Obsolete("Use GetRecords.InvokeAsync() instead")]
        public static Task<GetRecordsResult> GetRecords(GetRecordsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRecordsResult>("alicloud:dns/getRecords:getRecords", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetRecords
    {
        /// <summary>
        /// This data source provides a list of DNS Domain Records in an Alibaba Cloud account according to the specified filters.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/dns_records.html.markdown.
        /// </summary>
        public static Task<GetRecordsResult> InvokeAsync(GetRecordsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRecordsResult>("alicloud:dns/getRecords:getRecords", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetRecordsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain name associated to the records.
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        /// <summary>
        /// Host record regex. 
        /// </summary>
        [Input("hostRecordRegex")]
        public string? HostRecordRegex { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of record IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Whether the record is locked or not.
        /// </summary>
        [Input("isLocked")]
        public bool? IsLocked { get; set; }

        /// <summary>
        /// ISP line. Valid items are `default`, `telecom`, `unicom`, `mobile`, `oversea`, `edu`, `drpeng`, `btvn`, .etc. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm) 
        /// </summary>
        [Input("line")]
        public string? Line { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Record status. Valid items are `ENABLE` and `DISABLE`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// Record type. Valid items are `A`, `NS`, `MX`, `TXT`, `CNAME`, `SRV`, `AAAA`, `REDIRECT_URL`, `FORWORD_URL` .
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        /// <summary>
        /// Host record value regex. 
        /// </summary>
        [Input("valueRegex")]
        public string? ValueRegex { get; set; }

        public GetRecordsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetRecordsResult
    {
        /// <summary>
        /// Name of the domain the record belongs to.
        /// </summary>
        public readonly string DomainName;
        public readonly string? HostRecordRegex;
        /// <summary>
        /// A list of record IDs. 
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly bool? IsLocked;
        /// <summary>
        /// ISP line of the record. 
        /// </summary>
        public readonly string? Line;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of records. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRecordsRecordsResult> Records;
        /// <summary>
        /// Status of the record.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Type of the record.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// A list of entire URLs. Each item format as `&lt;host_record&gt;.&lt;domain_name&gt;`.
        /// </summary>
        public readonly ImmutableArray<string> Urls;
        public readonly string? ValueRegex;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetRecordsResult(
            string domainName,
            string? hostRecordRegex,
            ImmutableArray<string> ids,
            bool? isLocked,
            string? line,
            string? outputFile,
            ImmutableArray<Outputs.GetRecordsRecordsResult> records,
            string? status,
            string? type,
            ImmutableArray<string> urls,
            string? valueRegex,
            string id)
        {
            DomainName = domainName;
            HostRecordRegex = hostRecordRegex;
            Ids = ids;
            IsLocked = isLocked;
            Line = line;
            OutputFile = outputFile;
            Records = records;
            Status = status;
            Type = type;
            Urls = urls;
            ValueRegex = valueRegex;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetRecordsRecordsResult
    {
        /// <summary>
        /// The domain name associated to the records.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// Host record of the domain.
        /// </summary>
        public readonly string HostRecord;
        /// <summary>
        /// ISP line. Valid items are `default`, `telecom`, `unicom`, `mobile`, `oversea`, `edu`, `drpeng`, `btvn`, .etc. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm) 
        /// </summary>
        public readonly string Line;
        /// <summary>
        /// Indicates whether the record is locked.
        /// </summary>
        public readonly bool Locked;
        /// <summary>
        /// Priority of the `MX` record.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// ID of the record.
        /// </summary>
        public readonly string RecordId;
        /// <summary>
        /// Record status. Valid items are `ENABLE` and `DISABLE`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// TTL of the record.
        /// </summary>
        public readonly double Ttl;
        /// <summary>
        /// Record type. Valid items are `A`, `NS`, `MX`, `TXT`, `CNAME`, `SRV`, `AAAA`, `REDIRECT_URL`, `FORWORD_URL` .
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Host record value of the domain.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetRecordsRecordsResult(
            string domainName,
            string hostRecord,
            string line,
            bool locked,
            int priority,
            string recordId,
            string status,
            double ttl,
            string type,
            string value)
        {
            DomainName = domainName;
            HostRecord = hostRecord;
            Line = line;
            Locked = locked;
            Priority = priority;
            RecordId = recordId;
            Status = status;
            Ttl = ttl;
            Type = type;
            Value = value;
        }
    }
    }
}
