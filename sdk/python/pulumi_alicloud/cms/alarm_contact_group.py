# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AlarmContactGroupArgs', 'AlarmContactGroup']

@pulumi.input_type
class AlarmContactGroupArgs:
    def __init__(__self__, *,
                 alarm_contact_group_name: pulumi.Input[str],
                 contacts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 describe: Optional[pulumi.Input[str]] = None,
                 enable_subscribed: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AlarmContactGroup resource.
        :param pulumi.Input[str] alarm_contact_group_name: The name of the alarm group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contacts: The name of the alert contact.
        :param pulumi.Input[str] describe: The description of the alert group.
        :param pulumi.Input[bool] enable_subscribed: Whether to open weekly subscription.
        """
        pulumi.set(__self__, "alarm_contact_group_name", alarm_contact_group_name)
        if contacts is not None:
            pulumi.set(__self__, "contacts", contacts)
        if describe is not None:
            pulumi.set(__self__, "describe", describe)
        if enable_subscribed is not None:
            pulumi.set(__self__, "enable_subscribed", enable_subscribed)

    @property
    @pulumi.getter(name="alarmContactGroupName")
    def alarm_contact_group_name(self) -> pulumi.Input[str]:
        """
        The name of the alarm group.
        """
        return pulumi.get(self, "alarm_contact_group_name")

    @alarm_contact_group_name.setter
    def alarm_contact_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "alarm_contact_group_name", value)

    @property
    @pulumi.getter
    def contacts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The name of the alert contact.
        """
        return pulumi.get(self, "contacts")

    @contacts.setter
    def contacts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "contacts", value)

    @property
    @pulumi.getter
    def describe(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the alert group.
        """
        return pulumi.get(self, "describe")

    @describe.setter
    def describe(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "describe", value)

    @property
    @pulumi.getter(name="enableSubscribed")
    def enable_subscribed(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to open weekly subscription.
        """
        return pulumi.get(self, "enable_subscribed")

    @enable_subscribed.setter
    def enable_subscribed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_subscribed", value)


@pulumi.input_type
class _AlarmContactGroupState:
    def __init__(__self__, *,
                 alarm_contact_group_name: Optional[pulumi.Input[str]] = None,
                 contacts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 describe: Optional[pulumi.Input[str]] = None,
                 enable_subscribed: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering AlarmContactGroup resources.
        :param pulumi.Input[str] alarm_contact_group_name: The name of the alarm group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contacts: The name of the alert contact.
        :param pulumi.Input[str] describe: The description of the alert group.
        :param pulumi.Input[bool] enable_subscribed: Whether to open weekly subscription.
        """
        if alarm_contact_group_name is not None:
            pulumi.set(__self__, "alarm_contact_group_name", alarm_contact_group_name)
        if contacts is not None:
            pulumi.set(__self__, "contacts", contacts)
        if describe is not None:
            pulumi.set(__self__, "describe", describe)
        if enable_subscribed is not None:
            pulumi.set(__self__, "enable_subscribed", enable_subscribed)

    @property
    @pulumi.getter(name="alarmContactGroupName")
    def alarm_contact_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the alarm group.
        """
        return pulumi.get(self, "alarm_contact_group_name")

    @alarm_contact_group_name.setter
    def alarm_contact_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alarm_contact_group_name", value)

    @property
    @pulumi.getter
    def contacts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The name of the alert contact.
        """
        return pulumi.get(self, "contacts")

    @contacts.setter
    def contacts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "contacts", value)

    @property
    @pulumi.getter
    def describe(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the alert group.
        """
        return pulumi.get(self, "describe")

    @describe.setter
    def describe(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "describe", value)

    @property
    @pulumi.getter(name="enableSubscribed")
    def enable_subscribed(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to open weekly subscription.
        """
        return pulumi.get(self, "enable_subscribed")

    @enable_subscribed.setter
    def enable_subscribed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_subscribed", value)


class AlarmContactGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alarm_contact_group_name: Optional[pulumi.Input[str]] = None,
                 contacts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 describe: Optional[pulumi.Input[str]] = None,
                 enable_subscribed: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Provides a CMS Alarm Contact Group resource.

        For information about CMS Alarm Contact Group and how to use it, see [What is Alarm Contact Group](https://www.alibabacloud.com/help/en/doc-detail/114929.htm).

        > **NOTE:** Available in v1.101.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.cms.AlarmContactGroup("example", alarm_contact_group_name="tf-test")
        ```

        ## Import

        CMS Alarm Contact Group can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cms/alarmContactGroup:AlarmContactGroup example tf-testacc123
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alarm_contact_group_name: The name of the alarm group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contacts: The name of the alert contact.
        :param pulumi.Input[str] describe: The description of the alert group.
        :param pulumi.Input[bool] enable_subscribed: Whether to open weekly subscription.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AlarmContactGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CMS Alarm Contact Group resource.

        For information about CMS Alarm Contact Group and how to use it, see [What is Alarm Contact Group](https://www.alibabacloud.com/help/en/doc-detail/114929.htm).

        > **NOTE:** Available in v1.101.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.cms.AlarmContactGroup("example", alarm_contact_group_name="tf-test")
        ```

        ## Import

        CMS Alarm Contact Group can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cms/alarmContactGroup:AlarmContactGroup example tf-testacc123
        ```

        :param str resource_name: The name of the resource.
        :param AlarmContactGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AlarmContactGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alarm_contact_group_name: Optional[pulumi.Input[str]] = None,
                 contacts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 describe: Optional[pulumi.Input[str]] = None,
                 enable_subscribed: Optional[pulumi.Input[bool]] = None,
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
            __props__ = AlarmContactGroupArgs.__new__(AlarmContactGroupArgs)

            if alarm_contact_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'alarm_contact_group_name'")
            __props__.__dict__["alarm_contact_group_name"] = alarm_contact_group_name
            __props__.__dict__["contacts"] = contacts
            __props__.__dict__["describe"] = describe
            __props__.__dict__["enable_subscribed"] = enable_subscribed
        super(AlarmContactGroup, __self__).__init__(
            'alicloud:cms/alarmContactGroup:AlarmContactGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alarm_contact_group_name: Optional[pulumi.Input[str]] = None,
            contacts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            describe: Optional[pulumi.Input[str]] = None,
            enable_subscribed: Optional[pulumi.Input[bool]] = None) -> 'AlarmContactGroup':
        """
        Get an existing AlarmContactGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alarm_contact_group_name: The name of the alarm group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contacts: The name of the alert contact.
        :param pulumi.Input[str] describe: The description of the alert group.
        :param pulumi.Input[bool] enable_subscribed: Whether to open weekly subscription.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AlarmContactGroupState.__new__(_AlarmContactGroupState)

        __props__.__dict__["alarm_contact_group_name"] = alarm_contact_group_name
        __props__.__dict__["contacts"] = contacts
        __props__.__dict__["describe"] = describe
        __props__.__dict__["enable_subscribed"] = enable_subscribed
        return AlarmContactGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alarmContactGroupName")
    def alarm_contact_group_name(self) -> pulumi.Output[str]:
        """
        The name of the alarm group.
        """
        return pulumi.get(self, "alarm_contact_group_name")

    @property
    @pulumi.getter
    def contacts(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The name of the alert contact.
        """
        return pulumi.get(self, "contacts")

    @property
    @pulumi.getter
    def describe(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the alert group.
        """
        return pulumi.get(self, "describe")

    @property
    @pulumi.getter(name="enableSubscribed")
    def enable_subscribed(self) -> pulumi.Output[bool]:
        """
        Whether to open weekly subscription.
        """
        return pulumi.get(self, "enable_subscribed")

