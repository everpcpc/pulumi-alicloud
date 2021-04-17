# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RegistryEnterpriseSyncRuleArgs', 'RegistryEnterpriseSyncRule']

@pulumi.input_type
class RegistryEnterpriseSyncRuleArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 namespace_name: pulumi.Input[str],
                 tag_filter: pulumi.Input[str],
                 target_instance_id: pulumi.Input[str],
                 target_namespace_name: pulumi.Input[str],
                 target_region_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 repo_name: Optional[pulumi.Input[str]] = None,
                 target_repo_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RegistryEnterpriseSyncRule resource.
        :param pulumi.Input[str] instance_id: ID of Container Registry Enterprise Edition source instance.
        :param pulumi.Input[str] namespace_name: Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
        :param pulumi.Input[str] tag_filter: The regular expression used to filter image tags for synchronization in the source repository.
        :param pulumi.Input[str] target_instance_id: ID of Container Registry Enterprise Edition target instance to be synchronized.
        :param pulumi.Input[str] target_namespace_name: Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
        :param pulumi.Input[str] target_region_id: The target region to be synchronized.
        :param pulumi.Input[str] name: Name of Container Registry Enterprise Edition sync rule.
        :param pulumi.Input[str] repo_name: Name of the source repository which should be set together with `target_repo_name`, if empty means that the synchronization scope is the entire namespace level.
        :param pulumi.Input[str] target_repo_name: Name of the target repository.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "namespace_name", namespace_name)
        pulumi.set(__self__, "tag_filter", tag_filter)
        pulumi.set(__self__, "target_instance_id", target_instance_id)
        pulumi.set(__self__, "target_namespace_name", target_namespace_name)
        pulumi.set(__self__, "target_region_id", target_region_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if repo_name is not None:
            pulumi.set(__self__, "repo_name", repo_name)
        if target_repo_name is not None:
            pulumi.set(__self__, "target_repo_name", target_repo_name)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        ID of Container Registry Enterprise Edition source instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Input[str]:
        """
        Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="tagFilter")
    def tag_filter(self) -> pulumi.Input[str]:
        """
        The regular expression used to filter image tags for synchronization in the source repository.
        """
        return pulumi.get(self, "tag_filter")

    @tag_filter.setter
    def tag_filter(self, value: pulumi.Input[str]):
        pulumi.set(self, "tag_filter", value)

    @property
    @pulumi.getter(name="targetInstanceId")
    def target_instance_id(self) -> pulumi.Input[str]:
        """
        ID of Container Registry Enterprise Edition target instance to be synchronized.
        """
        return pulumi.get(self, "target_instance_id")

    @target_instance_id.setter
    def target_instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_instance_id", value)

    @property
    @pulumi.getter(name="targetNamespaceName")
    def target_namespace_name(self) -> pulumi.Input[str]:
        """
        Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
        """
        return pulumi.get(self, "target_namespace_name")

    @target_namespace_name.setter
    def target_namespace_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_namespace_name", value)

    @property
    @pulumi.getter(name="targetRegionId")
    def target_region_id(self) -> pulumi.Input[str]:
        """
        The target region to be synchronized.
        """
        return pulumi.get(self, "target_region_id")

    @target_region_id.setter
    def target_region_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_region_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Container Registry Enterprise Edition sync rule.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="repoName")
    def repo_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the source repository which should be set together with `target_repo_name`, if empty means that the synchronization scope is the entire namespace level.
        """
        return pulumi.get(self, "repo_name")

    @repo_name.setter
    def repo_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repo_name", value)

    @property
    @pulumi.getter(name="targetRepoName")
    def target_repo_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the target repository.
        """
        return pulumi.get(self, "target_repo_name")

    @target_repo_name.setter
    def target_repo_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_repo_name", value)


@pulumi.input_type
class _RegistryEnterpriseSyncRuleState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 repo_name: Optional[pulumi.Input[str]] = None,
                 rule_id: Optional[pulumi.Input[str]] = None,
                 sync_direction: Optional[pulumi.Input[str]] = None,
                 sync_scope: Optional[pulumi.Input[str]] = None,
                 tag_filter: Optional[pulumi.Input[str]] = None,
                 target_instance_id: Optional[pulumi.Input[str]] = None,
                 target_namespace_name: Optional[pulumi.Input[str]] = None,
                 target_region_id: Optional[pulumi.Input[str]] = None,
                 target_repo_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RegistryEnterpriseSyncRule resources.
        :param pulumi.Input[str] instance_id: ID of Container Registry Enterprise Edition source instance.
        :param pulumi.Input[str] name: Name of Container Registry Enterprise Edition sync rule.
        :param pulumi.Input[str] namespace_name: Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
        :param pulumi.Input[str] repo_name: Name of the source repository which should be set together with `target_repo_name`, if empty means that the synchronization scope is the entire namespace level.
        :param pulumi.Input[str] rule_id: The uuid of Container Registry Enterprise Edition sync rule.
        :param pulumi.Input[str] sync_direction: `FROM` or `TO`, the direction of synchronization. `FROM` means source instance, `TO` means target instance.
        :param pulumi.Input[str] sync_scope: `REPO` or `NAMESPACE`,the scope that the synchronization rule applies.
        :param pulumi.Input[str] tag_filter: The regular expression used to filter image tags for synchronization in the source repository.
        :param pulumi.Input[str] target_instance_id: ID of Container Registry Enterprise Edition target instance to be synchronized.
        :param pulumi.Input[str] target_namespace_name: Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
        :param pulumi.Input[str] target_region_id: The target region to be synchronized.
        :param pulumi.Input[str] target_repo_name: Name of the target repository.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace_name is not None:
            pulumi.set(__self__, "namespace_name", namespace_name)
        if repo_name is not None:
            pulumi.set(__self__, "repo_name", repo_name)
        if rule_id is not None:
            pulumi.set(__self__, "rule_id", rule_id)
        if sync_direction is not None:
            pulumi.set(__self__, "sync_direction", sync_direction)
        if sync_scope is not None:
            pulumi.set(__self__, "sync_scope", sync_scope)
        if tag_filter is not None:
            pulumi.set(__self__, "tag_filter", tag_filter)
        if target_instance_id is not None:
            pulumi.set(__self__, "target_instance_id", target_instance_id)
        if target_namespace_name is not None:
            pulumi.set(__self__, "target_namespace_name", target_namespace_name)
        if target_region_id is not None:
            pulumi.set(__self__, "target_region_id", target_region_id)
        if target_repo_name is not None:
            pulumi.set(__self__, "target_repo_name", target_repo_name)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of Container Registry Enterprise Edition source instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Container Registry Enterprise Edition sync rule.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="repoName")
    def repo_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the source repository which should be set together with `target_repo_name`, if empty means that the synchronization scope is the entire namespace level.
        """
        return pulumi.get(self, "repo_name")

    @repo_name.setter
    def repo_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repo_name", value)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        The uuid of Container Registry Enterprise Edition sync rule.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter(name="syncDirection")
    def sync_direction(self) -> Optional[pulumi.Input[str]]:
        """
        `FROM` or `TO`, the direction of synchronization. `FROM` means source instance, `TO` means target instance.
        """
        return pulumi.get(self, "sync_direction")

    @sync_direction.setter
    def sync_direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sync_direction", value)

    @property
    @pulumi.getter(name="syncScope")
    def sync_scope(self) -> Optional[pulumi.Input[str]]:
        """
        `REPO` or `NAMESPACE`,the scope that the synchronization rule applies.
        """
        return pulumi.get(self, "sync_scope")

    @sync_scope.setter
    def sync_scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sync_scope", value)

    @property
    @pulumi.getter(name="tagFilter")
    def tag_filter(self) -> Optional[pulumi.Input[str]]:
        """
        The regular expression used to filter image tags for synchronization in the source repository.
        """
        return pulumi.get(self, "tag_filter")

    @tag_filter.setter
    def tag_filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag_filter", value)

    @property
    @pulumi.getter(name="targetInstanceId")
    def target_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of Container Registry Enterprise Edition target instance to be synchronized.
        """
        return pulumi.get(self, "target_instance_id")

    @target_instance_id.setter
    def target_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_instance_id", value)

    @property
    @pulumi.getter(name="targetNamespaceName")
    def target_namespace_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
        """
        return pulumi.get(self, "target_namespace_name")

    @target_namespace_name.setter
    def target_namespace_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_namespace_name", value)

    @property
    @pulumi.getter(name="targetRegionId")
    def target_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The target region to be synchronized.
        """
        return pulumi.get(self, "target_region_id")

    @target_region_id.setter
    def target_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_region_id", value)

    @property
    @pulumi.getter(name="targetRepoName")
    def target_repo_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the target repository.
        """
        return pulumi.get(self, "target_repo_name")

    @target_repo_name.setter
    def target_repo_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_repo_name", value)


class RegistryEnterpriseSyncRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 repo_name: Optional[pulumi.Input[str]] = None,
                 tag_filter: Optional[pulumi.Input[str]] = None,
                 target_instance_id: Optional[pulumi.Input[str]] = None,
                 target_namespace_name: Optional[pulumi.Input[str]] = None,
                 target_region_id: Optional[pulumi.Input[str]] = None,
                 target_repo_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        This resource will help you to manager Container Registry Enterprise Edition sync rules.

        For information about Container Registry Enterprise Edition sync rules and how to use it, see [Create a Sync Rule](https://www.alibabacloud.com/help/doc-detail/145280.htm)

        > **NOTE:** Available in v1.90.0+.

        > **NOTE:** You need to set your registry password in Container Registry Enterprise Edition console before use this resource.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.cs.RegistryEnterpriseSyncRule("default",
            instance_id="my-source-instance-id",
            namespace_name="my-source-namespace",
            repo_name="my-source-repo",
            tag_filter=".*",
            target_instance_id="my-target-instance-id",
            target_namespace_name="my-target-namespace",
            target_region_id="cn-hangzhou",
            target_repo_name="my-target-repo")
        ```

        ## Import

        Container Registry Enterprise Edition sync rule can be imported using the id. Format to `{instance_id}:{namespace_name}:{rule_id}`, e.g.

        ```sh
         $ pulumi import alicloud:cs/registryEnterpriseSyncRule:RegistryEnterpriseSyncRule default `cri-xxx:my-namespace:crsr-yyy`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: ID of Container Registry Enterprise Edition source instance.
        :param pulumi.Input[str] name: Name of Container Registry Enterprise Edition sync rule.
        :param pulumi.Input[str] namespace_name: Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
        :param pulumi.Input[str] repo_name: Name of the source repository which should be set together with `target_repo_name`, if empty means that the synchronization scope is the entire namespace level.
        :param pulumi.Input[str] tag_filter: The regular expression used to filter image tags for synchronization in the source repository.
        :param pulumi.Input[str] target_instance_id: ID of Container Registry Enterprise Edition target instance to be synchronized.
        :param pulumi.Input[str] target_namespace_name: Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
        :param pulumi.Input[str] target_region_id: The target region to be synchronized.
        :param pulumi.Input[str] target_repo_name: Name of the target repository.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegistryEnterpriseSyncRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource will help you to manager Container Registry Enterprise Edition sync rules.

        For information about Container Registry Enterprise Edition sync rules and how to use it, see [Create a Sync Rule](https://www.alibabacloud.com/help/doc-detail/145280.htm)

        > **NOTE:** Available in v1.90.0+.

        > **NOTE:** You need to set your registry password in Container Registry Enterprise Edition console before use this resource.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.cs.RegistryEnterpriseSyncRule("default",
            instance_id="my-source-instance-id",
            namespace_name="my-source-namespace",
            repo_name="my-source-repo",
            tag_filter=".*",
            target_instance_id="my-target-instance-id",
            target_namespace_name="my-target-namespace",
            target_region_id="cn-hangzhou",
            target_repo_name="my-target-repo")
        ```

        ## Import

        Container Registry Enterprise Edition sync rule can be imported using the id. Format to `{instance_id}:{namespace_name}:{rule_id}`, e.g.

        ```sh
         $ pulumi import alicloud:cs/registryEnterpriseSyncRule:RegistryEnterpriseSyncRule default `cri-xxx:my-namespace:crsr-yyy`
        ```

        :param str resource_name: The name of the resource.
        :param RegistryEnterpriseSyncRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegistryEnterpriseSyncRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 repo_name: Optional[pulumi.Input[str]] = None,
                 tag_filter: Optional[pulumi.Input[str]] = None,
                 target_instance_id: Optional[pulumi.Input[str]] = None,
                 target_namespace_name: Optional[pulumi.Input[str]] = None,
                 target_region_id: Optional[pulumi.Input[str]] = None,
                 target_repo_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            __props__ = RegistryEnterpriseSyncRuleArgs.__new__(RegistryEnterpriseSyncRuleArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            if namespace_name is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_name'")
            __props__.__dict__["namespace_name"] = namespace_name
            __props__.__dict__["repo_name"] = repo_name
            if tag_filter is None and not opts.urn:
                raise TypeError("Missing required property 'tag_filter'")
            __props__.__dict__["tag_filter"] = tag_filter
            if target_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_instance_id'")
            __props__.__dict__["target_instance_id"] = target_instance_id
            if target_namespace_name is None and not opts.urn:
                raise TypeError("Missing required property 'target_namespace_name'")
            __props__.__dict__["target_namespace_name"] = target_namespace_name
            if target_region_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_region_id'")
            __props__.__dict__["target_region_id"] = target_region_id
            __props__.__dict__["target_repo_name"] = target_repo_name
            __props__.__dict__["rule_id"] = None
            __props__.__dict__["sync_direction"] = None
            __props__.__dict__["sync_scope"] = None
        super(RegistryEnterpriseSyncRule, __self__).__init__(
            'alicloud:cs/registryEnterpriseSyncRule:RegistryEnterpriseSyncRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace_name: Optional[pulumi.Input[str]] = None,
            repo_name: Optional[pulumi.Input[str]] = None,
            rule_id: Optional[pulumi.Input[str]] = None,
            sync_direction: Optional[pulumi.Input[str]] = None,
            sync_scope: Optional[pulumi.Input[str]] = None,
            tag_filter: Optional[pulumi.Input[str]] = None,
            target_instance_id: Optional[pulumi.Input[str]] = None,
            target_namespace_name: Optional[pulumi.Input[str]] = None,
            target_region_id: Optional[pulumi.Input[str]] = None,
            target_repo_name: Optional[pulumi.Input[str]] = None) -> 'RegistryEnterpriseSyncRule':
        """
        Get an existing RegistryEnterpriseSyncRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: ID of Container Registry Enterprise Edition source instance.
        :param pulumi.Input[str] name: Name of Container Registry Enterprise Edition sync rule.
        :param pulumi.Input[str] namespace_name: Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
        :param pulumi.Input[str] repo_name: Name of the source repository which should be set together with `target_repo_name`, if empty means that the synchronization scope is the entire namespace level.
        :param pulumi.Input[str] rule_id: The uuid of Container Registry Enterprise Edition sync rule.
        :param pulumi.Input[str] sync_direction: `FROM` or `TO`, the direction of synchronization. `FROM` means source instance, `TO` means target instance.
        :param pulumi.Input[str] sync_scope: `REPO` or `NAMESPACE`,the scope that the synchronization rule applies.
        :param pulumi.Input[str] tag_filter: The regular expression used to filter image tags for synchronization in the source repository.
        :param pulumi.Input[str] target_instance_id: ID of Container Registry Enterprise Edition target instance to be synchronized.
        :param pulumi.Input[str] target_namespace_name: Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
        :param pulumi.Input[str] target_region_id: The target region to be synchronized.
        :param pulumi.Input[str] target_repo_name: Name of the target repository.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RegistryEnterpriseSyncRuleState.__new__(_RegistryEnterpriseSyncRuleState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace_name"] = namespace_name
        __props__.__dict__["repo_name"] = repo_name
        __props__.__dict__["rule_id"] = rule_id
        __props__.__dict__["sync_direction"] = sync_direction
        __props__.__dict__["sync_scope"] = sync_scope
        __props__.__dict__["tag_filter"] = tag_filter
        __props__.__dict__["target_instance_id"] = target_instance_id
        __props__.__dict__["target_namespace_name"] = target_namespace_name
        __props__.__dict__["target_region_id"] = target_region_id
        __props__.__dict__["target_repo_name"] = target_repo_name
        return RegistryEnterpriseSyncRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        ID of Container Registry Enterprise Edition source instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of Container Registry Enterprise Edition sync rule.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Output[str]:
        """
        Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
        """
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="repoName")
    def repo_name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the source repository which should be set together with `target_repo_name`, if empty means that the synchronization scope is the entire namespace level.
        """
        return pulumi.get(self, "repo_name")

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Output[str]:
        """
        The uuid of Container Registry Enterprise Edition sync rule.
        """
        return pulumi.get(self, "rule_id")

    @property
    @pulumi.getter(name="syncDirection")
    def sync_direction(self) -> pulumi.Output[str]:
        """
        `FROM` or `TO`, the direction of synchronization. `FROM` means source instance, `TO` means target instance.
        """
        return pulumi.get(self, "sync_direction")

    @property
    @pulumi.getter(name="syncScope")
    def sync_scope(self) -> pulumi.Output[str]:
        """
        `REPO` or `NAMESPACE`,the scope that the synchronization rule applies.
        """
        return pulumi.get(self, "sync_scope")

    @property
    @pulumi.getter(name="tagFilter")
    def tag_filter(self) -> pulumi.Output[str]:
        """
        The regular expression used to filter image tags for synchronization in the source repository.
        """
        return pulumi.get(self, "tag_filter")

    @property
    @pulumi.getter(name="targetInstanceId")
    def target_instance_id(self) -> pulumi.Output[str]:
        """
        ID of Container Registry Enterprise Edition target instance to be synchronized.
        """
        return pulumi.get(self, "target_instance_id")

    @property
    @pulumi.getter(name="targetNamespaceName")
    def target_namespace_name(self) -> pulumi.Output[str]:
        """
        Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
        """
        return pulumi.get(self, "target_namespace_name")

    @property
    @pulumi.getter(name="targetRegionId")
    def target_region_id(self) -> pulumi.Output[str]:
        """
        The target region to be synchronized.
        """
        return pulumi.get(self, "target_region_id")

    @property
    @pulumi.getter(name="targetRepoName")
    def target_repo_name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the target repository.
        """
        return pulumi.get(self, "target_repo_name")

