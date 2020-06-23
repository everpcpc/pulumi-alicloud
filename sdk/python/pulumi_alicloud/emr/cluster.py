# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Cluster(pulumi.CustomResource):
    bootstrap_actions: pulumi.Output[list]
    charge_type: pulumi.Output[str]
    """
    Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global charge_type value.
    """
    cluster_type: pulumi.Output[str]
    """
    EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
    """
    deposit_type: pulumi.Output[str]
    """
    Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
    """
    eas_enable: pulumi.Output[bool]
    """
    High security cluster (true) or not. Default value is false.
    """
    emr_ver: pulumi.Output[str]
    """
    EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
    """
    high_availability_enable: pulumi.Output[bool]
    """
    High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
    """
    host_groups: pulumi.Output[list]
    """
    Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).

      * `auto_renew` (`bool`) - Auto renew for prepaid, true of false. Default is false.
      * `charge_type` (`str`) - Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global charge_type value.
      * `diskCapacity` (`str`) - Data disk capacity.
      * `diskCount` (`str`) - Data disk count.
      * `disk_type` (`str`) - Data disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,local_disk,cloud_essd.
      * `gpuDriver` (`str`)
      * `hostGroupName` (`str`) - host group name.
      * `hostGroupType` (`str`) - host group type, supported value: MASTER, CORE or TASK, supported 'GATEWAY' available in 1.61.0+.
      * `instanceList` (`str`) - Instance list for cluster scale down. This value follows the json format, e.g. ["instance_id1","instance_id2"]. escape character for " is \".
      * `instance_type` (`str`) - Host Ecs instance type.
      * `node_count` (`str`) - Host number in this group.
      * `period` (`float`) - If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
      * `sysDiskCapacity` (`str`) - System disk capacity.
      * `sysDiskType` (`str`) - System disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,cloud_essd.
    """
    is_open_public_ip: pulumi.Output[bool]
    key_pair_name: pulumi.Output[str]
    """
    Ssh key pair.
    """
    master_pwd: pulumi.Output[str]
    """
    Master ssh password.
    """
    name: pulumi.Output[str]
    """
    bootstrap action name.
    """
    option_software_lists: pulumi.Output[list]
    """
    Optional software list.
    """
    related_cluster_id: pulumi.Output[str]
    """
    This specify the related cluster id, if this cluster is a Gateway.
    """
    security_group_id: pulumi.Output[str]
    """
    Security Group ID for Cluster, you can also specify this key for each host group.
    """
    ssh_enable: pulumi.Output[bool]
    """
    If this is set true, we can ssh into cluster. Default value is false.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    use_local_metadb: pulumi.Output[bool]
    """
    Use local metadb. Default is false.
    """
    user_defined_emr_ecs_role: pulumi.Output[str]
    """
    Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
    """
    vswitch_id: pulumi.Output[str]
    """
    Global vswitch id, you can also specify it in host group.
    """
    zone_id: pulumi.Output[str]
    """
    Zone ID, e.g. cn-huhehaote-a
    """
    def __init__(__self__, resource_name, opts=None, bootstrap_actions=None, charge_type=None, cluster_type=None, deposit_type=None, eas_enable=None, emr_ver=None, high_availability_enable=None, host_groups=None, is_open_public_ip=None, key_pair_name=None, master_pwd=None, name=None, option_software_lists=None, related_cluster_id=None, security_group_id=None, ssh_enable=None, tags=None, use_local_metadb=None, user_defined_emr_ecs_role=None, vswitch_id=None, zone_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a EMR Cluster resource. With this you can create, read, and release  EMR Cluster. 

        > **NOTE:** Available in 1.57.0+.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] charge_type: Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global charge_type value.
        :param pulumi.Input[str] cluster_type: EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
        :param pulumi.Input[str] deposit_type: Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
        :param pulumi.Input[bool] eas_enable: High security cluster (true) or not. Default value is false.
        :param pulumi.Input[str] emr_ver: EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
        :param pulumi.Input[bool] high_availability_enable: High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
        :param pulumi.Input[list] host_groups: Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
        :param pulumi.Input[str] key_pair_name: Ssh key pair.
        :param pulumi.Input[str] master_pwd: Master ssh password.
        :param pulumi.Input[str] name: bootstrap action name.
        :param pulumi.Input[list] option_software_lists: Optional software list.
        :param pulumi.Input[str] related_cluster_id: This specify the related cluster id, if this cluster is a Gateway.
        :param pulumi.Input[str] security_group_id: Security Group ID for Cluster, you can also specify this key for each host group.
        :param pulumi.Input[bool] ssh_enable: If this is set true, we can ssh into cluster. Default value is false.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[bool] use_local_metadb: Use local metadb. Default is false.
        :param pulumi.Input[str] user_defined_emr_ecs_role: Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
        :param pulumi.Input[str] vswitch_id: Global vswitch id, you can also specify it in host group.
        :param pulumi.Input[str] zone_id: Zone ID, e.g. cn-huhehaote-a

        The **bootstrap_actions** object supports the following:

          * `arg` (`pulumi.Input[str]`) - bootstrap action args, e.g. "--a=b".
          * `name` (`pulumi.Input[str]`) - bootstrap action name.
          * `path` (`pulumi.Input[str]`) - bootstrap action path, e.g. "oss://bucket/path".

        The **host_groups** object supports the following:

          * `auto_renew` (`pulumi.Input[bool]`) - Auto renew for prepaid, true of false. Default is false.
          * `charge_type` (`pulumi.Input[str]`) - Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global charge_type value.
          * `diskCapacity` (`pulumi.Input[str]`) - Data disk capacity.
          * `diskCount` (`pulumi.Input[str]`) - Data disk count.
          * `disk_type` (`pulumi.Input[str]`) - Data disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,local_disk,cloud_essd.
          * `gpuDriver` (`pulumi.Input[str]`)
          * `hostGroupName` (`pulumi.Input[str]`) - host group name.
          * `hostGroupType` (`pulumi.Input[str]`) - host group type, supported value: MASTER, CORE or TASK, supported 'GATEWAY' available in 1.61.0+.
          * `instanceList` (`pulumi.Input[str]`) - Instance list for cluster scale down. This value follows the json format, e.g. ["instance_id1","instance_id2"]. escape character for " is \".
          * `instance_type` (`pulumi.Input[str]`) - Host Ecs instance type.
          * `node_count` (`pulumi.Input[str]`) - Host number in this group.
          * `period` (`pulumi.Input[float]`) - If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
          * `sysDiskCapacity` (`pulumi.Input[str]`) - System disk capacity.
          * `sysDiskType` (`pulumi.Input[str]`) - System disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,cloud_essd.
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

            __props__['bootstrap_actions'] = bootstrap_actions
            __props__['charge_type'] = charge_type
            if cluster_type is None:
                raise TypeError("Missing required property 'cluster_type'")
            __props__['cluster_type'] = cluster_type
            __props__['deposit_type'] = deposit_type
            __props__['eas_enable'] = eas_enable
            if emr_ver is None:
                raise TypeError("Missing required property 'emr_ver'")
            __props__['emr_ver'] = emr_ver
            __props__['high_availability_enable'] = high_availability_enable
            __props__['host_groups'] = host_groups
            __props__['is_open_public_ip'] = is_open_public_ip
            __props__['key_pair_name'] = key_pair_name
            __props__['master_pwd'] = master_pwd
            __props__['name'] = name
            __props__['option_software_lists'] = option_software_lists
            __props__['related_cluster_id'] = related_cluster_id
            __props__['security_group_id'] = security_group_id
            __props__['ssh_enable'] = ssh_enable
            __props__['tags'] = tags
            __props__['use_local_metadb'] = use_local_metadb
            __props__['user_defined_emr_ecs_role'] = user_defined_emr_ecs_role
            __props__['vswitch_id'] = vswitch_id
            if zone_id is None:
                raise TypeError("Missing required property 'zone_id'")
            __props__['zone_id'] = zone_id
        super(Cluster, __self__).__init__(
            'alicloud:emr/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bootstrap_actions=None, charge_type=None, cluster_type=None, deposit_type=None, eas_enable=None, emr_ver=None, high_availability_enable=None, host_groups=None, is_open_public_ip=None, key_pair_name=None, master_pwd=None, name=None, option_software_lists=None, related_cluster_id=None, security_group_id=None, ssh_enable=None, tags=None, use_local_metadb=None, user_defined_emr_ecs_role=None, vswitch_id=None, zone_id=None):
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] charge_type: Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global charge_type value.
        :param pulumi.Input[str] cluster_type: EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
        :param pulumi.Input[str] deposit_type: Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
        :param pulumi.Input[bool] eas_enable: High security cluster (true) or not. Default value is false.
        :param pulumi.Input[str] emr_ver: EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
        :param pulumi.Input[bool] high_availability_enable: High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
        :param pulumi.Input[list] host_groups: Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
        :param pulumi.Input[str] key_pair_name: Ssh key pair.
        :param pulumi.Input[str] master_pwd: Master ssh password.
        :param pulumi.Input[str] name: bootstrap action name.
        :param pulumi.Input[list] option_software_lists: Optional software list.
        :param pulumi.Input[str] related_cluster_id: This specify the related cluster id, if this cluster is a Gateway.
        :param pulumi.Input[str] security_group_id: Security Group ID for Cluster, you can also specify this key for each host group.
        :param pulumi.Input[bool] ssh_enable: If this is set true, we can ssh into cluster. Default value is false.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[bool] use_local_metadb: Use local metadb. Default is false.
        :param pulumi.Input[str] user_defined_emr_ecs_role: Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
        :param pulumi.Input[str] vswitch_id: Global vswitch id, you can also specify it in host group.
        :param pulumi.Input[str] zone_id: Zone ID, e.g. cn-huhehaote-a

        The **bootstrap_actions** object supports the following:

          * `arg` (`pulumi.Input[str]`) - bootstrap action args, e.g. "--a=b".
          * `name` (`pulumi.Input[str]`) - bootstrap action name.
          * `path` (`pulumi.Input[str]`) - bootstrap action path, e.g. "oss://bucket/path".

        The **host_groups** object supports the following:

          * `auto_renew` (`pulumi.Input[bool]`) - Auto renew for prepaid, true of false. Default is false.
          * `charge_type` (`pulumi.Input[str]`) - Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global charge_type value.
          * `diskCapacity` (`pulumi.Input[str]`) - Data disk capacity.
          * `diskCount` (`pulumi.Input[str]`) - Data disk count.
          * `disk_type` (`pulumi.Input[str]`) - Data disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,local_disk,cloud_essd.
          * `gpuDriver` (`pulumi.Input[str]`)
          * `hostGroupName` (`pulumi.Input[str]`) - host group name.
          * `hostGroupType` (`pulumi.Input[str]`) - host group type, supported value: MASTER, CORE or TASK, supported 'GATEWAY' available in 1.61.0+.
          * `instanceList` (`pulumi.Input[str]`) - Instance list for cluster scale down. This value follows the json format, e.g. ["instance_id1","instance_id2"]. escape character for " is \".
          * `instance_type` (`pulumi.Input[str]`) - Host Ecs instance type.
          * `node_count` (`pulumi.Input[str]`) - Host number in this group.
          * `period` (`pulumi.Input[float]`) - If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
          * `sysDiskCapacity` (`pulumi.Input[str]`) - System disk capacity.
          * `sysDiskType` (`pulumi.Input[str]`) - System disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,cloud_essd.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bootstrap_actions"] = bootstrap_actions
        __props__["charge_type"] = charge_type
        __props__["cluster_type"] = cluster_type
        __props__["deposit_type"] = deposit_type
        __props__["eas_enable"] = eas_enable
        __props__["emr_ver"] = emr_ver
        __props__["high_availability_enable"] = high_availability_enable
        __props__["host_groups"] = host_groups
        __props__["is_open_public_ip"] = is_open_public_ip
        __props__["key_pair_name"] = key_pair_name
        __props__["master_pwd"] = master_pwd
        __props__["name"] = name
        __props__["option_software_lists"] = option_software_lists
        __props__["related_cluster_id"] = related_cluster_id
        __props__["security_group_id"] = security_group_id
        __props__["ssh_enable"] = ssh_enable
        __props__["tags"] = tags
        __props__["use_local_metadb"] = use_local_metadb
        __props__["user_defined_emr_ecs_role"] = user_defined_emr_ecs_role
        __props__["vswitch_id"] = vswitch_id
        __props__["zone_id"] = zone_id
        return Cluster(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

