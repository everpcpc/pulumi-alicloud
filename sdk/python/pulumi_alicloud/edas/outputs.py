# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetApplicationsApplicationResult',
    'GetClustersClusterResult',
    'GetDeployGroupsGroupResult',
]

@pulumi.output_type
class GetApplicationsApplicationResult(dict):
    def __init__(__self__, *,
                 app_id: str,
                 app_name: str,
                 application_type: str,
                 build_package_id: float,
                 cluster_id: str,
                 cluster_type: float,
                 create_time: float,
                 health_check_url: str,
                 instance_count: float,
                 region_id: str,
                 running_instance_count: float,
                 slb_id: str,
                 slb_ip: str,
                 slb_port: float):
        """
        :param str app_id: The ID of the application that you want to deploy.
        :param str app_name: The name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
        :param str application_type: The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
        :param float build_package_id: The package ID of Enterprise Distributed Application Service (EDAS) Container.
        :param str cluster_id: The ID of the cluster that you want to create the application.
        :param float cluster_type: The type of the cluster that you want to create. Valid values: 1: Swarm cluster. 2: ECS cluster. 3: Kubernates cluster.
        :param float create_time: Application's creation time.
        :param str health_check_url: The URL for health checking of the application.
        :param float instance_count: Number of instances.
        :param str region_id: The ID of the namespace the application belongs to.
        :param float running_instance_count: Number of running instances.
        :param str slb_id: The ID of the SLB instance that is going to be bound.
        :param str slb_ip: The IP address that is allocated to the bound SLB instance.
        :param float slb_port: The port of intranet SLB.
        """
        pulumi.set(__self__, "app_id", app_id)
        pulumi.set(__self__, "app_name", app_name)
        pulumi.set(__self__, "application_type", application_type)
        pulumi.set(__self__, "build_package_id", build_package_id)
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "cluster_type", cluster_type)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "health_check_url", health_check_url)
        pulumi.set(__self__, "instance_count", instance_count)
        pulumi.set(__self__, "region_id", region_id)
        pulumi.set(__self__, "running_instance_count", running_instance_count)
        pulumi.set(__self__, "slb_id", slb_id)
        pulumi.set(__self__, "slb_ip", slb_ip)
        pulumi.set(__self__, "slb_port", slb_port)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> str:
        """
        The ID of the application that you want to deploy.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="appName")
    def app_name(self) -> str:
        """
        The name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
        """
        return pulumi.get(self, "app_name")

    @property
    @pulumi.getter(name="applicationType")
    def application_type(self) -> str:
        """
        The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
        """
        return pulumi.get(self, "application_type")

    @property
    @pulumi.getter(name="buildPackageId")
    def build_package_id(self) -> float:
        """
        The package ID of Enterprise Distributed Application Service (EDAS) Container.
        """
        return pulumi.get(self, "build_package_id")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        The ID of the cluster that you want to create the application.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> float:
        """
        The type of the cluster that you want to create. Valid values: 1: Swarm cluster. 2: ECS cluster. 3: Kubernates cluster.
        """
        return pulumi.get(self, "cluster_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> float:
        """
        Application's creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="healthCheckUrl")
    def health_check_url(self) -> str:
        """
        The URL for health checking of the application.
        """
        return pulumi.get(self, "health_check_url")

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> float:
        """
        Number of instances.
        """
        return pulumi.get(self, "instance_count")

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> str:
        """
        The ID of the namespace the application belongs to.
        """
        return pulumi.get(self, "region_id")

    @property
    @pulumi.getter(name="runningInstanceCount")
    def running_instance_count(self) -> float:
        """
        Number of running instances.
        """
        return pulumi.get(self, "running_instance_count")

    @property
    @pulumi.getter(name="slbId")
    def slb_id(self) -> str:
        """
        The ID of the SLB instance that is going to be bound.
        """
        return pulumi.get(self, "slb_id")

    @property
    @pulumi.getter(name="slbIp")
    def slb_ip(self) -> str:
        """
        The IP address that is allocated to the bound SLB instance.
        """
        return pulumi.get(self, "slb_ip")

    @property
    @pulumi.getter(name="slbPort")
    def slb_port(self) -> float:
        """
        The port of intranet SLB.
        """
        return pulumi.get(self, "slb_port")


