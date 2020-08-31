# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'PolicyStatementArgs',
]

@pulumi.input_type
class PolicyStatementArgs:
    def __init__(__self__, *,
                 actions: pulumi.Input[List[pulumi.Input[str]]],
                 effect: pulumi.Input[str],
                 resources: pulumi.Input[List[pulumi.Input[str]]]):
        """
        :param pulumi.Input[List[pulumi.Input[str]]] actions: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of operations for the `resource`. The format of each item in this list is `${service}:${action_name}`, such as `oss:ListBuckets` and `ecs:Describe*`. The `${service}` can be `ecs`, `oss`, `ots` and so on, the `${action_name}` refers to the name of an api interface which related to the `${service}`.
        :param pulumi.Input[str] effect: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) This parameter indicates whether or not the `action` is allowed. Valid values are `Allow` and `Deny`.
        :param pulumi.Input[List[pulumi.Input[str]]] resources: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of specific objects which will be authorized. The format of each item in this list is `acs:${service}:${region}:${account_id}:${relative_id}`, such as `acs:ecs:*:*:instance/inst-002` and `acs:oss:*:1234567890000:mybucket`. The `${service}` can be `ecs`, `oss`, `ots` and so on, the `${region}` is the region info which can use `*` replace when it is not supplied, the `${account_id}` refers to someone's Alicloud account id or you can use `*` to replace, the `${relative_id}` is the resource description section which related to the `${service}`.
        """
        pulumi.set(__self__, "actions", actions)
        pulumi.set(__self__, "effect", effect)
        pulumi.set(__self__, "resources", resources)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of operations for the `resource`. The format of each item in this list is `${service}:${action_name}`, such as `oss:ListBuckets` and `ecs:Describe*`. The `${service}` can be `ecs`, `oss`, `ots` and so on, the `${action_name}` refers to the name of an api interface which related to the `${service}`.
        """
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter
    def effect(self) -> pulumi.Input[str]:
        """
        (It has been deprecated from version 1.49.0, and use field 'document' to replace.) This parameter indicates whether or not the `action` is allowed. Valid values are `Allow` and `Deny`.
        """
        return pulumi.get(self, "effect")

    @effect.setter
    def effect(self, value: pulumi.Input[str]):
        pulumi.set(self, "effect", value)

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of specific objects which will be authorized. The format of each item in this list is `acs:${service}:${region}:${account_id}:${relative_id}`, such as `acs:ecs:*:*:instance/inst-002` and `acs:oss:*:1234567890000:mybucket`. The `${service}` can be `ecs`, `oss`, `ots` and so on, the `${region}` is the region info which can use `*` replace when it is not supplied, the `${account_id}` refers to someone's Alicloud account id or you can use `*` to replace, the `${relative_id}` is the resource description section which related to the `${service}`.
        """
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "resources", value)

