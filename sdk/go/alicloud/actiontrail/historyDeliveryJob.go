// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package actiontrail

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Actiontrail History Delivery Job resource.
//
// For information about Actiontrail History Delivery Job and how to use it, see [What is History Delivery Job](https://www.alibabacloud.com/help/doc-detail/199999.htm).
//
// > **NOTE:** Available in v1.139.0+.
//
// > **NOTE:** Make sure that you have called the `actiontrail.Trail` to create a single account trace that is delivered to Log Service SLS. An Alibaba cloud account can only have one running delivery history job at the same time.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/actiontrail"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/log"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := true
// 		exampleRegions, err := alicloud.GetRegions(ctx, &GetRegionsArgs{
// 			Current: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := alicloud.GetAccount(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleProject, err := log.NewProject(ctx, "exampleProject", &log.ProjectArgs{
// 			Description: pulumi.String("tf actiontrail test"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := "AliyunActionTrailDefaultRole"
// 		exampleRoles, err := ram.GetRoles(ctx, &ram.GetRolesArgs{
// 			NameRegex: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleTrail, err := actiontrail.NewTrail(ctx, "exampleTrail", &actiontrail.TrailArgs{
// 			TrailName:       pulumi.String("example_value"),
// 			SlsWriteRoleArn: pulumi.String(exampleRoles.Roles[0].Arn),
// 			SlsProjectArn: exampleProject.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v%v%v%v%v", "acs:log:", exampleRegions.Regions[0].Id, ":", exampleAccount.Id, ":project/", name), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = actiontrail.NewHistoryDeliveryJob(ctx, "exampleHistoryDeliveryJob", &actiontrail.HistoryDeliveryJobArgs{
// 			TrailName: exampleTrail.ID(),
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
// Actiontrail History Delivery Job can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:actiontrail/historyDeliveryJob:HistoryDeliveryJob example <id>
// ```
type HistoryDeliveryJob struct {
	pulumi.CustomResourceState

	// The status of the task. Valid values: `0`, `1`, `2`, `3`. `0`: The task is initializing. `1`: The task is delivering historical events. `2`: The delivery of historical events is complete. `3`: The task fails.
	Status pulumi.IntOutput `pulumi:"status"`
	// The name of the trail for which you want to create a historical event delivery task.
	TrailName pulumi.StringOutput `pulumi:"trailName"`
}

// NewHistoryDeliveryJob registers a new resource with the given unique name, arguments, and options.
func NewHistoryDeliveryJob(ctx *pulumi.Context,
	name string, args *HistoryDeliveryJobArgs, opts ...pulumi.ResourceOption) (*HistoryDeliveryJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TrailName == nil {
		return nil, errors.New("invalid value for required argument 'TrailName'")
	}
	var resource HistoryDeliveryJob
	err := ctx.RegisterResource("alicloud:actiontrail/historyDeliveryJob:HistoryDeliveryJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHistoryDeliveryJob gets an existing HistoryDeliveryJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHistoryDeliveryJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HistoryDeliveryJobState, opts ...pulumi.ResourceOption) (*HistoryDeliveryJob, error) {
	var resource HistoryDeliveryJob
	err := ctx.ReadResource("alicloud:actiontrail/historyDeliveryJob:HistoryDeliveryJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HistoryDeliveryJob resources.
type historyDeliveryJobState struct {
	// The status of the task. Valid values: `0`, `1`, `2`, `3`. `0`: The task is initializing. `1`: The task is delivering historical events. `2`: The delivery of historical events is complete. `3`: The task fails.
	Status *int `pulumi:"status"`
	// The name of the trail for which you want to create a historical event delivery task.
	TrailName *string `pulumi:"trailName"`
}

type HistoryDeliveryJobState struct {
	// The status of the task. Valid values: `0`, `1`, `2`, `3`. `0`: The task is initializing. `1`: The task is delivering historical events. `2`: The delivery of historical events is complete. `3`: The task fails.
	Status pulumi.IntPtrInput
	// The name of the trail for which you want to create a historical event delivery task.
	TrailName pulumi.StringPtrInput
}

func (HistoryDeliveryJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*historyDeliveryJobState)(nil)).Elem()
}

type historyDeliveryJobArgs struct {
	// The name of the trail for which you want to create a historical event delivery task.
	TrailName string `pulumi:"trailName"`
}

// The set of arguments for constructing a HistoryDeliveryJob resource.
type HistoryDeliveryJobArgs struct {
	// The name of the trail for which you want to create a historical event delivery task.
	TrailName pulumi.StringInput
}

func (HistoryDeliveryJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*historyDeliveryJobArgs)(nil)).Elem()
}

