# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class DomainNew(pulumi.CustomResource):
    cdn_type: pulumi.Output[str]
    """
    Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
    """
    certificate_config: pulumi.Output[dict]
    """
    Certificate config of the accelerated domain. It's a list and consist of at most 1 item.

      * `certName` (`str`) - The SSL certificate name.
      * `certType` (`str`) - The SSL certificate type, can be "upload", "cas" and "free".
      * `forceSet` (`str`) - Set `1` to ignore the repeated verification for certificate name, and cover the information of the origin certificate (with the same name). Set `0` to work the verification.
      * `private_key` (`str`) - The SSL private key. This is required if `server_certificate_status` is `on`
      * `server_certificate` (`str`) - The SSL server certificate string. This is required if `server_certificate_status` is `on`
      * `serverCertificateStatus` (`str`) - This parameter indicates whether or not enable https. Valid values are `on` and `off`. Default value is `on`.
    """
    domain_name: pulumi.Output[str]
    """
    Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
    """
    resource_group_id: pulumi.Output[str]
    """
    Resource group ID.
    """
    scope: pulumi.Output[str]
    """
    Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
    """
    sources: pulumi.Output[dict]
    """
    The source address list of the accelerated domain. Defaults to null. See Block Sources.

      * `content` (`str`) - The adress of source. Valid values can be ip or doaminName. Each item's `content` can not be repeated.
      * `port` (`float`) - The port of source. Valid values are `443` and `80`. Default value is `80`.
      * `priority` (`float`) - Priority of the source. Valid values are `0` and `100`. Default value is `20`.
      * `type` (`str`) - The type of the source. Valid values are `ipaddr`, `domain` and `oss`.
      * `weight` (`float`) - Weight of the source. Valid values are from `0` to `100`. Default value is `10`, but if type is `ipaddr`, the value can only be `10`. 
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, cdn_type=None, certificate_config=None, domain_name=None, resource_group_id=None, scope=None, sources=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a CDN Accelerated Domain resource. This resource is based on CDN's new version OpenAPI.

        For information about Cdn Domain New and how to use it, see [Add a domain](https://www.alibabacloud.com/help/doc-detail/91176.html).

        > **NOTE:** Available in v1.34.0+.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cdn_type: Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
        :param pulumi.Input[dict] certificate_config: Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
        :param pulumi.Input[str] domain_name: Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        :param pulumi.Input[str] resource_group_id: Resource group ID.
        :param pulumi.Input[str] scope: Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
        :param pulumi.Input[dict] sources: The source address list of the accelerated domain. Defaults to null. See Block Sources.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **certificate_config** object supports the following:

          * `certName` (`pulumi.Input[str]`) - The SSL certificate name.
          * `certType` (`pulumi.Input[str]`) - The SSL certificate type, can be "upload", "cas" and "free".
          * `forceSet` (`pulumi.Input[str]`) - Set `1` to ignore the repeated verification for certificate name, and cover the information of the origin certificate (with the same name). Set `0` to work the verification.
          * `private_key` (`pulumi.Input[str]`) - The SSL private key. This is required if `server_certificate_status` is `on`
          * `server_certificate` (`pulumi.Input[str]`) - The SSL server certificate string. This is required if `server_certificate_status` is `on`
          * `serverCertificateStatus` (`pulumi.Input[str]`) - This parameter indicates whether or not enable https. Valid values are `on` and `off`. Default value is `on`.

        The **sources** object supports the following:

          * `content` (`pulumi.Input[str]`) - The adress of source. Valid values can be ip or doaminName. Each item's `content` can not be repeated.
          * `port` (`pulumi.Input[float]`) - The port of source. Valid values are `443` and `80`. Default value is `80`.
          * `priority` (`pulumi.Input[float]`) - Priority of the source. Valid values are `0` and `100`. Default value is `20`.
          * `type` (`pulumi.Input[str]`) - The type of the source. Valid values are `ipaddr`, `domain` and `oss`.
          * `weight` (`pulumi.Input[float]`) - Weight of the source. Valid values are from `0` to `100`. Default value is `10`, but if type is `ipaddr`, the value can only be `10`. 
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

            if cdn_type is None:
                raise TypeError("Missing required property 'cdn_type'")
            __props__['cdn_type'] = cdn_type
            __props__['certificate_config'] = certificate_config
            if domain_name is None:
                raise TypeError("Missing required property 'domain_name'")
            __props__['domain_name'] = domain_name
            __props__['resource_group_id'] = resource_group_id
            __props__['scope'] = scope
            if sources is None:
                raise TypeError("Missing required property 'sources'")
            __props__['sources'] = sources
            __props__['tags'] = tags
        super(DomainNew, __self__).__init__(
            'alicloud:cdn/domainNew:DomainNew',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, cdn_type=None, certificate_config=None, domain_name=None, resource_group_id=None, scope=None, sources=None, tags=None):
        """
        Get an existing DomainNew resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cdn_type: Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
        :param pulumi.Input[dict] certificate_config: Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
        :param pulumi.Input[str] domain_name: Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        :param pulumi.Input[str] resource_group_id: Resource group ID.
        :param pulumi.Input[str] scope: Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
        :param pulumi.Input[dict] sources: The source address list of the accelerated domain. Defaults to null. See Block Sources.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **certificate_config** object supports the following:

          * `certName` (`pulumi.Input[str]`) - The SSL certificate name.
          * `certType` (`pulumi.Input[str]`) - The SSL certificate type, can be "upload", "cas" and "free".
          * `forceSet` (`pulumi.Input[str]`) - Set `1` to ignore the repeated verification for certificate name, and cover the information of the origin certificate (with the same name). Set `0` to work the verification.
          * `private_key` (`pulumi.Input[str]`) - The SSL private key. This is required if `server_certificate_status` is `on`
          * `server_certificate` (`pulumi.Input[str]`) - The SSL server certificate string. This is required if `server_certificate_status` is `on`
          * `serverCertificateStatus` (`pulumi.Input[str]`) - This parameter indicates whether or not enable https. Valid values are `on` and `off`. Default value is `on`.

        The **sources** object supports the following:

          * `content` (`pulumi.Input[str]`) - The adress of source. Valid values can be ip or doaminName. Each item's `content` can not be repeated.
          * `port` (`pulumi.Input[float]`) - The port of source. Valid values are `443` and `80`. Default value is `80`.
          * `priority` (`pulumi.Input[float]`) - Priority of the source. Valid values are `0` and `100`. Default value is `20`.
          * `type` (`pulumi.Input[str]`) - The type of the source. Valid values are `ipaddr`, `domain` and `oss`.
          * `weight` (`pulumi.Input[float]`) - Weight of the source. Valid values are from `0` to `100`. Default value is `10`, but if type is `ipaddr`, the value can only be `10`. 
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cdn_type"] = cdn_type
        __props__["certificate_config"] = certificate_config
        __props__["domain_name"] = domain_name
        __props__["resource_group_id"] = resource_group_id
        __props__["scope"] = scope
        __props__["sources"] = sources
        __props__["tags"] = tags
        return DomainNew(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

