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
    'DbClusterDbClusterAccessWhiteList',
    'GetAccountsAccountResult',
    'GetDbClustersClusterResult',
    'GetDbClustersClusterDbClusterAccessWhiteListResult',
    'GetDbClustersClusterScaleOutStatusResult',
    'GetRegionsRegionResult',
    'GetRegionsRegionZoneIdResult',
]

@pulumi.output_type
class DbClusterDbClusterAccessWhiteList(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dbClusterIpArrayAttribute":
            suggest = "db_cluster_ip_array_attribute"
        elif key == "dbClusterIpArrayName":
            suggest = "db_cluster_ip_array_name"
        elif key == "securityIpList":
            suggest = "security_ip_list"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DbClusterDbClusterAccessWhiteList. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DbClusterDbClusterAccessWhiteList.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DbClusterDbClusterAccessWhiteList.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 db_cluster_ip_array_attribute: Optional[str] = None,
                 db_cluster_ip_array_name: Optional[str] = None,
                 security_ip_list: Optional[str] = None):
        """
        :param str db_cluster_ip_array_attribute: Whitelist grouping attribute.
        :param str db_cluster_ip_array_name: Whitelist group name.
        :param str security_ip_list: The IP address list under the whitelist group.
        """
        if db_cluster_ip_array_attribute is not None:
            pulumi.set(__self__, "db_cluster_ip_array_attribute", db_cluster_ip_array_attribute)
        if db_cluster_ip_array_name is not None:
            pulumi.set(__self__, "db_cluster_ip_array_name", db_cluster_ip_array_name)
        if security_ip_list is not None:
            pulumi.set(__self__, "security_ip_list", security_ip_list)

    @property
    @pulumi.getter(name="dbClusterIpArrayAttribute")
    def db_cluster_ip_array_attribute(self) -> Optional[str]:
        """
        Whitelist grouping attribute.
        """
        return pulumi.get(self, "db_cluster_ip_array_attribute")

    @property
    @pulumi.getter(name="dbClusterIpArrayName")
    def db_cluster_ip_array_name(self) -> Optional[str]:
        """
        Whitelist group name.
        """
        return pulumi.get(self, "db_cluster_ip_array_name")

    @property
    @pulumi.getter(name="securityIpList")
    def security_ip_list(self) -> Optional[str]:
        """
        The IP address list under the whitelist group.
        """
        return pulumi.get(self, "security_ip_list")


