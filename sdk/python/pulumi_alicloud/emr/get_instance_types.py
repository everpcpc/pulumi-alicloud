# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetInstanceTypesResult:
    """
    A collection of values returned by getInstanceTypes.
    """
    def __init__(__self__, cluster_type=None, destination_resource=None, id=None, ids=None, instance_charge_type=None, instance_type=None, output_file=None, support_local_storage=None, support_node_types=None, types=None, zone_id=None):
        if cluster_type and not isinstance(cluster_type, str):
            raise TypeError("Expected argument 'cluster_type' to be a str")
        __self__.cluster_type = cluster_type
        if destination_resource and not isinstance(destination_resource, str):
            raise TypeError("Expected argument 'destination_resource' to be a str")
        __self__.destination_resource = destination_resource
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
        A list of emr instance types IDs. 
        """
        if instance_charge_type and not isinstance(instance_charge_type, str):
            raise TypeError("Expected argument 'instance_charge_type' to be a str")
        __self__.instance_charge_type = instance_charge_type
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        __self__.instance_type = instance_type
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if support_local_storage and not isinstance(support_local_storage, bool):
            raise TypeError("Expected argument 'support_local_storage' to be a bool")
        __self__.support_local_storage = support_local_storage
        if support_node_types and not isinstance(support_node_types, list):
            raise TypeError("Expected argument 'support_node_types' to be a list")
        __self__.support_node_types = support_node_types
        if types and not isinstance(types, list):
            raise TypeError("Expected argument 'types' to be a list")
        __self__.types = types
        """
        A list of emr instance types. Each element contains the following attributes:
        """
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        __self__.zone_id = zone_id
        """
        The available zone id in Alibaba Cloud account
        """
class AwaitableGetInstanceTypesResult(GetInstanceTypesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceTypesResult(
            cluster_type=self.cluster_type,
            destination_resource=self.destination_resource,
            id=self.id,
            ids=self.ids,
            instance_charge_type=self.instance_charge_type,
            instance_type=self.instance_type,
            output_file=self.output_file,
            support_local_storage=self.support_local_storage,
            support_node_types=self.support_node_types,
            types=self.types,
            zone_id=self.zone_id)

def get_instance_types(cluster_type=None,destination_resource=None,instance_charge_type=None,instance_type=None,output_file=None,support_local_storage=None,support_node_types=None,zone_id=None,opts=None):
    """
    The `emr.getInstanceTypes` data source provides a collection of ecs
    instance types available in Alibaba Cloud account when create a emr cluster.

    > **NOTE:** Available in 1.59.0+

    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/emr_instance_types.html.markdown.


    :param str cluster_type: The cluster type of the emr cluster instance. Possible values: `HADOOP`, `KAFKA`, `ZOOKEEPER`, `DRUID`.
    :param str destination_resource: The destination resource of emr cluster instance
    :param str instance_charge_type: Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
    :param str instance_type: Filter the specific ecs instance type to create emr cluster.
    :param bool support_local_storage: Whether the current storage disk is local or not.
    :param list support_node_types: The specific supported node type list.
           Possible values may be any one or combination of these: ["MASTER", "CORE", "TASK", "GATEWAY"]
    :param str zone_id: The supported resources of specific zoneId.
    """
    __args__ = dict()


    __args__['clusterType'] = cluster_type
    __args__['destinationResource'] = destination_resource
    __args__['instanceChargeType'] = instance_charge_type
    __args__['instanceType'] = instance_type
    __args__['outputFile'] = output_file
    __args__['supportLocalStorage'] = support_local_storage
    __args__['supportNodeTypes'] = support_node_types
    __args__['zoneId'] = zone_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:emr/getInstanceTypes:getInstanceTypes', __args__, opts=opts).value

    return AwaitableGetInstanceTypesResult(
        cluster_type=__ret__.get('clusterType'),
        destination_resource=__ret__.get('destinationResource'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        instance_charge_type=__ret__.get('instanceChargeType'),
        instance_type=__ret__.get('instanceType'),
        output_file=__ret__.get('outputFile'),
        support_local_storage=__ret__.get('supportLocalStorage'),
        support_node_types=__ret__.get('supportNodeTypes'),
        types=__ret__.get('types'),
        zone_id=__ret__.get('zoneId'))
