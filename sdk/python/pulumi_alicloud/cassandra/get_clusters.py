# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetClustersResult:
    """
    A collection of values returned by getClusters.
    """
    def __init__(__self__, clusters=None, id=None, ids=None, name_regex=None, names=None, output_file=None, tags=None):
        if clusters and not isinstance(clusters, list):
            raise TypeError("Expected argument 'clusters' to be a list")
        __self__.clusters = clusters
        """
        A list of Cassandra clusters. Its every element contains the following attributes:
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        The list of Cassandra cluster ids.
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        The name list of Cassandra clusters.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
class AwaitableGetClustersResult(GetClustersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClustersResult(
            clusters=self.clusters,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            tags=self.tags)

def get_clusters(ids=None,name_regex=None,output_file=None,tags=None,opts=None):
    """
    The `cassandra.getClusters` data source provides a collection of Cassandra clusters available in Alicloud account.
    Filters support regular expression for the cluster name, ids or tags.

    > **NOTE:**  Available in 1.88.0+.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    cassandra = alicloud.cassandra.get_clusters(name_regex="tf_testAccCassandra")
    ```



    :param list ids: The list of Cassandra cluster ids.
    :param str name_regex: A regex string to apply to the cluster name.
    :param dict tags: A mapping of tags to assign to the resource.
    """
    __args__ = dict()


    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:cassandra/getClusters:getClusters', __args__, opts=opts).value

    return AwaitableGetClustersResult(
        clusters=__ret__.get('clusters'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        tags=__ret__.get('tags'))
