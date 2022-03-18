# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DataLimitArgs', 'DataLimit']

@pulumi.input_type
class DataLimitArgs:
    def __init__(__self__, *,
                 resource_type: pulumi.Input[str],
                 audit_status: Optional[pulumi.Input[int]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 log_store_day: Optional[pulumi.Input[int]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 service_region_id: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DataLimit resource.
        :param pulumi.Input[str] resource_type: The type of the service to which the data asset belongs. Valid values: `MaxCompute`, `OSS`, `RDS`.
        :param pulumi.Input[int] audit_status: Whether to enable the log auditing feature. Valid values: `0`, `1`.
        :param pulumi.Input[str] engine_type: The type of the database. Valid values: `MySQL`, `SQLServer`.
        :param pulumi.Input[str] lang: The lang.
        :param pulumi.Input[int] log_store_day: The retention period of raw logs after you enable the log auditing feature. Unit: day. Valid values: `180`, `30`, `365`, `90`. **NOTE:** The`log_store_day` is valid when the `audit_status` is `1`.
        :param pulumi.Input[str] parent_id: The ID of the data asset.
        :param pulumi.Input[str] password: The password that is used to connect to the database.
        :param pulumi.Input[int] port: The port that is used to connect to the database.
        :param pulumi.Input[str] service_region_id: The region ID of the data asset.
        :param pulumi.Input[str] user_name: The name of the service to which the data asset belongs.
        """
        pulumi.set(__self__, "resource_type", resource_type)
        if audit_status is not None:
            pulumi.set(__self__, "audit_status", audit_status)
        if engine_type is not None:
            pulumi.set(__self__, "engine_type", engine_type)
        if lang is not None:
            pulumi.set(__self__, "lang", lang)
        if log_store_day is not None:
            pulumi.set(__self__, "log_store_day", log_store_day)
        if parent_id is not None:
            pulumi.set(__self__, "parent_id", parent_id)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if service_region_id is not None:
            pulumi.set(__self__, "service_region_id", service_region_id)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Input[str]:
        """
        The type of the service to which the data asset belongs. Valid values: `MaxCompute`, `OSS`, `RDS`.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter(name="auditStatus")
    def audit_status(self) -> Optional[pulumi.Input[int]]:
        """
        Whether to enable the log auditing feature. Valid values: `0`, `1`.
        """
        return pulumi.get(self, "audit_status")

    @audit_status.setter
    def audit_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "audit_status", value)

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the database. Valid values: `MySQL`, `SQLServer`.
        """
        return pulumi.get(self, "engine_type")

    @engine_type.setter
    def engine_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_type", value)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        The lang.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)

    @property
    @pulumi.getter(name="logStoreDay")
    def log_store_day(self) -> Optional[pulumi.Input[int]]:
        """
        The retention period of raw logs after you enable the log auditing feature. Unit: day. Valid values: `180`, `30`, `365`, `90`. **NOTE:** The`log_store_day` is valid when the `audit_status` is `1`.
        """
        return pulumi.get(self, "log_store_day")

    @log_store_day.setter
    def log_store_day(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "log_store_day", value)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the data asset.
        """
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_id", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password that is used to connect to the database.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port that is used to connect to the database.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serviceRegionId")
    def service_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The region ID of the data asset.
        """
        return pulumi.get(self, "service_region_id")

    @service_region_id.setter
    def service_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_region_id", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the service to which the data asset belongs.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


@pulumi.input_type
class _DataLimitState:
    def __init__(__self__, *,
                 audit_status: Optional[pulumi.Input[int]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 log_store_day: Optional[pulumi.Input[int]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 service_region_id: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DataLimit resources.
        :param pulumi.Input[int] audit_status: Whether to enable the log auditing feature. Valid values: `0`, `1`.
        :param pulumi.Input[str] engine_type: The type of the database. Valid values: `MySQL`, `SQLServer`.
        :param pulumi.Input[str] lang: The lang.
        :param pulumi.Input[int] log_store_day: The retention period of raw logs after you enable the log auditing feature. Unit: day. Valid values: `180`, `30`, `365`, `90`. **NOTE:** The`log_store_day` is valid when the `audit_status` is `1`.
        :param pulumi.Input[str] parent_id: The ID of the data asset.
        :param pulumi.Input[str] password: The password that is used to connect to the database.
        :param pulumi.Input[int] port: The port that is used to connect to the database.
        :param pulumi.Input[str] resource_type: The type of the service to which the data asset belongs. Valid values: `MaxCompute`, `OSS`, `RDS`.
        :param pulumi.Input[str] service_region_id: The region ID of the data asset.
        :param pulumi.Input[str] user_name: The name of the service to which the data asset belongs.
        """
        if audit_status is not None:
            pulumi.set(__self__, "audit_status", audit_status)
        if engine_type is not None:
            pulumi.set(__self__, "engine_type", engine_type)
        if lang is not None:
            pulumi.set(__self__, "lang", lang)
        if log_store_day is not None:
            pulumi.set(__self__, "log_store_day", log_store_day)
        if parent_id is not None:
            pulumi.set(__self__, "parent_id", parent_id)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if service_region_id is not None:
            pulumi.set(__self__, "service_region_id", service_region_id)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="auditStatus")
    def audit_status(self) -> Optional[pulumi.Input[int]]:
        """
        Whether to enable the log auditing feature. Valid values: `0`, `1`.
        """
        return pulumi.get(self, "audit_status")

    @audit_status.setter
    def audit_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "audit_status", value)

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the database. Valid values: `MySQL`, `SQLServer`.
        """
        return pulumi.get(self, "engine_type")

    @engine_type.setter
    def engine_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_type", value)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        The lang.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)

    @property
    @pulumi.getter(name="logStoreDay")
    def log_store_day(self) -> Optional[pulumi.Input[int]]:
        """
        The retention period of raw logs after you enable the log auditing feature. Unit: day. Valid values: `180`, `30`, `365`, `90`. **NOTE:** The`log_store_day` is valid when the `audit_status` is `1`.
        """
        return pulumi.get(self, "log_store_day")

    @log_store_day.setter
    def log_store_day(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "log_store_day", value)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the data asset.
        """
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_id", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password that is used to connect to the database.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port that is used to connect to the database.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the service to which the data asset belongs. Valid values: `MaxCompute`, `OSS`, `RDS`.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter(name="serviceRegionId")
    def service_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The region ID of the data asset.
        """
        return pulumi.get(self, "service_region_id")

    @service_region_id.setter
    def service_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_region_id", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the service to which the data asset belongs.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


class DataLimit(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audit_status: Optional[pulumi.Input[int]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 log_store_day: Optional[pulumi.Input[int]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 service_region_id: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Data Security Center Data Limit resource.

        For information about Data Security Center Data Limit and how to use it, see [What is Data Limit](https://www.alibabacloud.com/help/en/doc-detail/158987.html).

        > **NOTE:** Available in v1.159.0+.

        ## Import

        Data Security Center Data Limit can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:sddp/dataLimit:DataLimit example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] audit_status: Whether to enable the log auditing feature. Valid values: `0`, `1`.
        :param pulumi.Input[str] engine_type: The type of the database. Valid values: `MySQL`, `SQLServer`.
        :param pulumi.Input[str] lang: The lang.
        :param pulumi.Input[int] log_store_day: The retention period of raw logs after you enable the log auditing feature. Unit: day. Valid values: `180`, `30`, `365`, `90`. **NOTE:** The`log_store_day` is valid when the `audit_status` is `1`.
        :param pulumi.Input[str] parent_id: The ID of the data asset.
        :param pulumi.Input[str] password: The password that is used to connect to the database.
        :param pulumi.Input[int] port: The port that is used to connect to the database.
        :param pulumi.Input[str] resource_type: The type of the service to which the data asset belongs. Valid values: `MaxCompute`, `OSS`, `RDS`.
        :param pulumi.Input[str] service_region_id: The region ID of the data asset.
        :param pulumi.Input[str] user_name: The name of the service to which the data asset belongs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataLimitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Data Security Center Data Limit resource.

        For information about Data Security Center Data Limit and how to use it, see [What is Data Limit](https://www.alibabacloud.com/help/en/doc-detail/158987.html).

        > **NOTE:** Available in v1.159.0+.

        ## Import

        Data Security Center Data Limit can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:sddp/dataLimit:DataLimit example <id>
        ```

        :param str resource_name: The name of the resource.
        :param DataLimitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataLimitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audit_status: Optional[pulumi.Input[int]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 log_store_day: Optional[pulumi.Input[int]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 service_region_id: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = DataLimitArgs.__new__(DataLimitArgs)

            __props__.__dict__["audit_status"] = audit_status
            __props__.__dict__["engine_type"] = engine_type
            __props__.__dict__["lang"] = lang
            __props__.__dict__["log_store_day"] = log_store_day
            __props__.__dict__["parent_id"] = parent_id
            __props__.__dict__["password"] = password
            __props__.__dict__["port"] = port
            if resource_type is None and not opts.urn:
                raise TypeError("Missing required property 'resource_type'")
            __props__.__dict__["resource_type"] = resource_type
            __props__.__dict__["service_region_id"] = service_region_id
            __props__.__dict__["user_name"] = user_name
        super(DataLimit, __self__).__init__(
            'alicloud:sddp/dataLimit:DataLimit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            audit_status: Optional[pulumi.Input[int]] = None,
            engine_type: Optional[pulumi.Input[str]] = None,
            lang: Optional[pulumi.Input[str]] = None,
            log_store_day: Optional[pulumi.Input[int]] = None,
            parent_id: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            resource_type: Optional[pulumi.Input[str]] = None,
            service_region_id: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'DataLimit':
        """
        Get an existing DataLimit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] audit_status: Whether to enable the log auditing feature. Valid values: `0`, `1`.
        :param pulumi.Input[str] engine_type: The type of the database. Valid values: `MySQL`, `SQLServer`.
        :param pulumi.Input[str] lang: The lang.
        :param pulumi.Input[int] log_store_day: The retention period of raw logs after you enable the log auditing feature. Unit: day. Valid values: `180`, `30`, `365`, `90`. **NOTE:** The`log_store_day` is valid when the `audit_status` is `1`.
        :param pulumi.Input[str] parent_id: The ID of the data asset.
        :param pulumi.Input[str] password: The password that is used to connect to the database.
        :param pulumi.Input[int] port: The port that is used to connect to the database.
        :param pulumi.Input[str] resource_type: The type of the service to which the data asset belongs. Valid values: `MaxCompute`, `OSS`, `RDS`.
        :param pulumi.Input[str] service_region_id: The region ID of the data asset.
        :param pulumi.Input[str] user_name: The name of the service to which the data asset belongs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataLimitState.__new__(_DataLimitState)

        __props__.__dict__["audit_status"] = audit_status
        __props__.__dict__["engine_type"] = engine_type
        __props__.__dict__["lang"] = lang
        __props__.__dict__["log_store_day"] = log_store_day
        __props__.__dict__["parent_id"] = parent_id
        __props__.__dict__["password"] = password
        __props__.__dict__["port"] = port
        __props__.__dict__["resource_type"] = resource_type
        __props__.__dict__["service_region_id"] = service_region_id
        __props__.__dict__["user_name"] = user_name
        return DataLimit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="auditStatus")
    def audit_status(self) -> pulumi.Output[int]:
        """
        Whether to enable the log auditing feature. Valid values: `0`, `1`.
        """
        return pulumi.get(self, "audit_status")

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the database. Valid values: `MySQL`, `SQLServer`.
        """
        return pulumi.get(self, "engine_type")

    @property
    @pulumi.getter
    def lang(self) -> pulumi.Output[Optional[str]]:
        """
        The lang.
        """
        return pulumi.get(self, "lang")

    @property
    @pulumi.getter(name="logStoreDay")
    def log_store_day(self) -> pulumi.Output[Optional[int]]:
        """
        The retention period of raw logs after you enable the log auditing feature. Unit: day. Valid values: `180`, `30`, `365`, `90`. **NOTE:** The`log_store_day` is valid when the `audit_status` is `1`.
        """
        return pulumi.get(self, "log_store_day")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the data asset.
        """
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password that is used to connect to the database.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[int]]:
        """
        The port that is used to connect to the database.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[str]:
        """
        The type of the service to which the data asset belongs. Valid values: `MaxCompute`, `OSS`, `RDS`.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="serviceRegionId")
    def service_region_id(self) -> pulumi.Output[Optional[str]]:
        """
        The region ID of the data asset.
        """
        return pulumi.get(self, "service_region_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the service to which the data asset belongs.
        """
        return pulumi.get(self, "user_name")

