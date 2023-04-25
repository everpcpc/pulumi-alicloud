// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Log
{
    /// <summary>
    /// The dashboard is a real-time data analysis platform provided by the log service. You can display frequently used query and analysis statements in the form of charts and save statistical charts to the dashboard.
    /// [Refer to details](https://www.alibabacloud.com/help/doc-detail/102530.htm).
    /// 
    /// &gt; **NOTE:** Available in 1.86.0, parameter "action" in char_list is supported since 1.164.0+.
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
    ///     var defaultProject = new AliCloud.Log.Project("defaultProject", new()
    ///     {
    ///         Description = "tf unit test",
    ///     });
    /// 
    ///     var defaultStore = new AliCloud.Log.Store("defaultStore", new()
    ///     {
    ///         Project = "tf-project",
    ///         RetentionPeriod = 3000,
    ///         ShardCount = 1,
    ///     });
    /// 
    ///     var example = new AliCloud.Log.Dashboard("example", new()
    ///     {
    ///         Attribute = "{\"type\":\"grid\"}",
    ///         CharList = @"  [
    ///     {
    ///       ""action"": {},
    ///       ""title"":""new_title"",
    ///       ""type"":""map"",
    ///       ""search"":{
    ///         ""logstore"":""tf-logstore"",
    ///         ""topic"":""new_topic"",
    ///         ""query"":""* | SELECT COUNT(name) as ct_name, COUNT(product) as ct_product, name,product GROUP BY name,product"",
    ///         ""start"":""-86400s"",
    ///         ""end"":""now""
    ///       },
    ///       ""display"":{
    ///         ""xAxis"":[
    ///           ""ct_name""
    ///         ],
    ///         ""yAxis"":[
    ///           ""ct_product""
    ///         ],
    ///         ""xPos"":0,
    ///         ""yPos"":0,
    ///         ""width"":10,
    ///         ""height"":12,
    ///         ""displayName"":""xixihaha911""
    ///       }
    ///     }
    ///   ]
    /// 
    /// ",
    ///         DashboardName = "tf-dashboard",
    ///         ProjectName = "tf-project",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Log Dashboard can be imported using the id or name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:log/dashboard:Dashboard example tf-project:tf-logstore:tf-dashboard
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:log/dashboard:Dashboard")]
    public partial class Dashboard : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Dashboard attribute.
        /// </summary>
        [Output("attribute")]
        public Output<string> Attribute { get; private set; } = null!;

        /// <summary>
        /// Configuration of charts in the dashboard.
        /// </summary>
        [Output("charList")]
        public Output<string> CharList { get; private set; } = null!;

        /// <summary>
        /// The name of the Log Dashboard.
        /// </summary>
        [Output("dashboardName")]
        public Output<string> DashboardName { get; private set; } = null!;

        /// <summary>
        /// Dashboard alias.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The name of the log project. It is the only in one Alicloud account.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;


        /// <summary>
        /// Create a Dashboard resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dashboard(string name, DashboardArgs args, CustomResourceOptions? options = null)
            : base("alicloud:log/dashboard:Dashboard", name, args ?? new DashboardArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dashboard(string name, Input<string> id, DashboardState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:log/dashboard:Dashboard", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Dashboard resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dashboard Get(string name, Input<string> id, DashboardState? state = null, CustomResourceOptions? options = null)
        {
            return new Dashboard(name, id, state, options);
        }
    }

    public sealed class DashboardArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dashboard attribute.
        /// </summary>
        [Input("attribute")]
        public Input<string>? Attribute { get; set; }

        /// <summary>
        /// Configuration of charts in the dashboard.
        /// </summary>
        [Input("charList", required: true)]
        public Input<string> CharList { get; set; } = null!;

        /// <summary>
        /// The name of the Log Dashboard.
        /// </summary>
        [Input("dashboardName", required: true)]
        public Input<string> DashboardName { get; set; } = null!;

        /// <summary>
        /// Dashboard alias.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The name of the log project. It is the only in one Alicloud account.
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        public DashboardArgs()
        {
        }
        public static new DashboardArgs Empty => new DashboardArgs();
    }

    public sealed class DashboardState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dashboard attribute.
        /// </summary>
        [Input("attribute")]
        public Input<string>? Attribute { get; set; }

        /// <summary>
        /// Configuration of charts in the dashboard.
        /// </summary>
        [Input("charList")]
        public Input<string>? CharList { get; set; }

        /// <summary>
        /// The name of the Log Dashboard.
        /// </summary>
        [Input("dashboardName")]
        public Input<string>? DashboardName { get; set; }

        /// <summary>
        /// Dashboard alias.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The name of the log project. It is the only in one Alicloud account.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        public DashboardState()
        {
        }
        public static new DashboardState Empty => new DashboardState();
    }
}
