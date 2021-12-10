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

__all__ = ['AggregateCompliancePackArgs', 'AggregateCompliancePack']

@pulumi.input_type
class AggregateCompliancePackArgs:
    def __init__(__self__, *,
                 aggregate_compliance_pack_name: pulumi.Input[str],
                 aggregator_id: pulumi.Input[str],
                 description: pulumi.Input[str],
                 risk_level: pulumi.Input[int],
                 compliance_pack_template_id: Optional[pulumi.Input[str]] = None,
                 config_rule_ids: Optional[pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleIdArgs']]]] = None,
                 config_rules: Optional[pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleArgs']]]] = None):
        """
        The set of arguments for constructing a AggregateCompliancePack resource.
        :param pulumi.Input[str] aggregate_compliance_pack_name: The name of compliance package name. **NOTE:** the `aggregate_compliance_pack_name` supports modification since V1.145.0.
        :param pulumi.Input[str] aggregator_id: The ID of aggregator.
        :param pulumi.Input[str] description: The description of compliance package.
        :param pulumi.Input[int] risk_level: The Risk Level. Valid values: `1`, `2`, `3`.
        :param pulumi.Input[str] compliance_pack_template_id: The Template ID of compliance package.
        :param pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleIdArgs']]] config_rule_ids: A list of Config Rule IDs.
        :param pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleArgs']]] config_rules: A list of Config Rules.
        """
        pulumi.set(__self__, "aggregate_compliance_pack_name", aggregate_compliance_pack_name)
        pulumi.set(__self__, "aggregator_id", aggregator_id)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "risk_level", risk_level)
        if compliance_pack_template_id is not None:
            pulumi.set(__self__, "compliance_pack_template_id", compliance_pack_template_id)
        if config_rule_ids is not None:
            pulumi.set(__self__, "config_rule_ids", config_rule_ids)
        if config_rules is not None:
            warnings.warn("""Field 'config_rules' has been deprecated from provider version 1.141.0. New field 'config_rule_ids' instead.""", DeprecationWarning)
            pulumi.log.warn("""config_rules is deprecated: Field 'config_rules' has been deprecated from provider version 1.141.0. New field 'config_rule_ids' instead.""")
        if config_rules is not None:
            pulumi.set(__self__, "config_rules", config_rules)

    @property
    @pulumi.getter(name="aggregateCompliancePackName")
    def aggregate_compliance_pack_name(self) -> pulumi.Input[str]:
        """
        The name of compliance package name. **NOTE:** the `aggregate_compliance_pack_name` supports modification since V1.145.0.
        """
        return pulumi.get(self, "aggregate_compliance_pack_name")

    @aggregate_compliance_pack_name.setter
    def aggregate_compliance_pack_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "aggregate_compliance_pack_name", value)

    @property
    @pulumi.getter(name="aggregatorId")
    def aggregator_id(self) -> pulumi.Input[str]:
        """
        The ID of aggregator.
        """
        return pulumi.get(self, "aggregator_id")

    @aggregator_id.setter
    def aggregator_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "aggregator_id", value)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        The description of compliance package.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="riskLevel")
    def risk_level(self) -> pulumi.Input[int]:
        """
        The Risk Level. Valid values: `1`, `2`, `3`.
        """
        return pulumi.get(self, "risk_level")

    @risk_level.setter
    def risk_level(self, value: pulumi.Input[int]):
        pulumi.set(self, "risk_level", value)

    @property
    @pulumi.getter(name="compliancePackTemplateId")
    def compliance_pack_template_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Template ID of compliance package.
        """
        return pulumi.get(self, "compliance_pack_template_id")

    @compliance_pack_template_id.setter
    def compliance_pack_template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compliance_pack_template_id", value)

    @property
    @pulumi.getter(name="configRuleIds")
    def config_rule_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleIdArgs']]]]:
        """
        A list of Config Rule IDs.
        """
        return pulumi.get(self, "config_rule_ids")

    @config_rule_ids.setter
    def config_rule_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleIdArgs']]]]):
        pulumi.set(self, "config_rule_ids", value)

    @property
    @pulumi.getter(name="configRules")
    def config_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleArgs']]]]:
        """
        A list of Config Rules.
        """
        return pulumi.get(self, "config_rules")

    @config_rules.setter
    def config_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleArgs']]]]):
        pulumi.set(self, "config_rules", value)


@pulumi.input_type
class _AggregateCompliancePackState:
    def __init__(__self__, *,
                 aggregate_compliance_pack_name: Optional[pulumi.Input[str]] = None,
                 aggregator_id: Optional[pulumi.Input[str]] = None,
                 compliance_pack_template_id: Optional[pulumi.Input[str]] = None,
                 config_rule_ids: Optional[pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleIdArgs']]]] = None,
                 config_rules: Optional[pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 risk_level: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AggregateCompliancePack resources.
        :param pulumi.Input[str] aggregate_compliance_pack_name: The name of compliance package name. **NOTE:** the `aggregate_compliance_pack_name` supports modification since V1.145.0.
        :param pulumi.Input[str] aggregator_id: The ID of aggregator.
        :param pulumi.Input[str] compliance_pack_template_id: The Template ID of compliance package.
        :param pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleIdArgs']]] config_rule_ids: A list of Config Rule IDs.
        :param pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleArgs']]] config_rules: A list of Config Rules.
        :param pulumi.Input[str] description: The description of compliance package.
        :param pulumi.Input[int] risk_level: The Risk Level. Valid values: `1`, `2`, `3`.
        :param pulumi.Input[str] status: The status of the resource.
        """
        if aggregate_compliance_pack_name is not None:
            pulumi.set(__self__, "aggregate_compliance_pack_name", aggregate_compliance_pack_name)
        if aggregator_id is not None:
            pulumi.set(__self__, "aggregator_id", aggregator_id)
        if compliance_pack_template_id is not None:
            pulumi.set(__self__, "compliance_pack_template_id", compliance_pack_template_id)
        if config_rule_ids is not None:
            pulumi.set(__self__, "config_rule_ids", config_rule_ids)
        if config_rules is not None:
            warnings.warn("""Field 'config_rules' has been deprecated from provider version 1.141.0. New field 'config_rule_ids' instead.""", DeprecationWarning)
            pulumi.log.warn("""config_rules is deprecated: Field 'config_rules' has been deprecated from provider version 1.141.0. New field 'config_rule_ids' instead.""")
        if config_rules is not None:
            pulumi.set(__self__, "config_rules", config_rules)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if risk_level is not None:
            pulumi.set(__self__, "risk_level", risk_level)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="aggregateCompliancePackName")
    def aggregate_compliance_pack_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of compliance package name. **NOTE:** the `aggregate_compliance_pack_name` supports modification since V1.145.0.
        """
        return pulumi.get(self, "aggregate_compliance_pack_name")

    @aggregate_compliance_pack_name.setter
    def aggregate_compliance_pack_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aggregate_compliance_pack_name", value)

    @property
    @pulumi.getter(name="aggregatorId")
    def aggregator_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of aggregator.
        """
        return pulumi.get(self, "aggregator_id")

    @aggregator_id.setter
    def aggregator_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aggregator_id", value)

    @property
    @pulumi.getter(name="compliancePackTemplateId")
    def compliance_pack_template_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Template ID of compliance package.
        """
        return pulumi.get(self, "compliance_pack_template_id")

    @compliance_pack_template_id.setter
    def compliance_pack_template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compliance_pack_template_id", value)

    @property
    @pulumi.getter(name="configRuleIds")
    def config_rule_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleIdArgs']]]]:
        """
        A list of Config Rule IDs.
        """
        return pulumi.get(self, "config_rule_ids")

    @config_rule_ids.setter
    def config_rule_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleIdArgs']]]]):
        pulumi.set(self, "config_rule_ids", value)

    @property
    @pulumi.getter(name="configRules")
    def config_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleArgs']]]]:
        """
        A list of Config Rules.
        """
        return pulumi.get(self, "config_rules")

    @config_rules.setter
    def config_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleArgs']]]]):
        pulumi.set(self, "config_rules", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of compliance package.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="riskLevel")
    def risk_level(self) -> Optional[pulumi.Input[int]]:
        """
        The Risk Level. Valid values: `1`, `2`, `3`.
        """
        return pulumi.get(self, "risk_level")

    @risk_level.setter
    def risk_level(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "risk_level", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class AggregateCompliancePack(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aggregate_compliance_pack_name: Optional[pulumi.Input[str]] = None,
                 aggregator_id: Optional[pulumi.Input[str]] = None,
                 compliance_pack_template_id: Optional[pulumi.Input[str]] = None,
                 config_rule_ids: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AggregateCompliancePackConfigRuleIdArgs']]]]] = None,
                 config_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AggregateCompliancePackConfigRuleArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 risk_level: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a Cloud Config Aggregate Compliance Pack resource.

        For information about Cloud Config Aggregate Compliance Pack and how to use it, see [What is Aggregate Compliance Pack](https://www.alibabacloud.com/help/en/doc-detail/194753.html).

        > **NOTE:** Available in v1.124.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "example_name"
        default_resource_groups = alicloud.resourcemanager.get_resource_groups(status="OK")
        default_instances = alicloud.ecs.get_instances()
        default_aggregator = alicloud.cfg.Aggregator("defaultAggregator",
            aggregator_accounts=[alicloud.cfg.AggregatorAggregatorAccountArgs(
                account_id="140278452670****",
                account_name="test-2",
                account_type="ResourceDirectory",
            )],
            aggregator_name="tf-testaccaggregator",
            description="tf-testaccaggregator")
        default_aggregate_config_rule = alicloud.cfg.AggregateConfigRule("defaultAggregateConfigRule",
            aggregator_id=default_aggregator.id,
            aggregate_config_rule_name=name,
            source_owner="ALIYUN",
            source_identifier="ecs-cpu-min-count-limit",
            config_rule_trigger_types="ConfigurationItemChangeNotification",
            resource_types_scopes=["ACS::ECS::Instance"],
            risk_level=1,
            description=name,
            exclude_resource_ids_scope=default_instances.ids[0],
            input_parameters={
                "cpuCount": "4",
            },
            region_ids_scope="cn-hangzhou",
            resource_group_ids_scope=default_resource_groups.ids[0],
            tag_key_scope="tFTest",
            tag_value_scope="forTF 123")
        default_aggregate_compliance_pack = alicloud.cfg.AggregateCompliancePack("defaultAggregateCompliancePack",
            aggregate_compliance_pack_name="tf-testaccConfig1234",
            aggregator_id=default_aggregator.id,
            description="tf-testaccConfig1234",
            risk_level=1,
            config_rule_ids=[alicloud.cfg.AggregateCompliancePackConfigRuleIdArgs(
                config_rule_id=default_aggregate_config_rule.config_rule_id,
            )])
        ```

        ## Import

        Cloud Config Aggregate Compliance Pack can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cfg/aggregateCompliancePack:AggregateCompliancePack example <aggregator_id>:<aggregator_compliance_pack_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aggregate_compliance_pack_name: The name of compliance package name. **NOTE:** the `aggregate_compliance_pack_name` supports modification since V1.145.0.
        :param pulumi.Input[str] aggregator_id: The ID of aggregator.
        :param pulumi.Input[str] compliance_pack_template_id: The Template ID of compliance package.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AggregateCompliancePackConfigRuleIdArgs']]]] config_rule_ids: A list of Config Rule IDs.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AggregateCompliancePackConfigRuleArgs']]]] config_rules: A list of Config Rules.
        :param pulumi.Input[str] description: The description of compliance package.
        :param pulumi.Input[int] risk_level: The Risk Level. Valid values: `1`, `2`, `3`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AggregateCompliancePackArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Cloud Config Aggregate Compliance Pack resource.

        For information about Cloud Config Aggregate Compliance Pack and how to use it, see [What is Aggregate Compliance Pack](https://www.alibabacloud.com/help/en/doc-detail/194753.html).

        > **NOTE:** Available in v1.124.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "example_name"
        default_resource_groups = alicloud.resourcemanager.get_resource_groups(status="OK")
        default_instances = alicloud.ecs.get_instances()
        default_aggregator = alicloud.cfg.Aggregator("defaultAggregator",
            aggregator_accounts=[alicloud.cfg.AggregatorAggregatorAccountArgs(
                account_id="140278452670****",
                account_name="test-2",
                account_type="ResourceDirectory",
            )],
            aggregator_name="tf-testaccaggregator",
            description="tf-testaccaggregator")
        default_aggregate_config_rule = alicloud.cfg.AggregateConfigRule("defaultAggregateConfigRule",
            aggregator_id=default_aggregator.id,
            aggregate_config_rule_name=name,
            source_owner="ALIYUN",
            source_identifier="ecs-cpu-min-count-limit",
            config_rule_trigger_types="ConfigurationItemChangeNotification",
            resource_types_scopes=["ACS::ECS::Instance"],
            risk_level=1,
            description=name,
            exclude_resource_ids_scope=default_instances.ids[0],
            input_parameters={
                "cpuCount": "4",
            },
            region_ids_scope="cn-hangzhou",
            resource_group_ids_scope=default_resource_groups.ids[0],
            tag_key_scope="tFTest",
            tag_value_scope="forTF 123")
        default_aggregate_compliance_pack = alicloud.cfg.AggregateCompliancePack("defaultAggregateCompliancePack",
            aggregate_compliance_pack_name="tf-testaccConfig1234",
            aggregator_id=default_aggregator.id,
            description="tf-testaccConfig1234",
            risk_level=1,
            config_rule_ids=[alicloud.cfg.AggregateCompliancePackConfigRuleIdArgs(
                config_rule_id=default_aggregate_config_rule.config_rule_id,
            )])
        ```

        ## Import

        Cloud Config Aggregate Compliance Pack can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cfg/aggregateCompliancePack:AggregateCompliancePack example <aggregator_id>:<aggregator_compliance_pack_id>
        ```

        :param str resource_name: The name of the resource.
        :param AggregateCompliancePackArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AggregateCompliancePackArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aggregate_compliance_pack_name: Optional[pulumi.Input[str]] = None,
                 aggregator_id: Optional[pulumi.Input[str]] = None,
                 compliance_pack_template_id: Optional[pulumi.Input[str]] = None,
                 config_rule_ids: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AggregateCompliancePackConfigRuleIdArgs']]]]] = None,
                 config_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AggregateCompliancePackConfigRuleArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 risk_level: Optional[pulumi.Input[int]] = None,
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
            __props__ = AggregateCompliancePackArgs.__new__(AggregateCompliancePackArgs)

            if aggregate_compliance_pack_name is None and not opts.urn:
                raise TypeError("Missing required property 'aggregate_compliance_pack_name'")
            __props__.__dict__["aggregate_compliance_pack_name"] = aggregate_compliance_pack_name
            if aggregator_id is None and not opts.urn:
                raise TypeError("Missing required property 'aggregator_id'")
            __props__.__dict__["aggregator_id"] = aggregator_id
            __props__.__dict__["compliance_pack_template_id"] = compliance_pack_template_id
            __props__.__dict__["config_rule_ids"] = config_rule_ids
            if config_rules is not None and not opts.urn:
                warnings.warn("""Field 'config_rules' has been deprecated from provider version 1.141.0. New field 'config_rule_ids' instead.""", DeprecationWarning)
                pulumi.log.warn("""config_rules is deprecated: Field 'config_rules' has been deprecated from provider version 1.141.0. New field 'config_rule_ids' instead.""")
            __props__.__dict__["config_rules"] = config_rules
            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if risk_level is None and not opts.urn:
                raise TypeError("Missing required property 'risk_level'")
            __props__.__dict__["risk_level"] = risk_level
            __props__.__dict__["status"] = None
        super(AggregateCompliancePack, __self__).__init__(
            'alicloud:cfg/aggregateCompliancePack:AggregateCompliancePack',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            aggregate_compliance_pack_name: Optional[pulumi.Input[str]] = None,
            aggregator_id: Optional[pulumi.Input[str]] = None,
            compliance_pack_template_id: Optional[pulumi.Input[str]] = None,
            config_rule_ids: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AggregateCompliancePackConfigRuleIdArgs']]]]] = None,
            config_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AggregateCompliancePackConfigRuleArgs']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            risk_level: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'AggregateCompliancePack':
        """
        Get an existing AggregateCompliancePack resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aggregate_compliance_pack_name: The name of compliance package name. **NOTE:** the `aggregate_compliance_pack_name` supports modification since V1.145.0.
        :param pulumi.Input[str] aggregator_id: The ID of aggregator.
        :param pulumi.Input[str] compliance_pack_template_id: The Template ID of compliance package.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AggregateCompliancePackConfigRuleIdArgs']]]] config_rule_ids: A list of Config Rule IDs.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AggregateCompliancePackConfigRuleArgs']]]] config_rules: A list of Config Rules.
        :param pulumi.Input[str] description: The description of compliance package.
        :param pulumi.Input[int] risk_level: The Risk Level. Valid values: `1`, `2`, `3`.
        :param pulumi.Input[str] status: The status of the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AggregateCompliancePackState.__new__(_AggregateCompliancePackState)

        __props__.__dict__["aggregate_compliance_pack_name"] = aggregate_compliance_pack_name
        __props__.__dict__["aggregator_id"] = aggregator_id
        __props__.__dict__["compliance_pack_template_id"] = compliance_pack_template_id
        __props__.__dict__["config_rule_ids"] = config_rule_ids
        __props__.__dict__["config_rules"] = config_rules
        __props__.__dict__["description"] = description
        __props__.__dict__["risk_level"] = risk_level
        __props__.__dict__["status"] = status
        return AggregateCompliancePack(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aggregateCompliancePackName")
    def aggregate_compliance_pack_name(self) -> pulumi.Output[str]:
        """
        The name of compliance package name. **NOTE:** the `aggregate_compliance_pack_name` supports modification since V1.145.0.
        """
        return pulumi.get(self, "aggregate_compliance_pack_name")

    @property
    @pulumi.getter(name="aggregatorId")
    def aggregator_id(self) -> pulumi.Output[str]:
        """
        The ID of aggregator.
        """
        return pulumi.get(self, "aggregator_id")

    @property
    @pulumi.getter(name="compliancePackTemplateId")
    def compliance_pack_template_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Template ID of compliance package.
        """
        return pulumi.get(self, "compliance_pack_template_id")

    @property
    @pulumi.getter(name="configRuleIds")
    def config_rule_ids(self) -> pulumi.Output[Optional[Sequence['outputs.AggregateCompliancePackConfigRuleId']]]:
        """
        A list of Config Rule IDs.
        """
        return pulumi.get(self, "config_rule_ids")

    @property
    @pulumi.getter(name="configRules")
    def config_rules(self) -> pulumi.Output[Optional[Sequence['outputs.AggregateCompliancePackConfigRule']]]:
        """
        A list of Config Rules.
        """
        return pulumi.get(self, "config_rules")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of compliance package.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="riskLevel")
    def risk_level(self) -> pulumi.Output[int]:
        """
        The Risk Level. Valid values: `1`, `2`, `3`.
        """
        return pulumi.get(self, "risk_level")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

