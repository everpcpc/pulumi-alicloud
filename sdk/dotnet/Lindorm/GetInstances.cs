// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.Lindorm
{
    public static class GetInstances
    {
        /// <summary>
        /// This data source provides the Lindorm Instances of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.132.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var ids = Output.Create(AliCloud.Lindorm.GetInstances.InvokeAsync());
        ///         this.LindormInstanceId1 = ids.Apply(ids =&gt; ids.Instances?[0]?.Id);
        ///         var nameRegex = Output.Create(AliCloud.Lindorm.GetInstances.InvokeAsync(new AliCloud.Lindorm.GetInstancesArgs
        ///         {
        ///             NameRegex = "^my-Instance",
        ///         }));
        ///         this.LindormInstanceId2 = nameRegex.Apply(nameRegex =&gt; nameRegex.Instances?[0]?.Id);
        ///     }
        /// 
        ///     [Output("lindormInstanceId1")]
        ///     public Output&lt;string&gt; LindormInstanceId1 { get; set; }
        ///     [Output("lindormInstanceId2")]
        ///     public Output&lt;string&gt; LindormInstanceId2 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("alicloud:lindorm/getInstances:getInstances", args ?? new GetInstancesArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the Lindorm Instances of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.132.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var ids = Output.Create(AliCloud.Lindorm.GetInstances.InvokeAsync());
        ///         this.LindormInstanceId1 = ids.Apply(ids =&gt; ids.Instances?[0]?.Id);
        ///         var nameRegex = Output.Create(AliCloud.Lindorm.GetInstances.InvokeAsync(new AliCloud.Lindorm.GetInstancesArgs
        ///         {
        ///             NameRegex = "^my-Instance",
        ///         }));
        ///         this.LindormInstanceId2 = nameRegex.Apply(nameRegex =&gt; nameRegex.Instances?[0]?.Id);
        ///     }
        /// 
        ///     [Output("lindormInstanceId1")]
        ///     public Output&lt;string&gt; LindormInstanceId1 { get; set; }
        ///     [Output("lindormInstanceId2")]
        ///     public Output&lt;string&gt; LindormInstanceId2 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstancesResult> Invoke(GetInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstancesResult>("alicloud:lindorm/getInstances:getInstances", args ?? new GetInstancesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetInstancesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Instance IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Instance name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The query str, which can use `instance_name` keyword for fuzzy search.
        /// </summary>
        [Input("queryStr")]
        public string? QueryStr { get; set; }

        /// <summary>
        /// The status of Instance, enumerative: Valid values: `ACTIVATION`, `DELETED`, `CREATING`, `CLASS_CHANGING`, `LOCKED`, `INSTANCE_LEVEL_MODIFY`, `NET_MODIFYING`, `RESIZING`, `RESTARTING`, `MINOR_VERSION_TRANSING`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The support engine. Valid values: `1` to `7`.
        /// </summary>
        [Input("supportEngine")]
        public int? SupportEngine { get; set; }

        public GetInstancesArgs()
        {
        }
    }

    public sealed class GetInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Instance IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Instance name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The query str, which can use `instance_name` keyword for fuzzy search.
        /// </summary>
        [Input("queryStr")]
        public Input<string>? QueryStr { get; set; }

        /// <summary>
        /// The status of Instance, enumerative: Valid values: `ACTIVATION`, `DELETED`, `CREATING`, `CLASS_CHANGING`, `LOCKED`, `INSTANCE_LEVEL_MODIFY`, `NET_MODIFYING`, `RESIZING`, `RESTARTING`, `MINOR_VERSION_TRANSING`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The support engine. Valid values: `1` to `7`.
        /// </summary>
        [Input("supportEngine")]
        public Input<int>? SupportEngine { get; set; }

        public GetInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstancesResult
    {
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<Outputs.GetInstancesInstanceResult> Instances;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? QueryStr;
        public readonly string? Status;
        public readonly int? SupportEngine;

        [OutputConstructor]
        private GetInstancesResult(
            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetInstancesInstanceResult> instances,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? queryStr,

            string? status,

            int? supportEngine)
        {
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            Instances = instances;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            QueryStr = queryStr;
            Status = status;
            SupportEngine = supportEngine;
        }
    }
}
