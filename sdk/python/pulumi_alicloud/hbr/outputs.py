# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetEcsBackupClientsClientResult',
    'GetEcsBackupPlansPlanResult',
    'GetNasBackupPlansPlanResult',
    'GetOssBackupPlansPlanResult',
    'GetVaultsVaultResult',
]

@pulumi.output_type
class GetEcsBackupClientsClientResult(dict):
    def __init__(__self__, *,
                 arch_type: str,
                 backup_status: str,
                 client_type: str,
                 client_version: str,
                 create_time: str,
                 data_network_type: str,
                 data_proxy_setting: str,
                 ecs_backup_client_id: str,
                 hostname: str,
                 id: str,
                 instance_id: str,
                 instance_name: str,
                 last_heart_beat_time: str,
                 max_client_version: str,
                 max_cpu_core: str,
                 max_worker: str,
                 os_type: str,
                 private_ipv4: str,
                 proxy_host: str,
                 proxy_password: str,
                 proxy_port: str,
                 proxy_user: str,
                 status: str,
                 updated_time: str,
                 use_https: bool,
                 zone_id: str):
        """
        :param str arch_type: The Client System Architecture (Only the ECS File Backup Client Is Available. Possible Values: * AMD64 * 386.
        :param str backup_status: Client protected status.
        :param str client_type: The Client Type. Possible Values: * ECS_CLIENT (ECS File Backup Client).
        :param str client_version: Client Version.
        :param str create_time: The Client Creates a Time. Unix Time Seconds.
        :param str data_network_type: The Data Plane Data Access Point Type. Valid Values: * Public Internet * VPC VPC * Classic Network.
        :param str data_proxy_setting: The Data Plane Proxy Settings. Valid Values: * DISABLE  * USE_CONTROL_PROXY (Default, the same with Control Plane) * CUSTOM (Custom Configuration Items for the HTTP Protocol).
        :param str ecs_backup_client_id: The first ID of the resource.
        :param str hostname: The ECS Host Name.
        :param str id: The ID of the Ecs Backup Client.
        :param str instance_id: The ID of ECS Instance.
        :param str instance_name: ECS Instance Names.
        :param str last_heart_beat_time: Client Last Heartbeat Time. Unix Time Seconds.
        :param str max_client_version: The Latest Client Version.
        :param str max_cpu_core: A Single Backup Task Uses for Example, Instances Can Be Grouped According to CPU Core Count, 0 Means No Restrictions.
        :param str max_worker: A Single Backup Task Parallel Work, the Number of 0 Means No Restrictions.
        :param str os_type: The Client System Type (Only the ECS File Backup Client Is Available. Possible Values: * windows * linux.
        :param str private_ipv4: Instance Must Not Use the Intranet IP Address.
        :param str proxy_host: Custom Data Plane Proxy Server Host Address.
        :param str proxy_password: Custom Data Plane Proxy Password.
        :param str proxy_port: Custom Data Plane Proxy Server Host Port.
        :param str proxy_user: Custom Data Plane Proxy Server User Name.
        :param str status: The status of the resource.
        :param str updated_time: Client Update Time. Unix Time Seconds.
        :param bool use_https: Indicates Whether to Use the Https Transport Data Plane Data.
        :param str zone_id: The Zone ID.
        """
        pulumi.set(__self__, "arch_type", arch_type)
        pulumi.set(__self__, "backup_status", backup_status)
        pulumi.set(__self__, "client_type", client_type)
        pulumi.set(__self__, "client_version", client_version)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "data_network_type", data_network_type)
        pulumi.set(__self__, "data_proxy_setting", data_proxy_setting)
        pulumi.set(__self__, "ecs_backup_client_id", ecs_backup_client_id)
        pulumi.set(__self__, "hostname", hostname)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "last_heart_beat_time", last_heart_beat_time)
        pulumi.set(__self__, "max_client_version", max_client_version)
        pulumi.set(__self__, "max_cpu_core", max_cpu_core)
        pulumi.set(__self__, "max_worker", max_worker)
        pulumi.set(__self__, "os_type", os_type)
        pulumi.set(__self__, "private_ipv4", private_ipv4)
        pulumi.set(__self__, "proxy_host", proxy_host)
        pulumi.set(__self__, "proxy_password", proxy_password)
        pulumi.set(__self__, "proxy_port", proxy_port)
        pulumi.set(__self__, "proxy_user", proxy_user)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "updated_time", updated_time)
        pulumi.set(__self__, "use_https", use_https)
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="archType")
    def arch_type(self) -> str:
        """
        The Client System Architecture (Only the ECS File Backup Client Is Available. Possible Values: * AMD64 * 386.
        """
        return pulumi.get(self, "arch_type")

    @property
    @pulumi.getter(name="backupStatus")
    def backup_status(self) -> str:
        """
        Client protected status.
        """
        return pulumi.get(self, "backup_status")

    @property
    @pulumi.getter(name="clientType")
    def client_type(self) -> str:
        """
        The Client Type. Possible Values: * ECS_CLIENT (ECS File Backup Client).
        """
        return pulumi.get(self, "client_type")

    @property
    @pulumi.getter(name="clientVersion")
    def client_version(self) -> str:
        """
        Client Version.
        """
        return pulumi.get(self, "client_version")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The Client Creates a Time. Unix Time Seconds.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dataNetworkType")
    def data_network_type(self) -> str:
        """
        The Data Plane Data Access Point Type. Valid Values: * Public Internet * VPC VPC * Classic Network.
        """
        return pulumi.get(self, "data_network_type")

    @property
    @pulumi.getter(name="dataProxySetting")
    def data_proxy_setting(self) -> str:
        """
        The Data Plane Proxy Settings. Valid Values: * DISABLE  * USE_CONTROL_PROXY (Default, the same with Control Plane) * CUSTOM (Custom Configuration Items for the HTTP Protocol).
        """
        return pulumi.get(self, "data_proxy_setting")

    @property
    @pulumi.getter(name="ecsBackupClientId")
    def ecs_backup_client_id(self) -> str:
        """
        The first ID of the resource.
        """
        return pulumi.get(self, "ecs_backup_client_id")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        """
        The ECS Host Name.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Ecs Backup Client.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        The ID of ECS Instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        """
        ECS Instance Names.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="lastHeartBeatTime")
    def last_heart_beat_time(self) -> str:
        """
        Client Last Heartbeat Time. Unix Time Seconds.
        """
        return pulumi.get(self, "last_heart_beat_time")

    @property
    @pulumi.getter(name="maxClientVersion")
    def max_client_version(self) -> str:
        """
        The Latest Client Version.
        """
        return pulumi.get(self, "max_client_version")

    @property
    @pulumi.getter(name="maxCpuCore")
    def max_cpu_core(self) -> str:
        """
        A Single Backup Task Uses for Example, Instances Can Be Grouped According to CPU Core Count, 0 Means No Restrictions.
        """
        return pulumi.get(self, "max_cpu_core")

    @property
    @pulumi.getter(name="maxWorker")
    def max_worker(self) -> str:
        """
        A Single Backup Task Parallel Work, the Number of 0 Means No Restrictions.
        """
        return pulumi.get(self, "max_worker")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> str:
        """
        The Client System Type (Only the ECS File Backup Client Is Available. Possible Values: * windows * linux.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="privateIpv4")
    def private_ipv4(self) -> str:
        """
        Instance Must Not Use the Intranet IP Address.
        """
        return pulumi.get(self, "private_ipv4")

    @property
    @pulumi.getter(name="proxyHost")
    def proxy_host(self) -> str:
        """
        Custom Data Plane Proxy Server Host Address.
        """
        return pulumi.get(self, "proxy_host")

    @property
    @pulumi.getter(name="proxyPassword")
    def proxy_password(self) -> str:
        """
        Custom Data Plane Proxy Password.
        """
        return pulumi.get(self, "proxy_password")

    @property
    @pulumi.getter(name="proxyPort")
    def proxy_port(self) -> str:
        """
        Custom Data Plane Proxy Server Host Port.
        """
        return pulumi.get(self, "proxy_port")

    @property
    @pulumi.getter(name="proxyUser")
    def proxy_user(self) -> str:
        """
        Custom Data Plane Proxy Server User Name.
        """
        return pulumi.get(self, "proxy_user")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedTime")
    def updated_time(self) -> str:
        """
        Client Update Time. Unix Time Seconds.
        """
        return pulumi.get(self, "updated_time")

    @property
    @pulumi.getter(name="useHttps")
    def use_https(self) -> bool:
        """
        Indicates Whether to Use the Https Transport Data Plane Data.
        """
        return pulumi.get(self, "use_https")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> str:
        """
        The Zone ID.
        """
        return pulumi.get(self, "zone_id")


@pulumi.output_type
class GetEcsBackupPlansPlanResult(dict):
    def __init__(__self__, *,
                 backup_type: str,
                 create_time: str,
                 detail: str,
                 disabled: bool,
                 ecs_backup_plan_id: str,
                 ecs_backup_plan_name: str,
                 exclude: str,
                 id: str,
                 include: str,
                 instance_id: str,
                 options: str,
                 paths: Sequence[str],
                 retention: str,
                 schedule: str,
                 speed_limit: str,
                 vault_id: str):
        """
        :param str backup_type: Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
        :param str ecs_backup_plan_name: The Configuration Page of a Backup Plan Name. 1-64 Characters, requiring a Single Warehouse under Each of the Data Source Type Drop-down List of the Configuration Page of a Backup Plan Name Is Unique.
        :param str exclude: Exclude Path. String of Json List, most 255 Characters. e.g. `"[\"/home/work\"]"`
        :param str include: Include Path. String of Json List, most 255 Characters. e.g. `"[\"/var\"]"`
        :param str instance_id: The ECS Instance Id. Must Have Installed the Client.
        :param str options: Windows System with Application Consistency Using VSS. eg: {`UseVSS`:false}.
        :param Sequence[str] paths: Backup Path. e.g. `["/home", "/var"]`
        :param str retention: Backup Retention Period, the Minimum Value of 1.
        :param str schedule: Backup strategy. Optional format: I|{startTime}|{interval} * startTime Backup start time, UNIX time, in seconds. * interval ISO8601 time interval. E.g: ** PT1H, one hour apart. ** P1D, one day apart. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed, the next backup task will not be triggered.
        :param str speed_limit: flow control. The format is: {start}|{end}|{bandwidth} * start starting hour * end end hour * bandwidth limit rate, in KiB ** Use | to separate multiple flow control configurations; ** Multiple flow control configurations are not allowed to have overlapping times.
        :param str vault_id: Vault ID.
        """
        pulumi.set(__self__, "backup_type", backup_type)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "detail", detail)
        pulumi.set(__self__, "disabled", disabled)
        pulumi.set(__self__, "ecs_backup_plan_id", ecs_backup_plan_id)
        pulumi.set(__self__, "ecs_backup_plan_name", ecs_backup_plan_name)
        pulumi.set(__self__, "exclude", exclude)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "include", include)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "options", options)
        pulumi.set(__self__, "paths", paths)
        pulumi.set(__self__, "retention", retention)
        pulumi.set(__self__, "schedule", schedule)
        pulumi.set(__self__, "speed_limit", speed_limit)
        pulumi.set(__self__, "vault_id", vault_id)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> str:
        """
        Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
        """
        return pulumi.get(self, "backup_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def detail(self) -> str:
        return pulumi.get(self, "detail")

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="ecsBackupPlanId")
    def ecs_backup_plan_id(self) -> str:
        return pulumi.get(self, "ecs_backup_plan_id")

    @property
    @pulumi.getter(name="ecsBackupPlanName")
    def ecs_backup_plan_name(self) -> str:
        """
        The Configuration Page of a Backup Plan Name. 1-64 Characters, requiring a Single Warehouse under Each of the Data Source Type Drop-down List of the Configuration Page of a Backup Plan Name Is Unique.
        """
        return pulumi.get(self, "ecs_backup_plan_name")

    @property
    @pulumi.getter
    def exclude(self) -> str:
        """
        Exclude Path. String of Json List, most 255 Characters. e.g. `"[\"/home/work\"]"`
        """
        return pulumi.get(self, "exclude")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def include(self) -> str:
        """
        Include Path. String of Json List, most 255 Characters. e.g. `"[\"/var\"]"`
        """
        return pulumi.get(self, "include")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        The ECS Instance Id. Must Have Installed the Client.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def options(self) -> str:
        """
        Windows System with Application Consistency Using VSS. eg: {`UseVSS`:false}.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter
    def paths(self) -> Sequence[str]:
        """
        Backup Path. e.g. `["/home", "/var"]`
        """
        return pulumi.get(self, "paths")

    @property
    @pulumi.getter
    def retention(self) -> str:
        """
        Backup Retention Period, the Minimum Value of 1.
        """
        return pulumi.get(self, "retention")

    @property
    @pulumi.getter
    def schedule(self) -> str:
        """
        Backup strategy. Optional format: I|{startTime}|{interval} * startTime Backup start time, UNIX time, in seconds. * interval ISO8601 time interval. E.g: ** PT1H, one hour apart. ** P1D, one day apart. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed, the next backup task will not be triggered.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="speedLimit")
    def speed_limit(self) -> str:
        """
        flow control. The format is: {start}|{end}|{bandwidth} * start starting hour * end end hour * bandwidth limit rate, in KiB ** Use | to separate multiple flow control configurations; ** Multiple flow control configurations are not allowed to have overlapping times.
        """
        return pulumi.get(self, "speed_limit")

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> str:
        """
        Vault ID.
        """
        return pulumi.get(self, "vault_id")


@pulumi.output_type
class GetNasBackupPlansPlanResult(dict):
    def __init__(__self__, *,
                 backup_type: str,
                 create_time: str,
                 disabled: bool,
                 file_system_id: str,
                 id: str,
                 nas_backup_plan_id: str,
                 nas_backup_plan_name: str,
                 options: str,
                 paths: Sequence[str],
                 retention: str,
                 schedule: str,
                 vault_id: str):
        """
        :param str backup_type: Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
        :param str create_time: File System Creation Time. Unix Time Seconds.
        :param str file_system_id: The File System ID.
        :param str nas_backup_plan_name: The name of the resource.
        :param str options: Options. NAS Backup Plan Does Not Support Yet.
        :param Sequence[str] paths: Backup Path. Up to 65536 Characters. e.g.`["/home", "/var"]`
        :param str retention: Backup Retention Period, the Minimum Value of 1.
        :param str schedule: The Backup Policy. Formats: I | {Range Specified by the Starttime }|{ Interval}\n* The Time Range Specified by the Starttime Backup Start Time in Unix Time Seconds.\n* Interval ISO8601 Time Intervals. For Example:\n**PT1H Interval for an Hour.\n**P1D Interval Day.\nMeaning from {Range Specified by the Starttime} Every {Interval} of the Time Where We Took Backups Once a Task. Does Not Compensate the Has Elapsed Time the Backup Task. If the Last Backup Has Not Been Completed without Triggering the next Backup.
        :param str vault_id: The Vault ID of the EcsBackupPlan used.
        """
        pulumi.set(__self__, "backup_type", backup_type)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "disabled", disabled)
        pulumi.set(__self__, "file_system_id", file_system_id)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "nas_backup_plan_id", nas_backup_plan_id)
        pulumi.set(__self__, "nas_backup_plan_name", nas_backup_plan_name)
        pulumi.set(__self__, "options", options)
        pulumi.set(__self__, "paths", paths)
        pulumi.set(__self__, "retention", retention)
        pulumi.set(__self__, "schedule", schedule)
        pulumi.set(__self__, "vault_id", vault_id)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> str:
        """
        Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
        """
        return pulumi.get(self, "backup_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        File System Creation Time. Unix Time Seconds.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> str:
        """
        The File System ID.
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="nasBackupPlanId")
    def nas_backup_plan_id(self) -> str:
        return pulumi.get(self, "nas_backup_plan_id")

    @property
    @pulumi.getter(name="nasBackupPlanName")
    def nas_backup_plan_name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "nas_backup_plan_name")

    @property
    @pulumi.getter
    def options(self) -> str:
        """
        Options. NAS Backup Plan Does Not Support Yet.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter
    def paths(self) -> Sequence[str]:
        """
        Backup Path. Up to 65536 Characters. e.g.`["/home", "/var"]`
        """
        return pulumi.get(self, "paths")

    @property
    @pulumi.getter
    def retention(self) -> str:
        """
        Backup Retention Period, the Minimum Value of 1.
        """
        return pulumi.get(self, "retention")

    @property
    @pulumi.getter
    def schedule(self) -> str:
        """
        The Backup Policy. Formats: I | {Range Specified by the Starttime }|{ Interval}\n* The Time Range Specified by the Starttime Backup Start Time in Unix Time Seconds.\n* Interval ISO8601 Time Intervals. For Example:\n**PT1H Interval for an Hour.\n**P1D Interval Day.\nMeaning from {Range Specified by the Starttime} Every {Interval} of the Time Where We Took Backups Once a Task. Does Not Compensate the Has Elapsed Time the Backup Task. If the Last Backup Has Not Been Completed without Triggering the next Backup.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> str:
        """
        The Vault ID of the EcsBackupPlan used.
        """
        return pulumi.get(self, "vault_id")


