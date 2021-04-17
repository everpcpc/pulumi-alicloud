# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetRouteTablesResult',
    'AwaitableGetRouteTablesResult',
    'get_route_tables',
]

@pulumi.output_type
class GetRouteTablesResult:
    """
    A collection of values returned by getRouteTables.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, names=None, output_file=None, resource_group_id=None, route_table_name=None, router_id=None, router_type=None, status=None, tables=None, tags=None, vpc_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        if route_table_name and not isinstance(route_table_name, str):
            raise TypeError("Expected argument 'route_table_name' to be a str")
        pulumi.set(__self__, "route_table_name", route_table_name)
        if router_id and not isinstance(router_id, str):
            raise TypeError("Expected argument 'router_id' to be a str")
        pulumi.set(__self__, "router_id", router_id)
        if router_type and not isinstance(router_type, str):
            raise TypeError("Expected argument 'router_type' to be a str")
        pulumi.set(__self__, "router_type", router_type)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tables and not isinstance(tables, list):
            raise TypeError("Expected argument 'tables' to be a list")
        pulumi.set(__self__, "tables", tables)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        """
        (Optional) A list of Route Tables IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of Route Tables names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        """
        The Id of resource group which route tables belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="routeTableName")
    def route_table_name(self) -> Optional[str]:
        """
        The route table name.
        """
        return pulumi.get(self, "route_table_name")

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> Optional[str]:
        """
        Router Id of the route table.
        """
        return pulumi.get(self, "router_id")

    @property
    @pulumi.getter(name="routerType")
    def router_type(self) -> Optional[str]:
        """
        The route type.
        """
        return pulumi.get(self, "router_type")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of route table.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tables(self) -> Sequence['outputs.GetRouteTablesTableResult']:
        """
        A list of Route Tables. Each element contains the following attributes:
        """
        return pulumi.get(self, "tables")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        The VPC ID.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableGetRouteTablesResult(GetRouteTablesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouteTablesResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            resource_group_id=self.resource_group_id,
            route_table_name=self.route_table_name,
            router_id=self.router_id,
            router_type=self.router_type,
            status=self.status,
            tables=self.tables,
            tags=self.tags,
            vpc_id=self.vpc_id)


def get_route_tables(ids: Optional[Sequence[str]] = None,
                     name_regex: Optional[str] = None,
                     output_file: Optional[str] = None,
                     resource_group_id: Optional[str] = None,
                     route_table_name: Optional[str] = None,
                     router_id: Optional[str] = None,
                     router_type: Optional[str] = None,
                     status: Optional[str] = None,
                     tags: Optional[Mapping[str, Any]] = None,
                     vpc_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouteTablesResult:
    """
    This data source provides a list of Route Tables owned by an Alibaba Cloud account.

    > **NOTE:** Available in 1.36.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "route-tables-datasource-example-name"
    foo_network = alicloud.vpc.Network("fooNetwork",
        cidr_block="172.16.0.0/12",
        vpc_name=name)
    foo_route_table = alicloud.vpc.RouteTable("fooRouteTable",
        description=name,
        route_table_name=name,
        vpc_id=foo_network.id)
    foo_route_tables = foo_route_table.id.apply(lambda id: alicloud.vpc.get_route_tables(ids=[id]))
    pulumi.export("routeTableIds", foo_route_tables.ids)
    ```


    :param Sequence[str] ids: A list of Route Tables IDs.
    :param str name_regex: A regex string to filter route tables by name.
    :param str resource_group_id: The Id of resource group which route tables belongs.
    :param str route_table_name: The route table name.
    :param str router_id: The router ID.
    :param str router_type: The route type of route table. Valid values: `VRouter` and `VBR`.
    :param str status: The status of resource. Valid values: `Available` and `Pending`.
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    :param str vpc_id: Vpc id of the route table.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['resourceGroupId'] = resource_group_id
    __args__['routeTableName'] = route_table_name
    __args__['routerId'] = router_id
    __args__['routerType'] = router_type
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:vpc/getRouteTables:getRouteTables', __args__, opts=opts, typ=GetRouteTablesResult).value

    return AwaitableGetRouteTablesResult(
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        resource_group_id=__ret__.resource_group_id,
        route_table_name=__ret__.route_table_name,
        router_id=__ret__.router_id,
        router_type=__ret__.router_type,
        status=__ret__.status,
        tables=__ret__.tables,
        tags=__ret__.tags,
        vpc_id=__ret__.vpc_id)
