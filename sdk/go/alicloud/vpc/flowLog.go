// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Flow Log resource.
//
// For information about VPC Flow log and how to use it, see [Flow log overview](https://www.alibabacloud.com/help/doc-detail/127150.htm).
//
// > **NOTE:** Available in v1.117.0+
//
// > **NOTE:** While it uses `vpc.FlowLog` to build a vpc flow log resource, it will be active by default.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "terratest_vpc_flow_log"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			logStoreName := "vpc-flow-log-for-vpc"
//			if param := cfg.Get("logStoreName"); param != "" {
//				logStoreName = param
//			}
//			projectName := "vpc-flow-log-for-vpc"
//			if param := cfg.Get("projectName"); param != "" {
//				projectName = param
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				CidrBlock: pulumi.String("192.168.0.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewFlowLog(ctx, "defaultFlowLog", &vpc.FlowLogArgs{
//				ResourceId:   defaultNetwork.ID(),
//				ResourceType: pulumi.String("VPC"),
//				TrafficType:  pulumi.String("All"),
//				LogStoreName: pulumi.String(logStoreName),
//				ProjectName:  pulumi.String(projectName),
//				FlowLogName:  pulumi.String(name),
//				Status:       pulumi.String("Active"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				pulumi.Resource("alicloud_vpc.default"),
//			}))
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
// VPC Flow Log can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:vpc/flowLog:FlowLog example fl-abc123456
//
// ```
type FlowLog struct {
	pulumi.CustomResourceState

	// The Description of the VPC Flow Log.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Name of the VPC Flow Log.
	FlowLogName pulumi.StringPtrOutput `pulumi:"flowLogName"`
	// The name of the logstore.
	LogStoreName pulumi.StringOutput `pulumi:"logStoreName"`
	// The name of the project.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// The ID of the resource.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The type of the resource to capture traffic. Valid values `NetworkInterface`, `VPC`, and `VSwitch`.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// The status of the VPC Flow Log. Valid values `Active` and `Inactive`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The type of traffic collected. Valid values `All`, `Drop` and `Allow`.
	TrafficType pulumi.StringOutput `pulumi:"trafficType"`
}

// NewFlowLog registers a new resource with the given unique name, arguments, and options.
func NewFlowLog(ctx *pulumi.Context,
	name string, args *FlowLogArgs, opts ...pulumi.ResourceOption) (*FlowLog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LogStoreName == nil {
		return nil, errors.New("invalid value for required argument 'LogStoreName'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	if args.TrafficType == nil {
		return nil, errors.New("invalid value for required argument 'TrafficType'")
	}
	var resource FlowLog
	err := ctx.RegisterResource("alicloud:vpc/flowLog:FlowLog", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:vpc/flowLog:FlowLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlowLog resources.
type flowLogState struct {
	// The Description of the VPC Flow Log.
	Description *string `pulumi:"description"`
	// The Name of the VPC Flow Log.
	FlowLogName *string `pulumi:"flowLogName"`
	// The name of the logstore.
	LogStoreName *string `pulumi:"logStoreName"`
	// The name of the project.
	ProjectName *string `pulumi:"projectName"`
	// The ID of the resource.
	ResourceId *string `pulumi:"resourceId"`
	// The type of the resource to capture traffic. Valid values `NetworkInterface`, `VPC`, and `VSwitch`.
	ResourceType *string `pulumi:"resourceType"`
	// The status of the VPC Flow Log. Valid values `Active` and `Inactive`.
	Status *string `pulumi:"status"`
	// The type of traffic collected. Valid values `All`, `Drop` and `Allow`.
	TrafficType *string `pulumi:"trafficType"`
}

type FlowLogState struct {
	// The Description of the VPC Flow Log.
	Description pulumi.StringPtrInput
	// The Name of the VPC Flow Log.
	FlowLogName pulumi.StringPtrInput
	// The name of the logstore.
	LogStoreName pulumi.StringPtrInput
	// The name of the project.
	ProjectName pulumi.StringPtrInput
	// The ID of the resource.
	ResourceId pulumi.StringPtrInput
	// The type of the resource to capture traffic. Valid values `NetworkInterface`, `VPC`, and `VSwitch`.
	ResourceType pulumi.StringPtrInput
	// The status of the VPC Flow Log. Valid values `Active` and `Inactive`.
	Status pulumi.StringPtrInput
	// The type of traffic collected. Valid values `All`, `Drop` and `Allow`.
	TrafficType pulumi.StringPtrInput
}

func (FlowLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowLogState)(nil)).Elem()
}

type flowLogArgs struct {
	// The Description of the VPC Flow Log.
	Description *string `pulumi:"description"`
	// The Name of the VPC Flow Log.
	FlowLogName *string `pulumi:"flowLogName"`
	// The name of the logstore.
	LogStoreName string `pulumi:"logStoreName"`
	// The name of the project.
	ProjectName string `pulumi:"projectName"`
	// The ID of the resource.
	ResourceId string `pulumi:"resourceId"`
	// The type of the resource to capture traffic. Valid values `NetworkInterface`, `VPC`, and `VSwitch`.
	ResourceType string `pulumi:"resourceType"`
	// The status of the VPC Flow Log. Valid values `Active` and `Inactive`.
	Status *string `pulumi:"status"`
	// The type of traffic collected. Valid values `All`, `Drop` and `Allow`.
	TrafficType string `pulumi:"trafficType"`
}

// The set of arguments for constructing a FlowLog resource.
type FlowLogArgs struct {
	// The Description of the VPC Flow Log.
	Description pulumi.StringPtrInput
	// The Name of the VPC Flow Log.
	FlowLogName pulumi.StringPtrInput
	// The name of the logstore.
	LogStoreName pulumi.StringInput
	// The name of the project.
	ProjectName pulumi.StringInput
	// The ID of the resource.
	ResourceId pulumi.StringInput
	// The type of the resource to capture traffic. Valid values `NetworkInterface`, `VPC`, and `VSwitch`.
	ResourceType pulumi.StringInput
	// The status of the VPC Flow Log. Valid values `Active` and `Inactive`.
	Status pulumi.StringPtrInput
	// The type of traffic collected. Valid values `All`, `Drop` and `Allow`.
	TrafficType pulumi.StringInput
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
	return reflect.TypeOf((**FlowLog)(nil)).Elem()
}

func (i *FlowLog) ToFlowLogOutput() FlowLogOutput {
	return i.ToFlowLogOutputWithContext(context.Background())
}

func (i *FlowLog) ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowLogOutput)
}

