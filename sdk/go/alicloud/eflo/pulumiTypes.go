// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eflo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetVpdsVpd struct {
	// CIDR network segment
	Cidr string `pulumi:"cidr"`
	// The creation time of the resource
	CreateTime string `pulumi:"createTime"`
	// Modification time
	GmtModified string `pulumi:"gmtModified"`
	// The id of the vpd.
	Id string `pulumi:"id"`
	// The Resource group id
	ResourceGroupId string `pulumi:"resourceGroupId"`
	// The Vpd status. Valid values: `Available`, `Not Available`, `Executing`, `Deleting`,
	Status string `pulumi:"status"`
	// The id of the vpd.
	VpdId string `pulumi:"vpdId"`
	// The Name of the VPD.
	VpdName string `pulumi:"vpdName"`
}

// GetVpdsVpdInput is an input type that accepts GetVpdsVpdArgs and GetVpdsVpdOutput values.
// You can construct a concrete instance of `GetVpdsVpdInput` via:
//
//	GetVpdsVpdArgs{...}
type GetVpdsVpdInput interface {
	pulumi.Input

	ToGetVpdsVpdOutput() GetVpdsVpdOutput
	ToGetVpdsVpdOutputWithContext(context.Context) GetVpdsVpdOutput
}

type GetVpdsVpdArgs struct {
	// CIDR network segment
	Cidr pulumi.StringInput `pulumi:"cidr"`
	// The creation time of the resource
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// Modification time
	GmtModified pulumi.StringInput `pulumi:"gmtModified"`
	// The id of the vpd.
	Id pulumi.StringInput `pulumi:"id"`
	// The Resource group id
	ResourceGroupId pulumi.StringInput `pulumi:"resourceGroupId"`
	// The Vpd status. Valid values: `Available`, `Not Available`, `Executing`, `Deleting`,
	Status pulumi.StringInput `pulumi:"status"`
	// The id of the vpd.
	VpdId pulumi.StringInput `pulumi:"vpdId"`
	// The Name of the VPD.
	VpdName pulumi.StringInput `pulumi:"vpdName"`
}

func (GetVpdsVpdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpdsVpd)(nil)).Elem()
}

func (i GetVpdsVpdArgs) ToGetVpdsVpdOutput() GetVpdsVpdOutput {
	return i.ToGetVpdsVpdOutputWithContext(context.Background())
}

func (i GetVpdsVpdArgs) ToGetVpdsVpdOutputWithContext(ctx context.Context) GetVpdsVpdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVpdsVpdOutput)
}

// GetVpdsVpdArrayInput is an input type that accepts GetVpdsVpdArray and GetVpdsVpdArrayOutput values.
// You can construct a concrete instance of `GetVpdsVpdArrayInput` via:
//
//	GetVpdsVpdArray{ GetVpdsVpdArgs{...} }
type GetVpdsVpdArrayInput interface {
	pulumi.Input

	ToGetVpdsVpdArrayOutput() GetVpdsVpdArrayOutput
	ToGetVpdsVpdArrayOutputWithContext(context.Context) GetVpdsVpdArrayOutput
}

type GetVpdsVpdArray []GetVpdsVpdInput

func (GetVpdsVpdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVpdsVpd)(nil)).Elem()
}

func (i GetVpdsVpdArray) ToGetVpdsVpdArrayOutput() GetVpdsVpdArrayOutput {
	return i.ToGetVpdsVpdArrayOutputWithContext(context.Background())
}

func (i GetVpdsVpdArray) ToGetVpdsVpdArrayOutputWithContext(ctx context.Context) GetVpdsVpdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVpdsVpdArrayOutput)
}

type GetVpdsVpdOutput struct{ *pulumi.OutputState }

func (GetVpdsVpdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpdsVpd)(nil)).Elem()
}

func (o GetVpdsVpdOutput) ToGetVpdsVpdOutput() GetVpdsVpdOutput {
	return o
}

func (o GetVpdsVpdOutput) ToGetVpdsVpdOutputWithContext(ctx context.Context) GetVpdsVpdOutput {
	return o
}

// CIDR network segment
func (o GetVpdsVpdOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.Cidr }).(pulumi.StringOutput)
}

// The creation time of the resource
func (o GetVpdsVpdOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Modification time
func (o GetVpdsVpdOutput) GmtModified() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.GmtModified }).(pulumi.StringOutput)
}

// The id of the vpd.
func (o GetVpdsVpdOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.Id }).(pulumi.StringOutput)
}

// The Resource group id
func (o GetVpdsVpdOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The Vpd status. Valid values: `Available`, `Not Available`, `Executing`, `Deleting`,
func (o GetVpdsVpdOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.Status }).(pulumi.StringOutput)
}

// The id of the vpd.
func (o GetVpdsVpdOutput) VpdId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.VpdId }).(pulumi.StringOutput)
}

// The Name of the VPD.
func (o GetVpdsVpdOutput) VpdName() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.VpdName }).(pulumi.StringOutput)
}

type GetVpdsVpdArrayOutput struct{ *pulumi.OutputState }

func (GetVpdsVpdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVpdsVpd)(nil)).Elem()
}

func (o GetVpdsVpdArrayOutput) ToGetVpdsVpdArrayOutput() GetVpdsVpdArrayOutput {
	return o
}

func (o GetVpdsVpdArrayOutput) ToGetVpdsVpdArrayOutputWithContext(ctx context.Context) GetVpdsVpdArrayOutput {
	return o
}

func (o GetVpdsVpdArrayOutput) Index(i pulumi.IntInput) GetVpdsVpdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetVpdsVpd {
		return vs[0].([]GetVpdsVpd)[vs[1].(int)]
	}).(GetVpdsVpdOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetVpdsVpdInput)(nil)).Elem(), GetVpdsVpdArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetVpdsVpdArrayInput)(nil)).Elem(), GetVpdsVpdArray{})
	pulumi.RegisterOutputType(GetVpdsVpdOutput{})
	pulumi.RegisterOutputType(GetVpdsVpdArrayOutput{})
}