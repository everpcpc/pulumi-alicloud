// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public static class GetEcsNetworkInterfacePermissions
    {
        /// <summary>
        /// This data source provides the Ecs Network Interface Permissions of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.166.0+.
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
        ///         var ids = Output.Create(AliCloud.Ecs.GetEcsNetworkInterfacePermissions.InvokeAsync(new AliCloud.Ecs.GetEcsNetworkInterfacePermissionsArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_value",
        ///             },
        ///             NetworkInterfaceId = "example_value",
        ///         }));
        ///         this.EcsNetworkInterfacePermissionId1 = ids.Apply(ids =&gt; ids.Permissions?[0]?.Id);
        ///     }
        /// 
        ///     [Output("ecsNetworkInterfacePermissionId1")]
        ///     public Output&lt;string&gt; EcsNetworkInterfacePermissionId1 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEcsNetworkInterfacePermissionsResult> InvokeAsync(GetEcsNetworkInterfacePermissionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEcsNetworkInterfacePermissionsResult>("alicloud:ecs/getEcsNetworkInterfacePermissions:getEcsNetworkInterfacePermissions", args ?? new GetEcsNetworkInterfacePermissionsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Ecs Network Interface Permissions of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.166.0+.
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
        ///         var ids = Output.Create(AliCloud.Ecs.GetEcsNetworkInterfacePermissions.InvokeAsync(new AliCloud.Ecs.GetEcsNetworkInterfacePermissionsArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_value",
        ///             },
        ///             NetworkInterfaceId = "example_value",
        ///         }));
        ///         this.EcsNetworkInterfacePermissionId1 = ids.Apply(ids =&gt; ids.Permissions?[0]?.Id);
        ///     }
        /// 
        ///     [Output("ecsNetworkInterfacePermissionId1")]
        ///     public Output&lt;string&gt; EcsNetworkInterfacePermissionId1 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetEcsNetworkInterfacePermissionsResult> Invoke(GetEcsNetworkInterfacePermissionsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEcsNetworkInterfacePermissionsResult>("alicloud:ecs/getEcsNetworkInterfacePermissions:getEcsNetworkInterfacePermissions", args ?? new GetEcsNetworkInterfacePermissionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEcsNetworkInterfacePermissionsArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Network Interface Permission IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The ID of the network interface.
        /// </summary>
        [Input("networkInterfaceId", required: true)]
        public string NetworkInterfaceId { get; set; } = null!;

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        /// <summary>
        /// The Status of the Network Interface Permissions.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetEcsNetworkInterfacePermissionsArgs()
        {
        }
    }

    public sealed class GetEcsNetworkInterfacePermissionsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Network Interface Permission IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The ID of the network interface.
        /// </summary>
        [Input("networkInterfaceId", required: true)]
        public Input<string> NetworkInterfaceId { get; set; } = null!;

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        /// <summary>
        /// The Status of the Network Interface Permissions.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetEcsNetworkInterfacePermissionsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEcsNetworkInterfacePermissionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string NetworkInterfaceId;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;
        public readonly ImmutableArray<Outputs.GetEcsNetworkInterfacePermissionsPermissionResult> Permissions;
        public readonly string? Status;
        public readonly int TotalCount;

        [OutputConstructor]
        private GetEcsNetworkInterfacePermissionsResult(
            string id,

            ImmutableArray<string> ids,

            string networkInterfaceId,

            string? outputFile,

            int? pageNumber,

            int? pageSize,

            ImmutableArray<Outputs.GetEcsNetworkInterfacePermissionsPermissionResult> permissions,

            string? status,

            int totalCount)
        {
            Id = id;
            Ids = ids;
            NetworkInterfaceId = networkInterfaceId;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
            Permissions = permissions;
            Status = status;
            TotalCount = totalCount;
        }
    }
}
