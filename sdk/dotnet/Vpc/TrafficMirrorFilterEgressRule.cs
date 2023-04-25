// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    /// <summary>
    /// Provides a VPC Traffic Mirror Filter Egress Rule resource.
    /// 
    /// For information about VPC Traffic Mirror Filter Egress Rule and how to use it, see [What is Traffic Mirror Filter Egress Rule](https://www.alibabacloud.com/help/doc-detail/261357.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.140.0+.
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
    ///     var exampleTrafficMirrorFilter = new AliCloud.Vpc.TrafficMirrorFilter("exampleTrafficMirrorFilter", new()
    ///     {
    ///         TrafficMirrorFilterName = "example_value",
    ///     });
    /// 
    ///     var exampleTrafficMirrorFilterEgressRule = new AliCloud.Vpc.TrafficMirrorFilterEgressRule("exampleTrafficMirrorFilterEgressRule", new()
    ///     {
    ///         TrafficMirrorFilterId = exampleTrafficMirrorFilter.Id,
    ///         Priority = 1,
    ///         RuleAction = "accept",
    ///         Protocol = "UDP",
    ///         DestinationCidrBlock = "10.0.0.0/24",
    ///         SourceCidrBlock = "10.0.0.0/24",
    ///         DestinationPortRange = "1/120",
    ///         SourcePortRange = "1/120",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VPC Traffic Mirror Filter Egress Rule can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:vpc/trafficMirrorFilterEgressRule:TrafficMirrorFilterEgressRule example &lt;traffic_mirror_filter_id&gt;:&lt;traffic_mirror_filter_egress_rule_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/trafficMirrorFilterEgressRule:TrafficMirrorFilterEgressRule")]
    public partial class TrafficMirrorFilterEgressRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The destination CIDR block of the outbound traffic.
        /// </summary>
        [Output("destinationCidrBlock")]
        public Output<string> DestinationCidrBlock { get; private set; } = null!;

        /// <summary>
        /// The destination CIDR block of the outbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        /// </summary>
        [Output("destinationPortRange")]
        public Output<string> DestinationPortRange { get; private set; } = null!;

        /// <summary>
        /// Whether to pre-check this request only. Default to: `false`
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// The transport protocol used by outbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
        /// </summary>
        [Output("ruleAction")]
        public Output<string> RuleAction { get; private set; } = null!;

        /// <summary>
        /// The source CIDR block of the outbound traffic.
        /// </summary>
        [Output("sourceCidrBlock")]
        public Output<string> SourceCidrBlock { get; private set; } = null!;

        /// <summary>
        /// The source port range of the outbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        /// </summary>
        [Output("sourcePortRange")]
        public Output<string> SourcePortRange { get; private set; } = null!;

        /// <summary>
        /// The state of the inbound rule. Valid values:`Creating`, `Created`, `Modifying` and `Deleting`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the outbound rule.
        /// </summary>
        [Output("trafficMirrorFilterEgressRuleId")]
        public Output<string> TrafficMirrorFilterEgressRuleId { get; private set; } = null!;

        /// <summary>
        /// The ID of the filter.
        /// </summary>
        [Output("trafficMirrorFilterId")]
        public Output<string> TrafficMirrorFilterId { get; private set; } = null!;


        /// <summary>
        /// Create a TrafficMirrorFilterEgressRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TrafficMirrorFilterEgressRule(string name, TrafficMirrorFilterEgressRuleArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/trafficMirrorFilterEgressRule:TrafficMirrorFilterEgressRule", name, args ?? new TrafficMirrorFilterEgressRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TrafficMirrorFilterEgressRule(string name, Input<string> id, TrafficMirrorFilterEgressRuleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/trafficMirrorFilterEgressRule:TrafficMirrorFilterEgressRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TrafficMirrorFilterEgressRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TrafficMirrorFilterEgressRule Get(string name, Input<string> id, TrafficMirrorFilterEgressRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new TrafficMirrorFilterEgressRule(name, id, state, options);
        }
    }

    public sealed class TrafficMirrorFilterEgressRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination CIDR block of the outbound traffic.
        /// </summary>
        [Input("destinationCidrBlock", required: true)]
        public Input<string> DestinationCidrBlock { get; set; } = null!;

        /// <summary>
        /// The destination CIDR block of the outbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        /// </summary>
        [Input("destinationPortRange")]
        public Input<string>? DestinationPortRange { get; set; }

        /// <summary>
        /// Whether to pre-check this request only. Default to: `false`
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// The transport protocol used by outbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
        /// </summary>
        [Input("ruleAction", required: true)]
        public Input<string> RuleAction { get; set; } = null!;

        /// <summary>
        /// The source CIDR block of the outbound traffic.
        /// </summary>
        [Input("sourceCidrBlock", required: true)]
        public Input<string> SourceCidrBlock { get; set; } = null!;

        /// <summary>
        /// The source port range of the outbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        /// </summary>
        [Input("sourcePortRange")]
        public Input<string>? SourcePortRange { get; set; }

        /// <summary>
        /// The ID of the filter.
        /// </summary>
        [Input("trafficMirrorFilterId", required: true)]
        public Input<string> TrafficMirrorFilterId { get; set; } = null!;

        public TrafficMirrorFilterEgressRuleArgs()
        {
        }
        public static new TrafficMirrorFilterEgressRuleArgs Empty => new TrafficMirrorFilterEgressRuleArgs();
    }

    public sealed class TrafficMirrorFilterEgressRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination CIDR block of the outbound traffic.
        /// </summary>
        [Input("destinationCidrBlock")]
        public Input<string>? DestinationCidrBlock { get; set; }

        /// <summary>
        /// The destination CIDR block of the outbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        /// </summary>
        [Input("destinationPortRange")]
        public Input<string>? DestinationPortRange { get; set; }

        /// <summary>
        /// Whether to pre-check this request only. Default to: `false`
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The transport protocol used by outbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
        /// </summary>
        [Input("ruleAction")]
        public Input<string>? RuleAction { get; set; }

        /// <summary>
        /// The source CIDR block of the outbound traffic.
        /// </summary>
        [Input("sourceCidrBlock")]
        public Input<string>? SourceCidrBlock { get; set; }

        /// <summary>
        /// The source port range of the outbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        /// </summary>
        [Input("sourcePortRange")]
        public Input<string>? SourcePortRange { get; set; }

        /// <summary>
        /// The state of the inbound rule. Valid values:`Creating`, `Created`, `Modifying` and `Deleting`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the outbound rule.
        /// </summary>
        [Input("trafficMirrorFilterEgressRuleId")]
        public Input<string>? TrafficMirrorFilterEgressRuleId { get; set; }

        /// <summary>
        /// The ID of the filter.
        /// </summary>
        [Input("trafficMirrorFilterId")]
        public Input<string>? TrafficMirrorFilterId { get; set; }

        public TrafficMirrorFilterEgressRuleState()
        {
        }
        public static new TrafficMirrorFilterEgressRuleState Empty => new TrafficMirrorFilterEgressRuleState();
    }
}
