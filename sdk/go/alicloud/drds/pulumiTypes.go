// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package drds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type GetInstancesInstance struct {
	// Creation time of the instance.
	CreateTime int `pulumi:"createTime"`
	// The DRDS instance description.
	Description string `pulumi:"description"`
	// The ID of the DRDS instance.
	Id string `pulumi:"id"`
	// `Classic` for public classic network or `VPC` for private network.
	NetworkType string `pulumi:"networkType"`
	// Status of the instance.
	Status string `pulumi:"status"`
	// The DRDS Instance type.
	Type string `pulumi:"type"`
	// The DRDS Instance version.
	Version int `pulumi:"version"`
	// Zone ID the instance belongs to.
	ZoneId string `pulumi:"zoneId"`
}

// GetInstancesInstanceInput is an input type that accepts GetInstancesInstanceArgs and GetInstancesInstanceOutput values.
// You can construct a concrete instance of `GetInstancesInstanceInput` via:
//
// 		 GetInstancesInstanceArgs{...}
//
type GetInstancesInstanceInput interface {
	pulumi.Input

	ToGetInstancesInstanceOutput() GetInstancesInstanceOutput
	ToGetInstancesInstanceOutputWithContext(context.Context) GetInstancesInstanceOutput
}

type GetInstancesInstanceArgs struct {
	// Creation time of the instance.
	CreateTime pulumi.IntInput `pulumi:"createTime"`
	// The DRDS instance description.
	Description pulumi.StringInput `pulumi:"description"`
	// The ID of the DRDS instance.
	Id pulumi.StringInput `pulumi:"id"`
	// `Classic` for public classic network or `VPC` for private network.
	NetworkType pulumi.StringInput `pulumi:"networkType"`
	// Status of the instance.
	Status pulumi.StringInput `pulumi:"status"`
	// The DRDS Instance type.
	Type pulumi.StringInput `pulumi:"type"`
	// The DRDS Instance version.
	Version pulumi.IntInput `pulumi:"version"`
	// Zone ID the instance belongs to.
	ZoneId pulumi.StringInput `pulumi:"zoneId"`
}

func (GetInstancesInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesInstance)(nil)).Elem()
}

func (i GetInstancesInstanceArgs) ToGetInstancesInstanceOutput() GetInstancesInstanceOutput {
	return i.ToGetInstancesInstanceOutputWithContext(context.Background())
}

func (i GetInstancesInstanceArgs) ToGetInstancesInstanceOutputWithContext(ctx context.Context) GetInstancesInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstancesInstanceOutput)
}

// GetInstancesInstanceArrayInput is an input type that accepts GetInstancesInstanceArray and GetInstancesInstanceArrayOutput values.
// You can construct a concrete instance of `GetInstancesInstanceArrayInput` via:
//
// 		 GetInstancesInstanceArray{ GetInstancesInstanceArgs{...} }
//
type GetInstancesInstanceArrayInput interface {
	pulumi.Input

	ToGetInstancesInstanceArrayOutput() GetInstancesInstanceArrayOutput
	ToGetInstancesInstanceArrayOutputWithContext(context.Context) GetInstancesInstanceArrayOutput
}

type GetInstancesInstanceArray []GetInstancesInstanceInput

func (GetInstancesInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstancesInstance)(nil)).Elem()
}

func (i GetInstancesInstanceArray) ToGetInstancesInstanceArrayOutput() GetInstancesInstanceArrayOutput {
	return i.ToGetInstancesInstanceArrayOutputWithContext(context.Background())
}

func (i GetInstancesInstanceArray) ToGetInstancesInstanceArrayOutputWithContext(ctx context.Context) GetInstancesInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstancesInstanceArrayOutput)
}

type GetInstancesInstanceOutput struct{ *pulumi.OutputState }

func (GetInstancesInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesInstance)(nil)).Elem()
}

func (o GetInstancesInstanceOutput) ToGetInstancesInstanceOutput() GetInstancesInstanceOutput {
	return o
}

func (o GetInstancesInstanceOutput) ToGetInstancesInstanceOutputWithContext(ctx context.Context) GetInstancesInstanceOutput {
	return o
}

// Creation time of the instance.
func (o GetInstancesInstanceOutput) CreateTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstancesInstance) int { return v.CreateTime }).(pulumi.IntOutput)
}

// The DRDS instance description.
func (o GetInstancesInstanceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesInstance) string { return v.Description }).(pulumi.StringOutput)
}

// The ID of the DRDS instance.
func (o GetInstancesInstanceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesInstance) string { return v.Id }).(pulumi.StringOutput)
}

// `Classic` for public classic network or `VPC` for private network.
func (o GetInstancesInstanceOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesInstance) string { return v.NetworkType }).(pulumi.StringOutput)
}

// Status of the instance.
func (o GetInstancesInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesInstance) string { return v.Status }).(pulumi.StringOutput)
}

// The DRDS Instance type.
func (o GetInstancesInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesInstance) string { return v.Type }).(pulumi.StringOutput)
}

// The DRDS Instance version.
func (o GetInstancesInstanceOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstancesInstance) int { return v.Version }).(pulumi.IntOutput)
}

// Zone ID the instance belongs to.
func (o GetInstancesInstanceOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesInstance) string { return v.ZoneId }).(pulumi.StringOutput)
}

type GetInstancesInstanceArrayOutput struct{ *pulumi.OutputState }

func (GetInstancesInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstancesInstance)(nil)).Elem()
}

func (o GetInstancesInstanceArrayOutput) ToGetInstancesInstanceArrayOutput() GetInstancesInstanceArrayOutput {
	return o
}

func (o GetInstancesInstanceArrayOutput) ToGetInstancesInstanceArrayOutputWithContext(ctx context.Context) GetInstancesInstanceArrayOutput {
	return o
}

func (o GetInstancesInstanceArrayOutput) Index(i pulumi.IntInput) GetInstancesInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetInstancesInstance {
		return vs[0].([]GetInstancesInstance)[vs[1].(int)]
	}).(GetInstancesInstanceOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesInstanceOutput{})
	pulumi.RegisterOutputType(GetInstancesInstanceArrayOutput{})
}
