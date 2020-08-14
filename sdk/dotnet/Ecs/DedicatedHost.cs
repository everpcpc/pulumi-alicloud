// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public partial class DedicatedHost : Pulumi.CustomResource
    {
        /// <summary>
        /// The policy used to migrate the instances from the dedicated host when the dedicated host fails or needs to be repaired online. Valid values: `Migrate`, `Stop`.
        /// </summary>
        [Output("actionOnMaintenance")]
        public Output<string?> ActionOnMaintenance { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to add the dedicated host to the resource pool for automatic deployment. If you do not specify the DedicatedHostId parameter when you create an instance on a dedicated host, Alibaba Cloud automatically selects a dedicated host from the resource pool to host the instance. Valid values: `on`, `off`. Default: `on`.
        /// </summary>
        [Output("autoPlacement")]
        public Output<string?> AutoPlacement { get; private set; } = null!;

        /// <summary>
        /// The automatic release time of the dedicated host. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
        /// </summary>
        [Output("autoReleaseTime")]
        public Output<string?> AutoReleaseTime { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to automatically renew the subscription dedicated host.
        /// </summary>
        [Output("autoRenew")]
        public Output<bool?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// The auto-renewal period of the dedicated host. Unit: months. Valid values: `1`, `2`, `3`, `6`, and `12`. takes effect and is required only when the AutoRenew parameter is set to true.
        /// </summary>
        [Output("autoRenewPeriod")]
        public Output<int?> AutoRenewPeriod { get; private set; } = null!;

        /// <summary>
        /// The name of the dedicated host. The name must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        /// </summary>
        [Output("dedicatedHostName")]
        public Output<string?> DedicatedHostName { get; private set; } = null!;

        /// <summary>
        /// The type of the dedicated host. You can call the [DescribeDedicatedHostTypes](https://www.alibabacloud.com/help/doc-detail/134240.htm) operation to obtain the most recent list of dedicated host types.
        /// </summary>
        [Output("dedicatedHostType")]
        public Output<string> DedicatedHostType { get; private set; } = null!;

        /// <summary>
        /// The description of the dedicated host. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to return the billing details of the order when the billing method is changed from subscription to pay-as-you-go. Default: `false`.
        /// </summary>
        [Output("detailFee")]
        public Output<bool?> DetailFee { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to only validate the request. Default: `false`.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The subscription period of the dedicated host. The Period parameter takes effect and is required only when the ChargeType parameter is set to PrePaid.
        /// </summary>
        [Output("expiredTime")]
        public Output<string> ExpiredTime { get; private set; } = null!;

        /// <summary>
        /// dedicated host network parameters. contains the following attributes:
        /// </summary>
        [Output("networkAttributes")]
        public Output<ImmutableArray<Outputs.DedicatedHostNetworkAttribute>> NetworkAttributes { get; private set; } = null!;

        /// <summary>
        /// The billing method of the dedicated host. Valid values: `PrePaid`, `PostPaid`. Default: `PostPaid`.
        /// </summary>
        [Output("paymentType")]
        public Output<string?> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group to which the dedicated host belongs.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The unit of the subscription period of the dedicated host.
        /// </summary>
        [Output("saleCycle")]
        public Output<string> SaleCycle { get; private set; } = null!;

        /// <summary>
        /// The status of the dedicated host.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The zone ID of the dedicated host. This parameter is empty by default. If you do not specify this parameter, the system automatically selects a zone.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a DedicatedHost resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DedicatedHost(string name, DedicatedHostArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ecs/dedicatedHost:DedicatedHost", name, args ?? new DedicatedHostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DedicatedHost(string name, Input<string> id, DedicatedHostState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/dedicatedHost:DedicatedHost", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DedicatedHost resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DedicatedHost Get(string name, Input<string> id, DedicatedHostState? state = null, CustomResourceOptions? options = null)
        {
            return new DedicatedHost(name, id, state, options);
        }
    }

    public sealed class DedicatedHostArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy used to migrate the instances from the dedicated host when the dedicated host fails or needs to be repaired online. Valid values: `Migrate`, `Stop`.
        /// </summary>
        [Input("actionOnMaintenance")]
        public Input<string>? ActionOnMaintenance { get; set; }

        /// <summary>
        /// Specifies whether to add the dedicated host to the resource pool for automatic deployment. If you do not specify the DedicatedHostId parameter when you create an instance on a dedicated host, Alibaba Cloud automatically selects a dedicated host from the resource pool to host the instance. Valid values: `on`, `off`. Default: `on`.
        /// </summary>
        [Input("autoPlacement")]
        public Input<string>? AutoPlacement { get; set; }

        /// <summary>
        /// The automatic release time of the dedicated host. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
        /// </summary>
        [Input("autoReleaseTime")]
        public Input<string>? AutoReleaseTime { get; set; }

        /// <summary>
        /// Specifies whether to automatically renew the subscription dedicated host.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The auto-renewal period of the dedicated host. Unit: months. Valid values: `1`, `2`, `3`, `6`, and `12`. takes effect and is required only when the AutoRenew parameter is set to true.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// The name of the dedicated host. The name must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        /// </summary>
        [Input("dedicatedHostName")]
        public Input<string>? DedicatedHostName { get; set; }

        /// <summary>
        /// The type of the dedicated host. You can call the [DescribeDedicatedHostTypes](https://www.alibabacloud.com/help/doc-detail/134240.htm) operation to obtain the most recent list of dedicated host types.
        /// </summary>
        [Input("dedicatedHostType", required: true)]
        public Input<string> DedicatedHostType { get; set; } = null!;

        /// <summary>
        /// The description of the dedicated host. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether to return the billing details of the order when the billing method is changed from subscription to pay-as-you-go. Default: `false`.
        /// </summary>
        [Input("detailFee")]
        public Input<bool>? DetailFee { get; set; }

        /// <summary>
        /// Specifies whether to only validate the request. Default: `false`.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The subscription period of the dedicated host. The Period parameter takes effect and is required only when the ChargeType parameter is set to PrePaid.
        /// </summary>
        [Input("expiredTime")]
        public Input<string>? ExpiredTime { get; set; }

        [Input("networkAttributes")]
        private InputList<Inputs.DedicatedHostNetworkAttributeArgs>? _networkAttributes;

        /// <summary>
        /// dedicated host network parameters. contains the following attributes:
        /// </summary>
        public InputList<Inputs.DedicatedHostNetworkAttributeArgs> NetworkAttributes
        {
            get => _networkAttributes ?? (_networkAttributes = new InputList<Inputs.DedicatedHostNetworkAttributeArgs>());
            set => _networkAttributes = value;
        }

        /// <summary>
        /// The billing method of the dedicated host. Valid values: `PrePaid`, `PostPaid`. Default: `PostPaid`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The ID of the resource group to which the dedicated host belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The unit of the subscription period of the dedicated host.
        /// </summary>
        [Input("saleCycle")]
        public Input<string>? SaleCycle { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone ID of the dedicated host. This parameter is empty by default. If you do not specify this parameter, the system automatically selects a zone.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public DedicatedHostArgs()
        {
        }
    }

    public sealed class DedicatedHostState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy used to migrate the instances from the dedicated host when the dedicated host fails or needs to be repaired online. Valid values: `Migrate`, `Stop`.
        /// </summary>
        [Input("actionOnMaintenance")]
        public Input<string>? ActionOnMaintenance { get; set; }

        /// <summary>
        /// Specifies whether to add the dedicated host to the resource pool for automatic deployment. If you do not specify the DedicatedHostId parameter when you create an instance on a dedicated host, Alibaba Cloud automatically selects a dedicated host from the resource pool to host the instance. Valid values: `on`, `off`. Default: `on`.
        /// </summary>
        [Input("autoPlacement")]
        public Input<string>? AutoPlacement { get; set; }

        /// <summary>
        /// The automatic release time of the dedicated host. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
        /// </summary>
        [Input("autoReleaseTime")]
        public Input<string>? AutoReleaseTime { get; set; }

        /// <summary>
        /// Specifies whether to automatically renew the subscription dedicated host.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The auto-renewal period of the dedicated host. Unit: months. Valid values: `1`, `2`, `3`, `6`, and `12`. takes effect and is required only when the AutoRenew parameter is set to true.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// The name of the dedicated host. The name must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        /// </summary>
        [Input("dedicatedHostName")]
        public Input<string>? DedicatedHostName { get; set; }

        /// <summary>
        /// The type of the dedicated host. You can call the [DescribeDedicatedHostTypes](https://www.alibabacloud.com/help/doc-detail/134240.htm) operation to obtain the most recent list of dedicated host types.
        /// </summary>
        [Input("dedicatedHostType")]
        public Input<string>? DedicatedHostType { get; set; }

        /// <summary>
        /// The description of the dedicated host. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether to return the billing details of the order when the billing method is changed from subscription to pay-as-you-go. Default: `false`.
        /// </summary>
        [Input("detailFee")]
        public Input<bool>? DetailFee { get; set; }

        /// <summary>
        /// Specifies whether to only validate the request. Default: `false`.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The subscription period of the dedicated host. The Period parameter takes effect and is required only when the ChargeType parameter is set to PrePaid.
        /// </summary>
        [Input("expiredTime")]
        public Input<string>? ExpiredTime { get; set; }

        [Input("networkAttributes")]
        private InputList<Inputs.DedicatedHostNetworkAttributeGetArgs>? _networkAttributes;

        /// <summary>
        /// dedicated host network parameters. contains the following attributes:
        /// </summary>
        public InputList<Inputs.DedicatedHostNetworkAttributeGetArgs> NetworkAttributes
        {
            get => _networkAttributes ?? (_networkAttributes = new InputList<Inputs.DedicatedHostNetworkAttributeGetArgs>());
            set => _networkAttributes = value;
        }

        /// <summary>
        /// The billing method of the dedicated host. Valid values: `PrePaid`, `PostPaid`. Default: `PostPaid`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The ID of the resource group to which the dedicated host belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The unit of the subscription period of the dedicated host.
        /// </summary>
        [Input("saleCycle")]
        public Input<string>? SaleCycle { get; set; }

        /// <summary>
        /// The status of the dedicated host.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone ID of the dedicated host. This parameter is empty by default. If you do not specify this parameter, the system automatically selects a zone.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public DedicatedHostState()
        {
        }
    }
}
