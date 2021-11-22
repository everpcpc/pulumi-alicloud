# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetAckServiceResult',
    'AwaitableGetAckServiceResult',
    'get_ack_service',
    'get_ack_service_output',
]

@pulumi.output_type
class GetAckServiceResult:
    """
    A collection of values returned by getAckService.
    """
    def __init__(__self__, enable=None, id=None, status=None, type=None):
        if enable and not isinstance(enable, str):
            raise TypeError("Expected argument 'enable' to be a str")
        pulumi.set(__self__, "enable", enable)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def enable(self) -> Optional[str]:
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The current service enable status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


class AwaitableGetAckServiceResult(GetAckServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAckServiceResult(
            enable=self.enable,
            id=self.id,
            status=self.status,
            type=self.type)


def get_ack_service(enable: Optional[str] = None,
                    type: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAckServiceResult:
    """
    Using this data source can open Container Service (CS) service automatically. If the service has been opened, it will return opened.

    For information about Container Service (CS) and how to use it, see [What is Container Service (CS)](https://www.alibabacloud.com/help/en/product/85222.htm).

    > **NOTE:** Available in v1.113.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    open = alicloud.cs.get_ack_service(enable="On",
        type="propayasgo")
    ```


    :param str enable: Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
    :param str type: Types of services opened. Valid values: `propayasgo`: Container service ack Pro managed version, `edgepayasgo`: Edge container service, `gspayasgo`: Gene computing services.
    """
    __args__ = dict()
    __args__['enable'] = enable
    __args__['type'] = type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:cs/getAckService:getAckService', __args__, opts=opts, typ=GetAckServiceResult).value

    return AwaitableGetAckServiceResult(
        enable=__ret__.enable,
        id=__ret__.id,
        status=__ret__.status,
        type=__ret__.type)


@_utilities.lift_output_func(get_ack_service)
def get_ack_service_output(enable: Optional[pulumi.Input[Optional[str]]] = None,
                           type: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAckServiceResult]:
    """
    Using this data source can open Container Service (CS) service automatically. If the service has been opened, it will return opened.

    For information about Container Service (CS) and how to use it, see [What is Container Service (CS)](https://www.alibabacloud.com/help/en/product/85222.htm).

    > **NOTE:** Available in v1.113.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    open = alicloud.cs.get_ack_service(enable="On",
        type="propayasgo")
    ```


    :param str enable: Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
    :param str type: Types of services opened. Valid values: `propayasgo`: Container service ack Pro managed version, `edgepayasgo`: Edge container service, `gspayasgo`: Gene computing services.
    """
    ...
