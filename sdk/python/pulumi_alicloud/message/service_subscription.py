# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ServiceSubscriptionArgs', 'ServiceSubscription']

@pulumi.input_type
class ServiceSubscriptionArgs:
    def __init__(__self__, *,
                 endpoint: pulumi.Input[str],
                 push_type: pulumi.Input[str],
                 subscription_name: pulumi.Input[str],
                 topic_name: pulumi.Input[str],
                 filter_tag: Optional[pulumi.Input[str]] = None,
                 notify_content_format: Optional[pulumi.Input[str]] = None,
                 notify_strategy: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceSubscription resource.
        :param pulumi.Input[str] endpoint: The endpoint has three format. Available values format:
               - `HTTP Format`: http://xxx.com/xxx
               - `Queue Format`: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
               - `Email Format`: mail:directmail:{MailAddress}
        :param pulumi.Input[str] push_type: The Push type of Subscription. The Valid values: `http`, `queue`, `mpush`, `alisms` and `email`.
        :param pulumi.Input[str] subscription_name: Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
        :param pulumi.Input[str] topic_name: The topic which The subscription belongs to was named with the name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
        :param pulumi.Input[str] filter_tag: The tag that is used to filter messages. Only the messages that have the same tag can be pushed. A tag is a string that can be up to 16 characters in length. By default, no tag is specified to filter messages.
        :param pulumi.Input[str] notify_content_format: The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. Valid values: `XML`, `JSON` and `SIMPLIFIED`. Default value: `XML`.
        :param pulumi.Input[str] notify_strategy: The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. Default value: `BACKOFF_RETRY`. Valid values:
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "push_type", push_type)
        pulumi.set(__self__, "subscription_name", subscription_name)
        pulumi.set(__self__, "topic_name", topic_name)
        if filter_tag is not None:
            pulumi.set(__self__, "filter_tag", filter_tag)
        if notify_content_format is not None:
            pulumi.set(__self__, "notify_content_format", notify_content_format)
        if notify_strategy is not None:
            pulumi.set(__self__, "notify_strategy", notify_strategy)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Input[str]:
        """
        The endpoint has three format. Available values format:
        - `HTTP Format`: http://xxx.com/xxx
        - `Queue Format`: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
        - `Email Format`: mail:directmail:{MailAddress}
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="pushType")
    def push_type(self) -> pulumi.Input[str]:
        """
        The Push type of Subscription. The Valid values: `http`, `queue`, `mpush`, `alisms` and `email`.
        """
        return pulumi.get(self, "push_type")

    @push_type.setter
    def push_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "push_type", value)

    @property
    @pulumi.getter(name="subscriptionName")
    def subscription_name(self) -> pulumi.Input[str]:
        """
        Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
        """
        return pulumi.get(self, "subscription_name")

    @subscription_name.setter
    def subscription_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "subscription_name", value)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Input[str]:
        """
        The topic which The subscription belongs to was named with the name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_name", value)

    @property
    @pulumi.getter(name="filterTag")
    def filter_tag(self) -> Optional[pulumi.Input[str]]:
        """
        The tag that is used to filter messages. Only the messages that have the same tag can be pushed. A tag is a string that can be up to 16 characters in length. By default, no tag is specified to filter messages.
        """
        return pulumi.get(self, "filter_tag")

    @filter_tag.setter
    def filter_tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_tag", value)

    @property
    @pulumi.getter(name="notifyContentFormat")
    def notify_content_format(self) -> Optional[pulumi.Input[str]]:
        """
        The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. Valid values: `XML`, `JSON` and `SIMPLIFIED`. Default value: `XML`.
        """
        return pulumi.get(self, "notify_content_format")

    @notify_content_format.setter
    def notify_content_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_content_format", value)

    @property
    @pulumi.getter(name="notifyStrategy")
    def notify_strategy(self) -> Optional[pulumi.Input[str]]:
        """
        The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. Default value: `BACKOFF_RETRY`. Valid values:
        """
        return pulumi.get(self, "notify_strategy")

    @notify_strategy.setter
    def notify_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_strategy", value)