type HistoryDeliveryJobInput interface {
	pulumi.Input

	ToHistoryDeliveryJobOutput() HistoryDeliveryJobOutput
	ToHistoryDeliveryJobOutputWithContext(ctx context.Context) HistoryDeliveryJobOutput
}

func (*HistoryDeliveryJob) ElementType() reflect.Type {
	return reflect.TypeOf((*HistoryDeliveryJob)(nil))
}

func (i *HistoryDeliveryJob) ToHistoryDeliveryJobOutput() HistoryDeliveryJobOutput {
	return i.ToHistoryDeliveryJobOutputWithContext(context.Background())
}

func (i *HistoryDeliveryJob) ToHistoryDeliveryJobOutputWithContext(ctx context.Context) HistoryDeliveryJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HistoryDeliveryJobOutput)
}

func (i *HistoryDeliveryJob) ToHistoryDeliveryJobPtrOutput() HistoryDeliveryJobPtrOutput {
	return i.ToHistoryDeliveryJobPtrOutputWithContext(context.Background())
}

func (i *HistoryDeliveryJob) ToHistoryDeliveryJobPtrOutputWithContext(ctx context.Context) HistoryDeliveryJobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HistoryDeliveryJobPtrOutput)
}

type HistoryDeliveryJobPtrInput interface {
	pulumi.Input

	ToHistoryDeliveryJobPtrOutput() HistoryDeliveryJobPtrOutput
	ToHistoryDeliveryJobPtrOutputWithContext(ctx context.Context) HistoryDeliveryJobPtrOutput
}

type historyDeliveryJobPtrType HistoryDeliveryJobArgs

func (*historyDeliveryJobPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HistoryDeliveryJob)(nil))
}

func (i *historyDeliveryJobPtrType) ToHistoryDeliveryJobPtrOutput() HistoryDeliveryJobPtrOutput {
	return i.ToHistoryDeliveryJobPtrOutputWithContext(context.Background())
}

func (i *historyDeliveryJobPtrType) ToHistoryDeliveryJobPtrOutputWithContext(ctx context.Context) HistoryDeliveryJobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HistoryDeliveryJobPtrOutput)
}

// HistoryDeliveryJobArrayInput is an input type that accepts HistoryDeliveryJobArray and HistoryDeliveryJobArrayOutput values.
// You can construct a concrete instance of `HistoryDeliveryJobArrayInput` via:
//
//          HistoryDeliveryJobArray{ HistoryDeliveryJobArgs{...} }
type HistoryDeliveryJobArrayInput interface {
	pulumi.Input

	ToHistoryDeliveryJobArrayOutput() HistoryDeliveryJobArrayOutput
	ToHistoryDeliveryJobArrayOutputWithContext(context.Context) HistoryDeliveryJobArrayOutput
}

type HistoryDeliveryJobArray []HistoryDeliveryJobInput

func (HistoryDeliveryJobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HistoryDeliveryJob)(nil)).Elem()
}

func (i HistoryDeliveryJobArray) ToHistoryDeliveryJobArrayOutput() HistoryDeliveryJobArrayOutput {
	return i.ToHistoryDeliveryJobArrayOutputWithContext(context.Background())
}

func (i HistoryDeliveryJobArray) ToHistoryDeliveryJobArrayOutputWithContext(ctx context.Context) HistoryDeliveryJobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HistoryDeliveryJobArrayOutput)
}

// HistoryDeliveryJobMapInput is an input type that accepts HistoryDeliveryJobMap and HistoryDeliveryJobMapOutput values.
// You can construct a concrete instance of `HistoryDeliveryJobMapInput` via:
//
//          HistoryDeliveryJobMap{ "key": HistoryDeliveryJobArgs{...} }
type HistoryDeliveryJobMapInput interface {
	pulumi.Input

	ToHistoryDeliveryJobMapOutput() HistoryDeliveryJobMapOutput
	ToHistoryDeliveryJobMapOutputWithContext(context.Context) HistoryDeliveryJobMapOutput
}

type HistoryDeliveryJobMap map[string]HistoryDeliveryJobInput

func (HistoryDeliveryJobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HistoryDeliveryJob)(nil)).Elem()
}

func (i HistoryDeliveryJobMap) ToHistoryDeliveryJobMapOutput() HistoryDeliveryJobMapOutput {
	return i.ToHistoryDeliveryJobMapOutputWithContext(context.Background())
}

