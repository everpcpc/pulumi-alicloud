# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EipArgs', 'Eip']

@pulumi.input_type
class EipArgs:
    def __init__(__self__, *,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_charge_type: Optional[pulumi.Input[str]] = None,
                 internet_charge_type: Optional[pulumi.Input[str]] = None,
                 isp: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Eip resource.
        :param pulumi.Input[int] bandwidth: Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
        :param pulumi.Input[bool] deletion_protection: Whether enable the deletion protection or not. Default value: `false`.
               - true: Enable deletion protection.
               - false: Disable deletion protection.
        :param pulumi.Input[str] description: Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        :param pulumi.Input[str] instance_charge_type: Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
        :param pulumi.Input[str] internet_charge_type: Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. From version `1.7.1`, default to `PayByTraffic`. It is only PayByBandwidth when `instance_charge_type` is PrePaid.
        :param pulumi.Input[str] isp: The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
        :param pulumi.Input[str] name: The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the eip belongs.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if deletion_protection is not None:
            pulumi.set(__self__, "deletion_protection", deletion_protection)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_charge_type is not None:
            pulumi.set(__self__, "instance_charge_type", instance_charge_type)
        if internet_charge_type is not None:
            pulumi.set(__self__, "internet_charge_type", internet_charge_type)
        if isp is not None:
            pulumi.set(__self__, "isp", isp)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether enable the deletion protection or not. Default value: `false`.
        - true: Enable deletion protection.
        - false: Disable deletion protection.
        """
        return pulumi.get(self, "deletion_protection")

    @deletion_protection.setter
    def deletion_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deletion_protection", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceChargeType")
    def instance_charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
        """
        return pulumi.get(self, "instance_charge_type")

    @instance_charge_type.setter
    def instance_charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_charge_type", value)

    @property
    @pulumi.getter(name="internetChargeType")
    def internet_charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. From version `1.7.1`, default to `PayByTraffic`. It is only PayByBandwidth when `instance_charge_type` is PrePaid.
        """
        return pulumi.get(self, "internet_charge_type")

    @internet_charge_type.setter
    def internet_charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internet_charge_type", value)

    @property
    @pulumi.getter
    def isp(self) -> Optional[pulumi.Input[str]]:
        """
        The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
        """
        return pulumi.get(self, "isp")

    @isp.setter
    def isp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "isp", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of resource group which the eip belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _EipState:
    def __init__(__self__, *,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_charge_type: Optional[pulumi.Input[str]] = None,
                 internet_charge_type: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 isp: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Eip resources.
        :param pulumi.Input[int] bandwidth: Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
        :param pulumi.Input[bool] deletion_protection: Whether enable the deletion protection or not. Default value: `false`.
               - true: Enable deletion protection.
               - false: Disable deletion protection.
        :param pulumi.Input[str] description: Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        :param pulumi.Input[str] instance_charge_type: Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
        :param pulumi.Input[str] internet_charge_type: Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. From version `1.7.1`, default to `PayByTraffic`. It is only PayByBandwidth when `instance_charge_type` is PrePaid.
        :param pulumi.Input[str] ip_address: The elastic ip address
        :param pulumi.Input[str] isp: The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
        :param pulumi.Input[str] name: The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the eip belongs.
        :param pulumi.Input[str] status: The EIP current status.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if deletion_protection is not None:
            pulumi.set(__self__, "deletion_protection", deletion_protection)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_charge_type is not None:
            pulumi.set(__self__, "instance_charge_type", instance_charge_type)
        if internet_charge_type is not None:
            pulumi.set(__self__, "internet_charge_type", internet_charge_type)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if isp is not None:
            pulumi.set(__self__, "isp", isp)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether enable the deletion protection or not. Default value: `false`.
        - true: Enable deletion protection.
        - false: Disable deletion protection.
        """
        return pulumi.get(self, "deletion_protection")

    @deletion_protection.setter
    def deletion_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deletion_protection", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceChargeType")
    def instance_charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
        """
        return pulumi.get(self, "instance_charge_type")

    @instance_charge_type.setter
    def instance_charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_charge_type", value)

    @property
    @pulumi.getter(name="internetChargeType")
    def internet_charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. From version `1.7.1`, default to `PayByTraffic`. It is only PayByBandwidth when `instance_charge_type` is PrePaid.
        """
        return pulumi.get(self, "internet_charge_type")

    @internet_charge_type.setter
    def internet_charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internet_charge_type", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The elastic ip address
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter
    def isp(self) -> Optional[pulumi.Input[str]]:
        """
        The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
        """
        return pulumi.get(self, "isp")

    @isp.setter
    def isp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "isp", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of resource group which the eip belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The EIP current status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class Eip(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_charge_type: Optional[pulumi.Input[str]] = None,
                 internet_charge_type: Optional[pulumi.Input[str]] = None,
                 isp: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        ## Import

        Elastic IP address can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ecs/eip:Eip example eip-abc12345678
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
        :param pulumi.Input[bool] deletion_protection: Whether enable the deletion protection or not. Default value: `false`.
               - true: Enable deletion protection.
               - false: Disable deletion protection.
        :param pulumi.Input[str] description: Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        :param pulumi.Input[str] instance_charge_type: Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
        :param pulumi.Input[str] internet_charge_type: Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. From version `1.7.1`, default to `PayByTraffic`. It is only PayByBandwidth when `instance_charge_type` is PrePaid.
        :param pulumi.Input[str] isp: The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
        :param pulumi.Input[str] name: The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the eip belongs.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[EipArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Elastic IP address can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ecs/eip:Eip example eip-abc12345678
        ```

        :param str resource_name: The name of the resource.
        :param EipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_charge_type: Optional[pulumi.Input[str]] = None,
                 internet_charge_type: Optional[pulumi.Input[str]] = None,
                 isp: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
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
            __props__ = EipArgs.__new__(EipArgs)

            __props__.__dict__["bandwidth"] = bandwidth
            __props__.__dict__["deletion_protection"] = deletion_protection
            __props__.__dict__["description"] = description
            __props__.__dict__["instance_charge_type"] = instance_charge_type
            __props__.__dict__["internet_charge_type"] = internet_charge_type
            __props__.__dict__["isp"] = isp
            __props__.__dict__["name"] = name
            __props__.__dict__["period"] = period
            __props__.__dict__["resource_group_id"] = resource_group_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["ip_address"] = None
            __props__.__dict__["status"] = None
        super(Eip, __self__).__init__(
            'alicloud:ecs/eip:Eip',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bandwidth: Optional[pulumi.Input[int]] = None,
            deletion_protection: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            instance_charge_type: Optional[pulumi.Input[str]] = None,
            internet_charge_type: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            isp: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[int]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Eip':
        """
        Get an existing Eip resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
        :param pulumi.Input[bool] deletion_protection: Whether enable the deletion protection or not. Default value: `false`.
               - true: Enable deletion protection.
               - false: Disable deletion protection.
        :param pulumi.Input[str] description: Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        :param pulumi.Input[str] instance_charge_type: Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
        :param pulumi.Input[str] internet_charge_type: Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. From version `1.7.1`, default to `PayByTraffic`. It is only PayByBandwidth when `instance_charge_type` is PrePaid.
        :param pulumi.Input[str] ip_address: The elastic ip address
        :param pulumi.Input[str] isp: The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
        :param pulumi.Input[str] name: The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the eip belongs.
        :param pulumi.Input[str] status: The EIP current status.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EipState.__new__(_EipState)

        __props__.__dict__["bandwidth"] = bandwidth
        __props__.__dict__["deletion_protection"] = deletion_protection
        __props__.__dict__["description"] = description
        __props__.__dict__["instance_charge_type"] = instance_charge_type
        __props__.__dict__["internet_charge_type"] = internet_charge_type
        __props__.__dict__["ip_address"] = ip_address
        __props__.__dict__["isp"] = isp
        __props__.__dict__["name"] = name
        __props__.__dict__["period"] = period
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        return Eip(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Output[Optional[int]]:
        """
        Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> pulumi.Output[bool]:
        """
        Whether enable the deletion protection or not. Default value: `false`.
        - true: Enable deletion protection.
        - false: Disable deletion protection.
        """
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="instanceChargeType")
    def instance_charge_type(self) -> pulumi.Output[Optional[str]]:
        """
        Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
        """
        return pulumi.get(self, "instance_charge_type")

    @property
    @pulumi.getter(name="internetChargeType")
    def internet_charge_type(self) -> pulumi.Output[Optional[str]]:
        """
        Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. From version `1.7.1`, default to `PayByTraffic`. It is only PayByBandwidth when `instance_charge_type` is PrePaid.
        """
        return pulumi.get(self, "internet_charge_type")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        The elastic ip address
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def isp(self) -> pulumi.Output[str]:
        """
        The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
        """
        return pulumi.get(self, "isp")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        The Id of resource group which the eip belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The EIP current status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

