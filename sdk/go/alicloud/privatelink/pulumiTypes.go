// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type VpcEndpointServiceResource struct {
	// The id of service resources added to the endpoint service.
	ResourceId *string `pulumi:"resourceId"`
	// The type of service resource added to the endpoint service.
	ResourceType *string `pulumi:"resourceType"`
}

// VpcEndpointServiceResourceInput is an input type that accepts VpcEndpointServiceResourceArgs and VpcEndpointServiceResourceOutput values.
// You can construct a concrete instance of `VpcEndpointServiceResourceInput` via:
//
//          VpcEndpointServiceResourceArgs{...}
type VpcEndpointServiceResourceInput interface {
	pulumi.Input

	ToVpcEndpointServiceResourceOutput() VpcEndpointServiceResourceOutput
	ToVpcEndpointServiceResourceOutputWithContext(context.Context) VpcEndpointServiceResourceOutput
}

type VpcEndpointServiceResourceArgs struct {
	// The id of service resources added to the endpoint service.
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
	// The type of service resource added to the endpoint service.
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
}

func (VpcEndpointServiceResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointServiceResource)(nil)).Elem()
}

func (i VpcEndpointServiceResourceArgs) ToVpcEndpointServiceResourceOutput() VpcEndpointServiceResourceOutput {
	return i.ToVpcEndpointServiceResourceOutputWithContext(context.Background())
}

func (i VpcEndpointServiceResourceArgs) ToVpcEndpointServiceResourceOutputWithContext(ctx context.Context) VpcEndpointServiceResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceResourceOutput)
}

// VpcEndpointServiceResourceArrayInput is an input type that accepts VpcEndpointServiceResourceArray and VpcEndpointServiceResourceArrayOutput values.
// You can construct a concrete instance of `VpcEndpointServiceResourceArrayInput` via:
//
//          VpcEndpointServiceResourceArray{ VpcEndpointServiceResourceArgs{...} }
type VpcEndpointServiceResourceArrayInput interface {
	pulumi.Input

	ToVpcEndpointServiceResourceArrayOutput() VpcEndpointServiceResourceArrayOutput
	ToVpcEndpointServiceResourceArrayOutputWithContext(context.Context) VpcEndpointServiceResourceArrayOutput
}

type VpcEndpointServiceResourceArray []VpcEndpointServiceResourceInput

func (VpcEndpointServiceResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcEndpointServiceResource)(nil)).Elem()
}

func (i VpcEndpointServiceResourceArray) ToVpcEndpointServiceResourceArrayOutput() VpcEndpointServiceResourceArrayOutput {
	return i.ToVpcEndpointServiceResourceArrayOutputWithContext(context.Background())
}

func (i VpcEndpointServiceResourceArray) ToVpcEndpointServiceResourceArrayOutputWithContext(ctx context.Context) VpcEndpointServiceResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceResourceArrayOutput)
}

type VpcEndpointServiceResourceOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointServiceResource)(nil)).Elem()
}

func (o VpcEndpointServiceResourceOutput) ToVpcEndpointServiceResourceOutput() VpcEndpointServiceResourceOutput {
	return o
}

func (o VpcEndpointServiceResourceOutput) ToVpcEndpointServiceResourceOutputWithContext(ctx context.Context) VpcEndpointServiceResourceOutput {
	return o
}

// The id of service resources added to the endpoint service.
func (o VpcEndpointServiceResourceOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcEndpointServiceResource) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// The type of service resource added to the endpoint service.
func (o VpcEndpointServiceResourceOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcEndpointServiceResource) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

type VpcEndpointServiceResourceArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcEndpointServiceResource)(nil)).Elem()
}

func (o VpcEndpointServiceResourceArrayOutput) ToVpcEndpointServiceResourceArrayOutput() VpcEndpointServiceResourceArrayOutput {
	return o
}

func (o VpcEndpointServiceResourceArrayOutput) ToVpcEndpointServiceResourceArrayOutputWithContext(ctx context.Context) VpcEndpointServiceResourceArrayOutput {
	return o
}

func (o VpcEndpointServiceResourceArrayOutput) Index(i pulumi.IntInput) VpcEndpointServiceResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpcEndpointServiceResource {
		return vs[0].([]VpcEndpointServiceResource)[vs[1].(int)]
	}).(VpcEndpointServiceResourceOutput)
}

type VpcEndpointZone struct {
	// To create the vswitch of the terminal node network card in the available zone.
	VswitchId *string `pulumi:"vswitchId"`
	// Availability zone corresponding to terminal node service.
	ZoneId *string `pulumi:"zoneId"`
}

