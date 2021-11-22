// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.Rds
{
    public static class GetInstances
    {
        /// <summary>
        /// The `alicloud.rds.getInstances` data source provides a collection of RDS instances available in Alibaba Cloud account.
        /// Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
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
        ///         var dbInstancesDs = Output.Create(AliCloud.Rds.GetInstances.InvokeAsync(new AliCloud.Rds.GetInstancesArgs
        ///         {
        ///             NameRegex = "data-\\d+",
        ///             Status = "Running",
        ///             Tags = 
        ///             {
        ///                 { "size", "tiny" },
        ///                 { "type", "database" },
        ///             },
        ///         }));
        ///         this.FirstDbInstanceId = dbInstancesDs.Apply(dbInstancesDs =&gt; dbInstancesDs.Instances?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstDbInstanceId")]
        ///     public Output&lt;string&gt; FirstDbInstanceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("alicloud:rds/getInstances:getInstances", args ?? new GetInstancesArgs(), options.WithVersion());

        /// <summary>
        /// The `alicloud.rds.getInstances` data source provides a collection of RDS instances available in Alibaba Cloud account.
        /// Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
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
        ///         var dbInstancesDs = Output.Create(AliCloud.Rds.GetInstances.InvokeAsync(new AliCloud.Rds.GetInstancesArgs
        ///         {
        ///             NameRegex = "data-\\d+",
        ///             Status = "Running",
        ///             Tags = 
        ///             {
        ///                 { "size", "tiny" },
        ///                 { "type", "database" },
        ///             },
        ///         }));
        ///         this.FirstDbInstanceId = dbInstancesDs.Apply(dbInstancesDs =&gt; dbInstancesDs.Instances?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstDbInstanceId")]
        ///     public Output&lt;string&gt; FirstDbInstanceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstancesResult> Invoke(GetInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstancesResult>("alicloud:rds/getInstances:getInstances", args ?? new GetInstancesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetInstancesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// `Standard` for standard access mode and `Safe` for high security access mode.
        /// </summary>
        [Input("connectionMode")]
        public string? ConnectionMode { get; set; }

        /// <summary>
        /// `Primary` for primary instance, `Readonly` for read-only instance, `Guard` for disaster recovery instance, and `Temp` for temporary instance.
        /// </summary>
        [Input("dbType")]
        public string? DbType { get; set; }

        /// <summary>
        /// Default to `false`. Set it to `true` can output parameter template about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        /// <summary>
        /// Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL` and `PPAS`. If no value is specified, all types are returned.
        /// </summary>
        [Input("engine")]
        public string? Engine { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of RDS instance IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by instance name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Status of the instance.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A map of tags assigned to the DB instances. 
        /// Note: Before 1.60.0, the value's format is a `json` string which including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `"{\"key1\":\"value1\"}"`
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        /// <summary>
        /// Used to retrieve instances belong to specified VPC.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        /// <summary>
        /// Used to retrieve instances belong to specified `vswitch` resources.
        /// </summary>
        [Input("vswitchId")]
        public string? VswitchId { get; set; }

        public GetInstancesArgs()
        {
        }
    }

    public sealed class GetInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// `Standard` for standard access mode and `Safe` for high security access mode.
        /// </summary>
        [Input("connectionMode")]
        public Input<string>? ConnectionMode { get; set; }

        /// <summary>
        /// `Primary` for primary instance, `Readonly` for read-only instance, `Guard` for disaster recovery instance, and `Temp` for temporary instance.
        /// </summary>
        [Input("dbType")]
        public Input<string>? DbType { get; set; }

        /// <summary>
        /// Default to `false`. Set it to `true` can output parameter template about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        /// <summary>
        /// Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL` and `PPAS`. If no value is specified, all types are returned.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of RDS instance IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by instance name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Status of the instance.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags assigned to the DB instances. 
        /// Note: Before 1.60.0, the value's format is a `json` string which including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `"{\"key1\":\"value1\"}"`
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Used to retrieve instances belong to specified VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Used to retrieve instances belong to specified `vswitch` resources.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public GetInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstancesResult
    {
        /// <summary>
        /// `Standard` for standard access mode and `Safe` for high security access mode.
        /// </summary>
        public readonly string? ConnectionMode;
        /// <summary>
        /// `Primary` for primary instance, `Readonly` for read-only instance, `Guard` for disaster recovery instance, and `Temp` for temporary instance.
        /// </summary>
        public readonly string? DbType;
        public readonly bool? EnableDetails;
        /// <summary>
        /// Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL` and `PPAS`. If no value is specified, all types are returned.
        /// </summary>
        public readonly string? Engine;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of RDS instance IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of RDS instances. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstancesInstanceResult> Instances;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of RDS instance names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// Status of the instance.
        /// </summary>
        public readonly string? Status;
        public readonly ImmutableDictionary<string, object>? Tags;
        /// <summary>
        /// ID of the VPC the instance belongs to.
        /// </summary>
        public readonly string? VpcId;
        /// <summary>
        /// ID of the VSwitch the instance belongs to.
        /// </summary>
        public readonly string? VswitchId;

        [OutputConstructor]
        private GetInstancesResult(
            string? connectionMode,

            string? dbType,

            bool? enableDetails,

            string? engine,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetInstancesInstanceResult> instances,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? status,

            ImmutableDictionary<string, object>? tags,

            string? vpcId,

            string? vswitchId)
        {
            ConnectionMode = connectionMode;
            DbType = dbType;
            EnableDetails = enableDetails;
            Engine = engine;
            Id = id;
            Ids = ids;
            Instances = instances;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
            Tags = tags;
            VpcId = vpcId;
            VswitchId = vswitchId;
        }
    }
}
