// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AssumeRole struct {
	Policy            *string `pulumi:"policy"`
	RoleArn           string  `pulumi:"roleArn"`
	SessionExpiration *int    `pulumi:"sessionExpiration"`
	SessionName       *string `pulumi:"sessionName"`
}

// AssumeRoleInput is an input type that accepts AssumeRoleArgs and AssumeRoleOutput values.
// You can construct a concrete instance of `AssumeRoleInput` via:
//
//          AssumeRoleArgs{...}
type AssumeRoleInput interface {
	pulumi.Input

	ToAssumeRoleOutput() AssumeRoleOutput
	ToAssumeRoleOutputWithContext(context.Context) AssumeRoleOutput
}

type AssumeRoleArgs struct {
	Policy            pulumi.StringPtrInput `pulumi:"policy"`
	RoleArn           pulumi.StringInput    `pulumi:"roleArn"`
	SessionExpiration pulumi.IntPtrInput    `pulumi:"sessionExpiration"`
	SessionName       pulumi.StringPtrInput `pulumi:"sessionName"`
}

func (AssumeRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssumeRole)(nil)).Elem()
}

func (i AssumeRoleArgs) ToAssumeRoleOutput() AssumeRoleOutput {
	return i.ToAssumeRoleOutputWithContext(context.Background())
}

func (i AssumeRoleArgs) ToAssumeRoleOutputWithContext(ctx context.Context) AssumeRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssumeRoleOutput)
}

type AssumeRoleOutput struct{ *pulumi.OutputState }

func (AssumeRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssumeRole)(nil)).Elem()
}

func (o AssumeRoleOutput) ToAssumeRoleOutput() AssumeRoleOutput {
	return o
}

func (o AssumeRoleOutput) ToAssumeRoleOutputWithContext(ctx context.Context) AssumeRoleOutput {
	return o
}

func (o AssumeRoleOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssumeRole) *string { return v.Policy }).(pulumi.StringPtrOutput)
}

func (o AssumeRoleOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v AssumeRole) string { return v.RoleArn }).(pulumi.StringOutput)
}

func (o AssumeRoleOutput) SessionExpiration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AssumeRole) *int { return v.SessionExpiration }).(pulumi.IntPtrOutput)
}

func (o AssumeRoleOutput) SessionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssumeRole) *string { return v.SessionName }).(pulumi.StringPtrOutput)
}

type Endpoints struct {
	Actiontrail     *string `pulumi:"actiontrail"`
	Adb             *string `pulumi:"adb"`
	Alidns          *string `pulumi:"alidns"`
	Alikafka        *string `pulumi:"alikafka"`
	Apigateway      *string `pulumi:"apigateway"`
	BrainIndustrial *string `pulumi:"brainIndustrial"`
	Bssopenapi      *string `pulumi:"bssopenapi"`
	Cas             *string `pulumi:"cas"`
	Cassandra       *string `pulumi:"cassandra"`
	Cbn             *string `pulumi:"cbn"`
	Cdn             *string `pulumi:"cdn"`
	Cen             *string `pulumi:"cen"`
	Cms             *string `pulumi:"cms"`
	Config          *string `pulumi:"config"`
	Cr              *string `pulumi:"cr"`
	Cs              *string `pulumi:"cs"`
	Datahub         *string `pulumi:"datahub"`
	Dcdn            *string `pulumi:"dcdn"`
	Ddosbgp         *string `pulumi:"ddosbgp"`
	Ddoscoo         *string `pulumi:"ddoscoo"`
	Dds             *string `pulumi:"dds"`
	DmsEnterprise   *string `pulumi:"dmsEnterprise"`
	Dns             *string `pulumi:"dns"`
	Drds            *string `pulumi:"drds"`
	Eci             *string `pulumi:"eci"`
	Ecs             *string `pulumi:"ecs"`
	Eipanycast      *string `pulumi:"eipanycast"`
	Elasticsearch   *string `pulumi:"elasticsearch"`
	Emr             *string `pulumi:"emr"`
	Ess             *string `pulumi:"ess"`
	Fc              *string `pulumi:"fc"`
	Fnf             *string `pulumi:"fnf"`
	Ga              *string `pulumi:"ga"`
	Gpdb            *string `pulumi:"gpdb"`
	Hitsdb          *string `pulumi:"hitsdb"`
	Ims             *string `pulumi:"ims"`
	Kms             *string `pulumi:"kms"`
	Kvstore         *string `pulumi:"kvstore"`
	Location        *string `pulumi:"location"`
	Log             *string `pulumi:"log"`
	Market          *string `pulumi:"market"`
	Maxcompute      *string `pulumi:"maxcompute"`
	Mns             *string `pulumi:"mns"`
	Mse             *string `pulumi:"mse"`
	Nas             *string `pulumi:"nas"`
	Ons             *string `pulumi:"ons"`
	Oos             *string `pulumi:"oos"`
	Oss             *string `pulumi:"oss"`
	Ots             *string `pulumi:"ots"`
	Polardb         *string `pulumi:"polardb"`
	Privatelink     *string `pulumi:"privatelink"`
	Pvtz            *string `pulumi:"pvtz"`
	Quotas          *string `pulumi:"quotas"`
	RKvstore        *string `pulumi:"rKvstore"`
	Ram             *string `pulumi:"ram"`
	Rds             *string `pulumi:"rds"`
	Resourcemanager *string `pulumi:"resourcemanager"`
	Resourcesharing *string `pulumi:"resourcesharing"`
	Ros             *string `pulumi:"ros"`
	Slb             *string `pulumi:"slb"`
	Sts             *string `pulumi:"sts"`
	Vpc             *string `pulumi:"vpc"`
	WafOpenapi      *string `pulumi:"wafOpenapi"`
}

