// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.DataWorks
{
    /// <summary>
    /// Provides a Data Works Folder resource.
    /// 
    /// For information about Data Works Folder and how to use it, see [What is Folder](https://help.aliyun.com/document_detail/173940.html).
    /// 
    /// &gt; **NOTE:** Available in v1.131.0+.
    /// 
    /// ## Example Usage
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
    ///     var example = new AliCloud.DataWorks.Folder("example", new()
    ///     {
    ///         FolderPath = "Business Flow/tfTestAcc/folderDi/tftest1",
    ///         ProjectId = "320687",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Data Works Folder can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:dataworks/folder:Folder example &lt;folder_id&gt;:&lt;$.ProjectId&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dataworks/folder:Folder")]
    public partial class Folder : global::Pulumi.CustomResource
    {
        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// Folder Path. The folder path composed with for part: `Business Flow/{Business Flow Name}/[folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined]/{Directory Name}`. The first segment of path must be `Business Flow`, and sencond segment of path must be a Business Flow Name within the project. The third part of path must be one of those keywords:`folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined`. Then the finial part of folder path can be specified in yourself.
        /// </summary>
        [Output("folderPath")]
        public Output<string> FolderPath { get; private set; } = null!;

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;

        [Output("projectIdentifier")]
        public Output<string?> ProjectIdentifier { get; private set; } = null!;


        /// <summary>
        /// Create a Folder resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Folder(string name, FolderArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dataworks/folder:Folder", name, args ?? new FolderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Folder(string name, Input<string> id, FolderState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dataworks/folder:Folder", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Folder resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Folder Get(string name, Input<string> id, FolderState? state = null, CustomResourceOptions? options = null)
        {
            return new Folder(name, id, state, options);
        }
    }

    public sealed class FolderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Folder Path. The folder path composed with for part: `Business Flow/{Business Flow Name}/[folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined]/{Directory Name}`. The first segment of path must be `Business Flow`, and sencond segment of path must be a Business Flow Name within the project. The third part of path must be one of those keywords:`folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined`. Then the finial part of folder path can be specified in yourself.
        /// </summary>
        [Input("folderPath", required: true)]
        public Input<string> FolderPath { get; set; } = null!;

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("projectIdentifier")]
        public Input<string>? ProjectIdentifier { get; set; }

        public FolderArgs()
        {
        }
        public static new FolderArgs Empty => new FolderArgs();
    }

    public sealed class FolderState : global::Pulumi.ResourceArgs
    {
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// Folder Path. The folder path composed with for part: `Business Flow/{Business Flow Name}/[folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined]/{Directory Name}`. The first segment of path must be `Business Flow`, and sencond segment of path must be a Business Flow Name within the project. The third part of path must be one of those keywords:`folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined`. Then the finial part of folder path can be specified in yourself.
        /// </summary>
        [Input("folderPath")]
        public Input<string>? FolderPath { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("projectIdentifier")]
        public Input<string>? ProjectIdentifier { get; set; }

        public FolderState()
        {
        }
        public static new FolderState Empty => new FolderState();
    }
}
