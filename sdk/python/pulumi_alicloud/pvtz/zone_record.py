# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ZoneRecordArgs', 'ZoneRecord']

@pulumi.input_type
class ZoneRecordArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 value: pulumi.Input[str],
                 zone_id: pulumi.Input[str],
                 lang: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 resource_record: Optional[pulumi.Input[str]] = None,
                 rr: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 user_client_ip: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ZoneRecord resource.
        :param pulumi.Input[str] type: The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
        :param pulumi.Input[str] value: The value of the Private Zone Record.
        :param pulumi.Input[str] zone_id: The name of the Private Zone Record.
        :param pulumi.Input[str] lang: User language.
        :param pulumi.Input[int] priority: The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-99]. Default to 1.
        :param pulumi.Input[str] remark: The remark of the Private Zone Record.
        :param pulumi.Input[str] resource_record: The resource record of the Private Zone Record.
        :param pulumi.Input[str] rr: The rr of the Private Zone Record.
        :param pulumi.Input[str] status: Resolve record status. Value:
               - ENABLE: enable resolution.
               - DISABLE: pause parsing.
        :param pulumi.Input[int] ttl: The ttl of the Private Zone Record. Default to `60`.
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)
        pulumi.set(__self__, "zone_id", zone_id)
        if lang is not None:
            pulumi.set(__self__, "lang", lang)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if resource_record is not None:
            warnings.warn("""Field 'resource_record' has been deprecated from version 1.109.0. Use 'rr' instead.""", DeprecationWarning)
            pulumi.log.warn("""resource_record is deprecated: Field 'resource_record' has been deprecated from version 1.109.0. Use 'rr' instead.""")
        if resource_record is not None:
            pulumi.set(__self__, "resource_record", resource_record)
        if rr is not None:
            pulumi.set(__self__, "rr", rr)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if user_client_ip is not None:
            pulumi.set(__self__, "user_client_ip", user_client_ip)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value of the Private Zone Record.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        The name of the Private Zone Record.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        User language.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-99]. Default to 1.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        The remark of the Private Zone Record.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="resourceRecord")
    def resource_record(self) -> Optional[pulumi.Input[str]]:
        """
        The resource record of the Private Zone Record.
        """
        return pulumi.get(self, "resource_record")

    @resource_record.setter
    def resource_record(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_record", value)

    @property
    @pulumi.getter
    def rr(self) -> Optional[pulumi.Input[str]]:
        """
        The rr of the Private Zone Record.
        """
        return pulumi.get(self, "rr")

    @rr.setter
    def rr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rr", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Resolve record status. Value:
        - ENABLE: enable resolution.
        - DISABLE: pause parsing.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The ttl of the Private Zone Record. Default to `60`.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter(name="userClientIp")
    def user_client_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_client_ip")

    @user_client_ip.setter
    def user_client_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_client_ip", value)


@pulumi.input_type
class _ZoneRecordState:
    def __init__(__self__, *,
                 lang: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 record_id: Optional[pulumi.Input[int]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 resource_record: Optional[pulumi.Input[str]] = None,
                 rr: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_client_ip: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ZoneRecord resources.
        :param pulumi.Input[str] lang: User language.
        :param pulumi.Input[int] priority: The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-99]. Default to 1.
        :param pulumi.Input[int] record_id: The Private Zone Record ID.
        :param pulumi.Input[str] remark: The remark of the Private Zone Record.
        :param pulumi.Input[str] resource_record: The resource record of the Private Zone Record.
        :param pulumi.Input[str] rr: The rr of the Private Zone Record.
        :param pulumi.Input[str] status: Resolve record status. Value:
               - ENABLE: enable resolution.
               - DISABLE: pause parsing.
        :param pulumi.Input[int] ttl: The ttl of the Private Zone Record. Default to `60`.
        :param pulumi.Input[str] type: The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
        :param pulumi.Input[str] value: The value of the Private Zone Record.
        :param pulumi.Input[str] zone_id: The name of the Private Zone Record.
        """
        if lang is not None:
            pulumi.set(__self__, "lang", lang)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if record_id is not None:
            pulumi.set(__self__, "record_id", record_id)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if resource_record is not None:
            warnings.warn("""Field 'resource_record' has been deprecated from version 1.109.0. Use 'rr' instead.""", DeprecationWarning)
            pulumi.log.warn("""resource_record is deprecated: Field 'resource_record' has been deprecated from version 1.109.0. Use 'rr' instead.""")
        if resource_record is not None:
            pulumi.set(__self__, "resource_record", resource_record)
        if rr is not None:
            pulumi.set(__self__, "rr", rr)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_client_ip is not None:
            pulumi.set(__self__, "user_client_ip", user_client_ip)
        if value is not None:
            pulumi.set(__self__, "value", value)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        User language.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-99]. Default to 1.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="recordId")
    def record_id(self) -> Optional[pulumi.Input[int]]:
        """
        The Private Zone Record ID.
        """
        return pulumi.get(self, "record_id")

    @record_id.setter
    def record_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "record_id", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        The remark of the Private Zone Record.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="resourceRecord")
    def resource_record(self) -> Optional[pulumi.Input[str]]:
        """
        The resource record of the Private Zone Record.
        """
        return pulumi.get(self, "resource_record")

    @resource_record.setter
    def resource_record(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_record", value)

    @property
    @pulumi.getter
    def rr(self) -> Optional[pulumi.Input[str]]:
        """
        The rr of the Private Zone Record.
        """
        return pulumi.get(self, "rr")

    @rr.setter
    def rr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rr", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Resolve record status. Value:
        - ENABLE: enable resolution.
        - DISABLE: pause parsing.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The ttl of the Private Zone Record. Default to `60`.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userClientIp")
    def user_client_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_client_ip")

    @user_client_ip.setter
    def user_client_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_client_ip", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the Private Zone Record.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Private Zone Record.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class ZoneRecord(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 resource_record: Optional[pulumi.Input[str]] = None,
                 rr: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_client_ip: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        Private Zone Record can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:pvtz/zoneRecord:ZoneRecord example abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] lang: User language.
        :param pulumi.Input[int] priority: The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-99]. Default to 1.
        :param pulumi.Input[str] remark: The remark of the Private Zone Record.
        :param pulumi.Input[str] resource_record: The resource record of the Private Zone Record.
        :param pulumi.Input[str] rr: The rr of the Private Zone Record.
        :param pulumi.Input[str] status: Resolve record status. Value:
               - ENABLE: enable resolution.
               - DISABLE: pause parsing.
        :param pulumi.Input[int] ttl: The ttl of the Private Zone Record. Default to `60`.
        :param pulumi.Input[str] type: The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
        :param pulumi.Input[str] value: The value of the Private Zone Record.
        :param pulumi.Input[str] zone_id: The name of the Private Zone Record.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ZoneRecordArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Private Zone Record can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:pvtz/zoneRecord:ZoneRecord example abc123456
        ```

        :param str resource_name: The name of the resource.
        :param ZoneRecordArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ZoneRecordArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 resource_record: Optional[pulumi.Input[str]] = None,
                 rr: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_client_ip: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ZoneRecordArgs.__new__(ZoneRecordArgs)

            __props__.__dict__["lang"] = lang
            __props__.__dict__["priority"] = priority
            __props__.__dict__["remark"] = remark
            if resource_record is not None and not opts.urn:
                warnings.warn("""Field 'resource_record' has been deprecated from version 1.109.0. Use 'rr' instead.""", DeprecationWarning)
                pulumi.log.warn("""resource_record is deprecated: Field 'resource_record' has been deprecated from version 1.109.0. Use 'rr' instead.""")
            __props__.__dict__["resource_record"] = resource_record
            __props__.__dict__["rr"] = rr
            __props__.__dict__["status"] = status
            __props__.__dict__["ttl"] = ttl
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["user_client_ip"] = user_client_ip
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
            __props__.__dict__["record_id"] = None
        super(ZoneRecord, __self__).__init__(
            'alicloud:pvtz/zoneRecord:ZoneRecord',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            lang: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            record_id: Optional[pulumi.Input[int]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            resource_record: Optional[pulumi.Input[str]] = None,
            rr: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            user_client_ip: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'ZoneRecord':
        """
        Get an existing ZoneRecord resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] lang: User language.
        :param pulumi.Input[int] priority: The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-99]. Default to 1.
        :param pulumi.Input[int] record_id: The Private Zone Record ID.
        :param pulumi.Input[str] remark: The remark of the Private Zone Record.
        :param pulumi.Input[str] resource_record: The resource record of the Private Zone Record.
        :param pulumi.Input[str] rr: The rr of the Private Zone Record.
        :param pulumi.Input[str] status: Resolve record status. Value:
               - ENABLE: enable resolution.
               - DISABLE: pause parsing.
        :param pulumi.Input[int] ttl: The ttl of the Private Zone Record. Default to `60`.
        :param pulumi.Input[str] type: The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
        :param pulumi.Input[str] value: The value of the Private Zone Record.
        :param pulumi.Input[str] zone_id: The name of the Private Zone Record.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ZoneRecordState.__new__(_ZoneRecordState)

        __props__.__dict__["lang"] = lang
        __props__.__dict__["priority"] = priority
        __props__.__dict__["record_id"] = record_id
        __props__.__dict__["remark"] = remark
        __props__.__dict__["resource_record"] = resource_record
        __props__.__dict__["rr"] = rr
        __props__.__dict__["status"] = status
        __props__.__dict__["ttl"] = ttl
        __props__.__dict__["type"] = type
        __props__.__dict__["user_client_ip"] = user_client_ip
        __props__.__dict__["value"] = value
        __props__.__dict__["zone_id"] = zone_id
        return ZoneRecord(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def lang(self) -> pulumi.Output[Optional[str]]:
        """
        User language.
        """
        return pulumi.get(self, "lang")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[Optional[int]]:
        """
        The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-99]. Default to 1.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="recordId")
    def record_id(self) -> pulumi.Output[int]:
        """
        The Private Zone Record ID.
        """
        return pulumi.get(self, "record_id")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        The remark of the Private Zone Record.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="resourceRecord")
    def resource_record(self) -> pulumi.Output[str]:
        """
        The resource record of the Private Zone Record.
        """
        return pulumi.get(self, "resource_record")

    @property
    @pulumi.getter
    def rr(self) -> pulumi.Output[str]:
        """
        The rr of the Private Zone Record.
        """
        return pulumi.get(self, "rr")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        Resolve record status. Value:
        - ENABLE: enable resolution.
        - DISABLE: pause parsing.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[int]]:
        """
        The ttl of the Private Zone Record. Default to `60`.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userClientIp")
    def user_client_ip(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "user_client_ip")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        The value of the Private Zone Record.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The name of the Private Zone Record.
        """
        return pulumi.get(self, "zone_id")

