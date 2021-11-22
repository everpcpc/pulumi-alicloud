// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.ResourceManager
{
    public static class GetFolders
    {
        /// <summary>
        /// This data source provides the resource manager folders of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:**  Available in 1.84.0+.
        /// 
        /// &gt; **NOTE:**  You can view only the information of the first-level child folders of the specified folder.
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
        ///         var example = Output.Create(AliCloud.ResourceManager.GetFolders.InvokeAsync(new AliCloud.ResourceManager.GetFoldersArgs
        ///         {
        ///             NameRegex = "tftest",
        ///         }));
        ///         this.FirstFolderId = example.Apply(example =&gt; example.Folders?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstFolderId")]
        ///     public Output&lt;string&gt; FirstFolderId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFoldersResult> InvokeAsync(GetFoldersArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFoldersResult>("alicloud:resourcemanager/getFolders:getFolders", args ?? new GetFoldersArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the resource manager folders of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:**  Available in 1.84.0+.
        /// 
        /// &gt; **NOTE:**  You can view only the information of the first-level child folders of the specified folder.
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
        ///         var example = Output.Create(AliCloud.ResourceManager.GetFolders.InvokeAsync(new AliCloud.ResourceManager.GetFoldersArgs
        ///         {
        ///             NameRegex = "tftest",
        ///         }));
        ///         this.FirstFolderId = example.Apply(example =&gt; example.Folders?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstFolderId")]
        ///     public Output&lt;string&gt; FirstFolderId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFoldersResult> Invoke(GetFoldersInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFoldersResult>("alicloud:resourcemanager/getFolders:getFolders", args ?? new GetFoldersInvokeArgs(), options.WithVersion());
    }


    public sealed class GetFoldersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// -(Optional, Available in v1.114.0+) Default to `false`. Set it to true can output more details.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of resource manager folders IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by folder name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ID of the parent folder.
        /// </summary>
        [Input("parentFolderId")]
        public string? ParentFolderId { get; set; }

        /// <summary>
        /// The query keyword.
        /// </summary>
        [Input("queryKeyword")]
        public string? QueryKeyword { get; set; }

        public GetFoldersArgs()
        {
        }
    }

    public sealed class GetFoldersInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// -(Optional, Available in v1.114.0+) Default to `false`. Set it to true can output more details.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of resource manager folders IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by folder name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The ID of the parent folder.
        /// </summary>
        [Input("parentFolderId")]
        public Input<string>? ParentFolderId { get; set; }

        /// <summary>
        /// The query keyword.
        /// </summary>
        [Input("queryKeyword")]
        public Input<string>? QueryKeyword { get; set; }

        public GetFoldersInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFoldersResult
    {
        public readonly bool? EnableDetails;
        /// <summary>
        /// A list of folders. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFoldersFolderResult> Folders;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of folder IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of folder names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? ParentFolderId;
        public readonly string? QueryKeyword;

        [OutputConstructor]
        private GetFoldersResult(
            bool? enableDetails,

            ImmutableArray<Outputs.GetFoldersFolderResult> folders,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? parentFolderId,

            string? queryKeyword)
        {
            EnableDetails = enableDetails;
            Folders = folders;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ParentFolderId = parentFolderId;
            QueryKeyword = queryKeyword;
        }
    }
}
