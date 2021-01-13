# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from ._inputs import *

__all__ = ['Provider']


class Provider(pulumi.ProviderResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 assume_role: Optional[pulumi.Input[pulumi.InputType['ProviderAssumeRoleArgs']]] = None,
                 configuration_source: Optional[pulumi.Input[str]] = None,
                 ecs_role_name: Optional[pulumi.Input[str]] = None,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProviderEndpointArgs']]]]] = None,
                 fc: Optional[pulumi.Input[str]] = None,
                 log_endpoint: Optional[pulumi.Input[str]] = None,
                 mns_endpoint: Optional[pulumi.Input[str]] = None,
                 ots_instance_name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 security_token: Optional[pulumi.Input[str]] = None,
                 shared_credentials_file: Optional[pulumi.Input[str]] = None,
                 skip_region_validation: Optional[pulumi.Input[bool]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The provider type for the alicloud package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The access key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
               console.
        :param pulumi.Input[str] account_id: The account ID for some service API operations. You can retrieve this from the 'Security Settings' section of the
               Alibaba Cloud console.
        :param pulumi.Input[str] configuration_source: Use this to mark a terraform configuration file source.
        :param pulumi.Input[str] ecs_role_name: The RAM Role Name attached on a ECS instance for API operations. You can retrieve this from the 'Access Control' section
               of the Alibaba Cloud console.
        :param pulumi.Input[str] profile: The profile for API operations. If not set, the default profile created with `aliyun configure` will be used.
        :param pulumi.Input[str] region: The region where Alibaba Cloud operations will take place. Examples are cn-beijing, cn-hangzhou, eu-central-1, etc.
        :param pulumi.Input[str] secret_key: The secret key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
               console.
        :param pulumi.Input[str] security_token: security token. A security token is only required if you are using Security Token Service.
        :param pulumi.Input[str] shared_credentials_file: The path to the shared credentials file. If not set this defaults to ~/.aliyun/config.json
        :param pulumi.Input[bool] skip_region_validation: Skip static validation of region ID. Used by users of alternative AlibabaCloud-like APIs or users w/ access to regions
               that are not public (yet).
        :param pulumi.Input[str] source_ip: The access key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
               console.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if access_key is None:
                access_key = _utilities.get_env('ALICLOUD_ACCESS_KEY')
            __props__['access_key'] = access_key
            if account_id is None:
                account_id = _utilities.get_env('ALICLOUD_ACCOUNT_ID')
            __props__['account_id'] = account_id
            __props__['assume_role'] = pulumi.Output.from_input(assume_role).apply(pulumi.runtime.to_json) if assume_role is not None else None
            __props__['configuration_source'] = configuration_source
            if ecs_role_name is None:
                ecs_role_name = _utilities.get_env('ALICLOUD_ECS_ROLE_NAME')
            __props__['ecs_role_name'] = ecs_role_name
            __props__['endpoints'] = pulumi.Output.from_input(endpoints).apply(pulumi.runtime.to_json) if endpoints is not None else None
            if fc is not None and not opts.urn:
                warnings.warn("""Field 'fc' has been deprecated from provider version 1.28.0. New field 'fc' which in nested endpoints instead.""", DeprecationWarning)
                pulumi.log.warn("fc is deprecated: Field 'fc' has been deprecated from provider version 1.28.0. New field 'fc' which in nested endpoints instead.")
            __props__['fc'] = fc
            if log_endpoint is not None and not opts.urn:
                warnings.warn("""Field 'log_endpoint' has been deprecated from provider version 1.28.0. New field 'log' which in nested endpoints instead.""", DeprecationWarning)
                pulumi.log.warn("log_endpoint is deprecated: Field 'log_endpoint' has been deprecated from provider version 1.28.0. New field 'log' which in nested endpoints instead.")
            __props__['log_endpoint'] = log_endpoint
            if mns_endpoint is not None and not opts.urn:
                warnings.warn("""Field 'mns_endpoint' has been deprecated from provider version 1.28.0. New field 'mns' which in nested endpoints instead.""", DeprecationWarning)
                pulumi.log.warn("mns_endpoint is deprecated: Field 'mns_endpoint' has been deprecated from provider version 1.28.0. New field 'mns' which in nested endpoints instead.")
            __props__['mns_endpoint'] = mns_endpoint
            if ots_instance_name is not None and not opts.urn:
                warnings.warn("""Field 'ots_instance_name' has been deprecated from provider version 1.10.0. New field 'instance_name' of resource 'alicloud_ots_table' instead.""", DeprecationWarning)
                pulumi.log.warn("ots_instance_name is deprecated: Field 'ots_instance_name' has been deprecated from provider version 1.10.0. New field 'instance_name' of resource 'alicloud_ots_table' instead.")
            __props__['ots_instance_name'] = ots_instance_name
            if profile is None:
                profile = _utilities.get_env('ALICLOUD_PROFILE')
            __props__['profile'] = profile
            __props__['protocol'] = protocol
            if region is None:
                region = _utilities.get_env('ALICLOUD_REGION')
            __props__['region'] = region
            if secret_key is None:
                secret_key = _utilities.get_env('ALICLOUD_SECRET_KEY')
            __props__['secret_key'] = secret_key
            if security_token is None:
                security_token = _utilities.get_env('ALICLOUD_SECURITY_TOKEN')
            __props__['security_token'] = security_token
            if shared_credentials_file is None:
                shared_credentials_file = _utilities.get_env('ALICLOUD_SHARED_CREDENTIALS_FILE')
            __props__['shared_credentials_file'] = shared_credentials_file
            __props__['skip_region_validation'] = pulumi.Output.from_input(skip_region_validation).apply(pulumi.runtime.to_json) if skip_region_validation is not None else None
            __props__['source_ip'] = source_ip
        super(Provider, __self__).__init__(
            'alicloud',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

