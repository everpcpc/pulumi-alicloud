# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EcsBackupPlanArgs', 'EcsBackupPlan']

@pulumi.input_type
class EcsBackupPlanArgs:
    def __init__(__self__, *,
                 ecs_backup_plan_name: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 retention: pulumi.Input[str],
                 schedule: pulumi.Input[str],
                 backup_type: Optional[pulumi.Input[str]] = None,
                 detail: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 exclude: Optional[pulumi.Input[str]] = None,
                 include: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[str]] = None,
                 paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 speed_limit: Optional[pulumi.Input[str]] = None,
                 update_paths: Optional[pulumi.Input[bool]] = None,
                 vault_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EcsBackupPlan resource.
        :param pulumi.Input[str] ecs_backup_plan_name: The Configuration Page of a Backup Plan Name. 1-64 Characters, requiring a Single Warehouse under Each of the Data Source Type Drop-down List of the Configuration Page of a Backup Plan Name Is Unique.
        :param pulumi.Input[str] instance_id: The ECS Instance Id. Must Have Installed the Client.
        :param pulumi.Input[str] retention: Backup Retention Period, the Minimum Value of 1.
        :param pulumi.Input[str] schedule: Backup strategy. Optional format: I|{startTime}|{interval} * startTime Backup start time, UNIX time, in seconds. * interval ISO8601 time interval. E.g: ** PT1H, one hour apart. ** P1D, one day apart. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed, the next backup task will not be triggered.
        :param pulumi.Input[str] backup_type: Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
        :param pulumi.Input[bool] disabled: Whether to Disable the Backup Task. Valid Values: true, false.
        :param pulumi.Input[str] exclude: Exclude Path. String of Json List, most 255 Characters. e.g. `"[\"/home/work\"]"`
        :param pulumi.Input[str] include: Include Path. String of Json List, most 255 Characters. e.g. `"[\"/var\"]"`
        :param pulumi.Input[str] options: Windows System with Application Consistency Using VSS. eg: {`UseVSS`:false}.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] paths: Backup Path. e.g. `["/home", "/var"]`
        :param pulumi.Input[str] speed_limit: flow control. The format is: {start}|{end}|{bandwidth} * start starting hour * end end hour * bandwidth limit rate, in KiB ** Use | to separate multiple flow control configurations; ** Multiple flow control configurations are not allowed to have overlapping times.
        :param pulumi.Input[str] vault_id: Vault ID.
        """
        pulumi.set(__self__, "ecs_backup_plan_name", ecs_backup_plan_name)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "retention", retention)
        pulumi.set(__self__, "schedule", schedule)
        if backup_type is not None:
            pulumi.set(__self__, "backup_type", backup_type)
        if detail is not None:
            pulumi.set(__self__, "detail", detail)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if exclude is not None:
            pulumi.set(__self__, "exclude", exclude)
        if include is not None:
            pulumi.set(__self__, "include", include)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if paths is not None:
            pulumi.set(__self__, "paths", paths)
        if speed_limit is not None:
            pulumi.set(__self__, "speed_limit", speed_limit)
        if update_paths is not None:
            pulumi.set(__self__, "update_paths", update_paths)
        if vault_id is not None:
            pulumi.set(__self__, "vault_id", vault_id)

    @property
    @pulumi.getter(name="ecsBackupPlanName")
    def ecs_backup_plan_name(self) -> pulumi.Input[str]:
        """
        The Configuration Page of a Backup Plan Name. 1-64 Characters, requiring a Single Warehouse under Each of the Data Source Type Drop-down List of the Configuration Page of a Backup Plan Name Is Unique.
        """
        return pulumi.get(self, "ecs_backup_plan_name")

    @ecs_backup_plan_name.setter
    def ecs_backup_plan_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "ecs_backup_plan_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ECS Instance Id. Must Have Installed the Client.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def retention(self) -> pulumi.Input[str]:
        """
        Backup Retention Period, the Minimum Value of 1.
        """
        return pulumi.get(self, "retention")

    @retention.setter
    def retention(self, value: pulumi.Input[str]):
        pulumi.set(self, "retention", value)

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Input[str]:
        """
        Backup strategy. Optional format: I|{startTime}|{interval} * startTime Backup start time, UNIX time, in seconds. * interval ISO8601 time interval. E.g: ** PT1H, one hour apart. ** P1D, one day apart. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed, the next backup task will not be triggered.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: pulumi.Input[str]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> Optional[pulumi.Input[str]]:
        """
        Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
        """
        return pulumi.get(self, "backup_type")

    @backup_type.setter
    def backup_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_type", value)

    @property
    @pulumi.getter
    def detail(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "detail")

    @detail.setter
    def detail(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "detail", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to Disable the Backup Task. Valid Values: true, false.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def exclude(self) -> Optional[pulumi.Input[str]]:
        """
        Exclude Path. String of Json List, most 255 Characters. e.g. `"[\"/home/work\"]"`
        """
        return pulumi.get(self, "exclude")

    @exclude.setter
    def exclude(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclude", value)

    @property
    @pulumi.getter
    def include(self) -> Optional[pulumi.Input[str]]:
        """
        Include Path. String of Json List, most 255 Characters. e.g. `"[\"/var\"]"`
        """
        return pulumi.get(self, "include")

    @include.setter
    def include(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "include", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[str]]:
        """
        Windows System with Application Consistency Using VSS. eg: {`UseVSS`:false}.
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter
    def paths(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Backup Path. e.g. `["/home", "/var"]`
        """
        return pulumi.get(self, "paths")

    @paths.setter
    def paths(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "paths", value)

    @property
    @pulumi.getter(name="speedLimit")
    def speed_limit(self) -> Optional[pulumi.Input[str]]:
        """
        flow control. The format is: {start}|{end}|{bandwidth} * start starting hour * end end hour * bandwidth limit rate, in KiB ** Use | to separate multiple flow control configurations; ** Multiple flow control configurations are not allowed to have overlapping times.
        """
        return pulumi.get(self, "speed_limit")

    @speed_limit.setter
    def speed_limit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "speed_limit", value)

    @property
    @pulumi.getter(name="updatePaths")
    def update_paths(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "update_paths")

    @update_paths.setter
    def update_paths(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "update_paths", value)

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> Optional[pulumi.Input[str]]:
        """
        Vault ID.
        """
        return pulumi.get(self, "vault_id")

    @vault_id.setter
    def vault_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vault_id", value)


@pulumi.input_type
class _EcsBackupPlanState:
    def __init__(__self__, *,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 detail: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 ecs_backup_plan_name: Optional[pulumi.Input[str]] = None,
                 exclude: Optional[pulumi.Input[str]] = None,
                 include: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[str]] = None,
                 paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 speed_limit: Optional[pulumi.Input[str]] = None,
                 update_paths: Optional[pulumi.Input[bool]] = None,
                 vault_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EcsBackupPlan resources.
        :param pulumi.Input[str] backup_type: Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
        :param pulumi.Input[bool] disabled: Whether to Disable the Backup Task. Valid Values: true, false.
        :param pulumi.Input[str] ecs_backup_plan_name: The Configuration Page of a Backup Plan Name. 1-64 Characters, requiring a Single Warehouse under Each of the Data Source Type Drop-down List of the Configuration Page of a Backup Plan Name Is Unique.
        :param pulumi.Input[str] exclude: Exclude Path. String of Json List, most 255 Characters. e.g. `"[\"/home/work\"]"`
        :param pulumi.Input[str] include: Include Path. String of Json List, most 255 Characters. e.g. `"[\"/var\"]"`
        :param pulumi.Input[str] instance_id: The ECS Instance Id. Must Have Installed the Client.
        :param pulumi.Input[str] options: Windows System with Application Consistency Using VSS. eg: {`UseVSS`:false}.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] paths: Backup Path. e.g. `["/home", "/var"]`
        :param pulumi.Input[str] retention: Backup Retention Period, the Minimum Value of 1.
        :param pulumi.Input[str] schedule: Backup strategy. Optional format: I|{startTime}|{interval} * startTime Backup start time, UNIX time, in seconds. * interval ISO8601 time interval. E.g: ** PT1H, one hour apart. ** P1D, one day apart. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed, the next backup task will not be triggered.
        :param pulumi.Input[str] speed_limit: flow control. The format is: {start}|{end}|{bandwidth} * start starting hour * end end hour * bandwidth limit rate, in KiB ** Use | to separate multiple flow control configurations; ** Multiple flow control configurations are not allowed to have overlapping times.
        :param pulumi.Input[str] vault_id: Vault ID.
        """
        if backup_type is not None:
            pulumi.set(__self__, "backup_type", backup_type)
        if detail is not None:
            pulumi.set(__self__, "detail", detail)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if ecs_backup_plan_name is not None:
            pulumi.set(__self__, "ecs_backup_plan_name", ecs_backup_plan_name)
        if exclude is not None:
            pulumi.set(__self__, "exclude", exclude)
        if include is not None:
            pulumi.set(__self__, "include", include)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if paths is not None:
            pulumi.set(__self__, "paths", paths)
        if retention is not None:
            pulumi.set(__self__, "retention", retention)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if speed_limit is not None:
            pulumi.set(__self__, "speed_limit", speed_limit)
        if update_paths is not None:
            pulumi.set(__self__, "update_paths", update_paths)
        if vault_id is not None:
            pulumi.set(__self__, "vault_id", vault_id)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> Optional[pulumi.Input[str]]:
        """
        Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
        """
        return pulumi.get(self, "backup_type")

    @backup_type.setter
    def backup_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_type", value)

    @property
    @pulumi.getter
    def detail(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "detail")

    @detail.setter
    def detail(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "detail", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to Disable the Backup Task. Valid Values: true, false.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="ecsBackupPlanName")
    def ecs_backup_plan_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Configuration Page of a Backup Plan Name. 1-64 Characters, requiring a Single Warehouse under Each of the Data Source Type Drop-down List of the Configuration Page of a Backup Plan Name Is Unique.
        """
        return pulumi.get(self, "ecs_backup_plan_name")

    @ecs_backup_plan_name.setter
    def ecs_backup_plan_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ecs_backup_plan_name", value)

    @property
    @pulumi.getter
    def exclude(self) -> Optional[pulumi.Input[str]]:
        """
        Exclude Path. String of Json List, most 255 Characters. e.g. `"[\"/home/work\"]"`
        """
        return pulumi.get(self, "exclude")

    @exclude.setter
    def exclude(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclude", value)

    @property
    @pulumi.getter
    def include(self) -> Optional[pulumi.Input[str]]:
        """
        Include Path. String of Json List, most 255 Characters. e.g. `"[\"/var\"]"`
        """
        return pulumi.get(self, "include")

    @include.setter
    def include(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "include", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ECS Instance Id. Must Have Installed the Client.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[str]]:
        """
        Windows System with Application Consistency Using VSS. eg: {`UseVSS`:false}.
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter
    def paths(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Backup Path. e.g. `["/home", "/var"]`
        """
        return pulumi.get(self, "paths")

    @paths.setter
    def paths(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "paths", value)

    @property
    @pulumi.getter
    def retention(self) -> Optional[pulumi.Input[str]]:
        """
        Backup Retention Period, the Minimum Value of 1.
        """
        return pulumi.get(self, "retention")

    @retention.setter
    def retention(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "retention", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[str]]:
        """
        Backup strategy. Optional format: I|{startTime}|{interval} * startTime Backup start time, UNIX time, in seconds. * interval ISO8601 time interval. E.g: ** PT1H, one hour apart. ** P1D, one day apart. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed, the next backup task will not be triggered.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="speedLimit")
    def speed_limit(self) -> Optional[pulumi.Input[str]]:
        """
        flow control. The format is: {start}|{end}|{bandwidth} * start starting hour * end end hour * bandwidth limit rate, in KiB ** Use | to separate multiple flow control configurations; ** Multiple flow control configurations are not allowed to have overlapping times.
        """
        return pulumi.get(self, "speed_limit")

    @speed_limit.setter
    def speed_limit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "speed_limit", value)

    @property
    @pulumi.getter(name="updatePaths")
    def update_paths(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "update_paths")

    @update_paths.setter
    def update_paths(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "update_paths", value)

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> Optional[pulumi.Input[str]]:
        """
        Vault ID.
        """
        return pulumi.get(self, "vault_id")

    @vault_id.setter
    def vault_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vault_id", value)


class EcsBackupPlan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 detail: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 ecs_backup_plan_name: Optional[pulumi.Input[str]] = None,
                 exclude: Optional[pulumi.Input[str]] = None,
                 include: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[str]] = None,
                 paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 speed_limit: Optional[pulumi.Input[str]] = None,
                 update_paths: Optional[pulumi.Input[bool]] = None,
                 vault_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a HBR Ecs Backup Plan resource.

        For information about HBR Ecs Backup Plan and how to use it, see [What is Ecs Backup Plan](https://www.alibabacloud.com/help/doc-detail/186568.htm).

        > **NOTE:** Available in v1.132.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.hbr.EcsBackupPlan("example",
            backup_type="COMPLETE",
            ecs_backup_plan_name="example_value",
            exclude=\"\"\"  ["/home/exclude"]
          
        \"\"\",
            include=\"\"\"  ["/home/include"]
          
        \"\"\",
            instance_id="i-bp1567rc0oxxxxxxxxxx",
            paths=[
                "/home",
                "/var",
            ],
            retention="1",
            schedule="I|1602673264|PT2H",
            speed_limit="I|1602673264|PT2H",
            vault_id="v-0003gxoksflhxxxxxxxx")
        ```

        ## Import

        HBR Ecs Backup Plan can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:hbr/ecsBackupPlan:EcsBackupPlan example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_type: Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
        :param pulumi.Input[bool] disabled: Whether to Disable the Backup Task. Valid Values: true, false.
        :param pulumi.Input[str] ecs_backup_plan_name: The Configuration Page of a Backup Plan Name. 1-64 Characters, requiring a Single Warehouse under Each of the Data Source Type Drop-down List of the Configuration Page of a Backup Plan Name Is Unique.
        :param pulumi.Input[str] exclude: Exclude Path. String of Json List, most 255 Characters. e.g. `"[\"/home/work\"]"`
        :param pulumi.Input[str] include: Include Path. String of Json List, most 255 Characters. e.g. `"[\"/var\"]"`
        :param pulumi.Input[str] instance_id: The ECS Instance Id. Must Have Installed the Client.
        :param pulumi.Input[str] options: Windows System with Application Consistency Using VSS. eg: {`UseVSS`:false}.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] paths: Backup Path. e.g. `["/home", "/var"]`
        :param pulumi.Input[str] retention: Backup Retention Period, the Minimum Value of 1.
        :param pulumi.Input[str] schedule: Backup strategy. Optional format: I|{startTime}|{interval} * startTime Backup start time, UNIX time, in seconds. * interval ISO8601 time interval. E.g: ** PT1H, one hour apart. ** P1D, one day apart. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed, the next backup task will not be triggered.
        :param pulumi.Input[str] speed_limit: flow control. The format is: {start}|{end}|{bandwidth} * start starting hour * end end hour * bandwidth limit rate, in KiB ** Use | to separate multiple flow control configurations; ** Multiple flow control configurations are not allowed to have overlapping times.
        :param pulumi.Input[str] vault_id: Vault ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EcsBackupPlanArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a HBR Ecs Backup Plan resource.

        For information about HBR Ecs Backup Plan and how to use it, see [What is Ecs Backup Plan](https://www.alibabacloud.com/help/doc-detail/186568.htm).

        > **NOTE:** Available in v1.132.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.hbr.EcsBackupPlan("example",
            backup_type="COMPLETE",
            ecs_backup_plan_name="example_value",
            exclude=\"\"\"  ["/home/exclude"]
          
        \"\"\",
            include=\"\"\"  ["/home/include"]
          
        \"\"\",
            instance_id="i-bp1567rc0oxxxxxxxxxx",
            paths=[
                "/home",
                "/var",
            ],
            retention="1",
            schedule="I|1602673264|PT2H",
            speed_limit="I|1602673264|PT2H",
            vault_id="v-0003gxoksflhxxxxxxxx")
        ```

        ## Import

        HBR Ecs Backup Plan can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:hbr/ecsBackupPlan:EcsBackupPlan example <id>
        ```

        :param str resource_name: The name of the resource.
        :param EcsBackupPlanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EcsBackupPlanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 detail: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 ecs_backup_plan_name: Optional[pulumi.Input[str]] = None,
                 exclude: Optional[pulumi.Input[str]] = None,
                 include: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[str]] = None,
                 paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 speed_limit: Optional[pulumi.Input[str]] = None,
                 update_paths: Optional[pulumi.Input[bool]] = None,
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
            __props__ = EcsBackupPlanArgs.__new__(EcsBackupPlanArgs)

            __props__.__dict__["backup_type"] = backup_type
            __props__.__dict__["detail"] = detail
            __props__.__dict__["disabled"] = disabled
            if ecs_backup_plan_name is None and not opts.urn:
                raise TypeError("Missing required property 'ecs_backup_plan_name'")
            __props__.__dict__["ecs_backup_plan_name"] = ecs_backup_plan_name
            __props__.__dict__["exclude"] = exclude
            __props__.__dict__["include"] = include
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["options"] = options
            __props__.__dict__["paths"] = paths
            if retention is None and not opts.urn:
                raise TypeError("Missing required property 'retention'")
            __props__.__dict__["retention"] = retention
            if schedule is None and not opts.urn:
                raise TypeError("Missing required property 'schedule'")
            __props__.__dict__["schedule"] = schedule
            __props__.__dict__["speed_limit"] = speed_limit
            __props__.__dict__["update_paths"] = update_paths
            __props__.__dict__["vault_id"] = vault_id
        super(EcsBackupPlan, __self__).__init__(
            'alicloud:hbr/ecsBackupPlan:EcsBackupPlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_type: Optional[pulumi.Input[str]] = None,
            detail: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            ecs_backup_plan_name: Optional[pulumi.Input[str]] = None,
            exclude: Optional[pulumi.Input[str]] = None,
            include: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            options: Optional[pulumi.Input[str]] = None,
            paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            retention: Optional[pulumi.Input[str]] = None,
            schedule: Optional[pulumi.Input[str]] = None,
            speed_limit: Optional[pulumi.Input[str]] = None,
            update_paths: Optional[pulumi.Input[bool]] = None,
            vault_id: Optional[pulumi.Input[str]] = None) -> 'EcsBackupPlan':
        """
        Get an existing EcsBackupPlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_type: Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
        :param pulumi.Input[bool] disabled: Whether to Disable the Backup Task. Valid Values: true, false.
        :param pulumi.Input[str] ecs_backup_plan_name: The Configuration Page of a Backup Plan Name. 1-64 Characters, requiring a Single Warehouse under Each of the Data Source Type Drop-down List of the Configuration Page of a Backup Plan Name Is Unique.
        :param pulumi.Input[str] exclude: Exclude Path. String of Json List, most 255 Characters. e.g. `"[\"/home/work\"]"`
        :param pulumi.Input[str] include: Include Path. String of Json List, most 255 Characters. e.g. `"[\"/var\"]"`
        :param pulumi.Input[str] instance_id: The ECS Instance Id. Must Have Installed the Client.
        :param pulumi.Input[str] options: Windows System with Application Consistency Using VSS. eg: {`UseVSS`:false}.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] paths: Backup Path. e.g. `["/home", "/var"]`
        :param pulumi.Input[str] retention: Backup Retention Period, the Minimum Value of 1.
        :param pulumi.Input[str] schedule: Backup strategy. Optional format: I|{startTime}|{interval} * startTime Backup start time, UNIX time, in seconds. * interval ISO8601 time interval. E.g: ** PT1H, one hour apart. ** P1D, one day apart. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed, the next backup task will not be triggered.
        :param pulumi.Input[str] speed_limit: flow control. The format is: {start}|{end}|{bandwidth} * start starting hour * end end hour * bandwidth limit rate, in KiB ** Use | to separate multiple flow control configurations; ** Multiple flow control configurations are not allowed to have overlapping times.
        :param pulumi.Input[str] vault_id: Vault ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EcsBackupPlanState.__new__(_EcsBackupPlanState)

        __props__.__dict__["backup_type"] = backup_type
        __props__.__dict__["detail"] = detail
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["ecs_backup_plan_name"] = ecs_backup_plan_name
        __props__.__dict__["exclude"] = exclude
        __props__.__dict__["include"] = include
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["options"] = options
        __props__.__dict__["paths"] = paths
        __props__.__dict__["retention"] = retention
        __props__.__dict__["schedule"] = schedule
        __props__.__dict__["speed_limit"] = speed_limit
        __props__.__dict__["update_paths"] = update_paths
        __props__.__dict__["vault_id"] = vault_id
        return EcsBackupPlan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> pulumi.Output[str]:
        """
        Backup Type. Valid Values: * Complete. Valid values: `COMPLETE`.
        """
        return pulumi.get(self, "backup_type")

    @property
    @pulumi.getter
    def detail(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "detail")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        Whether to Disable the Backup Task. Valid Values: true, false.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="ecsBackupPlanName")
    def ecs_backup_plan_name(self) -> pulumi.Output[str]:
        """
        The Configuration Page of a Backup Plan Name. 1-64 Characters, requiring a Single Warehouse under Each of the Data Source Type Drop-down List of the Configuration Page of a Backup Plan Name Is Unique.
        """
        return pulumi.get(self, "ecs_backup_plan_name")

    @property
    @pulumi.getter
    def exclude(self) -> pulumi.Output[Optional[str]]:
        """
        Exclude Path. String of Json List, most 255 Characters. e.g. `"[\"/home/work\"]"`
        """
        return pulumi.get(self, "exclude")

    @property
    @pulumi.getter
    def include(self) -> pulumi.Output[Optional[str]]:
        """
        Include Path. String of Json List, most 255 Characters. e.g. `"[\"/var\"]"`
        """
        return pulumi.get(self, "include")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ECS Instance Id. Must Have Installed the Client.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def options(self) -> pulumi.Output[Optional[str]]:
        """
        Windows System with Application Consistency Using VSS. eg: {`UseVSS`:false}.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter
    def paths(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Backup Path. e.g. `["/home", "/var"]`
        """
        return pulumi.get(self, "paths")

    @property
    @pulumi.getter
    def retention(self) -> pulumi.Output[str]:
        """
        Backup Retention Period, the Minimum Value of 1.
        """
        return pulumi.get(self, "retention")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[str]:
        """
        Backup strategy. Optional format: I|{startTime}|{interval} * startTime Backup start time, UNIX time, in seconds. * interval ISO8601 time interval. E.g: ** PT1H, one hour apart. ** P1D, one day apart. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed, the next backup task will not be triggered.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="speedLimit")
    def speed_limit(self) -> pulumi.Output[Optional[str]]:
        """
        flow control. The format is: {start}|{end}|{bandwidth} * start starting hour * end end hour * bandwidth limit rate, in KiB ** Use | to separate multiple flow control configurations; ** Multiple flow control configurations are not allowed to have overlapping times.
        """
        return pulumi.get(self, "speed_limit")

    @property
    @pulumi.getter(name="updatePaths")
    def update_paths(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "update_paths")

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> pulumi.Output[Optional[str]]:
        """
        Vault ID.
        """
        return pulumi.get(self, "vault_id")

