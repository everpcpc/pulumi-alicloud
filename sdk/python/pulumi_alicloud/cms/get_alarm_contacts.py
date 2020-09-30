# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetAlarmContactsResult',
    'AwaitableGetAlarmContactsResult',
    'get_alarm_contacts',
]

@pulumi.output_type
class GetAlarmContactsResult:
    """
    A collection of values returned by getAlarmContacts.
    """
    def __init__(__self__, chanel_type=None, chanel_value=None, contacts=None, id=None, ids=None, name_regex=None, names=None, output_file=None):
        if chanel_type and not isinstance(chanel_type, str):
            raise TypeError("Expected argument 'chanel_type' to be a str")
        pulumi.set(__self__, "chanel_type", chanel_type)
        if chanel_value and not isinstance(chanel_value, str):
            raise TypeError("Expected argument 'chanel_value' to be a str")
        pulumi.set(__self__, "chanel_value", chanel_value)
        if contacts and not isinstance(contacts, list):
            raise TypeError("Expected argument 'contacts' to be a list")
        pulumi.set(__self__, "contacts", contacts)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)

    @property
    @pulumi.getter(name="chanelType")
    def chanel_type(self) -> Optional[str]:
        return pulumi.get(self, "chanel_type")

    @property
    @pulumi.getter(name="chanelValue")
    def chanel_value(self) -> Optional[str]:
        return pulumi.get(self, "chanel_value")

    @property
    @pulumi.getter
    def contacts(self) -> List['outputs.GetAlarmContactsContactResult']:
        """
        A list of alarm contacts. Each element contains the following attributes:
        """
        return pulumi.get(self, "contacts")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> List[str]:
        """
        A list of alarm contact IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> List[str]:
        """
        A list of alarm contact names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")


class AwaitableGetAlarmContactsResult(GetAlarmContactsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAlarmContactsResult(
            chanel_type=self.chanel_type,
            chanel_value=self.chanel_value,
            contacts=self.contacts,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file)


def get_alarm_contacts(chanel_type: Optional[str] = None,
                       chanel_value: Optional[str] = None,
                       ids: Optional[List[str]] = None,
                       name_regex: Optional[str] = None,
                       output_file: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAlarmContactsResult:
    """
    Provides a list of alarm contact owned by an Alibaba Cloud account.

    > **NOTE:** Available in v1.99.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.cms.get_alarm_contacts(ids=["tf-testAccCmsAlarmContact"])
    pulumi.export("first-contact", data["alicloud_cms_alarm_contacts"]["this"]["contacts"])
    ```


    :param str chanel_type: The alarm notification method. Alarm notifications can be sent by using `Email` or `DingWebHook`.
    :param str chanel_value: The alarm notification target.
    :param List[str] ids: A list of alarm contact IDs.
    :param str name_regex: A regex string to filter results by alarm contact name.
    """
    __args__ = dict()
    __args__['chanelType'] = chanel_type
    __args__['chanelValue'] = chanel_value
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:cms/getAlarmContacts:getAlarmContacts', __args__, opts=opts, typ=GetAlarmContactsResult).value

    return AwaitableGetAlarmContactsResult(
        chanel_type=__ret__.chanel_type,
        chanel_value=__ret__.chanel_value,
        contacts=__ret__.contacts,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file)