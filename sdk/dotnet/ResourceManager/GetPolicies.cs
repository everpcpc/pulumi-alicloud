// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ResourceManager
{
    public static class GetPolicies
    {
        /// <summary>
        /// This data source provides the Resource Manager Policies of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:**  Available in 1.86.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(AliCloud.ResourceManager.GetPolicies.InvokeAsync(new AliCloud.ResourceManager.GetPoliciesArgs
        ///         {
        ///             DescriptionRegex = "tftest_policy",
        ///             NameRegex = "tftest",
        ///             PolicyType = "Custom",
        ///         }));
        ///         this.FirstPolicyId = example.Apply(example =&gt; example.Policies[0].Id);
        ///     }
        /// 
        ///     [Output("firstPolicyId")]
        ///     public Output&lt;string&gt; FirstPolicyId { get; set; }
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPoliciesResult> InvokeAsync(GetPoliciesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPoliciesResult>("alicloud:resourcemanager/getPolicies:getPolicies", args ?? new GetPoliciesArgs(), options.WithVersion());
    }


    public sealed class GetPoliciesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Resource Manager Policy IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by policy name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The type of the policy. If you do not specify this parameter, the system lists all types of policies. Valid values: `Custom` and `System`.
        /// </summary>
        [Input("policyType")]
        public string? PolicyType { get; set; }

        public GetPoliciesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPoliciesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of policy IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of policy names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of policies. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPoliciesPolicyResult> Policies;
        public readonly string? PolicyType;

        [OutputConstructor]
        private GetPoliciesResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetPoliciesPolicyResult> policies,

            string? policyType)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Policies = policies;
            PolicyType = policyType;
        }
    }
}
