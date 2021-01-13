# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Network']


class Network(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a cloud connect network resource. Cloud Connect Network (CCN) is another important component of Smart Access Gateway. It is a device access matrix composed of Alibaba Cloud distributed access gateways. You can add multiple Smart Access Gateway (SAG) devices to a CCN instance and then attach the CCN instance to a Cloud Enterprise Network (CEN) instance to connect the local branches to the Alibaba Cloud.

        For information about cloud connect network and how to use it, see [What is Cloud Connect Network](https://www.alibabacloud.com/help/doc-detail/93667.htm).

        > **NOTE:** Available in 1.59.0+

        > **NOTE:** Only the following regions support create Cloud Connect Network. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.cloudconnect.Network("default",
            cidr_block="192.168.0.0/24",
            description="tf-testAccCloudConnectNetworkDescription",
            is_default=True)
        ```

        ## Import

        The cloud connect network instance can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cloudconnect/network:Network example ccn-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr_block: The CidrBlock of the CCN instance. Defaults to null.
        :param pulumi.Input[str] description: The description of the CCN instance. The description can contain 2 to 256 characters. The description must start with English letters, but cannot start with http:// or https://.
        :param pulumi.Input[bool] is_default: Created by default. If the client does not have ccn in the binding, it will create a ccn for the user to replace.
        :param pulumi.Input[str] name: The name of the CCN instance. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.
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

            __props__['cidr_block'] = cidr_block
            __props__['description'] = description
            if is_default is None and not opts.urn:
                raise TypeError("Missing required property 'is_default'")
            __props__['is_default'] = is_default
            __props__['name'] = name
        super(Network, __self__).__init__(
            'alicloud:cloudconnect/network:Network',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cidr_block: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            is_default: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'Network':
        """
        Get an existing Network resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr_block: The CidrBlock of the CCN instance. Defaults to null.
        :param pulumi.Input[str] description: The description of the CCN instance. The description can contain 2 to 256 characters. The description must start with English letters, but cannot start with http:// or https://.
        :param pulumi.Input[bool] is_default: Created by default. If the client does not have ccn in the binding, it will create a ccn for the user to replace.
        :param pulumi.Input[str] name: The name of the CCN instance. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cidr_block"] = cidr_block
        __props__["description"] = description
        __props__["is_default"] = is_default
        __props__["name"] = name
        return Network(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> pulumi.Output[Optional[str]]:
        """
        The CidrBlock of the CCN instance. Defaults to null.
        """
        return pulumi.get(self, "cidr_block")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the CCN instance. The description can contain 2 to 256 characters. The description must start with English letters, but cannot start with http:// or https://.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> pulumi.Output[bool]:
        """
        Created by default. If the client does not have ccn in the binding, it will create a ccn for the user to replace.
        """
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the CCN instance. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.
        """
        return pulumi.get(self, "name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

