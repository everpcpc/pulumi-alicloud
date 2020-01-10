# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetDomainExtensionsResult:
    """
    A collection of values returned by getDomainExtensions.
    """
    def __init__(__self__, extensions=None, frontend_port=None, ids=None, load_balancer_id=None, output_file=None, id=None):
        if extensions and not isinstance(extensions, list):
            raise TypeError("Expected argument 'extensions' to be a list")
        __self__.extensions = extensions
        """
        A list of SLB domain extension. Each element contains the following attributes:
        """
        if frontend_port and not isinstance(frontend_port, float):
            raise TypeError("Expected argument 'frontend_port' to be a float")
        __self__.frontend_port = frontend_port
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        if load_balancer_id and not isinstance(load_balancer_id, str):
            raise TypeError("Expected argument 'load_balancer_id' to be a str")
        __self__.load_balancer_id = load_balancer_id
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetDomainExtensionsResult(GetDomainExtensionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainExtensionsResult(
            extensions=self.extensions,
            frontend_port=self.frontend_port,
            ids=self.ids,
            load_balancer_id=self.load_balancer_id,
            output_file=self.output_file,
            id=self.id)

def get_domain_extensions(frontend_port=None,ids=None,load_balancer_id=None,output_file=None,opts=None):
    """
    This data source provides the domain extensions associated with a server load balancer listener.
    
    > **NOTE:** Available in 1.60.0+
    
    :param float frontend_port: The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.
    :param list ids: IDs of the SLB domain extensions.
    :param str load_balancer_id: The ID of the SLB instance.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/slb_domain_extensions.html.markdown.
    """
    __args__ = dict()

    __args__['frontendPort'] = frontend_port
    __args__['ids'] = ids
    __args__['loadBalancerId'] = load_balancer_id
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:slb/getDomainExtensions:getDomainExtensions', __args__, opts=opts).value

    return AwaitableGetDomainExtensionsResult(
        extensions=__ret__.get('extensions'),
        frontend_port=__ret__.get('frontendPort'),
        ids=__ret__.get('ids'),
        load_balancer_id=__ret__.get('loadBalancerId'),
        output_file=__ret__.get('outputFile'),
        id=__ret__.get('id'))
