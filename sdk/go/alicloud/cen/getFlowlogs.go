// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides CEN flow logs available to the user.
//
// > **NOTE:** Available in 1.78.0+
func GetFlowlogs(ctx *pulumi.Context, args *GetFlowlogsArgs, opts ...pulumi.InvokeOption) (*GetFlowlogsResult, error) {
	var rv GetFlowlogsResult
	err := ctx.Invoke("alicloud:cen/getFlowlogs:getFlowlogs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFlowlogs.
type GetFlowlogsArgs struct {
	// The ID of the CEN Instance.
	CenId *string `pulumi:"cenId"`
	// The description of flowlog.
	Description *string `pulumi:"description"`
	// A list of CEN flow log IDs.
	Ids []string `pulumi:"ids"`
	// The name of the log store which is in the  `projectName` SLS project.
	LogStoreName *string `pulumi:"logStoreName"`
	// A regex string to filter CEN flow logs by name.
	NameRegex *string `pulumi:"nameRegex"`
	// The name of the SLS project.
	ProjectName *string `pulumi:"projectName"`
	// The status of flowlog. Valid values: ["Active", "Inactive"]. Default to "Active".
	Status *string `pulumi:"status"`
}

// A collection of values returned by getFlowlogs.
type GetFlowlogsResult struct {
	// The ID of the CEN Instance.
	CenId *string `pulumi:"cenId"`
	// The description of flowlog.
	Description *string              `pulumi:"description"`
	Flowlogs    []GetFlowlogsFlowlog `pulumi:"flowlogs"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of CEN flow log IDs.
	Ids []string `pulumi:"ids"`
	// The name of the log store which is in the  `projectName` SLS project.
	LogStoreName *string `pulumi:"logStoreName"`
	NameRegex    *string `pulumi:"nameRegex"`
	// A list of CEN flow log names.
	Names []string `pulumi:"names"`
	// The name of the SLS project.
	ProjectName *string `pulumi:"projectName"`
	// The status of flowlog.
	Status *string `pulumi:"status"`
}
