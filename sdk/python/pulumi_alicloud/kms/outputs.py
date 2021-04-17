# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetAliasesAliasResult',
    'GetKeyVersionsVersionResult',
    'GetKeysKeyResult',
    'GetSecretVersionsVersionResult',
    'GetSecretsSecretResult',
]

@pulumi.output_type
class GetAliasesAliasResult(dict):
    def __init__(__self__, *,
                 alias_name: str,
                 id: str,
                 key_id: str):
        """
        :param str alias_name: The unique identifier of the alias.
        :param str id: ID of the alias. The value is same as KMS alias_name.
        :param str key_id: ID of the key.
        """
        pulumi.set(__self__, "alias_name", alias_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "key_id", key_id)

    @property
    @pulumi.getter(name="aliasName")
    def alias_name(self) -> str:
        """
        The unique identifier of the alias.
        """
        return pulumi.get(self, "alias_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the alias. The value is same as KMS alias_name.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> str:
        """
        ID of the key.
        """
        return pulumi.get(self, "key_id")


@pulumi.output_type
class GetKeyVersionsVersionResult(dict):
    def __init__(__self__, *,
                 creation_date: str,
                 id: str,
                 key_id: str,
                 key_version_id: str):
        """
        :param str creation_date: Date and time when the key version was created (UTC time).
        :param str id: ID of the KMS KeyVersion resource.
        :param str key_id: The id of kms key.
        :param str key_version_id: ID of the key version.
        """
        pulumi.set(__self__, "creation_date", creation_date)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "key_id", key_id)
        pulumi.set(__self__, "key_version_id", key_version_id)

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        Date and time when the key version was created (UTC time).
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the KMS KeyVersion resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> str:
        """
        The id of kms key.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="keyVersionId")
    def key_version_id(self) -> str:
        """
        ID of the key version.
        """
        return pulumi.get(self, "key_version_id")


@pulumi.output_type
class GetKeysKeyResult(dict):
    def __init__(__self__, *,
                 arn: str,
                 creation_date: str,
                 creator: str,
                 delete_date: str,
                 description: str,
                 id: str,
                 status: str):
        """
        :param str arn: The Alibaba Cloud Resource Name (ARN) of the key.
        :param str creation_date: Creation date of key.
        :param str creator: The owner of the key.
        :param str delete_date: Deletion date of key.
        :param str description: Description of the key.
        :param str id: ID of the key.
        :param str status: Filter the results by status of the KMS keys. Valid values: `Enabled`, `Disabled`, `PendingDeletion`.
        """
        pulumi.set(__self__, "arn", arn)
        pulumi.set(__self__, "creation_date", creation_date)
        pulumi.set(__self__, "creator", creator)
        pulumi.set(__self__, "delete_date", delete_date)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The Alibaba Cloud Resource Name (ARN) of the key.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        Creation date of key.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def creator(self) -> str:
        """
        The owner of the key.
        """
        return pulumi.get(self, "creator")

    @property
    @pulumi.getter(name="deleteDate")
    def delete_date(self) -> str:
        """
        Deletion date of key.
        """
        return pulumi.get(self, "delete_date")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the key.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the key.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Filter the results by status of the KMS keys. Valid values: `Enabled`, `Disabled`, `PendingDeletion`.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class GetSecretVersionsVersionResult(dict):
    def __init__(__self__, *,
                 secret_data: str,
                 secret_data_type: str,
                 secret_name: str,
                 version_id: str,
                 version_stages: Sequence[str]):
        """
        :param str secret_data: The secret value. Secrets Manager decrypts the stored secret value in ciphertext and returns it. (Returned when `enable_details` is true).
        :param str secret_data_type: The type of the secret value. (Returned when `enable_details` is true).
        :param str secret_name: The name of the secret.
        :param str version_id: The version number of the secret value.
        :param Sequence[str] version_stages: Stage labels that mark the secret version.
        """
        pulumi.set(__self__, "secret_data", secret_data)
        pulumi.set(__self__, "secret_data_type", secret_data_type)
        pulumi.set(__self__, "secret_name", secret_name)
        pulumi.set(__self__, "version_id", version_id)
        pulumi.set(__self__, "version_stages", version_stages)

    @property
    @pulumi.getter(name="secretData")
    def secret_data(self) -> str:
        """
        The secret value. Secrets Manager decrypts the stored secret value in ciphertext and returns it. (Returned when `enable_details` is true).
        """
        return pulumi.get(self, "secret_data")

    @property
    @pulumi.getter(name="secretDataType")
    def secret_data_type(self) -> str:
        """
        The type of the secret value. (Returned when `enable_details` is true).
        """
        return pulumi.get(self, "secret_data_type")

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> str:
        """
        The name of the secret.
        """
        return pulumi.get(self, "secret_name")

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> str:
        """
        The version number of the secret value.
        """
        return pulumi.get(self, "version_id")

    @property
    @pulumi.getter(name="versionStages")
    def version_stages(self) -> Sequence[str]:
        """
        Stage labels that mark the secret version.
        """
        return pulumi.get(self, "version_stages")


@pulumi.output_type
class GetSecretsSecretResult(dict):
    def __init__(__self__, *,
                 id: str,
                 planned_delete_time: str,
                 secret_name: str,
                 tags: Mapping[str, Any]):
        """
        :param str id: ID of the Kms Secret. The value is same as KMS secret_name.
        :param str planned_delete_time: Schedule deletion time.
        :param str secret_name: Name of the KMS Secret.
        :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "planned_delete_time", planned_delete_time)
        pulumi.set(__self__, "secret_name", secret_name)
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the Kms Secret. The value is same as KMS secret_name.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="plannedDeleteTime")
    def planned_delete_time(self) -> str:
        """
        Schedule deletion time.
        """
        return pulumi.get(self, "planned_delete_time")

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> str:
        """
        Name of the KMS Secret.
        """
        return pulumi.get(self, "secret_name")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, Any]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")


