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
    'GetSecurityGroupRulesResult',
    'AwaitableGetSecurityGroupRulesResult',
    'get_security_group_rules',
    'get_security_group_rules_output',
]

@pulumi.output_type
class GetSecurityGroupRulesResult:
    """
    A collection of values returned by getSecurityGroupRules.
    """
    def __init__(__self__, direction=None, group_desc=None, group_id=None, group_name=None, id=None, ip_protocol=None, nic_type=None, output_file=None, policy=None, rules=None):
        if direction and not isinstance(direction, str):
            raise TypeError("Expected argument 'direction' to be a str")
        pulumi.set(__self__, "direction", direction)
        if group_desc and not isinstance(group_desc, str):
            raise TypeError("Expected argument 'group_desc' to be a str")
        pulumi.set(__self__, "group_desc", group_desc)
        if group_id and not isinstance(group_id, str):
            raise TypeError("Expected argument 'group_id' to be a str")
        pulumi.set(__self__, "group_id", group_id)
        if group_name and not isinstance(group_name, str):
            raise TypeError("Expected argument 'group_name' to be a str")
        pulumi.set(__self__, "group_name", group_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_protocol and not isinstance(ip_protocol, str):
            raise TypeError("Expected argument 'ip_protocol' to be a str")
        pulumi.set(__self__, "ip_protocol", ip_protocol)
        if nic_type and not isinstance(nic_type, str):
            raise TypeError("Expected argument 'nic_type' to be a str")
        pulumi.set(__self__, "nic_type", nic_type)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if policy and not isinstance(policy, str):
            raise TypeError("Expected argument 'policy' to be a str")
        pulumi.set(__self__, "policy", policy)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter
    def direction(self) -> Optional[str]:
        """
        Authorization direction, `ingress` or `egress`.
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter(name="groupDesc")
    def group_desc(self) -> str:
        """
        The description of the security group that owns the rules.
        """
        return pulumi.get(self, "group_desc")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> str:
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> str:
        """
        The name of the security group that owns the rules.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> Optional[str]:
        """
        The protocol. Can be `tcp`, `udp`, `icmp`, `gre` or `all`.
        """
        return pulumi.get(self, "ip_protocol")

    @property
    @pulumi.getter(name="nicType")
    def nic_type(self) -> Optional[str]:
        """
        Network type, `internet` or `intranet`.
        """
        return pulumi.get(self, "nic_type")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def policy(self) -> Optional[str]:
        """
        Authorization policy. Can be either `accept` or `drop`.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetSecurityGroupRulesRuleResult']:
        """
        A list of security group rules. Each element contains the following attributes:
        """
        return pulumi.get(self, "rules")


class AwaitableGetSecurityGroupRulesResult(GetSecurityGroupRulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecurityGroupRulesResult(
            direction=self.direction,
            group_desc=self.group_desc,
            group_id=self.group_id,
            group_name=self.group_name,
            id=self.id,
            ip_protocol=self.ip_protocol,
            nic_type=self.nic_type,
            output_file=self.output_file,
            policy=self.policy,
            rules=self.rules)


def get_security_group_rules(direction: Optional[str] = None,
                             group_id: Optional[str] = None,
                             ip_protocol: Optional[str] = None,
                             nic_type: Optional[str] = None,
                             output_file: Optional[str] = None,
                             policy: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecurityGroupRulesResult:
    """
    The `ecs_get_security_group_rules` data source provides a collection of security permissions of a specific security group.
    Each collection item represents a single `ingress` or `egress` permission rule.
    The ID of the security group can be provided via a variable or the result from the other data source `ecs_get_security_groups`.

    ## Example Usage

    The following example shows how to obtain details about a security group rule and how to pass its data to an instance at launch time.

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    security_group_id = config.require_object("securityGroupId")
    groups_ds = alicloud.ecs.get_security_groups(name_regex="api")
    ingress_rules_ds = alicloud.ecs.get_security_group_rules(direction="ingress",
        group_id=groups_ds.groups[0].id,
        ip_protocol="tcp",
        nic_type="internet")
    # Pass port_range to the backend service
    backend = alicloud.ecs.Instance("backend", user_data=f"config_service.sh --portrange={ingress_rules_ds.rules[0].port_range}")
    ```


    :param str direction: Authorization direction. Valid values are: `ingress` or `egress`.
    :param str group_id: The ID of the security group that owns the rules.
    :param str ip_protocol: The IP protocol. Valid values are: `tcp`, `udp`, `icmp`, `gre` and `all`.
    :param str nic_type: Refers to the network type. Can be either `internet` or `intranet`. The default value is `internet`.
    :param str policy: Authorization policy. Can be either `accept` or `drop`. The default value is `accept`.
    """
    __args__ = dict()
    __args__['direction'] = direction
    __args__['groupId'] = group_id
    __args__['ipProtocol'] = ip_protocol
    __args__['nicType'] = nic_type
    __args__['outputFile'] = output_file
    __args__['policy'] = policy
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:ecs/getSecurityGroupRules:getSecurityGroupRules', __args__, opts=opts, typ=GetSecurityGroupRulesResult).value

    return AwaitableGetSecurityGroupRulesResult(
        direction=__ret__.direction,
        group_desc=__ret__.group_desc,
        group_id=__ret__.group_id,
        group_name=__ret__.group_name,
        id=__ret__.id,
        ip_protocol=__ret__.ip_protocol,
        nic_type=__ret__.nic_type,
        output_file=__ret__.output_file,
        policy=__ret__.policy,
        rules=__ret__.rules)


@_utilities.lift_output_func(get_security_group_rules)
def get_security_group_rules_output(direction: Optional[pulumi.Input[Optional[str]]] = None,
                                    group_id: Optional[pulumi.Input[str]] = None,
                                    ip_protocol: Optional[pulumi.Input[Optional[str]]] = None,
                                    nic_type: Optional[pulumi.Input[Optional[str]]] = None,
                                    output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                    policy: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecurityGroupRulesResult]:
    """
    The `ecs_get_security_group_rules` data source provides a collection of security permissions of a specific security group.
    Each collection item represents a single `ingress` or `egress` permission rule.
    The ID of the security group can be provided via a variable or the result from the other data source `ecs_get_security_groups`.

    ## Example Usage

    The following example shows how to obtain details about a security group rule and how to pass its data to an instance at launch time.

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    security_group_id = config.require_object("securityGroupId")
    groups_ds = alicloud.ecs.get_security_groups(name_regex="api")
    ingress_rules_ds = alicloud.ecs.get_security_group_rules(direction="ingress",
        group_id=groups_ds.groups[0].id,
        ip_protocol="tcp",
        nic_type="internet")
    # Pass port_range to the backend service
    backend = alicloud.ecs.Instance("backend", user_data=f"config_service.sh --portrange={ingress_rules_ds.rules[0].port_range}")
    ```


    :param str direction: Authorization direction. Valid values are: `ingress` or `egress`.
    :param str group_id: The ID of the security group that owns the rules.
    :param str ip_protocol: The IP protocol. Valid values are: `tcp`, `udp`, `icmp`, `gre` and `all`.
    :param str nic_type: Refers to the network type. Can be either `internet` or `intranet`. The default value is `internet`.
    :param str policy: Authorization policy. Can be either `accept` or `drop`. The default value is `accept`.
    """
    ...
