# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetDiskTypesResult:
    """
    A collection of values returned by getDiskTypes.
    """
    def __init__(__self__, cluster_type=None, destination_resource=None, ids=None, instance_charge_type=None, instance_type=None, output_file=None, types=None, zone_id=None, id=None):
        if cluster_type and not isinstance(cluster_type, str):
            raise TypeError("Expected argument 'cluster_type' to be a str")
        __self__.cluster_type = cluster_type
        if destination_resource and not isinstance(destination_resource, str):
            raise TypeError("Expected argument 'destination_resource' to be a str")
        __self__.destination_resource = destination_resource
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of data disk and system disk type IDs. 
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
        if types and not isinstance(types, list):
            raise TypeError("Expected argument 'types' to be a list")
        __self__.types = types
        """
        A list of emr instance types. Each element contains the following attributes:
        """
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        __self__.zone_id = zone_id
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetDiskTypesResult(GetDiskTypesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDiskTypesResult(
            cluster_type=self.cluster_type,
            destination_resource=self.destination_resource,
            ids=self.ids,
            instance_charge_type=self.instance_charge_type,
            instance_type=self.instance_type,
            output_file=self.output_file,
            types=self.types,
            zone_id=self.zone_id,
            id=self.id)

def get_disk_types(cluster_type=None,destination_resource=None,instance_charge_type=None,instance_type=None,output_file=None,zone_id=None,opts=None):
    """
    The `emr.getDiskTypes` data source provides a collection of data disk and 
    system disk types available in Alibaba Cloud account when create a emr cluster.
    
    > **NOTE:** Available in 1.60.0+
    
    :param str cluster_type: The cluster type of the emr cluster instance. Possible values: `HADOOP`, `KAFKA`, `ZOOKEEPER`, `DRUID`.
    :param str destination_resource: The destination resource of emr cluster instance
    :param str instance_charge_type: Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
    :param str instance_type: The ecs instance type of create emr cluster instance.
    :param str zone_id: The Zone to create emr cluster instance.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/emr_disk_types.html.markdown.
    """
    __args__ = dict()

    __args__['clusterType'] = cluster_type
    __args__['destinationResource'] = destination_resource
    __args__['instanceChargeType'] = instance_charge_type
    __args__['instanceType'] = instance_type
    __args__['outputFile'] = output_file
    __args__['zoneId'] = zone_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:emr/getDiskTypes:getDiskTypes', __args__, opts=opts).value

    return AwaitableGetDiskTypesResult(
        cluster_type=__ret__.get('clusterType'),
        destination_resource=__ret__.get('destinationResource'),
        ids=__ret__.get('ids'),
        instance_charge_type=__ret__.get('instanceChargeType'),
        instance_type=__ret__.get('instanceType'),
        output_file=__ret__.get('outputFile'),
        types=__ret__.get('types'),
        zone_id=__ret__.get('zoneId'),
        id=__ret__.get('id'))
