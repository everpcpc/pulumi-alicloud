# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Swarm']


class Swarm(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 disk_category: Optional[pulumi.Input[str]] = None,
                 disk_size: Optional[pulumi.Input[int]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 is_outdated: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 need_slb: Optional[pulumi.Input[bool]] = None,
                 node_number: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 release_eip: Optional[pulumi.Input[bool]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        > **DEPRECATED:** This resource manages swarm cluster, which is being deprecated and will be replaced by Kubernetes cluster.

        This resource will help you to manager a Swarm Cluster.

        > **NOTE:** Swarm cluster only supports VPC network and you can specify a VPC network by filed `vswitch_id`.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        my_cluster = alicloud.cs.Swarm("myCluster",
            password="Yourpassword1234",
            instance_type="ecs.n4.small",
            node_number=2,
            disk_category="cloud_efficiency",
            disk_size=20,
            cidr_block="172.18.0.0/24",
            image_id=var["image_id"],
            vswitch_id=var["vswitch_id"])
        ```

        ## Import

        Swarm cluster can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cs/swarm:Swarm foo cf123456789
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr_block: The CIDR block for the Container. It can not be same as the CIDR used by the VPC.
               Valid value:
               - 192.168.0.0/16
               - 172.19-30.0.0/16
               - 10.0.0.0/16
        :param pulumi.Input[str] disk_category: The data disk category of ECS instance node. Its valid value are `cloud`, `cloud_ssd`, `cloud_essd`, `ephemeral_essd` and `cloud_efficiency`. Default to `cloud_efficiency`.
        :param pulumi.Input[int] disk_size: The data disk size of ECS instance node. Its valid value is 20~32768 GB. Default to 20.
        :param pulumi.Input[str] image_id: The image ID of ECS instance node used. Default to System automate allocated.
        :param pulumi.Input[str] instance_type: The type of ECS instance node.
        :param pulumi.Input[bool] is_outdated: Whether to use outdated instance type. Default to false.
        :param pulumi.Input[str] name: The container cluster's name. It is the only in one Alicloud account.
        :param pulumi.Input[bool] need_slb: Whether to create the default simple routing Server Load Balancer instance for the cluster. The default value is true.
        :param pulumi.Input[int] node_number: The ECS node number of the container cluster. Its value choices are 1~50, and default to 1.
        :param pulumi.Input[str] password: The password of ECS instance node.
        :param pulumi.Input[bool] release_eip: Whether to release EIP after creating swarm cluster successfully. Default to false.
        :param pulumi.Input[int] size: Field 'size' has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.
        :param pulumi.Input[str] vswitch_id: The password of ECS instance node. If it is not specified, the container cluster's network mode will be `Classic`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'cidr_block'")
            __props__['cidr_block'] = cidr_block
            __props__['disk_category'] = disk_category
            __props__['disk_size'] = disk_size
            __props__['image_id'] = image_id
            if instance_type is None and not opts.urn:
                raise TypeError("Missing required property 'instance_type'")
            __props__['instance_type'] = instance_type
            __props__['is_outdated'] = is_outdated
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            __props__['need_slb'] = need_slb
            __props__['node_number'] = node_number
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__['password'] = password
            __props__['release_eip'] = release_eip
            if size is not None and not opts.urn:
                warnings.warn("""Field 'size' has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.""", DeprecationWarning)
                pulumi.log.warn("size is deprecated: Field 'size' has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.")
            __props__['size'] = size
            if vswitch_id is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__['vswitch_id'] = vswitch_id
            __props__['agent_version'] = None
            __props__['nodes'] = None
            __props__['security_group_id'] = None
            __props__['slb_id'] = None
            __props__['vpc_id'] = None
        super(Swarm, __self__).__init__(
            'alicloud:cs/swarm:Swarm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            agent_version: Optional[pulumi.Input[str]] = None,
            cidr_block: Optional[pulumi.Input[str]] = None,
            disk_category: Optional[pulumi.Input[str]] = None,
            disk_size: Optional[pulumi.Input[int]] = None,
            image_id: Optional[pulumi.Input[str]] = None,
            instance_type: Optional[pulumi.Input[str]] = None,
            is_outdated: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            need_slb: Optional[pulumi.Input[bool]] = None,
            node_number: Optional[pulumi.Input[int]] = None,
            nodes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwarmNodeArgs']]]]] = None,
            password: Optional[pulumi.Input[str]] = None,
            release_eip: Optional[pulumi.Input[bool]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            slb_id: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None) -> 'Swarm':
        """
        Get an existing Swarm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] agent_version: The nodes agent version.
        :param pulumi.Input[str] cidr_block: The CIDR block for the Container. It can not be same as the CIDR used by the VPC.
               Valid value:
               - 192.168.0.0/16
               - 172.19-30.0.0/16
               - 10.0.0.0/16
        :param pulumi.Input[str] disk_category: The data disk category of ECS instance node. Its valid value are `cloud`, `cloud_ssd`, `cloud_essd`, `ephemeral_essd` and `cloud_efficiency`. Default to `cloud_efficiency`.
        :param pulumi.Input[int] disk_size: The data disk size of ECS instance node. Its valid value is 20~32768 GB. Default to 20.
        :param pulumi.Input[str] image_id: The image ID of ECS instance node used. Default to System automate allocated.
        :param pulumi.Input[str] instance_type: The type of ECS instance node.
        :param pulumi.Input[bool] is_outdated: Whether to use outdated instance type. Default to false.
        :param pulumi.Input[str] name: The container cluster's name. It is the only in one Alicloud account.
        :param pulumi.Input[bool] need_slb: Whether to create the default simple routing Server Load Balancer instance for the cluster. The default value is true.
        :param pulumi.Input[int] node_number: The ECS node number of the container cluster. Its value choices are 1~50, and default to 1.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwarmNodeArgs']]]] nodes: List of cluster nodes. It contains several attributes to `Block Nodes`.
        :param pulumi.Input[str] password: The password of ECS instance node.
        :param pulumi.Input[bool] release_eip: Whether to release EIP after creating swarm cluster successfully. Default to false.
        :param pulumi.Input[str] security_group_id: The ID of security group where the current cluster worker node is located.
        :param pulumi.Input[int] size: Field 'size' has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.
        :param pulumi.Input[str] slb_id: The ID of load balancer where the current cluster worker node is located.
        :param pulumi.Input[str] vpc_id: The ID of VPC where the current cluster is located.
        :param pulumi.Input[str] vswitch_id: The password of ECS instance node. If it is not specified, the container cluster's network mode will be `Classic`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["agent_version"] = agent_version
        __props__["cidr_block"] = cidr_block
        __props__["disk_category"] = disk_category
        __props__["disk_size"] = disk_size
        __props__["image_id"] = image_id
        __props__["instance_type"] = instance_type
        __props__["is_outdated"] = is_outdated
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["need_slb"] = need_slb
        __props__["node_number"] = node_number
        __props__["nodes"] = nodes
        __props__["password"] = password
        __props__["release_eip"] = release_eip
        __props__["security_group_id"] = security_group_id
        __props__["size"] = size
        __props__["slb_id"] = slb_id
        __props__["vpc_id"] = vpc_id
        __props__["vswitch_id"] = vswitch_id
        return Swarm(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="agentVersion")
    def agent_version(self) -> pulumi.Output[str]:
        """
        The nodes agent version.
        """
        return pulumi.get(self, "agent_version")

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> pulumi.Output[str]:
        """
        The CIDR block for the Container. It can not be same as the CIDR used by the VPC.
        Valid value:
        - 192.168.0.0/16
        - 172.19-30.0.0/16
        - 10.0.0.0/16
        """
        return pulumi.get(self, "cidr_block")

    @property
    @pulumi.getter(name="diskCategory")
    def disk_category(self) -> pulumi.Output[Optional[str]]:
        """
        The data disk category of ECS instance node. Its valid value are `cloud`, `cloud_ssd`, `cloud_essd`, `ephemeral_essd` and `cloud_efficiency`. Default to `cloud_efficiency`.
        """
        return pulumi.get(self, "disk_category")

    @property
    @pulumi.getter(name="diskSize")
    def disk_size(self) -> pulumi.Output[Optional[int]]:
        """
        The data disk size of ECS instance node. Its valid value is 20~32768 GB. Default to 20.
        """
        return pulumi.get(self, "disk_size")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[Optional[str]]:
        """
        The image ID of ECS instance node used. Default to System automate allocated.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        """
        The type of ECS instance node.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="isOutdated")
    def is_outdated(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to use outdated instance type. Default to false.
        """
        return pulumi.get(self, "is_outdated")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The container cluster's name. It is the only in one Alicloud account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter(name="needSlb")
    def need_slb(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to create the default simple routing Server Load Balancer instance for the cluster. The default value is true.
        """
        return pulumi.get(self, "need_slb")

    @property
    @pulumi.getter(name="nodeNumber")
    def node_number(self) -> pulumi.Output[Optional[int]]:
        """
        The ECS node number of the container cluster. Its value choices are 1~50, and default to 1.
        """
        return pulumi.get(self, "node_number")

    @property
    @pulumi.getter
    def nodes(self) -> pulumi.Output[Sequence['outputs.SwarmNode']]:
        """
        List of cluster nodes. It contains several attributes to `Block Nodes`.
        """
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        The password of ECS instance node.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="releaseEip")
    def release_eip(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to release EIP after creating swarm cluster successfully. Default to false.
        """
        return pulumi.get(self, "release_eip")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[str]:
        """
        The ID of security group where the current cluster worker node is located.
        """
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[Optional[int]]:
        """
        Field 'size' has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="slbId")
    def slb_id(self) -> pulumi.Output[str]:
        """
        The ID of load balancer where the current cluster worker node is located.
        """
        return pulumi.get(self, "slb_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of VPC where the current cluster is located.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[str]:
        """
        The password of ECS instance node. If it is not specified, the container cluster's network mode will be `Classic`.
        """
        return pulumi.get(self, "vswitch_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

