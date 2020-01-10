# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetRegionsResult:
    """
    A collection of values returned by getRegions.
    """
    def __init__(__self__, current=None, ids=None, name=None, output_file=None, regions=None, id=None):
        if current and not isinstance(current, bool):
            raise TypeError("Expected argument 'current' to be a bool")
        __self__.current = current
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of region IDs.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if regions and not isinstance(regions, list):
            raise TypeError("Expected argument 'regions' to be a list")
        __self__.regions = regions
        """
        A list of regions. Each element contains the following attributes:
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetRegionsResult(GetRegionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionsResult(
            current=self.current,
            ids=self.ids,
            name=self.name,
            output_file=self.output_file,
            regions=self.regions,
            id=self.id)

def get_regions(current=None,name=None,output_file=None,opts=None):
    """
    This data source provides Alibaba Cloud regions.
    
    :param bool current: Set to true to match only the region configured in the provider.
    :param str name: The name of the region to select, such as `eu-central-1`.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/regions.html.markdown.
    """
    __args__ = dict()

    __args__['current'] = current
    __args__['name'] = name
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:index/getRegions:getRegions', __args__, opts=opts).value

    return AwaitableGetRegionsResult(
        current=__ret__.get('current'),
        ids=__ret__.get('ids'),
        name=__ret__.get('name'),
        output_file=__ret__.get('outputFile'),
        regions=__ret__.get('regions'),
        id=__ret__.get('id'))
