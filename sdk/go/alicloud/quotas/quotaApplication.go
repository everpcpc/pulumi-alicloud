// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quotas

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Quotas Quota Application resource.
//
// For information about Quotas Quota Application and how to use it, see [What is Quota Application](https://help.aliyun.com/document_detail/171289.html).
//
// > **NOTE:** Available in v1.117.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/quotas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := quotas.NewQuotaApplication(ctx, "example", &quotas.QuotaApplicationArgs{
//				DesireValue: pulumi.Float64(100),
//				Dimensions: quotas.QuotaApplicationDimensionArray{
//					&quotas.QuotaApplicationDimensionArgs{
//						Key:   pulumi.String("regionId"),
//						Value: pulumi.String("cn-hangzhou"),
//					},
//				},
//				NoticeType:      pulumi.Int(0),
//				ProductCode:     pulumi.String("ess"),
//				QuotaActionCode: pulumi.String("q_db_instance"),
//				Reason:          pulumi.String("For Terraform Test"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Quotas Application Info can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:quotas/quotaApplication:QuotaApplication example <id>
//
// ```
type QuotaApplication struct {
	pulumi.CustomResourceState

	// The approve value of the quota application.
	ApproveValue pulumi.StringOutput `pulumi:"approveValue"`
	// The audit mode. Valid values: `Async`, `Sync`. Default to: `Async`.
	AuditMode pulumi.StringPtrOutput `pulumi:"auditMode"`
	// The audit reason.
	AuditReason pulumi.StringOutput `pulumi:"auditReason"`
	// The desire value of the quota application.
	DesireValue pulumi.Float64Output `pulumi:"desireValue"`
	// The quota dimensions.
	Dimensions QuotaApplicationDimensionArrayOutput `pulumi:"dimensions"`
	// The effective time of the quota application.
	EffectiveTime pulumi.StringOutput `pulumi:"effectiveTime"`
	// The expire time of the quota application.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// The notice type. Valid values: `0`, `1`, `2`, `3`.
	NoticeType pulumi.IntPtrOutput `pulumi:"noticeType"`
	// The product code.
	ProductCode pulumi.StringOutput `pulumi:"productCode"`
	// The ID of quota action.
	QuotaActionCode pulumi.StringOutput `pulumi:"quotaActionCode"`
	// The quota category. Valid values: `CommonQuota`, `FlowControl`.
	QuotaCategory pulumi.StringPtrOutput `pulumi:"quotaCategory"`
	// The description of the quota application.
	QuotaDescription pulumi.StringOutput `pulumi:"quotaDescription"`
	// The name of the quota application.
	QuotaName pulumi.StringOutput `pulumi:"quotaName"`
	// The unit of the quota application.
	QuotaUnit pulumi.StringOutput `pulumi:"quotaUnit"`
	// The reason of the quota application.
	Reason pulumi.StringOutput `pulumi:"reason"`
	// The status of the quota application.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewQuotaApplication registers a new resource with the given unique name, arguments, and options.
func NewQuotaApplication(ctx *pulumi.Context,
	name string, args *QuotaApplicationArgs, opts ...pulumi.ResourceOption) (*QuotaApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DesireValue == nil {
		return nil, errors.New("invalid value for required argument 'DesireValue'")
	}
	if args.ProductCode == nil {
		return nil, errors.New("invalid value for required argument 'ProductCode'")
	}
	if args.QuotaActionCode == nil {
		return nil, errors.New("invalid value for required argument 'QuotaActionCode'")
	}
	if args.Reason == nil {
		return nil, errors.New("invalid value for required argument 'Reason'")
	}
	var resource QuotaApplication
	err := ctx.RegisterResource("alicloud:quotas/quotaApplication:QuotaApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuotaApplication gets an existing QuotaApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuotaApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QuotaApplicationState, opts ...pulumi.ResourceOption) (*QuotaApplication, error) {
	var resource QuotaApplication
	err := ctx.ReadResource("alicloud:quotas/quotaApplication:QuotaApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QuotaApplication resources.
type quotaApplicationState struct {
	// The approve value of the quota application.
	ApproveValue *string `pulumi:"approveValue"`
	// The audit mode. Valid values: `Async`, `Sync`. Default to: `Async`.
	AuditMode *string `pulumi:"auditMode"`
	// The audit reason.
	AuditReason *string `pulumi:"auditReason"`
	// The desire value of the quota application.
	DesireValue *float64 `pulumi:"desireValue"`
	// The quota dimensions.
	Dimensions []QuotaApplicationDimension `pulumi:"dimensions"`
	// The effective time of the quota application.
	EffectiveTime *string `pulumi:"effectiveTime"`
	// The expire time of the quota application.
	ExpireTime *string `pulumi:"expireTime"`
	// The notice type. Valid values: `0`, `1`, `2`, `3`.
	NoticeType *int `pulumi:"noticeType"`
	// The product code.
	ProductCode *string `pulumi:"productCode"`
	// The ID of quota action.
	QuotaActionCode *string `pulumi:"quotaActionCode"`
	// The quota category. Valid values: `CommonQuota`, `FlowControl`.
	QuotaCategory *string `pulumi:"quotaCategory"`
	// The description of the quota application.
	QuotaDescription *string `pulumi:"quotaDescription"`
	// The name of the quota application.
	QuotaName *string `pulumi:"quotaName"`
	// The unit of the quota application.
	QuotaUnit *string `pulumi:"quotaUnit"`
	// The reason of the quota application.
	Reason *string `pulumi:"reason"`
	// The status of the quota application.
	Status *string `pulumi:"status"`
}

type QuotaApplicationState struct {
	// The approve value of the quota application.
	ApproveValue pulumi.StringPtrInput
	// The audit mode. Valid values: `Async`, `Sync`. Default to: `Async`.
	AuditMode pulumi.StringPtrInput
	// The audit reason.
	AuditReason pulumi.StringPtrInput
	// The desire value of the quota application.
	DesireValue pulumi.Float64PtrInput
	// The quota dimensions.
	Dimensions QuotaApplicationDimensionArrayInput
	// The effective time of the quota application.
	EffectiveTime pulumi.StringPtrInput
	// The expire time of the quota application.
	ExpireTime pulumi.StringPtrInput
	// The notice type. Valid values: `0`, `1`, `2`, `3`.
	NoticeType pulumi.IntPtrInput
	// The product code.
	ProductCode pulumi.StringPtrInput
	// The ID of quota action.
	QuotaActionCode pulumi.StringPtrInput
	// The quota category. Valid values: `CommonQuota`, `FlowControl`.
	QuotaCategory pulumi.StringPtrInput
	// The description of the quota application.
	QuotaDescription pulumi.StringPtrInput
	// The name of the quota application.
	QuotaName pulumi.StringPtrInput
	// The unit of the quota application.
	QuotaUnit pulumi.StringPtrInput
	// The reason of the quota application.
	Reason pulumi.StringPtrInput
	// The status of the quota application.
	Status pulumi.StringPtrInput
}

func (QuotaApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaApplicationState)(nil)).Elem()
}

type quotaApplicationArgs struct {
	// The audit mode. Valid values: `Async`, `Sync`. Default to: `Async`.
	AuditMode *string `pulumi:"auditMode"`
	// The desire value of the quota application.
	DesireValue float64 `pulumi:"desireValue"`
	// The quota dimensions.
	Dimensions []QuotaApplicationDimension `pulumi:"dimensions"`
	// The notice type. Valid values: `0`, `1`, `2`, `3`.
	NoticeType *int `pulumi:"noticeType"`
	// The product code.
	ProductCode string `pulumi:"productCode"`
	// The ID of quota action.
	QuotaActionCode string `pulumi:"quotaActionCode"`
	// The quota category. Valid values: `CommonQuota`, `FlowControl`.
	QuotaCategory *string `pulumi:"quotaCategory"`
	// The reason of the quota application.
	Reason string `pulumi:"reason"`
}

// The set of arguments for constructing a QuotaApplication resource.
type QuotaApplicationArgs struct {
	// The audit mode. Valid values: `Async`, `Sync`. Default to: `Async`.
	AuditMode pulumi.StringPtrInput
	// The desire value of the quota application.
	DesireValue pulumi.Float64Input
	// The quota dimensions.
	Dimensions QuotaApplicationDimensionArrayInput
	// The notice type. Valid values: `0`, `1`, `2`, `3`.
	NoticeType pulumi.IntPtrInput
	// The product code.
	ProductCode pulumi.StringInput
	// The ID of quota action.
	QuotaActionCode pulumi.StringInput
	// The quota category. Valid values: `CommonQuota`, `FlowControl`.
	QuotaCategory pulumi.StringPtrInput
	// The reason of the quota application.
	Reason pulumi.StringInput
}

func (QuotaApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaApplicationArgs)(nil)).Elem()
}

