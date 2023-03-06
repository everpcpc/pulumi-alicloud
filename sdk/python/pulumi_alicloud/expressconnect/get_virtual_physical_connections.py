# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetVirtualPhysicalConnectionsResult',
    'AwaitableGetVirtualPhysicalConnectionsResult',
    'get_virtual_physical_connections',
    'get_virtual_physical_connections_output',
]

@pulumi.output_type
class GetVirtualPhysicalConnectionsResult:
    """
    A collection of values returned by getVirtualPhysicalConnections.
    """
    def __init__(__self__, business_status=None, connections=None, id=None, ids=None, is_confirmed=None, name_regex=None, names=None, output_file=None, parent_physical_connection_id=None, virtual_physical_connection_ids=None, virtual_physical_connection_status=None, vlan_ids=None, vpconn_ali_uid=None):
        if business_status and not isinstance(business_status, str):
            raise TypeError("Expected argument 'business_status' to be a str")
        pulumi.set(__self__, "business_status", business_status)
        if connections and not isinstance(connections, list):
            raise TypeError("Expected argument 'connections' to be a list")
        pulumi.set(__self__, "connections", connections)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if is_confirmed and not isinstance(is_confirmed, bool):
            raise TypeError("Expected argument 'is_confirmed' to be a bool")
        pulumi.set(__self__, "is_confirmed", is_confirmed)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if parent_physical_connection_id and not isinstance(parent_physical_connection_id, str):
            raise TypeError("Expected argument 'parent_physical_connection_id' to be a str")
        pulumi.set(__self__, "parent_physical_connection_id", parent_physical_connection_id)
        if virtual_physical_connection_ids and not isinstance(virtual_physical_connection_ids, list):
            raise TypeError("Expected argument 'virtual_physical_connection_ids' to be a list")
        pulumi.set(__self__, "virtual_physical_connection_ids", virtual_physical_connection_ids)
        if virtual_physical_connection_status and not isinstance(virtual_physical_connection_status, str):
            raise TypeError("Expected argument 'virtual_physical_connection_status' to be a str")
        pulumi.set(__self__, "virtual_physical_connection_status", virtual_physical_connection_status)
        if vlan_ids and not isinstance(vlan_ids, list):
            raise TypeError("Expected argument 'vlan_ids' to be a list")
        pulumi.set(__self__, "vlan_ids", vlan_ids)
        if vpconn_ali_uid and not isinstance(vpconn_ali_uid, str):
            raise TypeError("Expected argument 'vpconn_ali_uid' to be a str")
        pulumi.set(__self__, "vpconn_ali_uid", vpconn_ali_uid)

    @property
    @pulumi.getter(name="businessStatus")
    def business_status(self) -> Optional[str]:
        """
        The commercial status of the physical line. Value:-**Normal**: activated.-**Financialized**: Arrears locked.-**SecurityLocked**: locked for security reasons.
        """
        return pulumi.get(self, "business_status")

    @property
    @pulumi.getter
    def connections(self) -> Sequence['outputs.GetVirtualPhysicalConnectionsConnectionResult']:
        """
        A list of Virtual Physical Connection Entries. Each element contains the following attributes:
        """
        return pulumi.get(self, "connections")

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
        """
        A list of Virtual Physical Connection IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="isConfirmed")
    def is_confirmed(self) -> Optional[bool]:
        return pulumi.get(self, "is_confirmed")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of name of Virtual Physical Connections.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="parentPhysicalConnectionId")
    def parent_physical_connection_id(self) -> Optional[str]:
        """
        The ID of the instance of the physical connection.
        """
        return pulumi.get(self, "parent_physical_connection_id")

    @property
    @pulumi.getter(name="virtualPhysicalConnectionIds")
    def virtual_physical_connection_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "virtual_physical_connection_ids")

    @property
    @pulumi.getter(name="virtualPhysicalConnectionStatus")
    def virtual_physical_connection_status(self) -> Optional[str]:
        """
        The business status of the shared line. Value:-**Confirmed**: The shared line has been Confirmed to receive.-**UnConfirmed**: The shared line has not been confirmed to be received.-**Deleted**: The shared line has been Deleted.
        """
        return pulumi.get(self, "virtual_physical_connection_status")

    @property
    @pulumi.getter(name="vlanIds")
    def vlan_ids(self) -> Optional[Sequence[int]]:
        return pulumi.get(self, "vlan_ids")

    @property
    @pulumi.getter(name="vpconnAliUid")
    def vpconn_ali_uid(self) -> Optional[str]:
        """
        The ID of the Alibaba Cloud account (primary account) of the owner of the shared line.
        """
        return pulumi.get(self, "vpconn_ali_uid")


class AwaitableGetVirtualPhysicalConnectionsResult(GetVirtualPhysicalConnectionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualPhysicalConnectionsResult(
            business_status=self.business_status,
            connections=self.connections,
            id=self.id,
            ids=self.ids,
            is_confirmed=self.is_confirmed,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            parent_physical_connection_id=self.parent_physical_connection_id,
            virtual_physical_connection_ids=self.virtual_physical_connection_ids,
            virtual_physical_connection_status=self.virtual_physical_connection_status,
            vlan_ids=self.vlan_ids,
            vpconn_ali_uid=self.vpconn_ali_uid)


def get_virtual_physical_connections(business_status: Optional[str] = None,
                                     ids: Optional[Sequence[str]] = None,
                                     is_confirmed: Optional[bool] = None,
                                     name_regex: Optional[str] = None,
                                     output_file: Optional[str] = None,
                                     parent_physical_connection_id: Optional[str] = None,
                                     virtual_physical_connection_ids: Optional[Sequence[str]] = None,
                                     virtual_physical_connection_status: Optional[str] = None,
                                     vlan_ids: Optional[Sequence[int]] = None,
                                     vpconn_ali_uid: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualPhysicalConnectionsResult:
    """
    This data source provides Express Connect Virtual Physical Connection available to the user.

    > **NOTE:** Available in 1.196.0+


    :param str business_status: The commercial status of the physical line. Value:
           - **Normal**: activated.
           - **Financialized**: Arrears locked.
           - **SecurityLocked**: locked for security reasons.
    :param Sequence[str] ids: A list of Virtual Physical Connection IDs.
    :param str name_regex: A regex string to filter results by Group Metric Rule name.
    :param str parent_physical_connection_id: The ID of the instance of the physical connection.
    :param Sequence[str] virtual_physical_connection_ids: The ID of the hosted connection. You can specify multiple hosted connection IDs.
    :param str virtual_physical_connection_status: The business status of the shared line. Value:
           - **Confirmed**: The shared line has been Confirmed to receive.
           - **UnConfirmed**: The shared line has not been confirmed to be received.
           - **Deleted**: The shared line has been Deleted.
    :param Sequence[int] vlan_ids: The VLAN ID of the hosted connection. You can specify multiple VLAN IDs.
    :param str vpconn_ali_uid: The ID of the Alibaba Cloud account (primary account) of the owner of the shared line.
    """
    __args__ = dict()
    __args__['businessStatus'] = business_status
    __args__['ids'] = ids
    __args__['isConfirmed'] = is_confirmed
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['parentPhysicalConnectionId'] = parent_physical_connection_id
    __args__['virtualPhysicalConnectionIds'] = virtual_physical_connection_ids
    __args__['virtualPhysicalConnectionStatus'] = virtual_physical_connection_status
    __args__['vlanIds'] = vlan_ids
    __args__['vpconnAliUid'] = vpconn_ali_uid
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:expressconnect/getVirtualPhysicalConnections:getVirtualPhysicalConnections', __args__, opts=opts, typ=GetVirtualPhysicalConnectionsResult).value

    return AwaitableGetVirtualPhysicalConnectionsResult(
        business_status=__ret__.business_status,
        connections=__ret__.connections,
        id=__ret__.id,
        ids=__ret__.ids,
        is_confirmed=__ret__.is_confirmed,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        parent_physical_connection_id=__ret__.parent_physical_connection_id,
        virtual_physical_connection_ids=__ret__.virtual_physical_connection_ids,
        virtual_physical_connection_status=__ret__.virtual_physical_connection_status,
        vlan_ids=__ret__.vlan_ids,
        vpconn_ali_uid=__ret__.vpconn_ali_uid)


@_utilities.lift_output_func(get_virtual_physical_connections)
def get_virtual_physical_connections_output(business_status: Optional[pulumi.Input[Optional[str]]] = None,
                                            ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                            is_confirmed: Optional[pulumi.Input[Optional[bool]]] = None,
                                            name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                            output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                            parent_physical_connection_id: Optional[pulumi.Input[Optional[str]]] = None,
                                            virtual_physical_connection_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                            virtual_physical_connection_status: Optional[pulumi.Input[Optional[str]]] = None,
                                            vlan_ids: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                                            vpconn_ali_uid: Optional[pulumi.Input[Optional[str]]] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVirtualPhysicalConnectionsResult]:
    """
    This data source provides Express Connect Virtual Physical Connection available to the user.

    > **NOTE:** Available in 1.196.0+


    :param str business_status: The commercial status of the physical line. Value:
           - **Normal**: activated.
           - **Financialized**: Arrears locked.
           - **SecurityLocked**: locked for security reasons.
    :param Sequence[str] ids: A list of Virtual Physical Connection IDs.
    :param str name_regex: A regex string to filter results by Group Metric Rule name.
    :param str parent_physical_connection_id: The ID of the instance of the physical connection.
    :param Sequence[str] virtual_physical_connection_ids: The ID of the hosted connection. You can specify multiple hosted connection IDs.
    :param str virtual_physical_connection_status: The business status of the shared line. Value:
           - **Confirmed**: The shared line has been Confirmed to receive.
           - **UnConfirmed**: The shared line has not been confirmed to be received.
           - **Deleted**: The shared line has been Deleted.
    :param Sequence[int] vlan_ids: The VLAN ID of the hosted connection. You can specify multiple VLAN IDs.
    :param str vpconn_ali_uid: The ID of the Alibaba Cloud account (primary account) of the owner of the shared line.
    """
    ...