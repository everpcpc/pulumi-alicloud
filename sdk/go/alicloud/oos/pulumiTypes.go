// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oos

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetExecutionsExecution struct {
	// The category of template. Valid: `AlarmTrigger`, `EventTrigger`, `Other` and `TimerTrigger`.
	Category string `pulumi:"category"`
	// The counters of OOS Execution.
	Counters string `pulumi:"counters"`
	// The time when the execution was created.
	CreateDate string `pulumi:"createDate"`
	// The time when the execution was ended.
	EndDate string `pulumi:"endDate"`
	// The user who execute the template.
	ExecutedBy string `pulumi:"executedBy"`
	// ID of the OOS Executions.
	ExecutionId string `pulumi:"executionId"`
	// ID of the OOS Executions.
	Id string `pulumi:"id"`
	// Whether to include subtasks.
	IsParent bool `pulumi:"isParent"`
	// The mode of OOS Execution. Valid: `Automatic`, `Debug`.
	Mode string `pulumi:"mode"`
	// The outputs of OOS Executions.
	Outputs string `pulumi:"outputs"`
	// The parameters required by the template
	Parameters string `pulumi:"parameters"`
	// The id of parent OOS Execution.
	ParentExecutionId string `pulumi:"parentExecutionId"`
	// The role that executes the current template.
	RamRole string `pulumi:"ramRole"`
	// The time when the template was started.
	StartDate string `pulumi:"startDate"`
	// The Status of OOS Execution. Valid: `Cancelled`, `Failed`, `Queued`, `Running`, `Started`, `Success`, `Waiting`.
	Status string `pulumi:"status"`
	// The message of status.
	StatusMessage string `pulumi:"statusMessage"`
	// The reason of status.
	StatusReason string `pulumi:"statusReason"`
	// The id of execution template.
	TemplateId string `pulumi:"templateId"`
	// The name of execution template.
	TemplateName string `pulumi:"templateName"`
	// The version of execution template.
	TemplateVersion string `pulumi:"templateVersion"`
	// The time when the template was updated.
	UpdateDate string `pulumi:"updateDate"`
}

// GetExecutionsExecutionInput is an input type that accepts GetExecutionsExecutionArgs and GetExecutionsExecutionOutput values.
// You can construct a concrete instance of `GetExecutionsExecutionInput` via:
//
//          GetExecutionsExecutionArgs{...}
type GetExecutionsExecutionInput interface {
	pulumi.Input

	ToGetExecutionsExecutionOutput() GetExecutionsExecutionOutput
	ToGetExecutionsExecutionOutputWithContext(context.Context) GetExecutionsExecutionOutput
}

type GetExecutionsExecutionArgs struct {
	// The category of template. Valid: `AlarmTrigger`, `EventTrigger`, `Other` and `TimerTrigger`.
	Category pulumi.StringInput `pulumi:"category"`
	// The counters of OOS Execution.
	Counters pulumi.StringInput `pulumi:"counters"`
	// The time when the execution was created.
	CreateDate pulumi.StringInput `pulumi:"createDate"`
	// The time when the execution was ended.
	EndDate pulumi.StringInput `pulumi:"endDate"`
	// The user who execute the template.
	ExecutedBy pulumi.StringInput `pulumi:"executedBy"`
	// ID of the OOS Executions.
	ExecutionId pulumi.StringInput `pulumi:"executionId"`
	// ID of the OOS Executions.
	Id pulumi.StringInput `pulumi:"id"`
	// Whether to include subtasks.
	IsParent pulumi.BoolInput `pulumi:"isParent"`
	// The mode of OOS Execution. Valid: `Automatic`, `Debug`.
	Mode pulumi.StringInput `pulumi:"mode"`
	// The outputs of OOS Executions.
	Outputs pulumi.StringInput `pulumi:"outputs"`
	// The parameters required by the template
	Parameters pulumi.StringInput `pulumi:"parameters"`
	// The id of parent OOS Execution.
	ParentExecutionId pulumi.StringInput `pulumi:"parentExecutionId"`
	// The role that executes the current template.
	RamRole pulumi.StringInput `pulumi:"ramRole"`
	// The time when the template was started.
	StartDate pulumi.StringInput `pulumi:"startDate"`
	// The Status of OOS Execution. Valid: `Cancelled`, `Failed`, `Queued`, `Running`, `Started`, `Success`, `Waiting`.
	Status pulumi.StringInput `pulumi:"status"`
	// The message of status.
	StatusMessage pulumi.StringInput `pulumi:"statusMessage"`
	// The reason of status.
	StatusReason pulumi.StringInput `pulumi:"statusReason"`
	// The id of execution template.
	TemplateId pulumi.StringInput `pulumi:"templateId"`
	// The name of execution template.
	TemplateName pulumi.StringInput `pulumi:"templateName"`
	// The version of execution template.
	TemplateVersion pulumi.StringInput `pulumi:"templateVersion"`
	// The time when the template was updated.
	UpdateDate pulumi.StringInput `pulumi:"updateDate"`
}

