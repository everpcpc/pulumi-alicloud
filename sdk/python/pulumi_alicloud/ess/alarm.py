# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Alarm(pulumi.CustomResource):
    alarm_actions: pulumi.Output[list]
    """
    The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.
    """
    cloud_monitor_group_id: pulumi.Output[float]
    """
    Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.
    """
    comparison_operator: pulumi.Output[str]
    """
    The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Supported value: >=, <=, >, <. Defaults to >=.
    """
    description: pulumi.Output[str]
    """
    The description for the alarm.
    """
    dimensions: pulumi.Output[dict]
    """
    The dimension map for the alarm's associated metric (documented below). For all metrics, you can not set the dimension key as "scaling_group" or "userId", which is set by default, the second dimension for metric, such as "device" for "PackagesNetIn", need to be set by users.
    """
    enable: pulumi.Output[bool]
    """
    Whether to enable specific ess alarm. Default to true.
    """
    evaluation_count: pulumi.Output[float]
    """
    The number of times that needs to satisfies comparison condition before transition into ALARM state. Defaults to 3.
    """
    metric_name: pulumi.Output[str]
    """
    The name for the alarm's associated metric. See Block_metricNames_and_dimensions below for details.
    """
    metric_type: pulumi.Output[str]
    """
    The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system. 
    """
    name: pulumi.Output[str]
    """
    The name for ess alarm.
    """
    period: pulumi.Output[float]
    """
    The period in seconds over which the specified statistic is applied. Supported value: 60, 120, 300, 900. Defaults to 300.
    """
    scaling_group_id: pulumi.Output[str]
    """
    The scaling group associated with this alarm, the 'ForceNew' attribute is available in 1.56.0+.
    """
    state: pulumi.Output[str]
    """
    The state of specified alarm.  
    """
    statistics: pulumi.Output[str]
    """
    The statistic to apply to the alarm's associated metric. Supported value: Average, Minimum, Maximum. Defaults to Average.
    """
    threshold: pulumi.Output[str]
    """
    The value against which the specified statistics is compared.
    """
    def __init__(__self__, resource_name, opts=None, alarm_actions=None, cloud_monitor_group_id=None, comparison_operator=None, description=None, dimensions=None, enable=None, evaluation_count=None, metric_name=None, metric_type=None, name=None, period=None, scaling_group_id=None, statistics=None, threshold=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Alarm resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] alarm_actions: The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.
        :param pulumi.Input[float] cloud_monitor_group_id: Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.
        :param pulumi.Input[str] comparison_operator: The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Supported value: >=, <=, >, <. Defaults to >=.
        :param pulumi.Input[str] description: The description for the alarm.
        :param pulumi.Input[dict] dimensions: The dimension map for the alarm's associated metric (documented below). For all metrics, you can not set the dimension key as "scaling_group" or "userId", which is set by default, the second dimension for metric, such as "device" for "PackagesNetIn", need to be set by users.
        :param pulumi.Input[bool] enable: Whether to enable specific ess alarm. Default to true.
        :param pulumi.Input[float] evaluation_count: The number of times that needs to satisfies comparison condition before transition into ALARM state. Defaults to 3.
        :param pulumi.Input[str] metric_name: The name for the alarm's associated metric. See Block_metricNames_and_dimensions below for details.
        :param pulumi.Input[str] metric_type: The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system. 
        :param pulumi.Input[str] name: The name for ess alarm.
        :param pulumi.Input[float] period: The period in seconds over which the specified statistic is applied. Supported value: 60, 120, 300, 900. Defaults to 300.
        :param pulumi.Input[str] scaling_group_id: The scaling group associated with this alarm, the 'ForceNew' attribute is available in 1.56.0+.
        :param pulumi.Input[str] statistics: The statistic to apply to the alarm's associated metric. Supported value: Average, Minimum, Maximum. Defaults to Average.
        :param pulumi.Input[str] threshold: The value against which the specified statistics is compared.
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

            if alarm_actions is None:
                raise TypeError("Missing required property 'alarm_actions'")
            __props__['alarm_actions'] = alarm_actions
            __props__['cloud_monitor_group_id'] = cloud_monitor_group_id
            __props__['comparison_operator'] = comparison_operator
            __props__['description'] = description
            __props__['dimensions'] = dimensions
            __props__['enable'] = enable
            __props__['evaluation_count'] = evaluation_count
            if metric_name is None:
                raise TypeError("Missing required property 'metric_name'")
            __props__['metric_name'] = metric_name
            __props__['metric_type'] = metric_type
            __props__['name'] = name
            __props__['period'] = period
            if scaling_group_id is None:
                raise TypeError("Missing required property 'scaling_group_id'")
            __props__['scaling_group_id'] = scaling_group_id
            __props__['statistics'] = statistics
            if threshold is None:
                raise TypeError("Missing required property 'threshold'")
            __props__['threshold'] = threshold
            __props__['state'] = None
        super(Alarm, __self__).__init__(
            'alicloud:ess/alarm:Alarm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, alarm_actions=None, cloud_monitor_group_id=None, comparison_operator=None, description=None, dimensions=None, enable=None, evaluation_count=None, metric_name=None, metric_type=None, name=None, period=None, scaling_group_id=None, state=None, statistics=None, threshold=None):
        """
        Get an existing Alarm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] alarm_actions: The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.
        :param pulumi.Input[float] cloud_monitor_group_id: Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.
        :param pulumi.Input[str] comparison_operator: The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Supported value: >=, <=, >, <. Defaults to >=.
        :param pulumi.Input[str] description: The description for the alarm.
        :param pulumi.Input[dict] dimensions: The dimension map for the alarm's associated metric (documented below). For all metrics, you can not set the dimension key as "scaling_group" or "userId", which is set by default, the second dimension for metric, such as "device" for "PackagesNetIn", need to be set by users.
        :param pulumi.Input[bool] enable: Whether to enable specific ess alarm. Default to true.
        :param pulumi.Input[float] evaluation_count: The number of times that needs to satisfies comparison condition before transition into ALARM state. Defaults to 3.
        :param pulumi.Input[str] metric_name: The name for the alarm's associated metric. See Block_metricNames_and_dimensions below for details.
        :param pulumi.Input[str] metric_type: The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system. 
        :param pulumi.Input[str] name: The name for ess alarm.
        :param pulumi.Input[float] period: The period in seconds over which the specified statistic is applied. Supported value: 60, 120, 300, 900. Defaults to 300.
        :param pulumi.Input[str] scaling_group_id: The scaling group associated with this alarm, the 'ForceNew' attribute is available in 1.56.0+.
        :param pulumi.Input[str] state: The state of specified alarm.  
        :param pulumi.Input[str] statistics: The statistic to apply to the alarm's associated metric. Supported value: Average, Minimum, Maximum. Defaults to Average.
        :param pulumi.Input[str] threshold: The value against which the specified statistics is compared.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["alarm_actions"] = alarm_actions
        __props__["cloud_monitor_group_id"] = cloud_monitor_group_id
        __props__["comparison_operator"] = comparison_operator
        __props__["description"] = description
        __props__["dimensions"] = dimensions
        __props__["enable"] = enable
        __props__["evaluation_count"] = evaluation_count
        __props__["metric_name"] = metric_name
        __props__["metric_type"] = metric_type
        __props__["name"] = name
        __props__["period"] = period
        __props__["scaling_group_id"] = scaling_group_id
        __props__["state"] = state
        __props__["statistics"] = statistics
        __props__["threshold"] = threshold
        return Alarm(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

