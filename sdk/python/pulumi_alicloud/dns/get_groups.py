# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetGroupsResult:
    """
    A collection of values returned by getGroups.
    """
    def __init__(__self__, groups=None, id=None, ids=None, name_regex=None, names=None, output_file=None):
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        __self__.groups = groups
        """
        A list of groups. Each element contains the following attributes:
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
        A list of group IDs. 
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of group names.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
class AwaitableGetGroupsResult(GetGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupsResult(
            groups=self.groups,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file)

def get_groups(ids=None,name_regex=None,output_file=None,opts=None):
    """
    This data source provides a list of DNS Domain Groups in an Alibaba Cloud account according to the specified filters.




    :param list ids: A list of group IDs.
    :param str name_regex: A regex string to filter results by group name. 
    """
    __args__ = dict()


    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:dns/getGroups:getGroups', __args__, opts=opts).value

    return AwaitableGetGroupsResult(
        groups=__ret__.get('groups'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'))