@pulumi.output_type
class GetOssBackupPlansPlanResult(dict):
    def __init__(__self__, *,
                 backup_type: str,
                 bucket: str,
                 disabled: bool,
                 id: str,
                 oss_backup_plan_id: str,
                 oss_backup_plan_name: str,
                 prefix: str,
                 retention: str,
                 schedule: str,
                 vault_id: str):
        """
        :param str backup_type: Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
        :param str bucket: The OSS Bucket Name.
        :param str oss_backup_plan_name: The Configuration Page of a Backup Plan Name. 1-64 Characters, requiring a Single Warehouse under Each of the Data Source Type Drop-down List of the Configuration Page of a Backup Plan Name Is Unique.
        :param str retention: Backup Retention Period, the Minimum Value of 1.
        :param str schedule: Backup strategy. Optional format: I|{startTime}|{interval} * startTime Backup start time, UNIX time, in seconds. * interval ISO8601 time interval. E.g: ** PT1H, one hour apart. ** P1D, one day apart. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed, the next backup task will not be triggered.
        :param str vault_id: Vault ID.
        """
        pulumi.set(__self__, "backup_type", backup_type)
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "disabled", disabled)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "oss_backup_plan_id", oss_backup_plan_id)
        pulumi.set(__self__, "oss_backup_plan_name", oss_backup_plan_name)
        pulumi.set(__self__, "prefix", prefix)
        pulumi.set(__self__, "retention", retention)
        pulumi.set(__self__, "schedule", schedule)
        pulumi.set(__self__, "vault_id", vault_id)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> str:
        """
        Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
        """
        return pulumi.get(self, "backup_type")

    @property
    @pulumi.getter
    def bucket(self) -> str:
        """
        The OSS Bucket Name.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ossBackupPlanId")
    def oss_backup_plan_id(self) -> str:
        return pulumi.get(self, "oss_backup_plan_id")

    @property
    @pulumi.getter(name="ossBackupPlanName")
    def oss_backup_plan_name(self) -> str:
        """
        The Configuration Page of a Backup Plan Name. 1-64 Characters, requiring a Single Warehouse under Each of the Data Source Type Drop-down List of the Configuration Page of a Backup Plan Name Is Unique.
        """
        return pulumi.get(self, "oss_backup_plan_name")

    @property
    @pulumi.getter
    def prefix(self) -> str:
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter
    def retention(self) -> str:
        """
        Backup Retention Period, the Minimum Value of 1.
        """
        return pulumi.get(self, "retention")

    @property
    @pulumi.getter
    def schedule(self) -> str:
        """
        Backup strategy. Optional format: I|{startTime}|{interval} * startTime Backup start time, UNIX time, in seconds. * interval ISO8601 time interval. E.g: ** PT1H, one hour apart. ** P1D, one day apart. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed, the next backup task will not be triggered.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> str:
        """
        Vault ID.
        """
        return pulumi.get(self, "vault_id")


