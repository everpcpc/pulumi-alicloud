# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ScheduleArgs', 'Schedule']

@pulumi.input_type
class ScheduleArgs:
    def __init__(__self__, *,
                 cron_expression: pulumi.Input[str],
                 flow_name: pulumi.Input[str],
                 schedule_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 payload: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Schedule resource.
        :param pulumi.Input[str] cron_expression: The CRON expression of the time-based schedule to be created.
        :param pulumi.Input[str] flow_name: The name of the flow bound to the time-based schedule you want to create.
        :param pulumi.Input[str] schedule_name: The name of the time-based schedule to be created.
        :param pulumi.Input[str] description: The description of the time-based schedule to be created.
        :param pulumi.Input[bool] enable: Specifies whether to enable the time-based schedule you want to create. Valid values: `false`, `true`.
        :param pulumi.Input[str] payload: The trigger message of the time-based schedule to be created. It must be in JSON object format.
        """
        pulumi.set(__self__, "cron_expression", cron_expression)
        pulumi.set(__self__, "flow_name", flow_name)
        pulumi.set(__self__, "schedule_name", schedule_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if payload is not None:
            pulumi.set(__self__, "payload", payload)

    @property
    @pulumi.getter(name="cronExpression")
    def cron_expression(self) -> pulumi.Input[str]:
        """
        The CRON expression of the time-based schedule to be created.
        """
        return pulumi.get(self, "cron_expression")

    @cron_expression.setter
    def cron_expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "cron_expression", value)

    @property
    @pulumi.getter(name="flowName")
    def flow_name(self) -> pulumi.Input[str]:
        """
        The name of the flow bound to the time-based schedule you want to create.
        """
        return pulumi.get(self, "flow_name")

    @flow_name.setter
    def flow_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "flow_name", value)

    @property
    @pulumi.getter(name="scheduleName")
    def schedule_name(self) -> pulumi.Input[str]:
        """
        The name of the time-based schedule to be created.
        """
        return pulumi.get(self, "schedule_name")

    @schedule_name.setter
    def schedule_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "schedule_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the time-based schedule to be created.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable the time-based schedule you want to create. Valid values: `false`, `true`.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter
    def payload(self) -> Optional[pulumi.Input[str]]:
        """
        The trigger message of the time-based schedule to be created. It must be in JSON object format.
        """
        return pulumi.get(self, "payload")

    @payload.setter
    def payload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payload", value)


@pulumi.input_type
class _ScheduleState:
    def __init__(__self__, *,
                 cron_expression: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 flow_name: Optional[pulumi.Input[str]] = None,
                 last_modified_time: Optional[pulumi.Input[str]] = None,
                 payload: Optional[pulumi.Input[str]] = None,
                 schedule_id: Optional[pulumi.Input[str]] = None,
                 schedule_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Schedule resources.
        :param pulumi.Input[str] cron_expression: The CRON expression of the time-based schedule to be created.
        :param pulumi.Input[str] description: The description of the time-based schedule to be created.
        :param pulumi.Input[bool] enable: Specifies whether to enable the time-based schedule you want to create. Valid values: `false`, `true`.
        :param pulumi.Input[str] flow_name: The name of the flow bound to the time-based schedule you want to create.
        :param pulumi.Input[str] last_modified_time: The time when the time-based schedule was last updated.
        :param pulumi.Input[str] payload: The trigger message of the time-based schedule to be created. It must be in JSON object format.
        :param pulumi.Input[str] schedule_id: The ID of the time-based schedule.
        :param pulumi.Input[str] schedule_name: The name of the time-based schedule to be created.
        """
        if cron_expression is not None:
            pulumi.set(__self__, "cron_expression", cron_expression)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if flow_name is not None:
            pulumi.set(__self__, "flow_name", flow_name)
        if last_modified_time is not None:
            pulumi.set(__self__, "last_modified_time", last_modified_time)
        if payload is not None:
            pulumi.set(__self__, "payload", payload)
        if schedule_id is not None:
            pulumi.set(__self__, "schedule_id", schedule_id)
        if schedule_name is not None:
            pulumi.set(__self__, "schedule_name", schedule_name)

    @property
    @pulumi.getter(name="cronExpression")
    def cron_expression(self) -> Optional[pulumi.Input[str]]:
        """
        The CRON expression of the time-based schedule to be created.
        """
        return pulumi.get(self, "cron_expression")

    @cron_expression.setter
    def cron_expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cron_expression", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the time-based schedule to be created.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable the time-based schedule you want to create. Valid values: `false`, `true`.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="flowName")
    def flow_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the flow bound to the time-based schedule you want to create.
        """
        return pulumi.get(self, "flow_name")

    @flow_name.setter
    def flow_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flow_name", value)

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time when the time-based schedule was last updated.
        """
        return pulumi.get(self, "last_modified_time")

    @last_modified_time.setter
    def last_modified_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_modified_time", value)

    @property
    @pulumi.getter
    def payload(self) -> Optional[pulumi.Input[str]]:
        """
        The trigger message of the time-based schedule to be created. It must be in JSON object format.
        """
        return pulumi.get(self, "payload")

    @payload.setter
    def payload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payload", value)

    @property
    @pulumi.getter(name="scheduleId")
    def schedule_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the time-based schedule.
        """
        return pulumi.get(self, "schedule_id")

    @schedule_id.setter
    def schedule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_id", value)

    @property
    @pulumi.getter(name="scheduleName")
    def schedule_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the time-based schedule to be created.
        """
        return pulumi.get(self, "schedule_name")

    @schedule_name.setter
    def schedule_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_name", value)


class Schedule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cron_expression: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 flow_name: Optional[pulumi.Input[str]] = None,
                 payload: Optional[pulumi.Input[str]] = None,
                 schedule_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Serverless Workflow Schedule resource.

        For information about Serverless Workflow Schedule and how to use it, see [What is Schedule](https://www.alibabacloud.com/help/en/doc-detail/168934.htm).

        > **NOTE:** Available in v1.105.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_flow = alicloud.fnf.Flow("exampleFlow",
            definition=\"\"\"  version: v1beta1
          type: flow
          steps:
            - type: pass
              name: helloworld
        \"\"\",
            description="tf-testaccFnFFlow983041",
            type="FDL")
        example_schedule = alicloud.fnf.Schedule("exampleSchedule",
            cron_expression="30 9 * * * *",
            description="tf-testaccFnFSchedule983041",
            enable=True,
            flow_name=example_flow.name,
            payload="{\"tf-test\": \"test success\"}",
            schedule_name="tf-testaccFnFSchedule983041")
        ```

        ## Import

        Serverless Workflow Schedule can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:fnf/schedule:Schedule example <schedule_name>:<flow_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cron_expression: The CRON expression of the time-based schedule to be created.
        :param pulumi.Input[str] description: The description of the time-based schedule to be created.
        :param pulumi.Input[bool] enable: Specifies whether to enable the time-based schedule you want to create. Valid values: `false`, `true`.
        :param pulumi.Input[str] flow_name: The name of the flow bound to the time-based schedule you want to create.
        :param pulumi.Input[str] payload: The trigger message of the time-based schedule to be created. It must be in JSON object format.
        :param pulumi.Input[str] schedule_name: The name of the time-based schedule to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ScheduleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Serverless Workflow Schedule resource.

        For information about Serverless Workflow Schedule and how to use it, see [What is Schedule](https://www.alibabacloud.com/help/en/doc-detail/168934.htm).

        > **NOTE:** Available in v1.105.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_flow = alicloud.fnf.Flow("exampleFlow",
            definition=\"\"\"  version: v1beta1
          type: flow
          steps:
            - type: pass
              name: helloworld
        \"\"\",
            description="tf-testaccFnFFlow983041",
            type="FDL")
        example_schedule = alicloud.fnf.Schedule("exampleSchedule",
            cron_expression="30 9 * * * *",
            description="tf-testaccFnFSchedule983041",
            enable=True,
            flow_name=example_flow.name,
            payload="{\"tf-test\": \"test success\"}",
            schedule_name="tf-testaccFnFSchedule983041")
        ```

        ## Import

        Serverless Workflow Schedule can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:fnf/schedule:Schedule example <schedule_name>:<flow_name>
        ```

        :param str resource_name: The name of the resource.
        :param ScheduleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScheduleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cron_expression: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 flow_name: Optional[pulumi.Input[str]] = None,
                 payload: Optional[pulumi.Input[str]] = None,
                 schedule_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = ScheduleArgs.__new__(ScheduleArgs)

            if cron_expression is None and not opts.urn:
                raise TypeError("Missing required property 'cron_expression'")
            __props__.__dict__["cron_expression"] = cron_expression
            __props__.__dict__["description"] = description
            __props__.__dict__["enable"] = enable
            if flow_name is None and not opts.urn:
                raise TypeError("Missing required property 'flow_name'")
            __props__.__dict__["flow_name"] = flow_name
            __props__.__dict__["payload"] = payload
            if schedule_name is None and not opts.urn:
                raise TypeError("Missing required property 'schedule_name'")
            __props__.__dict__["schedule_name"] = schedule_name
            __props__.__dict__["last_modified_time"] = None
            __props__.__dict__["schedule_id"] = None
        super(Schedule, __self__).__init__(
            'alicloud:fnf/schedule:Schedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cron_expression: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enable: Optional[pulumi.Input[bool]] = None,
            flow_name: Optional[pulumi.Input[str]] = None,
            last_modified_time: Optional[pulumi.Input[str]] = None,
            payload: Optional[pulumi.Input[str]] = None,
            schedule_id: Optional[pulumi.Input[str]] = None,
            schedule_name: Optional[pulumi.Input[str]] = None) -> 'Schedule':
        """
        Get an existing Schedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cron_expression: The CRON expression of the time-based schedule to be created.
        :param pulumi.Input[str] description: The description of the time-based schedule to be created.
        :param pulumi.Input[bool] enable: Specifies whether to enable the time-based schedule you want to create. Valid values: `false`, `true`.
        :param pulumi.Input[str] flow_name: The name of the flow bound to the time-based schedule you want to create.
        :param pulumi.Input[str] last_modified_time: The time when the time-based schedule was last updated.
        :param pulumi.Input[str] payload: The trigger message of the time-based schedule to be created. It must be in JSON object format.
        :param pulumi.Input[str] schedule_id: The ID of the time-based schedule.
        :param pulumi.Input[str] schedule_name: The name of the time-based schedule to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ScheduleState.__new__(_ScheduleState)

        __props__.__dict__["cron_expression"] = cron_expression
        __props__.__dict__["description"] = description
        __props__.__dict__["enable"] = enable
        __props__.__dict__["flow_name"] = flow_name
        __props__.__dict__["last_modified_time"] = last_modified_time
        __props__.__dict__["payload"] = payload
        __props__.__dict__["schedule_id"] = schedule_id
        __props__.__dict__["schedule_name"] = schedule_name
        return Schedule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cronExpression")
    def cron_expression(self) -> pulumi.Output[str]:
        """
        The CRON expression of the time-based schedule to be created.
        """
        return pulumi.get(self, "cron_expression")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the time-based schedule to be created.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to enable the time-based schedule you want to create. Valid values: `false`, `true`.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter(name="flowName")
    def flow_name(self) -> pulumi.Output[str]:
        """
        The name of the flow bound to the time-based schedule you want to create.
        """
        return pulumi.get(self, "flow_name")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> pulumi.Output[str]:
        """
        The time when the time-based schedule was last updated.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def payload(self) -> pulumi.Output[Optional[str]]:
        """
        The trigger message of the time-based schedule to be created. It must be in JSON object format.
        """
        return pulumi.get(self, "payload")

    @property
    @pulumi.getter(name="scheduleId")
    def schedule_id(self) -> pulumi.Output[str]:
        """
        The ID of the time-based schedule.
        """
        return pulumi.get(self, "schedule_id")

    @property
    @pulumi.getter(name="scheduleName")
    def schedule_name(self) -> pulumi.Output[str]:
        """
        The name of the time-based schedule to be created.
        """
        return pulumi.get(self, "schedule_name")

