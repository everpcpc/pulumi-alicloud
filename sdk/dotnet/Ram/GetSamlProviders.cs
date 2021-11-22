// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.Ram
{
    public static class GetSamlProviders
    {
        /// <summary>
        /// This data source provides the Ram Saml Providers of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.114.0+.
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
        ///         var example = Output.Create(AliCloud.Ram.GetSamlProviders.InvokeAsync(new AliCloud.Ram.GetSamlProvidersArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "samlProviderName",
        ///             },
        ///             NameRegex = "tf-testAcc",
        ///         }));
        ///         this.FirstRamSamlProviderId = example.Apply(example =&gt; example.Providers?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstRamSamlProviderId")]
        ///     public Output&lt;string&gt; FirstRamSamlProviderId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSamlProvidersResult> InvokeAsync(GetSamlProvidersArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSamlProvidersResult>("alicloud:ram/getSamlProviders:getSamlProviders", args ?? new GetSamlProvidersArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the Ram Saml Providers of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.114.0+.
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
        ///         var example = Output.Create(AliCloud.Ram.GetSamlProviders.InvokeAsync(new AliCloud.Ram.GetSamlProvidersArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "samlProviderName",
        ///             },
        ///             NameRegex = "tf-testAcc",
        ///         }));
        ///         this.FirstRamSamlProviderId = example.Apply(example =&gt; example.Providers?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstRamSamlProviderId")]
        ///     public Output&lt;string&gt; FirstRamSamlProviderId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSamlProvidersResult> Invoke(GetSamlProvidersInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSamlProvidersResult>("alicloud:ram/getSamlProviders:getSamlProviders", args ?? new GetSamlProvidersInvokeArgs(), options.WithVersion());
    }


    public sealed class GetSamlProvidersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of SAML Provider IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by SAML Provider name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetSamlProvidersArgs()
        {
        }
    }

    public sealed class GetSamlProvidersInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of SAML Provider IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by SAML Provider name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetSamlProvidersInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSamlProvidersResult
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
        public readonly ImmutableArray<Outputs.GetSamlProvidersProviderResult> Providers;

        [OutputConstructor]
        private GetSamlProvidersResult(
            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetSamlProvidersProviderResult> providers)
        {
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Providers = providers;
        }
    }
}