// VpcEndpointZoneInput is an input type that accepts VpcEndpointZoneArgs and VpcEndpointZoneOutput values.
// You can construct a concrete instance of `VpcEndpointZoneInput` via:
//
//          VpcEndpointZoneArgs{...}
type VpcEndpointZoneInput interface {
	pulumi.Input

	ToVpcEndpointZoneOutput() VpcEndpointZoneOutput
	ToVpcEndpointZoneOutputWithContext(context.Context) VpcEndpointZoneOutput
}

type VpcEndpointZoneArgs struct {
	// To create the vswitch of the terminal node network card in the available zone.
	VswitchId pulumi.StringPtrInput `pulumi:"vswitchId"`
	// Availability zone corresponding to terminal node service.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (VpcEndpointZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointZone)(nil)).Elem()
}

func (i VpcEndpointZoneArgs) ToVpcEndpointZoneOutput() VpcEndpointZoneOutput {
	return i.ToVpcEndpointZoneOutputWithContext(context.Background())
}

func (i VpcEndpointZoneArgs) ToVpcEndpointZoneOutputWithContext(ctx context.Context) VpcEndpointZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointZoneOutput)
}

// VpcEndpointZoneArrayInput is an input type that accepts VpcEndpointZoneArray and VpcEndpointZoneArrayOutput values.
// You can construct a concrete instance of `VpcEndpointZoneArrayInput` via:
//
//          VpcEndpointZoneArray{ VpcEndpointZoneArgs{...} }
type VpcEndpointZoneArrayInput interface {
	pulumi.Input

	ToVpcEndpointZoneArrayOutput() VpcEndpointZoneArrayOutput
	ToVpcEndpointZoneArrayOutputWithContext(context.Context) VpcEndpointZoneArrayOutput
}

type VpcEndpointZoneArray []VpcEndpointZoneInput

func (VpcEndpointZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcEndpointZone)(nil)).Elem()
}

func (i VpcEndpointZoneArray) ToVpcEndpointZoneArrayOutput() VpcEndpointZoneArrayOutput {
	return i.ToVpcEndpointZoneArrayOutputWithContext(context.Background())
}

func (i VpcEndpointZoneArray) ToVpcEndpointZoneArrayOutputWithContext(ctx context.Context) VpcEndpointZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointZoneArrayOutput)
}

type VpcEndpointZoneOutput struct{ *pulumi.OutputState }

func (VpcEndpointZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointZone)(nil)).Elem()
}

func (o VpcEndpointZoneOutput) ToVpcEndpointZoneOutput() VpcEndpointZoneOutput {
	return o
}

func (o VpcEndpointZoneOutput) ToVpcEndpointZoneOutputWithContext(ctx context.Context) VpcEndpointZoneOutput {
	return o
}

// To create the vswitch of the terminal node network card in the available zone.
func (o VpcEndpointZoneOutput) VswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcEndpointZone) *string { return v.VswitchId }).(pulumi.StringPtrOutput)
}

// Availability zone corresponding to terminal node service.
func (o VpcEndpointZoneOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcEndpointZone) *string { return v.ZoneId }).(pulumi.StringPtrOutput)
}

type VpcEndpointZoneArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcEndpointZone)(nil)).Elem()
}

func (o VpcEndpointZoneArrayOutput) ToVpcEndpointZoneArrayOutput() VpcEndpointZoneArrayOutput {
	return o
}

func (o VpcEndpointZoneArrayOutput) ToVpcEndpointZoneArrayOutputWithContext(ctx context.Context) VpcEndpointZoneArrayOutput {
	return o
}

func (o VpcEndpointZoneArrayOutput) Index(i pulumi.IntInput) VpcEndpointZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpcEndpointZone {
		return vs[0].([]VpcEndpointZone)[vs[1].(int)]
	}).(VpcEndpointZoneOutput)
}

type GetVpcEndpointServicesService struct {
	// Whether to automatically accept terminal node connections..
	AutoAcceptConnection bool `pulumi:"autoAcceptConnection"`
	// The connection bandwidth.
	ConnectBandwidth int `pulumi:"connectBandwidth"`
	// The ID of the Vpc Endpoint Service.
	Id string `pulumi:"id"`
	// Service resources added to the endpoint service.
	Resources []GetVpcEndpointServicesServiceResource `pulumi:"resources"`
	// The business status of the terminal node service..
	ServiceBusinessStatus string `pulumi:"serviceBusinessStatus"`
	// The description of the terminal node service.
	ServiceDescription string `pulumi:"serviceDescription"`
	// The domain of service.
	ServiceDomain string `pulumi:"serviceDomain"`
	// The ID of the Vpc Endpoint Service.
	ServiceId string `pulumi:"serviceId"`
	// The Status of Vpc Endpoint Service.
	Status string `pulumi:"status"`
	// The name of Vpc Endpoint Service.
	VpcEndpointServiceName string `pulumi:"vpcEndpointServiceName"`
}

