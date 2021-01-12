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

__all__ = ['Listener']


class Listener(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accelerator_id: Optional[pulumi.Input[str]] = None,
                 certificates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ListenerCertificateArgs']]]]] = None,
                 client_affinity: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ListenerPortRangeArgs']]]]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 proxy_protocol: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Global Accelerator (GA) Listener resource.

        For information about Global Accelerator (GA) Listener and how to use it, see [What is Listener](https://help.aliyun.com/document_detail/153253.html).

        > **NOTE:** Available in v1.111.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_accelerator = alicloud.ga.Accelerator("exampleAccelerator",
            auto_use_coupon=True,
            duration=1,
            spec="1")
        example_listener = alicloud.ga.Listener("exampleListener",
            accelerator_id="alicloud_ga_accelerator.example.id",
            port_ranges=[alicloud.ga.ListenerPortRangeArgs(
                from_port=60,
                to_port=70,
            )])
        ```

        ## Import

        Ga Listener can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ga/listener:Listener example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accelerator_id: The accelerator id.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ListenerCertificateArgs']]]] certificates: The certificates of the listener.
        :param pulumi.Input[str] client_affinity: The clientAffinity of the listener. Default value is `NONE`. Valid values:
               `NONE`: client affinity is not maintained, that is, connection requests from the same client cannot always be directed to the same terminal node.
               `SOURCE_IP`: maintain client affinity. When a client accesses a stateful application, all requests from the same client can be directed to the same terminal node, regardless of the source port and protocol.
        :param pulumi.Input[str] description: The description of the listener.
        :param pulumi.Input[str] name: The name of the listener. The length of the name is 2-128 characters. It starts with uppercase and lowercase letters or Chinese characters. It can contain numbers and underscores and dashes.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ListenerPortRangeArgs']]]] port_ranges: The portRanges of the listener.
        :param pulumi.Input[str] protocol: Type of network transport protocol monitored. Default value is `TCP`. Valid values: `TCP`, `UDP`.
        :param pulumi.Input[bool] proxy_protocol: The proxy protocol of the listener.
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

            if accelerator_id is None:
                raise TypeError("Missing required property 'accelerator_id'")
            __props__['accelerator_id'] = accelerator_id
            __props__['certificates'] = certificates
            __props__['client_affinity'] = client_affinity
            __props__['description'] = description
            __props__['name'] = name
            if port_ranges is None:
                raise TypeError("Missing required property 'port_ranges'")
            __props__['port_ranges'] = port_ranges
            __props__['protocol'] = protocol
            __props__['proxy_protocol'] = proxy_protocol
            __props__['status'] = None
        super(Listener, __self__).__init__(
            'alicloud:ga/listener:Listener',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accelerator_id: Optional[pulumi.Input[str]] = None,
            certificates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ListenerCertificateArgs']]]]] = None,
            client_affinity: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ListenerPortRangeArgs']]]]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            proxy_protocol: Optional[pulumi.Input[bool]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'Listener':
        """
        Get an existing Listener resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accelerator_id: The accelerator id.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ListenerCertificateArgs']]]] certificates: The certificates of the listener.
        :param pulumi.Input[str] client_affinity: The clientAffinity of the listener. Default value is `NONE`. Valid values:
               `NONE`: client affinity is not maintained, that is, connection requests from the same client cannot always be directed to the same terminal node.
               `SOURCE_IP`: maintain client affinity. When a client accesses a stateful application, all requests from the same client can be directed to the same terminal node, regardless of the source port and protocol.
        :param pulumi.Input[str] description: The description of the listener.
        :param pulumi.Input[str] name: The name of the listener. The length of the name is 2-128 characters. It starts with uppercase and lowercase letters or Chinese characters. It can contain numbers and underscores and dashes.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ListenerPortRangeArgs']]]] port_ranges: The portRanges of the listener.
        :param pulumi.Input[str] protocol: Type of network transport protocol monitored. Default value is `TCP`. Valid values: `TCP`, `UDP`.
        :param pulumi.Input[bool] proxy_protocol: The proxy protocol of the listener.
        :param pulumi.Input[str] status: The status of the listener.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accelerator_id"] = accelerator_id
        __props__["certificates"] = certificates
        __props__["client_affinity"] = client_affinity
        __props__["description"] = description
        __props__["name"] = name
        __props__["port_ranges"] = port_ranges
        __props__["protocol"] = protocol
        __props__["proxy_protocol"] = proxy_protocol
        __props__["status"] = status
        return Listener(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceleratorId")
    def accelerator_id(self) -> pulumi.Output[str]:
        """
        The accelerator id.
        """
        return pulumi.get(self, "accelerator_id")

    @property
    @pulumi.getter
    def certificates(self) -> pulumi.Output[Optional[Sequence['outputs.ListenerCertificate']]]:
        """
        The certificates of the listener.
        """
        return pulumi.get(self, "certificates")

    @property
    @pulumi.getter(name="clientAffinity")
    def client_affinity(self) -> pulumi.Output[Optional[str]]:
        """
        The clientAffinity of the listener. Default value is `NONE`. Valid values:
        `NONE`: client affinity is not maintained, that is, connection requests from the same client cannot always be directed to the same terminal node.
        `SOURCE_IP`: maintain client affinity. When a client accesses a stateful application, all requests from the same client can be directed to the same terminal node, regardless of the source port and protocol.
        """
        return pulumi.get(self, "client_affinity")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the listener.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the listener. The length of the name is 2-128 characters. It starts with uppercase and lowercase letters or Chinese characters. It can contain numbers and underscores and dashes.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="portRanges")
    def port_ranges(self) -> pulumi.Output[Sequence['outputs.ListenerPortRange']]:
        """
        The portRanges of the listener.
        """
        return pulumi.get(self, "port_ranges")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[Optional[str]]:
        """
        Type of network transport protocol monitored. Default value is `TCP`. Valid values: `TCP`, `UDP`.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="proxyProtocol")
    def proxy_protocol(self) -> pulumi.Output[Optional[bool]]:
        """
        The proxy protocol of the listener.
        """
        return pulumi.get(self, "proxy_protocol")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the listener.
        """
        return pulumi.get(self, "status")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

