// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Api gateway api can be imported using the id.Format to `<API Group Id>:<API Id>` e.g.
//
// ```sh
//  $ pulumi import alicloud:apigateway/api:Api example "ab2351f2ce904edaa8d92a0510832b91:e4f728fca5a94148b023b99a3e5d0b62"
// ```
type Api struct {
	pulumi.CustomResourceState

	// The ID of the api of api gateway.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// The authorization Type including APP and ANONYMOUS. Defaults to null.
	AuthType pulumi.StringOutput `pulumi:"authType"`
	// constant_parameters defines the constant parameters of the api.
	ConstantParameters ApiConstantParameterArrayOutput `pulumi:"constantParameters"`
	// The description of Constant parameter.
	Description pulumi.StringOutput `pulumi:"description"`
	// fc_service_config defines the config when serviceType selected 'FunctionCompute'.
	FcServiceConfig ApiFcServiceConfigPtrOutput `pulumi:"fcServiceConfig"`
	// The api gateway that the api belongs to. Defaults to null.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// http_service_config defines the config when serviceType selected 'HTTP'.
	HttpServiceConfig ApiHttpServiceConfigPtrOutput `pulumi:"httpServiceConfig"`
	// http_vpc_service_config defines the config when serviceType selected 'HTTP-VPC'.
	HttpVpcServiceConfig ApiHttpVpcServiceConfigPtrOutput `pulumi:"httpVpcServiceConfig"`
	// http_service_config defines the config when serviceType selected 'MOCK'.
	MockServiceConfig ApiMockServiceConfigPtrOutput `pulumi:"mockServiceConfig"`
	// System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
	Name pulumi.StringOutput `pulumi:"name"`
	// Request_config defines how users can send requests to your API.
	RequestConfig ApiRequestConfigOutput `pulumi:"requestConfig"`
	// request_parameters defines the request parameters of the api.
	RequestParameters ApiRequestParameterArrayOutput `pulumi:"requestParameters"`
	// The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
	ServiceType pulumi.StringOutput `pulumi:"serviceType"`
	// Stages that the api need to be deployed. Valid value: RELEASE | PRE | TEST.
	StageNames pulumi.StringArrayOutput `pulumi:"stageNames"`
	// system_parameters defines the system parameters of the api.
	SystemParameters ApiSystemParameterArrayOutput `pulumi:"systemParameters"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthType == nil {
		return nil, errors.New("invalid value for required argument 'AuthType'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.RequestConfig == nil {
		return nil, errors.New("invalid value for required argument 'RequestConfig'")
	}
	if args.ServiceType == nil {
		return nil, errors.New("invalid value for required argument 'ServiceType'")
	}
	var resource Api
	err := ctx.RegisterResource("alicloud:apigateway/api:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApi gets an existing Api resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("alicloud:apigateway/api:Api", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Api resources.
type apiState struct {
	// The ID of the api of api gateway.
	ApiId *string `pulumi:"apiId"`
	// The authorization Type including APP and ANONYMOUS. Defaults to null.
	AuthType *string `pulumi:"authType"`
	// constant_parameters defines the constant parameters of the api.
	ConstantParameters []ApiConstantParameter `pulumi:"constantParameters"`
	// The description of Constant parameter.
	Description *string `pulumi:"description"`
	// fc_service_config defines the config when serviceType selected 'FunctionCompute'.
	FcServiceConfig *ApiFcServiceConfig `pulumi:"fcServiceConfig"`
	// The api gateway that the api belongs to. Defaults to null.
	GroupId *string `pulumi:"groupId"`
	// http_service_config defines the config when serviceType selected 'HTTP'.
	HttpServiceConfig *ApiHttpServiceConfig `pulumi:"httpServiceConfig"`
	// http_vpc_service_config defines the config when serviceType selected 'HTTP-VPC'.
	HttpVpcServiceConfig *ApiHttpVpcServiceConfig `pulumi:"httpVpcServiceConfig"`
	// http_service_config defines the config when serviceType selected 'MOCK'.
	MockServiceConfig *ApiMockServiceConfig `pulumi:"mockServiceConfig"`
	// System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
	Name *string `pulumi:"name"`
	// Request_config defines how users can send requests to your API.
	RequestConfig *ApiRequestConfig `pulumi:"requestConfig"`
	// request_parameters defines the request parameters of the api.
	RequestParameters []ApiRequestParameter `pulumi:"requestParameters"`
	// The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
	ServiceType *string `pulumi:"serviceType"`
	// Stages that the api need to be deployed. Valid value: RELEASE | PRE | TEST.
	StageNames []string `pulumi:"stageNames"`
	// system_parameters defines the system parameters of the api.
	SystemParameters []ApiSystemParameter `pulumi:"systemParameters"`
}

type ApiState struct {
	// The ID of the api of api gateway.
	ApiId pulumi.StringPtrInput
	// The authorization Type including APP and ANONYMOUS. Defaults to null.
	AuthType pulumi.StringPtrInput
	// constant_parameters defines the constant parameters of the api.
	ConstantParameters ApiConstantParameterArrayInput
	// The description of Constant parameter.
	Description pulumi.StringPtrInput
	// fc_service_config defines the config when serviceType selected 'FunctionCompute'.
	FcServiceConfig ApiFcServiceConfigPtrInput
	// The api gateway that the api belongs to. Defaults to null.
	GroupId pulumi.StringPtrInput
	// http_service_config defines the config when serviceType selected 'HTTP'.
	HttpServiceConfig ApiHttpServiceConfigPtrInput
	// http_vpc_service_config defines the config when serviceType selected 'HTTP-VPC'.
	HttpVpcServiceConfig ApiHttpVpcServiceConfigPtrInput
	// http_service_config defines the config when serviceType selected 'MOCK'.
	MockServiceConfig ApiMockServiceConfigPtrInput
	// System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
	Name pulumi.StringPtrInput
	// Request_config defines how users can send requests to your API.
	RequestConfig ApiRequestConfigPtrInput
	// request_parameters defines the request parameters of the api.
	RequestParameters ApiRequestParameterArrayInput
	// The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
	ServiceType pulumi.StringPtrInput
	// Stages that the api need to be deployed. Valid value: RELEASE | PRE | TEST.
	StageNames pulumi.StringArrayInput
	// system_parameters defines the system parameters of the api.
	SystemParameters ApiSystemParameterArrayInput
}

func (ApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiState)(nil)).Elem()
}

type apiArgs struct {
	// The authorization Type including APP and ANONYMOUS. Defaults to null.
	AuthType string `pulumi:"authType"`
	// constant_parameters defines the constant parameters of the api.
	ConstantParameters []ApiConstantParameter `pulumi:"constantParameters"`
	// The description of Constant parameter.
	Description string `pulumi:"description"`
	// fc_service_config defines the config when serviceType selected 'FunctionCompute'.
	FcServiceConfig *ApiFcServiceConfig `pulumi:"fcServiceConfig"`
	// The api gateway that the api belongs to. Defaults to null.
	GroupId string `pulumi:"groupId"`
	// http_service_config defines the config when serviceType selected 'HTTP'.
	HttpServiceConfig *ApiHttpServiceConfig `pulumi:"httpServiceConfig"`
	// http_vpc_service_config defines the config when serviceType selected 'HTTP-VPC'.
	HttpVpcServiceConfig *ApiHttpVpcServiceConfig `pulumi:"httpVpcServiceConfig"`
	// http_service_config defines the config when serviceType selected 'MOCK'.
	MockServiceConfig *ApiMockServiceConfig `pulumi:"mockServiceConfig"`
	// System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
	Name *string `pulumi:"name"`
	// Request_config defines how users can send requests to your API.
	RequestConfig ApiRequestConfig `pulumi:"requestConfig"`
	// request_parameters defines the request parameters of the api.
	RequestParameters []ApiRequestParameter `pulumi:"requestParameters"`
	// The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
	ServiceType string `pulumi:"serviceType"`
	// Stages that the api need to be deployed. Valid value: RELEASE | PRE | TEST.
	StageNames []string `pulumi:"stageNames"`
	// system_parameters defines the system parameters of the api.
	SystemParameters []ApiSystemParameter `pulumi:"systemParameters"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	// The authorization Type including APP and ANONYMOUS. Defaults to null.
	AuthType pulumi.StringInput
	// constant_parameters defines the constant parameters of the api.
	ConstantParameters ApiConstantParameterArrayInput
	// The description of Constant parameter.
	Description pulumi.StringInput
	// fc_service_config defines the config when serviceType selected 'FunctionCompute'.
	FcServiceConfig ApiFcServiceConfigPtrInput
	// The api gateway that the api belongs to. Defaults to null.
	GroupId pulumi.StringInput
	// http_service_config defines the config when serviceType selected 'HTTP'.
	HttpServiceConfig ApiHttpServiceConfigPtrInput
	// http_vpc_service_config defines the config when serviceType selected 'HTTP-VPC'.
	HttpVpcServiceConfig ApiHttpVpcServiceConfigPtrInput
	// http_service_config defines the config when serviceType selected 'MOCK'.
	MockServiceConfig ApiMockServiceConfigPtrInput
	// System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)
	Name pulumi.StringPtrInput
	// Request_config defines how users can send requests to your API.
	RequestConfig ApiRequestConfigInput
	// request_parameters defines the request parameters of the api.
	RequestParameters ApiRequestParameterArrayInput
	// The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
	ServiceType pulumi.StringInput
	// Stages that the api need to be deployed. Valid value: RELEASE | PRE | TEST.
	StageNames pulumi.StringArrayInput
	// system_parameters defines the system parameters of the api.
	SystemParameters ApiSystemParameterArrayInput
}

