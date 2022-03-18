# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NasBackupPlanArgs', 'NasBackupPlan']

@pulumi.input_type
class NasBackupPlanArgs:
    def __init__(__self__, *,
                 backup_type: pulumi.Input[str],
                 file_system_id: pulumi.Input[str],
                 nas_backup_plan_name: pulumi.Input[str],
                 paths: pulumi.Input[Sequence[pulumi.Input[str]]],
                 retention: pulumi.Input[str],
                 schedule: pulumi.Input[str],
                 vault_id: pulumi.Input[str],
                 create_time: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 options: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NasBackupPlan resource.
        :param pulumi.Input[str] backup_type: Backup type. Valid values: `COMPLETE`.
        :param pulumi.Input[str] file_system_id: The File System ID of Nas.
        :param pulumi.Input[str] nas_backup_plan_name: The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] paths: List of backup path. Up to 65536 characters. e.g.`["/home", "/var"]`. **Note** You should at least specify a backup path, empty array not allowed here.
        :param pulumi.Input[str] retention: Backup retention days, the minimum is 1.
        :param pulumi.Input[str] schedule: Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
        :param pulumi.Input[str] vault_id: The ID of Backup vault.
        :param pulumi.Input[str] create_time: This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
        :param pulumi.Input[bool] disabled: Whether to disable the backup task. Valid values: `true`, `false`.
        """
        pulumi.set(__self__, "backup_type", backup_type)
        pulumi.set(__self__, "file_system_id", file_system_id)
        pulumi.set(__self__, "nas_backup_plan_name", nas_backup_plan_name)
        pulumi.set(__self__, "paths", paths)
        pulumi.set(__self__, "retention", retention)
        pulumi.set(__self__, "schedule", schedule)
        pulumi.set(__self__, "vault_id", vault_id)
        if create_time is not None:
            warnings.warn("""Field 'create_time' has been deprecated from provider version 1.153.0.""", DeprecationWarning)
            pulumi.log.warn("""create_time is deprecated: Field 'create_time' has been deprecated from provider version 1.153.0.""")
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if options is not None:
            pulumi.set(__self__, "options", options)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> pulumi.Input[str]:
        """
        Backup type. Valid values: `COMPLETE`.
        """
        return pulumi.get(self, "backup_type")

    @backup_type.setter
    def backup_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_type", value)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Input[str]:
        """
        The File System ID of Nas.
        """
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter(name="nasBackupPlanName")
    def nas_backup_plan_name(self) -> pulumi.Input[str]:
        """
        The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
        """
        return pulumi.get(self, "nas_backup_plan_name")

    @nas_backup_plan_name.setter
    def nas_backup_plan_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "nas_backup_plan_name", value)

    @property
    @pulumi.getter
    def paths(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of backup path. Up to 65536 characters. e.g.`["/home", "/var"]`. **Note** You should at least specify a backup path, empty array not allowed here.
        """
        return pulumi.get(self, "paths")

    @paths.setter
    def paths(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "paths", value)

    @property
    @pulumi.getter
    def retention(self) -> pulumi.Input[str]:
        """
        Backup retention days, the minimum is 1.
        """
        return pulumi.get(self, "retention")

    @retention.setter
    def retention(self, value: pulumi.Input[str]):
        pulumi.set(self, "retention", value)

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Input[str]:
        """
        Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: pulumi.Input[str]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> pulumi.Input[str]:
        """
        The ID of Backup vault.
        """
        return pulumi.get(self, "vault_id")

    @vault_id.setter
    def vault_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vault_id", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to disable the backup task. Valid values: `true`, `false`.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "options", value)


@pulumi.input_type
class _NasBackupPlanState:
    def __init__(__self__, *,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 nas_backup_plan_name: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[str]] = None,
                 paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 vault_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NasBackupPlan resources.
        :param pulumi.Input[str] backup_type: Backup type. Valid values: `COMPLETE`.
        :param pulumi.Input[str] create_time: This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
        :param pulumi.Input[bool] disabled: Whether to disable the backup task. Valid values: `true`, `false`.
        :param pulumi.Input[str] file_system_id: The File System ID of Nas.
        :param pulumi.Input[str] nas_backup_plan_name: The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] paths: List of backup path. Up to 65536 characters. e.g.`["/home", "/var"]`. **Note** You should at least specify a backup path, empty array not allowed here.
        :param pulumi.Input[str] retention: Backup retention days, the minimum is 1.
        :param pulumi.Input[str] schedule: Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
        :param pulumi.Input[str] vault_id: The ID of Backup vault.
        """
        if backup_type is not None:
            pulumi.set(__self__, "backup_type", backup_type)
        if create_time is not None:
            warnings.warn("""Field 'create_time' has been deprecated from provider version 1.153.0.""", DeprecationWarning)
            pulumi.log.warn("""create_time is deprecated: Field 'create_time' has been deprecated from provider version 1.153.0.""")
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if file_system_id is not None:
            pulumi.set(__self__, "file_system_id", file_system_id)
        if nas_backup_plan_name is not None:
            pulumi.set(__self__, "nas_backup_plan_name", nas_backup_plan_name)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if paths is not None:
            pulumi.set(__self__, "paths", paths)
        if retention is not None:
            pulumi.set(__self__, "retention", retention)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if vault_id is not None:
            pulumi.set(__self__, "vault_id", vault_id)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> Optional[pulumi.Input[str]]:
        """
        Backup type. Valid values: `COMPLETE`.
        """
        return pulumi.get(self, "backup_type")

    @backup_type.setter
    def backup_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_type", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to disable the backup task. Valid values: `true`, `false`.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> Optional[pulumi.Input[str]]:
        """
        The File System ID of Nas.
        """
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter(name="nasBackupPlanName")
    def nas_backup_plan_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
        """
        return pulumi.get(self, "nas_backup_plan_name")

    @nas_backup_plan_name.setter
    def nas_backup_plan_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nas_backup_plan_name", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter
    def paths(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of backup path. Up to 65536 characters. e.g.`["/home", "/var"]`. **Note** You should at least specify a backup path, empty array not allowed here.
        """
        return pulumi.get(self, "paths")

    @paths.setter
    def paths(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "paths", value)

    @property
    @pulumi.getter
    def retention(self) -> Optional[pulumi.Input[str]]:
        """
        Backup retention days, the minimum is 1.
        """
        return pulumi.get(self, "retention")

    @retention.setter
    def retention(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "retention", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[str]]:
        """
        Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of Backup vault.
        """
        return pulumi.get(self, "vault_id")

    @vault_id.setter
    def vault_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vault_id", value)


