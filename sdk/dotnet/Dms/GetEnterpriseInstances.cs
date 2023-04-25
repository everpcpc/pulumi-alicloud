// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dms
{
    public static class GetEnterpriseInstances
    {
        /// <summary>
        /// This data source provides a list of DMS Enterprise Instances in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.88.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var dmsEnterpriseInstancesDs = AliCloud.Dms.GetEnterpriseInstances.Invoke(new()
        ///     {
        ///         EnvType = "test",
        ///         InstanceType = "mysql",
        ///         NameRegex = "tf_testAcc",
        ///         NetType = "CLASSIC",
        ///         OutputFile = "dms_enterprise_instances.json",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstDatabaseInstanceId"] = dmsEnterpriseInstancesDs.Apply(getEnterpriseInstancesResult =&gt; getEnterpriseInstancesResult.Instances[0]?.InstanceId),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEnterpriseInstancesResult> InvokeAsync(GetEnterpriseInstancesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEnterpriseInstancesResult>("alicloud:dms/getEnterpriseInstances:getEnterpriseInstances", args ?? new GetEnterpriseInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides a list of DMS Enterprise Instances in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.88.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var dmsEnterpriseInstancesDs = AliCloud.Dms.GetEnterpriseInstances.Invoke(new()
        ///     {
        ///         EnvType = "test",
        ///         InstanceType = "mysql",
        ///         NameRegex = "tf_testAcc",
        ///         NetType = "CLASSIC",
        ///         OutputFile = "dms_enterprise_instances.json",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstDatabaseInstanceId"] = dmsEnterpriseInstancesDs.Apply(getEnterpriseInstancesResult =&gt; getEnterpriseInstancesResult.Instances[0]?.InstanceId),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetEnterpriseInstancesResult> Invoke(GetEnterpriseInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnterpriseInstancesResult>("alicloud:dms/getEnterpriseInstances:getEnterpriseInstances", args ?? new GetEnterpriseInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnterpriseInstancesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The type of the environment to which the database instance belongs.
        /// </summary>
        [Input("envType")]
        public string? EnvType { get; set; }

        /// <summary>
        /// A regex string to filter the results by the DMS Enterprise Instance instance_alias.
        /// </summary>
        [Input("instanceAliasRegex")]
        public string? InstanceAliasRegex { get; set; }

        /// <summary>
        /// The source of the database instance.
        /// </summary>
        [Input("instanceSource")]
        public string? InstanceSource { get; set; }

        /// <summary>
        /// The ID of the database instance.
        /// </summary>
        [Input("instanceType")]
        public string? InstanceType { get; set; }

        /// <summary>
        /// A regex string to filter the results by the DMS Enterprise Instance instance_alias.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// The network type of the database instance. Valid values: CLASSIC and VPC. For more information about the valid values, see the description of the RegisterInstance operation.
        /// </summary>
        [Input("netType")]
        public string? NetType { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The keyword used to query database instances.
        /// </summary>
        [Input("searchKey")]
        public string? SearchKey { get; set; }

        /// <summary>
        /// Filter the results by status of the DMS Enterprise Instances. Valid values: `NORMAL`, `UNAVAILABLE`, `UNKNOWN`, `DELETED`, `DISABLE`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The ID of the tenant in Data Management (DMS) Enterprise.
        /// </summary>
        [Input("tid")]
        public int? Tid { get; set; }

        public GetEnterpriseInstancesArgs()
        {
        }
        public static new GetEnterpriseInstancesArgs Empty => new GetEnterpriseInstancesArgs();
    }

    public sealed class GetEnterpriseInstancesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The type of the environment to which the database instance belongs.
        /// </summary>
        [Input("envType")]
        public Input<string>? EnvType { get; set; }

        /// <summary>
        /// A regex string to filter the results by the DMS Enterprise Instance instance_alias.
        /// </summary>
        [Input("instanceAliasRegex")]
        public Input<string>? InstanceAliasRegex { get; set; }

        /// <summary>
        /// The source of the database instance.
        /// </summary>
        [Input("instanceSource")]
        public Input<string>? InstanceSource { get; set; }

        /// <summary>
        /// The ID of the database instance.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// A regex string to filter the results by the DMS Enterprise Instance instance_alias.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// The network type of the database instance. Valid values: CLASSIC and VPC. For more information about the valid values, see the description of the RegisterInstance operation.
        /// </summary>
        [Input("netType")]
        public Input<string>? NetType { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The keyword used to query database instances.
        /// </summary>
        [Input("searchKey")]
        public Input<string>? SearchKey { get; set; }

        /// <summary>
        /// Filter the results by status of the DMS Enterprise Instances. Valid values: `NORMAL`, `UNAVAILABLE`, `UNKNOWN`, `DELETED`, `DISABLE`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the tenant in Data Management (DMS) Enterprise.
        /// </summary>
        [Input("tid")]
        public Input<int>? Tid { get; set; }

        public GetEnterpriseInstancesInvokeArgs()
        {
        }
        public static new GetEnterpriseInstancesInvokeArgs Empty => new GetEnterpriseInstancesInvokeArgs();
    }


    [OutputType]
    public sealed class GetEnterpriseInstancesResult
    {
        /// <summary>
        /// The type of the environment to which the database instance belongs..
        /// </summary>
        public readonly string? EnvType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of DMS Enterprise IDs (Each of them consists of host:port).
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? InstanceAliasRegex;
        /// <summary>
        /// The ID of the database instance.
        /// </summary>
        public readonly string? InstanceSource;
        /// <summary>
        /// The ID of the database instance.
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// A list of KMS keys. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEnterpriseInstancesInstanceResult> Instances;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of DMS Enterprise names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? NetType;
        public readonly string? OutputFile;
        public readonly string? SearchKey;
        /// <summary>
        /// The status of the database instance.
        /// </summary>
        public readonly string? Status;
        public readonly int? Tid;

        [OutputConstructor]
        private GetEnterpriseInstancesResult(
            string? envType,

            string id,

            ImmutableArray<string> ids,

            string? instanceAliasRegex,

            string? instanceSource,

            string? instanceType,

            ImmutableArray<Outputs.GetEnterpriseInstancesInstanceResult> instances,

            string? nameRegex,

            ImmutableArray<string> names,

            string? netType,

            string? outputFile,

            string? searchKey,

            string? status,

            int? tid)
        {
            EnvType = envType;
            Id = id;
            Ids = ids;
            InstanceAliasRegex = instanceAliasRegex;
            InstanceSource = instanceSource;
            InstanceType = instanceType;
            Instances = instances;
            NameRegex = nameRegex;
            Names = names;
            NetType = netType;
            OutputFile = outputFile;
            SearchKey = searchKey;
            Status = status;
            Tid = tid;
        }
    }
}
