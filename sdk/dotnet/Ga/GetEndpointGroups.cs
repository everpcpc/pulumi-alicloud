// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga
{
    public static class GetEndpointGroups
    {
        /// <summary>
        /// This data source provides the Global Accelerator (GA) Endpoint Groups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.113.0+.
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
        ///     var example = AliCloud.Ga.GetEndpointGroups.Invoke(new()
        ///     {
        ///         AcceleratorId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value",
        ///         },
        ///         NameRegex = "the_resource_name",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstGaEndpointGroupId"] = example.Apply(getEndpointGroupsResult =&gt; getEndpointGroupsResult.Groups[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEndpointGroupsResult> InvokeAsync(GetEndpointGroupsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEndpointGroupsResult>("alicloud:ga/getEndpointGroups:getEndpointGroups", args ?? new GetEndpointGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Global Accelerator (GA) Endpoint Groups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.113.0+.
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
        ///     var example = AliCloud.Ga.GetEndpointGroups.Invoke(new()
        ///     {
        ///         AcceleratorId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value",
        ///         },
        ///         NameRegex = "the_resource_name",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstGaEndpointGroupId"] = example.Apply(getEndpointGroupsResult =&gt; getEndpointGroupsResult.Groups[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetEndpointGroupsResult> Invoke(GetEndpointGroupsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEndpointGroupsResult>("alicloud:ga/getEndpointGroups:getEndpointGroups", args ?? new GetEndpointGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEndpointGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Global Accelerator instance to which the endpoint group will be added.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public string AcceleratorId { get; set; } = null!;

        /// <summary>
        /// The endpoint group type. Valid values: `default`, `virtual`. Default value is `default`.
        /// </summary>
        [Input("endpointGroupType")]
        public string? EndpointGroupType { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Endpoint Group IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The ID of the listener that is associated with the endpoint group.
        /// </summary>
        [Input("listenerId")]
        public string? ListenerId { get; set; }

        /// <summary>
        /// A regex string to filter results by Endpoint Group name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the endpoint group.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetEndpointGroupsArgs()
        {
        }
        public static new GetEndpointGroupsArgs Empty => new GetEndpointGroupsArgs();
    }

    public sealed class GetEndpointGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Global Accelerator instance to which the endpoint group will be added.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public Input<string> AcceleratorId { get; set; } = null!;

        /// <summary>
        /// The endpoint group type. Valid values: `default`, `virtual`. Default value is `default`.
        /// </summary>
        [Input("endpointGroupType")]
        public Input<string>? EndpointGroupType { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Endpoint Group IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The ID of the listener that is associated with the endpoint group.
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        /// <summary>
        /// A regex string to filter results by Endpoint Group name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of the endpoint group.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetEndpointGroupsInvokeArgs()
        {
        }
        public static new GetEndpointGroupsInvokeArgs Empty => new GetEndpointGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetEndpointGroupsResult
    {
        public readonly string AcceleratorId;
        public readonly string? EndpointGroupType;
        public readonly ImmutableArray<Outputs.GetEndpointGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? ListenerId;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? Status;

        [OutputConstructor]
        private GetEndpointGroupsResult(
            string acceleratorId,

            string? endpointGroupType,

            ImmutableArray<Outputs.GetEndpointGroupsGroupResult> groups,

            string id,

            ImmutableArray<string> ids,

            string? listenerId,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? status)
        {
            AcceleratorId = acceleratorId;
            EndpointGroupType = endpointGroupType;
            Groups = groups;
            Id = id;
            Ids = ids;
            ListenerId = listenerId;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
        }
    }
}
