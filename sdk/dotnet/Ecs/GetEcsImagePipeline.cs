// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public static class GetEcsImagePipeline
    {
        /// <summary>
        /// This data source provides the Ecs Image Pipelines of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.163.0+.
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
        ///         var ids = Output.Create(AliCloud.Ecs.GetEcsImagePipeline.InvokeAsync(new AliCloud.Ecs.GetEcsImagePipelineArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_value",
        ///             },
        ///         }));
        ///         this.EcsImagePipelineId1 = ids.Apply(ids =&gt; ids.Pipelines?[0]?.Id);
        ///         var nameRegex = Output.Create(AliCloud.Ecs.GetEcsImagePipeline.InvokeAsync(new AliCloud.Ecs.GetEcsImagePipelineArgs
        ///         {
        ///             NameRegex = "^my-ImagePipeline",
        ///         }));
        ///         this.EcsImagePipelineId2 = nameRegex.Apply(nameRegex =&gt; nameRegex.Pipelines?[0]?.Id);
        ///     }
        /// 
        ///     [Output("ecsImagePipelineId1")]
        ///     public Output&lt;string&gt; EcsImagePipelineId1 { get; set; }
        ///     [Output("ecsImagePipelineId2")]
        ///     public Output&lt;string&gt; EcsImagePipelineId2 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEcsImagePipelineResult> InvokeAsync(GetEcsImagePipelineArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEcsImagePipelineResult>("alicloud:ecs/getEcsImagePipeline:getEcsImagePipeline", args ?? new GetEcsImagePipelineArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Ecs Image Pipelines of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.163.0+.
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
        ///         var ids = Output.Create(AliCloud.Ecs.GetEcsImagePipeline.InvokeAsync(new AliCloud.Ecs.GetEcsImagePipelineArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_value",
        ///             },
        ///         }));
        ///         this.EcsImagePipelineId1 = ids.Apply(ids =&gt; ids.Pipelines?[0]?.Id);
        ///         var nameRegex = Output.Create(AliCloud.Ecs.GetEcsImagePipeline.InvokeAsync(new AliCloud.Ecs.GetEcsImagePipelineArgs
        ///         {
        ///             NameRegex = "^my-ImagePipeline",
        ///         }));
        ///         this.EcsImagePipelineId2 = nameRegex.Apply(nameRegex =&gt; nameRegex.Pipelines?[0]?.Id);
        ///     }
        /// 
        ///     [Output("ecsImagePipelineId1")]
        ///     public Output&lt;string&gt; EcsImagePipelineId1 { get; set; }
        ///     [Output("ecsImagePipelineId2")]
        ///     public Output&lt;string&gt; EcsImagePipelineId2 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetEcsImagePipelineResult> Invoke(GetEcsImagePipelineInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEcsImagePipelineResult>("alicloud:ecs/getEcsImagePipeline:getEcsImagePipeline", args ?? new GetEcsImagePipelineInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEcsImagePipelineArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Image Pipeline ids.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of the image template.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// A regex string to filter results by Image Pipeline name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ID of the resource group to which the image template belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetEcsImagePipelineArgs()
        {
        }
    }

    public sealed class GetEcsImagePipelineInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Image Pipeline ids.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of the image template.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A regex string to filter results by Image Pipeline name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The ID of the resource group to which the image template belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetEcsImagePipelineInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEcsImagePipelineResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? Name;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetEcsImagePipelinePipelineResult> Pipelines;
        public readonly string? ResourceGroupId;
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetEcsImagePipelineResult(
            string id,

            ImmutableArray<string> ids,

            string? name,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetEcsImagePipelinePipelineResult> pipelines,

            string? resourceGroupId,

            ImmutableDictionary<string, object>? tags)
        {
            Id = id;
            Ids = ids;
            Name = name;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Pipelines = pipelines;
            ResourceGroupId = resourceGroupId;
            Tags = tags;
        }
    }
}
