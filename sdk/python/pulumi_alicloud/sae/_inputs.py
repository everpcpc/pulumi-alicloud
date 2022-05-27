# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ApplicationScalingRuleScalingRuleMetricArgs',
    'ApplicationScalingRuleScalingRuleMetricMetricArgs',
    'ApplicationScalingRuleScalingRuleMetricScaleDownRulesArgs',
    'ApplicationScalingRuleScalingRuleMetricScaleUpRulesArgs',
    'ApplicationScalingRuleScalingRuleTimerArgs',
    'ApplicationScalingRuleScalingRuleTimerScheduleArgs',
    'GreyTagRouteDubboRuleArgs',
    'GreyTagRouteDubboRuleItemArgs',
    'GreyTagRouteScRuleArgs',
    'GreyTagRouteScRuleItemArgs',
    'IngressDefaultRuleArgs',
    'IngressRuleArgs',
    'LoadBalancerInternetInternetArgs',
    'LoadBalancerIntranetIntranetArgs',
]

@pulumi.input_type
class ApplicationScalingRuleScalingRuleMetricArgs:
    def __init__(__self__, *,
                 max_replicas: Optional[pulumi.Input[int]] = None,
                 metrics: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationScalingRuleScalingRuleMetricMetricArgs']]]] = None,
                 min_replicas: Optional[pulumi.Input[int]] = None,
                 scale_down_rules: Optional[pulumi.Input['ApplicationScalingRuleScalingRuleMetricScaleDownRulesArgs']] = None,
                 scale_up_rules: Optional[pulumi.Input['ApplicationScalingRuleScalingRuleMetricScaleUpRulesArgs']] = None):
        """
        :param pulumi.Input[int] max_replicas: Maximum number of instances applied.
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationScalingRuleScalingRuleMetricMetricArgs']]] metrics: Indicator rule configuration. See the following `Block metrics`.
        :param pulumi.Input[int] min_replicas: Minimum number of instances applied.
        :param pulumi.Input['ApplicationScalingRuleScalingRuleMetricScaleDownRulesArgs'] scale_down_rules: Apply shrink rules. See the following `Block scale_down_rules`.
        :param pulumi.Input['ApplicationScalingRuleScalingRuleMetricScaleUpRulesArgs'] scale_up_rules: Apply expansion rules. See the following `Block scale_up_rules`.
        """
        if max_replicas is not None:
            pulumi.set(__self__, "max_replicas", max_replicas)
        if metrics is not None:
            pulumi.set(__self__, "metrics", metrics)
        if min_replicas is not None:
            pulumi.set(__self__, "min_replicas", min_replicas)
        if scale_down_rules is not None:
            pulumi.set(__self__, "scale_down_rules", scale_down_rules)
        if scale_up_rules is not None:
            pulumi.set(__self__, "scale_up_rules", scale_up_rules)

    @property
    @pulumi.getter(name="maxReplicas")
    def max_replicas(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of instances applied.
        """
        return pulumi.get(self, "max_replicas")

    @max_replicas.setter
    def max_replicas(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_replicas", value)

    @property
    @pulumi.getter
    def metrics(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationScalingRuleScalingRuleMetricMetricArgs']]]]:
        """
        Indicator rule configuration. See the following `Block metrics`.
        """
        return pulumi.get(self, "metrics")

    @metrics.setter
    def metrics(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationScalingRuleScalingRuleMetricMetricArgs']]]]):
        pulumi.set(self, "metrics", value)

    @property
    @pulumi.getter(name="minReplicas")
    def min_replicas(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of instances applied.
        """
        return pulumi.get(self, "min_replicas")

    @min_replicas.setter
    def min_replicas(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_replicas", value)

    @property
    @pulumi.getter(name="scaleDownRules")
    def scale_down_rules(self) -> Optional[pulumi.Input['ApplicationScalingRuleScalingRuleMetricScaleDownRulesArgs']]:
        """
        Apply shrink rules. See the following `Block scale_down_rules`.
        """
        return pulumi.get(self, "scale_down_rules")

    @scale_down_rules.setter
    def scale_down_rules(self, value: Optional[pulumi.Input['ApplicationScalingRuleScalingRuleMetricScaleDownRulesArgs']]):
        pulumi.set(self, "scale_down_rules", value)

    @property
    @pulumi.getter(name="scaleUpRules")
    def scale_up_rules(self) -> Optional[pulumi.Input['ApplicationScalingRuleScalingRuleMetricScaleUpRulesArgs']]:
        """
        Apply expansion rules. See the following `Block scale_up_rules`.
        """
        return pulumi.get(self, "scale_up_rules")

    @scale_up_rules.setter
    def scale_up_rules(self, value: Optional[pulumi.Input['ApplicationScalingRuleScalingRuleMetricScaleUpRulesArgs']]):
        pulumi.set(self, "scale_up_rules", value)


@pulumi.input_type
class ApplicationScalingRuleScalingRuleMetricMetricArgs:
    def __init__(__self__, *,
                 metric_target_average_utilization: Optional[pulumi.Input[int]] = None,
                 metric_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] metric_target_average_utilization: According to different `metric_type`, set the target value of the corresponding monitoring index.
        :param pulumi.Input[str] metric_type: Monitoring indicator trigger condition. Valid values: `CPU`, `MEMORY`, `tcpActiveConn`, `SLB_QPS` and `SLB_RT`. The values are described as follows:
               - CPU: CPU usage.
               - MEMORY: MEMORY usage.
               - tcpActiveConn: the average number of TCP active connections for a single instance in 30 seconds.
               - SLB_QPS: the average public network SLB QPS of a single instance within 15 seconds.
               - SLB_RT: the average response time of public network SLB within 15 seconds.
        """
        if metric_target_average_utilization is not None:
            pulumi.set(__self__, "metric_target_average_utilization", metric_target_average_utilization)
        if metric_type is not None:
            pulumi.set(__self__, "metric_type", metric_type)

    @property
    @pulumi.getter(name="metricTargetAverageUtilization")
    def metric_target_average_utilization(self) -> Optional[pulumi.Input[int]]:
        """
        According to different `metric_type`, set the target value of the corresponding monitoring index.
        """
        return pulumi.get(self, "metric_target_average_utilization")

    @metric_target_average_utilization.setter
    def metric_target_average_utilization(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "metric_target_average_utilization", value)

    @property
    @pulumi.getter(name="metricType")
    def metric_type(self) -> Optional[pulumi.Input[str]]:
        """
        Monitoring indicator trigger condition. Valid values: `CPU`, `MEMORY`, `tcpActiveConn`, `SLB_QPS` and `SLB_RT`. The values are described as follows:
        - CPU: CPU usage.
        - MEMORY: MEMORY usage.
        - tcpActiveConn: the average number of TCP active connections for a single instance in 30 seconds.
        - SLB_QPS: the average public network SLB QPS of a single instance within 15 seconds.
        - SLB_RT: the average response time of public network SLB within 15 seconds.
        """
        return pulumi.get(self, "metric_type")

    @metric_type.setter
    def metric_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metric_type", value)


@pulumi.input_type
class ApplicationScalingRuleScalingRuleMetricScaleDownRulesArgs:
    def __init__(__self__, *,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 stabilization_window_seconds: Optional[pulumi.Input[int]] = None,
                 step: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[bool] disabled: Whether shrinkage is prohibited.
        :param pulumi.Input[int] stabilization_window_seconds: Cooling time for expansion or contraction. Valid values: `0` to `3600`. Unit: seconds. The default is `0` seconds.
        :param pulumi.Input[int] step: Elastic expansion or contraction step size. the maximum number of instances to be scaled in per unit time.
        """
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if stabilization_window_seconds is not None:
            pulumi.set(__self__, "stabilization_window_seconds", stabilization_window_seconds)
        if step is not None:
            pulumi.set(__self__, "step", step)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether shrinkage is prohibited.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="stabilizationWindowSeconds")
    def stabilization_window_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Cooling time for expansion or contraction. Valid values: `0` to `3600`. Unit: seconds. The default is `0` seconds.
        """
        return pulumi.get(self, "stabilization_window_seconds")

    @stabilization_window_seconds.setter
    def stabilization_window_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "stabilization_window_seconds", value)

    @property
    @pulumi.getter
    def step(self) -> Optional[pulumi.Input[int]]:
        """
        Elastic expansion or contraction step size. the maximum number of instances to be scaled in per unit time.
        """
        return pulumi.get(self, "step")

    @step.setter
    def step(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "step", value)


@pulumi.input_type
class ApplicationScalingRuleScalingRuleMetricScaleUpRulesArgs:
    def __init__(__self__, *,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 stabilization_window_seconds: Optional[pulumi.Input[int]] = None,
                 step: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[bool] disabled: Whether shrinkage is prohibited.
        :param pulumi.Input[int] stabilization_window_seconds: Cooling time for expansion or contraction. Valid values: `0` to `3600`. Unit: seconds. The default is `0` seconds.
        :param pulumi.Input[int] step: Elastic expansion or contraction step size. the maximum number of instances to be scaled in per unit time.
        """
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if stabilization_window_seconds is not None:
            pulumi.set(__self__, "stabilization_window_seconds", stabilization_window_seconds)
        if step is not None:
            pulumi.set(__self__, "step", step)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether shrinkage is prohibited.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="stabilizationWindowSeconds")
    def stabilization_window_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Cooling time for expansion or contraction. Valid values: `0` to `3600`. Unit: seconds. The default is `0` seconds.
        """
        return pulumi.get(self, "stabilization_window_seconds")

    @stabilization_window_seconds.setter
    def stabilization_window_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "stabilization_window_seconds", value)

    @property
    @pulumi.getter
    def step(self) -> Optional[pulumi.Input[int]]:
        """
        Elastic expansion or contraction step size. the maximum number of instances to be scaled in per unit time.
        """
        return pulumi.get(self, "step")

    @step.setter
    def step(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "step", value)


@pulumi.input_type
class ApplicationScalingRuleScalingRuleTimerArgs:
    def __init__(__self__, *,
                 begin_date: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[str]] = None,
                 schedules: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationScalingRuleScalingRuleTimerScheduleArgs']]]] = None):
        """
        :param pulumi.Input[str] begin_date: The Start date. When the `begin_date` and `end_date` values are empty. it indicates long-term execution and is the default value.
        :param pulumi.Input[str] end_date: The End Date. When the `begin_date` and `end_date` values are empty. it indicates long-term execution and is the default value.
        :param pulumi.Input[str] period: The period in which a timed elastic scaling strategy is executed.
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationScalingRuleScalingRuleTimerScheduleArgs']]] schedules: Resilient Scaling Strategy Trigger Timing. See the following `Block schedules`.
        """
        if begin_date is not None:
            pulumi.set(__self__, "begin_date", begin_date)
        if end_date is not None:
            pulumi.set(__self__, "end_date", end_date)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if schedules is not None:
            pulumi.set(__self__, "schedules", schedules)

    @property
    @pulumi.getter(name="beginDate")
    def begin_date(self) -> Optional[pulumi.Input[str]]:
        """
        The Start date. When the `begin_date` and `end_date` values are empty. it indicates long-term execution and is the default value.
        """
        return pulumi.get(self, "begin_date")

    @begin_date.setter
    def begin_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "begin_date", value)

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[pulumi.Input[str]]:
        """
        The End Date. When the `begin_date` and `end_date` values are empty. it indicates long-term execution and is the default value.
        """
        return pulumi.get(self, "end_date")

    @end_date.setter
    def end_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[str]]:
        """
        The period in which a timed elastic scaling strategy is executed.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter
    def schedules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationScalingRuleScalingRuleTimerScheduleArgs']]]]:
        """
        Resilient Scaling Strategy Trigger Timing. See the following `Block schedules`.
        """
        return pulumi.get(self, "schedules")

    @schedules.setter
    def schedules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationScalingRuleScalingRuleTimerScheduleArgs']]]]):
        pulumi.set(self, "schedules", value)


