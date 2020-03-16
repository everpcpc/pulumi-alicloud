# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SiteMonitor(pulumi.CustomResource):
    address: pulumi.Output[str]
    alert_ids: pulumi.Output[list]
    create_time: pulumi.Output[str]
    interval: pulumi.Output[float]
    options_json: pulumi.Output[str]
    task_name: pulumi.Output[str]
    task_state: pulumi.Output[str]
    task_type: pulumi.Output[str]
    update_time: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, address=None, alert_ids=None, interval=None, options_json=None, task_name=None, task_type=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a SiteMonitor resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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

            if address is None:
                raise TypeError("Missing required property 'address'")
            __props__['address'] = address
            __props__['alert_ids'] = alert_ids
            __props__['interval'] = interval
            __props__['options_json'] = options_json
            if task_name is None:
                raise TypeError("Missing required property 'task_name'")
            __props__['task_name'] = task_name
            if task_type is None:
                raise TypeError("Missing required property 'task_type'")
            __props__['task_type'] = task_type
            __props__['create_time'] = None
            __props__['task_state'] = None
            __props__['update_time'] = None
        super(SiteMonitor, __self__).__init__(
            'alicloud:cms/siteMonitor:SiteMonitor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, address=None, alert_ids=None, create_time=None, interval=None, options_json=None, task_name=None, task_state=None, task_type=None, update_time=None):
        """
        Get an existing SiteMonitor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address"] = address
        __props__["alert_ids"] = alert_ids
        __props__["create_time"] = create_time
        __props__["interval"] = interval
        __props__["options_json"] = options_json
        __props__["task_name"] = task_name
        __props__["task_state"] = task_state
        __props__["task_type"] = task_type
        __props__["update_time"] = update_time
        return SiteMonitor(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

