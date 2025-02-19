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
from ._inputs import *

__all__ = ['FirewallVpcFirewallArgs', 'FirewallVpcFirewall']

@pulumi.input_type
class FirewallVpcFirewallArgs:
    def __init__(__self__, *,
                 local_vpc: pulumi.Input['FirewallVpcFirewallLocalVpcArgs'],
                 peer_vpc: pulumi.Input['FirewallVpcFirewallPeerVpcArgs'],
                 status: pulumi.Input[str],
                 vpc_firewall_name: pulumi.Input[str],
                 lang: Optional[pulumi.Input[str]] = None,
                 member_uid: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallVpcFirewall resource.
        :param pulumi.Input['FirewallVpcFirewallLocalVpcArgs'] local_vpc: The details of the local VPC. See the following `Block LocalVpc`.
        :param pulumi.Input['FirewallVpcFirewallPeerVpcArgs'] peer_vpc: The details of the peer VPC. See the following `Block PeerVpc`.
        :param pulumi.Input[str] status: The status of the resource
        :param pulumi.Input[str] vpc_firewall_name: The name of the VPC firewall instance.
        :param pulumi.Input[str] lang: The language type of the requested and received messages. Value:**zh** (default): Chinese.**en**: English.
        :param pulumi.Input[str] member_uid: The UID of the Alibaba Cloud member account.
        """
        pulumi.set(__self__, "local_vpc", local_vpc)
        pulumi.set(__self__, "peer_vpc", peer_vpc)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "vpc_firewall_name", vpc_firewall_name)
        if lang is not None:
            pulumi.set(__self__, "lang", lang)
        if member_uid is not None:
            pulumi.set(__self__, "member_uid", member_uid)

    @property
    @pulumi.getter(name="localVpc")
    def local_vpc(self) -> pulumi.Input['FirewallVpcFirewallLocalVpcArgs']:
        """
        The details of the local VPC. See the following `Block LocalVpc`.
        """
        return pulumi.get(self, "local_vpc")

    @local_vpc.setter
    def local_vpc(self, value: pulumi.Input['FirewallVpcFirewallLocalVpcArgs']):
        pulumi.set(self, "local_vpc", value)

    @property
    @pulumi.getter(name="peerVpc")
    def peer_vpc(self) -> pulumi.Input['FirewallVpcFirewallPeerVpcArgs']:
        """
        The details of the peer VPC. See the following `Block PeerVpc`.
        """
        return pulumi.get(self, "peer_vpc")

    @peer_vpc.setter
    def peer_vpc(self, value: pulumi.Input['FirewallVpcFirewallPeerVpcArgs']):
        pulumi.set(self, "peer_vpc", value)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[str]:
        """
        The status of the resource
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vpcFirewallName")
    def vpc_firewall_name(self) -> pulumi.Input[str]:
        """
        The name of the VPC firewall instance.
        """
        return pulumi.get(self, "vpc_firewall_name")

    @vpc_firewall_name.setter
    def vpc_firewall_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_firewall_name", value)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        The language type of the requested and received messages. Value:**zh** (default): Chinese.**en**: English.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)

    @property
    @pulumi.getter(name="memberUid")
    def member_uid(self) -> Optional[pulumi.Input[str]]:
        """
        The UID of the Alibaba Cloud member account.
        """
        return pulumi.get(self, "member_uid")

    @member_uid.setter
    def member_uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "member_uid", value)