func (GetExecutionsExecutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExecutionsExecution)(nil)).Elem()
}

func (i GetExecutionsExecutionArgs) ToGetExecutionsExecutionOutput() GetExecutionsExecutionOutput {
	return i.ToGetExecutionsExecutionOutputWithContext(context.Background())
}

func (i GetExecutionsExecutionArgs) ToGetExecutionsExecutionOutputWithContext(ctx context.Context) GetExecutionsExecutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetExecutionsExecutionOutput)
}

// GetExecutionsExecutionArrayInput is an input type that accepts GetExecutionsExecutionArray and GetExecutionsExecutionArrayOutput values.
// You can construct a concrete instance of `GetExecutionsExecutionArrayInput` via:
//
//          GetExecutionsExecutionArray{ GetExecutionsExecutionArgs{...} }
type GetExecutionsExecutionArrayInput interface {
	pulumi.Input

	ToGetExecutionsExecutionArrayOutput() GetExecutionsExecutionArrayOutput
	ToGetExecutionsExecutionArrayOutputWithContext(context.Context) GetExecutionsExecutionArrayOutput
}

type GetExecutionsExecutionArray []GetExecutionsExecutionInput

func (GetExecutionsExecutionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetExecutionsExecution)(nil)).Elem()
}

func (i GetExecutionsExecutionArray) ToGetExecutionsExecutionArrayOutput() GetExecutionsExecutionArrayOutput {
	return i.ToGetExecutionsExecutionArrayOutputWithContext(context.Background())
}

func (i GetExecutionsExecutionArray) ToGetExecutionsExecutionArrayOutputWithContext(ctx context.Context) GetExecutionsExecutionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetExecutionsExecutionArrayOutput)
}

type GetExecutionsExecutionOutput struct{ *pulumi.OutputState }

func (GetExecutionsExecutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExecutionsExecution)(nil)).Elem()
}

func (o GetExecutionsExecutionOutput) ToGetExecutionsExecutionOutput() GetExecutionsExecutionOutput {
	return o
}

func (o GetExecutionsExecutionOutput) ToGetExecutionsExecutionOutputWithContext(ctx context.Context) GetExecutionsExecutionOutput {
	return o
}

// The category of template. Valid: `AlarmTrigger`, `EventTrigger`, `Other` and `TimerTrigger`.
func (o GetExecutionsExecutionOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.Category }).(pulumi.StringOutput)
}

// The counters of OOS Execution.
func (o GetExecutionsExecutionOutput) Counters() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.Counters }).(pulumi.StringOutput)
}

// The time when the execution was created.
func (o GetExecutionsExecutionOutput) CreateDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.CreateDate }).(pulumi.StringOutput)
}

// The time when the execution was ended.
func (o GetExecutionsExecutionOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.EndDate }).(pulumi.StringOutput)
}

// The user who execute the template.
func (o GetExecutionsExecutionOutput) ExecutedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.ExecutedBy }).(pulumi.StringOutput)
}

// ID of the OOS Executions.
func (o GetExecutionsExecutionOutput) ExecutionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.ExecutionId }).(pulumi.StringOutput)
}

// ID of the OOS Executions.
func (o GetExecutionsExecutionOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.Id }).(pulumi.StringOutput)
}

// Whether to include subtasks.
func (o GetExecutionsExecutionOutput) IsParent() pulumi.BoolOutput {
	return o.ApplyT(func(v GetExecutionsExecution) bool { return v.IsParent }).(pulumi.BoolOutput)
}

