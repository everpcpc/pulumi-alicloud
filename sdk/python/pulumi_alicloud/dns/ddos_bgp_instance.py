# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class DdosBgpInstance(pulumi.CustomResource):
    bandwidth: pulumi.Output[float]
    """
    Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
    """
    base_bandwidth: pulumi.Output[float]
    """
    Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
    """
    ip_count: pulumi.Output[float]
    """
    IP count of the instance. Valid values: 100.
    """
    ip_type: pulumi.Output[str]
    """
    IP version of the instance. Valid values: IPv4,IPv6.
    """
    name: pulumi.Output[str]
    """
    Name of the instance. This name can have a string of 1 to 63 characters.
    """
    period: pulumi.Output[float]
    """
    The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
    """
    type: pulumi.Output[str]
    """
    Type of the instance. Valid values: Enterprise,Professional. Default to `Enterprise`  
    """
    def __init__(__self__, resource_name, opts=None, bandwidth=None, base_bandwidth=None, ip_count=None, ip_type=None, name=None, period=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Anti-DDoS Advanced instance resource. "Ddosbgp" is the short term of this product.

        > **NOTE:** The endpoint of bssopenapi used only support "business.aliyuncs.com" at present.

        > **NOTE:** Available in 1.57.0+ .



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] bandwidth: Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
        :param pulumi.Input[float] base_bandwidth: Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
        :param pulumi.Input[float] ip_count: IP count of the instance. Valid values: 100.
        :param pulumi.Input[str] ip_type: IP version of the instance. Valid values: IPv4,IPv6.
        :param pulumi.Input[str] name: Name of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[float] period: The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
        :param pulumi.Input[str] type: Type of the instance. Valid values: Enterprise,Professional. Default to `Enterprise`  
        """
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if bandwidth is None:
                raise TypeError("Missing required property 'bandwidth'")
            __props__['bandwidth'] = bandwidth
            __props__['base_bandwidth'] = base_bandwidth
            if ip_count is None:
                raise TypeError("Missing required property 'ip_count'")
            __props__['ip_count'] = ip_count
            if ip_type is None:
                raise TypeError("Missing required property 'ip_type'")
            __props__['ip_type'] = ip_type
            __props__['name'] = name
            __props__['period'] = period
            __props__['type'] = type
        super(DdosBgpInstance, __self__).__init__(
            'alicloud:dns/ddosBgpInstance:DdosBgpInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bandwidth=None, base_bandwidth=None, ip_count=None, ip_type=None, name=None, period=None, type=None):
        """
        Get an existing DdosBgpInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] bandwidth: Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
        :param pulumi.Input[float] base_bandwidth: Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
        :param pulumi.Input[float] ip_count: IP count of the instance. Valid values: 100.
        :param pulumi.Input[str] ip_type: IP version of the instance. Valid values: IPv4,IPv6.
        :param pulumi.Input[str] name: Name of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[float] period: The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
        :param pulumi.Input[str] type: Type of the instance. Valid values: Enterprise,Professional. Default to `Enterprise`  
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bandwidth"] = bandwidth
        __props__["base_bandwidth"] = base_bandwidth
        __props__["ip_count"] = ip_count
        __props__["ip_type"] = ip_type
        __props__["name"] = name
        __props__["period"] = period
        __props__["type"] = type
        return DdosBgpInstance(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

