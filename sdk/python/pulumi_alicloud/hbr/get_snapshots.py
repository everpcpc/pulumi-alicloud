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
    'GetSnapshotsResult',
    'AwaitableGetSnapshotsResult',
    'get_snapshots',
    'get_snapshots_output',
]

@pulumi.output_type
class GetSnapshotsResult:
    """
    A collection of values returned by getSnapshots.
    """
    def __init__(__self__, bucket=None, complete_time=None, complete_time_checker=None, create_time=None, file_system_id=None, id=None, ids=None, instance_id=None, limit=None, output_file=None, query=None, snapshots=None, source_type=None, status=None, vault_id=None):
        if bucket and not isinstance(bucket, str):
            raise TypeError("Expected argument 'bucket' to be a str")
        pulumi.set(__self__, "bucket", bucket)
        if complete_time and not isinstance(complete_time, str):
            raise TypeError("Expected argument 'complete_time' to be a str")
        pulumi.set(__self__, "complete_time", complete_time)
        if complete_time_checker and not isinstance(complete_time_checker, str):
            raise TypeError("Expected argument 'complete_time_checker' to be a str")
        pulumi.set(__self__, "complete_time_checker", complete_time_checker)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if file_system_id and not isinstance(file_system_id, str):
            raise TypeError("Expected argument 'file_system_id' to be a str")
        pulumi.set(__self__, "file_system_id", file_system_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if limit and not isinstance(limit, int):
            raise TypeError("Expected argument 'limit' to be a int")
        pulumi.set(__self__, "limit", limit)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if query and not isinstance(query, str):
            raise TypeError("Expected argument 'query' to be a str")
        pulumi.set(__self__, "query", query)
        if snapshots and not isinstance(snapshots, list):
            raise TypeError("Expected argument 'snapshots' to be a list")
        pulumi.set(__self__, "snapshots", snapshots)
        if source_type and not isinstance(source_type, str):
            raise TypeError("Expected argument 'source_type' to be a str")
        pulumi.set(__self__, "source_type", source_type)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vault_id and not isinstance(vault_id, str):
            raise TypeError("Expected argument 'vault_id' to be a str")
        pulumi.set(__self__, "vault_id", vault_id)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[str]:
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="completeTime")
    def complete_time(self) -> Optional[str]:
        return pulumi.get(self, "complete_time")

    @property
    @pulumi.getter(name="completeTimeChecker")
    def complete_time_checker(self) -> Optional[str]:
        return pulumi.get(self, "complete_time_checker")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[str]:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> Optional[str]:
        return pulumi.get(self, "file_system_id")

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
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def limit(self) -> Optional[int]:
        return pulumi.get(self, "limit")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def query(self) -> Optional[str]:
        return pulumi.get(self, "query")

    @property
    @pulumi.getter
    def snapshots(self) -> Sequence['outputs.GetSnapshotsSnapshotResult']:
        return pulumi.get(self, "snapshots")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> str:
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> str:
        return pulumi.get(self, "vault_id")