@pulumi.input_type
class ApplicationScalingRuleScalingRuleTimerScheduleArgs:
    def __init__(__self__, *,
                 at_time: Optional[pulumi.Input[str]] = None,
                 max_replicas: Optional[pulumi.Input[int]] = None,
                 min_replicas: Optional[pulumi.Input[int]] = None,
                 target_replicas: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] at_time: Trigger point in time. When supporting format: minutes, for example: `08:00`.
        :param pulumi.Input[int] max_replicas: Maximum number of instances applied.
        :param pulumi.Input[int] min_replicas: Minimum number of instances applied.
        :param pulumi.Input[int] target_replicas: This parameter can specify the number of instances to be applied or the minimum number of surviving instances per deployment. value range [1,50]. > **NOTE:** The attribute is valid when the attribute `scaling_rule_type` is `timing`.
        """
        if at_time is not None:
            pulumi.set(__self__, "at_time", at_time)
        if max_replicas is not None:
            pulumi.set(__self__, "max_replicas", max_replicas)
        if min_replicas is not None:
            pulumi.set(__self__, "min_replicas", min_replicas)
        if target_replicas is not None:
            pulumi.set(__self__, "target_replicas", target_replicas)

    @property
    @pulumi.getter(name="atTime")
    def at_time(self) -> Optional[pulumi.Input[str]]:
        """
        Trigger point in time. When supporting format: minutes, for example: `08:00`.
        """
        return pulumi.get(self, "at_time")

    @at_time.setter
    def at_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "at_time", value)

    @property
    @pulumi.getter(name="maxReplicas")
    def max_replicas(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of instances applied.
        """
        return pulumi.get(self, "max_replicas")

    @max_replicas.setter
    def max_replicas(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_replicas", value)

    @property
    @pulumi.getter(name="minReplicas")
    def min_replicas(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of instances applied.
        """
        return pulumi.get(self, "min_replicas")

    @min_replicas.setter
    def min_replicas(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_replicas", value)

    @property
    @pulumi.getter(name="targetReplicas")
    def target_replicas(self) -> Optional[pulumi.Input[int]]:
        """
        This parameter can specify the number of instances to be applied or the minimum number of surviving instances per deployment. value range [1,50]. > **NOTE:** The attribute is valid when the attribute `scaling_rule_type` is `timing`.
        """
        return pulumi.get(self, "target_replicas")

    @target_replicas.setter
    def target_replicas(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "target_replicas", value)


@pulumi.input_type
class GreyTagRouteDubboRuleArgs:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 items: Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteDubboRuleItemArgs']]]] = None,
                 method_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] condition: The Conditional Patterns for Grayscale Rules. Valid values: `AND`, `OR`.
        :param pulumi.Input[str] group: The service group.
        :param pulumi.Input[Sequence[pulumi.Input['GreyTagRouteDubboRuleItemArgs']]] items: A list of conditions items. The details see Block `dubbo_rules_items`.
        :param pulumi.Input[str] method_name: The method name
        :param pulumi.Input[str] service_name: The service name.
        :param pulumi.Input[str] version: The service version.
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if items is not None:
            pulumi.set(__self__, "items", items)
        if method_name is not None:
            pulumi.set(__self__, "method_name", method_name)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input[str]]:
        """
        The Conditional Patterns for Grayscale Rules. Valid values: `AND`, `OR`.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        The service group.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def items(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteDubboRuleItemArgs']]]]:
        """
        A list of conditions items. The details see Block `dubbo_rules_items`.
        """
        return pulumi.get(self, "items")

    @items.setter
    def items(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteDubboRuleItemArgs']]]]):
        pulumi.set(self, "items", value)

    @property
    @pulumi.getter(name="methodName")
    def method_name(self) -> Optional[pulumi.Input[str]]:
        """
        The method name
        """
        return pulumi.get(self, "method_name")

    @method_name.setter
    def method_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "method_name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The service name.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The service version.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class GreyTagRouteDubboRuleItemArgs:
    def __init__(__self__, *,
                 cond: Optional[pulumi.Input[str]] = None,
                 expr: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[int]] = None,
                 operator: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] cond: The comparison operator. Valid values: `>`, `<`, `>=`, `<=`, `==`, `!=`.
        :param pulumi.Input[str] expr: The parameter value gets the expression.
        :param pulumi.Input[int] index: The parameter number.
        :param pulumi.Input[str] operator: The operator. Valid values: `rawvalue`, `list`, `mod`, `deterministic_proportional_steaming_division`
        :param pulumi.Input[str] value: The value of the parameter.
        """
        if cond is not None:
            pulumi.set(__self__, "cond", cond)
        if expr is not None:
            pulumi.set(__self__, "expr", expr)
        if index is not None:
            pulumi.set(__self__, "index", index)
        if operator is not None:
            pulumi.set(__self__, "operator", operator)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def cond(self) -> Optional[pulumi.Input[str]]:
        """
        The comparison operator. Valid values: `>`, `<`, `>=`, `<=`, `==`, `!=`.
        """
        return pulumi.get(self, "cond")

    @cond.setter
    def cond(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cond", value)

    @property
    @pulumi.getter
    def expr(self) -> Optional[pulumi.Input[str]]:
        """
        The parameter value gets the expression.
        """
        return pulumi.get(self, "expr")

    @expr.setter
    def expr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expr", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[int]]:
        """
        The parameter number.
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def operator(self) -> Optional[pulumi.Input[str]]:
        """
        The operator. Valid values: `rawvalue`, `list`, `mod`, `deterministic_proportional_steaming_division`
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the parameter.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GreyTagRouteScRuleArgs:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input[str]] = None,
                 items: Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteScRuleItemArgs']]]] = None,
                 path: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] condition: The conditional Patterns for Grayscale Rules. Valid values: `AND`, `OR`.
        :param pulumi.Input[Sequence[pulumi.Input['GreyTagRouteScRuleItemArgs']]] items: A list of conditions items. The details see Block `sc_rules_items`.
        :param pulumi.Input[str] path: The path corresponding to the grayscale rule.
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if items is not None:
            pulumi.set(__self__, "items", items)
        if path is not None:
            pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input[str]]:
        """
        The conditional Patterns for Grayscale Rules. Valid values: `AND`, `OR`.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def items(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteScRuleItemArgs']]]]:
        """
        A list of conditions items. The details see Block `sc_rules_items`.
        """
        return pulumi.get(self, "items")

    @items.setter
    def items(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteScRuleItemArgs']]]]):
        pulumi.set(self, "items", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The path corresponding to the grayscale rule.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)


@pulumi.input_type
class GreyTagRouteScRuleItemArgs:
    def __init__(__self__, *,
                 cond: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 operator: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] cond: The comparison operator. Valid values: `>`, `<`, `>=`, `<=`, `==`, `!=`.
        :param pulumi.Input[str] name: The name of the parameter.
        :param pulumi.Input[str] operator: The operator. Valid values: `rawvalue`, `list`, `mod`, `deterministic_proportional_steaming_division`
        :param pulumi.Input[str] type: The compare types. Valid values: `param`, `cookie`, `header`.
        :param pulumi.Input[str] value: The value of the parameter.
        """
        if cond is not None:
            pulumi.set(__self__, "cond", cond)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if operator is not None:
            pulumi.set(__self__, "operator", operator)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def cond(self) -> Optional[pulumi.Input[str]]:
        """
        The comparison operator. Valid values: `>`, `<`, `>=`, `<=`, `==`, `!=`.
        """
        return pulumi.get(self, "cond")

    @cond.setter
    def cond(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cond", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the parameter.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def operator(self) -> Optional[pulumi.Input[str]]:
        """
        The operator. Valid values: `rawvalue`, `list`, `mod`, `deterministic_proportional_steaming_division`
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The compare types. Valid values: `param`, `cookie`, `header`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the parameter.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class IngressDefaultRuleArgs:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[str]] = None,
                 app_name: Optional[pulumi.Input[str]] = None,
                 container_port: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] app_id: Target application ID.
        :param pulumi.Input[str] app_name: Target application name.
        :param pulumi.Input[int] container_port: Application backend port.
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if app_name is not None:
            pulumi.set(__self__, "app_name", app_name)
        if container_port is not None:
            pulumi.set(__self__, "container_port", container_port)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[str]]:
        """
        Target application ID.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="appName")
    def app_name(self) -> Optional[pulumi.Input[str]]:
        """
        Target application name.
        """
        return pulumi.get(self, "app_name")

    @app_name.setter
    def app_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_name", value)

    @property
    @pulumi.getter(name="containerPort")
    def container_port(self) -> Optional[pulumi.Input[int]]:
        """
        Application backend port.
        """
        return pulumi.get(self, "container_port")

    @container_port.setter
    def container_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "container_port", value)


@pulumi.input_type
class IngressRuleArgs:
    def __init__(__self__, *,
                 app_id: pulumi.Input[str],
                 app_name: pulumi.Input[str],
                 container_port: pulumi.Input[int],
                 domain: pulumi.Input[str],
                 path: pulumi.Input[str]):
        """
        :param pulumi.Input[str] app_id: Target application ID.
        :param pulumi.Input[str] app_name: Target application name.
        :param pulumi.Input[int] container_port: Application backend port.
        :param pulumi.Input[str] domain: Application domain name.
        :param pulumi.Input[str] path: URL path.
        """
        pulumi.set(__self__, "app_id", app_id)
        pulumi.set(__self__, "app_name", app_name)
        pulumi.set(__self__, "container_port", container_port)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Input[str]:
        """
        Target application ID.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="appName")
    def app_name(self) -> pulumi.Input[str]:
        """
        Target application name.
        """
        return pulumi.get(self, "app_name")

    @app_name.setter
    def app_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_name", value)

    @property
    @pulumi.getter(name="containerPort")
    def container_port(self) -> pulumi.Input[int]:
        """
        Application backend port.
        """
        return pulumi.get(self, "container_port")

    @container_port.setter
    def container_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "container_port", value)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        Application domain name.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        URL path.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)


