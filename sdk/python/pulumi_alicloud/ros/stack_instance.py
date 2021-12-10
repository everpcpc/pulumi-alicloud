# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['StackInstanceArgs', 'StackInstance']

@pulumi.input_type
class StackInstanceArgs:
    def __init__(__self__, *,
                 stack_group_name: pulumi.Input[str],
                 stack_instance_account_id: pulumi.Input[str],
                 stack_instance_region_id: pulumi.Input[str],
                 operation_description: Optional[pulumi.Input[str]] = None,
                 operation_preferences: Optional[pulumi.Input[str]] = None,
                 parameter_overrides: Optional[pulumi.Input[Sequence[pulumi.Input['StackInstanceParameterOverrideArgs']]]] = None,
                 retain_stacks: Optional[pulumi.Input[bool]] = None,
                 timeout_in_minutes: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StackInstance resource.
        :param pulumi.Input[str] stack_group_name: The name of the stack group.
        :param pulumi.Input[str] stack_instance_account_id: The account to which the stack instance belongs.
        :param pulumi.Input[str] stack_instance_region_id: The region of the stack instance.
        :param pulumi.Input[str] operation_description: The operation description.
        :param pulumi.Input[str] operation_preferences: The operation preferences. The operation settings. The following fields are supported:
        :param pulumi.Input[Sequence[pulumi.Input['StackInstanceParameterOverrideArgs']]] parameter_overrides: ParameterOverrides. See the following `Block parameter_overrides`.
        :param pulumi.Input[bool] retain_stacks: Specifies whether to retain the stack corresponding to the stack instance.Default value `false`. **NOTE:** When `retain_stacks` is `true`, the stack is retained. If the stack is retained, the corresponding stack is not deleted when the stack instance is deleted from the stack group.
        :param pulumi.Input[str] timeout_in_minutes: The timeout period that is specified for the stack creation request. Default value: `60`. Unit: `minutes`.
        """
        pulumi.set(__self__, "stack_group_name", stack_group_name)
        pulumi.set(__self__, "stack_instance_account_id", stack_instance_account_id)
        pulumi.set(__self__, "stack_instance_region_id", stack_instance_region_id)
        if operation_description is not None:
            pulumi.set(__self__, "operation_description", operation_description)
        if operation_preferences is not None:
            pulumi.set(__self__, "operation_preferences", operation_preferences)
        if parameter_overrides is not None:
            pulumi.set(__self__, "parameter_overrides", parameter_overrides)
        if retain_stacks is not None:
            pulumi.set(__self__, "retain_stacks", retain_stacks)
        if timeout_in_minutes is not None:
            pulumi.set(__self__, "timeout_in_minutes", timeout_in_minutes)

    @property
    @pulumi.getter(name="stackGroupName")
    def stack_group_name(self) -> pulumi.Input[str]:
        """
        The name of the stack group.
        """
        return pulumi.get(self, "stack_group_name")

    @stack_group_name.setter
    def stack_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "stack_group_name", value)

    @property
    @pulumi.getter(name="stackInstanceAccountId")
    def stack_instance_account_id(self) -> pulumi.Input[str]:
        """
        The account to which the stack instance belongs.
        """
        return pulumi.get(self, "stack_instance_account_id")

    @stack_instance_account_id.setter
    def stack_instance_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "stack_instance_account_id", value)

    @property
    @pulumi.getter(name="stackInstanceRegionId")
    def stack_instance_region_id(self) -> pulumi.Input[str]:
        """
        The region of the stack instance.
        """
        return pulumi.get(self, "stack_instance_region_id")

    @stack_instance_region_id.setter
    def stack_instance_region_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "stack_instance_region_id", value)

    @property
    @pulumi.getter(name="operationDescription")
    def operation_description(self) -> Optional[pulumi.Input[str]]:
        """
        The operation description.
        """
        return pulumi.get(self, "operation_description")

    @operation_description.setter
    def operation_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operation_description", value)

    @property
    @pulumi.getter(name="operationPreferences")
    def operation_preferences(self) -> Optional[pulumi.Input[str]]:
        """
        The operation preferences. The operation settings. The following fields are supported:
        """
        return pulumi.get(self, "operation_preferences")

    @operation_preferences.setter
    def operation_preferences(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operation_preferences", value)

    @property
    @pulumi.getter(name="parameterOverrides")
    def parameter_overrides(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StackInstanceParameterOverrideArgs']]]]:
        """
        ParameterOverrides. See the following `Block parameter_overrides`.
        """
        return pulumi.get(self, "parameter_overrides")

    @parameter_overrides.setter
    def parameter_overrides(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StackInstanceParameterOverrideArgs']]]]):
        pulumi.set(self, "parameter_overrides", value)

    @property
    @pulumi.getter(name="retainStacks")
    def retain_stacks(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to retain the stack corresponding to the stack instance.Default value `false`. **NOTE:** When `retain_stacks` is `true`, the stack is retained. If the stack is retained, the corresponding stack is not deleted when the stack instance is deleted from the stack group.
        """
        return pulumi.get(self, "retain_stacks")

    @retain_stacks.setter
    def retain_stacks(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "retain_stacks", value)

    @property
    @pulumi.getter(name="timeoutInMinutes")
    def timeout_in_minutes(self) -> Optional[pulumi.Input[str]]:
        """
        The timeout period that is specified for the stack creation request. Default value: `60`. Unit: `minutes`.
        """
        return pulumi.get(self, "timeout_in_minutes")

    @timeout_in_minutes.setter
    def timeout_in_minutes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timeout_in_minutes", value)


@pulumi.input_type
class _StackInstanceState:
    def __init__(__self__, *,
                 operation_description: Optional[pulumi.Input[str]] = None,
                 operation_preferences: Optional[pulumi.Input[str]] = None,
                 parameter_overrides: Optional[pulumi.Input[Sequence[pulumi.Input['StackInstanceParameterOverrideArgs']]]] = None,
                 retain_stacks: Optional[pulumi.Input[bool]] = None,
                 stack_group_name: Optional[pulumi.Input[str]] = None,
                 stack_instance_account_id: Optional[pulumi.Input[str]] = None,
                 stack_instance_region_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 timeout_in_minutes: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StackInstance resources.
        :param pulumi.Input[str] operation_description: The operation description.
        :param pulumi.Input[str] operation_preferences: The operation preferences. The operation settings. The following fields are supported:
        :param pulumi.Input[Sequence[pulumi.Input['StackInstanceParameterOverrideArgs']]] parameter_overrides: ParameterOverrides. See the following `Block parameter_overrides`.
        :param pulumi.Input[bool] retain_stacks: Specifies whether to retain the stack corresponding to the stack instance.Default value `false`. **NOTE:** When `retain_stacks` is `true`, the stack is retained. If the stack is retained, the corresponding stack is not deleted when the stack instance is deleted from the stack group.
        :param pulumi.Input[str] stack_group_name: The name of the stack group.
        :param pulumi.Input[str] stack_instance_account_id: The account to which the stack instance belongs.
        :param pulumi.Input[str] stack_instance_region_id: The region of the stack instance.
        :param pulumi.Input[str] status: The status of the stack instance. Valid values: `CURRENT` or `OUTDATED`. 
               * `CURRENT`: The stack corresponding to the stack instance is up to date with the stack group.
               * `OUTDATED`: The stack corresponding to the stack instance is not up to date with the stack group. The `OUTDATED` state has the following possible causes:
               * When the CreateStackInstances operation is called to create stack instances, the corresponding stacks fail to be created.
               * When the UpdateStackInstances or UpdateStackGroup operation is called to update stack instances, the corresponding stacks fail to be updated, or only some of the stack instances are updated.
               * The create or update operation is not complete.
        :param pulumi.Input[str] timeout_in_minutes: The timeout period that is specified for the stack creation request. Default value: `60`. Unit: `minutes`.
        """
        if operation_description is not None:
            pulumi.set(__self__, "operation_description", operation_description)
        if operation_preferences is not None:
            pulumi.set(__self__, "operation_preferences", operation_preferences)
        if parameter_overrides is not None:
            pulumi.set(__self__, "parameter_overrides", parameter_overrides)
        if retain_stacks is not None:
            pulumi.set(__self__, "retain_stacks", retain_stacks)
        if stack_group_name is not None:
            pulumi.set(__self__, "stack_group_name", stack_group_name)
        if stack_instance_account_id is not None:
            pulumi.set(__self__, "stack_instance_account_id", stack_instance_account_id)
        if stack_instance_region_id is not None:
            pulumi.set(__self__, "stack_instance_region_id", stack_instance_region_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if timeout_in_minutes is not None:
            pulumi.set(__self__, "timeout_in_minutes", timeout_in_minutes)

    @property
    @pulumi.getter(name="operationDescription")
    def operation_description(self) -> Optional[pulumi.Input[str]]:
        """
        The operation description.
        """
        return pulumi.get(self, "operation_description")

    @operation_description.setter
    def operation_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operation_description", value)

    @property
    @pulumi.getter(name="operationPreferences")
    def operation_preferences(self) -> Optional[pulumi.Input[str]]:
        """
        The operation preferences. The operation settings. The following fields are supported:
        """
        return pulumi.get(self, "operation_preferences")

    @operation_preferences.setter
    def operation_preferences(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operation_preferences", value)

    @property
    @pulumi.getter(name="parameterOverrides")
    def parameter_overrides(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StackInstanceParameterOverrideArgs']]]]:
        """
        ParameterOverrides. See the following `Block parameter_overrides`.
        """
        return pulumi.get(self, "parameter_overrides")

    @parameter_overrides.setter
    def parameter_overrides(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StackInstanceParameterOverrideArgs']]]]):
        pulumi.set(self, "parameter_overrides", value)

    @property
    @pulumi.getter(name="retainStacks")
    def retain_stacks(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to retain the stack corresponding to the stack instance.Default value `false`. **NOTE:** When `retain_stacks` is `true`, the stack is retained. If the stack is retained, the corresponding stack is not deleted when the stack instance is deleted from the stack group.
        """
        return pulumi.get(self, "retain_stacks")

    @retain_stacks.setter
    def retain_stacks(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "retain_stacks", value)

    @property
    @pulumi.getter(name="stackGroupName")
    def stack_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the stack group.
        """
        return pulumi.get(self, "stack_group_name")

    @stack_group_name.setter
    def stack_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_group_name", value)

    @property
    @pulumi.getter(name="stackInstanceAccountId")
    def stack_instance_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The account to which the stack instance belongs.
        """
        return pulumi.get(self, "stack_instance_account_id")

    @stack_instance_account_id.setter
    def stack_instance_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_instance_account_id", value)

    @property
    @pulumi.getter(name="stackInstanceRegionId")
    def stack_instance_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The region of the stack instance.
        """
        return pulumi.get(self, "stack_instance_region_id")

    @stack_instance_region_id.setter
    def stack_instance_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stack_instance_region_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the stack instance. Valid values: `CURRENT` or `OUTDATED`. 
        * `CURRENT`: The stack corresponding to the stack instance is up to date with the stack group.
        * `OUTDATED`: The stack corresponding to the stack instance is not up to date with the stack group. The `OUTDATED` state has the following possible causes:
        * When the CreateStackInstances operation is called to create stack instances, the corresponding stacks fail to be created.
        * When the UpdateStackInstances or UpdateStackGroup operation is called to update stack instances, the corresponding stacks fail to be updated, or only some of the stack instances are updated.
        * The create or update operation is not complete.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="timeoutInMinutes")
    def timeout_in_minutes(self) -> Optional[pulumi.Input[str]]:
        """
        The timeout period that is specified for the stack creation request. Default value: `60`. Unit: `minutes`.
        """
        return pulumi.get(self, "timeout_in_minutes")

    @timeout_in_minutes.setter
    def timeout_in_minutes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timeout_in_minutes", value)


class StackInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 operation_description: Optional[pulumi.Input[str]] = None,
                 operation_preferences: Optional[pulumi.Input[str]] = None,
                 parameter_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StackInstanceParameterOverrideArgs']]]]] = None,
                 retain_stacks: Optional[pulumi.Input[bool]] = None,
                 stack_group_name: Optional[pulumi.Input[str]] = None,
                 stack_instance_account_id: Optional[pulumi.Input[str]] = None,
                 stack_instance_region_id: Optional[pulumi.Input[str]] = None,
                 timeout_in_minutes: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a ROS Stack Instance resource.

        For information about ROS Stack Instance and how to use it, see [What is Stack Instance](https://www.alibabacloud.com/help/en/doc-detail/151338.html).

        > **NOTE:** Available in v1.145.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_regions = alicloud.ros.get_regions()
        example_stack_group = alicloud.ros.StackGroup("exampleStackGroup",
            stack_group_name=var["name"],
            template_body="{\"ROSTemplateFormatVersion\":\"2015-09-01\", \"Parameters\": {\"VpcName\": {\"Type\": \"String\"},\"InstanceType\": {\"Type\": \"String\"}}}",
            description="test for stack groups",
            parameters=[
                alicloud.ros.StackGroupParameterArgs(
                    parameter_key="VpcName",
                    parameter_value="VpcName",
                ),
                alicloud.ros.StackGroupParameterArgs(
                    parameter_key="InstanceType",
                    parameter_value="InstanceType",
                ),
            ])
        example_stack_instance = alicloud.ros.StackInstance("exampleStackInstance",
            stack_group_name=example_stack_group.stack_group_name,
            stack_instance_account_id="example_value",
            stack_instance_region_id=example_regions.regions[0].region_id,
            operation_preferences="{\"FailureToleranceCount\": 1, \"MaxConcurrentCount\": 2}",
            parameter_overrides=[alicloud.ros.StackInstanceParameterOverrideArgs(
                parameter_value="VpcName",
                parameter_key="VpcName",
            )])
        ```

        ## Import

        ROS Stack Instance can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ros/stackInstance:StackInstance example <stack_group_name>:<stack_instance_account_id>:<stack_instance_region_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] operation_description: The operation description.
        :param pulumi.Input[str] operation_preferences: The operation preferences. The operation settings. The following fields are supported:
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StackInstanceParameterOverrideArgs']]]] parameter_overrides: ParameterOverrides. See the following `Block parameter_overrides`.
        :param pulumi.Input[bool] retain_stacks: Specifies whether to retain the stack corresponding to the stack instance.Default value `false`. **NOTE:** When `retain_stacks` is `true`, the stack is retained. If the stack is retained, the corresponding stack is not deleted when the stack instance is deleted from the stack group.
        :param pulumi.Input[str] stack_group_name: The name of the stack group.
        :param pulumi.Input[str] stack_instance_account_id: The account to which the stack instance belongs.
        :param pulumi.Input[str] stack_instance_region_id: The region of the stack instance.
        :param pulumi.Input[str] timeout_in_minutes: The timeout period that is specified for the stack creation request. Default value: `60`. Unit: `minutes`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StackInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a ROS Stack Instance resource.

        For information about ROS Stack Instance and how to use it, see [What is Stack Instance](https://www.alibabacloud.com/help/en/doc-detail/151338.html).

        > **NOTE:** Available in v1.145.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_regions = alicloud.ros.get_regions()
        example_stack_group = alicloud.ros.StackGroup("exampleStackGroup",
            stack_group_name=var["name"],
            template_body="{\"ROSTemplateFormatVersion\":\"2015-09-01\", \"Parameters\": {\"VpcName\": {\"Type\": \"String\"},\"InstanceType\": {\"Type\": \"String\"}}}",
            description="test for stack groups",
            parameters=[
                alicloud.ros.StackGroupParameterArgs(
                    parameter_key="VpcName",
                    parameter_value="VpcName",
                ),
                alicloud.ros.StackGroupParameterArgs(
                    parameter_key="InstanceType",
                    parameter_value="InstanceType",
                ),
            ])
        example_stack_instance = alicloud.ros.StackInstance("exampleStackInstance",
            stack_group_name=example_stack_group.stack_group_name,
            stack_instance_account_id="example_value",
            stack_instance_region_id=example_regions.regions[0].region_id,
            operation_preferences="{\"FailureToleranceCount\": 1, \"MaxConcurrentCount\": 2}",
            parameter_overrides=[alicloud.ros.StackInstanceParameterOverrideArgs(
                parameter_value="VpcName",
                parameter_key="VpcName",
            )])
        ```

        ## Import

        ROS Stack Instance can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ros/stackInstance:StackInstance example <stack_group_name>:<stack_instance_account_id>:<stack_instance_region_id>
        ```

        :param str resource_name: The name of the resource.
        :param StackInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StackInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 operation_description: Optional[pulumi.Input[str]] = None,
                 operation_preferences: Optional[pulumi.Input[str]] = None,
                 parameter_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StackInstanceParameterOverrideArgs']]]]] = None,
                 retain_stacks: Optional[pulumi.Input[bool]] = None,
                 stack_group_name: Optional[pulumi.Input[str]] = None,
                 stack_instance_account_id: Optional[pulumi.Input[str]] = None,
                 stack_instance_region_id: Optional[pulumi.Input[str]] = None,
                 timeout_in_minutes: Optional[pulumi.Input[str]] = None,
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
            __props__ = StackInstanceArgs.__new__(StackInstanceArgs)

            __props__.__dict__["operation_description"] = operation_description
            __props__.__dict__["operation_preferences"] = operation_preferences
            __props__.__dict__["parameter_overrides"] = parameter_overrides
            __props__.__dict__["retain_stacks"] = retain_stacks
            if stack_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'stack_group_name'")
            __props__.__dict__["stack_group_name"] = stack_group_name
            if stack_instance_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'stack_instance_account_id'")
            __props__.__dict__["stack_instance_account_id"] = stack_instance_account_id
            if stack_instance_region_id is None and not opts.urn:
                raise TypeError("Missing required property 'stack_instance_region_id'")
            __props__.__dict__["stack_instance_region_id"] = stack_instance_region_id
            __props__.__dict__["timeout_in_minutes"] = timeout_in_minutes
            __props__.__dict__["status"] = None
        super(StackInstance, __self__).__init__(
            'alicloud:ros/stackInstance:StackInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            operation_description: Optional[pulumi.Input[str]] = None,
            operation_preferences: Optional[pulumi.Input[str]] = None,
            parameter_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StackInstanceParameterOverrideArgs']]]]] = None,
            retain_stacks: Optional[pulumi.Input[bool]] = None,
            stack_group_name: Optional[pulumi.Input[str]] = None,
            stack_instance_account_id: Optional[pulumi.Input[str]] = None,
            stack_instance_region_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            timeout_in_minutes: Optional[pulumi.Input[str]] = None) -> 'StackInstance':
        """
        Get an existing StackInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] operation_description: The operation description.
        :param pulumi.Input[str] operation_preferences: The operation preferences. The operation settings. The following fields are supported:
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StackInstanceParameterOverrideArgs']]]] parameter_overrides: ParameterOverrides. See the following `Block parameter_overrides`.
        :param pulumi.Input[bool] retain_stacks: Specifies whether to retain the stack corresponding to the stack instance.Default value `false`. **NOTE:** When `retain_stacks` is `true`, the stack is retained. If the stack is retained, the corresponding stack is not deleted when the stack instance is deleted from the stack group.
        :param pulumi.Input[str] stack_group_name: The name of the stack group.
        :param pulumi.Input[str] stack_instance_account_id: The account to which the stack instance belongs.
        :param pulumi.Input[str] stack_instance_region_id: The region of the stack instance.
        :param pulumi.Input[str] status: The status of the stack instance. Valid values: `CURRENT` or `OUTDATED`. 
               * `CURRENT`: The stack corresponding to the stack instance is up to date with the stack group.
               * `OUTDATED`: The stack corresponding to the stack instance is not up to date with the stack group. The `OUTDATED` state has the following possible causes:
               * When the CreateStackInstances operation is called to create stack instances, the corresponding stacks fail to be created.
               * When the UpdateStackInstances or UpdateStackGroup operation is called to update stack instances, the corresponding stacks fail to be updated, or only some of the stack instances are updated.
               * The create or update operation is not complete.
        :param pulumi.Input[str] timeout_in_minutes: The timeout period that is specified for the stack creation request. Default value: `60`. Unit: `minutes`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StackInstanceState.__new__(_StackInstanceState)

        __props__.__dict__["operation_description"] = operation_description
        __props__.__dict__["operation_preferences"] = operation_preferences
        __props__.__dict__["parameter_overrides"] = parameter_overrides
        __props__.__dict__["retain_stacks"] = retain_stacks
        __props__.__dict__["stack_group_name"] = stack_group_name
        __props__.__dict__["stack_instance_account_id"] = stack_instance_account_id
        __props__.__dict__["stack_instance_region_id"] = stack_instance_region_id
        __props__.__dict__["status"] = status
        __props__.__dict__["timeout_in_minutes"] = timeout_in_minutes
        return StackInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="operationDescription")
    def operation_description(self) -> pulumi.Output[Optional[str]]:
        """
        The operation description.
        """
        return pulumi.get(self, "operation_description")

    @property
    @pulumi.getter(name="operationPreferences")
    def operation_preferences(self) -> pulumi.Output[Optional[str]]:
        """
        The operation preferences. The operation settings. The following fields are supported:
        """
        return pulumi.get(self, "operation_preferences")

    @property
    @pulumi.getter(name="parameterOverrides")
    def parameter_overrides(self) -> pulumi.Output[Optional[Sequence['outputs.StackInstanceParameterOverride']]]:
        """
        ParameterOverrides. See the following `Block parameter_overrides`.
        """
        return pulumi.get(self, "parameter_overrides")

    @property
    @pulumi.getter(name="retainStacks")
    def retain_stacks(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to retain the stack corresponding to the stack instance.Default value `false`. **NOTE:** When `retain_stacks` is `true`, the stack is retained. If the stack is retained, the corresponding stack is not deleted when the stack instance is deleted from the stack group.
        """
        return pulumi.get(self, "retain_stacks")

    @property
    @pulumi.getter(name="stackGroupName")
    def stack_group_name(self) -> pulumi.Output[str]:
        """
        The name of the stack group.
        """
        return pulumi.get(self, "stack_group_name")

    @property
    @pulumi.getter(name="stackInstanceAccountId")
    def stack_instance_account_id(self) -> pulumi.Output[str]:
        """
        The account to which the stack instance belongs.
        """
        return pulumi.get(self, "stack_instance_account_id")

    @property
    @pulumi.getter(name="stackInstanceRegionId")
    def stack_instance_region_id(self) -> pulumi.Output[str]:
        """
        The region of the stack instance.
        """
        return pulumi.get(self, "stack_instance_region_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the stack instance. Valid values: `CURRENT` or `OUTDATED`. 
        * `CURRENT`: The stack corresponding to the stack instance is up to date with the stack group.
        * `OUTDATED`: The stack corresponding to the stack instance is not up to date with the stack group. The `OUTDATED` state has the following possible causes:
        * When the CreateStackInstances operation is called to create stack instances, the corresponding stacks fail to be created.
        * When the UpdateStackInstances or UpdateStackGroup operation is called to update stack instances, the corresponding stacks fail to be updated, or only some of the stack instances are updated.
        * The create or update operation is not complete.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="timeoutInMinutes")
    def timeout_in_minutes(self) -> pulumi.Output[Optional[str]]:
        """
        The timeout period that is specified for the stack creation request. Default value: `60`. Unit: `minutes`.
        """
        return pulumi.get(self, "timeout_in_minutes")

