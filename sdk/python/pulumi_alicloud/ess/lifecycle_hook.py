# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class LifecycleHook(pulumi.CustomResource):
    default_result: pulumi.Output[str]
    """
    Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, default value: CONTINUE.
    """
    heartbeat_timeout: pulumi.Output[float]
    """
    Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter. Default value: 600.
    """
    lifecycle_transition: pulumi.Output[str]
    """
    Type of Scaling activity attached to lifecycle hook. Supported value: SCALE_OUT, SCALE_IN.
    """
    name: pulumi.Output[str]
    """
    The name of the lifecycle hook, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is lifecycle hook id.
    """
    notification_arn: pulumi.Output[str]
    """
    The Arn of notification target.
    """
    notification_metadata: pulumi.Output[str]
    """
    Additional information that you want to include when Auto Scaling sends a message to the notification target.
    """
    scaling_group_id: pulumi.Output[str]
    """
    The ID of the Auto Scaling group to which you want to assign the lifecycle hook.
    """
    def __init__(__self__, resource_name, opts=None, default_result=None, heartbeat_timeout=None, lifecycle_transition=None, name=None, notification_arn=None, notification_metadata=None, scaling_group_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a LifecycleHook resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_result: Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, default value: CONTINUE.
        :param pulumi.Input[float] heartbeat_timeout: Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter. Default value: 600.
        :param pulumi.Input[str] lifecycle_transition: Type of Scaling activity attached to lifecycle hook. Supported value: SCALE_OUT, SCALE_IN.
        :param pulumi.Input[str] name: The name of the lifecycle hook, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is lifecycle hook id.
        :param pulumi.Input[str] notification_arn: The Arn of notification target.
        :param pulumi.Input[str] notification_metadata: Additional information that you want to include when Auto Scaling sends a message to the notification target.
        :param pulumi.Input[str] scaling_group_id: The ID of the Auto Scaling group to which you want to assign the lifecycle hook.
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

            __props__['default_result'] = default_result
            __props__['heartbeat_timeout'] = heartbeat_timeout
            if lifecycle_transition is None:
                raise TypeError("Missing required property 'lifecycle_transition'")
            __props__['lifecycle_transition'] = lifecycle_transition
            __props__['name'] = name
            __props__['notification_arn'] = notification_arn
            __props__['notification_metadata'] = notification_metadata
            if scaling_group_id is None:
                raise TypeError("Missing required property 'scaling_group_id'")
            __props__['scaling_group_id'] = scaling_group_id
        super(LifecycleHook, __self__).__init__(
            'alicloud:ess/lifecycleHook:LifecycleHook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, default_result=None, heartbeat_timeout=None, lifecycle_transition=None, name=None, notification_arn=None, notification_metadata=None, scaling_group_id=None):
        """
        Get an existing LifecycleHook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_result: Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, default value: CONTINUE.
        :param pulumi.Input[float] heartbeat_timeout: Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter. Default value: 600.
        :param pulumi.Input[str] lifecycle_transition: Type of Scaling activity attached to lifecycle hook. Supported value: SCALE_OUT, SCALE_IN.
        :param pulumi.Input[str] name: The name of the lifecycle hook, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is lifecycle hook id.
        :param pulumi.Input[str] notification_arn: The Arn of notification target.
        :param pulumi.Input[str] notification_metadata: Additional information that you want to include when Auto Scaling sends a message to the notification target.
        :param pulumi.Input[str] scaling_group_id: The ID of the Auto Scaling group to which you want to assign the lifecycle hook.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["default_result"] = default_result
        __props__["heartbeat_timeout"] = heartbeat_timeout
        __props__["lifecycle_transition"] = lifecycle_transition
        __props__["name"] = name
        __props__["notification_arn"] = notification_arn
        __props__["notification_metadata"] = notification_metadata
        __props__["scaling_group_id"] = scaling_group_id
        return LifecycleHook(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

