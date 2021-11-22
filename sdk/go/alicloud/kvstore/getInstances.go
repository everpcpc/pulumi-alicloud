// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kvstore

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `kvstore.getInstances` data source provides a collection of kvstore instances available in Alicloud account.
// Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/kvstore"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "testname"
// 		_default, err := kvstore.GetInstances(ctx, &kvstore.GetInstancesArgs{
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstInstanceName", _default.Instances[0].Name)
// 		return nil
// 	})
// }
// ```
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("alicloud:kvstore/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// The type of the architecture. Valid values: `cluster`, `standard` and `SplitRW`.
	ArchitectureType *string `pulumi:"architectureType"`
	// Used to retrieve instances belong to specified `vswitch` resources.  Valid values: `Enterprise`, `Community`.
	EditionType *string `pulumi:"editionType"`
	// Default to `false`. Set it to true can output more details.
	EnableDetails *bool `pulumi:"enableDetails"`
	// The engine version. Valid values: `2.8`, `4.0`, `5.0`, `6.0`.
	EngineVersion *string `pulumi:"engineVersion"`
	// The expiration status of the instance.
	Expired *string `pulumi:"expired"`
	// Whether to create a distributed cache.
	GlobalInstance *bool `pulumi:"globalInstance"`
	// A list of KVStore DBInstance IDs.
	Ids []string `pulumi:"ids"`
	// Type of the applied ApsaraDB for Redis instance. For more information, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/61135.htm).
	InstanceClass *string `pulumi:"instanceClass"`
	// The engine type of the KVStore DBInstance. Options are `Memcache`, and `Redis`. If no value is specified, all types are returned.
	InstanceType *string `pulumi:"instanceType"`
	// A regex string to apply to the instance name.
	NameRegex *string `pulumi:"nameRegex"`
	// The type of the network. Valid values: `CLASSIC`, `VPC`.
	NetworkType *string `pulumi:"networkType"`
	OutputFile  *string `pulumi:"outputFile"`
	// The payment type. Valid values: `PostPaid`, `PrePaid`.
	PaymentType *string `pulumi:"paymentType"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The name of the instance.
	SearchKey *string `pulumi:"searchKey"`
	// The status of the KVStore DBInstance. Valid values: `Changing`, `CleaningUpExpiredData`, `Creating`, `Flushing`, `HASwitching`, `Inactive`, `MajorVersionUpgrading`, `Migrating`, `NetworkModifying`, `Normal`, `Rebooting`, `SSLModifying`, `Transforming`, `ZoneMigrating`.
	Status *string `pulumi:"status"`
	// Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{"key1":"value1"}`.
	Tags map[string]interface{} `pulumi:"tags"`
	// Used to retrieve instances belong to specified VPC.
	VpcId *string `pulumi:"vpcId"`
	// Used to retrieve instances belong to specified `vswitch` resources.
	VswitchId *string `pulumi:"vswitchId"`
	// The ID of the zone.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	ArchitectureType *string `pulumi:"architectureType"`
	EditionType      *string `pulumi:"editionType"`
	EnableDetails    *bool   `pulumi:"enableDetails"`
	// The engine version of the instance.
	EngineVersion  *string `pulumi:"engineVersion"`
	Expired        *string `pulumi:"expired"`
	GlobalInstance *bool   `pulumi:"globalInstance"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of KVStore Instance IDs.
	Ids           []string `pulumi:"ids"`
	InstanceClass *string  `pulumi:"instanceClass"`
	// (Optional) Database type. Valid Values: `Memcache`, `Redis`. If no value is specified, all types are returned.
	// * `instanceClass`- (Optional) Type of the applied ApsaraDB for instance.
	//   For more information, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/61135.htm).
	InstanceType *string `pulumi:"instanceType"`
	// A list of KVStore Instances. Its every element contains the following attributes:
	Instances []GetInstancesInstance `pulumi:"instances"`
	NameRegex *string                `pulumi:"nameRegex"`
	// A list of KVStore Instance names.
	Names []string `pulumi:"names"`
	// The network type of the instance.
	NetworkType *string `pulumi:"networkType"`
	OutputFile  *string `pulumi:"outputFile"`
	// Billing method. Valid Values: `PostPaid` for  Pay-As-You-Go and `PrePaid` for subscription.
	PaymentType     *string `pulumi:"paymentType"`
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	SearchKey       *string `pulumi:"searchKey"`
	// Status of the instance.
	Status *string                `pulumi:"status"`
	Tags   map[string]interface{} `pulumi:"tags"`
	// VPC ID the instance belongs to.
	VpcId *string `pulumi:"vpcId"`
	// VSwitch ID the instance belongs to.
	VswitchId *string `pulumi:"vswitchId"`
	// The ID of zone.
	ZoneId *string `pulumi:"zoneId"`
}

func GetInstancesOutput(ctx *pulumi.Context, args GetInstancesOutputArgs, opts ...pulumi.InvokeOption) GetInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstancesResult, error) {
			args := v.(GetInstancesArgs)
			r, err := GetInstances(ctx, &args, opts...)
			return *r, err
		}).(GetInstancesResultOutput)
}

// A collection of arguments for invoking getInstances.
type GetInstancesOutputArgs struct {
	// The type of the architecture. Valid values: `cluster`, `standard` and `SplitRW`.
	ArchitectureType pulumi.StringPtrInput `pulumi:"architectureType"`
	// Used to retrieve instances belong to specified `vswitch` resources.  Valid values: `Enterprise`, `Community`.
	EditionType pulumi.StringPtrInput `pulumi:"editionType"`
	// Default to `false`. Set it to true can output more details.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// The engine version. Valid values: `2.8`, `4.0`, `5.0`, `6.0`.
	EngineVersion pulumi.StringPtrInput `pulumi:"engineVersion"`
	// The expiration status of the instance.
	Expired pulumi.StringPtrInput `pulumi:"expired"`
	// Whether to create a distributed cache.
	GlobalInstance pulumi.BoolPtrInput `pulumi:"globalInstance"`
	// A list of KVStore DBInstance IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// Type of the applied ApsaraDB for Redis instance. For more information, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/61135.htm).
	InstanceClass pulumi.StringPtrInput `pulumi:"instanceClass"`
	// The engine type of the KVStore DBInstance. Options are `Memcache`, and `Redis`. If no value is specified, all types are returned.
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
	// A regex string to apply to the instance name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The type of the network. Valid values: `CLASSIC`, `VPC`.
	NetworkType pulumi.StringPtrInput `pulumi:"networkType"`
	OutputFile  pulumi.StringPtrInput `pulumi:"outputFile"`
	// The payment type. Valid values: `PostPaid`, `PrePaid`.
	PaymentType pulumi.StringPtrInput `pulumi:"paymentType"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The name of the instance.
	SearchKey pulumi.StringPtrInput `pulumi:"searchKey"`
	// The status of the KVStore DBInstance. Valid values: `Changing`, `CleaningUpExpiredData`, `Creating`, `Flushing`, `HASwitching`, `Inactive`, `MajorVersionUpgrading`, `Migrating`, `NetworkModifying`, `Normal`, `Rebooting`, `SSLModifying`, `Transforming`, `ZoneMigrating`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{"key1":"value1"}`.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Used to retrieve instances belong to specified VPC.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
	// Used to retrieve instances belong to specified `vswitch` resources.
	VswitchId pulumi.StringPtrInput `pulumi:"vswitchId"`
	// The ID of the zone.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (GetInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getInstances.
type GetInstancesResultOutput struct{ *pulumi.OutputState }

func (GetInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesResult)(nil)).Elem()
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutput() GetInstancesResultOutput {
	return o
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutputWithContext(ctx context.Context) GetInstancesResultOutput {
	return o
}

func (o GetInstancesResultOutput) ArchitectureType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.ArchitectureType }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) EditionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.EditionType }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The engine version of the instance.
func (o GetInstancesResultOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) Expired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.Expired }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) GlobalInstance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *bool { return v.GlobalInstance }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of KVStore Instance IDs.
func (o GetInstancesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetInstancesResultOutput) InstanceClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.InstanceClass }).(pulumi.StringPtrOutput)
}

