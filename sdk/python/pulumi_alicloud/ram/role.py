# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Role(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The role arn.
    """
    description: pulumi.Output[str]
    """
    Description of the RAM role. This name can have a string of 1 to 1024 characters.
    """
    document: pulumi.Output[str]
    """
    Authorization strategy of the RAM role. It is required when the `services` and `ram_users` are not specified.
    """
    force: pulumi.Output[bool]
    """
    This parameter is used for resource destroy. Default value is `false`.
    """
    name: pulumi.Output[str]
    """
    Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
    """
    ram_users: pulumi.Output[list]
    """
    (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
    """
    role_id: pulumi.Output[str]
    """
    The role ID.
    """
    services: pulumi.Output[list]
    """
    (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
    """
    version: pulumi.Output[str]
    """
    (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
    """
    def __init__(__self__, resource_name, opts=None, description=None, document=None, force=None, name=None, ram_users=None, services=None, version=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Role resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the RAM role. This name can have a string of 1 to 1024 characters.
        :param pulumi.Input[str] document: Authorization strategy of the RAM role. It is required when the `services` and `ram_users` are not specified.
        :param pulumi.Input[bool] force: This parameter is used for resource destroy. Default value is `false`.
        :param pulumi.Input[str] name: Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
        :param pulumi.Input[list] ram_users: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
        :param pulumi.Input[list] services: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
        :param pulumi.Input[str] version: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
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

            __props__['description'] = description
            __props__['document'] = document
            __props__['force'] = force
            __props__['name'] = name
            __props__['ram_users'] = ram_users
            __props__['services'] = services
            __props__['version'] = version
            __props__['arn'] = None
            __props__['role_id'] = None
        super(Role, __self__).__init__(
            'alicloud:ram/role:Role',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, description=None, document=None, force=None, name=None, ram_users=None, role_id=None, services=None, version=None):
        """
        Get an existing Role resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The role arn.
        :param pulumi.Input[str] description: Description of the RAM role. This name can have a string of 1 to 1024 characters.
        :param pulumi.Input[str] document: Authorization strategy of the RAM role. It is required when the `services` and `ram_users` are not specified.
        :param pulumi.Input[bool] force: This parameter is used for resource destroy. Default value is `false`.
        :param pulumi.Input[str] name: Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
        :param pulumi.Input[list] ram_users: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
        :param pulumi.Input[str] role_id: The role ID.
        :param pulumi.Input[list] services: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
        :param pulumi.Input[str] version: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["document"] = document
        __props__["force"] = force
        __props__["name"] = name
        __props__["ram_users"] = ram_users
        __props__["role_id"] = role_id
        __props__["services"] = services
        __props__["version"] = version
        return Role(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

