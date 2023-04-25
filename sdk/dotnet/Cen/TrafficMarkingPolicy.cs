// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    /// <summary>
    /// Provides a Cloud Enterprise Network (CEN) Traffic Marking Policy resource.
    /// 
    /// For information about Cloud Enterprise Network (CEN) Traffic Marking Policy and how to use it, see [What is Traffic Marking Policy](https://help.aliyun.com/document_detail/419025.html).
    /// 
    /// &gt; **NOTE:** Available in v1.173.0+.
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
    ///     var exampleInstance = new AliCloud.Cen.Instance("exampleInstance", new()
    ///     {
    ///         CenInstanceName = "example_value",
    ///     });
    /// 
    ///     var exampleTransitRouter = new AliCloud.Cen.TransitRouter("exampleTransitRouter", new()
    ///     {
    ///         CenId = exampleInstance.Id,
    ///         TransitRouterName = "example_value",
    ///     });
    /// 
    ///     var exampleTrafficMarkingPolicy = new AliCloud.Cen.TrafficMarkingPolicy("exampleTrafficMarkingPolicy", new()
    ///     {
    ///         MarkingDscp = 1,
    ///         Priority = 1,
    ///         TrafficMarkingPolicyName = "example_value",
    ///         TransitRouterId = exampleTransitRouter.TransitRouterId,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cloud Enterprise Network (CEN) Traffic Marking Policy can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cen/trafficMarkingPolicy:TrafficMarkingPolicy example &lt;transit_router_id&gt;:&lt;traffic_marking_policy_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cen/trafficMarkingPolicy:TrafficMarkingPolicy")]
    public partial class TrafficMarkingPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the Traffic Marking Policy. The description must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The dry run.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The DSCP(Differentiated Services Code Point) of the Traffic Marking Policy. Value range: 0~63.
        /// </summary>
        [Output("markingDscp")]
        public Output<int> MarkingDscp { get; private set; } = null!;

        /// <summary>
        /// The Priority of the Traffic Marking Policy. Value range: 1~100.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the Traffic Marking Policy.
        /// </summary>
        [Output("trafficMarkingPolicyId")]
        public Output<string> TrafficMarkingPolicyId { get; private set; } = null!;

        /// <summary>
        /// The name of the Traffic Marking Policy. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
        /// </summary>
        [Output("trafficMarkingPolicyName")]
        public Output<string?> TrafficMarkingPolicyName { get; private set; } = null!;

        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        [Output("transitRouterId")]
        public Output<string> TransitRouterId { get; private set; } = null!;


        /// <summary>
        /// Create a TrafficMarkingPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TrafficMarkingPolicy(string name, TrafficMarkingPolicyArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/trafficMarkingPolicy:TrafficMarkingPolicy", name, args ?? new TrafficMarkingPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TrafficMarkingPolicy(string name, Input<string> id, TrafficMarkingPolicyState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/trafficMarkingPolicy:TrafficMarkingPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TrafficMarkingPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TrafficMarkingPolicy Get(string name, Input<string> id, TrafficMarkingPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new TrafficMarkingPolicy(name, id, state, options);
        }
    }

    public sealed class TrafficMarkingPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Traffic Marking Policy. The description must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The dry run.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The DSCP(Differentiated Services Code Point) of the Traffic Marking Policy. Value range: 0~63.
        /// </summary>
        [Input("markingDscp", required: true)]
        public Input<int> MarkingDscp { get; set; } = null!;

        /// <summary>
        /// The Priority of the Traffic Marking Policy. Value range: 1~100.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// The name of the Traffic Marking Policy. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
        /// </summary>
        [Input("trafficMarkingPolicyName")]
        public Input<string>? TrafficMarkingPolicyName { get; set; }

        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        [Input("transitRouterId", required: true)]
        public Input<string> TransitRouterId { get; set; } = null!;

        public TrafficMarkingPolicyArgs()
        {
        }
        public static new TrafficMarkingPolicyArgs Empty => new TrafficMarkingPolicyArgs();
    }

    public sealed class TrafficMarkingPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Traffic Marking Policy. The description must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The dry run.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The DSCP(Differentiated Services Code Point) of the Traffic Marking Policy. Value range: 0~63.
        /// </summary>
        [Input("markingDscp")]
        public Input<int>? MarkingDscp { get; set; }

        /// <summary>
        /// The Priority of the Traffic Marking Policy. Value range: 1~100.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the Traffic Marking Policy.
        /// </summary>
        [Input("trafficMarkingPolicyId")]
        public Input<string>? TrafficMarkingPolicyId { get; set; }

        /// <summary>
        /// The name of the Traffic Marking Policy. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
        /// </summary>
        [Input("trafficMarkingPolicyName")]
        public Input<string>? TrafficMarkingPolicyName { get; set; }

        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        [Input("transitRouterId")]
        public Input<string>? TransitRouterId { get; set; }

        public TrafficMarkingPolicyState()
        {
        }
        public static new TrafficMarkingPolicyState Empty => new TrafficMarkingPolicyState();
    }
}
