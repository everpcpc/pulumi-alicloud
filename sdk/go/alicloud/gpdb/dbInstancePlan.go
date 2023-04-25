// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gpdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a GPDB DB Instance Plan resource.
//
// For information about GPDB DB Instance Plan and how to use it, see [What is DB Instance Plan](https://www.alibabacloud.com/help/zh/analyticdb-for-postgresql/latest/createdbinstanceplan).
//
// > **NOTE:** Available in v1.189.0+.
//
// ## Import
//
// GPDB DB Instance Plan can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:gpdb/dbInstancePlan:DbInstancePlan example <db_instance_id>:<plan_id>
//
// ```
type DbInstancePlan struct {
	pulumi.CustomResourceState

	// The ID of the Database instance.
	DbInstanceId pulumi.StringOutput `pulumi:"dbInstanceId"`
	// The name of the Plan.
	DbInstancePlanName pulumi.StringOutput `pulumi:"dbInstancePlanName"`
	// The plan config. See the following `Block planConfig`.
	PlanConfigs DbInstancePlanPlanConfigArrayOutput `pulumi:"planConfigs"`
	// The description of the Plan.
	PlanDesc pulumi.StringOutput `pulumi:"planDesc"`
	// The end time of the Plan.
	PlanEndDate pulumi.StringOutput `pulumi:"planEndDate"`
	// The ID of DB Instance Plan.
	PlanId pulumi.StringOutput `pulumi:"planId"`
	// Plan scheduling type. Valid values: `Postpone`, `Regular`.
	PlanScheduleType pulumi.StringOutput `pulumi:"planScheduleType"`
	// The start time of the Plan.
	PlanStartDate pulumi.StringOutput `pulumi:"planStartDate"`
	// The type of the Plan. Valid values: `PauseResume`, `Resize`.
	PlanType pulumi.StringOutput `pulumi:"planType"`
	// The Status of the Plan. Valid values: `active`, `cancel`.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewDbInstancePlan registers a new resource with the given unique name, arguments, and options.
func NewDbInstancePlan(ctx *pulumi.Context,
	name string, args *DbInstancePlanArgs, opts ...pulumi.ResourceOption) (*DbInstancePlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceId'")
	}
	if args.DbInstancePlanName == nil {
		return nil, errors.New("invalid value for required argument 'DbInstancePlanName'")
	}
	if args.PlanConfigs == nil {
		return nil, errors.New("invalid value for required argument 'PlanConfigs'")
	}
	if args.PlanScheduleType == nil {
		return nil, errors.New("invalid value for required argument 'PlanScheduleType'")
	}
	if args.PlanType == nil {
		return nil, errors.New("invalid value for required argument 'PlanType'")
	}
	var resource DbInstancePlan
	err := ctx.RegisterResource("alicloud:gpdb/dbInstancePlan:DbInstancePlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDbInstancePlan gets an existing DbInstancePlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDbInstancePlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DbInstancePlanState, opts ...pulumi.ResourceOption) (*DbInstancePlan, error) {
	var resource DbInstancePlan
	err := ctx.ReadResource("alicloud:gpdb/dbInstancePlan:DbInstancePlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DbInstancePlan resources.
type dbInstancePlanState struct {
	// The ID of the Database instance.
	DbInstanceId *string `pulumi:"dbInstanceId"`
	// The name of the Plan.
	DbInstancePlanName *string `pulumi:"dbInstancePlanName"`
	// The plan config. See the following `Block planConfig`.
	PlanConfigs []DbInstancePlanPlanConfig `pulumi:"planConfigs"`
	// The description of the Plan.
	PlanDesc *string `pulumi:"planDesc"`
	// The end time of the Plan.
	PlanEndDate *string `pulumi:"planEndDate"`
	// The ID of DB Instance Plan.
	PlanId *string `pulumi:"planId"`
	// Plan scheduling type. Valid values: `Postpone`, `Regular`.
	PlanScheduleType *string `pulumi:"planScheduleType"`
	// The start time of the Plan.
	PlanStartDate *string `pulumi:"planStartDate"`
	// The type of the Plan. Valid values: `PauseResume`, `Resize`.
	PlanType *string `pulumi:"planType"`
	// The Status of the Plan. Valid values: `active`, `cancel`.
	Status *string `pulumi:"status"`
}

type DbInstancePlanState struct {
	// The ID of the Database instance.
	DbInstanceId pulumi.StringPtrInput
	// The name of the Plan.
	DbInstancePlanName pulumi.StringPtrInput
	// The plan config. See the following `Block planConfig`.
	PlanConfigs DbInstancePlanPlanConfigArrayInput
	// The description of the Plan.
	PlanDesc pulumi.StringPtrInput
	// The end time of the Plan.
	PlanEndDate pulumi.StringPtrInput
	// The ID of DB Instance Plan.
	PlanId pulumi.StringPtrInput
	// Plan scheduling type. Valid values: `Postpone`, `Regular`.
	PlanScheduleType pulumi.StringPtrInput
	// The start time of the Plan.
	PlanStartDate pulumi.StringPtrInput
	// The type of the Plan. Valid values: `PauseResume`, `Resize`.
	PlanType pulumi.StringPtrInput
	// The Status of the Plan. Valid values: `active`, `cancel`.
	Status pulumi.StringPtrInput
}

func (DbInstancePlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbInstancePlanState)(nil)).Elem()
}

type dbInstancePlanArgs struct {
	// The ID of the Database instance.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// The name of the Plan.
	DbInstancePlanName string `pulumi:"dbInstancePlanName"`
	// The plan config. See the following `Block planConfig`.
	PlanConfigs []DbInstancePlanPlanConfig `pulumi:"planConfigs"`
	// The description of the Plan.
	PlanDesc *string `pulumi:"planDesc"`
	// The end time of the Plan.
	PlanEndDate *string `pulumi:"planEndDate"`
	// Plan scheduling type. Valid values: `Postpone`, `Regular`.
	PlanScheduleType string `pulumi:"planScheduleType"`
	// The start time of the Plan.
	PlanStartDate *string `pulumi:"planStartDate"`
	// The type of the Plan. Valid values: `PauseResume`, `Resize`.
	PlanType string `pulumi:"planType"`
	// The Status of the Plan. Valid values: `active`, `cancel`.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a DbInstancePlan resource.
type DbInstancePlanArgs struct {
	// The ID of the Database instance.
	DbInstanceId pulumi.StringInput
	// The name of the Plan.
	DbInstancePlanName pulumi.StringInput
	// The plan config. See the following `Block planConfig`.
	PlanConfigs DbInstancePlanPlanConfigArrayInput
	// The description of the Plan.
	PlanDesc pulumi.StringPtrInput
	// The end time of the Plan.
	PlanEndDate pulumi.StringPtrInput
	// Plan scheduling type. Valid values: `Postpone`, `Regular`.
	PlanScheduleType pulumi.StringInput
	// The start time of the Plan.
	PlanStartDate pulumi.StringPtrInput
	// The type of the Plan. Valid values: `PauseResume`, `Resize`.
	PlanType pulumi.StringInput
	// The Status of the Plan. Valid values: `active`, `cancel`.
	Status pulumi.StringPtrInput
}

func (DbInstancePlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbInstancePlanArgs)(nil)).Elem()
}

