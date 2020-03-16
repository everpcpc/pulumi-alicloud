# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetInstanceAttachmentsResult:
    """
    A collection of values returned by getInstanceAttachments.
    """
    def __init__(__self__, attachments=None, id=None, instance_name=None, name_regex=None, names=None, output_file=None, vpc_ids=None):
        if attachments and not isinstance(attachments, list):
            raise TypeError("Expected argument 'attachments' to be a list")
        __self__.attachments = attachments
        """
        A list of instance attachments. Each element contains the following attributes:
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if instance_name and not isinstance(instance_name, str):
            raise TypeError("Expected argument 'instance_name' to be a str")
        __self__.instance_name = instance_name
        """
        The instance name.
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of vpc names.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if vpc_ids and not isinstance(vpc_ids, list):
            raise TypeError("Expected argument 'vpc_ids' to be a list")
        __self__.vpc_ids = vpc_ids
        """
        A list of vpc ids.
        """
class AwaitableGetInstanceAttachmentsResult(GetInstanceAttachmentsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceAttachmentsResult(
            attachments=self.attachments,
            id=self.id,
            instance_name=self.instance_name,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            vpc_ids=self.vpc_ids)

def get_instance_attachments(instance_name=None,name_regex=None,output_file=None,opts=None):
    """
    This data source provides the ots instance attachments of the current Alibaba Cloud user.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ots_instance_attachments.html.markdown.


    :param str instance_name: The name of OTS instance.
    :param str name_regex: A regex string to filter results by vpc name.
    """
    __args__ = dict()


    __args__['instanceName'] = instance_name
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:oss/getInstanceAttachments:getInstanceAttachments', __args__, opts=opts).value

    return AwaitableGetInstanceAttachmentsResult(
        attachments=__ret__.get('attachments'),
        id=__ret__.get('id'),
        instance_name=__ret__.get('instanceName'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        vpc_ids=__ret__.get('vpcIds'))
