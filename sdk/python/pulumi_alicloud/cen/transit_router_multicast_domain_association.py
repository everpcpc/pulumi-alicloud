# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TransitRouterMulticastDomainAssociationArgs', 'TransitRouterMulticastDomainAssociation']

@pulumi.input_type
class TransitRouterMulticastDomainAssociationArgs:
    def __init__(__self__, *,
                 transit_router_attachment_id: pulumi.Input[str],
                 transit_router_multicast_domain_id: pulumi.Input[str],
                 vswitch_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a TransitRouterMulticastDomainAssociation resource.
        :param pulumi.Input[str] transit_router_attachment_id: The ID of the VPC connection.
        :param pulumi.Input[str] transit_router_multicast_domain_id: The ID of the multicast domain.
        :param pulumi.Input[str] vswitch_id: The ID of the vSwitch.
        """
        pulumi.set(__self__, "transit_router_attachment_id", transit_router_attachment_id)
        pulumi.set(__self__, "transit_router_multicast_domain_id", transit_router_multicast_domain_id)
        pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="transitRouterAttachmentId")
    def transit_router_attachment_id(self) -> pulumi.Input[str]:
        """
        The ID of the VPC connection.
        """
        return pulumi.get(self, "transit_router_attachment_id")

    @transit_router_attachment_id.setter
    def transit_router_attachment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "transit_router_attachment_id", value)

    @property
    @pulumi.getter(name="transitRouterMulticastDomainId")
    def transit_router_multicast_domain_id(self) -> pulumi.Input[str]:
        """
        The ID of the multicast domain.
        """
        return pulumi.get(self, "transit_router_multicast_domain_id")

    @transit_router_multicast_domain_id.setter
    def transit_router_multicast_domain_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "transit_router_multicast_domain_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Input[str]:
        """
        The ID of the vSwitch.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vswitch_id", value)


@pulumi.input_type
class _TransitRouterMulticastDomainAssociationState:
    def __init__(__self__, *,
                 status: Optional[pulumi.Input[str]] = None,
                 transit_router_attachment_id: Optional[pulumi.Input[str]] = None,
                 transit_router_multicast_domain_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TransitRouterMulticastDomainAssociation resources.
        :param pulumi.Input[str] status: The status of the Transit Router Multicast Domain Association.
        :param pulumi.Input[str] transit_router_attachment_id: The ID of the VPC connection.
        :param pulumi.Input[str] transit_router_multicast_domain_id: The ID of the multicast domain.
        :param pulumi.Input[str] vswitch_id: The ID of the vSwitch.
        """
        if status is not None:
            pulumi.set(__self__, "status", status)
        if transit_router_attachment_id is not None:
            pulumi.set(__self__, "transit_router_attachment_id", transit_router_attachment_id)
        if transit_router_multicast_domain_id is not None:
            pulumi.set(__self__, "transit_router_multicast_domain_id", transit_router_multicast_domain_id)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the Transit Router Multicast Domain Association.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="transitRouterAttachmentId")
    def transit_router_attachment_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPC connection.
        """
        return pulumi.get(self, "transit_router_attachment_id")

    @transit_router_attachment_id.setter
    def transit_router_attachment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_attachment_id", value)

    @property
    @pulumi.getter(name="transitRouterMulticastDomainId")
    def transit_router_multicast_domain_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the multicast domain.
        """
        return pulumi.get(self, "transit_router_multicast_domain_id")

    @transit_router_multicast_domain_id.setter
    def transit_router_multicast_domain_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_multicast_domain_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the vSwitch.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)


class TransitRouterMulticastDomainAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 transit_router_attachment_id: Optional[pulumi.Input[str]] = None,
                 transit_router_multicast_domain_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Cloud Enterprise Network (CEN) Transit Router Multicast Domain Association resource.

        For information about Cloud Enterprise Network (CEN) Transit Router Multicast Domain Association and how to use it, see [What is Transit Router Multicast Domain Association](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-associatetransitroutermulticastdomain).

        > **NOTE:** Available in v1.195.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_instance = alicloud.cen.Instance("defaultInstance", cen_instance_name="tf-example")
        default_transit_router = alicloud.cen.TransitRouter("defaultTransitRouter",
            cen_id=default_instance.id,
            support_multicast=True)
        default_transit_router_multicast_domain = alicloud.cen.TransitRouterMulticastDomain("defaultTransitRouterMulticastDomain", transit_router_id=default_transit_router.transit_router_id)
        default_transit_router_vpc_attachment = alicloud.cen.TransitRouterVpcAttachment("defaultTransitRouterVpcAttachment",
            cen_id=default_transit_router.cen_id,
            transit_router_id=default_transit_router_multicast_domain.transit_router_id,
            vpc_id="your_vpc_id",
            zone_mappings=[alicloud.cen.TransitRouterVpcAttachmentZoneMappingArgs(
                zone_id="your_zone_id",
                vswitch_id="your_vswitch_id",
            )])
        default_transit_router_multicast_domain_association = alicloud.cen.TransitRouterMulticastDomainAssociation("defaultTransitRouterMulticastDomainAssociation",
            transit_router_multicast_domain_id=default_transit_router_multicast_domain.id,
            transit_router_attachment_id=default_transit_router_vpc_attachment.transit_router_attachment_id,
            vswitch_id="your_vswitch_id")
        ```

        ## Import

        Cloud Enterprise Network (CEN) Transit Router Multicast Domain Association can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cen/transitRouterMulticastDomainAssociation:TransitRouterMulticastDomainAssociation example <transit_router_multicast_domain_id>:<transit_router_attachment_id>:<vswitch_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] transit_router_attachment_id: The ID of the VPC connection.
        :param pulumi.Input[str] transit_router_multicast_domain_id: The ID of the multicast domain.
        :param pulumi.Input[str] vswitch_id: The ID of the vSwitch.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TransitRouterMulticastDomainAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Cloud Enterprise Network (CEN) Transit Router Multicast Domain Association resource.

        For information about Cloud Enterprise Network (CEN) Transit Router Multicast Domain Association and how to use it, see [What is Transit Router Multicast Domain Association](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-associatetransitroutermulticastdomain).

        > **NOTE:** Available in v1.195.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_instance = alicloud.cen.Instance("defaultInstance", cen_instance_name="tf-example")
        default_transit_router = alicloud.cen.TransitRouter("defaultTransitRouter",
            cen_id=default_instance.id,
            support_multicast=True)
        default_transit_router_multicast_domain = alicloud.cen.TransitRouterMulticastDomain("defaultTransitRouterMulticastDomain", transit_router_id=default_transit_router.transit_router_id)
        default_transit_router_vpc_attachment = alicloud.cen.TransitRouterVpcAttachment("defaultTransitRouterVpcAttachment",
            cen_id=default_transit_router.cen_id,
            transit_router_id=default_transit_router_multicast_domain.transit_router_id,
            vpc_id="your_vpc_id",
            zone_mappings=[alicloud.cen.TransitRouterVpcAttachmentZoneMappingArgs(
                zone_id="your_zone_id",
                vswitch_id="your_vswitch_id",
            )])
        default_transit_router_multicast_domain_association = alicloud.cen.TransitRouterMulticastDomainAssociation("defaultTransitRouterMulticastDomainAssociation",
            transit_router_multicast_domain_id=default_transit_router_multicast_domain.id,
            transit_router_attachment_id=default_transit_router_vpc_attachment.transit_router_attachment_id,
            vswitch_id="your_vswitch_id")
        ```

        ## Import

        Cloud Enterprise Network (CEN) Transit Router Multicast Domain Association can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cen/transitRouterMulticastDomainAssociation:TransitRouterMulticastDomainAssociation example <transit_router_multicast_domain_id>:<transit_router_attachment_id>:<vswitch_id>
        ```

        :param str resource_name: The name of the resource.
        :param TransitRouterMulticastDomainAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TransitRouterMulticastDomainAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 transit_router_attachment_id: Optional[pulumi.Input[str]] = None,
                 transit_router_multicast_domain_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TransitRouterMulticastDomainAssociationArgs.__new__(TransitRouterMulticastDomainAssociationArgs)

            if transit_router_attachment_id is None and not opts.urn:
                raise TypeError("Missing required property 'transit_router_attachment_id'")
            __props__.__dict__["transit_router_attachment_id"] = transit_router_attachment_id
            if transit_router_multicast_domain_id is None and not opts.urn:
                raise TypeError("Missing required property 'transit_router_multicast_domain_id'")
            __props__.__dict__["transit_router_multicast_domain_id"] = transit_router_multicast_domain_id
            if vswitch_id is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__.__dict__["vswitch_id"] = vswitch_id
            __props__.__dict__["status"] = None
        super(TransitRouterMulticastDomainAssociation, __self__).__init__(
            'alicloud:cen/transitRouterMulticastDomainAssociation:TransitRouterMulticastDomainAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            status: Optional[pulumi.Input[str]] = None,
            transit_router_attachment_id: Optional[pulumi.Input[str]] = None,
            transit_router_multicast_domain_id: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None) -> 'TransitRouterMulticastDomainAssociation':
        """
        Get an existing TransitRouterMulticastDomainAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] status: The status of the Transit Router Multicast Domain Association.
        :param pulumi.Input[str] transit_router_attachment_id: The ID of the VPC connection.
        :param pulumi.Input[str] transit_router_multicast_domain_id: The ID of the multicast domain.
        :param pulumi.Input[str] vswitch_id: The ID of the vSwitch.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TransitRouterMulticastDomainAssociationState.__new__(_TransitRouterMulticastDomainAssociationState)

        __props__.__dict__["status"] = status
        __props__.__dict__["transit_router_attachment_id"] = transit_router_attachment_id
        __props__.__dict__["transit_router_multicast_domain_id"] = transit_router_multicast_domain_id
        __props__.__dict__["vswitch_id"] = vswitch_id
        return TransitRouterMulticastDomainAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the Transit Router Multicast Domain Association.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="transitRouterAttachmentId")
    def transit_router_attachment_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPC connection.
        """
        return pulumi.get(self, "transit_router_attachment_id")

    @property
    @pulumi.getter(name="transitRouterMulticastDomainId")
    def transit_router_multicast_domain_id(self) -> pulumi.Output[str]:
        """
        The ID of the multicast domain.
        """
        return pulumi.get(self, "transit_router_multicast_domain_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[str]:
        """
        The ID of the vSwitch.
        """
        return pulumi.get(self, "vswitch_id")

