# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetSecretVersionsResult:
    """
    A collection of values returned by getSecretVersions.
    """
    def __init__(__self__, enable_details=None, id=None, ids=None, include_deprecated=None, output_file=None, secret_name=None, version_stage=None, versions=None):
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        __self__.enable_details = enable_details
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of Kms Secret Version ids. 
        """
        if include_deprecated and not isinstance(include_deprecated, str):
            raise TypeError("Expected argument 'include_deprecated' to be a str")
        __self__.include_deprecated = include_deprecated
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if secret_name and not isinstance(secret_name, str):
            raise TypeError("Expected argument 'secret_name' to be a str")
        __self__.secret_name = secret_name
        """
        The name of the secret.
        """
        if version_stage and not isinstance(version_stage, str):
            raise TypeError("Expected argument 'version_stage' to be a str")
        __self__.version_stage = version_stage
        if versions and not isinstance(versions, list):
            raise TypeError("Expected argument 'versions' to be a list")
        __self__.versions = versions
        """
        A list of KMS Secret Versions. Each element contains the following attributes:
        """
class AwaitableGetSecretVersionsResult(GetSecretVersionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecretVersionsResult(
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            include_deprecated=self.include_deprecated,
            output_file=self.output_file,
            secret_name=self.secret_name,
            version_stage=self.version_stage,
            versions=self.versions)

def get_secret_versions(enable_details=None,ids=None,include_deprecated=None,output_file=None,secret_name=None,version_stage=None,opts=None):
    """
    This data source provides a list of KMS Secret Versions in an Alibaba Cloud account according to the specified filters.
     
    > **NOTE:** Available in v1.88.0+.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    kms_secret_versions_ds = alicloud.kms.get_secret_versions(enable_details=True,
        secret_name="secret_name")
    pulumi.export("firstSecretData", kms_secret_versions_ds.versions[0]["secret_data"])
    ```



    :param bool enable_details: Default to false and only output `secret_name`, `version_id`, `version_stages`. Set it to true can output more details.
    :param list ids: A list of KMS Secret Version ids.
    :param str include_deprecated: Specifies whether to return deprecated secret versions. Default to `false`.
    :param str secret_name: The name of the secret.
    :param str version_stage: The stage of the secret version.
    """
    __args__ = dict()


    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['includeDeprecated'] = include_deprecated
    __args__['outputFile'] = output_file
    __args__['secretName'] = secret_name
    __args__['versionStage'] = version_stage
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:kms/getSecretVersions:getSecretVersions', __args__, opts=opts).value

    return AwaitableGetSecretVersionsResult(
        enable_details=__ret__.get('enableDetails'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        include_deprecated=__ret__.get('includeDeprecated'),
        output_file=__ret__.get('outputFile'),
        secret_name=__ret__.get('secretName'),
        version_stage=__ret__.get('versionStage'),
        versions=__ret__.get('versions'))
