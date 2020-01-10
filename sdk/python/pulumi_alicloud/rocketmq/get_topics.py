# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetTopicsResult:
    """
    A collection of values returned by getTopics.
    """
    def __init__(__self__, instance_id=None, name_regex=None, names=None, output_file=None, topics=None, id=None):
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        __self__.instance_id = instance_id
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of topic names.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if topics and not isinstance(topics, list):
            raise TypeError("Expected argument 'topics' to be a list")
        __self__.topics = topics
        """
        A list of topics. Each element contains the following attributes:
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetTopicsResult(GetTopicsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTopicsResult(
            instance_id=self.instance_id,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            topics=self.topics,
            id=self.id)

def get_topics(instance_id=None,name_regex=None,output_file=None,opts=None):
    """
    This data source provides a list of ONS Topics in an Alibaba Cloud account according to the specified filters.
    
    > **NOTE:** Available in 1.53.0+
    
    :param str instance_id: ID of the ONS Instance that owns the topics.
    :param str name_regex: A regex string to filter results by the topic name. 

    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ons_topics.html.markdown.
    """
    __args__ = dict()

    __args__['instanceId'] = instance_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:rocketmq/getTopics:getTopics', __args__, opts=opts).value

    return AwaitableGetTopicsResult(
        instance_id=__ret__.get('instanceId'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        topics=__ret__.get('topics'),
        id=__ret__.get('id'))
