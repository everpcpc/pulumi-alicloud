// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package adb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type GetClustersCluster struct {
	// Billing method. Value options: `PostPaid` for Pay-As-You-Go and `PrePaid` for subscription.
	ChargeType string `pulumi:"chargeType"`
	// The CreateTime of the ADB cluster.
	CreateTime string `pulumi:"createTime"`
	// The DBNodeClass of the ADB cluster.
	DbNodeClass string `pulumi:"dbNodeClass"`
	// The DBNodeCount of the ADB cluster.
	DbNodeCount int `pulumi:"dbNodeCount"`
	// The DBNodeStorage of the ADB cluster.
	DbNodeStorage int `pulumi:"dbNodeStorage"`
	// The description of the ADB cluster.
	Description string `pulumi:"description"`
	// Expiration time. Pay-As-You-Go clusters never expire.
	ExpireTime string `pulumi:"expireTime"`
	// The expired of the ADB cluster.
	Expired string `pulumi:"expired"`
	// The ID of the ADB cluster.
	Id string `pulumi:"id"`
	// The LockMode of the ADB cluster.
	LockMode string `pulumi:"lockMode"`
	// The DBClusterNetworkType of the ADB cluster.
	NetworkType string `pulumi:"networkType"`
	// Region ID the cluster belongs to.
	RegionId string `pulumi:"regionId"`
	// status of the cluster.
	Status string `pulumi:"status"`
	// ID of the VPC the cluster belongs to.
	VpcId string `pulumi:"vpcId"`
	// The ZoneId of the ADB cluster.
	ZoneId string `pulumi:"zoneId"`
}

// GetClustersClusterInput is an input type that accepts GetClustersClusterArgs and GetClustersClusterOutput values.
// You can construct a concrete instance of `GetClustersClusterInput` via:
//
// 		 GetClustersClusterArgs{...}
//
type GetClustersClusterInput interface {
	pulumi.Input

	ToGetClustersClusterOutput() GetClustersClusterOutput
	ToGetClustersClusterOutputWithContext(context.Context) GetClustersClusterOutput
}

type GetClustersClusterArgs struct {
	// Billing method. Value options: `PostPaid` for Pay-As-You-Go and `PrePaid` for subscription.
	ChargeType pulumi.StringInput `pulumi:"chargeType"`
	// The CreateTime of the ADB cluster.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// The DBNodeClass of the ADB cluster.
	DbNodeClass pulumi.StringInput `pulumi:"dbNodeClass"`
	// The DBNodeCount of the ADB cluster.
	DbNodeCount pulumi.IntInput `pulumi:"dbNodeCount"`
	// The DBNodeStorage of the ADB cluster.
	DbNodeStorage pulumi.IntInput `pulumi:"dbNodeStorage"`
	// The description of the ADB cluster.
	Description pulumi.StringInput `pulumi:"description"`
	// Expiration time. Pay-As-You-Go clusters never expire.
	ExpireTime pulumi.StringInput `pulumi:"expireTime"`
	// The expired of the ADB cluster.
	Expired pulumi.StringInput `pulumi:"expired"`
	// The ID of the ADB cluster.
	Id pulumi.StringInput `pulumi:"id"`
	// The LockMode of the ADB cluster.
	LockMode pulumi.StringInput `pulumi:"lockMode"`
	// The DBClusterNetworkType of the ADB cluster.
	NetworkType pulumi.StringInput `pulumi:"networkType"`
	// Region ID the cluster belongs to.
	RegionId pulumi.StringInput `pulumi:"regionId"`
	// status of the cluster.
	Status pulumi.StringInput `pulumi:"status"`
	// ID of the VPC the cluster belongs to.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
	// The ZoneId of the ADB cluster.
	ZoneId pulumi.StringInput `pulumi:"zoneId"`
}

func (GetClustersClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClustersCluster)(nil)).Elem()
}

func (i GetClustersClusterArgs) ToGetClustersClusterOutput() GetClustersClusterOutput {
	return i.ToGetClustersClusterOutputWithContext(context.Background())
}

func (i GetClustersClusterArgs) ToGetClustersClusterOutputWithContext(ctx context.Context) GetClustersClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClustersClusterOutput)
}

