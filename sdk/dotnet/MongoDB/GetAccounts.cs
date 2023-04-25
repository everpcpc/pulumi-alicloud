// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.MongoDB
{
    public static class GetAccounts
    {
        /// <summary>
        /// This data source provides the Mongodb Accounts of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.148.0+.
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
        ///     var example = AliCloud.MongoDB.GetAccounts.Invoke(new()
        ///     {
        ///         InstanceId = "example_value",
        ///         AccountName = "root",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["mongodbAccountId1"] = example.Apply(getAccountsResult =&gt; getAccountsResult.Accounts[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccountsResult> InvokeAsync(GetAccountsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountsResult>("alicloud:mongodb/getAccounts:getAccounts", args ?? new GetAccountsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Mongodb Accounts of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.148.0+.
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
        ///     var example = AliCloud.MongoDB.GetAccounts.Invoke(new()
        ///     {
        ///         InstanceId = "example_value",
        ///         AccountName = "root",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["mongodbAccountId1"] = example.Apply(getAccountsResult =&gt; getAccountsResult.Accounts[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAccountsResult> Invoke(GetAccountsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountsResult>("alicloud:mongodb/getAccounts:getAccounts", args ?? new GetAccountsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the account.
        /// </summary>
        [Input("accountName")]
        public string? AccountName { get; set; }

        /// <summary>
        /// The id of the instance to which the account belongs.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetAccountsArgs()
        {
        }
        public static new GetAccountsArgs Empty => new GetAccountsArgs();
    }

    public sealed class GetAccountsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the account.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// The id of the instance to which the account belongs.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetAccountsInvokeArgs()
        {
        }
        public static new GetAccountsInvokeArgs Empty => new GetAccountsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccountsResult
    {
        public readonly string? AccountName;
        public readonly ImmutableArray<Outputs.GetAccountsAccountResult> Accounts;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetAccountsResult(
            string? accountName,

            ImmutableArray<Outputs.GetAccountsAccountResult> accounts,

            string id,

            string instanceId,

            string? outputFile)
        {
            AccountName = accountName;
            Accounts = accounts;
            Id = id;
            InstanceId = instanceId;
            OutputFile = outputFile;
        }
    }
}
