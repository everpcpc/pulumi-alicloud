# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CaCertificateArgs', 'CaCertificate']

@pulumi.input_type
class CaCertificateArgs:
    def __init__(__self__, *,
                 ca_certificate: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a CaCertificate resource.
        :param pulumi.Input[str] ca_certificate: the content of the CA certificate.
        :param pulumi.Input[str] name: Name of the CA Certificate.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the slb_ca certificate belongs.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "ca_certificate", ca_certificate)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="caCertificate")
    def ca_certificate(self) -> pulumi.Input[str]:
        """
        the content of the CA certificate.
        """
        return pulumi.get(self, "ca_certificate")

    @ca_certificate.setter
    def ca_certificate(self, value: pulumi.Input[str]):
        pulumi.set(self, "ca_certificate", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the CA Certificate.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of resource group which the slb_ca certificate belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

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
class _CaCertificateState:
    def __init__(__self__, *,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering CaCertificate resources.
        :param pulumi.Input[str] ca_certificate: the content of the CA certificate.
        :param pulumi.Input[str] name: Name of the CA Certificate.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the slb_ca certificate belongs.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        if ca_certificate is not None:
            pulumi.set(__self__, "ca_certificate", ca_certificate)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="caCertificate")
    def ca_certificate(self) -> Optional[pulumi.Input[str]]:
        """
        the content of the CA certificate.
        """
        return pulumi.get(self, "ca_certificate")

    @ca_certificate.setter
    def ca_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_certificate", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the CA Certificate.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of resource group which the slb_ca certificate belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

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


class CaCertificate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        A Load Balancer CA Certificate is used by the listener of the protocol https.

        For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).

        For information about CA Certificate and how to use it, see [Configure CA Certificate](https://www.alibabacloud.com/help/doc-detail/85968.htm).

        ## Example Usage

        * using CA certificate content

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        # create a CA certificate
        foo = alicloud.slb.CaCertificate("foo", ca_certificate=\"\"\"-----BEGIN CERTIFICATE-----
        MIIDRjCCAq+gAwIBAgIJAJnI******90EAxEG/bJJyOm5LqoiA=
        -----END CERTIFICATE-----
        \"\"\")
        ```

        * using CA certificate file

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        foo_file = alicloud.slb.CaCertificate("foo-file", ca_certificate=(lambda path: open(path).read())(f"{path['module']}/ca_certificate.pem"))
        ```

        ## Import

        Server Load balancer CA Certificate can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:slb/caCertificate:CaCertificate example abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ca_certificate: the content of the CA certificate.
        :param pulumi.Input[str] name: Name of the CA Certificate.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the slb_ca certificate belongs.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CaCertificateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Load Balancer CA Certificate is used by the listener of the protocol https.

        For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).

        For information about CA Certificate and how to use it, see [Configure CA Certificate](https://www.alibabacloud.com/help/doc-detail/85968.htm).

        ## Example Usage

        * using CA certificate content

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        # create a CA certificate
        foo = alicloud.slb.CaCertificate("foo", ca_certificate=\"\"\"-----BEGIN CERTIFICATE-----
        MIIDRjCCAq+gAwIBAgIJAJnI******90EAxEG/bJJyOm5LqoiA=
        -----END CERTIFICATE-----
        \"\"\")
        ```

        * using CA certificate file

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        foo_file = alicloud.slb.CaCertificate("foo-file", ca_certificate=(lambda path: open(path).read())(f"{path['module']}/ca_certificate.pem"))
        ```

        ## Import

        Server Load balancer CA Certificate can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:slb/caCertificate:CaCertificate example abc123456
        ```

        :param str resource_name: The name of the resource.
        :param CaCertificateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CaCertificateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = CaCertificateArgs.__new__(CaCertificateArgs)

            if ca_certificate is None and not opts.urn:
                raise TypeError("Missing required property 'ca_certificate'")
            __props__.__dict__["ca_certificate"] = ca_certificate
            __props__.__dict__["name"] = name
            __props__.__dict__["resource_group_id"] = resource_group_id
            __props__.__dict__["tags"] = tags
        super(CaCertificate, __self__).__init__(
            'alicloud:slb/caCertificate:CaCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ca_certificate: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'CaCertificate':
        """
        Get an existing CaCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ca_certificate: the content of the CA certificate.
        :param pulumi.Input[str] name: Name of the CA Certificate.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the slb_ca certificate belongs.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CaCertificateState.__new__(_CaCertificateState)

        __props__.__dict__["ca_certificate"] = ca_certificate
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["tags"] = tags
        return CaCertificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="caCertificate")
    def ca_certificate(self) -> pulumi.Output[str]:
        """
        the content of the CA certificate.
        """
        return pulumi.get(self, "ca_certificate")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the CA Certificate.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        The Id of resource group which the slb_ca certificate belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

