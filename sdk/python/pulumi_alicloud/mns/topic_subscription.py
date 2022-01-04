# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TopicSubscriptionArgs', 'TopicSubscription']

@pulumi.input_type
class TopicSubscriptionArgs:
    def __init__(__self__, *,
                 endpoint: pulumi.Input[str],
                 topic_name: pulumi.Input[str],
                 filter_tag: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_content_format: Optional[pulumi.Input[str]] = None,
                 notify_strategy: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TopicSubscription resource.
        :param pulumi.Input[str] endpoint: The endpoint has three format. Available values format:
               - `HTTP Format`: http://xxx.com/xxx
               - `Queue Format`: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
               - `Email Format`: mail:directmail:{MailAddress}
        :param pulumi.Input[str] topic_name: The topic which The subscription belongs to was named with the name.A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        :param pulumi.Input[str] filter_tag: The length should be shorter than 16.
        :param pulumi.Input[str] name: Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        :param pulumi.Input[str] notify_content_format: The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: `SIMPLIFIED`, `XML` and `JSON`. Default to `SIMPLIFIED`.
        :param pulumi.Input[str] notify_strategy: The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. The Valid values: `EXPONENTIAL_DECAY_RETRY` and `BACKOFF_RETRY`. Default value to `BACKOFF_RETRY` .
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "topic_name", topic_name)
        if filter_tag is not None:
            pulumi.set(__self__, "filter_tag", filter_tag)
        if name is not None:
            pulumi.set(__self__, "name", name)
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
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Input[str]:
        """
        The topic which The subscription belongs to was named with the name.A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_name", value)

    @property
    @pulumi.getter(name="filterTag")
    def filter_tag(self) -> Optional[pulumi.Input[str]]:
        """
        The length should be shorter than 16.
        """
        return pulumi.get(self, "filter_tag")

    @filter_tag.setter
    def filter_tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_tag", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notifyContentFormat")
    def notify_content_format(self) -> Optional[pulumi.Input[str]]:
        """
        The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: `SIMPLIFIED`, `XML` and `JSON`. Default to `SIMPLIFIED`.
        """
        return pulumi.get(self, "notify_content_format")

    @notify_content_format.setter
    def notify_content_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_content_format", value)

    @property
    @pulumi.getter(name="notifyStrategy")
    def notify_strategy(self) -> Optional[pulumi.Input[str]]:
        """
        The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. The Valid values: `EXPONENTIAL_DECAY_RETRY` and `BACKOFF_RETRY`. Default value to `BACKOFF_RETRY` .
        """
        return pulumi.get(self, "notify_strategy")

    @notify_strategy.setter
    def notify_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_strategy", value)


@pulumi.input_type
class _TopicSubscriptionState:
    def __init__(__self__, *,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 filter_tag: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_content_format: Optional[pulumi.Input[str]] = None,
                 notify_strategy: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TopicSubscription resources.
        :param pulumi.Input[str] endpoint: The endpoint has three format. Available values format:
               - `HTTP Format`: http://xxx.com/xxx
               - `Queue Format`: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
               - `Email Format`: mail:directmail:{MailAddress}
        :param pulumi.Input[str] filter_tag: The length should be shorter than 16.
        :param pulumi.Input[str] name: Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        :param pulumi.Input[str] notify_content_format: The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: `SIMPLIFIED`, `XML` and `JSON`. Default to `SIMPLIFIED`.
        :param pulumi.Input[str] notify_strategy: The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. The Valid values: `EXPONENTIAL_DECAY_RETRY` and `BACKOFF_RETRY`. Default value to `BACKOFF_RETRY` .
        :param pulumi.Input[str] topic_name: The topic which The subscription belongs to was named with the name.A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        """
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if filter_tag is not None:
            pulumi.set(__self__, "filter_tag", filter_tag)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notify_content_format is not None:
            pulumi.set(__self__, "notify_content_format", notify_content_format)
        if notify_strategy is not None:
            pulumi.set(__self__, "notify_strategy", notify_strategy)
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
        The length should be shorter than 16.
        """
        return pulumi.get(self, "filter_tag")

    @filter_tag.setter
    def filter_tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_tag", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notifyContentFormat")
    def notify_content_format(self) -> Optional[pulumi.Input[str]]:
        """
        The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: `SIMPLIFIED`, `XML` and `JSON`. Default to `SIMPLIFIED`.
        """
        return pulumi.get(self, "notify_content_format")

    @notify_content_format.setter
    def notify_content_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_content_format", value)

    @property
    @pulumi.getter(name="notifyStrategy")
    def notify_strategy(self) -> Optional[pulumi.Input[str]]:
        """
        The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. The Valid values: `EXPONENTIAL_DECAY_RETRY` and `BACKOFF_RETRY`. Default value to `BACKOFF_RETRY` .
        """
        return pulumi.get(self, "notify_strategy")

    @notify_strategy.setter
    def notify_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_strategy", value)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> Optional[pulumi.Input[str]]:
        """
        The topic which The subscription belongs to was named with the name.A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_name", value)


