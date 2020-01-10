# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class CustomerGateway(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The description of the VPN customer gateway instance.
    """
    ip_address: pulumi.Output[str]
    """
    The IP address of the customer gateway.
    """
    name: pulumi.Output[str]
    """
    The name of the VPN customer gateway. Defaults to null.
    """
    def __init__(__self__, resource_name, opts=None, description=None, ip_address=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a CustomerGateway resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the VPN customer gateway instance.
        :param pulumi.Input[str] ip_address: The IP address of the customer gateway.
        :param pulumi.Input[str] name: The name of the VPN customer gateway. Defaults to null.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/vpn_customer_gateway.html.markdown.
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

            __props__['description'] = description
            if ip_address is None:
                raise TypeError("Missing required property 'ip_address'")
            __props__['ip_address'] = ip_address
            __props__['name'] = name
        super(CustomerGateway, __self__).__init__(
            'alicloud:vpn/customerGateway:CustomerGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, ip_address=None, name=None):
        """
        Get an existing CustomerGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the VPN customer gateway instance.
        :param pulumi.Input[str] ip_address: The IP address of the customer gateway.
        :param pulumi.Input[str] name: The name of the VPN customer gateway. Defaults to null.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/vpn_customer_gateway.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["description"] = description
        __props__["ip_address"] = ip_address
        __props__["name"] = name
        return CustomerGateway(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

