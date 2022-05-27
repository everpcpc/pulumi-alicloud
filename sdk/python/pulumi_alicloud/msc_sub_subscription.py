# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MscSubSubscriptionArgs', 'MscSubSubscription']

@pulumi.input_type
class MscSubSubscriptionArgs:
    def __init__(__self__, *,
                 item_name: pulumi.Input[str],
                 contact_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 email_status: Optional[pulumi.Input[int]] = None,
                 pmsg_status: Optional[pulumi.Input[int]] = None,
                 sms_status: Optional[pulumi.Input[int]] = None,
                 tts_status: Optional[pulumi.Input[int]] = None,
                 webhook_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 webhook_status: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a MscSubSubscription resource.
        :param pulumi.Input[str] item_name: The name of the Subscription. **NOTE:**  You should use the `get_msc_sub_subscriptions` to query the available subscription item name.
        :param pulumi.Input[int] email_status: The status of email subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[int] pmsg_status: The status of pmsg subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[int] sms_status: The status of sms subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[int] tts_status: The status of tts subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] webhook_ids: The ids of subscribed webhooks.
        :param pulumi.Input[int] webhook_status: The status of webhook subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        pulumi.set(__self__, "item_name", item_name)
        if contact_ids is not None:
            pulumi.set(__self__, "contact_ids", contact_ids)
        if email_status is not None:
            pulumi.set(__self__, "email_status", email_status)
        if pmsg_status is not None:
            pulumi.set(__self__, "pmsg_status", pmsg_status)
        if sms_status is not None:
            pulumi.set(__self__, "sms_status", sms_status)
        if tts_status is not None:
            pulumi.set(__self__, "tts_status", tts_status)
        if webhook_ids is not None:
            pulumi.set(__self__, "webhook_ids", webhook_ids)
        if webhook_status is not None:
            pulumi.set(__self__, "webhook_status", webhook_status)

    @property
    @pulumi.getter(name="itemName")
    def item_name(self) -> pulumi.Input[str]:
        """
        The name of the Subscription. **NOTE:**  You should use the `get_msc_sub_subscriptions` to query the available subscription item name.
        """
        return pulumi.get(self, "item_name")

    @item_name.setter
    def item_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "item_name", value)

    @property
    @pulumi.getter(name="contactIds")
    def contact_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "contact_ids")

    @contact_ids.setter
    def contact_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "contact_ids", value)

    @property
    @pulumi.getter(name="emailStatus")
    def email_status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of email subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        return pulumi.get(self, "email_status")

    @email_status.setter
    def email_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "email_status", value)

    @property
    @pulumi.getter(name="pmsgStatus")
    def pmsg_status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of pmsg subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        return pulumi.get(self, "pmsg_status")

    @pmsg_status.setter
    def pmsg_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pmsg_status", value)

    @property
    @pulumi.getter(name="smsStatus")
    def sms_status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of sms subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        return pulumi.get(self, "sms_status")

    @sms_status.setter
    def sms_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sms_status", value)

    @property
    @pulumi.getter(name="ttsStatus")
    def tts_status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of tts subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        return pulumi.get(self, "tts_status")

    @tts_status.setter
    def tts_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tts_status", value)

    @property
    @pulumi.getter(name="webhookIds")
    def webhook_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The ids of subscribed webhooks.
        """
        return pulumi.get(self, "webhook_ids")

    @webhook_ids.setter
    def webhook_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "webhook_ids", value)

    @property
    @pulumi.getter(name="webhookStatus")
    def webhook_status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of webhook subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        return pulumi.get(self, "webhook_status")

    @webhook_status.setter
    def webhook_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "webhook_status", value)


@pulumi.input_type
class _MscSubSubscriptionState:
    def __init__(__self__, *,
                 channel: Optional[pulumi.Input[str]] = None,
                 contact_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 email_status: Optional[pulumi.Input[int]] = None,
                 item_name: Optional[pulumi.Input[str]] = None,
                 pmsg_status: Optional[pulumi.Input[int]] = None,
                 sms_status: Optional[pulumi.Input[int]] = None,
                 tts_status: Optional[pulumi.Input[int]] = None,
                 webhook_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 webhook_status: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering MscSubSubscription resources.
        :param pulumi.Input[str] channel: The channel the Subscription.
        :param pulumi.Input[str] description: The description of the Subscription.
        :param pulumi.Input[int] email_status: The status of email subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[str] item_name: The name of the Subscription. **NOTE:**  You should use the `get_msc_sub_subscriptions` to query the available subscription item name.
        :param pulumi.Input[int] pmsg_status: The status of pmsg subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[int] sms_status: The status of sms subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[int] tts_status: The status of tts subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] webhook_ids: The ids of subscribed webhooks.
        :param pulumi.Input[int] webhook_status: The status of webhook subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        if channel is not None:
            pulumi.set(__self__, "channel", channel)
        if contact_ids is not None:
            pulumi.set(__self__, "contact_ids", contact_ids)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if email_status is not None:
            pulumi.set(__self__, "email_status", email_status)
        if item_name is not None:
            pulumi.set(__self__, "item_name", item_name)
        if pmsg_status is not None:
            pulumi.set(__self__, "pmsg_status", pmsg_status)
        if sms_status is not None:
            pulumi.set(__self__, "sms_status", sms_status)
        if tts_status is not None:
            pulumi.set(__self__, "tts_status", tts_status)
        if webhook_ids is not None:
            pulumi.set(__self__, "webhook_ids", webhook_ids)
        if webhook_status is not None:
            pulumi.set(__self__, "webhook_status", webhook_status)

    @property
    @pulumi.getter
    def channel(self) -> Optional[pulumi.Input[str]]:
        """
        The channel the Subscription.
        """
        return pulumi.get(self, "channel")

    @channel.setter
    def channel(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "channel", value)

    @property
    @pulumi.getter(name="contactIds")
    def contact_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "contact_ids")

    @contact_ids.setter
    def contact_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "contact_ids", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the Subscription.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="emailStatus")
    def email_status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of email subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        return pulumi.get(self, "email_status")

    @email_status.setter
    def email_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "email_status", value)

    @property
    @pulumi.getter(name="itemName")
    def item_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Subscription. **NOTE:**  You should use the `get_msc_sub_subscriptions` to query the available subscription item name.
        """
        return pulumi.get(self, "item_name")

    @item_name.setter
    def item_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "item_name", value)

    @property
    @pulumi.getter(name="pmsgStatus")
    def pmsg_status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of pmsg subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        return pulumi.get(self, "pmsg_status")

    @pmsg_status.setter
    def pmsg_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pmsg_status", value)

    @property
    @pulumi.getter(name="smsStatus")
    def sms_status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of sms subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        return pulumi.get(self, "sms_status")

    @sms_status.setter
    def sms_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sms_status", value)

    @property
    @pulumi.getter(name="ttsStatus")
    def tts_status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of tts subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        return pulumi.get(self, "tts_status")

    @tts_status.setter
    def tts_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tts_status", value)

    @property
    @pulumi.getter(name="webhookIds")
    def webhook_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The ids of subscribed webhooks.
        """
        return pulumi.get(self, "webhook_ids")

    @webhook_ids.setter
    def webhook_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "webhook_ids", value)

    @property
    @pulumi.getter(name="webhookStatus")
    def webhook_status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of webhook subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        return pulumi.get(self, "webhook_status")

    @webhook_status.setter
    def webhook_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "webhook_status", value)


class MscSubSubscription(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 contact_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 email_status: Optional[pulumi.Input[int]] = None,
                 item_name: Optional[pulumi.Input[str]] = None,
                 pmsg_status: Optional[pulumi.Input[int]] = None,
                 sms_status: Optional[pulumi.Input[int]] = None,
                 tts_status: Optional[pulumi.Input[int]] = None,
                 webhook_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 webhook_status: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a Msc Sub Subscription resource.

        > **NOTE:** Available in v1.135.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.MscSubSubscription("example",
            email_status=1,
            item_name="Notifications of Product Expiration",
            pmsg_status=1,
            sms_status=1,
            tts_status=1,
            webhook_status=0)
        ```

        ## Import

        Msc Sub Subscription can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:index/mscSubSubscription:MscSubSubscription example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] email_status: The status of email subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[str] item_name: The name of the Subscription. **NOTE:**  You should use the `get_msc_sub_subscriptions` to query the available subscription item name.
        :param pulumi.Input[int] pmsg_status: The status of pmsg subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[int] sms_status: The status of sms subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[int] tts_status: The status of tts subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] webhook_ids: The ids of subscribed webhooks.
        :param pulumi.Input[int] webhook_status: The status of webhook subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MscSubSubscriptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Msc Sub Subscription resource.

        > **NOTE:** Available in v1.135.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.MscSubSubscription("example",
            email_status=1,
            item_name="Notifications of Product Expiration",
            pmsg_status=1,
            sms_status=1,
            tts_status=1,
            webhook_status=0)
        ```

        ## Import

        Msc Sub Subscription can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:index/mscSubSubscription:MscSubSubscription example <id>
        ```

        :param str resource_name: The name of the resource.
        :param MscSubSubscriptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MscSubSubscriptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 contact_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 email_status: Optional[pulumi.Input[int]] = None,
                 item_name: Optional[pulumi.Input[str]] = None,
                 pmsg_status: Optional[pulumi.Input[int]] = None,
                 sms_status: Optional[pulumi.Input[int]] = None,
                 tts_status: Optional[pulumi.Input[int]] = None,
                 webhook_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 webhook_status: Optional[pulumi.Input[int]] = None,
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
            __props__ = MscSubSubscriptionArgs.__new__(MscSubSubscriptionArgs)

            __props__.__dict__["contact_ids"] = contact_ids
            __props__.__dict__["email_status"] = email_status
            if item_name is None and not opts.urn:
                raise TypeError("Missing required property 'item_name'")
            __props__.__dict__["item_name"] = item_name
            __props__.__dict__["pmsg_status"] = pmsg_status
            __props__.__dict__["sms_status"] = sms_status
            __props__.__dict__["tts_status"] = tts_status
            __props__.__dict__["webhook_ids"] = webhook_ids
            __props__.__dict__["webhook_status"] = webhook_status
            __props__.__dict__["channel"] = None
            __props__.__dict__["description"] = None
        super(MscSubSubscription, __self__).__init__(
            'alicloud:index/mscSubSubscription:MscSubSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            channel: Optional[pulumi.Input[str]] = None,
            contact_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            email_status: Optional[pulumi.Input[int]] = None,
            item_name: Optional[pulumi.Input[str]] = None,
            pmsg_status: Optional[pulumi.Input[int]] = None,
            sms_status: Optional[pulumi.Input[int]] = None,
            tts_status: Optional[pulumi.Input[int]] = None,
            webhook_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            webhook_status: Optional[pulumi.Input[int]] = None) -> 'MscSubSubscription':
        """
        Get an existing MscSubSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] channel: The channel the Subscription.
        :param pulumi.Input[str] description: The description of the Subscription.
        :param pulumi.Input[int] email_status: The status of email subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[str] item_name: The name of the Subscription. **NOTE:**  You should use the `get_msc_sub_subscriptions` to query the available subscription item name.
        :param pulumi.Input[int] pmsg_status: The status of pmsg subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[int] sms_status: The status of sms subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[int] tts_status: The status of tts subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] webhook_ids: The ids of subscribed webhooks.
        :param pulumi.Input[int] webhook_status: The status of webhook subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MscSubSubscriptionState.__new__(_MscSubSubscriptionState)

        __props__.__dict__["channel"] = channel
        __props__.__dict__["contact_ids"] = contact_ids
        __props__.__dict__["description"] = description
        __props__.__dict__["email_status"] = email_status
        __props__.__dict__["item_name"] = item_name
        __props__.__dict__["pmsg_status"] = pmsg_status
        __props__.__dict__["sms_status"] = sms_status
        __props__.__dict__["tts_status"] = tts_status
        __props__.__dict__["webhook_ids"] = webhook_ids
        __props__.__dict__["webhook_status"] = webhook_status
        return MscSubSubscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def channel(self) -> pulumi.Output[str]:
        """
        The channel the Subscription.
        """
        return pulumi.get(self, "channel")

    @property
    @pulumi.getter(name="contactIds")
    def contact_ids(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "contact_ids")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the Subscription.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="emailStatus")
    def email_status(self) -> pulumi.Output[Optional[int]]:
        """
        The status of email subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        return pulumi.get(self, "email_status")

    @property
    @pulumi.getter(name="itemName")
    def item_name(self) -> pulumi.Output[str]:
        """
        The name of the Subscription. **NOTE:**  You should use the `get_msc_sub_subscriptions` to query the available subscription item name.
        """
        return pulumi.get(self, "item_name")

    @property
    @pulumi.getter(name="pmsgStatus")
    def pmsg_status(self) -> pulumi.Output[Optional[int]]:
        """
        The status of pmsg subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        return pulumi.get(self, "pmsg_status")

    @property
    @pulumi.getter(name="smsStatus")
    def sms_status(self) -> pulumi.Output[Optional[int]]:
        """
        The status of sms subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        return pulumi.get(self, "sms_status")

    @property
    @pulumi.getter(name="ttsStatus")
    def tts_status(self) -> pulumi.Output[Optional[int]]:
        """
        The status of tts subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        return pulumi.get(self, "tts_status")

    @property
    @pulumi.getter(name="webhookIds")
    def webhook_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The ids of subscribed webhooks.
        """
        return pulumi.get(self, "webhook_ids")

    @property
    @pulumi.getter(name="webhookStatus")
    def webhook_status(self) -> pulumi.Output[Optional[int]]:
        """
        The status of webhook subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        """
        return pulumi.get(self, "webhook_status")

