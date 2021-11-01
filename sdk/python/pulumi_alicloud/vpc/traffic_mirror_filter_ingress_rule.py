# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TrafficMirrorFilterIngressRuleArgs', 'TrafficMirrorFilterIngressRule']

@pulumi.input_type
class TrafficMirrorFilterIngressRuleArgs:
    def __init__(__self__, *,
                 destination_cidr_block: pulumi.Input[str],
                 priority: pulumi.Input[int],
                 protocol: pulumi.Input[str],
                 rule_action: pulumi.Input[str],
                 source_cidr_block: pulumi.Input[str],
                 traffic_mirror_filter_id: pulumi.Input[str],
                 destination_port_range: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 source_port_range: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TrafficMirrorFilterIngressRule resource.
        :param pulumi.Input[str] destination_cidr_block: The destination CIDR block of the inbound traffic.
        :param pulumi.Input[int] priority: The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
        :param pulumi.Input[str] protocol: The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
        :param pulumi.Input[str] rule_action: The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
        :param pulumi.Input[str] source_cidr_block: The source CIDR block of the inbound traffic.
        :param pulumi.Input[str] traffic_mirror_filter_id: The ID of the filter.
        :param pulumi.Input[str] destination_port_range: The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        :param pulumi.Input[bool] dry_run: Whether to pre-check this request only. Default to: `false`
        :param pulumi.Input[str] source_port_range: The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        """
        pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "rule_action", rule_action)
        pulumi.set(__self__, "source_cidr_block", source_cidr_block)
        pulumi.set(__self__, "traffic_mirror_filter_id", traffic_mirror_filter_id)
        if destination_port_range is not None:
            pulumi.set(__self__, "destination_port_range", destination_port_range)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if source_port_range is not None:
            pulumi.set(__self__, "source_port_range", source_port_range)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Input[str]:
        """
        The destination CIDR block of the inbound traffic.
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[int]:
        """
        The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: pulumi.Input[int]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="ruleAction")
    def rule_action(self) -> pulumi.Input[str]:
        """
        The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
        """
        return pulumi.get(self, "rule_action")

    @rule_action.setter
    def rule_action(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_action", value)

    @property
    @pulumi.getter(name="sourceCidrBlock")
    def source_cidr_block(self) -> pulumi.Input[str]:
        """
        The source CIDR block of the inbound traffic.
        """
        return pulumi.get(self, "source_cidr_block")

    @source_cidr_block.setter
    def source_cidr_block(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_cidr_block", value)

    @property
    @pulumi.getter(name="trafficMirrorFilterId")
    def traffic_mirror_filter_id(self) -> pulumi.Input[str]:
        """
        The ID of the filter.
        """
        return pulumi.get(self, "traffic_mirror_filter_id")

    @traffic_mirror_filter_id.setter
    def traffic_mirror_filter_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "traffic_mirror_filter_id", value)

    @property
    @pulumi.getter(name="destinationPortRange")
    def destination_port_range(self) -> Optional[pulumi.Input[str]]:
        """
        The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        """
        return pulumi.get(self, "destination_port_range")

    @destination_port_range.setter
    def destination_port_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_port_range", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to pre-check this request only. Default to: `false`
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter(name="sourcePortRange")
    def source_port_range(self) -> Optional[pulumi.Input[str]]:
        """
        The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        """
        return pulumi.get(self, "source_port_range")

    @source_port_range.setter
    def source_port_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_port_range", value)


@pulumi.input_type
class _TrafficMirrorFilterIngressRuleState:
    def __init__(__self__, *,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 destination_port_range: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 rule_action: Optional[pulumi.Input[str]] = None,
                 source_cidr_block: Optional[pulumi.Input[str]] = None,
                 source_port_range: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 traffic_mirror_filter_id: Optional[pulumi.Input[str]] = None,
                 traffic_mirror_filter_ingress_rule_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TrafficMirrorFilterIngressRule resources.
        :param pulumi.Input[str] destination_cidr_block: The destination CIDR block of the inbound traffic.
        :param pulumi.Input[str] destination_port_range: The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        :param pulumi.Input[bool] dry_run: Whether to pre-check this request only. Default to: `false`
        :param pulumi.Input[int] priority: The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
        :param pulumi.Input[str] protocol: The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
        :param pulumi.Input[str] rule_action: The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
        :param pulumi.Input[str] source_cidr_block: The source CIDR block of the inbound traffic.
        :param pulumi.Input[str] source_port_range: The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        :param pulumi.Input[str] status: The state of the inbound rule. Valid values:`Creating`, `Created`, `Modifying` and `Deleting`.
        :param pulumi.Input[str] traffic_mirror_filter_id: The ID of the filter.
        :param pulumi.Input[str] traffic_mirror_filter_ingress_rule_id: The ID of the inbound rule.
        """
        if destination_cidr_block is not None:
            pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        if destination_port_range is not None:
            pulumi.set(__self__, "destination_port_range", destination_port_range)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if rule_action is not None:
            pulumi.set(__self__, "rule_action", rule_action)
        if source_cidr_block is not None:
            pulumi.set(__self__, "source_cidr_block", source_cidr_block)
        if source_port_range is not None:
            pulumi.set(__self__, "source_port_range", source_port_range)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if traffic_mirror_filter_id is not None:
            pulumi.set(__self__, "traffic_mirror_filter_id", traffic_mirror_filter_id)
        if traffic_mirror_filter_ingress_rule_id is not None:
            pulumi.set(__self__, "traffic_mirror_filter_ingress_rule_id", traffic_mirror_filter_ingress_rule_id)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The destination CIDR block of the inbound traffic.
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter(name="destinationPortRange")
    def destination_port_range(self) -> Optional[pulumi.Input[str]]:
        """
        The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        """
        return pulumi.get(self, "destination_port_range")

    @destination_port_range.setter
    def destination_port_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_port_range", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to pre-check this request only. Default to: `false`
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="ruleAction")
    def rule_action(self) -> Optional[pulumi.Input[str]]:
        """
        The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
        """
        return pulumi.get(self, "rule_action")

    @rule_action.setter
    def rule_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_action", value)

    @property
    @pulumi.getter(name="sourceCidrBlock")
    def source_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The source CIDR block of the inbound traffic.
        """
        return pulumi.get(self, "source_cidr_block")

    @source_cidr_block.setter
    def source_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_cidr_block", value)

    @property
    @pulumi.getter(name="sourcePortRange")
    def source_port_range(self) -> Optional[pulumi.Input[str]]:
        """
        The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        """
        return pulumi.get(self, "source_port_range")

    @source_port_range.setter
    def source_port_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_port_range", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The state of the inbound rule. Valid values:`Creating`, `Created`, `Modifying` and `Deleting`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="trafficMirrorFilterId")
    def traffic_mirror_filter_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the filter.
        """
        return pulumi.get(self, "traffic_mirror_filter_id")

    @traffic_mirror_filter_id.setter
    def traffic_mirror_filter_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_mirror_filter_id", value)

    @property
    @pulumi.getter(name="trafficMirrorFilterIngressRuleId")
    def traffic_mirror_filter_ingress_rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the inbound rule.
        """
        return pulumi.get(self, "traffic_mirror_filter_ingress_rule_id")

    @traffic_mirror_filter_ingress_rule_id.setter
    def traffic_mirror_filter_ingress_rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_mirror_filter_ingress_rule_id", value)


class TrafficMirrorFilterIngressRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 destination_port_range: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 rule_action: Optional[pulumi.Input[str]] = None,
                 source_cidr_block: Optional[pulumi.Input[str]] = None,
                 source_port_range: Optional[pulumi.Input[str]] = None,
                 traffic_mirror_filter_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a VPC Traffic Mirror Filter Ingress Rule resource.

        For information about VPC Traffic Mirror Filter Ingress Rule and how to use it, see [What is Traffic Mirror Filter Ingress Rule](https://www.alibabacloud.com/help/doc-detail/261357.htm).

        > **NOTE:** Available in v1.141.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_traffic_mirror_filter = alicloud.vpc.TrafficMirrorFilter("exampleTrafficMirrorFilter", traffic_mirror_filter_name="example_value")
        example_traffic_mirror_filter_ingress_rule = alicloud.vpc.TrafficMirrorFilterIngressRule("exampleTrafficMirrorFilterIngressRule",
            traffic_mirror_filter_id=example_traffic_mirror_filter.id,
            priority=1,
            rule_action="accept",
            protocol="UDP",
            destination_cidr_block="10.0.0.0/24",
            source_cidr_block="10.0.0.0/24",
            destination_port_range="1/120",
            source_port_range="1/120")
        ```

        ## Import

        VPC Traffic Mirror Filter Ingress Rule can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/trafficMirrorFilterIngressRule:TrafficMirrorFilterIngressRule example <traffic_mirror_filter_id>:<traffic_mirror_filter_ingress_rule_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidr_block: The destination CIDR block of the inbound traffic.
        :param pulumi.Input[str] destination_port_range: The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        :param pulumi.Input[bool] dry_run: Whether to pre-check this request only. Default to: `false`
        :param pulumi.Input[int] priority: The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
        :param pulumi.Input[str] protocol: The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
        :param pulumi.Input[str] rule_action: The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
        :param pulumi.Input[str] source_cidr_block: The source CIDR block of the inbound traffic.
        :param pulumi.Input[str] source_port_range: The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        :param pulumi.Input[str] traffic_mirror_filter_id: The ID of the filter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TrafficMirrorFilterIngressRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a VPC Traffic Mirror Filter Ingress Rule resource.

        For information about VPC Traffic Mirror Filter Ingress Rule and how to use it, see [What is Traffic Mirror Filter Ingress Rule](https://www.alibabacloud.com/help/doc-detail/261357.htm).

        > **NOTE:** Available in v1.141.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_traffic_mirror_filter = alicloud.vpc.TrafficMirrorFilter("exampleTrafficMirrorFilter", traffic_mirror_filter_name="example_value")
        example_traffic_mirror_filter_ingress_rule = alicloud.vpc.TrafficMirrorFilterIngressRule("exampleTrafficMirrorFilterIngressRule",
            traffic_mirror_filter_id=example_traffic_mirror_filter.id,
            priority=1,
            rule_action="accept",
            protocol="UDP",
            destination_cidr_block="10.0.0.0/24",
            source_cidr_block="10.0.0.0/24",
            destination_port_range="1/120",
            source_port_range="1/120")
        ```

        ## Import

        VPC Traffic Mirror Filter Ingress Rule can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/trafficMirrorFilterIngressRule:TrafficMirrorFilterIngressRule example <traffic_mirror_filter_id>:<traffic_mirror_filter_ingress_rule_id>
        ```

        :param str resource_name: The name of the resource.
        :param TrafficMirrorFilterIngressRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrafficMirrorFilterIngressRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 destination_port_range: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 rule_action: Optional[pulumi.Input[str]] = None,
                 source_cidr_block: Optional[pulumi.Input[str]] = None,
                 source_port_range: Optional[pulumi.Input[str]] = None,
                 traffic_mirror_filter_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = TrafficMirrorFilterIngressRuleArgs.__new__(TrafficMirrorFilterIngressRuleArgs)

            if destination_cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'destination_cidr_block'")
            __props__.__dict__["destination_cidr_block"] = destination_cidr_block
            __props__.__dict__["destination_port_range"] = destination_port_range
            __props__.__dict__["dry_run"] = dry_run
            if priority is None and not opts.urn:
                raise TypeError("Missing required property 'priority'")
            __props__.__dict__["priority"] = priority
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            if rule_action is None and not opts.urn:
                raise TypeError("Missing required property 'rule_action'")
            __props__.__dict__["rule_action"] = rule_action
            if source_cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'source_cidr_block'")
            __props__.__dict__["source_cidr_block"] = source_cidr_block
            __props__.__dict__["source_port_range"] = source_port_range
            if traffic_mirror_filter_id is None and not opts.urn:
                raise TypeError("Missing required property 'traffic_mirror_filter_id'")
            __props__.__dict__["traffic_mirror_filter_id"] = traffic_mirror_filter_id
            __props__.__dict__["status"] = None
            __props__.__dict__["traffic_mirror_filter_ingress_rule_id"] = None
        super(TrafficMirrorFilterIngressRule, __self__).__init__(
            'alicloud:vpc/trafficMirrorFilterIngressRule:TrafficMirrorFilterIngressRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            destination_cidr_block: Optional[pulumi.Input[str]] = None,
            destination_port_range: Optional[pulumi.Input[str]] = None,
            dry_run: Optional[pulumi.Input[bool]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            rule_action: Optional[pulumi.Input[str]] = None,
            source_cidr_block: Optional[pulumi.Input[str]] = None,
            source_port_range: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            traffic_mirror_filter_id: Optional[pulumi.Input[str]] = None,
            traffic_mirror_filter_ingress_rule_id: Optional[pulumi.Input[str]] = None) -> 'TrafficMirrorFilterIngressRule':
        """
        Get an existing TrafficMirrorFilterIngressRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidr_block: The destination CIDR block of the inbound traffic.
        :param pulumi.Input[str] destination_port_range: The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        :param pulumi.Input[bool] dry_run: Whether to pre-check this request only. Default to: `false`
        :param pulumi.Input[int] priority: The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
        :param pulumi.Input[str] protocol: The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
        :param pulumi.Input[str] rule_action: The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
        :param pulumi.Input[str] source_cidr_block: The source CIDR block of the inbound traffic.
        :param pulumi.Input[str] source_port_range: The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        :param pulumi.Input[str] status: The state of the inbound rule. Valid values:`Creating`, `Created`, `Modifying` and `Deleting`.
        :param pulumi.Input[str] traffic_mirror_filter_id: The ID of the filter.
        :param pulumi.Input[str] traffic_mirror_filter_ingress_rule_id: The ID of the inbound rule.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TrafficMirrorFilterIngressRuleState.__new__(_TrafficMirrorFilterIngressRuleState)

        __props__.__dict__["destination_cidr_block"] = destination_cidr_block
        __props__.__dict__["destination_port_range"] = destination_port_range
        __props__.__dict__["dry_run"] = dry_run
        __props__.__dict__["priority"] = priority
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["rule_action"] = rule_action
        __props__.__dict__["source_cidr_block"] = source_cidr_block
        __props__.__dict__["source_port_range"] = source_port_range
        __props__.__dict__["status"] = status
        __props__.__dict__["traffic_mirror_filter_id"] = traffic_mirror_filter_id
        __props__.__dict__["traffic_mirror_filter_ingress_rule_id"] = traffic_mirror_filter_ingress_rule_id
        return TrafficMirrorFilterIngressRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Output[str]:
        """
        The destination CIDR block of the inbound traffic.
        """
        return pulumi.get(self, "destination_cidr_block")

    @property
    @pulumi.getter(name="destinationPortRange")
    def destination_port_range(self) -> pulumi.Output[str]:
        """
        The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        """
        return pulumi.get(self, "destination_port_range")

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to pre-check this request only. Default to: `false`
        """
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[int]:
        """
        The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="ruleAction")
    def rule_action(self) -> pulumi.Output[str]:
        """
        The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
        """
        return pulumi.get(self, "rule_action")

    @property
    @pulumi.getter(name="sourceCidrBlock")
    def source_cidr_block(self) -> pulumi.Output[str]:
        """
        The source CIDR block of the inbound traffic.
        """
        return pulumi.get(self, "source_cidr_block")

    @property
    @pulumi.getter(name="sourcePortRange")
    def source_port_range(self) -> pulumi.Output[str]:
        """
        The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
        """
        return pulumi.get(self, "source_port_range")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The state of the inbound rule. Valid values:`Creating`, `Created`, `Modifying` and `Deleting`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="trafficMirrorFilterId")
    def traffic_mirror_filter_id(self) -> pulumi.Output[str]:
        """
        The ID of the filter.
        """
        return pulumi.get(self, "traffic_mirror_filter_id")

    @property
    @pulumi.getter(name="trafficMirrorFilterIngressRuleId")
    def traffic_mirror_filter_ingress_rule_id(self) -> pulumi.Output[str]:
        """
        The ID of the inbound rule.
        """
        return pulumi.get(self, "traffic_mirror_filter_ingress_rule_id")

