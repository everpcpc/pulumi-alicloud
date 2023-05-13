# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DbInstanceEndpointArgs', 'DbInstanceEndpoint']

@pulumi.input_type
class DbInstanceEndpointArgs:
    def __init__(__self__, *,
                 connection_string_prefix: pulumi.Input[str],
                 db_instance_id: pulumi.Input[str],
                 node_items: pulumi.Input[Sequence[pulumi.Input['DbInstanceEndpointNodeItemArgs']]],
                 port: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 vswitch_id: pulumi.Input[str],
                 db_instance_endpoint_description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DbInstanceEndpoint resource.
        :param pulumi.Input[str] connection_string_prefix: The IP address of the internal endpoint.
        :param pulumi.Input[str] db_instance_id: The ID of the instance.
        :param pulumi.Input[Sequence[pulumi.Input['DbInstanceEndpointNodeItemArgs']]] node_items: The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
        :param pulumi.Input[str] port: The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
        :param pulumi.Input[str] vpc_id: The virtual private cloud (VPC) ID of the internal endpoint.
        :param pulumi.Input[str] vswitch_id: The vSwitch ID of the internal endpoint.
        :param pulumi.Input[str] db_instance_endpoint_description: The user-defined description of the endpoint.
        """
        pulumi.set(__self__, "connection_string_prefix", connection_string_prefix)
        pulumi.set(__self__, "db_instance_id", db_instance_id)
        pulumi.set(__self__, "node_items", node_items)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vswitch_id", vswitch_id)
        if db_instance_endpoint_description is not None:
            pulumi.set(__self__, "db_instance_endpoint_description", db_instance_endpoint_description)

    @property
    @pulumi.getter(name="connectionStringPrefix")
    def connection_string_prefix(self) -> pulumi.Input[str]:
        """
        The IP address of the internal endpoint.
        """
        return pulumi.get(self, "connection_string_prefix")

    @connection_string_prefix.setter
    def connection_string_prefix(self, value: pulumi.Input[str]):
        pulumi.set(self, "connection_string_prefix", value)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "db_instance_id")

    @db_instance_id.setter
    def db_instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_instance_id", value)

    @property
    @pulumi.getter(name="nodeItems")
    def node_items(self) -> pulumi.Input[Sequence[pulumi.Input['DbInstanceEndpointNodeItemArgs']]]:
        """
        The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
        """
        return pulumi.get(self, "node_items")

    @node_items.setter
    def node_items(self, value: pulumi.Input[Sequence[pulumi.Input['DbInstanceEndpointNodeItemArgs']]]):
        pulumi.set(self, "node_items", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[str]:
        """
        The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[str]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The virtual private cloud (VPC) ID of the internal endpoint.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Input[str]:
        """
        The vSwitch ID of the internal endpoint.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vswitch_id", value)

    @property
    @pulumi.getter(name="dbInstanceEndpointDescription")
    def db_instance_endpoint_description(self) -> Optional[pulumi.Input[str]]:
        """
        The user-defined description of the endpoint.
        """
        return pulumi.get(self, "db_instance_endpoint_description")

    @db_instance_endpoint_description.setter
    def db_instance_endpoint_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_endpoint_description", value)


@pulumi.input_type
class _DbInstanceEndpointState:
    def __init__(__self__, *,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 connection_string_prefix: Optional[pulumi.Input[str]] = None,
                 db_instance_endpoint_description: Optional[pulumi.Input[str]] = None,
                 db_instance_endpoint_id: Optional[pulumi.Input[str]] = None,
                 db_instance_endpoint_type: Optional[pulumi.Input[str]] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 node_items: Optional[pulumi.Input[Sequence[pulumi.Input['DbInstanceEndpointNodeItemArgs']]]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 private_ip_address: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DbInstanceEndpoint resources.
        :param pulumi.Input[str] connection_string: The internal endpoint.
        :param pulumi.Input[str] connection_string_prefix: The IP address of the internal endpoint.
        :param pulumi.Input[str] db_instance_endpoint_description: The user-defined description of the endpoint.
        :param pulumi.Input[str] db_instance_endpoint_id: The Endpoint ID of the instance.
        :param pulumi.Input[str] db_instance_endpoint_type: The type of the endpoint.
        :param pulumi.Input[str] db_instance_id: The ID of the instance.
        :param pulumi.Input[str] ip_type: The type of the IP address.
        :param pulumi.Input[Sequence[pulumi.Input['DbInstanceEndpointNodeItemArgs']]] node_items: The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
        :param pulumi.Input[str] port: The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
        :param pulumi.Input[str] private_ip_address: The IP address of the internal endpoint.
        :param pulumi.Input[str] vpc_id: The virtual private cloud (VPC) ID of the internal endpoint.
        :param pulumi.Input[str] vswitch_id: The vSwitch ID of the internal endpoint.
        """
        if connection_string is not None:
            pulumi.set(__self__, "connection_string", connection_string)
        if connection_string_prefix is not None:
            pulumi.set(__self__, "connection_string_prefix", connection_string_prefix)
        if db_instance_endpoint_description is not None:
            pulumi.set(__self__, "db_instance_endpoint_description", db_instance_endpoint_description)
        if db_instance_endpoint_id is not None:
            pulumi.set(__self__, "db_instance_endpoint_id", db_instance_endpoint_id)
        if db_instance_endpoint_type is not None:
            pulumi.set(__self__, "db_instance_endpoint_type", db_instance_endpoint_type)
        if db_instance_id is not None:
            pulumi.set(__self__, "db_instance_id", db_instance_id)
        if ip_type is not None:
            pulumi.set(__self__, "ip_type", ip_type)
        if node_items is not None:
            pulumi.set(__self__, "node_items", node_items)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if private_ip_address is not None:
            pulumi.set(__self__, "private_ip_address", private_ip_address)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> Optional[pulumi.Input[str]]:
        """
        The internal endpoint.
        """
        return pulumi.get(self, "connection_string")

    @connection_string.setter
    def connection_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_string", value)

    @property
    @pulumi.getter(name="connectionStringPrefix")
    def connection_string_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address of the internal endpoint.
        """
        return pulumi.get(self, "connection_string_prefix")

    @connection_string_prefix.setter
    def connection_string_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_string_prefix", value)

    @property
    @pulumi.getter(name="dbInstanceEndpointDescription")
    def db_instance_endpoint_description(self) -> Optional[pulumi.Input[str]]:
        """
        The user-defined description of the endpoint.
        """
        return pulumi.get(self, "db_instance_endpoint_description")

    @db_instance_endpoint_description.setter
    def db_instance_endpoint_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_endpoint_description", value)

    @property
    @pulumi.getter(name="dbInstanceEndpointId")
    def db_instance_endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Endpoint ID of the instance.
        """
        return pulumi.get(self, "db_instance_endpoint_id")

    @db_instance_endpoint_id.setter
    def db_instance_endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_endpoint_id", value)

    @property
    @pulumi.getter(name="dbInstanceEndpointType")
    def db_instance_endpoint_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the endpoint.
        """
        return pulumi.get(self, "db_instance_endpoint_type")

    @db_instance_endpoint_type.setter
    def db_instance_endpoint_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_endpoint_type", value)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "db_instance_id")

    @db_instance_id.setter
    def db_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_id", value)

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the IP address.
        """
        return pulumi.get(self, "ip_type")

    @ip_type.setter
    def ip_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_type", value)

    @property
    @pulumi.getter(name="nodeItems")
    def node_items(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DbInstanceEndpointNodeItemArgs']]]]:
        """
        The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
        """
        return pulumi.get(self, "node_items")

    @node_items.setter
    def node_items(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DbInstanceEndpointNodeItemArgs']]]]):
        pulumi.set(self, "node_items", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        """
        The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address of the internal endpoint.
        """
        return pulumi.get(self, "private_ip_address")

    @private_ip_address.setter
    def private_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip_address", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The virtual private cloud (VPC) ID of the internal endpoint.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The vSwitch ID of the internal endpoint.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)


class DbInstanceEndpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_string_prefix: Optional[pulumi.Input[str]] = None,
                 db_instance_endpoint_description: Optional[pulumi.Input[str]] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 node_items: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DbInstanceEndpointNodeItemArgs']]]]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provide RDS cluster instance endpoint connection resources.
        > **NOTE:** Available in 1.203.0+.

        ## Block node_items

        The node_items mapping supports the following:

        * `node_id` - (Required) The ID of the node.
        * `weight` - (Required) The weight of the node. Read requests are distributed based on the weight.Valid values: 0 to 100.

        ## Import

        RDS database endpoint feature can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:rds/dbInstanceEndpoint:DbInstanceEndpoint example <db_instance_id>:<db_instance_endpoint_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_string_prefix: The IP address of the internal endpoint.
        :param pulumi.Input[str] db_instance_endpoint_description: The user-defined description of the endpoint.
        :param pulumi.Input[str] db_instance_id: The ID of the instance.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DbInstanceEndpointNodeItemArgs']]]] node_items: The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
        :param pulumi.Input[str] port: The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
        :param pulumi.Input[str] vpc_id: The virtual private cloud (VPC) ID of the internal endpoint.
        :param pulumi.Input[str] vswitch_id: The vSwitch ID of the internal endpoint.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DbInstanceEndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provide RDS cluster instance endpoint connection resources.
        > **NOTE:** Available in 1.203.0+.

        ## Block node_items

        The node_items mapping supports the following:

        * `node_id` - (Required) The ID of the node.
        * `weight` - (Required) The weight of the node. Read requests are distributed based on the weight.Valid values: 0 to 100.

        ## Import

        RDS database endpoint feature can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:rds/dbInstanceEndpoint:DbInstanceEndpoint example <db_instance_id>:<db_instance_endpoint_id>
        ```

        :param str resource_name: The name of the resource.
        :param DbInstanceEndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DbInstanceEndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_string_prefix: Optional[pulumi.Input[str]] = None,
                 db_instance_endpoint_description: Optional[pulumi.Input[str]] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 node_items: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DbInstanceEndpointNodeItemArgs']]]]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DbInstanceEndpointArgs.__new__(DbInstanceEndpointArgs)

            if connection_string_prefix is None and not opts.urn:
                raise TypeError("Missing required property 'connection_string_prefix'")
            __props__.__dict__["connection_string_prefix"] = connection_string_prefix
            __props__.__dict__["db_instance_endpoint_description"] = db_instance_endpoint_description
            if db_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'db_instance_id'")
            __props__.__dict__["db_instance_id"] = db_instance_id
            if node_items is None and not opts.urn:
                raise TypeError("Missing required property 'node_items'")
            __props__.__dict__["node_items"] = node_items
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            if vswitch_id is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__.__dict__["vswitch_id"] = vswitch_id
            __props__.__dict__["connection_string"] = None
            __props__.__dict__["db_instance_endpoint_id"] = None
            __props__.__dict__["db_instance_endpoint_type"] = None
            __props__.__dict__["ip_type"] = None
            __props__.__dict__["private_ip_address"] = None
        super(DbInstanceEndpoint, __self__).__init__(
            'alicloud:rds/dbInstanceEndpoint:DbInstanceEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            connection_string: Optional[pulumi.Input[str]] = None,
            connection_string_prefix: Optional[pulumi.Input[str]] = None,
            db_instance_endpoint_description: Optional[pulumi.Input[str]] = None,
            db_instance_endpoint_id: Optional[pulumi.Input[str]] = None,
            db_instance_endpoint_type: Optional[pulumi.Input[str]] = None,
            db_instance_id: Optional[pulumi.Input[str]] = None,
            ip_type: Optional[pulumi.Input[str]] = None,
            node_items: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DbInstanceEndpointNodeItemArgs']]]]] = None,
            port: Optional[pulumi.Input[str]] = None,
            private_ip_address: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None) -> 'DbInstanceEndpoint':
        """
        Get an existing DbInstanceEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_string: The internal endpoint.
        :param pulumi.Input[str] connection_string_prefix: The IP address of the internal endpoint.
        :param pulumi.Input[str] db_instance_endpoint_description: The user-defined description of the endpoint.
        :param pulumi.Input[str] db_instance_endpoint_id: The Endpoint ID of the instance.
        :param pulumi.Input[str] db_instance_endpoint_type: The type of the endpoint.
        :param pulumi.Input[str] db_instance_id: The ID of the instance.
        :param pulumi.Input[str] ip_type: The type of the IP address.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DbInstanceEndpointNodeItemArgs']]]] node_items: The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
        :param pulumi.Input[str] port: The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
        :param pulumi.Input[str] private_ip_address: The IP address of the internal endpoint.
        :param pulumi.Input[str] vpc_id: The virtual private cloud (VPC) ID of the internal endpoint.
        :param pulumi.Input[str] vswitch_id: The vSwitch ID of the internal endpoint.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DbInstanceEndpointState.__new__(_DbInstanceEndpointState)

        __props__.__dict__["connection_string"] = connection_string
        __props__.__dict__["connection_string_prefix"] = connection_string_prefix
        __props__.__dict__["db_instance_endpoint_description"] = db_instance_endpoint_description
        __props__.__dict__["db_instance_endpoint_id"] = db_instance_endpoint_id
        __props__.__dict__["db_instance_endpoint_type"] = db_instance_endpoint_type
        __props__.__dict__["db_instance_id"] = db_instance_id
        __props__.__dict__["ip_type"] = ip_type
        __props__.__dict__["node_items"] = node_items
        __props__.__dict__["port"] = port
        __props__.__dict__["private_ip_address"] = private_ip_address
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["vswitch_id"] = vswitch_id
        return DbInstanceEndpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> pulumi.Output[str]:
        """
        The internal endpoint.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="connectionStringPrefix")
    def connection_string_prefix(self) -> pulumi.Output[str]:
        """
        The IP address of the internal endpoint.
        """
        return pulumi.get(self, "connection_string_prefix")

    @property
    @pulumi.getter(name="dbInstanceEndpointDescription")
    def db_instance_endpoint_description(self) -> pulumi.Output[str]:
        """
        The user-defined description of the endpoint.
        """
        return pulumi.get(self, "db_instance_endpoint_description")

    @property
    @pulumi.getter(name="dbInstanceEndpointId")
    def db_instance_endpoint_id(self) -> pulumi.Output[str]:
        """
        The Endpoint ID of the instance.
        """
        return pulumi.get(self, "db_instance_endpoint_id")

    @property
    @pulumi.getter(name="dbInstanceEndpointType")
    def db_instance_endpoint_type(self) -> pulumi.Output[str]:
        """
        The type of the endpoint.
        """
        return pulumi.get(self, "db_instance_endpoint_type")

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "db_instance_id")

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> pulumi.Output[str]:
        """
        The type of the IP address.
        """
        return pulumi.get(self, "ip_type")

    @property
    @pulumi.getter(name="nodeItems")
    def node_items(self) -> pulumi.Output[Sequence['outputs.DbInstanceEndpointNodeItem']]:
        """
        The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
        """
        return pulumi.get(self, "node_items")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[str]:
        """
        The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> pulumi.Output[str]:
        """
        The IP address of the internal endpoint.
        """
        return pulumi.get(self, "private_ip_address")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The virtual private cloud (VPC) ID of the internal endpoint.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[str]:
        """
        The vSwitch ID of the internal endpoint.
        """
        return pulumi.get(self, "vswitch_id")
