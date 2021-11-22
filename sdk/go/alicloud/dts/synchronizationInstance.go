// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dts

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DTS Synchronization Instance resource.
//
// For information about DTS Synchronization Instance and how to use it, see [What is Synchronization Instance](https://help.aliyun.com/document_detail/211599.html).
//
// > **NOTE:** Available in v1.138.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dts"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dts.NewSynchronizationInstance(ctx, "_default", &dts.SynchronizationInstanceArgs{
// 			DestinationEndpointEngineName: pulumi.String("ADB30"),
// 			DestinationEndpointRegion:     pulumi.String("cn-hangzhou"),
// 			InstanceClass:                 pulumi.String("small"),
// 			PaymentType:                   pulumi.String("PayAsYouGo"),
// 			SourceEndpointEngineName:      pulumi.String("PolarDB"),
// 			SourceEndpointRegion:          pulumi.String("cn-hangzhou"),
// 			SyncArchitecture:              pulumi.String("oneway"),
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
// DTS Synchronization Instance can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:dts/synchronizationInstance:SynchronizationInstance example <id>
// ```
type SynchronizationInstance struct {
	pulumi.CustomResourceState

	// [ETL specifications](https://help.aliyun.com/document_detail/212324.html). The unit is the computing unit ComputeUnit (CU), 1CU=1vCPU+4 GB memory. The value range is an integer greater than or equal to 2.
	ComputeUnit pulumi.IntPtrOutput `pulumi:"computeUnit"`
	// The number of private customized RDS instances under PolarDB-X. The default value is 1. This parameter needs to be passed only when `sourceEndpointEngineName` equals `drds`.
	DatabaseCount pulumi.IntPtrOutput `pulumi:"databaseCount"`
	// The type of destination engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
	DestinationEndpointEngineName pulumi.StringOutput `pulumi:"destinationEndpointEngineName"`
	// The region of destination instance. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).
	DestinationEndpointRegion pulumi.StringOutput `pulumi:"destinationEndpointRegion"`
	// The instance class. Valid values: `large`, `medium`, `micro`, `small`, `xlarge`, `xxlarge`. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).
	InstanceClass pulumi.StringOutput `pulumi:"instanceClass"`
	// The duration of prepaid instance purchase. When `paymentType` is `Subscription`, this parameter is valid and must be passed in.
	PaymentDuration pulumi.IntPtrOutput `pulumi:"paymentDuration"`
	// The payment duration unit. Valid values: `Month`, `Year`. When `paymentType` is `Subscription`, this parameter is valid and must be passed in.
	PaymentDurationUnit pulumi.StringPtrOutput `pulumi:"paymentDurationUnit"`
	// The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The number of instances purchased.
	Quantity pulumi.IntPtrOutput `pulumi:"quantity"`
	// The type of source endpoint engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
	SourceEndpointEngineName pulumi.StringOutput `pulumi:"sourceEndpointEngineName"`
	// The region of source instance.
	SourceEndpointRegion pulumi.StringOutput `pulumi:"sourceEndpointRegion"`
	// The status.
	Status pulumi.StringOutput `pulumi:"status"`
	// The sync architecture. Valid values: `oneway`, `bidirectional`.
	SyncArchitecture pulumi.StringPtrOutput `pulumi:"syncArchitecture"`
}