type QuotaApplicationInput interface {
	pulumi.Input

	ToQuotaApplicationOutput() QuotaApplicationOutput
	ToQuotaApplicationOutputWithContext(ctx context.Context) QuotaApplicationOutput
}

func (*QuotaApplication) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaApplication)(nil)).Elem()
}

func (i *QuotaApplication) ToQuotaApplicationOutput() QuotaApplicationOutput {
	return i.ToQuotaApplicationOutputWithContext(context.Background())
}

func (i *QuotaApplication) ToQuotaApplicationOutputWithContext(ctx context.Context) QuotaApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaApplicationOutput)
}

// QuotaApplicationArrayInput is an input type that accepts QuotaApplicationArray and QuotaApplicationArrayOutput values.
// You can construct a concrete instance of `QuotaApplicationArrayInput` via:
//
//	QuotaApplicationArray{ QuotaApplicationArgs{...} }
type QuotaApplicationArrayInput interface {
	pulumi.Input

	ToQuotaApplicationArrayOutput() QuotaApplicationArrayOutput
	ToQuotaApplicationArrayOutputWithContext(context.Context) QuotaApplicationArrayOutput
}

type QuotaApplicationArray []QuotaApplicationInput

func (QuotaApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QuotaApplication)(nil)).Elem()
}

