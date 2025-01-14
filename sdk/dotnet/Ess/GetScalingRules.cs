// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ess
{
    public static class GetScalingRules
    {
        /// <summary>
        /// This data source provides available scaling rule resources. 
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
        ///     var scalingrulesDs = AliCloud.Ess.GetScalingRules.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "scaling_rule_id1",
        ///             "scaling_rule_id2",
        ///         },
        ///         NameRegex = "scaling_rule_name",
        ///         ScalingGroupId = "scaling_group_id",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstScalingRule"] = scalingrulesDs.Apply(getScalingRulesResult =&gt; getScalingRulesResult.Rules[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetScalingRulesResult> InvokeAsync(GetScalingRulesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetScalingRulesResult>("alicloud:ess/getScalingRules:getScalingRules", args ?? new GetScalingRulesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides available scaling rule resources. 
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
        ///     var scalingrulesDs = AliCloud.Ess.GetScalingRules.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "scaling_rule_id1",
        ///             "scaling_rule_id2",
        ///         },
        ///         NameRegex = "scaling_rule_name",
        ///         ScalingGroupId = "scaling_group_id",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstScalingRule"] = scalingrulesDs.Apply(getScalingRulesResult =&gt; getScalingRulesResult.Rules[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetScalingRulesResult> Invoke(GetScalingRulesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetScalingRulesResult>("alicloud:ess/getScalingRules:getScalingRules", args ?? new GetScalingRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetScalingRulesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of scaling rule IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter resulting scaling rules by name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Scaling group id the scaling rules belong to.
        /// </summary>
        [Input("scalingGroupId")]
        public string? ScalingGroupId { get; set; }

        /// <summary>
        /// Type of scaling rule.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetScalingRulesArgs()
        {
        }
        public static new GetScalingRulesArgs Empty => new GetScalingRulesArgs();
    }

    public sealed class GetScalingRulesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of scaling rule IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter resulting scaling rules by name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Scaling group id the scaling rules belong to.
        /// </summary>
        [Input("scalingGroupId")]
        public Input<string>? ScalingGroupId { get; set; }

        /// <summary>
        /// Type of scaling rule.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetScalingRulesInvokeArgs()
        {
        }
        public static new GetScalingRulesInvokeArgs Empty => new GetScalingRulesInvokeArgs();
    }


    [OutputType]
    public sealed class GetScalingRulesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of scaling rule ids.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of scaling rule names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of scaling rules. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetScalingRulesRuleResult> Rules;
        /// <summary>
        /// ID of the scaling group.
        /// </summary>
        public readonly string? ScalingGroupId;
        /// <summary>
        /// Type of the scaling rule.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetScalingRulesResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetScalingRulesRuleResult> rules,

            string? scalingGroupId,

            string? type)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Rules = rules;
            ScalingGroupId = scalingGroupId;
            Type = type;
        }
    }
}