func (i HistoryDeliveryJobMap) ToHistoryDeliveryJobMapOutputWithContext(ctx context.Context) HistoryDeliveryJobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HistoryDeliveryJobMapOutput)
}

type HistoryDeliveryJobOutput struct{ *pulumi.OutputState }

func (HistoryDeliveryJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HistoryDeliveryJob)(nil))
}

func (o HistoryDeliveryJobOutput) ToHistoryDeliveryJobOutput() HistoryDeliveryJobOutput {
	return o
}

func (o HistoryDeliveryJobOutput) ToHistoryDeliveryJobOutputWithContext(ctx context.Context) HistoryDeliveryJobOutput {
	return o
}

func (o HistoryDeliveryJobOutput) ToHistoryDeliveryJobPtrOutput() HistoryDeliveryJobPtrOutput {
	return o.ToHistoryDeliveryJobPtrOutputWithContext(context.Background())
}

func (o HistoryDeliveryJobOutput) ToHistoryDeliveryJobPtrOutputWithContext(ctx context.Context) HistoryDeliveryJobPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HistoryDeliveryJob) *HistoryDeliveryJob {
		return &v
	}).(HistoryDeliveryJobPtrOutput)
}

type HistoryDeliveryJobPtrOutput struct{ *pulumi.OutputState }

func (HistoryDeliveryJobPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HistoryDeliveryJob)(nil))
}

func (o HistoryDeliveryJobPtrOutput) ToHistoryDeliveryJobPtrOutput() HistoryDeliveryJobPtrOutput {
	return o
}

func (o HistoryDeliveryJobPtrOutput) ToHistoryDeliveryJobPtrOutputWithContext(ctx context.Context) HistoryDeliveryJobPtrOutput {
	return o
}

func (o HistoryDeliveryJobPtrOutput) Elem() HistoryDeliveryJobOutput {
	return o.ApplyT(func(v *HistoryDeliveryJob) HistoryDeliveryJob {
		if v != nil {
			return *v
		}
		var ret HistoryDeliveryJob
		return ret
	}).(HistoryDeliveryJobOutput)
}

type HistoryDeliveryJobArrayOutput struct{ *pulumi.OutputState }

func (HistoryDeliveryJobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HistoryDeliveryJob)(nil))
}

func (o HistoryDeliveryJobArrayOutput) ToHistoryDeliveryJobArrayOutput() HistoryDeliveryJobArrayOutput {
	return o
}

func (o HistoryDeliveryJobArrayOutput) ToHistoryDeliveryJobArrayOutputWithContext(ctx context.Context) HistoryDeliveryJobArrayOutput {
	return o
}

func (o HistoryDeliveryJobArrayOutput) Index(i pulumi.IntInput) HistoryDeliveryJobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HistoryDeliveryJob {
		return vs[0].([]HistoryDeliveryJob)[vs[1].(int)]
	}).(HistoryDeliveryJobOutput)
}

type HistoryDeliveryJobMapOutput struct{ *pulumi.OutputState }

func (HistoryDeliveryJobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]HistoryDeliveryJob)(nil))
}

func (o HistoryDeliveryJobMapOutput) ToHistoryDeliveryJobMapOutput() HistoryDeliveryJobMapOutput {
	return o
}

func (o HistoryDeliveryJobMapOutput) ToHistoryDeliveryJobMapOutputWithContext(ctx context.Context) HistoryDeliveryJobMapOutput {
	return o
}

func (o HistoryDeliveryJobMapOutput) MapIndex(k pulumi.StringInput) HistoryDeliveryJobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) HistoryDeliveryJob {
		return vs[0].(map[string]HistoryDeliveryJob)[vs[1].(string)]
	}).(HistoryDeliveryJobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HistoryDeliveryJobInput)(nil)).Elem(), &HistoryDeliveryJob{})
	pulumi.RegisterInputType(reflect.TypeOf((*HistoryDeliveryJobPtrInput)(nil)).Elem(), &HistoryDeliveryJob{})
	pulumi.RegisterInputType(reflect.TypeOf((*HistoryDeliveryJobArrayInput)(nil)).Elem(), HistoryDeliveryJobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HistoryDeliveryJobMapInput)(nil)).Elem(), HistoryDeliveryJobMap{})
	pulumi.RegisterOutputType(HistoryDeliveryJobOutput{})
	pulumi.RegisterOutputType(HistoryDeliveryJobPtrOutput{})
	pulumi.RegisterOutputType(HistoryDeliveryJobArrayOutput{})
	pulumi.RegisterOutputType(HistoryDeliveryJobMapOutput{})
}
