// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cen

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This resource used to create a flow log function in Cloud Enterprise Network (CEN).
// By using the flow log function, you can capture the traffic data of the network instances in different regions of a CEN.
// You can also use the data aggregated in flow logs to analyze cross-region traffic flows, minimize traffic costs, and troubleshoot network faults.
//
// For information about CEN flow log and how to use it, see [Manage CEN flowlog](https://www.alibabacloud.com/help/doc-detail/123006.htm).
//
// > **NOTE:** Available in 1.73.0+
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cen_flowlog.html.markdown.
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
	if args == nil || args.CenId == nil {
		return nil, errors.New("missing required argument 'CenId'")
	}
	if args == nil || args.LogStoreName == nil {
		return nil, errors.New("missing required argument 'LogStoreName'")
	}
	if args == nil || args.ProjectName == nil {
		return nil, errors.New("missing required argument 'ProjectName'")
	}
	if args == nil {
		args = &FlowLogArgs{}
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