// The mode of OOS Execution. Valid: `Automatic`, `Debug`.
func (o GetExecutionsExecutionOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.Mode }).(pulumi.StringOutput)
}

// The outputs of OOS Executions.
func (o GetExecutionsExecutionOutput) Outputs() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.Outputs }).(pulumi.StringOutput)
}

// The parameters required by the template
func (o GetExecutionsExecutionOutput) Parameters() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.Parameters }).(pulumi.StringOutput)
}

// The id of parent OOS Execution.
func (o GetExecutionsExecutionOutput) ParentExecutionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.ParentExecutionId }).(pulumi.StringOutput)
}

// The role that executes the current template.
func (o GetExecutionsExecutionOutput) RamRole() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.RamRole }).(pulumi.StringOutput)
}

// The time when the template was started.
func (o GetExecutionsExecutionOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.StartDate }).(pulumi.StringOutput)
}

// The Status of OOS Execution. Valid: `Cancelled`, `Failed`, `Queued`, `Running`, `Started`, `Success`, `Waiting`.
func (o GetExecutionsExecutionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.Status }).(pulumi.StringOutput)
}

// The message of status.
func (o GetExecutionsExecutionOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.StatusMessage }).(pulumi.StringOutput)
}

// The reason of status.
func (o GetExecutionsExecutionOutput) StatusReason() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.StatusReason }).(pulumi.StringOutput)
}

// The id of execution template.
func (o GetExecutionsExecutionOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.TemplateId }).(pulumi.StringOutput)
}

// The name of execution template.
func (o GetExecutionsExecutionOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.TemplateName }).(pulumi.StringOutput)
}

// The version of execution template.
func (o GetExecutionsExecutionOutput) TemplateVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.TemplateVersion }).(pulumi.StringOutput)
}

// The time when the template was updated.
func (o GetExecutionsExecutionOutput) UpdateDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsExecution) string { return v.UpdateDate }).(pulumi.StringOutput)
}

type GetExecutionsExecutionArrayOutput struct{ *pulumi.OutputState }

func (GetExecutionsExecutionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetExecutionsExecution)(nil)).Elem()
}

func (o GetExecutionsExecutionArrayOutput) ToGetExecutionsExecutionArrayOutput() GetExecutionsExecutionArrayOutput {
	return o
}

func (o GetExecutionsExecutionArrayOutput) ToGetExecutionsExecutionArrayOutputWithContext(ctx context.Context) GetExecutionsExecutionArrayOutput {
	return o
}

func (o GetExecutionsExecutionArrayOutput) Index(i pulumi.IntInput) GetExecutionsExecutionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetExecutionsExecution {
		return vs[0].([]GetExecutionsExecution)[vs[1].(int)]
	}).(GetExecutionsExecutionOutput)
}

