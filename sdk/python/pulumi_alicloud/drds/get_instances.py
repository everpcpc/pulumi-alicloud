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
    def __init__(__self__, descriptions=None, id=None, ids=None, instances=None, name_regex=None, output_file=None):
        if descriptions and not isinstance(descriptions, list):
            raise TypeError("Expected argument 'descriptions' to be a list")
        __self__.descriptions = descriptions
        """
        A list of DRDS descriptions. 
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of DRDS instance IDs.
        """
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        __self__.instances = instances
        """
        A list of DRDS instances.
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
class AwaitableGetInstancesResult(GetInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancesResult(
            descriptions=self.descriptions,
            id=self.id,
            ids=self.ids,
            instances=self.instances,
            name_regex=self.name_regex,
            output_file=self.output_file)

def get_instances(ids=None,name_regex=None,output_file=None,opts=None):
    """
     The `drds.Instance` data source provides a collection of DRDS instances available in Alibaba Cloud account.
    Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.

    > **NOTE:** Available in 1.35.0+.




    :param list ids: A list of DRDS instance IDs.
    :param str name_regex: A regex string to filter results by instance name.
    """
    __args__ = dict()


    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:drds/getInstances:getInstances', __args__, opts=opts).value

    return AwaitableGetInstancesResult(
        descriptions=__ret__.get('descriptions'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        instances=__ret__.get('instances'),
        name_regex=__ret__.get('nameRegex'),
        output_file=__ret__.get('outputFile'))