func (ApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiArgs)(nil)).Elem()
}

type ApiInput interface {
	pulumi.Input

	ToApiOutput() ApiOutput
	ToApiOutputWithContext(ctx context.Context) ApiOutput
}

func (*Api) ElementType() reflect.Type {
	return reflect.TypeOf((*Api)(nil))
}

func (i *Api) ToApiOutput() ApiOutput {
	return i.ToApiOutputWithContext(context.Background())
}

func (i *Api) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOutput)
}

func (i *Api) ToApiPtrOutput() ApiPtrOutput {
	return i.ToApiPtrOutputWithContext(context.Background())
}

func (i *Api) ToApiPtrOutputWithContext(ctx context.Context) ApiPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPtrOutput)
}

type ApiPtrInput interface {
	pulumi.Input

	ToApiPtrOutput() ApiPtrOutput
	ToApiPtrOutputWithContext(ctx context.Context) ApiPtrOutput
}

type apiPtrType ApiArgs

func (*apiPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil))
}

func (i *apiPtrType) ToApiPtrOutput() ApiPtrOutput {
	return i.ToApiPtrOutputWithContext(context.Background())
}

func (i *apiPtrType) ToApiPtrOutputWithContext(ctx context.Context) ApiPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPtrOutput)
}

// ApiArrayInput is an input type that accepts ApiArray and ApiArrayOutput values.
// You can construct a concrete instance of `ApiArrayInput` via:
//
//          ApiArray{ ApiArgs{...} }
type ApiArrayInput interface {
	pulumi.Input

	ToApiArrayOutput() ApiArrayOutput
	ToApiArrayOutputWithContext(context.Context) ApiArrayOutput
}

