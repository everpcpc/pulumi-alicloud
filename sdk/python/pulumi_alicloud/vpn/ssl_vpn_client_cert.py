# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SslVpnClientCertArgs', 'SslVpnClientCert']

@pulumi.input_type
class SslVpnClientCertArgs:
    def __init__(__self__, *,
                 ssl_vpn_server_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SslVpnClientCert resource.
        :param pulumi.Input[str] ssl_vpn_server_id: The ID of the SSL-VPN server.
        :param pulumi.Input[str] name: The name of the client certificate.
        """
        pulumi.set(__self__, "ssl_vpn_server_id", ssl_vpn_server_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="sslVpnServerId")
    def ssl_vpn_server_id(self) -> pulumi.Input[str]:
        """
        The ID of the SSL-VPN server.
        """
        return pulumi.get(self, "ssl_vpn_server_id")

    @ssl_vpn_server_id.setter
    def ssl_vpn_server_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ssl_vpn_server_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the client certificate.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _SslVpnClientCertState:
    def __init__(__self__, *,
                 ca_cert: Optional[pulumi.Input[str]] = None,
                 client_cert: Optional[pulumi.Input[str]] = None,
                 client_config: Optional[pulumi.Input[str]] = None,
                 client_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ssl_vpn_server_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SslVpnClientCert resources.
        :param pulumi.Input[str] ca_cert: The client ca cert.
        :param pulumi.Input[str] client_cert: The client cert.
        :param pulumi.Input[str] client_config: The vpn client config.
        :param pulumi.Input[str] client_key: The client key.
        :param pulumi.Input[str] name: The name of the client certificate.
        :param pulumi.Input[str] ssl_vpn_server_id: The ID of the SSL-VPN server.
        :param pulumi.Input[str] status: The status of the client certificate.
        """
        if ca_cert is not None:
            pulumi.set(__self__, "ca_cert", ca_cert)
        if client_cert is not None:
            pulumi.set(__self__, "client_cert", client_cert)
        if client_config is not None:
            pulumi.set(__self__, "client_config", client_config)
        if client_key is not None:
            pulumi.set(__self__, "client_key", client_key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ssl_vpn_server_id is not None:
            pulumi.set(__self__, "ssl_vpn_server_id", ssl_vpn_server_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="caCert")
    def ca_cert(self) -> Optional[pulumi.Input[str]]:
        """
        The client ca cert.
        """
        return pulumi.get(self, "ca_cert")

    @ca_cert.setter
    def ca_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_cert", value)

    @property
    @pulumi.getter(name="clientCert")
    def client_cert(self) -> Optional[pulumi.Input[str]]:
        """
        The client cert.
        """
        return pulumi.get(self, "client_cert")

    @client_cert.setter
    def client_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_cert", value)

    @property
    @pulumi.getter(name="clientConfig")
    def client_config(self) -> Optional[pulumi.Input[str]]:
        """
        The vpn client config.
        """
        return pulumi.get(self, "client_config")

    @client_config.setter
    def client_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_config", value)

    @property
    @pulumi.getter(name="clientKey")
    def client_key(self) -> Optional[pulumi.Input[str]]:
        """
        The client key.
        """
        return pulumi.get(self, "client_key")

    @client_key.setter
    def client_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the client certificate.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sslVpnServerId")
    def ssl_vpn_server_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the SSL-VPN server.
        """
        return pulumi.get(self, "ssl_vpn_server_id")

    @ssl_vpn_server_id.setter
    def ssl_vpn_server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_vpn_server_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the client certificate.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class SslVpnClientCert(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ssl_vpn_server_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        SSL-VPN client certificates can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpn/sslVpnClientCert:SslVpnClientCert example vsc-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the client certificate.
        :param pulumi.Input[str] ssl_vpn_server_id: The ID of the SSL-VPN server.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SslVpnClientCertArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        SSL-VPN client certificates can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpn/sslVpnClientCert:SslVpnClientCert example vsc-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param SslVpnClientCertArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SslVpnClientCertArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ssl_vpn_server_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = SslVpnClientCertArgs.__new__(SslVpnClientCertArgs)

            __props__.__dict__["name"] = name
            if ssl_vpn_server_id is None and not opts.urn:
                raise TypeError("Missing required property 'ssl_vpn_server_id'")
            __props__.__dict__["ssl_vpn_server_id"] = ssl_vpn_server_id
            __props__.__dict__["ca_cert"] = None
            __props__.__dict__["client_cert"] = None
            __props__.__dict__["client_config"] = None
            __props__.__dict__["client_key"] = None
            __props__.__dict__["status"] = None
        super(SslVpnClientCert, __self__).__init__(
            'alicloud:vpn/sslVpnClientCert:SslVpnClientCert',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ca_cert: Optional[pulumi.Input[str]] = None,
            client_cert: Optional[pulumi.Input[str]] = None,
            client_config: Optional[pulumi.Input[str]] = None,
            client_key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            ssl_vpn_server_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'SslVpnClientCert':
        """
        Get an existing SslVpnClientCert resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ca_cert: The client ca cert.
        :param pulumi.Input[str] client_cert: The client cert.
        :param pulumi.Input[str] client_config: The vpn client config.
        :param pulumi.Input[str] client_key: The client key.
        :param pulumi.Input[str] name: The name of the client certificate.
        :param pulumi.Input[str] ssl_vpn_server_id: The ID of the SSL-VPN server.
        :param pulumi.Input[str] status: The status of the client certificate.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SslVpnClientCertState.__new__(_SslVpnClientCertState)

        __props__.__dict__["ca_cert"] = ca_cert
        __props__.__dict__["client_cert"] = client_cert
        __props__.__dict__["client_config"] = client_config
        __props__.__dict__["client_key"] = client_key
        __props__.__dict__["name"] = name
        __props__.__dict__["ssl_vpn_server_id"] = ssl_vpn_server_id
        __props__.__dict__["status"] = status
        return SslVpnClientCert(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="caCert")
    def ca_cert(self) -> pulumi.Output[str]:
        """
        The client ca cert.
        """
        return pulumi.get(self, "ca_cert")

    @property
    @pulumi.getter(name="clientCert")
    def client_cert(self) -> pulumi.Output[str]:
        """
        The client cert.
        """
        return pulumi.get(self, "client_cert")

    @property
    @pulumi.getter(name="clientConfig")
    def client_config(self) -> pulumi.Output[str]:
        """
        The vpn client config.
        """
        return pulumi.get(self, "client_config")

    @property
    @pulumi.getter(name="clientKey")
    def client_key(self) -> pulumi.Output[str]:
        """
        The client key.
        """
        return pulumi.get(self, "client_key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the client certificate.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sslVpnServerId")
    def ssl_vpn_server_id(self) -> pulumi.Output[str]:
        """
        The ID of the SSL-VPN server.
        """
        return pulumi.get(self, "ssl_vpn_server_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the client certificate.
        """
        return pulumi.get(self, "status")

