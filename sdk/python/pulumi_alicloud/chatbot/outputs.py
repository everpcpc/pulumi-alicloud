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
    'GetAgentsAgentResult',
]

@pulumi.output_type
class GetAgentsAgentResult(dict):
    def __init__(__self__, *,
                 agent_id: str,
                 agent_key: str,
                 agent_name: str,
                 id: str):
        """
        :param str agent_id: The agent id.
        :param str agent_key: Service space signature, which is used when PAAS interface specifies the service space.
        :param str agent_name: The name of the agent.
        :param str id: ID of the agent.
        """
        pulumi.set(__self__, "agent_id", agent_id)
        pulumi.set(__self__, "agent_key", agent_key)
        pulumi.set(__self__, "agent_name", agent_name)
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="agentId")
    def agent_id(self) -> str:
        """
        The agent id.
        """
        return pulumi.get(self, "agent_id")

    @property
    @pulumi.getter(name="agentKey")
    def agent_key(self) -> str:
        """
        Service space signature, which is used when PAAS interface specifies the service space.
        """
        return pulumi.get(self, "agent_key")

    @property
    @pulumi.getter(name="agentName")
    def agent_name(self) -> str:
        """
        The name of the agent.
        """
        return pulumi.get(self, "agent_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the agent.
        """
        return pulumi.get(self, "id")