type ApiArray []ApiInput

func (ApiArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Api)(nil))
}

func (i ApiArray) ToApiArrayOutput() ApiArrayOutput {
	return i.ToApiArrayOutputWithContext(context.Background())
}

func (i ApiArray) ToApiArrayOutputWithContext(ctx context.Context) ApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiArrayOutput)
}

// ApiMapInput is an input type that accepts ApiMap and ApiMapOutput values.
// You can construct a concrete instance of `ApiMapInput` via:
//
//          ApiMap{ "key": ApiArgs{...} }
type ApiMapInput interface {
	pulumi.Input

	ToApiMapOutput() ApiMapOutput
	ToApiMapOutputWithContext(context.Context) ApiMapOutput
}

type ApiMap map[string]ApiInput

func (ApiMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Api)(nil))
}

func (i ApiMap) ToApiMapOutput() ApiMapOutput {
	return i.ToApiMapOutputWithContext(context.Background())
}

func (i ApiMap) ToApiMapOutputWithContext(ctx context.Context) ApiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiMapOutput)
}

type ApiOutput struct {
	*pulumi.OutputState
}

func (ApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Api)(nil))
}

func (o ApiOutput) ToApiOutput() ApiOutput {
	return o
}

func (o ApiOutput) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return o
}

func (o ApiOutput) ToApiPtrOutput() ApiPtrOutput {
	return o.ToApiPtrOutputWithContext(context.Background())
}

func (o ApiOutput) ToApiPtrOutputWithContext(ctx context.Context) ApiPtrOutput {
	return o.ApplyT(func(v Api) *Api {
		return &v
	}).(ApiPtrOutput)
}

type ApiPtrOutput struct {
	*pulumi.OutputState
}

func (ApiPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil))
}

func (o ApiPtrOutput) ToApiPtrOutput() ApiPtrOutput {
	return o
}

func (o ApiPtrOutput) ToApiPtrOutputWithContext(ctx context.Context) ApiPtrOutput {
	return o
}

type ApiArrayOutput struct{ *pulumi.OutputState }

func (ApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Api)(nil))
}

func (o ApiArrayOutput) ToApiArrayOutput() ApiArrayOutput {
	return o
}

func (o ApiArrayOutput) ToApiArrayOutputWithContext(ctx context.Context) ApiArrayOutput {
	return o
}

func (o ApiArrayOutput) Index(i pulumi.IntInput) ApiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Api {
		return vs[0].([]Api)[vs[1].(int)]
	}).(ApiOutput)
}

type ApiMapOutput struct{ *pulumi.OutputState }

func (ApiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Api)(nil))
}

func (o ApiMapOutput) ToApiMapOutput() ApiMapOutput {
	return o
}

func (o ApiMapOutput) ToApiMapOutputWithContext(ctx context.Context) ApiMapOutput {
	return o
}

func (o ApiMapOutput) MapIndex(k pulumi.StringInput) ApiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Api {
		return vs[0].(map[string]Api)[vs[1].(string)]
	}).(ApiOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiOutput{})
	pulumi.RegisterOutputType(ApiPtrOutput{})
	pulumi.RegisterOutputType(ApiArrayOutput{})
	pulumi.RegisterOutputType(ApiMapOutput{})
}
