# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ChainChainConfigArgs',
    'ChainChainConfigNodeArgs',
    'ChainChainConfigNodeNodeConfigArgs',
    'ChainChainConfigNodeNodeConfigDenyPolicyArgs',
    'ChainChainConfigRouterArgs',
    'ChainChainConfigRouterFromArgs',
    'ChainChainConfigRouterToArgs',
    'RepoDomainListArgs',
]

@pulumi.input_type
class ChainChainConfigArgs:
    def __init__(__self__, *,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigNodeArgs']]]] = None,
                 routers: Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigRouterArgs']]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['ChainChainConfigNodeArgs']]] nodes: Each node in the delivery chain.
        :param pulumi.Input[Sequence[pulumi.Input['ChainChainConfigRouterArgs']]] routers: Execution sequence relationship between delivery chain nodes.
        """
        if nodes is not None:
            pulumi.set(__self__, "nodes", nodes)
        if routers is not None:
            pulumi.set(__self__, "routers", routers)

    @property
    @pulumi.getter
    def nodes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigNodeArgs']]]]:
        """
        Each node in the delivery chain.
        """
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigNodeArgs']]]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter
    def routers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigRouterArgs']]]]:
        """
        Execution sequence relationship between delivery chain nodes.
        """
        return pulumi.get(self, "routers")

    @routers.setter
    def routers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigRouterArgs']]]]):
        pulumi.set(self, "routers", value)


@pulumi.input_type
class ChainChainConfigNodeArgs:
    def __init__(__self__, *,
                 enable: Optional[pulumi.Input[bool]] = None,
                 node_configs: Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigNodeNodeConfigArgs']]]] = None,
                 node_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] enable: Whether to enable the delivery chain node. Valid values: `true`, `false`.
        :param pulumi.Input[Sequence[pulumi.Input['ChainChainConfigNodeNodeConfigArgs']]] node_configs: The configuration of delivery chain node.
        :param pulumi.Input[str] node_name: The name of node. Valid values: `DOCKER_IMAGE_BUILD`, `DOCKER_IMAGE_PUSH`, `VULNERABILITY_SCANNING`, `ACTIVATE_REPLICATION`, `TRIGGER`, `SNAPSHOT`, `TRIGGER_SNAPSHOT`.
        """
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if node_configs is not None:
            pulumi.set(__self__, "node_configs", node_configs)
        if node_name is not None:
            pulumi.set(__self__, "node_name", node_name)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the delivery chain node. Valid values: `true`, `false`.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="nodeConfigs")
    def node_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigNodeNodeConfigArgs']]]]:
        """
        The configuration of delivery chain node.
        """
        return pulumi.get(self, "node_configs")

    @node_configs.setter
    def node_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigNodeNodeConfigArgs']]]]):
        pulumi.set(self, "node_configs", value)

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of node. Valid values: `DOCKER_IMAGE_BUILD`, `DOCKER_IMAGE_PUSH`, `VULNERABILITY_SCANNING`, `ACTIVATE_REPLICATION`, `TRIGGER`, `SNAPSHOT`, `TRIGGER_SNAPSHOT`.
        """
        return pulumi.get(self, "node_name")

    @node_name.setter
    def node_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_name", value)


@pulumi.input_type
class ChainChainConfigNodeNodeConfigArgs:
    def __init__(__self__, *,
                 deny_policies: Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigNodeNodeConfigDenyPolicyArgs']]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['ChainChainConfigNodeNodeConfigDenyPolicyArgs']]] deny_policies: Blocking rules for scanning nodes in delivery chain nodes. **Note:** When `node_name` is `VULNERABILITY_SCANNING`, the parameters in `deny_policy` need to be filled in.
        """
        if deny_policies is not None:
            pulumi.set(__self__, "deny_policies", deny_policies)

    @property
    @pulumi.getter(name="denyPolicies")
    def deny_policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigNodeNodeConfigDenyPolicyArgs']]]]:
        """
        Blocking rules for scanning nodes in delivery chain nodes. **Note:** When `node_name` is `VULNERABILITY_SCANNING`, the parameters in `deny_policy` need to be filled in.
        """
        return pulumi.get(self, "deny_policies")

    @deny_policies.setter
    def deny_policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigNodeNodeConfigDenyPolicyArgs']]]]):
        pulumi.set(self, "deny_policies", value)


