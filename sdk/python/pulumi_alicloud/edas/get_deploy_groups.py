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
    'GetDeployGroupsResult',
    'AwaitableGetDeployGroupsResult',
    'get_deploy_groups',
    'get_deploy_groups_output',
]

@pulumi.output_type
class GetDeployGroupsResult:
    """
    A collection of values returned by getDeployGroups.
    """
    def __init__(__self__, app_id=None, groups=None, id=None, name_regex=None, names=None, output_file=None):
        if app_id and not isinstance(app_id, str):
            raise TypeError("Expected argument 'app_id' to be a str")
        pulumi.set(__self__, "app_id", app_id)
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> str:
        """
        The ID of the application that you want to deploy.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetDeployGroupsGroupResult']:
        """
        A list of consumer group ids.
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
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of deploy group names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")


class AwaitableGetDeployGroupsResult(GetDeployGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeployGroupsResult(
            app_id=self.app_id,
            groups=self.groups,
            id=self.id,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file)


def get_deploy_groups(app_id: Optional[str] = None,
                      name_regex: Optional[str] = None,
                      output_file: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeployGroupsResult:
    """
    This data source provides a list of EDAS deploy groups in an Alibaba Cloud account according to the specified filters.

    > **NOTE:** Available in 1.82.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    groups = alicloud.edas.get_deploy_groups(app_id="xxx",
        ids=["xxx"],
        output_file="groups.txt")
    pulumi.export("firstGroupName", groups.groups[0].group_name)
    ```


    :param str app_id: ID of the EDAS application.
    :param str name_regex: A regex string to filter results by the deploy group name.
    """
    __args__ = dict()
    __args__['appId'] = app_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:edas/getDeployGroups:getDeployGroups', __args__, opts=opts, typ=GetDeployGroupsResult).value

    return AwaitableGetDeployGroupsResult(
        app_id=__ret__.app_id,
        groups=__ret__.groups,
        id=__ret__.id,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file)


@_utilities.lift_output_func(get_deploy_groups)
def get_deploy_groups_output(app_id: Optional[pulumi.Input[str]] = None,
                             name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                             output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeployGroupsResult]:
    """
    This data source provides a list of EDAS deploy groups in an Alibaba Cloud account according to the specified filters.

    > **NOTE:** Available in 1.82.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    groups = alicloud.edas.get_deploy_groups(app_id="xxx",
        ids=["xxx"],
        output_file="groups.txt")
    pulumi.export("firstGroupName", groups.groups[0].group_name)
    ```


    :param str app_id: ID of the EDAS application.
    :param str name_regex: A regex string to filter results by the deploy group name.
    """
    ...
