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
    'GetEcsBackupClientsResult',
    'AwaitableGetEcsBackupClientsResult',
    'get_ecs_backup_clients',
    'get_ecs_backup_clients_output',
]

@pulumi.output_type
class GetEcsBackupClientsResult:
    """
    A collection of values returned by getEcsBackupClients.
    """
    def __init__(__self__, clients=None, id=None, ids=None, instance_ids=None, output_file=None, status=None):
        if clients and not isinstance(clients, list):
            raise TypeError("Expected argument 'clients' to be a list")
        pulumi.set(__self__, "clients", clients)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_ids and not isinstance(instance_ids, list):
            raise TypeError("Expected argument 'instance_ids' to be a list")
        pulumi.set(__self__, "instance_ids", instance_ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def clients(self) -> Sequence['outputs.GetEcsBackupClientsClientResult']:
        return pulumi.get(self, "clients")

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
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "instance_ids")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


class AwaitableGetEcsBackupClientsResult(GetEcsBackupClientsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEcsBackupClientsResult(
            clients=self.clients,
            id=self.id,
            ids=self.ids,
            instance_ids=self.instance_ids,
            output_file=self.output_file,
            status=self.status)


def get_ecs_backup_clients(ids: Optional[Sequence[str]] = None,
                           instance_ids: Optional[Sequence[str]] = None,
                           output_file: Optional[str] = None,
                           status: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEcsBackupClientsResult:
    """
    This data source provides the Hbr Ecs File Backup Clients of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.132.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.ecs.get_instances(name_regex="ecs_instance_name",
        status="Running")
    ids = alicloud.hbr.get_ecs_backup_clients(ids=[alicloud_hbr_ecs_backup_client["default"]["id"]],
        instance_ids=[alicloud_hbr_ecs_backup_client["default"]["instance_id"]])
    pulumi.export("hbrEcsBackupClientId1", ids.clients[0].id)
    ```


    :param Sequence[str] ids: A list of Ecs Backup Client IDs.
    :param Sequence[str] instance_ids: A list of ECS Instance IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the resource.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['instanceIds'] = instance_ids
    __args__['outputFile'] = output_file
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:hbr/getEcsBackupClients:getEcsBackupClients', __args__, opts=opts, typ=GetEcsBackupClientsResult).value

    return AwaitableGetEcsBackupClientsResult(
        clients=__ret__.clients,
        id=__ret__.id,
        ids=__ret__.ids,
        instance_ids=__ret__.instance_ids,
        output_file=__ret__.output_file,
        status=__ret__.status)


@_utilities.lift_output_func(get_ecs_backup_clients)
def get_ecs_backup_clients_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                  instance_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                  output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  status: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEcsBackupClientsResult]:
    """
    This data source provides the Hbr Ecs File Backup Clients of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.132.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.ecs.get_instances(name_regex="ecs_instance_name",
        status="Running")
    ids = alicloud.hbr.get_ecs_backup_clients(ids=[alicloud_hbr_ecs_backup_client["default"]["id"]],
        instance_ids=[alicloud_hbr_ecs_backup_client["default"]["instance_id"]])
    pulumi.export("hbrEcsBackupClientId1", ids.clients[0].id)
    ```


    :param Sequence[str] ids: A list of Ecs Backup Client IDs.
    :param Sequence[str] instance_ids: A list of ECS Instance IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the resource.
    """
    ...
