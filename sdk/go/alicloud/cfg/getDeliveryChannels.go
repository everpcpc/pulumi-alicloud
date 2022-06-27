// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cfg"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := cfg.GetDeliveryChannels(ctx, &cfg.GetDeliveryChannelsArgs{
// 			Ids: []string{
// 				"cdc-49a2ad756057********",
// 			},
// 			NameRegex: pulumi.StringRef("tftest"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstConfigDeliveryChannelId", example.Channels[0].Id)
// 		return nil
// 	})
// }
// ```
func GetDeliveryChannels(ctx *pulumi.Context, args *GetDeliveryChannelsArgs, opts ...pulumi.InvokeOption) (*GetDeliveryChannelsResult, error) {
	var rv GetDeliveryChannelsResult
	err := ctx.Invoke("alicloud:cfg/getDeliveryChannels:getDeliveryChannels", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDeliveryChannels.
type GetDeliveryChannelsArgs struct {
	// A list of Config Delivery Channel IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by delivery channel name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The status of the config delivery channel. Valid values `0`: Disable delivery channel, `1`: Enable delivery channel.
	Status *int `pulumi:"status"`
}

// A collection of values returned by getDeliveryChannels.
type GetDeliveryChannelsResult struct {
	// A list of Config Delivery Channels. Each element contains the following attributes:
	Channels []GetDeliveryChannelsChannel `pulumi:"channels"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Config Delivery Channel IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of Config Delivery Channel names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// The status of the delivery channel.
	Status *int `pulumi:"status"`
}

func GetDeliveryChannelsOutput(ctx *pulumi.Context, args GetDeliveryChannelsOutputArgs, opts ...pulumi.InvokeOption) GetDeliveryChannelsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDeliveryChannelsResult, error) {
			args := v.(GetDeliveryChannelsArgs)
			r, err := GetDeliveryChannels(ctx, &args, opts...)
			var s GetDeliveryChannelsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDeliveryChannelsResultOutput)
}

// A collection of arguments for invoking getDeliveryChannels.
type GetDeliveryChannelsOutputArgs struct {
	// A list of Config Delivery Channel IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by delivery channel name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the config delivery channel. Valid values `0`: Disable delivery channel, `1`: Enable delivery channel.
	Status pulumi.IntPtrInput `pulumi:"status"`
}

func (GetDeliveryChannelsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeliveryChannelsArgs)(nil)).Elem()
}

// A collection of values returned by getDeliveryChannels.
type GetDeliveryChannelsResultOutput struct{ *pulumi.OutputState }

func (GetDeliveryChannelsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeliveryChannelsResult)(nil)).Elem()
}

func (o GetDeliveryChannelsResultOutput) ToGetDeliveryChannelsResultOutput() GetDeliveryChannelsResultOutput {
	return o
}

func (o GetDeliveryChannelsResultOutput) ToGetDeliveryChannelsResultOutputWithContext(ctx context.Context) GetDeliveryChannelsResultOutput {
	return o
}

// A list of Config Delivery Channels. Each element contains the following attributes:
func (o GetDeliveryChannelsResultOutput) Channels() GetDeliveryChannelsChannelArrayOutput {
	return o.ApplyT(func(v GetDeliveryChannelsResult) []GetDeliveryChannelsChannel { return v.Channels }).(GetDeliveryChannelsChannelArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDeliveryChannelsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeliveryChannelsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of Config Delivery Channel IDs.
func (o GetDeliveryChannelsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDeliveryChannelsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetDeliveryChannelsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeliveryChannelsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of Config Delivery Channel names.
func (o GetDeliveryChannelsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDeliveryChannelsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetDeliveryChannelsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeliveryChannelsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The status of the delivery channel.
func (o GetDeliveryChannelsResultOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDeliveryChannelsResult) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDeliveryChannelsResultOutput{})
}
