# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TrailArgs', 'Trail']

@pulumi.input_type
class TrailArgs:
    def __init__(__self__, *,
                 event_rw: Optional[pulumi.Input[str]] = None,
                 is_organization_trail: Optional[pulumi.Input[bool]] = None,
                 mns_topic_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 oss_bucket_name: Optional[pulumi.Input[str]] = None,
                 oss_key_prefix: Optional[pulumi.Input[str]] = None,
                 oss_write_role_arn: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 sls_project_arn: Optional[pulumi.Input[str]] = None,
                 sls_write_role_arn: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 trail_name: Optional[pulumi.Input[str]] = None,
                 trail_region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Trail resource.
        :param pulumi.Input[str] event_rw: Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
        :param pulumi.Input[str] mns_topic_arn: Field `mns_topic_arn` has been deprecated from version 1.118.0.
        :param pulumi.Input[str] name: Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
        :param pulumi.Input[str] oss_bucket_name: The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
        :param pulumi.Input[str] oss_key_prefix: The prefix of the specified OSS bucket name. This parameter can be left empty.
        :param pulumi.Input[str] oss_write_role_arn: The unique ARN of the Oss role.
        :param pulumi.Input[str] role_name: Field `name` has been deprecated from version 1.118.0.
        :param pulumi.Input[str] sls_project_arn: The unique ARN of the Log Service project.
        :param pulumi.Input[str] sls_write_role_arn: The unique ARN of the Log Service role.
        :param pulumi.Input[str] status: The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
        :param pulumi.Input[str] trail_name: The name of the trail to be created, which must be unique for an account.
        :param pulumi.Input[str] trail_region: The regions to which the trail is applied. Valid values: `cn-beijing`, `cn-hangzhou`, and `All`. Default to `All`.
        """
        if event_rw is not None:
            pulumi.set(__self__, "event_rw", event_rw)
        if is_organization_trail is not None:
            pulumi.set(__self__, "is_organization_trail", is_organization_trail)
        if mns_topic_arn is not None:
            warnings.warn("""Field 'mns_topic_arn' has been deprecated from version 1.118.0""", DeprecationWarning)
            pulumi.log.warn("""mns_topic_arn is deprecated: Field 'mns_topic_arn' has been deprecated from version 1.118.0""")
        if mns_topic_arn is not None:
            pulumi.set(__self__, "mns_topic_arn", mns_topic_arn)
        if name is not None:
            warnings.warn("""Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.""", DeprecationWarning)
            pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.""")
        if name is not None:
            pulumi.set(__self__, "name", name)
        if oss_bucket_name is not None:
            pulumi.set(__self__, "oss_bucket_name", oss_bucket_name)
        if oss_key_prefix is not None:
            pulumi.set(__self__, "oss_key_prefix", oss_key_prefix)
        if oss_write_role_arn is not None:
            pulumi.set(__self__, "oss_write_role_arn", oss_write_role_arn)
        if role_name is not None:
            warnings.warn("""Field 'role_name' has been deprecated from version 1.118.0""", DeprecationWarning)
            pulumi.log.warn("""role_name is deprecated: Field 'role_name' has been deprecated from version 1.118.0""")
        if role_name is not None:
            pulumi.set(__self__, "role_name", role_name)
        if sls_project_arn is not None:
            pulumi.set(__self__, "sls_project_arn", sls_project_arn)
        if sls_write_role_arn is not None:
            pulumi.set(__self__, "sls_write_role_arn", sls_write_role_arn)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if trail_name is not None:
            pulumi.set(__self__, "trail_name", trail_name)
        if trail_region is not None:
            pulumi.set(__self__, "trail_region", trail_region)

    @property
    @pulumi.getter(name="eventRw")
    def event_rw(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
        """
        return pulumi.get(self, "event_rw")

    @event_rw.setter
    def event_rw(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_rw", value)

    @property
    @pulumi.getter(name="isOrganizationTrail")
    def is_organization_trail(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "is_organization_trail")

    @is_organization_trail.setter
    def is_organization_trail(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_organization_trail", value)

    @property
    @pulumi.getter(name="mnsTopicArn")
    def mns_topic_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Field `mns_topic_arn` has been deprecated from version 1.118.0.
        """
        return pulumi.get(self, "mns_topic_arn")

    @mns_topic_arn.setter
    def mns_topic_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mns_topic_arn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ossBucketName")
    def oss_bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
        """
        return pulumi.get(self, "oss_bucket_name")

    @oss_bucket_name.setter
    def oss_bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oss_bucket_name", value)

    @property
    @pulumi.getter(name="ossKeyPrefix")
    def oss_key_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The prefix of the specified OSS bucket name. This parameter can be left empty.
        """
        return pulumi.get(self, "oss_key_prefix")

    @oss_key_prefix.setter
    def oss_key_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oss_key_prefix", value)

    @property
    @pulumi.getter(name="ossWriteRoleArn")
    def oss_write_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ARN of the Oss role.
        """
        return pulumi.get(self, "oss_write_role_arn")

    @oss_write_role_arn.setter
    def oss_write_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oss_write_role_arn", value)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> Optional[pulumi.Input[str]]:
        """
        Field `name` has been deprecated from version 1.118.0.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_name", value)

    @property
    @pulumi.getter(name="slsProjectArn")
    def sls_project_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ARN of the Log Service project.
        """
        return pulumi.get(self, "sls_project_arn")

    @sls_project_arn.setter
    def sls_project_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sls_project_arn", value)

    @property
    @pulumi.getter(name="slsWriteRoleArn")
    def sls_write_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ARN of the Log Service role.
        """
        return pulumi.get(self, "sls_write_role_arn")

    @sls_write_role_arn.setter
    def sls_write_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sls_write_role_arn", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="trailName")
    def trail_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the trail to be created, which must be unique for an account.
        """
        return pulumi.get(self, "trail_name")

    @trail_name.setter
    def trail_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trail_name", value)

    @property
    @pulumi.getter(name="trailRegion")
    def trail_region(self) -> Optional[pulumi.Input[str]]:
        """
        The regions to which the trail is applied. Valid values: `cn-beijing`, `cn-hangzhou`, and `All`. Default to `All`.
        """
        return pulumi.get(self, "trail_region")

    @trail_region.setter
    def trail_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trail_region", value)


@pulumi.input_type
class _TrailState:
    def __init__(__self__, *,
                 event_rw: Optional[pulumi.Input[str]] = None,
                 is_organization_trail: Optional[pulumi.Input[bool]] = None,
                 mns_topic_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 oss_bucket_name: Optional[pulumi.Input[str]] = None,
                 oss_key_prefix: Optional[pulumi.Input[str]] = None,
                 oss_write_role_arn: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 sls_project_arn: Optional[pulumi.Input[str]] = None,
                 sls_write_role_arn: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 trail_name: Optional[pulumi.Input[str]] = None,
                 trail_region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Trail resources.
        :param pulumi.Input[str] event_rw: Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
        :param pulumi.Input[str] mns_topic_arn: Field `mns_topic_arn` has been deprecated from version 1.118.0.
        :param pulumi.Input[str] name: Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
        :param pulumi.Input[str] oss_bucket_name: The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
        :param pulumi.Input[str] oss_key_prefix: The prefix of the specified OSS bucket name. This parameter can be left empty.
        :param pulumi.Input[str] oss_write_role_arn: The unique ARN of the Oss role.
        :param pulumi.Input[str] role_name: Field `name` has been deprecated from version 1.118.0.
        :param pulumi.Input[str] sls_project_arn: The unique ARN of the Log Service project.
        :param pulumi.Input[str] sls_write_role_arn: The unique ARN of the Log Service role.
        :param pulumi.Input[str] status: The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
        :param pulumi.Input[str] trail_name: The name of the trail to be created, which must be unique for an account.
        :param pulumi.Input[str] trail_region: The regions to which the trail is applied. Valid values: `cn-beijing`, `cn-hangzhou`, and `All`. Default to `All`.
        """
        if event_rw is not None:
            pulumi.set(__self__, "event_rw", event_rw)
        if is_organization_trail is not None:
            pulumi.set(__self__, "is_organization_trail", is_organization_trail)
        if mns_topic_arn is not None:
            warnings.warn("""Field 'mns_topic_arn' has been deprecated from version 1.118.0""", DeprecationWarning)
            pulumi.log.warn("""mns_topic_arn is deprecated: Field 'mns_topic_arn' has been deprecated from version 1.118.0""")
        if mns_topic_arn is not None:
            pulumi.set(__self__, "mns_topic_arn", mns_topic_arn)
        if name is not None:
            warnings.warn("""Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.""", DeprecationWarning)
            pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.""")
        if name is not None:
            pulumi.set(__self__, "name", name)
        if oss_bucket_name is not None:
            pulumi.set(__self__, "oss_bucket_name", oss_bucket_name)
        if oss_key_prefix is not None:
            pulumi.set(__self__, "oss_key_prefix", oss_key_prefix)
        if oss_write_role_arn is not None:
            pulumi.set(__self__, "oss_write_role_arn", oss_write_role_arn)
        if role_name is not None:
            warnings.warn("""Field 'role_name' has been deprecated from version 1.118.0""", DeprecationWarning)
            pulumi.log.warn("""role_name is deprecated: Field 'role_name' has been deprecated from version 1.118.0""")
        if role_name is not None:
            pulumi.set(__self__, "role_name", role_name)
        if sls_project_arn is not None:
            pulumi.set(__self__, "sls_project_arn", sls_project_arn)
        if sls_write_role_arn is not None:
            pulumi.set(__self__, "sls_write_role_arn", sls_write_role_arn)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if trail_name is not None:
            pulumi.set(__self__, "trail_name", trail_name)
        if trail_region is not None:
            pulumi.set(__self__, "trail_region", trail_region)

    @property
    @pulumi.getter(name="eventRw")
    def event_rw(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
        """
        return pulumi.get(self, "event_rw")

    @event_rw.setter
    def event_rw(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_rw", value)

    @property
    @pulumi.getter(name="isOrganizationTrail")
    def is_organization_trail(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "is_organization_trail")

    @is_organization_trail.setter
    def is_organization_trail(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_organization_trail", value)

    @property
    @pulumi.getter(name="mnsTopicArn")
    def mns_topic_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Field `mns_topic_arn` has been deprecated from version 1.118.0.
        """
        return pulumi.get(self, "mns_topic_arn")

    @mns_topic_arn.setter
    def mns_topic_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mns_topic_arn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ossBucketName")
    def oss_bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
        """
        return pulumi.get(self, "oss_bucket_name")

    @oss_bucket_name.setter
    def oss_bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oss_bucket_name", value)

    @property
    @pulumi.getter(name="ossKeyPrefix")
    def oss_key_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The prefix of the specified OSS bucket name. This parameter can be left empty.
        """
        return pulumi.get(self, "oss_key_prefix")

    @oss_key_prefix.setter
    def oss_key_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oss_key_prefix", value)

    @property
    @pulumi.getter(name="ossWriteRoleArn")
    def oss_write_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ARN of the Oss role.
        """
        return pulumi.get(self, "oss_write_role_arn")

    @oss_write_role_arn.setter
    def oss_write_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oss_write_role_arn", value)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> Optional[pulumi.Input[str]]:
        """
        Field `name` has been deprecated from version 1.118.0.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_name", value)

    @property
    @pulumi.getter(name="slsProjectArn")
    def sls_project_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ARN of the Log Service project.
        """
        return pulumi.get(self, "sls_project_arn")

    @sls_project_arn.setter
    def sls_project_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sls_project_arn", value)

    @property
    @pulumi.getter(name="slsWriteRoleArn")
    def sls_write_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ARN of the Log Service role.
        """
        return pulumi.get(self, "sls_write_role_arn")

    @sls_write_role_arn.setter
    def sls_write_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sls_write_role_arn", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="trailName")
    def trail_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the trail to be created, which must be unique for an account.
        """
        return pulumi.get(self, "trail_name")

    @trail_name.setter
    def trail_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trail_name", value)

    @property
    @pulumi.getter(name="trailRegion")
    def trail_region(self) -> Optional[pulumi.Input[str]]:
        """
        The regions to which the trail is applied. Valid values: `cn-beijing`, `cn-hangzhou`, and `All`. Default to `All`.
        """
        return pulumi.get(self, "trail_region")

    @trail_region.setter
    def trail_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trail_region", value)


class Trail(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 event_rw: Optional[pulumi.Input[str]] = None,
                 is_organization_trail: Optional[pulumi.Input[bool]] = None,
                 mns_topic_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 oss_bucket_name: Optional[pulumi.Input[str]] = None,
                 oss_key_prefix: Optional[pulumi.Input[str]] = None,
                 oss_write_role_arn: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 sls_project_arn: Optional[pulumi.Input[str]] = None,
                 sls_write_role_arn: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 trail_name: Optional[pulumi.Input[str]] = None,
                 trail_region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a ActionTrail Trail resource. For information about alicloud actiontrail trail and how to use it, see [What is Resource Alicloud ActionTrail Trail](https://www.alibabacloud.com/help/doc-detail/28804.htm).

        > **NOTE:** Available in 1.95.0+

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        # Create a new actiontrail trail.
        default = alicloud.actiontrail.Trail("default",
            event_rw="All",
            oss_bucket_name="bucket_name",
            oss_write_role_arn="acs:ram::1182725xxxxxxxxxxx",
            trail_name="action-trail",
            trail_region="cn-hangzhou")
        ```

        ## Import

        Action trail can be imported using the id or trail_name, e.g.

        ```sh
         $ pulumi import alicloud:actiontrail/trail:Trail default abc12345678
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] event_rw: Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
        :param pulumi.Input[str] mns_topic_arn: Field `mns_topic_arn` has been deprecated from version 1.118.0.
        :param pulumi.Input[str] name: Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
        :param pulumi.Input[str] oss_bucket_name: The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
        :param pulumi.Input[str] oss_key_prefix: The prefix of the specified OSS bucket name. This parameter can be left empty.
        :param pulumi.Input[str] oss_write_role_arn: The unique ARN of the Oss role.
        :param pulumi.Input[str] role_name: Field `name` has been deprecated from version 1.118.0.
        :param pulumi.Input[str] sls_project_arn: The unique ARN of the Log Service project.
        :param pulumi.Input[str] sls_write_role_arn: The unique ARN of the Log Service role.
        :param pulumi.Input[str] status: The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
        :param pulumi.Input[str] trail_name: The name of the trail to be created, which must be unique for an account.
        :param pulumi.Input[str] trail_region: The regions to which the trail is applied. Valid values: `cn-beijing`, `cn-hangzhou`, and `All`. Default to `All`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TrailArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a ActionTrail Trail resource. For information about alicloud actiontrail trail and how to use it, see [What is Resource Alicloud ActionTrail Trail](https://www.alibabacloud.com/help/doc-detail/28804.htm).

        > **NOTE:** Available in 1.95.0+

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        # Create a new actiontrail trail.
        default = alicloud.actiontrail.Trail("default",
            event_rw="All",
            oss_bucket_name="bucket_name",
            oss_write_role_arn="acs:ram::1182725xxxxxxxxxxx",
            trail_name="action-trail",
            trail_region="cn-hangzhou")
        ```

        ## Import

        Action trail can be imported using the id or trail_name, e.g.

        ```sh
         $ pulumi import alicloud:actiontrail/trail:Trail default abc12345678
        ```

        :param str resource_name: The name of the resource.
        :param TrailArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrailArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 event_rw: Optional[pulumi.Input[str]] = None,
                 is_organization_trail: Optional[pulumi.Input[bool]] = None,
                 mns_topic_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 oss_bucket_name: Optional[pulumi.Input[str]] = None,
                 oss_key_prefix: Optional[pulumi.Input[str]] = None,
                 oss_write_role_arn: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 sls_project_arn: Optional[pulumi.Input[str]] = None,
                 sls_write_role_arn: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 trail_name: Optional[pulumi.Input[str]] = None,
                 trail_region: Optional[pulumi.Input[str]] = None,
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
            __props__ = TrailArgs.__new__(TrailArgs)

            __props__.__dict__["event_rw"] = event_rw
            __props__.__dict__["is_organization_trail"] = is_organization_trail
            if mns_topic_arn is not None and not opts.urn:
                warnings.warn("""Field 'mns_topic_arn' has been deprecated from version 1.118.0""", DeprecationWarning)
                pulumi.log.warn("""mns_topic_arn is deprecated: Field 'mns_topic_arn' has been deprecated from version 1.118.0""")
            __props__.__dict__["mns_topic_arn"] = mns_topic_arn
            if name is not None and not opts.urn:
                warnings.warn("""Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.""", DeprecationWarning)
                pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.""")
            __props__.__dict__["name"] = name
            __props__.__dict__["oss_bucket_name"] = oss_bucket_name
            __props__.__dict__["oss_key_prefix"] = oss_key_prefix
            __props__.__dict__["oss_write_role_arn"] = oss_write_role_arn
            if role_name is not None and not opts.urn:
                warnings.warn("""Field 'role_name' has been deprecated from version 1.118.0""", DeprecationWarning)
                pulumi.log.warn("""role_name is deprecated: Field 'role_name' has been deprecated from version 1.118.0""")
            __props__.__dict__["role_name"] = role_name
            __props__.__dict__["sls_project_arn"] = sls_project_arn
            __props__.__dict__["sls_write_role_arn"] = sls_write_role_arn
            __props__.__dict__["status"] = status
            __props__.__dict__["trail_name"] = trail_name
            __props__.__dict__["trail_region"] = trail_region
        super(Trail, __self__).__init__(
            'alicloud:actiontrail/trail:Trail',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            event_rw: Optional[pulumi.Input[str]] = None,
            is_organization_trail: Optional[pulumi.Input[bool]] = None,
            mns_topic_arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            oss_bucket_name: Optional[pulumi.Input[str]] = None,
            oss_key_prefix: Optional[pulumi.Input[str]] = None,
            oss_write_role_arn: Optional[pulumi.Input[str]] = None,
            role_name: Optional[pulumi.Input[str]] = None,
            sls_project_arn: Optional[pulumi.Input[str]] = None,
            sls_write_role_arn: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            trail_name: Optional[pulumi.Input[str]] = None,
            trail_region: Optional[pulumi.Input[str]] = None) -> 'Trail':
        """
        Get an existing Trail resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] event_rw: Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
        :param pulumi.Input[str] mns_topic_arn: Field `mns_topic_arn` has been deprecated from version 1.118.0.
        :param pulumi.Input[str] name: Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
        :param pulumi.Input[str] oss_bucket_name: The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
        :param pulumi.Input[str] oss_key_prefix: The prefix of the specified OSS bucket name. This parameter can be left empty.
        :param pulumi.Input[str] oss_write_role_arn: The unique ARN of the Oss role.
        :param pulumi.Input[str] role_name: Field `name` has been deprecated from version 1.118.0.
        :param pulumi.Input[str] sls_project_arn: The unique ARN of the Log Service project.
        :param pulumi.Input[str] sls_write_role_arn: The unique ARN of the Log Service role.
        :param pulumi.Input[str] status: The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
        :param pulumi.Input[str] trail_name: The name of the trail to be created, which must be unique for an account.
        :param pulumi.Input[str] trail_region: The regions to which the trail is applied. Valid values: `cn-beijing`, `cn-hangzhou`, and `All`. Default to `All`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TrailState.__new__(_TrailState)

        __props__.__dict__["event_rw"] = event_rw
        __props__.__dict__["is_organization_trail"] = is_organization_trail
        __props__.__dict__["mns_topic_arn"] = mns_topic_arn
        __props__.__dict__["name"] = name
        __props__.__dict__["oss_bucket_name"] = oss_bucket_name
        __props__.__dict__["oss_key_prefix"] = oss_key_prefix
        __props__.__dict__["oss_write_role_arn"] = oss_write_role_arn
        __props__.__dict__["role_name"] = role_name
        __props__.__dict__["sls_project_arn"] = sls_project_arn
        __props__.__dict__["sls_write_role_arn"] = sls_write_role_arn
        __props__.__dict__["status"] = status
        __props__.__dict__["trail_name"] = trail_name
        __props__.__dict__["trail_region"] = trail_region
        return Trail(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="eventRw")
    def event_rw(self) -> pulumi.Output[Optional[str]]:
        """
        Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
        """
        return pulumi.get(self, "event_rw")

    @property
    @pulumi.getter(name="isOrganizationTrail")
    def is_organization_trail(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "is_organization_trail")

    @property
    @pulumi.getter(name="mnsTopicArn")
    def mns_topic_arn(self) -> pulumi.Output[Optional[str]]:
        """
        Field `mns_topic_arn` has been deprecated from version 1.118.0.
        """
        return pulumi.get(self, "mns_topic_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ossBucketName")
    def oss_bucket_name(self) -> pulumi.Output[Optional[str]]:
        """
        The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
        """
        return pulumi.get(self, "oss_bucket_name")

    @property
    @pulumi.getter(name="ossKeyPrefix")
    def oss_key_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        The prefix of the specified OSS bucket name. This parameter can be left empty.
        """
        return pulumi.get(self, "oss_key_prefix")

    @property
    @pulumi.getter(name="ossWriteRoleArn")
    def oss_write_role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The unique ARN of the Oss role.
        """
        return pulumi.get(self, "oss_write_role_arn")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Output[str]:
        """
        Field `name` has been deprecated from version 1.118.0.
        """
        return pulumi.get(self, "role_name")

    @property
    @pulumi.getter(name="slsProjectArn")
    def sls_project_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The unique ARN of the Log Service project.
        """
        return pulumi.get(self, "sls_project_arn")

    @property
    @pulumi.getter(name="slsWriteRoleArn")
    def sls_write_role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The unique ARN of the Log Service role.
        """
        return pulumi.get(self, "sls_write_role_arn")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="trailName")
    def trail_name(self) -> pulumi.Output[str]:
        """
        The name of the trail to be created, which must be unique for an account.
        """
        return pulumi.get(self, "trail_name")

    @property
    @pulumi.getter(name="trailRegion")
    def trail_region(self) -> pulumi.Output[Optional[str]]:
        """
        The regions to which the trail is applied. Valid values: `cn-beijing`, `cn-hangzhou`, and `All`. Default to `All`.
        """
        return pulumi.get(self, "trail_region")

