// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae
{
    public static class GetGreyTagRoutes
    {
        /// <summary>
        /// This data source provides the Sae GreyTagRoutes of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.160.0+.
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
        ///         var nameRegex = Output.Create(AliCloud.Sae.GetGreyTagRoutes.InvokeAsync(new AliCloud.Sae.GetGreyTagRoutesArgs
        ///         {
        ///             AppId = "example_id",
        ///             NameRegex = "^my-GreyTagRoute",
        ///         }));
        ///         this.SaeGreyTagRoutesId = nameRegex.Apply(nameRegex =&gt; nameRegex.Routes?[0]?.Id);
        ///     }
        /// 
        ///     [Output("saeGreyTagRoutesId")]
        ///     public Output&lt;string&gt; SaeGreyTagRoutesId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGreyTagRoutesResult> InvokeAsync(GetGreyTagRoutesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGreyTagRoutesResult>("alicloud:sae/getGreyTagRoutes:getGreyTagRoutes", args ?? new GetGreyTagRoutesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Sae GreyTagRoutes of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.160.0+.
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
        ///         var nameRegex = Output.Create(AliCloud.Sae.GetGreyTagRoutes.InvokeAsync(new AliCloud.Sae.GetGreyTagRoutesArgs
        ///         {
        ///             AppId = "example_id",
        ///             NameRegex = "^my-GreyTagRoute",
        ///         }));
        ///         this.SaeGreyTagRoutesId = nameRegex.Apply(nameRegex =&gt; nameRegex.Routes?[0]?.Id);
        ///     }
        /// 
        ///     [Output("saeGreyTagRoutesId")]
        ///     public Output&lt;string&gt; SaeGreyTagRoutesId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetGreyTagRoutesResult> Invoke(GetGreyTagRoutesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGreyTagRoutesResult>("alicloud:sae/getGreyTagRoutes:getGreyTagRoutes", args ?? new GetGreyTagRoutesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGreyTagRoutesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID  of the SAE Application.
        /// </summary>
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of GreyTagRoute IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by GreyTagRoute name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetGreyTagRoutesArgs()
        {
        }
    }

    public sealed class GetGreyTagRoutesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID  of the SAE Application.
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of GreyTagRoute IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by GreyTagRoute name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetGreyTagRoutesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGreyTagRoutesResult
    {
        public readonly string AppId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetGreyTagRoutesRouteResult> Routes;

        [OutputConstructor]
        private GetGreyTagRoutesResult(
            string appId,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetGreyTagRoutesRouteResult> routes)
        {
            AppId = appId;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Routes = routes;
        }
    }
}