// EndpointsInput is an input type that accepts EndpointsArgs and EndpointsOutput values.
// You can construct a concrete instance of `EndpointsInput` via:
//
//          EndpointsArgs{...}
type EndpointsInput interface {
	pulumi.Input

	ToEndpointsOutput() EndpointsOutput
	ToEndpointsOutputWithContext(context.Context) EndpointsOutput
}

type EndpointsArgs struct {
	Actiontrail     pulumi.StringPtrInput `pulumi:"actiontrail"`
	Adb             pulumi.StringPtrInput `pulumi:"adb"`
	Alidns          pulumi.StringPtrInput `pulumi:"alidns"`
	Alikafka        pulumi.StringPtrInput `pulumi:"alikafka"`
	Apigateway      pulumi.StringPtrInput `pulumi:"apigateway"`
	BrainIndustrial pulumi.StringPtrInput `pulumi:"brainIndustrial"`
	Bssopenapi      pulumi.StringPtrInput `pulumi:"bssopenapi"`
	Cas             pulumi.StringPtrInput `pulumi:"cas"`
	Cassandra       pulumi.StringPtrInput `pulumi:"cassandra"`
	Cbn             pulumi.StringPtrInput `pulumi:"cbn"`
	Cdn             pulumi.StringPtrInput `pulumi:"cdn"`
	Cen             pulumi.StringPtrInput `pulumi:"cen"`
	Cms             pulumi.StringPtrInput `pulumi:"cms"`
	Config          pulumi.StringPtrInput `pulumi:"config"`
	Cr              pulumi.StringPtrInput `pulumi:"cr"`
	Cs              pulumi.StringPtrInput `pulumi:"cs"`
	Datahub         pulumi.StringPtrInput `pulumi:"datahub"`
	Dcdn            pulumi.StringPtrInput `pulumi:"dcdn"`
	Ddosbgp         pulumi.StringPtrInput `pulumi:"ddosbgp"`
	Ddoscoo         pulumi.StringPtrInput `pulumi:"ddoscoo"`
	Dds             pulumi.StringPtrInput `pulumi:"dds"`
	DmsEnterprise   pulumi.StringPtrInput `pulumi:"dmsEnterprise"`
	Dns             pulumi.StringPtrInput `pulumi:"dns"`
	Drds            pulumi.StringPtrInput `pulumi:"drds"`
	Eci             pulumi.StringPtrInput `pulumi:"eci"`
	Ecs             pulumi.StringPtrInput `pulumi:"ecs"`
	Eipanycast      pulumi.StringPtrInput `pulumi:"eipanycast"`
	Elasticsearch   pulumi.StringPtrInput `pulumi:"elasticsearch"`
	Emr             pulumi.StringPtrInput `pulumi:"emr"`
	Ess             pulumi.StringPtrInput `pulumi:"ess"`
	Fc              pulumi.StringPtrInput `pulumi:"fc"`
	Fnf             pulumi.StringPtrInput `pulumi:"fnf"`
	Ga              pulumi.StringPtrInput `pulumi:"ga"`
	Gpdb            pulumi.StringPtrInput `pulumi:"gpdb"`
	Hitsdb          pulumi.StringPtrInput `pulumi:"hitsdb"`
	Ims             pulumi.StringPtrInput `pulumi:"ims"`
	Kms             pulumi.StringPtrInput `pulumi:"kms"`
	Kvstore         pulumi.StringPtrInput `pulumi:"kvstore"`
	Location        pulumi.StringPtrInput `pulumi:"location"`
	Log             pulumi.StringPtrInput `pulumi:"log"`
	Market          pulumi.StringPtrInput `pulumi:"market"`
	Maxcompute      pulumi.StringPtrInput `pulumi:"maxcompute"`
	Mns             pulumi.StringPtrInput `pulumi:"mns"`
	Mse             pulumi.StringPtrInput `pulumi:"mse"`
	Nas             pulumi.StringPtrInput `pulumi:"nas"`
	Ons             pulumi.StringPtrInput `pulumi:"ons"`
	Oos             pulumi.StringPtrInput `pulumi:"oos"`
	Oss             pulumi.StringPtrInput `pulumi:"oss"`
	Ots             pulumi.StringPtrInput `pulumi:"ots"`
	Polardb         pulumi.StringPtrInput `pulumi:"polardb"`
	Privatelink     pulumi.StringPtrInput `pulumi:"privatelink"`
	Pvtz            pulumi.StringPtrInput `pulumi:"pvtz"`
	Quotas          pulumi.StringPtrInput `pulumi:"quotas"`
	RKvstore        pulumi.StringPtrInput `pulumi:"rKvstore"`
	Ram             pulumi.StringPtrInput `pulumi:"ram"`
	Rds             pulumi.StringPtrInput `pulumi:"rds"`
	Resourcemanager pulumi.StringPtrInput `pulumi:"resourcemanager"`
	Resourcesharing pulumi.StringPtrInput `pulumi:"resourcesharing"`
	Ros             pulumi.StringPtrInput `pulumi:"ros"`
	Slb             pulumi.StringPtrInput `pulumi:"slb"`
	Sts             pulumi.StringPtrInput `pulumi:"sts"`
	Vpc             pulumi.StringPtrInput `pulumi:"vpc"`
	WafOpenapi      pulumi.StringPtrInput `pulumi:"wafOpenapi"`
}