func (i QuotaApplicationArray) ToQuotaApplicationArrayOutput() QuotaApplicationArrayOutput {
	return i.ToQuotaApplicationArrayOutputWithContext(context.Background())
}

func (i QuotaApplicationArray) ToQuotaApplicationArrayOutputWithContext(ctx context.Context) QuotaApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaApplicationArrayOutput)
}

// QuotaApplicationMapInput is an input type that accepts QuotaApplicationMap and QuotaApplicationMapOutput values.
// You can construct a concrete instance of `QuotaApplicationMapInput` via:
//
//	QuotaApplicationMap{ "key": QuotaApplicationArgs{...} }
type QuotaApplicationMapInput interface {
	pulumi.Input

	ToQuotaApplicationMapOutput() QuotaApplicationMapOutput
	ToQuotaApplicationMapOutputWithContext(context.Context) QuotaApplicationMapOutput
}

type QuotaApplicationMap map[string]QuotaApplicationInput

func (QuotaApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QuotaApplication)(nil)).Elem()
}

func (i QuotaApplicationMap) ToQuotaApplicationMapOutput() QuotaApplicationMapOutput {
	return i.ToQuotaApplicationMapOutputWithContext(context.Background())
}

func (i QuotaApplicationMap) ToQuotaApplicationMapOutputWithContext(ctx context.Context) QuotaApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaApplicationMapOutput)
}

type QuotaApplicationOutput struct{ *pulumi.OutputState }

func (QuotaApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaApplication)(nil)).Elem()
}

func (o QuotaApplicationOutput) ToQuotaApplicationOutput() QuotaApplicationOutput {
	return o
}

func (o QuotaApplicationOutput) ToQuotaApplicationOutputWithContext(ctx context.Context) QuotaApplicationOutput {
	return o
}

// The approve value of the quota application.
func (o QuotaApplicationOutput) ApproveValue() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaApplication) pulumi.StringOutput { return v.ApproveValue }).(pulumi.StringOutput)
}

// The audit mode. Valid values: `Async`, `Sync`. Default to: `Async`.
func (o QuotaApplicationOutput) AuditMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuotaApplication) pulumi.StringPtrOutput { return v.AuditMode }).(pulumi.StringPtrOutput)
}

// The audit reason.
func (o QuotaApplicationOutput) AuditReason() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaApplication) pulumi.StringOutput { return v.AuditReason }).(pulumi.StringOutput)
}

// The desire value of the quota application.
func (o QuotaApplicationOutput) DesireValue() pulumi.Float64Output {
	return o.ApplyT(func(v *QuotaApplication) pulumi.Float64Output { return v.DesireValue }).(pulumi.Float64Output)
}

