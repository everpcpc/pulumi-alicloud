# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RouteTableAttachmentArgs', 'RouteTableAttachment']

@pulumi.input_type
class RouteTableAttachmentArgs:
    def __init__(__self__, *,
                 route_table_id: pulumi.Input[str],
                 vswitch_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a RouteTableAttachment resource.
        :param pulumi.Input[str] route_table_id: The route_table_id of the route table attachment, the field can't be changed.
        :param pulumi.Input[str] vswitch_id: The vswitch_id of the route table attachment, the field can't be changed.
        """
        pulumi.set(__self__, "route_table_id", route_table_id)
        pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> pulumi.Input[str]:
        """
        The route_table_id of the route table attachment, the field can't be changed.
        """
        return pulumi.get(self, "route_table_id")

    @route_table_id.setter
    def route_table_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "route_table_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Input[str]:
        """
        The vswitch_id of the route table attachment, the field can't be changed.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vswitch_id", value)


@pulumi.input_type
class _RouteTableAttachmentState:
    def __init__(__self__, *,
                 route_table_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RouteTableAttachment resources.
        :param pulumi.Input[str] route_table_id: The route_table_id of the route table attachment, the field can't be changed.
        :param pulumi.Input[str] vswitch_id: The vswitch_id of the route table attachment, the field can't be changed.
        """
        if route_table_id is not None:
            pulumi.set(__self__, "route_table_id", route_table_id)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        The route_table_id of the route table attachment, the field can't be changed.
        """
        return pulumi.get(self, "route_table_id")

    @route_table_id.setter
    def route_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_table_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The vswitch_id of the route table attachment, the field can't be changed.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)


class RouteTableAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 route_table_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        The route table attachemnt can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/routeTableAttachment:RouteTableAttachment foo vtb-abc123456:vsw-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] route_table_id: The route_table_id of the route table attachment, the field can't be changed.
        :param pulumi.Input[str] vswitch_id: The vswitch_id of the route table attachment, the field can't be changed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RouteTableAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        The route table attachemnt can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/routeTableAttachment:RouteTableAttachment foo vtb-abc123456:vsw-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param RouteTableAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouteTableAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 route_table_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = RouteTableAttachmentArgs.__new__(RouteTableAttachmentArgs)

            if route_table_id is None and not opts.urn:
                raise TypeError("Missing required property 'route_table_id'")
            __props__.__dict__["route_table_id"] = route_table_id
            if vswitch_id is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__.__dict__["vswitch_id"] = vswitch_id
        super(RouteTableAttachment, __self__).__init__(
            'alicloud:vpc/routeTableAttachment:RouteTableAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            route_table_id: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None) -> 'RouteTableAttachment':
        """
        Get an existing RouteTableAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] route_table_id: The route_table_id of the route table attachment, the field can't be changed.
        :param pulumi.Input[str] vswitch_id: The vswitch_id of the route table attachment, the field can't be changed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RouteTableAttachmentState.__new__(_RouteTableAttachmentState)

        __props__.__dict__["route_table_id"] = route_table_id
        __props__.__dict__["vswitch_id"] = vswitch_id
        return RouteTableAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> pulumi.Output[str]:
        """
        The route_table_id of the route table attachment, the field can't be changed.
        """
        return pulumi.get(self, "route_table_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[str]:
        """
        The vswitch_id of the route table attachment, the field can't be changed.
        """
        return pulumi.get(self, "vswitch_id")

