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
    'GetBandwidthLimitsResult',
    'AwaitableGetBandwidthLimitsResult',
    'get_bandwidth_limits',
]

@pulumi.output_type
class GetBandwidthLimitsResult:
    """
    A collection of values returned by getBandwidthLimits.
    """
    def __init__(__self__, id=None, instance_ids=None, limits=None, output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_ids and not isinstance(instance_ids, list):
            raise TypeError("Expected argument 'instance_ids' to be a list")
        pulumi.set(__self__, "instance_ids", instance_ids)
        if limits and not isinstance(limits, list):
            raise TypeError("Expected argument 'limits' to be a list")
        pulumi.set(__self__, "limits", limits)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "instance_ids")

    @property
    @pulumi.getter
    def limits(self) -> Sequence['outputs.GetBandwidthLimitsLimitResult']:
        """
        A list of CEN Bandwidth Limits. Each element contains the following attributes:
        """
        return pulumi.get(self, "limits")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")


class AwaitableGetBandwidthLimitsResult(GetBandwidthLimitsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBandwidthLimitsResult(
            id=self.id,
            instance_ids=self.instance_ids,
            limits=self.limits,
            output_file=self.output_file)


def get_bandwidth_limits(instance_ids: Optional[Sequence[str]] = None,
                         output_file: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBandwidthLimitsResult:
    """
    This data source provides CEN Bandwidth Limits available to the user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    bwl = alicloud.cen.get_bandwidth_limits(instance_ids=["cen-id1"])
    pulumi.export("firstCenBandwidthLimitsLocalRegionId", bwl.limits[0].local_region_id)
    ```


    :param Sequence[str] instance_ids: A list of CEN instances IDs.
    """
    __args__ = dict()
    __args__['instanceIds'] = instance_ids
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:cen/getBandwidthLimits:getBandwidthLimits', __args__, opts=opts, typ=GetBandwidthLimitsResult).value

    return AwaitableGetBandwidthLimitsResult(
        id=__ret__.id,
        instance_ids=__ret__.instance_ids,
        limits=__ret__.limits,
        output_file=__ret__.output_file)
