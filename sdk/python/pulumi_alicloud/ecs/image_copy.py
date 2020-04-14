# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ImageCopy(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
    """
    encrypted: pulumi.Output[bool]
    """
    Indicates whether to encrypt the image.
    """
    force: pulumi.Output[bool]
    """
    Indicates whether to force delete the custom image, Default is `false`. 
    - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
    - false：Verifies that the image is not currently in use by any other instances before deleting the image.
    """
    image_name: pulumi.Output[str]
    """
    The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
    """
    kms_key_id: pulumi.Output[str]
    """
    Key ID used to encrypt the image.
    """
    name: pulumi.Output[str]
    source_image_id: pulumi.Output[str]
    """
    The source image ID.
    """
    source_region_id: pulumi.Output[str]
    """
    The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
    """
    tags: pulumi.Output[dict]
    """
    The tag value of an image. The value of N ranges from 1 to 20.
    """
    def __init__(__self__, resource_name, opts=None, description=None, encrypted=None, force=None, image_name=None, kms_key_id=None, name=None, source_image_id=None, source_region_id=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Copies a custom image from one region to another. You can use copied images to perform operations in the target region, such as creating instances (RunInstances) and replacing system disks (ReplaceSystemDisk).

        > **NOTE:** You can only copy the custom image when it is in the Available state.

        > **NOTE:** You can only copy the image belonging to your Alibaba Cloud account. Images cannot be copied from one account to another.

        > **NOTE:** If the copying is not completed, you cannot call DeleteImage to delete the image but you can call CancelCopyImage to cancel the copying.

        > **NOTE:** Available in 1.66.0+.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
        :param pulumi.Input[bool] encrypted: Indicates whether to encrypt the image.
        :param pulumi.Input[bool] force: Indicates whether to force delete the custom image, Default is `false`. 
               - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
               - false：Verifies that the image is not currently in use by any other instances before deleting the image.
        :param pulumi.Input[str] image_name: The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
        :param pulumi.Input[str] kms_key_id: Key ID used to encrypt the image.
        :param pulumi.Input[str] source_image_id: The source image ID.
        :param pulumi.Input[str] source_region_id: The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
        :param pulumi.Input[dict] tags: The tag value of an image. The value of N ranges from 1 to 20.
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
        super(ImageCopy, __self__).__init__(
            'alicloud:ecs/imageCopy:ImageCopy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, encrypted=None, force=None, image_name=None, kms_key_id=None, name=None, source_image_id=None, source_region_id=None, tags=None):
        """
        Get an existing ImageCopy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
        :param pulumi.Input[bool] encrypted: Indicates whether to encrypt the image.
        :param pulumi.Input[bool] force: Indicates whether to force delete the custom image, Default is `false`. 
               - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
               - false：Verifies that the image is not currently in use by any other instances before deleting the image.
        :param pulumi.Input[str] image_name: The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
        :param pulumi.Input[str] kms_key_id: Key ID used to encrypt the image.
        :param pulumi.Input[str] source_image_id: The source image ID.
        :param pulumi.Input[str] source_region_id: The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
        :param pulumi.Input[dict] tags: The tag value of an image. The value of N ranges from 1 to 20.
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
        return ImageCopy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

