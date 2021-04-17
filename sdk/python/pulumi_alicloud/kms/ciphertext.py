# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CiphertextArgs', 'Ciphertext']

@pulumi.input_type
class CiphertextArgs:
    def __init__(__self__, *,
                 key_id: pulumi.Input[str],
                 plaintext: pulumi.Input[str],
                 encryption_context: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Ciphertext resource.
        :param pulumi.Input[str] key_id: The globally unique ID of the CMK.
        :param pulumi.Input[str] plaintext: The plaintext to be encrypted which must be encoded in Base64.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] encryption_context: -
               (Optional, ForceNew) The Encryption context. If you specify this parameter here, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm).
        """
        pulumi.set(__self__, "key_id", key_id)
        pulumi.set(__self__, "plaintext", plaintext)
        if encryption_context is not None:
            pulumi.set(__self__, "encryption_context", encryption_context)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Input[str]:
        """
        The globally unique ID of the CMK.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter
    def plaintext(self) -> pulumi.Input[str]:
        """
        The plaintext to be encrypted which must be encoded in Base64.
        """
        return pulumi.get(self, "plaintext")

    @plaintext.setter
    def plaintext(self, value: pulumi.Input[str]):
        pulumi.set(self, "plaintext", value)

    @property
    @pulumi.getter(name="encryptionContext")
    def encryption_context(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        -
        (Optional, ForceNew) The Encryption context. If you specify this parameter here, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm).
        """
        return pulumi.get(self, "encryption_context")

    @encryption_context.setter
    def encryption_context(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "encryption_context", value)


@pulumi.input_type
class _CiphertextState:
    def __init__(__self__, *,
                 ciphertext_blob: Optional[pulumi.Input[str]] = None,
                 encryption_context: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 plaintext: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ciphertext resources.
        :param pulumi.Input[str] ciphertext_blob: The ciphertext of the data key encrypted with the primary CMK version.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] encryption_context: -
               (Optional, ForceNew) The Encryption context. If you specify this parameter here, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm).
        :param pulumi.Input[str] key_id: The globally unique ID of the CMK.
        :param pulumi.Input[str] plaintext: The plaintext to be encrypted which must be encoded in Base64.
        """
        if ciphertext_blob is not None:
            pulumi.set(__self__, "ciphertext_blob", ciphertext_blob)
        if encryption_context is not None:
            pulumi.set(__self__, "encryption_context", encryption_context)
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if plaintext is not None:
            pulumi.set(__self__, "plaintext", plaintext)

    @property
    @pulumi.getter(name="ciphertextBlob")
    def ciphertext_blob(self) -> Optional[pulumi.Input[str]]:
        """
        The ciphertext of the data key encrypted with the primary CMK version.
        """
        return pulumi.get(self, "ciphertext_blob")

    @ciphertext_blob.setter
    def ciphertext_blob(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ciphertext_blob", value)

    @property
    @pulumi.getter(name="encryptionContext")
    def encryption_context(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        -
        (Optional, ForceNew) The Encryption context. If you specify this parameter here, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm).
        """
        return pulumi.get(self, "encryption_context")

    @encryption_context.setter
    def encryption_context(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "encryption_context", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The globally unique ID of the CMK.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter
    def plaintext(self) -> Optional[pulumi.Input[str]]:
        """
        The plaintext to be encrypted which must be encoded in Base64.
        """
        return pulumi.get(self, "plaintext")

    @plaintext.setter
    def plaintext(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plaintext", value)


class Ciphertext(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 encryption_context: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 plaintext: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Ciphertext resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] encryption_context: -
               (Optional, ForceNew) The Encryption context. If you specify this parameter here, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm).
        :param pulumi.Input[str] key_id: The globally unique ID of the CMK.
        :param pulumi.Input[str] plaintext: The plaintext to be encrypted which must be encoded in Base64.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CiphertextArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Ciphertext resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CiphertextArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CiphertextArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 encryption_context: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 plaintext: Optional[pulumi.Input[str]] = None,
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
            __props__ = CiphertextArgs.__new__(CiphertextArgs)

            __props__.__dict__["encryption_context"] = encryption_context
            if key_id is None and not opts.urn:
                raise TypeError("Missing required property 'key_id'")
            __props__.__dict__["key_id"] = key_id
            if plaintext is None and not opts.urn:
                raise TypeError("Missing required property 'plaintext'")
            __props__.__dict__["plaintext"] = plaintext
            __props__.__dict__["ciphertext_blob"] = None
        super(Ciphertext, __self__).__init__(
            'alicloud:kms/ciphertext:Ciphertext',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ciphertext_blob: Optional[pulumi.Input[str]] = None,
            encryption_context: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            key_id: Optional[pulumi.Input[str]] = None,
            plaintext: Optional[pulumi.Input[str]] = None) -> 'Ciphertext':
        """
        Get an existing Ciphertext resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ciphertext_blob: The ciphertext of the data key encrypted with the primary CMK version.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] encryption_context: -
               (Optional, ForceNew) The Encryption context. If you specify this parameter here, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm).
        :param pulumi.Input[str] key_id: The globally unique ID of the CMK.
        :param pulumi.Input[str] plaintext: The plaintext to be encrypted which must be encoded in Base64.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CiphertextState.__new__(_CiphertextState)

        __props__.__dict__["ciphertext_blob"] = ciphertext_blob
        __props__.__dict__["encryption_context"] = encryption_context
        __props__.__dict__["key_id"] = key_id
        __props__.__dict__["plaintext"] = plaintext
        return Ciphertext(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ciphertextBlob")
    def ciphertext_blob(self) -> pulumi.Output[str]:
        """
        The ciphertext of the data key encrypted with the primary CMK version.
        """
        return pulumi.get(self, "ciphertext_blob")

    @property
    @pulumi.getter(name="encryptionContext")
    def encryption_context(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        -
        (Optional, ForceNew) The Encryption context. If you specify this parameter here, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm).
        """
        return pulumi.get(self, "encryption_context")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[str]:
        """
        The globally unique ID of the CMK.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter
    def plaintext(self) -> pulumi.Output[str]:
        """
        The plaintext to be encrypted which must be encoded in Base64.
        """
        return pulumi.get(self, "plaintext")

