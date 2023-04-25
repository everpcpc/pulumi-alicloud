// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms
{
    public static class GetSlsGroups
    {
        /// <summary>
        /// This data source provides the Cms Sls Groups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.171.0+.
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
        ///     var ids = AliCloud.Cms.GetSlsGroups.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Cms.GetSlsGroups.Invoke(new()
        ///     {
        ///         NameRegex = "^my-SlsGroup",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cmsSlsGroupId1"] = ids.Apply(getSlsGroupsResult =&gt; getSlsGroupsResult.Groups[0]?.Id),
        ///         ["cmsSlsGroupId2"] = nameRegex.Apply(getSlsGroupsResult =&gt; getSlsGroupsResult.Groups[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSlsGroupsResult> InvokeAsync(GetSlsGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSlsGroupsResult>("alicloud:cms/getSlsGroups:getSlsGroups", args ?? new GetSlsGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Cms Sls Groups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.171.0+.
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
        ///     var ids = AliCloud.Cms.GetSlsGroups.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Cms.GetSlsGroups.Invoke(new()
        ///     {
        ///         NameRegex = "^my-SlsGroup",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cmsSlsGroupId1"] = ids.Apply(getSlsGroupsResult =&gt; getSlsGroupsResult.Groups[0]?.Id),
        ///         ["cmsSlsGroupId2"] = nameRegex.Apply(getSlsGroupsResult =&gt; getSlsGroupsResult.Groups[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSlsGroupsResult> Invoke(GetSlsGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSlsGroupsResult>("alicloud:cms/getSlsGroups:getSlsGroups", args ?? new GetSlsGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSlsGroupsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Sls Group IDs. Its element value is same as Sls Group Name.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The keywords of the `sls_group_name` or `sls_group_description` of the Sls Group.
        /// </summary>
        [Input("keyword")]
        public string? Keyword { get; set; }

        /// <summary>
        /// A regex string to filter results by Sls Group name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        public GetSlsGroupsArgs()
        {
        }
        public static new GetSlsGroupsArgs Empty => new GetSlsGroupsArgs();
    }

    public sealed class GetSlsGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Sls Group IDs. Its element value is same as Sls Group Name.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The keywords of the `sls_group_name` or `sls_group_description` of the Sls Group.
        /// </summary>
        [Input("keyword")]
        public Input<string>? Keyword { get; set; }

        /// <summary>
        /// A regex string to filter results by Sls Group name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        public GetSlsGroupsInvokeArgs()
        {
        }
        public static new GetSlsGroupsInvokeArgs Empty => new GetSlsGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSlsGroupsResult
    {
        public readonly ImmutableArray<Outputs.GetSlsGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? Keyword;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;

        [OutputConstructor]
        private GetSlsGroupsResult(
            ImmutableArray<Outputs.GetSlsGroupsGroupResult> groups,

            string id,

            ImmutableArray<string> ids,

            string? keyword,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            int? pageNumber,

            int? pageSize)
        {
            Groups = groups;
            Id = id;
            Ids = ids;
            Keyword = keyword;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
        }
    }
}
