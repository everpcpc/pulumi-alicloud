// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Kms
{
    public static class GetSecrets
    {
        /// <summary>
        /// This data source provides a list of KMS Secrets in an Alibaba Cloud account according to the specified filters.
        ///  
        /// &gt; **NOTE:** Available in v1.86.0+.
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
        ///     var kmsSecretsDs = AliCloud.Kms.GetSecrets.Invoke(new()
        ///     {
        ///         FetchTags = true,
        ///         NameRegex = "name_regex",
        ///         Tags = 
        ///         {
        ///             { "k-aa", "v-aa" },
        ///             { "k-bb", "v-bb" },
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstSecretId"] = kmsSecretsDs.Apply(getSecretsResult =&gt; getSecretsResult.Secrets[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSecretsResult> InvokeAsync(GetSecretsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecretsResult>("alicloud:kms/getSecrets:getSecrets", args ?? new GetSecretsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides a list of KMS Secrets in an Alibaba Cloud account according to the specified filters.
        ///  
        /// &gt; **NOTE:** Available in v1.86.0+.
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
        ///     var kmsSecretsDs = AliCloud.Kms.GetSecrets.Invoke(new()
        ///     {
        ///         FetchTags = true,
        ///         NameRegex = "name_regex",
        ///         Tags = 
        ///         {
        ///             { "k-aa", "v-aa" },
        ///             { "k-bb", "v-bb" },
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstSecretId"] = kmsSecretsDs.Apply(getSecretsResult =&gt; getSecretsResult.Secrets[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSecretsResult> Invoke(GetSecretsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecretsResult>("alicloud:kms/getSecrets:getSecrets", args ?? new GetSecretsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecretsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to true can output more details.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        /// <summary>
        /// Whether to include the predetermined resource tag in the return value. Default to `false`.
        /// </summary>
        [Input("fetchTags")]
        public bool? FetchTags { get; set; }

        /// <summary>
        /// The secret filter. The filter consists of one or more key-value pairs. 
        /// More details see API [ListSecrets](https://www.alibabacloud.com/help/en/key-management-service/latest/listsecrets).
        /// </summary>
        [Input("filters")]
        public string? Filters { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of KMS Secret ids. The value is same as KMS secret_name.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter the results by the KMS secret_name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

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

        public GetSecretsArgs()
        {
        }
        public static new GetSecretsArgs Empty => new GetSecretsArgs();
    }

    public sealed class GetSecretsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to true can output more details.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        /// <summary>
        /// Whether to include the predetermined resource tag in the return value. Default to `false`.
        /// </summary>
        [Input("fetchTags")]
        public Input<bool>? FetchTags { get; set; }

        /// <summary>
        /// The secret filter. The filter consists of one or more key-value pairs. 
        /// More details see API [ListSecrets](https://www.alibabacloud.com/help/en/key-management-service/latest/listsecrets).
        /// </summary>
        [Input("filters")]
        public Input<string>? Filters { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of KMS Secret ids. The value is same as KMS secret_name.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter the results by the KMS secret_name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetSecretsInvokeArgs()
        {
        }
        public static new GetSecretsInvokeArgs Empty => new GetSecretsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecretsResult
    {
        public readonly bool? EnableDetails;
        public readonly bool? FetchTags;
        public readonly string? Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Kms Secret ids. The value is same as KMS secret_name.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of KMS Secret names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of KMS Secrets. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSecretsSecretResult> Secrets;
        /// <summary>
        /// (Optional) A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetSecretsResult(
            bool? enableDetails,

            bool? fetchTags,

            string? filters,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetSecretsSecretResult> secrets,

            ImmutableDictionary<string, object>? tags)
        {
            EnableDetails = enableDetails;
            FetchTags = fetchTags;
            Filters = filters;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Secrets = secrets;
            Tags = tags;
        }
    }
}
