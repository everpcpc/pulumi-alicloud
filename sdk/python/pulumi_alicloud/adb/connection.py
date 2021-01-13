# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Connection']


class Connection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_prefix: Optional[pulumi.Input[str]] = None,
                 db_cluster_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an ADB connection resource to allocate an Internet connection string for ADB cluster.

        > **NOTE:** Each ADB instance will allocate a intranet connnection string automatically and its prifix is ADB instance ID.
         To avoid unnecessary conflict, please specified a internet connection prefix before applying the resource.

        > **NOTE:** Available in v1.81.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        creation = config.get("creation")
        if creation is None:
            creation = "ADB"
        name = config.get("name")
        if name is None:
            name = "adbaccountmysql"
        default_zones = alicloud.get_zones(available_resource_creation=creation)
        default_network = alicloud.vpc.Network("defaultNetwork", cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/24",
            availability_zone=default_zones.zones[0].id)
        cluster = alicloud.adb.Cluster("cluster",
            db_cluster_version="3.0",
            db_cluster_category="Cluster",
            db_node_class="C8",
            db_node_count=2,
            db_node_storage=200,
            pay_type="PostPaid",
            vswitch_id=default_switch.id,
            description=name)
        connection = alicloud.adb.Connection("connection",
            db_cluster_id=cluster.id,
            connection_prefix="testabc")
        ```

        ## Import

        ADB connection can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:adb/connection:Connection example am-12345678
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_prefix: Prefix of the cluster public endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter. Default to `<db_cluster_id> + tf`.
        :param pulumi.Input[str] db_cluster_id: The Id of cluster that can run database.
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

            __props__['connection_prefix'] = connection_prefix
            if db_cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'db_cluster_id'")
            __props__['db_cluster_id'] = db_cluster_id
            __props__['connection_string'] = None
            __props__['ip_address'] = None
            __props__['port'] = None
        super(Connection, __self__).__init__(
            'alicloud:adb/connection:Connection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            connection_prefix: Optional[pulumi.Input[str]] = None,
            connection_string: Optional[pulumi.Input[str]] = None,
            db_cluster_id: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[str]] = None) -> 'Connection':
        """
        Get an existing Connection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_prefix: Prefix of the cluster public endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter. Default to `<db_cluster_id> + tf`.
        :param pulumi.Input[str] connection_string: Connection cluster string.
        :param pulumi.Input[str] db_cluster_id: The Id of cluster that can run database.
        :param pulumi.Input[str] ip_address: The ip address of connection string.
        :param pulumi.Input[str] port: Connection cluster port.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["connection_prefix"] = connection_prefix
        __props__["connection_string"] = connection_string
        __props__["db_cluster_id"] = db_cluster_id
        __props__["ip_address"] = ip_address
        __props__["port"] = port
        return Connection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectionPrefix")
    def connection_prefix(self) -> pulumi.Output[str]:
        """
        Prefix of the cluster public endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter. Default to `<db_cluster_id> + tf`.
        """
        return pulumi.get(self, "connection_prefix")

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> pulumi.Output[str]:
        """
        Connection cluster string.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="dbClusterId")
    def db_cluster_id(self) -> pulumi.Output[str]:
        """
        The Id of cluster that can run database.
        """
        return pulumi.get(self, "db_cluster_id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        The ip address of connection string.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[str]:
        """
        Connection cluster port.
        """
        return pulumi.get(self, "port")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

