# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SaslAcl(pulumi.CustomResource):
    acl_operation_type: pulumi.Output[str]
    """
    Operation type for this acl. The operation type can only be "Write" and "Read".
    """
    acl_resource_name: pulumi.Output[str]
    """
    Resource name for this acl. The resource name should be a topic or consumer group name.
    """
    acl_resource_pattern_type: pulumi.Output[str]
    """
    Resource pattern type for this acl. The resource pattern support two types "LITERAL" and "PREFIXED". "LITERAL": A literal name defines the full name of a resource. The special wildcard character "*" can be used to represent a resource with any name. "PREFIXED": A prefixed name defines a prefix for a resource.
    """
    acl_resource_type: pulumi.Output[str]
    """
    Resource type for this acl. The resource type can only be "Topic" and "Group".
    """
    host: pulumi.Output[str]
    """
    The host of the acl.
    """
    instance_id: pulumi.Output[str]
    """
    ID of the ALIKAFKA Instance that owns the groups.
    """
    username: pulumi.Output[str]
    """
    Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.
    """
    def __init__(__self__, resource_name, opts=None, acl_operation_type=None, acl_resource_name=None, acl_resource_pattern_type=None, acl_resource_type=None, instance_id=None, username=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an ALIKAFKA sasl acl resource.

        > **NOTE:** Available in 1.66.0+

        > **NOTE:**  Only the following regions support create alikafka sasl user.
        [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`ap-southeast-1`,`ap-south-1`,`ap-southeast-5`]



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] acl_operation_type: Operation type for this acl. The operation type can only be "Write" and "Read".
        :param pulumi.Input[str] acl_resource_name: Resource name for this acl. The resource name should be a topic or consumer group name.
        :param pulumi.Input[str] acl_resource_pattern_type: Resource pattern type for this acl. The resource pattern support two types "LITERAL" and "PREFIXED". "LITERAL": A literal name defines the full name of a resource. The special wildcard character "*" can be used to represent a resource with any name. "PREFIXED": A prefixed name defines a prefix for a resource.
        :param pulumi.Input[str] acl_resource_type: Resource type for this acl. The resource type can only be "Topic" and "Group".
        :param pulumi.Input[str] instance_id: ID of the ALIKAFKA Instance that owns the groups.
        :param pulumi.Input[str] username: Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.
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

            if acl_operation_type is None:
                raise TypeError("Missing required property 'acl_operation_type'")
            __props__['acl_operation_type'] = acl_operation_type
            if acl_resource_name is None:
                raise TypeError("Missing required property 'acl_resource_name'")
            __props__['acl_resource_name'] = acl_resource_name
            if acl_resource_pattern_type is None:
                raise TypeError("Missing required property 'acl_resource_pattern_type'")
            __props__['acl_resource_pattern_type'] = acl_resource_pattern_type
            if acl_resource_type is None:
                raise TypeError("Missing required property 'acl_resource_type'")
            __props__['acl_resource_type'] = acl_resource_type
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            if username is None:
                raise TypeError("Missing required property 'username'")
            __props__['username'] = username
            __props__['host'] = None
        super(SaslAcl, __self__).__init__(
            'alicloud:alikafka/saslAcl:SaslAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, acl_operation_type=None, acl_resource_name=None, acl_resource_pattern_type=None, acl_resource_type=None, host=None, instance_id=None, username=None):
        """
        Get an existing SaslAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] acl_operation_type: Operation type for this acl. The operation type can only be "Write" and "Read".
        :param pulumi.Input[str] acl_resource_name: Resource name for this acl. The resource name should be a topic or consumer group name.
        :param pulumi.Input[str] acl_resource_pattern_type: Resource pattern type for this acl. The resource pattern support two types "LITERAL" and "PREFIXED". "LITERAL": A literal name defines the full name of a resource. The special wildcard character "*" can be used to represent a resource with any name. "PREFIXED": A prefixed name defines a prefix for a resource.
        :param pulumi.Input[str] acl_resource_type: Resource type for this acl. The resource type can only be "Topic" and "Group".
        :param pulumi.Input[str] host: The host of the acl.
        :param pulumi.Input[str] instance_id: ID of the ALIKAFKA Instance that owns the groups.
        :param pulumi.Input[str] username: Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["acl_operation_type"] = acl_operation_type
        __props__["acl_resource_name"] = acl_resource_name
        __props__["acl_resource_pattern_type"] = acl_resource_pattern_type
        __props__["acl_resource_type"] = acl_resource_type
        __props__["host"] = host
        __props__["instance_id"] = instance_id
        __props__["username"] = username
        return SaslAcl(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