// GetClustersClusterArrayInput is an input type that accepts GetClustersClusterArray and GetClustersClusterArrayOutput values.
// You can construct a concrete instance of `GetClustersClusterArrayInput` via:
//
// 		 GetClustersClusterArray{ GetClustersClusterArgs{...} }
//
type GetClustersClusterArrayInput interface {
	pulumi.Input

	ToGetClustersClusterArrayOutput() GetClustersClusterArrayOutput
	ToGetClustersClusterArrayOutputWithContext(context.Context) GetClustersClusterArrayOutput
}

type GetClustersClusterArray []GetClustersClusterInput

func (GetClustersClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClustersCluster)(nil)).Elem()
}

func (i GetClustersClusterArray) ToGetClustersClusterArrayOutput() GetClustersClusterArrayOutput {
	return i.ToGetClustersClusterArrayOutputWithContext(context.Background())
}

func (i GetClustersClusterArray) ToGetClustersClusterArrayOutputWithContext(ctx context.Context) GetClustersClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClustersClusterArrayOutput)
}

type GetClustersClusterOutput struct{ *pulumi.OutputState }

func (GetClustersClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClustersCluster)(nil)).Elem()
}

func (o GetClustersClusterOutput) ToGetClustersClusterOutput() GetClustersClusterOutput {
	return o
}

func (o GetClustersClusterOutput) ToGetClustersClusterOutputWithContext(ctx context.Context) GetClustersClusterOutput {
	return o
}

// Billing method. Value options: `PostPaid` for Pay-As-You-Go and `PrePaid` for subscription.
func (o GetClustersClusterOutput) ChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.ChargeType }).(pulumi.StringOutput)
}

// The CreateTime of the ADB cluster.
func (o GetClustersClusterOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The DBNodeClass of the ADB cluster.
func (o GetClustersClusterOutput) DbNodeClass() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.DbNodeClass }).(pulumi.StringOutput)
}

// The DBNodeCount of the ADB cluster.
func (o GetClustersClusterOutput) DbNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetClustersCluster) int { return v.DbNodeCount }).(pulumi.IntOutput)
}

// The DBNodeStorage of the ADB cluster.
func (o GetClustersClusterOutput) DbNodeStorage() pulumi.IntOutput {
	return o.ApplyT(func(v GetClustersCluster) int { return v.DbNodeStorage }).(pulumi.IntOutput)
}

// The description of the ADB cluster.
func (o GetClustersClusterOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.Description }).(pulumi.StringOutput)
}

// Expiration time. Pay-As-You-Go clusters never expire.
func (o GetClustersClusterOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.ExpireTime }).(pulumi.StringOutput)
}

// The expired of the ADB cluster.
func (o GetClustersClusterOutput) Expired() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.Expired }).(pulumi.StringOutput)
}

// The ID of the ADB cluster.
func (o GetClustersClusterOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.Id }).(pulumi.StringOutput)
}

// The LockMode of the ADB cluster.
func (o GetClustersClusterOutput) LockMode() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.LockMode }).(pulumi.StringOutput)
}

// The DBClusterNetworkType of the ADB cluster.
func (o GetClustersClusterOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.NetworkType }).(pulumi.StringOutput)
}

// Region ID the cluster belongs to.
func (o GetClustersClusterOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.RegionId }).(pulumi.StringOutput)
}

// status of the cluster.
func (o GetClustersClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.Status }).(pulumi.StringOutput)
}

// ID of the VPC the cluster belongs to.
func (o GetClustersClusterOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.VpcId }).(pulumi.StringOutput)
}

// The ZoneId of the ADB cluster.
func (o GetClustersClusterOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.ZoneId }).(pulumi.StringOutput)
}

type GetClustersClusterArrayOutput struct{ *pulumi.OutputState }

func (GetClustersClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClustersCluster)(nil)).Elem()
}

func (o GetClustersClusterArrayOutput) ToGetClustersClusterArrayOutput() GetClustersClusterArrayOutput {
	return o
}

func (o GetClustersClusterArrayOutput) ToGetClustersClusterArrayOutputWithContext(ctx context.Context) GetClustersClusterArrayOutput {
	return o
}

func (o GetClustersClusterArrayOutput) Index(i pulumi.IntInput) GetClustersClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetClustersCluster {
		return vs[0].([]GetClustersCluster)[vs[1].(int)]
	}).(GetClustersClusterOutput)
}

