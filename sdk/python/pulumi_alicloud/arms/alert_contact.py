# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AlertContactArgs', 'AlertContact']

@pulumi.input_type
class AlertContactArgs:
    def __init__(__self__, *,
                 alert_contact_name: Optional[pulumi.Input[str]] = None,
                 ding_robot_webhook_url: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 phone_num: Optional[pulumi.Input[str]] = None,
                 system_noc: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AlertContact resource.
        :param pulumi.Input[str] alert_contact_name: The name of the alert contact.
        :param pulumi.Input[str] ding_robot_webhook_url: The webhook URL of the DingTalk chatbot. For more information about how to obtain the URL, see Configure a DingTalk chatbot to send alert notifications: https://www.alibabacloud.com/help/en/doc-detail/106247.htm. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        :param pulumi.Input[str] email: The email address of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        :param pulumi.Input[str] phone_num: The mobile number of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        :param pulumi.Input[bool] system_noc: Specifies whether the alert contact receives system notifications. Valid values:  true: receives system notifications. false: does not receive system notifications.
        """
        if alert_contact_name is not None:
            pulumi.set(__self__, "alert_contact_name", alert_contact_name)
        if ding_robot_webhook_url is not None:
            pulumi.set(__self__, "ding_robot_webhook_url", ding_robot_webhook_url)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if phone_num is not None:
            pulumi.set(__self__, "phone_num", phone_num)
        if system_noc is not None:
            pulumi.set(__self__, "system_noc", system_noc)

    @property
    @pulumi.getter(name="alertContactName")
    def alert_contact_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the alert contact.
        """
        return pulumi.get(self, "alert_contact_name")

    @alert_contact_name.setter
    def alert_contact_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alert_contact_name", value)

    @property
    @pulumi.getter(name="dingRobotWebhookUrl")
    def ding_robot_webhook_url(self) -> Optional[pulumi.Input[str]]:
        """
        The webhook URL of the DingTalk chatbot. For more information about how to obtain the URL, see Configure a DingTalk chatbot to send alert notifications: https://www.alibabacloud.com/help/en/doc-detail/106247.htm. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        """
        return pulumi.get(self, "ding_robot_webhook_url")

    @ding_robot_webhook_url.setter
    def ding_robot_webhook_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ding_robot_webhook_url", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The email address of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="phoneNum")
    def phone_num(self) -> Optional[pulumi.Input[str]]:
        """
        The mobile number of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        """
        return pulumi.get(self, "phone_num")

    @phone_num.setter
    def phone_num(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone_num", value)

    @property
    @pulumi.getter(name="systemNoc")
    def system_noc(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the alert contact receives system notifications. Valid values:  true: receives system notifications. false: does not receive system notifications.
        """
        return pulumi.get(self, "system_noc")

    @system_noc.setter
    def system_noc(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "system_noc", value)


@pulumi.input_type
class _AlertContactState:
    def __init__(__self__, *,
                 alert_contact_name: Optional[pulumi.Input[str]] = None,
                 ding_robot_webhook_url: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 phone_num: Optional[pulumi.Input[str]] = None,
                 system_noc: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering AlertContact resources.
        :param pulumi.Input[str] alert_contact_name: The name of the alert contact.
        :param pulumi.Input[str] ding_robot_webhook_url: The webhook URL of the DingTalk chatbot. For more information about how to obtain the URL, see Configure a DingTalk chatbot to send alert notifications: https://www.alibabacloud.com/help/en/doc-detail/106247.htm. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        :param pulumi.Input[str] email: The email address of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        :param pulumi.Input[str] phone_num: The mobile number of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        :param pulumi.Input[bool] system_noc: Specifies whether the alert contact receives system notifications. Valid values:  true: receives system notifications. false: does not receive system notifications.
        """
        if alert_contact_name is not None:
            pulumi.set(__self__, "alert_contact_name", alert_contact_name)
        if ding_robot_webhook_url is not None:
            pulumi.set(__self__, "ding_robot_webhook_url", ding_robot_webhook_url)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if phone_num is not None:
            pulumi.set(__self__, "phone_num", phone_num)
        if system_noc is not None:
            pulumi.set(__self__, "system_noc", system_noc)

    @property
    @pulumi.getter(name="alertContactName")
    def alert_contact_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the alert contact.
        """
        return pulumi.get(self, "alert_contact_name")

    @alert_contact_name.setter
    def alert_contact_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alert_contact_name", value)

    @property
    @pulumi.getter(name="dingRobotWebhookUrl")
    def ding_robot_webhook_url(self) -> Optional[pulumi.Input[str]]:
        """
        The webhook URL of the DingTalk chatbot. For more information about how to obtain the URL, see Configure a DingTalk chatbot to send alert notifications: https://www.alibabacloud.com/help/en/doc-detail/106247.htm. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        """
        return pulumi.get(self, "ding_robot_webhook_url")

    @ding_robot_webhook_url.setter
    def ding_robot_webhook_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ding_robot_webhook_url", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The email address of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="phoneNum")
    def phone_num(self) -> Optional[pulumi.Input[str]]:
        """
        The mobile number of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        """
        return pulumi.get(self, "phone_num")

    @phone_num.setter
    def phone_num(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone_num", value)

    @property
    @pulumi.getter(name="systemNoc")
    def system_noc(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the alert contact receives system notifications. Valid values:  true: receives system notifications. false: does not receive system notifications.
        """
        return pulumi.get(self, "system_noc")

    @system_noc.setter
    def system_noc(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "system_noc", value)


class AlertContact(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_contact_name: Optional[pulumi.Input[str]] = None,
                 ding_robot_webhook_url: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 phone_num: Optional[pulumi.Input[str]] = None,
                 system_noc: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Provides a Application Real-Time Monitoring Service (ARMS) Alert Contact resource.

        For information about Application Real-Time Monitoring Service (ARMS) Alert Contact and how to use it, see [What is Alert Contact](https://www.alibabacloud.com/help/en/doc-detail/42953.htm).

        > **NOTE:** Available in v1.129.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.arms.AlertContact("example",
            alert_contact_name="example_value",
            ding_robot_webhook_url="https://oapi.dingtalk.com/robot/send?access_token=91f2f6****",
            email="someone@example.com",
            phone_num="1381111****")
        ```

        ## Import

        Application Real-Time Monitoring Service (ARMS) Alert Contact can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:arms/alertContact:AlertContact example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alert_contact_name: The name of the alert contact.
        :param pulumi.Input[str] ding_robot_webhook_url: The webhook URL of the DingTalk chatbot. For more information about how to obtain the URL, see Configure a DingTalk chatbot to send alert notifications: https://www.alibabacloud.com/help/en/doc-detail/106247.htm. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        :param pulumi.Input[str] email: The email address of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        :param pulumi.Input[str] phone_num: The mobile number of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        :param pulumi.Input[bool] system_noc: Specifies whether the alert contact receives system notifications. Valid values:  true: receives system notifications. false: does not receive system notifications.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AlertContactArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Application Real-Time Monitoring Service (ARMS) Alert Contact resource.

        For information about Application Real-Time Monitoring Service (ARMS) Alert Contact and how to use it, see [What is Alert Contact](https://www.alibabacloud.com/help/en/doc-detail/42953.htm).

        > **NOTE:** Available in v1.129.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.arms.AlertContact("example",
            alert_contact_name="example_value",
            ding_robot_webhook_url="https://oapi.dingtalk.com/robot/send?access_token=91f2f6****",
            email="someone@example.com",
            phone_num="1381111****")
        ```

        ## Import

        Application Real-Time Monitoring Service (ARMS) Alert Contact can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:arms/alertContact:AlertContact example <id>
        ```

        :param str resource_name: The name of the resource.
        :param AlertContactArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AlertContactArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_contact_name: Optional[pulumi.Input[str]] = None,
                 ding_robot_webhook_url: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 phone_num: Optional[pulumi.Input[str]] = None,
                 system_noc: Optional[pulumi.Input[bool]] = None,
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
            __props__ = AlertContactArgs.__new__(AlertContactArgs)

            __props__.__dict__["alert_contact_name"] = alert_contact_name
            __props__.__dict__["ding_robot_webhook_url"] = ding_robot_webhook_url
            __props__.__dict__["email"] = email
            __props__.__dict__["phone_num"] = phone_num
            __props__.__dict__["system_noc"] = system_noc
        super(AlertContact, __self__).__init__(
            'alicloud:arms/alertContact:AlertContact',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alert_contact_name: Optional[pulumi.Input[str]] = None,
            ding_robot_webhook_url: Optional[pulumi.Input[str]] = None,
            email: Optional[pulumi.Input[str]] = None,
            phone_num: Optional[pulumi.Input[str]] = None,
            system_noc: Optional[pulumi.Input[bool]] = None) -> 'AlertContact':
        """
        Get an existing AlertContact resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alert_contact_name: The name of the alert contact.
        :param pulumi.Input[str] ding_robot_webhook_url: The webhook URL of the DingTalk chatbot. For more information about how to obtain the URL, see Configure a DingTalk chatbot to send alert notifications: https://www.alibabacloud.com/help/en/doc-detail/106247.htm. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        :param pulumi.Input[str] email: The email address of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        :param pulumi.Input[str] phone_num: The mobile number of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        :param pulumi.Input[bool] system_noc: Specifies whether the alert contact receives system notifications. Valid values:  true: receives system notifications. false: does not receive system notifications.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AlertContactState.__new__(_AlertContactState)

        __props__.__dict__["alert_contact_name"] = alert_contact_name
        __props__.__dict__["ding_robot_webhook_url"] = ding_robot_webhook_url
        __props__.__dict__["email"] = email
        __props__.__dict__["phone_num"] = phone_num
        __props__.__dict__["system_noc"] = system_noc
        return AlertContact(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alertContactName")
    def alert_contact_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the alert contact.
        """
        return pulumi.get(self, "alert_contact_name")

    @property
    @pulumi.getter(name="dingRobotWebhookUrl")
    def ding_robot_webhook_url(self) -> pulumi.Output[Optional[str]]:
        """
        The webhook URL of the DingTalk chatbot. For more information about how to obtain the URL, see Configure a DingTalk chatbot to send alert notifications: https://www.alibabacloud.com/help/en/doc-detail/106247.htm. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        """
        return pulumi.get(self, "ding_robot_webhook_url")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[Optional[str]]:
        """
        The email address of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="phoneNum")
    def phone_num(self) -> pulumi.Output[Optional[str]]:
        """
        The mobile number of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
        """
        return pulumi.get(self, "phone_num")

    @property
    @pulumi.getter(name="systemNoc")
    def system_noc(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether the alert contact receives system notifications. Valid values:  true: receives system notifications. false: does not receive system notifications.
        """
        return pulumi.get(self, "system_noc")

