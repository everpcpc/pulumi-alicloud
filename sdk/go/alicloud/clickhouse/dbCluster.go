// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package clickhouse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Click House DBCluster resource.
//
// For information about Click House DBCluster and how to use it, see [What is DBCluster](https://www.alibabacloud.com/product/clickhouse).
//
// > **NOTE:** Available in v1.134.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/clickhouse"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := clickhouse.NewDbCluster(ctx, "_default", &clickhouse.DbClusterArgs{
// 			Category: pulumi.String("Basic"),
// 			DbClusterAccessWhiteLists: clickhouse.DbClusterDbClusterAccessWhiteListArray{
// 				&clickhouse.DbClusterDbClusterAccessWhiteListArgs{
// 					DbClusterIpArrayAttribute: pulumi.String("test"),
// 					DbClusterIpArrayName:      pulumi.String("test"),
// 					SecurityIpList:            pulumi.String("192.168.0.1"),
// 				},
// 			},
// 			DbClusterClass:       pulumi.String("S8"),
// 			DbClusterNetworkType: pulumi.String("vpc"),
// 			DbClusterVersion:     pulumi.String("20.3.10.75"),
// 			DbNodeGroupCount:     pulumi.Int(1),
// 			DbNodeStorage:        pulumi.String("500"),
// 			PaymentType:          pulumi.String("PayAsYouGo"),
// 			StorageType:          pulumi.String("cloud_essd"),
// 			VswitchId:            pulumi.String("your_vswitch_id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Click House DBCluster can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:clickhouse/dbCluster:DbCluster example <id>
// ```
type DbCluster struct {
	pulumi.CustomResourceState

	// The Category of DBCluster. Valid values: `Basic`,`HighAvailability`.
	Category pulumi.StringOutput `pulumi:"category"`
	// The db cluster access white list.
	DbClusterAccessWhiteLists DbClusterDbClusterAccessWhiteListArrayOutput `pulumi:"dbClusterAccessWhiteLists"`
	// The DBCluster class. According to the category, dbClusterClass has two value ranges:
	// * Under the condition that the category is the `Basic`, Valid values: `S4-NEW`, `S8`, `S16`, `S32`, `S64`, `S104`.
	// * Under the condition that the category is the `HighAvailability`, Valid values: `C4-NEW`, `C8`, `C16`, `C32`, `C64`, `C104`.
	DbClusterClass pulumi.StringOutput `pulumi:"dbClusterClass"`
	// The DBCluster description.
	DbClusterDescription pulumi.StringOutput `pulumi:"dbClusterDescription"`
	// The DBCluster network type. Valid values: `vpc`.
	DbClusterNetworkType pulumi.StringOutput `pulumi:"dbClusterNetworkType"`
	// The DBCluster version. Valid values: `19.15.2.2`, `20.3.10.75`, `20.8.7.15`.
	DbClusterVersion pulumi.StringOutput `pulumi:"dbClusterVersion"`
	// The db node group count. The number should between 1 and 48.
	DbNodeGroupCount pulumi.IntOutput `pulumi:"dbNodeGroupCount"`
	// The db node storage.
	DbNodeStorage pulumi.StringOutput `pulumi:"dbNodeStorage"`
	// Key management service KMS key ID.
	EncryptionKey pulumi.StringPtrOutput `pulumi:"encryptionKey"`
	// Currently only supports ECS disk encryption, with a value of CloudDisk, not encrypted when empty.
	EncryptionType pulumi.StringPtrOutput `pulumi:"encryptionType"`
	// The maintenance window of DBCluster. Valid format: `hh:mmZ-hh:mm Z`.
	MaintainTime pulumi.StringOutput `pulumi:"maintainTime"`
	// The payment type of the resource. Valid values: `PayAsYouGo`,`Subscription`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// Pre-paid cluster of the pay-as-you-go cycle. Valid values: `Month`, `Year`.
	Period pulumi.StringPtrOutput `pulumi:"period"`
	// The status of the resource. Valid values: `Running`,`Creating`,`Deleting`,`Restarting`,`Preparing`,.
	Status pulumi.StringOutput `pulumi:"status"`
	// Storage type of DBCluster. Valid values: `cloudEssd`, `cloudEfficiency`, `cloudEssdPl2`, `cloudEssdPl3`.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
	// The used time of DBCluster.
	UsedTime pulumi.StringPtrOutput `pulumi:"usedTime"`
	// The vswitch id of DBCluster.
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
}

