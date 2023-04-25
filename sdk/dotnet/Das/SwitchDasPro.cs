// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Das
{
    /// <summary>
    /// Provides a DAS Switch Das Pro resource.
    /// 
    /// For information about DAS Switch Das Pro and how to use it, see [What is Switch Das Pro](https://www.alibabacloud.com/help/en/database-autonomy-service/latest/enabledaspro).
    /// 
    /// &gt; **NOTE:** Available in v1.193.0+.
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
    ///     var @default = new AliCloud.Das.SwitchDasPro("default", new()
    ///     {
    ///         InstanceId = "your_sql_id",
    ///         SqlRetention = 30,
    ///         UserId = "your_account_id",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DAS Switch Das Pro can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:das/switchDasPro:SwitchDasPro example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:das/switchDasPro:SwitchDasPro")]
    public partial class SwitchDasPro : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the database instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The storage duration of SQL Explorer data. Valid values: `30`, `180`, `365`, `1095`, `1825`. Unit: days. Default value: `30`.
        /// </summary>
        [Output("sqlRetention")]
        public Output<int> SqlRetention { get; private set; } = null!;

        /// <summary>
        /// Whether the database instance has DAS professional.
        /// </summary>
        [Output("status")]
        public Output<bool> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the Alibaba Cloud account that is used to create the database instance.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a SwitchDasPro resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SwitchDasPro(string name, SwitchDasProArgs args, CustomResourceOptions? options = null)
            : base("alicloud:das/switchDasPro:SwitchDasPro", name, args ?? new SwitchDasProArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SwitchDasPro(string name, Input<string> id, SwitchDasProState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:das/switchDasPro:SwitchDasPro", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SwitchDasPro resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SwitchDasPro Get(string name, Input<string> id, SwitchDasProState? state = null, CustomResourceOptions? options = null)
        {
            return new SwitchDasPro(name, id, state, options);
        }
    }

    public sealed class SwitchDasProArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the database instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The storage duration of SQL Explorer data. Valid values: `30`, `180`, `365`, `1095`, `1825`. Unit: days. Default value: `30`.
        /// </summary>
        [Input("sqlRetention")]
        public Input<int>? SqlRetention { get; set; }

        /// <summary>
        /// The ID of the Alibaba Cloud account that is used to create the database instance.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public SwitchDasProArgs()
        {
        }
        public static new SwitchDasProArgs Empty => new SwitchDasProArgs();
    }

    public sealed class SwitchDasProState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the database instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The storage duration of SQL Explorer data. Valid values: `30`, `180`, `365`, `1095`, `1825`. Unit: days. Default value: `30`.
        /// </summary>
        [Input("sqlRetention")]
        public Input<int>? SqlRetention { get; set; }

        /// <summary>
        /// Whether the database instance has DAS professional.
        /// </summary>
        [Input("status")]
        public Input<bool>? Status { get; set; }

        /// <summary>
        /// The ID of the Alibaba Cloud account that is used to create the database instance.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public SwitchDasProState()
        {
        }
        public static new SwitchDasProState Empty => new SwitchDasProState();
    }
}