// NewSynchronizationInstance registers a new resource with the given unique name, arguments, and options.
func NewSynchronizationInstance(ctx *pulumi.Context,
	name string, args *SynchronizationInstanceArgs, opts ...pulumi.ResourceOption) (*SynchronizationInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationEndpointEngineName == nil {
		return nil, errors.New("invalid value for required argument 'DestinationEndpointEngineName'")
	}
	if args.DestinationEndpointRegion == nil {
		return nil, errors.New("invalid value for required argument 'DestinationEndpointRegion'")
	}
	if args.PaymentType == nil {
		return nil, errors.New("invalid value for required argument 'PaymentType'")
	}
	if args.SourceEndpointEngineName == nil {
		return nil, errors.New("invalid value for required argument 'SourceEndpointEngineName'")
	}
	if args.SourceEndpointRegion == nil {
		return nil, errors.New("invalid value for required argument 'SourceEndpointRegion'")
	}
	var resource SynchronizationInstance
	err := ctx.RegisterResource("alicloud:dts/synchronizationInstance:SynchronizationInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSynchronizationInstance gets an existing SynchronizationInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSynchronizationInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SynchronizationInstanceState, opts ...pulumi.ResourceOption) (*SynchronizationInstance, error) {
	var resource SynchronizationInstance
	err := ctx.ReadResource("alicloud:dts/synchronizationInstance:SynchronizationInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SynchronizationInstance resources.
type synchronizationInstanceState struct {
	// [ETL specifications](https://help.aliyun.com/document_detail/212324.html). The unit is the computing unit ComputeUnit (CU), 1CU=1vCPU+4 GB memory. The value range is an integer greater than or equal to 2.
	ComputeUnit *int `pulumi:"computeUnit"`
	// The number of private customized RDS instances under PolarDB-X. The default value is 1. This parameter needs to be passed only when `sourceEndpointEngineName` equals `drds`.
	DatabaseCount *int `pulumi:"databaseCount"`
	// The type of destination engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
	DestinationEndpointEngineName *string `pulumi:"destinationEndpointEngineName"`
	// The region of destination instance. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).
	DestinationEndpointRegion *string `pulumi:"destinationEndpointRegion"`
	// The instance class. Valid values: `large`, `medium`, `micro`, `small`, `xlarge`, `xxlarge`. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).
	InstanceClass *string `pulumi:"instanceClass"`
	// The duration of prepaid instance purchase. When `paymentType` is `Subscription`, this parameter is valid and must be passed in.
	PaymentDuration *int `pulumi:"paymentDuration"`
	// The payment duration unit. Valid values: `Month`, `Year`. When `paymentType` is `Subscription`, this parameter is valid and must be passed in.
	PaymentDurationUnit *string `pulumi:"paymentDurationUnit"`
	// The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`.
	PaymentType *string `pulumi:"paymentType"`
	// The number of instances purchased.
	Quantity *int `pulumi:"quantity"`
	// The type of source endpoint engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
	SourceEndpointEngineName *string `pulumi:"sourceEndpointEngineName"`
	// The region of source instance.
	SourceEndpointRegion *string `pulumi:"sourceEndpointRegion"`
	// The status.
	Status *string `pulumi:"status"`
	// The sync architecture. Valid values: `oneway`, `bidirectional`.
	SyncArchitecture *string `pulumi:"syncArchitecture"`
}

type SynchronizationInstanceState struct {
	// [ETL specifications](https://help.aliyun.com/document_detail/212324.html). The unit is the computing unit ComputeUnit (CU), 1CU=1vCPU+4 GB memory. The value range is an integer greater than or equal to 2.
	ComputeUnit pulumi.IntPtrInput
	// The number of private customized RDS instances under PolarDB-X. The default value is 1. This parameter needs to be passed only when `sourceEndpointEngineName` equals `drds`.
	DatabaseCount pulumi.IntPtrInput
	// The type of destination engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
	DestinationEndpointEngineName pulumi.StringPtrInput
	// The region of destination instance. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).
	DestinationEndpointRegion pulumi.StringPtrInput
	// The instance class. Valid values: `large`, `medium`, `micro`, `small`, `xlarge`, `xxlarge`. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).
	InstanceClass pulumi.StringPtrInput
	// The duration of prepaid instance purchase. When `paymentType` is `Subscription`, this parameter is valid and must be passed in.
	PaymentDuration pulumi.IntPtrInput
	// The payment duration unit. Valid values: `Month`, `Year`. When `paymentType` is `Subscription`, this parameter is valid and must be passed in.
	PaymentDurationUnit pulumi.StringPtrInput
	// The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`.
	PaymentType pulumi.StringPtrInput
	// The number of instances purchased.
	Quantity pulumi.IntPtrInput
	// The type of source endpoint engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
	SourceEndpointEngineName pulumi.StringPtrInput
	// The region of source instance.
	SourceEndpointRegion pulumi.StringPtrInput
	// The status.
	Status pulumi.StringPtrInput
	// The sync architecture. Valid values: `oneway`, `bidirectional`.
	SyncArchitecture pulumi.StringPtrInput
}

func (SynchronizationInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*synchronizationInstanceState)(nil)).Elem()
}

type synchronizationInstanceArgs struct {
	// [ETL specifications](https://help.aliyun.com/document_detail/212324.html). The unit is the computing unit ComputeUnit (CU), 1CU=1vCPU+4 GB memory. The value range is an integer greater than or equal to 2.
	ComputeUnit *int `pulumi:"computeUnit"`
	// The number of private customized RDS instances under PolarDB-X. The default value is 1. This parameter needs to be passed only when `sourceEndpointEngineName` equals `drds`.
	DatabaseCount *int `pulumi:"databaseCount"`
	// The type of destination engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
	DestinationEndpointEngineName string `pulumi:"destinationEndpointEngineName"`
	// The region of destination instance. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).
	DestinationEndpointRegion string `pulumi:"destinationEndpointRegion"`
	// The instance class. Valid values: `large`, `medium`, `micro`, `small`, `xlarge`, `xxlarge`. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).
	InstanceClass *string `pulumi:"instanceClass"`
	// The duration of prepaid instance purchase. When `paymentType` is `Subscription`, this parameter is valid and must be passed in.
	PaymentDuration *int `pulumi:"paymentDuration"`
	// The payment duration unit. Valid values: `Month`, `Year`. When `paymentType` is `Subscription`, this parameter is valid and must be passed in.
	PaymentDurationUnit *string `pulumi:"paymentDurationUnit"`
	// The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`.
	PaymentType string `pulumi:"paymentType"`
	// The number of instances purchased.
	Quantity *int `pulumi:"quantity"`
	// The type of source endpoint engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
	SourceEndpointEngineName string `pulumi:"sourceEndpointEngineName"`
	// The region of source instance.
	SourceEndpointRegion string `pulumi:"sourceEndpointRegion"`
	// The sync architecture. Valid values: `oneway`, `bidirectional`.
	SyncArchitecture *string `pulumi:"syncArchitecture"`
}

// The set of arguments for constructing a SynchronizationInstance resource.
type SynchronizationInstanceArgs struct {
	// [ETL specifications](https://help.aliyun.com/document_detail/212324.html). The unit is the computing unit ComputeUnit (CU), 1CU=1vCPU+4 GB memory. The value range is an integer greater than or equal to 2.
	ComputeUnit pulumi.IntPtrInput
	// The number of private customized RDS instances under PolarDB-X. The default value is 1. This parameter needs to be passed only when `sourceEndpointEngineName` equals `drds`.
	DatabaseCount pulumi.IntPtrInput
	// The type of destination engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
	DestinationEndpointEngineName pulumi.StringInput
	// The region of destination instance. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).
	DestinationEndpointRegion pulumi.StringInput
	// The instance class. Valid values: `large`, `medium`, `micro`, `small`, `xlarge`, `xxlarge`. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).
	InstanceClass pulumi.StringPtrInput
	// The duration of prepaid instance purchase. When `paymentType` is `Subscription`, this parameter is valid and must be passed in.
	PaymentDuration pulumi.IntPtrInput
	// The payment duration unit. Valid values: `Month`, `Year`. When `paymentType` is `Subscription`, this parameter is valid and must be passed in.
	PaymentDurationUnit pulumi.StringPtrInput
	// The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`.
	PaymentType pulumi.StringInput
	// The number of instances purchased.
	Quantity pulumi.IntPtrInput
	// The type of source endpoint engine. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardbO`, `polardbPg`, `tidb`. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).
	SourceEndpointEngineName pulumi.StringInput
	// The region of source instance.
	SourceEndpointRegion pulumi.StringInput
	// The sync architecture. Valid values: `oneway`, `bidirectional`.
	SyncArchitecture pulumi.StringPtrInput
}

func (SynchronizationInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*synchronizationInstanceArgs)(nil)).Elem()
}

