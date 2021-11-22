// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource used to create a flow log function in Cloud Enterprise Network (CEN).
// By using the flow log function, you can capture the traffic data of the network instances in different regions of a CEN.
// You can also use the data aggregated in flow logs to analyze cross-region traffic flows, minimize traffic costs, and troubleshoot network faults.
//
// For information about CEN flow log and how to use it, see [Manage CEN flowlog](https://www.alibabacloud.com/help/doc-detail/123006.htm).
//
// > **NOTE:** Available in 1.73.0+
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/log"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultInstance, err := cen.NewInstance(ctx, "defaultInstance", nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultProject, err := log.NewProject(ctx, "defaultProject", &log.ProjectArgs{
// 			Description: pulumi.String("create by terraform"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultStore, err := log.NewStore(ctx, "defaultStore", &log.StoreArgs{
// 			Project:            defaultProject.Name,
// 			RetentionPeriod:    pulumi.Int(3650),
// 			ShardCount:         pulumi.Int(3),
// 			AutoSplit:          pulumi.Bool(true),
// 			MaxSplitShardCount: pulumi.Int(60),
// 			AppendMeta:         pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cen.NewFlowLog(ctx, "defaultFlowLog", &cen.FlowLogArgs{
// 			FlowLogName:  pulumi.String("my-flowlog"),
// 			CenId:        defaultInstance.ID(),
// 			ProjectName:  defaultProject.Name,
// 			LogStoreName: defaultStore.Name,
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
// CEN flowlog can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:cen/flowLog:FlowLog default flowlog-tig1xxxxxx
// ```
type FlowLog struct {
	pulumi.CustomResourceState

	// The ID of the CEN Instance.
	CenId pulumi.StringOutput `pulumi:"cenId"`
	// The description of flowlog.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of flowlog.
	FlowLogName pulumi.StringPtrOutput `pulumi:"flowLogName"`
	// The name of the log store which is in the  `projectName` SLS project.
	LogStoreName pulumi.StringOutput `pulumi:"logStoreName"`
	// The name of the SLS project.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// The status of flowlog. Valid values: ["Active", "Inactive"]. Default to "Active".
	Status pulumi.StringPtrOutput `pulumi:"status"`
}

// NewFlowLog registers a new resource with the given unique name, arguments, and options.
func NewFlowLog(ctx *pulumi.Context,
	name string, args *FlowLogArgs, opts ...pulumi.ResourceOption) (*FlowLog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CenId == nil {
		return nil, errors.New("invalid value for required argument 'CenId'")
	}
	if args.LogStoreName == nil {
		return nil, errors.New("invalid value for required argument 'LogStoreName'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	var resource FlowLog
	err := ctx.RegisterResource("alicloud:cen/flowLog:FlowLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlowLog gets an existing FlowLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlowLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowLogState, opts ...pulumi.ResourceOption) (*FlowLog, error) {
	var resource FlowLog
	err := ctx.ReadResource("alicloud:cen/flowLog:FlowLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlowLog resources.
type flowLogState struct {
	// The ID of the CEN Instance.
	CenId *string `pulumi:"cenId"`
	// The description of flowlog.
	Description *string `pulumi:"description"`
	// The name of flowlog.
	FlowLogName *string `pulumi:"flowLogName"`
	// The name of the log store which is in the  `projectName` SLS project.
	LogStoreName *string `pulumi:"logStoreName"`
	// The name of the SLS project.
	ProjectName *string `pulumi:"projectName"`
	// The status of flowlog. Valid values: ["Active", "Inactive"]. Default to "Active".
	Status *string `pulumi:"status"`
}

type FlowLogState struct {
	// The ID of the CEN Instance.
	CenId pulumi.StringPtrInput
	// The description of flowlog.
	Description pulumi.StringPtrInput
	// The name of flowlog.
	FlowLogName pulumi.StringPtrInput
	// The name of the log store which is in the  `projectName` SLS project.
	LogStoreName pulumi.StringPtrInput
	// The name of the SLS project.
	ProjectName pulumi.StringPtrInput
	// The status of flowlog. Valid values: ["Active", "Inactive"]. Default to "Active".
	Status pulumi.StringPtrInput
}

func (FlowLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowLogState)(nil)).Elem()
}

type flowLogArgs struct {
	// The ID of the CEN Instance.
	CenId string `pulumi:"cenId"`
	// The description of flowlog.
	Description *string `pulumi:"description"`
	// The name of flowlog.
	FlowLogName *string `pulumi:"flowLogName"`
	// The name of the log store which is in the  `projectName` SLS project.
	LogStoreName string `pulumi:"logStoreName"`
	// The name of the SLS project.
	ProjectName string `pulumi:"projectName"`
	// The status of flowlog. Valid values: ["Active", "Inactive"]. Default to "Active".
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a FlowLog resource.
type FlowLogArgs struct {
	// The ID of the CEN Instance.
	CenId pulumi.StringInput
	// The description of flowlog.
	Description pulumi.StringPtrInput
	// The name of flowlog.
	FlowLogName pulumi.StringPtrInput
	// The name of the log store which is in the  `projectName` SLS project.
	LogStoreName pulumi.StringInput
	// The name of the SLS project.
	ProjectName pulumi.StringInput
	// The status of flowlog. Valid values: ["Active", "Inactive"]. Default to "Active".
	Status pulumi.StringPtrInput
}

func (FlowLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowLogArgs)(nil)).Elem()
}

type FlowLogInput interface {
	pulumi.Input

	ToFlowLogOutput() FlowLogOutput
	ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput
}

func (*FlowLog) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowLog)(nil))
}

func (i *FlowLog) ToFlowLogOutput() FlowLogOutput {
	return i.ToFlowLogOutputWithContext(context.Background())
}

