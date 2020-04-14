# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetPoliciesResult:
    """
    A collection of values returned by getPolicies.
    """
    def __init__(__self__, group_name=None, id=None, name_regex=None, names=None, output_file=None, policies=None, role_name=None, type=None, user_name=None):
        if group_name and not isinstance(group_name, str):
            raise TypeError("Expected argument 'group_name' to be a str")
        __self__.group_name = group_name
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of ram group names.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        __self__.policies = policies
        """
        A list of policies. Each element contains the following attributes:
        """
        if role_name and not isinstance(role_name, str):
            raise TypeError("Expected argument 'role_name' to be a str")
        __self__.role_name = role_name
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Type of the policy.
        """
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        __self__.user_name = user_name
class AwaitableGetPoliciesResult(GetPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPoliciesResult(
            group_name=self.group_name,
            id=self.id,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            policies=self.policies,
            role_name=self.role_name,
            type=self.type,
            user_name=self.user_name)

def get_policies(group_name=None,name_regex=None,output_file=None,role_name=None,type=None,user_name=None,opts=None):
    """
    This data source provides a list of RAM policies in an Alibaba Cloud account according to the specified filters.




    :param str group_name: Filter results by a specific group name. Returned policies are attached to the specified group.
    :param str name_regex: A regex string to filter resulting policies by name.
    :param str role_name: Filter results by a specific role name. Returned policies are attached to the specified role.
    :param str type: Filter results by a specific policy type. Valid values are `Custom` and `System`.
    :param str user_name: Filter results by a specific user name. Returned policies are attached to the specified user.
    """
    __args__ = dict()


    __args__['groupName'] = group_name
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['roleName'] = role_name
    __args__['type'] = type
    __args__['userName'] = user_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:ram/getPolicies:getPolicies', __args__, opts=opts).value

    return AwaitableGetPoliciesResult(
        group_name=__ret__.get('groupName'),
        id=__ret__.get('id'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        policies=__ret__.get('policies'),
        role_name=__ret__.get('roleName'),
        type=__ret__.get('type'),
        user_name=__ret__.get('userName'))
