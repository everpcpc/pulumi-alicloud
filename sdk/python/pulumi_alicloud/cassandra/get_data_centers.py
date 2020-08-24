# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetDataCentersResult:
    """
    A collection of values returned by getDataCenters.
    """
    def __init__(__self__, centers=None, cluster_id=None, id=None, ids=None, name_regex=None, names=None, output_file=None):
        if centers and not isinstance(centers, list):
            raise TypeError("Expected argument 'centers' to be a list")
        __self__.centers = centers
        """
        A list of Cassandra data centers. Its every element contains the following attributes:
        """
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        __self__.cluster_id = cluster_id
        """
        The ID of the Cassandra cluster.
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
        The list of Cassandra data center ids.
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        The name list of Cassandra data centers.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
class AwaitableGetDataCentersResult(GetDataCentersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDataCentersResult(
            centers=self.centers,
            cluster_id=self.cluster_id,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file)

def get_data_centers(cluster_id=None,ids=None,name_regex=None,output_file=None,opts=None):
    """
    The `cassandra.getDataCenters` data source provides a collection of Cassandra Data Centers available in Alicloud account.
    Filters support regular expression for the cluster name or ids.

    > **NOTE:**  Available in 1.88.0+.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    cassandra = alicloud.cassandra.get_data_centers(cluster_id="cds-xxxxx",
        name_regex="tf_testAccCassandra_dc")
    ```



    :param str cluster_id: The cluster id of dataCenters belongs to.
    :param list ids: The list of Cassandra data center ids.
    :param str name_regex: A regex string to apply to the cluster name.
    """
    __args__ = dict()


    __args__['clusterId'] = cluster_id
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:cassandra/getDataCenters:getDataCenters', __args__, opts=opts).value

    return AwaitableGetDataCentersResult(
        centers=__ret__.get('centers'),
        cluster_id=__ret__.get('clusterId'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'))