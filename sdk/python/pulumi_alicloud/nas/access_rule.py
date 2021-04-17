# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AccessRuleArgs', 'AccessRule']

@pulumi.input_type
class AccessRuleArgs:
    def __init__(__self__, *,
                 access_group_name: pulumi.Input[str],
                 source_cidr_ip: pulumi.Input[str],
                 priority: Optional[pulumi.Input[int]] = None,
                 rw_access_type: Optional[pulumi.Input[str]] = None,
                 user_access_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AccessRule resource.
        :param pulumi.Input[str] access_group_name: Permission group name.
        :param pulumi.Input[str] source_cidr_ip: Address or address segment.
        :param pulumi.Input[int] priority: Priority level. Range: 1-100. Default value: `1`.
        :param pulumi.Input[str] rw_access_type: Read-write permission type: `RDWR` (default), `RDONLY`.
        :param pulumi.Input[str] user_access_type: User permission type: `no_squash` (default), `root_squash`, `all_squash`.
        """
        pulumi.set(__self__, "access_group_name", access_group_name)
        pulumi.set(__self__, "source_cidr_ip", source_cidr_ip)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if rw_access_type is not None:
            pulumi.set(__self__, "rw_access_type", rw_access_type)
        if user_access_type is not None:
            pulumi.set(__self__, "user_access_type", user_access_type)

    @property
    @pulumi.getter(name="accessGroupName")
    def access_group_name(self) -> pulumi.Input[str]:
        """
        Permission group name.
        """
        return pulumi.get(self, "access_group_name")

    @access_group_name.setter
    def access_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_group_name", value)

    @property
    @pulumi.getter(name="sourceCidrIp")
    def source_cidr_ip(self) -> pulumi.Input[str]:
        """
        Address or address segment.
        """
        return pulumi.get(self, "source_cidr_ip")

    @source_cidr_ip.setter
    def source_cidr_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_cidr_ip", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        Priority level. Range: 1-100. Default value: `1`.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="rwAccessType")
    def rw_access_type(self) -> Optional[pulumi.Input[str]]:
        """
        Read-write permission type: `RDWR` (default), `RDONLY`.
        """
        return pulumi.get(self, "rw_access_type")

    @rw_access_type.setter
    def rw_access_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rw_access_type", value)

    @property
    @pulumi.getter(name="userAccessType")
    def user_access_type(self) -> Optional[pulumi.Input[str]]:
        """
        User permission type: `no_squash` (default), `root_squash`, `all_squash`.
        """
        return pulumi.get(self, "user_access_type")

    @user_access_type.setter
    def user_access_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_access_type", value)


@pulumi.input_type
class _AccessRuleState:
    def __init__(__self__, *,
                 access_group_name: Optional[pulumi.Input[str]] = None,
                 access_rule_id: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 rw_access_type: Optional[pulumi.Input[str]] = None,
                 source_cidr_ip: Optional[pulumi.Input[str]] = None,
                 user_access_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AccessRule resources.
        :param pulumi.Input[str] access_group_name: Permission group name.
        :param pulumi.Input[str] access_rule_id: The nas access rule ID.
        :param pulumi.Input[int] priority: Priority level. Range: 1-100. Default value: `1`.
        :param pulumi.Input[str] rw_access_type: Read-write permission type: `RDWR` (default), `RDONLY`.
        :param pulumi.Input[str] source_cidr_ip: Address or address segment.
        :param pulumi.Input[str] user_access_type: User permission type: `no_squash` (default), `root_squash`, `all_squash`.
        """
        if access_group_name is not None:
            pulumi.set(__self__, "access_group_name", access_group_name)
        if access_rule_id is not None:
            pulumi.set(__self__, "access_rule_id", access_rule_id)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if rw_access_type is not None:
            pulumi.set(__self__, "rw_access_type", rw_access_type)
        if source_cidr_ip is not None:
            pulumi.set(__self__, "source_cidr_ip", source_cidr_ip)
        if user_access_type is not None:
            pulumi.set(__self__, "user_access_type", user_access_type)

    @property
    @pulumi.getter(name="accessGroupName")
    def access_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Permission group name.
        """
        return pulumi.get(self, "access_group_name")

    @access_group_name.setter
    def access_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_group_name", value)

    @property
    @pulumi.getter(name="accessRuleId")
    def access_rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        The nas access rule ID.
        """
        return pulumi.get(self, "access_rule_id")

    @access_rule_id.setter
    def access_rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_rule_id", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        Priority level. Range: 1-100. Default value: `1`.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="rwAccessType")
    def rw_access_type(self) -> Optional[pulumi.Input[str]]:
        """
        Read-write permission type: `RDWR` (default), `RDONLY`.
        """
        return pulumi.get(self, "rw_access_type")

    @rw_access_type.setter
    def rw_access_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rw_access_type", value)

    @property
    @pulumi.getter(name="sourceCidrIp")
    def source_cidr_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Address or address segment.
        """
        return pulumi.get(self, "source_cidr_ip")

    @source_cidr_ip.setter
    def source_cidr_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_cidr_ip", value)

    @property
    @pulumi.getter(name="userAccessType")
    def user_access_type(self) -> Optional[pulumi.Input[str]]:
        """
        User permission type: `no_squash` (default), `root_squash`, `all_squash`.
        """
        return pulumi.get(self, "user_access_type")

    @user_access_type.setter
    def user_access_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_access_type", value)


class AccessRule(pulumi.CustomResource):
    @overload
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
            access_group_name="tf-NasConfigName",
            access_group_type="Vpc",
            description="tf-testAccNasConfig")
        foo_access_rule = alicloud.nas.AccessRule("fooAccessRule",
            access_group_name=foo_access_group.access_group_name,
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
        :param pulumi.Input[int] priority: Priority level. Range: 1-100. Default value: `1`.
        :param pulumi.Input[str] rw_access_type: Read-write permission type: `RDWR` (default), `RDONLY`.
        :param pulumi.Input[str] source_cidr_ip: Address or address segment.
        :param pulumi.Input[str] user_access_type: User permission type: `no_squash` (default), `root_squash`, `all_squash`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccessRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
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
            access_group_name="tf-NasConfigName",
            access_group_type="Vpc",
            description="tf-testAccNasConfig")
        foo_access_rule = alicloud.nas.AccessRule("fooAccessRule",
            access_group_name=foo_access_group.access_group_name,
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
        :param AccessRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccessRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
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
            __props__ = AccessRuleArgs.__new__(AccessRuleArgs)

            if access_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'access_group_name'")
            __props__.__dict__["access_group_name"] = access_group_name
            __props__.__dict__["priority"] = priority
            __props__.__dict__["rw_access_type"] = rw_access_type
            if source_cidr_ip is None and not opts.urn:
                raise TypeError("Missing required property 'source_cidr_ip'")
            __props__.__dict__["source_cidr_ip"] = source_cidr_ip
            __props__.__dict__["user_access_type"] = user_access_type
            __props__.__dict__["access_rule_id"] = None
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
        :param pulumi.Input[int] priority: Priority level. Range: 1-100. Default value: `1`.
        :param pulumi.Input[str] rw_access_type: Read-write permission type: `RDWR` (default), `RDONLY`.
        :param pulumi.Input[str] source_cidr_ip: Address or address segment.
        :param pulumi.Input[str] user_access_type: User permission type: `no_squash` (default), `root_squash`, `all_squash`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccessRuleState.__new__(_AccessRuleState)

        __props__.__dict__["access_group_name"] = access_group_name
        __props__.__dict__["access_rule_id"] = access_rule_id
        __props__.__dict__["priority"] = priority
        __props__.__dict__["rw_access_type"] = rw_access_type
        __props__.__dict__["source_cidr_ip"] = source_cidr_ip
        __props__.__dict__["user_access_type"] = user_access_type
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
        Priority level. Range: 1-100. Default value: `1`.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="rwAccessType")
    def rw_access_type(self) -> pulumi.Output[Optional[str]]:
        """
        Read-write permission type: `RDWR` (default), `RDONLY`.
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
        User permission type: `no_squash` (default), `root_squash`, `all_squash`.
        """
        return pulumi.get(self, "user_access_type")

