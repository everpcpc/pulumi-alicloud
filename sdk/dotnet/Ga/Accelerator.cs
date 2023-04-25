// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga
{
    /// <summary>
    /// Provides a Global Accelerator (GA) Accelerator resource.
    /// 
    /// For information about Global Accelerator (GA) Accelerator and how to use it, see [What is Accelerator](https://help.aliyun.com/document_detail/153235.html).
    /// 
    /// &gt; **NOTE:** At present, The `alicloud.ga.Accelerator` cannot be deleted. you need to wait until the resource is outdated and released automatically.
    /// 
    /// &gt; **NOTE:** Available in v1.111.0+.
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
    ///     var example = new AliCloud.Ga.Accelerator("example", new()
    ///     {
    ///         AutoUseCoupon = true,
    ///         Duration = 1,
    ///         Spec = "1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Ga Accelerator can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ga/accelerator:Accelerator example &lt;accelerator_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ga/accelerator:Accelerator")]
    public partial class Accelerator : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Name of the GA instance.
        /// </summary>
        [Output("acceleratorName")]
        public Output<string?> AcceleratorName { get; private set; } = null!;

        /// <summary>
        /// Auto renewal period of an instance, in the unit of month. The value range is 1-12.
        /// </summary>
        [Output("autoRenewDuration")]
        public Output<int?> AutoRenewDuration { get; private set; } = null!;

        /// <summary>
        /// Use coupons to pay bills automatically. Default value is `false`. Valid value: `true`: Use, `false`: Not used.
        /// </summary>
        [Output("autoUseCoupon")]
        public Output<bool?> AutoUseCoupon { get; private set; } = null!;

        /// <summary>
        /// Descriptive information of the global acceleration instance.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The subscription duration. **NOTE:** Starting from v1.150.0+, the `duration` and  `pricing_cycle` are both required.
        /// * If the `pricing_cycle` parameter is set to `Month`, the valid values for the `duration` parameter are 1 to 9.
        /// * If the `pricing_cycle` parameter is set to `Year`, the valid values for the `duration` parameter are 1 to 3.
        /// </summary>
        [Output("duration")]
        public Output<int> Duration { get; private set; } = null!;

        /// <summary>
        /// The billing cycle of the GA instance. Valid values: `Month`,`Year`. The default value: `Month`.
        /// </summary>
        [Output("pricingCycle")]
        public Output<string> PricingCycle { get; private set; } = null!;

        /// <summary>
        /// Whether to renew an accelerator automatically or not. Default to "Normal". Valid values:
        /// </summary>
        [Output("renewalStatus")]
        public Output<string> RenewalStatus { get; private set; } = null!;

        /// <summary>
        /// The instance type of the GA instance. Specification of global acceleration instance, value:
        /// `1`: Small 1.
        /// `2`: Small 2.
        /// `3`: Small 3.
        /// `5`: Medium 1.
        /// `8`: Medium 2.
        /// `10`: Medium 3.
        /// </summary>
        [Output("spec")]
        public Output<string> Spec { get; private set; } = null!;

        /// <summary>
        /// The status of the GA instance.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Accelerator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Accelerator(string name, AcceleratorArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ga/accelerator:Accelerator", name, args ?? new AcceleratorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Accelerator(string name, Input<string> id, AcceleratorState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ga/accelerator:Accelerator", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Accelerator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Accelerator Get(string name, Input<string> id, AcceleratorState? state = null, CustomResourceOptions? options = null)
        {
            return new Accelerator(name, id, state, options);
        }
    }

    public sealed class AcceleratorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Name of the GA instance.
        /// </summary>
        [Input("acceleratorName")]
        public Input<string>? AcceleratorName { get; set; }

        /// <summary>
        /// Auto renewal period of an instance, in the unit of month. The value range is 1-12.
        /// </summary>
        [Input("autoRenewDuration")]
        public Input<int>? AutoRenewDuration { get; set; }

        /// <summary>
        /// Use coupons to pay bills automatically. Default value is `false`. Valid value: `true`: Use, `false`: Not used.
        /// </summary>
        [Input("autoUseCoupon")]
        public Input<bool>? AutoUseCoupon { get; set; }

        /// <summary>
        /// Descriptive information of the global acceleration instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The subscription duration. **NOTE:** Starting from v1.150.0+, the `duration` and  `pricing_cycle` are both required.
        /// * If the `pricing_cycle` parameter is set to `Month`, the valid values for the `duration` parameter are 1 to 9.
        /// * If the `pricing_cycle` parameter is set to `Year`, the valid values for the `duration` parameter are 1 to 3.
        /// </summary>
        [Input("duration", required: true)]
        public Input<int> Duration { get; set; } = null!;

        /// <summary>
        /// The billing cycle of the GA instance. Valid values: `Month`,`Year`. The default value: `Month`.
        /// </summary>
        [Input("pricingCycle")]
        public Input<string>? PricingCycle { get; set; }

        /// <summary>
        /// Whether to renew an accelerator automatically or not. Default to "Normal". Valid values:
        /// </summary>
        [Input("renewalStatus")]
        public Input<string>? RenewalStatus { get; set; }

        /// <summary>
        /// The instance type of the GA instance. Specification of global acceleration instance, value:
        /// `1`: Small 1.
        /// `2`: Small 2.
        /// `3`: Small 3.
        /// `5`: Medium 1.
        /// `8`: Medium 2.
        /// `10`: Medium 3.
        /// </summary>
        [Input("spec", required: true)]
        public Input<string> Spec { get; set; } = null!;

        public AcceleratorArgs()
        {
        }
        public static new AcceleratorArgs Empty => new AcceleratorArgs();
    }

    public sealed class AcceleratorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Name of the GA instance.
        /// </summary>
        [Input("acceleratorName")]
        public Input<string>? AcceleratorName { get; set; }

        /// <summary>
        /// Auto renewal period of an instance, in the unit of month. The value range is 1-12.
        /// </summary>
        [Input("autoRenewDuration")]
        public Input<int>? AutoRenewDuration { get; set; }

        /// <summary>
        /// Use coupons to pay bills automatically. Default value is `false`. Valid value: `true`: Use, `false`: Not used.
        /// </summary>
        [Input("autoUseCoupon")]
        public Input<bool>? AutoUseCoupon { get; set; }

        /// <summary>
        /// Descriptive information of the global acceleration instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The subscription duration. **NOTE:** Starting from v1.150.0+, the `duration` and  `pricing_cycle` are both required.
        /// * If the `pricing_cycle` parameter is set to `Month`, the valid values for the `duration` parameter are 1 to 9.
        /// * If the `pricing_cycle` parameter is set to `Year`, the valid values for the `duration` parameter are 1 to 3.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// The billing cycle of the GA instance. Valid values: `Month`,`Year`. The default value: `Month`.
        /// </summary>
        [Input("pricingCycle")]
        public Input<string>? PricingCycle { get; set; }

        /// <summary>
        /// Whether to renew an accelerator automatically or not. Default to "Normal". Valid values:
        /// </summary>
        [Input("renewalStatus")]
        public Input<string>? RenewalStatus { get; set; }

        /// <summary>
        /// The instance type of the GA instance. Specification of global acceleration instance, value:
        /// `1`: Small 1.
        /// `2`: Small 2.
        /// `3`: Small 3.
        /// `5`: Medium 1.
        /// `8`: Medium 2.
        /// `10`: Medium 3.
        /// </summary>
        [Input("spec")]
        public Input<string>? Spec { get; set; }

        /// <summary>
        /// The status of the GA instance.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public AcceleratorState()
        {
        }
        public static new AcceleratorState Empty => new AcceleratorState();
    }
}
