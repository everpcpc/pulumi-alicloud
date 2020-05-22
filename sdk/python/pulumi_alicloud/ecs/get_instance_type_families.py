# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetInstanceTypeFamiliesResult:
    """
    A collection of values returned by getInstanceTypeFamilies.
    """
    def __init__(__self__, families=None, generation=None, id=None, ids=None, instance_charge_type=None, output_file=None, spot_strategy=None, zone_id=None):
        if families and not isinstance(families, list):
            raise TypeError("Expected argument 'families' to be a list")
        __self__.families = families
        if generation and not isinstance(generation, str):
            raise TypeError("Expected argument 'generation' to be a str")
        __self__.generation = generation
        """
        The generation of the instance type family.
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
        A list of instance type family IDs.
        """
        if instance_charge_type and not isinstance(instance_charge_type, str):
            raise TypeError("Expected argument 'instance_charge_type' to be a str")
        __self__.instance_charge_type = instance_charge_type
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if spot_strategy and not isinstance(spot_strategy, str):
            raise TypeError("Expected argument 'spot_strategy' to be a str")
        __self__.spot_strategy = spot_strategy
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        __self__.zone_id = zone_id
class AwaitableGetInstanceTypeFamiliesResult(GetInstanceTypeFamiliesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceTypeFamiliesResult(
            families=self.families,
            generation=self.generation,
            id=self.id,
            ids=self.ids,
            instance_charge_type=self.instance_charge_type,
            output_file=self.output_file,
            spot_strategy=self.spot_strategy,
            zone_id=self.zone_id)

def get_instance_type_families(generation=None,instance_charge_type=None,output_file=None,spot_strategy=None,zone_id=None,opts=None):
    """
    This data source provides the ECS instance type families of Alibaba Cloud.

    > **NOTE:** Available in 1.54.0+

    ## Example Usage



    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.ecs.get_instance_type_families(instance_charge_type="PrePaid")
    pulumi.export("firstInstanceTypeFamilyId", default.families[0]["id"])
    pulumi.export("instanceIds", default.ids)
    ```



    :param str generation: The generation of the instance type family, Valid values: `ecs-1`, `ecs-2`, `ecs-3` and `ecs-4`. For more information, see [Instance type families](https://www.alibabacloud.com/help/doc-detail/25378.htm). 
    :param str instance_charge_type: Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
    :param str spot_strategy: Filter the results by ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
    :param str zone_id: The Zone to launch the instance.
    """
    __args__ = dict()


    __args__['generation'] = generation
    __args__['instanceChargeType'] = instance_charge_type
    __args__['outputFile'] = output_file
    __args__['spotStrategy'] = spot_strategy
    __args__['zoneId'] = zone_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:ecs/getInstanceTypeFamilies:getInstanceTypeFamilies', __args__, opts=opts).value

    return AwaitableGetInstanceTypeFamiliesResult(
        families=__ret__.get('families'),
        generation=__ret__.get('generation'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        instance_charge_type=__ret__.get('instanceChargeType'),
        output_file=__ret__.get('outputFile'),
        spot_strategy=__ret__.get('spotStrategy'),
        zone_id=__ret__.get('zoneId'))