@pulumi.output_type
class GetAccountsAccountResult(dict):
    def __init__(__self__, *,
                 account_description: str,
                 account_name: str,
                 account_type: str,
                 db_cluster_id: str,
                 id: str,
                 status: str):
        """
        :param str account_description: In Chinese, English letter. May contain Chinese and English characters, lowercase letters, numbers, and underscores (_), the dash (-). Cannot start with http:// and https:// at the beginning. Length is from 2 to 256 characters.
        :param str account_name: Account name: lowercase letters, numbers, underscores, lowercase letter; length no more than 16 characters.
        :param str account_type: The Valid Account type: `Normal`, `Super`.
        :param str db_cluster_id: The DBCluster id.
        :param str id: The ID of the Account. Its value is same as Queue Name.
        :param str status: The status of the resource.
        """
        pulumi.set(__self__, "account_description", account_description)
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "account_type", account_type)
        pulumi.set(__self__, "db_cluster_id", db_cluster_id)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="accountDescription")
    def account_description(self) -> str:
        """
        In Chinese, English letter. May contain Chinese and English characters, lowercase letters, numbers, and underscores (_), the dash (-). Cannot start with http:// and https:// at the beginning. Length is from 2 to 256 characters.
        """
        return pulumi.get(self, "account_description")

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> str:
        """
        Account name: lowercase letters, numbers, underscores, lowercase letter; length no more than 16 characters.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="accountType")
    def account_type(self) -> str:
        """
        The Valid Account type: `Normal`, `Super`.
        """
        return pulumi.get(self, "account_type")

    @property
    @pulumi.getter(name="dbClusterId")
    def db_cluster_id(self) -> str:
        """
        The DBCluster id.
        """
        return pulumi.get(self, "db_cluster_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Account. Its value is same as Queue Name.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class GetDbClustersClusterResult(dict):
    def __init__(__self__, *,
                 ali_uid: str,
                 bid: str,
                 category: str,
                 commodity_code: str,
                 connection_string: str,
                 create_time: str,
                 db_cluster_access_white_lists: Sequence['outputs.GetDbClustersClusterDbClusterAccessWhiteListResult'],
                 db_cluster_description: str,
                 db_cluster_id: str,
                 db_cluster_network_type: str,
                 db_cluster_type: str,
                 db_node_class: str,
                 db_node_count: str,
                 db_node_storage: str,
                 encryption_key: str,
                 encryption_type: str,
                 engine: str,
                 engine_version: str,
                 expire_time: str,
                 id: str,
                 is_expired: str,
                 lock_mode: str,
                 lock_reason: str,
                 maintain_time: str,
                 payment_type: str,
                 port: int,
                 public_connection_string: str,
                 public_port: str,
                 scale_out_statuses: Sequence['outputs.GetDbClustersClusterScaleOutStatusResult'],
                 storage_type: str,
                 support_backup: int,
                 support_https_port: bool,
                 support_mysql_port: bool,
                 vpc_cloud_instance_id: str,
                 vpc_id: str,
                 vswitch_id: str,
                 zone_id: str):
        """
        :param str ali_uid: Alibaba Cloud account Id.
        :param str bid: The ID of the business process flow.
        :param str category: Instance family values include: Basic: Basic edition; HighAvailability: high availability edition.
        :param str commodity_code: The Commodity Code of the DBCluster.
        :param str connection_string: Connection string.
        :param str create_time: The creation time of the resource.
        :param Sequence['GetDbClustersClusterDbClusterAccessWhiteListArgs'] db_cluster_access_white_lists: The db cluster access white list.
        :param str db_cluster_description: The DBCluster description.
        :param str db_cluster_network_type: The DBCluster network type.
        :param str db_cluster_type: The DBCluster type.
        :param str db_node_class: The node class of the DBCluster.
        :param str db_node_count: The node count of the DBCluster.
        :param str db_node_storage: The node storage of the DBCluster.
        :param str encryption_key: Key management service KMS key ID.
        :param str encryption_type: Currently only supports ECS disk encryption, with a value of CloudDisk, not encrypted when empty.
        :param str engine: The Engine of the DBCluster.
        :param str engine_version: The engine version of the DBCluster.
        :param str expire_time: The expiration time of the DBCluster.
        :param str id: The ID of the DBCluster.
        :param str is_expired: If the instance has expired.
        :param str lock_mode: The lock mode of the DBCluster.
        :param str lock_reason: Lock reason of the DBCluster.
        :param str maintain_time: Examples of the maintenance window, in the format of hh:mmZ-hh:mm Z.
        :param str payment_type: The payment type of the resource. Valid values: `PayAsYouGo`,`Subscription`.
        :param int port: Connection port.
        :param str public_connection_string: A public IP address for the connection.
        :param str public_port: Public network port.
        :param Sequence['GetDbClustersClusterScaleOutStatusArgs'] scale_out_statuses: Scale state.
        :param str storage_type: Storage type of DBCluster. Valid values: `cloud_essd`, `cloud_efficiency`, `cloud_essd_pl2`, `cloud_essd_pl3`.
        :param int support_backup: Support fallback scheme.
        :param bool support_https_port: The system supports http port number.
        :param bool support_mysql_port: Supports Mysql, and those of the ports.
        :param str vpc_cloud_instance_id: Virtual Private Cloud (VPC cloud instance ID.
        :param str vpc_id: The VPC ID of the DBCluster.
        :param str vswitch_id: The vswitch id of the DBCluster.
        :param str zone_id: The zone ID of the DBCluster.
        """
        pulumi.set(__self__, "ali_uid", ali_uid)
        pulumi.set(__self__, "bid", bid)
        pulumi.set(__self__, "category", category)
        pulumi.set(__self__, "commodity_code", commodity_code)
        pulumi.set(__self__, "connection_string", connection_string)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "db_cluster_access_white_lists", db_cluster_access_white_lists)
        pulumi.set(__self__, "db_cluster_description", db_cluster_description)
        pulumi.set(__self__, "db_cluster_id", db_cluster_id)
        pulumi.set(__self__, "db_cluster_network_type", db_cluster_network_type)
        pulumi.set(__self__, "db_cluster_type", db_cluster_type)
        pulumi.set(__self__, "db_node_class", db_node_class)
        pulumi.set(__self__, "db_node_count", db_node_count)
        pulumi.set(__self__, "db_node_storage", db_node_storage)
        pulumi.set(__self__, "encryption_key", encryption_key)
        pulumi.set(__self__, "encryption_type", encryption_type)
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "engine_version", engine_version)
        pulumi.set(__self__, "expire_time", expire_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "is_expired", is_expired)
        pulumi.set(__self__, "lock_mode", lock_mode)
        pulumi.set(__self__, "lock_reason", lock_reason)
        pulumi.set(__self__, "maintain_time", maintain_time)
        pulumi.set(__self__, "payment_type", payment_type)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "public_connection_string", public_connection_string)
        pulumi.set(__self__, "public_port", public_port)
        pulumi.set(__self__, "scale_out_statuses", scale_out_statuses)
        pulumi.set(__self__, "storage_type", storage_type)
        pulumi.set(__self__, "support_backup", support_backup)
        pulumi.set(__self__, "support_https_port", support_https_port)
        pulumi.set(__self__, "support_mysql_port", support_mysql_port)
        pulumi.set(__self__, "vpc_cloud_instance_id", vpc_cloud_instance_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vswitch_id", vswitch_id)
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="aliUid")
    def ali_uid(self) -> str:
        """
        Alibaba Cloud account Id.
        """
        return pulumi.get(self, "ali_uid")

    @property
    @pulumi.getter
    def bid(self) -> str:
        """
        The ID of the business process flow.
        """
        return pulumi.get(self, "bid")

    @property
    @pulumi.getter
    def category(self) -> str:
        """
        Instance family values include: Basic: Basic edition; HighAvailability: high availability edition.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="commodityCode")
    def commodity_code(self) -> str:
        """
        The Commodity Code of the DBCluster.
        """
        return pulumi.get(self, "commodity_code")

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> str:
        """
        Connection string.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation time of the resource.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dbClusterAccessWhiteLists")
    def db_cluster_access_white_lists(self) -> Sequence['outputs.GetDbClustersClusterDbClusterAccessWhiteListResult']:
        """
        The db cluster access white list.
        """
        return pulumi.get(self, "db_cluster_access_white_lists")

    @property
    @pulumi.getter(name="dbClusterDescription")
    def db_cluster_description(self) -> str:
        """
        The DBCluster description.
        """
        return pulumi.get(self, "db_cluster_description")

    @property
    @pulumi.getter(name="dbClusterId")
    def db_cluster_id(self) -> str:
        return pulumi.get(self, "db_cluster_id")

    @property
    @pulumi.getter(name="dbClusterNetworkType")
    def db_cluster_network_type(self) -> str:
        """
        The DBCluster network type.
        """
        return pulumi.get(self, "db_cluster_network_type")

    @property
    @pulumi.getter(name="dbClusterType")
    def db_cluster_type(self) -> str:
        """
        The DBCluster type.
        """
        return pulumi.get(self, "db_cluster_type")

    @property
    @pulumi.getter(name="dbNodeClass")
    def db_node_class(self) -> str:
        """
        The node class of the DBCluster.
        """
        return pulumi.get(self, "db_node_class")

    @property
    @pulumi.getter(name="dbNodeCount")
    def db_node_count(self) -> str:
        """
        The node count of the DBCluster.
        """
        return pulumi.get(self, "db_node_count")

    @property
    @pulumi.getter(name="dbNodeStorage")
    def db_node_storage(self) -> str:
        """
        The node storage of the DBCluster.
        """
        return pulumi.get(self, "db_node_storage")

    @property
    @pulumi.getter(name="encryptionKey")
    def encryption_key(self) -> str:
        """
        Key management service KMS key ID.
        """
        return pulumi.get(self, "encryption_key")

    @property
    @pulumi.getter(name="encryptionType")
    def encryption_type(self) -> str:
        """
        Currently only supports ECS disk encryption, with a value of CloudDisk, not encrypted when empty.
        """
        return pulumi.get(self, "encryption_type")

    @property
    @pulumi.getter
    def engine(self) -> str:
        """
        The Engine of the DBCluster.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> str:
        """
        The engine version of the DBCluster.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        """
        The expiration time of the DBCluster.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the DBCluster.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isExpired")
    def is_expired(self) -> str:
        """
        If the instance has expired.
        """
        return pulumi.get(self, "is_expired")

    @property
    @pulumi.getter(name="lockMode")
    def lock_mode(self) -> str:
        """
        The lock mode of the DBCluster.
        """
        return pulumi.get(self, "lock_mode")

    @property
    @pulumi.getter(name="lockReason")
    def lock_reason(self) -> str:
        """
        Lock reason of the DBCluster.
        """
        return pulumi.get(self, "lock_reason")

    @property
    @pulumi.getter(name="maintainTime")
    def maintain_time(self) -> str:
        """
        Examples of the maintenance window, in the format of hh:mmZ-hh:mm Z.
        """
        return pulumi.get(self, "maintain_time")

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> str:
        """
        The payment type of the resource. Valid values: `PayAsYouGo`,`Subscription`.
        """
        return pulumi.get(self, "payment_type")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        Connection port.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="publicConnectionString")
    def public_connection_string(self) -> str:
        """
        A public IP address for the connection.
        """
        return pulumi.get(self, "public_connection_string")

    @property
    @pulumi.getter(name="publicPort")
    def public_port(self) -> str:
        """
        Public network port.
        """
        return pulumi.get(self, "public_port")

    @property
    @pulumi.getter(name="scaleOutStatuses")
    def scale_out_statuses(self) -> Sequence['outputs.GetDbClustersClusterScaleOutStatusResult']:
        """
        Scale state.
        """
        return pulumi.get(self, "scale_out_statuses")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> str:
        """
        Storage type of DBCluster. Valid values: `cloud_essd`, `cloud_efficiency`, `cloud_essd_pl2`, `cloud_essd_pl3`.
        """
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter(name="supportBackup")
    def support_backup(self) -> int:
        """
        Support fallback scheme.
        """
        return pulumi.get(self, "support_backup")

    @property
    @pulumi.getter(name="supportHttpsPort")
    def support_https_port(self) -> bool:
        """
        The system supports http port number.
        """
        return pulumi.get(self, "support_https_port")

    @property
    @pulumi.getter(name="supportMysqlPort")
    def support_mysql_port(self) -> bool:
        """
        Supports Mysql, and those of the ports.
        """
        return pulumi.get(self, "support_mysql_port")

    @property
    @pulumi.getter(name="vpcCloudInstanceId")
    def vpc_cloud_instance_id(self) -> str:
        """
        Virtual Private Cloud (VPC cloud instance ID.
        """
        return pulumi.get(self, "vpc_cloud_instance_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The VPC ID of the DBCluster.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> str:
        """
        The vswitch id of the DBCluster.
        """
        return pulumi.get(self, "vswitch_id")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> str:
        """
        The zone ID of the DBCluster.
        """
        return pulumi.get(self, "zone_id")


@pulumi.output_type
class GetDbClustersClusterDbClusterAccessWhiteListResult(dict):
    def __init__(__self__, *,
                 db_cluster_ip_array_attribute: str,
                 db_cluster_ip_array_name: str,
                 security_ip_list: str):
        """
        :param str db_cluster_ip_array_attribute: Whitelist grouping attribute.
        :param str db_cluster_ip_array_name: Whitelist group name.
        :param str security_ip_list: The IP address list under the whitelist group.
        """
        pulumi.set(__self__, "db_cluster_ip_array_attribute", db_cluster_ip_array_attribute)
        pulumi.set(__self__, "db_cluster_ip_array_name", db_cluster_ip_array_name)
        pulumi.set(__self__, "security_ip_list", security_ip_list)

    @property
    @pulumi.getter(name="dbClusterIpArrayAttribute")
    def db_cluster_ip_array_attribute(self) -> str:
        """
        Whitelist grouping attribute.
        """
        return pulumi.get(self, "db_cluster_ip_array_attribute")

    @property
    @pulumi.getter(name="dbClusterIpArrayName")
    def db_cluster_ip_array_name(self) -> str:
        """
        Whitelist group name.
        """
        return pulumi.get(self, "db_cluster_ip_array_name")

    @property
    @pulumi.getter(name="securityIpList")
    def security_ip_list(self) -> str:
        """
        The IP address list under the whitelist group.
        """
        return pulumi.get(self, "security_ip_list")


@pulumi.output_type
class GetDbClustersClusterScaleOutStatusResult(dict):
    def __init__(__self__, *,
                 progress: str,
                 ratio: str):
        """
        :param str progress: Process.
        :param str ratio: Efficiency.
        """
        pulumi.set(__self__, "progress", progress)
        pulumi.set(__self__, "ratio", ratio)

    @property
    @pulumi.getter
    def progress(self) -> str:
        """
        Process.
        """
        return pulumi.get(self, "progress")

    @property
    @pulumi.getter
    def ratio(self) -> str:
        """
        Efficiency.
        """
        return pulumi.get(self, "ratio")


@pulumi.output_type
class GetRegionsRegionResult(dict):
    def __init__(__self__, *,
                 region_id: str,
                 zone_ids: Sequence['outputs.GetRegionsRegionZoneIdResult']):
        """
        :param str region_id: The Region ID.
        :param Sequence['GetRegionsRegionZoneIdArgs'] zone_ids: A list of available zone ids in the region_id.
        """
        pulumi.set(__self__, "region_id", region_id)
        pulumi.set(__self__, "zone_ids", zone_ids)

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> str:
        """
        The Region ID.
        """
        return pulumi.get(self, "region_id")

    @property
    @pulumi.getter(name="zoneIds")
    def zone_ids(self) -> Sequence['outputs.GetRegionsRegionZoneIdResult']:
        """
        A list of available zone ids in the region_id.
        """
        return pulumi.get(self, "zone_ids")


@pulumi.output_type
class GetRegionsRegionZoneIdResult(dict):
    def __init__(__self__, *,
                 vpc_enabled: bool,
                 zone_id: str):
        """
        :param bool vpc_enabled: Whether to support vpc network.
        :param str zone_id: The zone ID.
        """
        pulumi.set(__self__, "vpc_enabled", vpc_enabled)
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="vpcEnabled")
    def vpc_enabled(self) -> bool:
        """
        Whether to support vpc network.
        """
        return pulumi.get(self, "vpc_enabled")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> str:
        """
        The zone ID.
        """
        return pulumi.get(self, "zone_id")