@pulumi.input_type
class LoadBalancerInternetInternetArgs:
    def __init__(__self__, *,
                 https_cert_id: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 target_port: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] https_cert_id: The SSL certificate. `https_cert_id` is required when HTTPS is selected
        :param pulumi.Input[int] port: The SLB Port.
        :param pulumi.Input[str] protocol: The Network protocol. Valid values: `TCP` ,`HTTP`,`HTTPS`.
        :param pulumi.Input[int] target_port: The Container port.
        """
        if https_cert_id is not None:
            pulumi.set(__self__, "https_cert_id", https_cert_id)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if target_port is not None:
            pulumi.set(__self__, "target_port", target_port)

    @property
    @pulumi.getter(name="httpsCertId")
    def https_cert_id(self) -> Optional[pulumi.Input[str]]:
        """
        The SSL certificate. `https_cert_id` is required when HTTPS is selected
        """
        return pulumi.get(self, "https_cert_id")

    @https_cert_id.setter
    def https_cert_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "https_cert_id", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The SLB Port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The Network protocol. Valid values: `TCP` ,`HTTP`,`HTTPS`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="targetPort")
    def target_port(self) -> Optional[pulumi.Input[int]]:
        """
        The Container port.
        """
        return pulumi.get(self, "target_port")

    @target_port.setter
    def target_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "target_port", value)


@pulumi.input_type
class LoadBalancerIntranetIntranetArgs:
    def __init__(__self__, *,
                 https_cert_id: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 target_port: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] https_cert_id: The SSL certificate. `https_cert_id` is required when HTTPS is selected
        :param pulumi.Input[int] port: The SLB Port.
        :param pulumi.Input[str] protocol: The Network protocol. Valid values: `TCP` ,`HTTP`,`HTTPS`.
        :param pulumi.Input[int] target_port: The Container port.
        """
        if https_cert_id is not None:
            pulumi.set(__self__, "https_cert_id", https_cert_id)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if target_port is not None:
            pulumi.set(__self__, "target_port", target_port)

    @property
    @pulumi.getter(name="httpsCertId")
    def https_cert_id(self) -> Optional[pulumi.Input[str]]:
        """
        The SSL certificate. `https_cert_id` is required when HTTPS is selected
        """
        return pulumi.get(self, "https_cert_id")

    @https_cert_id.setter
    def https_cert_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "https_cert_id", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The SLB Port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The Network protocol. Valid values: `TCP` ,`HTTP`,`HTTPS`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="targetPort")
    def target_port(self) -> Optional[pulumi.Input[int]]:
        """
        The Container port.
        """
        return pulumi.get(self, "target_port")

    @target_port.setter
    def target_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "target_port", value)


