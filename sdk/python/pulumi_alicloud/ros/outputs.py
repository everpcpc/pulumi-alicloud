# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'ChangeSetParameter',
    'StackGroupParameter',
    'StackParameter',
    'GetChangeSetsSetResult',
    'GetChangeSetsSetParameterResult',
    'GetStackGroupsGroupResult',
    'GetStackGroupsGroupParameterResult',
    'GetStacksStackResult',
    'GetStacksStackParameterResult',
    'GetTemplatesTemplateResult',
]

@pulumi.output_type
class ChangeSetParameter(dict):
    def __init__(__self__, *,
                 parameter_key: str,
                 parameter_value: str):
        """
        :param str parameter_key: The parameter key.
        :param str parameter_value: The parameter value.
        """
        pulumi.set(__self__, "parameter_key", parameter_key)
        pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="parameterKey")
    def parameter_key(self) -> str:
        """
        The parameter key.
        """
        return pulumi.get(self, "parameter_key")

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> str:
        """
        The parameter value.
        """
        return pulumi.get(self, "parameter_value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StackGroupParameter(dict):
    def __init__(__self__, *,
                 parameter_key: Optional[str] = None,
                 parameter_value: Optional[str] = None):
        """
        :param str parameter_key: The parameter key.
        :param str parameter_value: The parameter value.
        """
        if parameter_key is not None:
            pulumi.set(__self__, "parameter_key", parameter_key)
        if parameter_value is not None:
            pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="parameterKey")
    def parameter_key(self) -> Optional[str]:
        """
        The parameter key.
        """
        return pulumi.get(self, "parameter_key")

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> Optional[str]:
        """
        The parameter value.
        """
        return pulumi.get(self, "parameter_value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StackParameter(dict):
    def __init__(__self__, *,
                 parameter_value: str,
                 parameter_key: Optional[str] = None):
        """
        :param str parameter_value: The parameter value.
        :param str parameter_key: The parameter key.
        """
        pulumi.set(__self__, "parameter_value", parameter_value)
        if parameter_key is not None:
            pulumi.set(__self__, "parameter_key", parameter_key)

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> str:
        """
        The parameter value.
        """
        return pulumi.get(self, "parameter_value")

    @property
    @pulumi.getter(name="parameterKey")
    def parameter_key(self) -> Optional[str]:
        """
        The parameter key.
        """
        return pulumi.get(self, "parameter_key")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetChangeSetsSetResult(dict):
    def __init__(__self__, *,
                 change_set_id: str,
                 change_set_name: str,
                 change_set_type: str,
                 description: str,
                 disable_rollback: bool,
                 execution_status: str,
                 id: str,
                 parameters: Sequence['outputs.GetChangeSetsSetParameterResult'],
                 stack_id: str,
                 stack_name: str,
                 status: str,
                 template_body: str,
                 timeout_in_minutes: int):
        """
        :param str change_set_id: The ID of the change set.
        :param str change_set_name: The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
        :param str change_set_type: The type of the change set. Valid values:  CREATE: creates a change set for a new stack. UPDATE: creates a change set for an existing stack. IMPORT: creates a change set for a new stack or an existing stack to import non-ROS-managed resources. If you create a change set for a new stack, ROS creates a stack that has a unique stack ID. The stack is in the REVIEW_IN_PROGRESS state until you execute the change set.  You cannot use the UPDATE type to create a change set for a new stack or the CREATE type to create a change set for an existing stack.
        :param str description: The description of the change set. The description can be up to 1,024 bytes in length.
        :param bool disable_rollback: Specifies whether to disable rollback on stack creation failure. Default value: false.  Valid values:  true: disables rollback on stack creation failure. false: enables rollback on stack creation failure. Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
        :param str execution_status: The execution status of change set N. Maximum value of N: 5. Valid values:  UNAVAILABLE AVAILABLE EXECUTE_IN_PROGRESS EXECUTE_COMPLETE EXECUTE_FAILED OBSOLETE.
        :param str id: The ID of the Change Set.
        :param Sequence['GetChangeSetsSetParameterArgs'] parameters: Parameters.
        :param str stack_id: The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
        :param str stack_name: The name of the stack for which you want to create the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.  Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
        :param str status: The status of the change set.
        :param str template_body: The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You can specify one of TemplateBody or TemplateURL parameters, but you cannot specify both of them.
        :param int timeout_in_minutes: Timeout In Minutes.
        """
        pulumi.set(__self__, "change_set_id", change_set_id)
        pulumi.set(__self__, "change_set_name", change_set_name)
        pulumi.set(__self__, "change_set_type", change_set_type)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "disable_rollback", disable_rollback)
        pulumi.set(__self__, "execution_status", execution_status)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "parameters", parameters)
        pulumi.set(__self__, "stack_id", stack_id)
        pulumi.set(__self__, "stack_name", stack_name)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "template_body", template_body)
        pulumi.set(__self__, "timeout_in_minutes", timeout_in_minutes)

    @property
    @pulumi.getter(name="changeSetId")
    def change_set_id(self) -> str:
        """
        The ID of the change set.
        """
        return pulumi.get(self, "change_set_id")

    @property
    @pulumi.getter(name="changeSetName")
    def change_set_name(self) -> str:
        """
        The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
        """
        return pulumi.get(self, "change_set_name")

    @property
    @pulumi.getter(name="changeSetType")
    def change_set_type(self) -> str:
        """
        The type of the change set. Valid values:  CREATE: creates a change set for a new stack. UPDATE: creates a change set for an existing stack. IMPORT: creates a change set for a new stack or an existing stack to import non-ROS-managed resources. If you create a change set for a new stack, ROS creates a stack that has a unique stack ID. The stack is in the REVIEW_IN_PROGRESS state until you execute the change set.  You cannot use the UPDATE type to create a change set for a new stack or the CREATE type to create a change set for an existing stack.
        """
        return pulumi.get(self, "change_set_type")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the change set. The description can be up to 1,024 bytes in length.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disableRollback")
    def disable_rollback(self) -> bool:
        """
        Specifies whether to disable rollback on stack creation failure. Default value: false.  Valid values:  true: disables rollback on stack creation failure. false: enables rollback on stack creation failure. Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
        """
        return pulumi.get(self, "disable_rollback")

    @property
    @pulumi.getter(name="executionStatus")
    def execution_status(self) -> str:
        """
        The execution status of change set N. Maximum value of N: 5. Valid values:  UNAVAILABLE AVAILABLE EXECUTE_IN_PROGRESS EXECUTE_COMPLETE EXECUTE_FAILED OBSOLETE.
        """
        return pulumi.get(self, "execution_status")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Change Set.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def parameters(self) -> Sequence['outputs.GetChangeSetsSetParameterResult']:
        """
        Parameters.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> str:
        """
        The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter(name="stackName")
    def stack_name(self) -> str:
        """
        The name of the stack for which you want to create the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.  Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
        """
        return pulumi.get(self, "stack_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the change set.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="templateBody")
    def template_body(self) -> str:
        """
        The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You can specify one of TemplateBody or TemplateURL parameters, but you cannot specify both of them.
        """
        return pulumi.get(self, "template_body")

    @property
    @pulumi.getter(name="timeoutInMinutes")
    def timeout_in_minutes(self) -> int:
        """
        Timeout In Minutes.
        """
        return pulumi.get(self, "timeout_in_minutes")


@pulumi.output_type
class GetChangeSetsSetParameterResult(dict):
    def __init__(__self__, *,
                 parameter_key: str,
                 parameter_value: str):
        """
        :param str parameter_key: The parameters.
        :param str parameter_value: The parameters.
        """
        pulumi.set(__self__, "parameter_key", parameter_key)
        pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="parameterKey")
    def parameter_key(self) -> str:
        """
        The parameters.
        """
        return pulumi.get(self, "parameter_key")

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> str:
        """
        The parameters.
        """
        return pulumi.get(self, "parameter_value")


