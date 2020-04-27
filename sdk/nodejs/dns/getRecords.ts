// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of DNS Domain Records in an Alibaba Cloud account according to the specified filters.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const recordsDs = pulumi.output(alicloud.dns.getRecords({
 *     domainName: "xiaozhu.top",
 *     hostRecordRegex: "^@",
 *     isLocked: false,
 *     outputFile: "records.txt",
 *     type: "A",
 * }, { async: true }));
 * 
 * export const firstRecordId = recordsDs.records[0].recordId;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/dns_records.html.markdown.
 */
export function getRecords(args: GetRecordsArgs, opts?: pulumi.InvokeOptions): Promise<GetRecordsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:dns/getRecords:getRecords", {
        "domainName": args.domainName,
        "hostRecordRegex": args.hostRecordRegex,
        "ids": args.ids,
        "isLocked": args.isLocked,
        "line": args.line,
        "outputFile": args.outputFile,
        "status": args.status,
        "type": args.type,
        "valueRegex": args.valueRegex,
    }, opts);
}

/**
 * A collection of arguments for invoking getRecords.
 */
export interface GetRecordsArgs {
    /**
     * The domain name associated to the records.
     */
    readonly domainName: string;
    /**
     * Host record regex. 
     */
    readonly hostRecordRegex?: string;
    /**
     * A list of record IDs.
     */
    readonly ids?: string[];
    /**
     * Whether the record is locked or not.
     */
    readonly isLocked?: boolean;
    /**
     * ISP line. Valid items are `default`, `telecom`, `unicom`, `mobile`, `oversea`, `edu`, `drpeng`, `btvn`, .etc. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm) 
     */
    readonly line?: string;
    readonly outputFile?: string;
    /**
     * Record status. Valid items are `ENABLE` and `DISABLE`.
     */
    readonly status?: string;
    /**
     * Record type. Valid items are `A`, `NS`, `MX`, `TXT`, `CNAME`, `SRV`, `AAAA`, `REDIRECT_URL`, `FORWORD_URL` .
     */
    readonly type?: string;
    /**
     * Host record value regex. 
     */
    readonly valueRegex?: string;
}

/**
 * A collection of values returned by getRecords.
 */
export interface GetRecordsResult {
    /**
     * Name of the domain the record belongs to.
     */
    readonly domainName: string;
    readonly hostRecordRegex?: string;
    /**
     * A list of record IDs. 
     */
    readonly ids: string[];
    readonly isLocked?: boolean;
    /**
     * ISP line of the record. 
     */
    readonly line?: string;
    readonly outputFile?: string;
    /**
     * A list of records. Each element contains the following attributes:
     */
    readonly records: outputs.dns.GetRecordsRecord[];
    /**
     * Status of the record.
     */
    readonly status?: string;
    /**
     * Type of the record.
     */
    readonly type?: string;
    /**
     * A list of entire URLs. Each item format as `<host_record>.<domain_name>`.
     */
    readonly urls: string[];
    readonly valueRegex?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