type GetTemplatesTemplate struct {
	// The category of template.
	Category string `pulumi:"category"`
	// The creator of the template.
	CreatedBy string `pulumi:"createdBy"`
	// The template whose creation time is less than or equal to the specified time. The format is: YYYY-MM-DDThh:mm::ssZ.
	CreatedDate string `pulumi:"createdDate"`
	// Description of the OOS Template.
	Description string `pulumi:"description"`
	// Is it triggered successfully.
	HasTrigger bool `pulumi:"hasTrigger"`
	// ID of the OOS Template. The value is same as template_name.
	Id string `pulumi:"id"`
	// The sharing type of the template. Valid values: `Private`, `Public`.
	ShareType string `pulumi:"shareType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The format of the template. Valid values: `JSON`, `YAML`.
	TemplateFormat string `pulumi:"templateFormat"`
	// ID of the OOS Template resource.
	TemplateId string `pulumi:"templateId"`
	// Name of the OOS Template.
	TemplateName string `pulumi:"templateName"`
	// The type of OOS Template.
	TemplateType string `pulumi:"templateType"`
	// Version of the OOS Template.
	TemplateVersion string `pulumi:"templateVersion"`
	// The user who updated the template.
	UpdatedBy string `pulumi:"updatedBy"`
	// The time when the template was updated.
	UpdatedDate string `pulumi:"updatedDate"`
}

// GetTemplatesTemplateInput is an input type that accepts GetTemplatesTemplateArgs and GetTemplatesTemplateOutput values.
// You can construct a concrete instance of `GetTemplatesTemplateInput` via:
//
//          GetTemplatesTemplateArgs{...}
type GetTemplatesTemplateInput interface {
	pulumi.Input

	ToGetTemplatesTemplateOutput() GetTemplatesTemplateOutput
	ToGetTemplatesTemplateOutputWithContext(context.Context) GetTemplatesTemplateOutput
}

type GetTemplatesTemplateArgs struct {
	// The category of template.
	Category pulumi.StringInput `pulumi:"category"`
	// The creator of the template.
	CreatedBy pulumi.StringInput `pulumi:"createdBy"`
	// The template whose creation time is less than or equal to the specified time. The format is: YYYY-MM-DDThh:mm::ssZ.
	CreatedDate pulumi.StringInput `pulumi:"createdDate"`
	// Description of the OOS Template.
	Description pulumi.StringInput `pulumi:"description"`
	// Is it triggered successfully.
	HasTrigger pulumi.BoolInput `pulumi:"hasTrigger"`
	// ID of the OOS Template. The value is same as template_name.
	Id pulumi.StringInput `pulumi:"id"`
	// The sharing type of the template. Valid values: `Private`, `Public`.
	ShareType pulumi.StringInput `pulumi:"shareType"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The format of the template. Valid values: `JSON`, `YAML`.
	TemplateFormat pulumi.StringInput `pulumi:"templateFormat"`
	// ID of the OOS Template resource.
	TemplateId pulumi.StringInput `pulumi:"templateId"`
	// Name of the OOS Template.
	TemplateName pulumi.StringInput `pulumi:"templateName"`
	// The type of OOS Template.
	TemplateType pulumi.StringInput `pulumi:"templateType"`
	// Version of the OOS Template.
	TemplateVersion pulumi.StringInput `pulumi:"templateVersion"`
	// The user who updated the template.
	UpdatedBy pulumi.StringInput `pulumi:"updatedBy"`
	// The time when the template was updated.
	UpdatedDate pulumi.StringInput `pulumi:"updatedDate"`
}

func (GetTemplatesTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplatesTemplate)(nil)).Elem()
}

func (i GetTemplatesTemplateArgs) ToGetTemplatesTemplateOutput() GetTemplatesTemplateOutput {
	return i.ToGetTemplatesTemplateOutputWithContext(context.Background())
}

func (i GetTemplatesTemplateArgs) ToGetTemplatesTemplateOutputWithContext(ctx context.Context) GetTemplatesTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTemplatesTemplateOutput)
}

// GetTemplatesTemplateArrayInput is an input type that accepts GetTemplatesTemplateArray and GetTemplatesTemplateArrayOutput values.
// You can construct a concrete instance of `GetTemplatesTemplateArrayInput` via:
//
//          GetTemplatesTemplateArray{ GetTemplatesTemplateArgs{...} }
type GetTemplatesTemplateArrayInput interface {
	pulumi.Input

	ToGetTemplatesTemplateArrayOutput() GetTemplatesTemplateArrayOutput
	ToGetTemplatesTemplateArrayOutputWithContext(context.Context) GetTemplatesTemplateArrayOutput
}

type GetTemplatesTemplateArray []GetTemplatesTemplateInput

func (GetTemplatesTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTemplatesTemplate)(nil)).Elem()
}

func (i GetTemplatesTemplateArray) ToGetTemplatesTemplateArrayOutput() GetTemplatesTemplateArrayOutput {
	return i.ToGetTemplatesTemplateArrayOutputWithContext(context.Background())
}

func (i GetTemplatesTemplateArray) ToGetTemplatesTemplateArrayOutputWithContext(ctx context.Context) GetTemplatesTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTemplatesTemplateArrayOutput)
}

type GetTemplatesTemplateOutput struct{ *pulumi.OutputState }

func (GetTemplatesTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplatesTemplate)(nil)).Elem()
}

func (o GetTemplatesTemplateOutput) ToGetTemplatesTemplateOutput() GetTemplatesTemplateOutput {
	return o
}

