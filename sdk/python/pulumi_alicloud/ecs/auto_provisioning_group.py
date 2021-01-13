# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['AutoProvisioningGroup']


class AutoProvisioningGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_provisioning_group_name: Optional[pulumi.Input[str]] = None,
                 auto_provisioning_group_type: Optional[pulumi.Input[str]] = None,
                 default_target_capacity_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 excess_capacity_termination_policy: Optional[pulumi.Input[str]] = None,
                 launch_template_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoProvisioningGroupLaunchTemplateConfigArgs']]]]] = None,
                 launch_template_id: Optional[pulumi.Input[str]] = None,
                 launch_template_version: Optional[pulumi.Input[str]] = None,
                 max_spot_price: Optional[pulumi.Input[float]] = None,
                 pay_as_you_go_allocation_strategy: Optional[pulumi.Input[str]] = None,
                 pay_as_you_go_target_capacity: Optional[pulumi.Input[str]] = None,
                 spot_allocation_strategy: Optional[pulumi.Input[str]] = None,
                 spot_instance_interruption_behavior: Optional[pulumi.Input[str]] = None,
                 spot_instance_pools_to_use_count: Optional[pulumi.Input[int]] = None,
                 spot_target_capacity: Optional[pulumi.Input[str]] = None,
                 terminate_instances: Optional[pulumi.Input[bool]] = None,
                 terminate_instances_with_expiration: Optional[pulumi.Input[bool]] = None,
                 total_target_capacity: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 valid_until: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a ECS auto provisioning group resource which is a solution that uses preemptive instances and pay_as_you_go instances to rapidly deploy clusters.

        > **NOTE:** Available in 1.79.0+

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "auto_provisioning_group"
        default_zones = alicloud.get_zones(available_disk_category="cloud_efficiency",
            available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("defaultNetwork", cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/24",
            availability_zone=default_zones.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup", vpc_id=default_network.id)
        default_images = alicloud.ecs.get_images(name_regex="^ubuntu_18.*64",
            most_recent=True,
            owners="system")
        template = alicloud.ecs.LaunchTemplate("template",
            image_id=default_images.images[0].id,
            instance_type="ecs.n1.tiny",
            security_group_id=default_security_group.id)
        default_auto_provisioning_group = alicloud.ecs.AutoProvisioningGroup("defaultAutoProvisioningGroup",
            launch_template_id=template.id,
            total_target_capacity="4",
            pay_as_you_go_target_capacity="1",
            spot_target_capacity="2",
            launch_template_configs=[alicloud.ecs.AutoProvisioningGroupLaunchTemplateConfigArgs(
                instance_type="ecs.n1.small",
                vswitch_id=default_switch.id,
                weighted_capacity="2",
                max_price="2",
            )])
        ```
        ## Block config

        The config mapping supports the following:
        * `instance_type` - (Optional) The instance type of the Nth extended configurations of the launch template.
        * `max_price` - (Required) The maximum price of the instance type specified in the Nth extended configurations of the launch template.
        * `vswitch_id` - (Required) The ID of the VSwitch in the Nth extended configurations of the launch template.
        * `weighted_capacity` - (Required) The weight of the instance type specified in the Nth extended configurations of the launch template.
        * `priority` - (Optional) The priority of the instance type specified in the Nth extended configurations of the launch template. A value of 0 indicates the highest priority.

        ## Import

        ECS auto provisioning group can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ecs/autoProvisioningGroup:AutoProvisioningGroup example asg-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_provisioning_group_name: The name of the auto provisioning group to be created. It must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-)
        :param pulumi.Input[str] auto_provisioning_group_type: The type of the auto provisioning group. Valid values:`request` and `maintain`,Default value: `maintain`.
        :param pulumi.Input[str] default_target_capacity_type: The type of supplemental instances. When the total value of `PayAsYouGoTargetCapacity` and `SpotTargetCapacity` is smaller than the value of TotalTargetCapacity, the auto provisioning group will create instances of the specified type to meet the capacity requirements. Valid values:`PayAsYouGo`: Pay-as-you-go instances; `Spot`: Preemptible instances, Default value: `Spot`.
        :param pulumi.Input[str] description: The description of the auto provisioning group.
        :param pulumi.Input[str] excess_capacity_termination_policy: The shutdown policy for excess preemptible instances followed when the capacity of the auto provisioning group exceeds the target capacity. Valid values: `no-termination` and `termination`,Default value: `no-termination`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoProvisioningGroupLaunchTemplateConfigArgs']]]] launch_template_configs: DataDisk mappings to attach to ecs instance. See Block config below for details.
        :param pulumi.Input[str] launch_template_id: The ID of the instance launch template associated with the auto provisioning group.
        :param pulumi.Input[str] launch_template_version: The version of the instance launch template associated with the auto provisioning group.
        :param pulumi.Input[float] max_spot_price: The global maximum price for preemptible instances in the auto provisioning group. If both the `MaxSpotPrice` and `LaunchTemplateConfig.N.MaxPrice` parameters are specified, the maximum price is the lower value of the two.
        :param pulumi.Input[str] pay_as_you_go_allocation_strategy: The scale-out policy for pay-as-you-go instances. Valid values: `lowest-price` and `prioritized`,Default value: `lowest-price`.
        :param pulumi.Input[str] pay_as_you_go_target_capacity: The target capacity of pay-as-you-go instances in the auto provisioning group.
        :param pulumi.Input[str] spot_allocation_strategy: The scale-out policy for preemptible instances. Valid values:`lowest-price` and `diversified`,Default value: `lowest-price`.
        :param pulumi.Input[str] spot_instance_interruption_behavior: The default behavior after preemptible instances are shut down. Value values: `stop` and `terminate`,Default value: `stop`.
        :param pulumi.Input[int] spot_instance_pools_to_use_count: This parameter takes effect when the `SpotAllocationStrategy` parameter is set to `lowest-price`. The auto provisioning group selects instance types of the lowest cost to create instances.
        :param pulumi.Input[str] spot_target_capacity: The target capacity of preemptible instances in the auto provisioning group.
        :param pulumi.Input[bool] terminate_instances: Specifies whether to release instances of the auto provisioning group. Valid values:`false` and `true`, default value: `false`.
        :param pulumi.Input[bool] terminate_instances_with_expiration: The shutdown policy for preemptible instances when the auto provisioning group expires. Valid values: `false` and `true`, default value: `false`.
        :param pulumi.Input[str] total_target_capacity: The total target capacity of the auto provisioning group. The target capacity consists of the following three parts:PayAsYouGoTargetCapacity,SpotTargetCapacity and the supplemental capacity besides PayAsYouGoTargetCapacity and SpotTargetCapacity.
        :param pulumi.Input[str] valid_from: The time when the auto provisioning group is started. The period of time between this point in time and the point in time specified by the `valid_until` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group is immediately started after creation.
        :param pulumi.Input[str] valid_until: The time when the auto provisioning group expires. The period of time between this point in time and the point in time specified by the `valid_from` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group never expires.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['auto_provisioning_group_name'] = auto_provisioning_group_name
            __props__['auto_provisioning_group_type'] = auto_provisioning_group_type
            __props__['default_target_capacity_type'] = default_target_capacity_type
            __props__['description'] = description
            __props__['excess_capacity_termination_policy'] = excess_capacity_termination_policy
            if launch_template_configs is None and not opts.urn:
                raise TypeError("Missing required property 'launch_template_configs'")
            __props__['launch_template_configs'] = launch_template_configs
            if launch_template_id is None and not opts.urn:
                raise TypeError("Missing required property 'launch_template_id'")
            __props__['launch_template_id'] = launch_template_id
            __props__['launch_template_version'] = launch_template_version
            __props__['max_spot_price'] = max_spot_price
            __props__['pay_as_you_go_allocation_strategy'] = pay_as_you_go_allocation_strategy
            __props__['pay_as_you_go_target_capacity'] = pay_as_you_go_target_capacity
            __props__['spot_allocation_strategy'] = spot_allocation_strategy
            __props__['spot_instance_interruption_behavior'] = spot_instance_interruption_behavior
            __props__['spot_instance_pools_to_use_count'] = spot_instance_pools_to_use_count
            __props__['spot_target_capacity'] = spot_target_capacity
            __props__['terminate_instances'] = terminate_instances
            __props__['terminate_instances_with_expiration'] = terminate_instances_with_expiration
            if total_target_capacity is None and not opts.urn:
                raise TypeError("Missing required property 'total_target_capacity'")
            __props__['total_target_capacity'] = total_target_capacity
            __props__['valid_from'] = valid_from
            __props__['valid_until'] = valid_until
        super(AutoProvisioningGroup, __self__).__init__(
            'alicloud:ecs/autoProvisioningGroup:AutoProvisioningGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_provisioning_group_name: Optional[pulumi.Input[str]] = None,
            auto_provisioning_group_type: Optional[pulumi.Input[str]] = None,
            default_target_capacity_type: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            excess_capacity_termination_policy: Optional[pulumi.Input[str]] = None,
            launch_template_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoProvisioningGroupLaunchTemplateConfigArgs']]]]] = None,
            launch_template_id: Optional[pulumi.Input[str]] = None,
            launch_template_version: Optional[pulumi.Input[str]] = None,
            max_spot_price: Optional[pulumi.Input[float]] = None,
            pay_as_you_go_allocation_strategy: Optional[pulumi.Input[str]] = None,
            pay_as_you_go_target_capacity: Optional[pulumi.Input[str]] = None,
            spot_allocation_strategy: Optional[pulumi.Input[str]] = None,
            spot_instance_interruption_behavior: Optional[pulumi.Input[str]] = None,
            spot_instance_pools_to_use_count: Optional[pulumi.Input[int]] = None,
            spot_target_capacity: Optional[pulumi.Input[str]] = None,
            terminate_instances: Optional[pulumi.Input[bool]] = None,
            terminate_instances_with_expiration: Optional[pulumi.Input[bool]] = None,
            total_target_capacity: Optional[pulumi.Input[str]] = None,
            valid_from: Optional[pulumi.Input[str]] = None,
            valid_until: Optional[pulumi.Input[str]] = None) -> 'AutoProvisioningGroup':
        """
        Get an existing AutoProvisioningGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_provisioning_group_name: The name of the auto provisioning group to be created. It must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-)
        :param pulumi.Input[str] auto_provisioning_group_type: The type of the auto provisioning group. Valid values:`request` and `maintain`,Default value: `maintain`.
        :param pulumi.Input[str] default_target_capacity_type: The type of supplemental instances. When the total value of `PayAsYouGoTargetCapacity` and `SpotTargetCapacity` is smaller than the value of TotalTargetCapacity, the auto provisioning group will create instances of the specified type to meet the capacity requirements. Valid values:`PayAsYouGo`: Pay-as-you-go instances; `Spot`: Preemptible instances, Default value: `Spot`.
        :param pulumi.Input[str] description: The description of the auto provisioning group.
        :param pulumi.Input[str] excess_capacity_termination_policy: The shutdown policy for excess preemptible instances followed when the capacity of the auto provisioning group exceeds the target capacity. Valid values: `no-termination` and `termination`,Default value: `no-termination`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoProvisioningGroupLaunchTemplateConfigArgs']]]] launch_template_configs: DataDisk mappings to attach to ecs instance. See Block config below for details.
        :param pulumi.Input[str] launch_template_id: The ID of the instance launch template associated with the auto provisioning group.
        :param pulumi.Input[str] launch_template_version: The version of the instance launch template associated with the auto provisioning group.
        :param pulumi.Input[float] max_spot_price: The global maximum price for preemptible instances in the auto provisioning group. If both the `MaxSpotPrice` and `LaunchTemplateConfig.N.MaxPrice` parameters are specified, the maximum price is the lower value of the two.
        :param pulumi.Input[str] pay_as_you_go_allocation_strategy: The scale-out policy for pay-as-you-go instances. Valid values: `lowest-price` and `prioritized`,Default value: `lowest-price`.
        :param pulumi.Input[str] pay_as_you_go_target_capacity: The target capacity of pay-as-you-go instances in the auto provisioning group.
        :param pulumi.Input[str] spot_allocation_strategy: The scale-out policy for preemptible instances. Valid values:`lowest-price` and `diversified`,Default value: `lowest-price`.
        :param pulumi.Input[str] spot_instance_interruption_behavior: The default behavior after preemptible instances are shut down. Value values: `stop` and `terminate`,Default value: `stop`.
        :param pulumi.Input[int] spot_instance_pools_to_use_count: This parameter takes effect when the `SpotAllocationStrategy` parameter is set to `lowest-price`. The auto provisioning group selects instance types of the lowest cost to create instances.
        :param pulumi.Input[str] spot_target_capacity: The target capacity of preemptible instances in the auto provisioning group.
        :param pulumi.Input[bool] terminate_instances: Specifies whether to release instances of the auto provisioning group. Valid values:`false` and `true`, default value: `false`.
        :param pulumi.Input[bool] terminate_instances_with_expiration: The shutdown policy for preemptible instances when the auto provisioning group expires. Valid values: `false` and `true`, default value: `false`.
        :param pulumi.Input[str] total_target_capacity: The total target capacity of the auto provisioning group. The target capacity consists of the following three parts:PayAsYouGoTargetCapacity,SpotTargetCapacity and the supplemental capacity besides PayAsYouGoTargetCapacity and SpotTargetCapacity.
        :param pulumi.Input[str] valid_from: The time when the auto provisioning group is started. The period of time between this point in time and the point in time specified by the `valid_until` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group is immediately started after creation.
        :param pulumi.Input[str] valid_until: The time when the auto provisioning group expires. The period of time between this point in time and the point in time specified by the `valid_from` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group never expires.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_provisioning_group_name"] = auto_provisioning_group_name
        __props__["auto_provisioning_group_type"] = auto_provisioning_group_type
        __props__["default_target_capacity_type"] = default_target_capacity_type
        __props__["description"] = description
        __props__["excess_capacity_termination_policy"] = excess_capacity_termination_policy
        __props__["launch_template_configs"] = launch_template_configs
        __props__["launch_template_id"] = launch_template_id
        __props__["launch_template_version"] = launch_template_version
        __props__["max_spot_price"] = max_spot_price
        __props__["pay_as_you_go_allocation_strategy"] = pay_as_you_go_allocation_strategy
        __props__["pay_as_you_go_target_capacity"] = pay_as_you_go_target_capacity
        __props__["spot_allocation_strategy"] = spot_allocation_strategy
        __props__["spot_instance_interruption_behavior"] = spot_instance_interruption_behavior
        __props__["spot_instance_pools_to_use_count"] = spot_instance_pools_to_use_count
        __props__["spot_target_capacity"] = spot_target_capacity
        __props__["terminate_instances"] = terminate_instances
        __props__["terminate_instances_with_expiration"] = terminate_instances_with_expiration
        __props__["total_target_capacity"] = total_target_capacity
        __props__["valid_from"] = valid_from
        __props__["valid_until"] = valid_until
        return AutoProvisioningGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoProvisioningGroupName")
    def auto_provisioning_group_name(self) -> pulumi.Output[str]:
        """
        The name of the auto provisioning group to be created. It must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-)
        """
        return pulumi.get(self, "auto_provisioning_group_name")

    @property
    @pulumi.getter(name="autoProvisioningGroupType")
    def auto_provisioning_group_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the auto provisioning group. Valid values:`request` and `maintain`,Default value: `maintain`.
        """
        return pulumi.get(self, "auto_provisioning_group_type")

    @property
    @pulumi.getter(name="defaultTargetCapacityType")
    def default_target_capacity_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of supplemental instances. When the total value of `PayAsYouGoTargetCapacity` and `SpotTargetCapacity` is smaller than the value of TotalTargetCapacity, the auto provisioning group will create instances of the specified type to meet the capacity requirements. Valid values:`PayAsYouGo`: Pay-as-you-go instances; `Spot`: Preemptible instances, Default value: `Spot`.
        """
        return pulumi.get(self, "default_target_capacity_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the auto provisioning group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="excessCapacityTerminationPolicy")
    def excess_capacity_termination_policy(self) -> pulumi.Output[Optional[str]]:
        """
        The shutdown policy for excess preemptible instances followed when the capacity of the auto provisioning group exceeds the target capacity. Valid values: `no-termination` and `termination`,Default value: `no-termination`.
        """
        return pulumi.get(self, "excess_capacity_termination_policy")

    @property
    @pulumi.getter(name="launchTemplateConfigs")
    def launch_template_configs(self) -> pulumi.Output[Sequence['outputs.AutoProvisioningGroupLaunchTemplateConfig']]:
        """
        DataDisk mappings to attach to ecs instance. See Block config below for details.
        """
        return pulumi.get(self, "launch_template_configs")

    @property
    @pulumi.getter(name="launchTemplateId")
    def launch_template_id(self) -> pulumi.Output[str]:
        """
        The ID of the instance launch template associated with the auto provisioning group.
        """
        return pulumi.get(self, "launch_template_id")

    @property
    @pulumi.getter(name="launchTemplateVersion")
    def launch_template_version(self) -> pulumi.Output[str]:
        """
        The version of the instance launch template associated with the auto provisioning group.
        """
        return pulumi.get(self, "launch_template_version")

    @property
    @pulumi.getter(name="maxSpotPrice")
    def max_spot_price(self) -> pulumi.Output[float]:
        """
        The global maximum price for preemptible instances in the auto provisioning group. If both the `MaxSpotPrice` and `LaunchTemplateConfig.N.MaxPrice` parameters are specified, the maximum price is the lower value of the two.
        """
        return pulumi.get(self, "max_spot_price")

    @property
    @pulumi.getter(name="payAsYouGoAllocationStrategy")
    def pay_as_you_go_allocation_strategy(self) -> pulumi.Output[Optional[str]]:
        """
        The scale-out policy for pay-as-you-go instances. Valid values: `lowest-price` and `prioritized`,Default value: `lowest-price`.
        """
        return pulumi.get(self, "pay_as_you_go_allocation_strategy")

    @property
    @pulumi.getter(name="payAsYouGoTargetCapacity")
    def pay_as_you_go_target_capacity(self) -> pulumi.Output[Optional[str]]:
        """
        The target capacity of pay-as-you-go instances in the auto provisioning group.
        """
        return pulumi.get(self, "pay_as_you_go_target_capacity")

    @property
    @pulumi.getter(name="spotAllocationStrategy")
    def spot_allocation_strategy(self) -> pulumi.Output[Optional[str]]:
        """
        The scale-out policy for preemptible instances. Valid values:`lowest-price` and `diversified`,Default value: `lowest-price`.
        """
        return pulumi.get(self, "spot_allocation_strategy")

    @property
    @pulumi.getter(name="spotInstanceInterruptionBehavior")
    def spot_instance_interruption_behavior(self) -> pulumi.Output[Optional[str]]:
        """
        The default behavior after preemptible instances are shut down. Value values: `stop` and `terminate`,Default value: `stop`.
        """
        return pulumi.get(self, "spot_instance_interruption_behavior")

    @property
    @pulumi.getter(name="spotInstancePoolsToUseCount")
    def spot_instance_pools_to_use_count(self) -> pulumi.Output[int]:
        """
        This parameter takes effect when the `SpotAllocationStrategy` parameter is set to `lowest-price`. The auto provisioning group selects instance types of the lowest cost to create instances.
        """
        return pulumi.get(self, "spot_instance_pools_to_use_count")

    @property
    @pulumi.getter(name="spotTargetCapacity")
    def spot_target_capacity(self) -> pulumi.Output[Optional[str]]:
        """
        The target capacity of preemptible instances in the auto provisioning group.
        """
        return pulumi.get(self, "spot_target_capacity")

    @property
    @pulumi.getter(name="terminateInstances")
    def terminate_instances(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to release instances of the auto provisioning group. Valid values:`false` and `true`, default value: `false`.
        """
        return pulumi.get(self, "terminate_instances")

    @property
    @pulumi.getter(name="terminateInstancesWithExpiration")
    def terminate_instances_with_expiration(self) -> pulumi.Output[Optional[bool]]:
        """
        The shutdown policy for preemptible instances when the auto provisioning group expires. Valid values: `false` and `true`, default value: `false`.
        """
        return pulumi.get(self, "terminate_instances_with_expiration")

    @property
    @pulumi.getter(name="totalTargetCapacity")
    def total_target_capacity(self) -> pulumi.Output[str]:
        """
        The total target capacity of the auto provisioning group. The target capacity consists of the following three parts:PayAsYouGoTargetCapacity,SpotTargetCapacity and the supplemental capacity besides PayAsYouGoTargetCapacity and SpotTargetCapacity.
        """
        return pulumi.get(self, "total_target_capacity")

    @property
    @pulumi.getter(name="validFrom")
    def valid_from(self) -> pulumi.Output[str]:
        """
        The time when the auto provisioning group is started. The period of time between this point in time and the point in time specified by the `valid_until` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group is immediately started after creation.
        """
        return pulumi.get(self, "valid_from")

    @property
    @pulumi.getter(name="validUntil")
    def valid_until(self) -> pulumi.Output[str]:
        """
        The time when the auto provisioning group expires. The period of time between this point in time and the point in time specified by the `valid_from` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group never expires.
        """
        return pulumi.get(self, "valid_until")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

