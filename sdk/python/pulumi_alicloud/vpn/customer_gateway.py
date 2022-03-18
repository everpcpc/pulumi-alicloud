# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CustomerGatewayArgs', 'CustomerGateway']

@pulumi.input_type
class CustomerGatewayArgs:
    def __init__(__self__, *,
                 ip_address: pulumi.Input[str],
                 asn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CustomerGateway resource.
        :param pulumi.Input[str] ip_address: The IP address of the customer gateway.
        :param pulumi.Input[str] asn: The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
        :param pulumi.Input[str] description: The description of the VPN customer gateway instance.
        :param pulumi.Input[str] name: The name of the VPN customer gateway. Defaults to null.
        """
        pulumi.set(__self__, "ip_address", ip_address)
        if asn is not None:
            pulumi.set(__self__, "asn", asn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Input[str]:
        """
        The IP address of the customer gateway.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter
    def asn(self) -> Optional[pulumi.Input[str]]:
        """
        The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
        """
        return pulumi.get(self, "asn")

    @asn.setter
    def asn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "asn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the VPN customer gateway instance.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the VPN customer gateway. Defaults to null.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _CustomerGatewayState:
    def __init__(__self__, *,
                 asn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CustomerGateway resources.
        :param pulumi.Input[str] asn: The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
        :param pulumi.Input[str] description: The description of the VPN customer gateway instance.
        :param pulumi.Input[str] ip_address: The IP address of the customer gateway.
        :param pulumi.Input[str] name: The name of the VPN customer gateway. Defaults to null.
        """
        if asn is not None:
            pulumi.set(__self__, "asn", asn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def asn(self) -> Optional[pulumi.Input[str]]:
        """
        The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
        """
        return pulumi.get(self, "asn")

    @asn.setter
    def asn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "asn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the VPN customer gateway instance.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address of the customer gateway.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the VPN customer gateway. Defaults to null.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class CustomerGateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        foo = alicloud.vpn.CustomerGateway("foo",
            description="vpnCgwDescriptionExample",
            ip_address="43.104.22.228")
        ```

        ## Import

        VPN customer gateway can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpn/customerGateway:CustomerGateway example cgw-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] asn: The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
        :param pulumi.Input[str] description: The description of the VPN customer gateway instance.
        :param pulumi.Input[str] ip_address: The IP address of the customer gateway.
        :param pulumi.Input[str] name: The name of the VPN customer gateway. Defaults to null.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomerGatewayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        foo = alicloud.vpn.CustomerGateway("foo",
            description="vpnCgwDescriptionExample",
            ip_address="43.104.22.228")
        ```

        ## Import

        VPN customer gateway can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpn/customerGateway:CustomerGateway example cgw-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param CustomerGatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomerGatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomerGatewayArgs.__new__(CustomerGatewayArgs)

            __props__.__dict__["asn"] = asn
            __props__.__dict__["description"] = description
            if ip_address is None and not opts.urn:
                raise TypeError("Missing required property 'ip_address'")
            __props__.__dict__["ip_address"] = ip_address
            __props__.__dict__["name"] = name
        super(CustomerGateway, __self__).__init__(
            'alicloud:vpn/customerGateway:CustomerGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            asn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'CustomerGateway':
        """
        Get an existing CustomerGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] asn: The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
        :param pulumi.Input[str] description: The description of the VPN customer gateway instance.
        :param pulumi.Input[str] ip_address: The IP address of the customer gateway.
        :param pulumi.Input[str] name: The name of the VPN customer gateway. Defaults to null.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CustomerGatewayState.__new__(_CustomerGatewayState)

        __props__.__dict__["asn"] = asn
        __props__.__dict__["description"] = description
        __props__.__dict__["ip_address"] = ip_address
        __props__.__dict__["name"] = name
        return CustomerGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def asn(self) -> pulumi.Output[Optional[str]]:
        """
        The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
        """
        return pulumi.get(self, "asn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the VPN customer gateway instance.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        The IP address of the customer gateway.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the VPN customer gateway. Defaults to null.
        """
        return pulumi.get(self, "name")