func (o GetTemplatesTemplateOutput) ToGetTemplatesTemplateOutputWithContext(ctx context.Context) GetTemplatesTemplateOutput {
	return o
}

// The category of template.
func (o GetTemplatesTemplateOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesTemplate) string { return v.Category }).(pulumi.StringOutput)
}

// The creator of the template.
func (o GetTemplatesTemplateOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesTemplate) string { return v.CreatedBy }).(pulumi.StringOutput)
}

// The template whose creation time is less than or equal to the specified time. The format is: YYYY-MM-DDThh:mm::ssZ.
func (o GetTemplatesTemplateOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesTemplate) string { return v.CreatedDate }).(pulumi.StringOutput)
}

// Description of the OOS Template.
func (o GetTemplatesTemplateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesTemplate) string { return v.Description }).(pulumi.StringOutput)
}

// Is it triggered successfully.
func (o GetTemplatesTemplateOutput) HasTrigger() pulumi.BoolOutput {
	return o.ApplyT(func(v GetTemplatesTemplate) bool { return v.HasTrigger }).(pulumi.BoolOutput)
}

// ID of the OOS Template. The value is same as template_name.
func (o GetTemplatesTemplateOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesTemplate) string { return v.Id }).(pulumi.StringOutput)
}

// The sharing type of the template. Valid values: `Private`, `Public`.
func (o GetTemplatesTemplateOutput) ShareType() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesTemplate) string { return v.ShareType }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o GetTemplatesTemplateOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetTemplatesTemplate) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

// The format of the template. Valid values: `JSON`, `YAML`.
func (o GetTemplatesTemplateOutput) TemplateFormat() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesTemplate) string { return v.TemplateFormat }).(pulumi.StringOutput)
}

// ID of the OOS Template resource.
func (o GetTemplatesTemplateOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesTemplate) string { return v.TemplateId }).(pulumi.StringOutput)
}

// Name of the OOS Template.
func (o GetTemplatesTemplateOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesTemplate) string { return v.TemplateName }).(pulumi.StringOutput)
}

// The type of OOS Template.
func (o GetTemplatesTemplateOutput) TemplateType() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesTemplate) string { return v.TemplateType }).(pulumi.StringOutput)
}

// Version of the OOS Template.
func (o GetTemplatesTemplateOutput) TemplateVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesTemplate) string { return v.TemplateVersion }).(pulumi.StringOutput)
}

// The user who updated the template.
func (o GetTemplatesTemplateOutput) UpdatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesTemplate) string { return v.UpdatedBy }).(pulumi.StringOutput)
}

// The time when the template was updated.
func (o GetTemplatesTemplateOutput) UpdatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesTemplate) string { return v.UpdatedDate }).(pulumi.StringOutput)
}

type GetTemplatesTemplateArrayOutput struct{ *pulumi.OutputState }

func (GetTemplatesTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTemplatesTemplate)(nil)).Elem()
}

func (o GetTemplatesTemplateArrayOutput) ToGetTemplatesTemplateArrayOutput() GetTemplatesTemplateArrayOutput {
	return o
}

func (o GetTemplatesTemplateArrayOutput) ToGetTemplatesTemplateArrayOutputWithContext(ctx context.Context) GetTemplatesTemplateArrayOutput {
	return o
}

func (o GetTemplatesTemplateArrayOutput) Index(i pulumi.IntInput) GetTemplatesTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetTemplatesTemplate {
		return vs[0].([]GetTemplatesTemplate)[vs[1].(int)]
	}).(GetTemplatesTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetExecutionsExecutionInput)(nil)).Elem(), GetExecutionsExecutionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetExecutionsExecutionArrayInput)(nil)).Elem(), GetExecutionsExecutionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetTemplatesTemplateInput)(nil)).Elem(), GetTemplatesTemplateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetTemplatesTemplateArrayInput)(nil)).Elem(), GetTemplatesTemplateArray{})
	pulumi.RegisterOutputType(GetExecutionsExecutionOutput{})
	pulumi.RegisterOutputType(GetExecutionsExecutionArrayOutput{})
	pulumi.RegisterOutputType(GetTemplatesTemplateOutput{})
	pulumi.RegisterOutputType(GetTemplatesTemplateArrayOutput{})
}
