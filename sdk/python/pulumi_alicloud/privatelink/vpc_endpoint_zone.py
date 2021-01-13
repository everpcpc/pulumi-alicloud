# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['VpcEndpointZone']


class VpcEndpointZone(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Private Link Vpc Endpoint Zone resource.

        For information about Private Link Vpc Endpoint Zone and how to use it, see [What is Vpc Endpoint Zone](https://help.aliyun.com/document_detail/183561.html).

        > **NOTE:** Available in v1.111.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.privatelink.VpcEndpointZone("example",
            endpoint_id="ep-gw8boxxxxx",
            vswitch_id="vsw-rtycxxxxx",
            zone_id="eu-central-1a")
        ```

        ## Import

        Private Link Vpc Endpoint Zone can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:privatelink/vpcEndpointZone:VpcEndpointZone example <endpoint_id>:<zone_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] dry_run: The dry run.
        :param pulumi.Input[str] endpoint_id: The ID of the Vpc Endpoint.
        :param pulumi.Input[str] vswitch_id: The VSwitch id.
        :param pulumi.Input[str] zone_id: The Zone Id.
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

            __props__['dry_run'] = dry_run
            if endpoint_id is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_id'")
            __props__['endpoint_id'] = endpoint_id
            if vswitch_id is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__['vswitch_id'] = vswitch_id
            __props__['zone_id'] = zone_id
            __props__['status'] = None
        super(VpcEndpointZone, __self__).__init__(
            'alicloud:privatelink/vpcEndpointZone:VpcEndpointZone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dry_run: Optional[pulumi.Input[bool]] = None,
            endpoint_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'VpcEndpointZone':
        """
        Get an existing VpcEndpointZone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] dry_run: The dry run.
        :param pulumi.Input[str] endpoint_id: The ID of the Vpc Endpoint.
        :param pulumi.Input[str] status: Status.
        :param pulumi.Input[str] vswitch_id: The VSwitch id.
        :param pulumi.Input[str] zone_id: The Zone Id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dry_run"] = dry_run
        __props__["endpoint_id"] = endpoint_id
        __props__["status"] = status
        __props__["vswitch_id"] = vswitch_id
        __props__["zone_id"] = zone_id
        return VpcEndpointZone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> pulumi.Output[Optional[bool]]:
        """
        The dry run.
        """
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> pulumi.Output[str]:
        """
        The ID of the Vpc Endpoint.
        """
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[str]:
        """
        The VSwitch id.
        """
        return pulumi.get(self, "vswitch_id")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The Zone Id.
        """
        return pulumi.get(self, "zone_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

