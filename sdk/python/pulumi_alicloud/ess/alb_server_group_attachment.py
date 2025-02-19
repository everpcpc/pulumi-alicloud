# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AlbServerGroupAttachmentArgs', 'AlbServerGroupAttachment']

@pulumi.input_type
class AlbServerGroupAttachmentArgs:
    def __init__(__self__, *,
                 alb_server_group_id: pulumi.Input[str],
                 port: pulumi.Input[int],
                 scaling_group_id: pulumi.Input[str],
                 weight: pulumi.Input[int],
                 force_attach: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AlbServerGroupAttachment resource.
        :param pulumi.Input[str] alb_server_group_id: ID of Alb Server Group.
        :param pulumi.Input[int] port: The port will be used for Alb Server Group backend server.
        :param pulumi.Input[str] scaling_group_id: ID of the scaling group.
        :param pulumi.Input[int] weight: The weight of an ECS instance attached to the Alb Server Group.
        :param pulumi.Input[bool] force_attach: If instances of scaling group are attached/removed from slb backend server when attach/detach alb
               server group from scaling group. Default to false.
        """
        pulumi.set(__self__, "alb_server_group_id", alb_server_group_id)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "scaling_group_id", scaling_group_id)
        pulumi.set(__self__, "weight", weight)
        if force_attach is not None:
            pulumi.set(__self__, "force_attach", force_attach)

    @property
    @pulumi.getter(name="albServerGroupId")
    def alb_server_group_id(self) -> pulumi.Input[str]:
        """
        ID of Alb Server Group.
        """
        return pulumi.get(self, "alb_server_group_id")

    @alb_server_group_id.setter
    def alb_server_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "alb_server_group_id", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        The port will be used for Alb Server Group backend server.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> pulumi.Input[str]:
        """
        ID of the scaling group.
        """
        return pulumi.get(self, "scaling_group_id")

    @scaling_group_id.setter
    def scaling_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "scaling_group_id", value)

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Input[int]:
        """
        The weight of an ECS instance attached to the Alb Server Group.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: pulumi.Input[int]):
        pulumi.set(self, "weight", value)

    @property
    @pulumi.getter(name="forceAttach")
    def force_attach(self) -> Optional[pulumi.Input[bool]]:
        """
        If instances of scaling group are attached/removed from slb backend server when attach/detach alb
        server group from scaling group. Default to false.
        """
        return pulumi.get(self, "force_attach")

    @force_attach.setter
    def force_attach(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_attach", value)


@pulumi.input_type
class _AlbServerGroupAttachmentState:
    def __init__(__self__, *,
                 alb_server_group_id: Optional[pulumi.Input[str]] = None,
                 force_attach: Optional[pulumi.Input[bool]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering AlbServerGroupAttachment resources.
        :param pulumi.Input[str] alb_server_group_id: ID of Alb Server Group.
        :param pulumi.Input[bool] force_attach: If instances of scaling group are attached/removed from slb backend server when attach/detach alb
               server group from scaling group. Default to false.
        :param pulumi.Input[int] port: The port will be used for Alb Server Group backend server.
        :param pulumi.Input[str] scaling_group_id: ID of the scaling group.
        :param pulumi.Input[int] weight: The weight of an ECS instance attached to the Alb Server Group.
        """
        if alb_server_group_id is not None:
            pulumi.set(__self__, "alb_server_group_id", alb_server_group_id)
        if force_attach is not None:
            pulumi.set(__self__, "force_attach", force_attach)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if scaling_group_id is not None:
            pulumi.set(__self__, "scaling_group_id", scaling_group_id)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter(name="albServerGroupId")
    def alb_server_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of Alb Server Group.
        """
        return pulumi.get(self, "alb_server_group_id")

    @alb_server_group_id.setter
    def alb_server_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alb_server_group_id", value)

    @property
    @pulumi.getter(name="forceAttach")
    def force_attach(self) -> Optional[pulumi.Input[bool]]:
        """
        If instances of scaling group are attached/removed from slb backend server when attach/detach alb
        server group from scaling group. Default to false.
        """
        return pulumi.get(self, "force_attach")

    @force_attach.setter
    def force_attach(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_attach", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port will be used for Alb Server Group backend server.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the scaling group.
        """
        return pulumi.get(self, "scaling_group_id")

    @scaling_group_id.setter
    def scaling_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scaling_group_id", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        The weight of an ECS instance attached to the Alb Server Group.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


class AlbServerGroupAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alb_server_group_id: Optional[pulumi.Input[str]] = None,
                 force_attach: Optional[pulumi.Input[bool]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Attaches/Detaches alb server group to a specified scaling group.

        For information about alb server group attachment, see [AttachAlbServerGroups](https://www.alibabacloud.com/help/en/doc-detail/266800.html).

        > **NOTE:** If scaling group's network type is `VPC`, the alb server groups must be in the same `VPC`.

        > **NOTE:** Alb server group attachment is defined uniquely by `scaling_group_id`, `alb_server_group_id`, `port`.

        > **NOTE:** Resource `ess.AlbServerGroupAttachment` don't support modification.

        > **NOTE:** Resource `ess.AlbServerGroupAttachment` is available in 1.158.0+.

        ## Import

        ESS alb server groups can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ess/albServerGroupAttachment:AlbServerGroupAttachment example asg-xxx:sgp-xxx:5000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alb_server_group_id: ID of Alb Server Group.
        :param pulumi.Input[bool] force_attach: If instances of scaling group are attached/removed from slb backend server when attach/detach alb
               server group from scaling group. Default to false.
        :param pulumi.Input[int] port: The port will be used for Alb Server Group backend server.
        :param pulumi.Input[str] scaling_group_id: ID of the scaling group.
        :param pulumi.Input[int] weight: The weight of an ECS instance attached to the Alb Server Group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AlbServerGroupAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Attaches/Detaches alb server group to a specified scaling group.

        For information about alb server group attachment, see [AttachAlbServerGroups](https://www.alibabacloud.com/help/en/doc-detail/266800.html).

        > **NOTE:** If scaling group's network type is `VPC`, the alb server groups must be in the same `VPC`.

        > **NOTE:** Alb server group attachment is defined uniquely by `scaling_group_id`, `alb_server_group_id`, `port`.

        > **NOTE:** Resource `ess.AlbServerGroupAttachment` don't support modification.

        > **NOTE:** Resource `ess.AlbServerGroupAttachment` is available in 1.158.0+.

        ## Import

        ESS alb server groups can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ess/albServerGroupAttachment:AlbServerGroupAttachment example asg-xxx:sgp-xxx:5000
        ```

        :param str resource_name: The name of the resource.
        :param AlbServerGroupAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AlbServerGroupAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alb_server_group_id: Optional[pulumi.Input[str]] = None,
                 force_attach: Optional[pulumi.Input[bool]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AlbServerGroupAttachmentArgs.__new__(AlbServerGroupAttachmentArgs)

            if alb_server_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'alb_server_group_id'")
            __props__.__dict__["alb_server_group_id"] = alb_server_group_id
            __props__.__dict__["force_attach"] = force_attach
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            if scaling_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'scaling_group_id'")
            __props__.__dict__["scaling_group_id"] = scaling_group_id
            if weight is None and not opts.urn:
                raise TypeError("Missing required property 'weight'")
            __props__.__dict__["weight"] = weight
        super(AlbServerGroupAttachment, __self__).__init__(
            'alicloud:ess/albServerGroupAttachment:AlbServerGroupAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alb_server_group_id: Optional[pulumi.Input[str]] = None,
            force_attach: Optional[pulumi.Input[bool]] = None,
            port: Optional[pulumi.Input[int]] = None,
            scaling_group_id: Optional[pulumi.Input[str]] = None,
            weight: Optional[pulumi.Input[int]] = None) -> 'AlbServerGroupAttachment':
        """
        Get an existing AlbServerGroupAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alb_server_group_id: ID of Alb Server Group.
        :param pulumi.Input[bool] force_attach: If instances of scaling group are attached/removed from slb backend server when attach/detach alb
               server group from scaling group. Default to false.
        :param pulumi.Input[int] port: The port will be used for Alb Server Group backend server.
        :param pulumi.Input[str] scaling_group_id: ID of the scaling group.
        :param pulumi.Input[int] weight: The weight of an ECS instance attached to the Alb Server Group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AlbServerGroupAttachmentState.__new__(_AlbServerGroupAttachmentState)

        __props__.__dict__["alb_server_group_id"] = alb_server_group_id
        __props__.__dict__["force_attach"] = force_attach
        __props__.__dict__["port"] = port
        __props__.__dict__["scaling_group_id"] = scaling_group_id
        __props__.__dict__["weight"] = weight
        return AlbServerGroupAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="albServerGroupId")
    def alb_server_group_id(self) -> pulumi.Output[str]:
        """
        ID of Alb Server Group.
        """
        return pulumi.get(self, "alb_server_group_id")

    @property
    @pulumi.getter(name="forceAttach")
    def force_attach(self) -> pulumi.Output[Optional[bool]]:
        """
        If instances of scaling group are attached/removed from slb backend server when attach/detach alb
        server group from scaling group. Default to false.
        """
        return pulumi.get(self, "force_attach")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        The port will be used for Alb Server Group backend server.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> pulumi.Output[str]:
        """
        ID of the scaling group.
        """
        return pulumi.get(self, "scaling_group_id")

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Output[int]:
        """
        The weight of an ECS instance attached to the Alb Server Group.
        """
        return pulumi.get(self, "weight")

