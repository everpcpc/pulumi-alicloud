// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.SimpleApplicationServer
{
    public static class GetServerDisks
    {
        /// <summary>
        /// This data source provides the Simple Application Server Disks of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.143.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.SimpleApplicationServer.GetServerDisks.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.SimpleApplicationServer.GetServerDisks.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Disk",
        ///     });
        /// 
        ///     var status = AliCloud.SimpleApplicationServer.GetServerDisks.Invoke(new()
        ///     {
        ///         Status = "In_use",
        ///     });
        /// 
        ///     var instanceId = AliCloud.SimpleApplicationServer.GetServerDisks.Invoke(new()
        ///     {
        ///         InstanceId = "example_value",
        ///     });
        /// 
        ///     var diskType = AliCloud.SimpleApplicationServer.GetServerDisks.Invoke(new()
        ///     {
        ///         DiskType = "System",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["simpleApplicationServerDiskId1"] = ids.Apply(getServerDisksResult =&gt; getServerDisksResult.Disks[0]?.Id),
        ///         ["simpleApplicationServerDiskId2"] = nameRegex.Apply(getServerDisksResult =&gt; getServerDisksResult.Disks[0]?.Id),
        ///         ["simpleApplicationServerDiskId3"] = status.Apply(getServerDisksResult =&gt; getServerDisksResult.Disks[0]?.Id),
        ///         ["simpleApplicationServerDiskId4"] = instanceId.Apply(getServerDisksResult =&gt; getServerDisksResult.Disks[0]?.Id),
        ///         ["simpleApplicationServerDiskId5"] = diskType.Apply(getServerDisksResult =&gt; getServerDisksResult.Disks[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServerDisksResult> InvokeAsync(GetServerDisksArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServerDisksResult>("alicloud:simpleapplicationserver/getServerDisks:getServerDisks", args ?? new GetServerDisksArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Simple Application Server Disks of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.143.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.SimpleApplicationServer.GetServerDisks.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.SimpleApplicationServer.GetServerDisks.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Disk",
        ///     });
        /// 
        ///     var status = AliCloud.SimpleApplicationServer.GetServerDisks.Invoke(new()
        ///     {
        ///         Status = "In_use",
        ///     });
        /// 
        ///     var instanceId = AliCloud.SimpleApplicationServer.GetServerDisks.Invoke(new()
        ///     {
        ///         InstanceId = "example_value",
        ///     });
        /// 
        ///     var diskType = AliCloud.SimpleApplicationServer.GetServerDisks.Invoke(new()
        ///     {
        ///         DiskType = "System",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["simpleApplicationServerDiskId1"] = ids.Apply(getServerDisksResult =&gt; getServerDisksResult.Disks[0]?.Id),
        ///         ["simpleApplicationServerDiskId2"] = nameRegex.Apply(getServerDisksResult =&gt; getServerDisksResult.Disks[0]?.Id),
        ///         ["simpleApplicationServerDiskId3"] = status.Apply(getServerDisksResult =&gt; getServerDisksResult.Disks[0]?.Id),
        ///         ["simpleApplicationServerDiskId4"] = instanceId.Apply(getServerDisksResult =&gt; getServerDisksResult.Disks[0]?.Id),
        ///         ["simpleApplicationServerDiskId5"] = diskType.Apply(getServerDisksResult =&gt; getServerDisksResult.Disks[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServerDisksResult> Invoke(GetServerDisksInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerDisksResult>("alicloud:simpleapplicationserver/getServerDisks:getServerDisks", args ?? new GetServerDisksInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerDisksArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The type of the disk. Possible values: `System`, `Data`.
        /// </summary>
        [Input("diskType")]
        public string? DiskType { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Disk IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Alibaba Cloud simple application server instance ID.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// A regex string to filter results by Disk name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the disk. Valid values: `ReIniting`, `Creating`, `In_Use`, `Available`, `Attaching`, `Detaching`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetServerDisksArgs()
        {
        }
        public static new GetServerDisksArgs Empty => new GetServerDisksArgs();
    }

    public sealed class GetServerDisksInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The type of the disk. Possible values: `System`, `Data`.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Disk IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Alibaba Cloud simple application server instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// A regex string to filter results by Disk name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of the disk. Valid values: `ReIniting`, `Creating`, `In_Use`, `Available`, `Attaching`, `Detaching`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetServerDisksInvokeArgs()
        {
        }
        public static new GetServerDisksInvokeArgs Empty => new GetServerDisksInvokeArgs();
    }


    [OutputType]
    public sealed class GetServerDisksResult
    {
        public readonly string? DiskType;
        public readonly ImmutableArray<Outputs.GetServerDisksDiskResult> Disks;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? InstanceId;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? Status;

        [OutputConstructor]
        private GetServerDisksResult(
            string? diskType,

            ImmutableArray<Outputs.GetServerDisksDiskResult> disks,

            string id,

            ImmutableArray<string> ids,

            string? instanceId,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? status)
        {
            DiskType = diskType;
            Disks = disks;
            Id = id;
            Ids = ids;
            InstanceId = instanceId;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
        }
    }
}
