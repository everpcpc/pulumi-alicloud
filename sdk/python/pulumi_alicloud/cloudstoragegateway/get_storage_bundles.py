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
    'GetStorageBundlesResult',
    'AwaitableGetStorageBundlesResult',
    'get_storage_bundles',
    'get_storage_bundles_output',
]

@pulumi.output_type
class GetStorageBundlesResult:
    """
    A collection of values returned by getStorageBundles.
    """
    def __init__(__self__, backend_bucket_region_id=None, bundles=None, id=None, ids=None, name_regex=None, names=None, output_file=None):
        if backend_bucket_region_id and not isinstance(backend_bucket_region_id, str):
            raise TypeError("Expected argument 'backend_bucket_region_id' to be a str")
        pulumi.set(__self__, "backend_bucket_region_id", backend_bucket_region_id)
        if bundles and not isinstance(bundles, list):
            raise TypeError("Expected argument 'bundles' to be a list")
        pulumi.set(__self__, "bundles", bundles)
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

    @property
    @pulumi.getter(name="backendBucketRegionId")
    def backend_bucket_region_id(self) -> str:
        return pulumi.get(self, "backend_bucket_region_id")

    @property
    @pulumi.getter
    def bundles(self) -> Sequence['outputs.GetStorageBundlesBundleResult']:
        return pulumi.get(self, "bundles")

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


class AwaitableGetStorageBundlesResult(GetStorageBundlesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStorageBundlesResult(
            backend_bucket_region_id=self.backend_bucket_region_id,
            bundles=self.bundles,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file)


def get_storage_bundles(backend_bucket_region_id: Optional[str] = None,
                        ids: Optional[Sequence[str]] = None,
                        name_regex: Optional[str] = None,
                        output_file: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStorageBundlesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['backendBucketRegionId'] = backend_bucket_region_id
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:cloudstoragegateway/getStorageBundles:getStorageBundles', __args__, opts=opts, typ=GetStorageBundlesResult).value

    return AwaitableGetStorageBundlesResult(
        backend_bucket_region_id=__ret__.backend_bucket_region_id,
        bundles=__ret__.bundles,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file)


@_utilities.lift_output_func(get_storage_bundles)
def get_storage_bundles_output(backend_bucket_region_id: Optional[pulumi.Input[str]] = None,
                               ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                               name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                               output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetStorageBundlesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
