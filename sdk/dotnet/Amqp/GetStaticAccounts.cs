// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Amqp
{
    public static class GetStaticAccounts
    {
        /// <summary>
        /// This data source provides Amqp Static Account available to the user.[What is Static Account](https://help.aliyun.com/document_detail/184399.html)
        /// 
        /// &gt; **NOTE:** Available in 1.195.0+
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
        ///     var @default = AliCloud.Amqp.GetStaticAccounts.Invoke(new()
        ///     {
        ///         InstanceId = "amqp-cn-0ju2y01zs001",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudAmqpStaticAccountExampleId"] = @default.Apply(@default =&gt; @default.Apply(getStaticAccountsResult =&gt; getStaticAccountsResult.Accounts[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetStaticAccountsResult> InvokeAsync(GetStaticAccountsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStaticAccountsResult>("alicloud:amqp/getStaticAccounts:getStaticAccounts", args ?? new GetStaticAccountsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides Amqp Static Account available to the user.[What is Static Account](https://help.aliyun.com/document_detail/184399.html)
        /// 
        /// &gt; **NOTE:** Available in 1.195.0+
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
        ///     var @default = AliCloud.Amqp.GetStaticAccounts.Invoke(new()
        ///     {
        ///         InstanceId = "amqp-cn-0ju2y01zs001",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudAmqpStaticAccountExampleId"] = @default.Apply(@default =&gt; @default.Apply(getStaticAccountsResult =&gt; getStaticAccountsResult.Accounts[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetStaticAccountsResult> Invoke(GetStaticAccountsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStaticAccountsResult>("alicloud:amqp/getStaticAccounts:getStaticAccounts", args ?? new GetStaticAccountsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStaticAccountsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// The `key` of the resource supplied above.The value is formulated as `&lt;instance_id&gt;:&lt;access_key&gt;`.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// InstanceId
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetStaticAccountsArgs()
        {
        }
        public static new GetStaticAccountsArgs Empty => new GetStaticAccountsArgs();
    }

    public sealed class GetStaticAccountsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// The `key` of the resource supplied above.The value is formulated as `&lt;instance_id&gt;:&lt;access_key&gt;`.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// InstanceId
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetStaticAccountsInvokeArgs()
        {
        }
        public static new GetStaticAccountsInvokeArgs Empty => new GetStaticAccountsInvokeArgs();
    }


    [OutputType]
    public sealed class GetStaticAccountsResult
    {
        /// <summary>
        /// A list of Static Account Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetStaticAccountsAccountResult> Accounts;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// Amqp instance ID.
        /// </summary>
        public readonly string? InstanceId;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetStaticAccountsResult(
            ImmutableArray<Outputs.GetStaticAccountsAccountResult> accounts,

            string id,

            ImmutableArray<string> ids,

            string? instanceId,

            string? outputFile)
        {
            Accounts = accounts;
            Id = id;
            Ids = ids;
            InstanceId = instanceId;
            OutputFile = outputFile;
        }
    }
}