@pulumi.output_type
class GetVaultsVaultResult(dict):
    def __init__(__self__, *,
                 bucket_name: str,
                 bytes_done: str,
                 create_time: str,
                 dedup: bool,
                 description: str,
                 id: str,
                 index_available: bool,
                 index_level: str,
                 index_update_time: str,
                 latest_replication_time: str,
                 payment_type: str,
                 replication: bool,
                 replication_source_region_id: str,
                 replication_source_vault_id: str,
                 retention: str,
                 search_enabled: bool,
                 source_types: Sequence[str],
                 status: str,
                 storage_size: str,
                 updated_time: str,
                 vault_id: str,
                 vault_name: str,
                 vault_status_message: str,
                 vault_storage_class: str,
                 vault_type: str):
        """
        :param str bucket_name: The name of the OSS bucket of the Vault.
        :param str bytes_done: The amount of backup data. The unit is Byte.
        :param str create_time: The creation time of the Vault. UNIX time in seconds.
        :param bool dedup: Whether to enable the deduplication function for the database backup Vault.
        :param str description: The description of the Vault.
        :param str id: The ID of Vault.
        :param bool index_available: Index available.
        :param str index_level: Index level.
        :param str index_update_time: Index update time.
        :param str latest_replication_time: The time of the last remote backup synchronization.
        :param str payment_type: Billing model, possible values:
        :param bool replication: Whether it is a remote backup warehouse. It's a boolean value.
        :param str replication_source_region_id: The region ID to which the remote backup Vault belongs.
        :param str replication_source_vault_id: The source vault ID of the remote backup Vault.
        :param str retention: Warehouse-level data retention days, only valid for archive libraries.
        :param bool search_enabled: Whether to enable the backup search function.
        :param str status: The status of Vault. Valid values: `INITIALIZING`, `CREATED`, `ERROR`, `UNKNOWN`.
        :param str storage_size: Backup vault storage usage. The unit is Byte.
        :param str updated_time: The update time of the Vault. UNIX time in seconds.
        :param str vault_id: The ID of Vault.
        :param str vault_name: The name of Vault.
        :param str vault_status_message: Error status information of Vault.
        :param str vault_storage_class: The storage class of Vault. Valid values: `STANDARD`.
        :param str vault_type: The type of Vault. Valid values: `STANDARD`.
        """
        pulumi.set(__self__, "bucket_name", bucket_name)
        pulumi.set(__self__, "bytes_done", bytes_done)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "dedup", dedup)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "index_available", index_available)
        pulumi.set(__self__, "index_level", index_level)
        pulumi.set(__self__, "index_update_time", index_update_time)
        pulumi.set(__self__, "latest_replication_time", latest_replication_time)
        pulumi.set(__self__, "payment_type", payment_type)
        pulumi.set(__self__, "replication", replication)
        pulumi.set(__self__, "replication_source_region_id", replication_source_region_id)
        pulumi.set(__self__, "replication_source_vault_id", replication_source_vault_id)
        pulumi.set(__self__, "retention", retention)
        pulumi.set(__self__, "search_enabled", search_enabled)
        pulumi.set(__self__, "source_types", source_types)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "storage_size", storage_size)
        pulumi.set(__self__, "updated_time", updated_time)
        pulumi.set(__self__, "vault_id", vault_id)
        pulumi.set(__self__, "vault_name", vault_name)
        pulumi.set(__self__, "vault_status_message", vault_status_message)
        pulumi.set(__self__, "vault_storage_class", vault_storage_class)
        pulumi.set(__self__, "vault_type", vault_type)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> str:
        """
        The name of the OSS bucket of the Vault.
        """
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter(name="bytesDone")
    def bytes_done(self) -> str:
        """
        The amount of backup data. The unit is Byte.
        """
        return pulumi.get(self, "bytes_done")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation time of the Vault. UNIX time in seconds.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def dedup(self) -> bool:
        """
        Whether to enable the deduplication function for the database backup Vault.
        """
        return pulumi.get(self, "dedup")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the Vault.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of Vault.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="indexAvailable")
    def index_available(self) -> bool:
        """
        Index available.
        """
        return pulumi.get(self, "index_available")

    @property
    @pulumi.getter(name="indexLevel")
    def index_level(self) -> str:
        """
        Index level.
        """
        return pulumi.get(self, "index_level")

    @property
    @pulumi.getter(name="indexUpdateTime")
    def index_update_time(self) -> str:
        """
        Index update time.
        """
        return pulumi.get(self, "index_update_time")

    @property
    @pulumi.getter(name="latestReplicationTime")
    def latest_replication_time(self) -> str:
        """
        The time of the last remote backup synchronization.
        """
        return pulumi.get(self, "latest_replication_time")

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> str:
        """
        Billing model, possible values:
        """
        return pulumi.get(self, "payment_type")

    @property
    @pulumi.getter
    def replication(self) -> bool:
        """
        Whether it is a remote backup warehouse. It's a boolean value.
        """
        return pulumi.get(self, "replication")

    @property
    @pulumi.getter(name="replicationSourceRegionId")
    def replication_source_region_id(self) -> str:
        """
        The region ID to which the remote backup Vault belongs.
        """
        return pulumi.get(self, "replication_source_region_id")

    @property
    @pulumi.getter(name="replicationSourceVaultId")
    def replication_source_vault_id(self) -> str:
        """
        The source vault ID of the remote backup Vault.
        """
        return pulumi.get(self, "replication_source_vault_id")

    @property
    @pulumi.getter
    def retention(self) -> str:
        """
        Warehouse-level data retention days, only valid for archive libraries.
        """
        return pulumi.get(self, "retention")

    @property
    @pulumi.getter(name="searchEnabled")
    def search_enabled(self) -> bool:
        """
        Whether to enable the backup search function.
        """
        return pulumi.get(self, "search_enabled")

    @property
    @pulumi.getter(name="sourceTypes")
    def source_types(self) -> Sequence[str]:
        return pulumi.get(self, "source_types")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of Vault. Valid values: `INITIALIZING`, `CREATED`, `ERROR`, `UNKNOWN`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> str:
        """
        Backup vault storage usage. The unit is Byte.
        """
        return pulumi.get(self, "storage_size")

    @property
    @pulumi.getter(name="updatedTime")
    def updated_time(self) -> str:
        """
        The update time of the Vault. UNIX time in seconds.
        """
        return pulumi.get(self, "updated_time")

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> str:
        """
        The ID of Vault.
        """
        return pulumi.get(self, "vault_id")

    @property
    @pulumi.getter(name="vaultName")
    def vault_name(self) -> str:
        """
        The name of Vault.
        """
        return pulumi.get(self, "vault_name")

    @property
    @pulumi.getter(name="vaultStatusMessage")
    def vault_status_message(self) -> str:
        """
        Error status information of Vault.
        """
        return pulumi.get(self, "vault_status_message")

    @property
    @pulumi.getter(name="vaultStorageClass")
    def vault_storage_class(self) -> str:
        """
        The storage class of Vault. Valid values: `STANDARD`.
        """
        return pulumi.get(self, "vault_storage_class")

    @property
    @pulumi.getter(name="vaultType")
    def vault_type(self) -> str:
        """
        The type of Vault. Valid values: `STANDARD`.
        """
        return pulumi.get(self, "vault_type")


