# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetZoneRecordsResult:
    """
    A collection of values returned by getZoneRecords.
    """
    def __init__(__self__, ids=None, keyword=None, output_file=None, records=None, zone_id=None, id=None):
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of Private Zone Record IDs.
        """
        if keyword and not isinstance(keyword, str):
            raise TypeError("Expected argument 'keyword' to be a str")
        __self__.keyword = keyword
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if records and not isinstance(records, list):
            raise TypeError("Expected argument 'records' to be a list")
        __self__.records = records
        """
        A list of zone records. Each element contains the following attributes:
        """
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        __self__.zone_id = zone_id
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetZoneRecordsResult(GetZoneRecordsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZoneRecordsResult(
            ids=self.ids,
            keyword=self.keyword,
            output_file=self.output_file,
            records=self.records,
            zone_id=self.zone_id,
            id=self.id)

def get_zone_records(ids=None,keyword=None,output_file=None,zone_id=None,opts=None):
    """
    This data source provides Private Zone Records resource information owned by an Alibaba Cloud account.
    
    :param list ids: A list of Private Zone Record IDs.
    :param str keyword: Keyword for record rr and value.
    :param str zone_id: ID of the Private Zone.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/pvtz_zone_records.html.markdown.
    """
    __args__ = dict()

    __args__['ids'] = ids
    __args__['keyword'] = keyword
    __args__['outputFile'] = output_file
    __args__['zoneId'] = zone_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:pvtz/getZoneRecords:getZoneRecords', __args__, opts=opts).value

    return AwaitableGetZoneRecordsResult(
        ids=__ret__.get('ids'),
        keyword=__ret__.get('keyword'),
        output_file=__ret__.get('outputFile'),
        records=__ret__.get('records'),
        zone_id=__ret__.get('zoneId'),
        id=__ret__.get('id'))
