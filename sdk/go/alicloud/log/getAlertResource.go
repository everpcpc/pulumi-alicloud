// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Using this data source can init SLS Alert resources automatically.
//
// For information about SLS Alert and how to use it, see [SLS Alert Overview](https://www.alibabacloud.com/help/en/doc-detail/209202.html)
//
// > **NOTE:** Available in v1.161.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/log"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := log.GetAlertResource(ctx, &log.GetAlertResourceArgs{
// 			Lang: pulumi.StringRef("cn"),
// 			Type: "user",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = log.GetAlertResource(ctx, &log.GetAlertResourceArgs{
// 			Project: pulumi.StringRef("test-alert-tf"),
// 			Type:    "project",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetAlertResource(ctx *pulumi.Context, args *GetAlertResourceArgs, opts ...pulumi.InvokeOption) (*GetAlertResourceResult, error) {
	var rv GetAlertResourceResult
	err := ctx.Invoke("alicloud:log/getAlertResource:getAlertResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlertResource.
type GetAlertResourceArgs struct {
	// The lang of alert center resource when type is user.
	Lang *string `pulumi:"lang"`
	// The project of alert resource when type is project.
	Project *string `pulumi:"project"`
	// The type of alert resources, must be user or project, 'user' for init aliyuncloud account's alert center resource, including project named sls-alert-{uid}-{region} and some dashboards; 'project' for init project's alert resource, including logstore named internal-alert-history and alert dashboard.
	Type string `pulumi:"type"`
}

// A collection of values returned by getAlertResource.
type GetAlertResourceResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	Lang    *string `pulumi:"lang"`
	Project *string `pulumi:"project"`
	Type    string  `pulumi:"type"`
}

func GetAlertResourceOutput(ctx *pulumi.Context, args GetAlertResourceOutputArgs, opts ...pulumi.InvokeOption) GetAlertResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAlertResourceResult, error) {
			args := v.(GetAlertResourceArgs)
			r, err := GetAlertResource(ctx, &args, opts...)
			var s GetAlertResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAlertResourceResultOutput)
}

// A collection of arguments for invoking getAlertResource.
type GetAlertResourceOutputArgs struct {
	// The lang of alert center resource when type is user.
	Lang pulumi.StringPtrInput `pulumi:"lang"`
	// The project of alert resource when type is project.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// The type of alert resources, must be user or project, 'user' for init aliyuncloud account's alert center resource, including project named sls-alert-{uid}-{region} and some dashboards; 'project' for init project's alert resource, including logstore named internal-alert-history and alert dashboard.
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetAlertResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlertResourceArgs)(nil)).Elem()
}

// A collection of values returned by getAlertResource.
type GetAlertResourceResultOutput struct{ *pulumi.OutputState }

func (GetAlertResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlertResourceResult)(nil)).Elem()
}

func (o GetAlertResourceResultOutput) ToGetAlertResourceResultOutput() GetAlertResourceResultOutput {
	return o
}

func (o GetAlertResourceResultOutput) ToGetAlertResourceResultOutputWithContext(ctx context.Context) GetAlertResourceResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetAlertResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlertResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAlertResourceResultOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlertResourceResult) *string { return v.Lang }).(pulumi.StringPtrOutput)
}

func (o GetAlertResourceResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlertResourceResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o GetAlertResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlertResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAlertResourceResultOutput{})
}