func (EndpointsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Endpoints)(nil)).Elem()
}

func (i EndpointsArgs) ToEndpointsOutput() EndpointsOutput {
	return i.ToEndpointsOutputWithContext(context.Background())
}

func (i EndpointsArgs) ToEndpointsOutputWithContext(ctx context.Context) EndpointsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsOutput)
}

// EndpointsArrayInput is an input type that accepts EndpointsArray and EndpointsArrayOutput values.
// You can construct a concrete instance of `EndpointsArrayInput` via:
//
//          EndpointsArray{ EndpointsArgs{...} }
type EndpointsArrayInput interface {
	pulumi.Input

	ToEndpointsArrayOutput() EndpointsArrayOutput
	ToEndpointsArrayOutputWithContext(context.Context) EndpointsArrayOutput
}

type EndpointsArray []EndpointsInput

func (EndpointsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Endpoints)(nil)).Elem()
}

func (i EndpointsArray) ToEndpointsArrayOutput() EndpointsArrayOutput {
	return i.ToEndpointsArrayOutputWithContext(context.Background())
}

func (i EndpointsArray) ToEndpointsArrayOutputWithContext(ctx context.Context) EndpointsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsArrayOutput)
}

type EndpointsOutput struct{ *pulumi.OutputState }

func (EndpointsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Endpoints)(nil)).Elem()
}

func (o EndpointsOutput) ToEndpointsOutput() EndpointsOutput {
	return o
}

func (o EndpointsOutput) ToEndpointsOutputWithContext(ctx context.Context) EndpointsOutput {
	return o
}

func (o EndpointsOutput) Actiontrail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Actiontrail }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Adb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Adb }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Alidns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Alidns }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Alikafka() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Alikafka }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Apigateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Apigateway }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) BrainIndustrial() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.BrainIndustrial }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Bssopenapi() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Bssopenapi }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cas() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cas }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cassandra() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cassandra }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cbn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cbn }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cdn }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cen() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cen }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cms }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Config }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cr }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cs }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Datahub() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Datahub }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Dcdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Dcdn }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ddosbgp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ddosbgp }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ddoscoo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ddoscoo }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Dds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Dds }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) DmsEnterprise() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.DmsEnterprise }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Dns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Dns }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Drds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Drds }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Eci() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Eci }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ecs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ecs }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Eipanycast() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Eipanycast }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Elasticsearch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Elasticsearch }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Emr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Emr }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ess }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Fc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Fc }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Fnf() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Fnf }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ga() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ga }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Gpdb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Gpdb }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Hitsdb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Hitsdb }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ims() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ims }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Kms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Kms }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Kvstore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Kvstore }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Log() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Log }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Market() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Market }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Maxcompute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Maxcompute }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Mns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Mns }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Mse() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Mse }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Nas() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Nas }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ons() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ons }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Oos() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Oos }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Oss() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Oss }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ots() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ots }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Polardb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Polardb }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Privatelink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Privatelink }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Pvtz() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Pvtz }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Quotas() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Quotas }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) RKvstore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.RKvstore }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ram() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ram }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Rds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Rds }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Resourcemanager() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Resourcemanager }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Resourcesharing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Resourcesharing }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ros() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ros }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Slb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Slb }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Sts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Sts }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Vpc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Vpc }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) WafOpenapi() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.WafOpenapi }).(pulumi.StringPtrOutput)
}

type EndpointsArrayOutput struct{ *pulumi.OutputState }

func (EndpointsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Endpoints)(nil)).Elem()
}

func (o EndpointsArrayOutput) ToEndpointsArrayOutput() EndpointsArrayOutput {
	return o
}

func (o EndpointsArrayOutput) ToEndpointsArrayOutputWithContext(ctx context.Context) EndpointsArrayOutput {
	return o
}

func (o EndpointsArrayOutput) Index(i pulumi.IntInput) EndpointsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Endpoints {
		return vs[0].([]Endpoints)[vs[1].(int)]
	}).(EndpointsOutput)
}

func init() {
	pulumi.RegisterOutputType(AssumeRoleOutput{})
	pulumi.RegisterOutputType(EndpointsOutput{})
	pulumi.RegisterOutputType(EndpointsArrayOutput{})
}