// GetVpcEndpointServicesServiceInput is an input type that accepts GetVpcEndpointServicesServiceArgs and GetVpcEndpointServicesServiceOutput values.
// You can construct a concrete instance of `GetVpcEndpointServicesServiceInput` via:
//
//          GetVpcEndpointServicesServiceArgs{...}
type GetVpcEndpointServicesServiceInput interface {
	pulumi.Input

	ToGetVpcEndpointServicesServiceOutput() GetVpcEndpointServicesServiceOutput
	ToGetVpcEndpointServicesServiceOutputWithContext(context.Context) GetVpcEndpointServicesServiceOutput
}

type GetVpcEndpointServicesServiceArgs struct {
	// Whether to automatically accept terminal node connections..
	AutoAcceptConnection pulumi.BoolInput `pulumi:"autoAcceptConnection"`
	// The connection bandwidth.
	ConnectBandwidth pulumi.IntInput `pulumi:"connectBandwidth"`
	// The ID of the Vpc Endpoint Service.
	Id pulumi.StringInput `pulumi:"id"`
	// Service resources added to the endpoint service.
	Resources GetVpcEndpointServicesServiceResourceArrayInput `pulumi:"resources"`
	// The business status of the terminal node service..
	ServiceBusinessStatus pulumi.StringInput `pulumi:"serviceBusinessStatus"`
	// The description of the terminal node service.
	ServiceDescription pulumi.StringInput `pulumi:"serviceDescription"`
	// The domain of service.
	ServiceDomain pulumi.StringInput `pulumi:"serviceDomain"`
	// The ID of the Vpc Endpoint Service.
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
	// The Status of Vpc Endpoint Service.
	Status pulumi.StringInput `pulumi:"status"`
	// The name of Vpc Endpoint Service.
	VpcEndpointServiceName pulumi.StringInput `pulumi:"vpcEndpointServiceName"`
}

func (GetVpcEndpointServicesServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcEndpointServicesService)(nil)).Elem()
}

func (i GetVpcEndpointServicesServiceArgs) ToGetVpcEndpointServicesServiceOutput() GetVpcEndpointServicesServiceOutput {
	return i.ToGetVpcEndpointServicesServiceOutputWithContext(context.Background())
}

func (i GetVpcEndpointServicesServiceArgs) ToGetVpcEndpointServicesServiceOutputWithContext(ctx context.Context) GetVpcEndpointServicesServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVpcEndpointServicesServiceOutput)
}

// GetVpcEndpointServicesServiceArrayInput is an input type that accepts GetVpcEndpointServicesServiceArray and GetVpcEndpointServicesServiceArrayOutput values.
// You can construct a concrete instance of `GetVpcEndpointServicesServiceArrayInput` via:
//
//          GetVpcEndpointServicesServiceArray{ GetVpcEndpointServicesServiceArgs{...} }
type GetVpcEndpointServicesServiceArrayInput interface {
	pulumi.Input

	ToGetVpcEndpointServicesServiceArrayOutput() GetVpcEndpointServicesServiceArrayOutput
	ToGetVpcEndpointServicesServiceArrayOutputWithContext(context.Context) GetVpcEndpointServicesServiceArrayOutput
}

type GetVpcEndpointServicesServiceArray []GetVpcEndpointServicesServiceInput

func (GetVpcEndpointServicesServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVpcEndpointServicesService)(nil)).Elem()
}

func (i GetVpcEndpointServicesServiceArray) ToGetVpcEndpointServicesServiceArrayOutput() GetVpcEndpointServicesServiceArrayOutput {
	return i.ToGetVpcEndpointServicesServiceArrayOutputWithContext(context.Background())
}

func (i GetVpcEndpointServicesServiceArray) ToGetVpcEndpointServicesServiceArrayOutputWithContext(ctx context.Context) GetVpcEndpointServicesServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVpcEndpointServicesServiceArrayOutput)
}

type GetVpcEndpointServicesServiceOutput struct{ *pulumi.OutputState }

func (GetVpcEndpointServicesServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcEndpointServicesService)(nil)).Elem()
}

func (o GetVpcEndpointServicesServiceOutput) ToGetVpcEndpointServicesServiceOutput() GetVpcEndpointServicesServiceOutput {
	return o
}

func (o GetVpcEndpointServicesServiceOutput) ToGetVpcEndpointServicesServiceOutputWithContext(ctx context.Context) GetVpcEndpointServicesServiceOutput {
	return o
}

// Whether to automatically accept terminal node connections..
func (o GetVpcEndpointServicesServiceOutput) AutoAcceptConnection() pulumi.BoolOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesService) bool { return v.AutoAcceptConnection }).(pulumi.BoolOutput)
}

