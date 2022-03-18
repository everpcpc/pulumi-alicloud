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
    'GetGatewaysResult',
    'AwaitableGetGatewaysResult',
    'get_gateways',
    'get_gateways_output',
]

@pulumi.output_type
class GetGatewaysResult:
    """
    A collection of values returned by getGateways.
    """
    def __init__(__self__, enable_details=None, gateway_name=None, gateways=None, id=None, ids=None, name_regex=None, names=None, output_file=None, status=None, vpc_id=None):
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if gateway_name and not isinstance(gateway_name, str):
            raise TypeError("Expected argument 'gateway_name' to be a str")
        pulumi.set(__self__, "gateway_name", gateway_name)
        if gateways and not isinstance(gateways, list):
            raise TypeError("Expected argument 'gateways' to be a list")
        pulumi.set(__self__, "gateways", gateways)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

    @property
    @pulumi.getter(name="gatewayName")
    def gateway_name(self) -> Optional[str]:
        return pulumi.get(self, "gateway_name")

    @property
    @pulumi.getter
    def gateways(self) -> Sequence['outputs.GetGatewaysGatewayResult']:
        return pulumi.get(self, "gateways")

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
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")


class AwaitableGetGatewaysResult(GetGatewaysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewaysResult(
            enable_details=self.enable_details,
            gateway_name=self.gateway_name,
            gateways=self.gateways,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            status=self.status,
            vpc_id=self.vpc_id)


def get_gateways(enable_details: Optional[bool] = None,
                 gateway_name: Optional[str] = None,
                 ids: Optional[Sequence[str]] = None,
                 name_regex: Optional[str] = None,
                 output_file: Optional[str] = None,
                 status: Optional[str] = None,
                 vpc_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewaysResult:
    """
    This data source provides the Mse Gateways of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.157.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.mse.get_gateways(ids=["example_id"])
    pulumi.export("mseGatewayId1", ids.gateways[0].id)
    name_regex = alicloud.mse.get_gateways(name_regex="^my-Gateway")
    pulumi.export("mseGatewayId2", name_regex.gateways[0].id)
    status = alicloud.mse.get_gateways(status="2")
    pulumi.export("mseGatewayId3", status.gateways[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param str gateway_name: The name of the Gateway.
    :param Sequence[str] ids: A list of Gateway IDs.
    :param str name_regex: A regex string to filter results by Gateway name.
    :param str status: The status of the gateway.
    :param str vpc_id: The ID of the vpc.
    """
    __args__ = dict()
    __args__['enableDetails'] = enable_details
    __args__['gatewayName'] = gateway_name
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['status'] = status
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:mse/getGateways:getGateways', __args__, opts=opts, typ=GetGatewaysResult).value

    return AwaitableGetGatewaysResult(
        enable_details=__ret__.enable_details,
        gateway_name=__ret__.gateway_name,
        gateways=__ret__.gateways,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        status=__ret__.status,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(get_gateways)
def get_gateways_output(enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                        gateway_name: Optional[pulumi.Input[Optional[str]]] = None,
                        ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                        output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        status: Optional[pulumi.Input[Optional[str]]] = None,
                        vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGatewaysResult]:
    """
    This data source provides the Mse Gateways of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.157.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.mse.get_gateways(ids=["example_id"])
    pulumi.export("mseGatewayId1", ids.gateways[0].id)
    name_regex = alicloud.mse.get_gateways(name_regex="^my-Gateway")
    pulumi.export("mseGatewayId2", name_regex.gateways[0].id)
    status = alicloud.mse.get_gateways(status="2")
    pulumi.export("mseGatewayId3", status.gateways[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param str gateway_name: The name of the Gateway.
    :param Sequence[str] ids: A list of Gateway IDs.
    :param str name_regex: A regex string to filter results by Gateway name.
    :param str status: The status of the gateway.
    :param str vpc_id: The ID of the vpc.
    """
    ...
