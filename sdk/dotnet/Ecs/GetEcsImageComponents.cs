// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public static class GetEcsImageComponents
    {
        /// <summary>
        /// This data source provides the Ecs Image Components of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.159.0+.
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
        ///     var ids = AliCloud.Ecs.GetEcsImageComponents.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Ecs.GetEcsImageComponents.Invoke(new()
        ///     {
        ///         NameRegex = "^my-ImageComponent",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ecsImageComponentId1"] = ids.Apply(getEcsImageComponentsResult =&gt; getEcsImageComponentsResult.Components[0]?.Id),
        ///         ["ecsImageComponentId2"] = nameRegex.Apply(getEcsImageComponentsResult =&gt; getEcsImageComponentsResult.Components[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEcsImageComponentsResult> InvokeAsync(GetEcsImageComponentsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEcsImageComponentsResult>("alicloud:ecs/getEcsImageComponents:getEcsImageComponents", args ?? new GetEcsImageComponentsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Ecs Image Components of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.159.0+.
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
        ///     var ids = AliCloud.Ecs.GetEcsImageComponents.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Ecs.GetEcsImageComponents.Invoke(new()
        ///     {
        ///         NameRegex = "^my-ImageComponent",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ecsImageComponentId1"] = ids.Apply(getEcsImageComponentsResult =&gt; getEcsImageComponentsResult.Components[0]?.Id),
        ///         ["ecsImageComponentId2"] = nameRegex.Apply(getEcsImageComponentsResult =&gt; getEcsImageComponentsResult.Components[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetEcsImageComponentsResult> Invoke(GetEcsImageComponentsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEcsImageComponentsResult>("alicloud:ecs/getEcsImageComponents:getEcsImageComponents", args ?? new GetEcsImageComponentsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEcsImageComponentsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Image Component IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of the image component.
        /// </summary>
        [Input("imageComponentName")]
        public string? ImageComponentName { get; set; }

        /// <summary>
        /// A regex string to filter results by Image Component name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The type of the image component.
        /// </summary>
        [Input("owner")]
        public string? Owner { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// List of label key-value pairs.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetEcsImageComponentsArgs()
        {
        }
        public static new GetEcsImageComponentsArgs Empty => new GetEcsImageComponentsArgs();
    }

    public sealed class GetEcsImageComponentsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Image Component IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of the image component.
        /// </summary>
        [Input("imageComponentName")]
        public Input<string>? ImageComponentName { get; set; }

        /// <summary>
        /// A regex string to filter results by Image Component name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The type of the image component.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// List of label key-value pairs.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetEcsImageComponentsInvokeArgs()
        {
        }
        public static new GetEcsImageComponentsInvokeArgs Empty => new GetEcsImageComponentsInvokeArgs();
    }


    [OutputType]
    public sealed class GetEcsImageComponentsResult
    {
        public readonly ImmutableArray<Outputs.GetEcsImageComponentsComponentResult> Components;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? ImageComponentName;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? Owner;
        public readonly string? ResourceGroupId;
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetEcsImageComponentsResult(
            ImmutableArray<Outputs.GetEcsImageComponentsComponentResult> components,

            string id,

            ImmutableArray<string> ids,

            string? imageComponentName,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? owner,

            string? resourceGroupId,

            ImmutableDictionary<string, object>? tags)
        {
            Components = components;
            Id = id;
            Ids = ids;
            ImageComponentName = imageComponentName;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Owner = owner;
            ResourceGroupId = resourceGroupId;
            Tags = tags;
        }
    }
}