// The connection bandwidth.
func (o GetVpcEndpointServicesServiceOutput) ConnectBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesService) int { return v.ConnectBandwidth }).(pulumi.IntOutput)
}

// The ID of the Vpc Endpoint Service.
func (o GetVpcEndpointServicesServiceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesService) string { return v.Id }).(pulumi.StringOutput)
}

// Service resources added to the endpoint service.
func (o GetVpcEndpointServicesServiceOutput) Resources() GetVpcEndpointServicesServiceResourceArrayOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesService) []GetVpcEndpointServicesServiceResource { return v.Resources }).(GetVpcEndpointServicesServiceResourceArrayOutput)
}

// The business status of the terminal node service..
func (o GetVpcEndpointServicesServiceOutput) ServiceBusinessStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesService) string { return v.ServiceBusinessStatus }).(pulumi.StringOutput)
}

// The description of the terminal node service.
func (o GetVpcEndpointServicesServiceOutput) ServiceDescription() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesService) string { return v.ServiceDescription }).(pulumi.StringOutput)
}

// The domain of service.
func (o GetVpcEndpointServicesServiceOutput) ServiceDomain() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesService) string { return v.ServiceDomain }).(pulumi.StringOutput)
}

// The ID of the Vpc Endpoint Service.
func (o GetVpcEndpointServicesServiceOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesService) string { return v.ServiceId }).(pulumi.StringOutput)
}

// The Status of Vpc Endpoint Service.
func (o GetVpcEndpointServicesServiceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesService) string { return v.Status }).(pulumi.StringOutput)
}

// The name of Vpc Endpoint Service.
func (o GetVpcEndpointServicesServiceOutput) VpcEndpointServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesService) string { return v.VpcEndpointServiceName }).(pulumi.StringOutput)
}

type GetVpcEndpointServicesServiceArrayOutput struct{ *pulumi.OutputState }

func (GetVpcEndpointServicesServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVpcEndpointServicesService)(nil)).Elem()
}

func (o GetVpcEndpointServicesServiceArrayOutput) ToGetVpcEndpointServicesServiceArrayOutput() GetVpcEndpointServicesServiceArrayOutput {
	return o
}

func (o GetVpcEndpointServicesServiceArrayOutput) ToGetVpcEndpointServicesServiceArrayOutputWithContext(ctx context.Context) GetVpcEndpointServicesServiceArrayOutput {
	return o
}

func (o GetVpcEndpointServicesServiceArrayOutput) Index(i pulumi.IntInput) GetVpcEndpointServicesServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetVpcEndpointServicesService {
		return vs[0].([]GetVpcEndpointServicesService)[vs[1].(int)]
	}).(GetVpcEndpointServicesServiceOutput)
}

type GetVpcEndpointServicesServiceResource struct {
	// The id of service resources added to the endpoint service.
	ResourceId string `pulumi:"resourceId"`
	// The type of service resource added to the endpoint service.
	ResourceType string `pulumi:"resourceType"`
}

// GetVpcEndpointServicesServiceResourceInput is an input type that accepts GetVpcEndpointServicesServiceResourceArgs and GetVpcEndpointServicesServiceResourceOutput values.
// You can construct a concrete instance of `GetVpcEndpointServicesServiceResourceInput` via:
//
//          GetVpcEndpointServicesServiceResourceArgs{...}
type GetVpcEndpointServicesServiceResourceInput interface {
	pulumi.Input

	ToGetVpcEndpointServicesServiceResourceOutput() GetVpcEndpointServicesServiceResourceOutput
	ToGetVpcEndpointServicesServiceResourceOutputWithContext(context.Context) GetVpcEndpointServicesServiceResourceOutput
}

type GetVpcEndpointServicesServiceResourceArgs struct {
	// The id of service resources added to the endpoint service.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// The type of service resource added to the endpoint service.
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
}

func (GetVpcEndpointServicesServiceResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcEndpointServicesServiceResource)(nil)).Elem()
}

func (i GetVpcEndpointServicesServiceResourceArgs) ToGetVpcEndpointServicesServiceResourceOutput() GetVpcEndpointServicesServiceResourceOutput {
	return i.ToGetVpcEndpointServicesServiceResourceOutputWithContext(context.Background())
}

func (i GetVpcEndpointServicesServiceResourceArgs) ToGetVpcEndpointServicesServiceResourceOutputWithContext(ctx context.Context) GetVpcEndpointServicesServiceResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVpcEndpointServicesServiceResourceOutput)
}