class NasBackupPlan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 nas_backup_plan_name: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[str]] = None,
                 paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 vault_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a HBR Nas Backup Plan resource.

        For information about HBR Nas Backup Plan and how to use it, see [What is Nas Backup Plan](https://www.alibabacloud.com/help/doc-detail/132248.htm).

        > **NOTE:** Available in v1.132.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-testAccHBRNas"
        default_vault = alicloud.hbr.Vault("defaultVault", vault_name=name)
        default_file_system = alicloud.nas.FileSystem("defaultFileSystem",
            protocol_type="NFS",
            storage_type="Performance",
            description=name,
            encrypt_type=1)
        default_file_systems = default_file_system.description.apply(lambda description: alicloud.nas.get_file_systems_output(protocol_type="NFS",
            description_regex=description))
        default_nas_backup_plan = alicloud.hbr.NasBackupPlan("defaultNasBackupPlan",
            nas_backup_plan_name=name,
            file_system_id=default_file_system.id,
            schedule="I|1602673264|PT2H",
            backup_type="COMPLETE",
            vault_id=default_vault.id,
            create_time=default_file_systems.systems[0].create_time,
            retention="2",
            paths=["/"],
            opts=pulumi.ResourceOptions(depends_on=["alicloud_nas_file_system.default"]))
        ```

        ## Import

        HBR Nas Backup Plan can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:hbr/nasBackupPlan:NasBackupPlan example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_type: Backup type. Valid values: `COMPLETE`.
        :param pulumi.Input[str] create_time: This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
        :param pulumi.Input[bool] disabled: Whether to disable the backup task. Valid values: `true`, `false`.
        :param pulumi.Input[str] file_system_id: The File System ID of Nas.
        :param pulumi.Input[str] nas_backup_plan_name: The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] paths: List of backup path. Up to 65536 characters. e.g.`["/home", "/var"]`. **Note** You should at least specify a backup path, empty array not allowed here.
        :param pulumi.Input[str] retention: Backup retention days, the minimum is 1.
        :param pulumi.Input[str] schedule: Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
        :param pulumi.Input[str] vault_id: The ID of Backup vault.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NasBackupPlanArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a HBR Nas Backup Plan resource.

        For information about HBR Nas Backup Plan and how to use it, see [What is Nas Backup Plan](https://www.alibabacloud.com/help/doc-detail/132248.htm).

        > **NOTE:** Available in v1.132.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-testAccHBRNas"
        default_vault = alicloud.hbr.Vault("defaultVault", vault_name=name)
        default_file_system = alicloud.nas.FileSystem("defaultFileSystem",
            protocol_type="NFS",
            storage_type="Performance",
            description=name,
            encrypt_type=1)
        default_file_systems = default_file_system.description.apply(lambda description: alicloud.nas.get_file_systems_output(protocol_type="NFS",
            description_regex=description))
        default_nas_backup_plan = alicloud.hbr.NasBackupPlan("defaultNasBackupPlan",
            nas_backup_plan_name=name,
            file_system_id=default_file_system.id,
            schedule="I|1602673264|PT2H",
            backup_type="COMPLETE",
            vault_id=default_vault.id,
            create_time=default_file_systems.systems[0].create_time,
            retention="2",
            paths=["/"],
            opts=pulumi.ResourceOptions(depends_on=["alicloud_nas_file_system.default"]))
        ```

        ## Import

        HBR Nas Backup Plan can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:hbr/nasBackupPlan:NasBackupPlan example <id>
        ```

        :param str resource_name: The name of the resource.
        :param NasBackupPlanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NasBackupPlanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 nas_backup_plan_name: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[str]] = None,
                 paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 vault_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NasBackupPlanArgs.__new__(NasBackupPlanArgs)

            if backup_type is None and not opts.urn:
                raise TypeError("Missing required property 'backup_type'")
            __props__.__dict__["backup_type"] = backup_type
            if create_time is not None and not opts.urn:
                warnings.warn("""Field 'create_time' has been deprecated from provider version 1.153.0.""", DeprecationWarning)
                pulumi.log.warn("""create_time is deprecated: Field 'create_time' has been deprecated from provider version 1.153.0.""")
            __props__.__dict__["create_time"] = create_time
            __props__.__dict__["disabled"] = disabled
            if file_system_id is None and not opts.urn:
                raise TypeError("Missing required property 'file_system_id'")
            __props__.__dict__["file_system_id"] = file_system_id
            if nas_backup_plan_name is None and not opts.urn:
                raise TypeError("Missing required property 'nas_backup_plan_name'")
            __props__.__dict__["nas_backup_plan_name"] = nas_backup_plan_name
            __props__.__dict__["options"] = options
            if paths is None and not opts.urn:
                raise TypeError("Missing required property 'paths'")
            __props__.__dict__["paths"] = paths
            if retention is None and not opts.urn:
                raise TypeError("Missing required property 'retention'")
            __props__.__dict__["retention"] = retention
            if schedule is None and not opts.urn:
                raise TypeError("Missing required property 'schedule'")
            __props__.__dict__["schedule"] = schedule
            if vault_id is None and not opts.urn:
                raise TypeError("Missing required property 'vault_id'")
            __props__.__dict__["vault_id"] = vault_id
        super(NasBackupPlan, __self__).__init__(
            'alicloud:hbr/nasBackupPlan:NasBackupPlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_type: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            file_system_id: Optional[pulumi.Input[str]] = None,
            nas_backup_plan_name: Optional[pulumi.Input[str]] = None,
            options: Optional[pulumi.Input[str]] = None,
            paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            retention: Optional[pulumi.Input[str]] = None,
            schedule: Optional[pulumi.Input[str]] = None,
            vault_id: Optional[pulumi.Input[str]] = None) -> 'NasBackupPlan':
        """
        Get an existing NasBackupPlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_type: Backup type. Valid values: `COMPLETE`.
        :param pulumi.Input[str] create_time: This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
        :param pulumi.Input[bool] disabled: Whether to disable the backup task. Valid values: `true`, `false`.
        :param pulumi.Input[str] file_system_id: The File System ID of Nas.
        :param pulumi.Input[str] nas_backup_plan_name: The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] paths: List of backup path. Up to 65536 characters. e.g.`["/home", "/var"]`. **Note** You should at least specify a backup path, empty array not allowed here.
        :param pulumi.Input[str] retention: Backup retention days, the minimum is 1.
        :param pulumi.Input[str] schedule: Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
        :param pulumi.Input[str] vault_id: The ID of Backup vault.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NasBackupPlanState.__new__(_NasBackupPlanState)

        __props__.__dict__["backup_type"] = backup_type
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["file_system_id"] = file_system_id
        __props__.__dict__["nas_backup_plan_name"] = nas_backup_plan_name
        __props__.__dict__["options"] = options
        __props__.__dict__["paths"] = paths
        __props__.__dict__["retention"] = retention
        __props__.__dict__["schedule"] = schedule
        __props__.__dict__["vault_id"] = vault_id
        return NasBackupPlan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> pulumi.Output[str]:
        """
        Backup type. Valid values: `COMPLETE`.
        """
        return pulumi.get(self, "backup_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        Whether to disable the backup task. Valid values: `true`, `false`.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Output[str]:
        """
        The File System ID of Nas.
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter(name="nasBackupPlanName")
    def nas_backup_plan_name(self) -> pulumi.Output[str]:
        """
        The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
        """
        return pulumi.get(self, "nas_backup_plan_name")

    @property
    @pulumi.getter
    def options(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "options")

    @property
    @pulumi.getter
    def paths(self) -> pulumi.Output[Sequence[str]]:
        """
        List of backup path. Up to 65536 characters. e.g.`["/home", "/var"]`. **Note** You should at least specify a backup path, empty array not allowed here.
        """
        return pulumi.get(self, "paths")

    @property
    @pulumi.getter
    def retention(self) -> pulumi.Output[str]:
        """
        Backup retention days, the minimum is 1.
        """
        return pulumi.get(self, "retention")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[str]:
        """
        Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> pulumi.Output[str]:
        """
        The ID of Backup vault.
        """
        return pulumi.get(self, "vault_id")

