# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ImageCacheArgs', 'ImageCache']

@pulumi.input_type
class ImageCacheArgs:
    def __init__(__self__, *,
                 image_cache_name: pulumi.Input[str],
                 images: pulumi.Input[Sequence[pulumi.Input[str]]],
                 security_group_id: pulumi.Input[str],
                 vswitch_id: pulumi.Input[str],
                 eip_instance_id: Optional[pulumi.Input[str]] = None,
                 image_cache_size: Optional[pulumi.Input[int]] = None,
                 image_registry_credentials: Optional[pulumi.Input[Sequence[pulumi.Input['ImageCacheImageRegistryCredentialArgs']]]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ImageCache resource.
        :param pulumi.Input[str] image_cache_name: The name of the image cache.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] images: The images to be cached. The image name must be versioned.
        :param pulumi.Input[str] security_group_id: The ID of the security group. You do not need to specify the same security group as the container group.
        :param pulumi.Input[str] vswitch_id: The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
        :param pulumi.Input[str] eip_instance_id: The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
        :param pulumi.Input[int] image_cache_size: The size of the image cache. Default to `20`. Unit: GiB.
        :param pulumi.Input[Sequence[pulumi.Input['ImageCacheImageRegistryCredentialArgs']]] image_registry_credentials: The Image Registry parameters about the image to be cached.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[int] retention_days: The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
        :param pulumi.Input[str] zone_id: The zone id to cache image.
        """
        pulumi.set(__self__, "image_cache_name", image_cache_name)
        pulumi.set(__self__, "images", images)
        pulumi.set(__self__, "security_group_id", security_group_id)
        pulumi.set(__self__, "vswitch_id", vswitch_id)
        if eip_instance_id is not None:
            pulumi.set(__self__, "eip_instance_id", eip_instance_id)
        if image_cache_size is not None:
            pulumi.set(__self__, "image_cache_size", image_cache_size)
        if image_registry_credentials is not None:
            pulumi.set(__self__, "image_registry_credentials", image_registry_credentials)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="imageCacheName")
    def image_cache_name(self) -> pulumi.Input[str]:
        """
        The name of the image cache.
        """
        return pulumi.get(self, "image_cache_name")

    @image_cache_name.setter
    def image_cache_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "image_cache_name", value)

    @property
    @pulumi.getter
    def images(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The images to be cached. The image name must be versioned.
        """
        return pulumi.get(self, "images")

    @images.setter
    def images(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "images", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Input[str]:
        """
        The ID of the security group. You do not need to specify the same security group as the container group.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Input[str]:
        """
        The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vswitch_id", value)

    @property
    @pulumi.getter(name="eipInstanceId")
    def eip_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
        """
        return pulumi.get(self, "eip_instance_id")

    @eip_instance_id.setter
    def eip_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip_instance_id", value)

    @property
    @pulumi.getter(name="imageCacheSize")
    def image_cache_size(self) -> Optional[pulumi.Input[int]]:
        """
        The size of the image cache. Default to `20`. Unit: GiB.
        """
        return pulumi.get(self, "image_cache_size")

    @image_cache_size.setter
    def image_cache_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "image_cache_size", value)

    @property
    @pulumi.getter(name="imageRegistryCredentials")
    def image_registry_credentials(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ImageCacheImageRegistryCredentialArgs']]]]:
        """
        The Image Registry parameters about the image to be cached.
        """
        return pulumi.get(self, "image_registry_credentials")

    @image_registry_credentials.setter
    def image_registry_credentials(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ImageCacheImageRegistryCredentialArgs']]]]):
        pulumi.set(self, "image_registry_credentials", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The zone id to cache image.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


@pulumi.input_type
class _ImageCacheState:
    def __init__(__self__, *,
                 container_group_id: Optional[pulumi.Input[str]] = None,
                 eip_instance_id: Optional[pulumi.Input[str]] = None,
                 image_cache_name: Optional[pulumi.Input[str]] = None,
                 image_cache_size: Optional[pulumi.Input[int]] = None,
                 image_registry_credentials: Optional[pulumi.Input[Sequence[pulumi.Input['ImageCacheImageRegistryCredentialArgs']]]] = None,
                 images: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ImageCache resources.
        :param pulumi.Input[str] container_group_id: The ID of the container group job that is used to create the image cache.
        :param pulumi.Input[str] eip_instance_id: The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
        :param pulumi.Input[str] image_cache_name: The name of the image cache.
        :param pulumi.Input[int] image_cache_size: The size of the image cache. Default to `20`. Unit: GiB.
        :param pulumi.Input[Sequence[pulumi.Input['ImageCacheImageRegistryCredentialArgs']]] image_registry_credentials: The Image Registry parameters about the image to be cached.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] images: The images to be cached. The image name must be versioned.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[int] retention_days: The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
        :param pulumi.Input[str] security_group_id: The ID of the security group. You do not need to specify the same security group as the container group.
        :param pulumi.Input[str] status: The status of the image cache.
        :param pulumi.Input[str] vswitch_id: The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
        :param pulumi.Input[str] zone_id: The zone id to cache image.
        """
        if container_group_id is not None:
            pulumi.set(__self__, "container_group_id", container_group_id)
        if eip_instance_id is not None:
            pulumi.set(__self__, "eip_instance_id", eip_instance_id)
        if image_cache_name is not None:
            pulumi.set(__self__, "image_cache_name", image_cache_name)
        if image_cache_size is not None:
            pulumi.set(__self__, "image_cache_size", image_cache_size)
        if image_registry_credentials is not None:
            pulumi.set(__self__, "image_registry_credentials", image_registry_credentials)
        if images is not None:
            pulumi.set(__self__, "images", images)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="containerGroupId")
    def container_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the container group job that is used to create the image cache.
        """
        return pulumi.get(self, "container_group_id")

    @container_group_id.setter
    def container_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_group_id", value)

    @property
    @pulumi.getter(name="eipInstanceId")
    def eip_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
        """
        return pulumi.get(self, "eip_instance_id")

    @eip_instance_id.setter
    def eip_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip_instance_id", value)

    @property
    @pulumi.getter(name="imageCacheName")
    def image_cache_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the image cache.
        """
        return pulumi.get(self, "image_cache_name")

    @image_cache_name.setter
    def image_cache_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_cache_name", value)

    @property
    @pulumi.getter(name="imageCacheSize")
    def image_cache_size(self) -> Optional[pulumi.Input[int]]:
        """
        The size of the image cache. Default to `20`. Unit: GiB.
        """
        return pulumi.get(self, "image_cache_size")

    @image_cache_size.setter
    def image_cache_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "image_cache_size", value)

    @property
    @pulumi.getter(name="imageRegistryCredentials")
    def image_registry_credentials(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ImageCacheImageRegistryCredentialArgs']]]]:
        """
        The Image Registry parameters about the image to be cached.
        """
        return pulumi.get(self, "image_registry_credentials")

    @image_registry_credentials.setter
    def image_registry_credentials(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ImageCacheImageRegistryCredentialArgs']]]]):
        pulumi.set(self, "image_registry_credentials", value)

    @property
    @pulumi.getter
    def images(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The images to be cached. The image name must be versioned.
        """
        return pulumi.get(self, "images")

    @images.setter
    def images(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "images", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the security group. You do not need to specify the same security group as the container group.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the image cache.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The zone id to cache image.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class ImageCache(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eip_instance_id: Optional[pulumi.Input[str]] = None,
                 image_cache_name: Optional[pulumi.Input[str]] = None,
                 image_cache_size: Optional[pulumi.Input[int]] = None,
                 image_registry_credentials: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ImageCacheImageRegistryCredentialArgs']]]]] = None,
                 images: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        An ECI Image Cache can help user to solve the time-consuming problem of image pull. For information about Alicloud ECI Image Cache and how to use it, see [What is Resource Alicloud ECI Image Cache](https://www.alibabacloud.com/help/doc-detail/146891.htm).

        > **NOTE:** Available in v1.89.0+.

        > **NOTE:** Each image cache corresponds to a snapshot, and the user does not delete the snapshot directly, otherwise the cache will fail.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.eci.ImageCache("example",
            eip_instance_id="eip-uf60c7cqb2pcrkgxhxxxx",
            image_cache_name="tf-test",
            images=["registry.cn-beijing.aliyuncs.com/sceneplatform/sae-image-xxxx:latest"],
            security_group_id="sg-2zeef68b66fxxxx",
            vswitch_id="vsw-2zef9k7ng82xxxx")
        ```

        ## Import

        ECI Image Cache can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:eci/imageCache:ImageCache example abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eip_instance_id: The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
        :param pulumi.Input[str] image_cache_name: The name of the image cache.
        :param pulumi.Input[int] image_cache_size: The size of the image cache. Default to `20`. Unit: GiB.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ImageCacheImageRegistryCredentialArgs']]]] image_registry_credentials: The Image Registry parameters about the image to be cached.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] images: The images to be cached. The image name must be versioned.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[int] retention_days: The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
        :param pulumi.Input[str] security_group_id: The ID of the security group. You do not need to specify the same security group as the container group.
        :param pulumi.Input[str] vswitch_id: The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
        :param pulumi.Input[str] zone_id: The zone id to cache image.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ImageCacheArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An ECI Image Cache can help user to solve the time-consuming problem of image pull. For information about Alicloud ECI Image Cache and how to use it, see [What is Resource Alicloud ECI Image Cache](https://www.alibabacloud.com/help/doc-detail/146891.htm).

        > **NOTE:** Available in v1.89.0+.

        > **NOTE:** Each image cache corresponds to a snapshot, and the user does not delete the snapshot directly, otherwise the cache will fail.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.eci.ImageCache("example",
            eip_instance_id="eip-uf60c7cqb2pcrkgxhxxxx",
            image_cache_name="tf-test",
            images=["registry.cn-beijing.aliyuncs.com/sceneplatform/sae-image-xxxx:latest"],
            security_group_id="sg-2zeef68b66fxxxx",
            vswitch_id="vsw-2zef9k7ng82xxxx")
        ```

        ## Import

        ECI Image Cache can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:eci/imageCache:ImageCache example abc123456
        ```

        :param str resource_name: The name of the resource.
        :param ImageCacheArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ImageCacheArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eip_instance_id: Optional[pulumi.Input[str]] = None,
                 image_cache_name: Optional[pulumi.Input[str]] = None,
                 image_cache_size: Optional[pulumi.Input[int]] = None,
                 image_registry_credentials: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ImageCacheImageRegistryCredentialArgs']]]]] = None,
                 images: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ImageCacheArgs.__new__(ImageCacheArgs)

            __props__.__dict__["eip_instance_id"] = eip_instance_id
            if image_cache_name is None and not opts.urn:
                raise TypeError("Missing required property 'image_cache_name'")
            __props__.__dict__["image_cache_name"] = image_cache_name
            __props__.__dict__["image_cache_size"] = image_cache_size
            __props__.__dict__["image_registry_credentials"] = image_registry_credentials
            if images is None and not opts.urn:
                raise TypeError("Missing required property 'images'")
            __props__.__dict__["images"] = images
            __props__.__dict__["resource_group_id"] = resource_group_id
            __props__.__dict__["retention_days"] = retention_days
            if security_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_group_id'")
            __props__.__dict__["security_group_id"] = security_group_id
            if vswitch_id is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__.__dict__["vswitch_id"] = vswitch_id
            __props__.__dict__["zone_id"] = zone_id
            __props__.__dict__["container_group_id"] = None
            __props__.__dict__["status"] = None
        super(ImageCache, __self__).__init__(
            'alicloud:eci/imageCache:ImageCache',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            container_group_id: Optional[pulumi.Input[str]] = None,
            eip_instance_id: Optional[pulumi.Input[str]] = None,
            image_cache_name: Optional[pulumi.Input[str]] = None,
            image_cache_size: Optional[pulumi.Input[int]] = None,
            image_registry_credentials: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ImageCacheImageRegistryCredentialArgs']]]]] = None,
            images: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            retention_days: Optional[pulumi.Input[int]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'ImageCache':
        """
        Get an existing ImageCache resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_group_id: The ID of the container group job that is used to create the image cache.
        :param pulumi.Input[str] eip_instance_id: The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
        :param pulumi.Input[str] image_cache_name: The name of the image cache.
        :param pulumi.Input[int] image_cache_size: The size of the image cache. Default to `20`. Unit: GiB.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ImageCacheImageRegistryCredentialArgs']]]] image_registry_credentials: The Image Registry parameters about the image to be cached.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] images: The images to be cached. The image name must be versioned.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[int] retention_days: The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
        :param pulumi.Input[str] security_group_id: The ID of the security group. You do not need to specify the same security group as the container group.
        :param pulumi.Input[str] status: The status of the image cache.
        :param pulumi.Input[str] vswitch_id: The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
        :param pulumi.Input[str] zone_id: The zone id to cache image.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ImageCacheState.__new__(_ImageCacheState)

        __props__.__dict__["container_group_id"] = container_group_id
        __props__.__dict__["eip_instance_id"] = eip_instance_id
        __props__.__dict__["image_cache_name"] = image_cache_name
        __props__.__dict__["image_cache_size"] = image_cache_size
        __props__.__dict__["image_registry_credentials"] = image_registry_credentials
        __props__.__dict__["images"] = images
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["retention_days"] = retention_days
        __props__.__dict__["security_group_id"] = security_group_id
        __props__.__dict__["status"] = status
        __props__.__dict__["vswitch_id"] = vswitch_id
        __props__.__dict__["zone_id"] = zone_id
        return ImageCache(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="containerGroupId")
    def container_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the container group job that is used to create the image cache.
        """
        return pulumi.get(self, "container_group_id")

    @property
    @pulumi.getter(name="eipInstanceId")
    def eip_instance_id(self) -> pulumi.Output[Optional[str]]:
        """
        The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
        """
        return pulumi.get(self, "eip_instance_id")

    @property
    @pulumi.getter(name="imageCacheName")
    def image_cache_name(self) -> pulumi.Output[str]:
        """
        The name of the image cache.
        """
        return pulumi.get(self, "image_cache_name")

    @property
    @pulumi.getter(name="imageCacheSize")
    def image_cache_size(self) -> pulumi.Output[Optional[int]]:
        """
        The size of the image cache. Default to `20`. Unit: GiB.
        """
        return pulumi.get(self, "image_cache_size")

    @property
    @pulumi.getter(name="imageRegistryCredentials")
    def image_registry_credentials(self) -> pulumi.Output[Optional[Sequence['outputs.ImageCacheImageRegistryCredential']]]:
        """
        The Image Registry parameters about the image to be cached.
        """
        return pulumi.get(self, "image_registry_credentials")

    @property
    @pulumi.getter
    def images(self) -> pulumi.Output[Sequence[str]]:
        """
        The images to be cached. The image name must be versioned.
        """
        return pulumi.get(self, "images")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> pulumi.Output[Optional[int]]:
        """
        The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
        """
        return pulumi.get(self, "retention_days")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the security group. You do not need to specify the same security group as the container group.
        """
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the image cache.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[str]:
        """
        The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
        """
        return pulumi.get(self, "vswitch_id")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[Optional[str]]:
        """
        The zone id to cache image.
        """
        return pulumi.get(self, "zone_id")

