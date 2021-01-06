# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Stack']


class Stack(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_option: Optional[pulumi.Input[str]] = None,
                 deletion_protection: Optional[pulumi.Input[str]] = None,
                 disable_rollback: Optional[pulumi.Input[bool]] = None,
                 notification_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StackParameterArgs']]]]] = None,
                 ram_role_name: Optional[pulumi.Input[str]] = None,
                 replacement_option: Optional[pulumi.Input[str]] = None,
                 retain_all_resources: Optional[pulumi.Input[bool]] = None,
                 retain_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 stack_name: Optional[pulumi.Input[str]] = None,
                 stack_policy_body: Optional[pulumi.Input[str]] = None,
                 stack_policy_during_update_body: Optional[pulumi.Input[str]] = None,
                 stack_policy_during_update_url: Optional[pulumi.Input[str]] = None,
                 stack_policy_url: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 template_body: Optional[pulumi.Input[str]] = None,
                 template_url: Optional[pulumi.Input[str]] = None,
                 template_version: Optional[pulumi.Input[str]] = None,
                 timeout_in_minutes: Optional[pulumi.Input[int]] = None,
                 use_previous_parameters: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a ROS Stack resource.

        For information about ROS Stack and how to use it, see [What is Stack](https://www.alibabacloud.com/help/en/doc-detail/132086.htm).

        > **NOTE:** Available in v1.106.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.ros.Stack("example",
            stack_name="tf-testaccstack",
            stack_policy_body=\"\"\"    {
            	"Statement": [{
            		"Action": "Update:Delete",
            		"Resource": "*",
            		"Effect": "Allow",
            		"Principal": "*"
            	}]
            }
            
        \"\"\",
            template_body=\"\"\"    {
            	"ROSTemplateFormatVersion": "2015-09-01"
            }
            
        \"\"\")
        ```

        ## Import

        ROS Stack can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ros/stack:Stack example <stack_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_option: Specifies whether to delete the stack after it is created.
        :param pulumi.Input[str] deletion_protection: Specifies whether to enable deletion protection on the stack. Valid values: `Disabled`, `Enabled`. Default to: `Disabled`
        :param pulumi.Input[bool] disable_rollback: Specifies whether to disable rollback on stack creation failure. Default to: `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_urls: The callback URL for receiving stack event N. Only HTTP POST is supported. Maximum value of N: 5.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StackParameterArgs']]]] parameters: The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.
        :param pulumi.Input[str] ram_role_name: The name of the RAM role. ROS assumes the specified RAM role to create the stack and call API operations by using the credentials of the role.
        :param pulumi.Input[str] replacement_option: Specifies whether to enable replacement update after a resource attribute that does not support modification update is changed. Modification update keeps the physical ID of the resource unchanged. However, the resource is deleted and then recreated, and its physical ID is changed if replacement update is enabled.
        :param pulumi.Input[bool] retain_all_resources: The retain all resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] retain_resources: Specifies whether to retain the resources in the stack.
        :param pulumi.Input[str] stack_name: The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
        :param pulumi.Input[str] stack_policy_body: The structure that contains the stack policy body. The stack policy body must be 1 to 16,384 bytes in length.
        :param pulumi.Input[str] stack_policy_during_update_body: The structure that contains the body of the temporary overriding stack policy. The stack policy body must be 1 to 16,384 bytes in length.
        :param pulumi.Input[str] stack_policy_during_update_url: The URL of the file that contains the temporary overriding stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
        :param pulumi.Input[str] stack_policy_url: The URL of the file that contains the stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] template_body: The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.
        :param pulumi.Input[str] template_url: The URL of the file that contains the template body. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template must be 1 to 524,288 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
        :param pulumi.Input[str] template_version: The version of the template.
        :param pulumi.Input[int] timeout_in_minutes: The timeout period that is specified for the stack creation request. Default to: `60`.
        :param pulumi.Input[bool] use_previous_parameters: Specifies whether to use the values that were passed last time for the parameters that you do not specify in the current request.
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

            __props__['create_option'] = create_option
            __props__['deletion_protection'] = deletion_protection
            __props__['disable_rollback'] = disable_rollback
            __props__['notification_urls'] = notification_urls
            __props__['parameters'] = parameters
            __props__['ram_role_name'] = ram_role_name
            __props__['replacement_option'] = replacement_option
            __props__['retain_all_resources'] = retain_all_resources
            __props__['retain_resources'] = retain_resources
            if stack_name is None:
                raise TypeError("Missing required property 'stack_name'")
            __props__['stack_name'] = stack_name
            __props__['stack_policy_body'] = stack_policy_body
            __props__['stack_policy_during_update_body'] = stack_policy_during_update_body
            __props__['stack_policy_during_update_url'] = stack_policy_during_update_url
            __props__['stack_policy_url'] = stack_policy_url
            __props__['tags'] = tags
            __props__['template_body'] = template_body
            __props__['template_url'] = template_url
            __props__['template_version'] = template_version
            __props__['timeout_in_minutes'] = timeout_in_minutes
            __props__['use_previous_parameters'] = use_previous_parameters
            __props__['status'] = None
        super(Stack, __self__).__init__(
            'alicloud:ros/stack:Stack',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_option: Optional[pulumi.Input[str]] = None,
            deletion_protection: Optional[pulumi.Input[str]] = None,
            disable_rollback: Optional[pulumi.Input[bool]] = None,
            notification_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StackParameterArgs']]]]] = None,
            ram_role_name: Optional[pulumi.Input[str]] = None,
            replacement_option: Optional[pulumi.Input[str]] = None,
            retain_all_resources: Optional[pulumi.Input[bool]] = None,
            retain_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            stack_name: Optional[pulumi.Input[str]] = None,
            stack_policy_body: Optional[pulumi.Input[str]] = None,
            stack_policy_during_update_body: Optional[pulumi.Input[str]] = None,
            stack_policy_during_update_url: Optional[pulumi.Input[str]] = None,
            stack_policy_url: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            template_body: Optional[pulumi.Input[str]] = None,
            template_url: Optional[pulumi.Input[str]] = None,
            template_version: Optional[pulumi.Input[str]] = None,
            timeout_in_minutes: Optional[pulumi.Input[int]] = None,
            use_previous_parameters: Optional[pulumi.Input[bool]] = None) -> 'Stack':
        """
        Get an existing Stack resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_option: Specifies whether to delete the stack after it is created.
        :param pulumi.Input[str] deletion_protection: Specifies whether to enable deletion protection on the stack. Valid values: `Disabled`, `Enabled`. Default to: `Disabled`
        :param pulumi.Input[bool] disable_rollback: Specifies whether to disable rollback on stack creation failure. Default to: `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_urls: The callback URL for receiving stack event N. Only HTTP POST is supported. Maximum value of N: 5.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StackParameterArgs']]]] parameters: The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.
        :param pulumi.Input[str] ram_role_name: The name of the RAM role. ROS assumes the specified RAM role to create the stack and call API operations by using the credentials of the role.
        :param pulumi.Input[str] replacement_option: Specifies whether to enable replacement update after a resource attribute that does not support modification update is changed. Modification update keeps the physical ID of the resource unchanged. However, the resource is deleted and then recreated, and its physical ID is changed if replacement update is enabled.
        :param pulumi.Input[bool] retain_all_resources: The retain all resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] retain_resources: Specifies whether to retain the resources in the stack.
        :param pulumi.Input[str] stack_name: The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
        :param pulumi.Input[str] stack_policy_body: The structure that contains the stack policy body. The stack policy body must be 1 to 16,384 bytes in length.
        :param pulumi.Input[str] stack_policy_during_update_body: The structure that contains the body of the temporary overriding stack policy. The stack policy body must be 1 to 16,384 bytes in length.
        :param pulumi.Input[str] stack_policy_during_update_url: The URL of the file that contains the temporary overriding stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
        :param pulumi.Input[str] stack_policy_url: The URL of the file that contains the stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
        :param pulumi.Input[str] status: The status of Stack.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] template_body: The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.
        :param pulumi.Input[str] template_url: The URL of the file that contains the template body. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template must be 1 to 524,288 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
        :param pulumi.Input[str] template_version: The version of the template.
        :param pulumi.Input[int] timeout_in_minutes: The timeout period that is specified for the stack creation request. Default to: `60`.
        :param pulumi.Input[bool] use_previous_parameters: Specifies whether to use the values that were passed last time for the parameters that you do not specify in the current request.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["create_option"] = create_option
        __props__["deletion_protection"] = deletion_protection
        __props__["disable_rollback"] = disable_rollback
        __props__["notification_urls"] = notification_urls
        __props__["parameters"] = parameters
        __props__["ram_role_name"] = ram_role_name
        __props__["replacement_option"] = replacement_option
        __props__["retain_all_resources"] = retain_all_resources
        __props__["retain_resources"] = retain_resources
        __props__["stack_name"] = stack_name
        __props__["stack_policy_body"] = stack_policy_body
        __props__["stack_policy_during_update_body"] = stack_policy_during_update_body
        __props__["stack_policy_during_update_url"] = stack_policy_during_update_url
        __props__["stack_policy_url"] = stack_policy_url
        __props__["status"] = status
        __props__["tags"] = tags
        __props__["template_body"] = template_body
        __props__["template_url"] = template_url
        __props__["template_version"] = template_version
        __props__["timeout_in_minutes"] = timeout_in_minutes
        __props__["use_previous_parameters"] = use_previous_parameters
        return Stack(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createOption")
    def create_option(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether to delete the stack after it is created.
        """
        return pulumi.get(self, "create_option")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether to enable deletion protection on the stack. Valid values: `Disabled`, `Enabled`. Default to: `Disabled`
        """
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter(name="disableRollback")
    def disable_rollback(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to disable rollback on stack creation failure. Default to: `false`.
        """
        return pulumi.get(self, "disable_rollback")

    @property
    @pulumi.getter(name="notificationUrls")
    def notification_urls(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The callback URL for receiving stack event N. Only HTTP POST is supported. Maximum value of N: 5.
        """
        return pulumi.get(self, "notification_urls")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Sequence['outputs.StackParameter']]]:
        """
        The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="ramRoleName")
    def ram_role_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the RAM role. ROS assumes the specified RAM role to create the stack and call API operations by using the credentials of the role.
        """
        return pulumi.get(self, "ram_role_name")

    @property
    @pulumi.getter(name="replacementOption")
    def replacement_option(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether to enable replacement update after a resource attribute that does not support modification update is changed. Modification update keeps the physical ID of the resource unchanged. However, the resource is deleted and then recreated, and its physical ID is changed if replacement update is enabled.
        """
        return pulumi.get(self, "replacement_option")

    @property
    @pulumi.getter(name="retainAllResources")
    def retain_all_resources(self) -> pulumi.Output[Optional[bool]]:
        """
        The retain all resources.
        """
        return pulumi.get(self, "retain_all_resources")

    @property
    @pulumi.getter(name="retainResources")
    def retain_resources(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Specifies whether to retain the resources in the stack.
        """
        return pulumi.get(self, "retain_resources")

    @property
    @pulumi.getter(name="stackName")
    def stack_name(self) -> pulumi.Output[str]:
        """
        The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
        """
        return pulumi.get(self, "stack_name")

    @property
    @pulumi.getter(name="stackPolicyBody")
    def stack_policy_body(self) -> pulumi.Output[Optional[str]]:
        """
        The structure that contains the stack policy body. The stack policy body must be 1 to 16,384 bytes in length.
        """
        return pulumi.get(self, "stack_policy_body")

    @property
    @pulumi.getter(name="stackPolicyDuringUpdateBody")
    def stack_policy_during_update_body(self) -> pulumi.Output[Optional[str]]:
        """
        The structure that contains the body of the temporary overriding stack policy. The stack policy body must be 1 to 16,384 bytes in length.
        """
        return pulumi.get(self, "stack_policy_during_update_body")

    @property
    @pulumi.getter(name="stackPolicyDuringUpdateUrl")
    def stack_policy_during_update_url(self) -> pulumi.Output[Optional[str]]:
        """
        The URL of the file that contains the temporary overriding stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
        """
        return pulumi.get(self, "stack_policy_during_update_url")

    @property
    @pulumi.getter(name="stackPolicyUrl")
    def stack_policy_url(self) -> pulumi.Output[Optional[str]]:
        """
        The URL of the file that contains the stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
        """
        return pulumi.get(self, "stack_policy_url")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of Stack.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateBody")
    def template_body(self) -> pulumi.Output[Optional[str]]:
        """
        The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.
        """
        return pulumi.get(self, "template_body")

    @property
    @pulumi.getter(name="templateUrl")
    def template_url(self) -> pulumi.Output[Optional[str]]:
        """
        The URL of the file that contains the template body. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template must be 1 to 524,288 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.
        """
        return pulumi.get(self, "template_url")

    @property
    @pulumi.getter(name="templateVersion")
    def template_version(self) -> pulumi.Output[Optional[str]]:
        """
        The version of the template.
        """
        return pulumi.get(self, "template_version")

    @property
    @pulumi.getter(name="timeoutInMinutes")
    def timeout_in_minutes(self) -> pulumi.Output[Optional[int]]:
        """
        The timeout period that is specified for the stack creation request. Default to: `60`.
        """
        return pulumi.get(self, "timeout_in_minutes")

    @property
    @pulumi.getter(name="usePreviousParameters")
    def use_previous_parameters(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to use the values that were passed last time for the parameters that you do not specify in the current request.
        """
        return pulumi.get(self, "use_previous_parameters")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
