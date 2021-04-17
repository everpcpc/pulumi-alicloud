# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DnatEntryArgs', 'DnatEntry']

@pulumi.input_type
class DnatEntryArgs:
    def __init__(__self__, *,
                 external_port: pulumi.Input[str],
                 internal_ip: pulumi.Input[str],
                 internal_port: pulumi.Input[str],
                 ip_protocol: pulumi.Input[str],
                 sag_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 external_ip: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DnatEntry resource.
        :param pulumi.Input[str] external_port: The public port.Value range: 1 to 65535 or "any".
        :param pulumi.Input[str] internal_ip: The destination private IP address.
        :param pulumi.Input[str] internal_port: The destination private port.Value range: 1 to 65535 or "any".
        :param pulumi.Input[str] ip_protocol: The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.
        :param pulumi.Input[str] sag_id: The ID of the SAG instance.
        :param pulumi.Input[str] type: The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses
        :param pulumi.Input[str] external_ip: The external public IP address.when "type" is "Internet",automatically identify the external ip.
        """
        pulumi.set(__self__, "external_port", external_port)
        pulumi.set(__self__, "internal_ip", internal_ip)
        pulumi.set(__self__, "internal_port", internal_port)
        pulumi.set(__self__, "ip_protocol", ip_protocol)
        pulumi.set(__self__, "sag_id", sag_id)
        pulumi.set(__self__, "type", type)
        if external_ip is not None:
            pulumi.set(__self__, "external_ip", external_ip)

    @property
    @pulumi.getter(name="externalPort")
    def external_port(self) -> pulumi.Input[str]:
        """
        The public port.Value range: 1 to 65535 or "any".
        """
        return pulumi.get(self, "external_port")

    @external_port.setter
    def external_port(self, value: pulumi.Input[str]):
        pulumi.set(self, "external_port", value)

    @property
    @pulumi.getter(name="internalIp")
    def internal_ip(self) -> pulumi.Input[str]:
        """
        The destination private IP address.
        """
        return pulumi.get(self, "internal_ip")

    @internal_ip.setter
    def internal_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "internal_ip", value)

    @property
    @pulumi.getter(name="internalPort")
    def internal_port(self) -> pulumi.Input[str]:
        """
        The destination private port.Value range: 1 to 65535 or "any".
        """
        return pulumi.get(self, "internal_port")

    @internal_port.setter
    def internal_port(self, value: pulumi.Input[str]):
        pulumi.set(self, "internal_port", value)

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> pulumi.Input[str]:
        """
        The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.
        """
        return pulumi.get(self, "ip_protocol")

    @ip_protocol.setter
    def ip_protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_protocol", value)

    @property
    @pulumi.getter(name="sagId")
    def sag_id(self) -> pulumi.Input[str]:
        """
        The ID of the SAG instance.
        """
        return pulumi.get(self, "sag_id")

    @sag_id.setter
    def sag_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "sag_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="externalIp")
    def external_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The external public IP address.when "type" is "Internet",automatically identify the external ip.
        """
        return pulumi.get(self, "external_ip")

    @external_ip.setter
    def external_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_ip", value)


@pulumi.input_type
class _DnatEntryState:
    def __init__(__self__, *,
                 external_ip: Optional[pulumi.Input[str]] = None,
                 external_port: Optional[pulumi.Input[str]] = None,
                 internal_ip: Optional[pulumi.Input[str]] = None,
                 internal_port: Optional[pulumi.Input[str]] = None,
                 ip_protocol: Optional[pulumi.Input[str]] = None,
                 sag_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DnatEntry resources.
        :param pulumi.Input[str] external_ip: The external public IP address.when "type" is "Internet",automatically identify the external ip.
        :param pulumi.Input[str] external_port: The public port.Value range: 1 to 65535 or "any".
        :param pulumi.Input[str] internal_ip: The destination private IP address.
        :param pulumi.Input[str] internal_port: The destination private port.Value range: 1 to 65535 or "any".
        :param pulumi.Input[str] ip_protocol: The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.
        :param pulumi.Input[str] sag_id: The ID of the SAG instance.
        :param pulumi.Input[str] type: The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses
        """
        if external_ip is not None:
            pulumi.set(__self__, "external_ip", external_ip)
        if external_port is not None:
            pulumi.set(__self__, "external_port", external_port)
        if internal_ip is not None:
            pulumi.set(__self__, "internal_ip", internal_ip)
        if internal_port is not None:
            pulumi.set(__self__, "internal_port", internal_port)
        if ip_protocol is not None:
            pulumi.set(__self__, "ip_protocol", ip_protocol)
        if sag_id is not None:
            pulumi.set(__self__, "sag_id", sag_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="externalIp")
    def external_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The external public IP address.when "type" is "Internet",automatically identify the external ip.
        """
        return pulumi.get(self, "external_ip")

    @external_ip.setter
    def external_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_ip", value)

    @property
    @pulumi.getter(name="externalPort")
    def external_port(self) -> Optional[pulumi.Input[str]]:
        """
        The public port.Value range: 1 to 65535 or "any".
        """
        return pulumi.get(self, "external_port")

    @external_port.setter
    def external_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_port", value)

    @property
    @pulumi.getter(name="internalIp")
    def internal_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The destination private IP address.
        """
        return pulumi.get(self, "internal_ip")

    @internal_ip.setter
    def internal_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_ip", value)

    @property
    @pulumi.getter(name="internalPort")
    def internal_port(self) -> Optional[pulumi.Input[str]]:
        """
        The destination private port.Value range: 1 to 65535 or "any".
        """
        return pulumi.get(self, "internal_port")

    @internal_port.setter
    def internal_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_port", value)

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.
        """
        return pulumi.get(self, "ip_protocol")

    @ip_protocol.setter
    def ip_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_protocol", value)

    @property
    @pulumi.getter(name="sagId")
    def sag_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the SAG instance.
        """
        return pulumi.get(self, "sag_id")

    @sag_id.setter
    def sag_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sag_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class DnatEntry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 external_ip: Optional[pulumi.Input[str]] = None,
                 external_port: Optional[pulumi.Input[str]] = None,
                 internal_ip: Optional[pulumi.Input[str]] = None,
                 internal_port: Optional[pulumi.Input[str]] = None,
                 ip_protocol: Optional[pulumi.Input[str]] = None,
                 sag_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Sag DnatEntry resource. This topic describes how to add a DNAT entry to a Smart Access Gateway (SAG) instance to enable the DNAT function. By using the DNAT function, you can forward requests received by public IP addresses to Alibaba Cloud instances according to custom mapping rules.

        For information about Sag DnatEntry and how to use it, see [What is Sag DnatEntry](https://www.alibabacloud.com/help/doc-detail/124312.htm).

        > **NOTE:** Available in 1.63.0+

        > **NOTE:** Only the following regions suppor. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.rocketmq.DnatEntry("default",
            external_ip="1.0.0.2",
            external_port="1",
            internal_ip="10.0.0.2",
            internal_port="20",
            ip_protocol="tcp",
            sag_id="sag-3rb1t3iagy3w0zgwy9",
            type="Intranet")
        ```

        ## Import

        The Sag DnatEntry can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:rocketmq/dnatEntry:DnatEntry example sag-abc123456:dnat-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] external_ip: The external public IP address.when "type" is "Internet",automatically identify the external ip.
        :param pulumi.Input[str] external_port: The public port.Value range: 1 to 65535 or "any".
        :param pulumi.Input[str] internal_ip: The destination private IP address.
        :param pulumi.Input[str] internal_port: The destination private port.Value range: 1 to 65535 or "any".
        :param pulumi.Input[str] ip_protocol: The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.
        :param pulumi.Input[str] sag_id: The ID of the SAG instance.
        :param pulumi.Input[str] type: The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DnatEntryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Sag DnatEntry resource. This topic describes how to add a DNAT entry to a Smart Access Gateway (SAG) instance to enable the DNAT function. By using the DNAT function, you can forward requests received by public IP addresses to Alibaba Cloud instances according to custom mapping rules.

        For information about Sag DnatEntry and how to use it, see [What is Sag DnatEntry](https://www.alibabacloud.com/help/doc-detail/124312.htm).

        > **NOTE:** Available in 1.63.0+

        > **NOTE:** Only the following regions suppor. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.rocketmq.DnatEntry("default",
            external_ip="1.0.0.2",
            external_port="1",
            internal_ip="10.0.0.2",
            internal_port="20",
            ip_protocol="tcp",
            sag_id="sag-3rb1t3iagy3w0zgwy9",
            type="Intranet")
        ```

        ## Import

        The Sag DnatEntry can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:rocketmq/dnatEntry:DnatEntry example sag-abc123456:dnat-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param DnatEntryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DnatEntryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 external_ip: Optional[pulumi.Input[str]] = None,
                 external_port: Optional[pulumi.Input[str]] = None,
                 internal_ip: Optional[pulumi.Input[str]] = None,
                 internal_port: Optional[pulumi.Input[str]] = None,
                 ip_protocol: Optional[pulumi.Input[str]] = None,
                 sag_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DnatEntryArgs.__new__(DnatEntryArgs)

            __props__.__dict__["external_ip"] = external_ip
            if external_port is None and not opts.urn:
                raise TypeError("Missing required property 'external_port'")
            __props__.__dict__["external_port"] = external_port
            if internal_ip is None and not opts.urn:
                raise TypeError("Missing required property 'internal_ip'")
            __props__.__dict__["internal_ip"] = internal_ip
            if internal_port is None and not opts.urn:
                raise TypeError("Missing required property 'internal_port'")
            __props__.__dict__["internal_port"] = internal_port
            if ip_protocol is None and not opts.urn:
                raise TypeError("Missing required property 'ip_protocol'")
            __props__.__dict__["ip_protocol"] = ip_protocol
            if sag_id is None and not opts.urn:
                raise TypeError("Missing required property 'sag_id'")
            __props__.__dict__["sag_id"] = sag_id
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(DnatEntry, __self__).__init__(
            'alicloud:rocketmq/dnatEntry:DnatEntry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            external_ip: Optional[pulumi.Input[str]] = None,
            external_port: Optional[pulumi.Input[str]] = None,
            internal_ip: Optional[pulumi.Input[str]] = None,
            internal_port: Optional[pulumi.Input[str]] = None,
            ip_protocol: Optional[pulumi.Input[str]] = None,
            sag_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'DnatEntry':
        """
        Get an existing DnatEntry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] external_ip: The external public IP address.when "type" is "Internet",automatically identify the external ip.
        :param pulumi.Input[str] external_port: The public port.Value range: 1 to 65535 or "any".
        :param pulumi.Input[str] internal_ip: The destination private IP address.
        :param pulumi.Input[str] internal_port: The destination private port.Value range: 1 to 65535 or "any".
        :param pulumi.Input[str] ip_protocol: The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.
        :param pulumi.Input[str] sag_id: The ID of the SAG instance.
        :param pulumi.Input[str] type: The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DnatEntryState.__new__(_DnatEntryState)

        __props__.__dict__["external_ip"] = external_ip
        __props__.__dict__["external_port"] = external_port
        __props__.__dict__["internal_ip"] = internal_ip
        __props__.__dict__["internal_port"] = internal_port
        __props__.__dict__["ip_protocol"] = ip_protocol
        __props__.__dict__["sag_id"] = sag_id
        __props__.__dict__["type"] = type
        return DnatEntry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="externalIp")
    def external_ip(self) -> pulumi.Output[Optional[str]]:
        """
        The external public IP address.when "type" is "Internet",automatically identify the external ip.
        """
        return pulumi.get(self, "external_ip")

    @property
    @pulumi.getter(name="externalPort")
    def external_port(self) -> pulumi.Output[str]:
        """
        The public port.Value range: 1 to 65535 or "any".
        """
        return pulumi.get(self, "external_port")

    @property
    @pulumi.getter(name="internalIp")
    def internal_ip(self) -> pulumi.Output[str]:
        """
        The destination private IP address.
        """
        return pulumi.get(self, "internal_ip")

    @property
    @pulumi.getter(name="internalPort")
    def internal_port(self) -> pulumi.Output[str]:
        """
        The destination private port.Value range: 1 to 65535 or "any".
        """
        return pulumi.get(self, "internal_port")

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> pulumi.Output[str]:
        """
        The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.
        """
        return pulumi.get(self, "ip_protocol")

    @property
    @pulumi.getter(name="sagId")
    def sag_id(self) -> pulumi.Output[str]:
        """
        The ID of the SAG instance.
        """
        return pulumi.get(self, "sag_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses
        """
        return pulumi.get(self, "type")

