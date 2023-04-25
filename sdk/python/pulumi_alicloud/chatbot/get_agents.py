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
    'GetAgentsResult',
    'AwaitableGetAgentsResult',
    'get_agents',
    'get_agents_output',
]

@pulumi.output_type
class GetAgentsResult:
    """
    A collection of values returned by getAgents.
    """
    def __init__(__self__, agent_name=None, agents=None, id=None, ids=None, name_regex=None, names=None, output_file=None, page_number=None, page_size=None):
        if agent_name and not isinstance(agent_name, str):
            raise TypeError("Expected argument 'agent_name' to be a str")
        pulumi.set(__self__, "agent_name", agent_name)
        if agents and not isinstance(agents, list):
            raise TypeError("Expected argument 'agents' to be a list")
        pulumi.set(__self__, "agents", agents)
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
        if page_number and not isinstance(page_number, int):
            raise TypeError("Expected argument 'page_number' to be a int")
        pulumi.set(__self__, "page_number", page_number)
        if page_size and not isinstance(page_size, int):
            raise TypeError("Expected argument 'page_size' to be a int")
        pulumi.set(__self__, "page_size", page_size)

    @property
    @pulumi.getter(name="agentName")
    def agent_name(self) -> str:
        """
        The agent Name.
        """
        return pulumi.get(self, "agent_name")

    @property
    @pulumi.getter
    def agents(self) -> Sequence['outputs.GetAgentsAgentResult']:
        """
        A list of availability zones. Each element contains the following attributes:
        """
        return pulumi.get(self, "agents")

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
        """
        A list of chatbot agents names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="pageNumber")
    def page_number(self) -> Optional[int]:
        return pulumi.get(self, "page_number")

    @property
    @pulumi.getter(name="pageSize")
    def page_size(self) -> Optional[int]:
        return pulumi.get(self, "page_size")


class AwaitableGetAgentsResult(GetAgentsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAgentsResult(
            agent_name=self.agent_name,
            agents=self.agents,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            page_number=self.page_number,
            page_size=self.page_size)


def get_agents(agent_name: Optional[str] = None,
               name_regex: Optional[str] = None,
               output_file: Optional[str] = None,
               page_number: Optional[int] = None,
               page_size: Optional[int] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAgentsResult:
    """
    This data source provides the Chatbot Agents of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.203.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    name_regex = alicloud.chatbot.get_agents(name_regex="^my-Agent")
    pulumi.export("alicloudChatbotAgentsId1", name_regex.agents[0].id)
    ```


    :param str agent_name: The name of the agent.
    :param str name_regex: A regex string to filter resulting chatbot agents by name.
    """
    __args__ = dict()
    __args__['agentName'] = agent_name
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['pageNumber'] = page_number
    __args__['pageSize'] = page_size
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:chatbot/getAgents:getAgents', __args__, opts=opts, typ=GetAgentsResult).value

    return AwaitableGetAgentsResult(
        agent_name=__ret__.agent_name,
        agents=__ret__.agents,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        page_number=__ret__.page_number,
        page_size=__ret__.page_size)


@_utilities.lift_output_func(get_agents)
def get_agents_output(agent_name: Optional[pulumi.Input[Optional[str]]] = None,
                      name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                      output_file: Optional[pulumi.Input[Optional[str]]] = None,
                      page_number: Optional[pulumi.Input[Optional[int]]] = None,
                      page_size: Optional[pulumi.Input[Optional[int]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAgentsResult]:
    """
    This data source provides the Chatbot Agents of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.203.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    name_regex = alicloud.chatbot.get_agents(name_regex="^my-Agent")
    pulumi.export("alicloudChatbotAgentsId1", name_regex.agents[0].id)
    ```


    :param str agent_name: The name of the agent.
    :param str name_regex: A regex string to filter resulting chatbot agents by name.
    """
    ...
