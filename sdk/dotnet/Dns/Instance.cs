// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dns
{
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// DNS security level. Valid values: `no`, `basic`, `advanced`.
        /// </summary>
        [Output("dnsSecurity")]
        public Output<string> DnsSecurity { get; private set; } = null!;

        /// <summary>
        /// Number of domain names bound.
        /// </summary>
        [Output("domainNumbers")]
        public Output<string> DomainNumbers { get; private set; } = null!;

        [Output("paymentType")]
        public Output<string?> PaymentType { get; private set; } = null!;

        /// <summary>
        /// Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// Automatic renewal period, the unit is month. When setting RenewalStatus to AutoRenewal, it must be set.
        /// </summary>
        [Output("renewPeriod")]
        public Output<int?> RenewPeriod { get; private set; } = null!;

        /// <summary>
        /// Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, default to `ManualRenewal`.
        /// </summary>
        [Output("renewalStatus")]
        public Output<string?> RenewalStatus { get; private set; } = null!;

        /// <summary>
        /// Paid package version. Valid values: `version_personal`, `version_enterprise_basic`, `version_enterprise_advanced`.
        /// </summary>
        [Output("versionCode")]
        public Output<string> VersionCode { get; private set; } = null!;

        /// <summary>
        /// Paid package version name.
        /// </summary>
        [Output("versionName")]
        public Output<string> VersionName { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dns/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dns/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS security level. Valid values: `no`, `basic`, `advanced`.
        /// </summary>
        [Input("dnsSecurity", required: true)]
        public Input<string> DnsSecurity { get; set; } = null!;

        /// <summary>
        /// Number of domain names bound.
        /// </summary>
        [Input("domainNumbers", required: true)]
        public Input<string> DomainNumbers { get; set; } = null!;

        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Automatic renewal period, the unit is month. When setting RenewalStatus to AutoRenewal, it must be set.
        /// </summary>
        [Input("renewPeriod")]
        public Input<int>? RenewPeriod { get; set; }

        /// <summary>
        /// Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, default to `ManualRenewal`.
        /// </summary>
        [Input("renewalStatus")]
        public Input<string>? RenewalStatus { get; set; }

        /// <summary>
        /// Paid package version. Valid values: `version_personal`, `version_enterprise_basic`, `version_enterprise_advanced`.
        /// </summary>
        [Input("versionCode", required: true)]
        public Input<string> VersionCode { get; set; } = null!;

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS security level. Valid values: `no`, `basic`, `advanced`.
        /// </summary>
        [Input("dnsSecurity")]
        public Input<string>? DnsSecurity { get; set; }

        /// <summary>
        /// Number of domain names bound.
        /// </summary>
        [Input("domainNumbers")]
        public Input<string>? DomainNumbers { get; set; }

        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Automatic renewal period, the unit is month. When setting RenewalStatus to AutoRenewal, it must be set.
        /// </summary>
        [Input("renewPeriod")]
        public Input<int>? RenewPeriod { get; set; }

        /// <summary>
        /// Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, default to `ManualRenewal`.
        /// </summary>
        [Input("renewalStatus")]
        public Input<string>? RenewalStatus { get; set; }

        /// <summary>
        /// Paid package version. Valid values: `version_personal`, `version_enterprise_basic`, `version_enterprise_advanced`.
        /// </summary>
        [Input("versionCode")]
        public Input<string>? VersionCode { get; set; }

        /// <summary>
        /// Paid package version name.
        /// </summary>
        [Input("versionName")]
        public Input<string>? VersionName { get; set; }

        public InstanceState()
        {
        }
    }
}