// FlowLogArrayInput is an input type that accepts FlowLogArray and FlowLogArrayOutput values.
// You can construct a concrete instance of `FlowLogArrayInput` via:
//
//	FlowLogArray{ FlowLogArgs{...} }
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
//	FlowLogMap{ "key": FlowLogArgs{...} }
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
	return reflect.TypeOf((**FlowLog)(nil)).Elem()
}

func (o FlowLogOutput) ToFlowLogOutput() FlowLogOutput {
	return o
}

func (o FlowLogOutput) ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput {
	return o
}

// The Description of the VPC Flow Log.
func (o FlowLogOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Name of the VPC Flow Log.
func (o FlowLogOutput) FlowLogName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.FlowLogName }).(pulumi.StringPtrOutput)
}

// The name of the logstore.
func (o FlowLogOutput) LogStoreName() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.LogStoreName }).(pulumi.StringOutput)
}

// The name of the project.
func (o FlowLogOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// The ID of the resource.
func (o FlowLogOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// The type of the resource to capture traffic. Valid values `NetworkInterface`, `VPC`, and `VSwitch`.
func (o FlowLogOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// The status of the VPC Flow Log. Valid values `Active` and `Inactive`.
func (o FlowLogOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The type of traffic collected. Valid values `All`, `Drop` and `Allow`.
func (o FlowLogOutput) TrafficType() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.TrafficType }).(pulumi.StringOutput)
}

type FlowLogArrayOutput struct{ *pulumi.OutputState }

func (FlowLogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlowLog)(nil)).Elem()
}

func (o FlowLogArrayOutput) ToFlowLogArrayOutput() FlowLogArrayOutput {
	return o
}

func (o FlowLogArrayOutput) ToFlowLogArrayOutputWithContext(ctx context.Context) FlowLogArrayOutput {
	return o
}

func (o FlowLogArrayOutput) Index(i pulumi.IntInput) FlowLogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FlowLog {
		return vs[0].([]*FlowLog)[vs[1].(int)]
	}).(FlowLogOutput)
}

type FlowLogMapOutput struct{ *pulumi.OutputState }

func (FlowLogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlowLog)(nil)).Elem()
}

func (o FlowLogMapOutput) ToFlowLogMapOutput() FlowLogMapOutput {
	return o
}

func (o FlowLogMapOutput) ToFlowLogMapOutputWithContext(ctx context.Context) FlowLogMapOutput {
	return o
}

func (o FlowLogMapOutput) MapIndex(k pulumi.StringInput) FlowLogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FlowLog {
		return vs[0].(map[string]*FlowLog)[vs[1].(string)]
	}).(FlowLogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlowLogInput)(nil)).Elem(), &FlowLog{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowLogArrayInput)(nil)).Elem(), FlowLogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowLogMapInput)(nil)).Elem(), FlowLogMap{})
	pulumi.RegisterOutputType(FlowLogOutput{})
	pulumi.RegisterOutputType(FlowLogArrayOutput{})
	pulumi.RegisterOutputType(FlowLogMapOutput{})
}
