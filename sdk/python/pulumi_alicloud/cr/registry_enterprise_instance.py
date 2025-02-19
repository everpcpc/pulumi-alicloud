# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RegistryEnterpriseInstanceArgs', 'RegistryEnterpriseInstance']

@pulumi.input_type
class RegistryEnterpriseInstanceArgs:
    def __init__(__self__, *,
                 instance_name: pulumi.Input[str],
                 instance_type: pulumi.Input[str],
                 custom_oss_bucket: Optional[pulumi.Input[str]] = None,
                 kms_encrypted_password: Optional[pulumi.Input[str]] = None,
                 kms_encryption_context: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 payment_type: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 renew_period: Optional[pulumi.Input[int]] = None,
                 renewal_status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RegistryEnterpriseInstance resource.
        :param pulumi.Input[str] instance_name: Name of Container Registry Enterprise Edition instance.
        :param pulumi.Input[str] instance_type: Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
        :param pulumi.Input[str] custom_oss_bucket: Name of your customized oss bucket. Use this bucket as instance storage if set.
        :param pulumi.Input[str] kms_encrypted_password: An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
        :param pulumi.Input[Mapping[str, Any]] kms_encryption_context: An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        :param pulumi.Input[str] password: The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
        :param pulumi.Input[str] payment_type: Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
        :param pulumi.Input[int] period: Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
        :param pulumi.Input[int] renew_period: Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
        :param pulumi.Input[str] renewal_status: Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
        """
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "instance_type", instance_type)
        if custom_oss_bucket is not None:
            pulumi.set(__self__, "custom_oss_bucket", custom_oss_bucket)
        if kms_encrypted_password is not None:
            pulumi.set(__self__, "kms_encrypted_password", kms_encrypted_password)
        if kms_encryption_context is not None:
            pulumi.set(__self__, "kms_encryption_context", kms_encryption_context)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if payment_type is not None:
            pulumi.set(__self__, "payment_type", payment_type)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if renew_period is not None:
            pulumi.set(__self__, "renew_period", renew_period)
        if renewal_status is not None:
            pulumi.set(__self__, "renewal_status", renewal_status)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Input[str]:
        """
        Name of Container Registry Enterprise Edition instance.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        """
        Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="customOssBucket")
    def custom_oss_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        Name of your customized oss bucket. Use this bucket as instance storage if set.
        """
        return pulumi.get(self, "custom_oss_bucket")

    @custom_oss_bucket.setter
    def custom_oss_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_oss_bucket", value)

    @property
    @pulumi.getter(name="kmsEncryptedPassword")
    def kms_encrypted_password(self) -> Optional[pulumi.Input[str]]:
        """
        An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
        """
        return pulumi.get(self, "kms_encrypted_password")

    @kms_encrypted_password.setter
    def kms_encrypted_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_encrypted_password", value)

    @property
    @pulumi.getter(name="kmsEncryptionContext")
    def kms_encryption_context(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        """
        return pulumi.get(self, "kms_encryption_context")

    @kms_encryption_context.setter
    def kms_encryption_context(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "kms_encryption_context", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> Optional[pulumi.Input[str]]:
        """
        Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
        """
        return pulumi.get(self, "payment_type")

    @payment_type.setter
    def payment_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payment_type", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="renewPeriod")
    def renew_period(self) -> Optional[pulumi.Input[int]]:
        """
        Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
        """
        return pulumi.get(self, "renew_period")

    @renew_period.setter
    def renew_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "renew_period", value)

    @property
    @pulumi.getter(name="renewalStatus")
    def renewal_status(self) -> Optional[pulumi.Input[str]]:
        """
        Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
        """
        return pulumi.get(self, "renewal_status")

    @renewal_status.setter
    def renewal_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "renewal_status", value)


@pulumi.input_type
class _RegistryEnterpriseInstanceState:
    def __init__(__self__, *,
                 created_time: Optional[pulumi.Input[str]] = None,
                 custom_oss_bucket: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 kms_encrypted_password: Optional[pulumi.Input[str]] = None,
                 kms_encryption_context: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 payment_type: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 renew_period: Optional[pulumi.Input[int]] = None,
                 renewal_status: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RegistryEnterpriseInstance resources.
        :param pulumi.Input[str] created_time: Time of Container Registry Enterprise Edition instance creation.
        :param pulumi.Input[str] custom_oss_bucket: Name of your customized oss bucket. Use this bucket as instance storage if set.
        :param pulumi.Input[str] end_time: Time of Container Registry Enterprise Edition instance expiration.
        :param pulumi.Input[str] instance_name: Name of Container Registry Enterprise Edition instance.
        :param pulumi.Input[str] instance_type: Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
        :param pulumi.Input[str] kms_encrypted_password: An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
        :param pulumi.Input[Mapping[str, Any]] kms_encryption_context: An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        :param pulumi.Input[str] password: The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
        :param pulumi.Input[str] payment_type: Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
        :param pulumi.Input[int] period: Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
        :param pulumi.Input[int] renew_period: Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
        :param pulumi.Input[str] renewal_status: Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
        :param pulumi.Input[str] status: Status of Container Registry Enterprise Edition instance.
        """
        if created_time is not None:
            pulumi.set(__self__, "created_time", created_time)
        if custom_oss_bucket is not None:
            pulumi.set(__self__, "custom_oss_bucket", custom_oss_bucket)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if kms_encrypted_password is not None:
            pulumi.set(__self__, "kms_encrypted_password", kms_encrypted_password)
        if kms_encryption_context is not None:
            pulumi.set(__self__, "kms_encryption_context", kms_encryption_context)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if payment_type is not None:
            pulumi.set(__self__, "payment_type", payment_type)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if renew_period is not None:
            pulumi.set(__self__, "renew_period", renew_period)
        if renewal_status is not None:
            pulumi.set(__self__, "renewal_status", renewal_status)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time of Container Registry Enterprise Edition instance creation.
        """
        return pulumi.get(self, "created_time")

    @created_time.setter
    def created_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_time", value)

    @property
    @pulumi.getter(name="customOssBucket")
    def custom_oss_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        Name of your customized oss bucket. Use this bucket as instance storage if set.
        """
        return pulumi.get(self, "custom_oss_bucket")

    @custom_oss_bucket.setter
    def custom_oss_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_oss_bucket", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time of Container Registry Enterprise Edition instance expiration.
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Container Registry Enterprise Edition instance.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="kmsEncryptedPassword")
    def kms_encrypted_password(self) -> Optional[pulumi.Input[str]]:
        """
        An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
        """
        return pulumi.get(self, "kms_encrypted_password")

    @kms_encrypted_password.setter
    def kms_encrypted_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_encrypted_password", value)

    @property
    @pulumi.getter(name="kmsEncryptionContext")
    def kms_encryption_context(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        """
        return pulumi.get(self, "kms_encryption_context")

    @kms_encryption_context.setter
    def kms_encryption_context(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "kms_encryption_context", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> Optional[pulumi.Input[str]]:
        """
        Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
        """
        return pulumi.get(self, "payment_type")

    @payment_type.setter
    def payment_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payment_type", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="renewPeriod")
    def renew_period(self) -> Optional[pulumi.Input[int]]:
        """
        Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
        """
        return pulumi.get(self, "renew_period")

    @renew_period.setter
    def renew_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "renew_period", value)

    @property
    @pulumi.getter(name="renewalStatus")
    def renewal_status(self) -> Optional[pulumi.Input[str]]:
        """
        Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
        """
        return pulumi.get(self, "renewal_status")

    @renewal_status.setter
    def renewal_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "renewal_status", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of Container Registry Enterprise Edition instance.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class RegistryEnterpriseInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_oss_bucket: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 kms_encrypted_password: Optional[pulumi.Input[str]] = None,
                 kms_encryption_context: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 payment_type: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 renew_period: Optional[pulumi.Input[int]] = None,
                 renewal_status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource will help you to manager Container Registry Enterprise Edition instances.

        For information about Container Registry Enterprise Edition instances and how to use it, see [Create a Instance](https://www.alibabacloud.com/help/en/doc-detail/208144.htm)

        > **NOTE:** Available in v1.124.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        my_instance = alicloud.cr.RegistryEnterpriseInstance("my-instance",
            instance_name="test",
            instance_type="Advanced",
            payment_type="Subscription",
            period=1,
            renew_period=1,
            renewal_status="AutoRenewal")
        ```

        ## Import

        Container Registry Enterprise Edition instance can be imported using the `id`, e.g.

        ```sh
         $ pulumi import alicloud:cr/registryEnterpriseInstance:RegistryEnterpriseInstance default cri-test
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] custom_oss_bucket: Name of your customized oss bucket. Use this bucket as instance storage if set.
        :param pulumi.Input[str] instance_name: Name of Container Registry Enterprise Edition instance.
        :param pulumi.Input[str] instance_type: Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
        :param pulumi.Input[str] kms_encrypted_password: An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
        :param pulumi.Input[Mapping[str, Any]] kms_encryption_context: An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        :param pulumi.Input[str] password: The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
        :param pulumi.Input[str] payment_type: Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
        :param pulumi.Input[int] period: Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
        :param pulumi.Input[int] renew_period: Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
        :param pulumi.Input[str] renewal_status: Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegistryEnterpriseInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource will help you to manager Container Registry Enterprise Edition instances.

        For information about Container Registry Enterprise Edition instances and how to use it, see [Create a Instance](https://www.alibabacloud.com/help/en/doc-detail/208144.htm)

        > **NOTE:** Available in v1.124.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        my_instance = alicloud.cr.RegistryEnterpriseInstance("my-instance",
            instance_name="test",
            instance_type="Advanced",
            payment_type="Subscription",
            period=1,
            renew_period=1,
            renewal_status="AutoRenewal")
        ```

        ## Import

        Container Registry Enterprise Edition instance can be imported using the `id`, e.g.

        ```sh
         $ pulumi import alicloud:cr/registryEnterpriseInstance:RegistryEnterpriseInstance default cri-test
        ```

        :param str resource_name: The name of the resource.
        :param RegistryEnterpriseInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegistryEnterpriseInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_oss_bucket: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 kms_encrypted_password: Optional[pulumi.Input[str]] = None,
                 kms_encryption_context: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 payment_type: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 renew_period: Optional[pulumi.Input[int]] = None,
                 renewal_status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegistryEnterpriseInstanceArgs.__new__(RegistryEnterpriseInstanceArgs)

            __props__.__dict__["custom_oss_bucket"] = custom_oss_bucket
            if instance_name is None and not opts.urn:
                raise TypeError("Missing required property 'instance_name'")
            __props__.__dict__["instance_name"] = instance_name
            if instance_type is None and not opts.urn:
                raise TypeError("Missing required property 'instance_type'")
            __props__.__dict__["instance_type"] = instance_type
            __props__.__dict__["kms_encrypted_password"] = kms_encrypted_password
            __props__.__dict__["kms_encryption_context"] = kms_encryption_context
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["payment_type"] = payment_type
            __props__.__dict__["period"] = period
            __props__.__dict__["renew_period"] = renew_period
            __props__.__dict__["renewal_status"] = renewal_status
            __props__.__dict__["created_time"] = None
            __props__.__dict__["end_time"] = None
            __props__.__dict__["status"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(RegistryEnterpriseInstance, __self__).__init__(
            'alicloud:cr/registryEnterpriseInstance:RegistryEnterpriseInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_time: Optional[pulumi.Input[str]] = None,
            custom_oss_bucket: Optional[pulumi.Input[str]] = None,
            end_time: Optional[pulumi.Input[str]] = None,
            instance_name: Optional[pulumi.Input[str]] = None,
            instance_type: Optional[pulumi.Input[str]] = None,
            kms_encrypted_password: Optional[pulumi.Input[str]] = None,
            kms_encryption_context: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            password: Optional[pulumi.Input[str]] = None,
            payment_type: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[int]] = None,
            renew_period: Optional[pulumi.Input[int]] = None,
            renewal_status: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'RegistryEnterpriseInstance':
        """
        Get an existing RegistryEnterpriseInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_time: Time of Container Registry Enterprise Edition instance creation.
        :param pulumi.Input[str] custom_oss_bucket: Name of your customized oss bucket. Use this bucket as instance storage if set.
        :param pulumi.Input[str] end_time: Time of Container Registry Enterprise Edition instance expiration.
        :param pulumi.Input[str] instance_name: Name of Container Registry Enterprise Edition instance.
        :param pulumi.Input[str] instance_type: Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
        :param pulumi.Input[str] kms_encrypted_password: An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
        :param pulumi.Input[Mapping[str, Any]] kms_encryption_context: An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        :param pulumi.Input[str] password: The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
        :param pulumi.Input[str] payment_type: Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
        :param pulumi.Input[int] period: Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
        :param pulumi.Input[int] renew_period: Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
        :param pulumi.Input[str] renewal_status: Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
        :param pulumi.Input[str] status: Status of Container Registry Enterprise Edition instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RegistryEnterpriseInstanceState.__new__(_RegistryEnterpriseInstanceState)

        __props__.__dict__["created_time"] = created_time
        __props__.__dict__["custom_oss_bucket"] = custom_oss_bucket
        __props__.__dict__["end_time"] = end_time
        __props__.__dict__["instance_name"] = instance_name
        __props__.__dict__["instance_type"] = instance_type
        __props__.__dict__["kms_encrypted_password"] = kms_encrypted_password
        __props__.__dict__["kms_encryption_context"] = kms_encryption_context
        __props__.__dict__["password"] = password
        __props__.__dict__["payment_type"] = payment_type
        __props__.__dict__["period"] = period
        __props__.__dict__["renew_period"] = renew_period
        __props__.__dict__["renewal_status"] = renewal_status
        __props__.__dict__["status"] = status
        return RegistryEnterpriseInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        """
        Time of Container Registry Enterprise Edition instance creation.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="customOssBucket")
    def custom_oss_bucket(self) -> pulumi.Output[Optional[str]]:
        """
        Name of your customized oss bucket. Use this bucket as instance storage if set.
        """
        return pulumi.get(self, "custom_oss_bucket")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[str]:
        """
        Time of Container Registry Enterprise Edition instance expiration.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Output[str]:
        """
        Name of Container Registry Enterprise Edition instance.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        """
        Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="kmsEncryptedPassword")
    def kms_encrypted_password(self) -> pulumi.Output[Optional[str]]:
        """
        An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
        """
        return pulumi.get(self, "kms_encrypted_password")

    @property
    @pulumi.getter(name="kmsEncryptionContext")
    def kms_encryption_context(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        """
        return pulumi.get(self, "kms_encryption_context")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> pulumi.Output[Optional[str]]:
        """
        Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
        """
        return pulumi.get(self, "payment_type")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[int]]:
        """
        Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="renewPeriod")
    def renew_period(self) -> pulumi.Output[Optional[int]]:
        """
        Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
        """
        return pulumi.get(self, "renew_period")

    @property
    @pulumi.getter(name="renewalStatus")
    def renewal_status(self) -> pulumi.Output[Optional[str]]:
        """
        Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
        """
        return pulumi.get(self, "renewal_status")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of Container Registry Enterprise Edition instance.
        """
        return pulumi.get(self, "status")

