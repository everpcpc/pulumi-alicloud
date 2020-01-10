# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AccessKey(pulumi.CustomResource):
    encrypted_secret: pulumi.Output[str]
    key_fingerprint: pulumi.Output[str]
    """
    The fingerprint of the PGP key used to encrypt the secret
    """
    pgp_key: pulumi.Output[str]
    """
    Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`
    """
    secret_file: pulumi.Output[str]
    """
    The name of file that can save access key id and access key secret. Strongly suggest you to specified it when you creating access key, otherwise, you wouldn't get its secret ever.
    """
    status: pulumi.Output[str]
    """
    Status of access key. It must be `Active` or `Inactive`. Default value is `Active`.
    """
    user_name: pulumi.Output[str]
    """
    Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
    """
    def __init__(__self__, resource_name, opts=None, pgp_key=None, secret_file=None, status=None, user_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a RAM User access key resource.
        
        > **NOTE:**  You should set the `secret_file` if you want to get the access key.  
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] pgp_key: Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`
        :param pulumi.Input[str] secret_file: The name of file that can save access key id and access key secret. Strongly suggest you to specified it when you creating access key, otherwise, you wouldn't get its secret ever.
        :param pulumi.Input[str] status: Status of access key. It must be `Active` or `Inactive`. Default value is `Active`.
        :param pulumi.Input[str] user_name: Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/ram_access_key.html.markdown.
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

            __props__['pgp_key'] = pgp_key
            __props__['secret_file'] = secret_file
            __props__['status'] = status
            __props__['user_name'] = user_name
            __props__['encrypted_secret'] = None
            __props__['key_fingerprint'] = None
        super(AccessKey, __self__).__init__(
            'alicloud:ram/accessKey:AccessKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, encrypted_secret=None, key_fingerprint=None, pgp_key=None, secret_file=None, status=None, user_name=None):
        """
        Get an existing AccessKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_fingerprint: The fingerprint of the PGP key used to encrypt the secret
        :param pulumi.Input[str] pgp_key: Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`
        :param pulumi.Input[str] secret_file: The name of file that can save access key id and access key secret. Strongly suggest you to specified it when you creating access key, otherwise, you wouldn't get its secret ever.
        :param pulumi.Input[str] status: Status of access key. It must be `Active` or `Inactive`. Default value is `Active`.
        :param pulumi.Input[str] user_name: Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/ram_access_key.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["encrypted_secret"] = encrypted_secret
        __props__["key_fingerprint"] = key_fingerprint
        __props__["pgp_key"] = pgp_key
        __props__["secret_file"] = secret_file
        __props__["status"] = status
        __props__["user_name"] = user_name
        return AccessKey(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

