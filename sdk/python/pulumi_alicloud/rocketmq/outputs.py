# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetGroupsGroupResult',
    'GetInstancesInstanceResult',
    'GetTopicsTopicResult',
]

@pulumi.output_type
class GetGroupsGroupResult(dict):
    def __init__(__self__, *,
                 id: str,
                 independent_naming: bool,
                 owner: str,
                 remark: str):
        """
        :param str id: The name of the group.
        :param bool independent_naming: Indicates whether namespaces are available. Read [Fields in SubscribeInfoDo](https://www.alibabacloud.com/help/doc-detail/29619.html) for further details.
        :param str owner: The ID of the group owner, which is the Alibaba Cloud UID.
        :param str remark: Remark of the group.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "independent_naming", independent_naming)
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "remark", remark)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The name of the group.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="independentNaming")
    def independent_naming(self) -> bool:
        """
        Indicates whether namespaces are available. Read [Fields in SubscribeInfoDo](https://www.alibabacloud.com/help/doc-detail/29619.html) for further details.
        """
        return pulumi.get(self, "independent_naming")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        The ID of the group owner, which is the Alibaba Cloud UID.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def remark(self) -> str:
        """
        Remark of the group.
        """
        return pulumi.get(self, "remark")


@pulumi.output_type
class GetInstancesInstanceResult(dict):
    def __init__(__self__, *,
                 id: str,
                 instance_id: str,
                 instance_name: str,
                 instance_status: float,
                 instance_type: float,
                 release_time: float):
        """
        :param str id: ID of the instance.
        :param str instance_id: ID of the instance.
        :param str instance_name: Name of the instance.
        :param float instance_status: The status of the instance. Read [Fields in InstanceVO](https://www.alibabacloud.com/help/doc-detail/106351.html) for further details.
        :param float instance_type: The type of the instance. Read [Fields in InstanceVO](https://www.alibabacloud.com/help/doc-detail/106351.html) for further details.
        :param float release_time: The automatic release time of an Enterprise Platinum Edition instance.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "instance_status", instance_status)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "release_time", release_time)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the instance.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        ID of the instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        """
        Name of the instance.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="instanceStatus")
    def instance_status(self) -> float:
        """
        The status of the instance. Read [Fields in InstanceVO](https://www.alibabacloud.com/help/doc-detail/106351.html) for further details.
        """
        return pulumi.get(self, "instance_status")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> float:
        """
        The type of the instance. Read [Fields in InstanceVO](https://www.alibabacloud.com/help/doc-detail/106351.html) for further details.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="releaseTime")
    def release_time(self) -> float:
        """
        The automatic release time of an Enterprise Platinum Edition instance.
        """
        return pulumi.get(self, "release_time")


@pulumi.output_type
class GetTopicsTopicResult(dict):
    def __init__(__self__, *,
                 create_time: str,
                 independent_naming: bool,
                 message_type: float,
                 owner: str,
                 relation: float,
                 relation_name: str,
                 remark: str,
                 topic: str):
        """
        :param str create_time: Time of creation.
        :param bool independent_naming: Indicates whether namespaces are available. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
        :param float message_type: The type of the message. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
        :param str owner: The ID of the topic owner, which is the Alibaba Cloud UID.
        :param float relation: The relation ID. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
        :param str relation_name: The name of the relation, for example, owner, publishable, subscribable, and publishable and subscribable.
        :param str remark: Remark of the topic.
        :param str topic: The name of the topic.
        """
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "independent_naming", independent_naming)
        pulumi.set(__self__, "message_type", message_type)
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "relation", relation)
        pulumi.set(__self__, "relation_name", relation_name)
        pulumi.set(__self__, "remark", remark)
        pulumi.set(__self__, "topic", topic)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Time of creation.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="independentNaming")
    def independent_naming(self) -> bool:
        """
        Indicates whether namespaces are available. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
        """
        return pulumi.get(self, "independent_naming")

    @property
    @pulumi.getter(name="messageType")
    def message_type(self) -> float:
        """
        The type of the message. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
        """
        return pulumi.get(self, "message_type")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        The ID of the topic owner, which is the Alibaba Cloud UID.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def relation(self) -> float:
        """
        The relation ID. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
        """
        return pulumi.get(self, "relation")

    @property
    @pulumi.getter(name="relationName")
    def relation_name(self) -> str:
        """
        The name of the relation, for example, owner, publishable, subscribable, and publishable and subscribable.
        """
        return pulumi.get(self, "relation_name")

    @property
    @pulumi.getter
    def remark(self) -> str:
        """
        Remark of the topic.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter
    def topic(self) -> str:
        """
        The name of the topic.
        """
        return pulumi.get(self, "topic")

