// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Quotas
{
    /// <summary>
    /// Provides a Quotas Quota Application resource.
    /// 
    /// For information about Quotas Quota Application and how to use it, see [What is Quota Application](https://help.aliyun.com/document_detail/171289.html).
    /// 
    /// &gt; **NOTE:** Available in v1.117.0+.
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
    ///     var example = new AliCloud.Quotas.QuotaApplication("example", new()
    ///     {
    ///         DesireValue = 100,
    ///         Dimensions = new[]
    ///         {
    ///             new AliCloud.Quotas.Inputs.QuotaApplicationDimensionArgs
    ///             {
    ///                 Key = "regionId",
    ///                 Value = "cn-hangzhou",
    ///             },
    ///         },
    ///         NoticeType = 0,
    ///         ProductCode = "ess",
    ///         QuotaActionCode = "q_db_instance",
    ///         Reason = "For Terraform Test",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Quotas Application Info can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:quotas/quotaApplication:QuotaApplication example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:quotas/quotaApplication:QuotaApplication")]
    public partial class QuotaApplication : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The approve value of the quota application.
        /// </summary>
        [Output("approveValue")]
        public Output<string> ApproveValue { get; private set; } = null!;

        /// <summary>
        /// The audit mode. Valid values: `Async`, `Sync`. Default to: `Async`.
        /// </summary>
        [Output("auditMode")]
        public Output<string?> AuditMode { get; private set; } = null!;

        /// <summary>
        /// The audit reason.
        /// </summary>
        [Output("auditReason")]
        public Output<string> AuditReason { get; private set; } = null!;

        /// <summary>
        /// The desire value of the quota application.
        /// </summary>
        [Output("desireValue")]
        public Output<double> DesireValue { get; private set; } = null!;

        /// <summary>
        /// The quota dimensions.
        /// </summary>
        [Output("dimensions")]
        public Output<ImmutableArray<Outputs.QuotaApplicationDimension>> Dimensions { get; private set; } = null!;

        /// <summary>
        /// The effective time of the quota application.
        /// </summary>
        [Output("effectiveTime")]
        public Output<string> EffectiveTime { get; private set; } = null!;

        /// <summary>
        /// The expire time of the quota application.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// The notice type. Valid values: `0`, `1`, `2`, `3`.
        /// </summary>
        [Output("noticeType")]
        public Output<int?> NoticeType { get; private set; } = null!;

        /// <summary>
        /// The product code.
        /// </summary>
        [Output("productCode")]
        public Output<string> ProductCode { get; private set; } = null!;

        /// <summary>
        /// The ID of quota action.
        /// </summary>
        [Output("quotaActionCode")]
        public Output<string> QuotaActionCode { get; private set; } = null!;

        /// <summary>
        /// The quota category. Valid values: `CommonQuota`, `FlowControl`.
        /// </summary>
        [Output("quotaCategory")]
        public Output<string?> QuotaCategory { get; private set; } = null!;

        /// <summary>
        /// The description of the quota application.
        /// </summary>
        [Output("quotaDescription")]
        public Output<string> QuotaDescription { get; private set; } = null!;

        /// <summary>
        /// The name of the quota application.
        /// </summary>
        [Output("quotaName")]
        public Output<string> QuotaName { get; private set; } = null!;

        /// <summary>
        /// The unit of the quota application.
        /// </summary>
        [Output("quotaUnit")]
        public Output<string> QuotaUnit { get; private set; } = null!;

        /// <summary>
        /// The reason of the quota application.
        /// </summary>
        [Output("reason")]
        public Output<string> Reason { get; private set; } = null!;

        /// <summary>
        /// The status of the quota application.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a QuotaApplication resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QuotaApplication(string name, QuotaApplicationArgs args, CustomResourceOptions? options = null)
            : base("alicloud:quotas/quotaApplication:QuotaApplication", name, args ?? new QuotaApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private QuotaApplication(string name, Input<string> id, QuotaApplicationState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:quotas/quotaApplication:QuotaApplication", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing QuotaApplication resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QuotaApplication Get(string name, Input<string> id, QuotaApplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new QuotaApplication(name, id, state, options);
        }
    }

    public sealed class QuotaApplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The audit mode. Valid values: `Async`, `Sync`. Default to: `Async`.
        /// </summary>
        [Input("auditMode")]
        public Input<string>? AuditMode { get; set; }

        /// <summary>
        /// The desire value of the quota application.
        /// </summary>
        [Input("desireValue", required: true)]
        public Input<double> DesireValue { get; set; } = null!;

        [Input("dimensions")]
        private InputList<Inputs.QuotaApplicationDimensionArgs>? _dimensions;

        /// <summary>
        /// The quota dimensions.
        /// </summary>
        public InputList<Inputs.QuotaApplicationDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.QuotaApplicationDimensionArgs>());
            set => _dimensions = value;
        }

        /// <summary>
        /// The notice type. Valid values: `0`, `1`, `2`, `3`.
        /// </summary>
        [Input("noticeType")]
        public Input<int>? NoticeType { get; set; }

        /// <summary>
        /// The product code.
        /// </summary>
        [Input("productCode", required: true)]
        public Input<string> ProductCode { get; set; } = null!;

        /// <summary>
        /// The ID of quota action.
        /// </summary>
        [Input("quotaActionCode", required: true)]
        public Input<string> QuotaActionCode { get; set; } = null!;

        /// <summary>
        /// The quota category. Valid values: `CommonQuota`, `FlowControl`.
        /// </summary>
        [Input("quotaCategory")]
        public Input<string>? QuotaCategory { get; set; }

        /// <summary>
        /// The reason of the quota application.
        /// </summary>
        [Input("reason", required: true)]
        public Input<string> Reason { get; set; } = null!;

        public QuotaApplicationArgs()
        {
        }
        public static new QuotaApplicationArgs Empty => new QuotaApplicationArgs();
    }

    public sealed class QuotaApplicationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The approve value of the quota application.
        /// </summary>
        [Input("approveValue")]
        public Input<string>? ApproveValue { get; set; }

        /// <summary>
        /// The audit mode. Valid values: `Async`, `Sync`. Default to: `Async`.
        /// </summary>
        [Input("auditMode")]
        public Input<string>? AuditMode { get; set; }

        /// <summary>
        /// The audit reason.
        /// </summary>
        [Input("auditReason")]
        public Input<string>? AuditReason { get; set; }

        /// <summary>
        /// The desire value of the quota application.
        /// </summary>
        [Input("desireValue")]
        public Input<double>? DesireValue { get; set; }

        [Input("dimensions")]
        private InputList<Inputs.QuotaApplicationDimensionGetArgs>? _dimensions;

        /// <summary>
        /// The quota dimensions.
        /// </summary>
        public InputList<Inputs.QuotaApplicationDimensionGetArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.QuotaApplicationDimensionGetArgs>());
            set => _dimensions = value;
        }

        /// <summary>
        /// The effective time of the quota application.
        /// </summary>
        [Input("effectiveTime")]
        public Input<string>? EffectiveTime { get; set; }

        /// <summary>
        /// The expire time of the quota application.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// The notice type. Valid values: `0`, `1`, `2`, `3`.
        /// </summary>
        [Input("noticeType")]
        public Input<int>? NoticeType { get; set; }

        /// <summary>
        /// The product code.
        /// </summary>
        [Input("productCode")]
        public Input<string>? ProductCode { get; set; }

        /// <summary>
        /// The ID of quota action.
        /// </summary>
        [Input("quotaActionCode")]
        public Input<string>? QuotaActionCode { get; set; }

        /// <summary>
        /// The quota category. Valid values: `CommonQuota`, `FlowControl`.
        /// </summary>
        [Input("quotaCategory")]
        public Input<string>? QuotaCategory { get; set; }

        /// <summary>
        /// The description of the quota application.
        /// </summary>
        [Input("quotaDescription")]
        public Input<string>? QuotaDescription { get; set; }

        /// <summary>
        /// The name of the quota application.
        /// </summary>
        [Input("quotaName")]
        public Input<string>? QuotaName { get; set; }

        /// <summary>
        /// The unit of the quota application.
        /// </summary>
        [Input("quotaUnit")]
        public Input<string>? QuotaUnit { get; set; }

        /// <summary>
        /// The reason of the quota application.
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        /// <summary>
        /// The status of the quota application.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public QuotaApplicationState()
        {
        }
        public static new QuotaApplicationState Empty => new QuotaApplicationState();
    }
}
