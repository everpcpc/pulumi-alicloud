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
    'GetEipsResult',
    'AwaitableGetEipsResult',
    'get_eips',
    'get_eips_output',
]

warnings.warn("""This function has been deprecated in favour of the getEipAddresses function""", DeprecationWarning)

@pulumi.output_type
class GetEipsResult:
    """
    A collection of values returned by getEips.
    """
    def __init__(__self__, address_name=None, addresses=None, associated_instance_id=None, associated_instance_type=None, dry_run=None, eips=None, enable_details=None, id=None, ids=None, include_reservation_data=None, ip_address=None, ip_addresses=None, isp=None, lock_reason=None, name_regex=None, names=None, output_file=None, payment_type=None, resource_group_id=None, segment_instance_id=None, status=None, tags=None):
        if address_name and not isinstance(address_name, str):
            raise TypeError("Expected argument 'address_name' to be a str")
        pulumi.set(__self__, "address_name", address_name)
        if addresses and not isinstance(addresses, list):
            raise TypeError("Expected argument 'addresses' to be a list")
        pulumi.set(__self__, "addresses", addresses)
        if associated_instance_id and not isinstance(associated_instance_id, str):
            raise TypeError("Expected argument 'associated_instance_id' to be a str")
        pulumi.set(__self__, "associated_instance_id", associated_instance_id)
        if associated_instance_type and not isinstance(associated_instance_type, str):
            raise TypeError("Expected argument 'associated_instance_type' to be a str")
        pulumi.set(__self__, "associated_instance_type", associated_instance_type)
        if dry_run and not isinstance(dry_run, bool):
            raise TypeError("Expected argument 'dry_run' to be a bool")
        pulumi.set(__self__, "dry_run", dry_run)
        if eips and not isinstance(eips, list):
            raise TypeError("Expected argument 'eips' to be a list")
        if eips is not None:
            warnings.warn("""Field 'eips' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'addresses' instead.""", DeprecationWarning)
            pulumi.log.warn("""eips is deprecated: Field 'eips' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'addresses' instead.""")

        pulumi.set(__self__, "eips", eips)
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if include_reservation_data and not isinstance(include_reservation_data, bool):
            raise TypeError("Expected argument 'include_reservation_data' to be a bool")
        pulumi.set(__self__, "include_reservation_data", include_reservation_data)
        if ip_address and not isinstance(ip_address, str):
            raise TypeError("Expected argument 'ip_address' to be a str")
        pulumi.set(__self__, "ip_address", ip_address)
        if ip_addresses and not isinstance(ip_addresses, list):
            raise TypeError("Expected argument 'ip_addresses' to be a list")
        if ip_addresses is not None:
            warnings.warn("""Field 'ip_addresses' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'ip_address' instead.""", DeprecationWarning)
            pulumi.log.warn("""ip_addresses is deprecated: Field 'ip_addresses' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'ip_address' instead.""")

        pulumi.set(__self__, "ip_addresses", ip_addresses)
        if isp and not isinstance(isp, str):
            raise TypeError("Expected argument 'isp' to be a str")
        pulumi.set(__self__, "isp", isp)
        if lock_reason and not isinstance(lock_reason, str):
            raise TypeError("Expected argument 'lock_reason' to be a str")
        pulumi.set(__self__, "lock_reason", lock_reason)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if payment_type and not isinstance(payment_type, str):
            raise TypeError("Expected argument 'payment_type' to be a str")
        pulumi.set(__self__, "payment_type", payment_type)
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        if segment_instance_id and not isinstance(segment_instance_id, str):
            raise TypeError("Expected argument 'segment_instance_id' to be a str")
        pulumi.set(__self__, "segment_instance_id", segment_instance_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="addressName")
    def address_name(self) -> Optional[str]:
        return pulumi.get(self, "address_name")

    @property
    @pulumi.getter
    def addresses(self) -> Sequence['outputs.GetEipsAddressResult']:
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter(name="associatedInstanceId")
    def associated_instance_id(self) -> Optional[str]:
        return pulumi.get(self, "associated_instance_id")

    @property
    @pulumi.getter(name="associatedInstanceType")
    def associated_instance_type(self) -> Optional[str]:
        return pulumi.get(self, "associated_instance_type")

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[bool]:
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter
    def eips(self) -> Sequence['outputs.GetEipsEipResult']:
        """
        A list of EIPs. Each element contains the following attributes:
        """
        return pulumi.get(self, "eips")

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

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
        """
        (Optional) A list of EIP IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="includeReservationData")
    def include_reservation_data(self) -> Optional[bool]:
        return pulumi.get(self, "include_reservation_data")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[str]:
        """
        Public IP Address of the the EIP.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter
    def isp(self) -> Optional[str]:
        return pulumi.get(self, "isp")

    @property
    @pulumi.getter(name="lockReason")
    def lock_reason(self) -> Optional[str]:
        return pulumi.get(self, "lock_reason")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        (Optional) A list of EIP names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> Optional[str]:
        return pulumi.get(self, "payment_type")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        """
        The Id of resource group which the eips belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="segmentInstanceId")
    def segment_instance_id(self) -> Optional[str]:
        return pulumi.get(self, "segment_instance_id")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        EIP status. Possible values are: `Associating`, `Unassociating`, `InUse` and `Available`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")


class AwaitableGetEipsResult(GetEipsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEipsResult(
            address_name=self.address_name,
            addresses=self.addresses,
            associated_instance_id=self.associated_instance_id,
            associated_instance_type=self.associated_instance_type,
            dry_run=self.dry_run,
            eips=self.eips,
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            include_reservation_data=self.include_reservation_data,
            ip_address=self.ip_address,
            ip_addresses=self.ip_addresses,
            isp=self.isp,
            lock_reason=self.lock_reason,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            payment_type=self.payment_type,
            resource_group_id=self.resource_group_id,
            segment_instance_id=self.segment_instance_id,
            status=self.status,
            tags=self.tags)


def get_eips(address_name: Optional[str] = None,
             associated_instance_id: Optional[str] = None,
             associated_instance_type: Optional[str] = None,
             dry_run: Optional[bool] = None,
             enable_details: Optional[bool] = None,
             ids: Optional[Sequence[str]] = None,
             include_reservation_data: Optional[bool] = None,
             ip_address: Optional[str] = None,
             ip_addresses: Optional[Sequence[str]] = None,
             isp: Optional[str] = None,
             lock_reason: Optional[str] = None,
             name_regex: Optional[str] = None,
             output_file: Optional[str] = None,
             payment_type: Optional[str] = None,
             resource_group_id: Optional[str] = None,
             segment_instance_id: Optional[str] = None,
             status: Optional[str] = None,
             tags: Optional[Mapping[str, Any]] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEipsResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    eips_ds = alicloud.ecs.get_eips()
    pulumi.export("firstEipId", eips_ds.eips[0].id)
    ```


    :param Sequence[str] ids: A list of EIP IDs.
    :param str ip_address: Public IP Address of the the EIP.
    :param Sequence[str] ip_addresses: A list of EIP public IP addresses.
    :param str resource_group_id: The Id of resource group which the eips belongs.
    :param str status: EIP status. Possible values are: `Associating`, `Unassociating`, `InUse` and `Available`.
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    """
    pulumi.log.warn("""get_eips is deprecated: This function has been deprecated in favour of the getEipAddresses function""")
    __args__ = dict()
    __args__['addressName'] = address_name
    __args__['associatedInstanceId'] = associated_instance_id
    __args__['associatedInstanceType'] = associated_instance_type
    __args__['dryRun'] = dry_run
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['includeReservationData'] = include_reservation_data
    __args__['ipAddress'] = ip_address
    __args__['ipAddresses'] = ip_addresses
    __args__['isp'] = isp
    __args__['lockReason'] = lock_reason
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['paymentType'] = payment_type
    __args__['resourceGroupId'] = resource_group_id
    __args__['segmentInstanceId'] = segment_instance_id
    __args__['status'] = status
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:ecs/getEips:getEips', __args__, opts=opts, typ=GetEipsResult).value

    return AwaitableGetEipsResult(
        address_name=__ret__.address_name,
        addresses=__ret__.addresses,
        associated_instance_id=__ret__.associated_instance_id,
        associated_instance_type=__ret__.associated_instance_type,
        dry_run=__ret__.dry_run,
        eips=__ret__.eips,
        enable_details=__ret__.enable_details,
        id=__ret__.id,
        ids=__ret__.ids,
        include_reservation_data=__ret__.include_reservation_data,
        ip_address=__ret__.ip_address,
        ip_addresses=__ret__.ip_addresses,
        isp=__ret__.isp,
        lock_reason=__ret__.lock_reason,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        payment_type=__ret__.payment_type,
        resource_group_id=__ret__.resource_group_id,
        segment_instance_id=__ret__.segment_instance_id,
        status=__ret__.status,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_eips)
