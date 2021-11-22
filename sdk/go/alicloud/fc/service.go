// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Function Compute Service can be imported using the id or name, e.g.
//
// ```sh
//  $ pulumi import alicloud:fc/service:Service foo my-fc-service
// ```
type Service struct {
	pulumi.CustomResourceState

	// The Function Compute Service description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether to allow the Service to access Internet. Default to "true".
	InternetAccess pulumi.BoolPtrOutput `pulumi:"internetAccess"`
	// The date this resource was last modified.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// Provide this to store your Function Compute Service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm).
	LogConfig ServiceLogConfigPtrOutput `pulumi:"logConfig"`
	// The Function Compute Service name. It is the only in one Alicloud account and is conflict with `namePrefix`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Setting a prefix to get a only name. It is conflict with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// Provide [NAS configuration](https://www.alibabacloud.com/help/doc-detail/87401.htm) to allow Function Compute Service to access your NAS resources.
	NasConfig ServiceNasConfigPtrOutput `pulumi:"nasConfig"`
	// Whether to publish creation/change as new Function Compute Service Version. Defaults to `false`.
	Publish pulumi.BoolPtrOutput `pulumi:"publish"`
	// RAM role arn attached to the Function Compute Service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// The Function Compute Service ID.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// The latest published version of your Function Compute Service.
	Version pulumi.StringOutput `pulumi:"version"`
	// Provide this to allow your Function Compute Service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm).
	VpcConfig ServiceVpcConfigPtrOutput `pulumi:"vpcConfig"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		args = &ServiceArgs{}
	}

	var resource Service
	err := ctx.RegisterResource("alicloud:fc/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("alicloud:fc/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// The Function Compute Service description.
	Description *string `pulumi:"description"`
	// Whether to allow the Service to access Internet. Default to "true".
	InternetAccess *bool `pulumi:"internetAccess"`
	// The date this resource was last modified.
	LastModified *string `pulumi:"lastModified"`
	// Provide this to store your Function Compute Service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm).
	LogConfig *ServiceLogConfig `pulumi:"logConfig"`
	// The Function Compute Service name. It is the only in one Alicloud account and is conflict with `namePrefix`.
	Name *string `pulumi:"name"`
	// Setting a prefix to get a only name. It is conflict with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Provide [NAS configuration](https://www.alibabacloud.com/help/doc-detail/87401.htm) to allow Function Compute Service to access your NAS resources.
	NasConfig *ServiceNasConfig `pulumi:"nasConfig"`
	// Whether to publish creation/change as new Function Compute Service Version. Defaults to `false`.
	Publish *bool `pulumi:"publish"`
	// RAM role arn attached to the Function Compute Service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
	Role *string `pulumi:"role"`
	// The Function Compute Service ID.
	ServiceId *string `pulumi:"serviceId"`
	// The latest published version of your Function Compute Service.
	Version *string `pulumi:"version"`
	// Provide this to allow your Function Compute Service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm).
	VpcConfig *ServiceVpcConfig `pulumi:"vpcConfig"`
}

type ServiceState struct {
	// The Function Compute Service description.
	Description pulumi.StringPtrInput
	// Whether to allow the Service to access Internet. Default to "true".
	InternetAccess pulumi.BoolPtrInput
	// The date this resource was last modified.
	LastModified pulumi.StringPtrInput
	// Provide this to store your Function Compute Service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm).
	LogConfig ServiceLogConfigPtrInput
	// The Function Compute Service name. It is the only in one Alicloud account and is conflict with `namePrefix`.
	Name pulumi.StringPtrInput
	// Setting a prefix to get a only name. It is conflict with `name`.
	NamePrefix pulumi.StringPtrInput
	// Provide [NAS configuration](https://www.alibabacloud.com/help/doc-detail/87401.htm) to allow Function Compute Service to access your NAS resources.
	NasConfig ServiceNasConfigPtrInput
	// Whether to publish creation/change as new Function Compute Service Version. Defaults to `false`.
	Publish pulumi.BoolPtrInput
	// RAM role arn attached to the Function Compute Service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
	Role pulumi.StringPtrInput
	// The Function Compute Service ID.
	ServiceId pulumi.StringPtrInput
	// The latest published version of your Function Compute Service.
	Version pulumi.StringPtrInput
	// Provide this to allow your Function Compute Service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm).
	VpcConfig ServiceVpcConfigPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// The Function Compute Service description.
	Description *string `pulumi:"description"`
	// Whether to allow the Service to access Internet. Default to "true".
	InternetAccess *bool `pulumi:"internetAccess"`
	// Provide this to store your Function Compute Service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm).
	LogConfig *ServiceLogConfig `pulumi:"logConfig"`
	// The Function Compute Service name. It is the only in one Alicloud account and is conflict with `namePrefix`.
	Name *string `pulumi:"name"`
	// Setting a prefix to get a only name. It is conflict with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Provide [NAS configuration](https://www.alibabacloud.com/help/doc-detail/87401.htm) to allow Function Compute Service to access your NAS resources.
	NasConfig *ServiceNasConfig `pulumi:"nasConfig"`
	// Whether to publish creation/change as new Function Compute Service Version. Defaults to `false`.
	Publish *bool `pulumi:"publish"`
	// RAM role arn attached to the Function Compute Service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
	Role *string `pulumi:"role"`
	// Provide this to allow your Function Compute Service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm).
	VpcConfig *ServiceVpcConfig `pulumi:"vpcConfig"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// The Function Compute Service description.
	Description pulumi.StringPtrInput
	// Whether to allow the Service to access Internet. Default to "true".
	InternetAccess pulumi.BoolPtrInput
	// Provide this to store your Function Compute Service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm).
	LogConfig ServiceLogConfigPtrInput
	// The Function Compute Service name. It is the only in one Alicloud account and is conflict with `namePrefix`.
	Name pulumi.StringPtrInput
	// Setting a prefix to get a only name. It is conflict with `name`.
	NamePrefix pulumi.StringPtrInput
	// Provide [NAS configuration](https://www.alibabacloud.com/help/doc-detail/87401.htm) to allow Function Compute Service to access your NAS resources.
	NasConfig ServiceNasConfigPtrInput
	// Whether to publish creation/change as new Function Compute Service Version. Defaults to `false`.
	Publish pulumi.BoolPtrInput
	// RAM role arn attached to the Function Compute Service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.
	Role pulumi.StringPtrInput
	// Provide this to allow your Function Compute Service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm).
	VpcConfig ServiceVpcConfigPtrInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil))
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

