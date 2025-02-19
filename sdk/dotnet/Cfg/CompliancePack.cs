// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cfg
{
    /// <summary>
    /// Provides a Cloud Config Compliance Pack resource.
    /// 
    /// For information about Cloud Config Compliance Pack and how to use it, see [What is Compliance Pack](https://www.alibabacloud.com/help/en/doc-detail/194753.html).
    /// 
    /// &gt; **NOTE:** Available in v1.124.0+.
    /// 
    /// ## Example Usage
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "example_name";
    ///     var defaultInstances = AliCloud.Ecs.GetInstances.Invoke();
    /// 
    ///     var defaultResourceGroups = AliCloud.ResourceManager.GetResourceGroups.Invoke(new()
    ///     {
    ///         Status = "OK",
    ///     });
    /// 
    ///     var defaultRule = new AliCloud.Cfg.Rule("defaultRule", new()
    ///     {
    ///         RuleName = name,
    ///         Description = name,
    ///         SourceIdentifier = "ecs-instances-in-vpc",
    ///         SourceOwner = "ALIYUN",
    ///         ResourceTypesScopes = new[]
    ///         {
    ///             "ACS::ECS::Instance",
    ///         },
    ///         RiskLevel = 1,
    ///         ConfigRuleTriggerTypes = "ConfigurationItemChangeNotification",
    ///         TagKeyScope = "tfTest",
    ///         TagValueScope = "tfTest 123",
    ///         ResourceGroupIdsScope = defaultResourceGroups.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Ids[0]),
    ///         ExcludeResourceIdsScope = defaultInstances.Apply(getInstancesResult =&gt; getInstancesResult.Instances[0]?.Id),
    ///         RegionIdsScope = "cn-hangzhou",
    ///         InputParameters = 
    ///         {
    ///             { "vpcIds", defaultInstances.Apply(getInstancesResult =&gt; getInstancesResult.Instances[0]?.VpcId) },
    ///         },
    ///     });
    /// 
    ///     var defaultCompliancePack = new AliCloud.Cfg.CompliancePack("defaultCompliancePack", new()
    ///     {
    ///         CompliancePackName = "tf-testaccConfig1234",
    ///         Description = "tf-testaccConfig1234",
    ///         RiskLevel = 1,
    ///         ConfigRuleIds = new[]
    ///         {
    ///             new AliCloud.Cfg.Inputs.CompliancePackConfigRuleIdArgs
    ///             {
    ///                 ConfigRuleId = defaultRule.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cloud Config Compliance Pack can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cfg/compliancePack:CompliancePack example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cfg/compliancePack:CompliancePack")]
    public partial class CompliancePack : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Compliance Package Name. . **NOTE:** the `compliance_pack_name` supports modification since V1.146.0.
        /// </summary>
        [Output("compliancePackName")]
        public Output<string> CompliancePackName { get; private set; } = null!;

        /// <summary>
        /// Compliance Package Template Id.
        /// </summary>
        [Output("compliancePackTemplateId")]
        public Output<string?> CompliancePackTemplateId { get; private set; } = null!;

        /// <summary>
        /// A list of Config Rule IDs.
        /// </summary>
        [Output("configRuleIds")]
        public Output<ImmutableArray<Outputs.CompliancePackConfigRuleId>> ConfigRuleIds { get; private set; } = null!;

        /// <summary>
        /// A list of Config Rules.
        /// </summary>
        [Output("configRules")]
        public Output<ImmutableArray<Outputs.CompliancePackConfigRule>> ConfigRules { get; private set; } = null!;

        /// <summary>
        /// The Description of compliance pack.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The Risk Level. Valid values:  `1`: critical, `2`: warning, `3`: info.
        /// </summary>
        [Output("riskLevel")]
        public Output<int> RiskLevel { get; private set; } = null!;

        /// <summary>
        /// The status of the resource. The valid values: `CREATING`, `ACTIVE`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a CompliancePack resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CompliancePack(string name, CompliancePackArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cfg/compliancePack:CompliancePack", name, args ?? new CompliancePackArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CompliancePack(string name, Input<string> id, CompliancePackState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cfg/compliancePack:CompliancePack", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CompliancePack resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CompliancePack Get(string name, Input<string> id, CompliancePackState? state = null, CustomResourceOptions? options = null)
        {
            return new CompliancePack(name, id, state, options);
        }
    }

    public sealed class CompliancePackArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Compliance Package Name. . **NOTE:** the `compliance_pack_name` supports modification since V1.146.0.
        /// </summary>
        [Input("compliancePackName", required: true)]
        public Input<string> CompliancePackName { get; set; } = null!;

        /// <summary>
        /// Compliance Package Template Id.
        /// </summary>
        [Input("compliancePackTemplateId")]
        public Input<string>? CompliancePackTemplateId { get; set; }

        [Input("configRuleIds")]
        private InputList<Inputs.CompliancePackConfigRuleIdArgs>? _configRuleIds;

        /// <summary>
        /// A list of Config Rule IDs.
        /// </summary>
        public InputList<Inputs.CompliancePackConfigRuleIdArgs> ConfigRuleIds
        {
            get => _configRuleIds ?? (_configRuleIds = new InputList<Inputs.CompliancePackConfigRuleIdArgs>());
            set => _configRuleIds = value;
        }

        [Input("configRules")]
        private InputList<Inputs.CompliancePackConfigRuleArgs>? _configRules;

        /// <summary>
        /// A list of Config Rules.
        /// </summary>
        [Obsolete(@"Field 'config_rules' has been deprecated from provider version 1.141.0. New field 'config_rule_ids' instead.")]
        public InputList<Inputs.CompliancePackConfigRuleArgs> ConfigRules
        {
            get => _configRules ?? (_configRules = new InputList<Inputs.CompliancePackConfigRuleArgs>());
            set => _configRules = value;
        }

        /// <summary>
        /// The Description of compliance pack.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The Risk Level. Valid values:  `1`: critical, `2`: warning, `3`: info.
        /// </summary>
        [Input("riskLevel", required: true)]
        public Input<int> RiskLevel { get; set; } = null!;

        public CompliancePackArgs()
        {
        }
        public static new CompliancePackArgs Empty => new CompliancePackArgs();
    }

    public sealed class CompliancePackState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Compliance Package Name. . **NOTE:** the `compliance_pack_name` supports modification since V1.146.0.
        /// </summary>
        [Input("compliancePackName")]
        public Input<string>? CompliancePackName { get; set; }

        /// <summary>
        /// Compliance Package Template Id.
        /// </summary>
        [Input("compliancePackTemplateId")]
        public Input<string>? CompliancePackTemplateId { get; set; }

        [Input("configRuleIds")]
        private InputList<Inputs.CompliancePackConfigRuleIdGetArgs>? _configRuleIds;

        /// <summary>
        /// A list of Config Rule IDs.
        /// </summary>
        public InputList<Inputs.CompliancePackConfigRuleIdGetArgs> ConfigRuleIds
        {
            get => _configRuleIds ?? (_configRuleIds = new InputList<Inputs.CompliancePackConfigRuleIdGetArgs>());
            set => _configRuleIds = value;
        }

        [Input("configRules")]
        private InputList<Inputs.CompliancePackConfigRuleGetArgs>? _configRules;

        /// <summary>
        /// A list of Config Rules.
        /// </summary>
        [Obsolete(@"Field 'config_rules' has been deprecated from provider version 1.141.0. New field 'config_rule_ids' instead.")]
        public InputList<Inputs.CompliancePackConfigRuleGetArgs> ConfigRules
        {
            get => _configRules ?? (_configRules = new InputList<Inputs.CompliancePackConfigRuleGetArgs>());
            set => _configRules = value;
        }

        /// <summary>
        /// The Description of compliance pack.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Risk Level. Valid values:  `1`: critical, `2`: warning, `3`: info.
        /// </summary>
        [Input("riskLevel")]
        public Input<int>? RiskLevel { get; set; }

        /// <summary>
        /// The status of the resource. The valid values: `CREATING`, `ACTIVE`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public CompliancePackState()
        {
        }
        public static new CompliancePackState Empty => new CompliancePackState();
    }
}
