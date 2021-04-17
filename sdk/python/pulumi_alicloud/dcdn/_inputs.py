# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'DomainSourceArgs',
]

@pulumi.input_type
class DomainSourceArgs:
    def __init__(__self__, *,
                 content: pulumi.Input[str],
                 type: pulumi.Input[str],
                 port: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] content: The origin address.
        :param pulumi.Input[str] type: The type of the origin. Valid values:
               `ipaddr`: The origin is configured using an IP address.
               `domain`: The origin is configured using a domain name.
               `oss`: The origin is configured using the Internet domain name of an Alibaba Cloud Object Storage Service (OSS) bucket.
        :param pulumi.Input[int] port: The port number. Valid values: `443` and `80`. Default to `80`.
        :param pulumi.Input[str] priority: The priority of the origin if multiple origins are specified. Default to `20`.
        :param pulumi.Input[str] weight: The weight of the origin if multiple origins are specified. Default to `10`.
        """
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "type", type)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        """
        The origin address.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the origin. Valid values:
        `ipaddr`: The origin is configured using an IP address.
        `domain`: The origin is configured using a domain name.
        `oss`: The origin is configured using the Internet domain name of an Alibaba Cloud Object Storage Service (OSS) bucket.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port number. Valid values: `443` and `80`. Default to `80`.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[str]]:
        """
        The priority of the origin if multiple origins are specified. Default to `20`.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[str]]:
        """
        The weight of the origin if multiple origins are specified. Default to `10`.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "weight", value)


