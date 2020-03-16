# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Domain(pulumi.CustomResource):
    dns_servers: pulumi.Output[list]
    """
    A list of the dns server name.
    """
    domain_id: pulumi.Output[str]
    """
    The domain ID.
    """
    group_id: pulumi.Output[str]
    """
    Id of the group in which the domain will add. If not supplied, then use default group.
    """
    name: pulumi.Output[str]
    """
    Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
    """
    resource_group_id: pulumi.Output[str]
    """
    The Id of resource group which the dns belongs.
    """
    def __init__(__self__, resource_name, opts=None, group_id=None, name=None, resource_group_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a DNS resource.

        > **NOTE:** The domain name which you want to add must be already registered and had not added by another account. Every domain name can only exist in a unique group.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/dns.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_id: Id of the group in which the domain will add. If not supplied, then use default group.
        :param pulumi.Input[str] name: Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the dns belongs.
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

            __props__['group_id'] = group_id
            __props__['name'] = name
            __props__['resource_group_id'] = resource_group_id
            __props__['dns_servers'] = None
            __props__['domain_id'] = None
        super(Domain, __self__).__init__(
            'alicloud:dns/domain:Domain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, dns_servers=None, domain_id=None, group_id=None, name=None, resource_group_id=None):
        """
        Get an existing Domain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] dns_servers: A list of the dns server name.
        :param pulumi.Input[str] domain_id: The domain ID.
        :param pulumi.Input[str] group_id: Id of the group in which the domain will add. If not supplied, then use default group.
        :param pulumi.Input[str] name: Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the dns belongs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dns_servers"] = dns_servers
        __props__["domain_id"] = domain_id
        __props__["group_id"] = group_id
        __props__["name"] = name
        __props__["resource_group_id"] = resource_group_id
        return Domain(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

