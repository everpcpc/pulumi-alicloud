# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SnatEntryArgs', 'SnatEntry']

@pulumi.input_type
class SnatEntryArgs:
    def __init__(__self__, *,
                 snat_ip: pulumi.Input[str],
                 snat_table_id: pulumi.Input[str],
                 snat_entry_name: Optional[pulumi.Input[str]] = None,
                 source_cidr: Optional[pulumi.Input[str]] = None,
                 source_vswitch_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SnatEntry resource.
        :param pulumi.Input[str] snat_ip: The SNAT ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        :param pulumi.Input[str] snat_table_id: The value can get from `vpc.NatGateway` Attributes "snat_table_ids".
        :param pulumi.Input[str] snat_entry_name: The name of snat entry.
        :param pulumi.Input[str] source_cidr: The private network segment of Ecs. This parameter and the `source_vswitch_id` parameter are mutually exclusive and cannot appear at the same time.
        :param pulumi.Input[str] source_vswitch_id: The vswitch ID.
        """
        pulumi.set(__self__, "snat_ip", snat_ip)
        pulumi.set(__self__, "snat_table_id", snat_table_id)
        if snat_entry_name is not None:
            pulumi.set(__self__, "snat_entry_name", snat_entry_name)
        if source_cidr is not None:
            pulumi.set(__self__, "source_cidr", source_cidr)
        if source_vswitch_id is not None:
            pulumi.set(__self__, "source_vswitch_id", source_vswitch_id)

    @property
    @pulumi.getter(name="snatIp")
    def snat_ip(self) -> pulumi.Input[str]:
        """
        The SNAT ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        """
        return pulumi.get(self, "snat_ip")

    @snat_ip.setter
    def snat_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "snat_ip", value)

    @property
    @pulumi.getter(name="snatTableId")
    def snat_table_id(self) -> pulumi.Input[str]:
        """
        The value can get from `vpc.NatGateway` Attributes "snat_table_ids".
        """
        return pulumi.get(self, "snat_table_id")

    @snat_table_id.setter
    def snat_table_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "snat_table_id", value)

    @property
    @pulumi.getter(name="snatEntryName")
    def snat_entry_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of snat entry.
        """
        return pulumi.get(self, "snat_entry_name")

    @snat_entry_name.setter
    def snat_entry_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snat_entry_name", value)

    @property
    @pulumi.getter(name="sourceCidr")
    def source_cidr(self) -> Optional[pulumi.Input[str]]:
        """
        The private network segment of Ecs. This parameter and the `source_vswitch_id` parameter are mutually exclusive and cannot appear at the same time.
        """
        return pulumi.get(self, "source_cidr")

    @source_cidr.setter
    def source_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_cidr", value)

    @property
    @pulumi.getter(name="sourceVswitchId")
    def source_vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The vswitch ID.
        """
        return pulumi.get(self, "source_vswitch_id")

    @source_vswitch_id.setter
    def source_vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_vswitch_id", value)


@pulumi.input_type
class _SnatEntryState:
    def __init__(__self__, *,
                 snat_entry_id: Optional[pulumi.Input[str]] = None,
                 snat_entry_name: Optional[pulumi.Input[str]] = None,
                 snat_ip: Optional[pulumi.Input[str]] = None,
                 snat_table_id: Optional[pulumi.Input[str]] = None,
                 source_cidr: Optional[pulumi.Input[str]] = None,
                 source_vswitch_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SnatEntry resources.
        :param pulumi.Input[str] snat_entry_id: The id of the snat entry on the server.
        :param pulumi.Input[str] snat_entry_name: The name of snat entry.
        :param pulumi.Input[str] snat_ip: The SNAT ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        :param pulumi.Input[str] snat_table_id: The value can get from `vpc.NatGateway` Attributes "snat_table_ids".
        :param pulumi.Input[str] source_cidr: The private network segment of Ecs. This parameter and the `source_vswitch_id` parameter are mutually exclusive and cannot appear at the same time.
        :param pulumi.Input[str] source_vswitch_id: The vswitch ID.
        :param pulumi.Input[str] status: (Available in 1.119.1+) The status of snat entry.
        """
        if snat_entry_id is not None:
            pulumi.set(__self__, "snat_entry_id", snat_entry_id)
        if snat_entry_name is not None:
            pulumi.set(__self__, "snat_entry_name", snat_entry_name)
        if snat_ip is not None:
            pulumi.set(__self__, "snat_ip", snat_ip)
        if snat_table_id is not None:
            pulumi.set(__self__, "snat_table_id", snat_table_id)
        if source_cidr is not None:
            pulumi.set(__self__, "source_cidr", source_cidr)
        if source_vswitch_id is not None:
            pulumi.set(__self__, "source_vswitch_id", source_vswitch_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="snatEntryId")
    def snat_entry_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the snat entry on the server.
        """
        return pulumi.get(self, "snat_entry_id")

    @snat_entry_id.setter
    def snat_entry_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snat_entry_id", value)

    @property
    @pulumi.getter(name="snatEntryName")
    def snat_entry_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of snat entry.
        """
        return pulumi.get(self, "snat_entry_name")

    @snat_entry_name.setter
    def snat_entry_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snat_entry_name", value)

    @property
    @pulumi.getter(name="snatIp")
    def snat_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The SNAT ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        """
        return pulumi.get(self, "snat_ip")

    @snat_ip.setter
    def snat_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snat_ip", value)

    @property
    @pulumi.getter(name="snatTableId")
    def snat_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        The value can get from `vpc.NatGateway` Attributes "snat_table_ids".
        """
        return pulumi.get(self, "snat_table_id")

    @snat_table_id.setter
    def snat_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snat_table_id", value)

    @property
    @pulumi.getter(name="sourceCidr")
    def source_cidr(self) -> Optional[pulumi.Input[str]]:
        """
        The private network segment of Ecs. This parameter and the `source_vswitch_id` parameter are mutually exclusive and cannot appear at the same time.
        """
        return pulumi.get(self, "source_cidr")

    @source_cidr.setter
    def source_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_cidr", value)

    @property
    @pulumi.getter(name="sourceVswitchId")
    def source_vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The vswitch ID.
        """
        return pulumi.get(self, "source_vswitch_id")

    @source_vswitch_id.setter
    def source_vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_vswitch_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        (Available in 1.119.1+) The status of snat entry.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class SnatEntry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 snat_entry_name: Optional[pulumi.Input[str]] = None,
                 snat_ip: Optional[pulumi.Input[str]] = None,
                 snat_table_id: Optional[pulumi.Input[str]] = None,
                 source_cidr: Optional[pulumi.Input[str]] = None,
                 source_vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a snat resource.

        ## Import

        Snat Entry can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/snatEntry:SnatEntry foo stb-1aece3:snat-232ce2
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] snat_entry_name: The name of snat entry.
        :param pulumi.Input[str] snat_ip: The SNAT ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        :param pulumi.Input[str] snat_table_id: The value can get from `vpc.NatGateway` Attributes "snat_table_ids".
        :param pulumi.Input[str] source_cidr: The private network segment of Ecs. This parameter and the `source_vswitch_id` parameter are mutually exclusive and cannot appear at the same time.
        :param pulumi.Input[str] source_vswitch_id: The vswitch ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnatEntryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a snat resource.

        ## Import

        Snat Entry can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/snatEntry:SnatEntry foo stb-1aece3:snat-232ce2
        ```

        :param str resource_name: The name of the resource.
        :param SnatEntryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnatEntryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 snat_entry_name: Optional[pulumi.Input[str]] = None,
                 snat_ip: Optional[pulumi.Input[str]] = None,
                 snat_table_id: Optional[pulumi.Input[str]] = None,
                 source_cidr: Optional[pulumi.Input[str]] = None,
                 source_vswitch_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = SnatEntryArgs.__new__(SnatEntryArgs)

            __props__.__dict__["snat_entry_name"] = snat_entry_name
            if snat_ip is None and not opts.urn:
                raise TypeError("Missing required property 'snat_ip'")
            __props__.__dict__["snat_ip"] = snat_ip
            if snat_table_id is None and not opts.urn:
                raise TypeError("Missing required property 'snat_table_id'")
            __props__.__dict__["snat_table_id"] = snat_table_id
            __props__.__dict__["source_cidr"] = source_cidr
            __props__.__dict__["source_vswitch_id"] = source_vswitch_id
            __props__.__dict__["snat_entry_id"] = None
            __props__.__dict__["status"] = None
        super(SnatEntry, __self__).__init__(
            'alicloud:vpc/snatEntry:SnatEntry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            snat_entry_id: Optional[pulumi.Input[str]] = None,
            snat_entry_name: Optional[pulumi.Input[str]] = None,
            snat_ip: Optional[pulumi.Input[str]] = None,
            snat_table_id: Optional[pulumi.Input[str]] = None,
            source_cidr: Optional[pulumi.Input[str]] = None,
            source_vswitch_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'SnatEntry':
        """
        Get an existing SnatEntry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] snat_entry_id: The id of the snat entry on the server.
        :param pulumi.Input[str] snat_entry_name: The name of snat entry.
        :param pulumi.Input[str] snat_ip: The SNAT ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        :param pulumi.Input[str] snat_table_id: The value can get from `vpc.NatGateway` Attributes "snat_table_ids".
        :param pulumi.Input[str] source_cidr: The private network segment of Ecs. This parameter and the `source_vswitch_id` parameter are mutually exclusive and cannot appear at the same time.
        :param pulumi.Input[str] source_vswitch_id: The vswitch ID.
        :param pulumi.Input[str] status: (Available in 1.119.1+) The status of snat entry.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnatEntryState.__new__(_SnatEntryState)

        __props__.__dict__["snat_entry_id"] = snat_entry_id
        __props__.__dict__["snat_entry_name"] = snat_entry_name
        __props__.__dict__["snat_ip"] = snat_ip
        __props__.__dict__["snat_table_id"] = snat_table_id
        __props__.__dict__["source_cidr"] = source_cidr
        __props__.__dict__["source_vswitch_id"] = source_vswitch_id
        __props__.__dict__["status"] = status
        return SnatEntry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="snatEntryId")
    def snat_entry_id(self) -> pulumi.Output[str]:
        """
        The id of the snat entry on the server.
        """
        return pulumi.get(self, "snat_entry_id")

    @property
    @pulumi.getter(name="snatEntryName")
    def snat_entry_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of snat entry.
        """
        return pulumi.get(self, "snat_entry_name")

    @property
    @pulumi.getter(name="snatIp")
    def snat_ip(self) -> pulumi.Output[str]:
        """
        The SNAT ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        """
        return pulumi.get(self, "snat_ip")

    @property
    @pulumi.getter(name="snatTableId")
    def snat_table_id(self) -> pulumi.Output[str]:
        """
        The value can get from `vpc.NatGateway` Attributes "snat_table_ids".
        """
        return pulumi.get(self, "snat_table_id")

    @property
    @pulumi.getter(name="sourceCidr")
    def source_cidr(self) -> pulumi.Output[Optional[str]]:
        """
        The private network segment of Ecs. This parameter and the `source_vswitch_id` parameter are mutually exclusive and cannot appear at the same time.
        """
        return pulumi.get(self, "source_cidr")

    @property
    @pulumi.getter(name="sourceVswitchId")
    def source_vswitch_id(self) -> pulumi.Output[Optional[str]]:
        """
        The vswitch ID.
        """
        return pulumi.get(self, "source_vswitch_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        (Available in 1.119.1+) The status of snat entry.
        """
        return pulumi.get(self, "status")

