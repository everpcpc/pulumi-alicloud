// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SiteMonitorIspCity struct {
	City string `pulumi:"city"`
	Isp  string `pulumi:"isp"`
}

// SiteMonitorIspCityInput is an input type that accepts SiteMonitorIspCityArgs and SiteMonitorIspCityOutput values.
// You can construct a concrete instance of `SiteMonitorIspCityInput` via:
//
// 		 SiteMonitorIspCityArgs{...}
//
type SiteMonitorIspCityInput interface {
	pulumi.Input

	ToSiteMonitorIspCityOutput() SiteMonitorIspCityOutput
	ToSiteMonitorIspCityOutputWithContext(context.Context) SiteMonitorIspCityOutput
}

type SiteMonitorIspCityArgs struct {
	City pulumi.StringInput `pulumi:"city"`
	Isp  pulumi.StringInput `pulumi:"isp"`
}

func (SiteMonitorIspCityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteMonitorIspCity)(nil)).Elem()
}

func (i SiteMonitorIspCityArgs) ToSiteMonitorIspCityOutput() SiteMonitorIspCityOutput {
	return i.ToSiteMonitorIspCityOutputWithContext(context.Background())
}

func (i SiteMonitorIspCityArgs) ToSiteMonitorIspCityOutputWithContext(ctx context.Context) SiteMonitorIspCityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMonitorIspCityOutput)
}

// SiteMonitorIspCityArrayInput is an input type that accepts SiteMonitorIspCityArray and SiteMonitorIspCityArrayOutput values.
// You can construct a concrete instance of `SiteMonitorIspCityArrayInput` via:
//
// 		 SiteMonitorIspCityArray{ SiteMonitorIspCityArgs{...} }
//
type SiteMonitorIspCityArrayInput interface {
	pulumi.Input

	ToSiteMonitorIspCityArrayOutput() SiteMonitorIspCityArrayOutput
	ToSiteMonitorIspCityArrayOutputWithContext(context.Context) SiteMonitorIspCityArrayOutput
}

type SiteMonitorIspCityArray []SiteMonitorIspCityInput

func (SiteMonitorIspCityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SiteMonitorIspCity)(nil)).Elem()
}

func (i SiteMonitorIspCityArray) ToSiteMonitorIspCityArrayOutput() SiteMonitorIspCityArrayOutput {
	return i.ToSiteMonitorIspCityArrayOutputWithContext(context.Background())
}

func (i SiteMonitorIspCityArray) ToSiteMonitorIspCityArrayOutputWithContext(ctx context.Context) SiteMonitorIspCityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMonitorIspCityArrayOutput)
}

type SiteMonitorIspCityOutput struct{ *pulumi.OutputState }

func (SiteMonitorIspCityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteMonitorIspCity)(nil)).Elem()
}

func (o SiteMonitorIspCityOutput) ToSiteMonitorIspCityOutput() SiteMonitorIspCityOutput {
	return o
}

func (o SiteMonitorIspCityOutput) ToSiteMonitorIspCityOutputWithContext(ctx context.Context) SiteMonitorIspCityOutput {
	return o
}

func (o SiteMonitorIspCityOutput) City() pulumi.StringOutput {
	return o.ApplyT(func(v SiteMonitorIspCity) string { return v.City }).(pulumi.StringOutput)
}

func (o SiteMonitorIspCityOutput) Isp() pulumi.StringOutput {
	return o.ApplyT(func(v SiteMonitorIspCity) string { return v.Isp }).(pulumi.StringOutput)
}

type SiteMonitorIspCityArrayOutput struct{ *pulumi.OutputState }

func (SiteMonitorIspCityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SiteMonitorIspCity)(nil)).Elem()
}

func (o SiteMonitorIspCityArrayOutput) ToSiteMonitorIspCityArrayOutput() SiteMonitorIspCityArrayOutput {
	return o
}

func (o SiteMonitorIspCityArrayOutput) ToSiteMonitorIspCityArrayOutputWithContext(ctx context.Context) SiteMonitorIspCityArrayOutput {
	return o
}

func (o SiteMonitorIspCityArrayOutput) Index(i pulumi.IntInput) SiteMonitorIspCityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SiteMonitorIspCity {
		return vs[0].([]SiteMonitorIspCity)[vs[1].(int)]
	}).(SiteMonitorIspCityOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteMonitorIspCityOutput{})
	pulumi.RegisterOutputType(SiteMonitorIspCityArrayOutput{})
}
