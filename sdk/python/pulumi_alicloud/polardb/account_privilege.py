# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AccountPrivilege(pulumi.CustomResource):
    account_name: pulumi.Output[str]
    """
    A specified account name.
    """
    account_privilege: pulumi.Output[str]
    """
    The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".
    """
    db_cluster_id: pulumi.Output[str]
    """
    The Id of cluster in which account belongs.
    """
    db_names: pulumi.Output[list]
    """
    List of specified database name.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, account_privilege=None, db_cluster_id=None, db_names=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a PolarDB account privilege resource and used to grant several database some access privilege. A database can be granted by multiple account.
        
        > **NOTE:** Available in v1.67.0+.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: A specified account name.
        :param pulumi.Input[str] account_privilege: The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".
        :param pulumi.Input[str] db_cluster_id: The Id of cluster in which account belongs.
        :param pulumi.Input[list] db_names: List of specified database name.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/polardb_account_privilege.html.markdown.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['account_privilege'] = account_privilege
            if db_cluster_id is None:
                raise TypeError("Missing required property 'db_cluster_id'")
            __props__['db_cluster_id'] = db_cluster_id
            if db_names is None:
                raise TypeError("Missing required property 'db_names'")
            __props__['db_names'] = db_names
        super(AccountPrivilege, __self__).__init__(
            'alicloud:polardb/accountPrivilege:AccountPrivilege',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_name=None, account_privilege=None, db_cluster_id=None, db_names=None):
        """
        Get an existing AccountPrivilege resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: A specified account name.
        :param pulumi.Input[str] account_privilege: The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".
        :param pulumi.Input[str] db_cluster_id: The Id of cluster in which account belongs.
        :param pulumi.Input[list] db_names: List of specified database name.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/polardb_account_privilege.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["account_name"] = account_name
        __props__["account_privilege"] = account_privilege
        __props__["db_cluster_id"] = db_cluster_id
        __props__["db_names"] = db_names
        return AccountPrivilege(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

