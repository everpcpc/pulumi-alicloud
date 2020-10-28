// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.PolarDB
{
    public static class GetNodeClasses
    {
        /// <summary>
        /// This data source provides the PolarDB node classes resource available info of Alibaba Cloud.
        /// 
        /// &gt; **NOTE:** Available in v1.81.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var resourcesZones = Output.Create(AliCloud.GetZones.InvokeAsync(new AliCloud.GetZonesArgs
        ///         {
        ///             AvailableResourceCreation = "PolarDB",
        ///         }));
        ///         var resourcesNodeClasses = resourcesZones.Apply(resourcesZones =&gt; Output.Create(AliCloud.PolarDB.GetNodeClasses.InvokeAsync(new AliCloud.PolarDB.GetNodeClassesArgs
        ///         {
        ///             ZoneId = resourcesZones.Zones[0].Id,
        ///             PayType = "PostPaid",
        ///             DbType = "MySQL",
        ///             DbVersion = "5.6",
        ///         })));
        ///         this.PolardbNodeClasses = resourcesNodeClasses.Apply(resourcesNodeClasses =&gt; resourcesNodeClasses.Classes);
        ///     }
        /// 
        ///     [Output("polardbNodeClasses")]
        ///     public Output&lt;string&gt; PolardbNodeClasses { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNodeClassesResult> InvokeAsync(GetNodeClassesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNodeClassesResult>("alicloud:polardb/getNodeClasses:getNodeClasses", args ?? new GetNodeClassesArgs(), options.WithVersion());
    }


    public sealed class GetNodeClassesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The PolarDB node class type by the user.
        /// </summary>
        [Input("dbNodeClass")]
        public string? DbNodeClass { get; set; }

        /// <summary>
        /// Database type. Options are `MySQL`, `PostgreSQL`, `Oracle`. If db_type is set, db_version also needs to be set.
        /// </summary>
        [Input("dbType")]
        public string? DbType { get; set; }

        /// <summary>
        /// Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/98169.htm) `DBVersion`. If db_version is set, db_type also needs to be set.
        /// </summary>
        [Input("dbVersion")]
        public string? DbVersion { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`.
        /// </summary>
        [Input("payType", required: true)]
        public string PayType { get; set; } = null!;

        /// <summary>
        /// The Region to launch the PolarDB cluster.
        /// </summary>
        [Input("regionId")]
        public string? RegionId { get; set; }

        /// <summary>
        /// The Zone to launch the PolarDB cluster.
        /// </summary>
        [Input("zoneId")]
        public string? ZoneId { get; set; }

        public GetNodeClassesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNodeClassesResult
    {
        /// <summary>
        /// A list of PolarDB node classes. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNodeClassesClassResult> Classes;
        /// <summary>
        /// PolarDB node available class.
        /// </summary>
        public readonly string? DbNodeClass;
        public readonly string? DbType;
        public readonly string? DbVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;
        public readonly string PayType;
        public readonly string? RegionId;
        /// <summary>
        /// The Zone to launch the PolarDB cluster.
        /// </summary>
        public readonly string? ZoneId;

        [OutputConstructor]
        private GetNodeClassesResult(
            ImmutableArray<Outputs.GetNodeClassesClassResult> classes,

            string? dbNodeClass,

            string? dbType,

            string? dbVersion,

            string id,

            string? outputFile,

            string payType,

            string? regionId,

            string? zoneId)
        {
            Classes = classes;
            DbNodeClass = dbNodeClass;
            DbType = dbType;
            DbVersion = dbVersion;
            Id = id;
            OutputFile = outputFile;
            PayType = payType;
            RegionId = regionId;
            ZoneId = zoneId;
        }
    }
}