type DbInstancePlanInput interface {
	pulumi.Input

	ToDbInstancePlanOutput() DbInstancePlanOutput
	ToDbInstancePlanOutputWithContext(ctx context.Context) DbInstancePlanOutput
}

func (*DbInstancePlan) ElementType() reflect.Type {
	return reflect.TypeOf((**DbInstancePlan)(nil)).Elem()
}

func (i *DbInstancePlan) ToDbInstancePlanOutput() DbInstancePlanOutput {
	return i.ToDbInstancePlanOutputWithContext(context.Background())
}

func (i *DbInstancePlan) ToDbInstancePlanOutputWithContext(ctx context.Context) DbInstancePlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbInstancePlanOutput)
}

// DbInstancePlanArrayInput is an input type that accepts DbInstancePlanArray and DbInstancePlanArrayOutput values.
// You can construct a concrete instance of `DbInstancePlanArrayInput` via:
//
//	DbInstancePlanArray{ DbInstancePlanArgs{...} }
type DbInstancePlanArrayInput interface {
	pulumi.Input

	ToDbInstancePlanArrayOutput() DbInstancePlanArrayOutput
	ToDbInstancePlanArrayOutputWithContext(context.Context) DbInstancePlanArrayOutput
}

type DbInstancePlanArray []DbInstancePlanInput

func (DbInstancePlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DbInstancePlan)(nil)).Elem()
}

func (i DbInstancePlanArray) ToDbInstancePlanArrayOutput() DbInstancePlanArrayOutput {
	return i.ToDbInstancePlanArrayOutputWithContext(context.Background())
}

func (i DbInstancePlanArray) ToDbInstancePlanArrayOutputWithContext(ctx context.Context) DbInstancePlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbInstancePlanArrayOutput)
}

// DbInstancePlanMapInput is an input type that accepts DbInstancePlanMap and DbInstancePlanMapOutput values.
// You can construct a concrete instance of `DbInstancePlanMapInput` via:
//
//	DbInstancePlanMap{ "key": DbInstancePlanArgs{...} }
type DbInstancePlanMapInput interface {
	pulumi.Input

	ToDbInstancePlanMapOutput() DbInstancePlanMapOutput
	ToDbInstancePlanMapOutputWithContext(context.Context) DbInstancePlanMapOutput
}

type DbInstancePlanMap map[string]DbInstancePlanInput

