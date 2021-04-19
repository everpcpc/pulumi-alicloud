# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AnycastEipAddressAttachmentArgs', 'AnycastEipAddressAttachment']

@pulumi.input_type
class AnycastEipAddressAttachmentArgs:
    def __init__(__self__, *,
                 anycast_id: pulumi.Input[str],
                 bind_instance_id: pulumi.Input[str],
                 bind_instance_region_id: pulumi.Input[str],
                 bind_instance_type: pulumi.Input[str]):
        """
        The set of arguments for constructing a AnycastEipAddressAttachment resource.
        :param pulumi.Input[str] anycast_id: The ID of Anycast EIP.
        :param pulumi.Input[str] bind_instance_id: The ID of bound instance.
        :param pulumi.Input[str] bind_instance_region_id: The region ID of bound instance.
        :param pulumi.Input[str] bind_instance_type: The type of bound instance. Valid value: `SlbInstance`.
        """
        pulumi.set(__self__, "anycast_id", anycast_id)
        pulumi.set(__self__, "bind_instance_id", bind_instance_id)
        pulumi.set(__self__, "bind_instance_region_id", bind_instance_region_id)
        pulumi.set(__self__, "bind_instance_type", bind_instance_type)

    @property
    @pulumi.getter(name="anycastId")
    def anycast_id(self) -> pulumi.Input[str]:
        """
        The ID of Anycast EIP.
        """
        return pulumi.get(self, "anycast_id")

    @anycast_id.setter
    def anycast_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "anycast_id", value)

    @property
    @pulumi.getter(name="bindInstanceId")
    def bind_instance_id(self) -> pulumi.Input[str]:
        """
        The ID of bound instance.
        """
        return pulumi.get(self, "bind_instance_id")

    @bind_instance_id.setter
    def bind_instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "bind_instance_id", value)

    @property
    @pulumi.getter(name="bindInstanceRegionId")
    def bind_instance_region_id(self) -> pulumi.Input[str]:
        """
        The region ID of bound instance.
        """
        return pulumi.get(self, "bind_instance_region_id")

    @bind_instance_region_id.setter
    def bind_instance_region_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "bind_instance_region_id", value)

    @property
    @pulumi.getter(name="bindInstanceType")
    def bind_instance_type(self) -> pulumi.Input[str]:
        """
        The type of bound instance. Valid value: `SlbInstance`.
        """
        return pulumi.get(self, "bind_instance_type")

    @bind_instance_type.setter
    def bind_instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "bind_instance_type", value)


@pulumi.input_type
class _AnycastEipAddressAttachmentState:
    def __init__(__self__, *,
                 anycast_id: Optional[pulumi.Input[str]] = None,
                 bind_instance_id: Optional[pulumi.Input[str]] = None,
                 bind_instance_region_id: Optional[pulumi.Input[str]] = None,
                 bind_instance_type: Optional[pulumi.Input[str]] = None,
                 bind_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AnycastEipAddressAttachment resources.
        :param pulumi.Input[str] anycast_id: The ID of Anycast EIP.
        :param pulumi.Input[str] bind_instance_id: The ID of bound instance.
        :param pulumi.Input[str] bind_instance_region_id: The region ID of bound instance.
        :param pulumi.Input[str] bind_instance_type: The type of bound instance. Valid value: `SlbInstance`.
        :param pulumi.Input[str] bind_time: The time of bound instance.
        """
        if anycast_id is not None:
            pulumi.set(__self__, "anycast_id", anycast_id)
        if bind_instance_id is not None:
            pulumi.set(__self__, "bind_instance_id", bind_instance_id)
        if bind_instance_region_id is not None:
            pulumi.set(__self__, "bind_instance_region_id", bind_instance_region_id)
        if bind_instance_type is not None:
            pulumi.set(__self__, "bind_instance_type", bind_instance_type)
        if bind_time is not None:
            pulumi.set(__self__, "bind_time", bind_time)

    @property
    @pulumi.getter(name="anycastId")
    def anycast_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of Anycast EIP.
        """
        return pulumi.get(self, "anycast_id")

    @anycast_id.setter
    def anycast_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "anycast_id", value)

    @property
    @pulumi.getter(name="bindInstanceId")
    def bind_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of bound instance.
        """
        return pulumi.get(self, "bind_instance_id")

    @bind_instance_id.setter
    def bind_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bind_instance_id", value)

    @property
    @pulumi.getter(name="bindInstanceRegionId")
    def bind_instance_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The region ID of bound instance.
        """
        return pulumi.get(self, "bind_instance_region_id")

    @bind_instance_region_id.setter
    def bind_instance_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bind_instance_region_id", value)

    @property
    @pulumi.getter(name="bindInstanceType")
    def bind_instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of bound instance. Valid value: `SlbInstance`.
        """
        return pulumi.get(self, "bind_instance_type")

    @bind_instance_type.setter
    def bind_instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bind_instance_type", value)

    @property
    @pulumi.getter(name="bindTime")
    def bind_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time of bound instance.
        """
        return pulumi.get(self, "bind_time")

    @bind_time.setter
    def bind_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bind_time", value)


class AnycastEipAddressAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 anycast_id: Optional[pulumi.Input[str]] = None,
                 bind_instance_id: Optional[pulumi.Input[str]] = None,
                 bind_instance_region_id: Optional[pulumi.Input[str]] = None,
                 bind_instance_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Eipanycast Anycast Eip Address Attachment resource.

        For information about Eipanycast Anycast Eip Address Attachment and how to use it, see [What is Anycast Eip Address Attachment](https://help.aliyun.com/document_detail/171857.html).

        > **NOTE:** Available in v1.113.0+.

        > **NOTE:** The following regions support currently while Slb instance support bound.
        [eu-west-1-gb33-a01,cn-hongkong-am4-c04,ap-southeast-os30-a01,us-west-ot7-a01,ap-south-in73-a01,ap-southeast-my88-a01]

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_anycast_eip_address = alicloud.eipanycast.AnycastEipAddress("exampleAnycastEipAddress", service_location="international")
        example_anycast_eip_address_attachment = alicloud.eipanycast.AnycastEipAddressAttachment("exampleAnycastEipAddressAttachment",
            anycast_id=example_anycast_eip_address.id,
            bind_instance_id="lb-j6chlcr8lffy7********",
            bind_instance_region_id="cn-hongkong",
            bind_instance_type="SlbInstance")
        ```

        ## Import

        Eipanycast Anycast Eip Address Attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment example `anycast_id`:`bind_instance_id`:`bind_instance_region_id`:`bind_instance_type`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] anycast_id: The ID of Anycast EIP.
        :param pulumi.Input[str] bind_instance_id: The ID of bound instance.
        :param pulumi.Input[str] bind_instance_region_id: The region ID of bound instance.
        :param pulumi.Input[str] bind_instance_type: The type of bound instance. Valid value: `SlbInstance`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AnycastEipAddressAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Eipanycast Anycast Eip Address Attachment resource.

        For information about Eipanycast Anycast Eip Address Attachment and how to use it, see [What is Anycast Eip Address Attachment](https://help.aliyun.com/document_detail/171857.html).

        > **NOTE:** Available in v1.113.0+.

        > **NOTE:** The following regions support currently while Slb instance support bound.
        [eu-west-1-gb33-a01,cn-hongkong-am4-c04,ap-southeast-os30-a01,us-west-ot7-a01,ap-south-in73-a01,ap-southeast-my88-a01]

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_anycast_eip_address = alicloud.eipanycast.AnycastEipAddress("exampleAnycastEipAddress", service_location="international")
        example_anycast_eip_address_attachment = alicloud.eipanycast.AnycastEipAddressAttachment("exampleAnycastEipAddressAttachment",
            anycast_id=example_anycast_eip_address.id,
            bind_instance_id="lb-j6chlcr8lffy7********",
            bind_instance_region_id="cn-hongkong",
            bind_instance_type="SlbInstance")
        ```

        ## Import

        Eipanycast Anycast Eip Address Attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment example `anycast_id`:`bind_instance_id`:`bind_instance_region_id`:`bind_instance_type`
        ```

        :param str resource_name: The name of the resource.
        :param AnycastEipAddressAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AnycastEipAddressAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 anycast_id: Optional[pulumi.Input[str]] = None,
                 bind_instance_id: Optional[pulumi.Input[str]] = None,
                 bind_instance_region_id: Optional[pulumi.Input[str]] = None,
                 bind_instance_type: Optional[pulumi.Input[str]] = None,
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
            __props__ = AnycastEipAddressAttachmentArgs.__new__(AnycastEipAddressAttachmentArgs)

            if anycast_id is None and not opts.urn:
                raise TypeError("Missing required property 'anycast_id'")
            __props__.__dict__["anycast_id"] = anycast_id
            if bind_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'bind_instance_id'")
            __props__.__dict__["bind_instance_id"] = bind_instance_id
            if bind_instance_region_id is None and not opts.urn:
                raise TypeError("Missing required property 'bind_instance_region_id'")
            __props__.__dict__["bind_instance_region_id"] = bind_instance_region_id
            if bind_instance_type is None and not opts.urn:
                raise TypeError("Missing required property 'bind_instance_type'")
            __props__.__dict__["bind_instance_type"] = bind_instance_type
            __props__.__dict__["bind_time"] = None
        super(AnycastEipAddressAttachment, __self__).__init__(
            'alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            anycast_id: Optional[pulumi.Input[str]] = None,
            bind_instance_id: Optional[pulumi.Input[str]] = None,
            bind_instance_region_id: Optional[pulumi.Input[str]] = None,
            bind_instance_type: Optional[pulumi.Input[str]] = None,
            bind_time: Optional[pulumi.Input[str]] = None) -> 'AnycastEipAddressAttachment':
        """
        Get an existing AnycastEipAddressAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] anycast_id: The ID of Anycast EIP.
        :param pulumi.Input[str] bind_instance_id: The ID of bound instance.
        :param pulumi.Input[str] bind_instance_region_id: The region ID of bound instance.
        :param pulumi.Input[str] bind_instance_type: The type of bound instance. Valid value: `SlbInstance`.
        :param pulumi.Input[str] bind_time: The time of bound instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AnycastEipAddressAttachmentState.__new__(_AnycastEipAddressAttachmentState)

        __props__.__dict__["anycast_id"] = anycast_id
        __props__.__dict__["bind_instance_id"] = bind_instance_id
        __props__.__dict__["bind_instance_region_id"] = bind_instance_region_id
        __props__.__dict__["bind_instance_type"] = bind_instance_type
        __props__.__dict__["bind_time"] = bind_time
        return AnycastEipAddressAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="anycastId")
    def anycast_id(self) -> pulumi.Output[str]:
        """
        The ID of Anycast EIP.
        """
        return pulumi.get(self, "anycast_id")

    @property
    @pulumi.getter(name="bindInstanceId")
    def bind_instance_id(self) -> pulumi.Output[str]:
        """
        The ID of bound instance.
        """
        return pulumi.get(self, "bind_instance_id")

    @property
    @pulumi.getter(name="bindInstanceRegionId")
    def bind_instance_region_id(self) -> pulumi.Output[str]:
        """
        The region ID of bound instance.
        """
        return pulumi.get(self, "bind_instance_region_id")

    @property
    @pulumi.getter(name="bindInstanceType")
    def bind_instance_type(self) -> pulumi.Output[str]:
        """
        The type of bound instance. Valid value: `SlbInstance`.
        """
        return pulumi.get(self, "bind_instance_type")

    @property
    @pulumi.getter(name="bindTime")
    def bind_time(self) -> pulumi.Output[str]:
        """
        The time of bound instance.
        """
        return pulumi.get(self, "bind_time")