// The quota dimensions.
func (o QuotaApplicationOutput) Dimensions() QuotaApplicationDimensionArrayOutput {
	return o.ApplyT(func(v *QuotaApplication) QuotaApplicationDimensionArrayOutput { return v.Dimensions }).(QuotaApplicationDimensionArrayOutput)
}

// The effective time of the quota application.
func (o QuotaApplicationOutput) EffectiveTime() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaApplication) pulumi.StringOutput { return v.EffectiveTime }).(pulumi.StringOutput)
}

// The expire time of the quota application.
func (o QuotaApplicationOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaApplication) pulumi.StringOutput { return v.ExpireTime }).(pulumi.StringOutput)
}

// The notice type. Valid values: `0`, `1`, `2`, `3`.
func (o QuotaApplicationOutput) NoticeType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QuotaApplication) pulumi.IntPtrOutput { return v.NoticeType }).(pulumi.IntPtrOutput)
}

// The product code.
func (o QuotaApplicationOutput) ProductCode() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaApplication) pulumi.StringOutput { return v.ProductCode }).(pulumi.StringOutput)
}

// The ID of quota action.
func (o QuotaApplicationOutput) QuotaActionCode() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaApplication) pulumi.StringOutput { return v.QuotaActionCode }).(pulumi.StringOutput)
}

// The quota category. Valid values: `CommonQuota`, `FlowControl`.
func (o QuotaApplicationOutput) QuotaCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuotaApplication) pulumi.StringPtrOutput { return v.QuotaCategory }).(pulumi.StringPtrOutput)
}

// The description of the quota application.
func (o QuotaApplicationOutput) QuotaDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaApplication) pulumi.StringOutput { return v.QuotaDescription }).(pulumi.StringOutput)
}

// The name of the quota application.
func (o QuotaApplicationOutput) QuotaName() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaApplication) pulumi.StringOutput { return v.QuotaName }).(pulumi.StringOutput)
}

// The unit of the quota application.
func (o QuotaApplicationOutput) QuotaUnit() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaApplication) pulumi.StringOutput { return v.QuotaUnit }).(pulumi.StringOutput)
}

// The reason of the quota application.
func (o QuotaApplicationOutput) Reason() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaApplication) pulumi.StringOutput { return v.Reason }).(pulumi.StringOutput)
}

// The status of the quota application.
func (o QuotaApplicationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaApplication) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type QuotaApplicationArrayOutput struct{ *pulumi.OutputState }

func (QuotaApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QuotaApplication)(nil)).Elem()
}

func (o QuotaApplicationArrayOutput) ToQuotaApplicationArrayOutput() QuotaApplicationArrayOutput {
	return o
}

func (o QuotaApplicationArrayOutput) ToQuotaApplicationArrayOutputWithContext(ctx context.Context) QuotaApplicationArrayOutput {
	return o
}

func (o QuotaApplicationArrayOutput) Index(i pulumi.IntInput) QuotaApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QuotaApplication {
		return vs[0].([]*QuotaApplication)[vs[1].(int)]
	}).(QuotaApplicationOutput)
}

type QuotaApplicationMapOutput struct{ *pulumi.OutputState }

func (QuotaApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QuotaApplication)(nil)).Elem()
}

func (o QuotaApplicationMapOutput) ToQuotaApplicationMapOutput() QuotaApplicationMapOutput {
	return o
}

func (o QuotaApplicationMapOutput) ToQuotaApplicationMapOutputWithContext(ctx context.Context) QuotaApplicationMapOutput {
	return o
}

func (o QuotaApplicationMapOutput) MapIndex(k pulumi.StringInput) QuotaApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QuotaApplication {
		return vs[0].(map[string]*QuotaApplication)[vs[1].(string)]
	}).(QuotaApplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QuotaApplicationInput)(nil)).Elem(), &QuotaApplication{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuotaApplicationArrayInput)(nil)).Elem(), QuotaApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuotaApplicationMapInput)(nil)).Elem(), QuotaApplicationMap{})
	pulumi.RegisterOutputType(QuotaApplicationOutput{})
	pulumi.RegisterOutputType(QuotaApplicationArrayOutput{})
	pulumi.RegisterOutputType(QuotaApplicationMapOutput{})
}
