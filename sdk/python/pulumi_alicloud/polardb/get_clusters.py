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
    def __init__(__self__, clusters=None, db_type=None, description_regex=None, descriptions=None, id=None, ids=None, output_file=None, status=None, tags=None):
        if clusters and not isinstance(clusters, list):
            raise TypeError("Expected argument 'clusters' to be a list")
        __self__.clusters = clusters
        """
        A list of PolarDB clusters. Each element contains the following attributes:
        """
        if db_type and not isinstance(db_type, str):
            raise TypeError("Expected argument 'db_type' to be a str")
        __self__.db_type = db_type
        """
        `Primary` for primary cluster, `ReadOnly` for read-only cluster, `Guard` for disaster recovery cluster, and `Temp` for temporary cluster.
        """
        if description_regex and not isinstance(description_regex, str):
            raise TypeError("Expected argument 'description_regex' to be a str")
        __self__.description_regex = description_regex
        if descriptions and not isinstance(descriptions, list):
            raise TypeError("Expected argument 'descriptions' to be a list")
        __self__.descriptions = descriptions
        """
        A list of RDS cluster descriptions. 
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of RDS cluster IDs. 
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        Status of the cluster.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
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

def get_clusters(db_type=None,description_regex=None,ids=None,output_file=None,status=None,tags=None,opts=None):
    """
    The `polardb.getClusters` data source provides a collection of PolarDB clusters available in Alibaba Cloud account.
    Filters support regular expression for the cluster description, searches by tags, and other filters which are listed below.

    > **NOTE:** Available in v1.66.0+.




    :param str db_type: Database type. Options are `MySQL`, `Oracle` and `PostgreSQL`. If no value is specified, all types are returned.
    :param str description_regex: A regex string to filter results by cluster description.
    :param list ids: A list of PolarDB cluster IDs. 
    :param str status: status of the cluster.
    :param dict tags: A mapping of tags to assign to the resource.
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
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:polardb/getClusters:getClusters', __args__, opts=opts).value

    return AwaitableGetClustersResult(
        clusters=__ret__.get('clusters'),
        db_type=__ret__.get('dbType'),
        description_regex=__ret__.get('descriptionRegex'),
        descriptions=__ret__.get('descriptions'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        output_file=__ret__.get('outputFile'),
        status=__ret__.get('status'),
        tags=__ret__.get('tags'))
