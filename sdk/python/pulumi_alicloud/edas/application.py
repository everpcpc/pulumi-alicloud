# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApplicationArgs', 'Application']

@pulumi.input_type
class ApplicationArgs:
    def __init__(__self__, *,
                 application_name: pulumi.Input[str],
                 cluster_id: pulumi.Input[str],
                 package_type: pulumi.Input[str],
                 build_pack_id: Optional[pulumi.Input[int]] = None,
                 descriotion: Optional[pulumi.Input[str]] = None,
                 ecu_infos: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 health_check_url: Optional[pulumi.Input[str]] = None,
                 logical_region_id: Optional[pulumi.Input[str]] = None,
                 package_version: Optional[pulumi.Input[str]] = None,
                 war_url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Application resource.
        :param pulumi.Input[str] application_name: Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
        :param pulumi.Input[str] cluster_id: The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
        :param pulumi.Input[str] package_type: The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
        :param pulumi.Input[int] build_pack_id: The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
        :param pulumi.Input[str] descriotion: The description of the application that you want to create.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ecu_infos: The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
        :param pulumi.Input[str] group_id: The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        :param pulumi.Input[str] health_check_url: The URL for health checking of the application.
        :param pulumi.Input[str] logical_region_id: The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
        :param pulumi.Input[str] package_version: The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        :param pulumi.Input[str] war_url: The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        """
        pulumi.set(__self__, "application_name", application_name)
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "package_type", package_type)
        if build_pack_id is not None:
            pulumi.set(__self__, "build_pack_id", build_pack_id)
        if descriotion is not None:
            pulumi.set(__self__, "descriotion", descriotion)
        if ecu_infos is not None:
            pulumi.set(__self__, "ecu_infos", ecu_infos)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if health_check_url is not None:
            pulumi.set(__self__, "health_check_url", health_check_url)
        if logical_region_id is not None:
            pulumi.set(__self__, "logical_region_id", logical_region_id)
        if package_version is not None:
            pulumi.set(__self__, "package_version", package_version)
        if war_url is not None:
            pulumi.set(__self__, "war_url", war_url)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> pulumi.Input[str]:
        """
        Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
        """
        return pulumi.get(self, "application_name")

    @application_name.setter
    def application_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_name", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> pulumi.Input[str]:
        """
        The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
        """
        return pulumi.get(self, "package_type")

    @package_type.setter
    def package_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "package_type", value)

    @property
    @pulumi.getter(name="buildPackId")
    def build_pack_id(self) -> Optional[pulumi.Input[int]]:
        """
        The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
        """
        return pulumi.get(self, "build_pack_id")

    @build_pack_id.setter
    def build_pack_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "build_pack_id", value)

    @property
    @pulumi.getter
    def descriotion(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the application that you want to create.
        """
        return pulumi.get(self, "descriotion")

    @descriotion.setter
    def descriotion(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "descriotion", value)

    @property
    @pulumi.getter(name="ecuInfos")
    def ecu_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
        """
        return pulumi.get(self, "ecu_infos")

    @ecu_infos.setter
    def ecu_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ecu_infos", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="healthCheckUrl")
    def health_check_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL for health checking of the application.
        """
        return pulumi.get(self, "health_check_url")

    @health_check_url.setter
    def health_check_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "health_check_url", value)

    @property
    @pulumi.getter(name="logicalRegionId")
    def logical_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
        """
        return pulumi.get(self, "logical_region_id")

    @logical_region_id.setter
    def logical_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logical_region_id", value)

    @property
    @pulumi.getter(name="packageVersion")
    def package_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        """
        return pulumi.get(self, "package_version")

    @package_version.setter
    def package_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "package_version", value)

    @property
    @pulumi.getter(name="warUrl")
    def war_url(self) -> Optional[pulumi.Input[str]]:
        """
        The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        """
        return pulumi.get(self, "war_url")

    @war_url.setter
    def war_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "war_url", value)


