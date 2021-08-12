// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae
{
    public static class GetNamespaces
    {
        /// <summary>
        /// This data source provides the Sae Namespaces of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.129.0+.
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
        ///         var nameRegex = Output.Create(AliCloud.Sae.GetNamespaces.InvokeAsync(new AliCloud.Sae.GetNamespacesArgs
        ///         {
        ///             NameRegex = "^my-Namespace",
        ///         }));
        ///         this.SaeNamespaceId = nameRegex.Apply(nameRegex =&gt; nameRegex.Namespaces[0].Id);
        ///     }
        /// 
        ///     [Output("saeNamespaceId")]
        ///     public Output&lt;string&gt; SaeNamespaceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNamespacesResult> InvokeAsync(GetNamespacesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNamespacesResult>("alicloud:sae/getNamespaces:getNamespaces", args ?? new GetNamespacesArgs(), options.WithVersion());
    }


    public sealed class GetNamespacesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Namespace IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Namespace name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetNamespacesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNamespacesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly ImmutableArray<Outputs.GetNamespacesNamespaceResult> Namespaces;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetNamespacesResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            ImmutableArray<Outputs.GetNamespacesNamespaceResult> namespaces,

            string? outputFile)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            Namespaces = namespaces;
            OutputFile = outputFile;
        }
    }
}