// GetVpcEndpointServicesServiceResourceArrayInput is an input type that accepts GetVpcEndpointServicesServiceResourceArray and GetVpcEndpointServicesServiceResourceArrayOutput values.
// You can construct a concrete instance of `GetVpcEndpointServicesServiceResourceArrayInput` via:
//
//          GetVpcEndpointServicesServiceResourceArray{ GetVpcEndpointServicesServiceResourceArgs{...} }
type GetVpcEndpointServicesServiceResourceArrayInput interface {
	pulumi.Input

	ToGetVpcEndpointServicesServiceResourceArrayOutput() GetVpcEndpointServicesServiceResourceArrayOutput
	ToGetVpcEndpointServicesServiceResourceArrayOutputWithContext(context.Context) GetVpcEndpointServicesServiceResourceArrayOutput
}

type GetVpcEndpointServicesServiceResourceArray []GetVpcEndpointServicesServiceResourceInput

func (GetVpcEndpointServicesServiceResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVpcEndpointServicesServiceResource)(nil)).Elem()
}

func (i GetVpcEndpointServicesServiceResourceArray) ToGetVpcEndpointServicesServiceResourceArrayOutput() GetVpcEndpointServicesServiceResourceArrayOutput {
	return i.ToGetVpcEndpointServicesServiceResourceArrayOutputWithContext(context.Background())
}

func (i GetVpcEndpointServicesServiceResourceArray) ToGetVpcEndpointServicesServiceResourceArrayOutputWithContext(ctx context.Context) GetVpcEndpointServicesServiceResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVpcEndpointServicesServiceResourceArrayOutput)
}

type GetVpcEndpointServicesServiceResourceOutput struct{ *pulumi.OutputState }

func (GetVpcEndpointServicesServiceResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcEndpointServicesServiceResource)(nil)).Elem()
}

func (o GetVpcEndpointServicesServiceResourceOutput) ToGetVpcEndpointServicesServiceResourceOutput() GetVpcEndpointServicesServiceResourceOutput {
	return o
}

func (o GetVpcEndpointServicesServiceResourceOutput) ToGetVpcEndpointServicesServiceResourceOutputWithContext(ctx context.Context) GetVpcEndpointServicesServiceResourceOutput {
	return o
}

// The id of service resources added to the endpoint service.
func (o GetVpcEndpointServicesServiceResourceOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesServiceResource) string { return v.ResourceId }).(pulumi.StringOutput)
}

// The type of service resource added to the endpoint service.
func (o GetVpcEndpointServicesServiceResourceOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesServiceResource) string { return v.ResourceType }).(pulumi.StringOutput)
}

type GetVpcEndpointServicesServiceResourceArrayOutput struct{ *pulumi.OutputState }

func (GetVpcEndpointServicesServiceResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVpcEndpointServicesServiceResource)(nil)).Elem()
}

func (o GetVpcEndpointServicesServiceResourceArrayOutput) ToGetVpcEndpointServicesServiceResourceArrayOutput() GetVpcEndpointServicesServiceResourceArrayOutput {
	return o
}

func (o GetVpcEndpointServicesServiceResourceArrayOutput) ToGetVpcEndpointServicesServiceResourceArrayOutputWithContext(ctx context.Context) GetVpcEndpointServicesServiceResourceArrayOutput {
	return o
}

func (o GetVpcEndpointServicesServiceResourceArrayOutput) Index(i pulumi.IntInput) GetVpcEndpointServicesServiceResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetVpcEndpointServicesServiceResource {
		return vs[0].([]GetVpcEndpointServicesServiceResource)[vs[1].(int)]
	}).(GetVpcEndpointServicesServiceResourceOutput)
}

type GetVpcEndpointsEndpoint struct {
	// The Bandwidth.
	Bandwidth int `pulumi:"bandwidth"`
	// The status of Connection.
	ConnectionStatus string `pulumi:"connectionStatus"`
	// The status of Endpoint Business.
	EndpointBusinessStatus string `pulumi:"endpointBusinessStatus"`
	// The description of Vpc Endpoint.
	EndpointDescription string `pulumi:"endpointDescription"`
	// The Endpoint Domain.
	EndpointDomain string `pulumi:"endpointDomain"`
	// The ID of the Vpc Endpoint.
	EndpointId string `pulumi:"endpointId"`
	// The ID of the Vpc Endpoint.
	Id string `pulumi:"id"`
	// The security group associated with the terminal node network card.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The terminal node service associated with the terminal node..
	ServiceId string `pulumi:"serviceId"`
	// The name of the terminal node service associated with the terminal node.
	ServiceName string `pulumi:"serviceName"`
	// The status of Vpc Endpoint.
	Status string `pulumi:"status"`
	// The name of Vpc Endpoint.
	VpcEndpointName string `pulumi:"vpcEndpointName"`
	// The private network to which the terminal node belongs.
	VpcId string `pulumi:"vpcId"`
	// Availability zone.
	Zones []GetVpcEndpointsEndpointZone `pulumi:"zones"`
}

