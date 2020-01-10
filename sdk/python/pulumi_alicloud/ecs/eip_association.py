# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class EipAssociation(pulumi.CustomResource):
    allocation_id: pulumi.Output[str]
    """
    The allocation EIP ID.
    """
    instance_id: pulumi.Output[str]
    """
    The ID of the ECS or SLB instance or Nat Gateway.
    """
    instance_type: pulumi.Output[str]
    """
    The type of cloud product that the eip instance to bind.
    """
    private_ip_address: pulumi.Output[str]
    """
    The private IP address in the network segment of the vswitch which has been assigned.
    """
    def __init__(__self__, resource_name, opts=None, allocation_id=None, instance_id=None, instance_type=None, private_ip_address=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an Alicloud EIP Association resource for associating Elastic IP to ECS Instance, SLB Instance or Nat Gateway.
        
        > **NOTE:** `ecs.EipAssociation` is useful in scenarios where EIPs are either
         pre-existing or distributed to customers or users and therefore cannot be changed.
        
        > **NOTE:** From version 1.7.1, the resource support to associate EIP to SLB Instance or Nat Gateway.
        
        > **NOTE:** One EIP can only be associated with ECS or SLB instance which in the VPC.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_id: The allocation EIP ID.
        :param pulumi.Input[str] instance_id: The ID of the ECS or SLB instance or Nat Gateway.
        :param pulumi.Input[str] instance_type: The type of cloud product that the eip instance to bind.
        :param pulumi.Input[str] private_ip_address: The private IP address in the network segment of the vswitch which has been assigned.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/eip_association.html.markdown.
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

            if allocation_id is None:
                raise TypeError("Missing required property 'allocation_id'")
            __props__['allocation_id'] = allocation_id
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            __props__['instance_type'] = instance_type
            __props__['private_ip_address'] = private_ip_address
        super(EipAssociation, __self__).__init__(
            'alicloud:ecs/eipAssociation:EipAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, allocation_id=None, instance_id=None, instance_type=None, private_ip_address=None):
        """
        Get an existing EipAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_id: The allocation EIP ID.
        :param pulumi.Input[str] instance_id: The ID of the ECS or SLB instance or Nat Gateway.
        :param pulumi.Input[str] instance_type: The type of cloud product that the eip instance to bind.
        :param pulumi.Input[str] private_ip_address: The private IP address in the network segment of the vswitch which has been assigned.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/eip_association.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["allocation_id"] = allocation_id
        __props__["instance_id"] = instance_id
        __props__["instance_type"] = instance_type
        __props__["private_ip_address"] = private_ip_address
        return EipAssociation(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

