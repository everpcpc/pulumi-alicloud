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
    'GetApplicationsResult',
    'AwaitableGetApplicationsResult',
    'get_applications',
    'get_applications_output',
]

@pulumi.output_type
class GetApplicationsResult:
    """
    A collection of values returned by getApplications.
    """
    def __init__(__self__, app_name=None, applications=None, enable_details=None, field_type=None, field_value=None, id=None, ids=None, namespace_id=None, order_by=None, output_file=None, reverse=None, status=None):
        if app_name and not isinstance(app_name, str):
            raise TypeError("Expected argument 'app_name' to be a str")
        pulumi.set(__self__, "app_name", app_name)
        if applications and not isinstance(applications, list):
            raise TypeError("Expected argument 'applications' to be a list")
        pulumi.set(__self__, "applications", applications)
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if field_type and not isinstance(field_type, str):
            raise TypeError("Expected argument 'field_type' to be a str")
        pulumi.set(__self__, "field_type", field_type)
        if field_value and not isinstance(field_value, str):
            raise TypeError("Expected argument 'field_value' to be a str")
        pulumi.set(__self__, "field_value", field_value)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if namespace_id and not isinstance(namespace_id, str):
            raise TypeError("Expected argument 'namespace_id' to be a str")
        pulumi.set(__self__, "namespace_id", namespace_id)
        if order_by and not isinstance(order_by, str):
            raise TypeError("Expected argument 'order_by' to be a str")
        pulumi.set(__self__, "order_by", order_by)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if reverse and not isinstance(reverse, bool):
            raise TypeError("Expected argument 'reverse' to be a bool")
        pulumi.set(__self__, "reverse", reverse)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="appName")
    def app_name(self) -> Optional[str]:
        return pulumi.get(self, "app_name")

    @property
    @pulumi.getter
    def applications(self) -> Sequence['outputs.GetApplicationsApplicationResult']:
        return pulumi.get(self, "applications")

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

    @property
    @pulumi.getter(name="fieldType")
    def field_type(self) -> Optional[str]:
        return pulumi.get(self, "field_type")

    @property
    @pulumi.getter(name="fieldValue")
    def field_value(self) -> Optional[str]:
        return pulumi.get(self, "field_value")

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
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[str]:
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional[str]:
        return pulumi.get(self, "order_by")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def reverse(self) -> Optional[bool]:
        return pulumi.get(self, "reverse")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


