// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package quotas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Quotas Quota Alarm resource.
//
// For information about Quotas Quota Alarm and how to use it, see [What is Quota Alarm](https://help.aliyun.com/document_detail/184343.html).
//
// > **NOTE:** Available in v1.116.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/quotas"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := quotas.NewQuotaAlarm(ctx, "example", &quotas.QuotaAlarmArgs{
// 			ProductCode:     pulumi.String("ecs"),
// 			QuotaActionCode: pulumi.String("q_prepaid-instance-count-per-once-purchase"),
// 			QuotaAlarmName:  pulumi.String("tf-testAcc"),
// 			QuotaDimensions: quotas.QuotaAlarmQuotaDimensionArray{
// 				&quotas.QuotaAlarmQuotaDimensionArgs{
// 					Key:   pulumi.String("regionId"),
// 					Value: pulumi.String("cn-hangzhou"),
// 				},
// 			},
// 			Threshold: pulumi.Float64(100),
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
// Quotas Quota Alarm can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:quotas/quotaAlarm:QuotaAlarm example <id>
// ```
type QuotaAlarm struct {
	pulumi.CustomResourceState

	// The Product Code.
	ProductCode pulumi.StringOutput `pulumi:"productCode"`
	// The Quota Action Code.
	QuotaActionCode pulumi.StringOutput `pulumi:"quotaActionCode"`
	// The name of Quota Alarm.
	QuotaAlarmName pulumi.StringOutput `pulumi:"quotaAlarmName"`
	// The Quota Dimensions.
	QuotaDimensions QuotaAlarmQuotaDimensionArrayOutput `pulumi:"quotaDimensions"`
	// The threshold of Quota Alarm.
	Threshold pulumi.Float64PtrOutput `pulumi:"threshold"`
	// The threshold percent of Quota Alarm.
	ThresholdPercent pulumi.Float64PtrOutput `pulumi:"thresholdPercent"`
	// The WebHook of Quota Alarm.
	WebHook pulumi.StringPtrOutput `pulumi:"webHook"`
}

// NewQuotaAlarm registers a new resource with the given unique name, arguments, and options.
func NewQuotaAlarm(ctx *pulumi.Context,
	name string, args *QuotaAlarmArgs, opts ...pulumi.ResourceOption) (*QuotaAlarm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductCode == nil {
		return nil, errors.New("invalid value for required argument 'ProductCode'")
	}
	if args.QuotaActionCode == nil {
		return nil, errors.New("invalid value for required argument 'QuotaActionCode'")
	}
	if args.QuotaAlarmName == nil {
		return nil, errors.New("invalid value for required argument 'QuotaAlarmName'")
	}
	var resource QuotaAlarm
	err := ctx.RegisterResource("alicloud:quotas/quotaAlarm:QuotaAlarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuotaAlarm gets an existing QuotaAlarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuotaAlarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QuotaAlarmState, opts ...pulumi.ResourceOption) (*QuotaAlarm, error) {
	var resource QuotaAlarm
	err := ctx.ReadResource("alicloud:quotas/quotaAlarm:QuotaAlarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QuotaAlarm resources.
type quotaAlarmState struct {
	// The Product Code.
	ProductCode *string `pulumi:"productCode"`
	// The Quota Action Code.
	QuotaActionCode *string `pulumi:"quotaActionCode"`
	// The name of Quota Alarm.
	QuotaAlarmName *string `pulumi:"quotaAlarmName"`
	// The Quota Dimensions.
	QuotaDimensions []QuotaAlarmQuotaDimension `pulumi:"quotaDimensions"`
	// The threshold of Quota Alarm.
	Threshold *float64 `pulumi:"threshold"`
	// The threshold percent of Quota Alarm.
	ThresholdPercent *float64 `pulumi:"thresholdPercent"`
	// The WebHook of Quota Alarm.
	WebHook *string `pulumi:"webHook"`
}

type QuotaAlarmState struct {
	// The Product Code.
	ProductCode pulumi.StringPtrInput
	// The Quota Action Code.
	QuotaActionCode pulumi.StringPtrInput
	// The name of Quota Alarm.
	QuotaAlarmName pulumi.StringPtrInput
	// The Quota Dimensions.
	QuotaDimensions QuotaAlarmQuotaDimensionArrayInput
	// The threshold of Quota Alarm.
	Threshold pulumi.Float64PtrInput
	// The threshold percent of Quota Alarm.
	ThresholdPercent pulumi.Float64PtrInput
	// The WebHook of Quota Alarm.
	WebHook pulumi.StringPtrInput
}

func (QuotaAlarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaAlarmState)(nil)).Elem()
}

type quotaAlarmArgs struct {
	// The Product Code.
	ProductCode string `pulumi:"productCode"`
	// The Quota Action Code.
	QuotaActionCode string `pulumi:"quotaActionCode"`
	// The name of Quota Alarm.
	QuotaAlarmName string `pulumi:"quotaAlarmName"`
	// The Quota Dimensions.
	QuotaDimensions []QuotaAlarmQuotaDimension `pulumi:"quotaDimensions"`
	// The threshold of Quota Alarm.
	Threshold *float64 `pulumi:"threshold"`
	// The threshold percent of Quota Alarm.
	ThresholdPercent *float64 `pulumi:"thresholdPercent"`
	// The WebHook of Quota Alarm.
	WebHook *string `pulumi:"webHook"`
}

// The set of arguments for constructing a QuotaAlarm resource.
type QuotaAlarmArgs struct {
	// The Product Code.
	ProductCode pulumi.StringInput
	// The Quota Action Code.
	QuotaActionCode pulumi.StringInput
	// The name of Quota Alarm.
	QuotaAlarmName pulumi.StringInput
	// The Quota Dimensions.
	QuotaDimensions QuotaAlarmQuotaDimensionArrayInput
	// The threshold of Quota Alarm.
	Threshold pulumi.Float64PtrInput
	// The threshold percent of Quota Alarm.
	ThresholdPercent pulumi.Float64PtrInput
	// The WebHook of Quota Alarm.
	WebHook pulumi.StringPtrInput
}

func (QuotaAlarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaAlarmArgs)(nil)).Elem()
}