@pulumi.output_type
class GetClustersClusterResult(dict):
    def __init__(__self__, *,
                 cluster_id: str,
                 cluster_name: str,
                 cluster_type: float,
                 cpu: float,
                 cpu_used: float,
                 create_time: float,
                 mem: float,
                 mem_used: float,
                 network_mode: float,
                 node_num: float,
                 region_id: str,
                 update_time: float,
                 vpc_id: str):
        """
        :param str cluster_id: The ID of the cluster that you want to create the application.
        :param str cluster_name: The name of the cluster.
        :param float cluster_type: The type of the cluster, Valid values: 1: Swarm cluster. 2: ECS cluster. 3: Kubernates cluster.
        :param float cpu: The total number of CPUs in the cluster.
        :param float cpu_used: The number of used CPUs in the cluster.
        :param float create_time: Cluster's creation time.
        :param float mem: The total amount of memory in the cluser. Unit: MB.
        :param float mem_used: The amount of used memory in the cluser. Unit: MB.
        :param float network_mode: The network type of the cluster. Valid values: 1: classic network. 2: VPC.
        :param float node_num: The number of the Elastic Compute Service (ECS) instances that are deployed to the cluster.
        :param str region_id: The ID of the namespace the application belongs to.
        :param float update_time: The time when the cluster was last updated.
        :param str vpc_id: The ID of the Virtual Private Cloud (VPC) for the cluster.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "cluster_type", cluster_type)
        pulumi.set(__self__, "cpu", cpu)
        pulumi.set(__self__, "cpu_used", cpu_used)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "mem", mem)
        pulumi.set(__self__, "mem_used", mem_used)
        pulumi.set(__self__, "network_mode", network_mode)
        pulumi.set(__self__, "node_num", node_num)
        pulumi.set(__self__, "region_id", region_id)
        pulumi.set(__self__, "update_time", update_time)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        The ID of the cluster that you want to create the application.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> str:
        """
        The name of the cluster.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> float:
        """
        The type of the cluster, Valid values: 1: Swarm cluster. 2: ECS cluster. 3: Kubernates cluster.
        """
        return pulumi.get(self, "cluster_type")

    @property
    @pulumi.getter
    def cpu(self) -> float:
        """
        The total number of CPUs in the cluster.
        """
        return pulumi.get(self, "cpu")

    @property
    @pulumi.getter(name="cpuUsed")
    def cpu_used(self) -> float:
        """
        The number of used CPUs in the cluster.
        """
        return pulumi.get(self, "cpu_used")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> float:
        """
        Cluster's creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def mem(self) -> float:
        """
        The total amount of memory in the cluser. Unit: MB.
        """
        return pulumi.get(self, "mem")

    @property
    @pulumi.getter(name="memUsed")
    def mem_used(self) -> float:
        """
        The amount of used memory in the cluser. Unit: MB.
        """
        return pulumi.get(self, "mem_used")

    @property
    @pulumi.getter(name="networkMode")
    def network_mode(self) -> float:
        """
        The network type of the cluster. Valid values: 1: classic network. 2: VPC.
        """
        return pulumi.get(self, "network_mode")

    @property
    @pulumi.getter(name="nodeNum")
    def node_num(self) -> float:
        """
        The number of the Elastic Compute Service (ECS) instances that are deployed to the cluster.
        """
        return pulumi.get(self, "node_num")

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> str:
        """
        The ID of the namespace the application belongs to.
        """
        return pulumi.get(self, "region_id")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> float:
        """
        The time when the cluster was last updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The ID of the Virtual Private Cloud (VPC) for the cluster.
        """
        return pulumi.get(self, "vpc_id")


@pulumi.output_type
class GetDeployGroupsGroupResult(dict):
    def __init__(__self__, *,
                 app_id: str,
                 app_version_id: str,
                 cluster_id: str,
                 create_time: float,
                 group_id: str,
                 group_name: str,
                 group_type: float,
                 package_version_id: str,
                 update_time: float):
        """
        :param str app_id: ID of the EDAS application.
        :param str app_version_id: The version of the deployment package for the application.
        :param str cluster_id: The ID of the cluster that you want to create the application.
        :param float create_time: The time when the instance group was created.
        :param str group_id: The ID of the instance group.
        :param str group_name: The name of the instance group. The length cannot exceed 64 characters.
        :param float group_type: The type of the instance group. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.
        :param str package_version_id: The version of the deployment package for the instance group that was created.
        :param float update_time: The time when the instance group was updated.
        """
        pulumi.set(__self__, "app_id", app_id)
        pulumi.set(__self__, "app_version_id", app_version_id)
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "group_name", group_name)
        pulumi.set(__self__, "group_type", group_type)
        pulumi.set(__self__, "package_version_id", package_version_id)
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> str:
        """
        ID of the EDAS application.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="appVersionId")
    def app_version_id(self) -> str:
        """
        The version of the deployment package for the application.
        """
        return pulumi.get(self, "app_version_id")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        The ID of the cluster that you want to create the application.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> float:
        """
        The time when the instance group was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> str:
        """
        The ID of the instance group.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> str:
        """
        The name of the instance group. The length cannot exceed 64 characters.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter(name="groupType")
    def group_type(self) -> float:
        """
        The type of the instance group. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.
        """
        return pulumi.get(self, "group_type")

    @property
    @pulumi.getter(name="packageVersionId")
    def package_version_id(self) -> str:
        """
        The version of the deployment package for the instance group that was created.
        """
        return pulumi.get(self, "package_version_id")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> float:
        """
        The time when the instance group was updated.
        """
        return pulumi.get(self, "update_time")