class TopicSubscription(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 filter_tag: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_content_format: Optional[pulumi.Input[str]] = None,
                 notify_strategy: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        MNS Topic subscription can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:mns/topicSubscription:TopicSubscription subscription tf-example-mnstopic:tf-example-mnstopic-sub
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint: The endpoint has three format. Available values format:
               - `HTTP Format`: http://xxx.com/xxx
               - `Queue Format`: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
               - `Email Format`: mail:directmail:{MailAddress}
        :param pulumi.Input[str] filter_tag: The length should be shorter than 16.
        :param pulumi.Input[str] name: Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        :param pulumi.Input[str] notify_content_format: The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: `SIMPLIFIED`, `XML` and `JSON`. Default to `SIMPLIFIED`.
        :param pulumi.Input[str] notify_strategy: The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. The Valid values: `EXPONENTIAL_DECAY_RETRY` and `BACKOFF_RETRY`. Default value to `BACKOFF_RETRY` .
        :param pulumi.Input[str] topic_name: The topic which The subscription belongs to was named with the name.A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TopicSubscriptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        MNS Topic subscription can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:mns/topicSubscription:TopicSubscription subscription tf-example-mnstopic:tf-example-mnstopic-sub
        ```

        :param str resource_name: The name of the resource.
        :param TopicSubscriptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TopicSubscriptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 filter_tag: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_content_format: Optional[pulumi.Input[str]] = None,
                 notify_strategy: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = TopicSubscriptionArgs.__new__(TopicSubscriptionArgs)

            if endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint'")
            __props__.__dict__["endpoint"] = endpoint
            __props__.__dict__["filter_tag"] = filter_tag
            __props__.__dict__["name"] = name
            __props__.__dict__["notify_content_format"] = notify_content_format
            __props__.__dict__["notify_strategy"] = notify_strategy
            if topic_name is None and not opts.urn:
                raise TypeError("Missing required property 'topic_name'")
            __props__.__dict__["topic_name"] = topic_name
        super(TopicSubscription, __self__).__init__(
            'alicloud:mns/topicSubscription:TopicSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            filter_tag: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notify_content_format: Optional[pulumi.Input[str]] = None,
            notify_strategy: Optional[pulumi.Input[str]] = None,
            topic_name: Optional[pulumi.Input[str]] = None) -> 'TopicSubscription':
        """
        Get an existing TopicSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint: The endpoint has three format. Available values format:
               - `HTTP Format`: http://xxx.com/xxx
               - `Queue Format`: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
               - `Email Format`: mail:directmail:{MailAddress}
        :param pulumi.Input[str] filter_tag: The length should be shorter than 16.
        :param pulumi.Input[str] name: Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        :param pulumi.Input[str] notify_content_format: The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: `SIMPLIFIED`, `XML` and `JSON`. Default to `SIMPLIFIED`.
        :param pulumi.Input[str] notify_strategy: The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. The Valid values: `EXPONENTIAL_DECAY_RETRY` and `BACKOFF_RETRY`. Default value to `BACKOFF_RETRY` .
        :param pulumi.Input[str] topic_name: The topic which The subscription belongs to was named with the name.A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TopicSubscriptionState.__new__(_TopicSubscriptionState)

        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["filter_tag"] = filter_tag
        __props__.__dict__["name"] = name
        __props__.__dict__["notify_content_format"] = notify_content_format
        __props__.__dict__["notify_strategy"] = notify_strategy
        __props__.__dict__["topic_name"] = topic_name
        return TopicSubscription(resource_name, opts=opts, __props__=__props__)

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
        The length should be shorter than 16.
        """
        return pulumi.get(self, "filter_tag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notifyContentFormat")
    def notify_content_format(self) -> pulumi.Output[Optional[str]]:
        """
        The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: `SIMPLIFIED`, `XML` and `JSON`. Default to `SIMPLIFIED`.
        """
        return pulumi.get(self, "notify_content_format")

    @property
    @pulumi.getter(name="notifyStrategy")
    def notify_strategy(self) -> pulumi.Output[Optional[str]]:
        """
        The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. The Valid values: `EXPONENTIAL_DECAY_RETRY` and `BACKOFF_RETRY`. Default value to `BACKOFF_RETRY` .
        """
        return pulumi.get(self, "notify_strategy")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Output[str]:
        """
        The topic which The subscription belongs to was named with the name.A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        """
        return pulumi.get(self, "topic_name")

