// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms
{
    /// <summary>
    /// Provides a Cloud Monitor Service Group Metric Rule resource.
    /// 
    /// For information about Cloud Monitor Service Group Metric Rule and how to use it, see [What is Group Metric Rule](https://www.alibabacloud.com/help/en/doc-detail/114943.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.104.0+.
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
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var thisRandomUuid = new Random.RandomUuid("thisRandomUuid");
    /// 
    ///     var thisGroupMetricRule = new AliCloud.Cms.GroupMetricRule("thisGroupMetricRule", new()
    ///     {
    ///         GroupId = "539****",
    ///         RuleId = thisRandomUuid.Id,
    ///         Category = "ecs",
    ///         Namespace = "acs_ecs_dashboard",
    ///         MetricName = "cpu_total",
    ///         Period = 60,
    ///         GroupMetricRuleName = "tf-testacc-rule-name",
    ///         EmailSubject = "tf-testacc-rule-name-warning",
    ///         Interval = "3600",
    ///         SilenceTime = 85800,
    ///         NoEffectiveInterval = "00:00-05:30",
    ///         Webhook = "http://www.aliyun.com",
    ///         Escalations = new AliCloud.Cms.Inputs.GroupMetricRuleEscalationsArgs
    ///         {
    ///             Warn = new AliCloud.Cms.Inputs.GroupMetricRuleEscalationsWarnArgs
    ///             {
    ///                 ComparisonOperator = "GreaterThanOrEqualToThreshold",
    ///                 Statistics = "Average",
    ///                 Threshold = "90",
    ///                 Times = 3,
    ///             },
    ///             Info = new AliCloud.Cms.Inputs.GroupMetricRuleEscalationsInfoArgs
    ///             {
    ///                 ComparisonOperator = "LessThanLastWeek",
    ///                 Statistics = "Average",
    ///                 Threshold = "90",
    ///                 Times = 5,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cloud Monitor Service Group Metric Rule can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cms/groupMetricRule:GroupMetricRule example &lt;rule_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cms/groupMetricRule:GroupMetricRule")]
    public partial class GroupMetricRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The abbreviation of the service name.
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// Alarm contact group.
        /// </summary>
        [Output("contactGroups")]
        public Output<string> ContactGroups { get; private set; } = null!;

        /// <summary>
        /// The dimensions that specify the resources to be associated with the alert rule.
        /// </summary>
        [Output("dimensions")]
        public Output<string> Dimensions { get; private set; } = null!;

        /// <summary>
        /// The time period during which the alert rule is effective.
        /// </summary>
        [Output("effectiveInterval")]
        public Output<string?> EffectiveInterval { get; private set; } = null!;

        /// <summary>
        /// The subject of the alert notification email.                                         .
        /// </summary>
        [Output("emailSubject")]
        public Output<string> EmailSubject { get; private set; } = null!;

        /// <summary>
        /// Alarm level. See the following `Block escalations`.
        /// </summary>
        [Output("escalations")]
        public Output<Outputs.GroupMetricRuleEscalations> Escalations { get; private set; } = null!;

        /// <summary>
        /// The ID of the application group.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the alert rule.
        /// </summary>
        [Output("groupMetricRuleName")]
        public Output<string> GroupMetricRuleName { get; private set; } = null!;

        /// <summary>
        /// The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
        /// </summary>
        [Output("interval")]
        public Output<string?> Interval { get; private set; } = null!;

        /// <summary>
        /// The name of the metric.
        /// </summary>
        [Output("metricName")]
        public Output<string> MetricName { get; private set; } = null!;

        /// <summary>
        /// The namespace of the service.
        /// </summary>
        [Output("namespace")]
        public Output<string> Namespace { get; private set; } = null!;

        /// <summary>
        /// The time period during which the alert rule is ineffective.
        /// </summary>
        [Output("noEffectiveInterval")]
        public Output<string?> NoEffectiveInterval { get; private set; } = null!;

        /// <summary>
        /// The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
        /// </summary>
        [Output("period")]
        public Output<int> Period { get; private set; } = null!;

        /// <summary>
        /// The ID of the alert rule.
        /// </summary>
        [Output("ruleId")]
        public Output<string> RuleId { get; private set; } = null!;

        /// <summary>
        /// The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
        /// </summary>
        [Output("silenceTime")]
        public Output<int?> SilenceTime { get; private set; } = null!;

        /// <summary>
        /// The status of Group Metric Rule.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The information about the resource for which alerts are triggered. See the following `Block targets`.
        /// </summary>
        [Output("targets")]
        public Output<ImmutableArray<Outputs.GroupMetricRuleTarget>> Targets { get; private set; } = null!;

        /// <summary>
        /// The callback URL.
        /// </summary>
        [Output("webhook")]
        public Output<string?> Webhook { get; private set; } = null!;


        /// <summary>
        /// Create a GroupMetricRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupMetricRule(string name, GroupMetricRuleArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cms/groupMetricRule:GroupMetricRule", name, args ?? new GroupMetricRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupMetricRule(string name, Input<string> id, GroupMetricRuleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cms/groupMetricRule:GroupMetricRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupMetricRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupMetricRule Get(string name, Input<string> id, GroupMetricRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupMetricRule(name, id, state, options);
        }
    }

    public sealed class GroupMetricRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The abbreviation of the service name.
        /// </summary>
        [Input("category", required: true)]
        public Input<string> Category { get; set; } = null!;

        /// <summary>
        /// Alarm contact group.
        /// </summary>
        [Input("contactGroups")]
        public Input<string>? ContactGroups { get; set; }

        /// <summary>
        /// The dimensions that specify the resources to be associated with the alert rule.
        /// </summary>
        [Input("dimensions")]
        public Input<string>? Dimensions { get; set; }

        /// <summary>
        /// The time period during which the alert rule is effective.
        /// </summary>
        [Input("effectiveInterval")]
        public Input<string>? EffectiveInterval { get; set; }

        /// <summary>
        /// The subject of the alert notification email.                                         .
        /// </summary>
        [Input("emailSubject")]
        public Input<string>? EmailSubject { get; set; }

        /// <summary>
        /// Alarm level. See the following `Block escalations`.
        /// </summary>
        [Input("escalations", required: true)]
        public Input<Inputs.GroupMetricRuleEscalationsArgs> Escalations { get; set; } = null!;

        /// <summary>
        /// The ID of the application group.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The name of the alert rule.
        /// </summary>
        [Input("groupMetricRuleName", required: true)]
        public Input<string> GroupMetricRuleName { get; set; } = null!;

        /// <summary>
        /// The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
        /// </summary>
        [Input("interval")]
        public Input<string>? Interval { get; set; }

        /// <summary>
        /// The name of the metric.
        /// </summary>
        [Input("metricName", required: true)]
        public Input<string> MetricName { get; set; } = null!;

        /// <summary>
        /// The namespace of the service.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        /// <summary>
        /// The time period during which the alert rule is ineffective.
        /// </summary>
        [Input("noEffectiveInterval")]
        public Input<string>? NoEffectiveInterval { get; set; }

        /// <summary>
        /// The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The ID of the alert rule.
        /// </summary>
        [Input("ruleId", required: true)]
        public Input<string> RuleId { get; set; } = null!;

        /// <summary>
        /// The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
        /// </summary>
        [Input("silenceTime")]
        public Input<int>? SilenceTime { get; set; }

        [Input("targets")]
        private InputList<Inputs.GroupMetricRuleTargetArgs>? _targets;

        /// <summary>
        /// The information about the resource for which alerts are triggered. See the following `Block targets`.
        /// </summary>
        public InputList<Inputs.GroupMetricRuleTargetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.GroupMetricRuleTargetArgs>());
            set => _targets = value;
        }

        /// <summary>
        /// The callback URL.
        /// </summary>
        [Input("webhook")]
        public Input<string>? Webhook { get; set; }

        public GroupMetricRuleArgs()
        {
        }
        public static new GroupMetricRuleArgs Empty => new GroupMetricRuleArgs();
    }

    public sealed class GroupMetricRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The abbreviation of the service name.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Alarm contact group.
        /// </summary>
        [Input("contactGroups")]
        public Input<string>? ContactGroups { get; set; }

        /// <summary>
        /// The dimensions that specify the resources to be associated with the alert rule.
        /// </summary>
        [Input("dimensions")]
        public Input<string>? Dimensions { get; set; }

        /// <summary>
        /// The time period during which the alert rule is effective.
        /// </summary>
        [Input("effectiveInterval")]
        public Input<string>? EffectiveInterval { get; set; }

        /// <summary>
        /// The subject of the alert notification email.                                         .
        /// </summary>
        [Input("emailSubject")]
        public Input<string>? EmailSubject { get; set; }

        /// <summary>
        /// Alarm level. See the following `Block escalations`.
        /// </summary>
        [Input("escalations")]
        public Input<Inputs.GroupMetricRuleEscalationsGetArgs>? Escalations { get; set; }

        /// <summary>
        /// The ID of the application group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the alert rule.
        /// </summary>
        [Input("groupMetricRuleName")]
        public Input<string>? GroupMetricRuleName { get; set; }

        /// <summary>
        /// The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
        /// </summary>
        [Input("interval")]
        public Input<string>? Interval { get; set; }

        /// <summary>
        /// The name of the metric.
        /// </summary>
        [Input("metricName")]
        public Input<string>? MetricName { get; set; }

        /// <summary>
        /// The namespace of the service.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The time period during which the alert rule is ineffective.
        /// </summary>
        [Input("noEffectiveInterval")]
        public Input<string>? NoEffectiveInterval { get; set; }

        /// <summary>
        /// The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The ID of the alert rule.
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        /// <summary>
        /// The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
        /// </summary>
        [Input("silenceTime")]
        public Input<int>? SilenceTime { get; set; }

        /// <summary>
        /// The status of Group Metric Rule.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("targets")]
        private InputList<Inputs.GroupMetricRuleTargetGetArgs>? _targets;

        /// <summary>
        /// The information about the resource for which alerts are triggered. See the following `Block targets`.
        /// </summary>
        public InputList<Inputs.GroupMetricRuleTargetGetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.GroupMetricRuleTargetGetArgs>());
            set => _targets = value;
        }

        /// <summary>
        /// The callback URL.
        /// </summary>
        [Input("webhook")]
        public Input<string>? Webhook { get; set; }

        public GroupMetricRuleState()
        {
        }
        public static new GroupMetricRuleState Empty => new GroupMetricRuleState();
    }
}
