# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DashboardArgs', 'Dashboard']

@pulumi.input_type
class DashboardArgs:
    def __init__(__self__, *,
                 char_list: pulumi.Input[str],
                 dashboard_name: pulumi.Input[str],
                 project_name: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Dashboard resource.
        :param pulumi.Input[str] char_list: Configuration of charts in the dashboard.
        :param pulumi.Input[str] dashboard_name: The name of the Log Dashboard.
        :param pulumi.Input[str] project_name: The name of the log project. It is the only in one Alicloud account.
        :param pulumi.Input[str] display_name: Dashboard alias.
        """
        pulumi.set(__self__, "char_list", char_list)
        pulumi.set(__self__, "dashboard_name", dashboard_name)
        pulumi.set(__self__, "project_name", project_name)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)

    @property
    @pulumi.getter(name="charList")
    def char_list(self) -> pulumi.Input[str]:
        """
        Configuration of charts in the dashboard.
        """
        return pulumi.get(self, "char_list")

    @char_list.setter
    def char_list(self, value: pulumi.Input[str]):
        pulumi.set(self, "char_list", value)

    @property
    @pulumi.getter(name="dashboardName")
    def dashboard_name(self) -> pulumi.Input[str]:
        """
        The name of the Log Dashboard.
        """
        return pulumi.get(self, "dashboard_name")

    @dashboard_name.setter
    def dashboard_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "dashboard_name", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Input[str]:
        """
        The name of the log project. It is the only in one Alicloud account.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Dashboard alias.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)


