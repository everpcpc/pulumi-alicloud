# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .get_service import *
from .project import *
from .subscription import *
from .topic import *

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "alicloud:datahub/project:Project":
                return Project(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:datahub/subscription:Subscription":
                return Subscription(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:datahub/topic:Topic":
                return Topic(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("alicloud", "datahub/project", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "datahub/subscription", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "datahub/topic", _module_instance)

_register_module()
