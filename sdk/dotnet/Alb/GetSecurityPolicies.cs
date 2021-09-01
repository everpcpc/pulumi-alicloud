// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb
{
    public static class GetSecurityPolicies
    {
        /// <summary>
        /// This data source provides the Alb Security Policies of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.130.0+.
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
        ///         var ids = Output.Create(AliCloud.Alb.GetSecurityPolicies.InvokeAsync());
        ///         this.AlbSecurityPolicyId1 = ids.Apply(ids =&gt; ids.Policies[0].Id);
        ///         var nameRegex = Output.Create(AliCloud.Alb.GetSecurityPolicies.InvokeAsync(new AliCloud.Alb.GetSecurityPoliciesArgs
        ///         {
        ///             NameRegex = "^my-SecurityPolicy",
        ///         }));
        ///         this.AlbSecurityPolicyId2 = nameRegex.Apply(nameRegex =&gt; nameRegex.Policies[0].Id);
        ///     }
        /// 
        ///     [Output("albSecurityPolicyId1")]
        ///     public Output&lt;string&gt; AlbSecurityPolicyId1 { get; set; }
        ///     [Output("albSecurityPolicyId2")]
        ///     public Output&lt;string&gt; AlbSecurityPolicyId2 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSecurityPoliciesResult> InvokeAsync(GetSecurityPoliciesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecurityPoliciesResult>("alicloud:alb/getSecurityPolicies:getSecurityPolicies", args ?? new GetSecurityPoliciesArgs(), options.WithVersion());
    }


    public sealed class GetSecurityPoliciesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Security Policy IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Security Policy name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        [Input("securityPolicyIds")]
        private List<string>? _securityPolicyIds;

        /// <summary>
        /// The security policy ids.
        /// </summary>
        public List<string> SecurityPolicyIds
        {
            get => _securityPolicyIds ?? (_securityPolicyIds = new List<string>());
            set => _securityPolicyIds = value;
        }

        /// <summary>
        /// The name of the resource. The name must be 2 to 128 characters in length and must start with a letter. It can contain digits, periods (.), underscores (_), and hyphens (-).
        /// </summary>
        [Input("securityPolicyName")]
        public string? SecurityPolicyName { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetSecurityPoliciesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSecurityPoliciesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetSecurityPoliciesPolicyResult> Policies;
        public readonly string? ResourceGroupId;
        public readonly ImmutableArray<string> SecurityPolicyIds;
        public readonly string? SecurityPolicyName;
        public readonly string? Status;
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetSecurityPoliciesResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetSecurityPoliciesPolicyResult> policies,

            string? resourceGroupId,

            ImmutableArray<string> securityPolicyIds,

            string? securityPolicyName,

            string? status,

            ImmutableDictionary<string, object>? tags)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Policies = policies;
            ResourceGroupId = resourceGroupId;
            SecurityPolicyIds = securityPolicyIds;
            SecurityPolicyName = securityPolicyName;
            Status = status;
            Tags = tags;
        }
    }
}