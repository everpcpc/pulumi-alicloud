# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceAttachmentArgs', 'InstanceAttachment']

@pulumi.input_type
class InstanceAttachmentArgs:
    def __init__(__self__, *,
                 ecs_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a InstanceAttachment resource.
        :param pulumi.Input[str] ecs_id: The ID of the ECS instance.
        :param pulumi.Input[str] instance_id: The ID of the database file system.
        """
        pulumi.set(__self__, "ecs_id", ecs_id)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="ecsId")
    def ecs_id(self) -> pulumi.Input[str]:
        """
        The ID of the ECS instance.
        """
        return pulumi.get(self, "ecs_id")

    @ecs_id.setter
    def ecs_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ecs_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the database file system.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _InstanceAttachmentState:
    def __init__(__self__, *,
                 ecs_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceAttachment resources.
        :param pulumi.Input[str] ecs_id: The ID of the ECS instance.
        :param pulumi.Input[str] instance_id: The ID of the database file system.
        :param pulumi.Input[str] status: The status of Database file system. Valid values: `attached`, `attaching`, `unattached`, `detaching`.
        """
        if ecs_id is not None:
            pulumi.set(__self__, "ecs_id", ecs_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="ecsId")
    def ecs_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the ECS instance.
        """
        return pulumi.get(self, "ecs_id")

    @ecs_id.setter
    def ecs_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ecs_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the database file system.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of Database file system. Valid values: `attached`, `attaching`, `unattached`, `detaching`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class InstanceAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ecs_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a DBFS Instance Attachment resource.

        For information about DBFS Instance Attachment and how to use it, see [What is Instance Attachment](https://help.aliyun.com/document_detail/149726.html).

        > **NOTE:** Available in v1.156.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_networks = alicloud.vpc.get_networks(name_regex="default-NODELETING")
        zone_id = "cn-hangzhou-i"
        default_switches = alicloud.vpc.get_switches(vpc_id=default_networks.ids[0],
            zone_id=zone_id)
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup",
            description="tf test",
            vpc_id=default_networks.ids[0])
        default_images = alicloud.ecs.get_images(owners="system",
            name_regex="^centos_8",
            most_recent=True)
        default_instance = alicloud.ecs.Instance("defaultInstance",
            image_id=default_images.images[0].id,
            instance_name=var["name"],
            instance_type="ecs.g7se.large",
            availability_zone=zone_id,
            vswitch_id=default_switches.ids[0],
            system_disk_category="cloud_essd",
            security_groups=[default_security_group.id])
        default_databasefilesystem_instance_instance = alicloud.databasefilesystem.Instance("defaultDatabasefilesystem/instanceInstance",
            category="standard",
            zone_id=default_instance.availability_zone,
            performance_level="PL1",
            instance_name=var["name"],
            size=100)
        example = alicloud.databasefilesystem.InstanceAttachment("example",
            ecs_id=default_instance.id,
            instance_id=default_databasefilesystem / instance_instance["id"])
        ```

        ## Import

        DBFS Instance Attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:databasefilesystem/instanceAttachment:InstanceAttachment example <instance_id>:<ecs_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ecs_id: The ID of the ECS instance.
        :param pulumi.Input[str] instance_id: The ID of the database file system.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a DBFS Instance Attachment resource.

        For information about DBFS Instance Attachment and how to use it, see [What is Instance Attachment](https://help.aliyun.com/document_detail/149726.html).

        > **NOTE:** Available in v1.156.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_networks = alicloud.vpc.get_networks(name_regex="default-NODELETING")
        zone_id = "cn-hangzhou-i"
        default_switches = alicloud.vpc.get_switches(vpc_id=default_networks.ids[0],
            zone_id=zone_id)
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup",
            description="tf test",
            vpc_id=default_networks.ids[0])
        default_images = alicloud.ecs.get_images(owners="system",
            name_regex="^centos_8",
            most_recent=True)
        default_instance = alicloud.ecs.Instance("defaultInstance",
            image_id=default_images.images[0].id,
            instance_name=var["name"],
            instance_type="ecs.g7se.large",
            availability_zone=zone_id,
            vswitch_id=default_switches.ids[0],
            system_disk_category="cloud_essd",
            security_groups=[default_security_group.id])
        default_databasefilesystem_instance_instance = alicloud.databasefilesystem.Instance("defaultDatabasefilesystem/instanceInstance",
            category="standard",
            zone_id=default_instance.availability_zone,
            performance_level="PL1",
            instance_name=var["name"],
            size=100)
        example = alicloud.databasefilesystem.InstanceAttachment("example",
            ecs_id=default_instance.id,
            instance_id=default_databasefilesystem / instance_instance["id"])
        ```

        ## Import

        DBFS Instance Attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:databasefilesystem/instanceAttachment:InstanceAttachment example <instance_id>:<ecs_id>
        ```

        :param str resource_name: The name of the resource.
        :param InstanceAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ecs_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceAttachmentArgs.__new__(InstanceAttachmentArgs)

            if ecs_id is None and not opts.urn:
                raise TypeError("Missing required property 'ecs_id'")
            __props__.__dict__["ecs_id"] = ecs_id
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["status"] = None
        super(InstanceAttachment, __self__).__init__(
            'alicloud:databasefilesystem/instanceAttachment:InstanceAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ecs_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'InstanceAttachment':
        """
        Get an existing InstanceAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ecs_id: The ID of the ECS instance.
        :param pulumi.Input[str] instance_id: The ID of the database file system.
        :param pulumi.Input[str] status: The status of Database file system. Valid values: `attached`, `attaching`, `unattached`, `detaching`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceAttachmentState.__new__(_InstanceAttachmentState)

        __props__.__dict__["ecs_id"] = ecs_id
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["status"] = status
        return InstanceAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ecsId")
    def ecs_id(self) -> pulumi.Output[str]:
        """
        The ID of the ECS instance.
        """
        return pulumi.get(self, "ecs_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the database file system.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of Database file system. Valid values: `attached`, `attaching`, `unattached`, `detaching`.
        """
        return pulumi.get(self, "status")

