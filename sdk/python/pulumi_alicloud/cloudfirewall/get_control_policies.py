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
    'GetControlPoliciesResult',
    'AwaitableGetControlPoliciesResult',
    'get_control_policies',
    'get_control_policies_output',
]

@pulumi.output_type
class GetControlPoliciesResult:
    """
    A collection of values returned by getControlPolicies.
    """
    def __init__(__self__, acl_action=None, acl_uuid=None, description=None, destination=None, direction=None, id=None, ids=None, ip_version=None, lang=None, output_file=None, policies=None, proto=None, source=None, source_ip=None):
        if acl_action and not isinstance(acl_action, str):
            raise TypeError("Expected argument 'acl_action' to be a str")
        pulumi.set(__self__, "acl_action", acl_action)
        if acl_uuid and not isinstance(acl_uuid, str):
            raise TypeError("Expected argument 'acl_uuid' to be a str")
        pulumi.set(__self__, "acl_uuid", acl_uuid)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if destination and not isinstance(destination, str):
            raise TypeError("Expected argument 'destination' to be a str")
        pulumi.set(__self__, "destination", destination)
        if direction and not isinstance(direction, str):
            raise TypeError("Expected argument 'direction' to be a str")
        pulumi.set(__self__, "direction", direction)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if ip_version and not isinstance(ip_version, str):
            raise TypeError("Expected argument 'ip_version' to be a str")
        pulumi.set(__self__, "ip_version", ip_version)
        if lang and not isinstance(lang, str):
            raise TypeError("Expected argument 'lang' to be a str")
        pulumi.set(__self__, "lang", lang)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)
        if proto and not isinstance(proto, str):
            raise TypeError("Expected argument 'proto' to be a str")
        pulumi.set(__self__, "proto", proto)
        if source and not isinstance(source, str):
            raise TypeError("Expected argument 'source' to be a str")
        pulumi.set(__self__, "source", source)
        if source_ip and not isinstance(source_ip, str):
            raise TypeError("Expected argument 'source_ip' to be a str")
        pulumi.set(__self__, "source_ip", source_ip)

    @property
    @pulumi.getter(name="aclAction")
    def acl_action(self) -> Optional[str]:
        return pulumi.get(self, "acl_action")

    @property
    @pulumi.getter(name="aclUuid")
    def acl_uuid(self) -> Optional[str]:
        return pulumi.get(self, "acl_uuid")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def destination(self) -> Optional[str]:
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter
    def direction(self) -> str:
        return pulumi.get(self, "direction")

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
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[str]:
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter
    def lang(self) -> Optional[str]:
        return pulumi.get(self, "lang")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def policies(self) -> Sequence['outputs.GetControlPoliciesPolicyResult']:
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def proto(self) -> Optional[str]:
        return pulumi.get(self, "proto")

    @property
    @pulumi.getter
    def source(self) -> Optional[str]:
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[str]:
        return pulumi.get(self, "source_ip")


class AwaitableGetControlPoliciesResult(GetControlPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetControlPoliciesResult(
            acl_action=self.acl_action,
            acl_uuid=self.acl_uuid,
            description=self.description,
            destination=self.destination,
            direction=self.direction,
            id=self.id,
            ids=self.ids,
            ip_version=self.ip_version,
            lang=self.lang,
            output_file=self.output_file,
            policies=self.policies,
            proto=self.proto,
            source=self.source,
            source_ip=self.source_ip)


def get_control_policies(acl_action: Optional[str] = None,
                         acl_uuid: Optional[str] = None,
                         description: Optional[str] = None,
                         destination: Optional[str] = None,
                         direction: Optional[str] = None,
                         ip_version: Optional[str] = None,
                         lang: Optional[str] = None,
                         output_file: Optional[str] = None,
                         proto: Optional[str] = None,
                         source: Optional[str] = None,
                         source_ip: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetControlPoliciesResult:
    """
    This data source provides the Cloud Firewall Control Policies of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.129.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.cloudfirewall.get_control_policies(direction="in")
    ```


    :param str acl_action: The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
    :param str acl_uuid: The unique ID of the access control policy.
    :param str description: The description of the access control policy.
    :param str destination: The destination address defined in the access control policy.
    :param str direction: The direction of traffic to which the access control policy applies. Valid values: `in`, `out`.
    :param str ip_version: The ip version.
    :param str lang: DestPortGroupPorts. Valid values: `en`, `zh`.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str proto: The protocol type of traffic to which the access control policy applies. Valid values: If `direction` is `in`, the valid value is `ANY`. If `direction` is `out`, the valid values are `ANY`, `TCP`, `UDP`, `ICMP`.
    :param str source: The source address defined in the access control policy.
    :param str source_ip: The source IP address of the request.
    """
    __args__ = dict()
    __args__['aclAction'] = acl_action
    __args__['aclUuid'] = acl_uuid
    __args__['description'] = description
    __args__['destination'] = destination
    __args__['direction'] = direction
    __args__['ipVersion'] = ip_version
    __args__['lang'] = lang
    __args__['outputFile'] = output_file
    __args__['proto'] = proto
    __args__['source'] = source
    __args__['sourceIp'] = source_ip
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cloudfirewall/getControlPolicies:getControlPolicies', __args__, opts=opts, typ=GetControlPoliciesResult).value

    return AwaitableGetControlPoliciesResult(
        acl_action=__ret__.acl_action,
        acl_uuid=__ret__.acl_uuid,
        description=__ret__.description,
        destination=__ret__.destination,
        direction=__ret__.direction,
        id=__ret__.id,
        ids=__ret__.ids,
        ip_version=__ret__.ip_version,
        lang=__ret__.lang,
        output_file=__ret__.output_file,
        policies=__ret__.policies,
        proto=__ret__.proto,
        source=__ret__.source,
        source_ip=__ret__.source_ip)


@_utilities.lift_output_func(get_control_policies)
def get_control_policies_output(acl_action: Optional[pulumi.Input[Optional[str]]] = None,
                                acl_uuid: Optional[pulumi.Input[Optional[str]]] = None,
                                description: Optional[pulumi.Input[Optional[str]]] = None,
                                destination: Optional[pulumi.Input[Optional[str]]] = None,
                                direction: Optional[pulumi.Input[str]] = None,
                                ip_version: Optional[pulumi.Input[Optional[str]]] = None,
                                lang: Optional[pulumi.Input[Optional[str]]] = None,
                                output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                proto: Optional[pulumi.Input[Optional[str]]] = None,
                                source: Optional[pulumi.Input[Optional[str]]] = None,
                                source_ip: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetControlPoliciesResult]:
    """
    This data source provides the Cloud Firewall Control Policies of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.129.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.cloudfirewall.get_control_policies(direction="in")
    ```


    :param str acl_action: The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
    :param str acl_uuid: The unique ID of the access control policy.
    :param str description: The description of the access control policy.
    :param str destination: The destination address defined in the access control policy.
    :param str direction: The direction of traffic to which the access control policy applies. Valid values: `in`, `out`.
    :param str ip_version: The ip version.
    :param str lang: DestPortGroupPorts. Valid values: `en`, `zh`.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str proto: The protocol type of traffic to which the access control policy applies. Valid values: If `direction` is `in`, the valid value is `ANY`. If `direction` is `out`, the valid values are `ANY`, `TCP`, `UDP`, `ICMP`.
    :param str source: The source address defined in the access control policy.
    :param str source_ip: The source IP address of the request.
    """
    ...
