# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetRulesResult',
    'AwaitableGetRulesResult',
    'get_rules',
]

@pulumi.output_type
class GetRulesResult:
    """
    A collection of values returned by getRules.
    """
    def __init__(__self__, event_bus_name=None, id=None, ids=None, name_regex=None, names=None, output_file=None, rule_name_prefix=None, rules=None, status=None):
        if event_bus_name and not isinstance(event_bus_name, str):
            raise TypeError("Expected argument 'event_bus_name' to be a str")
        pulumi.set(__self__, "event_bus_name", event_bus_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if rule_name_prefix and not isinstance(rule_name_prefix, str):
            raise TypeError("Expected argument 'rule_name_prefix' to be a str")
        pulumi.set(__self__, "rule_name_prefix", rule_name_prefix)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="eventBusName")
    def event_bus_name(self) -> str:
        return pulumi.get(self, "event_bus_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="ruleNamePrefix")
    def rule_name_prefix(self) -> Optional[str]:
        return pulumi.get(self, "rule_name_prefix")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetRulesRuleResult']:
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


class AwaitableGetRulesResult(GetRulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRulesResult(
            event_bus_name=self.event_bus_name,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            rule_name_prefix=self.rule_name_prefix,
            rules=self.rules,
            status=self.status)


def get_rules(event_bus_name: Optional[str] = None,
              ids: Optional[Sequence[str]] = None,
              name_regex: Optional[str] = None,
              output_file: Optional[str] = None,
              rule_name_prefix: Optional[str] = None,
              status: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRulesResult:
    """
    This data source provides the Event Bridge Rules of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.129.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.eventbridge.get_rules(event_bus_name="example_value",
        ids=["example_value"],
        name_regex="the_resource_name")
    pulumi.export("firstEventBridgeRuleId", example.rules[0].id)
    ```


    :param str event_bus_name: The name of event bus.
    :param Sequence[str] ids: A list of Rule IDs.
    :param str name_regex: A regex string to filter results by Rule name.
    :param str rule_name_prefix: The rule name prefix.
    :param str status: Rule status, either Enable or Disable.
    """
    __args__ = dict()
    __args__['eventBusName'] = event_bus_name
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['ruleNamePrefix'] = rule_name_prefix
    __args__['status'] = status
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:eventbridge/getRules:getRules', __args__, opts=opts, typ=GetRulesResult).value

    return AwaitableGetRulesResult(
        event_bus_name=__ret__.event_bus_name,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        rule_name_prefix=__ret__.rule_name_prefix,
        rules=__ret__.rules,
        status=__ret__.status)