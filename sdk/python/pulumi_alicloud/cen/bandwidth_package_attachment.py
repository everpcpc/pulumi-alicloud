# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['BandwidthPackageAttachment']


class BandwidthPackageAttachment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth_package_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a CEN bandwidth package attachment resource. The resource can be used to bind a bandwidth package to a specified CEN instance.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        # Create a new bandwidth package attachment and use it to attach a bandwidth package to a new CEN
        cen = alicloud.cen.Instance("cen", description="tf-testAccCenBandwidthPackageAttachmentDescription")
        bwp = alicloud.cen.BandwidthPackage("bwp",
            bandwidth=20,
            geographic_region_ids=[
                "China",
                "Asia-Pacific",
            ])
        foo = alicloud.cen.BandwidthPackageAttachment("foo",
            instance_id=cen.id,
            bandwidth_package_id=bwp.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bandwidth_package_id: The ID of the bandwidth package.
        :param pulumi.Input[str] instance_id: The ID of the CEN.
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

            if bandwidth_package_id is None:
                raise TypeError("Missing required property 'bandwidth_package_id'")
            __props__['bandwidth_package_id'] = bandwidth_package_id
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
        super(BandwidthPackageAttachment, __self__).__init__(
            'alicloud:cen/bandwidthPackageAttachment:BandwidthPackageAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bandwidth_package_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'BandwidthPackageAttachment':
        """
        Get an existing BandwidthPackageAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bandwidth_package_id: The ID of the bandwidth package.
        :param pulumi.Input[str] instance_id: The ID of the CEN.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bandwidth_package_id"] = bandwidth_package_id
        __props__["instance_id"] = instance_id
        return BandwidthPackageAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bandwidthPackageId")
    def bandwidth_package_id(self) -> pulumi.Output[str]:
        """
        The ID of the bandwidth package.
        """
        return pulumi.get(self, "bandwidth_package_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the CEN.
        """
        return pulumi.get(self, "instance_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