// GetVpcEndpointsEndpointInput is an input type that accepts GetVpcEndpointsEndpointArgs and GetVpcEndpointsEndpointOutput values.
// You can construct a concrete instance of `GetVpcEndpointsEndpointInput` via:
//
//          GetVpcEndpointsEndpointArgs{...}
type GetVpcEndpointsEndpointInput interface {
	pulumi.Input

	ToGetVpcEndpointsEndpointOutput() GetVpcEndpointsEndpointOutput
	ToGetVpcEndpointsEndpointOutputWithContext(context.Context) GetVpcEndpointsEndpointOutput
}

type GetVpcEndpointsEndpointArgs struct {
	// The Bandwidth.
	Bandwidth pulumi.IntInput `pulumi:"bandwidth"`
	// The status of Connection.
	ConnectionStatus pulumi.StringInput `pulumi:"connectionStatus"`
	// The status of Endpoint Business.
	EndpointBusinessStatus pulumi.StringInput `pulumi:"endpointBusinessStatus"`
	// The description of Vpc Endpoint.
	EndpointDescription pulumi.StringInput `pulumi:"endpointDescription"`
	// The Endpoint Domain.
	EndpointDomain pulumi.StringInput `pulumi:"endpointDomain"`
	// The ID of the Vpc Endpoint.
	EndpointId pulumi.StringInput `pulumi:"endpointId"`
	// The ID of the Vpc Endpoint.
	Id pulumi.StringInput `pulumi:"id"`
	// The security group associated with the terminal node network card.
	SecurityGroupIds pulumi.StringArrayInput `pulumi:"securityGroupIds"`
	// The terminal node service associated with the terminal node..
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
	// The name of the terminal node service associated with the terminal node.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// The status of Vpc Endpoint.
	Status pulumi.StringInput `pulumi:"status"`
	// The name of Vpc Endpoint.
	VpcEndpointName pulumi.StringInput `pulumi:"vpcEndpointName"`
	// The private network to which the terminal node belongs.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
	// Availability zone.
	Zones GetVpcEndpointsEndpointZoneArrayInput `pulumi:"zones"`
}

func (GetVpcEndpointsEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcEndpointsEndpoint)(nil)).Elem()
}

func (i GetVpcEndpointsEndpointArgs) ToGetVpcEndpointsEndpointOutput() GetVpcEndpointsEndpointOutput {
	return i.ToGetVpcEndpointsEndpointOutputWithContext(context.Background())
}

func (i GetVpcEndpointsEndpointArgs) ToGetVpcEndpointsEndpointOutputWithContext(ctx context.Context) GetVpcEndpointsEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVpcEndpointsEndpointOutput)
}

// GetVpcEndpointsEndpointArrayInput is an input type that accepts GetVpcEndpointsEndpointArray and GetVpcEndpointsEndpointArrayOutput values.
// You can construct a concrete instance of `GetVpcEndpointsEndpointArrayInput` via:
//
//          GetVpcEndpointsEndpointArray{ GetVpcEndpointsEndpointArgs{...} }
type GetVpcEndpointsEndpointArrayInput interface {
	pulumi.Input

	ToGetVpcEndpointsEndpointArrayOutput() GetVpcEndpointsEndpointArrayOutput
	ToGetVpcEndpointsEndpointArrayOutputWithContext(context.Context) GetVpcEndpointsEndpointArrayOutput
}

type GetVpcEndpointsEndpointArray []GetVpcEndpointsEndpointInput

func (GetVpcEndpointsEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVpcEndpointsEndpoint)(nil)).Elem()
}

func (i GetVpcEndpointsEndpointArray) ToGetVpcEndpointsEndpointArrayOutput() GetVpcEndpointsEndpointArrayOutput {
	return i.ToGetVpcEndpointsEndpointArrayOutputWithContext(context.Background())
}

func (i GetVpcEndpointsEndpointArray) ToGetVpcEndpointsEndpointArrayOutputWithContext(ctx context.Context) GetVpcEndpointsEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVpcEndpointsEndpointArrayOutput)
}

type GetVpcEndpointsEndpointOutput struct{ *pulumi.OutputState }

func (GetVpcEndpointsEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcEndpointsEndpoint)(nil)).Elem()
}

func (o GetVpcEndpointsEndpointOutput) ToGetVpcEndpointsEndpointOutput() GetVpcEndpointsEndpointOutput {
	return o
}

func (o GetVpcEndpointsEndpointOutput) ToGetVpcEndpointsEndpointOutputWithContext(ctx context.Context) GetVpcEndpointsEndpointOutput {
	return o
}

// The Bandwidth.
func (o GetVpcEndpointsEndpointOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpoint) int { return v.Bandwidth }).(pulumi.IntOutput)
}

// The status of Connection.
func (o GetVpcEndpointsEndpointOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpoint) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

