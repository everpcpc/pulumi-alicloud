// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ess

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides available notification resources.
//
// > **NOTE:** Available in 1.72.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ess"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ds, err := ess.GetNotifications(ctx, &ess.GetNotificationsArgs{
// 			ScalingGroupId: "scaling_group_id",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstNotification", ds.Notifications[0].Id)
// 		return nil
// 	})
// }
// ```
func GetNotifications(ctx *pulumi.Context, args *GetNotificationsArgs, opts ...pulumi.InvokeOption) (*GetNotificationsResult, error) {
	var rv GetNotificationsResult
	err := ctx.Invoke("alicloud:ess/getNotifications:getNotifications", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNotifications.
type GetNotificationsArgs struct {
	// A list of notification ids.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// Scaling group id the notifications belong to.
	ScalingGroupId string `pulumi:"scalingGroupId"`
}

// A collection of values returned by getNotifications.
type GetNotificationsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of notification ids.
	Ids []string `pulumi:"ids"`
	// A list of notifications. Each element contains the following attributes:
	Notifications []GetNotificationsNotification `pulumi:"notifications"`
	OutputFile    *string                        `pulumi:"outputFile"`
	// ID of the scaling group.
	ScalingGroupId string `pulumi:"scalingGroupId"`
}

func GetNotificationsOutput(ctx *pulumi.Context, args GetNotificationsOutputArgs, opts ...pulumi.InvokeOption) GetNotificationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNotificationsResult, error) {
			args := v.(GetNotificationsArgs)
			r, err := GetNotifications(ctx, &args, opts...)
			return *r, err
		}).(GetNotificationsResultOutput)
}

// A collection of arguments for invoking getNotifications.
type GetNotificationsOutputArgs struct {
	// A list of notification ids.
	Ids        pulumi.StringArrayInput `pulumi:"ids"`
	OutputFile pulumi.StringPtrInput   `pulumi:"outputFile"`
	// Scaling group id the notifications belong to.
	ScalingGroupId pulumi.StringInput `pulumi:"scalingGroupId"`
}

func (GetNotificationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNotificationsArgs)(nil)).Elem()
}

// A collection of values returned by getNotifications.
type GetNotificationsResultOutput struct{ *pulumi.OutputState }

func (GetNotificationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNotificationsResult)(nil)).Elem()
}

func (o GetNotificationsResultOutput) ToGetNotificationsResultOutput() GetNotificationsResultOutput {
	return o
}

func (o GetNotificationsResultOutput) ToGetNotificationsResultOutputWithContext(ctx context.Context) GetNotificationsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetNotificationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of notification ids.
func (o GetNotificationsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNotificationsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A list of notifications. Each element contains the following attributes:
func (o GetNotificationsResultOutput) Notifications() GetNotificationsNotificationArrayOutput {
	return o.ApplyT(func(v GetNotificationsResult) []GetNotificationsNotification { return v.Notifications }).(GetNotificationsNotificationArrayOutput)
}

func (o GetNotificationsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNotificationsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// ID of the scaling group.
func (o GetNotificationsResultOutput) ScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationsResult) string { return v.ScalingGroupId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNotificationsResultOutput{})
}
