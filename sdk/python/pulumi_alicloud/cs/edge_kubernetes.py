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

__all__ = ['EdgeKubernetes']


class EdgeKubernetes(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addons: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EdgeKubernetesAddonArgs']]]]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 client_cert: Optional[pulumi.Input[str]] = None,
                 client_key: Optional[pulumi.Input[str]] = None,
                 cluster_ca_cert: Optional[pulumi.Input[str]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 install_cloud_monitor: Optional[pulumi.Input[bool]] = None,
                 is_enterprise_security_group: Optional[pulumi.Input[bool]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 kube_config: Optional[pulumi.Input[str]] = None,
                 log_config: Optional[pulumi.Input[pulumi.InputType['EdgeKubernetesLogConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 new_nat_gateway: Optional[pulumi.Input[bool]] = None,
                 node_cidr_mask: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 pod_cidr: Optional[pulumi.Input[str]] = None,
                 proxy_mode: Optional[pulumi.Input[str]] = None,
                 rds_instances: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 service_cidr: Optional[pulumi.Input[str]] = None,
                 slb_internet_enabled: Optional[pulumi.Input[bool]] = None,
                 user_data: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 worker_data_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EdgeKubernetesWorkerDataDiskArgs']]]]] = None,
                 worker_disk_category: Optional[pulumi.Input[str]] = None,
                 worker_disk_size: Optional[pulumi.Input[int]] = None,
                 worker_instance_charge_type: Optional[pulumi.Input[str]] = None,
                 worker_instance_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 worker_number: Optional[pulumi.Input[int]] = None,
                 worker_vswitch_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        Kubernetes cluster can be imported using the id, e.g. Then complete the main.tf accords to the result of `terraform plan`

        ```sh
         $ pulumi import alicloud:cs/edgeKubernetes:EdgeKubernetes alicloud_cs_edge_kubernetes.main cluster-id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: The ID of availability zone.
        :param pulumi.Input[str] client_cert: The path of client certificate, like `~/.kube/client-cert.pem`.
        :param pulumi.Input[str] client_key: The path of client key, like `~/.kube/client-key.pem`.
        :param pulumi.Input[str] cluster_ca_cert: The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
        :param pulumi.Input[bool] deletion_protection: Whether to enable cluster deletion protection.
        :param pulumi.Input[bool] install_cloud_monitor: Install cloud monitor agent on ECS. default: `true`.
        :param pulumi.Input[bool] is_enterprise_security_group: Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
        :param pulumi.Input[str] key_name: The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        :param pulumi.Input[str] kube_config: The path of kube config, like `~/.kube/config`.
        :param pulumi.Input[str] name: The kubernetes cluster's name. It is unique in one Alicloud account.
        :param pulumi.Input[bool] new_nat_gateway: Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
        :param pulumi.Input[int] node_cidr_mask: The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
        :param pulumi.Input[str] password: The password of ssh login cluster node. You have to specify one of `password`, `key_name` `kms_encrypted_password` fields.
        :param pulumi.Input[str] pod_cidr: [Flannel Specific] The CIDR block for the pod network when using Flannel.
        :param pulumi.Input[str] proxy_mode: Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
        :param pulumi.Input[str] security_group_id: The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
        :param pulumi.Input[str] service_cidr: The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
        :param pulumi.Input[bool] slb_internet_enabled: Whether to create internet load balancer for API Server. Default to true.
        :param pulumi.Input[str] user_data: Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
        :param pulumi.Input[str] version: Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EdgeKubernetesWorkerDataDiskArgs']]]] worker_data_disks: The data disk configurations of worker nodes, such as the disk type and disk size.
               * `category`: the type of the data disks. Valid values:
               * cloud : basic disks.
               * cloud_efficiency : ultra disks.
               * cloud_ssd : SSDs.
               * cloud_essd : ESSDs.
               * `size`: the size of a data disk, at least 40. Unit: GiB.
               * `encrypted`: specifies whether to encrypt data disks. Valid values: true and false.
        :param pulumi.Input[str] worker_disk_category: The system disk category of worker node. Its valid value are `cloud_efficiency`, `cloud_ssd` and `cloud_essd` and . Default to `cloud_efficiency`.
        :param pulumi.Input[int] worker_disk_size: The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] worker_instance_types: The instance types of worker node, you can set multiple types to avoid NoStock of a certain type
        :param pulumi.Input[int] worker_number: The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
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

            __props__['addons'] = addons
            __props__['availability_zone'] = availability_zone
            __props__['client_cert'] = client_cert
            __props__['client_key'] = client_key
            __props__['cluster_ca_cert'] = cluster_ca_cert
            __props__['deletion_protection'] = deletion_protection
            __props__['install_cloud_monitor'] = install_cloud_monitor
            __props__['is_enterprise_security_group'] = is_enterprise_security_group
            __props__['key_name'] = key_name
            __props__['kube_config'] = kube_config
            __props__['log_config'] = log_config
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            __props__['new_nat_gateway'] = new_nat_gateway
            __props__['node_cidr_mask'] = node_cidr_mask
            __props__['password'] = password
            __props__['pod_cidr'] = pod_cidr
            __props__['proxy_mode'] = proxy_mode
            __props__['rds_instances'] = rds_instances
            __props__['resource_group_id'] = resource_group_id
            __props__['security_group_id'] = security_group_id
            __props__['service_cidr'] = service_cidr
            __props__['slb_internet_enabled'] = slb_internet_enabled
            __props__['user_data'] = user_data
            __props__['version'] = version
            __props__['worker_data_disks'] = worker_data_disks
            __props__['worker_disk_category'] = worker_disk_category
            __props__['worker_disk_size'] = worker_disk_size
            __props__['worker_instance_charge_type'] = worker_instance_charge_type
            if worker_instance_types is None:
                raise TypeError("Missing required property 'worker_instance_types'")
            __props__['worker_instance_types'] = worker_instance_types
            if worker_number is None:
                raise TypeError("Missing required property 'worker_number'")
            __props__['worker_number'] = worker_number
            if worker_vswitch_ids is None:
                raise TypeError("Missing required property 'worker_vswitch_ids'")
            __props__['worker_vswitch_ids'] = worker_vswitch_ids
            __props__['certificate_authority'] = None
            __props__['connections'] = None
            __props__['nat_gateway_id'] = None
            __props__['slb_internet'] = None
            __props__['slb_intranet'] = None
            __props__['vpc_id'] = None
            __props__['worker_nodes'] = None
        super(EdgeKubernetes, __self__).__init__(
            'alicloud:cs/edgeKubernetes:EdgeKubernetes',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            addons: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EdgeKubernetesAddonArgs']]]]] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            certificate_authority: Optional[pulumi.Input[pulumi.InputType['EdgeKubernetesCertificateAuthorityArgs']]] = None,
            client_cert: Optional[pulumi.Input[str]] = None,
            client_key: Optional[pulumi.Input[str]] = None,
            cluster_ca_cert: Optional[pulumi.Input[str]] = None,
            connections: Optional[pulumi.Input[pulumi.InputType['EdgeKubernetesConnectionsArgs']]] = None,
            deletion_protection: Optional[pulumi.Input[bool]] = None,
            install_cloud_monitor: Optional[pulumi.Input[bool]] = None,
            is_enterprise_security_group: Optional[pulumi.Input[bool]] = None,
            key_name: Optional[pulumi.Input[str]] = None,
            kube_config: Optional[pulumi.Input[str]] = None,
            log_config: Optional[pulumi.Input[pulumi.InputType['EdgeKubernetesLogConfigArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            nat_gateway_id: Optional[pulumi.Input[str]] = None,
            new_nat_gateway: Optional[pulumi.Input[bool]] = None,
            node_cidr_mask: Optional[pulumi.Input[int]] = None,
            password: Optional[pulumi.Input[str]] = None,
            pod_cidr: Optional[pulumi.Input[str]] = None,
            proxy_mode: Optional[pulumi.Input[str]] = None,
            rds_instances: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None,
            service_cidr: Optional[pulumi.Input[str]] = None,
            slb_internet: Optional[pulumi.Input[str]] = None,
            slb_internet_enabled: Optional[pulumi.Input[bool]] = None,
            slb_intranet: Optional[pulumi.Input[str]] = None,
            user_data: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            worker_data_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EdgeKubernetesWorkerDataDiskArgs']]]]] = None,
            worker_disk_category: Optional[pulumi.Input[str]] = None,
            worker_disk_size: Optional[pulumi.Input[int]] = None,
            worker_instance_charge_type: Optional[pulumi.Input[str]] = None,
            worker_instance_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            worker_nodes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EdgeKubernetesWorkerNodeArgs']]]]] = None,
            worker_number: Optional[pulumi.Input[int]] = None,
            worker_vswitch_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'EdgeKubernetes':
        """
        Get an existing EdgeKubernetes resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: The ID of availability zone.
        :param pulumi.Input[pulumi.InputType['EdgeKubernetesCertificateAuthorityArgs']] certificate_authority: (Available in 1.105.0+) Nested attribute containing certificate authority data for your cluster.
        :param pulumi.Input[str] client_cert: The path of client certificate, like `~/.kube/client-cert.pem`.
        :param pulumi.Input[str] client_key: The path of client key, like `~/.kube/client-key.pem`.
        :param pulumi.Input[str] cluster_ca_cert: The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
        :param pulumi.Input[bool] deletion_protection: Whether to enable cluster deletion protection.
        :param pulumi.Input[bool] install_cloud_monitor: Install cloud monitor agent on ECS. default: `true`.
        :param pulumi.Input[bool] is_enterprise_security_group: Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
        :param pulumi.Input[str] key_name: The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        :param pulumi.Input[str] kube_config: The path of kube config, like `~/.kube/config`.
        :param pulumi.Input[str] name: The kubernetes cluster's name. It is unique in one Alicloud account.
        :param pulumi.Input[str] nat_gateway_id: The ID of nat gateway used to launch kubernetes cluster.
        :param pulumi.Input[bool] new_nat_gateway: Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
        :param pulumi.Input[int] node_cidr_mask: The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
        :param pulumi.Input[str] password: The password of ssh login cluster node. You have to specify one of `password`, `key_name` `kms_encrypted_password` fields.
        :param pulumi.Input[str] pod_cidr: [Flannel Specific] The CIDR block for the pod network when using Flannel.
        :param pulumi.Input[str] proxy_mode: Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
        :param pulumi.Input[str] security_group_id: The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
        :param pulumi.Input[str] service_cidr: The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
        :param pulumi.Input[bool] slb_internet_enabled: Whether to create internet load balancer for API Server. Default to true.
        :param pulumi.Input[str] slb_intranet: The ID of private load balancer where the current cluster master node is located.
        :param pulumi.Input[str] user_data: Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
        :param pulumi.Input[str] version: Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
        :param pulumi.Input[str] vpc_id: The ID of VPC where the current cluster is located.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EdgeKubernetesWorkerDataDiskArgs']]]] worker_data_disks: The data disk configurations of worker nodes, such as the disk type and disk size.
               * `category`: the type of the data disks. Valid values:
               * cloud : basic disks.
               * cloud_efficiency : ultra disks.
               * cloud_ssd : SSDs.
               * cloud_essd : ESSDs.
               * `size`: the size of a data disk, at least 40. Unit: GiB.
               * `encrypted`: specifies whether to encrypt data disks. Valid values: true and false.
        :param pulumi.Input[str] worker_disk_category: The system disk category of worker node. Its valid value are `cloud_efficiency`, `cloud_ssd` and `cloud_essd` and . Default to `cloud_efficiency`.
        :param pulumi.Input[int] worker_disk_size: The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] worker_instance_types: The instance types of worker node, you can set multiple types to avoid NoStock of a certain type
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EdgeKubernetesWorkerNodeArgs']]]] worker_nodes: List of cluster worker nodes.
        :param pulumi.Input[int] worker_number: The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["addons"] = addons
        __props__["availability_zone"] = availability_zone
        __props__["certificate_authority"] = certificate_authority
        __props__["client_cert"] = client_cert
        __props__["client_key"] = client_key
        __props__["cluster_ca_cert"] = cluster_ca_cert
        __props__["connections"] = connections
        __props__["deletion_protection"] = deletion_protection
        __props__["install_cloud_monitor"] = install_cloud_monitor
        __props__["is_enterprise_security_group"] = is_enterprise_security_group
        __props__["key_name"] = key_name
        __props__["kube_config"] = kube_config
        __props__["log_config"] = log_config
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["nat_gateway_id"] = nat_gateway_id
        __props__["new_nat_gateway"] = new_nat_gateway
        __props__["node_cidr_mask"] = node_cidr_mask
        __props__["password"] = password
        __props__["pod_cidr"] = pod_cidr
        __props__["proxy_mode"] = proxy_mode
        __props__["rds_instances"] = rds_instances
        __props__["resource_group_id"] = resource_group_id
        __props__["security_group_id"] = security_group_id
        __props__["service_cidr"] = service_cidr
        __props__["slb_internet"] = slb_internet
        __props__["slb_internet_enabled"] = slb_internet_enabled
        __props__["slb_intranet"] = slb_intranet
        __props__["user_data"] = user_data
        __props__["version"] = version
        __props__["vpc_id"] = vpc_id
        __props__["worker_data_disks"] = worker_data_disks
        __props__["worker_disk_category"] = worker_disk_category
        __props__["worker_disk_size"] = worker_disk_size
        __props__["worker_instance_charge_type"] = worker_instance_charge_type
        __props__["worker_instance_types"] = worker_instance_types
        __props__["worker_nodes"] = worker_nodes
        __props__["worker_number"] = worker_number
        __props__["worker_vswitch_ids"] = worker_vswitch_ids
        return EdgeKubernetes(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def addons(self) -> pulumi.Output[Optional[Sequence['outputs.EdgeKubernetesAddon']]]:
        return pulumi.get(self, "addons")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        """
        The ID of availability zone.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="certificateAuthority")
    def certificate_authority(self) -> pulumi.Output['outputs.EdgeKubernetesCertificateAuthority']:
        """
        (Available in 1.105.0+) Nested attribute containing certificate authority data for your cluster.
        """
        return pulumi.get(self, "certificate_authority")

    @property
    @pulumi.getter(name="clientCert")
    def client_cert(self) -> pulumi.Output[Optional[str]]:
        """
        The path of client certificate, like `~/.kube/client-cert.pem`.
        """
        return pulumi.get(self, "client_cert")

    @property
    @pulumi.getter(name="clientKey")
    def client_key(self) -> pulumi.Output[Optional[str]]:
        """
        The path of client key, like `~/.kube/client-key.pem`.
        """
        return pulumi.get(self, "client_key")

    @property
    @pulumi.getter(name="clusterCaCert")
    def cluster_ca_cert(self) -> pulumi.Output[Optional[str]]:
        """
        The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
        """
        return pulumi.get(self, "cluster_ca_cert")

    @property
    @pulumi.getter
    def connections(self) -> pulumi.Output['outputs.EdgeKubernetesConnections']:
        return pulumi.get(self, "connections")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable cluster deletion protection.
        """
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter(name="installCloudMonitor")
    def install_cloud_monitor(self) -> pulumi.Output[Optional[bool]]:
        """
        Install cloud monitor agent on ECS. default: `true`.
        """
        return pulumi.get(self, "install_cloud_monitor")

    @property
    @pulumi.getter(name="isEnterpriseSecurityGroup")
    def is_enterprise_security_group(self) -> pulumi.Output[bool]:
        """
        Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
        """
        return pulumi.get(self, "is_enterprise_security_group")

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> pulumi.Output[Optional[str]]:
        """
        The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter(name="kubeConfig")
    def kube_config(self) -> pulumi.Output[Optional[str]]:
        """
        The path of kube config, like `~/.kube/config`.
        """
        return pulumi.get(self, "kube_config")

    @property
    @pulumi.getter(name="logConfig")
    def log_config(self) -> pulumi.Output[Optional['outputs.EdgeKubernetesLogConfig']]:
        return pulumi.get(self, "log_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The kubernetes cluster's name. It is unique in one Alicloud account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> pulumi.Output[str]:
        """
        The ID of nat gateway used to launch kubernetes cluster.
        """
        return pulumi.get(self, "nat_gateway_id")

    @property
    @pulumi.getter(name="newNatGateway")
    def new_nat_gateway(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
        """
        return pulumi.get(self, "new_nat_gateway")

    @property
    @pulumi.getter(name="nodeCidrMask")
    def node_cidr_mask(self) -> pulumi.Output[Optional[int]]:
        """
        The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
        """
        return pulumi.get(self, "node_cidr_mask")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password of ssh login cluster node. You have to specify one of `password`, `key_name` `kms_encrypted_password` fields.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="podCidr")
    def pod_cidr(self) -> pulumi.Output[Optional[str]]:
        """
        [Flannel Specific] The CIDR block for the pod network when using Flannel.
        """
        return pulumi.get(self, "pod_cidr")

    @property
    @pulumi.getter(name="proxyMode")
    def proxy_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
        """
        return pulumi.get(self, "proxy_mode")

    @property
    @pulumi.getter(name="rdsInstances")
    def rds_instances(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "rds_instances")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
        """
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="serviceCidr")
    def service_cidr(self) -> pulumi.Output[Optional[str]]:
        """
        The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
        """
        return pulumi.get(self, "service_cidr")

    @property
    @pulumi.getter(name="slbInternet")
    def slb_internet(self) -> pulumi.Output[str]:
        return pulumi.get(self, "slb_internet")

    @property
    @pulumi.getter(name="slbInternetEnabled")
    def slb_internet_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to create internet load balancer for API Server. Default to true.
        """
        return pulumi.get(self, "slb_internet_enabled")

    @property
    @pulumi.getter(name="slbIntranet")
    def slb_intranet(self) -> pulumi.Output[str]:
        """
        The ID of private load balancer where the current cluster master node is located.
        """
        return pulumi.get(self, "slb_intranet")

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> pulumi.Output[Optional[str]]:
        """
        Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
        """
        return pulumi.get(self, "user_data")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of VPC where the current cluster is located.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="workerDataDisks")
    def worker_data_disks(self) -> pulumi.Output[Optional[Sequence['outputs.EdgeKubernetesWorkerDataDisk']]]:
        """
        The data disk configurations of worker nodes, such as the disk type and disk size.
        * `category`: the type of the data disks. Valid values:
        * cloud : basic disks.
        * cloud_efficiency : ultra disks.
        * cloud_ssd : SSDs.
        * cloud_essd : ESSDs.
        * `size`: the size of a data disk, at least 40. Unit: GiB.
        * `encrypted`: specifies whether to encrypt data disks. Valid values: true and false.
        """
        return pulumi.get(self, "worker_data_disks")

    @property
    @pulumi.getter(name="workerDiskCategory")
    def worker_disk_category(self) -> pulumi.Output[Optional[str]]:
        """
        The system disk category of worker node. Its valid value are `cloud_efficiency`, `cloud_ssd` and `cloud_essd` and . Default to `cloud_efficiency`.
        """
        return pulumi.get(self, "worker_disk_category")

    @property
    @pulumi.getter(name="workerDiskSize")
    def worker_disk_size(self) -> pulumi.Output[Optional[int]]:
        """
        The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
        """
        return pulumi.get(self, "worker_disk_size")

    @property
    @pulumi.getter(name="workerInstanceChargeType")
    def worker_instance_charge_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "worker_instance_charge_type")

    @property
    @pulumi.getter(name="workerInstanceTypes")
    def worker_instance_types(self) -> pulumi.Output[Sequence[str]]:
        """
        The instance types of worker node, you can set multiple types to avoid NoStock of a certain type
        """
        return pulumi.get(self, "worker_instance_types")

    @property
    @pulumi.getter(name="workerNodes")
    def worker_nodes(self) -> pulumi.Output[Sequence['outputs.EdgeKubernetesWorkerNode']]:
        """
        List of cluster worker nodes.
        """
        return pulumi.get(self, "worker_nodes")

    @property
    @pulumi.getter(name="workerNumber")
    def worker_number(self) -> pulumi.Output[int]:
        """
        The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
        """
        return pulumi.get(self, "worker_number")

    @property
    @pulumi.getter(name="workerVswitchIds")
    def worker_vswitch_ids(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "worker_vswitch_ids")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