// NewDbCluster registers a new resource with the given unique name, arguments, and options.
func NewDbCluster(ctx *pulumi.Context,
	name string, args *DbClusterArgs, opts ...pulumi.ResourceOption) (*DbCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	if args.DbClusterClass == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterClass'")
	}
	if args.DbClusterNetworkType == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterNetworkType'")
	}
	if args.DbClusterVersion == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterVersion'")
	}
	if args.DbNodeGroupCount == nil {
		return nil, errors.New("invalid value for required argument 'DbNodeGroupCount'")
	}
	if args.DbNodeStorage == nil {
		return nil, errors.New("invalid value for required argument 'DbNodeStorage'")
	}
	if args.PaymentType == nil {
		return nil, errors.New("invalid value for required argument 'PaymentType'")
	}
	if args.StorageType == nil {
		return nil, errors.New("invalid value for required argument 'StorageType'")
	}
	var resource DbCluster
	err := ctx.RegisterResource("alicloud:clickhouse/dbCluster:DbCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDbCluster gets an existing DbCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDbCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DbClusterState, opts ...pulumi.ResourceOption) (*DbCluster, error) {
	var resource DbCluster
	err := ctx.ReadResource("alicloud:clickhouse/dbCluster:DbCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DbCluster resources.
type dbClusterState struct {
	// The Category of DBCluster. Valid values: `Basic`,`HighAvailability`.
	Category *string `pulumi:"category"`
	// The db cluster access white list.
	DbClusterAccessWhiteLists []DbClusterDbClusterAccessWhiteList `pulumi:"dbClusterAccessWhiteLists"`
	// The DBCluster class. According to the category, dbClusterClass has two value ranges:
	// * Under the condition that the category is the `Basic`, Valid values: `S4-NEW`, `S8`, `S16`, `S32`, `S64`, `S104`.
	// * Under the condition that the category is the `HighAvailability`, Valid values: `C4-NEW`, `C8`, `C16`, `C32`, `C64`, `C104`.
	DbClusterClass *string `pulumi:"dbClusterClass"`
	// The DBCluster description.
	DbClusterDescription *string `pulumi:"dbClusterDescription"`
	// The DBCluster network type. Valid values: `vpc`.
	DbClusterNetworkType *string `pulumi:"dbClusterNetworkType"`
	// The DBCluster version. Valid values: `19.15.2.2`, `20.3.10.75`, `20.8.7.15`.
	DbClusterVersion *string `pulumi:"dbClusterVersion"`
	// The db node group count. The number should between 1 and 48.
	DbNodeGroupCount *int `pulumi:"dbNodeGroupCount"`
	// The db node storage.
	DbNodeStorage *string `pulumi:"dbNodeStorage"`
	// Key management service KMS key ID.
	EncryptionKey *string `pulumi:"encryptionKey"`
	// Currently only supports ECS disk encryption, with a value of CloudDisk, not encrypted when empty.
	EncryptionType *string `pulumi:"encryptionType"`
	// The maintenance window of DBCluster. Valid format: `hh:mmZ-hh:mm Z`.
	MaintainTime *string `pulumi:"maintainTime"`
	// The payment type of the resource. Valid values: `PayAsYouGo`,`Subscription`.
	PaymentType *string `pulumi:"paymentType"`
	// Pre-paid cluster of the pay-as-you-go cycle. Valid values: `Month`, `Year`.
	Period *string `pulumi:"period"`
	// The status of the resource. Valid values: `Running`,`Creating`,`Deleting`,`Restarting`,`Preparing`,.
	Status *string `pulumi:"status"`
	// Storage type of DBCluster. Valid values: `cloudEssd`, `cloudEfficiency`, `cloudEssdPl2`, `cloudEssdPl3`.
	StorageType *string `pulumi:"storageType"`
	// The used time of DBCluster.
	UsedTime *string `pulumi:"usedTime"`
	// The vswitch id of DBCluster.
	VswitchId *string `pulumi:"vswitchId"`
}

type DbClusterState struct {
	// The Category of DBCluster. Valid values: `Basic`,`HighAvailability`.
	Category pulumi.StringPtrInput
	// The db cluster access white list.
	DbClusterAccessWhiteLists DbClusterDbClusterAccessWhiteListArrayInput
	// The DBCluster class. According to the category, dbClusterClass has two value ranges:
	// * Under the condition that the category is the `Basic`, Valid values: `S4-NEW`, `S8`, `S16`, `S32`, `S64`, `S104`.
	// * Under the condition that the category is the `HighAvailability`, Valid values: `C4-NEW`, `C8`, `C16`, `C32`, `C64`, `C104`.
	DbClusterClass pulumi.StringPtrInput
	// The DBCluster description.
	DbClusterDescription pulumi.StringPtrInput
	// The DBCluster network type. Valid values: `vpc`.
	DbClusterNetworkType pulumi.StringPtrInput
	// The DBCluster version. Valid values: `19.15.2.2`, `20.3.10.75`, `20.8.7.15`.
	DbClusterVersion pulumi.StringPtrInput
	// The db node group count. The number should between 1 and 48.
	DbNodeGroupCount pulumi.IntPtrInput
	// The db node storage.
	DbNodeStorage pulumi.StringPtrInput
	// Key management service KMS key ID.
	EncryptionKey pulumi.StringPtrInput
	// Currently only supports ECS disk encryption, with a value of CloudDisk, not encrypted when empty.
	EncryptionType pulumi.StringPtrInput
	// The maintenance window of DBCluster. Valid format: `hh:mmZ-hh:mm Z`.
	MaintainTime pulumi.StringPtrInput
	// The payment type of the resource. Valid values: `PayAsYouGo`,`Subscription`.
	PaymentType pulumi.StringPtrInput
	// Pre-paid cluster of the pay-as-you-go cycle. Valid values: `Month`, `Year`.
	Period pulumi.StringPtrInput
	// The status of the resource. Valid values: `Running`,`Creating`,`Deleting`,`Restarting`,`Preparing`,.
	Status pulumi.StringPtrInput
	// Storage type of DBCluster. Valid values: `cloudEssd`, `cloudEfficiency`, `cloudEssdPl2`, `cloudEssdPl3`.
	StorageType pulumi.StringPtrInput
	// The used time of DBCluster.
	UsedTime pulumi.StringPtrInput
	// The vswitch id of DBCluster.
	VswitchId pulumi.StringPtrInput
}

func (DbClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbClusterState)(nil)).Elem()
}

