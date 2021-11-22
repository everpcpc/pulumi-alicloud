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
    'GetDomainRecordsResult',
    'AwaitableGetDomainRecordsResult',
    'get_domain_records',
    'get_domain_records_output',
]

@pulumi.output_type
class GetDomainRecordsResult:
    """
    A collection of values returned by getDomainRecords.
    """
    def __init__(__self__, domain_name=None, host_record_regex=None, id=None, ids=None, is_locked=None, line=None, output_file=None, records=None, status=None, type=None, urls=None, value_regex=None):
        if domain_name and not isinstance(domain_name, str):
            raise TypeError("Expected argument 'domain_name' to be a str")
        pulumi.set(__self__, "domain_name", domain_name)
        if host_record_regex and not isinstance(host_record_regex, str):
            raise TypeError("Expected argument 'host_record_regex' to be a str")
        pulumi.set(__self__, "host_record_regex", host_record_regex)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if is_locked and not isinstance(is_locked, bool):
            raise TypeError("Expected argument 'is_locked' to be a bool")
        pulumi.set(__self__, "is_locked", is_locked)
        if line and not isinstance(line, str):
            raise TypeError("Expected argument 'line' to be a str")
        pulumi.set(__self__, "line", line)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if records and not isinstance(records, list):
            raise TypeError("Expected argument 'records' to be a list")
        pulumi.set(__self__, "records", records)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if urls and not isinstance(urls, list):
            raise TypeError("Expected argument 'urls' to be a list")
        pulumi.set(__self__, "urls", urls)
        if value_regex and not isinstance(value_regex, str):
            raise TypeError("Expected argument 'value_regex' to be a str")
        pulumi.set(__self__, "value_regex", value_regex)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> str:
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="hostRecordRegex")
    def host_record_regex(self) -> Optional[str]:
        return pulumi.get(self, "host_record_regex")

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
    @pulumi.getter(name="isLocked")
    def is_locked(self) -> Optional[bool]:
        return pulumi.get(self, "is_locked")

    @property
    @pulumi.getter
    def line(self) -> Optional[str]:
        return pulumi.get(self, "line")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def records(self) -> Sequence['outputs.GetDomainRecordsRecordResult']:
        return pulumi.get(self, "records")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def urls(self) -> Sequence[str]:
        return pulumi.get(self, "urls")

    @property
    @pulumi.getter(name="valueRegex")
    def value_regex(self) -> Optional[str]:
        return pulumi.get(self, "value_regex")


class AwaitableGetDomainRecordsResult(GetDomainRecordsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainRecordsResult(
            domain_name=self.domain_name,
            host_record_regex=self.host_record_regex,
            id=self.id,
            ids=self.ids,
            is_locked=self.is_locked,
            line=self.line,
            output_file=self.output_file,
            records=self.records,
            status=self.status,
            type=self.type,
            urls=self.urls,
            value_regex=self.value_regex)


def get_domain_records(domain_name: Optional[str] = None,
                       host_record_regex: Optional[str] = None,
                       ids: Optional[Sequence[str]] = None,
                       is_locked: Optional[bool] = None,
                       line: Optional[str] = None,
                       output_file: Optional[str] = None,
                       status: Optional[str] = None,
                       type: Optional[str] = None,
                       value_regex: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainRecordsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['domainName'] = domain_name
    __args__['hostRecordRegex'] = host_record_regex
    __args__['ids'] = ids
    __args__['isLocked'] = is_locked
    __args__['line'] = line
    __args__['outputFile'] = output_file
    __args__['status'] = status
    __args__['type'] = type
    __args__['valueRegex'] = value_regex
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:dns/getDomainRecords:getDomainRecords', __args__, opts=opts, typ=GetDomainRecordsResult).value

    return AwaitableGetDomainRecordsResult(
        domain_name=__ret__.domain_name,
        host_record_regex=__ret__.host_record_regex,
        id=__ret__.id,
        ids=__ret__.ids,
        is_locked=__ret__.is_locked,
        line=__ret__.line,
        output_file=__ret__.output_file,
        records=__ret__.records,
        status=__ret__.status,
        type=__ret__.type,
        urls=__ret__.urls,
        value_regex=__ret__.value_regex)


@_utilities.lift_output_func(get_domain_records)
def get_domain_records_output(domain_name: Optional[pulumi.Input[str]] = None,
                              host_record_regex: Optional[pulumi.Input[Optional[str]]] = None,
                              ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              is_locked: Optional[pulumi.Input[Optional[bool]]] = None,
                              line: Optional[pulumi.Input[Optional[str]]] = None,
                              output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              status: Optional[pulumi.Input[Optional[str]]] = None,
                              type: Optional[pulumi.Input[Optional[str]]] = None,
                              value_regex: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDomainRecordsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
