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
    'GetEcsNetworkInterfacePermissionsResult',
    'AwaitableGetEcsNetworkInterfacePermissionsResult',
    'get_ecs_network_interface_permissions',
    'get_ecs_network_interface_permissions_output',
]

@pulumi.output_type
class GetEcsNetworkInterfacePermissionsResult:
    """
    A collection of values returned by getEcsNetworkInterfacePermissions.
    """
    def __init__(__self__, id=None, ids=None, network_interface_id=None, output_file=None, page_number=None, page_size=None, permissions=None, status=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if network_interface_id and not isinstance(network_interface_id, str):
            raise TypeError("Expected argument 'network_interface_id' to be a str")
        pulumi.set(__self__, "network_interface_id", network_interface_id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if page_number and not isinstance(page_number, int):
            raise TypeError("Expected argument 'page_number' to be a int")
        pulumi.set(__self__, "page_number", page_number)
        if page_size and not isinstance(page_size, int):
            raise TypeError("Expected argument 'page_size' to be a int")
        pulumi.set(__self__, "page_size", page_size)
        if permissions and not isinstance(permissions, list):
            raise TypeError("Expected argument 'permissions' to be a list")
        pulumi.set(__self__, "permissions", permissions)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

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
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> str:
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="pageNumber")
    def page_number(self) -> Optional[int]:
        return pulumi.get(self, "page_number")

    @property
    @pulumi.getter(name="pageSize")
    def page_size(self) -> Optional[int]:
        return pulumi.get(self, "page_size")

    @property
    @pulumi.getter
    def permissions(self) -> Sequence['outputs.GetEcsNetworkInterfacePermissionsPermissionResult']:
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        return pulumi.get(self, "total_count")


class AwaitableGetEcsNetworkInterfacePermissionsResult(GetEcsNetworkInterfacePermissionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEcsNetworkInterfacePermissionsResult(
            id=self.id,
            ids=self.ids,
            network_interface_id=self.network_interface_id,
            output_file=self.output_file,
            page_number=self.page_number,
            page_size=self.page_size,
            permissions=self.permissions,
            status=self.status,
            total_count=self.total_count)


def get_ecs_network_interface_permissions(ids: Optional[Sequence[str]] = None,
                                          network_interface_id: Optional[str] = None,
                                          output_file: Optional[str] = None,
                                          page_number: Optional[int] = None,
                                          page_size: Optional[int] = None,
                                          status: Optional[str] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEcsNetworkInterfacePermissionsResult:
    """
    This data source provides the Ecs Network Interface Permissions of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.166.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.ecs.get_ecs_network_interface_permissions(ids=["example_value"],
        network_interface_id="example_value")
    pulumi.export("ecsNetworkInterfacePermissionId1", ids.permissions[0].id)
    ```


    :param Sequence[str] ids: A list of Network Interface Permission IDs.
    :param str network_interface_id: The ID of the network interface.
    :param str status: The Status of the Network Interface Permissions.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['networkInterfaceId'] = network_interface_id
    __args__['outputFile'] = output_file
    __args__['pageNumber'] = page_number
    __args__['pageSize'] = page_size
    __args__['status'] = status
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:ecs/getEcsNetworkInterfacePermissions:getEcsNetworkInterfacePermissions', __args__, opts=opts, typ=GetEcsNetworkInterfacePermissionsResult).value

    return AwaitableGetEcsNetworkInterfacePermissionsResult(
        id=__ret__.id,
        ids=__ret__.ids,
        network_interface_id=__ret__.network_interface_id,
        output_file=__ret__.output_file,
        page_number=__ret__.page_number,
        page_size=__ret__.page_size,
        permissions=__ret__.permissions,
        status=__ret__.status,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(get_ecs_network_interface_permissions)
def get_ecs_network_interface_permissions_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                                 network_interface_id: Optional[pulumi.Input[str]] = None,
                                                 output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                                 page_number: Optional[pulumi.Input[Optional[int]]] = None,
                                                 page_size: Optional[pulumi.Input[Optional[int]]] = None,
                                                 status: Optional[pulumi.Input[Optional[str]]] = None,
                                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEcsNetworkInterfacePermissionsResult]:
    """
    This data source provides the Ecs Network Interface Permissions of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.166.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.ecs.get_ecs_network_interface_permissions(ids=["example_value"],
        network_interface_id="example_value")
    pulumi.export("ecsNetworkInterfacePermissionId1", ids.permissions[0].id)
    ```


    :param Sequence[str] ids: A list of Network Interface Permission IDs.
    :param str network_interface_id: The ID of the network interface.
    :param str status: The Status of the Network Interface Permissions.
    """
    ...