@pulumi.output_type
class GetStackGroupsGroupResult(dict):
    def __init__(__self__, *,
                 administration_role_name: str,
                 description: str,
                 execution_role_name: str,
                 id: str,
                 parameters: Sequence['outputs.GetStackGroupsGroupParameterResult'],
                 stack_group_id: str,
                 stack_group_name: str,
                 status: str,
                 template_body: str):
        """
        :param str administration_role_name: The name of the RAM administrator role assumed by ROS.
        :param str description: The description of the stack group.
        :param str execution_role_name: The name of the RAM execution role assumed by the administrator role.
        :param str id: The ID of the Stack Group.
        :param Sequence['GetStackGroupsGroupParameterArgs'] parameters: The parameters.
        :param str stack_group_id: The id of Stack Group.
        :param str stack_group_name: The name of the stack group..
        :param str status: The status of Stack Group.
        :param str template_body: The structure that contains the template body.
        """
        pulumi.set(__self__, "administration_role_name", administration_role_name)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "execution_role_name", execution_role_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "parameters", parameters)
        pulumi.set(__self__, "stack_group_id", stack_group_id)
        pulumi.set(__self__, "stack_group_name", stack_group_name)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "template_body", template_body)

    @property
    @pulumi.getter(name="administrationRoleName")
    def administration_role_name(self) -> str:
        """
        The name of the RAM administrator role assumed by ROS.
        """
        return pulumi.get(self, "administration_role_name")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the stack group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="executionRoleName")
    def execution_role_name(self) -> str:
        """
        The name of the RAM execution role assumed by the administrator role.
        """
        return pulumi.get(self, "execution_role_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Stack Group.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def parameters(self) -> Sequence['outputs.GetStackGroupsGroupParameterResult']:
        """
        The parameters.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="stackGroupId")
    def stack_group_id(self) -> str:
        """
        The id of Stack Group.
        """
        return pulumi.get(self, "stack_group_id")

    @property
    @pulumi.getter(name="stackGroupName")
    def stack_group_name(self) -> str:
        """
        The name of the stack group..
        """
        return pulumi.get(self, "stack_group_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of Stack Group.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="templateBody")
    def template_body(self) -> str:
        """
        The structure that contains the template body.
        """
        return pulumi.get(self, "template_body")


@pulumi.output_type
class GetStackGroupsGroupParameterResult(dict):
    def __init__(__self__, *,
                 parameter_key: str,
                 parameter_value: str):
        """
        :param str parameter_key: The parameter key.
        :param str parameter_value: The parameter value.
        """
        pulumi.set(__self__, "parameter_key", parameter_key)
        pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="parameterKey")
    def parameter_key(self) -> str:
        """
        The parameter key.
        """
        return pulumi.get(self, "parameter_key")

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> str:
        """
        The parameter value.
        """
        return pulumi.get(self, "parameter_value")


@pulumi.output_type
class GetStacksStackResult(dict):
    def __init__(__self__, *,
                 deletion_protection: str,
                 description: str,
                 disable_rollback: bool,
                 drift_detection_time: str,
                 id: str,
                 parameters: Sequence['outputs.GetStacksStackParameterResult'],
                 parent_stack_id: str,
                 ram_role_name: str,
                 root_stack_id: str,
                 stack_drift_status: str,
                 stack_id: str,
                 stack_name: str,
                 stack_policy_body: str,
                 status: str,
                 status_reason: str,
                 tags: Mapping[str, Any],
                 template_description: str,
                 timeout_in_minutes: int):
        """
        :param str deletion_protection: Specifies whether to enable deletion protection on the stack.
        :param str description: The Description of the Stack.
        :param bool disable_rollback: Specifies whether to disable rollback on stack creation failure..
        :param str drift_detection_time: Drift DetectionTime.
        :param str id: The ID of the Stack.
        :param Sequence['GetStacksStackParameterArgs'] parameters: The parameters.
        :param str parent_stack_id: Parent Stack Id.
        :param str ram_role_name: The RamRoleName.
        :param str root_stack_id: Root Stack Id.
        :param str stack_drift_status: Stack DriftStatus.
        :param str stack_id: Stack Id.
        :param str stack_name: Stack Name.
        :param str stack_policy_body: The structure that contains the stack policy body.
        :param str status: The status of Stack. Valid Values: `CREATE_COMPLETE`, `CREATE_FAILED`, `CREATE_IN_PROGRESS`, `DELETE_COMPLETE`, `DELETE_FAILED`, `DELETE_IN_PROGRESS`, `ROLLBACK_COMPLETE`, `ROLLBACK_FAILED`, `ROLLBACK_IN_PROGRESS`.
        :param str status_reason: Status Reason.
        :param Mapping[str, Any] tags: Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{"key1":"value1"}`.
        :param str template_description: Template Description.
        :param int timeout_in_minutes: Specifies whether to use the values that were passed last time for the parameters that you do not specify in the current request.
        """
        pulumi.set(__self__, "deletion_protection", deletion_protection)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "disable_rollback", disable_rollback)
        pulumi.set(__self__, "drift_detection_time", drift_detection_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "parameters", parameters)
        pulumi.set(__self__, "parent_stack_id", parent_stack_id)
        pulumi.set(__self__, "ram_role_name", ram_role_name)
        pulumi.set(__self__, "root_stack_id", root_stack_id)
        pulumi.set(__self__, "stack_drift_status", stack_drift_status)
        pulumi.set(__self__, "stack_id", stack_id)
        pulumi.set(__self__, "stack_name", stack_name)
        pulumi.set(__self__, "stack_policy_body", stack_policy_body)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "status_reason", status_reason)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "template_description", template_description)
        pulumi.set(__self__, "timeout_in_minutes", timeout_in_minutes)

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> str:
        """
        Specifies whether to enable deletion protection on the stack.
        """
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The Description of the Stack.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disableRollback")
    def disable_rollback(self) -> bool:
        """
        Specifies whether to disable rollback on stack creation failure..
        """
        return pulumi.get(self, "disable_rollback")

    @property
    @pulumi.getter(name="driftDetectionTime")
    def drift_detection_time(self) -> str:
        """
        Drift DetectionTime.
        """
        return pulumi.get(self, "drift_detection_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Stack.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def parameters(self) -> Sequence['outputs.GetStacksStackParameterResult']:
        """
        The parameters.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="parentStackId")
    def parent_stack_id(self) -> str:
        """
        Parent Stack Id.
        """
        return pulumi.get(self, "parent_stack_id")

    @property
    @pulumi.getter(name="ramRoleName")
    def ram_role_name(self) -> str:
        """
        The RamRoleName.
        """
        return pulumi.get(self, "ram_role_name")

    @property
    @pulumi.getter(name="rootStackId")
    def root_stack_id(self) -> str:
        """
        Root Stack Id.
        """
        return pulumi.get(self, "root_stack_id")

    @property
    @pulumi.getter(name="stackDriftStatus")
    def stack_drift_status(self) -> str:
        """
        Stack DriftStatus.
        """
        return pulumi.get(self, "stack_drift_status")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> str:
        """
        Stack Id.
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter(name="stackName")
    def stack_name(self) -> str:
        """
        Stack Name.
        """
        return pulumi.get(self, "stack_name")

    @property
    @pulumi.getter(name="stackPolicyBody")
    def stack_policy_body(self) -> str:
        """
        The structure that contains the stack policy body.
        """
        return pulumi.get(self, "stack_policy_body")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of Stack. Valid Values: `CREATE_COMPLETE`, `CREATE_FAILED`, `CREATE_IN_PROGRESS`, `DELETE_COMPLETE`, `DELETE_FAILED`, `DELETE_IN_PROGRESS`, `ROLLBACK_COMPLETE`, `ROLLBACK_FAILED`, `ROLLBACK_IN_PROGRESS`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusReason")
    def status_reason(self) -> str:
        """
        Status Reason.
        """
        return pulumi.get(self, "status_reason")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, Any]:
        """
        Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{"key1":"value1"}`.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateDescription")
    def template_description(self) -> str:
        """
        Template Description.
        """
        return pulumi.get(self, "template_description")

    @property
    @pulumi.getter(name="timeoutInMinutes")
    def timeout_in_minutes(self) -> int:
        """
        Specifies whether to use the values that were passed last time for the parameters that you do not specify in the current request.
        """
        return pulumi.get(self, "timeout_in_minutes")


@pulumi.output_type
class GetStacksStackParameterResult(dict):
    def __init__(__self__, *,
                 parameter_key: str,
                 parameter_value: str):
        """
        :param str parameter_key: The key of parameters.
        :param str parameter_value: The value of parameters.
        """
        pulumi.set(__self__, "parameter_key", parameter_key)
        pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="parameterKey")
    def parameter_key(self) -> str:
        """
        The key of parameters.
        """
        return pulumi.get(self, "parameter_key")

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> str:
        """
        The value of parameters.
        """
        return pulumi.get(self, "parameter_value")


