# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetBackupJobsResult',
    'AwaitableGetBackupJobsResult',
    'get_backup_jobs',
    'get_backup_jobs_output',
]

@pulumi.output_type
class GetBackupJobsResult:
    """
    A collection of values returned by getBackupJobs.
    """
    def __init__(__self__, filters=None, id=None, ids=None, jobs=None, output_file=None, sort_direction=None, source_type=None, status=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if jobs and not isinstance(jobs, list):
            raise TypeError("Expected argument 'jobs' to be a list")
        pulumi.set(__self__, "jobs", jobs)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if sort_direction and not isinstance(sort_direction, str):
            raise TypeError("Expected argument 'sort_direction' to be a str")
        pulumi.set(__self__, "sort_direction", sort_direction)
        if source_type and not isinstance(source_type, str):
            raise TypeError("Expected argument 'source_type' to be a str")
        pulumi.set(__self__, "source_type", source_type)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetBackupJobsFilterResult']]:
        return pulumi.get(self, "filters")

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
    @pulumi.getter
    def jobs(self) -> Sequence['outputs.GetBackupJobsJobResult']:
        return pulumi.get(self, "jobs")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="sortDirection")
    def sort_direction(self) -> Optional[str]:
        return pulumi.get(self, "sort_direction")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> str:
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


class AwaitableGetBackupJobsResult(GetBackupJobsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackupJobsResult(
            filters=self.filters,
            id=self.id,
            ids=self.ids,
            jobs=self.jobs,
            output_file=self.output_file,
            sort_direction=self.sort_direction,
            source_type=self.source_type,
            status=self.status)


def get_backup_jobs(filters: Optional[Sequence[pulumi.InputType['GetBackupJobsFilterArgs']]] = None,
                    ids: Optional[Sequence[str]] = None,
                    output_file: Optional[str] = None,
                    sort_direction: Optional[str] = None,
                    source_type: Optional[str] = None,
                    status: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackupJobsResult:
    """
    This data source provides the Hbr Backup Jobs of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.138.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_ecs_backup_plans = alicloud.hbr.get_ecs_backup_plans(name_regex="plan-name")
    default_backup_jobs = alicloud.hbr.get_backup_jobs(source_type="ECS_FILE",
        filters=[
            alicloud.hbr.GetBackupJobsFilterArgs(
                key="VaultId",
                operator="IN",
                values=[default_ecs_backup_plans.plans[0].vault_id],
            ),
            alicloud.hbr.GetBackupJobsFilterArgs(
                key="InstanceId",
                operator="IN",
                values=[default_ecs_backup_plans.plans[0].instance_id],
            ),
            alicloud.hbr.GetBackupJobsFilterArgs(
                key="CompleteTime",
                operator="BETWEEN",
                values=[
                    "2021-08-23T14:17:15CST",
                    "2021-08-24T14:17:15CST",
                ],
            ),
        ])
    example = alicloud.hbr.get_backup_jobs(source_type="ECS_FILE",
        status="COMPLETE",
        filters=[
            alicloud.hbr.GetBackupJobsFilterArgs(
                key="VaultId",
                operator="IN",
                values=[default_ecs_backup_plans.plans[0].vault_id],
            ),
            alicloud.hbr.GetBackupJobsFilterArgs(
                key="InstanceId",
                operator="IN",
                values=[default_ecs_backup_plans.plans[0].instance_id],
            ),
            alicloud.hbr.GetBackupJobsFilterArgs(
                key="CompleteTime",
                operator="LESS_THAN",
                values=["2021-10-20T20:20:20CST"],
            ),
        ])
    pulumi.export("alicloudHbrBackupJobsDefault1", default_backup_jobs.jobs[0].id)
    pulumi.export("alicloudHbrBackupJobsExample1", example.jobs[0].id)
    ```


    :param Sequence[str] ids: A list of Backup Job IDs.
    :param str sort_direction: The sort direction, sort results by ascending or descending order based on the value jobs id. Valid values: `ASCEND`, `DESCEND`.
    :param str source_type: The type of data source. Valid Values: `ECS_FILE`, `OSS`, `NAS`, `UDM_DISK`.
    :param str status: The status of restore job. Valid values: `COMPLETE` , `PARTIAL_COMPLETE`, `FAILED`.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['sortDirection'] = sort_direction
    __args__['sourceType'] = source_type
    __args__['status'] = status
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:hbr/getBackupJobs:getBackupJobs', __args__, opts=opts, typ=GetBackupJobsResult).value

    return AwaitableGetBackupJobsResult(
        filters=__ret__.filters,
        id=__ret__.id,
        ids=__ret__.ids,
        jobs=__ret__.jobs,
        output_file=__ret__.output_file,
        sort_direction=__ret__.sort_direction,
        source_type=__ret__.source_type,
        status=__ret__.status)


@_utilities.lift_output_func(get_backup_jobs)
def get_backup_jobs_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetBackupJobsFilterArgs']]]]] = None,
                           ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           sort_direction: Optional[pulumi.Input[Optional[str]]] = None,
                           source_type: Optional[pulumi.Input[str]] = None,
                           status: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackupJobsResult]:
    """
    This data source provides the Hbr Backup Jobs of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.138.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_ecs_backup_plans = alicloud.hbr.get_ecs_backup_plans(name_regex="plan-name")
    default_backup_jobs = alicloud.hbr.get_backup_jobs(source_type="ECS_FILE",
        filters=[
            alicloud.hbr.GetBackupJobsFilterArgs(
                key="VaultId",
                operator="IN",
                values=[default_ecs_backup_plans.plans[0].vault_id],
            ),
            alicloud.hbr.GetBackupJobsFilterArgs(
                key="InstanceId",
                operator="IN",
                values=[default_ecs_backup_plans.plans[0].instance_id],
            ),
            alicloud.hbr.GetBackupJobsFilterArgs(
                key="CompleteTime",
                operator="BETWEEN",
                values=[
                    "2021-08-23T14:17:15CST",
                    "2021-08-24T14:17:15CST",
                ],
            ),
        ])
    example = alicloud.hbr.get_backup_jobs(source_type="ECS_FILE",
        status="COMPLETE",
        filters=[
            alicloud.hbr.GetBackupJobsFilterArgs(
                key="VaultId",
                operator="IN",
                values=[default_ecs_backup_plans.plans[0].vault_id],
            ),
            alicloud.hbr.GetBackupJobsFilterArgs(
                key="InstanceId",
                operator="IN",
                values=[default_ecs_backup_plans.plans[0].instance_id],
            ),
            alicloud.hbr.GetBackupJobsFilterArgs(
                key="CompleteTime",
                operator="LESS_THAN",
                values=["2021-10-20T20:20:20CST"],
            ),
        ])
    pulumi.export("alicloudHbrBackupJobsDefault1", default_backup_jobs.jobs[0].id)
    pulumi.export("alicloudHbrBackupJobsExample1", example.jobs[0].id)
    ```


    :param Sequence[str] ids: A list of Backup Job IDs.
    :param str sort_direction: The sort direction, sort results by ascending or descending order based on the value jobs id. Valid values: `ASCEND`, `DESCEND`.
    :param str source_type: The type of data source. Valid Values: `ECS_FILE`, `OSS`, `NAS`, `UDM_DISK`.
    :param str status: The status of restore job. Valid values: `COMPLETE` , `PARTIAL_COMPLETE`, `FAILED`.
    """
    ...
