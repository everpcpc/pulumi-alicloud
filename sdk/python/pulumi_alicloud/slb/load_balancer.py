# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['LoadBalancer']


class LoadBalancer(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 address_ip_version: Optional[pulumi.Input[str]] = None,
                 address_type: Optional[pulumi.Input[str]] = None,
                 bandwidth: Optional[pulumi.Input[float]] = None,
                 delete_protection: Optional[pulumi.Input[str]] = None,
                 instance_charge_type: Optional[pulumi.Input[str]] = None,
                 internet: Optional[pulumi.Input[bool]] = None,
                 internet_charge_type: Optional[pulumi.Input[str]] = None,
                 master_zone_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[float]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 slave_zone_id: Optional[pulumi.Input[str]] = None,
                 specification: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an Application Load Balancer resource.

        > **NOTE:** At present, to avoid some unnecessary regulation confusion, SLB can not support alicloud international account to create "paybybandwidth" instance.

        > **NOTE:** The supported specifications vary by region. Currently not all regions support guaranteed-performance instances.
        For more details about guaranteed-performance instance, see [Guaranteed-performance instances](https://www.alibabacloud.com/help/doc-detail/27657.htm).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraformtestslbconfig"
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("defaultNetwork", cidr_block="172.16.0.0/12")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/21",
            availability_zone=default_zones.zones[0].id)
        default_load_balancer = alicloud.slb.LoadBalancer("defaultLoadBalancer",
            specification="slb.s2.small",
            vswitch_id=default_switch.id,
            tags={
                "tag_a": 1,
                "tag_b": 2,
                "tag_c": 3,
                "tag_d": 4,
                "tag_e": 5,
                "tag_f": 6,
                "tag_g": 7,
                "tag_h": 8,
                "tag_i": 9,
                "tag_j": 10,
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
        :param pulumi.Input[str] address_ip_version: The IP version of the SLB instance to be created, which can be set to ipv4 or ipv6 . Default to "ipv4". Now, only internet instance support ipv6 address.
        :param pulumi.Input[str] address_type: The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be "intranet".
               - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
               - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
        :param pulumi.Input[float] bandwidth: Valid
               value is between 1 and 1000, If argument "internet_charge_type" is "paybytraffic", then this value will be ignore.
        :param pulumi.Input[str] delete_protection: Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
        :param pulumi.Input[str] instance_charge_type: The billing method of the load balancer. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
        :param pulumi.Input[bool] internet: Field 'internet' has been deprecated from provider version 1.55.3. Use 'address_type' replaces it.
        :param pulumi.Input[str] internet_charge_type: Valid
               values are `PayByBandwidth`, `PayByTraffic`. If this value is "PayByBandwidth", then argument "internet" must be "true". Default is "PayByTraffic". If load balancer launched in VPC, this value must be "PayByTraffic".
               Before version 1.10.1, the valid values are "paybybandwidth" and "paybytraffic".
        :param pulumi.Input[str] master_zone_id: The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
        :param pulumi.Input[float] period: The duration that you will buy the resource, in month. It is valid when `instance_charge_type` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36].
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the SLB belongs.
        :param pulumi.Input[str] slave_zone_id: The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
        :param pulumi.Input[str] specification: The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.
               Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: "slb.s1.small", "slb.s2.small", "slb.s2.medium",
               "slb.s3.small", "slb.s3.medium", "slb.s3.large" and "slb.s4.large".
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
        :param pulumi.Input[str] vswitch_id: The VSwitch ID to launch in. If `address_type` is internet, it will be ignore.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['address'] = address
            __props__['address_ip_version'] = address_ip_version
            __props__['address_type'] = address_type
            __props__['bandwidth'] = bandwidth
            __props__['delete_protection'] = delete_protection
            __props__['instance_charge_type'] = instance_charge_type
            if internet is not None:
                warnings.warn("Field 'internet' has been deprecated from provider version 1.55.3. Use 'address_type' replaces it.", DeprecationWarning)
                pulumi.log.warn("internet is deprecated: Field 'internet' has been deprecated from provider version 1.55.3. Use 'address_type' replaces it.")
            __props__['internet'] = internet
            __props__['internet_charge_type'] = internet_charge_type
            __props__['master_zone_id'] = master_zone_id
            __props__['name'] = name
            __props__['period'] = period
            __props__['resource_group_id'] = resource_group_id
            __props__['slave_zone_id'] = slave_zone_id
            __props__['specification'] = specification
            __props__['tags'] = tags
            __props__['vswitch_id'] = vswitch_id
        super(LoadBalancer, __self__).__init__(
            'alicloud:slb/loadBalancer:LoadBalancer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            address_ip_version: Optional[pulumi.Input[str]] = None,
            address_type: Optional[pulumi.Input[str]] = None,
            bandwidth: Optional[pulumi.Input[float]] = None,
            delete_protection: Optional[pulumi.Input[str]] = None,
            instance_charge_type: Optional[pulumi.Input[str]] = None,
            internet: Optional[pulumi.Input[bool]] = None,
            internet_charge_type: Optional[pulumi.Input[str]] = None,
            master_zone_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[float]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            slave_zone_id: Optional[pulumi.Input[str]] = None,
            specification: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None) -> 'LoadBalancer':
        """
        Get an existing LoadBalancer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
        :param pulumi.Input[str] address_ip_version: The IP version of the SLB instance to be created, which can be set to ipv4 or ipv6 . Default to "ipv4". Now, only internet instance support ipv6 address.
        :param pulumi.Input[str] address_type: The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be "intranet".
               - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
               - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
        :param pulumi.Input[float] bandwidth: Valid
               value is between 1 and 1000, If argument "internet_charge_type" is "paybytraffic", then this value will be ignore.
        :param pulumi.Input[str] delete_protection: Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
        :param pulumi.Input[str] instance_charge_type: The billing method of the load balancer. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
        :param pulumi.Input[bool] internet: Field 'internet' has been deprecated from provider version 1.55.3. Use 'address_type' replaces it.
        :param pulumi.Input[str] internet_charge_type: Valid
               values are `PayByBandwidth`, `PayByTraffic`. If this value is "PayByBandwidth", then argument "internet" must be "true". Default is "PayByTraffic". If load balancer launched in VPC, this value must be "PayByTraffic".
               Before version 1.10.1, the valid values are "paybybandwidth" and "paybytraffic".
        :param pulumi.Input[str] master_zone_id: The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
        :param pulumi.Input[float] period: The duration that you will buy the resource, in month. It is valid when `instance_charge_type` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36].
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the SLB belongs.
        :param pulumi.Input[str] slave_zone_id: The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
        :param pulumi.Input[str] specification: The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.
               Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: "slb.s1.small", "slb.s2.small", "slb.s2.medium",
               "slb.s3.small", "slb.s3.medium", "slb.s3.large" and "slb.s4.large".
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
        :param pulumi.Input[str] vswitch_id: The VSwitch ID to launch in. If `address_type` is internet, it will be ignore.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address"] = address
        __props__["address_ip_version"] = address_ip_version
        __props__["address_type"] = address_type
        __props__["bandwidth"] = bandwidth
        __props__["delete_protection"] = delete_protection
        __props__["instance_charge_type"] = instance_charge_type
        __props__["internet"] = internet
        __props__["internet_charge_type"] = internet_charge_type
        __props__["master_zone_id"] = master_zone_id
        __props__["name"] = name
        __props__["period"] = period
        __props__["resource_group_id"] = resource_group_id
        __props__["slave_zone_id"] = slave_zone_id
        __props__["specification"] = specification
        __props__["tags"] = tags
        __props__["vswitch_id"] = vswitch_id
        return LoadBalancer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="addressIpVersion")
    def address_ip_version(self) -> pulumi.Output[Optional[str]]:
        """
        The IP version of the SLB instance to be created, which can be set to ipv4 or ipv6 . Default to "ipv4". Now, only internet instance support ipv6 address.
        """
        return pulumi.get(self, "address_ip_version")

    @property
    @pulumi.getter(name="addressType")
    def address_type(self) -> pulumi.Output[str]:
        """
        The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be "intranet".
        - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
        - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
        """
        return pulumi.get(self, "address_type")

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Output[Optional[float]]:
        """
        Valid
        value is between 1 and 1000, If argument "internet_charge_type" is "paybytraffic", then this value will be ignore.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="deleteProtection")
    def delete_protection(self) -> pulumi.Output[Optional[str]]:
        """
        Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
        """
        return pulumi.get(self, "delete_protection")

    @property
    @pulumi.getter(name="instanceChargeType")
    def instance_charge_type(self) -> pulumi.Output[Optional[str]]:
        """
        The billing method of the load balancer. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
        """
        return pulumi.get(self, "instance_charge_type")

    @property
    @pulumi.getter
    def internet(self) -> pulumi.Output[bool]:
        """
        Field 'internet' has been deprecated from provider version 1.55.3. Use 'address_type' replaces it.
        """
        return pulumi.get(self, "internet")

    @property
    @pulumi.getter(name="internetChargeType")
    def internet_charge_type(self) -> pulumi.Output[Optional[str]]:
        """
        Valid
        values are `PayByBandwidth`, `PayByTraffic`. If this value is "PayByBandwidth", then argument "internet" must be "true". Default is "PayByTraffic". If load balancer launched in VPC, this value must be "PayByTraffic".
        Before version 1.10.1, the valid values are "paybybandwidth" and "paybytraffic".
        """
        return pulumi.get(self, "internet_charge_type")

    @property
    @pulumi.getter(name="masterZoneId")
    def master_zone_id(self) -> pulumi.Output[str]:
        """
        The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
        """
        return pulumi.get(self, "master_zone_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[float]]:
        """
        The duration that you will buy the resource, in month. It is valid when `instance_charge_type` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36].
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        The Id of resource group which the SLB belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="slaveZoneId")
    def slave_zone_id(self) -> pulumi.Output[str]:
        """
        The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
        """
        return pulumi.get(self, "slave_zone_id")

    @property
    @pulumi.getter
    def specification(self) -> pulumi.Output[Optional[str]]:
        """
        The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.
        Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: "slb.s1.small", "slb.s2.small", "slb.s2.medium",
        "slb.s3.small", "slb.s3.medium", "slb.s3.large" and "slb.s4.large".
        """
        return pulumi.get(self, "specification")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[Optional[str]]:
        """
        The VSwitch ID to launch in. If `address_type` is internet, it will be ignore.
        """
        return pulumi.get(self, "vswitch_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

