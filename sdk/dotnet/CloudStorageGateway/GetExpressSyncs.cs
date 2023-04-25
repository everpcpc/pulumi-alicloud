// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CloudStorageGateway
{
    public static class GetExpressSyncs
    {
        /// <summary>
        /// This data source provides the Cloud Storage Gateway Express Syncs of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.144.0+.
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
        ///     var ids = AliCloud.CloudStorageGateway.GetExpressSyncs.Invoke();
        /// 
        ///     var nameRegex = AliCloud.CloudStorageGateway.GetExpressSyncs.Invoke(new()
        ///     {
        ///         NameRegex = "^my-ExpressSync",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cloudStorageGatewayExpressSyncId1"] = ids.Apply(getExpressSyncsResult =&gt; getExpressSyncsResult.Syncs[0]?.Id),
        ///         ["cloudStorageGatewayExpressSyncId2"] = nameRegex.Apply(getExpressSyncsResult =&gt; getExpressSyncsResult.Syncs[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetExpressSyncsResult> InvokeAsync(GetExpressSyncsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetExpressSyncsResult>("alicloud:cloudstoragegateway/getExpressSyncs:getExpressSyncs", args ?? new GetExpressSyncsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Cloud Storage Gateway Express Syncs of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.144.0+.
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
        ///     var ids = AliCloud.CloudStorageGateway.GetExpressSyncs.Invoke();
        /// 
        ///     var nameRegex = AliCloud.CloudStorageGateway.GetExpressSyncs.Invoke(new()
        ///     {
        ///         NameRegex = "^my-ExpressSync",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cloudStorageGatewayExpressSyncId1"] = ids.Apply(getExpressSyncsResult =&gt; getExpressSyncsResult.Syncs[0]?.Id),
        ///         ["cloudStorageGatewayExpressSyncId2"] = nameRegex.Apply(getExpressSyncsResult =&gt; getExpressSyncsResult.Syncs[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetExpressSyncsResult> Invoke(GetExpressSyncsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetExpressSyncsResult>("alicloud:cloudstoragegateway/getExpressSyncs:getExpressSyncs", args ?? new GetExpressSyncsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExpressSyncsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Express Sync IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Express Sync name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetExpressSyncsArgs()
        {
        }
        public static new GetExpressSyncsArgs Empty => new GetExpressSyncsArgs();
    }

    public sealed class GetExpressSyncsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Express Sync IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Express Sync name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetExpressSyncsInvokeArgs()
        {
        }
        public static new GetExpressSyncsInvokeArgs Empty => new GetExpressSyncsInvokeArgs();
    }


    [OutputType]
    public sealed class GetExpressSyncsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetExpressSyncsSyncResult> Syncs;

        [OutputConstructor]
        private GetExpressSyncsResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetExpressSyncsSyncResult> syncs)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Syncs = syncs;
        }
    }
}
