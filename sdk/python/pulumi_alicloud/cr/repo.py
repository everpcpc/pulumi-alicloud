# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Repo(pulumi.CustomResource):
    detail: pulumi.Output[str]
    """
    The repository specific information. MarkDown format is supported, and the length limit is 2000.
    """
    domain_list: pulumi.Output[dict]
    """
    The repository domain list.

      * `internal` (`str`) - Domain of internal endpoint, only in some regions.
      * `public` (`str`) - Domain of public endpoint.
      * `vpc` (`str`) - Domain of vpc endpoint.
    """
    name: pulumi.Output[str]
    """
    Name of container registry repository.
    """
    namespace: pulumi.Output[str]
    """
    Name of container registry namespace where repository is located.
    """
    repo_type: pulumi.Output[str]
    """
    `PUBLIC` or `PRIVATE`, repo's visibility.
    """
    summary: pulumi.Output[str]
    """
    The repository general information. It can contain 1 to 80 characters.
    """
    def __init__(__self__, resource_name, opts=None, detail=None, name=None, namespace=None, repo_type=None, summary=None, __props__=None, __name__=None, __opts__=None):
        """
        This resource will help you to manager Container Registry repositories.

        > **NOTE:** Available in v1.35.0+.

        > **NOTE:** You need to set your registry password in Container Registry console before use this resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cr_repo.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] detail: The repository specific information. MarkDown format is supported, and the length limit is 2000.
        :param pulumi.Input[str] name: Name of container registry repository.
        :param pulumi.Input[str] namespace: Name of container registry namespace where repository is located.
        :param pulumi.Input[str] repo_type: `PUBLIC` or `PRIVATE`, repo's visibility.
        :param pulumi.Input[str] summary: The repository general information. It can contain 1 to 80 characters.
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

            __props__['detail'] = detail
            __props__['name'] = name
            if namespace is None:
                raise TypeError("Missing required property 'namespace'")
            __props__['namespace'] = namespace
            if repo_type is None:
                raise TypeError("Missing required property 'repo_type'")
            __props__['repo_type'] = repo_type
            if summary is None:
                raise TypeError("Missing required property 'summary'")
            __props__['summary'] = summary
            __props__['domain_list'] = None
        super(Repo, __self__).__init__(
            'alicloud:cr/repo:Repo',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, detail=None, domain_list=None, name=None, namespace=None, repo_type=None, summary=None):
        """
        Get an existing Repo resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] detail: The repository specific information. MarkDown format is supported, and the length limit is 2000.
        :param pulumi.Input[dict] domain_list: The repository domain list.
        :param pulumi.Input[str] name: Name of container registry repository.
        :param pulumi.Input[str] namespace: Name of container registry namespace where repository is located.
        :param pulumi.Input[str] repo_type: `PUBLIC` or `PRIVATE`, repo's visibility.
        :param pulumi.Input[str] summary: The repository general information. It can contain 1 to 80 characters.

        The **domain_list** object supports the following:

          * `internal` (`pulumi.Input[str]`) - Domain of internal endpoint, only in some regions.
          * `public` (`pulumi.Input[str]`) - Domain of public endpoint.
          * `vpc` (`pulumi.Input[str]`) - Domain of vpc endpoint.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["detail"] = detail
        __props__["domain_list"] = domain_list
        __props__["name"] = name
        __props__["namespace"] = namespace
        __props__["repo_type"] = repo_type
        __props__["summary"] = summary
        return Repo(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