@pulumi.input_type
class _FirewallVpcFirewallState:
    def __init__(__self__, *,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 connect_type: Optional[pulumi.Input[str]] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 local_vpc: Optional[pulumi.Input['FirewallVpcFirewallLocalVpcArgs']] = None,
                 member_uid: Optional[pulumi.Input[str]] = None,
                 peer_vpc: Optional[pulumi.Input['FirewallVpcFirewallPeerVpcArgs']] = None,
                 region_status: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vpc_firewall_id: Optional[pulumi.Input[str]] = None,
                 vpc_firewall_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallVpcFirewall resources.
        :param pulumi.Input[int] bandwidth: Bandwidth specifications for high-speed channels. Unit: Mbps.
        :param pulumi.Input[str] connect_type: The communication type of the VPC firewall. Valid value: **expressconnect**, which indicates Express Connect.
        :param pulumi.Input[str] lang: The language type of the requested and received messages. Value:**zh** (default): Chinese.**en**: English.
        :param pulumi.Input['FirewallVpcFirewallLocalVpcArgs'] local_vpc: The details of the local VPC. See the following `Block LocalVpc`.
        :param pulumi.Input[str] member_uid: The UID of the Alibaba Cloud member account.
        :param pulumi.Input['FirewallVpcFirewallPeerVpcArgs'] peer_vpc: The details of the peer VPC. See the following `Block PeerVpc`.
        :param pulumi.Input[str] region_status: The region is open. Value:-**enable**: is enabled, indicating that VPC firewall can be configured in this region.-**disable**: indicates that VPC firewall cannot be configured in this region.
        :param pulumi.Input[str] status: The status of the resource
        :param pulumi.Input[str] vpc_firewall_id: The ID of the VPC firewall instance.
        :param pulumi.Input[str] vpc_firewall_name: The name of the VPC firewall instance.
        """
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if connect_type is not None:
            pulumi.set(__self__, "connect_type", connect_type)
        if lang is not None:
            pulumi.set(__self__, "lang", lang)
        if local_vpc is not None:
            pulumi.set(__self__, "local_vpc", local_vpc)
        if member_uid is not None:
            pulumi.set(__self__, "member_uid", member_uid)
        if peer_vpc is not None:
            pulumi.set(__self__, "peer_vpc", peer_vpc)
        if region_status is not None:
            pulumi.set(__self__, "region_status", region_status)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vpc_firewall_id is not None:
            pulumi.set(__self__, "vpc_firewall_id", vpc_firewall_id)
        if vpc_firewall_name is not None:
            pulumi.set(__self__, "vpc_firewall_name", vpc_firewall_name)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Bandwidth specifications for high-speed channels. Unit: Mbps.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="connectType")
    def connect_type(self) -> Optional[pulumi.Input[str]]:
        """
        The communication type of the VPC firewall. Valid value: **expressconnect**, which indicates Express Connect.
        """
        return pulumi.get(self, "connect_type")

    @connect_type.setter
    def connect_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connect_type", value)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        The language type of the requested and received messages. Value:**zh** (default): Chinese.**en**: English.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)

    @property
    @pulumi.getter(name="localVpc")
    def local_vpc(self) -> Optional[pulumi.Input['FirewallVpcFirewallLocalVpcArgs']]:
        """
        The details of the local VPC. See the following `Block LocalVpc`.
        """
        return pulumi.get(self, "local_vpc")

    @local_vpc.setter
    def local_vpc(self, value: Optional[pulumi.Input['FirewallVpcFirewallLocalVpcArgs']]):
        pulumi.set(self, "local_vpc", value)

    @property
    @pulumi.getter(name="memberUid")
    def member_uid(self) -> Optional[pulumi.Input[str]]:
        """
        The UID of the Alibaba Cloud member account.
        """
        return pulumi.get(self, "member_uid")

    @member_uid.setter
    def member_uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "member_uid", value)

    @property
    @pulumi.getter(name="peerVpc")
    def peer_vpc(self) -> Optional[pulumi.Input['FirewallVpcFirewallPeerVpcArgs']]:
        """
        The details of the peer VPC. See the following `Block PeerVpc`.
        """
        return pulumi.get(self, "peer_vpc")

    @peer_vpc.setter
    def peer_vpc(self, value: Optional[pulumi.Input['FirewallVpcFirewallPeerVpcArgs']]):
        pulumi.set(self, "peer_vpc", value)

    @property
    @pulumi.getter(name="regionStatus")
    def region_status(self) -> Optional[pulumi.Input[str]]:
        """
        The region is open. Value:-**enable**: is enabled, indicating that VPC firewall can be configured in this region.-**disable**: indicates that VPC firewall cannot be configured in this region.
        """
        return pulumi.get(self, "region_status")

    @region_status.setter
    def region_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_status", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vpcFirewallId")
    def vpc_firewall_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPC firewall instance.
        """
        return pulumi.get(self, "vpc_firewall_id")

    @vpc_firewall_id.setter
    def vpc_firewall_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_firewall_id", value)

    @property
    @pulumi.getter(name="vpcFirewallName")
    def vpc_firewall_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the VPC firewall instance.
        """
        return pulumi.get(self, "vpc_firewall_name")

    @vpc_firewall_name.setter
    def vpc_firewall_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_firewall_name", value)


class FirewallVpcFirewall(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 local_vpc: Optional[pulumi.Input[pulumi.InputType['FirewallVpcFirewallLocalVpcArgs']]] = None,
                 member_uid: Optional[pulumi.Input[str]] = None,
                 peer_vpc: Optional[pulumi.Input[pulumi.InputType['FirewallVpcFirewallPeerVpcArgs']]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vpc_firewall_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Cloud Firewall Vpc Firewall resource.

        For information about Cloud Firewall Vpc Firewall and how to use it, see [What is Vpc Firewall](https://help.aliyun.com/document_detail/342893.html).

        > **NOTE:** Available in v1.194.0+.

        ## Import

        Cloud Firewall Vpc Firewall can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cloudfirewall/firewallVpcFirewall:FirewallVpcFirewall example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] lang: The language type of the requested and received messages. Value:**zh** (default): Chinese.**en**: English.
        :param pulumi.Input[pulumi.InputType['FirewallVpcFirewallLocalVpcArgs']] local_vpc: The details of the local VPC. See the following `Block LocalVpc`.
        :param pulumi.Input[str] member_uid: The UID of the Alibaba Cloud member account.
        :param pulumi.Input[pulumi.InputType['FirewallVpcFirewallPeerVpcArgs']] peer_vpc: The details of the peer VPC. See the following `Block PeerVpc`.
        :param pulumi.Input[str] status: The status of the resource
        :param pulumi.Input[str] vpc_firewall_name: The name of the VPC firewall instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallVpcFirewallArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Cloud Firewall Vpc Firewall resource.

        For information about Cloud Firewall Vpc Firewall and how to use it, see [What is Vpc Firewall](https://help.aliyun.com/document_detail/342893.html).

        > **NOTE:** Available in v1.194.0+.

        ## Import

        Cloud Firewall Vpc Firewall can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cloudfirewall/firewallVpcFirewall:FirewallVpcFirewall example <id>
        ```

        :param str resource_name: The name of the resource.
        :param FirewallVpcFirewallArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallVpcFirewallArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 local_vpc: Optional[pulumi.Input[pulumi.InputType['FirewallVpcFirewallLocalVpcArgs']]] = None,
                 member_uid: Optional[pulumi.Input[str]] = None,
                 peer_vpc: Optional[pulumi.Input[pulumi.InputType['FirewallVpcFirewallPeerVpcArgs']]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vpc_firewall_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallVpcFirewallArgs.__new__(FirewallVpcFirewallArgs)

            __props__.__dict__["lang"] = lang
            if local_vpc is None and not opts.urn:
                raise TypeError("Missing required property 'local_vpc'")
            __props__.__dict__["local_vpc"] = local_vpc
            __props__.__dict__["member_uid"] = member_uid
            if peer_vpc is None and not opts.urn:
                raise TypeError("Missing required property 'peer_vpc'")
            __props__.__dict__["peer_vpc"] = peer_vpc
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
            if vpc_firewall_name is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_firewall_name'")
            __props__.__dict__["vpc_firewall_name"] = vpc_firewall_name
            __props__.__dict__["bandwidth"] = None
            __props__.__dict__["connect_type"] = None
            __props__.__dict__["region_status"] = None
            __props__.__dict__["vpc_firewall_id"] = None
        super(FirewallVpcFirewall, __self__).__init__(
            'alicloud:cloudfirewall/firewallVpcFirewall:FirewallVpcFirewall',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bandwidth: Optional[pulumi.Input[int]] = None,
            connect_type: Optional[pulumi.Input[str]] = None,
            lang: Optional[pulumi.Input[str]] = None,
            local_vpc: Optional[pulumi.Input[pulumi.InputType['FirewallVpcFirewallLocalVpcArgs']]] = None,
            member_uid: Optional[pulumi.Input[str]] = None,
            peer_vpc: Optional[pulumi.Input[pulumi.InputType['FirewallVpcFirewallPeerVpcArgs']]] = None,
            region_status: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vpc_firewall_id: Optional[pulumi.Input[str]] = None,
            vpc_firewall_name: Optional[pulumi.Input[str]] = None) -> 'FirewallVpcFirewall':
        """
        Get an existing FirewallVpcFirewall resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: Bandwidth specifications for high-speed channels. Unit: Mbps.
        :param pulumi.Input[str] connect_type: The communication type of the VPC firewall. Valid value: **expressconnect**, which indicates Express Connect.
        :param pulumi.Input[str] lang: The language type of the requested and received messages. Value:**zh** (default): Chinese.**en**: English.
        :param pulumi.Input[pulumi.InputType['FirewallVpcFirewallLocalVpcArgs']] local_vpc: The details of the local VPC. See the following `Block LocalVpc`.
        :param pulumi.Input[str] member_uid: The UID of the Alibaba Cloud member account.
        :param pulumi.Input[pulumi.InputType['FirewallVpcFirewallPeerVpcArgs']] peer_vpc: The details of the peer VPC. See the following `Block PeerVpc`.
        :param pulumi.Input[str] region_status: The region is open. Value:-**enable**: is enabled, indicating that VPC firewall can be configured in this region.-**disable**: indicates that VPC firewall cannot be configured in this region.
        :param pulumi.Input[str] status: The status of the resource
        :param pulumi.Input[str] vpc_firewall_id: The ID of the VPC firewall instance.
        :param pulumi.Input[str] vpc_firewall_name: The name of the VPC firewall instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallVpcFirewallState.__new__(_FirewallVpcFirewallState)

        __props__.__dict__["bandwidth"] = bandwidth
        __props__.__dict__["connect_type"] = connect_type
        __props__.__dict__["lang"] = lang
        __props__.__dict__["local_vpc"] = local_vpc
        __props__.__dict__["member_uid"] = member_uid
        __props__.__dict__["peer_vpc"] = peer_vpc
        __props__.__dict__["region_status"] = region_status
        __props__.__dict__["status"] = status
        __props__.__dict__["vpc_firewall_id"] = vpc_firewall_id
        __props__.__dict__["vpc_firewall_name"] = vpc_firewall_name
        return FirewallVpcFirewall(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Output[int]:
        """
        Bandwidth specifications for high-speed channels. Unit: Mbps.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="connectType")
    def connect_type(self) -> pulumi.Output[str]:
        """
        The communication type of the VPC firewall. Valid value: **expressconnect**, which indicates Express Connect.
        """
        return pulumi.get(self, "connect_type")

    @property
    @pulumi.getter
    def lang(self) -> pulumi.Output[str]:
        """
        The language type of the requested and received messages. Value:**zh** (default): Chinese.**en**: English.
        """
        return pulumi.get(self, "lang")

    @property
    @pulumi.getter(name="localVpc")
    def local_vpc(self) -> pulumi.Output['outputs.FirewallVpcFirewallLocalVpc']:
        """
        The details of the local VPC. See the following `Block LocalVpc`.
        """
        return pulumi.get(self, "local_vpc")

    @property
    @pulumi.getter(name="memberUid")
    def member_uid(self) -> pulumi.Output[Optional[str]]:
        """
        The UID of the Alibaba Cloud member account.
        """
        return pulumi.get(self, "member_uid")

    @property
    @pulumi.getter(name="peerVpc")
    def peer_vpc(self) -> pulumi.Output['outputs.FirewallVpcFirewallPeerVpc']:
        """
        The details of the peer VPC. See the following `Block PeerVpc`.
        """
        return pulumi.get(self, "peer_vpc")

    @property
    @pulumi.getter(name="regionStatus")
    def region_status(self) -> pulumi.Output[str]:
        """
        The region is open. Value:-**enable**: is enabled, indicating that VPC firewall can be configured in this region.-**disable**: indicates that VPC firewall cannot be configured in this region.
        """
        return pulumi.get(self, "region_status")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpcFirewallId")
    def vpc_firewall_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPC firewall instance.
        """
        return pulumi.get(self, "vpc_firewall_id")

    @property
    @pulumi.getter(name="vpcFirewallName")
    def vpc_firewall_name(self) -> pulumi.Output[str]:
        """
        The name of the VPC firewall instance.
        """
        return pulumi.get(self, "vpc_firewall_name")

