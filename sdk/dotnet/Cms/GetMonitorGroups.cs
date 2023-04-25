// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms
{
    public static class GetMonitorGroups
    {
        /// <summary>
        /// This data source provides the Cms Monitor Groups of the current Alibaba Cloud user.
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
        ///     var example = AliCloud.Cms.GetMonitorGroups.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_value",
        ///         },
        ///         NameRegex = "the_resource_name",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstCmsMonitorGroupId"] = example.Apply(getMonitorGroupsResult =&gt; getMonitorGroupsResult.Groups[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMonitorGroupsResult> InvokeAsync(GetMonitorGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMonitorGroupsResult>("alicloud:cms/getMonitorGroups:getMonitorGroups", args ?? new GetMonitorGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Cms Monitor Groups of the current Alibaba Cloud user.
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
        ///     var example = AliCloud.Cms.GetMonitorGroups.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_value",
        ///         },
        ///         NameRegex = "the_resource_name",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstCmsMonitorGroupId"] = example.Apply(getMonitorGroupsResult =&gt; getMonitorGroupsResult.Groups[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetMonitorGroupsResult> Invoke(GetMonitorGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMonitorGroupsResult>("alicloud:cms/getMonitorGroups:getMonitorGroups", args ?? new GetMonitorGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMonitorGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the tag rule.
        /// </summary>
        [Input("dynamicTagRuleId")]
        public string? DynamicTagRuleId { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Monitor Group IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The include template history.
        /// </summary>
        [Input("includeTemplateHistory")]
        public bool? IncludeTemplateHistory { get; set; }

        /// <summary>
        /// The keyword to be matched.
        /// </summary>
        [Input("keyword")]
        public string? Keyword { get; set; }

        /// <summary>
        /// The name of the application group.
        /// </summary>
        [Input("monitorGroupName")]
        public string? MonitorGroupName { get; set; }

        /// <summary>
        /// A regex string to filter results by Monitor Group name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The select contact groups.
        /// </summary>
        [Input("selectContactGroups")]
        public bool? SelectContactGroups { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A map of tags assigned to the Cms Monitor Group.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the application group.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetMonitorGroupsArgs()
        {
        }
        public static new GetMonitorGroupsArgs Empty => new GetMonitorGroupsArgs();
    }

    public sealed class GetMonitorGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the tag rule.
        /// </summary>
        [Input("dynamicTagRuleId")]
        public Input<string>? DynamicTagRuleId { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Monitor Group IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The include template history.
        /// </summary>
        [Input("includeTemplateHistory")]
        public Input<bool>? IncludeTemplateHistory { get; set; }

        /// <summary>
        /// The keyword to be matched.
        /// </summary>
        [Input("keyword")]
        public Input<string>? Keyword { get; set; }

        /// <summary>
        /// The name of the application group.
        /// </summary>
        [Input("monitorGroupName")]
        public Input<string>? MonitorGroupName { get; set; }

        /// <summary>
        /// A regex string to filter results by Monitor Group name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The select contact groups.
        /// </summary>
        [Input("selectContactGroups")]
        public Input<bool>? SelectContactGroups { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags assigned to the Cms Monitor Group.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the application group.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetMonitorGroupsInvokeArgs()
        {
        }
        public static new GetMonitorGroupsInvokeArgs Empty => new GetMonitorGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetMonitorGroupsResult
    {
        public readonly string? DynamicTagRuleId;
        public readonly ImmutableArray<Outputs.GetMonitorGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly bool? IncludeTemplateHistory;
        public readonly string? Keyword;
        public readonly string? MonitorGroupName;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly bool? SelectContactGroups;
        public readonly ImmutableDictionary<string, object>? Tags;
        public readonly string? Type;

        [OutputConstructor]
        private GetMonitorGroupsResult(
            string? dynamicTagRuleId,

            ImmutableArray<Outputs.GetMonitorGroupsGroupResult> groups,

            string id,

            ImmutableArray<string> ids,

            bool? includeTemplateHistory,

            string? keyword,

            string? monitorGroupName,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            bool? selectContactGroups,

            ImmutableDictionary<string, object>? tags,

            string? type)
        {
            DynamicTagRuleId = dynamicTagRuleId;
            Groups = groups;
            Id = id;
            Ids = ids;
            IncludeTemplateHistory = includeTemplateHistory;
            Keyword = keyword;
            MonitorGroupName = monitorGroupName;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            SelectContactGroups = selectContactGroups;
            Tags = tags;
            Type = type;
        }
    }
}
