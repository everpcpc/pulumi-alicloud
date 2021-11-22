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
    'GetEnterpriseInstancesResult',
    'AwaitableGetEnterpriseInstancesResult',
    'get_enterprise_instances',
    'get_enterprise_instances_output',
]

@pulumi.output_type
class GetEnterpriseInstancesResult:
    """
    A collection of values returned by getEnterpriseInstances.
    """
    def __init__(__self__, env_type=None, id=None, ids=None, instance_alias_regex=None, instance_source=None, instance_type=None, instances=None, name_regex=None, names=None, net_type=None, output_file=None, search_key=None, status=None, tid=None):
        if env_type and not isinstance(env_type, str):
            raise TypeError("Expected argument 'env_type' to be a str")
        pulumi.set(__self__, "env_type", env_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_alias_regex and not isinstance(instance_alias_regex, str):
            raise TypeError("Expected argument 'instance_alias_regex' to be a str")
        pulumi.set(__self__, "instance_alias_regex", instance_alias_regex)
        if instance_source and not isinstance(instance_source, str):
            raise TypeError("Expected argument 'instance_source' to be a str")
        pulumi.set(__self__, "instance_source", instance_source)
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        pulumi.set(__self__, "instance_type", instance_type)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if net_type and not isinstance(net_type, str):
            raise TypeError("Expected argument 'net_type' to be a str")
        pulumi.set(__self__, "net_type", net_type)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if search_key and not isinstance(search_key, str):
            raise TypeError("Expected argument 'search_key' to be a str")
        pulumi.set(__self__, "search_key", search_key)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tid and not isinstance(tid, int):
            raise TypeError("Expected argument 'tid' to be a int")
        pulumi.set(__self__, "tid", tid)

    @property
    @pulumi.getter(name="envType")
    def env_type(self) -> Optional[str]:
        """
        The type of the environment to which the database instance belongs..
        """
        return pulumi.get(self, "env_type")

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
        A list of DMS Enterprise IDs (Each of them consists of host:port).
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="instanceAliasRegex")
    def instance_alias_regex(self) -> Optional[str]:
        return pulumi.get(self, "instance_alias_regex")

    @property
    @pulumi.getter(name="instanceSource")
    def instance_source(self) -> Optional[str]:
        """
        The ID of the database instance.
        """
        return pulumi.get(self, "instance_source")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[str]:
        """
        The ID of the database instance.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def instances(self) -> Sequence['outputs.GetEnterpriseInstancesInstanceResult']:
        """
        A list of KMS keys. Each element contains the following attributes:
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of DMS Enterprise names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="netType")
    def net_type(self) -> Optional[str]:
        return pulumi.get(self, "net_type")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="searchKey")
    def search_key(self) -> Optional[str]:
        return pulumi.get(self, "search_key")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the database instance.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tid(self) -> Optional[int]:
        return pulumi.get(self, "tid")


class AwaitableGetEnterpriseInstancesResult(GetEnterpriseInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnterpriseInstancesResult(
            env_type=self.env_type,
            id=self.id,
            ids=self.ids,
            instance_alias_regex=self.instance_alias_regex,
            instance_source=self.instance_source,
            instance_type=self.instance_type,
            instances=self.instances,
            name_regex=self.name_regex,
            names=self.names,
            net_type=self.net_type,
            output_file=self.output_file,
            search_key=self.search_key,
            status=self.status,
            tid=self.tid)


def get_enterprise_instances(env_type: Optional[str] = None,
                             instance_alias_regex: Optional[str] = None,
                             instance_source: Optional[str] = None,
                             instance_type: Optional[str] = None,
                             name_regex: Optional[str] = None,
                             net_type: Optional[str] = None,
                             output_file: Optional[str] = None,
                             search_key: Optional[str] = None,
                             status: Optional[str] = None,
                             tid: Optional[int] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnterpriseInstancesResult:
    """
    This data source provides a list of DMS Enterprise Instances in an Alibaba Cloud account according to the specified filters.

    > **NOTE:** Available in 1.88.0+


    :param str env_type: The type of the environment to which the database instance belongs.
    :param str instance_alias_regex: A regex string to filter the results by the DMS Enterprise Instance instance_alias.
    :param str instance_source: The source of the database instance.
    :param str instance_type: The ID of the database instance.
    :param str name_regex: A regex string to filter the results by the DMS Enterprise Instance instance_alias.
    :param str net_type: The network type of the database instance. Valid values: CLASSIC and VPC. For more information about the valid values, see the description of the RegisterInstance operation.
    :param str search_key: The keyword used to query database instances.
    :param str status: Filter the results by status of the DMS Enterprise Instances. Valid values: `NORMAL`, `UNAVAILABLE`, `UNKNOWN`, `DELETED`, `DISABLE`.
    :param int tid: The ID of the tenant in Data Management (DMS) Enterprise.
    """
    __args__ = dict()
    __args__['envType'] = env_type
    __args__['instanceAliasRegex'] = instance_alias_regex
    __args__['instanceSource'] = instance_source
    __args__['instanceType'] = instance_type
    __args__['nameRegex'] = name_regex
    __args__['netType'] = net_type
    __args__['outputFile'] = output_file
    __args__['searchKey'] = search_key
    __args__['status'] = status
    __args__['tid'] = tid
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:dms/getEnterpriseInstances:getEnterpriseInstances', __args__, opts=opts, typ=GetEnterpriseInstancesResult).value

    return AwaitableGetEnterpriseInstancesResult(
        env_type=__ret__.env_type,
        id=__ret__.id,
        ids=__ret__.ids,
        instance_alias_regex=__ret__.instance_alias_regex,
        instance_source=__ret__.instance_source,
        instance_type=__ret__.instance_type,
        instances=__ret__.instances,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        net_type=__ret__.net_type,
        output_file=__ret__.output_file,
        search_key=__ret__.search_key,
        status=__ret__.status,
        tid=__ret__.tid)


@_utilities.lift_output_func(get_enterprise_instances)
def get_enterprise_instances_output(env_type: Optional[pulumi.Input[Optional[str]]] = None,
                                    instance_alias_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                    instance_source: Optional[pulumi.Input[Optional[str]]] = None,
                                    instance_type: Optional[pulumi.Input[Optional[str]]] = None,
                                    name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                    net_type: Optional[pulumi.Input[Optional[str]]] = None,
                                    output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                    search_key: Optional[pulumi.Input[Optional[str]]] = None,
                                    status: Optional[pulumi.Input[Optional[str]]] = None,
                                    tid: Optional[pulumi.Input[Optional[int]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEnterpriseInstancesResult]:
    """
    This data source provides a list of DMS Enterprise Instances in an Alibaba Cloud account according to the specified filters.

    > **NOTE:** Available in 1.88.0+


    :param str env_type: The type of the environment to which the database instance belongs.
    :param str instance_alias_regex: A regex string to filter the results by the DMS Enterprise Instance instance_alias.
    :param str instance_source: The source of the database instance.
    :param str instance_type: The ID of the database instance.
    :param str name_regex: A regex string to filter the results by the DMS Enterprise Instance instance_alias.
    :param str net_type: The network type of the database instance. Valid values: CLASSIC and VPC. For more information about the valid values, see the description of the RegisterInstance operation.
    :param str search_key: The keyword used to query database instances.
    :param str status: Filter the results by status of the DMS Enterprise Instances. Valid values: `NORMAL`, `UNAVAILABLE`, `UNKNOWN`, `DELETED`, `DISABLE`.
    :param int tid: The ID of the tenant in Data Management (DMS) Enterprise.
    """
    ...