type SynchronizationInstanceInput interface {
	pulumi.Input

	ToSynchronizationInstanceOutput() SynchronizationInstanceOutput
	ToSynchronizationInstanceOutputWithContext(ctx context.Context) SynchronizationInstanceOutput
}

func (*SynchronizationInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*SynchronizationInstance)(nil))
}

func (i *SynchronizationInstance) ToSynchronizationInstanceOutput() SynchronizationInstanceOutput {
	return i.ToSynchronizationInstanceOutputWithContext(context.Background())
}

func (i *SynchronizationInstance) ToSynchronizationInstanceOutputWithContext(ctx context.Context) SynchronizationInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationInstanceOutput)
}

func (i *SynchronizationInstance) ToSynchronizationInstancePtrOutput() SynchronizationInstancePtrOutput {
	return i.ToSynchronizationInstancePtrOutputWithContext(context.Background())
}

func (i *SynchronizationInstance) ToSynchronizationInstancePtrOutputWithContext(ctx context.Context) SynchronizationInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationInstancePtrOutput)
}

type SynchronizationInstancePtrInput interface {
	pulumi.Input

	ToSynchronizationInstancePtrOutput() SynchronizationInstancePtrOutput
	ToSynchronizationInstancePtrOutputWithContext(ctx context.Context) SynchronizationInstancePtrOutput
}

