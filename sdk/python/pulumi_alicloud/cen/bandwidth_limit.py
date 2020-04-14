# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class BandwidthLimit(pulumi.CustomResource):
    bandwidth_limit: pulumi.Output[float]
    """
    The bandwidth configured for the interconnected regions communication.
    """
    instance_id: pulumi.Output[str]
    """
    The ID of the CEN.
    """
    region_ids: pulumi.Output[list]
    """
    List of the two regions to interconnect. Must be two different regions.
    """
    def __init__(__self__, resource_name, opts=None, bandwidth_limit=None, instance_id=None, region_ids=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a CEN cross-regional interconnection bandwidth resource. To connect networks in different regions, you must set cross-region interconnection bandwidth after buying a bandwidth package. The total bandwidth set for all the interconnected regions of a bandwidth package cannot exceed the bandwidth of the bandwidth package. By default, 1 Kbps bandwidth is provided for connectivity test. To run normal business, you must buy a bandwidth package and set a proper interconnection bandwidth.

        For example, a CEN instance is bound to a bandwidth package of 20 Mbps and  the interconnection areas are Mainland China and North America. You can set the cross-region interconnection bandwidth between US West 1 and China East 1, China East 2, China South 1, and so on. However, the total bandwidth set for all the interconnected regions cannot exceed 20  Mbps.

        For information about CEN and how to use it, see [Cross-region interconnection bandwidth](https://www.alibabacloud.com/help/doc-detail/65983.htm)



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] bandwidth_limit: The bandwidth configured for the interconnected regions communication.
        :param pulumi.Input[str] instance_id: The ID of the CEN.
        :param pulumi.Input[list] region_ids: List of the two regions to interconnect. Must be two different regions.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if bandwidth_limit is None:
                raise TypeError("Missing required property 'bandwidth_limit'")
            __props__['bandwidth_limit'] = bandwidth_limit
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            if region_ids is None:
                raise TypeError("Missing required property 'region_ids'")
            __props__['region_ids'] = region_ids
        super(BandwidthLimit, __self__).__init__(
            'alicloud:cen/bandwidthLimit:BandwidthLimit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bandwidth_limit=None, instance_id=None, region_ids=None):
        """
        Get an existing BandwidthLimit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] bandwidth_limit: The bandwidth configured for the interconnected regions communication.
        :param pulumi.Input[str] instance_id: The ID of the CEN.
        :param pulumi.Input[list] region_ids: List of the two regions to interconnect. Must be two different regions.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bandwidth_limit"] = bandwidth_limit
        __props__["instance_id"] = instance_id
        __props__["region_ids"] = region_ids
        return BandwidthLimit(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