class AwaitableGetSnapshotsResult(GetSnapshotsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSnapshotsResult(
            bucket=self.bucket,
            complete_time=self.complete_time,
            complete_time_checker=self.complete_time_checker,
            create_time=self.create_time,
            file_system_id=self.file_system_id,
            id=self.id,
            ids=self.ids,
            instance_id=self.instance_id,
            limit=self.limit,
            output_file=self.output_file,
            query=self.query,
            snapshots=self.snapshots,
            source_type=self.source_type,
            status=self.status,
            vault_id=self.vault_id)


def get_snapshots(bucket: Optional[str] = None,
                  complete_time: Optional[str] = None,
                  complete_time_checker: Optional[str] = None,
                  create_time: Optional[str] = None,
                  file_system_id: Optional[str] = None,
                  ids: Optional[Sequence[str]] = None,
                  instance_id: Optional[str] = None,
                  limit: Optional[int] = None,
                  output_file: Optional[str] = None,
                  query: Optional[str] = None,
                  source_type: Optional[str] = None,
                  status: Optional[str] = None,
                  vault_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSnapshotsResult:
    """
    This data source provides the Hbr Snapshots of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.133.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_ecs_backup_plans = alicloud.hbr.get_ecs_backup_plans(name_regex="plan-tf-used-dont-delete")
    default_oss_backup_plans = alicloud.hbr.get_oss_backup_plans(name_regex="plan-tf-used-dont-delete")
    default_nas_backup_plans = alicloud.hbr.get_nas_backup_plans(name_regex="plan-tf-used-dont-delete")
    ecs_snapshots = alicloud.hbr.get_snapshots(source_type="ECS_FILE",
        vault_id=default_ecs_backup_plans.plans[0].vault_id,
        instance_id=default_ecs_backup_plans.plans[0].instance_id)
    oss_snapshots = alicloud.hbr.get_snapshots(source_type="OSS",
        vault_id=default_oss_backup_plans.plans[0].vault_id,
        bucket=default_oss_backup_plans.plans[0].bucket,
        complete_time="2021-07-20T14:17:15CST,2021-07-24T14:17:15CST",
        complete_time_checker="BETWEEN")
    nas_snapshots = alicloud.hbr.get_snapshots(source_type="NAS",
        vault_id=default_nas_backup_plans.plans[0].vault_id,
        file_system_id=default_nas_backup_plans.plans[0].file_system_id,
        create_time=default_nas_backup_plans.plans[0].create_time,
        complete_time="2021-08-23T14:17:15CST",
        complete_time_checker="GREATER_THAN_OR_EQUAL")
    pulumi.export("hbrSnapshotId1", nas_snapshots.snapshots[0].id)
    ```


    :param str bucket: The name of OSS bucket.
    :param str complete_time: The time when the snapshot completed. UNIX time in seconds.
    :param str complete_time_checker: Complete time filter operator. Optional values: `MATCH_TERM`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `BETWEEN`.
    :param str create_time: File System Creation Time of Nas. Unix Time Seconds.
    :param str file_system_id: The ID of NAS File system.
    :param Sequence[str] ids: A list of Snapshot IDs.
    :param str instance_id: The ID of ECS instance.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str source_type: Data source type, optional values: `ECS_FILE`, `OSS`, `NAS`.
    :param str status: The status of snapshot execution. Possible values: `COMPLETE`, `PARTIAL_COMPLETE`, `FAILED`.
    :param str vault_id: The ID of Vault.
    """
    __args__ = dict()
    __args__['bucket'] = bucket
    __args__['completeTime'] = complete_time
    __args__['completeTimeChecker'] = complete_time_checker
    __args__['createTime'] = create_time
    __args__['fileSystemId'] = file_system_id
    __args__['ids'] = ids
    __args__['instanceId'] = instance_id
    __args__['limit'] = limit
    __args__['outputFile'] = output_file
    __args__['query'] = query
    __args__['sourceType'] = source_type
    __args__['status'] = status
    __args__['vaultId'] = vault_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:hbr/getSnapshots:getSnapshots', __args__, opts=opts, typ=GetSnapshotsResult).value

    return AwaitableGetSnapshotsResult(
        bucket=__ret__.bucket,
        complete_time=__ret__.complete_time,
        complete_time_checker=__ret__.complete_time_checker,
        create_time=__ret__.create_time,
        file_system_id=__ret__.file_system_id,
        id=__ret__.id,
        ids=__ret__.ids,
        instance_id=__ret__.instance_id,
        limit=__ret__.limit,
        output_file=__ret__.output_file,
        query=__ret__.query,
        snapshots=__ret__.snapshots,
        source_type=__ret__.source_type,
        status=__ret__.status,
        vault_id=__ret__.vault_id)


@_utilities.lift_output_func(get_snapshots)
def get_snapshots_output(bucket: Optional[pulumi.Input[Optional[str]]] = None,
                         complete_time: Optional[pulumi.Input[Optional[str]]] = None,
                         complete_time_checker: Optional[pulumi.Input[Optional[str]]] = None,
                         create_time: Optional[pulumi.Input[Optional[str]]] = None,
                         file_system_id: Optional[pulumi.Input[Optional[str]]] = None,
                         ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                         limit: Optional[pulumi.Input[Optional[int]]] = None,
                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         query: Optional[pulumi.Input[Optional[str]]] = None,
                         source_type: Optional[pulumi.Input[str]] = None,
                         status: Optional[pulumi.Input[Optional[str]]] = None,
                         vault_id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSnapshotsResult]:
    """
    This data source provides the Hbr Snapshots of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.133.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_ecs_backup_plans = alicloud.hbr.get_ecs_backup_plans(name_regex="plan-tf-used-dont-delete")
    default_oss_backup_plans = alicloud.hbr.get_oss_backup_plans(name_regex="plan-tf-used-dont-delete")
    default_nas_backup_plans = alicloud.hbr.get_nas_backup_plans(name_regex="plan-tf-used-dont-delete")
    ecs_snapshots = alicloud.hbr.get_snapshots(source_type="ECS_FILE",
        vault_id=default_ecs_backup_plans.plans[0].vault_id,
        instance_id=default_ecs_backup_plans.plans[0].instance_id)
    oss_snapshots = alicloud.hbr.get_snapshots(source_type="OSS",
        vault_id=default_oss_backup_plans.plans[0].vault_id,
        bucket=default_oss_backup_plans.plans[0].bucket,
        complete_time="2021-07-20T14:17:15CST,2021-07-24T14:17:15CST",
        complete_time_checker="BETWEEN")
    nas_snapshots = alicloud.hbr.get_snapshots(source_type="NAS",
        vault_id=default_nas_backup_plans.plans[0].vault_id,
        file_system_id=default_nas_backup_plans.plans[0].file_system_id,
        create_time=default_nas_backup_plans.plans[0].create_time,
        complete_time="2021-08-23T14:17:15CST",
        complete_time_checker="GREATER_THAN_OR_EQUAL")
    pulumi.export("hbrSnapshotId1", nas_snapshots.snapshots[0].id)
    ```


    :param str bucket: The name of OSS bucket.
    :param str complete_time: The time when the snapshot completed. UNIX time in seconds.
    :param str complete_time_checker: Complete time filter operator. Optional values: `MATCH_TERM`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `BETWEEN`.
    :param str create_time: File System Creation Time of Nas. Unix Time Seconds.
    :param str file_system_id: The ID of NAS File system.
    :param Sequence[str] ids: A list of Snapshot IDs.
    :param str instance_id: The ID of ECS instance.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str source_type: Data source type, optional values: `ECS_FILE`, `OSS`, `NAS`.
    :param str status: The status of snapshot execution. Possible values: `COMPLETE`, `PARTIAL_COMPLETE`, `FAILED`.
    :param str vault_id: The ID of Vault.
    """
    ...
