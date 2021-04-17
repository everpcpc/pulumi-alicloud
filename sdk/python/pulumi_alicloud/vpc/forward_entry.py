# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ForwardEntryArgs', 'ForwardEntry']

@pulumi.input_type
class ForwardEntryArgs:
    def __init__(__self__, *,
                 external_ip: pulumi.Input[str],
                 external_port: pulumi.Input[str],
                 forward_table_id: pulumi.Input[str],
                 internal_ip: pulumi.Input[str],
                 internal_port: pulumi.Input[str],
                 ip_protocol: pulumi.Input[str],
                 forward_entry_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port_break: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ForwardEntry resource.
        :param pulumi.Input[str] external_ip: The external ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        :param pulumi.Input[str] external_port: The external port, valid value is 1~65535|any.
        :param pulumi.Input[str] forward_table_id: The value can get from `vpc.NatGateway` Attributes "forward_table_ids".
        :param pulumi.Input[str] internal_ip: The internal ip, must a private ip.
        :param pulumi.Input[str] internal_port: The internal port, valid value is 1~65535|any.
        :param pulumi.Input[str] ip_protocol: The ip protocal, valid value is tcp|udp|any.
        :param pulumi.Input[str] forward_entry_name: The name of forward entry.
        :param pulumi.Input[str] name: Field `name` has been deprecated from provider version 1.119.1. New field `forward_entry_name` instead.
        :param pulumi.Input[bool] port_break: Specifies whether to remove limits on the port range. Default value is `false`.
        """
        pulumi.set(__self__, "external_ip", external_ip)
        pulumi.set(__self__, "external_port", external_port)
        pulumi.set(__self__, "forward_table_id", forward_table_id)
        pulumi.set(__self__, "internal_ip", internal_ip)
        pulumi.set(__self__, "internal_port", internal_port)
        pulumi.set(__self__, "ip_protocol", ip_protocol)
        if forward_entry_name is not None:
            pulumi.set(__self__, "forward_entry_name", forward_entry_name)
        if name is not None:
            warnings.warn("""Field 'name' has been deprecated from provider version 1.119.1. New field 'forward_entry_name' instead.""", DeprecationWarning)
            pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from provider version 1.119.1. New field 'forward_entry_name' instead.""")
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port_break is not None:
            pulumi.set(__self__, "port_break", port_break)

    @property
    @pulumi.getter(name="externalIp")
    def external_ip(self) -> pulumi.Input[str]:
        """
        The external ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        """
        return pulumi.get(self, "external_ip")

    @external_ip.setter
    def external_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "external_ip", value)

    @property
    @pulumi.getter(name="externalPort")
    def external_port(self) -> pulumi.Input[str]:
        """
        The external port, valid value is 1~65535|any.
        """
        return pulumi.get(self, "external_port")

    @external_port.setter
    def external_port(self, value: pulumi.Input[str]):
        pulumi.set(self, "external_port", value)

    @property
    @pulumi.getter(name="forwardTableId")
    def forward_table_id(self) -> pulumi.Input[str]:
        """
        The value can get from `vpc.NatGateway` Attributes "forward_table_ids".
        """
        return pulumi.get(self, "forward_table_id")

    @forward_table_id.setter
    def forward_table_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "forward_table_id", value)

    @property
    @pulumi.getter(name="internalIp")
    def internal_ip(self) -> pulumi.Input[str]:
        """
        The internal ip, must a private ip.
        """
        return pulumi.get(self, "internal_ip")

    @internal_ip.setter
    def internal_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "internal_ip", value)

    @property
    @pulumi.getter(name="internalPort")
    def internal_port(self) -> pulumi.Input[str]:
        """
        The internal port, valid value is 1~65535|any.
        """
        return pulumi.get(self, "internal_port")

    @internal_port.setter
    def internal_port(self, value: pulumi.Input[str]):
        pulumi.set(self, "internal_port", value)

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> pulumi.Input[str]:
        """
        The ip protocal, valid value is tcp|udp|any.
        """
        return pulumi.get(self, "ip_protocol")

    @ip_protocol.setter
    def ip_protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_protocol", value)

    @property
    @pulumi.getter(name="forwardEntryName")
    def forward_entry_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of forward entry.
        """
        return pulumi.get(self, "forward_entry_name")

    @forward_entry_name.setter
    def forward_entry_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forward_entry_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Field `name` has been deprecated from provider version 1.119.1. New field `forward_entry_name` instead.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="portBreak")
    def port_break(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to remove limits on the port range. Default value is `false`.
        """
        return pulumi.get(self, "port_break")

    @port_break.setter
    def port_break(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "port_break", value)


