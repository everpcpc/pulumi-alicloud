// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterBootstrapAction struct {
	// bootstrap action args, e.g. "--a=b".
	Arg *string `pulumi:"arg"`
	// bootstrap action name.
	Name *string `pulumi:"name"`
	// bootstrap action path, e.g. "oss://bucket/path".
	Path *string `pulumi:"path"`
}

// ClusterBootstrapActionInput is an input type that accepts ClusterBootstrapActionArgs and ClusterBootstrapActionOutput values.
// You can construct a concrete instance of `ClusterBootstrapActionInput` via:
//
//          ClusterBootstrapActionArgs{...}
type ClusterBootstrapActionInput interface {
	pulumi.Input

	ToClusterBootstrapActionOutput() ClusterBootstrapActionOutput
	ToClusterBootstrapActionOutputWithContext(context.Context) ClusterBootstrapActionOutput
}

type ClusterBootstrapActionArgs struct {
	// bootstrap action args, e.g. "--a=b".
	Arg pulumi.StringPtrInput `pulumi:"arg"`
	// bootstrap action name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// bootstrap action path, e.g. "oss://bucket/path".
	Path pulumi.StringPtrInput `pulumi:"path"`
}

func (ClusterBootstrapActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterBootstrapAction)(nil)).Elem()
}

func (i ClusterBootstrapActionArgs) ToClusterBootstrapActionOutput() ClusterBootstrapActionOutput {
	return i.ToClusterBootstrapActionOutputWithContext(context.Background())
}

func (i ClusterBootstrapActionArgs) ToClusterBootstrapActionOutputWithContext(ctx context.Context) ClusterBootstrapActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterBootstrapActionOutput)
}

// ClusterBootstrapActionArrayInput is an input type that accepts ClusterBootstrapActionArray and ClusterBootstrapActionArrayOutput values.
// You can construct a concrete instance of `ClusterBootstrapActionArrayInput` via:
//
//          ClusterBootstrapActionArray{ ClusterBootstrapActionArgs{...} }
type ClusterBootstrapActionArrayInput interface {
	pulumi.Input

	ToClusterBootstrapActionArrayOutput() ClusterBootstrapActionArrayOutput
	ToClusterBootstrapActionArrayOutputWithContext(context.Context) ClusterBootstrapActionArrayOutput
}

type ClusterBootstrapActionArray []ClusterBootstrapActionInput

func (ClusterBootstrapActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterBootstrapAction)(nil)).Elem()
}

func (i ClusterBootstrapActionArray) ToClusterBootstrapActionArrayOutput() ClusterBootstrapActionArrayOutput {
	return i.ToClusterBootstrapActionArrayOutputWithContext(context.Background())
}

func (i ClusterBootstrapActionArray) ToClusterBootstrapActionArrayOutputWithContext(ctx context.Context) ClusterBootstrapActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterBootstrapActionArrayOutput)
}

type ClusterBootstrapActionOutput struct{ *pulumi.OutputState }

func (ClusterBootstrapActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterBootstrapAction)(nil)).Elem()
}

func (o ClusterBootstrapActionOutput) ToClusterBootstrapActionOutput() ClusterBootstrapActionOutput {
	return o
}

func (o ClusterBootstrapActionOutput) ToClusterBootstrapActionOutputWithContext(ctx context.Context) ClusterBootstrapActionOutput {
	return o
}

// bootstrap action args, e.g. "--a=b".
func (o ClusterBootstrapActionOutput) Arg() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterBootstrapAction) *string { return v.Arg }).(pulumi.StringPtrOutput)
}

// bootstrap action name.
func (o ClusterBootstrapActionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterBootstrapAction) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// bootstrap action path, e.g. "oss://bucket/path".
func (o ClusterBootstrapActionOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterBootstrapAction) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type ClusterBootstrapActionArrayOutput struct{ *pulumi.OutputState }

func (ClusterBootstrapActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterBootstrapAction)(nil)).Elem()
}

func (o ClusterBootstrapActionArrayOutput) ToClusterBootstrapActionArrayOutput() ClusterBootstrapActionArrayOutput {
	return o
}

func (o ClusterBootstrapActionArrayOutput) ToClusterBootstrapActionArrayOutputWithContext(ctx context.Context) ClusterBootstrapActionArrayOutput {
	return o
}

