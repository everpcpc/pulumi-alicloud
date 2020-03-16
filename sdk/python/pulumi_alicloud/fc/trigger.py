# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Trigger(pulumi.CustomResource):
    config: pulumi.Output[str]
    """
    The config of Function Compute trigger.It is valid when `type` is not "mns_topic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
    """
    config_mns: pulumi.Output[str]
    """
    The config of Function Compute trigger when the type is "mns_topic".It is conflict with `config`.
    """
    function: pulumi.Output[str]
    """
    The Function Compute function name.
    """
    last_modified: pulumi.Output[str]
    """
    The date this resource was last modified.
    """
    name: pulumi.Output[str]
    """
    The Function Compute trigger name. It is the only in one service and is conflict with "name_prefix".
    """
    name_prefix: pulumi.Output[str]
    """
    Setting a prefix to get a only trigger name. It is conflict with "name".
    """
    role: pulumi.Output[str]
    """
    RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
    """
    service: pulumi.Output[str]
    """
    The Function Compute service name.
    """
    source_arn: pulumi.Output[str]
    """
    Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
    """
    trigger_id: pulumi.Output[str]
    """
    The Function Compute trigger ID.
    """
    type: pulumi.Output[str]
    """
    The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mns_topic", "cdn_events"].
    """
    def __init__(__self__, resource_name, opts=None, config=None, config_mns=None, function=None, name=None, name_prefix=None, role=None, service=None, source_arn=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Trigger resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: The config of Function Compute trigger.It is valid when `type` is not "mns_topic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
        :param pulumi.Input[str] config_mns: The config of Function Compute trigger when the type is "mns_topic".It is conflict with `config`.
        :param pulumi.Input[str] function: The Function Compute function name.
        :param pulumi.Input[str] name: The Function Compute trigger name. It is the only in one service and is conflict with "name_prefix".
        :param pulumi.Input[str] name_prefix: Setting a prefix to get a only trigger name. It is conflict with "name".
        :param pulumi.Input[str] role: RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        :param pulumi.Input[str] service: The Function Compute service name.
        :param pulumi.Input[str] source_arn: Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        :param pulumi.Input[str] type: The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mns_topic", "cdn_events"].
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['config'] = config
            __props__['config_mns'] = config_mns
            if function is None:
                raise TypeError("Missing required property 'function'")
            __props__['function'] = function
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            __props__['role'] = role
            if service is None:
                raise TypeError("Missing required property 'service'")
            __props__['service'] = service
            __props__['source_arn'] = source_arn
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['last_modified'] = None
            __props__['trigger_id'] = None
        super(Trigger, __self__).__init__(
            'alicloud:fc/trigger:Trigger',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, config=None, config_mns=None, function=None, last_modified=None, name=None, name_prefix=None, role=None, service=None, source_arn=None, trigger_id=None, type=None):
        """
        Get an existing Trigger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: The config of Function Compute trigger.It is valid when `type` is not "mns_topic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
        :param pulumi.Input[str] config_mns: The config of Function Compute trigger when the type is "mns_topic".It is conflict with `config`.
        :param pulumi.Input[str] function: The Function Compute function name.
        :param pulumi.Input[str] last_modified: The date this resource was last modified.
        :param pulumi.Input[str] name: The Function Compute trigger name. It is the only in one service and is conflict with "name_prefix".
        :param pulumi.Input[str] name_prefix: Setting a prefix to get a only trigger name. It is conflict with "name".
        :param pulumi.Input[str] role: RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        :param pulumi.Input[str] service: The Function Compute service name.
        :param pulumi.Input[str] source_arn: Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        :param pulumi.Input[str] trigger_id: The Function Compute trigger ID.
        :param pulumi.Input[str] type: The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mns_topic", "cdn_events"].
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["config"] = config
        __props__["config_mns"] = config_mns
        __props__["function"] = function
        __props__["last_modified"] = last_modified
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["role"] = role
        __props__["service"] = service
        __props__["source_arn"] = source_arn
        __props__["trigger_id"] = trigger_id
        __props__["type"] = type
        return Trigger(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