type dbClusterArgs struct {
	// The Category of DBCluster. Valid values: `Basic`,`HighAvailability`.
	Category string `pulumi:"category"`
	// The db cluster access white list.
	DbClusterAccessWhiteLists []DbClusterDbClusterAccessWhiteList `pulumi:"dbClusterAccessWhiteLists"`
	// The DBCluster class. According to the category, dbClusterClass has two value ranges:
	// * Under the condition that the category is the `Basic`, Valid values: `S4-NEW`, `S8`, `S16`, `S32`, `S64`, `S104`.
	// * Under the condition that the category is the `HighAvailability`, Valid values: `C4-NEW`, `C8`, `C16`, `C32`, `C64`, `C104`.
	DbClusterClass string `pulumi:"dbClusterClass"`
	// The DBCluster description.
	DbClusterDescription *string `pulumi:"dbClusterDescription"`
	// The DBCluster network type. Valid values: `vpc`.
	DbClusterNetworkType string `pulumi:"dbClusterNetworkType"`
	// The DBCluster version. Valid values: `19.15.2.2`, `20.3.10.75`, `20.8.7.15`.
	DbClusterVersion string `pulumi:"dbClusterVersion"`
	// The db node group count. The number should between 1 and 48.
	DbNodeGroupCount int `pulumi:"dbNodeGroupCount"`
	// The db node storage.
	DbNodeStorage string `pulumi:"dbNodeStorage"`
	// Key management service KMS key ID.
	EncryptionKey *string `pulumi:"encryptionKey"`
	// Currently only supports ECS disk encryption, with a value of CloudDisk, not encrypted when empty.
	EncryptionType *string `pulumi:"encryptionType"`
	// The maintenance window of DBCluster. Valid format: `hh:mmZ-hh:mm Z`.
	MaintainTime *string `pulumi:"maintainTime"`
	// The payment type of the resource. Valid values: `PayAsYouGo`,`Subscription`.
	PaymentType string `pulumi:"paymentType"`
	// Pre-paid cluster of the pay-as-you-go cycle. Valid values: `Month`, `Year`.
	Period *string `pulumi:"period"`
	// The status of the resource. Valid values: `Running`,`Creating`,`Deleting`,`Restarting`,`Preparing`,.
	Status *string `pulumi:"status"`
	// Storage type of DBCluster. Valid values: `cloudEssd`, `cloudEfficiency`, `cloudEssdPl2`, `cloudEssdPl3`.
	StorageType string `pulumi:"storageType"`
	// The used time of DBCluster.
	UsedTime *string `pulumi:"usedTime"`
	// The vswitch id of DBCluster.
	VswitchId *string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a DbCluster resource.
type DbClusterArgs struct {
	// The Category of DBCluster. Valid values: `Basic`,`HighAvailability`.
	Category pulumi.StringInput
	// The db cluster access white list.
	DbClusterAccessWhiteLists DbClusterDbClusterAccessWhiteListArrayInput
	// The DBCluster class. According to the category, dbClusterClass has two value ranges:
	// * Under the condition that the category is the `Basic`, Valid values: `S4-NEW`, `S8`, `S16`, `S32`, `S64`, `S104`.
	// * Under the condition that the category is the `HighAvailability`, Valid values: `C4-NEW`, `C8`, `C16`, `C32`, `C64`, `C104`.
	DbClusterClass pulumi.StringInput
	// The DBCluster description.
	DbClusterDescription pulumi.StringPtrInput
	// The DBCluster network type. Valid values: `vpc`.
	DbClusterNetworkType pulumi.StringInput
	// The DBCluster version. Valid values: `19.15.2.2`, `20.3.10.75`, `20.8.7.15`.
	DbClusterVersion pulumi.StringInput
	// The db node group count. The number should between 1 and 48.
	DbNodeGroupCount pulumi.IntInput
	// The db node storage.
	DbNodeStorage pulumi.StringInput
	// Key management service KMS key ID.
	EncryptionKey pulumi.StringPtrInput
	// Currently only supports ECS disk encryption, with a value of CloudDisk, not encrypted when empty.
	EncryptionType pulumi.StringPtrInput
	// The maintenance window of DBCluster. Valid format: `hh:mmZ-hh:mm Z`.
	MaintainTime pulumi.StringPtrInput
	// The payment type of the resource. Valid values: `PayAsYouGo`,`Subscription`.
	PaymentType pulumi.StringInput
	// Pre-paid cluster of the pay-as-you-go cycle. Valid values: `Month`, `Year`.
	Period pulumi.StringPtrInput
	// The status of the resource. Valid values: `Running`,`Creating`,`Deleting`,`Restarting`,`Preparing`,.
	Status pulumi.StringPtrInput
	// Storage type of DBCluster. Valid values: `cloudEssd`, `cloudEfficiency`, `cloudEssdPl2`, `cloudEssdPl3`.
	StorageType pulumi.StringInput
	// The used time of DBCluster.
	UsedTime pulumi.StringPtrInput
	// The vswitch id of DBCluster.
	VswitchId pulumi.StringPtrInput
}

func (DbClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbClusterArgs)(nil)).Elem()
}