func (o ClusterBootstrapActionArrayOutput) Index(i pulumi.IntInput) ClusterBootstrapActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterBootstrapAction {
		return vs[0].([]ClusterBootstrapAction)[vs[1].(int)]
	}).(ClusterBootstrapActionOutput)
}

type ClusterHostGroup struct {
	// Auto renew for prepaid, true of false. Default is false.
	AutoRenew *bool `pulumi:"autoRenew"`
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType *string `pulumi:"chargeType"`
	// Data disk capacity.
	DiskCapacity *string `pulumi:"diskCapacity"`
	// Data disk count.
	DiskCount *string `pulumi:"diskCount"`
	// Data disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,local_disk,cloud_essd.
	DiskType  *string `pulumi:"diskType"`
	GpuDriver *string `pulumi:"gpuDriver"`
	// host group name.
	HostGroupName *string `pulumi:"hostGroupName"`
	// host group type, supported value: MASTER, CORE or TASK, supported 'GATEWAY' available in 1.61.0+.
	HostGroupType *string `pulumi:"hostGroupType"`
	// Instance list for cluster scale down. This value follows the json format, e.g. ["instanceId1","instanceId2"]. escape character for " is \".
	InstanceList *string `pulumi:"instanceList"`
	// Host Ecs instance type.
	InstanceType *string `pulumi:"instanceType"`
	// Host number in this group.
	NodeCount *string `pulumi:"nodeCount"`
	// If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
	Period *int `pulumi:"period"`
	// System disk capacity.
	SysDiskCapacity *string `pulumi:"sysDiskCapacity"`
	// System disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,cloud_essd.
	SysDiskType *string `pulumi:"sysDiskType"`
}

// ClusterHostGroupInput is an input type that accepts ClusterHostGroupArgs and ClusterHostGroupOutput values.
// You can construct a concrete instance of `ClusterHostGroupInput` via:
//
//          ClusterHostGroupArgs{...}
type ClusterHostGroupInput interface {
	pulumi.Input

	ToClusterHostGroupOutput() ClusterHostGroupOutput
	ToClusterHostGroupOutputWithContext(context.Context) ClusterHostGroupOutput
}

type ClusterHostGroupArgs struct {
	// Auto renew for prepaid, true of false. Default is false.
	AutoRenew pulumi.BoolPtrInput `pulumi:"autoRenew"`
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType pulumi.StringPtrInput `pulumi:"chargeType"`
	// Data disk capacity.
	DiskCapacity pulumi.StringPtrInput `pulumi:"diskCapacity"`
	// Data disk count.
	DiskCount pulumi.StringPtrInput `pulumi:"diskCount"`
	// Data disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,local_disk,cloud_essd.
	DiskType  pulumi.StringPtrInput `pulumi:"diskType"`
	GpuDriver pulumi.StringPtrInput `pulumi:"gpuDriver"`
	// host group name.
	HostGroupName pulumi.StringPtrInput `pulumi:"hostGroupName"`
	// host group type, supported value: MASTER, CORE or TASK, supported 'GATEWAY' available in 1.61.0+.
	HostGroupType pulumi.StringPtrInput `pulumi:"hostGroupType"`
	// Instance list for cluster scale down. This value follows the json format, e.g. ["instanceId1","instanceId2"]. escape character for " is \".
	InstanceList pulumi.StringPtrInput `pulumi:"instanceList"`
	// Host Ecs instance type.
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
	// Host number in this group.
	NodeCount pulumi.StringPtrInput `pulumi:"nodeCount"`
	// If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
	Period pulumi.IntPtrInput `pulumi:"period"`
	// System disk capacity.
	SysDiskCapacity pulumi.StringPtrInput `pulumi:"sysDiskCapacity"`
	// System disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,cloud_essd.
	SysDiskType pulumi.StringPtrInput `pulumi:"sysDiskType"`
}

func (ClusterHostGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterHostGroup)(nil)).Elem()
}

func (i ClusterHostGroupArgs) ToClusterHostGroupOutput() ClusterHostGroupOutput {
	return i.ToClusterHostGroupOutputWithContext(context.Background())
}

func (i ClusterHostGroupArgs) ToClusterHostGroupOutputWithContext(ctx context.Context) ClusterHostGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterHostGroupOutput)
}

