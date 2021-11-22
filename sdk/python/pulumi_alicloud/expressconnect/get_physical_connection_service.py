# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetPhysicalConnectionServiceResult',
    'AwaitableGetPhysicalConnectionServiceResult',
    'get_physical_connection_service',
    'get_physical_connection_service_output',
]

@pulumi.output_type
class GetPhysicalConnectionServiceResult:
    """
    A collection of values returned by getPhysicalConnectionService.
    """
    def __init__(__self__, enable=None, id=None, status=None):
        if enable and not isinstance(enable, str):
            raise TypeError("Expected argument 'enable' to be a str")
        pulumi.set(__self__, "enable", enable)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

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


class AwaitableGetPhysicalConnectionServiceResult(GetPhysicalConnectionServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPhysicalConnectionServiceResult(
            enable=self.enable,
            id=self.id,
            status=self.status)


def get_physical_connection_service(enable: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPhysicalConnectionServiceResult:
    """
    Using this data source can enable outbound traffic for an Express Connect circuit automatically. If the service has been opened, it will return opened.

    For information about Express Connect and how to use it, see [What is Express Connect](https://www.alibabacloud.com/help/doc-detail/275179.htm).

    > **NOTE:** Available in v1.132.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    open = alicloud.expressconnect.get_physical_connection_service(enable="On")
    ```


    :param str enable: Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
    """
    __args__ = dict()
    __args__['enable'] = enable
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:expressconnect/getPhysicalConnectionService:getPhysicalConnectionService', __args__, opts=opts, typ=GetPhysicalConnectionServiceResult).value

    return AwaitableGetPhysicalConnectionServiceResult(
        enable=__ret__.enable,
        id=__ret__.id,
        status=__ret__.status)


@_utilities.lift_output_func(get_physical_connection_service)
def get_physical_connection_service_output(enable: Optional[pulumi.Input[Optional[str]]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPhysicalConnectionServiceResult]:
    """
    Using this data source can enable outbound traffic for an Express Connect circuit automatically. If the service has been opened, it will return opened.

    For information about Express Connect and how to use it, see [What is Express Connect](https://www.alibabacloud.com/help/doc-detail/275179.htm).

    > **NOTE:** Available in v1.132.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    open = alicloud.expressconnect.get_physical_connection_service(enable="On")
    ```


    :param str enable: Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
    """
    ...
