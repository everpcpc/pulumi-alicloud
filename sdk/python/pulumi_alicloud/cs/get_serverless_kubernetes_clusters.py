# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetServerlessKubernetesClustersResult:
    """
    A collection of values returned by getServerlessKubernetesClusters.
    """
    def __init__(__self__, clusters=None, enable_details=None, id=None, ids=None, name_regex=None, names=None, output_file=None):
        if clusters and not isinstance(clusters, list):
            raise TypeError("Expected argument 'clusters' to be a list")
        __self__.clusters = clusters
        """
        A list of matched Kubernetes clusters. Each element contains the following attributes:
        """
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        __self__.enable_details = enable_details
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
        A list of matched Kubernetes clusters' ids.
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of matched Kubernetes clusters' names.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
class AwaitableGetServerlessKubernetesClustersResult(GetServerlessKubernetesClustersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerlessKubernetesClustersResult(
            clusters=self.clusters,
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file)

def get_serverless_kubernetes_clusters(enable_details=None,ids=None,name_regex=None,output_file=None,opts=None):
    """
    This data source provides a list Container Service Serverless Kubernetes Clusters on Alibaba Cloud.

    > **NOTE:** Available in 1.58.0+




    :param list ids: Cluster IDs to filter.
    :param str name_regex: A regex string to filter results by cluster name.
    """
    __args__ = dict()


    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:cs/getServerlessKubernetesClusters:getServerlessKubernetesClusters', __args__, opts=opts).value

    return AwaitableGetServerlessKubernetesClustersResult(
        clusters=__ret__.get('clusters'),
        enable_details=__ret__.get('enableDetails'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'))
