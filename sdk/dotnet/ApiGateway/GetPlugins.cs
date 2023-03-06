// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ApiGateway
{
    public static class GetPlugins
    {
        /// <summary>
        /// This data source provides the Api Gateway Plugins of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.187.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.ApiGateway.GetPlugins.Invoke();
        /// 
        ///     var nameRegex = AliCloud.ApiGateway.GetPlugins.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Plugin",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["apiGatewayPluginId1"] = ids.Apply(getPluginsResult =&gt; getPluginsResult.Plugins[0]?.Id),
        ///         ["apiGatewayPluginId2"] = nameRegex.Apply(getPluginsResult =&gt; getPluginsResult.Plugins[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPluginsResult> InvokeAsync(GetPluginsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPluginsResult>("alicloud:apigateway/getPlugins:getPlugins", args ?? new GetPluginsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Api Gateway Plugins of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.187.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.ApiGateway.GetPlugins.Invoke();
        /// 
        ///     var nameRegex = AliCloud.ApiGateway.GetPlugins.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Plugin",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["apiGatewayPluginId1"] = ids.Apply(getPluginsResult =&gt; getPluginsResult.Plugins[0]?.Id),
        ///         ["apiGatewayPluginId2"] = nameRegex.Apply(getPluginsResult =&gt; getPluginsResult.Plugins[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPluginsResult> Invoke(GetPluginsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPluginsResult>("alicloud:apigateway/getPlugins:getPlugins", args ?? new GetPluginsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPluginsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Plugin IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Plugin name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        /// <summary>
        /// The name of the plug-in that you want to create.
        /// </summary>
        [Input("pluginName")]
        public string? PluginName { get; set; }

        /// <summary>
        /// The type of the plug-in.
        /// </summary>
        [Input("pluginType")]
        public string? PluginType { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetPluginsArgs()
        {
        }
        public static new GetPluginsArgs Empty => new GetPluginsArgs();
    }

    public sealed class GetPluginsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Plugin IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Plugin name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        /// <summary>
        /// The name of the plug-in that you want to create.
        /// </summary>
        [Input("pluginName")]
        public Input<string>? PluginName { get; set; }

        /// <summary>
        /// The type of the plug-in.
        /// </summary>
        [Input("pluginType")]
        public Input<string>? PluginType { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetPluginsInvokeArgs()
        {
        }
        public static new GetPluginsInvokeArgs Empty => new GetPluginsInvokeArgs();
    }


    [OutputType]
    public sealed class GetPluginsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;
        public readonly string? PluginName;
        public readonly string? PluginType;
        public readonly ImmutableArray<Outputs.GetPluginsPluginResult> Plugins;
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetPluginsResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            int? pageNumber,

            int? pageSize,

            string? pluginName,

            string? pluginType,

            ImmutableArray<Outputs.GetPluginsPluginResult> plugins,

            ImmutableDictionary<string, object>? tags)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
            PluginName = pluginName;
            PluginType = pluginType;
            Plugins = plugins;
            Tags = tags;
        }
    }
}