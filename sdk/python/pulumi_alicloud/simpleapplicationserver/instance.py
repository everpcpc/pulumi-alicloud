# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceArgs', 'Instance']

@pulumi.input_type
class InstanceArgs:
    def __init__(__self__, *,
                 image_id: pulumi.Input[str],
                 period: pulumi.Input[int],
                 plan_id: pulumi.Input[str],
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 auto_renew_period: Optional[pulumi.Input[int]] = None,
                 data_disk_size: Optional[pulumi.Input[int]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 payment_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[str] image_id: The ID of the image.  You can use the `simpleapplicationserver.get_images` to query the available images in the specified region. The value must be an integral multiple of 20.
        :param pulumi.Input[int] period: The period. Unit: months. Valid values: `1`,`3`, `6`, `12`, `24`, `36`.
        :param pulumi.Input[str] plan_id: The ID of the plan. You can use the `simpleapplicationserver.get_server_plans`  to query all the plans provided by Simple Application Server in the specified region.
        :param pulumi.Input[bool] auto_renew: Specifies whether to enable auto-renewal. Unit: months. Valid values: `true` and `false`.
        :param pulumi.Input[int] auto_renew_period: The auto renew period. Valid values: `1`,`3`, `6`, `12`, `24`, `36`. **NOTE:** The attribute `auto_renew` is valid when the attribute is `true`.
        :param pulumi.Input[int] data_disk_size: The size of the data disk. Unit: GB. Valid values: `0` to `16380`.
        :param pulumi.Input[str] instance_name: The name of the simple application server.
        :param pulumi.Input[str] password: The password of the simple application server. The password must be 8 to 30 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include: `( ) ~ ! @ # $ % ^ & * - + = | { } [ ] : ; < > , . ? /`.
        :param pulumi.Input[str] payment_type: The paymen type of the resource. Valid values: `Subscription`.
        :param pulumi.Input[str] status: The status of the simple application server. Valid values: `Resetting`, `Running`, `Stopped`.
        """
        pulumi.set(__self__, "image_id", image_id)
        pulumi.set(__self__, "period", period)
        pulumi.set(__self__, "plan_id", plan_id)
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if auto_renew_period is not None:
            pulumi.set(__self__, "auto_renew_period", auto_renew_period)
        if data_disk_size is not None:
            pulumi.set(__self__, "data_disk_size", data_disk_size)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if payment_type is not None:
            pulumi.set(__self__, "payment_type", payment_type)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Input[str]:
        """
        The ID of the image.  You can use the `simpleapplicationserver.get_images` to query the available images in the specified region. The value must be an integral multiple of 20.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter
    def period(self) -> pulumi.Input[int]:
        """
        The period. Unit: months. Valid values: `1`,`3`, `6`, `12`, `24`, `36`.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: pulumi.Input[int]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> pulumi.Input[str]:
        """
        The ID of the plan. You can use the `simpleapplicationserver.get_server_plans`  to query all the plans provided by Simple Application Server in the specified region.
        """
        return pulumi.get(self, "plan_id")

    @plan_id.setter
    def plan_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "plan_id", value)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable auto-renewal. Unit: months. Valid values: `true` and `false`.
        """
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter(name="autoRenewPeriod")
    def auto_renew_period(self) -> Optional[pulumi.Input[int]]:
        """
        The auto renew period. Valid values: `1`,`3`, `6`, `12`, `24`, `36`. **NOTE:** The attribute `auto_renew` is valid when the attribute is `true`.
        """
        return pulumi.get(self, "auto_renew_period")

    @auto_renew_period.setter
    def auto_renew_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_renew_period", value)

    @property
    @pulumi.getter(name="dataDiskSize")
    def data_disk_size(self) -> Optional[pulumi.Input[int]]:
        """
        The size of the data disk. Unit: GB. Valid values: `0` to `16380`.
        """
        return pulumi.get(self, "data_disk_size")

    @data_disk_size.setter
    def data_disk_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "data_disk_size", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the simple application server.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the simple application server. The password must be 8 to 30 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include: `( ) ~ ! @ # $ % ^ & * - + = | { } [ ] : ; < > , . ? /`.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> Optional[pulumi.Input[str]]:
        """
        The paymen type of the resource. Valid values: `Subscription`.
        """
        return pulumi.get(self, "payment_type")

    @payment_type.setter
    def payment_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payment_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the simple application server. Valid values: `Resetting`, `Running`, `Stopped`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _InstanceState:
    def __init__(__self__, *,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 auto_renew_period: Optional[pulumi.Input[int]] = None,
                 data_disk_size: Optional[pulumi.Input[int]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 payment_type: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 plan_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Instance resources.
        :param pulumi.Input[bool] auto_renew: Specifies whether to enable auto-renewal. Unit: months. Valid values: `true` and `false`.
        :param pulumi.Input[int] auto_renew_period: The auto renew period. Valid values: `1`,`3`, `6`, `12`, `24`, `36`. **NOTE:** The attribute `auto_renew` is valid when the attribute is `true`.
        :param pulumi.Input[int] data_disk_size: The size of the data disk. Unit: GB. Valid values: `0` to `16380`.
        :param pulumi.Input[str] image_id: The ID of the image.  You can use the `simpleapplicationserver.get_images` to query the available images in the specified region. The value must be an integral multiple of 20.
        :param pulumi.Input[str] instance_name: The name of the simple application server.
        :param pulumi.Input[str] password: The password of the simple application server. The password must be 8 to 30 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include: `( ) ~ ! @ # $ % ^ & * - + = | { } [ ] : ; < > , . ? /`.
        :param pulumi.Input[str] payment_type: The paymen type of the resource. Valid values: `Subscription`.
        :param pulumi.Input[int] period: The period. Unit: months. Valid values: `1`,`3`, `6`, `12`, `24`, `36`.
        :param pulumi.Input[str] plan_id: The ID of the plan. You can use the `simpleapplicationserver.get_server_plans`  to query all the plans provided by Simple Application Server in the specified region.
        :param pulumi.Input[str] status: The status of the simple application server. Valid values: `Resetting`, `Running`, `Stopped`.
        """
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if auto_renew_period is not None:
            pulumi.set(__self__, "auto_renew_period", auto_renew_period)
        if data_disk_size is not None:
            pulumi.set(__self__, "data_disk_size", data_disk_size)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if payment_type is not None:
            pulumi.set(__self__, "payment_type", payment_type)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if plan_id is not None:
            pulumi.set(__self__, "plan_id", plan_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable auto-renewal. Unit: months. Valid values: `true` and `false`.
        """
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter(name="autoRenewPeriod")
    def auto_renew_period(self) -> Optional[pulumi.Input[int]]:
        """
        The auto renew period. Valid values: `1`,`3`, `6`, `12`, `24`, `36`. **NOTE:** The attribute `auto_renew` is valid when the attribute is `true`.
        """
        return pulumi.get(self, "auto_renew_period")

    @auto_renew_period.setter
    def auto_renew_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_renew_period", value)

    @property
    @pulumi.getter(name="dataDiskSize")
    def data_disk_size(self) -> Optional[pulumi.Input[int]]:
        """
        The size of the data disk. Unit: GB. Valid values: `0` to `16380`.
        """
        return pulumi.get(self, "data_disk_size")

    @data_disk_size.setter
    def data_disk_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "data_disk_size", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the image.  You can use the `simpleapplicationserver.get_images` to query the available images in the specified region. The value must be an integral multiple of 20.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the simple application server.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the simple application server. The password must be 8 to 30 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include: `( ) ~ ! @ # $ % ^ & * - + = | { } [ ] : ; < > , . ? /`.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> Optional[pulumi.Input[str]]:
        """
        The paymen type of the resource. Valid values: `Subscription`.
        """
        return pulumi.get(self, "payment_type")

    @payment_type.setter
    def payment_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payment_type", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        The period. Unit: months. Valid values: `1`,`3`, `6`, `12`, `24`, `36`.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the plan. You can use the `simpleapplicationserver.get_server_plans`  to query all the plans provided by Simple Application Server in the specified region.
        """
        return pulumi.get(self, "plan_id")

    @plan_id.setter
    def plan_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the simple application server. Valid values: `Resetting`, `Running`, `Stopped`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 auto_renew_period: Optional[pulumi.Input[int]] = None,
                 data_disk_size: Optional[pulumi.Input[int]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 payment_type: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 plan_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Simple Application Server Instance resource.

        For information about Simple Application Server Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/doc-detail/190440.htm).

        > **NOTE:** Available in v1.135.0+.

        ## Import

        Simple Application Server Instance can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:simpleapplicationserver/instance:Instance example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renew: Specifies whether to enable auto-renewal. Unit: months. Valid values: `true` and `false`.
        :param pulumi.Input[int] auto_renew_period: The auto renew period. Valid values: `1`,`3`, `6`, `12`, `24`, `36`. **NOTE:** The attribute `auto_renew` is valid when the attribute is `true`.
        :param pulumi.Input[int] data_disk_size: The size of the data disk. Unit: GB. Valid values: `0` to `16380`.
        :param pulumi.Input[str] image_id: The ID of the image.  You can use the `simpleapplicationserver.get_images` to query the available images in the specified region. The value must be an integral multiple of 20.
        :param pulumi.Input[str] instance_name: The name of the simple application server.
        :param pulumi.Input[str] password: The password of the simple application server. The password must be 8 to 30 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include: `( ) ~ ! @ # $ % ^ & * - + = | { } [ ] : ; < > , . ? /`.
        :param pulumi.Input[str] payment_type: The paymen type of the resource. Valid values: `Subscription`.
        :param pulumi.Input[int] period: The period. Unit: months. Valid values: `1`,`3`, `6`, `12`, `24`, `36`.
        :param pulumi.Input[str] plan_id: The ID of the plan. You can use the `simpleapplicationserver.get_server_plans`  to query all the plans provided by Simple Application Server in the specified region.
        :param pulumi.Input[str] status: The status of the simple application server. Valid values: `Resetting`, `Running`, `Stopped`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Simple Application Server Instance resource.

        For information about Simple Application Server Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/doc-detail/190440.htm).

        > **NOTE:** Available in v1.135.0+.

        ## Import

        Simple Application Server Instance can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:simpleapplicationserver/instance:Instance example <id>
        ```

        :param str resource_name: The name of the resource.
        :param InstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 auto_renew_period: Optional[pulumi.Input[int]] = None,
                 data_disk_size: Optional[pulumi.Input[int]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 payment_type: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 plan_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
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
            __props__ = InstanceArgs.__new__(InstanceArgs)

            __props__.__dict__["auto_renew"] = auto_renew
            __props__.__dict__["auto_renew_period"] = auto_renew_period
            __props__.__dict__["data_disk_size"] = data_disk_size
            if image_id is None and not opts.urn:
                raise TypeError("Missing required property 'image_id'")
            __props__.__dict__["image_id"] = image_id
            __props__.__dict__["instance_name"] = instance_name
            __props__.__dict__["password"] = password
            __props__.__dict__["payment_type"] = payment_type
            if period is None and not opts.urn:
                raise TypeError("Missing required property 'period'")
            __props__.__dict__["period"] = period
            if plan_id is None and not opts.urn:
                raise TypeError("Missing required property 'plan_id'")
            __props__.__dict__["plan_id"] = plan_id
            __props__.__dict__["status"] = status
        super(Instance, __self__).__init__(
            'alicloud:simpleapplicationserver/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_renew: Optional[pulumi.Input[bool]] = None,
            auto_renew_period: Optional[pulumi.Input[int]] = None,
            data_disk_size: Optional[pulumi.Input[int]] = None,
            image_id: Optional[pulumi.Input[str]] = None,
            instance_name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            payment_type: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[int]] = None,
            plan_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renew: Specifies whether to enable auto-renewal. Unit: months. Valid values: `true` and `false`.
        :param pulumi.Input[int] auto_renew_period: The auto renew period. Valid values: `1`,`3`, `6`, `12`, `24`, `36`. **NOTE:** The attribute `auto_renew` is valid when the attribute is `true`.
        :param pulumi.Input[int] data_disk_size: The size of the data disk. Unit: GB. Valid values: `0` to `16380`.
        :param pulumi.Input[str] image_id: The ID of the image.  You can use the `simpleapplicationserver.get_images` to query the available images in the specified region. The value must be an integral multiple of 20.
        :param pulumi.Input[str] instance_name: The name of the simple application server.
        :param pulumi.Input[str] password: The password of the simple application server. The password must be 8 to 30 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include: `( ) ~ ! @ # $ % ^ & * - + = | { } [ ] : ; < > , . ? /`.
        :param pulumi.Input[str] payment_type: The paymen type of the resource. Valid values: `Subscription`.
        :param pulumi.Input[int] period: The period. Unit: months. Valid values: `1`,`3`, `6`, `12`, `24`, `36`.
        :param pulumi.Input[str] plan_id: The ID of the plan. You can use the `simpleapplicationserver.get_server_plans`  to query all the plans provided by Simple Application Server in the specified region.
        :param pulumi.Input[str] status: The status of the simple application server. Valid values: `Resetting`, `Running`, `Stopped`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceState.__new__(_InstanceState)

        __props__.__dict__["auto_renew"] = auto_renew
        __props__.__dict__["auto_renew_period"] = auto_renew_period
        __props__.__dict__["data_disk_size"] = data_disk_size
        __props__.__dict__["image_id"] = image_id
        __props__.__dict__["instance_name"] = instance_name
        __props__.__dict__["password"] = password
        __props__.__dict__["payment_type"] = payment_type
        __props__.__dict__["period"] = period
        __props__.__dict__["plan_id"] = plan_id
        __props__.__dict__["status"] = status
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to enable auto-renewal. Unit: months. Valid values: `true` and `false`.
        """
        return pulumi.get(self, "auto_renew")

    @property
    @pulumi.getter(name="autoRenewPeriod")
    def auto_renew_period(self) -> pulumi.Output[Optional[int]]:
        """
        The auto renew period. Valid values: `1`,`3`, `6`, `12`, `24`, `36`. **NOTE:** The attribute `auto_renew` is valid when the attribute is `true`.
        """
        return pulumi.get(self, "auto_renew_period")

    @property
    @pulumi.getter(name="dataDiskSize")
    def data_disk_size(self) -> pulumi.Output[Optional[int]]:
        """
        The size of the data disk. Unit: GB. Valid values: `0` to `16380`.
        """
        return pulumi.get(self, "data_disk_size")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[str]:
        """
        The ID of the image.  You can use the `simpleapplicationserver.get_images` to query the available images in the specified region. The value must be an integral multiple of 20.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the simple application server.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password of the simple application server. The password must be 8 to 30 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include: `( ) ~ ! @ # $ % ^ & * - + = | { } [ ] : ; < > , . ? /`.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> pulumi.Output[str]:
        """
        The paymen type of the resource. Valid values: `Subscription`.
        """
        return pulumi.get(self, "payment_type")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[int]:
        """
        The period. Unit: months. Valid values: `1`,`3`, `6`, `12`, `24`, `36`.
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> pulumi.Output[str]:
        """
        The ID of the plan. You can use the `simpleapplicationserver.get_server_plans`  to query all the plans provided by Simple Application Server in the specified region.
        """
        return pulumi.get(self, "plan_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the simple application server. Valid values: `Resetting`, `Running`, `Stopped`.
        """
        return pulumi.get(self, "status")