@pulumi.input_type
class _DashboardState:
    def __init__(__self__, *,
                 char_list: Optional[pulumi.Input[str]] = None,
                 dashboard_name: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Dashboard resources.
        :param pulumi.Input[str] char_list: Configuration of charts in the dashboard.
        :param pulumi.Input[str] dashboard_name: The name of the Log Dashboard.
        :param pulumi.Input[str] display_name: Dashboard alias.
        :param pulumi.Input[str] project_name: The name of the log project. It is the only in one Alicloud account.
        """
        if char_list is not None:
            pulumi.set(__self__, "char_list", char_list)
        if dashboard_name is not None:
            pulumi.set(__self__, "dashboard_name", dashboard_name)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)

    @property
    @pulumi.getter(name="charList")
    def char_list(self) -> Optional[pulumi.Input[str]]:
        """
        Configuration of charts in the dashboard.
        """
        return pulumi.get(self, "char_list")

    @char_list.setter
    def char_list(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "char_list", value)

    @property
    @pulumi.getter(name="dashboardName")
    def dashboard_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Log Dashboard.
        """
        return pulumi.get(self, "dashboard_name")

    @dashboard_name.setter
    def dashboard_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dashboard_name", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Dashboard alias.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the log project. It is the only in one Alicloud account.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)


class Dashboard(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 char_list: Optional[pulumi.Input[str]] = None,
                 dashboard_name: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The dashboard is a real-time data analysis platform provided by the log service. You can display frequently used query and analysis statements in the form of charts and save statistical charts to the dashboard.
        [Refer to details](https://www.alibabacloud.com/help/doc-detail/102530.htm).

        > **NOTE:** Available in 1.86.0

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_project = alicloud.log.Project("defaultProject", description="tf unit test")
        default_store = alicloud.log.Store("defaultStore",
            project="tf-project",
            retention_period=3000,
            shard_count=1)
        example = alicloud.log.Dashboard("example",
            char_list=\"\"\"  [
            {
              "title":"new_title",
              "type":"map",
              "search":{
                "logstore":"tf-logstore",
                "topic":"new_topic",
                "query":"* | SELECT COUNT(name) as ct_name, COUNT(product) as ct_product, name,product GROUP BY name,product",
                "start":"-86400s",
                "end":"now"
              },
              "display":{
                "xAxis":[
                  "ct_name"
                ],
                "yAxis":[
                  "ct_product"
                ],
                "xPos":0,
                "yPos":0,
                "width":10,
                "height":12,
                "displayName":"xixihaha911"
              }
            }
          ]

        \"\"\",
            dashboard_name="tf-dashboard",
            project_name="tf-project")
        ```

        ## Import

        Log Dashboard can be imported using the id or name, e.g.

        ```sh
         $ pulumi import alicloud:log/dashboard:Dashboard example tf-project:tf-logstore:tf-dashboard
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] char_list: Configuration of charts in the dashboard.
        :param pulumi.Input[str] dashboard_name: The name of the Log Dashboard.
        :param pulumi.Input[str] display_name: Dashboard alias.
        :param pulumi.Input[str] project_name: The name of the log project. It is the only in one Alicloud account.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DashboardArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The dashboard is a real-time data analysis platform provided by the log service. You can display frequently used query and analysis statements in the form of charts and save statistical charts to the dashboard.
        [Refer to details](https://www.alibabacloud.com/help/doc-detail/102530.htm).

        > **NOTE:** Available in 1.86.0

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_project = alicloud.log.Project("defaultProject", description="tf unit test")
        default_store = alicloud.log.Store("defaultStore",
            project="tf-project",
            retention_period=3000,
            shard_count=1)
        example = alicloud.log.Dashboard("example",
            char_list=\"\"\"  [
            {
              "title":"new_title",
              "type":"map",
              "search":{
                "logstore":"tf-logstore",
                "topic":"new_topic",
                "query":"* | SELECT COUNT(name) as ct_name, COUNT(product) as ct_product, name,product GROUP BY name,product",
                "start":"-86400s",
                "end":"now"
              },
              "display":{
                "xAxis":[
                  "ct_name"
                ],
                "yAxis":[
                  "ct_product"
                ],
                "xPos":0,
                "yPos":0,
                "width":10,
                "height":12,
                "displayName":"xixihaha911"
              }
            }
          ]

        \"\"\",
            dashboard_name="tf-dashboard",
            project_name="tf-project")
        ```

        ## Import

        Log Dashboard can be imported using the id or name, e.g.

        ```sh
         $ pulumi import alicloud:log/dashboard:Dashboard example tf-project:tf-logstore:tf-dashboard
        ```

        :param str resource_name: The name of the resource.
        :param DashboardArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DashboardArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 char_list: Optional[pulumi.Input[str]] = None,
                 dashboard_name: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = DashboardArgs.__new__(DashboardArgs)

            if char_list is None and not opts.urn:
                raise TypeError("Missing required property 'char_list'")
            __props__.__dict__["char_list"] = char_list
            if dashboard_name is None and not opts.urn:
                raise TypeError("Missing required property 'dashboard_name'")
            __props__.__dict__["dashboard_name"] = dashboard_name
            __props__.__dict__["display_name"] = display_name
            if project_name is None and not opts.urn:
                raise TypeError("Missing required property 'project_name'")
            __props__.__dict__["project_name"] = project_name
        super(Dashboard, __self__).__init__(
            'alicloud:log/dashboard:Dashboard',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            char_list: Optional[pulumi.Input[str]] = None,
            dashboard_name: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None) -> 'Dashboard':
        """
        Get an existing Dashboard resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] char_list: Configuration of charts in the dashboard.
        :param pulumi.Input[str] dashboard_name: The name of the Log Dashboard.
        :param pulumi.Input[str] display_name: Dashboard alias.
        :param pulumi.Input[str] project_name: The name of the log project. It is the only in one Alicloud account.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DashboardState.__new__(_DashboardState)

        __props__.__dict__["char_list"] = char_list
        __props__.__dict__["dashboard_name"] = dashboard_name
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["project_name"] = project_name
        return Dashboard(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="charList")
    def char_list(self) -> pulumi.Output[str]:
        """
        Configuration of charts in the dashboard.
        """
        return pulumi.get(self, "char_list")

    @property
    @pulumi.getter(name="dashboardName")
    def dashboard_name(self) -> pulumi.Output[str]:
        """
        The name of the Log Dashboard.
        """
        return pulumi.get(self, "dashboard_name")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        Dashboard alias.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[str]:
        """
        The name of the log project. It is the only in one Alicloud account.
        """
        return pulumi.get(self, "project_name")