// ClusterHostGroupArrayInput is an input type that accepts ClusterHostGroupArray and ClusterHostGroupArrayOutput values.
// You can construct a concrete instance of `ClusterHostGroupArrayInput` via:
//
//          ClusterHostGroupArray{ ClusterHostGroupArgs{...} }
type ClusterHostGroupArrayInput interface {
	pulumi.Input

	ToClusterHostGroupArrayOutput() ClusterHostGroupArrayOutput
	ToClusterHostGroupArrayOutputWithContext(context.Context) ClusterHostGroupArrayOutput
}

type ClusterHostGroupArray []ClusterHostGroupInput

func (ClusterHostGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterHostGroup)(nil)).Elem()
}

func (i ClusterHostGroupArray) ToClusterHostGroupArrayOutput() ClusterHostGroupArrayOutput {
	return i.ToClusterHostGroupArrayOutputWithContext(context.Background())
}

func (i ClusterHostGroupArray) ToClusterHostGroupArrayOutputWithContext(ctx context.Context) ClusterHostGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterHostGroupArrayOutput)
}

type ClusterHostGroupOutput struct{ *pulumi.OutputState }

func (ClusterHostGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterHostGroup)(nil)).Elem()
}

func (o ClusterHostGroupOutput) ToClusterHostGroupOutput() ClusterHostGroupOutput {
	return o
}

func (o ClusterHostGroupOutput) ToClusterHostGroupOutputWithContext(ctx context.Context) ClusterHostGroupOutput {
	return o
}

// Auto renew for prepaid, true of false. Default is false.
func (o ClusterHostGroupOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClusterHostGroup) *bool { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
func (o ClusterHostGroupOutput) ChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterHostGroup) *string { return v.ChargeType }).(pulumi.StringPtrOutput)
}

// Data disk capacity.
func (o ClusterHostGroupOutput) DiskCapacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterHostGroup) *string { return v.DiskCapacity }).(pulumi.StringPtrOutput)
}

// Data disk count.
func (o ClusterHostGroupOutput) DiskCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterHostGroup) *string { return v.DiskCount }).(pulumi.StringPtrOutput)
}

// Data disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,local_disk,cloud_essd.
func (o ClusterHostGroupOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterHostGroup) *string { return v.DiskType }).(pulumi.StringPtrOutput)
}

func (o ClusterHostGroupOutput) GpuDriver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterHostGroup) *string { return v.GpuDriver }).(pulumi.StringPtrOutput)
}

// host group name.
func (o ClusterHostGroupOutput) HostGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterHostGroup) *string { return v.HostGroupName }).(pulumi.StringPtrOutput)
}

// host group type, supported value: MASTER, CORE or TASK, supported 'GATEWAY' available in 1.61.0+.
func (o ClusterHostGroupOutput) HostGroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterHostGroup) *string { return v.HostGroupType }).(pulumi.StringPtrOutput)
}

// Instance list for cluster scale down. This value follows the json format, e.g. ["instanceId1","instanceId2"]. escape character for " is \".
func (o ClusterHostGroupOutput) InstanceList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterHostGroup) *string { return v.InstanceList }).(pulumi.StringPtrOutput)
}

// Host Ecs instance type.
func (o ClusterHostGroupOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterHostGroup) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

// Host number in this group.
func (o ClusterHostGroupOutput) NodeCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterHostGroup) *string { return v.NodeCount }).(pulumi.StringPtrOutput)
}

// If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
func (o ClusterHostGroupOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterHostGroup) *int { return v.Period }).(pulumi.IntPtrOutput)
}

// System disk capacity.
func (o ClusterHostGroupOutput) SysDiskCapacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterHostGroup) *string { return v.SysDiskCapacity }).(pulumi.StringPtrOutput)
}

// System disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,cloud_essd.
func (o ClusterHostGroupOutput) SysDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterHostGroup) *string { return v.SysDiskType }).(pulumi.StringPtrOutput)
}

type ClusterHostGroupArrayOutput struct{ *pulumi.OutputState }

func (ClusterHostGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterHostGroup)(nil)).Elem()
}

func (o ClusterHostGroupArrayOutput) ToClusterHostGroupArrayOutput() ClusterHostGroupArrayOutput {
	return o
}

func (o ClusterHostGroupArrayOutput) ToClusterHostGroupArrayOutputWithContext(ctx context.Context) ClusterHostGroupArrayOutput {
	return o
}

