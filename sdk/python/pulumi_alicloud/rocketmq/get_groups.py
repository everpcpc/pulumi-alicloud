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
    'GetGroupsResult',
    'AwaitableGetGroupsResult',
    'get_groups',
]

@pulumi.output_type
class GetGroupsResult:
    """
    A collection of values returned by getGroups.
    """
    def __init__(__self__, group_id_regex=None, group_type=None, groups=None, id=None, ids=None, instance_id=None, name_regex=None, names=None, output_file=None, tags=None):
        if group_id_regex and not isinstance(group_id_regex, str):
            raise TypeError("Expected argument 'group_id_regex' to be a str")
        pulumi.set(__self__, "group_id_regex", group_id_regex)
        if group_type and not isinstance(group_type, str):
            raise TypeError("Expected argument 'group_type' to be a str")
        pulumi.set(__self__, "group_type", group_type)
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
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="groupIdRegex")
    def group_id_regex(self) -> Optional[str]:
        return pulumi.get(self, "group_id_regex")

    @property
    @pulumi.getter(name="groupType")
    def group_type(self) -> Optional[str]:
        """
        Specify the protocol applicable to the created Group ID.
        """
        return pulumi.get(self, "group_type")

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetGroupsGroupResult']:
        """
        A list of groups. Each element contains the following attributes:
        """
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
        """
        A list of group names.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
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
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        """
        A map of tags assigned to the Ons group.
        """
        return pulumi.get(self, "tags")


class AwaitableGetGroupsResult(GetGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupsResult(
            group_id_regex=self.group_id_regex,
            group_type=self.group_type,
            groups=self.groups,
            id=self.id,
            ids=self.ids,
            instance_id=self.instance_id,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            tags=self.tags)


def get_groups(group_id_regex: Optional[str] = None,
               group_type: Optional[str] = None,
               ids: Optional[Sequence[str]] = None,
               instance_id: Optional[str] = None,
               name_regex: Optional[str] = None,
               output_file: Optional[str] = None,
               tags: Optional[Mapping[str, Any]] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupsResult:
    """
    This data source provides a list of ONS Groups in an Alibaba Cloud account according to the specified filters.

    > **NOTE:** Available in 1.53.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "onsInstanceName"
    group_name = config.get("groupName")
    if group_name is None:
        group_name = "GID-onsGroupDatasourceName"
    default_instance = alicloud.rocketmq.Instance("defaultInstance",
        instance_name=name,
        remark="default_ons_instance_remark")
    default_group = alicloud.rocketmq.Group("defaultGroup",
        group_name=group_name,
        instance_id=default_instance.id,
        remark="dafault_ons_group_remark")
    groups_ds = default_group.instance_id.apply(lambda instance_id: alicloud.rocketmq.get_groups(instance_id=instance_id,
        name_regex=var["group_id"],
        output_file="groups.txt"))
    pulumi.export("firstGroupName", groups_ds.groups[0].group_name)
    ```


    :param str group_id_regex: A regex string to filter results by the group name.
    :param str group_type: Specify the protocol applicable to the created Group ID. Valid values: `tcp`, `http`. Default to `tcp`.
    :param Sequence[str] ids: A list of group names.
    :param str instance_id: ID of the ONS Instance that owns the groups.
    :param Mapping[str, Any] tags: A map of tags assigned to the Ons instance.
    """
    __args__ = dict()
    __args__['groupIdRegex'] = group_id_regex
    __args__['groupType'] = group_type
    __args__['ids'] = ids
    __args__['instanceId'] = instance_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:rocketmq/getGroups:getGroups', __args__, opts=opts, typ=GetGroupsResult).value

    return AwaitableGetGroupsResult(
        group_id_regex=__ret__.group_id_regex,
        group_type=__ret__.group_type,
        groups=__ret__.groups,
        id=__ret__.id,
        ids=__ret__.ids,
        instance_id=__ret__.instance_id,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        tags=__ret__.tags)
