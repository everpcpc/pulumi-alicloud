# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ReadWriteSplittingConnection(pulumi.CustomResource):
    connection_prefix: pulumi.Output[str]
    """
    Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
    """
    connection_string: pulumi.Output[str]
    """
    Connection instance string.
    """
    distribution_type: pulumi.Output[str]
    """
    Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution. 
    """
    instance_id: pulumi.Output[str]
    """
    The Id of instance that can run database.
    """
    max_delay_time: pulumi.Output[float]
    """
    Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.  
    """
    port: pulumi.Output[float]
    """
    Intranet connection port. Valid value: [3001-3999]. Default to 3306.
    """
    weight: pulumi.Output[dict]
    """
    Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distribution_type is set to Custom. 
    """
    def __init__(__self__, resource_name, opts=None, connection_prefix=None, distribution_type=None, instance_id=None, max_delay_time=None, port=None, weight=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an RDS read write splitting connection resource to allocate an Intranet connection string for RDS instance.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_prefix: Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
        :param pulumi.Input[str] distribution_type: Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution. 
        :param pulumi.Input[str] instance_id: The Id of instance that can run database.
        :param pulumi.Input[float] max_delay_time: Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.  
        :param pulumi.Input[float] port: Intranet connection port. Valid value: [3001-3999]. Default to 3306.
        :param pulumi.Input[dict] weight: Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distribution_type is set to Custom. 
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

            __props__['connection_prefix'] = connection_prefix
            if distribution_type is None:
                raise TypeError("Missing required property 'distribution_type'")
            __props__['distribution_type'] = distribution_type
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            __props__['max_delay_time'] = max_delay_time
            __props__['port'] = port
            __props__['weight'] = weight
            __props__['connection_string'] = None
        super(ReadWriteSplittingConnection, __self__).__init__(
            'alicloud:rds/readWriteSplittingConnection:ReadWriteSplittingConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, connection_prefix=None, connection_string=None, distribution_type=None, instance_id=None, max_delay_time=None, port=None, weight=None):
        """
        Get an existing ReadWriteSplittingConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_prefix: Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
        :param pulumi.Input[str] connection_string: Connection instance string.
        :param pulumi.Input[str] distribution_type: Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution. 
        :param pulumi.Input[str] instance_id: The Id of instance that can run database.
        :param pulumi.Input[float] max_delay_time: Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.  
        :param pulumi.Input[float] port: Intranet connection port. Valid value: [3001-3999]. Default to 3306.
        :param pulumi.Input[dict] weight: Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distribution_type is set to Custom. 
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["connection_prefix"] = connection_prefix
        __props__["connection_string"] = connection_string
        __props__["distribution_type"] = distribution_type
        __props__["instance_id"] = instance_id
        __props__["max_delay_time"] = max_delay_time
        __props__["port"] = port
        __props__["weight"] = weight
        return ReadWriteSplittingConnection(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