func (o ClusterHostGroupArrayOutput) Index(i pulumi.IntInput) ClusterHostGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterHostGroup {
		return vs[0].([]ClusterHostGroup)[vs[1].(int)]
	}).(ClusterHostGroupOutput)
}

type GetDiskTypesType struct {
	// The maximum value of the data disk to supported the specific instance type
	Max int `pulumi:"max"`
	// The mininum value of the data disk to supported the specific instance type
	Min int `pulumi:"min"`
	// The value of the data disk or system disk
	Value string `pulumi:"value"`
}

// GetDiskTypesTypeInput is an input type that accepts GetDiskTypesTypeArgs and GetDiskTypesTypeOutput values.
// You can construct a concrete instance of `GetDiskTypesTypeInput` via:
//
//          GetDiskTypesTypeArgs{...}
type GetDiskTypesTypeInput interface {
	pulumi.Input

	ToGetDiskTypesTypeOutput() GetDiskTypesTypeOutput
	ToGetDiskTypesTypeOutputWithContext(context.Context) GetDiskTypesTypeOutput
}

type GetDiskTypesTypeArgs struct {
	// The maximum value of the data disk to supported the specific instance type
	Max pulumi.IntInput `pulumi:"max"`
	// The mininum value of the data disk to supported the specific instance type
	Min pulumi.IntInput `pulumi:"min"`
	// The value of the data disk or system disk
	Value pulumi.StringInput `pulumi:"value"`
}

func (GetDiskTypesTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDiskTypesType)(nil)).Elem()
}

func (i GetDiskTypesTypeArgs) ToGetDiskTypesTypeOutput() GetDiskTypesTypeOutput {
	return i.ToGetDiskTypesTypeOutputWithContext(context.Background())
}

func (i GetDiskTypesTypeArgs) ToGetDiskTypesTypeOutputWithContext(ctx context.Context) GetDiskTypesTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDiskTypesTypeOutput)
}

// GetDiskTypesTypeArrayInput is an input type that accepts GetDiskTypesTypeArray and GetDiskTypesTypeArrayOutput values.
// You can construct a concrete instance of `GetDiskTypesTypeArrayInput` via:
//
//          GetDiskTypesTypeArray{ GetDiskTypesTypeArgs{...} }
type GetDiskTypesTypeArrayInput interface {
	pulumi.Input

	ToGetDiskTypesTypeArrayOutput() GetDiskTypesTypeArrayOutput
	ToGetDiskTypesTypeArrayOutputWithContext(context.Context) GetDiskTypesTypeArrayOutput
}

type GetDiskTypesTypeArray []GetDiskTypesTypeInput

func (GetDiskTypesTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDiskTypesType)(nil)).Elem()
}

func (i GetDiskTypesTypeArray) ToGetDiskTypesTypeArrayOutput() GetDiskTypesTypeArrayOutput {
	return i.ToGetDiskTypesTypeArrayOutputWithContext(context.Background())
}

func (i GetDiskTypesTypeArray) ToGetDiskTypesTypeArrayOutputWithContext(ctx context.Context) GetDiskTypesTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDiskTypesTypeArrayOutput)
}

type GetDiskTypesTypeOutput struct{ *pulumi.OutputState }

func (GetDiskTypesTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDiskTypesType)(nil)).Elem()
}

func (o GetDiskTypesTypeOutput) ToGetDiskTypesTypeOutput() GetDiskTypesTypeOutput {
	return o
}

func (o GetDiskTypesTypeOutput) ToGetDiskTypesTypeOutputWithContext(ctx context.Context) GetDiskTypesTypeOutput {
	return o
}

// The maximum value of the data disk to supported the specific instance type
func (o GetDiskTypesTypeOutput) Max() pulumi.IntOutput {
	return o.ApplyT(func(v GetDiskTypesType) int { return v.Max }).(pulumi.IntOutput)
}

// The mininum value of the data disk to supported the specific instance type
func (o GetDiskTypesTypeOutput) Min() pulumi.IntOutput {
	return o.ApplyT(func(v GetDiskTypesType) int { return v.Min }).(pulumi.IntOutput)
}

// The value of the data disk or system disk
func (o GetDiskTypesTypeOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiskTypesType) string { return v.Value }).(pulumi.StringOutput)
}

type GetDiskTypesTypeArrayOutput struct{ *pulumi.OutputState }

func (GetDiskTypesTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDiskTypesType)(nil)).Elem()
}

