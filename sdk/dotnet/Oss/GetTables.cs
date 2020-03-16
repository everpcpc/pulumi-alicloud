// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oss
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides the ots tables of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.40.0+.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ots_tables.html.markdown.
        /// </summary>
        public static Task<GetTablesResult> GetTables(GetTablesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTablesResult>("alicloud:oss/getTables:getTables", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetTablesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of table IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of OTS instance.
        /// </summary>
        [Input("instanceName", required: true)]
        public string InstanceName { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by table name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetTablesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetTablesResult
    {
        /// <summary>
        /// A list of table IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The OTS instance name.
        /// </summary>
        public readonly string InstanceName;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of table names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of tables. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTablesTablesResult> Tables;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetTablesResult(
            ImmutableArray<string> ids,
            string instanceName,
            string? nameRegex,
            ImmutableArray<string> names,
            string? outputFile,
            ImmutableArray<Outputs.GetTablesTablesResult> tables,
            string id)
        {
            Ids = ids;
            InstanceName = instanceName;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Tables = tables;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetTablesTablesPrimaryKeysResult
    {
        public readonly string Name;
        public readonly string Type;

        [OutputConstructor]
        private GetTablesTablesPrimaryKeysResult(
            string name,
            string type)
        {
            Name = name;
            Type = type;
        }
    }

    [OutputType]
    public sealed class GetTablesTablesResult
    {
        /// <summary>
        /// ID of the table. The value is `&lt;instance_name&gt;:&lt;table_name&gt;`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of OTS instance.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// The maximum number of versions stored in this table.
        /// </summary>
        public readonly int MaxVersion;
        /// <summary>
        /// The property of `TableMeta` which indicates the structure information of a table.
        /// </summary>
        public readonly ImmutableArray<GetTablesTablesPrimaryKeysResult> PrimaryKeys;
        /// <summary>
        /// The table name of the OTS which could not be changed.
        /// </summary>
        public readonly string TableName;
        /// <summary>
        /// The retention time of data stored in this table.
        /// </summary>
        public readonly int TimeToLive;

        [OutputConstructor]
        private GetTablesTablesResult(
            string id,
            string instanceName,
            int maxVersion,
            ImmutableArray<GetTablesTablesPrimaryKeysResult> primaryKeys,
            string tableName,
            int timeToLive)
        {
            Id = id;
            InstanceName = instanceName;
            MaxVersion = maxVersion;
            PrimaryKeys = primaryKeys;
            TableName = tableName;
            TimeToLive = timeToLive;
        }
    }
    }
}
