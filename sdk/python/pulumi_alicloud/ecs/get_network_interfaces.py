# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetNetworkInterfacesResult:
    """
    A collection of values returned by getNetworkInterfaces.
    """
    def __init__(__self__, id=None, ids=None, instance_id=None, interfaces=None, name_regex=None, names=None, output_file=None, private_ip=None, resource_group_id=None, security_group_id=None, tags=None, type=None, vpc_id=None, vswitch_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        __self__.instance_id = instance_id
        """
        ID of the instance that the ENI is attached to.
        """
        if interfaces and not isinstance(interfaces, list):
            raise TypeError("Expected argument 'interfaces' to be a list")
        __self__.interfaces = interfaces
        """
        A list of ENIs. Each element contains the following attributes:
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if private_ip and not isinstance(private_ip, str):
            raise TypeError("Expected argument 'private_ip' to be a str")
        __self__.private_ip = private_ip
        """
        Primary private IP of the ENI.
        """
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        __self__.resource_group_id = resource_group_id
        """
        The Id of resource group.
        """
        if security_group_id and not isinstance(security_group_id, str):
            raise TypeError("Expected argument 'security_group_id' to be a str")
        __self__.security_group_id = security_group_id
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A map of tags assigned to the ENI.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        __self__.vpc_id = vpc_id
        """
        ID of the VPC that the ENI belongs to.
        """
        if vswitch_id and not isinstance(vswitch_id, str):
            raise TypeError("Expected argument 'vswitch_id' to be a str")
        __self__.vswitch_id = vswitch_id
        """
        ID of the VSwitch that the ENI is linked to.
        """
class AwaitableGetNetworkInterfacesResult(GetNetworkInterfacesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkInterfacesResult(
            id=self.id,
            ids=self.ids,
            instance_id=self.instance_id,
            interfaces=self.interfaces,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            private_ip=self.private_ip,
            resource_group_id=self.resource_group_id,
            security_group_id=self.security_group_id,
            tags=self.tags,
            type=self.type,
            vpc_id=self.vpc_id,
            vswitch_id=self.vswitch_id)

def get_network_interfaces(ids=None,instance_id=None,name_regex=None,output_file=None,private_ip=None,resource_group_id=None,security_group_id=None,tags=None,type=None,vpc_id=None,vswitch_id=None,opts=None):
    """
    Use this data source to get a list of elastic network interfaces according to the specified filters in an Alibaba Cloud account.

    For information about elastic network interface and how to use it, see [Elastic Network Interface](https://www.alibabacloud.com/help/doc-detail/58496.html)

    ##  Argument Reference

    The following arguments are supported:

    * `ids` - (Optional)  A list of ENI IDs.
    * `name_regex` - (Optional) A regex string to filter results by ENI name.
    * `vpc_id` - (Optional) The VPC ID linked to ENIs.
    * `vswitch_id` - (Optional) The VSwitch ID linked to ENIs.
    * `private_ip` - (Optional) The primary private IP address of the ENI.
    * `security_group_id` - (Optional) The security group ID linked to ENIs.
    * `name` - (Optional) The name of the ENIs.
    * `type` - (Optional) The type of ENIs, Only support for "Primary" or "Secondary".
    * `instance_id` - (Optional) The ECS instance ID that the ENI is attached to.
    * `tags` - (Optional) A map of tags assigned to ENIs.
    * `output_file` - (Optional) The name of output file that saves the filter results.
    * `resource_group_id` - (Optional, ForceNew, Available in 1.57.0+) The Id of resource group which the network interface belongs.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/network_interfaces.html.markdown.
    """
    __args__ = dict()


    __args__['ids'] = ids
    __args__['instanceId'] = instance_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['privateIp'] = private_ip
    __args__['resourceGroupId'] = resource_group_id
    __args__['securityGroupId'] = security_group_id
    __args__['tags'] = tags
    __args__['type'] = type
    __args__['vpcId'] = vpc_id
    __args__['vswitchId'] = vswitch_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:ecs/getNetworkInterfaces:getNetworkInterfaces', __args__, opts=opts).value

    return AwaitableGetNetworkInterfacesResult(
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        instance_id=__ret__.get('instanceId'),
        interfaces=__ret__.get('interfaces'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        private_ip=__ret__.get('privateIp'),
        resource_group_id=__ret__.get('resourceGroupId'),
        security_group_id=__ret__.get('securityGroupId'),
        tags=__ret__.get('tags'),
        type=__ret__.get('type'),
        vpc_id=__ret__.get('vpcId'),
        vswitch_id=__ret__.get('vswitchId'))
