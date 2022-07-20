# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ClusterArgs', 'Cluster']

@pulumi.input_type
class ClusterArgs:
    def __init__(__self__, *,
                 cluster_specification: pulumi.Input[str],
                 cluster_type: pulumi.Input[str],
                 cluster_version: pulumi.Input[str],
                 instance_count: pulumi.Input[int],
                 net_type: pulumi.Input[str],
                 pub_network_flow: pulumi.Input[str],
                 acl_entry_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cluster_alias_name: Optional[pulumi.Input[str]] = None,
                 disk_type: Optional[pulumi.Input[str]] = None,
                 mse_version: Optional[pulumi.Input[str]] = None,
                 private_slb_specification: Optional[pulumi.Input[str]] = None,
                 pub_slb_specification: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Cluster resource.
        :param pulumi.Input[str] cluster_specification: The engine specification of MSE Cluster. Valid values:
               `MSE_SC_1_2_200_c`：1C2G
               `MSE_SC_2_4_200_c`：2C4G
               `MSE_SC_4_8_200_c`：4C8G
               `MSE_SC_8_16_200_c`：8C16G
        :param pulumi.Input[str] cluster_type: The type of MSE Cluster.
        :param pulumi.Input[str] cluster_version: The version of MSE Cluster.
        :param pulumi.Input[int] instance_count: The count of instance.
        :param pulumi.Input[str] net_type: The type of network. Valid values: "privatenet" and "pubnet".
        :param pulumi.Input[str] pub_network_flow: The public network bandwidth. `0` means no access to the public network.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] acl_entry_lists: The whitelist. **NOTE:** This attribute is invalid when the value of `pub_network_flow` is `0` and the value of `net_type` is `privatenet`.
        :param pulumi.Input[str] cluster_alias_name: The alias of MSE Cluster.
        :param pulumi.Input[str] disk_type: The type of Disk.
        :param pulumi.Input[str] mse_version: The version of MSE. Valid values: `mse_basic` or `mse_pro`.
        :param pulumi.Input[str] private_slb_specification: The specification of private network SLB.
        :param pulumi.Input[str] pub_slb_specification: The specification of public network SLB.
        :param pulumi.Input[str] vswitch_id: The id of VSwitch.
        """
        pulumi.set(__self__, "cluster_specification", cluster_specification)
        pulumi.set(__self__, "cluster_type", cluster_type)
        pulumi.set(__self__, "cluster_version", cluster_version)
        pulumi.set(__self__, "instance_count", instance_count)
        pulumi.set(__self__, "net_type", net_type)
        pulumi.set(__self__, "pub_network_flow", pub_network_flow)
        if acl_entry_lists is not None:
            pulumi.set(__self__, "acl_entry_lists", acl_entry_lists)
        if cluster_alias_name is not None:
            pulumi.set(__self__, "cluster_alias_name", cluster_alias_name)
        if disk_type is not None:
            pulumi.set(__self__, "disk_type", disk_type)
        if mse_version is not None:
            pulumi.set(__self__, "mse_version", mse_version)
        if private_slb_specification is not None:
            pulumi.set(__self__, "private_slb_specification", private_slb_specification)
        if pub_slb_specification is not None:
            pulumi.set(__self__, "pub_slb_specification", pub_slb_specification)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="clusterSpecification")
    def cluster_specification(self) -> pulumi.Input[str]:
        """
        The engine specification of MSE Cluster. Valid values:
        `MSE_SC_1_2_200_c`：1C2G
        `MSE_SC_2_4_200_c`：2C4G
        `MSE_SC_4_8_200_c`：4C8G
        `MSE_SC_8_16_200_c`：8C16G
        """
        return pulumi.get(self, "cluster_specification")

    @cluster_specification.setter
    def cluster_specification(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_specification", value)

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> pulumi.Input[str]:
        """
        The type of MSE Cluster.
        """
        return pulumi.get(self, "cluster_type")

    @cluster_type.setter
    def cluster_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_type", value)

    @property
    @pulumi.getter(name="clusterVersion")
    def cluster_version(self) -> pulumi.Input[str]:
        """
        The version of MSE Cluster.
        """
        return pulumi.get(self, "cluster_version")

    @cluster_version.setter
    def cluster_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_version", value)

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> pulumi.Input[int]:
        """
        The count of instance.
        """
        return pulumi.get(self, "instance_count")

    @instance_count.setter
    def instance_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "instance_count", value)

    @property
    @pulumi.getter(name="netType")
    def net_type(self) -> pulumi.Input[str]:
        """
        The type of network. Valid values: "privatenet" and "pubnet".
        """
        return pulumi.get(self, "net_type")

    @net_type.setter
    def net_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "net_type", value)

    @property
    @pulumi.getter(name="pubNetworkFlow")
    def pub_network_flow(self) -> pulumi.Input[str]:
        """
        The public network bandwidth. `0` means no access to the public network.
        """
        return pulumi.get(self, "pub_network_flow")

    @pub_network_flow.setter
    def pub_network_flow(self, value: pulumi.Input[str]):
        pulumi.set(self, "pub_network_flow", value)

    @property
    @pulumi.getter(name="aclEntryLists")
    def acl_entry_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The whitelist. **NOTE:** This attribute is invalid when the value of `pub_network_flow` is `0` and the value of `net_type` is `privatenet`.
        """
        return pulumi.get(self, "acl_entry_lists")

    @acl_entry_lists.setter
    def acl_entry_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "acl_entry_lists", value)

    @property
    @pulumi.getter(name="clusterAliasName")
    def cluster_alias_name(self) -> Optional[pulumi.Input[str]]:
        """
        The alias of MSE Cluster.
        """
        return pulumi.get(self, "cluster_alias_name")

    @cluster_alias_name.setter
    def cluster_alias_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_alias_name", value)

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of Disk.
        """
        return pulumi.get(self, "disk_type")

    @disk_type.setter
    def disk_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_type", value)

    @property
    @pulumi.getter(name="mseVersion")
    def mse_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of MSE. Valid values: `mse_basic` or `mse_pro`.
        """
        return pulumi.get(self, "mse_version")

    @mse_version.setter
    def mse_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mse_version", value)

    @property
    @pulumi.getter(name="privateSlbSpecification")
    def private_slb_specification(self) -> Optional[pulumi.Input[str]]:
        """
        The specification of private network SLB.
        """
        return pulumi.get(self, "private_slb_specification")

    @private_slb_specification.setter
    def private_slb_specification(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_slb_specification", value)

    @property
    @pulumi.getter(name="pubSlbSpecification")
    def pub_slb_specification(self) -> Optional[pulumi.Input[str]]:
        """
        The specification of public network SLB.
        """
        return pulumi.get(self, "pub_slb_specification")

    @pub_slb_specification.setter
    def pub_slb_specification(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pub_slb_specification", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of VSwitch.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)


@pulumi.input_type
class _ClusterState:
    def __init__(__self__, *,
                 acl_entry_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cluster_alias_name: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 cluster_specification: Optional[pulumi.Input[str]] = None,
                 cluster_type: Optional[pulumi.Input[str]] = None,
                 cluster_version: Optional[pulumi.Input[str]] = None,
                 disk_type: Optional[pulumi.Input[str]] = None,
                 instance_count: Optional[pulumi.Input[int]] = None,
                 mse_version: Optional[pulumi.Input[str]] = None,
                 net_type: Optional[pulumi.Input[str]] = None,
                 private_slb_specification: Optional[pulumi.Input[str]] = None,
                 pub_network_flow: Optional[pulumi.Input[str]] = None,
                 pub_slb_specification: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Cluster resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] acl_entry_lists: The whitelist. **NOTE:** This attribute is invalid when the value of `pub_network_flow` is `0` and the value of `net_type` is `privatenet`.
        :param pulumi.Input[str] cluster_alias_name: The alias of MSE Cluster.
        :param pulumi.Input[str] cluster_id: (Available in v1.162.0+)  The id of Cluster.
        :param pulumi.Input[str] cluster_specification: The engine specification of MSE Cluster. Valid values:
               `MSE_SC_1_2_200_c`：1C2G
               `MSE_SC_2_4_200_c`：2C4G
               `MSE_SC_4_8_200_c`：4C8G
               `MSE_SC_8_16_200_c`：8C16G
        :param pulumi.Input[str] cluster_type: The type of MSE Cluster.
        :param pulumi.Input[str] cluster_version: The version of MSE Cluster.
        :param pulumi.Input[str] disk_type: The type of Disk.
        :param pulumi.Input[int] instance_count: The count of instance.
        :param pulumi.Input[str] mse_version: The version of MSE. Valid values: `mse_basic` or `mse_pro`.
        :param pulumi.Input[str] net_type: The type of network. Valid values: "privatenet" and "pubnet".
        :param pulumi.Input[str] private_slb_specification: The specification of private network SLB.
        :param pulumi.Input[str] pub_network_flow: The public network bandwidth. `0` means no access to the public network.
        :param pulumi.Input[str] pub_slb_specification: The specification of public network SLB.
        :param pulumi.Input[str] status: The status of MSE Cluster.
        :param pulumi.Input[str] vswitch_id: The id of VSwitch.
        """
        if acl_entry_lists is not None:
            pulumi.set(__self__, "acl_entry_lists", acl_entry_lists)
        if cluster_alias_name is not None:
            pulumi.set(__self__, "cluster_alias_name", cluster_alias_name)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_specification is not None:
            pulumi.set(__self__, "cluster_specification", cluster_specification)
        if cluster_type is not None:
            pulumi.set(__self__, "cluster_type", cluster_type)
        if cluster_version is not None:
            pulumi.set(__self__, "cluster_version", cluster_version)
        if disk_type is not None:
            pulumi.set(__self__, "disk_type", disk_type)
        if instance_count is not None:
            pulumi.set(__self__, "instance_count", instance_count)
        if mse_version is not None:
            pulumi.set(__self__, "mse_version", mse_version)
        if net_type is not None:
            pulumi.set(__self__, "net_type", net_type)
        if private_slb_specification is not None:
            pulumi.set(__self__, "private_slb_specification", private_slb_specification)
        if pub_network_flow is not None:
            pulumi.set(__self__, "pub_network_flow", pub_network_flow)
        if pub_slb_specification is not None:
            pulumi.set(__self__, "pub_slb_specification", pub_slb_specification)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="aclEntryLists")
    def acl_entry_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The whitelist. **NOTE:** This attribute is invalid when the value of `pub_network_flow` is `0` and the value of `net_type` is `privatenet`.
        """
        return pulumi.get(self, "acl_entry_lists")

    @acl_entry_lists.setter
    def acl_entry_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "acl_entry_lists", value)

    @property
    @pulumi.getter(name="clusterAliasName")
    def cluster_alias_name(self) -> Optional[pulumi.Input[str]]:
        """
        The alias of MSE Cluster.
        """
        return pulumi.get(self, "cluster_alias_name")

    @cluster_alias_name.setter
    def cluster_alias_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_alias_name", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        (Available in v1.162.0+)  The id of Cluster.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="clusterSpecification")
    def cluster_specification(self) -> Optional[pulumi.Input[str]]:
        """
        The engine specification of MSE Cluster. Valid values:
        `MSE_SC_1_2_200_c`：1C2G
        `MSE_SC_2_4_200_c`：2C4G
        `MSE_SC_4_8_200_c`：4C8G
        `MSE_SC_8_16_200_c`：8C16G
        """
        return pulumi.get(self, "cluster_specification")

    @cluster_specification.setter
    def cluster_specification(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_specification", value)

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of MSE Cluster.
        """
        return pulumi.get(self, "cluster_type")

    @cluster_type.setter
    def cluster_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_type", value)

    @property
    @pulumi.getter(name="clusterVersion")
    def cluster_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of MSE Cluster.
        """
        return pulumi.get(self, "cluster_version")

    @cluster_version.setter
    def cluster_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_version", value)

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of Disk.
        """
        return pulumi.get(self, "disk_type")

    @disk_type.setter
    def disk_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_type", value)

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> Optional[pulumi.Input[int]]:
        """
        The count of instance.
        """
        return pulumi.get(self, "instance_count")

    @instance_count.setter
    def instance_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_count", value)

    @property
    @pulumi.getter(name="mseVersion")
    def mse_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of MSE. Valid values: `mse_basic` or `mse_pro`.
        """
        return pulumi.get(self, "mse_version")

    @mse_version.setter
    def mse_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mse_version", value)

    @property
    @pulumi.getter(name="netType")
    def net_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of network. Valid values: "privatenet" and "pubnet".
        """
        return pulumi.get(self, "net_type")

    @net_type.setter
    def net_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "net_type", value)

    @property
    @pulumi.getter(name="privateSlbSpecification")
    def private_slb_specification(self) -> Optional[pulumi.Input[str]]:
        """
        The specification of private network SLB.
        """
        return pulumi.get(self, "private_slb_specification")

    @private_slb_specification.setter
    def private_slb_specification(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_slb_specification", value)

    @property
    @pulumi.getter(name="pubNetworkFlow")
    def pub_network_flow(self) -> Optional[pulumi.Input[str]]:
        """
        The public network bandwidth. `0` means no access to the public network.
        """
        return pulumi.get(self, "pub_network_flow")

    @pub_network_flow.setter
    def pub_network_flow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pub_network_flow", value)

    @property
    @pulumi.getter(name="pubSlbSpecification")
    def pub_slb_specification(self) -> Optional[pulumi.Input[str]]:
        """
        The specification of public network SLB.
        """
        return pulumi.get(self, "pub_slb_specification")

    @pub_slb_specification.setter
    def pub_slb_specification(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pub_slb_specification", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of MSE Cluster.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of VSwitch.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)


class Cluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl_entry_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cluster_alias_name: Optional[pulumi.Input[str]] = None,
                 cluster_specification: Optional[pulumi.Input[str]] = None,
                 cluster_type: Optional[pulumi.Input[str]] = None,
                 cluster_version: Optional[pulumi.Input[str]] = None,
                 disk_type: Optional[pulumi.Input[str]] = None,
                 instance_count: Optional[pulumi.Input[int]] = None,
                 mse_version: Optional[pulumi.Input[str]] = None,
                 net_type: Optional[pulumi.Input[str]] = None,
                 private_slb_specification: Optional[pulumi.Input[str]] = None,
                 pub_network_flow: Optional[pulumi.Input[str]] = None,
                 pub_slb_specification: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a MSE Cluster resource. It is a one-stop microservice platform for the industry's mainstream open source microservice frameworks Spring Cloud and Dubbo, providing governance center, managed registry and managed configuration center.

        > **NOTE:** Available in 1.94.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.mse.Cluster("example",
            acl_entry_lists=["127.0.0.1/32"],
            cluster_alias_name="tf-testAccMseCluster",
            cluster_specification="MSE_SC_1_2_200_c",
            cluster_type="Nacos-Ans",
            cluster_version="NACOS_ANS_1_2_1",
            instance_count=1,
            net_type="privatenet",
            pub_network_flow="1",
            vswitch_id="vsw-123456")
        ```

        ## Import

        MSE Cluster can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:mse/cluster:Cluster example mse-cn-0d9xxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] acl_entry_lists: The whitelist. **NOTE:** This attribute is invalid when the value of `pub_network_flow` is `0` and the value of `net_type` is `privatenet`.
        :param pulumi.Input[str] cluster_alias_name: The alias of MSE Cluster.
        :param pulumi.Input[str] cluster_specification: The engine specification of MSE Cluster. Valid values:
               `MSE_SC_1_2_200_c`：1C2G
               `MSE_SC_2_4_200_c`：2C4G
               `MSE_SC_4_8_200_c`：4C8G
               `MSE_SC_8_16_200_c`：8C16G
        :param pulumi.Input[str] cluster_type: The type of MSE Cluster.
        :param pulumi.Input[str] cluster_version: The version of MSE Cluster.
        :param pulumi.Input[str] disk_type: The type of Disk.
        :param pulumi.Input[int] instance_count: The count of instance.
        :param pulumi.Input[str] mse_version: The version of MSE. Valid values: `mse_basic` or `mse_pro`.
        :param pulumi.Input[str] net_type: The type of network. Valid values: "privatenet" and "pubnet".
        :param pulumi.Input[str] private_slb_specification: The specification of private network SLB.
        :param pulumi.Input[str] pub_network_flow: The public network bandwidth. `0` means no access to the public network.
        :param pulumi.Input[str] pub_slb_specification: The specification of public network SLB.
        :param pulumi.Input[str] vswitch_id: The id of VSwitch.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a MSE Cluster resource. It is a one-stop microservice platform for the industry's mainstream open source microservice frameworks Spring Cloud and Dubbo, providing governance center, managed registry and managed configuration center.

        > **NOTE:** Available in 1.94.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.mse.Cluster("example",
            acl_entry_lists=["127.0.0.1/32"],
            cluster_alias_name="tf-testAccMseCluster",
            cluster_specification="MSE_SC_1_2_200_c",
            cluster_type="Nacos-Ans",
            cluster_version="NACOS_ANS_1_2_1",
            instance_count=1,
            net_type="privatenet",
            pub_network_flow="1",
            vswitch_id="vsw-123456")
        ```

        ## Import

        MSE Cluster can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:mse/cluster:Cluster example mse-cn-0d9xxxx
        ```

        :param str resource_name: The name of the resource.
        :param ClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl_entry_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cluster_alias_name: Optional[pulumi.Input[str]] = None,
                 cluster_specification: Optional[pulumi.Input[str]] = None,
                 cluster_type: Optional[pulumi.Input[str]] = None,
                 cluster_version: Optional[pulumi.Input[str]] = None,
                 disk_type: Optional[pulumi.Input[str]] = None,
                 instance_count: Optional[pulumi.Input[int]] = None,
                 mse_version: Optional[pulumi.Input[str]] = None,
                 net_type: Optional[pulumi.Input[str]] = None,
                 private_slb_specification: Optional[pulumi.Input[str]] = None,
                 pub_network_flow: Optional[pulumi.Input[str]] = None,
                 pub_slb_specification: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ClusterArgs.__new__(ClusterArgs)

            __props__.__dict__["acl_entry_lists"] = acl_entry_lists
            __props__.__dict__["cluster_alias_name"] = cluster_alias_name
            if cluster_specification is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_specification'")
            __props__.__dict__["cluster_specification"] = cluster_specification
            if cluster_type is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_type'")
            __props__.__dict__["cluster_type"] = cluster_type
            if cluster_version is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_version'")
            __props__.__dict__["cluster_version"] = cluster_version
            __props__.__dict__["disk_type"] = disk_type
            if instance_count is None and not opts.urn:
                raise TypeError("Missing required property 'instance_count'")
            __props__.__dict__["instance_count"] = instance_count
            __props__.__dict__["mse_version"] = mse_version
            if net_type is None and not opts.urn:
                raise TypeError("Missing required property 'net_type'")
            __props__.__dict__["net_type"] = net_type
            __props__.__dict__["private_slb_specification"] = private_slb_specification
            if pub_network_flow is None and not opts.urn:
                raise TypeError("Missing required property 'pub_network_flow'")
            __props__.__dict__["pub_network_flow"] = pub_network_flow
            __props__.__dict__["pub_slb_specification"] = pub_slb_specification
            __props__.__dict__["vswitch_id"] = vswitch_id
            __props__.__dict__["cluster_id"] = None
            __props__.__dict__["status"] = None
        super(Cluster, __self__).__init__(
            'alicloud:mse/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl_entry_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            cluster_alias_name: Optional[pulumi.Input[str]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            cluster_specification: Optional[pulumi.Input[str]] = None,
            cluster_type: Optional[pulumi.Input[str]] = None,
            cluster_version: Optional[pulumi.Input[str]] = None,
            disk_type: Optional[pulumi.Input[str]] = None,
            instance_count: Optional[pulumi.Input[int]] = None,
            mse_version: Optional[pulumi.Input[str]] = None,
            net_type: Optional[pulumi.Input[str]] = None,
            private_slb_specification: Optional[pulumi.Input[str]] = None,
            pub_network_flow: Optional[pulumi.Input[str]] = None,
            pub_slb_specification: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] acl_entry_lists: The whitelist. **NOTE:** This attribute is invalid when the value of `pub_network_flow` is `0` and the value of `net_type` is `privatenet`.
        :param pulumi.Input[str] cluster_alias_name: The alias of MSE Cluster.
        :param pulumi.Input[str] cluster_id: (Available in v1.162.0+)  The id of Cluster.
        :param pulumi.Input[str] cluster_specification: The engine specification of MSE Cluster. Valid values:
               `MSE_SC_1_2_200_c`：1C2G
               `MSE_SC_2_4_200_c`：2C4G
               `MSE_SC_4_8_200_c`：4C8G
               `MSE_SC_8_16_200_c`：8C16G
        :param pulumi.Input[str] cluster_type: The type of MSE Cluster.
        :param pulumi.Input[str] cluster_version: The version of MSE Cluster.
        :param pulumi.Input[str] disk_type: The type of Disk.
        :param pulumi.Input[int] instance_count: The count of instance.
        :param pulumi.Input[str] mse_version: The version of MSE. Valid values: `mse_basic` or `mse_pro`.
        :param pulumi.Input[str] net_type: The type of network. Valid values: "privatenet" and "pubnet".
        :param pulumi.Input[str] private_slb_specification: The specification of private network SLB.
        :param pulumi.Input[str] pub_network_flow: The public network bandwidth. `0` means no access to the public network.
        :param pulumi.Input[str] pub_slb_specification: The specification of public network SLB.
        :param pulumi.Input[str] status: The status of MSE Cluster.
        :param pulumi.Input[str] vswitch_id: The id of VSwitch.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClusterState.__new__(_ClusterState)

        __props__.__dict__["acl_entry_lists"] = acl_entry_lists
        __props__.__dict__["cluster_alias_name"] = cluster_alias_name
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["cluster_specification"] = cluster_specification
        __props__.__dict__["cluster_type"] = cluster_type
        __props__.__dict__["cluster_version"] = cluster_version
        __props__.__dict__["disk_type"] = disk_type
        __props__.__dict__["instance_count"] = instance_count
        __props__.__dict__["mse_version"] = mse_version
        __props__.__dict__["net_type"] = net_type
        __props__.__dict__["private_slb_specification"] = private_slb_specification
        __props__.__dict__["pub_network_flow"] = pub_network_flow
        __props__.__dict__["pub_slb_specification"] = pub_slb_specification
        __props__.__dict__["status"] = status
        __props__.__dict__["vswitch_id"] = vswitch_id
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aclEntryLists")
    def acl_entry_lists(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The whitelist. **NOTE:** This attribute is invalid when the value of `pub_network_flow` is `0` and the value of `net_type` is `privatenet`.
        """
        return pulumi.get(self, "acl_entry_lists")

    @property
    @pulumi.getter(name="clusterAliasName")
    def cluster_alias_name(self) -> pulumi.Output[Optional[str]]:
        """
        The alias of MSE Cluster.
        """
        return pulumi.get(self, "cluster_alias_name")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        (Available in v1.162.0+)  The id of Cluster.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterSpecification")
    def cluster_specification(self) -> pulumi.Output[str]:
        """
        The engine specification of MSE Cluster. Valid values:
        `MSE_SC_1_2_200_c`：1C2G
        `MSE_SC_2_4_200_c`：2C4G
        `MSE_SC_4_8_200_c`：4C8G
        `MSE_SC_8_16_200_c`：8C16G
        """
        return pulumi.get(self, "cluster_specification")

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> pulumi.Output[str]:
        """
        The type of MSE Cluster.
        """
        return pulumi.get(self, "cluster_type")

    @property
    @pulumi.getter(name="clusterVersion")
    def cluster_version(self) -> pulumi.Output[str]:
        """
        The version of MSE Cluster.
        """
        return pulumi.get(self, "cluster_version")

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of Disk.
        """
        return pulumi.get(self, "disk_type")

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> pulumi.Output[int]:
        """
        The count of instance.
        """
        return pulumi.get(self, "instance_count")

    @property
    @pulumi.getter(name="mseVersion")
    def mse_version(self) -> pulumi.Output[str]:
        """
        The version of MSE. Valid values: `mse_basic` or `mse_pro`.
        """
        return pulumi.get(self, "mse_version")

    @property
    @pulumi.getter(name="netType")
    def net_type(self) -> pulumi.Output[str]:
        """
        The type of network. Valid values: "privatenet" and "pubnet".
        """
        return pulumi.get(self, "net_type")

    @property
    @pulumi.getter(name="privateSlbSpecification")
    def private_slb_specification(self) -> pulumi.Output[Optional[str]]:
        """
        The specification of private network SLB.
        """
        return pulumi.get(self, "private_slb_specification")

    @property
    @pulumi.getter(name="pubNetworkFlow")
    def pub_network_flow(self) -> pulumi.Output[str]:
        """
        The public network bandwidth. `0` means no access to the public network.
        """
        return pulumi.get(self, "pub_network_flow")

    @property
    @pulumi.getter(name="pubSlbSpecification")
    def pub_slb_specification(self) -> pulumi.Output[Optional[str]]:
        """
        The specification of public network SLB.
        """
        return pulumi.get(self, "pub_slb_specification")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of MSE Cluster.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[Optional[str]]:
        """
        The id of VSwitch.
        """
        return pulumi.get(self, "vswitch_id")

