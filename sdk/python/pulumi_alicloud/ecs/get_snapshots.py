# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetSnapshotsResult:
    """
    A collection of values returned by getSnapshots.
    """
    def __init__(__self__, disk_id=None, encrypted=None, id=None, ids=None, instance_id=None, name_regex=None, names=None, output_file=None, snapshots=None, source_disk_type=None, status=None, tags=None, type=None, usage=None):
        if disk_id and not isinstance(disk_id, str):
            raise TypeError("Expected argument 'disk_id' to be a str")
        __self__.disk_id = disk_id
        if encrypted and not isinstance(encrypted, bool):
            raise TypeError("Expected argument 'encrypted' to be a bool")
        __self__.encrypted = encrypted
        """
        Whether the snapshot is encrypted or not.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of snapshot IDs.
        """
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        __self__.instance_id = instance_id
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of snapshots names.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if snapshots and not isinstance(snapshots, list):
            raise TypeError("Expected argument 'snapshots' to be a list")
        __self__.snapshots = snapshots
        """
        A list of snapshots. Each element contains the following attributes:
        """
        if source_disk_type and not isinstance(source_disk_type, str):
            raise TypeError("Expected argument 'source_disk_type' to be a str")
        __self__.source_disk_type = source_disk_type
        """
        Source disk attribute. Value range:
        * System
        * Data
        """
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        The snapshot status. Value range:
        * progressing
        * accomplished
        * failed
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A map of tags assigned to the snapshot.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        if usage and not isinstance(usage, str):
            raise TypeError("Expected argument 'usage' to be a str")
        __self__.usage = usage
        """
        Whether the snapshots are used to create resources or not. Value range:
        * image
        * disk
        * image_disk
        * none
        """
class AwaitableGetSnapshotsResult(GetSnapshotsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSnapshotsResult(
            disk_id=self.disk_id,
            encrypted=self.encrypted,
            id=self.id,
            ids=self.ids,
            instance_id=self.instance_id,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            snapshots=self.snapshots,
            source_disk_type=self.source_disk_type,
            status=self.status,
            tags=self.tags,
            type=self.type,
            usage=self.usage)

def get_snapshots(disk_id=None,encrypted=None,ids=None,instance_id=None,name_regex=None,output_file=None,source_disk_type=None,status=None,tags=None,type=None,usage=None,opts=None):
    """
    Use this data source to get a list of snapshot according to the specified filters in an Alibaba Cloud account.

    For information about snapshot and how to use it, see [Snapshot](https://www.alibabacloud.com/help/doc-detail/25460.html).

    > **NOTE:**  Available in 1.40.0+.

    ##  Argument Reference

    The following arguments are supported:

    * `instance_id` - (Optional) The specified instance ID.
    * `disk_id` - (Optional) The specified disk ID.
    * `encrypted` - (Optional) Queries the encrypted snapshots. Optional values:
      * true: Encrypted snapshots.
      * false: No encryption attribute limit.
      
      Default value: false.
    * `ids` - (Optional)  A list of snapshot IDs.
    * `name_regex` - (Optional) A regex string to filter results by snapshot name.
    * `status` - (Optional) The specified snapshot status.
      * The snapshot status. Optional values:
      * progressing: The snapshots are being created.
      * accomplished: The snapshots are ready to use.
      * failed: The snapshot creation failed.
      * all: All status.
      
      Default value: all.

    * `type` - (Optional) The snapshot category. Optional values:
      * auto: Auto snapshots.
      * user: Manual snapshots.
      * all: Auto and manual snapshots.
      
      Default value: all.
    * `source_disk_type` - (Optional) The type of source disk:
      * System: The snapshots are created for system disks.
      * Data: The snapshots are created for data disks.
      
    * `usage` - (Optional) The usage of the snapshot:
      * image: The snapshots are used to create custom images.
      * disk: The snapshots are used to CreateDisk.
      * mage_disk: The snapshots are used to create custom images and data disks.
      * none: The snapshots are not used yet.
    * `tags` - (Optional) A map of tags assigned to snapshots.
    * `output_file` - (Optional) The name of output file that saves the filter results.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/snapshots.html.markdown.
    """
    __args__ = dict()


    __args__['diskId'] = disk_id
    __args__['encrypted'] = encrypted
    __args__['ids'] = ids
    __args__['instanceId'] = instance_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['sourceDiskType'] = source_disk_type
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['type'] = type
    __args__['usage'] = usage
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:ecs/getSnapshots:getSnapshots', __args__, opts=opts).value

    return AwaitableGetSnapshotsResult(
        disk_id=__ret__.get('diskId'),
        encrypted=__ret__.get('encrypted'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        instance_id=__ret__.get('instanceId'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        snapshots=__ret__.get('snapshots'),
        source_disk_type=__ret__.get('sourceDiskType'),
        status=__ret__.get('status'),
        tags=__ret__.get('tags'),
        type=__ret__.get('type'),
        usage=__ret__.get('usage'))
