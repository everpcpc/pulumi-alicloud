# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetServiceResult',
    'AwaitableGetServiceResult',
    'get_service',
]

@pulumi.output_type
class GetServiceResult:
    """
    A collection of values returned by getService.
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


class AwaitableGetServiceResult(GetServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceResult(
            enable=self.enable,
            id=self.id,
            status=self.status)


def get_service(enable: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceResult:
    """
    Using this data source can enable Log service automatically. If the service has been enabled, it will return `Opened`.

    For information about Log service and how to use it, see [What is Log Service](https://www.alibabacloud.com/help/product/28958.htm).

    > **NOTE:** Available in v1.96.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    open = alicloud.log.get_service(enable="On")
    ```


    :param str enable: Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: "On" or "Off".
    """
    __args__ = dict()
    __args__['enable'] = enable
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:log/getService:getService', __args__, opts=opts, typ=GetServiceResult).value

    return AwaitableGetServiceResult(
        enable=__ret__.enable,
        id=__ret__.id,
        status=__ret__.status)
