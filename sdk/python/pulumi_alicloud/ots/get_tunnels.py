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
    'GetTunnelsResult',
    'AwaitableGetTunnelsResult',
    'get_tunnels',
    'get_tunnels_output',
]

@pulumi.output_type
class GetTunnelsResult:
    """
    A collection of values returned by getTunnels.
    """
    def __init__(__self__, id=None, ids=None, instance_name=None, name_regex=None, names=None, output_file=None, table_name=None, tunnels=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_name and not isinstance(instance_name, str):
            raise TypeError("Expected argument 'instance_name' to be a str")
        pulumi.set(__self__, "instance_name", instance_name)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if table_name and not isinstance(table_name, str):
            raise TypeError("Expected argument 'table_name' to be a str")
        pulumi.set(__self__, "table_name", table_name)
        if tunnels and not isinstance(tunnels, list):
            raise TypeError("Expected argument 'tunnels' to be a list")
        pulumi.set(__self__, "tunnels", tunnels)

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
        A list of tunnel IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        """
        The OTS instance name.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of tunnel names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> str:
        """
        The table name of the OTS which could not be changed.
        """
        return pulumi.get(self, "table_name")

    @property
    @pulumi.getter
    def tunnels(self) -> Sequence['outputs.GetTunnelsTunnelResult']:
        """
        A list of tunnels. Each element contains the following attributes:
        """
        return pulumi.get(self, "tunnels")


class AwaitableGetTunnelsResult(GetTunnelsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTunnelsResult(
            id=self.id,
            ids=self.ids,
            instance_name=self.instance_name,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            table_name=self.table_name,
            tunnels=self.tunnels)


def get_tunnels(ids: Optional[Sequence[str]] = None,
                instance_name: Optional[str] = None,
                name_regex: Optional[str] = None,
                output_file: Optional[str] = None,
                table_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTunnelsResult:
    """
    This data source provides the ots tunnels of the current Alibaba Cloud user.

    For information about OTS tunnel and how to use it, see [Tunnel overview](https://www.alibabacloud.com/help/en/tablestore/latest/tunnel-service-overview).

    > **NOTE:** Available in v1.172.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    tunnels_ds = alicloud.ots.get_tunnels(instance_name="sample-instance",
        name_regex="sample-tunnel",
        output_file="tunnels.txt",
        table_name="sample-table")
    pulumi.export("firstTunnelId", tunnels_ds.tunnels[0].id)
    ```


    :param Sequence[str] ids: A list of tunnel IDs.
    :param str instance_name: The name of OTS instance.
    :param str name_regex: A regex string to filter results by tunnel name.
    :param str table_name: The name of OTS table.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['instanceName'] = instance_name
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['tableName'] = table_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:ots/getTunnels:getTunnels', __args__, opts=opts, typ=GetTunnelsResult).value

    return AwaitableGetTunnelsResult(
        id=__ret__.id,
        ids=__ret__.ids,
        instance_name=__ret__.instance_name,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        table_name=__ret__.table_name,
        tunnels=__ret__.tunnels)


@_utilities.lift_output_func(get_tunnels)
def get_tunnels_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                       instance_name: Optional[pulumi.Input[str]] = None,
                       name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                       output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       table_name: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTunnelsResult]:
    """
    This data source provides the ots tunnels of the current Alibaba Cloud user.

    For information about OTS tunnel and how to use it, see [Tunnel overview](https://www.alibabacloud.com/help/en/tablestore/latest/tunnel-service-overview).

    > **NOTE:** Available in v1.172.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    tunnels_ds = alicloud.ots.get_tunnels(instance_name="sample-instance",
        name_regex="sample-tunnel",
        output_file="tunnels.txt",
        table_name="sample-table")
    pulumi.export("firstTunnelId", tunnels_ds.tunnels[0].id)
    ```


    :param Sequence[str] ids: A list of tunnel IDs.
    :param str instance_name: The name of OTS instance.
    :param str name_regex: A regex string to filter results by tunnel name.
    :param str table_name: The name of OTS table.
    """
    ...
