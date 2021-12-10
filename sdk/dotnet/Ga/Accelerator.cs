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
    /// &gt; **NOTE:** Available in v1.111.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new AliCloud.Ga.Accelerator("example", new AliCloud.Ga.AcceleratorArgs
    ///         {
    ///             AutoUseCoupon = true,
    ///             Duration = 1,
    ///             Spec = "1",
    ///         });
    ///     }
    /// 
    /// }
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
    public partial class Accelerator : Pulumi.CustomResource
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
        /// The duration. The value range is 1-9.
        /// </summary>
        [Output("duration")]
        public Output<int> Duration { get; private set; } = null!;

        /// <summary>
        /// Whether to renew an accelerator automatically or not. Default to "Normal". Valid values:
        /// - `AutoRenewal`: Enable auto renewal.
        /// - `Normal`: Disable auto renewal.
        /// - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
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

    public sealed class AcceleratorArgs : Pulumi.ResourceArgs
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
        /// The duration. The value range is 1-9.
        /// </summary>
        [Input("duration", required: true)]
        public Input<int> Duration { get; set; } = null!;

        /// <summary>
        /// Whether to renew an accelerator automatically or not. Default to "Normal". Valid values:
        /// - `AutoRenewal`: Enable auto renewal.
        /// - `Normal`: Disable auto renewal.
        /// - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
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
    }

    public sealed class AcceleratorState : Pulumi.ResourceArgs
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
        /// The duration. The value range is 1-9.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// Whether to renew an accelerator automatically or not. Default to "Normal". Valid values:
        /// - `AutoRenewal`: Enable auto renewal.
        /// - `Normal`: Disable auto renewal.
        /// - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
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
    }
}