type synchronizationInstancePtrType SynchronizationInstanceArgs

func (*synchronizationInstancePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SynchronizationInstance)(nil))
}

func (i *synchronizationInstancePtrType) ToSynchronizationInstancePtrOutput() SynchronizationInstancePtrOutput {
	return i.ToSynchronizationInstancePtrOutputWithContext(context.Background())
}

func (i *synchronizationInstancePtrType) ToSynchronizationInstancePtrOutputWithContext(ctx context.Context) SynchronizationInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationInstancePtrOutput)
}

// SynchronizationInstanceArrayInput is an input type that accepts SynchronizationInstanceArray and SynchronizationInstanceArrayOutput values.
// You can construct a concrete instance of `SynchronizationInstanceArrayInput` via:
//
//          SynchronizationInstanceArray{ SynchronizationInstanceArgs{...} }
type SynchronizationInstanceArrayInput interface {
	pulumi.Input

	ToSynchronizationInstanceArrayOutput() SynchronizationInstanceArrayOutput
	ToSynchronizationInstanceArrayOutputWithContext(context.Context) SynchronizationInstanceArrayOutput
}

type SynchronizationInstanceArray []SynchronizationInstanceInput

func (SynchronizationInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SynchronizationInstance)(nil)).Elem()
}

func (i SynchronizationInstanceArray) ToSynchronizationInstanceArrayOutput() SynchronizationInstanceArrayOutput {
	return i.ToSynchronizationInstanceArrayOutputWithContext(context.Background())
}

func (i SynchronizationInstanceArray) ToSynchronizationInstanceArrayOutputWithContext(ctx context.Context) SynchronizationInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationInstanceArrayOutput)
}

// SynchronizationInstanceMapInput is an input type that accepts SynchronizationInstanceMap and SynchronizationInstanceMapOutput values.
// You can construct a concrete instance of `SynchronizationInstanceMapInput` via:
//
//          SynchronizationInstanceMap{ "key": SynchronizationInstanceArgs{...} }
type SynchronizationInstanceMapInput interface {
	pulumi.Input

	ToSynchronizationInstanceMapOutput() SynchronizationInstanceMapOutput
	ToSynchronizationInstanceMapOutputWithContext(context.Context) SynchronizationInstanceMapOutput
}

type SynchronizationInstanceMap map[string]SynchronizationInstanceInput

func (SynchronizationInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SynchronizationInstance)(nil)).Elem()
}

func (i SynchronizationInstanceMap) ToSynchronizationInstanceMapOutput() SynchronizationInstanceMapOutput {
	return i.ToSynchronizationInstanceMapOutputWithContext(context.Background())
}

