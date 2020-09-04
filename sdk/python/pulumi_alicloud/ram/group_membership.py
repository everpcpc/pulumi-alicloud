# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['GroupMembership']


class GroupMembership(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 user_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a RAM Group membership resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        # Create a RAM Group membership.
        group = alicloud.ram.Group("group",
            comments="this is a group comments.",
            force=True)
        user = alicloud.ram.User("user",
            display_name="user_display_name",
            mobile="86-18688888888",
            email="hello.uuu@aaa.com",
            comments="yoyoyo",
            force=True)
        user1 = alicloud.ram.User("user1",
            display_name="user_display_name1",
            mobile="86-18688888889",
            email="hello.uuu@aaa.com",
            comments="yoyoyo",
            force=True)
        membership = alicloud.ram.GroupMembership("membership",
            group_name=group.name,
            user_names=[
                user.name,
                user1.name,
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_name: Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        :param pulumi.Input[List[pulumi.Input[str]]] user_names: Set of user name which will be added to group. Each name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if group_name is None:
                raise TypeError("Missing required property 'group_name'")
            __props__['group_name'] = group_name
            if user_names is None:
                raise TypeError("Missing required property 'user_names'")
            __props__['user_names'] = user_names
        super(GroupMembership, __self__).__init__(
            'alicloud:ram/groupMembership:GroupMembership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group_name: Optional[pulumi.Input[str]] = None,
            user_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'GroupMembership':
        """
        Get an existing GroupMembership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_name: Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        :param pulumi.Input[List[pulumi.Input[str]]] user_names: Set of user name which will be added to group. Each name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["group_name"] = group_name
        __props__["user_names"] = user_names
        return GroupMembership(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Output[str]:
        """
        Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter(name="userNames")
    def user_names(self) -> pulumi.Output[List[str]]:
        """
        Set of user name which will be added to group. Each name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
        """
        return pulumi.get(self, "user_names")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

