# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['VpcEndpointService']


class VpcEndpointService(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_accept_connection: Optional[pulumi.Input[bool]] = None,
                 connect_bandwidth: Optional[pulumi.Input[int]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 payer: Optional[pulumi.Input[str]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcEndpointServiceResourceArgs']]]]] = None,
                 service_description: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Private Link Vpc Endpoint Service resource.

        For information about Private Link Vpc Endpoint Service and how to use it, see [What is Vpc Endpoint Service](https://help.aliyun.com/document_detail/183540.html).

        > **NOTE:** Available in v1.109.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.privatelink.VpcEndpointService("example",
            auto_accept_connection=False,
            connect_bandwidth=103,
            resources=[alicloud.privatelink.VpcEndpointServiceResourceArgs(
                resource_id="lb-gw8nxxxxxxx",
                resource_type="slb",
            )],
            service_description="tftest")
        ```

        ## Import

        Private Link Vpc Endpoint Service can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:privatelink/vpcEndpointService:VpcEndpointService example <service_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_accept_connection: Whether to automatically accept terminal node connections.
        :param pulumi.Input[int] connect_bandwidth: The connection bandwidth.
        :param pulumi.Input[bool] dry_run: Whether to pre-check this request only. Default to: `false`
        :param pulumi.Input[str] payer: The payer type. Valid Value: `EndpointService`, `Endpoint`. Default to: `Endpoint`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcEndpointServiceResourceArgs']]]] resources: Service resources added to the endpoint service.
        :param pulumi.Input[str] service_description: The description of the terminal node service.
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

            __props__['auto_accept_connection'] = auto_accept_connection
            __props__['connect_bandwidth'] = connect_bandwidth
            __props__['dry_run'] = dry_run
            __props__['payer'] = payer
            __props__['resources'] = resources
            __props__['service_description'] = service_description
            __props__['service_business_status'] = None
            __props__['service_domain'] = None
            __props__['status'] = None
        super(VpcEndpointService, __self__).__init__(
            'alicloud:privatelink/vpcEndpointService:VpcEndpointService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_accept_connection: Optional[pulumi.Input[bool]] = None,
            connect_bandwidth: Optional[pulumi.Input[int]] = None,
            dry_run: Optional[pulumi.Input[bool]] = None,
            payer: Optional[pulumi.Input[str]] = None,
            resources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcEndpointServiceResourceArgs']]]]] = None,
            service_business_status: Optional[pulumi.Input[str]] = None,
            service_description: Optional[pulumi.Input[str]] = None,
            service_domain: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'VpcEndpointService':
        """
        Get an existing VpcEndpointService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_accept_connection: Whether to automatically accept terminal node connections.
        :param pulumi.Input[int] connect_bandwidth: The connection bandwidth.
        :param pulumi.Input[bool] dry_run: Whether to pre-check this request only. Default to: `false`
        :param pulumi.Input[str] payer: The payer type. Valid Value: `EndpointService`, `Endpoint`. Default to: `Endpoint`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcEndpointServiceResourceArgs']]]] resources: Service resources added to the endpoint service.
        :param pulumi.Input[str] service_business_status: The business status of Vpc Endpoint Service.
        :param pulumi.Input[str] service_description: The description of the terminal node service.
        :param pulumi.Input[str] service_domain: Service Domain.
        :param pulumi.Input[str] status: The status of Vpc Endpoint Service.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_accept_connection"] = auto_accept_connection
        __props__["connect_bandwidth"] = connect_bandwidth
        __props__["dry_run"] = dry_run
        __props__["payer"] = payer
        __props__["resources"] = resources
        __props__["service_business_status"] = service_business_status
        __props__["service_description"] = service_description
        __props__["service_domain"] = service_domain
        __props__["status"] = status
        return VpcEndpointService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoAcceptConnection")
    def auto_accept_connection(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to automatically accept terminal node connections.
        """
        return pulumi.get(self, "auto_accept_connection")

    @property
    @pulumi.getter(name="connectBandwidth")
    def connect_bandwidth(self) -> pulumi.Output[Optional[int]]:
        """
        The connection bandwidth.
        """
        return pulumi.get(self, "connect_bandwidth")

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to pre-check this request only. Default to: `false`
        """
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter
    def payer(self) -> pulumi.Output[Optional[str]]:
        """
        The payer type. Valid Value: `EndpointService`, `Endpoint`. Default to: `Endpoint`.
        """
        return pulumi.get(self, "payer")

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Output[Optional[Sequence['outputs.VpcEndpointServiceResource']]]:
        """
        Service resources added to the endpoint service.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="serviceBusinessStatus")
    def service_business_status(self) -> pulumi.Output[str]:
        """
        The business status of Vpc Endpoint Service.
        """
        return pulumi.get(self, "service_business_status")

    @property
    @pulumi.getter(name="serviceDescription")
    def service_description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the terminal node service.
        """
        return pulumi.get(self, "service_description")

    @property
    @pulumi.getter(name="serviceDomain")
    def service_domain(self) -> pulumi.Output[str]:
        """
        Service Domain.
        """
        return pulumi.get(self, "service_domain")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of Vpc Endpoint Service.
        """
        return pulumi.get(self, "status")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