type DbClusterInput interface {
	pulumi.Input

	ToDbClusterOutput() DbClusterOutput
	ToDbClusterOutputWithContext(ctx context.Context) DbClusterOutput
}

func (*DbCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*DbCluster)(nil))
}

func (i *DbCluster) ToDbClusterOutput() DbClusterOutput {
	return i.ToDbClusterOutputWithContext(context.Background())
}

func (i *DbCluster) ToDbClusterOutputWithContext(ctx context.Context) DbClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbClusterOutput)
}

func (i *DbCluster) ToDbClusterPtrOutput() DbClusterPtrOutput {
	return i.ToDbClusterPtrOutputWithContext(context.Background())
}

func (i *DbCluster) ToDbClusterPtrOutputWithContext(ctx context.Context) DbClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbClusterPtrOutput)
}

type DbClusterPtrInput interface {
	pulumi.Input

	ToDbClusterPtrOutput() DbClusterPtrOutput
	ToDbClusterPtrOutputWithContext(ctx context.Context) DbClusterPtrOutput
}

type dbClusterPtrType DbClusterArgs

func (*dbClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DbCluster)(nil))
}

func (i *dbClusterPtrType) ToDbClusterPtrOutput() DbClusterPtrOutput {
	return i.ToDbClusterPtrOutputWithContext(context.Background())
}

func (i *dbClusterPtrType) ToDbClusterPtrOutputWithContext(ctx context.Context) DbClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbClusterPtrOutput)
}

// DbClusterArrayInput is an input type that accepts DbClusterArray and DbClusterArrayOutput values.
// You can construct a concrete instance of `DbClusterArrayInput` via:
//
//          DbClusterArray{ DbClusterArgs{...} }
type DbClusterArrayInput interface {
	pulumi.Input

	ToDbClusterArrayOutput() DbClusterArrayOutput
	ToDbClusterArrayOutputWithContext(context.Context) DbClusterArrayOutput
}

