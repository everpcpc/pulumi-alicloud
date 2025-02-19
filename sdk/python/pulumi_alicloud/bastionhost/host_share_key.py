# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['HostShareKeyArgs', 'HostShareKey']

@pulumi.input_type
class HostShareKeyArgs:
    def __init__(__self__, *,
                 host_share_key_name: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 private_key: pulumi.Input[str],
                 pass_phrase: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a HostShareKey resource.
        :param pulumi.Input[str] host_share_key_name: The name of the host shared key to be added. The name can be a maximum of 128 characters in length.
        :param pulumi.Input[str] instance_id: The ID of the Bastion instance.
        :param pulumi.Input[str] private_key: The private key. The value is a Base64-encoded string.
        :param pulumi.Input[str] pass_phrase: The password of the private key. The value is a Base64-encoded string.
        """
        pulumi.set(__self__, "host_share_key_name", host_share_key_name)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "private_key", private_key)
        if pass_phrase is not None:
            pulumi.set(__self__, "pass_phrase", pass_phrase)

    @property
    @pulumi.getter(name="hostShareKeyName")
    def host_share_key_name(self) -> pulumi.Input[str]:
        """
        The name of the host shared key to be added. The name can be a maximum of 128 characters in length.
        """
        return pulumi.get(self, "host_share_key_name")

    @host_share_key_name.setter
    def host_share_key_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "host_share_key_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the Bastion instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Input[str]:
        """
        The private key. The value is a Base64-encoded string.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="passPhrase")
    def pass_phrase(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the private key. The value is a Base64-encoded string.
        """
        return pulumi.get(self, "pass_phrase")

    @pass_phrase.setter
    def pass_phrase(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pass_phrase", value)


@pulumi.input_type
class _HostShareKeyState:
    def __init__(__self__, *,
                 host_share_key_id: Optional[pulumi.Input[str]] = None,
                 host_share_key_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 pass_phrase: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 private_key_finger_print: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering HostShareKey resources.
        :param pulumi.Input[str] host_share_key_id: The first ID of the resource.
        :param pulumi.Input[str] host_share_key_name: The name of the host shared key to be added. The name can be a maximum of 128 characters in length.
        :param pulumi.Input[str] instance_id: The ID of the Bastion instance.
        :param pulumi.Input[str] pass_phrase: The password of the private key. The value is a Base64-encoded string.
        :param pulumi.Input[str] private_key: The private key. The value is a Base64-encoded string.
        :param pulumi.Input[str] private_key_finger_print: The fingerprint of the private key.
        """
        if host_share_key_id is not None:
            pulumi.set(__self__, "host_share_key_id", host_share_key_id)
        if host_share_key_name is not None:
            pulumi.set(__self__, "host_share_key_name", host_share_key_name)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if pass_phrase is not None:
            pulumi.set(__self__, "pass_phrase", pass_phrase)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if private_key_finger_print is not None:
            pulumi.set(__self__, "private_key_finger_print", private_key_finger_print)

    @property
    @pulumi.getter(name="hostShareKeyId")
    def host_share_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The first ID of the resource.
        """
        return pulumi.get(self, "host_share_key_id")

    @host_share_key_id.setter
    def host_share_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_share_key_id", value)

    @property
    @pulumi.getter(name="hostShareKeyName")
    def host_share_key_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the host shared key to be added. The name can be a maximum of 128 characters in length.
        """
        return pulumi.get(self, "host_share_key_name")

    @host_share_key_name.setter
    def host_share_key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_share_key_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Bastion instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="passPhrase")
    def pass_phrase(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the private key. The value is a Base64-encoded string.
        """
        return pulumi.get(self, "pass_phrase")

    @pass_phrase.setter
    def pass_phrase(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pass_phrase", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        The private key. The value is a Base64-encoded string.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="privateKeyFingerPrint")
    def private_key_finger_print(self) -> Optional[pulumi.Input[str]]:
        """
        The fingerprint of the private key.
        """
        return pulumi.get(self, "private_key_finger_print")

    @private_key_finger_print.setter
    def private_key_finger_print(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_finger_print", value)


class HostShareKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 host_share_key_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 pass_phrase: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Bastion Host Share Key resource.

        For information about Bastion Host Host Share Key and how to use it, see [What is Host Share Key](https://www.alibabacloud.com/help/en/bastion-host/latest/createhostsharekey).

        > **NOTE:** Available in v1.165.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_instances = alicloud.bastionhost.get_instances()
        default_host_share_key = alicloud.bastionhost.HostShareKey("defaultHostShareKey",
            host_share_key_name="example_name",
            instance_id=default_instances.instances[0].id,
            pass_phrase="example_value",
            private_key="example_value")
        ```

        ## Import

        Bastion Host Share Key can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:bastionhost/hostShareKey:HostShareKey example <instance_id>:<host_share_key_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host_share_key_name: The name of the host shared key to be added. The name can be a maximum of 128 characters in length.
        :param pulumi.Input[str] instance_id: The ID of the Bastion instance.
        :param pulumi.Input[str] pass_phrase: The password of the private key. The value is a Base64-encoded string.
        :param pulumi.Input[str] private_key: The private key. The value is a Base64-encoded string.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HostShareKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Bastion Host Share Key resource.

        For information about Bastion Host Host Share Key and how to use it, see [What is Host Share Key](https://www.alibabacloud.com/help/en/bastion-host/latest/createhostsharekey).

        > **NOTE:** Available in v1.165.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_instances = alicloud.bastionhost.get_instances()
        default_host_share_key = alicloud.bastionhost.HostShareKey("defaultHostShareKey",
            host_share_key_name="example_name",
            instance_id=default_instances.instances[0].id,
            pass_phrase="example_value",
            private_key="example_value")
        ```

        ## Import

        Bastion Host Share Key can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:bastionhost/hostShareKey:HostShareKey example <instance_id>:<host_share_key_id>
        ```

        :param str resource_name: The name of the resource.
        :param HostShareKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HostShareKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 host_share_key_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 pass_phrase: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HostShareKeyArgs.__new__(HostShareKeyArgs)

            if host_share_key_name is None and not opts.urn:
                raise TypeError("Missing required property 'host_share_key_name'")
            __props__.__dict__["host_share_key_name"] = host_share_key_name
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["pass_phrase"] = None if pass_phrase is None else pulumi.Output.secret(pass_phrase)
            if private_key is None and not opts.urn:
                raise TypeError("Missing required property 'private_key'")
            __props__.__dict__["private_key"] = None if private_key is None else pulumi.Output.secret(private_key)
            __props__.__dict__["host_share_key_id"] = None
            __props__.__dict__["private_key_finger_print"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["passPhrase", "privateKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(HostShareKey, __self__).__init__(
            'alicloud:bastionhost/hostShareKey:HostShareKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            host_share_key_id: Optional[pulumi.Input[str]] = None,
            host_share_key_name: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            pass_phrase: Optional[pulumi.Input[str]] = None,
            private_key: Optional[pulumi.Input[str]] = None,
            private_key_finger_print: Optional[pulumi.Input[str]] = None) -> 'HostShareKey':
        """
        Get an existing HostShareKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host_share_key_id: The first ID of the resource.
        :param pulumi.Input[str] host_share_key_name: The name of the host shared key to be added. The name can be a maximum of 128 characters in length.
        :param pulumi.Input[str] instance_id: The ID of the Bastion instance.
        :param pulumi.Input[str] pass_phrase: The password of the private key. The value is a Base64-encoded string.
        :param pulumi.Input[str] private_key: The private key. The value is a Base64-encoded string.
        :param pulumi.Input[str] private_key_finger_print: The fingerprint of the private key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HostShareKeyState.__new__(_HostShareKeyState)

        __props__.__dict__["host_share_key_id"] = host_share_key_id
        __props__.__dict__["host_share_key_name"] = host_share_key_name
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["pass_phrase"] = pass_phrase
        __props__.__dict__["private_key"] = private_key
        __props__.__dict__["private_key_finger_print"] = private_key_finger_print
        return HostShareKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="hostShareKeyId")
    def host_share_key_id(self) -> pulumi.Output[str]:
        """
        The first ID of the resource.
        """
        return pulumi.get(self, "host_share_key_id")

    @property
    @pulumi.getter(name="hostShareKeyName")
    def host_share_key_name(self) -> pulumi.Output[str]:
        """
        The name of the host shared key to be added. The name can be a maximum of 128 characters in length.
        """
        return pulumi.get(self, "host_share_key_name")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the Bastion instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="passPhrase")
    def pass_phrase(self) -> pulumi.Output[Optional[str]]:
        """
        The password of the private key. The value is a Base64-encoded string.
        """
        return pulumi.get(self, "pass_phrase")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[str]:
        """
        The private key. The value is a Base64-encoded string.
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="privateKeyFingerPrint")
    def private_key_finger_print(self) -> pulumi.Output[str]:
        """
        The fingerprint of the private key.
        """
        return pulumi.get(self, "private_key_finger_print")

