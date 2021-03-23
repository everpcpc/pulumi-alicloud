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

__all__ = ['DomainNew']


class DomainNew(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cdn_type: Optional[pulumi.Input[str]] = None,
                 certificate_config: Optional[pulumi.Input[pulumi.InputType['DomainNewCertificateConfigArgs']]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainNewSourceArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a CDN Accelerated Domain resource. This resource is based on CDN's new version OpenAPI.

        For information about Cdn Domain New and how to use it, see [Add a domain](https://www.alibabacloud.com/help/doc-detail/91176.html).

        > **NOTE:** Available in v1.34.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        # Create a new Domain.
        domain = alicloud.cdn.DomainNew("domain",
            cdn_type="web",
            domain_name="terraform.test.com",
            scope="overseas",
            sources=[alicloud.cdn.DomainNewSourceArgs(
                content="1.1.1.1",
                port=80,
                priority=20,
                type="ipaddr",
                weight=10,
            )])
        ```

        ## Import

        CDN domain can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cdn/domainNew:DomainNew example xxxx.com
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cdn_type: Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
        :param pulumi.Input[pulumi.InputType['DomainNewCertificateConfigArgs']] certificate_config: Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
        :param pulumi.Input[str] domain_name: Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        :param pulumi.Input[str] resource_group_id: Resource group ID.
        :param pulumi.Input[str] scope: Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainNewSourceArgs']]]] sources: The source address list of the accelerated domain. Defaults to null. See Block Sources.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
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

            if cdn_type is None and not opts.urn:
                raise TypeError("Missing required property 'cdn_type'")
            __props__['cdn_type'] = cdn_type
            __props__['certificate_config'] = certificate_config
            if domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name'")
            __props__['domain_name'] = domain_name
            __props__['resource_group_id'] = resource_group_id
            __props__['scope'] = scope
            if sources is None and not opts.urn:
                raise TypeError("Missing required property 'sources'")
            __props__['sources'] = sources
            __props__['tags'] = tags
            __props__['cname'] = None
        super(DomainNew, __self__).__init__(
            'alicloud:cdn/domainNew:DomainNew',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cdn_type: Optional[pulumi.Input[str]] = None,
            certificate_config: Optional[pulumi.Input[pulumi.InputType['DomainNewCertificateConfigArgs']]] = None,
            cname: Optional[pulumi.Input[str]] = None,
            domain_name: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainNewSourceArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'DomainNew':
        """
        Get an existing DomainNew resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cdn_type: Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
        :param pulumi.Input[pulumi.InputType['DomainNewCertificateConfigArgs']] certificate_config: Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
        :param pulumi.Input[str] cname: (Available in v1.90.0+) The CNAME of the CDN domain.
        :param pulumi.Input[str] domain_name: Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        :param pulumi.Input[str] resource_group_id: Resource group ID.
        :param pulumi.Input[str] scope: Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainNewSourceArgs']]]] sources: The source address list of the accelerated domain. Defaults to null. See Block Sources.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cdn_type"] = cdn_type
        __props__["certificate_config"] = certificate_config
        __props__["cname"] = cname
        __props__["domain_name"] = domain_name
        __props__["resource_group_id"] = resource_group_id
        __props__["scope"] = scope
        __props__["sources"] = sources
        __props__["tags"] = tags
        return DomainNew(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cdnType")
    def cdn_type(self) -> pulumi.Output[str]:
        """
        Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
        """
        return pulumi.get(self, "cdn_type")

    @property
    @pulumi.getter(name="certificateConfig")
    def certificate_config(self) -> pulumi.Output['outputs.DomainNewCertificateConfig']:
        """
        Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
        """
        return pulumi.get(self, "certificate_config")

    @property
    @pulumi.getter
    def cname(self) -> pulumi.Output[str]:
        """
        (Available in v1.90.0+) The CNAME of the CDN domain.
        """
        return pulumi.get(self, "cname")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        Resource group ID.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def sources(self) -> pulumi.Output[Sequence['outputs.DomainNewSource']]:
        """
        The source address list of the accelerated domain. Defaults to null. See Block Sources.
        """
        return pulumi.get(self, "sources")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

