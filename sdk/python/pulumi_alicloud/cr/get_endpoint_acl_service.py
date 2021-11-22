# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetEndpointAclServiceResult',
    'AwaitableGetEndpointAclServiceResult',
    'get_endpoint_acl_service',
    'get_endpoint_acl_service_output',
]

@pulumi.output_type
class GetEndpointAclServiceResult:
    """
    A collection of values returned by getEndpointAclService.
    """
    def __init__(__self__, enable=None, endpoint_type=None, id=None, instance_id=None, module_name=None, status=None):
        if enable and not isinstance(enable, bool):
            raise TypeError("Expected argument 'enable' to be a bool")
        pulumi.set(__self__, "enable", enable)
        if endpoint_type and not isinstance(endpoint_type, str):
            raise TypeError("Expected argument 'endpoint_type' to be a str")
        pulumi.set(__self__, "endpoint_type", endpoint_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if module_name and not isinstance(module_name, str):
            raise TypeError("Expected argument 'module_name' to be a str")
        pulumi.set(__self__, "module_name", module_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def enable(self) -> bool:
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> str:
        return pulumi.get(self, "endpoint_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="moduleName")
    def module_name(self) -> Optional[str]:
        return pulumi.get(self, "module_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")


class AwaitableGetEndpointAclServiceResult(GetEndpointAclServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEndpointAclServiceResult(
            enable=self.enable,
            endpoint_type=self.endpoint_type,
            id=self.id,
            instance_id=self.instance_id,
            module_name=self.module_name,
            status=self.status)


def get_endpoint_acl_service(enable: Optional[bool] = None,
                             endpoint_type: Optional[str] = None,
                             instance_id: Optional[str] = None,
                             module_name: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEndpointAclServiceResult:
    """
    This data source provides the CR Endpoint Acl Service of the current Alibaba Cloud user.

    For information about Event Bridge and how to use it, see [What is CR Endpoint Acl](https://www.alibabacloud.com/help/en/doc-detail/142246.htm).

    > **NOTE:** Available in v1.139.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.cr.get_endpoint_acl_service(enable=True,
        endpoint_type="internet",
        instance_id="example_id",
        module_name="Registry")
    ```


    :param bool enable: Whether to enable Acl Service.  Valid values: `true` and `false`.
    :param str endpoint_type: The type of endpoint. Valid values: `internet`.
    :param str instance_id: The ID of the CR Instance.
    :param str module_name: The ModuleName. Valid values: `Registry`.
    """
    __args__ = dict()
    __args__['enable'] = enable
    __args__['endpointType'] = endpoint_type
    __args__['instanceId'] = instance_id
    __args__['moduleName'] = module_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:cr/getEndpointAclService:getEndpointAclService', __args__, opts=opts, typ=GetEndpointAclServiceResult).value

    return AwaitableGetEndpointAclServiceResult(
        enable=__ret__.enable,
        endpoint_type=__ret__.endpoint_type,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        module_name=__ret__.module_name,
        status=__ret__.status)


@_utilities.lift_output_func(get_endpoint_acl_service)
def get_endpoint_acl_service_output(enable: Optional[pulumi.Input[bool]] = None,
                                    endpoint_type: Optional[pulumi.Input[str]] = None,
                                    instance_id: Optional[pulumi.Input[str]] = None,
                                    module_name: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEndpointAclServiceResult]:
    """
    This data source provides the CR Endpoint Acl Service of the current Alibaba Cloud user.

    For information about Event Bridge and how to use it, see [What is CR Endpoint Acl](https://www.alibabacloud.com/help/en/doc-detail/142246.htm).

    > **NOTE:** Available in v1.139.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.cr.get_endpoint_acl_service(enable=True,
        endpoint_type="internet",
        instance_id="example_id",
        module_name="Registry")
    ```


    :param bool enable: Whether to enable Acl Service.  Valid values: `true` and `false`.
    :param str endpoint_type: The type of endpoint. Valid values: `internet`.
    :param str instance_id: The ID of the CR Instance.
    :param str module_name: The ModuleName. Valid values: `Registry`.
    """
    ...