type QuotaAlarmInput interface {
	pulumi.Input

	ToQuotaAlarmOutput() QuotaAlarmOutput
	ToQuotaAlarmOutputWithContext(ctx context.Context) QuotaAlarmOutput
}

func (*QuotaAlarm) ElementType() reflect.Type {
	return reflect.TypeOf((*QuotaAlarm)(nil))
}

func (i *QuotaAlarm) ToQuotaAlarmOutput() QuotaAlarmOutput {
	return i.ToQuotaAlarmOutputWithContext(context.Background())
}

func (i *QuotaAlarm) ToQuotaAlarmOutputWithContext(ctx context.Context) QuotaAlarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaAlarmOutput)
}

func (i *QuotaAlarm) ToQuotaAlarmPtrOutput() QuotaAlarmPtrOutput {
	return i.ToQuotaAlarmPtrOutputWithContext(context.Background())
}

func (i *QuotaAlarm) ToQuotaAlarmPtrOutputWithContext(ctx context.Context) QuotaAlarmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaAlarmPtrOutput)
}

type QuotaAlarmPtrInput interface {
	pulumi.Input

	ToQuotaAlarmPtrOutput() QuotaAlarmPtrOutput
	ToQuotaAlarmPtrOutputWithContext(ctx context.Context) QuotaAlarmPtrOutput
}

type quotaAlarmPtrType QuotaAlarmArgs

func (*quotaAlarmPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaAlarm)(nil))
}

func (i *quotaAlarmPtrType) ToQuotaAlarmPtrOutput() QuotaAlarmPtrOutput {
	return i.ToQuotaAlarmPtrOutputWithContext(context.Background())
}

func (i *quotaAlarmPtrType) ToQuotaAlarmPtrOutputWithContext(ctx context.Context) QuotaAlarmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaAlarmPtrOutput)
}

// QuotaAlarmArrayInput is an input type that accepts QuotaAlarmArray and QuotaAlarmArrayOutput values.
// You can construct a concrete instance of `QuotaAlarmArrayInput` via:
//
//          QuotaAlarmArray{ QuotaAlarmArgs{...} }
type QuotaAlarmArrayInput interface {
	pulumi.Input

	ToQuotaAlarmArrayOutput() QuotaAlarmArrayOutput
	ToQuotaAlarmArrayOutputWithContext(context.Context) QuotaAlarmArrayOutput
}

type QuotaAlarmArray []QuotaAlarmInput

func (QuotaAlarmArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*QuotaAlarm)(nil))
}

func (i QuotaAlarmArray) ToQuotaAlarmArrayOutput() QuotaAlarmArrayOutput {
	return i.ToQuotaAlarmArrayOutputWithContext(context.Background())
}

