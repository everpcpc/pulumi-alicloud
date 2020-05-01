# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class EnterpriseInstance(pulumi.CustomResource):
    data_link_name: pulumi.Output[str]
    """
    Cross-database query datalink name.
    """
    database_password: pulumi.Output[str]
    """
    Database access password.
    """
    database_user: pulumi.Output[str]
    """
    Database access account.
    """
    dba_id: pulumi.Output[str]
    dba_nick_name: pulumi.Output[str]
    """
    The instance dba nickname.
    """
    dba_uid: pulumi.Output[float]
    """
    The DBA of the instance is passed into the Alibaba Cloud uid of the DBA.
    """
    ddl_online: pulumi.Output[float]
    """
    Whether to use online services, currently only supports MySQL and PolarDB. Valid values: `0` Not used, `1` Native online DDL priority, `2` DMS lock-free table structure change priority.
    """
    ecs_instance_id: pulumi.Output[str]
    """
    ECS instance ID. The value of InstanceSource is the ECS self-built library. This value must be passed.
    """
    ecs_region: pulumi.Output[str]
    """
    The region where the instance is located. This value must be passed when the value of InstanceSource is RDS, ECS self-built library, and VPC dedicated line IDC.
    """
    env_type: pulumi.Output[str]
    """
    Environment type. Valid values: `product` production environment, `dev` development environment, `pre` pre-release environment, `test` test environment, `sit` SIT environment, `uat` UAT environment, `pet` pressure test environment, `stag` STAG environment.
    """
    export_timeout: pulumi.Output[float]
    """
    Export timeout, unit: s (seconds).
    """
    host: pulumi.Output[str]
    """
    Host address of the target database.
    """
    instance_alias: pulumi.Output[str]
    """
    Instance alias, to help users quickly distinguish positioning.
    """
    instance_id: pulumi.Output[str]
    instance_source: pulumi.Output[str]
    """
    The source of the database instance. Valid values: `PUBLIC_OWN`, `RDS`, `ECS_OWN`, `VPC_IDC`.
    """
    instance_type: pulumi.Output[str]
    """
    Database type. Valid values: `MySQL`, `SQLServer`, `PostgreSQL`, `Oracle,` `DRDS`, `OceanBase`, `Mongo`, `Redis`.
    """
    network_type: pulumi.Output[str]
    """
    Network type. Valid values: `CLASSIC`, `VPC`.
    """
    port: pulumi.Output[float]
    """
    Access port of the target database.
    """
    query_timeout: pulumi.Output[float]
    """
    Query timeout time, unit: s (seconds).
    """
    safe_rule: pulumi.Output[str]
    """
    The security rule of the instance is passed into the name of the security rule in the enterprise.
    """
    safe_rule_id: pulumi.Output[str]
    sid: pulumi.Output[str]
    """
    The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.
    """
    state: pulumi.Output[str]
    """
    The instance status.
    """
    tid: pulumi.Output[float]
    """
    The tenant ID. 
    """
    use_dsql: pulumi.Output[float]
    """
    Whether to enable cross-instance query. Valid values: `0` not open, `1` open.
    """
    vpc_id: pulumi.Output[str]
    """
    VPC ID. This value must be passed when the value of InstanceSource is VPC dedicated line IDC.
    """
    def __init__(__self__, resource_name, opts=None, data_link_name=None, database_password=None, database_user=None, dba_id=None, dba_uid=None, ddl_online=None, ecs_instance_id=None, ecs_region=None, env_type=None, export_timeout=None, host=None, instance_alias=None, instance_id=None, instance_source=None, instance_type=None, network_type=None, port=None, query_timeout=None, safe_rule=None, safe_rule_id=None, sid=None, tid=None, use_dsql=None, vpc_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a DMS Enterprise Instance resource.

        > **NOTE:** API users must first register in DMS.
        > **NOTE:** Available in 1.81.0+.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_link_name: Cross-database query datalink name.
        :param pulumi.Input[str] database_password: Database access password.
        :param pulumi.Input[str] database_user: Database access account.
        :param pulumi.Input[float] dba_uid: The DBA of the instance is passed into the Alibaba Cloud uid of the DBA.
        :param pulumi.Input[float] ddl_online: Whether to use online services, currently only supports MySQL and PolarDB. Valid values: `0` Not used, `1` Native online DDL priority, `2` DMS lock-free table structure change priority.
        :param pulumi.Input[str] ecs_instance_id: ECS instance ID. The value of InstanceSource is the ECS self-built library. This value must be passed.
        :param pulumi.Input[str] ecs_region: The region where the instance is located. This value must be passed when the value of InstanceSource is RDS, ECS self-built library, and VPC dedicated line IDC.
        :param pulumi.Input[str] env_type: Environment type. Valid values: `product` production environment, `dev` development environment, `pre` pre-release environment, `test` test environment, `sit` SIT environment, `uat` UAT environment, `pet` pressure test environment, `stag` STAG environment.
        :param pulumi.Input[float] export_timeout: Export timeout, unit: s (seconds).
        :param pulumi.Input[str] host: Host address of the target database.
        :param pulumi.Input[str] instance_alias: Instance alias, to help users quickly distinguish positioning.
        :param pulumi.Input[str] instance_source: The source of the database instance. Valid values: `PUBLIC_OWN`, `RDS`, `ECS_OWN`, `VPC_IDC`.
        :param pulumi.Input[str] instance_type: Database type. Valid values: `MySQL`, `SQLServer`, `PostgreSQL`, `Oracle,` `DRDS`, `OceanBase`, `Mongo`, `Redis`.
        :param pulumi.Input[str] network_type: Network type. Valid values: `CLASSIC`, `VPC`.
        :param pulumi.Input[float] port: Access port of the target database.
        :param pulumi.Input[float] query_timeout: Query timeout time, unit: s (seconds).
        :param pulumi.Input[str] safe_rule: The security rule of the instance is passed into the name of the security rule in the enterprise.
        :param pulumi.Input[str] sid: The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.
        :param pulumi.Input[float] tid: The tenant ID. 
        :param pulumi.Input[float] use_dsql: Whether to enable cross-instance query. Valid values: `0` not open, `1` open.
        :param pulumi.Input[str] vpc_id: VPC ID. This value must be passed when the value of InstanceSource is VPC dedicated line IDC.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['data_link_name'] = data_link_name
            if database_password is None:
                raise TypeError("Missing required property 'database_password'")
            __props__['database_password'] = database_password
            if database_user is None:
                raise TypeError("Missing required property 'database_user'")
            __props__['database_user'] = database_user
            __props__['dba_id'] = dba_id
            if dba_uid is None:
                raise TypeError("Missing required property 'dba_uid'")
            __props__['dba_uid'] = dba_uid
            __props__['ddl_online'] = ddl_online
            __props__['ecs_instance_id'] = ecs_instance_id
            __props__['ecs_region'] = ecs_region
            if env_type is None:
                raise TypeError("Missing required property 'env_type'")
            __props__['env_type'] = env_type
            if export_timeout is None:
                raise TypeError("Missing required property 'export_timeout'")
            __props__['export_timeout'] = export_timeout
            if host is None:
                raise TypeError("Missing required property 'host'")
            __props__['host'] = host
            if instance_alias is None:
                raise TypeError("Missing required property 'instance_alias'")
            __props__['instance_alias'] = instance_alias
            __props__['instance_id'] = instance_id
            if instance_source is None:
                raise TypeError("Missing required property 'instance_source'")
            __props__['instance_source'] = instance_source
            if instance_type is None:
                raise TypeError("Missing required property 'instance_type'")
            __props__['instance_type'] = instance_type
            if network_type is None:
                raise TypeError("Missing required property 'network_type'")
            __props__['network_type'] = network_type
            if port is None:
                raise TypeError("Missing required property 'port'")
            __props__['port'] = port
            if query_timeout is None:
                raise TypeError("Missing required property 'query_timeout'")
            __props__['query_timeout'] = query_timeout
            if safe_rule is None:
                raise TypeError("Missing required property 'safe_rule'")
            __props__['safe_rule'] = safe_rule
            __props__['safe_rule_id'] = safe_rule_id
            __props__['sid'] = sid
            __props__['tid'] = tid
            __props__['use_dsql'] = use_dsql
            __props__['vpc_id'] = vpc_id
            __props__['dba_nick_name'] = None
            __props__['state'] = None
        super(EnterpriseInstance, __self__).__init__(
            'alicloud:dms/enterpriseInstance:EnterpriseInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, data_link_name=None, database_password=None, database_user=None, dba_id=None, dba_nick_name=None, dba_uid=None, ddl_online=None, ecs_instance_id=None, ecs_region=None, env_type=None, export_timeout=None, host=None, instance_alias=None, instance_id=None, instance_source=None, instance_type=None, network_type=None, port=None, query_timeout=None, safe_rule=None, safe_rule_id=None, sid=None, state=None, tid=None, use_dsql=None, vpc_id=None):
        """
        Get an existing EnterpriseInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_link_name: Cross-database query datalink name.
        :param pulumi.Input[str] database_password: Database access password.
        :param pulumi.Input[str] database_user: Database access account.
        :param pulumi.Input[str] dba_nick_name: The instance dba nickname.
        :param pulumi.Input[float] dba_uid: The DBA of the instance is passed into the Alibaba Cloud uid of the DBA.
        :param pulumi.Input[float] ddl_online: Whether to use online services, currently only supports MySQL and PolarDB. Valid values: `0` Not used, `1` Native online DDL priority, `2` DMS lock-free table structure change priority.
        :param pulumi.Input[str] ecs_instance_id: ECS instance ID. The value of InstanceSource is the ECS self-built library. This value must be passed.
        :param pulumi.Input[str] ecs_region: The region where the instance is located. This value must be passed when the value of InstanceSource is RDS, ECS self-built library, and VPC dedicated line IDC.
        :param pulumi.Input[str] env_type: Environment type. Valid values: `product` production environment, `dev` development environment, `pre` pre-release environment, `test` test environment, `sit` SIT environment, `uat` UAT environment, `pet` pressure test environment, `stag` STAG environment.
        :param pulumi.Input[float] export_timeout: Export timeout, unit: s (seconds).
        :param pulumi.Input[str] host: Host address of the target database.
        :param pulumi.Input[str] instance_alias: Instance alias, to help users quickly distinguish positioning.
        :param pulumi.Input[str] instance_source: The source of the database instance. Valid values: `PUBLIC_OWN`, `RDS`, `ECS_OWN`, `VPC_IDC`.
        :param pulumi.Input[str] instance_type: Database type. Valid values: `MySQL`, `SQLServer`, `PostgreSQL`, `Oracle,` `DRDS`, `OceanBase`, `Mongo`, `Redis`.
        :param pulumi.Input[str] network_type: Network type. Valid values: `CLASSIC`, `VPC`.
        :param pulumi.Input[float] port: Access port of the target database.
        :param pulumi.Input[float] query_timeout: Query timeout time, unit: s (seconds).
        :param pulumi.Input[str] safe_rule: The security rule of the instance is passed into the name of the security rule in the enterprise.
        :param pulumi.Input[str] sid: The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.
        :param pulumi.Input[str] state: The instance status.
        :param pulumi.Input[float] tid: The tenant ID. 
        :param pulumi.Input[float] use_dsql: Whether to enable cross-instance query. Valid values: `0` not open, `1` open.
        :param pulumi.Input[str] vpc_id: VPC ID. This value must be passed when the value of InstanceSource is VPC dedicated line IDC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["data_link_name"] = data_link_name
        __props__["database_password"] = database_password
        __props__["database_user"] = database_user
        __props__["dba_id"] = dba_id
        __props__["dba_nick_name"] = dba_nick_name
        __props__["dba_uid"] = dba_uid
        __props__["ddl_online"] = ddl_online
        __props__["ecs_instance_id"] = ecs_instance_id
        __props__["ecs_region"] = ecs_region
        __props__["env_type"] = env_type
        __props__["export_timeout"] = export_timeout
        __props__["host"] = host
        __props__["instance_alias"] = instance_alias
        __props__["instance_id"] = instance_id
        __props__["instance_source"] = instance_source
        __props__["instance_type"] = instance_type
        __props__["network_type"] = network_type
        __props__["port"] = port
        __props__["query_timeout"] = query_timeout
        __props__["safe_rule"] = safe_rule
        __props__["safe_rule_id"] = safe_rule_id
        __props__["sid"] = sid
        __props__["state"] = state
        __props__["tid"] = tid
        __props__["use_dsql"] = use_dsql
        __props__["vpc_id"] = vpc_id
        return EnterpriseInstance(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