type DbClusterArray []DbClusterInput

func (DbClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DbCluster)(nil)).Elem()
}

func (i DbClusterArray) ToDbClusterArrayOutput() DbClusterArrayOutput {
	return i.ToDbClusterArrayOutputWithContext(context.Background())
}

func (i DbClusterArray) ToDbClusterArrayOutputWithContext(ctx context.Context) DbClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbClusterArrayOutput)
}

// DbClusterMapInput is an input type that accepts DbClusterMap and DbClusterMapOutput values.
// You can construct a concrete instance of `DbClusterMapInput` via:
//
//          DbClusterMap{ "key": DbClusterArgs{...} }
type DbClusterMapInput interface {
	pulumi.Input

	ToDbClusterMapOutput() DbClusterMapOutput
	ToDbClusterMapOutputWithContext(context.Context) DbClusterMapOutput
}

type DbClusterMap map[string]DbClusterInput

func (DbClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DbCluster)(nil)).Elem()
}

func (i DbClusterMap) ToDbClusterMapOutput() DbClusterMapOutput {
	return i.ToDbClusterMapOutputWithContext(context.Background())
}

func (i DbClusterMap) ToDbClusterMapOutputWithContext(ctx context.Context) DbClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbClusterMapOutput)
}

type DbClusterOutput struct{ *pulumi.OutputState }

func (DbClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DbCluster)(nil))
}

func (o DbClusterOutput) ToDbClusterOutput() DbClusterOutput {
	return o
}

func (o DbClusterOutput) ToDbClusterOutputWithContext(ctx context.Context) DbClusterOutput {
	return o
}

func (o DbClusterOutput) ToDbClusterPtrOutput() DbClusterPtrOutput {
	return o.ToDbClusterPtrOutputWithContext(context.Background())
}

func (o DbClusterOutput) ToDbClusterPtrOutputWithContext(ctx context.Context) DbClusterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DbCluster) *DbCluster {
		return &v
	}).(DbClusterPtrOutput)
}

type DbClusterPtrOutput struct{ *pulumi.OutputState }

func (DbClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DbCluster)(nil))
}

func (o DbClusterPtrOutput) ToDbClusterPtrOutput() DbClusterPtrOutput {
	return o
}

func (o DbClusterPtrOutput) ToDbClusterPtrOutputWithContext(ctx context.Context) DbClusterPtrOutput {
	return o
}

func (o DbClusterPtrOutput) Elem() DbClusterOutput {
	return o.ApplyT(func(v *DbCluster) DbCluster {
		if v != nil {
			return *v
		}
		var ret DbCluster
		return ret
	}).(DbClusterOutput)
}

type DbClusterArrayOutput struct{ *pulumi.OutputState }

func (DbClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DbCluster)(nil))
}

func (o DbClusterArrayOutput) ToDbClusterArrayOutput() DbClusterArrayOutput {
	return o
}

func (o DbClusterArrayOutput) ToDbClusterArrayOutputWithContext(ctx context.Context) DbClusterArrayOutput {
	return o
}

func (o DbClusterArrayOutput) Index(i pulumi.IntInput) DbClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DbCluster {
		return vs[0].([]DbCluster)[vs[1].(int)]
	}).(DbClusterOutput)
}

type DbClusterMapOutput struct{ *pulumi.OutputState }

func (DbClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DbCluster)(nil))
}

func (o DbClusterMapOutput) ToDbClusterMapOutput() DbClusterMapOutput {
	return o
}

func (o DbClusterMapOutput) ToDbClusterMapOutputWithContext(ctx context.Context) DbClusterMapOutput {
	return o
}

func (o DbClusterMapOutput) MapIndex(k pulumi.StringInput) DbClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DbCluster {
		return vs[0].(map[string]DbCluster)[vs[1].(string)]
	}).(DbClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DbClusterInput)(nil)).Elem(), &DbCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbClusterPtrInput)(nil)).Elem(), &DbCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbClusterArrayInput)(nil)).Elem(), DbClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbClusterMapInput)(nil)).Elem(), DbClusterMap{})
	pulumi.RegisterOutputType(DbClusterOutput{})
	pulumi.RegisterOutputType(DbClusterPtrOutput{})
	pulumi.RegisterOutputType(DbClusterArrayOutput{})
	pulumi.RegisterOutputType(DbClusterMapOutput{})
}
