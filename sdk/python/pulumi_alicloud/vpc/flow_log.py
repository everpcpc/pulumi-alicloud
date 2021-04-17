# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FlowLogArgs', 'FlowLog']

@pulumi.input_type
class FlowLogArgs:
    def __init__(__self__, *,
                 log_store_name: pulumi.Input[str],
                 project_name: pulumi.Input[str],
                 resource_id: pulumi.Input[str],
                 resource_type: pulumi.Input[str],
                 traffic_type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 flow_log_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FlowLog resource.
        :param pulumi.Input[str] log_store_name: The name of the logstore.
        :param pulumi.Input[str] project_name: The name of the project.
        :param pulumi.Input[str] resource_id: The ID of the resource.
        :param pulumi.Input[str] resource_type: The type of the resource to capture traffic.
        :param pulumi.Input[str] traffic_type: The type of traffic collected.
        :param pulumi.Input[str] description: The Description of the VPC Flow Log.
        :param pulumi.Input[str] flow_log_name: The Name of the VPC Flow Log.
        :param pulumi.Input[str] status: The status of the VPC Flow Log.
        """
        pulumi.set(__self__, "log_store_name", log_store_name)
        pulumi.set(__self__, "project_name", project_name)
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "resource_type", resource_type)
        pulumi.set(__self__, "traffic_type", traffic_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if flow_log_name is not None:
            pulumi.set(__self__, "flow_log_name", flow_log_name)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="logStoreName")
    def log_store_name(self) -> pulumi.Input[str]:
        """
        The name of the logstore.
        """
        return pulumi.get(self, "log_store_name")

    @log_store_name.setter
    def log_store_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_store_name", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Input[str]:
        """
        The name of the project.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The ID of the resource.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Input[str]:
        """
        The type of the resource to capture traffic.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter(name="trafficType")
    def traffic_type(self) -> pulumi.Input[str]:
        """
        The type of traffic collected.
        """
        return pulumi.get(self, "traffic_type")

    @traffic_type.setter
    def traffic_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "traffic_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The Description of the VPC Flow Log.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="flowLogName")
    def flow_log_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name of the VPC Flow Log.
        """
        return pulumi.get(self, "flow_log_name")

    @flow_log_name.setter
    def flow_log_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flow_log_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the VPC Flow Log.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _FlowLogState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 flow_log_name: Optional[pulumi.Input[str]] = None,
                 log_store_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 traffic_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FlowLog resources.
        :param pulumi.Input[str] description: The Description of the VPC Flow Log.
        :param pulumi.Input[str] flow_log_name: The Name of the VPC Flow Log.
        :param pulumi.Input[str] log_store_name: The name of the logstore.
        :param pulumi.Input[str] project_name: The name of the project.
        :param pulumi.Input[str] resource_id: The ID of the resource.
        :param pulumi.Input[str] resource_type: The type of the resource to capture traffic.
        :param pulumi.Input[str] status: The status of the VPC Flow Log.
        :param pulumi.Input[str] traffic_type: The type of traffic collected.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if flow_log_name is not None:
            pulumi.set(__self__, "flow_log_name", flow_log_name)
        if log_store_name is not None:
            pulumi.set(__self__, "log_store_name", log_store_name)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if traffic_type is not None:
            pulumi.set(__self__, "traffic_type", traffic_type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The Description of the VPC Flow Log.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="flowLogName")
    def flow_log_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name of the VPC Flow Log.
        """
        return pulumi.get(self, "flow_log_name")

    @flow_log_name.setter
    def flow_log_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flow_log_name", value)

    @property
    @pulumi.getter(name="logStoreName")
    def log_store_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the logstore.
        """
        return pulumi.get(self, "log_store_name")

    @log_store_name.setter
    def log_store_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_store_name", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the project.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the resource to capture traffic.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the VPC Flow Log.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="trafficType")
    def traffic_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of traffic collected.
        """
        return pulumi.get(self, "traffic_type")

    @traffic_type.setter
    def traffic_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_type", value)


class FlowLog(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 flow_log_name: Optional[pulumi.Input[str]] = None,
                 log_store_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 traffic_type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a VPC Flow Log resource.

        For information about VPC Flow log and how to use it, see [Flow log overview](https://www.alibabacloud.com/help/doc-detail/127150.htm).

        > **NOTE:** Available in v1.117.0+

        > **NOTE:** While it uses `vpc.FlowLog` to build a vpc flow log resource, it will be active by default.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terratest_vpc_flow_log"
        log_store_name = config.get("logStoreName")
        if log_store_name is None:
            log_store_name = "vpc-flow-log-for-vpc"
        project_name = config.get("projectName")
        if project_name is None:
            project_name = "vpc-flow-log-for-vpc"
        default_network = alicloud.vpc.Network("defaultNetwork", cidr_block="192.168.0.0/24")
        default_flow_log = alicloud.vpc.FlowLog("defaultFlowLog",
            resource_id=default_network.id,
            resource_type="VPC",
            traffic_type="All",
            log_store_name=log_store_name,
            project_name=project_name,
            flow_log_name=name,
            status="Active",
            opts=pulumi.ResourceOptions(depends_on=["alicloud_vpc.default"]))
        ```

        ## Import

        VPC Flow Log can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/flowLog:FlowLog example fl-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The Description of the VPC Flow Log.
        :param pulumi.Input[str] flow_log_name: The Name of the VPC Flow Log.
        :param pulumi.Input[str] log_store_name: The name of the logstore.
        :param pulumi.Input[str] project_name: The name of the project.
        :param pulumi.Input[str] resource_id: The ID of the resource.
        :param pulumi.Input[str] resource_type: The type of the resource to capture traffic.
        :param pulumi.Input[str] status: The status of the VPC Flow Log.
        :param pulumi.Input[str] traffic_type: The type of traffic collected.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FlowLogArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a VPC Flow Log resource.

        For information about VPC Flow log and how to use it, see [Flow log overview](https://www.alibabacloud.com/help/doc-detail/127150.htm).

        > **NOTE:** Available in v1.117.0+

        > **NOTE:** While it uses `vpc.FlowLog` to build a vpc flow log resource, it will be active by default.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terratest_vpc_flow_log"
        log_store_name = config.get("logStoreName")
        if log_store_name is None:
            log_store_name = "vpc-flow-log-for-vpc"
        project_name = config.get("projectName")
        if project_name is None:
            project_name = "vpc-flow-log-for-vpc"
        default_network = alicloud.vpc.Network("defaultNetwork", cidr_block="192.168.0.0/24")
        default_flow_log = alicloud.vpc.FlowLog("defaultFlowLog",
            resource_id=default_network.id,
            resource_type="VPC",
            traffic_type="All",
            log_store_name=log_store_name,
            project_name=project_name,
            flow_log_name=name,
            status="Active",
            opts=pulumi.ResourceOptions(depends_on=["alicloud_vpc.default"]))
        ```

        ## Import

        VPC Flow Log can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/flowLog:FlowLog example fl-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param FlowLogArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FlowLogArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 flow_log_name: Optional[pulumi.Input[str]] = None,
                 log_store_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 traffic_type: Optional[pulumi.Input[str]] = None,
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
            __props__ = FlowLogArgs.__new__(FlowLogArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["flow_log_name"] = flow_log_name
            if log_store_name is None and not opts.urn:
                raise TypeError("Missing required property 'log_store_name'")
            __props__.__dict__["log_store_name"] = log_store_name
            if project_name is None and not opts.urn:
                raise TypeError("Missing required property 'project_name'")
            __props__.__dict__["project_name"] = project_name
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
            if resource_type is None and not opts.urn:
                raise TypeError("Missing required property 'resource_type'")
            __props__.__dict__["resource_type"] = resource_type
            __props__.__dict__["status"] = status
            if traffic_type is None and not opts.urn:
                raise TypeError("Missing required property 'traffic_type'")
            __props__.__dict__["traffic_type"] = traffic_type
        super(FlowLog, __self__).__init__(
            'alicloud:vpc/flowLog:FlowLog',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            flow_log_name: Optional[pulumi.Input[str]] = None,
            log_store_name: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            resource_type: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            traffic_type: Optional[pulumi.Input[str]] = None) -> 'FlowLog':
        """
        Get an existing FlowLog resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The Description of the VPC Flow Log.
        :param pulumi.Input[str] flow_log_name: The Name of the VPC Flow Log.
        :param pulumi.Input[str] log_store_name: The name of the logstore.
        :param pulumi.Input[str] project_name: The name of the project.
        :param pulumi.Input[str] resource_id: The ID of the resource.
        :param pulumi.Input[str] resource_type: The type of the resource to capture traffic.
        :param pulumi.Input[str] status: The status of the VPC Flow Log.
        :param pulumi.Input[str] traffic_type: The type of traffic collected.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FlowLogState.__new__(_FlowLogState)

        __props__.__dict__["description"] = description
        __props__.__dict__["flow_log_name"] = flow_log_name
        __props__.__dict__["log_store_name"] = log_store_name
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["resource_type"] = resource_type
        __props__.__dict__["status"] = status
        __props__.__dict__["traffic_type"] = traffic_type
        return FlowLog(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The Description of the VPC Flow Log.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="flowLogName")
    def flow_log_name(self) -> pulumi.Output[Optional[str]]:
        """
        The Name of the VPC Flow Log.
        """
        return pulumi.get(self, "flow_log_name")

    @property
    @pulumi.getter(name="logStoreName")
    def log_store_name(self) -> pulumi.Output[str]:
        """
        The name of the logstore.
        """
        return pulumi.get(self, "log_store_name")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[str]:
        """
        The name of the project.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[str]:
        """
        The type of the resource to capture traffic.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the VPC Flow Log.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="trafficType")
    def traffic_type(self) -> pulumi.Output[str]:
        """
        The type of traffic collected.
        """
        return pulumi.get(self, "traffic_type")

