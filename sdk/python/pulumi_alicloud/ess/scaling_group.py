# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ScalingGroup(pulumi.CustomResource):
    db_instance_ids: pulumi.Output[list]
    """
    If an RDS instance is specified in the scaling group, the scaling group automatically attaches the Intranet IP addresses of its ECS instances to the RDS access whitelist.
    - The specified RDS instance must be in running status.
    - The specified RDS instance’s whitelist must have room for more IP addresses.
    """
    default_cooldown: pulumi.Output[float]
    """
    Default cool-down time (in seconds) of the scaling group. Value range: [0, 86400]. The default value is 300s.
    """
    desired_capacity: pulumi.Output[float]
    """
    Expected number of ECS instances in the scaling group. Value range: [min_size, max_size].
    """
    loadbalancer_ids: pulumi.Output[list]
    """
    If a Server Load Balancer instance is specified in the scaling group, the scaling group automatically attaches its ECS instances to the Server Load Balancer instance.
    - The Server Load Balancer instance must be enabled.
    - At least one listener must be configured for each Server Load Balancer and it HealthCheck must be on. Otherwise, creation will fail (it may be useful to add a `depends_on` argument
    targeting your `slb.Listener` in order to make sure the listener with its HealthCheck configuration is ready before creating your scaling group).
    - The Server Load Balancer instance attached with VPC-type ECS instances cannot be attached to the scaling group.
    - The default weight of an ECS instance attached to the Server Load Balancer instance is 50.
    """
    max_size: pulumi.Output[float]
    """
    Maximum number of ECS instances in the scaling group. Value range: [0, 1000].
    """
    min_size: pulumi.Output[float]
    """
    Minimum number of ECS instances in the scaling group. Value range: [0, 1000].
    """
    multi_az_policy: pulumi.Output[str]
    """
    Multi-AZ scaling group ECS instance expansion and contraction strategy. PRIORITY, BALANCE or COST_OPTIMIZED(Available in 1.54.0+).
    """
    on_demand_base_capacity: pulumi.Output[float]
    """
    The minimum amount of the Auto Scaling group's capacity that must be fulfilled by On-Demand Instances. This base portion is provisioned first as your group scales.
    """
    on_demand_percentage_above_base_capacity: pulumi.Output[float]
    """
    Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond OnDemandBaseCapacity.  
    """
    removal_policies: pulumi.Output[list]
    """
    RemovalPolicy is used to select the ECS instances you want to remove from the scaling group when multiple candidates for removal exist. Optional values:
    - OldestInstance: removes the ECS instance that is added to the scaling group at the earliest point in time.
    - NewestInstance: removes the ECS instance that is added to the scaling group at the latest point in time.
    - OldestScalingConfiguration: removes the ECS instance that is created based on the earliest scaling configuration.
    - Default values: Default value of RemovalPolicy.1: OldestScalingConfiguration. Default value of RemovalPolicy.2: OldestInstance.
    """
    scaling_group_name: pulumi.Output[str]
    """
    Name shown for the scaling group, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain numbers, underscores `_`, hyphens `-`, and decimal points `.`. If this parameter is not specified, the default value is ScalingGroupId.
    """
    spot_instance_pools: pulumi.Output[float]
    """
    The number of Spot pools to use to allocate your Spot capacity. The Spot pools is composed of instance types of lowest price.
    """
    spot_instance_remedy: pulumi.Output[bool]
    """
    Whether to replace spot instances with newly created spot/onDemand instance when receive a spot recycling message.            
    """
    vswitch_id: pulumi.Output[str]
    """
    It has been deprecated from version 1.7.1 and new field 'vswitch_ids' replaces it.
    """
    vswitch_ids: pulumi.Output[list]
    """
    List of virtual switch IDs in which the ecs instances to be launched.
    """
    def __init__(__self__, resource_name, opts=None, db_instance_ids=None, default_cooldown=None, desired_capacity=None, loadbalancer_ids=None, max_size=None, min_size=None, multi_az_policy=None, on_demand_base_capacity=None, on_demand_percentage_above_base_capacity=None, removal_policies=None, scaling_group_name=None, spot_instance_pools=None, spot_instance_remedy=None, vswitch_id=None, vswitch_ids=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a ScalingGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] db_instance_ids: If an RDS instance is specified in the scaling group, the scaling group automatically attaches the Intranet IP addresses of its ECS instances to the RDS access whitelist.
               - The specified RDS instance must be in running status.
               - The specified RDS instance’s whitelist must have room for more IP addresses.
        :param pulumi.Input[float] default_cooldown: Default cool-down time (in seconds) of the scaling group. Value range: [0, 86400]. The default value is 300s.
        :param pulumi.Input[float] desired_capacity: Expected number of ECS instances in the scaling group. Value range: [min_size, max_size].
        :param pulumi.Input[list] loadbalancer_ids: If a Server Load Balancer instance is specified in the scaling group, the scaling group automatically attaches its ECS instances to the Server Load Balancer instance.
               - The Server Load Balancer instance must be enabled.
               - At least one listener must be configured for each Server Load Balancer and it HealthCheck must be on. Otherwise, creation will fail (it may be useful to add a `depends_on` argument
               targeting your `slb.Listener` in order to make sure the listener with its HealthCheck configuration is ready before creating your scaling group).
               - The Server Load Balancer instance attached with VPC-type ECS instances cannot be attached to the scaling group.
               - The default weight of an ECS instance attached to the Server Load Balancer instance is 50.
        :param pulumi.Input[float] max_size: Maximum number of ECS instances in the scaling group. Value range: [0, 1000].
        :param pulumi.Input[float] min_size: Minimum number of ECS instances in the scaling group. Value range: [0, 1000].
        :param pulumi.Input[str] multi_az_policy: Multi-AZ scaling group ECS instance expansion and contraction strategy. PRIORITY, BALANCE or COST_OPTIMIZED(Available in 1.54.0+).
        :param pulumi.Input[float] on_demand_base_capacity: The minimum amount of the Auto Scaling group's capacity that must be fulfilled by On-Demand Instances. This base portion is provisioned first as your group scales.
        :param pulumi.Input[float] on_demand_percentage_above_base_capacity: Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond OnDemandBaseCapacity.  
        :param pulumi.Input[list] removal_policies: RemovalPolicy is used to select the ECS instances you want to remove from the scaling group when multiple candidates for removal exist. Optional values:
               - OldestInstance: removes the ECS instance that is added to the scaling group at the earliest point in time.
               - NewestInstance: removes the ECS instance that is added to the scaling group at the latest point in time.
               - OldestScalingConfiguration: removes the ECS instance that is created based on the earliest scaling configuration.
               - Default values: Default value of RemovalPolicy.1: OldestScalingConfiguration. Default value of RemovalPolicy.2: OldestInstance.
        :param pulumi.Input[str] scaling_group_name: Name shown for the scaling group, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain numbers, underscores `_`, hyphens `-`, and decimal points `.`. If this parameter is not specified, the default value is ScalingGroupId.
        :param pulumi.Input[float] spot_instance_pools: The number of Spot pools to use to allocate your Spot capacity. The Spot pools is composed of instance types of lowest price.
        :param pulumi.Input[bool] spot_instance_remedy: Whether to replace spot instances with newly created spot/onDemand instance when receive a spot recycling message.            
        :param pulumi.Input[str] vswitch_id: It has been deprecated from version 1.7.1 and new field 'vswitch_ids' replaces it.
        :param pulumi.Input[list] vswitch_ids: List of virtual switch IDs in which the ecs instances to be launched.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['db_instance_ids'] = db_instance_ids
            __props__['default_cooldown'] = default_cooldown
            __props__['desired_capacity'] = desired_capacity
            __props__['loadbalancer_ids'] = loadbalancer_ids
            if max_size is None:
                raise TypeError("Missing required property 'max_size'")
            __props__['max_size'] = max_size
            if min_size is None:
                raise TypeError("Missing required property 'min_size'")
            __props__['min_size'] = min_size
            __props__['multi_az_policy'] = multi_az_policy
            __props__['on_demand_base_capacity'] = on_demand_base_capacity
            __props__['on_demand_percentage_above_base_capacity'] = on_demand_percentage_above_base_capacity
            __props__['removal_policies'] = removal_policies
            __props__['scaling_group_name'] = scaling_group_name
            __props__['spot_instance_pools'] = spot_instance_pools
            __props__['spot_instance_remedy'] = spot_instance_remedy
            __props__['vswitch_id'] = vswitch_id
            __props__['vswitch_ids'] = vswitch_ids
        super(ScalingGroup, __self__).__init__(
            'alicloud:ess/scalingGroup:ScalingGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, db_instance_ids=None, default_cooldown=None, desired_capacity=None, loadbalancer_ids=None, max_size=None, min_size=None, multi_az_policy=None, on_demand_base_capacity=None, on_demand_percentage_above_base_capacity=None, removal_policies=None, scaling_group_name=None, spot_instance_pools=None, spot_instance_remedy=None, vswitch_id=None, vswitch_ids=None):
        """
        Get an existing ScalingGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] db_instance_ids: If an RDS instance is specified in the scaling group, the scaling group automatically attaches the Intranet IP addresses of its ECS instances to the RDS access whitelist.
               - The specified RDS instance must be in running status.
               - The specified RDS instance’s whitelist must have room for more IP addresses.
        :param pulumi.Input[float] default_cooldown: Default cool-down time (in seconds) of the scaling group. Value range: [0, 86400]. The default value is 300s.
        :param pulumi.Input[float] desired_capacity: Expected number of ECS instances in the scaling group. Value range: [min_size, max_size].
        :param pulumi.Input[list] loadbalancer_ids: If a Server Load Balancer instance is specified in the scaling group, the scaling group automatically attaches its ECS instances to the Server Load Balancer instance.
               - The Server Load Balancer instance must be enabled.
               - At least one listener must be configured for each Server Load Balancer and it HealthCheck must be on. Otherwise, creation will fail (it may be useful to add a `depends_on` argument
               targeting your `slb.Listener` in order to make sure the listener with its HealthCheck configuration is ready before creating your scaling group).
               - The Server Load Balancer instance attached with VPC-type ECS instances cannot be attached to the scaling group.
               - The default weight of an ECS instance attached to the Server Load Balancer instance is 50.
        :param pulumi.Input[float] max_size: Maximum number of ECS instances in the scaling group. Value range: [0, 1000].
        :param pulumi.Input[float] min_size: Minimum number of ECS instances in the scaling group. Value range: [0, 1000].
        :param pulumi.Input[str] multi_az_policy: Multi-AZ scaling group ECS instance expansion and contraction strategy. PRIORITY, BALANCE or COST_OPTIMIZED(Available in 1.54.0+).
        :param pulumi.Input[float] on_demand_base_capacity: The minimum amount of the Auto Scaling group's capacity that must be fulfilled by On-Demand Instances. This base portion is provisioned first as your group scales.
        :param pulumi.Input[float] on_demand_percentage_above_base_capacity: Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond OnDemandBaseCapacity.  
        :param pulumi.Input[list] removal_policies: RemovalPolicy is used to select the ECS instances you want to remove from the scaling group when multiple candidates for removal exist. Optional values:
               - OldestInstance: removes the ECS instance that is added to the scaling group at the earliest point in time.
               - NewestInstance: removes the ECS instance that is added to the scaling group at the latest point in time.
               - OldestScalingConfiguration: removes the ECS instance that is created based on the earliest scaling configuration.
               - Default values: Default value of RemovalPolicy.1: OldestScalingConfiguration. Default value of RemovalPolicy.2: OldestInstance.
        :param pulumi.Input[str] scaling_group_name: Name shown for the scaling group, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain numbers, underscores `_`, hyphens `-`, and decimal points `.`. If this parameter is not specified, the default value is ScalingGroupId.
        :param pulumi.Input[float] spot_instance_pools: The number of Spot pools to use to allocate your Spot capacity. The Spot pools is composed of instance types of lowest price.
        :param pulumi.Input[bool] spot_instance_remedy: Whether to replace spot instances with newly created spot/onDemand instance when receive a spot recycling message.            
        :param pulumi.Input[str] vswitch_id: It has been deprecated from version 1.7.1 and new field 'vswitch_ids' replaces it.
        :param pulumi.Input[list] vswitch_ids: List of virtual switch IDs in which the ecs instances to be launched.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["db_instance_ids"] = db_instance_ids
        __props__["default_cooldown"] = default_cooldown
        __props__["desired_capacity"] = desired_capacity
        __props__["loadbalancer_ids"] = loadbalancer_ids
        __props__["max_size"] = max_size
        __props__["min_size"] = min_size
        __props__["multi_az_policy"] = multi_az_policy
        __props__["on_demand_base_capacity"] = on_demand_base_capacity
        __props__["on_demand_percentage_above_base_capacity"] = on_demand_percentage_above_base_capacity
        __props__["removal_policies"] = removal_policies
        __props__["scaling_group_name"] = scaling_group_name
        __props__["spot_instance_pools"] = spot_instance_pools
        __props__["spot_instance_remedy"] = spot_instance_remedy
        __props__["vswitch_id"] = vswitch_id
        __props__["vswitch_ids"] = vswitch_ids
        return ScalingGroup(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