@pulumi.output_type
class GetTemplatesTemplateResult(dict):
    def __init__(__self__, *,
                 change_set_id: str,
                 description: str,
                 id: str,
                 share_type: str,
                 stack_group_name: str,
                 stack_id: str,
                 tags: Mapping[str, Any],
                 template_body: str,
                 template_id: str,
                 template_name: str,
                 template_version: str):
        """
        :param str change_set_id: The ID of the change set.
        :param str description: The description of the template. The description can be up to 256 characters in length.
        :param str id: The ID of the Template.
        :param str share_type: Share Type.
        :param str stack_group_name: The name of the stack group. The name must be unique in a region.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
        :param str stack_id: The ID of the stack.
        :param Mapping[str, Any] tags: Tags.
        :param str template_body: The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You must specify one of the TemplateBody and TemplateURL parameters, but you cannot specify both of them.
        :param str template_id: The ID of the template.
        :param str template_name: The name of the template.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
        :param str template_version: Template Version.
        """
        pulumi.set(__self__, "change_set_id", change_set_id)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "share_type", share_type)
        pulumi.set(__self__, "stack_group_name", stack_group_name)
        pulumi.set(__self__, "stack_id", stack_id)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "template_body", template_body)
        pulumi.set(__self__, "template_id", template_id)
        pulumi.set(__self__, "template_name", template_name)
        pulumi.set(__self__, "template_version", template_version)

    @property
    @pulumi.getter(name="changeSetId")
    def change_set_id(self) -> str:
        """
        The ID of the change set.
        """
        return pulumi.get(self, "change_set_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the template. The description can be up to 256 characters in length.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Template.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="shareType")
    def share_type(self) -> str:
        """
        Share Type.
        """
        return pulumi.get(self, "share_type")

    @property
    @pulumi.getter(name="stackGroupName")
    def stack_group_name(self) -> str:
        """
        The name of the stack group. The name must be unique in a region.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
        """
        return pulumi.get(self, "stack_group_name")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> str:
        """
        The ID of the stack.
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, Any]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateBody")
    def template_body(self) -> str:
        """
        The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You must specify one of the TemplateBody and TemplateURL parameters, but you cannot specify both of them.
        """
        return pulumi.get(self, "template_body")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> str:
        """
        The ID of the template.
        """
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> str:
        """
        The name of the template.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
        """
        return pulumi.get(self, "template_name")

    @property
    @pulumi.getter(name="templateVersion")
    def template_version(self) -> str:
        """
        Template Version.
        """
        return pulumi.get(self, "template_version")


