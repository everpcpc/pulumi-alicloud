// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ServiceCatalog
{
    public static class GetLaunchOptions
    {
        /// <summary>
        /// This data source provides Service Catalog Launch Option available to the user.[What is Launch Option](https://www.alibabacloud.com/help/en/servicecatalog/latest/api-doc-servicecatalog-2021-09-01-api-doc-listlaunchoptions)
        /// 
        /// &gt; **NOTE:** Available in 1.196.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var defaultEndUserProducts = AliCloud.ServiceCatalog.GetEndUserProducts.Invoke(new()
        ///     {
        ///         NameRegex = "ram模板创建",
        ///     });
        /// 
        ///     var defaultLaunchOptions = AliCloud.ServiceCatalog.GetLaunchOptions.Invoke(new()
        ///     {
        ///         ProductId = "data.alicloud_service_catalog_end_user_products.default.end_user_products.0.id",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudServiceCatalogLaunchOptionExampleId"] = defaultLaunchOptions.Apply(getLaunchOptionsResult =&gt; getLaunchOptionsResult.LaunchOptions[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLaunchOptionsResult> InvokeAsync(GetLaunchOptionsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLaunchOptionsResult>("alicloud:servicecatalog/getLaunchOptions:getLaunchOptions", args ?? new GetLaunchOptionsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides Service Catalog Launch Option available to the user.[What is Launch Option](https://www.alibabacloud.com/help/en/servicecatalog/latest/api-doc-servicecatalog-2021-09-01-api-doc-listlaunchoptions)
        /// 
        /// &gt; **NOTE:** Available in 1.196.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var defaultEndUserProducts = AliCloud.ServiceCatalog.GetEndUserProducts.Invoke(new()
        ///     {
        ///         NameRegex = "ram模板创建",
        ///     });
        /// 
        ///     var defaultLaunchOptions = AliCloud.ServiceCatalog.GetLaunchOptions.Invoke(new()
        ///     {
        ///         ProductId = "data.alicloud_service_catalog_end_user_products.default.end_user_products.0.id",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudServiceCatalogLaunchOptionExampleId"] = defaultLaunchOptions.Apply(getLaunchOptionsResult =&gt; getLaunchOptionsResult.LaunchOptions[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetLaunchOptionsResult> Invoke(GetLaunchOptionsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLaunchOptionsResult>("alicloud:servicecatalog/getLaunchOptions:getLaunchOptions", args ?? new GetLaunchOptionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLaunchOptionsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by portfolio name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Product ID.
        /// </summary>
        [Input("productId", required: true)]
        public string ProductId { get; set; } = null!;

        public GetLaunchOptionsArgs()
        {
        }
        public static new GetLaunchOptionsArgs Empty => new GetLaunchOptionsArgs();
    }

    public sealed class GetLaunchOptionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by portfolio name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Product ID.
        /// </summary>
        [Input("productId", required: true)]
        public Input<string> ProductId { get; set; } = null!;

        public GetLaunchOptionsInvokeArgs()
        {
        }
        public static new GetLaunchOptionsInvokeArgs Empty => new GetLaunchOptionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetLaunchOptionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of Launch Option Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLaunchOptionsLaunchOptionResult> LaunchOptions;
        public readonly string? NameRegex;
        public readonly ImmutableArray<Outputs.GetLaunchOptionsOptionResult> Options;
        public readonly string? OutputFile;
        public readonly string ProductId;

        [OutputConstructor]
        private GetLaunchOptionsResult(
            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetLaunchOptionsLaunchOptionResult> launchOptions,

            string? nameRegex,

            ImmutableArray<Outputs.GetLaunchOptionsOptionResult> options,

            string? outputFile,

            string productId)
        {
            Id = id;
            Ids = ids;
            LaunchOptions = launchOptions;
            NameRegex = nameRegex;
            Options = options;
            OutputFile = outputFile;
            ProductId = productId;
        }
    }
}