class AwaitableGetApplicationsResult(GetApplicationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationsResult(
            app_name=self.app_name,
            applications=self.applications,
            enable_details=self.enable_details,
            field_type=self.field_type,
            field_value=self.field_value,
            id=self.id,
            ids=self.ids,
            namespace_id=self.namespace_id,
            order_by=self.order_by,
            output_file=self.output_file,
            reverse=self.reverse,
            status=self.status)


def get_applications(app_name: Optional[str] = None,
                     enable_details: Optional[bool] = None,
                     field_type: Optional[str] = None,
                     field_value: Optional[str] = None,
                     ids: Optional[Sequence[str]] = None,
                     namespace_id: Optional[str] = None,
                     order_by: Optional[str] = None,
                     output_file: Optional[str] = None,
                     reverse: Optional[bool] = None,
                     status: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationsResult:
    """
    This data source provides the Sae Applications of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.161.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "tf-testacc"
    default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
    vpc = alicloud.vpc.Network("vpc",
        vpc_name="tf_testacc",
        cidr_block="172.16.0.0/12")
    vsw = alicloud.vpc.Switch("vsw",
        vpc_id=vpc.id,
        cidr_block="172.16.0.0/24",
        zone_id=default_zones.zones[0].id,
        vswitch_name=name)
    default_namespace = alicloud.sae.Namespace("defaultNamespace",
        namespace_description=name,
        namespace_id="cn-hangzhou:tfacctest",
        namespace_name=name)
    default_application = alicloud.sae.Application("defaultApplication",
        app_description="tf-testaccDescription",
        app_name="tf-testaccAppName131",
        namespace_id=default_namespace.id,
        image_url="registry-vpc.cn-hangzhou.aliyuncs.com/lxepoo/apache-php5",
        package_type="Image",
        vswitch_id=vsw.id,
        timezone="Asia/Beijing",
        replicas=5,
        cpu=500,
        memory=2048)
    default_applications = alicloud.sae.get_applications_output(ids=[default_application.id])
    pulumi.export("saeApplicationId", default_applications.applications[0].id)
    ```


    :param str app_name: Application Name. Combinations of numbers, letters, and dashes (-) are allowed. It must start with a letter and the maximum length is 36 characters.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param str field_type: The field type. Valid values:`appName`, `appIds`, `slbIps`, `instanceIps`
    :param str field_value: The field value.
    :param Sequence[str] ids: A list of Application IDs.
    :param str namespace_id: SAE namespace ID. Only namespaces whose names are lowercase letters and dashes (-) are supported, and must start with a letter. The namespace can be obtained by calling the DescribeNamespaceList interface.
    :param str order_by: The order by.Valid values:`running`,`instances`.
    :param bool reverse: The reverse.
    :param str status: The status of the resource.
    """
    __args__ = dict()
    __args__['appName'] = app_name
    __args__['enableDetails'] = enable_details
    __args__['fieldType'] = field_type
    __args__['fieldValue'] = field_value
    __args__['ids'] = ids
    __args__['namespaceId'] = namespace_id
    __args__['orderBy'] = order_by
    __args__['outputFile'] = output_file
    __args__['reverse'] = reverse
    __args__['status'] = status
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:sae/getApplications:getApplications', __args__, opts=opts, typ=GetApplicationsResult).value

    return AwaitableGetApplicationsResult(
        app_name=__ret__.app_name,
        applications=__ret__.applications,
        enable_details=__ret__.enable_details,
        field_type=__ret__.field_type,
        field_value=__ret__.field_value,
        id=__ret__.id,
        ids=__ret__.ids,
        namespace_id=__ret__.namespace_id,
        order_by=__ret__.order_by,
        output_file=__ret__.output_file,
        reverse=__ret__.reverse,
        status=__ret__.status)


@_utilities.lift_output_func(get_applications)
def get_applications_output(app_name: Optional[pulumi.Input[Optional[str]]] = None,
                            enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                            field_type: Optional[pulumi.Input[Optional[str]]] = None,
                            field_value: Optional[pulumi.Input[Optional[str]]] = None,
                            ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            namespace_id: Optional[pulumi.Input[Optional[str]]] = None,
                            order_by: Optional[pulumi.Input[Optional[str]]] = None,
                            output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            reverse: Optional[pulumi.Input[Optional[bool]]] = None,
                            status: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationsResult]:
    """
    This data source provides the Sae Applications of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.161.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "tf-testacc"
    default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
    vpc = alicloud.vpc.Network("vpc",
        vpc_name="tf_testacc",
        cidr_block="172.16.0.0/12")
    vsw = alicloud.vpc.Switch("vsw",
        vpc_id=vpc.id,
        cidr_block="172.16.0.0/24",
        zone_id=default_zones.zones[0].id,
        vswitch_name=name)
    default_namespace = alicloud.sae.Namespace("defaultNamespace",
        namespace_description=name,
        namespace_id="cn-hangzhou:tfacctest",
        namespace_name=name)
    default_application = alicloud.sae.Application("defaultApplication",
        app_description="tf-testaccDescription",
        app_name="tf-testaccAppName131",
        namespace_id=default_namespace.id,
        image_url="registry-vpc.cn-hangzhou.aliyuncs.com/lxepoo/apache-php5",
        package_type="Image",
        vswitch_id=vsw.id,
        timezone="Asia/Beijing",
        replicas=5,
        cpu=500,
        memory=2048)
    default_applications = alicloud.sae.get_applications_output(ids=[default_application.id])
    pulumi.export("saeApplicationId", default_applications.applications[0].id)
    ```


    :param str app_name: Application Name. Combinations of numbers, letters, and dashes (-) are allowed. It must start with a letter and the maximum length is 36 characters.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param str field_type: The field type. Valid values:`appName`, `appIds`, `slbIps`, `instanceIps`
    :param str field_value: The field value.
    :param Sequence[str] ids: A list of Application IDs.
    :param str namespace_id: SAE namespace ID. Only namespaces whose names are lowercase letters and dashes (-) are supported, and must start with a letter. The namespace can be obtained by calling the DescribeNamespaceList interface.
    :param str order_by: The order by.Valid values:`running`,`instances`.
    :param bool reverse: The reverse.
    :param str status: The status of the resource.
    """
    ...
