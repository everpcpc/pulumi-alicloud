# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['KeyPairArgs', 'KeyPair']

@pulumi.input_type
class KeyPairArgs:
    def __init__(__self__, *,
                 key_file: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 key_name_prefix: Optional[pulumi.Input[str]] = None,
                 key_pair_name: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a KeyPair resource.
        :param pulumi.Input[str] key_file: The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
        :param pulumi.Input[str] key_name: The key pair's name. It is the only in one Alicloud account.
        :param pulumi.Input[str] public_key: You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resource_group_id` is the key pair belongs.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the key pair belongs.
        """
        if key_file is not None:
            pulumi.set(__self__, "key_file", key_file)
        if key_name is not None:
            warnings.warn("""Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.""", DeprecationWarning)
            pulumi.log.warn("""key_name is deprecated: Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.""")
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)
        if key_name_prefix is not None:
            pulumi.set(__self__, "key_name_prefix", key_name_prefix)
        if key_pair_name is not None:
            pulumi.set(__self__, "key_pair_name", key_pair_name)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="keyFile")
    def key_file(self) -> Optional[pulumi.Input[str]]:
        """
        The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
        """
        return pulumi.get(self, "key_file")

    @key_file.setter
    def key_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_file", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[pulumi.Input[str]]:
        """
        The key pair's name. It is the only in one Alicloud account.
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_name", value)

    @property
    @pulumi.getter(name="keyNamePrefix")
    def key_name_prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key_name_prefix")

    @key_name_prefix.setter
    def key_name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_name_prefix", value)

    @property
    @pulumi.getter(name="keyPairName")
    def key_pair_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key_pair_name")

    @key_pair_name.setter
    def key_pair_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_pair_name", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resource_group_id` is the key pair belongs.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of resource group which the key pair belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _KeyPairState:
    def __init__(__self__, *,
                 finger_print: Optional[pulumi.Input[str]] = None,
                 key_file: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 key_name_prefix: Optional[pulumi.Input[str]] = None,
                 key_pair_name: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering KeyPair resources.
        :param pulumi.Input[str] key_file: The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
        :param pulumi.Input[str] key_name: The key pair's name. It is the only in one Alicloud account.
        :param pulumi.Input[str] public_key: You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resource_group_id` is the key pair belongs.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the key pair belongs.
        """
        if finger_print is not None:
            pulumi.set(__self__, "finger_print", finger_print)
        if key_file is not None:
            pulumi.set(__self__, "key_file", key_file)
        if key_name is not None:
            warnings.warn("""Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.""", DeprecationWarning)
            pulumi.log.warn("""key_name is deprecated: Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.""")
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)
        if key_name_prefix is not None:
            pulumi.set(__self__, "key_name_prefix", key_name_prefix)
        if key_pair_name is not None:
            pulumi.set(__self__, "key_pair_name", key_pair_name)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="fingerPrint")
    def finger_print(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "finger_print")

    @finger_print.setter
    def finger_print(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "finger_print", value)

    @property
    @pulumi.getter(name="keyFile")
    def key_file(self) -> Optional[pulumi.Input[str]]:
        """
        The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
        """
        return pulumi.get(self, "key_file")

    @key_file.setter
    def key_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_file", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[pulumi.Input[str]]:
        """
        The key pair's name. It is the only in one Alicloud account.
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_name", value)

    @property
    @pulumi.getter(name="keyNamePrefix")
    def key_name_prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key_name_prefix")

    @key_name_prefix.setter
    def key_name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_name_prefix", value)

    @property
    @pulumi.getter(name="keyPairName")
    def key_pair_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key_pair_name")

    @key_pair_name.setter
    def key_pair_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_pair_name", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resource_group_id` is the key pair belongs.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of resource group which the key pair belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class KeyPair(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_file: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 key_name_prefix: Optional[pulumi.Input[str]] = None,
                 key_pair_name: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        Key pair can be imported using the name, e.g.

        ```sh
         $ pulumi import alicloud:ecs/keyPair:KeyPair example my_public_key
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_file: The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
        :param pulumi.Input[str] key_name: The key pair's name. It is the only in one Alicloud account.
        :param pulumi.Input[str] public_key: You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resource_group_id` is the key pair belongs.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the key pair belongs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[KeyPairArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Key pair can be imported using the name, e.g.

        ```sh
         $ pulumi import alicloud:ecs/keyPair:KeyPair example my_public_key
        ```

        :param str resource_name: The name of the resource.
        :param KeyPairArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeyPairArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_file: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 key_name_prefix: Optional[pulumi.Input[str]] = None,
                 key_pair_name: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            __props__ = KeyPairArgs.__new__(KeyPairArgs)

            __props__.__dict__["key_file"] = key_file
            if key_name is not None and not opts.urn:
                warnings.warn("""Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.""", DeprecationWarning)
                pulumi.log.warn("""key_name is deprecated: Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.""")
            __props__.__dict__["key_name"] = key_name
            __props__.__dict__["key_name_prefix"] = key_name_prefix
            __props__.__dict__["key_pair_name"] = key_pair_name
            __props__.__dict__["public_key"] = public_key
            __props__.__dict__["resource_group_id"] = resource_group_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["finger_print"] = None
        super(KeyPair, __self__).__init__(
            'alicloud:ecs/keyPair:KeyPair',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            finger_print: Optional[pulumi.Input[str]] = None,
            key_file: Optional[pulumi.Input[str]] = None,
            key_name: Optional[pulumi.Input[str]] = None,
            key_name_prefix: Optional[pulumi.Input[str]] = None,
            key_pair_name: Optional[pulumi.Input[str]] = None,
            public_key: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'KeyPair':
        """
        Get an existing KeyPair resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_file: The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
        :param pulumi.Input[str] key_name: The key pair's name. It is the only in one Alicloud account.
        :param pulumi.Input[str] public_key: You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resource_group_id` is the key pair belongs.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the key pair belongs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KeyPairState.__new__(_KeyPairState)

        __props__.__dict__["finger_print"] = finger_print
        __props__.__dict__["key_file"] = key_file
        __props__.__dict__["key_name"] = key_name
        __props__.__dict__["key_name_prefix"] = key_name_prefix
        __props__.__dict__["key_pair_name"] = key_pair_name
        __props__.__dict__["public_key"] = public_key
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["tags"] = tags
        return KeyPair(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fingerPrint")
    def finger_print(self) -> pulumi.Output[str]:
        return pulumi.get(self, "finger_print")

    @property
    @pulumi.getter(name="keyFile")
    def key_file(self) -> pulumi.Output[Optional[str]]:
        """
        The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
        """
        return pulumi.get(self, "key_file")

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> pulumi.Output[str]:
        """
        The key pair's name. It is the only in one Alicloud account.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter(name="keyNamePrefix")
    def key_name_prefix(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "key_name_prefix")

    @property
    @pulumi.getter(name="keyPairName")
    def key_pair_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "key_pair_name")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[Optional[str]]:
        """
        You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resource_group_id` is the key pair belongs.
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Id of resource group which the key pair belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "tags")

