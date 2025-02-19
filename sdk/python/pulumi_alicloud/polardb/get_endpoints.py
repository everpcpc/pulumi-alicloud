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
    'GetEndpointsResult',
    'AwaitableGetEndpointsResult',
    'get_endpoints',
    'get_endpoints_output',
]

@pulumi.output_type
class GetEndpointsResult:
    """
    A collection of values returned by getEndpoints.
    """
    def __init__(__self__, db_cluster_id=None, db_endpoint_id=None, endpoints=None, id=None):
        if db_cluster_id and not isinstance(db_cluster_id, str):
            raise TypeError("Expected argument 'db_cluster_id' to be a str")
        pulumi.set(__self__, "db_cluster_id", db_cluster_id)
        if db_endpoint_id and not isinstance(db_endpoint_id, str):
            raise TypeError("Expected argument 'db_endpoint_id' to be a str")
        pulumi.set(__self__, "db_endpoint_id", db_endpoint_id)
        if endpoints and not isinstance(endpoints, list):
            raise TypeError("Expected argument 'endpoints' to be a list")
        pulumi.set(__self__, "endpoints", endpoints)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="dbClusterId")
    def db_cluster_id(self) -> str:
        return pulumi.get(self, "db_cluster_id")

    @property
    @pulumi.getter(name="dbEndpointId")
    def db_endpoint_id(self) -> Optional[str]:
        """
        The endpoint ID.
        """
        return pulumi.get(self, "db_endpoint_id")

    @property
    @pulumi.getter
    def endpoints(self) -> Sequence['outputs.GetEndpointsEndpointResult']:
        """
        A list of PolarDB cluster endpoints. Each element contains the following attributes:
        """
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetEndpointsResult(GetEndpointsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEndpointsResult(
            db_cluster_id=self.db_cluster_id,
            db_endpoint_id=self.db_endpoint_id,
            endpoints=self.endpoints,
            id=self.id)


def get_endpoints(db_cluster_id: Optional[str] = None,
                  db_endpoint_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEndpointsResult:
    """
    The `polardb_get_endpoints` data source provides a collection of PolarDB endpoints available in Alibaba Cloud account.
    Filters support regular expression for the cluster name, searches by clusterId, and other filters which are listed below.

    > **NOTE:** Available in v1.68.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    polardb_clusters_ds = alicloud.polardb.get_clusters(description_regex="pc-\\\\w+",
        status="Running")
    default = alicloud.polardb.get_endpoints(db_cluster_id=polardb_clusters_ds.clusters[0].id)
    pulumi.export("endpoint", default.endpoints[0].db_endpoint_id)
    ```


    :param str db_cluster_id: PolarDB cluster ID.
    :param str db_endpoint_id: endpoint of the cluster.
    """
    __args__ = dict()
    __args__['dbClusterId'] = db_cluster_id
    __args__['dbEndpointId'] = db_endpoint_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:polardb/getEndpoints:getEndpoints', __args__, opts=opts, typ=GetEndpointsResult).value

    return AwaitableGetEndpointsResult(
        db_cluster_id=__ret__.db_cluster_id,
        db_endpoint_id=__ret__.db_endpoint_id,
        endpoints=__ret__.endpoints,
        id=__ret__.id)


@_utilities.lift_output_func(get_endpoints)
def get_endpoints_output(db_cluster_id: Optional[pulumi.Input[str]] = None,
                         db_endpoint_id: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEndpointsResult]:
    """
    The `polardb_get_endpoints` data source provides a collection of PolarDB endpoints available in Alibaba Cloud account.
    Filters support regular expression for the cluster name, searches by clusterId, and other filters which are listed below.

    > **NOTE:** Available in v1.68.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    polardb_clusters_ds = alicloud.polardb.get_clusters(description_regex="pc-\\\\w+",
        status="Running")
    default = alicloud.polardb.get_endpoints(db_cluster_id=polardb_clusters_ds.clusters[0].id)
    pulumi.export("endpoint", default.endpoints[0].db_endpoint_id)
    ```


    :param str db_cluster_id: PolarDB cluster ID.
    :param str db_endpoint_id: endpoint of the cluster.
    """
    ...
