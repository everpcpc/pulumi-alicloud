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
    'GetResourceDirectoriesResult',
    'AwaitableGetResourceDirectoriesResult',
    'get_resource_directories',
]

@pulumi.output_type
class GetResourceDirectoriesResult:
    """
    A collection of values returned by getResourceDirectories.
    """
    def __init__(__self__, directories=None, id=None, output_file=None):
        if directories and not isinstance(directories, list):
            raise TypeError("Expected argument 'directories' to be a list")
        pulumi.set(__self__, "directories", directories)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)

    @property
    @pulumi.getter
    def directories(self) -> Sequence['outputs.GetResourceDirectoriesDirectoryResult']:
        """
        A list of resource directories. Each element contains the following attributes:
        """
        return pulumi.get(self, "directories")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")


class AwaitableGetResourceDirectoriesResult(GetResourceDirectoriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourceDirectoriesResult(
            directories=self.directories,
            id=self.id,
            output_file=self.output_file)


def get_resource_directories(output_file: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResourceDirectoriesResult:
    """
    This data source provides the Resource Manager Resource Directories of the current Alibaba Cloud user.

    > **NOTE:**  Available in 1.86.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.resourcemanager.get_resource_directories()
    pulumi.export("resourceDirectoryId", default.directories[0].id)
    ```
    """
    __args__ = dict()
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:resourcemanager/getResourceDirectories:getResourceDirectories', __args__, opts=opts, typ=GetResourceDirectoriesResult).value

    return AwaitableGetResourceDirectoriesResult(
        directories=__ret__.directories,
        id=__ret__.id,
        output_file=__ret__.output_file)
