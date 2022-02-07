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
    'GetAccessStrategiesResult',
    'AwaitableGetAccessStrategiesResult',
    'get_access_strategies',
    'get_access_strategies_output',
]

@pulumi.output_type
class GetAccessStrategiesResult:
    """
    A collection of values returned by getAccessStrategies.
    """
    def __init__(__self__, enable_details=None, id=None, ids=None, instance_id=None, lang=None, name_regex=None, names=None, output_file=None, strategies=None, strategy_mode=None):
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if lang and not isinstance(lang, str):
            raise TypeError("Expected argument 'lang' to be a str")
        pulumi.set(__self__, "lang", lang)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if strategies and not isinstance(strategies, list):
            raise TypeError("Expected argument 'strategies' to be a list")
        pulumi.set(__self__, "strategies", strategies)
        if strategy_mode and not isinstance(strategy_mode, str):
            raise TypeError("Expected argument 'strategy_mode' to be a str")
        pulumi.set(__self__, "strategy_mode", strategy_mode)

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

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
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def lang(self) -> Optional[str]:
        return pulumi.get(self, "lang")

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
    @pulumi.getter
    def strategies(self) -> Sequence['outputs.GetAccessStrategiesStrategyResult']:
        return pulumi.get(self, "strategies")

    @property
    @pulumi.getter(name="strategyMode")
    def strategy_mode(self) -> str:
        return pulumi.get(self, "strategy_mode")


class AwaitableGetAccessStrategiesResult(GetAccessStrategiesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccessStrategiesResult(
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            instance_id=self.instance_id,
            lang=self.lang,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            strategies=self.strategies,
            strategy_mode=self.strategy_mode)


def get_access_strategies(enable_details: Optional[bool] = None,
                          ids: Optional[Sequence[str]] = None,
                          instance_id: Optional[str] = None,
                          lang: Optional[str] = None,
                          name_regex: Optional[str] = None,
                          output_file: Optional[str] = None,
                          strategy_mode: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccessStrategiesResult:
    """
    This data source provides the Alidns Access Strategies of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.152.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.dns.get_access_strategies(instance_id="example_value",
        strategy_mode="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ],
        name_regex="the_resource_name")
    pulumi.export("alidnsAccessStrategyId1", ids.strategies[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Access Strategy IDs.
    :param str instance_id: The Id of the associated instance.
    :param str lang: The lang.
    :param str name_regex: A regex string to filter results by Access Strategy name.
    :param str strategy_mode: The type of the access policy.
    """
    __args__ = dict()
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['instanceId'] = instance_id
    __args__['lang'] = lang
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['strategyMode'] = strategy_mode
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:dns/getAccessStrategies:getAccessStrategies', __args__, opts=opts, typ=GetAccessStrategiesResult).value

    return AwaitableGetAccessStrategiesResult(
        enable_details=__ret__.enable_details,
        id=__ret__.id,
        ids=__ret__.ids,
        instance_id=__ret__.instance_id,
        lang=__ret__.lang,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        strategies=__ret__.strategies,
        strategy_mode=__ret__.strategy_mode)


@_utilities.lift_output_func(get_access_strategies)
def get_access_strategies_output(enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                                 ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                 instance_id: Optional[pulumi.Input[str]] = None,
                                 lang: Optional[pulumi.Input[Optional[str]]] = None,
                                 name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                 output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 strategy_mode: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccessStrategiesResult]:
    """
    This data source provides the Alidns Access Strategies of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.152.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.dns.get_access_strategies(instance_id="example_value",
        strategy_mode="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ],
        name_regex="the_resource_name")
    pulumi.export("alidnsAccessStrategyId1", ids.strategies[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Access Strategy IDs.
    :param str instance_id: The Id of the associated instance.
    :param str lang: The lang.
    :param str name_regex: A regex string to filter results by Access Strategy name.
    :param str strategy_mode: The type of the access policy.
    """
    ...