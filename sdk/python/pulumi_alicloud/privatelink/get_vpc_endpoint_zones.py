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
    'GetVpcEndpointZonesResult',
    'AwaitableGetVpcEndpointZonesResult',
    'get_vpc_endpoint_zones',
    'get_vpc_endpoint_zones_output',
]

@pulumi.output_type
class GetVpcEndpointZonesResult:
    """
    A collection of values returned by getVpcEndpointZones.
    """
    def __init__(__self__, endpoint_id=None, id=None, ids=None, output_file=None, status=None, zones=None):
        if endpoint_id and not isinstance(endpoint_id, str):
            raise TypeError("Expected argument 'endpoint_id' to be a str")
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> str:
        return pulumi.get(self, "endpoint_id")

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
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def zones(self) -> Sequence['outputs.GetVpcEndpointZonesZoneResult']:
        return pulumi.get(self, "zones")


class AwaitableGetVpcEndpointZonesResult(GetVpcEndpointZonesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcEndpointZonesResult(
            endpoint_id=self.endpoint_id,
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            status=self.status,
            zones=self.zones)


def get_vpc_endpoint_zones(endpoint_id: Optional[str] = None,
                           output_file: Optional[str] = None,
                           status: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcEndpointZonesResult:
    """
    This data source provides the Privatelink Vpc Endpoint Zones of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.111.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.privatelink.get_vpc_endpoint_zones(endpoint_id="ep-gw8boxxxxx")
    pulumi.export("firstPrivatelinkVpcEndpointZoneId", example.zones[0].id)
    ```


    :param str endpoint_id: The ID of the Vpc Endpoint.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The Status of Vpc Endpoint Zone..
    """
    __args__ = dict()
    __args__['endpointId'] = endpoint_id
    __args__['outputFile'] = output_file
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:privatelink/getVpcEndpointZones:getVpcEndpointZones', __args__, opts=opts, typ=GetVpcEndpointZonesResult).value

    return AwaitableGetVpcEndpointZonesResult(
        endpoint_id=__ret__.endpoint_id,
        id=__ret__.id,
        ids=__ret__.ids,
        output_file=__ret__.output_file,
        status=__ret__.status,
        zones=__ret__.zones)


@_utilities.lift_output_func(get_vpc_endpoint_zones)
def get_vpc_endpoint_zones_output(endpoint_id: Optional[pulumi.Input[str]] = None,
                                  output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  status: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcEndpointZonesResult]:
    """
    This data source provides the Privatelink Vpc Endpoint Zones of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.111.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.privatelink.get_vpc_endpoint_zones(endpoint_id="ep-gw8boxxxxx")
    pulumi.export("firstPrivatelinkVpcEndpointZoneId", example.zones[0].id)
    ```


    :param str endpoint_id: The ID of the Vpc Endpoint.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The Status of Vpc Endpoint Zone..
    """
    ...