func (o GetDiskTypesTypeArrayOutput) ToGetDiskTypesTypeArrayOutput() GetDiskTypesTypeArrayOutput {
	return o
}

func (o GetDiskTypesTypeArrayOutput) ToGetDiskTypesTypeArrayOutputWithContext(ctx context.Context) GetDiskTypesTypeArrayOutput {
	return o
}

func (o GetDiskTypesTypeArrayOutput) Index(i pulumi.IntInput) GetDiskTypesTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDiskTypesType {
		return vs[0].([]GetDiskTypesType)[vs[1].(int)]
	}).(GetDiskTypesTypeOutput)
}

type GetInstanceTypesType struct {
	// The ID of the instance type.
	Id string `pulumi:"id"`
	// Local capacity of the applied ecs instance for emr cluster. Unit: GB.
	LocalStorageCapacity int `pulumi:"localStorageCapacity"`
	// The supported resources of specific zoneId.
	ZoneId string `pulumi:"zoneId"`
}

// GetInstanceTypesTypeInput is an input type that accepts GetInstanceTypesTypeArgs and GetInstanceTypesTypeOutput values.
// You can construct a concrete instance of `GetInstanceTypesTypeInput` via:
//
//          GetInstanceTypesTypeArgs{...}
type GetInstanceTypesTypeInput interface {
	pulumi.Input

	ToGetInstanceTypesTypeOutput() GetInstanceTypesTypeOutput
	ToGetInstanceTypesTypeOutputWithContext(context.Context) GetInstanceTypesTypeOutput
}

type GetInstanceTypesTypeArgs struct {
	// The ID of the instance type.
	Id pulumi.StringInput `pulumi:"id"`
	// Local capacity of the applied ecs instance for emr cluster. Unit: GB.
	LocalStorageCapacity pulumi.IntInput `pulumi:"localStorageCapacity"`
	// The supported resources of specific zoneId.
	ZoneId pulumi.StringInput `pulumi:"zoneId"`
}

func (GetInstanceTypesTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceTypesType)(nil)).Elem()
}

func (i GetInstanceTypesTypeArgs) ToGetInstanceTypesTypeOutput() GetInstanceTypesTypeOutput {
	return i.ToGetInstanceTypesTypeOutputWithContext(context.Background())
}

func (i GetInstanceTypesTypeArgs) ToGetInstanceTypesTypeOutputWithContext(ctx context.Context) GetInstanceTypesTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceTypesTypeOutput)
}

// GetInstanceTypesTypeArrayInput is an input type that accepts GetInstanceTypesTypeArray and GetInstanceTypesTypeArrayOutput values.
// You can construct a concrete instance of `GetInstanceTypesTypeArrayInput` via:
//
//          GetInstanceTypesTypeArray{ GetInstanceTypesTypeArgs{...} }
type GetInstanceTypesTypeArrayInput interface {
	pulumi.Input

	ToGetInstanceTypesTypeArrayOutput() GetInstanceTypesTypeArrayOutput
	ToGetInstanceTypesTypeArrayOutputWithContext(context.Context) GetInstanceTypesTypeArrayOutput
}

type GetInstanceTypesTypeArray []GetInstanceTypesTypeInput

func (GetInstanceTypesTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceTypesType)(nil)).Elem()
}

func (i GetInstanceTypesTypeArray) ToGetInstanceTypesTypeArrayOutput() GetInstanceTypesTypeArrayOutput {
	return i.ToGetInstanceTypesTypeArrayOutputWithContext(context.Background())
}

func (i GetInstanceTypesTypeArray) ToGetInstanceTypesTypeArrayOutputWithContext(ctx context.Context) GetInstanceTypesTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceTypesTypeArrayOutput)
}

type GetInstanceTypesTypeOutput struct{ *pulumi.OutputState }

func (GetInstanceTypesTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceTypesType)(nil)).Elem()
}

func (o GetInstanceTypesTypeOutput) ToGetInstanceTypesTypeOutput() GetInstanceTypesTypeOutput {
	return o
}

func (o GetInstanceTypesTypeOutput) ToGetInstanceTypesTypeOutputWithContext(ctx context.Context) GetInstanceTypesTypeOutput {
	return o
}

// The ID of the instance type.
func (o GetInstanceTypesTypeOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceTypesType) string { return v.Id }).(pulumi.StringOutput)
}

