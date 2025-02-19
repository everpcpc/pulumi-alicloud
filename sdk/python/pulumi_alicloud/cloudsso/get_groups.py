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
    'GetGroupsResult',
    'AwaitableGetGroupsResult',
    'get_groups',
    'get_groups_output',
]

@pulumi.output_type
class GetGroupsResult:
    """
    A collection of values returned by getGroups.
    """
    def __init__(__self__, directory_id=None, groups=None, id=None, ids=None, name_regex=None, names=None, output_file=None, provision_type=None):
        if directory_id and not isinstance(directory_id, str):
            raise TypeError("Expected argument 'directory_id' to be a str")
        pulumi.set(__self__, "directory_id", directory_id)
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
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
        if provision_type and not isinstance(provision_type, str):
            raise TypeError("Expected argument 'provision_type' to be a str")
        pulumi.set(__self__, "provision_type", provision_type)

    @property
    @pulumi.getter(name="directoryId")
    def directory_id(self) -> str:
        return pulumi.get(self, "directory_id")

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetGroupsGroupResult']:
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
    @pulumi.getter(name="provisionType")
    def provision_type(self) -> Optional[str]:
        return pulumi.get(self, "provision_type")


class AwaitableGetGroupsResult(GetGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupsResult(
            directory_id=self.directory_id,
            groups=self.groups,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            provision_type=self.provision_type)


def get_groups(directory_id: Optional[str] = None,
               ids: Optional[Sequence[str]] = None,
               name_regex: Optional[str] = None,
               output_file: Optional[str] = None,
               provision_type: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupsResult:
    """
    This data source provides the Cloud Sso Groups of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.138.0+.

    > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cloudsso.get_groups(directory_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("cloudSsoGroupId1", ids.groups[0].id)
    name_regex = alicloud.cloudsso.get_groups(directory_id="example_value",
        name_regex="^my-Group")
    pulumi.export("cloudSsoGroupId2", name_regex.groups[0].id)
    ```


    :param str directory_id: The ID of the Directory.
    :param Sequence[str] ids: A list of Group IDs.
    :param str name_regex: A regex string to filter results by Group name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str provision_type: The Provision Type of the Group. Valid values: `Manual`, `Synchronized`.
    """
    __args__ = dict()
    __args__['directoryId'] = directory_id
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['provisionType'] = provision_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cloudsso/getGroups:getGroups', __args__, opts=opts, typ=GetGroupsResult).value

    return AwaitableGetGroupsResult(
        directory_id=__ret__.directory_id,
        groups=__ret__.groups,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        provision_type=__ret__.provision_type)


@_utilities.lift_output_func(get_groups)
def get_groups_output(directory_id: Optional[pulumi.Input[str]] = None,
                      ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                      name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                      output_file: Optional[pulumi.Input[Optional[str]]] = None,
                      provision_type: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGroupsResult]:
    """
    This data source provides the Cloud Sso Groups of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.138.0+.

    > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cloudsso.get_groups(directory_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("cloudSsoGroupId1", ids.groups[0].id)
    name_regex = alicloud.cloudsso.get_groups(directory_id="example_value",
        name_regex="^my-Group")
    pulumi.export("cloudSsoGroupId2", name_regex.groups[0].id)
    ```


    :param str directory_id: The ID of the Directory.
    :param Sequence[str] ids: A list of Group IDs.
    :param str name_regex: A regex string to filter results by Group name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str provision_type: The Provision Type of the Group. Valid values: `Manual`, `Synchronized`.
    """
    ...
