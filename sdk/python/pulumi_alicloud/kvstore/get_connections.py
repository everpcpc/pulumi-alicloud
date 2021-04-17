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
]

@pulumi.output_type
class GetConnectionsResult:
    """
    A collection of values returned by getConnections.
    """
    def __init__(__self__, connections=None, id=None, ids=None, output_file=None):
        if connections and not isinstance(connections, list):
            raise TypeError("Expected argument 'connections' to be a list")
        pulumi.set(__self__, "connections", connections)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, str):
            raise TypeError("Expected argument 'ids' to be a str")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)

    @property
    @pulumi.getter
    def connections(self) -> Sequence['outputs.GetConnectionsConnectionResult']:
        """
        Public network details of the specified resource. contains the following attributes:
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
    def ids(self) -> str:
        """
        A list of KVStore DBInstance ids.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")


class AwaitableGetConnectionsResult(GetConnectionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConnectionsResult(
            connections=self.connections,
            id=self.id,
            ids=self.ids,
            output_file=self.output_file)


def get_connections(ids: Optional[str] = None,
                    output_file: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConnectionsResult:
    """
    This data source can query the public IP of the specified KVStore DBInstance.

    > **NOTE:** Available in v1.101.0+.


    :param str ids: A list of KVStore DBInstance ids, only support one item.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:kvstore/getConnections:getConnections', __args__, opts=opts, typ=GetConnectionsResult).value

    return AwaitableGetConnectionsResult(
        connections=__ret__.connections,
        id=__ret__.id,
        ids=__ret__.ids,
        output_file=__ret__.output_file)
