# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class FileSystem(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The File System description.
    """
    protocol_type: pulumi.Output[str]
    """
    The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
    """
    storage_type: pulumi.Output[str]
    """
    The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
    """
    def __init__(__self__, resource_name, opts=None, description=None, protocol_type=None, storage_type=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Nas File System resource.
        
        After activating NAS, you can create a file system and purchase a storage package for it in the NAS console. The NAS console also enables you to view the file system details and remove unnecessary file systems.
        
        For information about NAS file system and how to use it, see [Manage file systems](https://www.alibabacloud.com/help/doc-detail/27530.htm)
        
        > **NOTE:** Available in v1.33.0+.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The File System description.
        :param pulumi.Input[str] protocol_type: The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
        :param pulumi.Input[str] storage_type: The Storage Type of a File System. Valid values: `Capacity` and `Performance`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/nas_file_system.html.markdown.
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
            if protocol_type is None:
                raise TypeError("Missing required property 'protocol_type'")
            __props__['protocol_type'] = protocol_type
            if storage_type is None:
                raise TypeError("Missing required property 'storage_type'")
            __props__['storage_type'] = storage_type
        super(FileSystem, __self__).__init__(
            'alicloud:nas/fileSystem:FileSystem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, protocol_type=None, storage_type=None):
        """
        Get an existing FileSystem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The File System description.
        :param pulumi.Input[str] protocol_type: The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
        :param pulumi.Input[str] storage_type: The Storage Type of a File System. Valid values: `Capacity` and `Performance`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/nas_file_system.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["description"] = description
        __props__["protocol_type"] = protocol_type
        __props__["storage_type"] = storage_type
        return FileSystem(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

