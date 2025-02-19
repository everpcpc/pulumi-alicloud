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
    'GetGatewaysResult',
    'AwaitableGetGatewaysResult',
    'get_gateways',
    'get_gateways_output',
]

@pulumi.output_type
class GetGatewaysResult:
    """
    A collection of values returned by getGateways.
    """
    def __init__(__self__, gateways=None, id=None, ids=None, name_regex=None, names=None, output_file=None, page_number=None, page_size=None, status=None, storage_bundle_id=None, total_count=None):
        if gateways and not isinstance(gateways, list):
            raise TypeError("Expected argument 'gateways' to be a list")
        pulumi.set(__self__, "gateways", gateways)
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
        if page_number and not isinstance(page_number, int):
            raise TypeError("Expected argument 'page_number' to be a int")
        pulumi.set(__self__, "page_number", page_number)
        if page_size and not isinstance(page_size, int):
            raise TypeError("Expected argument 'page_size' to be a int")
        pulumi.set(__self__, "page_size", page_size)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if storage_bundle_id and not isinstance(storage_bundle_id, str):
            raise TypeError("Expected argument 'storage_bundle_id' to be a str")
        pulumi.set(__self__, "storage_bundle_id", storage_bundle_id)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter
    def gateways(self) -> Sequence['outputs.GetGatewaysGatewayResult']:
        return pulumi.get(self, "gateways")

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
    @pulumi.getter(name="pageNumber")
    def page_number(self) -> Optional[int]:
        return pulumi.get(self, "page_number")

    @property
    @pulumi.getter(name="pageSize")
    def page_size(self) -> Optional[int]:
        return pulumi.get(self, "page_size")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageBundleId")
    def storage_bundle_id(self) -> str:
        return pulumi.get(self, "storage_bundle_id")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        return pulumi.get(self, "total_count")


class AwaitableGetGatewaysResult(GetGatewaysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewaysResult(
            gateways=self.gateways,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            page_number=self.page_number,
            page_size=self.page_size,
            status=self.status,
            storage_bundle_id=self.storage_bundle_id,
            total_count=self.total_count)


def get_gateways(ids: Optional[Sequence[str]] = None,
                 name_regex: Optional[str] = None,
                 output_file: Optional[str] = None,
                 page_number: Optional[int] = None,
                 page_size: Optional[int] = None,
                 status: Optional[str] = None,
                 storage_bundle_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewaysResult:
    """
    This data source provides the Cloud Storage Gateway Gateways of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.132.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.cloudstoragegateway.StorageBundle("example", storage_bundle_name="example_value")
    name_regex = alicloud.cloudstoragegateway.get_gateways_output(storage_bundle_id=example.id,
        name_regex="^my-Gateway")
    pulumi.export("cloudStorageGatewayGatewayId", name_regex.gateways[0].id)
    ```


    :param Sequence[str] ids: A list of Gateway IDs.
    :param str name_regex: A regex string to filter results by Gateway name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: gateway status.
    :param str storage_bundle_id: storage bundle id.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['pageNumber'] = page_number
    __args__['pageSize'] = page_size
    __args__['status'] = status
    __args__['storageBundleId'] = storage_bundle_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cloudstoragegateway/getGateways:getGateways', __args__, opts=opts, typ=GetGatewaysResult).value

    return AwaitableGetGatewaysResult(
        gateways=__ret__.gateways,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        page_number=__ret__.page_number,
        page_size=__ret__.page_size,
        status=__ret__.status,
        storage_bundle_id=__ret__.storage_bundle_id,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(get_gateways)
def get_gateways_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                        output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        page_number: Optional[pulumi.Input[Optional[int]]] = None,
                        page_size: Optional[pulumi.Input[Optional[int]]] = None,
                        status: Optional[pulumi.Input[Optional[str]]] = None,
                        storage_bundle_id: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGatewaysResult]:
    """
    This data source provides the Cloud Storage Gateway Gateways of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.132.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.cloudstoragegateway.StorageBundle("example", storage_bundle_name="example_value")
    name_regex = alicloud.cloudstoragegateway.get_gateways_output(storage_bundle_id=example.id,
        name_regex="^my-Gateway")
    pulumi.export("cloudStorageGatewayGatewayId", name_regex.gateways[0].id)
    ```


    :param Sequence[str] ids: A list of Gateway IDs.
    :param str name_regex: A regex string to filter results by Gateway name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: gateway status.
    :param str storage_bundle_id: storage bundle id.
    """
    ...
