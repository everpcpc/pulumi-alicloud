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
                 description: pulumi.Input[str],
                 license_code: pulumi.Input[str],
                 security_group_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 vswitch_id: pulumi.Input[str],
                 period: Optional[pulumi.Input[int]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[str] description: Description of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[str] license_code: The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: security group IDs configured to Bastionhost.
        :param pulumi.Input[str] vswitch_id: VSwitch ID configured to Bastionhost.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "license_code", license_code)
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "vswitch_id", vswitch_id)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        Description of the instance. This name can have a string of 1 to 63 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="licenseCode")
    def license_code(self) -> pulumi.Input[str]:
        """
        The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
        """
        return pulumi.get(self, "license_code")

    @license_code.setter
    def license_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "license_code", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        security group IDs configured to Bastionhost.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Input[str]:
        """
        VSwitch ID configured to Bastionhost.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vswitch_id", value)

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
        The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
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
class _InstanceState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 license_code: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Instance resources.
        :param pulumi.Input[str] description: Description of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[str] license_code: The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: security group IDs configured to Bastionhost.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vswitch_id: VSwitch ID configured to Bastionhost.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if license_code is not None:
            pulumi.set(__self__, "license_code", license_code)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the instance. This name can have a string of 1 to 63 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="licenseCode")
    def license_code(self) -> Optional[pulumi.Input[str]]:
        """
        The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
        """
        return pulumi.get(self, "license_code")

    @license_code.setter
    def license_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "license_code", value)

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
        The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        security group IDs configured to Bastionhost.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

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

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        VSwitch ID configured to Bastionhost.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 license_code: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        > **NOTE:** From the version 1.132.0, the resource has been renamed to `bastionhost.Instance`.

        Cloud Bastion Host instance resource ("Yundun_bastionhost" is the short term of this product).
        For information about Resource Manager Resource Directory and how to use it, see [What is Bastionhost](https://www.alibabacloud.com/help/en/doc-detail/52922.htm).

        > **NOTE:** The endpoint of bssopenapi used only support "business.aliyuncs.com" at present.

        > **NOTE:** Available in 1.63.0+ .

        > **NOTE:** In order to destroy Cloud Bastionhost instance , users are required to apply for white list first

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.bastionhost.Instance("default",
            description="Terraform-test",
            license_code="bhah_ent_50_asset",
            period=1,
            security_group_ids=[
                "sg-test",
                "sg-12345",
            ],
            vswitch_id="v-testVswitch")
        ```

        ## Import

        Yundun_bastionhost instance can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:bastionhost/instance:Instance example bastionhost-exampe123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[str] license_code: The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: security group IDs configured to Bastionhost.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vswitch_id: VSwitch ID configured to Bastionhost.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        > **NOTE:** From the version 1.132.0, the resource has been renamed to `bastionhost.Instance`.

        Cloud Bastion Host instance resource ("Yundun_bastionhost" is the short term of this product).
        For information about Resource Manager Resource Directory and how to use it, see [What is Bastionhost](https://www.alibabacloud.com/help/en/doc-detail/52922.htm).

        > **NOTE:** The endpoint of bssopenapi used only support "business.aliyuncs.com" at present.

        > **NOTE:** Available in 1.63.0+ .

        > **NOTE:** In order to destroy Cloud Bastionhost instance , users are required to apply for white list first

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.bastionhost.Instance("default",
            description="Terraform-test",
            license_code="bhah_ent_50_asset",
            period=1,
            security_group_ids=[
                "sg-test",
                "sg-12345",
            ],
            vswitch_id="v-testVswitch")
        ```

        ## Import

        Yundun_bastionhost instance can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:bastionhost/instance:Instance example bastionhost-exampe123456
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
                 description: Optional[pulumi.Input[str]] = None,
                 license_code: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
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

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if license_code is None and not opts.urn:
                raise TypeError("Missing required property 'license_code'")
            __props__.__dict__["license_code"] = license_code
            __props__.__dict__["period"] = period
            __props__.__dict__["resource_group_id"] = resource_group_id
            if security_group_ids is None and not opts.urn:
                raise TypeError("Missing required property 'security_group_ids'")
            __props__.__dict__["security_group_ids"] = security_group_ids
            __props__.__dict__["tags"] = tags
            if vswitch_id is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__.__dict__["vswitch_id"] = vswitch_id
        super(Instance, __self__).__init__(
            'alicloud:bastionhost/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            license_code: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[int]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[str] license_code: The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: security group IDs configured to Bastionhost.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vswitch_id: VSwitch ID configured to Bastionhost.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceState.__new__(_InstanceState)

        __props__.__dict__["description"] = description
        __props__.__dict__["license_code"] = license_code
        __props__.__dict__["period"] = period
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["security_group_ids"] = security_group_ids
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vswitch_id"] = vswitch_id
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of the instance. This name can have a string of 1 to 63 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="licenseCode")
    def license_code(self) -> pulumi.Output[str]:
        """
        The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
        """
        return pulumi.get(self, "license_code")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        security group IDs configured to Bastionhost.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[str]:
        """
        VSwitch ID configured to Bastionhost.
        """
        return pulumi.get(self, "vswitch_id")

