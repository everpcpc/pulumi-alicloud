# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetSecretsResult:
    """
    A collection of values returned by getSecrets.
    """
    def __init__(__self__, fetch_tags=None, id=None, ids=None, name_regex=None, names=None, output_file=None, secrets=None, tags=None):
        if fetch_tags and not isinstance(fetch_tags, bool):
            raise TypeError("Expected argument 'fetch_tags' to be a bool")
        __self__.fetch_tags = fetch_tags
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
        A list of Kms Secret ids. The value is same as KMS secret_name. 
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of KMS Secret names.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if secrets and not isinstance(secrets, list):
            raise TypeError("Expected argument 'secrets' to be a list")
        __self__.secrets = secrets
        """
        A list of KMS Secrets. Each element contains the following attributes:
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
class AwaitableGetSecretsResult(GetSecretsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecretsResult(
            fetch_tags=self.fetch_tags,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            secrets=self.secrets,
            tags=self.tags)

def get_secrets(fetch_tags=None,ids=None,name_regex=None,output_file=None,tags=None,opts=None):
    """
    This data source provides a list of KMS Secrets in an Alibaba Cloud account according to the specified filters.
     
    > **NOTE:** Available in v1.86.0+.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    kms_secrets_ds = alicloud.kms.get_secrets(fetch_tags=True,
        name_regex="name_regex",
        tags={
            "k-aa": "v-aa",
            "k-bb": "v-bb",
        })
    pulumi.export("firstSecretId", kms_secrets_ds.secrets[0]["id"])
    ```



    :param bool fetch_tags: Whether to include the predetermined resource tag in the return value. Default to `false`.
    :param list ids: A list of KMS Secret ids. The value is same as KMS secret_name.
    :param str name_regex: A regex string to filter the results by the KMS secret_name.
    :param dict tags: A mapping of tags to assign to the resource.
    """
    __args__ = dict()


    __args__['fetchTags'] = fetch_tags
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:kms/getSecrets:getSecrets', __args__, opts=opts).value

    return AwaitableGetSecretsResult(
        fetch_tags=__ret__.get('fetchTags'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        secrets=__ret__.get('secrets'),
        tags=__ret__.get('tags'))
