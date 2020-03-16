# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Store(pulumi.CustomResource):
    append_meta: pulumi.Output[bool]
    """
    Determines whether to append log meta automatically. The meta includes log receive time and client IP address. Default to true.
    """
    auto_split: pulumi.Output[bool]
    """
    Determines whether to automatically split a shard. Default to true.
    """
    enable_web_tracking: pulumi.Output[bool]
    """
    Determines whether to enable Web Tracking. Default false.
    """
    max_split_shard_count: pulumi.Output[float]
    """
    The maximum number of shards for automatic split, which is in the range of 1 to 64. You must specify this parameter when autoSplit is true.
    """
    name: pulumi.Output[str]
    """
    The log store, which is unique in the same project.
    """
    project: pulumi.Output[str]
    """
    The project name to the log store belongs.
    """
    retention_period: pulumi.Output[float]
    """
    The data retention time (in days). Valid values: [1-3650]. Default to 30. Log store data will be stored permanently when the value is "3650".
    """
    shard_count: pulumi.Output[float]
    """
    The number of shards in this log store. Default to 2. You can modify it by "Split" or "Merge" operations. [Refer to details](https://www.alibabacloud.com/help/doc-detail/28976.htm)
    """
    shards: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, append_meta=None, auto_split=None, enable_web_tracking=None, max_split_shard_count=None, name=None, project=None, retention_period=None, shard_count=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Store resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] append_meta: Determines whether to append log meta automatically. The meta includes log receive time and client IP address. Default to true.
        :param pulumi.Input[bool] auto_split: Determines whether to automatically split a shard. Default to true.
        :param pulumi.Input[bool] enable_web_tracking: Determines whether to enable Web Tracking. Default false.
        :param pulumi.Input[float] max_split_shard_count: The maximum number of shards for automatic split, which is in the range of 1 to 64. You must specify this parameter when autoSplit is true.
        :param pulumi.Input[str] name: The log store, which is unique in the same project.
        :param pulumi.Input[str] project: The project name to the log store belongs.
        :param pulumi.Input[float] retention_period: The data retention time (in days). Valid values: [1-3650]. Default to 30. Log store data will be stored permanently when the value is "3650".
        :param pulumi.Input[float] shard_count: The number of shards in this log store. Default to 2. You can modify it by "Split" or "Merge" operations. [Refer to details](https://www.alibabacloud.com/help/doc-detail/28976.htm)
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

            __props__['append_meta'] = append_meta
            __props__['auto_split'] = auto_split
            __props__['enable_web_tracking'] = enable_web_tracking
            __props__['max_split_shard_count'] = max_split_shard_count
            __props__['name'] = name
            if project is None:
                raise TypeError("Missing required property 'project'")
            __props__['project'] = project
            __props__['retention_period'] = retention_period
            __props__['shard_count'] = shard_count
            __props__['shards'] = None
        super(Store, __self__).__init__(
            'alicloud:log/store:Store',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, append_meta=None, auto_split=None, enable_web_tracking=None, max_split_shard_count=None, name=None, project=None, retention_period=None, shard_count=None, shards=None):
        """
        Get an existing Store resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] append_meta: Determines whether to append log meta automatically. The meta includes log receive time and client IP address. Default to true.
        :param pulumi.Input[bool] auto_split: Determines whether to automatically split a shard. Default to true.
        :param pulumi.Input[bool] enable_web_tracking: Determines whether to enable Web Tracking. Default false.
        :param pulumi.Input[float] max_split_shard_count: The maximum number of shards for automatic split, which is in the range of 1 to 64. You must specify this parameter when autoSplit is true.
        :param pulumi.Input[str] name: The log store, which is unique in the same project.
        :param pulumi.Input[str] project: The project name to the log store belongs.
        :param pulumi.Input[float] retention_period: The data retention time (in days). Valid values: [1-3650]. Default to 30. Log store data will be stored permanently when the value is "3650".
        :param pulumi.Input[float] shard_count: The number of shards in this log store. Default to 2. You can modify it by "Split" or "Merge" operations. [Refer to details](https://www.alibabacloud.com/help/doc-detail/28976.htm)

        The **shards** object supports the following:

          * `beginKey` (`pulumi.Input[str]`)
          * `endKey` (`pulumi.Input[str]`)
          * `id` (`pulumi.Input[float]`) - The ID of the log project. It formats of `<project>:<name>`.
          * `status` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["append_meta"] = append_meta
        __props__["auto_split"] = auto_split
        __props__["enable_web_tracking"] = enable_web_tracking
        __props__["max_split_shard_count"] = max_split_shard_count
        __props__["name"] = name
        __props__["project"] = project
        __props__["retention_period"] = retention_period
        __props__["shard_count"] = shard_count
        __props__["shards"] = shards
        return Store(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

