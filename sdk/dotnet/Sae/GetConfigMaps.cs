// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae
{
    public static class GetConfigMaps
    {
        /// <summary>
        /// This data source provides the Sae Config Maps of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.130.0+.
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
        /// using System.Text.Json;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var config = new Config();
        ///     var configMapName = config.Get("configMapName") ?? "examplename";
        ///     var exampleNamespace = new AliCloud.Sae.Namespace("exampleNamespace", new()
        ///     {
        ///         NamespaceId = "cn-hangzhou:yourname",
        ///         NamespaceName = "example_value",
        ///         NamespaceDescription = "your_description",
        ///     });
        /// 
        ///     var exampleConfigMap = new AliCloud.Sae.ConfigMap("exampleConfigMap", new()
        ///     {
        ///         Data = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["env.home"] = "/root",
        ///             ["env.shell"] = "/bin/sh",
        ///         }),
        ///         NamespaceId = exampleNamespace.NamespaceId,
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Sae.GetConfigMaps.Invoke(new()
        ///     {
        ///         NamespaceId = exampleNamespace.NamespaceId,
        ///         NameRegex = "^example",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["saeConfigMapId"] = nameRegex.Apply(getConfigMapsResult =&gt; getConfigMapsResult.Maps[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetConfigMapsResult> InvokeAsync(GetConfigMapsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConfigMapsResult>("alicloud:sae/getConfigMaps:getConfigMaps", args ?? new GetConfigMapsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Sae Config Maps of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.130.0+.
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
        /// using System.Text.Json;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var config = new Config();
        ///     var configMapName = config.Get("configMapName") ?? "examplename";
        ///     var exampleNamespace = new AliCloud.Sae.Namespace("exampleNamespace", new()
        ///     {
        ///         NamespaceId = "cn-hangzhou:yourname",
        ///         NamespaceName = "example_value",
        ///         NamespaceDescription = "your_description",
        ///     });
        /// 
        ///     var exampleConfigMap = new AliCloud.Sae.ConfigMap("exampleConfigMap", new()
        ///     {
        ///         Data = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["env.home"] = "/root",
        ///             ["env.shell"] = "/bin/sh",
        ///         }),
        ///         NamespaceId = exampleNamespace.NamespaceId,
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Sae.GetConfigMaps.Invoke(new()
        ///     {
        ///         NamespaceId = exampleNamespace.NamespaceId,
        ///         NameRegex = "^example",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["saeConfigMapId"] = nameRegex.Apply(getConfigMapsResult =&gt; getConfigMapsResult.Maps[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetConfigMapsResult> Invoke(GetConfigMapsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigMapsResult>("alicloud:sae/getConfigMaps:getConfigMaps", args ?? new GetConfigMapsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigMapsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Config Map IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Config Map name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// The NamespaceId of Config Maps.
        /// </summary>
        [Input("namespaceId", required: true)]
        public string NamespaceId { get; set; } = null!;

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetConfigMapsArgs()
        {
        }
        public static new GetConfigMapsArgs Empty => new GetConfigMapsArgs();
    }

    public sealed class GetConfigMapsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Config Map IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Config Map name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// The NamespaceId of Config Maps.
        /// </summary>
        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetConfigMapsInvokeArgs()
        {
        }
        public static new GetConfigMapsInvokeArgs Empty => new GetConfigMapsInvokeArgs();
    }


    [OutputType]
    public sealed class GetConfigMapsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<Outputs.GetConfigMapsMapResult> Maps;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string NamespaceId;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetConfigMapsResult(
            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetConfigMapsMapResult> maps,

            string? nameRegex,

            ImmutableArray<string> names,

            string namespaceId,

            string? outputFile)
        {
            Id = id;
            Ids = ids;
            Maps = maps;
            NameRegex = nameRegex;
            Names = names;
            NamespaceId = namespaceId;
            OutputFile = outputFile;
        }
    }
}
