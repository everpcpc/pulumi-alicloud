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
    'GetVbrHealthChecksResult',
    'AwaitableGetVbrHealthChecksResult',
    'get_vbr_health_checks',
    'get_vbr_health_checks_output',
]

@pulumi.output_type
class GetVbrHealthChecksResult:
    """
    A collection of values returned by getVbrHealthChecks.
    """
    def __init__(__self__, cen_id=None, checks=None, id=None, ids=None, output_file=None, vbr_instance_id=None, vbr_instance_owner_id=None, vbr_instance_region_id=None):
        if cen_id and not isinstance(cen_id, str):
            raise TypeError("Expected argument 'cen_id' to be a str")
        pulumi.set(__self__, "cen_id", cen_id)
        if checks and not isinstance(checks, list):
            raise TypeError("Expected argument 'checks' to be a list")
        pulumi.set(__self__, "checks", checks)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if vbr_instance_id and not isinstance(vbr_instance_id, str):
            raise TypeError("Expected argument 'vbr_instance_id' to be a str")
        pulumi.set(__self__, "vbr_instance_id", vbr_instance_id)
        if vbr_instance_owner_id and not isinstance(vbr_instance_owner_id, int):
            raise TypeError("Expected argument 'vbr_instance_owner_id' to be a int")
        pulumi.set(__self__, "vbr_instance_owner_id", vbr_instance_owner_id)
        if vbr_instance_region_id and not isinstance(vbr_instance_region_id, str):
            raise TypeError("Expected argument 'vbr_instance_region_id' to be a str")
        pulumi.set(__self__, "vbr_instance_region_id", vbr_instance_region_id)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> Optional[str]:
        """
        The ID of the Cloud Enterprise Network (CEN) instance.
        """
        return pulumi.get(self, "cen_id")

    @property
    @pulumi.getter
    def checks(self) -> Sequence['outputs.GetVbrHealthChecksCheckResult']:
        """
        A list of CEN VBR Heath Checks. Each element contains the following attributes:
        """
        return pulumi.get(self, "checks")

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
        A list of the CEN VBR Heath Check IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="vbrInstanceId")
    def vbr_instance_id(self) -> Optional[str]:
        """
        The ID of the VBR instance.
        """
        return pulumi.get(self, "vbr_instance_id")

    @property
    @pulumi.getter(name="vbrInstanceOwnerId")
    def vbr_instance_owner_id(self) -> Optional[int]:
        return pulumi.get(self, "vbr_instance_owner_id")

    @property
    @pulumi.getter(name="vbrInstanceRegionId")
    def vbr_instance_region_id(self) -> str:
        """
        The ID of the region where the VBR instance is deployed.
        """
        return pulumi.get(self, "vbr_instance_region_id")


class AwaitableGetVbrHealthChecksResult(GetVbrHealthChecksResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVbrHealthChecksResult(
            cen_id=self.cen_id,
            checks=self.checks,
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            vbr_instance_id=self.vbr_instance_id,
            vbr_instance_owner_id=self.vbr_instance_owner_id,
            vbr_instance_region_id=self.vbr_instance_region_id)


def get_vbr_health_checks(cen_id: Optional[str] = None,
                          output_file: Optional[str] = None,
                          vbr_instance_id: Optional[str] = None,
                          vbr_instance_owner_id: Optional[int] = None,
                          vbr_instance_region_id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVbrHealthChecksResult:
    """
    This data source provides CEN VBR Health Checks available to the user.

    > **NOTE:** Available in 1.98.0+


    :param str cen_id: The ID of the Cloud Enterprise Network (CEN) instance.
    :param str vbr_instance_id: The ID of the VBR instance.
    :param int vbr_instance_owner_id: The User ID (UID) of the account to which the VBR instance belongs.
    :param str vbr_instance_region_id: The ID of the region where the VBR instance is deployed.
    """
    __args__ = dict()
    __args__['cenId'] = cen_id
    __args__['outputFile'] = output_file
    __args__['vbrInstanceId'] = vbr_instance_id
    __args__['vbrInstanceOwnerId'] = vbr_instance_owner_id
    __args__['vbrInstanceRegionId'] = vbr_instance_region_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:cen/getVbrHealthChecks:getVbrHealthChecks', __args__, opts=opts, typ=GetVbrHealthChecksResult).value

    return AwaitableGetVbrHealthChecksResult(
        cen_id=__ret__.cen_id,
        checks=__ret__.checks,
        id=__ret__.id,
        ids=__ret__.ids,
        output_file=__ret__.output_file,
        vbr_instance_id=__ret__.vbr_instance_id,
        vbr_instance_owner_id=__ret__.vbr_instance_owner_id,
        vbr_instance_region_id=__ret__.vbr_instance_region_id)


@_utilities.lift_output_func(get_vbr_health_checks)
def get_vbr_health_checks_output(cen_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 vbr_instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 vbr_instance_owner_id: Optional[pulumi.Input[Optional[int]]] = None,
                                 vbr_instance_region_id: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVbrHealthChecksResult]:
    """
    This data source provides CEN VBR Health Checks available to the user.

    > **NOTE:** Available in 1.98.0+


    :param str cen_id: The ID of the Cloud Enterprise Network (CEN) instance.
    :param str vbr_instance_id: The ID of the VBR instance.
    :param int vbr_instance_owner_id: The User ID (UID) of the account to which the VBR instance belongs.
    :param str vbr_instance_region_id: The ID of the region where the VBR instance is deployed.
    """
    ...