def get_eips_output(address_name: Optional[pulumi.Input[Optional[str]]] = None,
                    associated_instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                    associated_instance_type: Optional[pulumi.Input[Optional[str]]] = None,
                    dry_run: Optional[pulumi.Input[Optional[bool]]] = None,
                    enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                    ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                    include_reservation_data: Optional[pulumi.Input[Optional[bool]]] = None,
                    ip_address: Optional[pulumi.Input[Optional[str]]] = None,
                    ip_addresses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                    isp: Optional[pulumi.Input[Optional[str]]] = None,
                    lock_reason: Optional[pulumi.Input[Optional[str]]] = None,
                    name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                    output_file: Optional[pulumi.Input[Optional[str]]] = None,
                    payment_type: Optional[pulumi.Input[Optional[str]]] = None,
                    resource_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                    segment_instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                    status: Optional[pulumi.Input[Optional[str]]] = None,
                    tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEipsResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    eips_ds = alicloud.ecs.get_eips()
    pulumi.export("firstEipId", eips_ds.eips[0].id)
    ```


    :param Sequence[str] ids: A list of EIP IDs.
    :param str ip_address: Public IP Address of the the EIP.
    :param Sequence[str] ip_addresses: A list of EIP public IP addresses.
    :param str resource_group_id: The Id of resource group which the eips belongs.
    :param str status: EIP status. Possible values are: `Associating`, `Unassociating`, `InUse` and `Available`.
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    """
    pulumi.log.warn("""get_eips is deprecated: This function has been deprecated in favour of the getEipAddresses function""")
    ...
