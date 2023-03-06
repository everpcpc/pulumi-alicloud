// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ServiceCatalog
{
    public static class GetProvisionedProducts
    {
        /// <summary>
        /// This data source provides Service Catalog Provisioned Product available to the user.[What is Provisioned Product](https://www.alibabacloud.com/help/en/servicecatalog/latest/api-doc-servicecatalog-2021-09-01-api-doc-launchproduct)
        /// 
        /// &gt; **NOTE:** Available in 1.196.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = AliCloud.ServiceCatalog.GetProvisionedProducts.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "IdExample",
        ///         },
        ///         NameRegex = "NameRegexExample",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudServiceCatalogProvisionedProductExampleId"] = @default.Apply(@default =&gt; @default.Apply(getProvisionedProductsResult =&gt; getProvisionedProductsResult.ProvisionedProducts[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProvisionedProductsResult> InvokeAsync(GetProvisionedProductsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProvisionedProductsResult>("alicloud:servicecatalog/getProvisionedProducts:getProvisionedProducts", args ?? new GetProvisionedProductsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides Service Catalog Provisioned Product available to the user.[What is Provisioned Product](https://www.alibabacloud.com/help/en/servicecatalog/latest/api-doc-servicecatalog-2021-09-01-api-doc-launchproduct)
        /// 
        /// &gt; **NOTE:** Available in 1.196.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = AliCloud.ServiceCatalog.GetProvisionedProducts.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "IdExample",
        ///         },
        ///         NameRegex = "NameRegexExample",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudServiceCatalogProvisionedProductExampleId"] = @default.Apply(@default =&gt; @default.Apply(getProvisionedProductsResult =&gt; getProvisionedProductsResult.ProvisionedProducts[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProvisionedProductsResult> Invoke(GetProvisionedProductsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProvisionedProductsResult>("alicloud:servicecatalog/getProvisionedProducts:getProvisionedProducts", args ?? new GetProvisionedProductsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProvisionedProductsArgs : global::Pulumi.InvokeArgs
    {
        [Input("accessLevelFilter")]
        public string? AccessLevelFilter { get; set; }

        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Provisioned Product IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Product name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        [Input("sortBy")]
        public string? SortBy { get; set; }

        [Input("sortOrder")]
        public string? SortOrder { get; set; }

        public GetProvisionedProductsArgs()
        {
        }
        public static new GetProvisionedProductsArgs Empty => new GetProvisionedProductsArgs();
    }

    public sealed class GetProvisionedProductsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("accessLevelFilter")]
        public Input<string>? AccessLevelFilter { get; set; }

        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Provisioned Product IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Product name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        [Input("sortBy")]
        public Input<string>? SortBy { get; set; }

        [Input("sortOrder")]
        public Input<string>? SortOrder { get; set; }

        public GetProvisionedProductsInvokeArgs()
        {
        }
        public static new GetProvisionedProductsInvokeArgs Empty => new GetProvisionedProductsInvokeArgs();
    }


    [OutputType]
    public sealed class GetProvisionedProductsResult
    {
        public readonly string? AccessLevelFilter;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Provisioned Product IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of name of Provisioned Products.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;
        public readonly ImmutableArray<Outputs.GetProvisionedProductsProductResult> Products;
        /// <summary>
        /// A list of Provisioned Product Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProvisionedProductsProvisionedProductResult> ProvisionedProducts;
        public readonly string? SortBy;
        public readonly string? SortOrder;

        [OutputConstructor]
        private GetProvisionedProductsResult(
            string? accessLevelFilter,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            int? pageNumber,

            int? pageSize,

            ImmutableArray<Outputs.GetProvisionedProductsProductResult> products,

            ImmutableArray<Outputs.GetProvisionedProductsProvisionedProductResult> provisionedProducts,

            string? sortBy,

            string? sortOrder)
        {
            AccessLevelFilter = accessLevelFilter;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
            Products = products;
            ProvisionedProducts = provisionedProducts;
            SortBy = sortBy;
            SortOrder = sortOrder;
        }
    }
}