@pulumi.input_type
class _ForwardEntryState:
    def __init__(__self__, *,
                 external_ip: Optional[pulumi.Input[str]] = None,
                 external_port: Optional[pulumi.Input[str]] = None,
                 forward_entry_id: Optional[pulumi.Input[str]] = None,
                 forward_entry_name: Optional[pulumi.Input[str]] = None,
                 forward_table_id: Optional[pulumi.Input[str]] = None,
                 internal_ip: Optional[pulumi.Input[str]] = None,
                 internal_port: Optional[pulumi.Input[str]] = None,
                 ip_protocol: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port_break: Optional[pulumi.Input[bool]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ForwardEntry resources.
        :param pulumi.Input[str] external_ip: The external ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        :param pulumi.Input[str] external_port: The external port, valid value is 1~65535|any.
        :param pulumi.Input[str] forward_entry_id: The id of the forward entry on the server.
        :param pulumi.Input[str] forward_entry_name: The name of forward entry.
        :param pulumi.Input[str] forward_table_id: The value can get from `vpc.NatGateway` Attributes "forward_table_ids".
        :param pulumi.Input[str] internal_ip: The internal ip, must a private ip.
        :param pulumi.Input[str] internal_port: The internal port, valid value is 1~65535|any.
        :param pulumi.Input[str] ip_protocol: The ip protocal, valid value is tcp|udp|any.
        :param pulumi.Input[str] name: Field `name` has been deprecated from provider version 1.119.1. New field `forward_entry_name` instead.
        :param pulumi.Input[bool] port_break: Specifies whether to remove limits on the port range. Default value is `false`.
        :param pulumi.Input[str] status: (Available in 1.119.1+) The status of forward entry.
        """
        if external_ip is not None:
            pulumi.set(__self__, "external_ip", external_ip)
        if external_port is not None:
            pulumi.set(__self__, "external_port", external_port)
        if forward_entry_id is not None:
            pulumi.set(__self__, "forward_entry_id", forward_entry_id)
        if forward_entry_name is not None:
            pulumi.set(__self__, "forward_entry_name", forward_entry_name)
        if forward_table_id is not None:
            pulumi.set(__self__, "forward_table_id", forward_table_id)
        if internal_ip is not None:
            pulumi.set(__self__, "internal_ip", internal_ip)
        if internal_port is not None:
            pulumi.set(__self__, "internal_port", internal_port)
        if ip_protocol is not None:
            pulumi.set(__self__, "ip_protocol", ip_protocol)
        if name is not None:
            warnings.warn("""Field 'name' has been deprecated from provider version 1.119.1. New field 'forward_entry_name' instead.""", DeprecationWarning)
            pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from provider version 1.119.1. New field 'forward_entry_name' instead.""")
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port_break is not None:
            pulumi.set(__self__, "port_break", port_break)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="externalIp")
    def external_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The external ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        """
        return pulumi.get(self, "external_ip")

    @external_ip.setter
    def external_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_ip", value)

    @property
    @pulumi.getter(name="externalPort")
    def external_port(self) -> Optional[pulumi.Input[str]]:
        """
        The external port, valid value is 1~65535|any.
        """
        return pulumi.get(self, "external_port")

    @external_port.setter
    def external_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_port", value)

    @property
    @pulumi.getter(name="forwardEntryId")
    def forward_entry_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the forward entry on the server.
        """
        return pulumi.get(self, "forward_entry_id")

    @forward_entry_id.setter
    def forward_entry_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forward_entry_id", value)

    @property
    @pulumi.getter(name="forwardEntryName")
    def forward_entry_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of forward entry.
        """
        return pulumi.get(self, "forward_entry_name")

    @forward_entry_name.setter
    def forward_entry_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forward_entry_name", value)

    @property
    @pulumi.getter(name="forwardTableId")
    def forward_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        The value can get from `vpc.NatGateway` Attributes "forward_table_ids".
        """
        return pulumi.get(self, "forward_table_id")

    @forward_table_id.setter
    def forward_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forward_table_id", value)

    @property
    @pulumi.getter(name="internalIp")
    def internal_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The internal ip, must a private ip.
        """
        return pulumi.get(self, "internal_ip")

    @internal_ip.setter
    def internal_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_ip", value)

    @property
    @pulumi.getter(name="internalPort")
    def internal_port(self) -> Optional[pulumi.Input[str]]:
        """
        The internal port, valid value is 1~65535|any.
        """
        return pulumi.get(self, "internal_port")

    @internal_port.setter
    def internal_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_port", value)

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The ip protocal, valid value is tcp|udp|any.
        """
        return pulumi.get(self, "ip_protocol")

    @ip_protocol.setter
    def ip_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_protocol", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Field `name` has been deprecated from provider version 1.119.1. New field `forward_entry_name` instead.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="portBreak")
    def port_break(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to remove limits on the port range. Default value is `false`.
        """
        return pulumi.get(self, "port_break")

    @port_break.setter
    def port_break(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "port_break", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        (Available in 1.119.1+) The status of forward entry.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class ForwardEntry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 external_ip: Optional[pulumi.Input[str]] = None,
                 external_port: Optional[pulumi.Input[str]] = None,
                 forward_entry_name: Optional[pulumi.Input[str]] = None,
                 forward_table_id: Optional[pulumi.Input[str]] = None,
                 internal_ip: Optional[pulumi.Input[str]] = None,
                 internal_port: Optional[pulumi.Input[str]] = None,
                 ip_protocol: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port_break: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a forward resource.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "forward-entry-example-name"
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name=name,
            cidr_block="172.16.0.0/12")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/21",
            zone_id=default_zones.zones[0].id,
            vswitch_name=name)
        default_nat_gateway = alicloud.vpc.NatGateway("defaultNatGateway",
            vpc_id=default_network.id,
            specification="Small")
        default_eip = alicloud.ecs.Eip("defaultEip")
        default_eip_association = alicloud.ecs.EipAssociation("defaultEipAssociation",
            allocation_id=default_eip.id,
            instance_id=default_nat_gateway.id)
        default_forward_entry = alicloud.vpc.ForwardEntry("defaultForwardEntry",
            forward_table_id=default_nat_gateway.forward_table_ids,
            external_ip=default_eip.ip_address,
            external_port="80",
            ip_protocol="tcp",
            internal_ip="172.16.0.3",
            internal_port="8080")
        ```

        ## Import

        Forward Entry can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/forwardEntry:ForwardEntry foo ftb-1aece3:fwd-232ce2
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] external_ip: The external ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        :param pulumi.Input[str] external_port: The external port, valid value is 1~65535|any.
        :param pulumi.Input[str] forward_entry_name: The name of forward entry.
        :param pulumi.Input[str] forward_table_id: The value can get from `vpc.NatGateway` Attributes "forward_table_ids".
        :param pulumi.Input[str] internal_ip: The internal ip, must a private ip.
        :param pulumi.Input[str] internal_port: The internal port, valid value is 1~65535|any.
        :param pulumi.Input[str] ip_protocol: The ip protocal, valid value is tcp|udp|any.
        :param pulumi.Input[str] name: Field `name` has been deprecated from provider version 1.119.1. New field `forward_entry_name` instead.
        :param pulumi.Input[bool] port_break: Specifies whether to remove limits on the port range. Default value is `false`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ForwardEntryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a forward resource.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "forward-entry-example-name"
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name=name,
            cidr_block="172.16.0.0/12")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/21",
            zone_id=default_zones.zones[0].id,
            vswitch_name=name)
        default_nat_gateway = alicloud.vpc.NatGateway("defaultNatGateway",
            vpc_id=default_network.id,
            specification="Small")
        default_eip = alicloud.ecs.Eip("defaultEip")
        default_eip_association = alicloud.ecs.EipAssociation("defaultEipAssociation",
            allocation_id=default_eip.id,
            instance_id=default_nat_gateway.id)
        default_forward_entry = alicloud.vpc.ForwardEntry("defaultForwardEntry",
            forward_table_id=default_nat_gateway.forward_table_ids,
            external_ip=default_eip.ip_address,
            external_port="80",
            ip_protocol="tcp",
            internal_ip="172.16.0.3",
            internal_port="8080")
        ```

        ## Import

        Forward Entry can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/forwardEntry:ForwardEntry foo ftb-1aece3:fwd-232ce2
        ```

        :param str resource_name: The name of the resource.
        :param ForwardEntryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ForwardEntryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 external_ip: Optional[pulumi.Input[str]] = None,
                 external_port: Optional[pulumi.Input[str]] = None,
                 forward_entry_name: Optional[pulumi.Input[str]] = None,
                 forward_table_id: Optional[pulumi.Input[str]] = None,
                 internal_ip: Optional[pulumi.Input[str]] = None,
                 internal_port: Optional[pulumi.Input[str]] = None,
                 ip_protocol: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port_break: Optional[pulumi.Input[bool]] = None,
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
            __props__ = ForwardEntryArgs.__new__(ForwardEntryArgs)

            if external_ip is None and not opts.urn:
                raise TypeError("Missing required property 'external_ip'")
            __props__.__dict__["external_ip"] = external_ip
            if external_port is None and not opts.urn:
                raise TypeError("Missing required property 'external_port'")
            __props__.__dict__["external_port"] = external_port
            __props__.__dict__["forward_entry_name"] = forward_entry_name
            if forward_table_id is None and not opts.urn:
                raise TypeError("Missing required property 'forward_table_id'")
            __props__.__dict__["forward_table_id"] = forward_table_id
            if internal_ip is None and not opts.urn:
                raise TypeError("Missing required property 'internal_ip'")
            __props__.__dict__["internal_ip"] = internal_ip
            if internal_port is None and not opts.urn:
                raise TypeError("Missing required property 'internal_port'")
            __props__.__dict__["internal_port"] = internal_port
            if ip_protocol is None and not opts.urn:
                raise TypeError("Missing required property 'ip_protocol'")
            __props__.__dict__["ip_protocol"] = ip_protocol
            if name is not None and not opts.urn:
                warnings.warn("""Field 'name' has been deprecated from provider version 1.119.1. New field 'forward_entry_name' instead.""", DeprecationWarning)
                pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from provider version 1.119.1. New field 'forward_entry_name' instead.""")
            __props__.__dict__["name"] = name
            __props__.__dict__["port_break"] = port_break
            __props__.__dict__["forward_entry_id"] = None
            __props__.__dict__["status"] = None
        super(ForwardEntry, __self__).__init__(
            'alicloud:vpc/forwardEntry:ForwardEntry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            external_ip: Optional[pulumi.Input[str]] = None,
            external_port: Optional[pulumi.Input[str]] = None,
            forward_entry_id: Optional[pulumi.Input[str]] = None,
            forward_entry_name: Optional[pulumi.Input[str]] = None,
            forward_table_id: Optional[pulumi.Input[str]] = None,
            internal_ip: Optional[pulumi.Input[str]] = None,
            internal_port: Optional[pulumi.Input[str]] = None,
            ip_protocol: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port_break: Optional[pulumi.Input[bool]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'ForwardEntry':
        """
        Get an existing ForwardEntry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] external_ip: The external ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        :param pulumi.Input[str] external_port: The external port, valid value is 1~65535|any.
        :param pulumi.Input[str] forward_entry_id: The id of the forward entry on the server.
        :param pulumi.Input[str] forward_entry_name: The name of forward entry.
        :param pulumi.Input[str] forward_table_id: The value can get from `vpc.NatGateway` Attributes "forward_table_ids".
        :param pulumi.Input[str] internal_ip: The internal ip, must a private ip.
        :param pulumi.Input[str] internal_port: The internal port, valid value is 1~65535|any.
        :param pulumi.Input[str] ip_protocol: The ip protocal, valid value is tcp|udp|any.
        :param pulumi.Input[str] name: Field `name` has been deprecated from provider version 1.119.1. New field `forward_entry_name` instead.
        :param pulumi.Input[bool] port_break: Specifies whether to remove limits on the port range. Default value is `false`.
        :param pulumi.Input[str] status: (Available in 1.119.1+) The status of forward entry.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ForwardEntryState.__new__(_ForwardEntryState)

        __props__.__dict__["external_ip"] = external_ip
        __props__.__dict__["external_port"] = external_port
        __props__.__dict__["forward_entry_id"] = forward_entry_id
        __props__.__dict__["forward_entry_name"] = forward_entry_name
        __props__.__dict__["forward_table_id"] = forward_table_id
        __props__.__dict__["internal_ip"] = internal_ip
        __props__.__dict__["internal_port"] = internal_port
        __props__.__dict__["ip_protocol"] = ip_protocol
        __props__.__dict__["name"] = name
        __props__.__dict__["port_break"] = port_break
        __props__.__dict__["status"] = status
        return ForwardEntry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="externalIp")
    def external_ip(self) -> pulumi.Output[str]:
        """
        The external ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        """
        return pulumi.get(self, "external_ip")

    @property
    @pulumi.getter(name="externalPort")
    def external_port(self) -> pulumi.Output[str]:
        """
        The external port, valid value is 1~65535|any.
        """
        return pulumi.get(self, "external_port")

    @property
    @pulumi.getter(name="forwardEntryId")
    def forward_entry_id(self) -> pulumi.Output[str]:
        """
        The id of the forward entry on the server.
        """
        return pulumi.get(self, "forward_entry_id")

    @property
    @pulumi.getter(name="forwardEntryName")
    def forward_entry_name(self) -> pulumi.Output[str]:
        """
        The name of forward entry.
        """
        return pulumi.get(self, "forward_entry_name")

    @property
    @pulumi.getter(name="forwardTableId")
    def forward_table_id(self) -> pulumi.Output[str]:
        """
        The value can get from `vpc.NatGateway` Attributes "forward_table_ids".
        """
        return pulumi.get(self, "forward_table_id")

    @property
    @pulumi.getter(name="internalIp")
    def internal_ip(self) -> pulumi.Output[str]:
        """
        The internal ip, must a private ip.
        """
        return pulumi.get(self, "internal_ip")

    @property
    @pulumi.getter(name="internalPort")
    def internal_port(self) -> pulumi.Output[str]:
        """
        The internal port, valid value is 1~65535|any.
        """
        return pulumi.get(self, "internal_port")

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> pulumi.Output[str]:
        """
        The ip protocal, valid value is tcp|udp|any.
        """
        return pulumi.get(self, "ip_protocol")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Field `name` has been deprecated from provider version 1.119.1. New field `forward_entry_name` instead.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="portBreak")
    def port_break(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to remove limits on the port range. Default value is `false`.
        """
        return pulumi.get(self, "port_break")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        (Available in 1.119.1+) The status of forward entry.
        """
        return pulumi.get(self, "status")

