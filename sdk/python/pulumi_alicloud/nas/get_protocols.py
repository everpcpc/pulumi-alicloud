# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetProtocolsResult:
    """
    A collection of values returned by getProtocols.
    """
    def __init__(__self__, id=None, output_file=None, protocols=None, type=None, zone_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if protocols and not isinstance(protocols, list):
            raise TypeError("Expected argument 'protocols' to be a list")
        __self__.protocols = protocols
        """
        A list of supported protocol type..
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        __self__.zone_id = zone_id
class AwaitableGetProtocolsResult(GetProtocolsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProtocolsResult(
            id=self.id,
            output_file=self.output_file,
            protocols=self.protocols,
            type=self.type,
            zone_id=self.zone_id)

def get_protocols(output_file=None,type=None,zone_id=None,opts=None):
    """
    Provide  a data source to retrieve the type of protocol used to create NAS file system.

    > **NOTE:** Available in 1.42.0

    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/nas_protocols.html.markdown.


    :param str type: The file system type. Valid Values: Performance and Capacity.  
    :param str zone_id: String to filter results by zone id. 
    """
    __args__ = dict()


    __args__['outputFile'] = output_file
    __args__['type'] = type
    __args__['zoneId'] = zone_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:nas/getProtocols:getProtocols', __args__, opts=opts).value

    return AwaitableGetProtocolsResult(
        id=__ret__.get('id'),
        output_file=__ret__.get('outputFile'),
        protocols=__ret__.get('protocols'),
        type=__ret__.get('type'),
        zone_id=__ret__.get('zoneId'))
