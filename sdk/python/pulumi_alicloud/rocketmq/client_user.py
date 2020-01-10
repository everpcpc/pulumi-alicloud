# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ClientUser(pulumi.CustomResource):
    bandwidth: pulumi.Output[float]
    """
    The SAG APP bandwidth that the user can use. Unit: Kbit/s. Maximum value: 2000 Kbit/s.
    """
    client_ip: pulumi.Output[str]
    """
    The IP address of the SAG APP. If you specify this parameter, the current account always uses the specified IP address.Note The IP address must be in the private CIDR block of the SAG client.If you do not specify this parameter, the system automatically allocates an IP address from the private CIDR block of the SAG client. In this case, each re-connection uses a different IP address.
    """
    kms_encrypted_password: pulumi.Output[str]
    kms_encryption_context: pulumi.Output[dict]
    password: pulumi.Output[str]
    """
    The password used to log on to the SAG APP.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
    """
    sag_id: pulumi.Output[str]
    """
    The ID of the SAG instance created for the SAG APP.
    """
    user_mail: pulumi.Output[str]
    """
    The email address of the user. The administrator uses this address to send the account information for logging on to the APP to the user.
    """
    user_name: pulumi.Output[str]
    """
    The user name. User names in the same SAG APP must be unique.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
    """
    def __init__(__self__, resource_name, opts=None, bandwidth=None, client_ip=None, kms_encrypted_password=None, kms_encryption_context=None, password=None, sag_id=None, user_mail=None, user_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Sag ClientUser resource. This topic describes how to manage accounts as an administrator. After you configure the network, you can create multiple accounts and distribute them to end users so that clients can access Alibaba Cloud.
        
        For information about Sag ClientUser and how to use it, see [What is Sag ClientUser](https://www.alibabacloud.com/help/doc-detail/108326.htm).
        
        > **NOTE:** Available in 1.65.0+
        
        > **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] bandwidth: The SAG APP bandwidth that the user can use. Unit: Kbit/s. Maximum value: 2000 Kbit/s.
        :param pulumi.Input[str] client_ip: The IP address of the SAG APP. If you specify this parameter, the current account always uses the specified IP address.Note The IP address must be in the private CIDR block of the SAG client.If you do not specify this parameter, the system automatically allocates an IP address from the private CIDR block of the SAG client. In this case, each re-connection uses a different IP address.
        :param pulumi.Input[str] password: The password used to log on to the SAG APP.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
        :param pulumi.Input[str] sag_id: The ID of the SAG instance created for the SAG APP.
        :param pulumi.Input[str] user_mail: The email address of the user. The administrator uses this address to send the account information for logging on to the APP to the user.
        :param pulumi.Input[str] user_name: The user name. User names in the same SAG APP must be unique.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/sag_client_user.html.markdown.
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

            if bandwidth is None:
                raise TypeError("Missing required property 'bandwidth'")
            __props__['bandwidth'] = bandwidth
            __props__['client_ip'] = client_ip
            __props__['kms_encrypted_password'] = kms_encrypted_password
            __props__['kms_encryption_context'] = kms_encryption_context
            __props__['password'] = password
            if sag_id is None:
                raise TypeError("Missing required property 'sag_id'")
            __props__['sag_id'] = sag_id
            if user_mail is None:
                raise TypeError("Missing required property 'user_mail'")
            __props__['user_mail'] = user_mail
            __props__['user_name'] = user_name
        super(ClientUser, __self__).__init__(
            'alicloud:rocketmq/clientUser:ClientUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bandwidth=None, client_ip=None, kms_encrypted_password=None, kms_encryption_context=None, password=None, sag_id=None, user_mail=None, user_name=None):
        """
        Get an existing ClientUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] bandwidth: The SAG APP bandwidth that the user can use. Unit: Kbit/s. Maximum value: 2000 Kbit/s.
        :param pulumi.Input[str] client_ip: The IP address of the SAG APP. If you specify this parameter, the current account always uses the specified IP address.Note The IP address must be in the private CIDR block of the SAG client.If you do not specify this parameter, the system automatically allocates an IP address from the private CIDR block of the SAG client. In this case, each re-connection uses a different IP address.
        :param pulumi.Input[str] password: The password used to log on to the SAG APP.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
        :param pulumi.Input[str] sag_id: The ID of the SAG instance created for the SAG APP.
        :param pulumi.Input[str] user_mail: The email address of the user. The administrator uses this address to send the account information for logging on to the APP to the user.
        :param pulumi.Input[str] user_name: The user name. User names in the same SAG APP must be unique.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/sag_client_user.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["bandwidth"] = bandwidth
        __props__["client_ip"] = client_ip
        __props__["kms_encrypted_password"] = kms_encrypted_password
        __props__["kms_encryption_context"] = kms_encryption_context
        __props__["password"] = password
        __props__["sag_id"] = sag_id
        __props__["user_mail"] = user_mail
        __props__["user_name"] = user_name
        return ClientUser(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

