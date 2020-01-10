// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Rds
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides the RDS instance engines resource available info of Alibaba Cloud.
        /// 
        /// &gt; **NOTE:** Available in v1.46.0+
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/db_instance_engines.html.markdown.
        /// </summary>
        public static Task<GetInstanceEnginesResult> GetInstanceEngines(GetInstanceEnginesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceEnginesResult>("alicloud:rds/getInstanceEngines:getInstanceEngines", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetInstanceEnginesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL` and `PPAS`. If no value is specified, all types are returned.
        /// </summary>
        [Input("engine")]
        public string? Engine { get; set; }

        /// <summary>
        /// Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
        /// </summary>
        [Input("engineVersion")]
        public string? EngineVersion { get; set; }

        /// <summary>
        /// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
        /// </summary>
        [Input("instanceChargeType")]
        public string? InstanceChargeType { get; set; }

        /// <summary>
        /// Whether to show multi available zone. Default false to not show multi availability zone.
        /// </summary>
        [Input("multiZone")]
        public bool? MultiZone { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The Zone to launch the DB instance.
        /// </summary>
        [Input("zoneId")]
        public string? ZoneId { get; set; }

        public GetInstanceEnginesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetInstanceEnginesResult
    {
        /// <summary>
        /// Database type.
        /// </summary>
        public readonly string? Engine;
        /// <summary>
        /// DB Instance version.
        /// </summary>
        public readonly string? EngineVersion;
        public readonly string? InstanceChargeType;
        /// <summary>
        /// A list of Rds available resource. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceEnginesInstanceEnginesResult> InstanceEngines;
        public readonly bool? MultiZone;
        public readonly string? OutputFile;
        public readonly string? ZoneId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetInstanceEnginesResult(
            string? engine,
            string? engineVersion,
            string? instanceChargeType,
            ImmutableArray<Outputs.GetInstanceEnginesInstanceEnginesResult> instanceEngines,
            bool? multiZone,
            string? outputFile,
            string? zoneId,
            string id)
        {
            Engine = engine;
            EngineVersion = engineVersion;
            InstanceChargeType = instanceChargeType;
            InstanceEngines = instanceEngines;
            MultiZone = multiZone;
            OutputFile = outputFile;
            ZoneId = zoneId;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetInstanceEnginesInstanceEnginesResult
    {
        /// <summary>
        /// DB Instance category.
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL` and `PPAS`. If no value is specified, all types are returned.
        /// </summary>
        public readonly string Engine;
        /// <summary>
        /// Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
        /// </summary>
        public readonly string EngineVersion;
        /// <summary>
        /// A list of Zone to launch the DB instance.
        /// </summary>
        public readonly ImmutableArray<GetInstanceEnginesInstanceEnginesZoneIdsResult> ZoneIds;

        [OutputConstructor]
        private GetInstanceEnginesInstanceEnginesResult(
            string category,
            string engine,
            string engineVersion,
            ImmutableArray<GetInstanceEnginesInstanceEnginesZoneIdsResult> zoneIds)
        {
            Category = category;
            Engine = engine;
            EngineVersion = engineVersion;
            ZoneIds = zoneIds;
        }
    }

    [OutputType]
    public sealed class GetInstanceEnginesInstanceEnginesZoneIdsResult
    {
        /// <summary>
        /// The Zone to launch the DB instance
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of sub zone ids which in the id - e.g If `id` is `cn-beijing-MAZ5(a,b)`, `sub_zone_ids` will be `["cn-beijing-a", "cn-beijing-b"]`.
        /// </summary>
        public readonly ImmutableArray<string> SubZoneIds;

        [OutputConstructor]
        private GetInstanceEnginesInstanceEnginesZoneIdsResult(
            string id,
            ImmutableArray<string> subZoneIds)
        {
            Id = id;
            SubZoneIds = subZoneIds;
        }
    }
    }
}