@pulumi.input_type
class _ApplicationState:
    def __init__(__self__, *,
                 application_name: Optional[pulumi.Input[str]] = None,
                 build_pack_id: Optional[pulumi.Input[int]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 descriotion: Optional[pulumi.Input[str]] = None,
                 ecu_infos: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 health_check_url: Optional[pulumi.Input[str]] = None,
                 logical_region_id: Optional[pulumi.Input[str]] = None,
                 package_type: Optional[pulumi.Input[str]] = None,
                 package_version: Optional[pulumi.Input[str]] = None,
                 war_url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Application resources.
        :param pulumi.Input[str] application_name: Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
        :param pulumi.Input[int] build_pack_id: The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
        :param pulumi.Input[str] cluster_id: The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
        :param pulumi.Input[str] descriotion: The description of the application that you want to create.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ecu_infos: The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
        :param pulumi.Input[str] group_id: The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        :param pulumi.Input[str] health_check_url: The URL for health checking of the application.
        :param pulumi.Input[str] logical_region_id: The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
        :param pulumi.Input[str] package_type: The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
        :param pulumi.Input[str] package_version: The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        :param pulumi.Input[str] war_url: The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        """
        if application_name is not None:
            pulumi.set(__self__, "application_name", application_name)
        if build_pack_id is not None:
            pulumi.set(__self__, "build_pack_id", build_pack_id)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if descriotion is not None:
            pulumi.set(__self__, "descriotion", descriotion)
        if ecu_infos is not None:
            pulumi.set(__self__, "ecu_infos", ecu_infos)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if health_check_url is not None:
            pulumi.set(__self__, "health_check_url", health_check_url)
        if logical_region_id is not None:
            pulumi.set(__self__, "logical_region_id", logical_region_id)
        if package_type is not None:
            pulumi.set(__self__, "package_type", package_type)
        if package_version is not None:
            pulumi.set(__self__, "package_version", package_version)
        if war_url is not None:
            pulumi.set(__self__, "war_url", war_url)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
        """
        return pulumi.get(self, "application_name")

    @application_name.setter
    def application_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_name", value)

    @property
    @pulumi.getter(name="buildPackId")
    def build_pack_id(self) -> Optional[pulumi.Input[int]]:
        """
        The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
        """
        return pulumi.get(self, "build_pack_id")

    @build_pack_id.setter
    def build_pack_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "build_pack_id", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def descriotion(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the application that you want to create.
        """
        return pulumi.get(self, "descriotion")

    @descriotion.setter
    def descriotion(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "descriotion", value)

    @property
    @pulumi.getter(name="ecuInfos")
    def ecu_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
        """
        return pulumi.get(self, "ecu_infos")

    @ecu_infos.setter
    def ecu_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ecu_infos", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="healthCheckUrl")
    def health_check_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL for health checking of the application.
        """
        return pulumi.get(self, "health_check_url")

    @health_check_url.setter
    def health_check_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "health_check_url", value)

    @property
    @pulumi.getter(name="logicalRegionId")
    def logical_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
        """
        return pulumi.get(self, "logical_region_id")

    @logical_region_id.setter
    def logical_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logical_region_id", value)

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
        """
        return pulumi.get(self, "package_type")

    @package_type.setter
    def package_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "package_type", value)

    @property
    @pulumi.getter(name="packageVersion")
    def package_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        """
        return pulumi.get(self, "package_version")

    @package_version.setter
    def package_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "package_version", value)

    @property
    @pulumi.getter(name="warUrl")
    def war_url(self) -> Optional[pulumi.Input[str]]:
        """
        The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        """
        return pulumi.get(self, "war_url")

    @war_url.setter
    def war_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "war_url", value)


class Application(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_name: Optional[pulumi.Input[str]] = None,
                 build_pack_id: Optional[pulumi.Input[int]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 descriotion: Optional[pulumi.Input[str]] = None,
                 ecu_infos: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 health_check_url: Optional[pulumi.Input[str]] = None,
                 logical_region_id: Optional[pulumi.Input[str]] = None,
                 package_type: Optional[pulumi.Input[str]] = None,
                 package_version: Optional[pulumi.Input[str]] = None,
                 war_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an EDAS ecs application on EDAS. The application will be deployed when `group_id` and `war_url` are given.

        > **NOTE:** Available in 1.82.0+

        ## Import

        EDAS application can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:edas/application:Application app app_Id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_name: Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
        :param pulumi.Input[int] build_pack_id: The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
        :param pulumi.Input[str] cluster_id: The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
        :param pulumi.Input[str] descriotion: The description of the application that you want to create.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ecu_infos: The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
        :param pulumi.Input[str] group_id: The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        :param pulumi.Input[str] health_check_url: The URL for health checking of the application.
        :param pulumi.Input[str] logical_region_id: The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
        :param pulumi.Input[str] package_type: The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
        :param pulumi.Input[str] package_version: The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        :param pulumi.Input[str] war_url: The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an EDAS ecs application on EDAS. The application will be deployed when `group_id` and `war_url` are given.

        > **NOTE:** Available in 1.82.0+

        ## Import

        EDAS application can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:edas/application:Application app app_Id
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_name: Optional[pulumi.Input[str]] = None,
                 build_pack_id: Optional[pulumi.Input[int]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 descriotion: Optional[pulumi.Input[str]] = None,
                 ecu_infos: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 health_check_url: Optional[pulumi.Input[str]] = None,
                 logical_region_id: Optional[pulumi.Input[str]] = None,
                 package_type: Optional[pulumi.Input[str]] = None,
                 package_version: Optional[pulumi.Input[str]] = None,
                 war_url: Optional[pulumi.Input[str]] = None,
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
            __props__ = ApplicationArgs.__new__(ApplicationArgs)

            if application_name is None and not opts.urn:
                raise TypeError("Missing required property 'application_name'")
            __props__.__dict__["application_name"] = application_name
            __props__.__dict__["build_pack_id"] = build_pack_id
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["descriotion"] = descriotion
            __props__.__dict__["ecu_infos"] = ecu_infos
            __props__.__dict__["group_id"] = group_id
            __props__.__dict__["health_check_url"] = health_check_url
            __props__.__dict__["logical_region_id"] = logical_region_id
            if package_type is None and not opts.urn:
                raise TypeError("Missing required property 'package_type'")
            __props__.__dict__["package_type"] = package_type
            __props__.__dict__["package_version"] = package_version
            __props__.__dict__["war_url"] = war_url
        super(Application, __self__).__init__(
            'alicloud:edas/application:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_name: Optional[pulumi.Input[str]] = None,
            build_pack_id: Optional[pulumi.Input[int]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            descriotion: Optional[pulumi.Input[str]] = None,
            ecu_infos: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            health_check_url: Optional[pulumi.Input[str]] = None,
            logical_region_id: Optional[pulumi.Input[str]] = None,
            package_type: Optional[pulumi.Input[str]] = None,
            package_version: Optional[pulumi.Input[str]] = None,
            war_url: Optional[pulumi.Input[str]] = None) -> 'Application':
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_name: Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
        :param pulumi.Input[int] build_pack_id: The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
        :param pulumi.Input[str] cluster_id: The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
        :param pulumi.Input[str] descriotion: The description of the application that you want to create.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ecu_infos: The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
        :param pulumi.Input[str] group_id: The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        :param pulumi.Input[str] health_check_url: The URL for health checking of the application.
        :param pulumi.Input[str] logical_region_id: The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
        :param pulumi.Input[str] package_type: The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
        :param pulumi.Input[str] package_version: The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        :param pulumi.Input[str] war_url: The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationState.__new__(_ApplicationState)

        __props__.__dict__["application_name"] = application_name
        __props__.__dict__["build_pack_id"] = build_pack_id
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["descriotion"] = descriotion
        __props__.__dict__["ecu_infos"] = ecu_infos
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["health_check_url"] = health_check_url
        __props__.__dict__["logical_region_id"] = logical_region_id
        __props__.__dict__["package_type"] = package_type
        __props__.__dict__["package_version"] = package_version
        __props__.__dict__["war_url"] = war_url
        return Application(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> pulumi.Output[str]:
        """
        Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
        """
        return pulumi.get(self, "application_name")

    @property
    @pulumi.getter(name="buildPackId")
    def build_pack_id(self) -> pulumi.Output[Optional[int]]:
        """
        The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
        """
        return pulumi.get(self, "build_pack_id")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def descriotion(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the application that you want to create.
        """
        return pulumi.get(self, "descriotion")

    @property
    @pulumi.getter(name="ecuInfos")
    def ecu_infos(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
        """
        return pulumi.get(self, "ecu_infos")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="healthCheckUrl")
    def health_check_url(self) -> pulumi.Output[Optional[str]]:
        """
        The URL for health checking of the application.
        """
        return pulumi.get(self, "health_check_url")

    @property
    @pulumi.getter(name="logicalRegionId")
    def logical_region_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
        """
        return pulumi.get(self, "logical_region_id")

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> pulumi.Output[str]:
        """
        The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
        """
        return pulumi.get(self, "package_type")

    @property
    @pulumi.getter(name="packageVersion")
    def package_version(self) -> pulumi.Output[Optional[str]]:
        """
        The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        """
        return pulumi.get(self, "package_version")

    @property
    @pulumi.getter(name="warUrl")
    def war_url(self) -> pulumi.Output[Optional[str]]:
        """
        The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        """
        return pulumi.get(self, "war_url")

