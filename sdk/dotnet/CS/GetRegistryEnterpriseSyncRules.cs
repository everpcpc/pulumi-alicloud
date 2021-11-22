// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.CS
{
    public static class GetRegistryEnterpriseSyncRules
    {
        /// <summary>
        /// This data source provides a list Container Registry Enterprise Edition sync rules on Alibaba Cloud.
        /// 
        /// &gt; **NOTE:** Available in v1.90.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var mySyncRules = Output.Create(AliCloud.CS.GetRegistryEnterpriseSyncRules.InvokeAsync(new AliCloud.CS.GetRegistryEnterpriseSyncRulesArgs
        ///         {
        ///             InstanceId = "cri-xxx",
        ///             NamespaceName = "test-namespace",
        ///             RepoName = "test-repo",
        ///             TargetInstanceId = "cri-yyy",
        ///             NameRegex = "test-rule",
        ///         }));
        ///         this.Output = 
        ///         {
        ///             mySyncRules.Apply(mySyncRules =&gt; mySyncRules.Rules),
        ///         }.Select(__item =&gt; __item?.Id).ToList();
        ///     }
        /// 
        ///     [Output("output")]
        ///     public Output&lt;string&gt; Output { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRegistryEnterpriseSyncRulesResult> InvokeAsync(GetRegistryEnterpriseSyncRulesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegistryEnterpriseSyncRulesResult>("alicloud:cs/getRegistryEnterpriseSyncRules:getRegistryEnterpriseSyncRules", args ?? new GetRegistryEnterpriseSyncRulesArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides a list Container Registry Enterprise Edition sync rules on Alibaba Cloud.
        /// 
        /// &gt; **NOTE:** Available in v1.90.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var mySyncRules = Output.Create(AliCloud.CS.GetRegistryEnterpriseSyncRules.InvokeAsync(new AliCloud.CS.GetRegistryEnterpriseSyncRulesArgs
        ///         {
        ///             InstanceId = "cri-xxx",
        ///             NamespaceName = "test-namespace",
        ///             RepoName = "test-repo",
        ///             TargetInstanceId = "cri-yyy",
        ///             NameRegex = "test-rule",
        ///         }));
        ///         this.Output = 
        ///         {
        ///             mySyncRules.Apply(mySyncRules =&gt; mySyncRules.Rules),
        ///         }.Select(__item =&gt; __item?.Id).ToList();
        ///     }
        /// 
        ///     [Output("output")]
        ///     public Output&lt;string&gt; Output { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRegistryEnterpriseSyncRulesResult> Invoke(GetRegistryEnterpriseSyncRulesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegistryEnterpriseSyncRulesResult>("alicloud:cs/getRegistryEnterpriseSyncRules:getRegistryEnterpriseSyncRules", args ?? new GetRegistryEnterpriseSyncRulesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetRegistryEnterpriseSyncRulesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of ids to filter results by sync rule id.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// ID of Container Registry Enterprise Edition local instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by sync rule name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// Name of Container Registry Enterprise Edition local namespace.
        /// </summary>
        [Input("namespaceName")]
        public string? NamespaceName { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Name of Container Registry Enterprise Edition local repo.
        /// </summary>
        [Input("repoName")]
        public string? RepoName { get; set; }

        /// <summary>
        /// ID of Container Registry Enterprise Edition target instance.
        /// </summary>
        [Input("targetInstanceId")]
        public string? TargetInstanceId { get; set; }

        public GetRegistryEnterpriseSyncRulesArgs()
        {
        }
    }

    public sealed class GetRegistryEnterpriseSyncRulesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of ids to filter results by sync rule id.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// ID of Container Registry Enterprise Edition local instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by sync rule name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// Name of Container Registry Enterprise Edition local namespace.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Name of Container Registry Enterprise Edition local repo.
        /// </summary>
        [Input("repoName")]
        public Input<string>? RepoName { get; set; }

        /// <summary>
        /// ID of Container Registry Enterprise Edition target instance.
        /// </summary>
        [Input("targetInstanceId")]
        public Input<string>? TargetInstanceId { get; set; }

        public GetRegistryEnterpriseSyncRulesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegistryEnterpriseSyncRulesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of matched Container Registry Enterprise Edition sync rules. Its element is a sync rule uuid.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// ID of Container Registry Enterprise Edition local instance.
        /// </summary>
        public readonly string InstanceId;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of sync rule names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        /// <summary>
        /// Name of Container Registry Enterprise Edition local namespace.
        /// </summary>
        public readonly string? NamespaceName;
        public readonly string? OutputFile;
        /// <summary>
        /// Name of Container Registry Enterprise Edition local repo.
        /// </summary>
        public readonly string? RepoName;
        /// <summary>
        /// A list of matched Container Registry Enterprise Edition sync rules. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRegistryEnterpriseSyncRulesRuleResult> Rules;
        /// <summary>
        /// ID of Container Registry Enterprise Edition target instance.
        /// </summary>
        public readonly string? TargetInstanceId;

        [OutputConstructor]
        private GetRegistryEnterpriseSyncRulesResult(
            string id,

            ImmutableArray<string> ids,

            string instanceId,

            string? nameRegex,

            ImmutableArray<string> names,

            string? namespaceName,

            string? outputFile,

            string? repoName,

            ImmutableArray<Outputs.GetRegistryEnterpriseSyncRulesRuleResult> rules,

            string? targetInstanceId)
        {
            Id = id;
            Ids = ids;
            InstanceId = instanceId;
            NameRegex = nameRegex;
            Names = names;
            NamespaceName = namespaceName;
            OutputFile = outputFile;
            RepoName = repoName;
            Rules = rules;
            TargetInstanceId = targetInstanceId;
        }
    }
}
