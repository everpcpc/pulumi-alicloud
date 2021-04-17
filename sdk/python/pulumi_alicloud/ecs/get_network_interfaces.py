# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetNetworkInterfacesResult',
    'AwaitableGetNetworkInterfacesResult',
    'get_network_interfaces',
]

@pulumi.output_type
class GetNetworkInterfacesResult:
    """
    A collection of values returned by getNetworkInterfaces.
    """
    def __init__(__self__, id=None, ids=None, instance_id=None, interfaces=None, name_regex=None, names=None, output_file=None, private_ip=None, resource_group_id=None, security_group_id=None, tags=None, type=None, vpc_id=None, vswitch_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if interfaces and not isinstance(interfaces, list):
            raise TypeError("Expected argument 'interfaces' to be a list")
        pulumi.set(__self__, "interfaces", interfaces)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if private_ip and not isinstance(private_ip, str):
            raise TypeError("Expected argument 'private_ip' to be a str")
        pulumi.set(__self__, "private_ip", private_ip)
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        if security_group_id and not isinstance(security_group_id, str):
            raise TypeError("Expected argument 'security_group_id' to be a str")
        pulumi.set(__self__, "security_group_id", security_group_id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)
        if vswitch_id and not isinstance(vswitch_id, str):
            raise TypeError("Expected argument 'vswitch_id' to be a str")
        pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        """
        ID of the instance that the ENI is attached to.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def interfaces(self) -> Sequence['outputs.GetNetworkInterfacesInterfaceResult']:
        """
        A list of ENIs. Each element contains the following attributes:
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> Optional[str]:
        """
        Primary private IP of the ENI.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        """
        The Id of resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[str]:
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        """
        A map of tags assigned to the ENI.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        ID of the VPC that the ENI belongs to.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[str]:
        """
        ID of the VSwitch that the ENI is linked to.
        """
        return pulumi.get(self, "vswitch_id")


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


def get_network_interfaces(ids: Optional[Sequence[str]] = None,
                           instance_id: Optional[str] = None,
                           name_regex: Optional[str] = None,
                           output_file: Optional[str] = None,
                           private_ip: Optional[str] = None,
                           resource_group_id: Optional[str] = None,
                           security_group_id: Optional[str] = None,
                           tags: Optional[Mapping[str, Any]] = None,
                           type: Optional[str] = None,
                           vpc_id: Optional[str] = None,
                           vswitch_id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkInterfacesResult:
    """
    Use this data source to get a list of elastic network interfaces according to the specified filters in an Alibaba Cloud account.

    For information about elastic network interface and how to use it, see [Elastic Network Interface](https://www.alibabacloud.com/help/doc-detail/58496.html)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "networkInterfacesName"
    vpc = alicloud.vpc.Network("vpc",
        cidr_block="192.168.0.0/24",
        vpc_name=name)
    default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
    vswitch = alicloud.vpc.Switch("vswitch",
        availability_zone=default_zones.zones[0].id,
        cidr_block="192.168.0.0/24",
        vpc_id=vpc.id,
        vswitch_name=name)
    group = alicloud.ecs.SecurityGroup("group", vpc_id=vpc.id)
    interface = alicloud.vpc.NetworkInterface("interface",
        description="Basic test",
        private_ip="192.168.0.2",
        security_groups=[group.id],
        tags={
            "TF-VER": "0.11.3",
        },
        vswitch_id=vswitch.id)
    instance = alicloud.ecs.Instance("instance",
        availability_zone=default_zones.zones[0].id,
        image_id="centos_7_04_64_20G_alibase_201701015.vhd",
        instance_name=name,
        instance_type="ecs.e3.xlarge",
        internet_max_bandwidth_out=10,
        security_groups=[group.id],
        system_disk_category="cloud_efficiency",
        vswitch_id=vswitch.id)
    attachment = alicloud.vpc.NetworkInterfaceAttachment("attachment",
        instance_id=instance.id,
        network_interface_id=interface.id)
    default_network_interfaces = pulumi.Output.all(attachment.network_interface_id, instance.id, group.id, vpc.id, vswitch.id).apply(lambda network_interface_id, instanceId, groupId, vpcId, vswitchId: alicloud.ecs.get_network_interfaces(ids=[network_interface_id],
        instance_id=instance_id,
        name_regex=name,
        private_ip="192.168.0.2",
        security_group_id=group_id,
        tags={
            "TF-VER": "0.11.3",
        },
        type="Secondary",
        vpc_id=vpc_id,
        vswitch_id=vswitch_id))
    pulumi.export("eni0Name", default_network_interfaces.interfaces[0].name)
    ```
    ## Argument Reference

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


    :param str instance_id: ID of the instance that the ENI is attached to.
    :param str private_ip: Primary private IP of the ENI.
    :param str resource_group_id: The Id of resource group.
    :param Mapping[str, Any] tags: A map of tags assigned to the ENI.
    :param str vpc_id: ID of the VPC that the ENI belongs to.
    :param str vswitch_id: ID of the VSwitch that the ENI is linked to.
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
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:ecs/getNetworkInterfaces:getNetworkInterfaces', __args__, opts=opts, typ=GetNetworkInterfacesResult).value

    return AwaitableGetNetworkInterfacesResult(
        id=__ret__.id,
        ids=__ret__.ids,
        instance_id=__ret__.instance_id,
        interfaces=__ret__.interfaces,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        private_ip=__ret__.private_ip,
        resource_group_id=__ret__.resource_group_id,
        security_group_id=__ret__.security_group_id,
        tags=__ret__.tags,
        type=__ret__.type,
        vpc_id=__ret__.vpc_id,
        vswitch_id=__ret__.vswitch_id)
