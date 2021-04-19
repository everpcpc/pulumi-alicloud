# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SwitchArgs', 'Switch']

@pulumi.input_type
class SwitchArgs:
    def __init__(__self__, *,
                 cidr_block: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vswitch_name: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Switch resource.
        :param pulumi.Input[str] cidr_block: The CIDR block for the switch.
        :param pulumi.Input[str] vpc_id: The VPC ID.
        :param pulumi.Input[str] availability_zone: Field `availability_zone` has been deprecated from provider version 1.119.0. New field `zone_id` instead.
        :param pulumi.Input[str] description: The switch description. Defaults to null.
        :param pulumi.Input[str] name: Field `name` has been deprecated from provider version 1.119.0. New field `vswitch_name` instead.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vswitch_name: The name of the switch. Defaults to null.
        :param pulumi.Input[str] zone_id: The AZ for the switch.
        """
        pulumi.set(__self__, "cidr_block", cidr_block)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if availability_zone is not None:
            warnings.warn("""Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.""", DeprecationWarning)
            pulumi.log.warn("""availability_zone is deprecated: Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.""")
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            warnings.warn("""Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.""", DeprecationWarning)
            pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.""")
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vswitch_name is not None:
            pulumi.set(__self__, "vswitch_name", vswitch_name)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> pulumi.Input[str]:
        """
        The CIDR block for the switch.
        """
        return pulumi.get(self, "cidr_block")

    @cidr_block.setter
    def cidr_block(self, value: pulumi.Input[str]):
        pulumi.set(self, "cidr_block", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The VPC ID.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        """
        Field `availability_zone` has been deprecated from provider version 1.119.0. New field `zone_id` instead.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The switch description. Defaults to null.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Field `name` has been deprecated from provider version 1.119.0. New field `vswitch_name` instead.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vswitchName")
    def vswitch_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the switch. Defaults to null.
        """
        return pulumi.get(self, "vswitch_name")

    @vswitch_name.setter
    def vswitch_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_name", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The AZ for the switch.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


@pulumi.input_type
class _SwitchState:
    def __init__(__self__, *,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_name: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Switch resources.
        :param pulumi.Input[str] availability_zone: Field `availability_zone` has been deprecated from provider version 1.119.0. New field `zone_id` instead.
        :param pulumi.Input[str] cidr_block: The CIDR block for the switch.
        :param pulumi.Input[str] description: The switch description. Defaults to null.
        :param pulumi.Input[str] name: Field `name` has been deprecated from provider version 1.119.0. New field `vswitch_name` instead.
        :param pulumi.Input[str] status: (Available in 1.119.0+) The status of the switch.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID.
        :param pulumi.Input[str] vswitch_name: The name of the switch. Defaults to null.
        :param pulumi.Input[str] zone_id: The AZ for the switch.
        """
        if availability_zone is not None:
            warnings.warn("""Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.""", DeprecationWarning)
            pulumi.log.warn("""availability_zone is deprecated: Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.""")
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if cidr_block is not None:
            pulumi.set(__self__, "cidr_block", cidr_block)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            warnings.warn("""Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.""", DeprecationWarning)
            pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.""")
        if name is not None:
            pulumi.set(__self__, "name", name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vswitch_name is not None:
            pulumi.set(__self__, "vswitch_name", vswitch_name)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        """
        Field `availability_zone` has been deprecated from provider version 1.119.0. New field `zone_id` instead.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The CIDR block for the switch.
        """
        return pulumi.get(self, "cidr_block")

    @cidr_block.setter
    def cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr_block", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The switch description. Defaults to null.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Field `name` has been deprecated from provider version 1.119.0. New field `vswitch_name` instead.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        (Available in 1.119.0+) The status of the switch.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The VPC ID.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vswitchName")
    def vswitch_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the switch. Defaults to null.
        """
        return pulumi.get(self, "vswitch_name")

    @vswitch_name.setter
    def vswitch_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_name", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The AZ for the switch.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class Switch(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_name: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        Vswitch can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/switch:Switch example vsw-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: Field `availability_zone` has been deprecated from provider version 1.119.0. New field `zone_id` instead.
        :param pulumi.Input[str] cidr_block: The CIDR block for the switch.
        :param pulumi.Input[str] description: The switch description. Defaults to null.
        :param pulumi.Input[str] name: Field `name` has been deprecated from provider version 1.119.0. New field `vswitch_name` instead.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID.
        :param pulumi.Input[str] vswitch_name: The name of the switch. Defaults to null.
        :param pulumi.Input[str] zone_id: The AZ for the switch.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SwitchArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Vswitch can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/switch:Switch example vsw-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param SwitchArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_name: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = SwitchArgs.__new__(SwitchArgs)

            if availability_zone is not None and not opts.urn:
                warnings.warn("""Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.""", DeprecationWarning)
                pulumi.log.warn("""availability_zone is deprecated: Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.""")
            __props__.__dict__["availability_zone"] = availability_zone
            if cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'cidr_block'")
            __props__.__dict__["cidr_block"] = cidr_block
            __props__.__dict__["description"] = description
            if name is not None and not opts.urn:
                warnings.warn("""Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.""", DeprecationWarning)
                pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.""")
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["vswitch_name"] = vswitch_name
            __props__.__dict__["zone_id"] = zone_id
            __props__.__dict__["status"] = None
        super(Switch, __self__).__init__(
            'alicloud:vpc/switch:Switch',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            cidr_block: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vswitch_name: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'Switch':
        """
        Get an existing Switch resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: Field `availability_zone` has been deprecated from provider version 1.119.0. New field `zone_id` instead.
        :param pulumi.Input[str] cidr_block: The CIDR block for the switch.
        :param pulumi.Input[str] description: The switch description. Defaults to null.
        :param pulumi.Input[str] name: Field `name` has been deprecated from provider version 1.119.0. New field `vswitch_name` instead.
        :param pulumi.Input[str] status: (Available in 1.119.0+) The status of the switch.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID.
        :param pulumi.Input[str] vswitch_name: The name of the switch. Defaults to null.
        :param pulumi.Input[str] zone_id: The AZ for the switch.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchState.__new__(_SwitchState)

        __props__.__dict__["availability_zone"] = availability_zone
        __props__.__dict__["cidr_block"] = cidr_block
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["vswitch_name"] = vswitch_name
        __props__.__dict__["zone_id"] = zone_id
        return Switch(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        """
        Field `availability_zone` has been deprecated from provider version 1.119.0. New field `zone_id` instead.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> pulumi.Output[str]:
        """
        The CIDR block for the switch.
        """
        return pulumi.get(self, "cidr_block")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The switch description. Defaults to null.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Field `name` has been deprecated from provider version 1.119.0. New field `vswitch_name` instead.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        (Available in 1.119.0+) The status of the switch.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The VPC ID.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchName")
    def vswitch_name(self) -> pulumi.Output[str]:
        """
        The name of the switch. Defaults to null.
        """
        return pulumi.get(self, "vswitch_name")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The AZ for the switch.
        """
        return pulumi.get(self, "zone_id")