// The status of Endpoint Business.
func (o GetVpcEndpointsEndpointOutput) EndpointBusinessStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpoint) string { return v.EndpointBusinessStatus }).(pulumi.StringOutput)
}

// The description of Vpc Endpoint.
func (o GetVpcEndpointsEndpointOutput) EndpointDescription() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpoint) string { return v.EndpointDescription }).(pulumi.StringOutput)
}

// The Endpoint Domain.
func (o GetVpcEndpointsEndpointOutput) EndpointDomain() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpoint) string { return v.EndpointDomain }).(pulumi.StringOutput)
}

// The ID of the Vpc Endpoint.
func (o GetVpcEndpointsEndpointOutput) EndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpoint) string { return v.EndpointId }).(pulumi.StringOutput)
}

// The ID of the Vpc Endpoint.
func (o GetVpcEndpointsEndpointOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpoint) string { return v.Id }).(pulumi.StringOutput)
}

// The security group associated with the terminal node network card.
func (o GetVpcEndpointsEndpointOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpoint) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The terminal node service associated with the terminal node..
func (o GetVpcEndpointsEndpointOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpoint) string { return v.ServiceId }).(pulumi.StringOutput)
}

// The name of the terminal node service associated with the terminal node.
func (o GetVpcEndpointsEndpointOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpoint) string { return v.ServiceName }).(pulumi.StringOutput)
}

// The status of Vpc Endpoint.
func (o GetVpcEndpointsEndpointOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpoint) string { return v.Status }).(pulumi.StringOutput)
}

// The name of Vpc Endpoint.
func (o GetVpcEndpointsEndpointOutput) VpcEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpoint) string { return v.VpcEndpointName }).(pulumi.StringOutput)
}

// The private network to which the terminal node belongs.
func (o GetVpcEndpointsEndpointOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpoint) string { return v.VpcId }).(pulumi.StringOutput)
}

// Availability zone.
func (o GetVpcEndpointsEndpointOutput) Zones() GetVpcEndpointsEndpointZoneArrayOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpoint) []GetVpcEndpointsEndpointZone { return v.Zones }).(GetVpcEndpointsEndpointZoneArrayOutput)
}

type GetVpcEndpointsEndpointArrayOutput struct{ *pulumi.OutputState }

func (GetVpcEndpointsEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVpcEndpointsEndpoint)(nil)).Elem()
}

func (o GetVpcEndpointsEndpointArrayOutput) ToGetVpcEndpointsEndpointArrayOutput() GetVpcEndpointsEndpointArrayOutput {
	return o
}

func (o GetVpcEndpointsEndpointArrayOutput) ToGetVpcEndpointsEndpointArrayOutputWithContext(ctx context.Context) GetVpcEndpointsEndpointArrayOutput {
	return o
}

func (o GetVpcEndpointsEndpointArrayOutput) Index(i pulumi.IntInput) GetVpcEndpointsEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetVpcEndpointsEndpoint {
		return vs[0].([]GetVpcEndpointsEndpoint)[vs[1].(int)]
	}).(GetVpcEndpointsEndpointOutput)
}

type GetVpcEndpointsEndpointZone struct {
	// To create the vswitch of the terminal node network card in the available zone.
	VswitchId string `pulumi:"vswitchId"`
	// Availability zone corresponding to terminal node service.
	ZoneId string `pulumi:"zoneId"`
}

// GetVpcEndpointsEndpointZoneInput is an input type that accepts GetVpcEndpointsEndpointZoneArgs and GetVpcEndpointsEndpointZoneOutput values.
// You can construct a concrete instance of `GetVpcEndpointsEndpointZoneInput` via:
//
//          GetVpcEndpointsEndpointZoneArgs{...}
type GetVpcEndpointsEndpointZoneInput interface {
	pulumi.Input

	ToGetVpcEndpointsEndpointZoneOutput() GetVpcEndpointsEndpointZoneOutput
	ToGetVpcEndpointsEndpointZoneOutputWithContext(context.Context) GetVpcEndpointsEndpointZoneOutput
}

type GetVpcEndpointsEndpointZoneArgs struct {
	// To create the vswitch of the terminal node network card in the available zone.
	VswitchId pulumi.StringInput `pulumi:"vswitchId"`
	// Availability zone corresponding to terminal node service.
	ZoneId pulumi.StringInput `pulumi:"zoneId"`
}

func (GetVpcEndpointsEndpointZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcEndpointsEndpointZone)(nil)).Elem()
}

func (i GetVpcEndpointsEndpointZoneArgs) ToGetVpcEndpointsEndpointZoneOutput() GetVpcEndpointsEndpointZoneOutput {
	return i.ToGetVpcEndpointsEndpointZoneOutputWithContext(context.Background())
}

