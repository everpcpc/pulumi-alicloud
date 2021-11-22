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
    'GetConnectionsResult',
    'AwaitableGetConnectionsResult',
    'get_connections',
    'get_connections_output',
]

@pulumi.output_type
class GetConnectionsResult:
    """
    A collection of values returned by getConnections.
    """
    def __init__(__self__, connections=None, customer_gateway_id=None, id=None, ids=None, name_regex=None, names=None, output_file=None, vpn_gateway_id=None):
        if connections and not isinstance(connections, list):
            raise TypeError("Expected argument 'connections' to be a list")
        pulumi.set(__self__, "connections", connections)
        if customer_gateway_id and not isinstance(customer_gateway_id, str):
            raise TypeError("Expected argument 'customer_gateway_id' to be a str")
        pulumi.set(__self__, "customer_gateway_id", customer_gateway_id)
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
        if vpn_gateway_id and not isinstance(vpn_gateway_id, str):
            raise TypeError("Expected argument 'vpn_gateway_id' to be a str")
        pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)

    @property
    @pulumi.getter
    def connections(self) -> Sequence['outputs.GetConnectionsConnectionResult']:
        """
        A list of VPN connections. Each element contains the following attributes:
        """
        return pulumi.get(self, "connections")

    @property
    @pulumi.getter(name="customerGatewayId")
    def customer_gateway_id(self) -> Optional[str]:
        """
        ID of the VPN customer gateway.
        """
        return pulumi.get(self, "customer_gateway_id")

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
        (Optional) IDs of the VPN connections.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        (Optional) names of the VPN connections.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> Optional[str]:
        """
        ID of the VPN gateway.
        """
        return pulumi.get(self, "vpn_gateway_id")


class AwaitableGetConnectionsResult(GetConnectionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConnectionsResult(
            connections=self.connections,
            customer_gateway_id=self.customer_gateway_id,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            vpn_gateway_id=self.vpn_gateway_id)


def get_connections(customer_gateway_id: Optional[str] = None,
                    ids: Optional[Sequence[str]] = None,
                    name_regex: Optional[str] = None,
                    output_file: Optional[str] = None,
                    vpn_gateway_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConnectionsResult:
    """
    The VPN connections data source lists lots of VPN connections resource information owned by an Alicloud account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    foo = alicloud.vpn.get_connections(customer_gateway_id="fake-cgw-id",
        ids=["fake-conn-id"],
        output_file="/tmp/vpnconn",
        vpn_gateway_id="fake-vpn-id")
    ```


    :param str customer_gateway_id: Use the VPN customer gateway ID as the search key.
    :param Sequence[str] ids: IDs of the VPN connections.
    :param str name_regex: A regex string of VPN connection name.
    :param str output_file: Save the result to the file.
    :param str vpn_gateway_id: Use the VPN gateway ID as the search key.
    """
    __args__ = dict()
    __args__['customerGatewayId'] = customer_gateway_id
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['vpnGatewayId'] = vpn_gateway_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:vpn/getConnections:getConnections', __args__, opts=opts, typ=GetConnectionsResult).value

    return AwaitableGetConnectionsResult(
        connections=__ret__.connections,
        customer_gateway_id=__ret__.customer_gateway_id,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        vpn_gateway_id=__ret__.vpn_gateway_id)


@_utilities.lift_output_func(get_connections)
def get_connections_output(customer_gateway_id: Optional[pulumi.Input[Optional[str]]] = None,
                           ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                           output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           vpn_gateway_id: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConnectionsResult]:
    """
    The VPN connections data source lists lots of VPN connections resource information owned by an Alicloud account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    foo = alicloud.vpn.get_connections(customer_gateway_id="fake-cgw-id",
        ids=["fake-conn-id"],
        output_file="/tmp/vpnconn",
        vpn_gateway_id="fake-vpn-id")
    ```


    :param str customer_gateway_id: Use the VPN customer gateway ID as the search key.
    :param Sequence[str] ids: IDs of the VPN connections.
    :param str name_regex: A regex string of VPN connection name.
    :param str output_file: Save the result to the file.
    :param str vpn_gateway_id: Use the VPN gateway ID as the search key.
    """
    ...