// Local capacity of the applied ecs instance for emr cluster. Unit: GB.
func (o GetInstanceTypesTypeOutput) LocalStorageCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstanceTypesType) int { return v.LocalStorageCapacity }).(pulumi.IntOutput)
}

// The supported resources of specific zoneId.
func (o GetInstanceTypesTypeOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceTypesType) string { return v.ZoneId }).(pulumi.StringOutput)
}

type GetInstanceTypesTypeArrayOutput struct{ *pulumi.OutputState }

func (GetInstanceTypesTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceTypesType)(nil)).Elem()
}

func (o GetInstanceTypesTypeArrayOutput) ToGetInstanceTypesTypeArrayOutput() GetInstanceTypesTypeArrayOutput {
	return o
}

func (o GetInstanceTypesTypeArrayOutput) ToGetInstanceTypesTypeArrayOutputWithContext(ctx context.Context) GetInstanceTypesTypeArrayOutput {
	return o
}

func (o GetInstanceTypesTypeArrayOutput) Index(i pulumi.IntInput) GetInstanceTypesTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetInstanceTypesType {
		return vs[0].([]GetInstanceTypesType)[vs[1].(int)]
	}).(GetInstanceTypesTypeOutput)
}

type GetMainVersionsMainVersion struct {
	// A list of cluster types the emr cluster supported. Possible values: `HADOOP`, `ZOOKEEPER`, `KAFKA`, `DRUID`.
	ClusterTypes []string `pulumi:"clusterTypes"`
	// The version of the emr cluster instance. Possible values: `EMR-4.0.0`, `EMR-3.23.0`, `EMR-3.22.0`.
	EmrVersion string `pulumi:"emrVersion"`
	// The image id of the emr cluster instance.
	ImageId string `pulumi:"imageId"`
}

// GetMainVersionsMainVersionInput is an input type that accepts GetMainVersionsMainVersionArgs and GetMainVersionsMainVersionOutput values.
// You can construct a concrete instance of `GetMainVersionsMainVersionInput` via:
//
//          GetMainVersionsMainVersionArgs{...}
type GetMainVersionsMainVersionInput interface {
	pulumi.Input

	ToGetMainVersionsMainVersionOutput() GetMainVersionsMainVersionOutput
	ToGetMainVersionsMainVersionOutputWithContext(context.Context) GetMainVersionsMainVersionOutput
}

type GetMainVersionsMainVersionArgs struct {
	// A list of cluster types the emr cluster supported. Possible values: `HADOOP`, `ZOOKEEPER`, `KAFKA`, `DRUID`.
	ClusterTypes pulumi.StringArrayInput `pulumi:"clusterTypes"`
	// The version of the emr cluster instance. Possible values: `EMR-4.0.0`, `EMR-3.23.0`, `EMR-3.22.0`.
	EmrVersion pulumi.StringInput `pulumi:"emrVersion"`
	// The image id of the emr cluster instance.
	ImageId pulumi.StringInput `pulumi:"imageId"`
}

func (GetMainVersionsMainVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMainVersionsMainVersion)(nil)).Elem()
}

func (i GetMainVersionsMainVersionArgs) ToGetMainVersionsMainVersionOutput() GetMainVersionsMainVersionOutput {
	return i.ToGetMainVersionsMainVersionOutputWithContext(context.Background())
}

func (i GetMainVersionsMainVersionArgs) ToGetMainVersionsMainVersionOutputWithContext(ctx context.Context) GetMainVersionsMainVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetMainVersionsMainVersionOutput)
}

// GetMainVersionsMainVersionArrayInput is an input type that accepts GetMainVersionsMainVersionArray and GetMainVersionsMainVersionArrayOutput values.
// You can construct a concrete instance of `GetMainVersionsMainVersionArrayInput` via:
//
//          GetMainVersionsMainVersionArray{ GetMainVersionsMainVersionArgs{...} }
type GetMainVersionsMainVersionArrayInput interface {
	pulumi.Input

	ToGetMainVersionsMainVersionArrayOutput() GetMainVersionsMainVersionArrayOutput
	ToGetMainVersionsMainVersionArrayOutputWithContext(context.Context) GetMainVersionsMainVersionArrayOutput
}

type GetMainVersionsMainVersionArray []GetMainVersionsMainVersionInput

func (GetMainVersionsMainVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetMainVersionsMainVersion)(nil)).Elem()
}

