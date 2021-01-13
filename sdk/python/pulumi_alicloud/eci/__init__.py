# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .container_group import *
from .get_container_groups import *
from .get_image_caches import *
from .image_cache import *
from .open_api_image_cache import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "alicloud:eci/containerGroup:ContainerGroup":
                return ContainerGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:eci/imageCache:ImageCache":
                return ImageCache(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:eci/openApiImageCache:OpenApiImageCache":
                return OpenApiImageCache(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("alicloud", "eci/containerGroup", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "eci/imageCache", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "eci/openApiImageCache", _module_instance)

_register_module()
