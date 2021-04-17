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
    'GetClustersResult',
    'AwaitableGetClustersResult',
    'get_clusters',
]

@pulumi.output_type
class GetClustersResult:
    """
    A collection of values returned by getClusters.
    """
    def __init__(__self__, clusters=None, description=None, description_regex=None, descriptions=None, enable_details=None, id=None, ids=None, output_file=None, resource_group_id=None, status=None, tags=None):
        if clusters and not isinstance(clusters, list):
            raise TypeError("Expected argument 'clusters' to be a list")
        pulumi.set(__self__, "clusters", clusters)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if description_regex and not isinstance(description_regex, str):
            raise TypeError("Expected argument 'description_regex' to be a str")
        pulumi.set(__self__, "description_regex", description_regex)
        if descriptions and not isinstance(descriptions, list):
            raise TypeError("Expected argument 'descriptions' to be a list")
        pulumi.set(__self__, "descriptions", descriptions)
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def clusters(self) -> Sequence['outputs.GetClustersClusterResult']:
        """
        A list of ADB clusters. Each element contains the following attributes:
        """
        return pulumi.get(self, "clusters")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the ADB cluster.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="descriptionRegex")
    def description_regex(self) -> Optional[str]:
        return pulumi.get(self, "description_regex")

    @property
    @pulumi.getter
    def descriptions(self) -> Sequence[str]:
        """
        A list of ADB cluster descriptions.
        """
        return pulumi.get(self, "descriptions")

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
        """
        A list of ADB cluster IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Status of the cluster.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")


class AwaitableGetClustersResult(GetClustersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClustersResult(
            clusters=self.clusters,
            description=self.description,
            description_regex=self.description_regex,
            descriptions=self.descriptions,
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            resource_group_id=self.resource_group_id,
            status=self.status,
            tags=self.tags)


def get_clusters(description: Optional[str] = None,
                 description_regex: Optional[str] = None,
                 enable_details: Optional[bool] = None,
                 ids: Optional[Sequence[str]] = None,
                 output_file: Optional[str] = None,
                 resource_group_id: Optional[str] = None,
                 status: Optional[str] = None,
                 tags: Optional[Mapping[str, Any]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClustersResult:
    """
    Use this data source to access information about an existing resource.

    :param str description: The description of the ADB cluster.
    :param str description_regex: A regex string to filter results by cluster description.
    :param Sequence[str] ids: A list of ADB cluster IDs.
    :param str status: The status of the cluster. Valid values: `Preparing`, `Creating`, `Restoring`, `Running`, `Deleting`, `ClassChanging`, `NetAddressCreating`, `NetAddressDeleting`. For more information, see [Cluster status](https://www.alibabacloud.com/help/doc-detail/143075.htm).
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
           - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
           - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['descriptionRegex'] = description_regex
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['resourceGroupId'] = resource_group_id
    __args__['status'] = status
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:adb/getClusters:getClusters', __args__, opts=opts, typ=GetClustersResult).value

    return AwaitableGetClustersResult(
        clusters=__ret__.clusters,
        description=__ret__.description,
        description_regex=__ret__.description_regex,
        descriptions=__ret__.descriptions,
        enable_details=__ret__.enable_details,
        id=__ret__.id,
        ids=__ret__.ids,
        output_file=__ret__.output_file,
        resource_group_id=__ret__.resource_group_id,
        status=__ret__.status,
        tags=__ret__.tags)
