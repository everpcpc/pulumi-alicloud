# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .get_instances import *
from .get_user_groups import *
from .get_users import *
from .instance import *
from .user import *
from .user_group import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "alicloud:bastionhost/instance:Instance":
                return Instance(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:bastionhost/user:User":
                return User(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:bastionhost/userGroup:UserGroup":
                return UserGroup(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("alicloud", "bastionhost/instance", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "bastionhost/user", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "bastionhost/userGroup", _module_instance)

_register_module()