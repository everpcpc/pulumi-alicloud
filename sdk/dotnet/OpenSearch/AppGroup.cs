// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.OpenSearch
{
    /// <summary>
    /// Provides a Open Search App Group resource.
    /// 
    /// For information about Open Search App Group and how to use it, see [What is App Group](https://www.aliyun.com/product/opensearch).
    /// 
    /// &gt; **NOTE:** Available in v1.136.0+.
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
    ///     var name = config.Get("name") ?? "name";
    ///     var @default = new AliCloud.OpenSearch.AppGroup("default", new()
    ///     {
    ///         AppGroupName = name,
    ///         PaymentType = "PayAsYouGo",
    ///         Type = "standard",
    ///         Quota = new AliCloud.OpenSearch.Inputs.AppGroupQuotaArgs
    ///         {
    ///             DocSize = 1,
    ///             ComputeResource = 20,
    ///             Spec = "opensearch.share.common",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Open Search App Group can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:opensearch/appGroup:AppGroup example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:opensearch/appGroup:AppGroup")]
    public partial class AppGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Application Group Name.
        /// </summary>
        [Output("appGroupName")]
        public Output<string> AppGroupName { get; private set; } = null!;

        /// <summary>
        /// Billing model. Valid values:`compute_resource` and `qps`.
        /// </summary>
        [Output("chargeWay")]
        public Output<string> ChargeWay { get; private set; } = null!;

        /// <summary>
        /// The version of Application Group Name.
        /// </summary>
        [Output("currentVersion")]
        public Output<string?> CurrentVersion { get; private set; } = null!;

        /// <summary>
        /// The description of the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Order change type. Valid values: `UPGRADE` and `DOWNGRADE`.
        /// </summary>
        [Output("orderType")]
        public Output<string> OrderType { get; private set; } = null!;

        /// <summary>
        /// Order cycle information. The details see Block order.
        /// </summary>
        [Output("orders")]
        public Output<ImmutableArray<Outputs.AppGroupOrder>> Orders { get; private set; } = null!;

        /// <summary>
        /// The billing method of the resource. Valid values: `Subscription` and `PayAsYouGo`.
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// Quota information.  The details see Block quota.
        /// </summary>
        [Output("quota")]
        public Output<Outputs.AppGroupQuota> Quota { get; private set; } = null!;

        /// <summary>
        /// The status of the resource. Valid values: `producing`,`review_pending`,`config_pending`,`normal`,`frozen`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Application type. Valid Values: `standard`, `enhanced`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AppGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppGroup(string name, AppGroupArgs args, CustomResourceOptions? options = null)
            : base("alicloud:opensearch/appGroup:AppGroup", name, args ?? new AppGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppGroup(string name, Input<string> id, AppGroupState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:opensearch/appGroup:AppGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppGroup Get(string name, Input<string> id, AppGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new AppGroup(name, id, state, options);
        }
    }

    public sealed class AppGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application Group Name.
        /// </summary>
        [Input("appGroupName", required: true)]
        public Input<string> AppGroupName { get; set; } = null!;

        /// <summary>
        /// Billing model. Valid values:`compute_resource` and `qps`.
        /// </summary>
        [Input("chargeWay")]
        public Input<string>? ChargeWay { get; set; }

        /// <summary>
        /// The version of Application Group Name.
        /// </summary>
        [Input("currentVersion")]
        public Input<string>? CurrentVersion { get; set; }

        /// <summary>
        /// The description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Order change type. Valid values: `UPGRADE` and `DOWNGRADE`.
        /// </summary>
        [Input("orderType")]
        public Input<string>? OrderType { get; set; }

        [Input("orders")]
        private InputList<Inputs.AppGroupOrderArgs>? _orders;

        /// <summary>
        /// Order cycle information. The details see Block order.
        /// </summary>
        public InputList<Inputs.AppGroupOrderArgs> Orders
        {
            get => _orders ?? (_orders = new InputList<Inputs.AppGroupOrderArgs>());
            set => _orders = value;
        }

        /// <summary>
        /// The billing method of the resource. Valid values: `Subscription` and `PayAsYouGo`.
        /// </summary>
        [Input("paymentType", required: true)]
        public Input<string> PaymentType { get; set; } = null!;

        /// <summary>
        /// Quota information.  The details see Block quota.
        /// </summary>
        [Input("quota", required: true)]
        public Input<Inputs.AppGroupQuotaArgs> Quota { get; set; } = null!;

        /// <summary>
        /// Application type. Valid Values: `standard`, `enhanced`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AppGroupArgs()
        {
        }
        public static new AppGroupArgs Empty => new AppGroupArgs();
    }

    public sealed class AppGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application Group Name.
        /// </summary>
        [Input("appGroupName")]
        public Input<string>? AppGroupName { get; set; }

        /// <summary>
        /// Billing model. Valid values:`compute_resource` and `qps`.
        /// </summary>
        [Input("chargeWay")]
        public Input<string>? ChargeWay { get; set; }

        /// <summary>
        /// The version of Application Group Name.
        /// </summary>
        [Input("currentVersion")]
        public Input<string>? CurrentVersion { get; set; }

        /// <summary>
        /// The description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Order change type. Valid values: `UPGRADE` and `DOWNGRADE`.
        /// </summary>
        [Input("orderType")]
        public Input<string>? OrderType { get; set; }

        [Input("orders")]
        private InputList<Inputs.AppGroupOrderGetArgs>? _orders;

        /// <summary>
        /// Order cycle information. The details see Block order.
        /// </summary>
        public InputList<Inputs.AppGroupOrderGetArgs> Orders
        {
            get => _orders ?? (_orders = new InputList<Inputs.AppGroupOrderGetArgs>());
            set => _orders = value;
        }

        /// <summary>
        /// The billing method of the resource. Valid values: `Subscription` and `PayAsYouGo`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// Quota information.  The details see Block quota.
        /// </summary>
        [Input("quota")]
        public Input<Inputs.AppGroupQuotaGetArgs>? Quota { get; set; }

        /// <summary>
        /// The status of the resource. Valid values: `producing`,`review_pending`,`config_pending`,`normal`,`frozen`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Application type. Valid Values: `standard`, `enhanced`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AppGroupState()
        {
        }
        public static new AppGroupState Empty => new AppGroupState();
    }
}
