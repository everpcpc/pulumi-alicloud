# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['AccessRule']


class AccessRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_group_name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 rw_access_type: Optional[pulumi.Input[str]] = None,
                 source_cidr_ip: Optional[pulumi.Input[str]] = None,
                 user_access_type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Nas Access Rule resource.

        When NAS is activated, the Default VPC Permission Group is automatically generated. It allows all IP addresses in a VPC to access the mount point with full permissions. Full permissions include Read/Write permission with no restriction on root users.

        > **NOTE:** Available in v1.34.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        foo_access_group = alicloud.nas.AccessGroup("fooAccessGroup",
            type="Vpc",
            description="tf-testAccNasConfig")
        foo_access_rule = alicloud.nas.AccessRule("fooAccessRule",
            access_group_name=foo_access_group.id,
            source_cidr_ip="168.1.1.0/16",
            rw_access_type="RDWR",
            user_access_type="no_squash",
            priority=2)
        ```

        ## Import

        Nas Access Rule can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:nas/accessRule:AccessRule foo tf-testAccNasConfigName:1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_group_name: Permission group name.
        :param pulumi.Input[int] priority: Priority level. Range: 1-100. Default value: 1.
        :param pulumi.Input[str] rw_access_type: Read-write permission type: RDWR (default), RDONLY.
        :param pulumi.Input[str] source_cidr_ip: Address or address segment.
        :param pulumi.Input[str] user_access_type: User permission type: no_squash (default), root_squash, all_squash.
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

            if access_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'access_group_name'")
            __props__['access_group_name'] = access_group_name
            __props__['priority'] = priority
            __props__['rw_access_type'] = rw_access_type
            if source_cidr_ip is None and not opts.urn:
                raise TypeError("Missing required property 'source_cidr_ip'")
            __props__['source_cidr_ip'] = source_cidr_ip
            __props__['user_access_type'] = user_access_type
            __props__['access_rule_id'] = None
        super(AccessRule, __self__).__init__(
            'alicloud:nas/accessRule:AccessRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_group_name: Optional[pulumi.Input[str]] = None,
            access_rule_id: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            rw_access_type: Optional[pulumi.Input[str]] = None,
            source_cidr_ip: Optional[pulumi.Input[str]] = None,
            user_access_type: Optional[pulumi.Input[str]] = None) -> 'AccessRule':
        """
        Get an existing AccessRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_group_name: Permission group name.
        :param pulumi.Input[str] access_rule_id: The nas access rule ID.
        :param pulumi.Input[int] priority: Priority level. Range: 1-100. Default value: 1.
        :param pulumi.Input[str] rw_access_type: Read-write permission type: RDWR (default), RDONLY.
        :param pulumi.Input[str] source_cidr_ip: Address or address segment.
        :param pulumi.Input[str] user_access_type: User permission type: no_squash (default), root_squash, all_squash.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_group_name"] = access_group_name
        __props__["access_rule_id"] = access_rule_id
        __props__["priority"] = priority
        __props__["rw_access_type"] = rw_access_type
        __props__["source_cidr_ip"] = source_cidr_ip
        __props__["user_access_type"] = user_access_type
        return AccessRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessGroupName")
    def access_group_name(self) -> pulumi.Output[str]:
        """
        Permission group name.
        """
        return pulumi.get(self, "access_group_name")

    @property
    @pulumi.getter(name="accessRuleId")
    def access_rule_id(self) -> pulumi.Output[str]:
        """
        The nas access rule ID.
        """
        return pulumi.get(self, "access_rule_id")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[Optional[int]]:
        """
        Priority level. Range: 1-100. Default value: 1.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="rwAccessType")
    def rw_access_type(self) -> pulumi.Output[Optional[str]]:
        """
        Read-write permission type: RDWR (default), RDONLY.
        """
        return pulumi.get(self, "rw_access_type")

    @property
    @pulumi.getter(name="sourceCidrIp")
    def source_cidr_ip(self) -> pulumi.Output[str]:
        """
        Address or address segment.
        """
        return pulumi.get(self, "source_cidr_ip")

    @property
    @pulumi.getter(name="userAccessType")
    def user_access_type(self) -> pulumi.Output[Optional[str]]:
        """
        User permission type: no_squash (default), root_squash, all_squash.
        """
        return pulumi.get(self, "user_access_type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