func (i GetMainVersionsMainVersionArray) ToGetMainVersionsMainVersionArrayOutput() GetMainVersionsMainVersionArrayOutput {
	return i.ToGetMainVersionsMainVersionArrayOutputWithContext(context.Background())
}

func (i GetMainVersionsMainVersionArray) ToGetMainVersionsMainVersionArrayOutputWithContext(ctx context.Context) GetMainVersionsMainVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetMainVersionsMainVersionArrayOutput)
}

type GetMainVersionsMainVersionOutput struct{ *pulumi.OutputState }

func (GetMainVersionsMainVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMainVersionsMainVersion)(nil)).Elem()
}

func (o GetMainVersionsMainVersionOutput) ToGetMainVersionsMainVersionOutput() GetMainVersionsMainVersionOutput {
	return o
}

func (o GetMainVersionsMainVersionOutput) ToGetMainVersionsMainVersionOutputWithContext(ctx context.Context) GetMainVersionsMainVersionOutput {
	return o
}

// A list of cluster types the emr cluster supported. Possible values: `HADOOP`, `ZOOKEEPER`, `KAFKA`, `DRUID`.
func (o GetMainVersionsMainVersionOutput) ClusterTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMainVersionsMainVersion) []string { return v.ClusterTypes }).(pulumi.StringArrayOutput)
}

// The version of the emr cluster instance. Possible values: `EMR-4.0.0`, `EMR-3.23.0`, `EMR-3.22.0`.
func (o GetMainVersionsMainVersionOutput) EmrVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetMainVersionsMainVersion) string { return v.EmrVersion }).(pulumi.StringOutput)
}

// The image id of the emr cluster instance.
func (o GetMainVersionsMainVersionOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMainVersionsMainVersion) string { return v.ImageId }).(pulumi.StringOutput)
}

type GetMainVersionsMainVersionArrayOutput struct{ *pulumi.OutputState }

func (GetMainVersionsMainVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetMainVersionsMainVersion)(nil)).Elem()
}

func (o GetMainVersionsMainVersionArrayOutput) ToGetMainVersionsMainVersionArrayOutput() GetMainVersionsMainVersionArrayOutput {
	return o
}

func (o GetMainVersionsMainVersionArrayOutput) ToGetMainVersionsMainVersionArrayOutputWithContext(ctx context.Context) GetMainVersionsMainVersionArrayOutput {
	return o
}

func (o GetMainVersionsMainVersionArrayOutput) Index(i pulumi.IntInput) GetMainVersionsMainVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetMainVersionsMainVersion {
		return vs[0].([]GetMainVersionsMainVersion)[vs[1].(int)]
	}).(GetMainVersionsMainVersionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterBootstrapActionInput)(nil)).Elem(), ClusterBootstrapActionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterBootstrapActionArrayInput)(nil)).Elem(), ClusterBootstrapActionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterHostGroupInput)(nil)).Elem(), ClusterHostGroupArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterHostGroupArrayInput)(nil)).Elem(), ClusterHostGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDiskTypesTypeInput)(nil)).Elem(), GetDiskTypesTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDiskTypesTypeArrayInput)(nil)).Elem(), GetDiskTypesTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInstanceTypesTypeInput)(nil)).Elem(), GetInstanceTypesTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInstanceTypesTypeArrayInput)(nil)).Elem(), GetInstanceTypesTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetMainVersionsMainVersionInput)(nil)).Elem(), GetMainVersionsMainVersionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetMainVersionsMainVersionArrayInput)(nil)).Elem(), GetMainVersionsMainVersionArray{})
	pulumi.RegisterOutputType(ClusterBootstrapActionOutput{})
	pulumi.RegisterOutputType(ClusterBootstrapActionArrayOutput{})
	pulumi.RegisterOutputType(ClusterHostGroupOutput{})
	pulumi.RegisterOutputType(ClusterHostGroupArrayOutput{})
	pulumi.RegisterOutputType(GetDiskTypesTypeOutput{})
	pulumi.RegisterOutputType(GetDiskTypesTypeArrayOutput{})
	pulumi.RegisterOutputType(GetInstanceTypesTypeOutput{})
	pulumi.RegisterOutputType(GetInstanceTypesTypeArrayOutput{})
	pulumi.RegisterOutputType(GetMainVersionsMainVersionOutput{})
	pulumi.RegisterOutputType(GetMainVersionsMainVersionArrayOutput{})
}