func (i GetVpcEndpointsEndpointZoneArgs) ToGetVpcEndpointsEndpointZoneOutputWithContext(ctx context.Context) GetVpcEndpointsEndpointZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVpcEndpointsEndpointZoneOutput)
}

// GetVpcEndpointsEndpointZoneArrayInput is an input type that accepts GetVpcEndpointsEndpointZoneArray and GetVpcEndpointsEndpointZoneArrayOutput values.
// You can construct a concrete instance of `GetVpcEndpointsEndpointZoneArrayInput` via:
//
//          GetVpcEndpointsEndpointZoneArray{ GetVpcEndpointsEndpointZoneArgs{...} }
type GetVpcEndpointsEndpointZoneArrayInput interface {
	pulumi.Input

	ToGetVpcEndpointsEndpointZoneArrayOutput() GetVpcEndpointsEndpointZoneArrayOutput
	ToGetVpcEndpointsEndpointZoneArrayOutputWithContext(context.Context) GetVpcEndpointsEndpointZoneArrayOutput
}

type GetVpcEndpointsEndpointZoneArray []GetVpcEndpointsEndpointZoneInput

func (GetVpcEndpointsEndpointZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVpcEndpointsEndpointZone)(nil)).Elem()
}

func (i GetVpcEndpointsEndpointZoneArray) ToGetVpcEndpointsEndpointZoneArrayOutput() GetVpcEndpointsEndpointZoneArrayOutput {
	return i.ToGetVpcEndpointsEndpointZoneArrayOutputWithContext(context.Background())
}

func (i GetVpcEndpointsEndpointZoneArray) ToGetVpcEndpointsEndpointZoneArrayOutputWithContext(ctx context.Context) GetVpcEndpointsEndpointZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVpcEndpointsEndpointZoneArrayOutput)
}

type GetVpcEndpointsEndpointZoneOutput struct{ *pulumi.OutputState }

func (GetVpcEndpointsEndpointZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcEndpointsEndpointZone)(nil)).Elem()
}

func (o GetVpcEndpointsEndpointZoneOutput) ToGetVpcEndpointsEndpointZoneOutput() GetVpcEndpointsEndpointZoneOutput {
	return o
}

func (o GetVpcEndpointsEndpointZoneOutput) ToGetVpcEndpointsEndpointZoneOutputWithContext(ctx context.Context) GetVpcEndpointsEndpointZoneOutput {
	return o
}

// To create the vswitch of the terminal node network card in the available zone.
func (o GetVpcEndpointsEndpointZoneOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpointZone) string { return v.VswitchId }).(pulumi.StringOutput)
}

// Availability zone corresponding to terminal node service.
func (o GetVpcEndpointsEndpointZoneOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointsEndpointZone) string { return v.ZoneId }).(pulumi.StringOutput)
}

type GetVpcEndpointsEndpointZoneArrayOutput struct{ *pulumi.OutputState }

func (GetVpcEndpointsEndpointZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVpcEndpointsEndpointZone)(nil)).Elem()
}

func (o GetVpcEndpointsEndpointZoneArrayOutput) ToGetVpcEndpointsEndpointZoneArrayOutput() GetVpcEndpointsEndpointZoneArrayOutput {
	return o
}

func (o GetVpcEndpointsEndpointZoneArrayOutput) ToGetVpcEndpointsEndpointZoneArrayOutputWithContext(ctx context.Context) GetVpcEndpointsEndpointZoneArrayOutput {
	return o
}

func (o GetVpcEndpointsEndpointZoneArrayOutput) Index(i pulumi.IntInput) GetVpcEndpointsEndpointZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetVpcEndpointsEndpointZone {
		return vs[0].([]GetVpcEndpointsEndpointZone)[vs[1].(int)]
	}).(GetVpcEndpointsEndpointZoneOutput)
}

func init() {
	pulumi.RegisterOutputType(VpcEndpointServiceResourceOutput{})
	pulumi.RegisterOutputType(VpcEndpointServiceResourceArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointZoneOutput{})
	pulumi.RegisterOutputType(VpcEndpointZoneArrayOutput{})
	pulumi.RegisterOutputType(GetVpcEndpointServicesServiceOutput{})
	pulumi.RegisterOutputType(GetVpcEndpointServicesServiceArrayOutput{})
	pulumi.RegisterOutputType(GetVpcEndpointServicesServiceResourceOutput{})
	pulumi.RegisterOutputType(GetVpcEndpointServicesServiceResourceArrayOutput{})
	pulumi.RegisterOutputType(GetVpcEndpointsEndpointOutput{})
	pulumi.RegisterOutputType(GetVpcEndpointsEndpointArrayOutput{})
	pulumi.RegisterOutputType(GetVpcEndpointsEndpointZoneOutput{})
	pulumi.RegisterOutputType(GetVpcEndpointsEndpointZoneArrayOutput{})
}