func (DbInstancePlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DbInstancePlan)(nil)).Elem()
}

func (i DbInstancePlanMap) ToDbInstancePlanMapOutput() DbInstancePlanMapOutput {
	return i.ToDbInstancePlanMapOutputWithContext(context.Background())
}

func (i DbInstancePlanMap) ToDbInstancePlanMapOutputWithContext(ctx context.Context) DbInstancePlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbInstancePlanMapOutput)
}

type DbInstancePlanOutput struct{ *pulumi.OutputState }

func (DbInstancePlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DbInstancePlan)(nil)).Elem()
}

func (o DbInstancePlanOutput) ToDbInstancePlanOutput() DbInstancePlanOutput {
	return o
}

func (o DbInstancePlanOutput) ToDbInstancePlanOutputWithContext(ctx context.Context) DbInstancePlanOutput {
	return o
}

// The ID of the Database instance.
func (o DbInstancePlanOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstancePlan) pulumi.StringOutput { return v.DbInstanceId }).(pulumi.StringOutput)
}

// The name of the Plan.
func (o DbInstancePlanOutput) DbInstancePlanName() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstancePlan) pulumi.StringOutput { return v.DbInstancePlanName }).(pulumi.StringOutput)
}

// The plan config. See the following `Block planConfig`.
func (o DbInstancePlanOutput) PlanConfigs() DbInstancePlanPlanConfigArrayOutput {
	return o.ApplyT(func(v *DbInstancePlan) DbInstancePlanPlanConfigArrayOutput { return v.PlanConfigs }).(DbInstancePlanPlanConfigArrayOutput)
}

// The description of the Plan.
func (o DbInstancePlanOutput) PlanDesc() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstancePlan) pulumi.StringOutput { return v.PlanDesc }).(pulumi.StringOutput)
}

// The end time of the Plan.
func (o DbInstancePlanOutput) PlanEndDate() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstancePlan) pulumi.StringOutput { return v.PlanEndDate }).(pulumi.StringOutput)
}

// The ID of DB Instance Plan.
func (o DbInstancePlanOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstancePlan) pulumi.StringOutput { return v.PlanId }).(pulumi.StringOutput)
}

// Plan scheduling type. Valid values: `Postpone`, `Regular`.
func (o DbInstancePlanOutput) PlanScheduleType() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstancePlan) pulumi.StringOutput { return v.PlanScheduleType }).(pulumi.StringOutput)
}

// The start time of the Plan.
func (o DbInstancePlanOutput) PlanStartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstancePlan) pulumi.StringOutput { return v.PlanStartDate }).(pulumi.StringOutput)
}

// The type of the Plan. Valid values: `PauseResume`, `Resize`.
func (o DbInstancePlanOutput) PlanType() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstancePlan) pulumi.StringOutput { return v.PlanType }).(pulumi.StringOutput)
}

// The Status of the Plan. Valid values: `active`, `cancel`.
func (o DbInstancePlanOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstancePlan) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type DbInstancePlanArrayOutput struct{ *pulumi.OutputState }

func (DbInstancePlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DbInstancePlan)(nil)).Elem()
}

func (o DbInstancePlanArrayOutput) ToDbInstancePlanArrayOutput() DbInstancePlanArrayOutput {
	return o
}

func (o DbInstancePlanArrayOutput) ToDbInstancePlanArrayOutputWithContext(ctx context.Context) DbInstancePlanArrayOutput {
	return o
}

func (o DbInstancePlanArrayOutput) Index(i pulumi.IntInput) DbInstancePlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DbInstancePlan {
		return vs[0].([]*DbInstancePlan)[vs[1].(int)]
	}).(DbInstancePlanOutput)
}

type DbInstancePlanMapOutput struct{ *pulumi.OutputState }

func (DbInstancePlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DbInstancePlan)(nil)).Elem()
}

func (o DbInstancePlanMapOutput) ToDbInstancePlanMapOutput() DbInstancePlanMapOutput {
	return o
}

func (o DbInstancePlanMapOutput) ToDbInstancePlanMapOutputWithContext(ctx context.Context) DbInstancePlanMapOutput {
	return o
}

func (o DbInstancePlanMapOutput) MapIndex(k pulumi.StringInput) DbInstancePlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DbInstancePlan {
		return vs[0].(map[string]*DbInstancePlan)[vs[1].(string)]
	}).(DbInstancePlanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DbInstancePlanInput)(nil)).Elem(), &DbInstancePlan{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbInstancePlanArrayInput)(nil)).Elem(), DbInstancePlanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbInstancePlanMapInput)(nil)).Elem(), DbInstancePlanMap{})
	pulumi.RegisterOutputType(DbInstancePlanOutput{})
	pulumi.RegisterOutputType(DbInstancePlanArrayOutput{})
	pulumi.RegisterOutputType(DbInstancePlanMapOutput{})
}
