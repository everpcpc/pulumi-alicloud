# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Acl(pulumi.CustomResource):
    entry_lists: pulumi.Output[list]
    """
    A list of entry (IP addresses or CIDR blocks) to be added. At most 50 etnry can be supported in one resource. It contains two sub-fields as `Entry Block` follows.

      * `comment` (`str`)
      * `entry` (`str`)
    """
    ip_version: pulumi.Output[str]
    """
    The IP Version of access control list is the type of its entry (IP addresses or CIDR blocks). It values ipv4/ipv6. Our plugin provides a default ip_version: "ipv4".
    """
    name: pulumi.Output[str]
    """
    Name of the access control list.
    """
    resource_group_id: pulumi.Output[str]
    """
    Resource group ID.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, entry_lists=None, ip_version=None, name=None, resource_group_id=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        An access control list contains multiple IP addresses or CIDR blocks.
        The access control list can help you to define multiple instance listening dimension,
        and to meet the multiple usage for single access control list.

        Server Load Balancer allows you to configure access control for listeners.
        You can configure different whitelists or blacklists for different listeners.

        You can configure access control
        when you create a listener or change access control configuration after a listener is created.

        > **NOTE:** One access control list can be attached to many Listeners in different load balancer as whitelists or blacklists.

        > **NOTE:** The maximum number of access control lists per region  is 50.

        > **NOTE:** The maximum number of IP addresses added each time is 50.

        > **NOTE:** The maximum number of entries per access control list is 300.

        > **NOTE:** The maximum number of listeners that an access control list can be added to is 50.

        For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).

        For information about acl and how to use it, see [Configure an access control list](https://www.alibabacloud.com/help/doc-detail/85978.htm).


        ## Entry Block

        The entry mapping supports the following:

        * `entry` - (Required) An IP addresses or CIDR blocks.
        * `comment` - (Optional) the comment of the entry.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/slb_acl.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] entry_lists: A list of entry (IP addresses or CIDR blocks) to be added. At most 50 etnry can be supported in one resource. It contains two sub-fields as `Entry Block` follows.
        :param pulumi.Input[str] ip_version: The IP Version of access control list is the type of its entry (IP addresses or CIDR blocks). It values ipv4/ipv6. Our plugin provides a default ip_version: "ipv4".
        :param pulumi.Input[str] name: Name of the access control list.
        :param pulumi.Input[str] resource_group_id: Resource group ID.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **entry_lists** object supports the following:

          * `comment` (`pulumi.Input[str]`)
          * `entry` (`pulumi.Input[str]`)
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

            __props__['entry_lists'] = entry_lists
            __props__['ip_version'] = ip_version
            __props__['name'] = name
            __props__['resource_group_id'] = resource_group_id
            __props__['tags'] = tags
        super(Acl, __self__).__init__(
            'alicloud:slb/acl:Acl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, entry_lists=None, ip_version=None, name=None, resource_group_id=None, tags=None):
        """
        Get an existing Acl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] entry_lists: A list of entry (IP addresses or CIDR blocks) to be added. At most 50 etnry can be supported in one resource. It contains two sub-fields as `Entry Block` follows.
        :param pulumi.Input[str] ip_version: The IP Version of access control list is the type of its entry (IP addresses or CIDR blocks). It values ipv4/ipv6. Our plugin provides a default ip_version: "ipv4".
        :param pulumi.Input[str] name: Name of the access control list.
        :param pulumi.Input[str] resource_group_id: Resource group ID.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **entry_lists** object supports the following:

          * `comment` (`pulumi.Input[str]`)
          * `entry` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["entry_lists"] = entry_lists
        __props__["ip_version"] = ip_version
        __props__["name"] = name
        __props__["resource_group_id"] = resource_group_id
        __props__["tags"] = tags
        return Acl(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

