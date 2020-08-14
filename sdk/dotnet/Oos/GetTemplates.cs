// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oos
{
    public static class GetTemplates
    {
        /// <summary>
        /// This data source provides a list of OOS Templates in an Alibaba Cloud account according to the specified filters.
        ///  
        /// &gt; **NOTE:** Available in v1.92.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(AliCloud.Oos.GetTemplates.InvokeAsync(new AliCloud.Oos.GetTemplatesArgs
        ///         {
        ///             HasTrigger = false,
        ///             NameRegex = "test",
        ///             ShareType = "Private",
        ///             Tags = 
        ///             {
        ///                 { "Created", "TF" },
        ///                 { "For", "template Test" },
        ///             },
        ///         }));
        ///         this.FirstTemplateName = example.Apply(example =&gt; example.Templates[0].TemplateName);
        ///     }
        /// 
        ///     [Output("firstTemplateName")]
        ///     public Output&lt;string&gt; FirstTemplateName { get; set; }
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTemplatesResult> InvokeAsync(GetTemplatesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTemplatesResult>("alicloud:oos/getTemplates:getTemplates", args ?? new GetTemplatesArgs(), options.WithVersion());
    }


    public sealed class GetTemplatesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The category of template.
        /// </summary>
        [Input("category")]
        public string? Category { get; set; }

        /// <summary>
        /// The creator of the template.
        /// </summary>
        [Input("createdBy")]
        public string? CreatedBy { get; set; }

        /// <summary>
        /// The template whose creation time is less than or equal to the specified time. The format is: YYYY-MM-DDThh:mm::ssZ.
        /// </summary>
        [Input("createdDate")]
        public string? CreatedDate { get; set; }

        /// <summary>
        /// Create a template whose time is greater than or equal to the specified time. The format is: YYYY-MM-DDThh:mm:ssZ.
        /// </summary>
        [Input("createdDateAfter")]
        public string? CreatedDateAfter { get; set; }

        /// <summary>
        /// Is it triggered successfully.
        /// </summary>
        [Input("hasTrigger")]
        public bool? HasTrigger { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of OOS Template ids. Each element in the list is same as template_name.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter the results by the template_name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The sharing type of the template. Valid values: `Private`, `Public`.
        /// </summary>
        [Input("shareType")]
        public string? ShareType { get; set; }

        /// <summary>
        /// Sort field. Valid values: `TotalExecutionCount`, `Popularity`, `TemplateName` and `CreatedDate`. Default to `TotalExecutionCount`.
        /// </summary>
        [Input("sortField")]
        public string? SortField { get; set; }

        /// <summary>
        /// Sort order. Valid values: `Ascending`, `Descending`. Default to `Descending`
        /// </summary>
        [Input("sortOrder")]
        public string? SortOrder { get; set; }

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

        /// <summary>
        /// The format of the template. Valid values: `JSON`, `YAML`.
        /// </summary>
        [Input("templateFormat")]
        public string? TemplateFormat { get; set; }

        /// <summary>
        /// The type of OOS Template.
        /// </summary>
        [Input("templateType")]
        public string? TemplateType { get; set; }

        public GetTemplatesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTemplatesResult
    {
        public readonly string? Category;
        public readonly string? CreatedBy;
        public readonly string? CreatedDate;
        public readonly string? CreatedDateAfter;
        public readonly bool? HasTrigger;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of OOS Template ids. Each element in the list is same as template_name.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        public readonly string? ShareType;
        public readonly string? SortField;
        public readonly string? SortOrder;
        public readonly ImmutableDictionary<string, object>? Tags;
        public readonly string? TemplateFormat;
        public readonly string? TemplateType;
        /// <summary>
        /// A list of OOS Templates. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTemplatesTemplateResult> Templates;

        [OutputConstructor]
        private GetTemplatesResult(
            string? category,

            string? createdBy,

            string? createdDate,

            string? createdDateAfter,

            bool? hasTrigger,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            string? shareType,

            string? sortField,

            string? sortOrder,

            ImmutableDictionary<string, object>? tags,

            string? templateFormat,

            string? templateType,

            ImmutableArray<Outputs.GetTemplatesTemplateResult> templates)
        {
            Category = category;
            CreatedBy = createdBy;
            CreatedDate = createdDate;
            CreatedDateAfter = createdDateAfter;
            HasTrigger = hasTrigger;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            ShareType = shareType;
            SortField = sortField;
            SortOrder = sortOrder;
            Tags = tags;
            TemplateFormat = templateFormat;
            TemplateType = templateType;
            Templates = templates;
        }
    }
}