func (i QuotaAlarmArray) ToQuotaAlarmArrayOutputWithContext(ctx context.Context) QuotaAlarmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaAlarmArrayOutput)
}

// QuotaAlarmMapInput is an input type that accepts QuotaAlarmMap and QuotaAlarmMapOutput values.
// You can construct a concrete instance of `QuotaAlarmMapInput` via:
//
//          QuotaAlarmMap{ "key": QuotaAlarmArgs{...} }
type QuotaAlarmMapInput interface {
	pulumi.Input

	ToQuotaAlarmMapOutput() QuotaAlarmMapOutput
	ToQuotaAlarmMapOutputWithContext(context.Context) QuotaAlarmMapOutput
}

type QuotaAlarmMap map[string]QuotaAlarmInput

func (QuotaAlarmMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*QuotaAlarm)(nil))
}

func (i QuotaAlarmMap) ToQuotaAlarmMapOutput() QuotaAlarmMapOutput {
	return i.ToQuotaAlarmMapOutputWithContext(context.Background())
}

func (i QuotaAlarmMap) ToQuotaAlarmMapOutputWithContext(ctx context.Context) QuotaAlarmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaAlarmMapOutput)
}

type QuotaAlarmOutput struct {
	*pulumi.OutputState
}

func (QuotaAlarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuotaAlarm)(nil))
}

func (o QuotaAlarmOutput) ToQuotaAlarmOutput() QuotaAlarmOutput {
	return o
}

func (o QuotaAlarmOutput) ToQuotaAlarmOutputWithContext(ctx context.Context) QuotaAlarmOutput {
	return o
}

func (o QuotaAlarmOutput) ToQuotaAlarmPtrOutput() QuotaAlarmPtrOutput {
	return o.ToQuotaAlarmPtrOutputWithContext(context.Background())
}

func (o QuotaAlarmOutput) ToQuotaAlarmPtrOutputWithContext(ctx context.Context) QuotaAlarmPtrOutput {
	return o.ApplyT(func(v QuotaAlarm) *QuotaAlarm {
		return &v
	}).(QuotaAlarmPtrOutput)
}

type QuotaAlarmPtrOutput struct {
	*pulumi.OutputState
}

func (QuotaAlarmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaAlarm)(nil))
}

func (o QuotaAlarmPtrOutput) ToQuotaAlarmPtrOutput() QuotaAlarmPtrOutput {
	return o
}

func (o QuotaAlarmPtrOutput) ToQuotaAlarmPtrOutputWithContext(ctx context.Context) QuotaAlarmPtrOutput {
	return o
}

type QuotaAlarmArrayOutput struct{ *pulumi.OutputState }

func (QuotaAlarmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QuotaAlarm)(nil))
}

func (o QuotaAlarmArrayOutput) ToQuotaAlarmArrayOutput() QuotaAlarmArrayOutput {
	return o
}

func (o QuotaAlarmArrayOutput) ToQuotaAlarmArrayOutputWithContext(ctx context.Context) QuotaAlarmArrayOutput {
	return o
}

func (o QuotaAlarmArrayOutput) Index(i pulumi.IntInput) QuotaAlarmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QuotaAlarm {
		return vs[0].([]QuotaAlarm)[vs[1].(int)]
	}).(QuotaAlarmOutput)
}

type QuotaAlarmMapOutput struct{ *pulumi.OutputState }

func (QuotaAlarmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]QuotaAlarm)(nil))
}

func (o QuotaAlarmMapOutput) ToQuotaAlarmMapOutput() QuotaAlarmMapOutput {
	return o
}

func (o QuotaAlarmMapOutput) ToQuotaAlarmMapOutputWithContext(ctx context.Context) QuotaAlarmMapOutput {
	return o
}

func (o QuotaAlarmMapOutput) MapIndex(k pulumi.StringInput) QuotaAlarmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) QuotaAlarm {
		return vs[0].(map[string]QuotaAlarm)[vs[1].(string)]
	}).(QuotaAlarmOutput)
}

func init() {
	pulumi.RegisterOutputType(QuotaAlarmOutput{})
	pulumi.RegisterOutputType(QuotaAlarmPtrOutput{})
	pulumi.RegisterOutputType(QuotaAlarmArrayOutput{})
	pulumi.RegisterOutputType(QuotaAlarmMapOutput{})
}