func (i SynchronizationInstanceMap) ToSynchronizationInstanceMapOutputWithContext(ctx context.Context) SynchronizationInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationInstanceMapOutput)
}

type SynchronizationInstanceOutput struct{ *pulumi.OutputState }

func (SynchronizationInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SynchronizationInstance)(nil))
}

func (o SynchronizationInstanceOutput) ToSynchronizationInstanceOutput() SynchronizationInstanceOutput {
	return o
}

func (o SynchronizationInstanceOutput) ToSynchronizationInstanceOutputWithContext(ctx context.Context) SynchronizationInstanceOutput {
	return o
}

func (o SynchronizationInstanceOutput) ToSynchronizationInstancePtrOutput() SynchronizationInstancePtrOutput {
	return o.ToSynchronizationInstancePtrOutputWithContext(context.Background())
}

func (o SynchronizationInstanceOutput) ToSynchronizationInstancePtrOutputWithContext(ctx context.Context) SynchronizationInstancePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SynchronizationInstance) *SynchronizationInstance {
		return &v
	}).(SynchronizationInstancePtrOutput)
}

type SynchronizationInstancePtrOutput struct{ *pulumi.OutputState }

func (SynchronizationInstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SynchronizationInstance)(nil))
}

func (o SynchronizationInstancePtrOutput) ToSynchronizationInstancePtrOutput() SynchronizationInstancePtrOutput {
	return o
}

func (o SynchronizationInstancePtrOutput) ToSynchronizationInstancePtrOutputWithContext(ctx context.Context) SynchronizationInstancePtrOutput {
	return o
}

func (o SynchronizationInstancePtrOutput) Elem() SynchronizationInstanceOutput {
	return o.ApplyT(func(v *SynchronizationInstance) SynchronizationInstance {
		if v != nil {
			return *v
		}
		var ret SynchronizationInstance
		return ret
	}).(SynchronizationInstanceOutput)
}

type SynchronizationInstanceArrayOutput struct{ *pulumi.OutputState }

func (SynchronizationInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SynchronizationInstance)(nil))
}

func (o SynchronizationInstanceArrayOutput) ToSynchronizationInstanceArrayOutput() SynchronizationInstanceArrayOutput {
	return o
}

func (o SynchronizationInstanceArrayOutput) ToSynchronizationInstanceArrayOutputWithContext(ctx context.Context) SynchronizationInstanceArrayOutput {
	return o
}

func (o SynchronizationInstanceArrayOutput) Index(i pulumi.IntInput) SynchronizationInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SynchronizationInstance {
		return vs[0].([]SynchronizationInstance)[vs[1].(int)]
	}).(SynchronizationInstanceOutput)
}

type SynchronizationInstanceMapOutput struct{ *pulumi.OutputState }

func (SynchronizationInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SynchronizationInstance)(nil))
}

func (o SynchronizationInstanceMapOutput) ToSynchronizationInstanceMapOutput() SynchronizationInstanceMapOutput {
	return o
}

func (o SynchronizationInstanceMapOutput) ToSynchronizationInstanceMapOutputWithContext(ctx context.Context) SynchronizationInstanceMapOutput {
	return o
}

func (o SynchronizationInstanceMapOutput) MapIndex(k pulumi.StringInput) SynchronizationInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SynchronizationInstance {
		return vs[0].(map[string]SynchronizationInstance)[vs[1].(string)]
	}).(SynchronizationInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SynchronizationInstanceInput)(nil)).Elem(), &SynchronizationInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*SynchronizationInstancePtrInput)(nil)).Elem(), &SynchronizationInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*SynchronizationInstanceArrayInput)(nil)).Elem(), SynchronizationInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SynchronizationInstanceMapInput)(nil)).Elem(), SynchronizationInstanceMap{})
	pulumi.RegisterOutputType(SynchronizationInstanceOutput{})
	pulumi.RegisterOutputType(SynchronizationInstancePtrOutput{})
	pulumi.RegisterOutputType(SynchronizationInstanceArrayOutput{})
	pulumi.RegisterOutputType(SynchronizationInstanceMapOutput{})
}
