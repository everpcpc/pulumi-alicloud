// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.Ros
{
    public static class GetTemplates
    {
        /// <summary>
        /// This data source provides the Ros Templates of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.108.0+.
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
        ///         var example = Output.Create(AliCloud.Ros.GetTemplates.InvokeAsync(new AliCloud.Ros.GetTemplatesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_value",
        ///             },
        ///             NameRegex = "the_resource_name",
        ///         }));
        ///         this.FirstRosTemplateId = example.Apply(example =&gt; example.Templates?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstRosTemplateId")]
        ///     public Output&lt;string&gt; FirstRosTemplateId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTemplatesResult> InvokeAsync(GetTemplatesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTemplatesResult>("alicloud:ros/getTemplates:getTemplates", args ?? new GetTemplatesArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the Ros Templates of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.108.0+.
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
        ///         var example = Output.Create(AliCloud.Ros.GetTemplates.InvokeAsync(new AliCloud.Ros.GetTemplatesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_value",
        ///             },
        ///             NameRegex = "the_resource_name",
        ///         }));
        ///         this.FirstRosTemplateId = example.Apply(example =&gt; example.Templates?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstRosTemplateId")]
        ///     public Output&lt;string&gt; FirstRosTemplateId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTemplatesResult> Invoke(GetTemplatesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTemplatesResult>("alicloud:ros/getTemplates:getTemplates", args ?? new GetTemplatesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetTemplatesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Template IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Template name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Share Type.
        /// </summary>
        [Input("shareType")]
        public string? ShareType { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the template.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
        /// </summary>
        [Input("templateName")]
        public string? TemplateName { get; set; }

        public GetTemplatesArgs()
        {
        }
    }

    public sealed class GetTemplatesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Template IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Template name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Share Type.
        /// </summary>
        [Input("shareType")]
        public Input<string>? ShareType { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the template.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        public GetTemplatesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTemplatesResult
    {
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? ShareType;
        public readonly ImmutableDictionary<string, object>? Tags;
        public readonly string? TemplateName;
        public readonly ImmutableArray<Outputs.GetTemplatesTemplateResult> Templates;

        [OutputConstructor]
        private GetTemplatesResult(
            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? shareType,

            ImmutableDictionary<string, object>? tags,

            string? templateName,

            ImmutableArray<Outputs.GetTemplatesTemplateResult> templates)
        {
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ShareType = shareType;
            Tags = tags;
            TemplateName = templateName;
            Templates = templates;
        }
    }
}
