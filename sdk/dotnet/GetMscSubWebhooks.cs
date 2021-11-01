// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud
{
    public static class GetMscSubWebhooks
    {
        /// <summary>
        /// This data source provides the Msc Sub Webhooks of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.141.0+.
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
        ///         var ids = Output.Create(AliCloud.GetMscSubWebhooks.InvokeAsync(new AliCloud.GetMscSubWebhooksArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_id",
        ///             },
        ///         }));
        ///         this.MscSubWebhookId1 = ids.Apply(ids =&gt; ids.Webhooks[0].Id);
        ///         var nameRegex = Output.Create(AliCloud.GetMscSubWebhooks.InvokeAsync(new AliCloud.GetMscSubWebhooksArgs
        ///         {
        ///             NameRegex = "^my-Webhook",
        ///         }));
        ///         this.MscSubWebhookId2 = nameRegex.Apply(nameRegex =&gt; nameRegex.Webhooks[0].Id);
        ///     }
        /// 
        ///     [Output("mscSubWebhookId1")]
        ///     public Output&lt;string&gt; MscSubWebhookId1 { get; set; }
        ///     [Output("mscSubWebhookId2")]
        ///     public Output&lt;string&gt; MscSubWebhookId2 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMscSubWebhooksResult> InvokeAsync(GetMscSubWebhooksArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMscSubWebhooksResult>("alicloud:index/getMscSubWebhooks:getMscSubWebhooks", args ?? new GetMscSubWebhooksArgs(), options.WithVersion());
    }


    public sealed class GetMscSubWebhooksArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Webhook IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Webhook name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetMscSubWebhooksArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMscSubWebhooksResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetMscSubWebhooksWebhookResult> Webhooks;

        [OutputConstructor]
        private GetMscSubWebhooksResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetMscSubWebhooksWebhookResult> webhooks)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Webhooks = webhooks;
        }
    }
}
