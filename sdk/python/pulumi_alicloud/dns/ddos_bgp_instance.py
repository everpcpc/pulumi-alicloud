# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DdosBgpInstanceArgs', 'DdosBgpInstance']

@pulumi.input_type
class DdosBgpInstanceArgs:
    def __init__(__self__, *,
                 bandwidth: pulumi.Input[int],
                 ip_count: pulumi.Input[int],
                 ip_type: pulumi.Input[str],
                 base_bandwidth: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DdosBgpInstance resource.
        :param pulumi.Input[int] bandwidth: Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
        :param pulumi.Input[int] ip_count: IP count of the instance. Valid values: 100.
        :param pulumi.Input[str] ip_type: IP version of the instance. Valid values: IPv4,IPv6.
        :param pulumi.Input[int] base_bandwidth: Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
        :param pulumi.Input[str] name: Name of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[int] period: The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
        :param pulumi.Input[str] type: Type of the instance. Valid values: Enterprise,Professional. Default to `Enterprise`
        """
        pulumi.set(__self__, "bandwidth", bandwidth)
        pulumi.set(__self__, "ip_count", ip_count)
        pulumi.set(__self__, "ip_type", ip_type)
        if base_bandwidth is not None:
            pulumi.set(__self__, "base_bandwidth", base_bandwidth)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Input[int]:
        """
        Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: pulumi.Input[int]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="ipCount")
    def ip_count(self) -> pulumi.Input[int]:
        """
        IP count of the instance. Valid values: 100.
        """
        return pulumi.get(self, "ip_count")

    @ip_count.setter
    def ip_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "ip_count", value)

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> pulumi.Input[str]:
        """
        IP version of the instance. Valid values: IPv4,IPv6.
        """
        return pulumi.get(self, "ip_type")

    @ip_type.setter
    def ip_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_type", value)

    @property
    @pulumi.getter(name="baseBandwidth")
    def base_bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
        """
        return pulumi.get(self, "base_bandwidth")

    @base_bandwidth.setter
    def base_bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "base_bandwidth", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the instance. This name can have a string of 1 to 63 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the instance. Valid values: Enterprise,Professional. Default to `Enterprise`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _DdosBgpInstanceState:
    def __init__(__self__, *,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 base_bandwidth: Optional[pulumi.Input[int]] = None,
                 ip_count: Optional[pulumi.Input[int]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DdosBgpInstance resources.
        :param pulumi.Input[int] bandwidth: Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
        :param pulumi.Input[int] base_bandwidth: Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
        :param pulumi.Input[int] ip_count: IP count of the instance. Valid values: 100.
        :param pulumi.Input[str] ip_type: IP version of the instance. Valid values: IPv4,IPv6.
        :param pulumi.Input[str] name: Name of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[int] period: The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
        :param pulumi.Input[str] type: Type of the instance. Valid values: Enterprise,Professional. Default to `Enterprise`
        """
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if base_bandwidth is not None:
            pulumi.set(__self__, "base_bandwidth", base_bandwidth)
        if ip_count is not None:
            pulumi.set(__self__, "ip_count", ip_count)
        if ip_type is not None:
            pulumi.set(__self__, "ip_type", ip_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="baseBandwidth")
    def base_bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
        """
        return pulumi.get(self, "base_bandwidth")

    @base_bandwidth.setter
    def base_bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "base_bandwidth", value)

    @property
    @pulumi.getter(name="ipCount")
    def ip_count(self) -> Optional[pulumi.Input[int]]:
        """
        IP count of the instance. Valid values: 100.
        """
        return pulumi.get(self, "ip_count")

    @ip_count.setter
    def ip_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ip_count", value)

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> Optional[pulumi.Input[str]]:
        """
        IP version of the instance. Valid values: IPv4,IPv6.
        """
        return pulumi.get(self, "ip_type")

    @ip_type.setter
    def ip_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the instance. This name can have a string of 1 to 63 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the instance. Valid values: Enterprise,Professional. Default to `Enterprise`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


warnings.warn("""alicloud.dns.DdosBgpInstance has been deprecated in favor of alicloud.ddos.DdosBgpInstance""", DeprecationWarning)


class DdosBgpInstance(pulumi.CustomResource):
    warnings.warn("""alicloud.dns.DdosBgpInstance has been deprecated in favor of alicloud.ddos.DdosBgpInstance""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 base_bandwidth: Optional[pulumi.Input[int]] = None,
                 ip_count: Optional[pulumi.Input[int]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Anti-DDoS Advanced instance resource. "Ddosbgp" is the short term of this product.

        > **NOTE:** The endpoint of bssopenapi used only support "business.aliyuncs.com" at present.

        > **NOTE:** Available in 1.57.0+ .

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        instance = alicloud.ddos.DdosBgpInstance("instance",
            bandwidth=201,
            base_bandwidth=20,
            ip_count=100,
            ip_type="IPv4")
        ```

        ## Import

        Ddosbgp instance can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:dns/ddosBgpInstance:DdosBgpInstance example ddosbgp-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
        :param pulumi.Input[int] base_bandwidth: Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
        :param pulumi.Input[int] ip_count: IP count of the instance. Valid values: 100.
        :param pulumi.Input[str] ip_type: IP version of the instance. Valid values: IPv4,IPv6.
        :param pulumi.Input[str] name: Name of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[int] period: The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
        :param pulumi.Input[str] type: Type of the instance. Valid values: Enterprise,Professional. Default to `Enterprise`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DdosBgpInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Anti-DDoS Advanced instance resource. "Ddosbgp" is the short term of this product.

        > **NOTE:** The endpoint of bssopenapi used only support "business.aliyuncs.com" at present.

        > **NOTE:** Available in 1.57.0+ .

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        instance = alicloud.ddos.DdosBgpInstance("instance",
            bandwidth=201,
            base_bandwidth=20,
            ip_count=100,
            ip_type="IPv4")
        ```

        ## Import

        Ddosbgp instance can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:dns/ddosBgpInstance:DdosBgpInstance example ddosbgp-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param DdosBgpInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DdosBgpInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 base_bandwidth: Optional[pulumi.Input[int]] = None,
                 ip_count: Optional[pulumi.Input[int]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""DdosBgpInstance is deprecated: alicloud.dns.DdosBgpInstance has been deprecated in favor of alicloud.ddos.DdosBgpInstance""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DdosBgpInstanceArgs.__new__(DdosBgpInstanceArgs)

            if bandwidth is None and not opts.urn:
                raise TypeError("Missing required property 'bandwidth'")
            __props__.__dict__["bandwidth"] = bandwidth
            __props__.__dict__["base_bandwidth"] = base_bandwidth
            if ip_count is None and not opts.urn:
                raise TypeError("Missing required property 'ip_count'")
            __props__.__dict__["ip_count"] = ip_count
            if ip_type is None and not opts.urn:
                raise TypeError("Missing required property 'ip_type'")
            __props__.__dict__["ip_type"] = ip_type
            __props__.__dict__["name"] = name
            __props__.__dict__["period"] = period
            __props__.__dict__["type"] = type
        super(DdosBgpInstance, __self__).__init__(
            'alicloud:dns/ddosBgpInstance:DdosBgpInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bandwidth: Optional[pulumi.Input[int]] = None,
            base_bandwidth: Optional[pulumi.Input[int]] = None,
            ip_count: Optional[pulumi.Input[int]] = None,
            ip_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'DdosBgpInstance':
        """
        Get an existing DdosBgpInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
        :param pulumi.Input[int] base_bandwidth: Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
        :param pulumi.Input[int] ip_count: IP count of the instance. Valid values: 100.
        :param pulumi.Input[str] ip_type: IP version of the instance. Valid values: IPv4,IPv6.
        :param pulumi.Input[str] name: Name of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[int] period: The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
        :param pulumi.Input[str] type: Type of the instance. Valid values: Enterprise,Professional. Default to `Enterprise`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DdosBgpInstanceState.__new__(_DdosBgpInstanceState)

        __props__.__dict__["bandwidth"] = bandwidth
        __props__.__dict__["base_bandwidth"] = base_bandwidth
        __props__.__dict__["ip_count"] = ip_count
        __props__.__dict__["ip_type"] = ip_type
        __props__.__dict__["name"] = name
        __props__.__dict__["period"] = period
        __props__.__dict__["type"] = type
        return DdosBgpInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Output[int]:
        """
        Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="baseBandwidth")
    def base_bandwidth(self) -> pulumi.Output[Optional[int]]:
        """
        Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
        """
        return pulumi.get(self, "base_bandwidth")

    @property
    @pulumi.getter(name="ipCount")
    def ip_count(self) -> pulumi.Output[int]:
        """
        IP count of the instance. Valid values: 100.
        """
        return pulumi.get(self, "ip_count")

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> pulumi.Output[str]:
        """
        IP version of the instance. Valid values: IPv4,IPv6.
        """
        return pulumi.get(self, "ip_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the instance. This name can have a string of 1 to 63 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[int]]:
        """
        The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        Type of the instance. Valid values: Enterprise,Professional. Default to `Enterprise`
        """
        return pulumi.get(self, "type")

