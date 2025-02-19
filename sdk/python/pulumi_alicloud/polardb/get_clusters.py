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
    'GetClustersResult',
    'AwaitableGetClustersResult',
    'get_clusters',
    'get_clusters_output',
]

@pulumi.output_type
class GetClustersResult:
    """
    A collection of values returned by getClusters.
    """
    def __init__(__self__, clusters=None, db_type=None, description_regex=None, descriptions=None, id=None, ids=None, output_file=None, status=None, tags=None):
        if clusters and not isinstance(clusters, list):
            raise TypeError("Expected argument 'clusters' to be a list")
        pulumi.set(__self__, "clusters", clusters)
        if db_type and not isinstance(db_type, str):
            raise TypeError("Expected argument 'db_type' to be a str")
        pulumi.set(__self__, "db_type", db_type)
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
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
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
        A list of PolarDB clusters. Each element contains the following attributes:
        """
        return pulumi.get(self, "clusters")

    @property
    @pulumi.getter(name="dbType")
    def db_type(self) -> Optional[str]:
        """
        `Primary` for primary cluster, `ReadOnly` for read-only cluster, `Guard` for disaster recovery cluster, and `Temp` for temporary cluster.
        """
        return pulumi.get(self, "db_type")

    @property
    @pulumi.getter(name="descriptionRegex")
    def description_regex(self) -> Optional[str]:
        return pulumi.get(self, "description_regex")

    @property
    @pulumi.getter
    def descriptions(self) -> Sequence[str]:
        """
        A list of RDS cluster descriptions.
        """
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
        """
        A list of RDS cluster IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

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
            db_type=self.db_type,
            description_regex=self.description_regex,
            descriptions=self.descriptions,
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            status=self.status,
            tags=self.tags)


def get_clusters(db_type: Optional[str] = None,
                 description_regex: Optional[str] = None,
                 ids: Optional[Sequence[str]] = None,
                 output_file: Optional[str] = None,
                 status: Optional[str] = None,
                 tags: Optional[Mapping[str, Any]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClustersResult:
    """
    The `polardb_get_clusters` data source provides a collection of PolarDB clusters available in Alibaba Cloud account.
    Filters support regular expression for the cluster description, searches by tags, and other filters which are listed below.

    > **NOTE:** Available in v1.66.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    polardb_clusters_ds = alicloud.polardb.get_clusters(description_regex="pc-\\\\w+",
        status="Running")
    pulumi.export("firstPolardbClusterId", polardb_clusters_ds.clusters[0].id)
    ```


    :param str db_type: Database type. Options are `MySQL`, `Oracle` and `PostgreSQL`. If no value is specified, all types are returned.
    :param str description_regex: A regex string to filter results by cluster description.
    :param Sequence[str] ids: A list of PolarDB cluster IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: status of the cluster.
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
           - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
           - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
    """
    __args__ = dict()
    __args__['dbType'] = db_type
    __args__['descriptionRegex'] = description_regex
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['status'] = status
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:polardb/getClusters:getClusters', __args__, opts=opts, typ=GetClustersResult).value

    return AwaitableGetClustersResult(
        clusters=__ret__.clusters,
        db_type=__ret__.db_type,
        description_regex=__ret__.description_regex,
        descriptions=__ret__.descriptions,
        id=__ret__.id,
        ids=__ret__.ids,
        output_file=__ret__.output_file,
        status=__ret__.status,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_clusters)
def get_clusters_output(db_type: Optional[pulumi.Input[Optional[str]]] = None,
                        description_regex: Optional[pulumi.Input[Optional[str]]] = None,
                        ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        status: Optional[pulumi.Input[Optional[str]]] = None,
                        tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClustersResult]:
    """
    The `polardb_get_clusters` data source provides a collection of PolarDB clusters available in Alibaba Cloud account.
    Filters support regular expression for the cluster description, searches by tags, and other filters which are listed below.

    > **NOTE:** Available in v1.66.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    polardb_clusters_ds = alicloud.polardb.get_clusters(description_regex="pc-\\\\w+",
        status="Running")
    pulumi.export("firstPolardbClusterId", polardb_clusters_ds.clusters[0].id)
    ```


    :param str db_type: Database type. Options are `MySQL`, `Oracle` and `PostgreSQL`. If no value is specified, all types are returned.
    :param str description_regex: A regex string to filter results by cluster description.
    :param Sequence[str] ids: A list of PolarDB cluster IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: status of the cluster.
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
           - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
           - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
    """
    ...