func (i *FlowLog) ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowLogOutput)
}

func (i *FlowLog) ToFlowLogPtrOutput() FlowLogPtrOutput {
	return i.ToFlowLogPtrOutputWithContext(context.Background())
}

func (i *FlowLog) ToFlowLogPtrOutputWithContext(ctx context.Context) FlowLogPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowLogPtrOutput)
}

type FlowLogPtrInput interface {
	pulumi.Input

	ToFlowLogPtrOutput() FlowLogPtrOutput
	ToFlowLogPtrOutputWithContext(ctx context.Context) FlowLogPtrOutput
}

type flowLogPtrType FlowLogArgs

func (*flowLogPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowLog)(nil))
}

func (i *flowLogPtrType) ToFlowLogPtrOutput() FlowLogPtrOutput {
	return i.ToFlowLogPtrOutputWithContext(context.Background())
}

func (i *flowLogPtrType) ToFlowLogPtrOutputWithContext(ctx context.Context) FlowLogPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowLogPtrOutput)
}

// FlowLogArrayInput is an input type that accepts FlowLogArray and FlowLogArrayOutput values.
// You can construct a concrete instance of `FlowLogArrayInput` via:
//
//          FlowLogArray{ FlowLogArgs{...} }
type FlowLogArrayInput interface {
	pulumi.Input

	ToFlowLogArrayOutput() FlowLogArrayOutput
	ToFlowLogArrayOutputWithContext(context.Context) FlowLogArrayOutput
}

type FlowLogArray []FlowLogInput

func (FlowLogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlowLog)(nil)).Elem()
}

func (i FlowLogArray) ToFlowLogArrayOutput() FlowLogArrayOutput {
	return i.ToFlowLogArrayOutputWithContext(context.Background())
}

func (i FlowLogArray) ToFlowLogArrayOutputWithContext(ctx context.Context) FlowLogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowLogArrayOutput)
}

// FlowLogMapInput is an input type that accepts FlowLogMap and FlowLogMapOutput values.
// You can construct a concrete instance of `FlowLogMapInput` via:
//
//          FlowLogMap{ "key": FlowLogArgs{...} }
type FlowLogMapInput interface {
	pulumi.Input

	ToFlowLogMapOutput() FlowLogMapOutput
	ToFlowLogMapOutputWithContext(context.Context) FlowLogMapOutput
}

type FlowLogMap map[string]FlowLogInput

func (FlowLogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlowLog)(nil)).Elem()
}

func (i FlowLogMap) ToFlowLogMapOutput() FlowLogMapOutput {
	return i.ToFlowLogMapOutputWithContext(context.Background())
}

func (i FlowLogMap) ToFlowLogMapOutputWithContext(ctx context.Context) FlowLogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowLogMapOutput)
}

type FlowLogOutput struct{ *pulumi.OutputState }

func (FlowLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowLog)(nil))
}

func (o FlowLogOutput) ToFlowLogOutput() FlowLogOutput {
	return o
}

func (o FlowLogOutput) ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput {
	return o
}

func (o FlowLogOutput) ToFlowLogPtrOutput() FlowLogPtrOutput {
	return o.ToFlowLogPtrOutputWithContext(context.Background())
}

func (o FlowLogOutput) ToFlowLogPtrOutputWithContext(ctx context.Context) FlowLogPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FlowLog) *FlowLog {
		return &v
	}).(FlowLogPtrOutput)
}

type FlowLogPtrOutput struct{ *pulumi.OutputState }

func (FlowLogPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowLog)(nil))
}

func (o FlowLogPtrOutput) ToFlowLogPtrOutput() FlowLogPtrOutput {
	return o
}

func (o FlowLogPtrOutput) ToFlowLogPtrOutputWithContext(ctx context.Context) FlowLogPtrOutput {
	return o
}

func (o FlowLogPtrOutput) Elem() FlowLogOutput {
	return o.ApplyT(func(v *FlowLog) FlowLog {
		if v != nil {
			return *v
		}
		var ret FlowLog
		return ret
	}).(FlowLogOutput)
}

type FlowLogArrayOutput struct{ *pulumi.OutputState }

func (FlowLogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FlowLog)(nil))
}

func (o FlowLogArrayOutput) ToFlowLogArrayOutput() FlowLogArrayOutput {
	return o
}

func (o FlowLogArrayOutput) ToFlowLogArrayOutputWithContext(ctx context.Context) FlowLogArrayOutput {
	return o
}

func (o FlowLogArrayOutput) Index(i pulumi.IntInput) FlowLogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FlowLog {
		return vs[0].([]FlowLog)[vs[1].(int)]
	}).(FlowLogOutput)
}

type FlowLogMapOutput struct{ *pulumi.OutputState }

func (FlowLogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FlowLog)(nil))
}

func (o FlowLogMapOutput) ToFlowLogMapOutput() FlowLogMapOutput {
	return o
}

func (o FlowLogMapOutput) ToFlowLogMapOutputWithContext(ctx context.Context) FlowLogMapOutput {
	return o
}

func (o FlowLogMapOutput) MapIndex(k pulumi.StringInput) FlowLogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FlowLog {
		return vs[0].(map[string]FlowLog)[vs[1].(string)]
	}).(FlowLogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlowLogInput)(nil)).Elem(), &FlowLog{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowLogPtrInput)(nil)).Elem(), &FlowLog{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowLogArrayInput)(nil)).Elem(), FlowLogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowLogMapInput)(nil)).Elem(), FlowLogMap{})
	pulumi.RegisterOutputType(FlowLogOutput{})
	pulumi.RegisterOutputType(FlowLogPtrOutput{})
	pulumi.RegisterOutputType(FlowLogArrayOutput{})
	pulumi.RegisterOutputType(FlowLogMapOutput{})
}
