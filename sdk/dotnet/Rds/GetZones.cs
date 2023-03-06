// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Rds
{
    public static class GetZones
    {
        /// <summary>
        /// This data source provides availability zones for RDS that can be accessed by an Alibaba Cloud account within the region configured in the provider.
        /// 
        /// &gt; **NOTE:** Available in v1.73.0+.
        /// </summary>
        public static Task<GetZonesResult> InvokeAsync(GetZonesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetZonesResult>("alicloud:rds/getZones:getZones", args ?? new GetZonesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides availability zones for RDS that can be accessed by an Alibaba Cloud account within the region configured in the provider.
        /// 
        /// &gt; **NOTE:** Available in v1.73.0+.
        /// </summary>
        public static Output<GetZonesResult> Invoke(GetZonesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetZonesResult>("alicloud:rds/getZones:getZones", args ?? new GetZonesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetZonesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// DB Instance category. the value like [`Basic`, `HighAvailability`, `Finance`, `AlwaysOn`, `serverless_basic`], [detail info](https://www.alibabacloud.com/help/doc-detail/69795.htm).
        /// </summary>
        [Input("category")]
        public string? Category { get; set; }

        [Input("dbInstanceClass")]
        public string? DbInstanceClass { get; set; }

        /// <summary>
        /// The DB instance storage space required by the user. Valid values: "cloud_ssd", "local_ssd", "cloud_essd", "cloud_essd2", "cloud_essd3".
        /// </summary>
        [Input("dbInstanceStorageType")]
        public string? DbInstanceStorageType { get; set; }

        /// <summary>
        /// Database type. Valid values: "MySQL", "SQLServer", "PostgreSQL", "PPAS", "MariaDB". If not set, it will match all of engines.
        /// </summary>
        [Input("engine")]
        public string? Engine { get; set; }

        /// <summary>
        /// Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
        /// </summary>
        [Input("engineVersion")]
        public string? EngineVersion { get; set; }

        /// <summary>
        /// Filter the results by a specific instance charge type. Valid values: `PrePaid` and `PostPaid` and `Serverless`. Default to `PostPaid`.
        /// </summary>
        [Input("instanceChargeType")]
        public string? InstanceChargeType { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.137.0 and using `multi_zone` instead.
        /// </summary>
        [Input("multi")]
        public bool? Multi { get; set; }

        /// <summary>
        /// Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch RDS instances.
        /// </summary>
        [Input("multiZone")]
        public bool? MultiZone { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetZonesArgs()
        {
        }
        public static new GetZonesArgs Empty => new GetZonesArgs();
    }

    public sealed class GetZonesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// DB Instance category. the value like [`Basic`, `HighAvailability`, `Finance`, `AlwaysOn`, `serverless_basic`], [detail info](https://www.alibabacloud.com/help/doc-detail/69795.htm).
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        [Input("dbInstanceClass")]
        public Input<string>? DbInstanceClass { get; set; }

        /// <summary>
        /// The DB instance storage space required by the user. Valid values: "cloud_ssd", "local_ssd", "cloud_essd", "cloud_essd2", "cloud_essd3".
        /// </summary>
        [Input("dbInstanceStorageType")]
        public Input<string>? DbInstanceStorageType { get; set; }

        /// <summary>
        /// Database type. Valid values: "MySQL", "SQLServer", "PostgreSQL", "PPAS", "MariaDB". If not set, it will match all of engines.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// Filter the results by a specific instance charge type. Valid values: `PrePaid` and `PostPaid` and `Serverless`. Default to `PostPaid`.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.137.0 and using `multi_zone` instead.
        /// </summary>
        [Input("multi")]
        public Input<bool>? Multi { get; set; }

        /// <summary>
        /// Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch RDS instances.
        /// </summary>
        [Input("multiZone")]
        public Input<bool>? MultiZone { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetZonesInvokeArgs()
        {
        }
        public static new GetZonesInvokeArgs Empty => new GetZonesInvokeArgs();
    }


    [OutputType]
    public sealed class GetZonesResult
    {
        public readonly string? Category;
        public readonly string? DbInstanceClass;
        public readonly string? DbInstanceStorageType;
        public readonly string? Engine;
        public readonly string? EngineVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of zone IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? InstanceChargeType;
        public readonly bool? Multi;
        public readonly bool? MultiZone;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of availability zones. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZonesZoneResult> Zones;

        [OutputConstructor]
        private GetZonesResult(
            string? category,

            string? dbInstanceClass,

            string? dbInstanceStorageType,

            string? engine,

            string? engineVersion,

            string id,

            ImmutableArray<string> ids,

            string? instanceChargeType,

            bool? multi,

            bool? multiZone,

            string? outputFile,

            ImmutableArray<Outputs.GetZonesZoneResult> zones)
        {
            Category = category;
            DbInstanceClass = dbInstanceClass;
            DbInstanceStorageType = dbInstanceStorageType;
            Engine = engine;
            EngineVersion = engineVersion;
            Id = id;
            Ids = ids;
            InstanceChargeType = instanceChargeType;
            Multi = multi;
            MultiZone = multiZone;
            OutputFile = outputFile;
            Zones = zones;
        }
    }
}
