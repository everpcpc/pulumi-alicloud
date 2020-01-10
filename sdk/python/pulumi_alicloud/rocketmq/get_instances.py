# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetInstancesResult:
    """
    A collection of values returned by getInstances.
    """
    def __init__(__self__, ids=None, instances=None, name_regex=None, names=None, output_file=None, id=None):
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of instance IDs.
        """
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        __self__.instances = instances
        """
        A list of instances. Each element contains the following attributes:
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of instance names.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetInstancesResult(GetInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancesResult(
            ids=self.ids,
            instances=self.instances,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            id=self.id)

def get_instances(ids=None,name_regex=None,output_file=None,opts=None):
    """
    This data source provides a list of ONS Instances in an Alibaba Cloud account according to the specified filters.
    
    > **NOTE:** Available in 1.52.0+
    
    :param list ids: A list of instance IDs to filter results.
    :param str name_regex: A regex string to filter results by the instance name. 

    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ons_instances.html.markdown.
    """
    __args__ = dict()

    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:rocketmq/getInstances:getInstances', __args__, opts=opts).value

    return AwaitableGetInstancesResult(
        ids=__ret__.get('ids'),
        instances=__ret__.get('instances'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        id=__ret__.get('id'))
