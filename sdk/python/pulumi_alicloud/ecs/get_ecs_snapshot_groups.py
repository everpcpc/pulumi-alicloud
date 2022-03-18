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
    'GetEcsSnapshotGroupsResult',
    'AwaitableGetEcsSnapshotGroupsResult',
    'get_ecs_snapshot_groups',
    'get_ecs_snapshot_groups_output',
]

@pulumi.output_type
class GetEcsSnapshotGroupsResult:
    """
    A collection of values returned by getEcsSnapshotGroups.
    """
    def __init__(__self__, groups=None, id=None, ids=None, instance_id=None, name_regex=None, names=None, output_file=None, snapshot_group_name=None, status=None, tags=None):
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if snapshot_group_name and not isinstance(snapshot_group_name, str):
            raise TypeError("Expected argument 'snapshot_group_name' to be a str")
        pulumi.set(__self__, "snapshot_group_name", snapshot_group_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetEcsSnapshotGroupsGroupResult']:
        return pulumi.get(self, "groups")

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
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="snapshotGroupName")
    def snapshot_group_name(self) -> Optional[str]:
        return pulumi.get(self, "snapshot_group_name")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")


class AwaitableGetEcsSnapshotGroupsResult(GetEcsSnapshotGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEcsSnapshotGroupsResult(
            groups=self.groups,
            id=self.id,
            ids=self.ids,
            instance_id=self.instance_id,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            snapshot_group_name=self.snapshot_group_name,
            status=self.status,
            tags=self.tags)


def get_ecs_snapshot_groups(ids: Optional[Sequence[str]] = None,
                            instance_id: Optional[str] = None,
                            name_regex: Optional[str] = None,
                            output_file: Optional[str] = None,
                            snapshot_group_name: Optional[str] = None,
                            status: Optional[str] = None,
                            tags: Optional[Mapping[str, Any]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEcsSnapshotGroupsResult:
    """
    This data source provides the Ecs Snapshot Groups of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.160.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.ecs.get_ecs_snapshot_groups(ids=["example-id"])
    pulumi.export("ecsSnapshotGroupId1", ids.groups[0].id)
    name_regex = alicloud.ecs.get_ecs_snapshot_groups(name_regex="^my-SnapshotGroup")
    pulumi.export("ecsSnapshotGroupId2", name_regex.groups[0].id)
    status = alicloud.ecs.get_ecs_snapshot_groups(status="accomplished")
    pulumi.export("ecsSnapshotGroupId3", status.groups[0].id)
    instance_id = alicloud.ecs.get_ecs_snapshot_groups(instance_id="example-instance_id")
    pulumi.export("ecsSnapshotGroupId4", instance_id.groups[0].id)
    ```


    :param Sequence[str] ids: A list of Snapshot Group IDs.
    :param str instance_id: The ID of the instance.
    :param str name_regex: A regex string to filter results by Snapshot Group name.
    :param str snapshot_group_name: The name of the snapshot-consistent group.
    :param str status: The status of the resource.
    :param Mapping[str, Any] tags: List of label key-value pairs.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['instanceId'] = instance_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['snapshotGroupName'] = snapshot_group_name
    __args__['status'] = status
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:ecs/getEcsSnapshotGroups:getEcsSnapshotGroups', __args__, opts=opts, typ=GetEcsSnapshotGroupsResult).value

    return AwaitableGetEcsSnapshotGroupsResult(
        groups=__ret__.groups,
        id=__ret__.id,
        ids=__ret__.ids,
        instance_id=__ret__.instance_id,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        snapshot_group_name=__ret__.snapshot_group_name,
        status=__ret__.status,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_ecs_snapshot_groups)
def get_ecs_snapshot_groups_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                   instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                   output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                   snapshot_group_name: Optional[pulumi.Input[Optional[str]]] = None,
                                   status: Optional[pulumi.Input[Optional[str]]] = None,
                                   tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEcsSnapshotGroupsResult]:
    """
    This data source provides the Ecs Snapshot Groups of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.160.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.ecs.get_ecs_snapshot_groups(ids=["example-id"])
    pulumi.export("ecsSnapshotGroupId1", ids.groups[0].id)
    name_regex = alicloud.ecs.get_ecs_snapshot_groups(name_regex="^my-SnapshotGroup")
    pulumi.export("ecsSnapshotGroupId2", name_regex.groups[0].id)
    status = alicloud.ecs.get_ecs_snapshot_groups(status="accomplished")
    pulumi.export("ecsSnapshotGroupId3", status.groups[0].id)
    instance_id = alicloud.ecs.get_ecs_snapshot_groups(instance_id="example-instance_id")
    pulumi.export("ecsSnapshotGroupId4", instance_id.groups[0].id)
    ```


    :param Sequence[str] ids: A list of Snapshot Group IDs.
    :param str instance_id: The ID of the instance.
    :param str name_regex: A regex string to filter results by Snapshot Group name.
    :param str snapshot_group_name: The name of the snapshot-consistent group.
    :param str status: The status of the resource.
    :param Mapping[str, Any] tags: List of label key-value pairs.
    """
    ...