@pulumi.input_type
class _ServiceSubscriptionState:
    def __init__(__self__, *,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 filter_tag: Optional[pulumi.Input[str]] = None,
                 notify_content_format: Optional[pulumi.Input[str]] = None,
                 notify_strategy: Optional[pulumi.Input[str]] = None,
                 push_type: Optional[pulumi.Input[str]] = None,
                 subscription_name: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceSubscription resources.
        :param pulumi.Input[str] endpoint: The endpoint has three format. Available values format:
               - `HTTP Format`: http://xxx.com/xxx
               - `Queue Format`: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
               - `Email Format`: mail:directmail:{MailAddress}
        :param pulumi.Input[str] filter_tag: The tag that is used to filter messages. Only the messages that have the same tag can be pushed. A tag is a string that can be up to 16 characters in length. By default, no tag is specified to filter messages.
        :param pulumi.Input[str] notify_content_format: The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. Valid values: `XML`, `JSON` and `SIMPLIFIED`. Default value: `XML`.
        :param pulumi.Input[str] notify_strategy: The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. Default value: `BACKOFF_RETRY`. Valid values:
        :param pulumi.Input[str] push_type: The Push type of Subscription. The Valid values: `http`, `queue`, `mpush`, `alisms` and `email`.
        :param pulumi.Input[str] subscription_name: Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
        :param pulumi.Input[str] topic_name: The topic which The subscription belongs to was named with the name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
        """
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if filter_tag is not None:
            pulumi.set(__self__, "filter_tag", filter_tag)
        if notify_content_format is not None:
            pulumi.set(__self__, "notify_content_format", notify_content_format)
        if notify_strategy is not None:
            pulumi.set(__self__, "notify_strategy", notify_strategy)
        if push_type is not None:
            pulumi.set(__self__, "push_type", push_type)
        if subscription_name is not None:
            pulumi.set(__self__, "subscription_name", subscription_name)
        if topic_name is not None:
            pulumi.set(__self__, "topic_name", topic_name)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint has three format. Available values format:
        - `HTTP Format`: http://xxx.com/xxx
        - `Queue Format`: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
        - `Email Format`: mail:directmail:{MailAddress}
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="filterTag")
    def filter_tag(self) -> Optional[pulumi.Input[str]]:
        """
        The tag that is used to filter messages. Only the messages that have the same tag can be pushed. A tag is a string that can be up to 16 characters in length. By default, no tag is specified to filter messages.
        """
        return pulumi.get(self, "filter_tag")

    @filter_tag.setter
    def filter_tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_tag", value)

    @property
    @pulumi.getter(name="notifyContentFormat")
    def notify_content_format(self) -> Optional[pulumi.Input[str]]:
        """
        The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. Valid values: `XML`, `JSON` and `SIMPLIFIED`. Default value: `XML`.
        """
        return pulumi.get(self, "notify_content_format")

    @notify_content_format.setter
    def notify_content_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_content_format", value)

    @property
    @pulumi.getter(name="notifyStrategy")
    def notify_strategy(self) -> Optional[pulumi.Input[str]]:
        """
        The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. Default value: `BACKOFF_RETRY`. Valid values:
        """
        return pulumi.get(self, "notify_strategy")

    @notify_strategy.setter
    def notify_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_strategy", value)

    @property
    @pulumi.getter(name="pushType")
    def push_type(self) -> Optional[pulumi.Input[str]]:
        """
        The Push type of Subscription. The Valid values: `http`, `queue`, `mpush`, `alisms` and `email`.
        """
        return pulumi.get(self, "push_type")

    @push_type.setter
    def push_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "push_type", value)

    @property
    @pulumi.getter(name="subscriptionName")
    def subscription_name(self) -> Optional[pulumi.Input[str]]:
        """
        Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
        """
        return pulumi.get(self, "subscription_name")

    @subscription_name.setter
    def subscription_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subscription_name", value)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> Optional[pulumi.Input[str]]:
        """
        The topic which The subscription belongs to was named with the name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_name", value)


class ServiceSubscription(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 filter_tag: Optional[pulumi.Input[str]] = None,
                 notify_content_format: Optional[pulumi.Input[str]] = None,
                 notify_strategy: Optional[pulumi.Input[str]] = None,
                 push_type: Optional[pulumi.Input[str]] = None,
                 subscription_name: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Message Notification Service Subscription resource.

        For information about Message Notification Service Subscription and how to use it, see [What is Subscription](https://www.alibabacloud.com/help/en/message-service/latest/subscribe-1).

        > **NOTE:** Available in v1.188.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_service_topic = alicloud.message.ServiceTopic("defaultServiceTopic",
            topic_name="tf-example-value",
            max_message_size=12357,
            logging_enabled=True)
        default_service_subscription = alicloud.message.ServiceSubscription("defaultServiceSubscription",
            topic_name=default_service_topic.topic_name,
            subscription_name="tf-example-value",
            endpoint="http://www.test.com/test",
            push_type="http",
            filter_tag="tf-test",
            notify_content_format="XML",
            notify_strategy="BACKOFF_RETRY")
        ```

        ## Import

        Message Notification Service Subscription can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:message/serviceSubscription:ServiceSubscription example <topic_name>:<subscription_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint: The endpoint has three format. Available values format:
               - `HTTP Format`: http://xxx.com/xxx
               - `Queue Format`: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
               - `Email Format`: mail:directmail:{MailAddress}
        :param pulumi.Input[str] filter_tag: The tag that is used to filter messages. Only the messages that have the same tag can be pushed. A tag is a string that can be up to 16 characters in length. By default, no tag is specified to filter messages.
        :param pulumi.Input[str] notify_content_format: The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. Valid values: `XML`, `JSON` and `SIMPLIFIED`. Default value: `XML`.
        :param pulumi.Input[str] notify_strategy: The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. Default value: `BACKOFF_RETRY`. Valid values:
        :param pulumi.Input[str] push_type: The Push type of Subscription. The Valid values: `http`, `queue`, `mpush`, `alisms` and `email`.
        :param pulumi.Input[str] subscription_name: Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
        :param pulumi.Input[str] topic_name: The topic which The subscription belongs to was named with the name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceSubscriptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Message Notification Service Subscription resource.

        For information about Message Notification Service Subscription and how to use it, see [What is Subscription](https://www.alibabacloud.com/help/en/message-service/latest/subscribe-1).

        > **NOTE:** Available in v1.188.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_service_topic = alicloud.message.ServiceTopic("defaultServiceTopic",
            topic_name="tf-example-value",
            max_message_size=12357,
            logging_enabled=True)
        default_service_subscription = alicloud.message.ServiceSubscription("defaultServiceSubscription",
            topic_name=default_service_topic.topic_name,
            subscription_name="tf-example-value",
            endpoint="http://www.test.com/test",
            push_type="http",
            filter_tag="tf-test",
            notify_content_format="XML",
            notify_strategy="BACKOFF_RETRY")
        ```

        ## Import

        Message Notification Service Subscription can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:message/serviceSubscription:ServiceSubscription example <topic_name>:<subscription_name>
        ```

        :param str resource_name: The name of the resource.
        :param ServiceSubscriptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceSubscriptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 filter_tag: Optional[pulumi.Input[str]] = None,
                 notify_content_format: Optional[pulumi.Input[str]] = None,
                 notify_strategy: Optional[pulumi.Input[str]] = None,
                 push_type: Optional[pulumi.Input[str]] = None,
                 subscription_name: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceSubscriptionArgs.__new__(ServiceSubscriptionArgs)

            if endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint'")
            __props__.__dict__["endpoint"] = endpoint
            __props__.__dict__["filter_tag"] = filter_tag
            __props__.__dict__["notify_content_format"] = notify_content_format
            __props__.__dict__["notify_strategy"] = notify_strategy
            if push_type is None and not opts.urn:
                raise TypeError("Missing required property 'push_type'")
            __props__.__dict__["push_type"] = push_type
            if subscription_name is None and not opts.urn:
                raise TypeError("Missing required property 'subscription_name'")
            __props__.__dict__["subscription_name"] = subscription_name
            if topic_name is None and not opts.urn:
                raise TypeError("Missing required property 'topic_name'")
            __props__.__dict__["topic_name"] = topic_name
        super(ServiceSubscription, __self__).__init__(
            'alicloud:message/serviceSubscription:ServiceSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            filter_tag: Optional[pulumi.Input[str]] = None,
            notify_content_format: Optional[pulumi.Input[str]] = None,
            notify_strategy: Optional[pulumi.Input[str]] = None,
            push_type: Optional[pulumi.Input[str]] = None,
            subscription_name: Optional[pulumi.Input[str]] = None,
            topic_name: Optional[pulumi.Input[str]] = None) -> 'ServiceSubscription':
        """
        Get an existing ServiceSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint: The endpoint has three format. Available values format:
               - `HTTP Format`: http://xxx.com/xxx
               - `Queue Format`: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
               - `Email Format`: mail:directmail:{MailAddress}
        :param pulumi.Input[str] filter_tag: The tag that is used to filter messages. Only the messages that have the same tag can be pushed. A tag is a string that can be up to 16 characters in length. By default, no tag is specified to filter messages.
        :param pulumi.Input[str] notify_content_format: The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. Valid values: `XML`, `JSON` and `SIMPLIFIED`. Default value: `XML`.
        :param pulumi.Input[str] notify_strategy: The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. Default value: `BACKOFF_RETRY`. Valid values:
        :param pulumi.Input[str] push_type: The Push type of Subscription. The Valid values: `http`, `queue`, `mpush`, `alisms` and `email`.
        :param pulumi.Input[str] subscription_name: Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
        :param pulumi.Input[str] topic_name: The topic which The subscription belongs to was named with the name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceSubscriptionState.__new__(_ServiceSubscriptionState)

        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["filter_tag"] = filter_tag
        __props__.__dict__["notify_content_format"] = notify_content_format
        __props__.__dict__["notify_strategy"] = notify_strategy
        __props__.__dict__["push_type"] = push_type
        __props__.__dict__["subscription_name"] = subscription_name
        __props__.__dict__["topic_name"] = topic_name
        return ServiceSubscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The endpoint has three format. Available values format:
        - `HTTP Format`: http://xxx.com/xxx
        - `Queue Format`: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
        - `Email Format`: mail:directmail:{MailAddress}
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="filterTag")
    def filter_tag(self) -> pulumi.Output[Optional[str]]:
        """
        The tag that is used to filter messages. Only the messages that have the same tag can be pushed. A tag is a string that can be up to 16 characters in length. By default, no tag is specified to filter messages.
        """
        return pulumi.get(self, "filter_tag")

    @property
    @pulumi.getter(name="notifyContentFormat")
    def notify_content_format(self) -> pulumi.Output[str]:
        """
        The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. Valid values: `XML`, `JSON` and `SIMPLIFIED`. Default value: `XML`.
        """
        return pulumi.get(self, "notify_content_format")

    @property
    @pulumi.getter(name="notifyStrategy")
    def notify_strategy(self) -> pulumi.Output[str]:
        """
        The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. Default value: `BACKOFF_RETRY`. Valid values:
        """
        return pulumi.get(self, "notify_strategy")

    @property
    @pulumi.getter(name="pushType")
    def push_type(self) -> pulumi.Output[str]:
        """
        The Push type of Subscription. The Valid values: `http`, `queue`, `mpush`, `alisms` and `email`.
        """
        return pulumi.get(self, "push_type")

    @property
    @pulumi.getter(name="subscriptionName")
    def subscription_name(self) -> pulumi.Output[str]:
        """
        Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
        """
        return pulumi.get(self, "subscription_name")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Output[str]:
        """
        The topic which The subscription belongs to was named with the name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
        """
        return pulumi.get(self, "topic_name")