// (Optional) Database type. Valid Values: `Memcache`, `Redis`. If no value is specified, all types are returned.
// * `instanceClass`- (Optional) Type of the applied ApsaraDB for instance.
//   For more information, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/61135.htm).
func (o GetInstancesResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

// A list of KVStore Instances. Its every element contains the following attributes:
func (o GetInstancesResultOutput) Instances() GetInstancesInstanceArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesInstance { return v.Instances }).(GetInstancesInstanceArrayOutput)
}

func (o GetInstancesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of KVStore Instance names.
func (o GetInstancesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

// The network type of the instance.
func (o GetInstancesResultOutput) NetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.NetworkType }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// Billing method. Valid Values: `PostPaid` for  Pay-As-You-Go and `PrePaid` for subscription.
func (o GetInstancesResultOutput) PaymentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.PaymentType }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) SearchKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.SearchKey }).(pulumi.StringPtrOutput)
}

// Status of the instance.
func (o GetInstancesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetInstancesResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

// VPC ID the instance belongs to.
func (o GetInstancesResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

// VSwitch ID the instance belongs to.
func (o GetInstancesResultOutput) VswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.VswitchId }).(pulumi.StringPtrOutput)
}

// The ID of zone.
func (o GetInstancesResultOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.ZoneId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}
