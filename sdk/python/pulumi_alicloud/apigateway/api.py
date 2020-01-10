# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Api(pulumi.CustomResource):
    api_id: pulumi.Output[str]
    """
    The ID of the api of api gateway.
    """
    auth_type: pulumi.Output[str]
    """
    The authorization Type including APP and ANONYMOUS. Defaults to null.
    """
    constant_parameters: pulumi.Output[list]
    """
    constant_parameters defines the constant parameters of the api.
    
      * `description` (`str`) - The description of Constant parameter.
      * `in` (`str`) - System parameter location; values: 'HEAD' and 'QUERY'.
      * `name` (`str`) - System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
      * `value` (`str`) - Constant parameter value.
    """
    description: pulumi.Output[str]
    """
    The description of Constant parameter.
    """
    fc_service_config: pulumi.Output[dict]
    """
    fc_service_config defines the config when service_type selected 'FunctionCompute'.
    
      * `arnRole` (`str`) - RAM role arn attached to the Function Compute service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
      * `functionName` (`str`) - The function name of function compute service.
      * `region` (`str`) - The region that the function compute service belongs to.
      * `serviceName` (`str`) - The service name of function compute service.
      * `timeout` (`float`) - Backend service time-out time; unit: millisecond.
    """
    group_id: pulumi.Output[str]
    """
    The api gateway that the api belongs to. Defaults to null.
    """
    http_service_config: pulumi.Output[dict]
    """
    http_service_config defines the config when service_type selected 'HTTP'.
    
      * `address` (`str`) - The address of backend service.
      * `aoneName` (`str`)
      * `method` (`str`) - The http method of backend service.
      * `path` (`str`) - The path of backend service.
      * `timeout` (`float`) - Backend service time-out time; unit: millisecond.
    """
    http_vpc_service_config: pulumi.Output[dict]
    """
    http_vpc_service_config defines the config when service_type selected 'HTTP-VPC'.
    
      * `aoneName` (`str`)
      * `method` (`str`) - The http method of backend service.
      * `name` (`str`) - System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
      * `path` (`str`) - The path of backend service.
      * `timeout` (`float`) - Backend service time-out time; unit: millisecond.
    """
    mock_service_config: pulumi.Output[dict]
    """
    http_service_config defines the config when service_type selected 'MOCK'.
    
      * `aoneName` (`str`)
      * `result` (`str`) - The result of the mock service.
    """
    name: pulumi.Output[str]
    """
    System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
    """
    request_config: pulumi.Output[dict]
    """
    Request_config defines how users can send requests to your API.
    
      * `bodyFormat` (`str`) - The body format of the api, which support the values of 'STREAM' and 'FORM'
      * `method` (`str`) - The http method of backend service.
      * `mode` (`str`) - The mode of the parameters between request parameters and service parameters, which support the values of 'MAPPING' and 'PASSTHROUGH'
      * `path` (`str`) - The path of backend service.
      * `protocol` (`str`) - The protocol of api which supports values of 'HTTP','HTTPS' or 'HTTP,HTTPS'
    """
    request_parameters: pulumi.Output[list]
    """
    request_parameters defines the request parameters of the api.
    
      * `defaultValue` (`str`) - The default value of the parameter.
      * `description` (`str`) - The description of Constant parameter.
      * `in` (`str`) - System parameter location; values: 'HEAD' and 'QUERY'.
      * `inService` (`str`) - Backend service's parameter location; values: BODY, HEAD, QUERY, and PATH.
      * `name` (`str`) - System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
      * `nameService` (`str`) - Backend service's parameter name.
      * `required` (`str`) - Parameter required or not; values: REQUIRED and OPTIONAL.
      * `type` (`str`) - Parameter type which supports values of 'STRING','INT','BOOLEAN','LONG',"FLOAT" and "DOUBLE"
    """
    service_type: pulumi.Output[str]
    """
    The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
    """
    stage_names: pulumi.Output[list]
    """
    Stages that the api need to be deployed. Valid value: RELEASE | PRE | TEST.
    """
    system_parameters: pulumi.Output[list]
    """
    system_parameters defines the system parameters of the api.
    
      * `in` (`str`) - System parameter location; values: 'HEAD' and 'QUERY'.
      * `name` (`str`) - System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
      * `nameService` (`str`) - Backend service's parameter name.
    """
    def __init__(__self__, resource_name, opts=None, auth_type=None, constant_parameters=None, description=None, fc_service_config=None, group_id=None, http_service_config=None, http_vpc_service_config=None, mock_service_config=None, name=None, request_config=None, request_parameters=None, service_type=None, stage_names=None, system_parameters=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Api resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_type: The authorization Type including APP and ANONYMOUS. Defaults to null.
        :param pulumi.Input[list] constant_parameters: constant_parameters defines the constant parameters of the api.
        :param pulumi.Input[str] description: The description of Constant parameter.
        :param pulumi.Input[dict] fc_service_config: fc_service_config defines the config when service_type selected 'FunctionCompute'.
        :param pulumi.Input[str] group_id: The api gateway that the api belongs to. Defaults to null.
        :param pulumi.Input[dict] http_service_config: http_service_config defines the config when service_type selected 'HTTP'.
        :param pulumi.Input[dict] http_vpc_service_config: http_vpc_service_config defines the config when service_type selected 'HTTP-VPC'.
        :param pulumi.Input[dict] mock_service_config: http_service_config defines the config when service_type selected 'MOCK'.
        :param pulumi.Input[str] name: System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
        :param pulumi.Input[dict] request_config: Request_config defines how users can send requests to your API.
        :param pulumi.Input[list] request_parameters: request_parameters defines the request parameters of the api.
        :param pulumi.Input[str] service_type: The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
        :param pulumi.Input[list] stage_names: Stages that the api need to be deployed. Valid value: RELEASE | PRE | TEST.
        :param pulumi.Input[list] system_parameters: system_parameters defines the system parameters of the api.
        
        The **constant_parameters** object supports the following:
        
          * `description` (`pulumi.Input[str]`) - The description of Constant parameter.
          * `in` (`pulumi.Input[str]`) - System parameter location; values: 'HEAD' and 'QUERY'.
          * `name` (`pulumi.Input[str]`) - System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
          * `value` (`pulumi.Input[str]`) - Constant parameter value.
        
        The **fc_service_config** object supports the following:
        
          * `arnRole` (`pulumi.Input[str]`) - RAM role arn attached to the Function Compute service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
          * `functionName` (`pulumi.Input[str]`) - The function name of function compute service.
          * `region` (`pulumi.Input[str]`) - The region that the function compute service belongs to.
          * `serviceName` (`pulumi.Input[str]`) - The service name of function compute service.
          * `timeout` (`pulumi.Input[float]`) - Backend service time-out time; unit: millisecond.
        
        The **http_service_config** object supports the following:
        
          * `address` (`pulumi.Input[str]`) - The address of backend service.
          * `aoneName` (`pulumi.Input[str]`)
          * `method` (`pulumi.Input[str]`) - The http method of backend service.
          * `path` (`pulumi.Input[str]`) - The path of backend service.
          * `timeout` (`pulumi.Input[float]`) - Backend service time-out time; unit: millisecond.
        
        The **http_vpc_service_config** object supports the following:
        
          * `aoneName` (`pulumi.Input[str]`)
          * `method` (`pulumi.Input[str]`) - The http method of backend service.
          * `name` (`pulumi.Input[str]`) - System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
          * `path` (`pulumi.Input[str]`) - The path of backend service.
          * `timeout` (`pulumi.Input[float]`) - Backend service time-out time; unit: millisecond.
        
        The **mock_service_config** object supports the following:
        
          * `aoneName` (`pulumi.Input[str]`)
          * `result` (`pulumi.Input[str]`) - The result of the mock service.
        
        The **request_config** object supports the following:
        
          * `bodyFormat` (`pulumi.Input[str]`) - The body format of the api, which support the values of 'STREAM' and 'FORM'
          * `method` (`pulumi.Input[str]`) - The http method of backend service.
          * `mode` (`pulumi.Input[str]`) - The mode of the parameters between request parameters and service parameters, which support the values of 'MAPPING' and 'PASSTHROUGH'
          * `path` (`pulumi.Input[str]`) - The path of backend service.
          * `protocol` (`pulumi.Input[str]`) - The protocol of api which supports values of 'HTTP','HTTPS' or 'HTTP,HTTPS'
        
        The **request_parameters** object supports the following:
        
          * `defaultValue` (`pulumi.Input[str]`) - The default value of the parameter.
          * `description` (`pulumi.Input[str]`) - The description of Constant parameter.
          * `in` (`pulumi.Input[str]`) - System parameter location; values: 'HEAD' and 'QUERY'.
          * `inService` (`pulumi.Input[str]`) - Backend service's parameter location; values: BODY, HEAD, QUERY, and PATH.
          * `name` (`pulumi.Input[str]`) - System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
          * `nameService` (`pulumi.Input[str]`) - Backend service's parameter name.
          * `required` (`pulumi.Input[str]`) - Parameter required or not; values: REQUIRED and OPTIONAL.
          * `type` (`pulumi.Input[str]`) - Parameter type which supports values of 'STRING','INT','BOOLEAN','LONG',"FLOAT" and "DOUBLE"
        
        The **system_parameters** object supports the following:
        
          * `in` (`pulumi.Input[str]`) - System parameter location; values: 'HEAD' and 'QUERY'.
          * `name` (`pulumi.Input[str]`) - System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
          * `nameService` (`pulumi.Input[str]`) - Backend service's parameter name.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/api_gateway_api.html.markdown.
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

            if auth_type is None:
                raise TypeError("Missing required property 'auth_type'")
            __props__['auth_type'] = auth_type
            __props__['constant_parameters'] = constant_parameters
            if description is None:
                raise TypeError("Missing required property 'description'")
            __props__['description'] = description
            __props__['fc_service_config'] = fc_service_config
            if group_id is None:
                raise TypeError("Missing required property 'group_id'")
            __props__['group_id'] = group_id
            __props__['http_service_config'] = http_service_config
            __props__['http_vpc_service_config'] = http_vpc_service_config
            __props__['mock_service_config'] = mock_service_config
            __props__['name'] = name
            if request_config is None:
                raise TypeError("Missing required property 'request_config'")
            __props__['request_config'] = request_config
            __props__['request_parameters'] = request_parameters
            if service_type is None:
                raise TypeError("Missing required property 'service_type'")
            __props__['service_type'] = service_type
            __props__['stage_names'] = stage_names
            __props__['system_parameters'] = system_parameters
            __props__['api_id'] = None
        super(Api, __self__).__init__(
            'alicloud:apigateway/api:Api',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, api_id=None, auth_type=None, constant_parameters=None, description=None, fc_service_config=None, group_id=None, http_service_config=None, http_vpc_service_config=None, mock_service_config=None, name=None, request_config=None, request_parameters=None, service_type=None, stage_names=None, system_parameters=None):
        """
        Get an existing Api resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: The ID of the api of api gateway.
        :param pulumi.Input[str] auth_type: The authorization Type including APP and ANONYMOUS. Defaults to null.
        :param pulumi.Input[list] constant_parameters: constant_parameters defines the constant parameters of the api.
        :param pulumi.Input[str] description: The description of Constant parameter.
        :param pulumi.Input[dict] fc_service_config: fc_service_config defines the config when service_type selected 'FunctionCompute'.
        :param pulumi.Input[str] group_id: The api gateway that the api belongs to. Defaults to null.
        :param pulumi.Input[dict] http_service_config: http_service_config defines the config when service_type selected 'HTTP'.
        :param pulumi.Input[dict] http_vpc_service_config: http_vpc_service_config defines the config when service_type selected 'HTTP-VPC'.
        :param pulumi.Input[dict] mock_service_config: http_service_config defines the config when service_type selected 'MOCK'.
        :param pulumi.Input[str] name: System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
        :param pulumi.Input[dict] request_config: Request_config defines how users can send requests to your API.
        :param pulumi.Input[list] request_parameters: request_parameters defines the request parameters of the api.
        :param pulumi.Input[str] service_type: The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
        :param pulumi.Input[list] stage_names: Stages that the api need to be deployed. Valid value: RELEASE | PRE | TEST.
        :param pulumi.Input[list] system_parameters: system_parameters defines the system parameters of the api.
        
        The **constant_parameters** object supports the following:
        
          * `description` (`pulumi.Input[str]`) - The description of Constant parameter.
          * `in` (`pulumi.Input[str]`) - System parameter location; values: 'HEAD' and 'QUERY'.
          * `name` (`pulumi.Input[str]`) - System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
          * `value` (`pulumi.Input[str]`) - Constant parameter value.
        
        The **fc_service_config** object supports the following:
        
          * `arnRole` (`pulumi.Input[str]`) - RAM role arn attached to the Function Compute service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
          * `functionName` (`pulumi.Input[str]`) - The function name of function compute service.
          * `region` (`pulumi.Input[str]`) - The region that the function compute service belongs to.
          * `serviceName` (`pulumi.Input[str]`) - The service name of function compute service.
          * `timeout` (`pulumi.Input[float]`) - Backend service time-out time; unit: millisecond.
        
        The **http_service_config** object supports the following:
        
          * `address` (`pulumi.Input[str]`) - The address of backend service.
          * `aoneName` (`pulumi.Input[str]`)
          * `method` (`pulumi.Input[str]`) - The http method of backend service.
          * `path` (`pulumi.Input[str]`) - The path of backend service.
          * `timeout` (`pulumi.Input[float]`) - Backend service time-out time; unit: millisecond.
        
        The **http_vpc_service_config** object supports the following:
        
          * `aoneName` (`pulumi.Input[str]`)
          * `method` (`pulumi.Input[str]`) - The http method of backend service.
          * `name` (`pulumi.Input[str]`) - System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
          * `path` (`pulumi.Input[str]`) - The path of backend service.
          * `timeout` (`pulumi.Input[float]`) - Backend service time-out time; unit: millisecond.
        
        The **mock_service_config** object supports the following:
        
          * `aoneName` (`pulumi.Input[str]`)
          * `result` (`pulumi.Input[str]`) - The result of the mock service.
        
        The **request_config** object supports the following:
        
          * `bodyFormat` (`pulumi.Input[str]`) - The body format of the api, which support the values of 'STREAM' and 'FORM'
          * `method` (`pulumi.Input[str]`) - The http method of backend service.
          * `mode` (`pulumi.Input[str]`) - The mode of the parameters between request parameters and service parameters, which support the values of 'MAPPING' and 'PASSTHROUGH'
          * `path` (`pulumi.Input[str]`) - The path of backend service.
          * `protocol` (`pulumi.Input[str]`) - The protocol of api which supports values of 'HTTP','HTTPS' or 'HTTP,HTTPS'
        
        The **request_parameters** object supports the following:
        
          * `defaultValue` (`pulumi.Input[str]`) - The default value of the parameter.
          * `description` (`pulumi.Input[str]`) - The description of Constant parameter.
          * `in` (`pulumi.Input[str]`) - System parameter location; values: 'HEAD' and 'QUERY'.
          * `inService` (`pulumi.Input[str]`) - Backend service's parameter location; values: BODY, HEAD, QUERY, and PATH.
          * `name` (`pulumi.Input[str]`) - System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
          * `nameService` (`pulumi.Input[str]`) - Backend service's parameter name.
          * `required` (`pulumi.Input[str]`) - Parameter required or not; values: REQUIRED and OPTIONAL.
          * `type` (`pulumi.Input[str]`) - Parameter type which supports values of 'STRING','INT','BOOLEAN','LONG',"FLOAT" and "DOUBLE"
        
        The **system_parameters** object supports the following:
        
          * `in` (`pulumi.Input[str]`) - System parameter location; values: 'HEAD' and 'QUERY'.
          * `name` (`pulumi.Input[str]`) - System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
          * `nameService` (`pulumi.Input[str]`) - Backend service's parameter name.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/api_gateway_api.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["api_id"] = api_id
        __props__["auth_type"] = auth_type
        __props__["constant_parameters"] = constant_parameters
        __props__["description"] = description
        __props__["fc_service_config"] = fc_service_config
        __props__["group_id"] = group_id
        __props__["http_service_config"] = http_service_config
        __props__["http_vpc_service_config"] = http_vpc_service_config
        __props__["mock_service_config"] = mock_service_config
        __props__["name"] = name
        __props__["request_config"] = request_config
        __props__["request_parameters"] = request_parameters
        __props__["service_type"] = service_type
        __props__["stage_names"] = stage_names
        __props__["system_parameters"] = system_parameters
        return Api(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

