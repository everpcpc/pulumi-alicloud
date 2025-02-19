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
    'GetServiceMeshesResult',
    'AwaitableGetServiceMeshesResult',
    'get_service_meshes',
    'get_service_meshes_output',
]

@pulumi.output_type
class GetServiceMeshesResult:
    """
    A collection of values returned by getServiceMeshes.
    """
    def __init__(__self__, enable_details=None, id=None, ids=None, meshes=None, name_regex=None, names=None, output_file=None, status=None):
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if meshes and not isinstance(meshes, list):
            raise TypeError("Expected argument 'meshes' to be a list")
        pulumi.set(__self__, "meshes", meshes)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

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
    def meshes(self) -> Sequence['outputs.GetServiceMeshesMeshResult']:
        """
        A list of Service Mesh Service Meshes. Each element contains the following attributes:
        """
        return pulumi.get(self, "meshes")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of Service Mesh names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")


class AwaitableGetServiceMeshesResult(GetServiceMeshesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceMeshesResult(
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            meshes=self.meshes,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            status=self.status)


def get_service_meshes(enable_details: Optional[bool] = None,
                       ids: Optional[Sequence[str]] = None,
                       name_regex: Optional[str] = None,
                       output_file: Optional[str] = None,
                       status: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceMeshesResult:
    """
    This data source provides the Service Mesh Service Meshes of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.138.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.servicemesh.get_service_meshes(ids=["example_id"])
    pulumi.export("serviceMeshServiceMeshId1", ids.meshes[0].id)
    name_regex = alicloud.servicemesh.get_service_meshes(name_regex="^my-ServiceMesh")
    pulumi.export("serviceMeshServiceMeshId2", name_regex.meshes[0].id)
    status = alicloud.servicemesh.get_service_meshes(ids=["example_id"],
        status="running")
    pulumi.export("serviceMeshServiceMeshId3", status.meshes[0].id)
    ```


    :param Sequence[str] ids: A list of Service Mesh IDs.
    :param str name_regex: A regex string to filter results by Service Mesh name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the resource. Valid values: `running` or `initial`.
    """
    __args__ = dict()
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:servicemesh/getServiceMeshes:getServiceMeshes', __args__, opts=opts, typ=GetServiceMeshesResult).value

    return AwaitableGetServiceMeshesResult(
        enable_details=__ret__.enable_details,
        id=__ret__.id,
        ids=__ret__.ids,
        meshes=__ret__.meshes,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        status=__ret__.status)


@_utilities.lift_output_func(get_service_meshes)
def get_service_meshes_output(enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                              ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                              output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              status: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceMeshesResult]:
    """
    This data source provides the Service Mesh Service Meshes of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.138.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.servicemesh.get_service_meshes(ids=["example_id"])
    pulumi.export("serviceMeshServiceMeshId1", ids.meshes[0].id)
    name_regex = alicloud.servicemesh.get_service_meshes(name_regex="^my-ServiceMesh")
    pulumi.export("serviceMeshServiceMeshId2", name_regex.meshes[0].id)
    status = alicloud.servicemesh.get_service_meshes(ids=["example_id"],
        status="running")
    pulumi.export("serviceMeshServiceMeshId3", status.meshes[0].id)
    ```


    :param Sequence[str] ids: A list of Service Mesh IDs.
    :param str name_regex: A regex string to filter results by Service Mesh name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the resource. Valid values: `running` or `initial`.
    """
    ...
