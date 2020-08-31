# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'ShardingInstanceMongoListArgs',
    'ShardingInstanceShardListArgs',
]

@pulumi.input_type
class ShardingInstanceMongoListArgs:
    def __init__(__self__, *,
                 node_class: pulumi.Input[str],
                 connect_string: Optional[pulumi.Input[str]] = None,
                 node_id: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[float]] = None):
        """
        :param pulumi.Input[str] node_class: -(Required) Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
        :param pulumi.Input[str] connect_string: Mongo node connection string
        :param pulumi.Input[str] node_id: The ID of the shard-node.
        :param pulumi.Input[float] port: Mongo node port
               * `shard_list`
        """
        pulumi.set(__self__, "node_class", node_class)
        if connect_string is not None:
            pulumi.set(__self__, "connect_string", connect_string)
        if node_id is not None:
            pulumi.set(__self__, "node_id", node_id)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="nodeClass")
    def node_class(self) -> pulumi.Input[str]:
        """
        -(Required) Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
        """
        return pulumi.get(self, "node_class")

    @node_class.setter
    def node_class(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_class", value)

    @property
    @pulumi.getter(name="connectString")
    def connect_string(self) -> Optional[pulumi.Input[str]]:
        """
        Mongo node connection string
        """
        return pulumi.get(self, "connect_string")

    @connect_string.setter
    def connect_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connect_string", value)

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the shard-node.
        """
        return pulumi.get(self, "node_id")

    @node_id.setter
    def node_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_id", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[float]]:
        """
        Mongo node port
        * `shard_list`
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class ShardingInstanceShardListArgs:
    def __init__(__self__, *,
                 node_class: pulumi.Input[str],
                 node_storage: pulumi.Input[float],
                 node_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] node_class: -(Required) Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
        :param pulumi.Input[float] node_storage: - Custom storage space; value range: [10, 1,000]
               - 10-GB increments. Unit: GB.
        :param pulumi.Input[str] node_id: The ID of the shard-node.
        """
        pulumi.set(__self__, "node_class", node_class)
        pulumi.set(__self__, "node_storage", node_storage)
        if node_id is not None:
            pulumi.set(__self__, "node_id", node_id)

    @property
    @pulumi.getter(name="nodeClass")
    def node_class(self) -> pulumi.Input[str]:
        """
        -(Required) Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
        """
        return pulumi.get(self, "node_class")

    @node_class.setter
    def node_class(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_class", value)

    @property
    @pulumi.getter(name="nodeStorage")
    def node_storage(self) -> pulumi.Input[float]:
        """
        - Custom storage space; value range: [10, 1,000]
        - 10-GB increments. Unit: GB.
        """
        return pulumi.get(self, "node_storage")

    @node_storage.setter
    def node_storage(self, value: pulumi.Input[float]):
        pulumi.set(self, "node_storage", value)

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the shard-node.
        """
        return pulumi.get(self, "node_id")

    @node_id.setter
    def node_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_id", value)

