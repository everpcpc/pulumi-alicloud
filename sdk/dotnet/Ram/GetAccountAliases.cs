// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ram
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides an alias for the Alibaba Cloud account.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ram_account_aliases.html.markdown.
        /// </summary>
        public static Task<GetAccountAliasesResult> GetAccountAliases(GetAccountAliasesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountAliasesResult>("alicloud:ram/getAccountAliases:getAccountAliases", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetAccountAliasesArgs : Pulumi.InvokeArgs
    {
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetAccountAliasesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetAccountAliasesResult
    {
        /// <summary>
        /// Alias of the account.
        /// </summary>
        public readonly string AccountAlias;
        public readonly string? OutputFile;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAccountAliasesResult(
            string accountAlias,
            string? outputFile,
            string id)
        {
            AccountAlias = accountAlias;
            OutputFile = outputFile;
            Id = id;
        }
    }
}
