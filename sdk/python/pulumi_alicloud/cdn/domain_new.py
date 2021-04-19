# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DomainNewArgs', 'DomainNew']

@pulumi.input_type
class DomainNewArgs:
    def __init__(__self__, *,
                 cdn_type: pulumi.Input[str],
                 domain_name: pulumi.Input[str],
                 sources: pulumi.Input[Sequence[pulumi.Input['DomainNewSourceArgs']]],
                 certificate_config: Optional[pulumi.Input['DomainNewCertificateConfigArgs']] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a DomainNew resource.
        :param pulumi.Input[str] cdn_type: Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
        :param pulumi.Input[str] domain_name: Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        :param pulumi.Input[Sequence[pulumi.Input['DomainNewSourceArgs']]] sources: The source address list of the accelerated domain. Defaults to null. See Block Sources.
        :param pulumi.Input['DomainNewCertificateConfigArgs'] certificate_config: Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
        :param pulumi.Input[str] resource_group_id: Resource group ID.
        :param pulumi.Input[str] scope: Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "cdn_type", cdn_type)
        pulumi.set(__self__, "domain_name", domain_name)
        pulumi.set(__self__, "sources", sources)
        if certificate_config is not None:
            pulumi.set(__self__, "certificate_config", certificate_config)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="cdnType")
    def cdn_type(self) -> pulumi.Input[str]:
        """
        Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
        """
        return pulumi.get(self, "cdn_type")

    @cdn_type.setter
    def cdn_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "cdn_type", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Input[str]:
        """
        Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter
    def sources(self) -> pulumi.Input[Sequence[pulumi.Input['DomainNewSourceArgs']]]:
        """
        The source address list of the accelerated domain. Defaults to null. See Block Sources.
        """
        return pulumi.get(self, "sources")

    @sources.setter
    def sources(self, value: pulumi.Input[Sequence[pulumi.Input['DomainNewSourceArgs']]]):
        pulumi.set(self, "sources", value)

    @property
    @pulumi.getter(name="certificateConfig")
    def certificate_config(self) -> Optional[pulumi.Input['DomainNewCertificateConfigArgs']]:
        """
        Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
        """
        return pulumi.get(self, "certificate_config")

    @certificate_config.setter
    def certificate_config(self, value: Optional[pulumi.Input['DomainNewCertificateConfigArgs']]):
        pulumi.set(self, "certificate_config", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource group ID.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _DomainNewState:
    def __init__(__self__, *,
                 cdn_type: Optional[pulumi.Input[str]] = None,
                 certificate_config: Optional[pulumi.Input['DomainNewCertificateConfigArgs']] = None,
                 cname: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input['DomainNewSourceArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering DomainNew resources.
        :param pulumi.Input[str] cdn_type: Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
        :param pulumi.Input['DomainNewCertificateConfigArgs'] certificate_config: Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
        :param pulumi.Input[str] cname: (Available in v1.90.0+) The CNAME of the CDN domain.
        :param pulumi.Input[str] domain_name: Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        :param pulumi.Input[str] resource_group_id: Resource group ID.
        :param pulumi.Input[str] scope: Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
        :param pulumi.Input[Sequence[pulumi.Input['DomainNewSourceArgs']]] sources: The source address list of the accelerated domain. Defaults to null. See Block Sources.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        if cdn_type is not None:
            pulumi.set(__self__, "cdn_type", cdn_type)
        if certificate_config is not None:
            pulumi.set(__self__, "certificate_config", certificate_config)
        if cname is not None:
            pulumi.set(__self__, "cname", cname)
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if sources is not None:
            pulumi.set(__self__, "sources", sources)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="cdnType")
    def cdn_type(self) -> Optional[pulumi.Input[str]]:
        """
        Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
        """
        return pulumi.get(self, "cdn_type")

    @cdn_type.setter
    def cdn_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cdn_type", value)

    @property
    @pulumi.getter(name="certificateConfig")
    def certificate_config(self) -> Optional[pulumi.Input['DomainNewCertificateConfigArgs']]:
        """
        Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
        """
        return pulumi.get(self, "certificate_config")

    @certificate_config.setter
    def certificate_config(self, value: Optional[pulumi.Input['DomainNewCertificateConfigArgs']]):
        pulumi.set(self, "certificate_config", value)

    @property
    @pulumi.getter
    def cname(self) -> Optional[pulumi.Input[str]]:
        """
        (Available in v1.90.0+) The CNAME of the CDN domain.
        """
        return pulumi.get(self, "cname")

    @cname.setter
    def cname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cname", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource group ID.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def sources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DomainNewSourceArgs']]]]:
        """
        The source address list of the accelerated domain. Defaults to null. See Block Sources.
        """
        return pulumi.get(self, "sources")

    @sources.setter
    def sources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DomainNewSourceArgs']]]]):
        pulumi.set(self, "sources", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class DomainNew(pulumi.CustomResource):
    @overload
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
                 __props__=None):
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
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainNewArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
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
        :param DomainNewArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainNewArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cdn_type: Optional[pulumi.Input[str]] = None,
                 certificate_config: Optional[pulumi.Input[pulumi.InputType['DomainNewCertificateConfigArgs']]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainNewSourceArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DomainNewArgs.__new__(DomainNewArgs)

            if cdn_type is None and not opts.urn:
                raise TypeError("Missing required property 'cdn_type'")
            __props__.__dict__["cdn_type"] = cdn_type
            __props__.__dict__["certificate_config"] = certificate_config
            if domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name'")
            __props__.__dict__["domain_name"] = domain_name
            __props__.__dict__["resource_group_id"] = resource_group_id
            __props__.__dict__["scope"] = scope
            if sources is None and not opts.urn:
                raise TypeError("Missing required property 'sources'")
            __props__.__dict__["sources"] = sources
            __props__.__dict__["tags"] = tags
            __props__.__dict__["cname"] = None
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

        __props__ = _DomainNewState.__new__(_DomainNewState)

        __props__.__dict__["cdn_type"] = cdn_type
        __props__.__dict__["certificate_config"] = certificate_config
        __props__.__dict__["cname"] = cname
        __props__.__dict__["domain_name"] = domain_name
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["scope"] = scope
        __props__.__dict__["sources"] = sources
        __props__.__dict__["tags"] = tags
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

