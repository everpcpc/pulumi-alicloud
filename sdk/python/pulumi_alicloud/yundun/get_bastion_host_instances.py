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
    'GetBastionHostInstancesResult',
    'AwaitableGetBastionHostInstancesResult',
    'get_bastion_host_instances',
    'get_bastion_host_instances_output',
]

@pulumi.output_type
class GetBastionHostInstancesResult:
    """
    A collection of values returned by getBastionHostInstances.
    """
    def __init__(__self__, description_regex=None, descriptions=None, id=None, ids=None, instances=None, output_file=None, tags=None):
        if description_regex and not isinstance(description_regex, str):
            raise TypeError("Expected argument 'description_regex' to be a str")
        pulumi.set(__self__, "description_regex", description_regex)
        if descriptions and not isinstance(descriptions, list):
            raise TypeError("Expected argument 'descriptions' to be a list")
        pulumi.set(__self__, "descriptions", descriptions)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="descriptionRegex")
    def description_regex(self) -> Optional[str]:
        return pulumi.get(self, "description_regex")

    @property
    @pulumi.getter
    def descriptions(self) -> Sequence[str]:
        return pulumi.get(self, "descriptions")

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
    def instances(self) -> Sequence['outputs.GetBastionHostInstancesInstanceResult']:
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")


class AwaitableGetBastionHostInstancesResult(GetBastionHostInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBastionHostInstancesResult(
            description_regex=self.description_regex,
            descriptions=self.descriptions,
            id=self.id,
            ids=self.ids,
            instances=self.instances,
            output_file=self.output_file,
            tags=self.tags)


def get_bastion_host_instances(description_regex: Optional[str] = None,
                               ids: Optional[Sequence[str]] = None,
                               output_file: Optional[str] = None,
                               tags: Optional[Mapping[str, Any]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBastionHostInstancesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['descriptionRegex'] = description_regex
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:yundun/getBastionHostInstances:getBastionHostInstances', __args__, opts=opts, typ=GetBastionHostInstancesResult).value

    return AwaitableGetBastionHostInstancesResult(
        description_regex=__ret__.description_regex,
        descriptions=__ret__.descriptions,
        id=__ret__.id,
        ids=__ret__.ids,
        instances=__ret__.instances,
        output_file=__ret__.output_file,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_bastion_host_instances)
def get_bastion_host_instances_output(description_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                      ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                      output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                      tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBastionHostInstancesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