@pulumi.input_type
class ChainChainConfigNodeNodeConfigDenyPolicyArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 issue_count: Optional[pulumi.Input[str]] = None,
                 issue_level: Optional[pulumi.Input[str]] = None,
                 logic: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] action: The action of trigger blocking. Valid values: `BLOCK`, `BLOCK_RETAG`, `BLOCK_DELETE_TAG`. While `Block` means block the delivery chain from continuing to execute, `BLOCK_RETAG` means block overwriting push image tag, `BLOCK_DELETE_TAG` means block deletion of mirror tags.
               
               > **NOTE:** The `from` and `to` fields are all fixed, and their structure and the value of `node_name` are fixed. You can refer to the template given in the example for configuration.
        :param pulumi.Input[str] issue_count: The count of scanning vulnerabilities that triggers blocking.
        :param pulumi.Input[str] issue_level: The level of scanning vulnerability that triggers blocking. Valid values: `LOW`, `MEDIUM`, `HIGH`, `UNKNOWN`.
        :param pulumi.Input[str] logic: The logic of trigger blocking. Valid values: `AND`, `OR`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if issue_count is not None:
            pulumi.set(__self__, "issue_count", issue_count)
        if issue_level is not None:
            pulumi.set(__self__, "issue_level", issue_level)
        if logic is not None:
            pulumi.set(__self__, "logic", logic)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        The action of trigger blocking. Valid values: `BLOCK`, `BLOCK_RETAG`, `BLOCK_DELETE_TAG`. While `Block` means block the delivery chain from continuing to execute, `BLOCK_RETAG` means block overwriting push image tag, `BLOCK_DELETE_TAG` means block deletion of mirror tags.

        > **NOTE:** The `from` and `to` fields are all fixed, and their structure and the value of `node_name` are fixed. You can refer to the template given in the example for configuration.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="issueCount")
    def issue_count(self) -> Optional[pulumi.Input[str]]:
        """
        The count of scanning vulnerabilities that triggers blocking.
        """
        return pulumi.get(self, "issue_count")

    @issue_count.setter
    def issue_count(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issue_count", value)

    @property
    @pulumi.getter(name="issueLevel")
    def issue_level(self) -> Optional[pulumi.Input[str]]:
        """
        The level of scanning vulnerability that triggers blocking. Valid values: `LOW`, `MEDIUM`, `HIGH`, `UNKNOWN`.
        """
        return pulumi.get(self, "issue_level")

    @issue_level.setter
    def issue_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issue_level", value)

    @property
    @pulumi.getter
    def logic(self) -> Optional[pulumi.Input[str]]:
        """
        The logic of trigger blocking. Valid values: `AND`, `OR`.
        """
        return pulumi.get(self, "logic")

    @logic.setter
    def logic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logic", value)


@pulumi.input_type
class ChainChainConfigRouterArgs:
    def __init__(__self__, *,
                 froms: Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigRouterFromArgs']]]] = None,
                 tos: Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigRouterToArgs']]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['ChainChainConfigRouterFromArgs']]] froms: Source node.
        :param pulumi.Input[Sequence[pulumi.Input['ChainChainConfigRouterToArgs']]] tos: Destination node.
        """
        if froms is not None:
            pulumi.set(__self__, "froms", froms)
        if tos is not None:
            pulumi.set(__self__, "tos", tos)

    @property
    @pulumi.getter
    def froms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigRouterFromArgs']]]]:
        """
        Source node.
        """
        return pulumi.get(self, "froms")

    @froms.setter
    def froms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigRouterFromArgs']]]]):
        pulumi.set(self, "froms", value)

    @property
    @pulumi.getter
    def tos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigRouterToArgs']]]]:
        """
        Destination node.
        """
        return pulumi.get(self, "tos")

    @tos.setter
    def tos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ChainChainConfigRouterToArgs']]]]):
        pulumi.set(self, "tos", value)


@pulumi.input_type
class ChainChainConfigRouterFromArgs:
    def __init__(__self__, *,
                 node_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] node_name: The name of node. Valid values: `DOCKER_IMAGE_BUILD`, `DOCKER_IMAGE_PUSH`, `VULNERABILITY_SCANNING`, `ACTIVATE_REPLICATION`, `TRIGGER`, `SNAPSHOT`, `TRIGGER_SNAPSHOT`.
        """
        if node_name is not None:
            pulumi.set(__self__, "node_name", node_name)

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of node. Valid values: `DOCKER_IMAGE_BUILD`, `DOCKER_IMAGE_PUSH`, `VULNERABILITY_SCANNING`, `ACTIVATE_REPLICATION`, `TRIGGER`, `SNAPSHOT`, `TRIGGER_SNAPSHOT`.
        """
        return pulumi.get(self, "node_name")

    @node_name.setter
    def node_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_name", value)


@pulumi.input_type
class ChainChainConfigRouterToArgs:
    def __init__(__self__, *,
                 node_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] node_name: The name of node. Valid values: `DOCKER_IMAGE_BUILD`, `DOCKER_IMAGE_PUSH`, `VULNERABILITY_SCANNING`, `ACTIVATE_REPLICATION`, `TRIGGER`, `SNAPSHOT`, `TRIGGER_SNAPSHOT`.
        """
        if node_name is not None:
            pulumi.set(__self__, "node_name", node_name)

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of node. Valid values: `DOCKER_IMAGE_BUILD`, `DOCKER_IMAGE_PUSH`, `VULNERABILITY_SCANNING`, `ACTIVATE_REPLICATION`, `TRIGGER`, `SNAPSHOT`, `TRIGGER_SNAPSHOT`.
        """
        return pulumi.get(self, "node_name")

    @node_name.setter
    def node_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_name", value)


@pulumi.input_type
class RepoDomainListArgs:
    def __init__(__self__, *,
                 internal: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[str]] = None,
                 vpc: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] internal: Domain of internal endpoint, only in some regions.
        :param pulumi.Input[str] public: Domain of public endpoint.
        :param pulumi.Input[str] vpc: Domain of vpc endpoint.
        """
        if internal is not None:
            pulumi.set(__self__, "internal", internal)
        if public is not None:
            pulumi.set(__self__, "public", public)
        if vpc is not None:
            pulumi.set(__self__, "vpc", vpc)

    @property
    @pulumi.getter
    def internal(self) -> Optional[pulumi.Input[str]]:
        """
        Domain of internal endpoint, only in some regions.
        """
        return pulumi.get(self, "internal")

    @internal.setter
    def internal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal", value)

    @property
    @pulumi.getter
    def public(self) -> Optional[pulumi.Input[str]]:
        """
        Domain of public endpoint.
        """
        return pulumi.get(self, "public")

    @public.setter
    def public(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public", value)

    @property
    @pulumi.getter
    def vpc(self) -> Optional[pulumi.Input[str]]:
        """
        Domain of vpc endpoint.
        """
        return pulumi.get(self, "vpc")

    @vpc.setter
    def vpc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc", value)