type GetZonesZone struct {
	// ID of the zone.
	Id string `pulumi:"id"`
	// A list of zone ids in which the multi zone.
	MultiZoneIds []string `pulumi:"multiZoneIds"`
}

// GetZonesZoneInput is an input type that accepts GetZonesZoneArgs and GetZonesZoneOutput values.
// You can construct a concrete instance of `GetZonesZoneInput` via:
//
// 		 GetZonesZoneArgs{...}
//
type GetZonesZoneInput interface {
	pulumi.Input

	ToGetZonesZoneOutput() GetZonesZoneOutput
	ToGetZonesZoneOutputWithContext(context.Context) GetZonesZoneOutput
}

type GetZonesZoneArgs struct {
	// ID of the zone.
	Id pulumi.StringInput `pulumi:"id"`
	// A list of zone ids in which the multi zone.
	MultiZoneIds pulumi.StringArrayInput `pulumi:"multiZoneIds"`
}

func (GetZonesZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZonesZone)(nil)).Elem()
}

func (i GetZonesZoneArgs) ToGetZonesZoneOutput() GetZonesZoneOutput {
	return i.ToGetZonesZoneOutputWithContext(context.Background())
}

func (i GetZonesZoneArgs) ToGetZonesZoneOutputWithContext(ctx context.Context) GetZonesZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetZonesZoneOutput)
}

// GetZonesZoneArrayInput is an input type that accepts GetZonesZoneArray and GetZonesZoneArrayOutput values.
// You can construct a concrete instance of `GetZonesZoneArrayInput` via:
//
// 		 GetZonesZoneArray{ GetZonesZoneArgs{...} }
//
type GetZonesZoneArrayInput interface {
	pulumi.Input

	ToGetZonesZoneArrayOutput() GetZonesZoneArrayOutput
	ToGetZonesZoneArrayOutputWithContext(context.Context) GetZonesZoneArrayOutput
}

type GetZonesZoneArray []GetZonesZoneInput

func (GetZonesZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetZonesZone)(nil)).Elem()
}

func (i GetZonesZoneArray) ToGetZonesZoneArrayOutput() GetZonesZoneArrayOutput {
	return i.ToGetZonesZoneArrayOutputWithContext(context.Background())
}

func (i GetZonesZoneArray) ToGetZonesZoneArrayOutputWithContext(ctx context.Context) GetZonesZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetZonesZoneArrayOutput)
}

type GetZonesZoneOutput struct{ *pulumi.OutputState }

func (GetZonesZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZonesZone)(nil)).Elem()
}

func (o GetZonesZoneOutput) ToGetZonesZoneOutput() GetZonesZoneOutput {
	return o
}

func (o GetZonesZoneOutput) ToGetZonesZoneOutputWithContext(ctx context.Context) GetZonesZoneOutput {
	return o
}

// ID of the zone.
func (o GetZonesZoneOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetZonesZone) string { return v.Id }).(pulumi.StringOutput)
}

// A list of zone ids in which the multi zone.
func (o GetZonesZoneOutput) MultiZoneIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetZonesZone) []string { return v.MultiZoneIds }).(pulumi.StringArrayOutput)
}

type GetZonesZoneArrayOutput struct{ *pulumi.OutputState }

func (GetZonesZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetZonesZone)(nil)).Elem()
}

func (o GetZonesZoneArrayOutput) ToGetZonesZoneArrayOutput() GetZonesZoneArrayOutput {
	return o
}

func (o GetZonesZoneArrayOutput) ToGetZonesZoneArrayOutputWithContext(ctx context.Context) GetZonesZoneArrayOutput {
	return o
}

func (o GetZonesZoneArrayOutput) Index(i pulumi.IntInput) GetZonesZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetZonesZone {
		return vs[0].([]GetZonesZone)[vs[1].(int)]
	}).(GetZonesZoneOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClustersClusterOutput{})
	pulumi.RegisterOutputType(GetClustersClusterArrayOutput{})
	pulumi.RegisterOutputType(GetZonesZoneOutput{})
	pulumi.RegisterOutputType(GetZonesZoneArrayOutput{})
}
