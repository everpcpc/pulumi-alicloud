# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetImagesResult',
    'AwaitableGetImagesResult',
    'get_images',
    'get_images_output',
]

@pulumi.output_type
class GetImagesResult:
    """
    A collection of values returned by getImages.
    """
    def __init__(__self__, id=None, ids=None, image_type=None, images=None, name_regex=None, names=None, output_file=None, platform=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if image_type and not isinstance(image_type, str):
            raise TypeError("Expected argument 'image_type' to be a str")
        pulumi.set(__self__, "image_type", image_type)
        if images and not isinstance(images, list):
            raise TypeError("Expected argument 'images' to be a list")
        pulumi.set(__self__, "images", images)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if platform and not isinstance(platform, str):
            raise TypeError("Expected argument 'platform' to be a str")
        pulumi.set(__self__, "platform", platform)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="imageType")
    def image_type(self) -> Optional[str]:
        return pulumi.get(self, "image_type")

    @property
    @pulumi.getter
    def images(self) -> Sequence['outputs.GetImagesImageResult']:
        return pulumi.get(self, "images")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def platform(self) -> Optional[str]:
        return pulumi.get(self, "platform")


class AwaitableGetImagesResult(GetImagesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetImagesResult(
            id=self.id,
            ids=self.ids,
            image_type=self.image_type,
            images=self.images,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            platform=self.platform)


def get_images(ids: Optional[Sequence[str]] = None,
               image_type: Optional[str] = None,
               name_regex: Optional[str] = None,
               output_file: Optional[str] = None,
               platform: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetImagesResult:
    """
    Use this data source to access information about an existing resource.

    :param str image_type: The type of the image. Valid values: `app`, `custom`, `system`.
           * `system`: operating system (OS) image.
           * `app`: application image.
           * `custom`: custom image.
    :param str platform: The platform of Plan supported.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['imageType'] = image_type
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['platform'] = platform
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:simpleapplicationserver/getImages:getImages', __args__, opts=opts, typ=GetImagesResult).value

    return AwaitableGetImagesResult(
        id=__ret__.id,
        ids=__ret__.ids,
        image_type=__ret__.image_type,
        images=__ret__.images,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        platform=__ret__.platform)


@_utilities.lift_output_func(get_images)
def get_images_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                      image_type: Optional[pulumi.Input[Optional[str]]] = None,
                      name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                      output_file: Optional[pulumi.Input[Optional[str]]] = None,
                      platform: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetImagesResult]:
    """
    Use this data source to access information about an existing resource.

    :param str image_type: The type of the image. Valid values: `app`, `custom`, `system`.
           * `system`: operating system (OS) image.
           * `app`: application image.
           * `custom`: custom image.
    :param str platform: The platform of Plan supported.
    """
    ...