func (i *Service) ToServicePtrOutput() ServicePtrOutput {
	return i.ToServicePtrOutputWithContext(context.Background())
}

func (i *Service) ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePtrOutput)
}

type ServicePtrInput interface {
	pulumi.Input

	ToServicePtrOutput() ServicePtrOutput
	ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput
}

type servicePtrType ServiceArgs

func (*servicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil))
}

func (i *servicePtrType) ToServicePtrOutput() ServicePtrOutput {
	return i.ToServicePtrOutputWithContext(context.Background())
}

func (i *servicePtrType) ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePtrOutput)
}

// ServiceArrayInput is an input type that accepts ServiceArray and ServiceArrayOutput values.
// You can construct a concrete instance of `ServiceArrayInput` via:
//
//          ServiceArray{ ServiceArgs{...} }
type ServiceArrayInput interface {
	pulumi.Input

	ToServiceArrayOutput() ServiceArrayOutput
	ToServiceArrayOutputWithContext(context.Context) ServiceArrayOutput
}

type ServiceArray []ServiceInput

func (ServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (i ServiceArray) ToServiceArrayOutput() ServiceArrayOutput {
	return i.ToServiceArrayOutputWithContext(context.Background())
}

func (i ServiceArray) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceArrayOutput)
}

// ServiceMapInput is an input type that accepts ServiceMap and ServiceMapOutput values.
// You can construct a concrete instance of `ServiceMapInput` via:
//
//          ServiceMap{ "key": ServiceArgs{...} }
type ServiceMapInput interface {
	pulumi.Input

	ToServiceMapOutput() ServiceMapOutput
	ToServiceMapOutputWithContext(context.Context) ServiceMapOutput
}

type ServiceMap map[string]ServiceInput

func (ServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (i ServiceMap) ToServiceMapOutput() ServiceMapOutput {
	return i.ToServiceMapOutputWithContext(context.Background())
}

func (i ServiceMap) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMapOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil))
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

func (o ServiceOutput) ToServicePtrOutput() ServicePtrOutput {
	return o.ToServicePtrOutputWithContext(context.Background())
}

func (o ServiceOutput) ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Service) *Service {
		return &v
	}).(ServicePtrOutput)
}

type ServicePtrOutput struct{ *pulumi.OutputState }

func (ServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil))
}

func (o ServicePtrOutput) ToServicePtrOutput() ServicePtrOutput {
	return o
}

func (o ServicePtrOutput) ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput {
	return o
}

func (o ServicePtrOutput) Elem() ServiceOutput {
	return o.ApplyT(func(v *Service) Service {
		if v != nil {
			return *v
		}
		var ret Service
		return ret
	}).(ServiceOutput)
}

type ServiceArrayOutput struct{ *pulumi.OutputState }

func (ServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Service)(nil))
}

func (o ServiceArrayOutput) ToServiceArrayOutput() ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) Index(i pulumi.IntInput) ServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Service {
		return vs[0].([]Service)[vs[1].(int)]
	}).(ServiceOutput)
}

type ServiceMapOutput struct{ *pulumi.OutputState }

func (ServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Service)(nil))
}

func (o ServiceMapOutput) ToServiceMapOutput() ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) MapIndex(k pulumi.StringInput) ServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Service {
		return vs[0].(map[string]Service)[vs[1].(string)]
	}).(ServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePtrInput)(nil)).Elem(), &Service{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceArrayInput)(nil)).Elem(), ServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceMapInput)(nil)).Elem(), ServiceMap{})
	pulumi.RegisterOutputType(ServiceOutput{})
	pulumi.RegisterOutputType(ServicePtrOutput{})
	pulumi.RegisterOutputType(ServiceArrayOutput{})
	pulumi.RegisterOutputType(ServiceMapOutput{})
}
