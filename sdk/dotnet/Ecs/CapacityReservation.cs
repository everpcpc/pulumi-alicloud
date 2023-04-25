// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    /// <summary>
    /// Provides a Ecs Capacity Reservation resource.
    /// 
    /// For information about Ecs Capacity Reservation and how to use it, see [What is Capacity Reservation](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/createcapacityreservation).
    /// 
    /// &gt; **NOTE:** Available in v1.195.0+.
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
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "Instance",
    ///     });
    /// 
    ///     var defaultResourceGroups = AliCloud.ResourceManager.GetResourceGroups.Invoke(new()
    ///     {
    ///         Status = "OK",
    ///     });
    /// 
    ///     var defaultCapacityReservation = new AliCloud.Ecs.CapacityReservation("defaultCapacityReservation", new()
    ///     {
    ///         Description = @var.Name,
    ///         Platform = "linux",
    ///         CapacityReservationName = @var.Name,
    ///         EndTimeType = "Unlimited",
    ///         ResourceGroupId = defaultResourceGroups.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Ids[0]),
    ///         InstanceAmount = 1,
    ///         InstanceType = "ecs.c5.2xlarge",
    ///         MatchCriteria = "Open",
    ///         Tags = 
    ///         {
    ///             { "Created", "tfTestAcc0" },
    ///             { "For", "Tftestacc 0" },
    ///         },
    ///         ZoneIds = new[]
    ///         {
    ///             defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Ecs Capacity Reservation can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ecs/capacityReservation:CapacityReservation example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ecs/capacityReservation:CapacityReservation")]
    public partial class CapacityReservation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Capacity reservation service name.
        /// </summary>
        [Output("capacityReservationName")]
        public Output<string> CapacityReservationName { get; private set; } = null!;

        /// <summary>
        /// description of the capacity reservation instance.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to pre-check the API request. Valid values: `true` and `false`.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// end time of the capacity reservation. the capacity reservation will be  released at the end time automatically if set. otherwise it will last until manually released
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// Release mode of capacity reservation service. Value range:Limited: release at specified time. The EndTime parameter must be specified at the same time.Unlimited: manual release. No time limit.
        /// </summary>
        [Output("endTimeType")]
        public Output<string> EndTimeType { get; private set; } = null!;

        /// <summary>
        /// The total number of instances that need to be reserved within the capacity reservation.
        /// </summary>
        [Output("instanceAmount")]
        public Output<int> InstanceAmount { get; private set; } = null!;

        /// <summary>
        /// Instance type. Currently, you can only set the capacity reservation service for one instance type.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The type of private resource pool generated after the capacity reservation service takes effect. Value range:Open: Open mode.Target: dedicated mode.Default value: Open
        /// </summary>
        [Output("matchCriteria")]
        public Output<string> MatchCriteria { get; private set; } = null!;

        /// <summary>
        /// The payment type of the resource
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// platform of the capacity reservation, value range `windows`, `linux`.
        /// </summary>
        [Output("platform")]
        public Output<string> Platform { get; private set; } = null!;

        /// <summary>
        /// The resource group id.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string?> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// time of the capacity reservation which become active.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// The capacity is scheduled to take effect. Possible values:-Now: Effective immediately.-Later: the specified time takes effect.
        /// </summary>
        [Output("startTimeType")]
        public Output<string> StartTimeType { get; private set; } = null!;

        /// <summary>
        /// The status of the capacity reservation.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// This parameter is under test and is not yet open for use.
        /// </summary>
        [Output("timeSlot")]
        public Output<string> TimeSlot { get; private set; } = null!;

        /// <summary>
        /// The ID of the zone in the region to which the capacity reservation service belongs. Currently, it is only supported to create a capacity reservation service in one zone.
        /// </summary>
        [Output("zoneIds")]
        public Output<ImmutableArray<string>> ZoneIds { get; private set; } = null!;


        /// <summary>
        /// Create a CapacityReservation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CapacityReservation(string name, CapacityReservationArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ecs/capacityReservation:CapacityReservation", name, args ?? new CapacityReservationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CapacityReservation(string name, Input<string> id, CapacityReservationState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/capacityReservation:CapacityReservation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CapacityReservation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CapacityReservation Get(string name, Input<string> id, CapacityReservationState? state = null, CustomResourceOptions? options = null)
        {
            return new CapacityReservation(name, id, state, options);
        }
    }

    public sealed class CapacityReservationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Capacity reservation service name.
        /// </summary>
        [Input("capacityReservationName")]
        public Input<string>? CapacityReservationName { get; set; }

        /// <summary>
        /// description of the capacity reservation instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether to pre-check the API request. Valid values: `true` and `false`.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// end time of the capacity reservation. the capacity reservation will be  released at the end time automatically if set. otherwise it will last until manually released
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Release mode of capacity reservation service. Value range:Limited: release at specified time. The EndTime parameter must be specified at the same time.Unlimited: manual release. No time limit.
        /// </summary>
        [Input("endTimeType")]
        public Input<string>? EndTimeType { get; set; }

        /// <summary>
        /// The total number of instances that need to be reserved within the capacity reservation.
        /// </summary>
        [Input("instanceAmount", required: true)]
        public Input<int> InstanceAmount { get; set; } = null!;

        /// <summary>
        /// Instance type. Currently, you can only set the capacity reservation service for one instance type.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The type of private resource pool generated after the capacity reservation service takes effect. Value range:Open: Open mode.Target: dedicated mode.Default value: Open
        /// </summary>
        [Input("matchCriteria")]
        public Input<string>? MatchCriteria { get; set; }

        /// <summary>
        /// platform of the capacity reservation, value range `windows`, `linux`.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// The resource group id.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("zoneIds", required: true)]
        private InputList<string>? _zoneIds;

        /// <summary>
        /// The ID of the zone in the region to which the capacity reservation service belongs. Currently, it is only supported to create a capacity reservation service in one zone.
        /// </summary>
        public InputList<string> ZoneIds
        {
            get => _zoneIds ?? (_zoneIds = new InputList<string>());
            set => _zoneIds = value;
        }

        public CapacityReservationArgs()
        {
        }
        public static new CapacityReservationArgs Empty => new CapacityReservationArgs();
    }

    public sealed class CapacityReservationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Capacity reservation service name.
        /// </summary>
        [Input("capacityReservationName")]
        public Input<string>? CapacityReservationName { get; set; }

        /// <summary>
        /// description of the capacity reservation instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether to pre-check the API request. Valid values: `true` and `false`.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// end time of the capacity reservation. the capacity reservation will be  released at the end time automatically if set. otherwise it will last until manually released
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Release mode of capacity reservation service. Value range:Limited: release at specified time. The EndTime parameter must be specified at the same time.Unlimited: manual release. No time limit.
        /// </summary>
        [Input("endTimeType")]
        public Input<string>? EndTimeType { get; set; }

        /// <summary>
        /// The total number of instances that need to be reserved within the capacity reservation.
        /// </summary>
        [Input("instanceAmount")]
        public Input<int>? InstanceAmount { get; set; }

        /// <summary>
        /// Instance type. Currently, you can only set the capacity reservation service for one instance type.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The type of private resource pool generated after the capacity reservation service takes effect. Value range:Open: Open mode.Target: dedicated mode.Default value: Open
        /// </summary>
        [Input("matchCriteria")]
        public Input<string>? MatchCriteria { get; set; }

        /// <summary>
        /// The payment type of the resource
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// platform of the capacity reservation, value range `windows`, `linux`.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// The resource group id.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// time of the capacity reservation which become active.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// The capacity is scheduled to take effect. Possible values:-Now: Effective immediately.-Later: the specified time takes effect.
        /// </summary>
        [Input("startTimeType")]
        public Input<string>? StartTimeType { get; set; }

        /// <summary>
        /// The status of the capacity reservation.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// This parameter is under test and is not yet open for use.
        /// </summary>
        [Input("timeSlot")]
        public Input<string>? TimeSlot { get; set; }

        [Input("zoneIds")]
        private InputList<string>? _zoneIds;

        /// <summary>
        /// The ID of the zone in the region to which the capacity reservation service belongs. Currently, it is only supported to create a capacity reservation service in one zone.
        /// </summary>
        public InputList<string> ZoneIds
        {
            get => _zoneIds ?? (_zoneIds = new InputList<string>());
            set => _zoneIds = value;
        }

        public CapacityReservationState()
        {
        }
        public static new CapacityReservationState Empty => new CapacityReservationState();
    }
}
