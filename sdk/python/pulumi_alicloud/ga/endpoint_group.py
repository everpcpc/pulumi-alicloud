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

__all__ = ['EndpointGroup']


class EndpointGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accelerator_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupEndpointConfigurationArgs']]]]] = None,
                 endpoint_group_region: Optional[pulumi.Input[str]] = None,
                 endpoint_group_type: Optional[pulumi.Input[str]] = None,
                 endpoint_request_protocol: Optional[pulumi.Input[str]] = None,
                 health_check_interval_seconds: Optional[pulumi.Input[int]] = None,
                 health_check_path: Optional[pulumi.Input[str]] = None,
                 health_check_port: Optional[pulumi.Input[int]] = None,
                 health_check_protocol: Optional[pulumi.Input[str]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port_overrides: Optional[pulumi.Input[pulumi.InputType['EndpointGroupPortOverridesArgs']]] = None,
                 threshold_count: Optional[pulumi.Input[int]] = None,
                 traffic_percentage: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Global Accelerator (GA) Endpoint Group resource.

        For information about Global Accelerator (GA) Endpoint Group and how to use it, see [What is Endpoint Group](https://www.alibabacloud.com/help/en/doc-detail/153259.htm).

        > **NOTE:** Available in v1.113.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_accelerator = alicloud.ga.Accelerator("exampleAccelerator",
            duration=1,
            auto_use_coupon=True,
            spec="1")
        example_listener = alicloud.ga.Listener("exampleListener",
            accelerator_id="alicloud_ga_accelerator.example.id",
            port_ranges=[alicloud.ga.ListenerPortRangeArgs(
                from_port=60,
                to_port=70,
            )])
        example_eip = alicloud.ecs.Eip("exampleEip",
            bandwidth=10,
            internet_charge_type="PayByBandwidth")
        example_endpoint_group = alicloud.ga.EndpointGroup("exampleEndpointGroup",
            accelerator_id=alicloud_ga_accelerators["example"]["id"],
            endpoint_configurations=[alicloud.ga.EndpointGroupEndpointConfigurationArgs(
                endpoint=example_eip.ip_address,
                type="PublicIp",
                weight=20,
            )],
            endpoint_group_region="cn-hangzhou",
            listener_id=example_listener.id)
        ```

        ## Import

        Ga Endpoint Group can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ga/endpointGroup:EndpointGroup example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accelerator_id: The ID of the Global Accelerator instance to which the endpoint group will be added.
        :param pulumi.Input[str] description: The description of the endpoint group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupEndpointConfigurationArgs']]]] endpoint_configurations: The endpointConfigurations of the endpoint group.
        :param pulumi.Input[str] endpoint_group_region: The ID of the region where the endpoint group is deployed.
        :param pulumi.Input[str] endpoint_group_type: The endpoint group type. Valid values: `default`, `virtual`. Default value is `default`.
        :param pulumi.Input[str] endpoint_request_protocol: The endpoint request protocol.
        :param pulumi.Input[int] health_check_interval_seconds: The interval between two consecutive health checks. Unit: seconds.
        :param pulumi.Input[str] health_check_path: The path specified as the destination of the targets for health checks.
        :param pulumi.Input[int] health_check_port: The port that is used for health checks.
        :param pulumi.Input[str] health_check_protocol: The protocol that is used to connect to the targets for health checks. Valid values: `http`, `https`, `tcp`.
        :param pulumi.Input[str] listener_id: The ID of the listener that is associated with the endpoint group.
        :param pulumi.Input[str] name: The name of the endpoint group.
        :param pulumi.Input[pulumi.InputType['EndpointGroupPortOverridesArgs']] port_overrides: Mapping between listening port and forwarding port of boarding point.
        :param pulumi.Input[int] threshold_count: The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value is `3`.
        :param pulumi.Input[int] traffic_percentage: The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
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

            if accelerator_id is None:
                raise TypeError("Missing required property 'accelerator_id'")
            __props__['accelerator_id'] = accelerator_id
            __props__['description'] = description
            if endpoint_configurations is None:
                raise TypeError("Missing required property 'endpoint_configurations'")
            __props__['endpoint_configurations'] = endpoint_configurations
            if endpoint_group_region is None:
                raise TypeError("Missing required property 'endpoint_group_region'")
            __props__['endpoint_group_region'] = endpoint_group_region
            __props__['endpoint_group_type'] = endpoint_group_type
            __props__['endpoint_request_protocol'] = endpoint_request_protocol
            __props__['health_check_interval_seconds'] = health_check_interval_seconds
            __props__['health_check_path'] = health_check_path
            __props__['health_check_port'] = health_check_port
            __props__['health_check_protocol'] = health_check_protocol
            if listener_id is None:
                raise TypeError("Missing required property 'listener_id'")
            __props__['listener_id'] = listener_id
            __props__['name'] = name
            __props__['port_overrides'] = port_overrides
            __props__['threshold_count'] = threshold_count
            __props__['traffic_percentage'] = traffic_percentage
            __props__['status'] = None
        super(EndpointGroup, __self__).__init__(
            'alicloud:ga/endpointGroup:EndpointGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accelerator_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            endpoint_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupEndpointConfigurationArgs']]]]] = None,
            endpoint_group_region: Optional[pulumi.Input[str]] = None,
            endpoint_group_type: Optional[pulumi.Input[str]] = None,
            endpoint_request_protocol: Optional[pulumi.Input[str]] = None,
            health_check_interval_seconds: Optional[pulumi.Input[int]] = None,
            health_check_path: Optional[pulumi.Input[str]] = None,
            health_check_port: Optional[pulumi.Input[int]] = None,
            health_check_protocol: Optional[pulumi.Input[str]] = None,
            listener_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port_overrides: Optional[pulumi.Input[pulumi.InputType['EndpointGroupPortOverridesArgs']]] = None,
            status: Optional[pulumi.Input[str]] = None,
            threshold_count: Optional[pulumi.Input[int]] = None,
            traffic_percentage: Optional[pulumi.Input[int]] = None) -> 'EndpointGroup':
        """
        Get an existing EndpointGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accelerator_id: The ID of the Global Accelerator instance to which the endpoint group will be added.
        :param pulumi.Input[str] description: The description of the endpoint group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupEndpointConfigurationArgs']]]] endpoint_configurations: The endpointConfigurations of the endpoint group.
        :param pulumi.Input[str] endpoint_group_region: The ID of the region where the endpoint group is deployed.
        :param pulumi.Input[str] endpoint_group_type: The endpoint group type. Valid values: `default`, `virtual`. Default value is `default`.
        :param pulumi.Input[str] endpoint_request_protocol: The endpoint request protocol.
        :param pulumi.Input[int] health_check_interval_seconds: The interval between two consecutive health checks. Unit: seconds.
        :param pulumi.Input[str] health_check_path: The path specified as the destination of the targets for health checks.
        :param pulumi.Input[int] health_check_port: The port that is used for health checks.
        :param pulumi.Input[str] health_check_protocol: The protocol that is used to connect to the targets for health checks. Valid values: `http`, `https`, `tcp`.
        :param pulumi.Input[str] listener_id: The ID of the listener that is associated with the endpoint group.
        :param pulumi.Input[str] name: The name of the endpoint group.
        :param pulumi.Input[pulumi.InputType['EndpointGroupPortOverridesArgs']] port_overrides: Mapping between listening port and forwarding port of boarding point.
        :param pulumi.Input[str] status: The status of the endpoint group.
        :param pulumi.Input[int] threshold_count: The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value is `3`.
        :param pulumi.Input[int] traffic_percentage: The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accelerator_id"] = accelerator_id
        __props__["description"] = description
        __props__["endpoint_configurations"] = endpoint_configurations
        __props__["endpoint_group_region"] = endpoint_group_region
        __props__["endpoint_group_type"] = endpoint_group_type
        __props__["endpoint_request_protocol"] = endpoint_request_protocol
        __props__["health_check_interval_seconds"] = health_check_interval_seconds
        __props__["health_check_path"] = health_check_path
        __props__["health_check_port"] = health_check_port
        __props__["health_check_protocol"] = health_check_protocol
        __props__["listener_id"] = listener_id
        __props__["name"] = name
        __props__["port_overrides"] = port_overrides
        __props__["status"] = status
        __props__["threshold_count"] = threshold_count
        __props__["traffic_percentage"] = traffic_percentage
        return EndpointGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceleratorId")
    def accelerator_id(self) -> pulumi.Output[str]:
        """
        The ID of the Global Accelerator instance to which the endpoint group will be added.
        """
        return pulumi.get(self, "accelerator_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the endpoint group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endpointConfigurations")
    def endpoint_configurations(self) -> pulumi.Output[Sequence['outputs.EndpointGroupEndpointConfiguration']]:
        """
        The endpointConfigurations of the endpoint group.
        """
        return pulumi.get(self, "endpoint_configurations")

    @property
    @pulumi.getter(name="endpointGroupRegion")
    def endpoint_group_region(self) -> pulumi.Output[str]:
        """
        The ID of the region where the endpoint group is deployed.
        """
        return pulumi.get(self, "endpoint_group_region")

    @property
    @pulumi.getter(name="endpointGroupType")
    def endpoint_group_type(self) -> pulumi.Output[Optional[str]]:
        """
        The endpoint group type. Valid values: `default`, `virtual`. Default value is `default`.
        """
        return pulumi.get(self, "endpoint_group_type")

    @property
    @pulumi.getter(name="endpointRequestProtocol")
    def endpoint_request_protocol(self) -> pulumi.Output[Optional[str]]:
        """
        The endpoint request protocol.
        """
        return pulumi.get(self, "endpoint_request_protocol")

    @property
    @pulumi.getter(name="healthCheckIntervalSeconds")
    def health_check_interval_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        The interval between two consecutive health checks. Unit: seconds.
        """
        return pulumi.get(self, "health_check_interval_seconds")

    @property
    @pulumi.getter(name="healthCheckPath")
    def health_check_path(self) -> pulumi.Output[Optional[str]]:
        """
        The path specified as the destination of the targets for health checks.
        """
        return pulumi.get(self, "health_check_path")

    @property
    @pulumi.getter(name="healthCheckPort")
    def health_check_port(self) -> pulumi.Output[Optional[int]]:
        """
        The port that is used for health checks.
        """
        return pulumi.get(self, "health_check_port")

    @property
    @pulumi.getter(name="healthCheckProtocol")
    def health_check_protocol(self) -> pulumi.Output[Optional[str]]:
        """
        The protocol that is used to connect to the targets for health checks. Valid values: `http`, `https`, `tcp`.
        """
        return pulumi.get(self, "health_check_protocol")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> pulumi.Output[str]:
        """
        The ID of the listener that is associated with the endpoint group.
        """
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the endpoint group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="portOverrides")
    def port_overrides(self) -> pulumi.Output[Optional['outputs.EndpointGroupPortOverrides']]:
        """
        Mapping between listening port and forwarding port of boarding point.
        """
        return pulumi.get(self, "port_overrides")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the endpoint group.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="thresholdCount")
    def threshold_count(self) -> pulumi.Output[Optional[int]]:
        """
        The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value is `3`.
        """
        return pulumi.get(self, "threshold_count")

    @property
    @pulumi.getter(name="trafficPercentage")
    def traffic_percentage(self) -> pulumi.Output[Optional[int]]:
        """
        The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
        """
        return pulumi.get(self, "traffic_percentage")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

