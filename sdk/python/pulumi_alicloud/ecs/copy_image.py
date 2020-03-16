# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class CopyImage(pulumi.CustomResource):
    description: pulumi.Output[str]
    encrypted: pulumi.Output[bool]
    force: pulumi.Output[bool]
    image_name: pulumi.Output[str]
    kms_key_id: pulumi.Output[str]
    name: pulumi.Output[str]
    source_image_id: pulumi.Output[str]
    source_region_id: pulumi.Output[str]
    tags: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, description=None, encrypted=None, force=None, image_name=None, kms_key_id=None, name=None, source_image_id=None, source_region_id=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a CopyImage resource with the given unique name, props, and options.
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

            __props__['description'] = description
            __props__['encrypted'] = encrypted
            __props__['force'] = force
            __props__['image_name'] = image_name
            __props__['kms_key_id'] = kms_key_id
            __props__['name'] = name
            if source_image_id is None:
                raise TypeError("Missing required property 'source_image_id'")
            __props__['source_image_id'] = source_image_id
            if source_region_id is None:
                raise TypeError("Missing required property 'source_region_id'")
            __props__['source_region_id'] = source_region_id
            __props__['tags'] = tags
        super(CopyImage, __self__).__init__(
            'alicloud:ecs/copyImage:CopyImage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, encrypted=None, force=None, image_name=None, kms_key_id=None, name=None, source_image_id=None, source_region_id=None, tags=None):
        """
        Get an existing CopyImage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["encrypted"] = encrypted
        __props__["force"] = force
        __props__["image_name"] = image_name
        __props__["kms_key_id"] = kms_key_id
        __props__["name"] = name
        __props__["source_image_id"] = source_image_id
        __props__["source_region_id"] = source_region_id
        __props__["tags"] = tags
        return CopyImage(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

