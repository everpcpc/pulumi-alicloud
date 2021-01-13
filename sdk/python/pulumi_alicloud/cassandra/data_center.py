# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['DataCenter']


class DataCenter(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 auto_renew_period: Optional[pulumi.Input[int]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 data_center_name: Optional[pulumi.Input[str]] = None,
                 disk_size: Optional[pulumi.Input[int]] = None,
                 disk_type: Optional[pulumi.Input[str]] = None,
                 enable_public: Optional[pulumi.Input[bool]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 node_count: Optional[pulumi.Input[int]] = None,
                 pay_type: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 period_unit: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Cassandra dataCenter resource supports replica set dataCenters only. The Cassandra provides stable, reliable, and automatic scalable database services.
        It offers a full range of database solutions, such as disaster recovery, backup, recovery, monitoring, and alarms.
        You can see detail product introduction [here](https://www.alibabacloud.com/help/product/49055.htm).

        > **NOTE:**  Available in 1.88.0+.

        > **NOTE:**  Create a cassandra dataCenter need a clusterId,so need create a cassandra cluster first.

        > **NOTE:**  The following regions support create Vpc network Cassandra cluster.
        The official website mark  more regions. Or you can call [DescribeRegions](https://help.aliyun.com/document_detail/157540.html).

        > **NOTE:**  Create Cassandra dataCenter or change dataCenter type and storage would cost 30 minutes. Please make full preparation.

        ## Example Usage
        ### Create a cassandra dataCenter

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_cluster = alicloud.cassandra.Cluster("defaultCluster",
            cluster_name="cassandra-cluster-name-tf",
            data_center_name="dc-1",
            auto_renew=False,
            instance_type="cassandra.c.large",
            major_version="3.11",
            node_count=2,
            pay_type="PayAsYouGo",
            vswitch_id="vsw-xxxx1",
            disk_size=160,
            disk_type="cloud_ssd",
            maintain_start_time="18:00Z",
            maintain_end_time="20:00Z",
            ip_white="127.0.0.1")
        default_data_center = alicloud.cassandra.DataCenter("defaultDataCenter",
            cluster_id=default_cluster.id,
            data_center_name="dc-2",
            auto_renew=False,
            instance_type="cassandra.c.large",
            node_count=2,
            pay_type="PayAsYouGo",
            vswitch_id="vsw-xxxx2",
            disk_size=160,
            disk_type="cloud_ssd")
        ```

        This is a example for class netType dataCenter. You can find more detail with the examples/cassandra_data_center dir.

        ## Import

        If you need full function, please import Cassandra cluster first. Cassandra dataCenter can be imported using the dcId:clusterId, e.g.

        ```sh
         $ pulumi import alicloud:cassandra/dataCenter:DataCenter dc_2 cn-shenxxxx-x:cds-wz933ryoaurxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renew: Auto renew of dataCenter-2,`true` or `false`. System default to `false`, valid when pay_type = Subscription.
        :param pulumi.Input[int] auto_renew_period: Period of dataCenter-2 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when pay_type = Subscription. Unit: month.
        :param pulumi.Input[str] cluster_id: Cassandra cluster id of dataCenter-2 belongs to.
        :param pulumi.Input[str] data_center_name: Cassandra dataCenter-2 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
        :param pulumi.Input[int] disk_size: User-defined Cassandra dataCenter one core node's storage space.Unit: GB. Value range:
               - Custom storage space; value range: [160, 2000].
               - 80-GB increments.
        :param pulumi.Input[str] disk_type: The disk type of Cassandra dataCenter-2. Valid values are `cloud_ssd`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`, local_disk size is fixed.
        :param pulumi.Input[str] instance_type: Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/157445.html). Or you can call describeInstanceType api.
        :param pulumi.Input[int] node_count: The node count of Cassandra dataCenter-2, default to 2.
        :param pulumi.Input[str] pay_type: The pay type of Cassandra dataCenter-2. Valid values are `Subscription`, `PayAsYouGo`. System default to `PayAsYouGo`.
        :param pulumi.Input[str] vswitch_id: The vswitch_id of dataCenter-2, mast different of vswitch_id(dc-1), can not empty.
        :param pulumi.Input[str] zone_id: The Zone to launch the Cassandra dataCenter-2. If vswitch_id is not empty, this zone_id can be "" or consistent.
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

            __props__['auto_renew'] = auto_renew
            __props__['auto_renew_period'] = auto_renew_period
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__['cluster_id'] = cluster_id
            __props__['data_center_name'] = data_center_name
            __props__['disk_size'] = disk_size
            __props__['disk_type'] = disk_type
            __props__['enable_public'] = enable_public
            if instance_type is None and not opts.urn:
                raise TypeError("Missing required property 'instance_type'")
            __props__['instance_type'] = instance_type
            if node_count is None and not opts.urn:
                raise TypeError("Missing required property 'node_count'")
            __props__['node_count'] = node_count
            if pay_type is None and not opts.urn:
                raise TypeError("Missing required property 'pay_type'")
            __props__['pay_type'] = pay_type
            __props__['period'] = period
            __props__['period_unit'] = period_unit
            if vswitch_id is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__['vswitch_id'] = vswitch_id
            __props__['zone_id'] = zone_id
            __props__['data_center_id'] = None
            __props__['public_points'] = None
            __props__['status'] = None
        super(DataCenter, __self__).__init__(
            'alicloud:cassandra/dataCenter:DataCenter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_renew: Optional[pulumi.Input[bool]] = None,
            auto_renew_period: Optional[pulumi.Input[int]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            data_center_id: Optional[pulumi.Input[str]] = None,
            data_center_name: Optional[pulumi.Input[str]] = None,
            disk_size: Optional[pulumi.Input[int]] = None,
            disk_type: Optional[pulumi.Input[str]] = None,
            enable_public: Optional[pulumi.Input[bool]] = None,
            instance_type: Optional[pulumi.Input[str]] = None,
            node_count: Optional[pulumi.Input[int]] = None,
            pay_type: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[int]] = None,
            period_unit: Optional[pulumi.Input[str]] = None,
            public_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'DataCenter':
        """
        Get an existing DataCenter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renew: Auto renew of dataCenter-2,`true` or `false`. System default to `false`, valid when pay_type = Subscription.
        :param pulumi.Input[int] auto_renew_period: Period of dataCenter-2 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when pay_type = Subscription. Unit: month.
        :param pulumi.Input[str] cluster_id: Cassandra cluster id of dataCenter-2 belongs to.
        :param pulumi.Input[str] data_center_name: Cassandra dataCenter-2 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
        :param pulumi.Input[int] disk_size: User-defined Cassandra dataCenter one core node's storage space.Unit: GB. Value range:
               - Custom storage space; value range: [160, 2000].
               - 80-GB increments.
        :param pulumi.Input[str] disk_type: The disk type of Cassandra dataCenter-2. Valid values are `cloud_ssd`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`, local_disk size is fixed.
        :param pulumi.Input[str] instance_type: Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/157445.html). Or you can call describeInstanceType api.
        :param pulumi.Input[int] node_count: The node count of Cassandra dataCenter-2, default to 2.
        :param pulumi.Input[str] pay_type: The pay type of Cassandra dataCenter-2. Valid values are `Subscription`, `PayAsYouGo`. System default to `PayAsYouGo`.
        :param pulumi.Input[str] vswitch_id: The vswitch_id of dataCenter-2, mast different of vswitch_id(dc-1), can not empty.
        :param pulumi.Input[str] zone_id: The Zone to launch the Cassandra dataCenter-2. If vswitch_id is not empty, this zone_id can be "" or consistent.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_renew"] = auto_renew
        __props__["auto_renew_period"] = auto_renew_period
        __props__["cluster_id"] = cluster_id
        __props__["data_center_id"] = data_center_id
        __props__["data_center_name"] = data_center_name
        __props__["disk_size"] = disk_size
        __props__["disk_type"] = disk_type
        __props__["enable_public"] = enable_public
        __props__["instance_type"] = instance_type
        __props__["node_count"] = node_count
        __props__["pay_type"] = pay_type
        __props__["period"] = period
        __props__["period_unit"] = period_unit
        __props__["public_points"] = public_points
        __props__["status"] = status
        __props__["vswitch_id"] = vswitch_id
        __props__["zone_id"] = zone_id
        return DataCenter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> pulumi.Output[Optional[bool]]:
        """
        Auto renew of dataCenter-2,`true` or `false`. System default to `false`, valid when pay_type = Subscription.
        """
        return pulumi.get(self, "auto_renew")

    @property
    @pulumi.getter(name="autoRenewPeriod")
    def auto_renew_period(self) -> pulumi.Output[Optional[int]]:
        """
        Period of dataCenter-2 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when pay_type = Subscription. Unit: month.
        """
        return pulumi.get(self, "auto_renew_period")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Cassandra cluster id of dataCenter-2 belongs to.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="dataCenterId")
    def data_center_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "data_center_id")

    @property
    @pulumi.getter(name="dataCenterName")
    def data_center_name(self) -> pulumi.Output[Optional[str]]:
        """
        Cassandra dataCenter-2 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
        """
        return pulumi.get(self, "data_center_name")

    @property
    @pulumi.getter(name="diskSize")
    def disk_size(self) -> pulumi.Output[Optional[int]]:
        """
        User-defined Cassandra dataCenter one core node's storage space.Unit: GB. Value range:
        - Custom storage space; value range: [160, 2000].
        - 80-GB increments.
        """
        return pulumi.get(self, "disk_size")

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> pulumi.Output[Optional[str]]:
        """
        The disk type of Cassandra dataCenter-2. Valid values are `cloud_ssd`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`, local_disk size is fixed.
        """
        return pulumi.get(self, "disk_type")

    @property
    @pulumi.getter(name="enablePublic")
    def enable_public(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enable_public")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        """
        Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/157445.html). Or you can call describeInstanceType api.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> pulumi.Output[int]:
        """
        The node count of Cassandra dataCenter-2, default to 2.
        """
        return pulumi.get(self, "node_count")

    @property
    @pulumi.getter(name="payType")
    def pay_type(self) -> pulumi.Output[str]:
        """
        The pay type of Cassandra dataCenter-2. Valid values are `Subscription`, `PayAsYouGo`. System default to `PayAsYouGo`.
        """
        return pulumi.get(self, "pay_type")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="periodUnit")
    def period_unit(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "period_unit")

    @property
    @pulumi.getter(name="publicPoints")
    def public_points(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "public_points")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[str]:
        """
        The vswitch_id of dataCenter-2, mast different of vswitch_id(dc-1), can not empty.
        """
        return pulumi.get(self, "vswitch_id")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The Zone to launch the Cassandra dataCenter-2. If vswitch_id is not empty, this zone_id can be "" or consistent.
        """
        return pulumi.get(self, "zone_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

