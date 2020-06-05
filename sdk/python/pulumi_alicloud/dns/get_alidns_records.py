# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetAlidnsRecordsResult:
    """
    A collection of values returned by getAlidnsRecords.
    """
    def __init__(__self__, direction=None, domain_name=None, group_id=None, id=None, ids=None, key_word=None, lang=None, line=None, order_by=None, output_file=None, records=None, rr_key_word=None, rr_regex=None, search_mode=None, status=None, type=None, type_key_word=None, value_key_word=None, value_regex=None):
        if direction and not isinstance(direction, str):
            raise TypeError("Expected argument 'direction' to be a str")
        __self__.direction = direction
        if domain_name and not isinstance(domain_name, str):
            raise TypeError("Expected argument 'domain_name' to be a str")
        __self__.domain_name = domain_name
        """
        Name of the domain record belongs to.
        """
        if group_id and not isinstance(group_id, float):
            raise TypeError("Expected argument 'group_id' to be a float")
        __self__.group_id = group_id
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of record IDs. 
        """
        if key_word and not isinstance(key_word, str):
            raise TypeError("Expected argument 'key_word' to be a str")
        __self__.key_word = key_word
        if lang and not isinstance(lang, str):
            raise TypeError("Expected argument 'lang' to be a str")
        __self__.lang = lang
        if line and not isinstance(line, str):
            raise TypeError("Expected argument 'line' to be a str")
        __self__.line = line
        """
        ISP line of the record. 
        """
        if order_by and not isinstance(order_by, str):
            raise TypeError("Expected argument 'order_by' to be a str")
        __self__.order_by = order_by
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if records and not isinstance(records, list):
            raise TypeError("Expected argument 'records' to be a list")
        __self__.records = records
        """
        A list of records. Each element contains the following attributes:
        """
        if rr_key_word and not isinstance(rr_key_word, str):
            raise TypeError("Expected argument 'rr_key_word' to be a str")
        __self__.rr_key_word = rr_key_word
        if rr_regex and not isinstance(rr_regex, str):
            raise TypeError("Expected argument 'rr_regex' to be a str")
        __self__.rr_regex = rr_regex
        if search_mode and not isinstance(search_mode, str):
            raise TypeError("Expected argument 'search_mode' to be a str")
        __self__.search_mode = search_mode
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        Status of the record.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Type of the record.
        """
        if type_key_word and not isinstance(type_key_word, str):
            raise TypeError("Expected argument 'type_key_word' to be a str")
        __self__.type_key_word = type_key_word
        if value_key_word and not isinstance(value_key_word, str):
            raise TypeError("Expected argument 'value_key_word' to be a str")
        __self__.value_key_word = value_key_word
        if value_regex and not isinstance(value_regex, str):
            raise TypeError("Expected argument 'value_regex' to be a str")
        __self__.value_regex = value_regex
class AwaitableGetAlidnsRecordsResult(GetAlidnsRecordsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAlidnsRecordsResult(
            direction=self.direction,
            domain_name=self.domain_name,
            group_id=self.group_id,
            id=self.id,
            ids=self.ids,
            key_word=self.key_word,
            lang=self.lang,
            line=self.line,
            order_by=self.order_by,
            output_file=self.output_file,
            records=self.records,
            rr_key_word=self.rr_key_word,
            rr_regex=self.rr_regex,
            search_mode=self.search_mode,
            status=self.status,
            type=self.type,
            type_key_word=self.type_key_word,
            value_key_word=self.value_key_word,
            value_regex=self.value_regex)

def get_alidns_records(direction=None,domain_name=None,group_id=None,ids=None,key_word=None,lang=None,line=None,order_by=None,output_file=None,rr_key_word=None,rr_regex=None,search_mode=None,status=None,type=None,type_key_word=None,value_key_word=None,value_regex=None,opts=None):
    """
    This data source provides a list of Alidns Domain Records in an Alibaba Cloud account according to the specified filters.

    > **NOTE:**  Available in 1.86.0+.




    :param str direction: Sorting direction. Valid values: `DESC`,`ASC`. Default to `AESC`.
    :param str domain_name: The domain name associated to the records. 
    :param float group_id: Domain name group ID.
    :param list ids: A list of record IDs.
    :param str key_word: Keywords.
    :param str lang: User language.
    :param str line: ISP line. Valid values: `default`, `telecom`, `unicom`, `mobile`, `oversea`, `edu`, `drpeng`, `btvn`, .etc. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm) 
    :param str order_by: Sort by. Sort from newest to oldest according to the time added by resolution.
    :param str rr_key_word: The keywords recorded by the host are searched according to the `%RRKeyWord%` mode, and are not case sensitive.
    :param str rr_regex: Host record regex. 
    :param str search_mode: Search mode, Valid values: `LIKE`, `EXACT`, `ADVANCED`, `LIKE` (fuzzy), `EXACT` (accurate) search supports KeyWord field, `ADVANCED` (advanced) mode supports other fields.
    :param str status: Record status. Valid values: `ENABLE` and `DISABLE`.
    :param str type: Record type. Valid values: `A`, `NS`, `MX`, `TXT`, `CNAME`, `SRV`, `AAAA`, `REDIRECT_URL`, `FORWORD_URL` .
    :param str type_key_word: Analyze type keywords, search by full match, not case sensitive.
    :param str value_key_word: The keywords of the recorded value are searched according to the `%ValueKeyWord%` mode, and are not case sensitive.
    :param str value_regex: Host record value regex. 
    """
    __args__ = dict()


    __args__['direction'] = direction
    __args__['domainName'] = domain_name
    __args__['groupId'] = group_id
    __args__['ids'] = ids
    __args__['keyWord'] = key_word
    __args__['lang'] = lang
    __args__['line'] = line
    __args__['orderBy'] = order_by
    __args__['outputFile'] = output_file
    __args__['rrKeyWord'] = rr_key_word
    __args__['rrRegex'] = rr_regex
    __args__['searchMode'] = search_mode
    __args__['status'] = status
    __args__['type'] = type
    __args__['typeKeyWord'] = type_key_word
    __args__['valueKeyWord'] = value_key_word
    __args__['valueRegex'] = value_regex
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:dns/getAlidnsRecords:getAlidnsRecords', __args__, opts=opts).value

    return AwaitableGetAlidnsRecordsResult(
        direction=__ret__.get('direction'),
        domain_name=__ret__.get('domainName'),
        group_id=__ret__.get('groupId'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        key_word=__ret__.get('keyWord'),
        lang=__ret__.get('lang'),
        line=__ret__.get('line'),
        order_by=__ret__.get('orderBy'),
        output_file=__ret__.get('outputFile'),
        records=__ret__.get('records'),
        rr_key_word=__ret__.get('rrKeyWord'),
        rr_regex=__ret__.get('rrRegex'),
        search_mode=__ret__.get('searchMode'),
        status=__ret__.get('status'),
        type=__ret__.get('type'),
        type_key_word=__ret__.get('typeKeyWord'),
        value_key_word=__ret__.get('valueKeyWord'),
        value_regex=__ret__.get('valueRegex'))
